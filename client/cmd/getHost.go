/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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
			print(err.Error(), color.CRed)
			return
		}

		print(hostname, color.CGreen)
	},
}

func init() {
	rootCmd.AddCommand(getHostCmd)
}
