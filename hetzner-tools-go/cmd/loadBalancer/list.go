package cmdLoadBalancer

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
	Short: "List loadBalancers",
	Run: func(c *cobra.Command, args []string) {

		ListLoadBalancers(HCloudApiToken)

	},
}

func init() {
}

func ListLoadBalancers(hcloudToken string) {
	client := hcloud.NewClient(hcloud.WithToken(hcloudToken))
	loadBalancers, _, err := client.LoadBalancer.List(context.Background(), hcloud.LoadBalancerListOpts{ListOpts: hcloud.ListOpts{}})
	if err != nil {
		log.Fatalf("error retrieving severs:%v", err)
	}
	if loadBalancers == nil {
		fmt.Printf("error retrieving severs:%v\n", loadBalancers)

	} else {
		if len(loadBalancers) == 0 {
			fmt.Println("None")
		} else {
			listLoadBalancers(loadBalancers)
		}
	}
}
func listLoadBalancers(loadBalancers []*hcloud.LoadBalancer) {
	for _, loadBalancer := range loadBalancers {
		fmt.Printf("ID: %v Name: %v, IPv4: %v\n", loadBalancer.ID, loadBalancer.Name, loadBalancer.PublicNet.IPv4.IP)
	}
}
