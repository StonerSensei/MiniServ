package utils

import (
	"net"
	"strings"
	"urlShortner/models"
)

func ValidDomain(domain string) string {
	domain = strings.TrimPrefix(domain, "https://")
	domain = strings.TrimPrefix(domain, "http://")
	parts := strings.Split(domain, "/")
	return parts[0]
}

func DNSLookup(domain string) (*models.DNSInfo, error) {
	domain = ValidDomain(domain)
	dnsInfo := models.DNSInfo{Domain: domain}

	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		dnsInfo.IPs = append(dnsInfo.IPs, ip.String())
	}

	nsRecords, _ := net.LookupNS(domain)
	for _, ns := range nsRecords {
		dnsInfo.NS = append(dnsInfo.NS, ns.Host)
	}

	mxRecords, _ := net.LookupMX(domain)
	for _, mx := range mxRecords {
		dnsInfo.MX = append(dnsInfo.MX, mx.Host)
	}

	cname, _ := net.LookupCNAME(domain)
	dnsInfo.CNAME = cname

	txtRecords, _ := net.LookupTXT(domain)
	dnsInfo.TXT = txtRecords

	return &dnsInfo, nil
}
