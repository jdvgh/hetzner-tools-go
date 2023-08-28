package cmdServer

import (
	"context"
	"fmt"
	"log"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List servers",
	Run: func(c *cobra.Command, args []string) {

		ListServers(HCloudApiToken)

	},
}

func init() {
}

func ListServers(hcloudToken string) {
	client := hcloud.NewClient(hcloud.WithToken(hcloudToken))
	servers, _, err := client.Server.List(context.Background(), hcloud.ServerListOpts{ListOpts: hcloud.ListOpts{}})
	if err != nil {
		log.Fatalf("error retrieving severs:%v", err)
	}
	if servers == nil {
		fmt.Printf("error retrieving severs:%v\n", servers)

	} else {
		if len(servers) == 0 {
			fmt.Println("None")
		} else {
			listServers(servers)
		}
	}
}
func listServers(servers []*hcloud.Server) {
	for _, server := range servers {
		fmt.Printf("ID: %v Name: %v, IPv4: %v\n", server.ID, server.Name, server.PublicNet.IPv4.IP)
	}
}
