package hcloudHelpers

import (
	"context"
	"log"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func GetLoadBalancerIP(token string) string {
	return ""
}
func GetLoadBalancers(hcloudToken string) []*hcloud.LoadBalancer {
	if hcloudToken == "" {
		log.Fatalln("Please supply filled HcloudToken")
	}
	client := hcloud.NewClient(hcloud.WithToken(hcloudToken))
	servers, _, _ := client.LoadBalancer.List(context.Background(), hcloud.LoadBalancerListOpts{ListOpts: hcloud.ListOpts{}})
	return servers

}

func GetFirstLoadBalancerIP(hcloudToken string) string {
	if servers := GetLoadBalancers(hcloudToken); servers != nil {

		return getLoadBalancerIP(servers[0])
	}
	return ""
}

func getLoadBalancerIP(server *hcloud.LoadBalancer) string {
	return server.PublicNet.IPv4.IP.String()
}
