package cmdDns

import (
	"github.com/jdvgh/hetzner-tools-go/dns"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Hetzner DNS Record per zone",
	Run: func(c *cobra.Command, args []string) {

		listDNSRecords(HetznerDnsApiToken, HetznerDnsZoneId)
	},
}

func init() {
}

func listDNSRecords(token, zoneID string) {
	dns.SendGetRecords(token, zoneID)
}
