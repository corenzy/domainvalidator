package providers

import (
	"strings"

	"github.com/corenzy/domainvalidator/models"
)

// knownProvider maps a keyword match to provider details.
type knownProvider struct {
	Keywords []string
	Info     models.ProviderInfo
}

// registry holds all known DNS providers and their nameserver patterns.
var registry = []knownProvider{
	// --- Major Cloud Providers ---
	{
		Keywords: []string{"cloudflare"},
		Info:     models.ProviderInfo{Name: "Cloudflare", Website: "https://www.cloudflare.com"},
	},
	{
		Keywords: []string{"awsdns", "amazonaws"},
		Info:     models.ProviderInfo{Name: "Amazon Route 53", Website: "https://aws.amazon.com/route53"},
	},
	{
		Keywords: []string{"googledomains", "google.com"},
		Info:     models.ProviderInfo{Name: "Google Cloud DNS", Website: "https://cloud.google.com/dns"},
	},
	{
		Keywords: []string{"azure-dns", "microsoft.com"},
		Info:     models.ProviderInfo{Name: "Azure DNS", Website: "https://azure.microsoft.com/products/dns"},
	},
	{
		Keywords: []string{"digitalocean"},
		Info:     models.ProviderInfo{Name: "DigitalOcean", Website: "https://www.digitalocean.com/products/dns"},
	},
	{
		Keywords: []string{"corenzydns"},
		Info:     models.ProviderInfo{Name: "Corenzy DNS", Website: "https://corenzy.com/dns"},
	},
	{
		Keywords: []string{"linode", "akamai"},
		Info:     models.ProviderInfo{Name: "Akamai / Linode", Website: "https://www.linode.com"},
	},
	{
		Keywords: []string{"vultr"},
		Info:     models.ProviderInfo{Name: "Vultr", Website: "https://www.vultr.com"},
	},
	{
		Keywords: []string{"hetzner"},
		Info:     models.ProviderInfo{Name: "Hetzner", Website: "https://www.hetzner.com"},
	},
	{
		Keywords: []string{"oracle", "oraclecloud"},
		Info:     models.ProviderInfo{Name: "Oracle Cloud DNS", Website: "https://www.oracle.com/cloud"},
	},
	{
		Keywords: []string{"ibm", "softlayer"},
		Info:     models.ProviderInfo{Name: "IBM Cloud DNS", Website: "https://www.ibm.com/cloud"},
	},
	{
		Keywords: []string{"aliyun", "alibaba", "alidns", "hichina"},
		Info:     models.ProviderInfo{Name: "Alibaba Cloud DNS", Website: "https://www.alibabacloud.com"},
	},
	{
		Keywords: []string{"tencentcloud", "dnspod"},
		Info:     models.ProviderInfo{Name: "Tencent Cloud DNS", Website: "https://cloud.tencent.com"},
	},

	// --- Hosting & Platform Providers ---
	{
		Keywords: []string{"vercel-dns"},
		Info:     models.ProviderInfo{Name: "Vercel", Website: "https://vercel.com"},
	},
	{
		Keywords: []string{"netlify"},
		Info:     models.ProviderInfo{Name: "Netlify", Website: "https://www.netlify.com"},
	},
	{
		Keywords: []string{"wpengine", "wp-engine"},
		Info:     models.ProviderInfo{Name: "WP Engine", Website: "https://wpengine.com"},
	},
	{
		Keywords: []string{"squarespace"},
		Info:     models.ProviderInfo{Name: "Squarespace", Website: "https://www.squarespace.com"},
	},
	{
		Keywords: []string{"wix"},
		Info:     models.ProviderInfo{Name: "Wix", Website: "https://www.wix.com"},
	},
	{
		Keywords: []string{"shopify"},
		Info:     models.ProviderInfo{Name: "Shopify", Website: "https://www.shopify.com"},
	},

	// --- Domain Registrars ---
	{
		Keywords: []string{"namecheap", "registrar-servers"},
		Info:     models.ProviderInfo{Name: "Namecheap", Website: "https://www.namecheap.com"},
	},
	{
		Keywords: []string{"godaddy", "domaincontrol"},
		Info:     models.ProviderInfo{Name: "GoDaddy", Website: "https://www.godaddy.com"},
	},
	{
		Keywords: []string{"name.com"},
		Info:     models.ProviderInfo{Name: "Name.com", Website: "https://www.name.com"},
	},
	{
		Keywords: []string{"hover"},
		Info:     models.ProviderInfo{Name: "Hover", Website: "https://www.hover.com"},
	},
	{
		Keywords: []string{"gandi"},
		Info:     models.ProviderInfo{Name: "Gandi", Website: "https://www.gandi.net"},
	},
	{
		Keywords: []string{"dynadot"},
		Info:     models.ProviderInfo{Name: "Dynadot", Website: "https://www.dynadot.com"},
	},
	{
		Keywords: []string{"enom"},
		Info:     models.ProviderInfo{Name: "eNom", Website: "https://www.enom.com"},
	},
	{
		Keywords: []string{"porkbun"},
		Info:     models.ProviderInfo{Name: "Porkbun", Website: "https://porkbun.com"},
	},
	{
		Keywords: []string{"namesilo"},
		Info:     models.ProviderInfo{Name: "NameSilo", Website: "https://www.namesilo.com"},
	},
	{
		Keywords: []string{"epik.com"},
		Info:     models.ProviderInfo{Name: "Epik", Website: "https://www.epik.com"},
	},
	{
		Keywords: []string{"ionos", "ui-dns"},
		Info:     models.ProviderInfo{Name: "IONOS (1&1)", Website: "https://www.ionos.com"},
	},
	{
		Keywords: []string{"register.com"},
		Info:     models.ProviderInfo{Name: "Register.com", Website: "https://www.register.com"},
	},

	// --- DNS-Specific Providers ---
	{
		Keywords: []string{"dnsimple"},
		Info:     models.ProviderInfo{Name: "DNSimple", Website: "https://dnsimple.com"},
	},
	{
		Keywords: []string{"dnsmadeeasy"},
		Info:     models.ProviderInfo{Name: "DNS Made Easy", Website: "https://dnsmadeeasy.com"},
	},
	{
		Keywords: []string{"ns1.com", "nsone"},
		Info:     models.ProviderInfo{Name: "NS1 (IBM)", Website: "https://ns1.com"},
	},
	{
		Keywords: []string{"ultradns", "neustar"},
		Info:     models.ProviderInfo{Name: "UltraDNS (Neustar)", Website: "https://www.home.neustar/dns-services"},
	},
	{
		Keywords: []string{"dyn.com", "dynect"},
		Info:     models.ProviderInfo{Name: "Dyn (Oracle)", Website: "https://dyn.com"},
	},
	{
		Keywords: []string{"constellix"},
		Info:     models.ProviderInfo{Name: "Constellix", Website: "https://constellix.com"},
	},
	{
		Keywords: []string{"easydns"},
		Info:     models.ProviderInfo{Name: "easyDNS", Website: "https://easydns.com"},
	},
	{
		Keywords: []string{"he.net", "hurricane"},
		Info:     models.ProviderInfo{Name: "Hurricane Electric", Website: "https://dns.he.net"},
	},
	{
		Keywords: []string{"afraid.org", "freedns"},
		Info:     models.ProviderInfo{Name: "FreeDNS", Website: "https://freedns.afraid.org"},
	},
	{
		Keywords: []string{"bunny", "bunnycdn"},
		Info:     models.ProviderInfo{Name: "Bunny DNS", Website: "https://bunny.net"},
	},

	// --- Hosting Providers ---
	{
		Keywords: []string{"ovh"},
		Info:     models.ProviderInfo{Name: "OVH", Website: "https://www.ovhcloud.com"},
	},
	{
		Keywords: []string{"hostinger"},
		Info:     models.ProviderInfo{Name: "Hostinger", Website: "https://www.hostinger.com"},
	},
	{
		Keywords: []string{"bluehost"},
		Info:     models.ProviderInfo{Name: "Bluehost", Website: "https://www.bluehost.com"},
	},
	{
		Keywords: []string{"hostgator"},
		Info:     models.ProviderInfo{Name: "HostGator", Website: "https://www.hostgator.com"},
	},
	{
		Keywords: []string{"dreamhost"},
		Info:     models.ProviderInfo{Name: "DreamHost", Website: "https://www.dreamhost.com"},
	},
	{
		Keywords: []string{"siteground"},
		Info:     models.ProviderInfo{Name: "SiteGround", Website: "https://www.siteground.com"},
	},
	{
		Keywords: []string{"inmotionhosting", "inmotion"},
		Info:     models.ProviderInfo{Name: "InMotion Hosting", Website: "https://www.inmotionhosting.com"},
	},
	{
		Keywords: []string{"a2hosting"},
		Info:     models.ProviderInfo{Name: "A2 Hosting", Website: "https://www.a2hosting.com"},
	},
	{
		Keywords: []string{"liquidweb"},
		Info:     models.ProviderInfo{Name: "Liquid Web", Website: "https://www.liquidweb.com"},
	},
	{
		Keywords: []string{"fasthosts"},
		Info:     models.ProviderInfo{Name: "Fasthosts", Website: "https://www.fasthosts.co.uk"},
	},
	{
		Keywords: []string{"pair.com"},
		Info:     models.ProviderInfo{Name: "pair Networks", Website: "https://www.pair.com"},
	},
	{
		Keywords: []string{"rackspace"},
		Info:     models.ProviderInfo{Name: "Rackspace", Website: "https://www.rackspace.com"},
	},

	// --- CDN / Security ---
	{
		Keywords: []string{"incapdns", "imperva"},
		Info:     models.ProviderInfo{Name: "Imperva (Incapsula)", Website: "https://www.imperva.com"},
	},
	{
		Keywords: []string{"sucuri"},
		Info:     models.ProviderInfo{Name: "Sucuri", Website: "https://sucuri.net"},
	},
	{
		Keywords: []string{"stackpath"},
		Info:     models.ProviderInfo{Name: "StackPath", Website: "https://www.stackpath.com"},
	},
	{
		Keywords: []string{"fastly"},
		Info:     models.ProviderInfo{Name: "Fastly", Website: "https://www.fastly.com"},
	},

	// --- Country / Regional ---
	{
		Keywords: []string{"yandex"},
		Info:     models.ProviderInfo{Name: "Yandex DNS", Website: "https://connect.yandex.com"},
	},
	{
		Keywords: []string{"selectel"},
		Info:     models.ProviderInfo{Name: "Selectel", Website: "https://selectel.ru"},
	},
	{
		Keywords: []string{"arvancloud"},
		Info:     models.ProviderInfo{Name: "ArvanCloud", Website: "https://www.arvancloud.ir"},
	},
	{
		Keywords: []string{"transip"},
		Info:     models.ProviderInfo{Name: "TransIP", Website: "https://www.transip.nl"},
	},
	{
		Keywords: []string{"strato"},
		Info:     models.ProviderInfo{Name: "Strato", Website: "https://www.strato.de"},
	},
	{
		Keywords: []string{"netcup"},
		Info:     models.ProviderInfo{Name: "Netcup", Website: "https://www.netcup.de"},
	},
	{
		Keywords: []string{"scaleway"},
		Info:     models.ProviderInfo{Name: "Scaleway", Website: "https://www.scaleway.com"},
	},
	{
		Keywords: []string{"zeit", "isimtescil"},
		Info:     models.ProviderInfo{Name: "İsimTescil", Website: "https://www.isimtescil.net"},
	},
	{
		Keywords: []string{"natro"},
		Info:     models.ProviderInfo{Name: "Natro", Website: "https://www.natro.com"},
	},
}

// Detect identifies the DNS provider from the given nameservers.
// It returns a best-match ProviderInfo.
func Detect(nameservers []string) models.ProviderInfo {
	for _, ns := range nameservers {
		nsLower := strings.ToLower(ns)
		for _, p := range registry {
			for _, kw := range p.Keywords {
				if strings.Contains(nsLower, kw) {
					return p.Info
				}
			}
		}
	}

	return models.ProviderInfo{
		Name:    "Unknown",
		Website: "",
	}
}
