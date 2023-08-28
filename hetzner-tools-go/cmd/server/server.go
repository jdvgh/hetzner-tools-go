package cmdServer

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const HCLOUD_API_TOKEN_KEY = "HCLOUD_API_TOKEN"
const HCLOUD_API_TOKEN_FLAG = "hcloud-api-token"

// serverCmd represents the server command
var serverCmd = &cobra.Command{

	Use:   "server",
	Short: "Manage hcloud servers",
	Run: func(c *cobra.Command, args []string) {
		c.HelpFunc()(c, args)
		os.Exit(1)
	},
}
var HCloudApiToken string

func NewServerCommand() *cobra.Command {
	return serverCmd
}
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
	hcloudApiToken, _ := os.LookupEnv(HCLOUD_API_TOKEN_KEY)
	serverCmd.AddCommand(listCmd)
	serverCmd.PersistentFlags().StringVar(&HCloudApiToken, HCLOUD_API_TOKEN_FLAG, hcloudApiToken, "API Token to authenticate against hcloud api")
}
