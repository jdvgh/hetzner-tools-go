package cmdDns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jdvgh/hetzner-tools-go/dns"
	"github.com/jdvgh/hetzner-tools-go/hcloudHelpers"
	"github.com/spf13/cobra"
)

// updateBulkCmd represents the updateBulk command
var updateBulkCmd = &cobra.Command{

	Use:   "updateBulk",
	Short: "Update Bulk all current existing DNS Entries of given Zone",
	Long:  "Retrieve Server IP and update all existing DNS Entries of Type A in this Zone with the new IP",

	Run: func(c *cobra.Command, args []string) {
		retrieveUpdateRecords(HetznerDnsZoneId, HetznerDnsApiToken, HCloudApiToken)
	},
}

func init() {
}
func retrieveUpdateRecords(zoneID, dnsToken, hcloudToken string) {
	newIP := hcloudHelpers.GetFirstLoadBalancerIP(hcloudToken)
	if newIP == "" {
		log.Println("Could not retrieve IP from LoadBalancer - trying server next")
		newIP = hcloudHelpers.GetFirstServerIP(hcloudToken)
		if newIP == "" {
			log.Fatalln("Could not retrieve IP from Server")
		}
	}
	dnsRecords := dns.GetDnsRecordsByType(dnsToken, zoneID, dns.A)
	if dnsRecords.Records == nil {
		log.Fatalf("No DNS Records available in zone %v to update \n", zoneID)
	}
	exchangeIpInRecords(newIP, &dnsRecords)
	sendBulkUpdateRecords(dnsToken, dnsRecords)
}

func exchangeIpInRecords(newIp string, dnsRecords *dns.DNSRecords) {

	for index := range dnsRecords.Records {
		dnsRecords.Records[index].Value = newIp
	}

}
func sendBulkUpdateRecords(dnsToken string, dnsRecords dns.DNSRecords) {

	// Bulk Update Records (PUT https://dns.hetzner.com/api/v1/records/bulk)
	//
	jsonBody, _ := json.Marshal(dnsRecords)
	fmt.Println(string(jsonBody))
	body := bytes.NewBuffer(jsonBody)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("PUT", "https://dns.hetzner.com/api/v1/records/bulk", body)

	// Headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", dnsToken)

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := io.ReadAll(resp.Body)
	var recordResponses dns.DNSRecordResponses
	json.Unmarshal(respBody, &recordResponses)

	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", recordResponses)
}
