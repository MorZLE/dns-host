/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dns-host/client/service"
	"dns-host/pkg/color"
	"fmt"

	"github.com/spf13/cobra"
)

// allDnsCmd represents the allDns command
var allDnsCmd = &cobra.Command{
	Use:   "allDns",
	Short: "Показывает список действующих серверов",
	Long:  `Показывает список действующих серверов установленных в файле resolv.conf`,
	Run: func(cmd *cobra.Command, args []string) {
		dns, err := service.GetAllDNS()
		if err != nil {
			color.Print(err.Error(), color.CRed)
			return
		}
		for _, v := range dns {
			fmt.Println(v.NameServer, v.Ip)
		}
	},
}

func init() {
	rootCmd.AddCommand(allDnsCmd)

}
