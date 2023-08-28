package cmdDns

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{

	Use:   "dns",
	Short: "Hetzner DNS Management",

	Run: func(c *cobra.Command, args []string) {
		c.HelpFunc()(c, args)
		os.Exit(1)
	},
}
var (
	HCloudApiToken     string
	HetznerDnsApiToken string
	HetznerDnsZoneId   string
)

func NewDnsCommand() *cobra.Command {
	return dnsCmd
}

const HETZNER_DNS_API_TOKEN_KEY = "HETZNER_DNS_API_TOKEN"
const HETZNER_DNS_ZONE_ID_KEY = "HETZNER_DNS_ZONE_ID"
const HETZNER_DNS_API_TOKEN_FLAG = "hetzner-dns-api-token"
const HETZNER_DNS_ZONE_ID_FLAG = "hetzner-dns-zone-id"
const HCLOUD_API_TOKEN_FLAG = "hcloud-api-token"
const HCLOUD_API_TOKEN_KEY = "HCLOUD_API_TOKEN"

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		err := godotenv.Load("../.env")
		if err != nil {
			err = godotenv.Load()
			if err != nil {
				log.Printf("Could not load .env file at `../../.env` or `../.env` or `.env`")
			}
		}
	}

	hetznerDnsApiKey, _ := os.LookupEnv(HETZNER_DNS_API_TOKEN_KEY)
	hcloudApiToken, _ := os.LookupEnv(HCLOUD_API_TOKEN_KEY)
	hetznerDnsZone, _ := os.LookupEnv(HETZNER_DNS_ZONE_ID_KEY)

	dnsCmd.PersistentFlags().StringVar(&HetznerDnsApiToken, HETZNER_DNS_API_TOKEN_FLAG, hetznerDnsApiKey, "API Key to authenticate against hetzner DNS")
	dnsCmd.PersistentFlags().StringVar(&HCloudApiToken, HCLOUD_API_TOKEN_FLAG, hcloudApiToken, "API Key to authenticate against hcloud")
	dnsCmd.PersistentFlags().StringVar(&HetznerDnsZoneId, HETZNER_DNS_ZONE_ID_FLAG, hetznerDnsZone, "Hetzner DNS Zone")
	dnsCmd.AddCommand(listCmd)
	dnsCmd.AddCommand(updateBulkCmd)

}
