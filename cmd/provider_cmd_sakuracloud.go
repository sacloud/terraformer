package cmd

import (
	sakuracloud_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/sakuracloud"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

func newCmdSakuraCloudImporter(options ImportOptions) *cobra.Command {
	var apiKey, appKey string
	cmd := &cobra.Command{
		Use:   "sakuracloud",
		Short: "Import current State to terraform configuration from sakuracloud",
		Long:  "Import current State to terraform configuration from sakuracloud",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newSakuraCloudProvider()
			err := Import(provider, options, []string{apiKey, appKey})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newSakuraCloudProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "server,disk", "sakuracloud_server=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&apiKey, "token", "", "", "YOUR_SAKURACLOUD_TOKEN or env param SAKURACLOUD_ACCESS_TOKEN")
	cmd.PersistentFlags().StringVarP(&appKey, "secret", "", "", "YOUR_SAKURACLOUD_SECRET or env param SAKURACLOUD_ACCESS_TOKEN_SECRET")
	return cmd
}

func newSakuraCloudProvider() terraform_utils.ProviderGenerator {
	return &sakuracloud_terraforming.SakuraCloudProvider{}
}
