package models

import (
	"github.com/lib/pq"
	"urlShortner/config"
)

type DNSInfo struct {
	Domain string   `json:"domain"`
	IPs    []string `json:"ip_addresses"`
	NS     []string `json:"name_servers"`
	MX     []string `json:"mail_servers"`
	CNAME  string   `json:"cname"`
	TXT    []string `json:"txt_records"`
}

func StoreDnsInfo(dnsInfo *DNSInfo) error {
	_, err := config.DB.Exec(
		`INSERT INTO dns_lookups (domain, ip_addresses, name_servers, mail_servers, cname, txt_records) 
		 VALUES ($1, $2, $3, $4, $5, $6) 
		 ON CONFLICT (domain) DO UPDATE 
		 SET ip_addresses = EXCLUDED.ip_addresses, 
		     name_servers = EXCLUDED.name_servers,  
		     mail_servers = EXCLUDED.mail_servers, 
		     cname = EXCLUDED.cname, 
		     txt_records = EXCLUDED.txt_records, 
		     created_at = CURRENT_TIMESTAMP`,
		dnsInfo.Domain,
		pq.Array(dnsInfo.IPs),
		pq.Array(dnsInfo.NS),
		pq.Array(dnsInfo.MX),
		dnsInfo.CNAME,
		pq.Array(dnsInfo.TXT),
	)
	return err
}
