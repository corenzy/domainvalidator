package models

// LookupRequest represents the incoming POST body.
type LookupRequest struct {
	Domain string `json:"domain" validate:"required"`
}

// DNSRecord holds a single DNS record entry.
type DNSRecord struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Priority uint16 `json:"priority,omitempty"`
	TTL      uint32 `json:"ttl,omitempty"`
}

// ProviderInfo holds details about the detected DNS provider.
type ProviderInfo struct {
	Name    string `json:"name"`
	Website string `json:"website,omitempty"`
}

// LookupResult contains all DNS data for the queried domain.
type LookupResult struct {
	Domain      string       `json:"domain"`
	Provider    ProviderInfo `json:"provider"`
	Nameservers []string     `json:"nameservers"`
	ARecords    []DNSRecord  `json:"a_records"`
	AAAARecords []DNSRecord  `json:"aaaa_records"`
	MXRecords   []DNSRecord  `json:"mx_records"`
	TXTRecords  []string     `json:"txt_records"`
	CNAMERecord string       `json:"cname_record,omitempty"`
}

// APIResponse is the standard JSON envelope for all responses.
type APIResponse struct {
	Success      bool        `json:"success"`
	Message      string      `json:"message"`
	LookupTimeMs int64       `json:"lookup_time_ms"`
	Result       interface{} `json:"result,omitempty"`
}
