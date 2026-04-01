package services

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/corenzy/domainvalidator/models"
	"github.com/corenzy/domainvalidator/providers"
)

// Lookup performs a full DNS lookup for the given domain.
// It uses Cloudflare (1.1.1.1) directly to bypass potential local resolver issues
// and runs all queries in parallel for maximum speed.
func Lookup(domain string) (*models.LookupResult, error) {
	domain = normalizeDomain(domain)
	if domain == "" {
		return nil, fmt.Errorf("domain cannot be empty")
	}

	// Use a 5-second timeout for the whole process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use a custom resolver to bypass local OS issues (common on macOS)
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Second * 2,
			}
			return d.DialContext(ctx, network, "1.1.1.1:53")
		},
	}

	var (
		wg          sync.WaitGroup
		mu          sync.Mutex
		nsList      []string
		aRecords    []models.DNSRecord
		aaaaRecords []models.DNSRecord
		mxRecords   []models.DNSRecord
		txtRecords  []string
		cnameRecord string
	)

	// Define lookups
	lookups := []func(){
		// NS
		func() {
			nss, err := resolver.LookupNS(ctx, domain)
			if err == nil {
				var local []string
				for _, ns := range nss {
					local = append(local, strings.TrimRight(ns.Host, "."))
				}
				mu.Lock()
				nsList = local
				mu.Unlock()
			}
		},
		// IP (A & AAAA)
		func() {
			ips, err := resolver.LookupIP(ctx, "ip", domain)
			if err == nil {
				var a, aaaa []models.DNSRecord
				for _, ip := range ips {
					if ip.To4() != nil {
						a = append(a, models.DNSRecord{Type: "A", Value: ip.String()})
					} else {
						aaaa = append(aaaa, models.DNSRecord{Type: "AAAA", Value: ip.String()})
					}
				}
				mu.Lock()
				aRecords, aaaaRecords = a, aaaa
				mu.Unlock()
			}
		},
		// MX
		func() {
			mxs, err := resolver.LookupMX(ctx, domain)
			if err == nil {
				var local []models.DNSRecord
				for _, mx := range mxs {
					local = append(local, models.DNSRecord{
						Type:     "MX",
						Value:    strings.TrimRight(mx.Host, "."),
						Priority: mx.Pref,
					})
				}
				mu.Lock()
				mxRecords = local
				mu.Unlock()
			}
		},
		// TXT
		func() {
			txts, err := resolver.LookupTXT(ctx, domain)
			if err == nil {
				mu.Lock()
				txtRecords = txts
				mu.Unlock()
			}
		},
		// CNAME
		func() {
			cname, err := resolver.LookupCNAME(ctx, domain)
			if err == nil {
				resolved := strings.TrimRight(cname, ".")
				if resolved != domain {
					mu.Lock()
					cnameRecord = resolved
					mu.Unlock()
				}
			}
		},
	}

	// Execution
	for _, f := range lookups {
		wg.Add(1)
		go func(fn func()) {
			defer wg.Done()
			fn()
		}(f)
	}

	// Wait with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Success
	case <-ctx.Done():
		// Timeout
	}

	provider := providers.Detect(nsList)

	// Sane defaults for JSON
	if nsList == nil { nsList = []string{} }
	if aRecords == nil { aRecords = []models.DNSRecord{} }
	if aaaaRecords == nil { aaaaRecords = []models.DNSRecord{} }
	if mxRecords == nil { mxRecords = []models.DNSRecord{} }
	if txtRecords == nil { txtRecords = []string{} }

	return &models.LookupResult{
		Domain:      domain,
		Provider:    provider,
		Nameservers: nsList,
		ARecords:    aRecords,
		AAAARecords: aaaaRecords,
		MXRecords:   mxRecords,
		TXTRecords:  txtRecords,
		CNAMERecord: cnameRecord,
	}, nil
}

func normalizeDomain(domain string) string {
	domain = strings.TrimSpace(domain)
	domain = strings.TrimPrefix(domain, "https://")
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.Split(domain, "/")[0]
	domain = strings.TrimRight(domain, ".")
	return strings.ToLower(domain)
}
