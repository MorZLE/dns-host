package cmd

import (
	"dns-host/client/service"
	"dns-host/pkg/color"
	"github.com/spf13/cobra"
)

// getHostCmd represents the getHost command
var getHostCmd = &cobra.Command{
	Use:   "getHost",
	Short: "Получить хост",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		hostname, err := service.GetHostname()
		if err != nil {
			color.Print(err.Error(), color.CRed)
			return
		}

		color.Print(hostname, color.CGreen)
	},
}

func init() {
	rootCmd.AddCommand(getHostCmd)
}
