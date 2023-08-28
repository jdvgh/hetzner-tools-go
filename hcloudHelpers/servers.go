package hcloudHelpers

import (
	"context"
	"log"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func GetServers(hcloudToken string) []*hcloud.Server {
	if hcloudToken == "" {
		log.Fatalln("Please supply filled HcloudToken")
	}
	client := hcloud.NewClient(hcloud.WithToken(hcloudToken))
	servers, _, _ := client.Server.List(context.Background(), hcloud.ServerListOpts{ListOpts: hcloud.ListOpts{}})
	return servers

}

func GetFirstServerIP(hcloudToken string) string {
	if servers := GetServers(hcloudToken); servers != nil {

		return getServerIP(servers[0])
	}
	return ""
}

func getServerIP(server *hcloud.Server) string {
	return server.PublicNet.IPv4.IP.String()
}
