package dns

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
)

type RecordTypeCreatable string

const (
	A     RecordTypeCreatable = "A"
	AAAA  RecordTypeCreatable = "AAAA"
	NS    RecordTypeCreatable = "NS"
	MX    RecordTypeCreatable = "MX"
	CNAME RecordTypeCreatable = "CNAME"
	RP    RecordTypeCreatable = "RP"
	TXT   RecordTypeCreatable = "TXT"
	SOA   RecordTypeCreatable = "SOA"
	HINFO RecordTypeCreatable = "HINFO"
	SRV   RecordTypeCreatable = "SRV"
	DANE  RecordTypeCreatable = "DANE"
	TLSA  RecordTypeCreatable = "TLSA"
	DS    RecordTypeCreatable = "DS"
	CAA   RecordTypeCreatable = "CAA"
)

type DNSRecord struct {
	Id     string              `json:"id"`
	Type   RecordTypeCreatable `json:"type"`
	Name   string              `json:"name"`
	Value  string              `json:"value"`
	Ttl    uint64              `json:"ttl"`
	ZoneId string              `json:"zone_id"`
}
type DNSRecords struct {
	Records []DNSRecord `json:"records"`
}

type DNSRecordResponse struct {
	DNSRecord
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

type DNSRecordResponses struct {
	Records []DNSRecordResponse `json:"records"`
}

const HCLOUD_DNS_API_TOKEN_KEY = "HCLOUD_DNS_API_TOKEN"
const HCLOUD_DNS_ZONE_ID_KEY = "HCLOUD_DNS_ZONE_ID"

func SendGetRecords(dnsToken, zoneID string) DNSRecordResponses {
	// Get Records (GET https://dns.hetzner.com/api/v1/records?zone_id={ZoneID})

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/records?zone_id=%v", zoneID), nil)

	// Headers
	req.Header.Add("Auth-API-Token", dnsToken)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := io.ReadAll(resp.Body)
	var recordResponses DNSRecordResponses
	json.Unmarshal(respBody, &recordResponses)
	// Display Results
	log.Println("response Status : ", resp.Status)
	log.Println("response Headers : ", resp.Header)
	log.Println("response Body : ", recordResponses)

	return recordResponses
}

func GetDnsRecordsByType(dnsToken string, zoneID string, recordTypes ...RecordTypeCreatable) DNSRecords {
	dnsRecordResponses := SendGetRecords(dnsToken, zoneID)
	dnsRecords := DNSRecords{}
	if dnsRecordResponses.Records != nil {
		for _, dnsRecordResponse := range dnsRecordResponses.Records {
			if recordTypes == nil || slices.Contains(recordTypes, dnsRecordResponse.DNSRecord.Type) {
				dnsRecords.Records = append(dnsRecords.Records, dnsRecordResponse.DNSRecord)

			}
		}
	}
	return dnsRecords
}
