package cmd

import (
	"os"

	cmdDns "github.com/jdvgh/hetzner-tools-go/hetzner-tools-go/cmd/dns"
	cmdLoadBalancer "github.com/jdvgh/hetzner-tools-go/hetzner-tools-go/cmd/loadBalancer"
	cmdServer "github.com/jdvgh/hetzner-tools-go/hetzner-tools-go/cmd/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hetzner-tools-go",
	Short: "hetzner-tools-go handles API calls to multiple different Hetzner specific services",
	Run: func(c *cobra.Command, args []string) {
		c.HelpFunc()(c, args)
		os.Exit(1)
	},
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmdServer.NewServerCommand())
	rootCmd.AddCommand(cmdLoadBalancer.NewLoadBalancerCommand())
	rootCmd.AddCommand(cmdDns.NewDnsCommand())

}
