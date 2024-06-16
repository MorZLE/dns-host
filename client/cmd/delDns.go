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

// delDnsCmd represents the delDns command
var delDnsCmd = &cobra.Command{
	Use:   "delDns",
	Short: "Удаление DNS",
	Long:  `Удаляет DNS из resolv.conf, удаляет по любому из параметров`,
	Run: func(cmd *cobra.Command, args []string) {
		hostname := cmd.Flag("server").Value.String()
		ip := cmd.Flag("ip").Value.String()
		if ip == "" || hostname == "" {
			print("Некорректные данные", color.CRed)
			return
		}

		res, err := service.DeleteDNS(hostname, ip)
		if err != nil && !res {
			print(fmt.Sprintf("Сервер %s IP %s не удален, ошибка %s", hostname, ip, err), color.CRed)
			return
		}
		print(fmt.Sprintf("Сервер %s IP %s удален", hostname, ip), color.CGreen)
	},
}

func init() {
	rootCmd.AddCommand(delDnsCmd)
	delDnsCmd.Flags().String("server", "localhost", "server")
	delDnsCmd.Flags().String("ip", "127.0.0.1", "IP")

}
