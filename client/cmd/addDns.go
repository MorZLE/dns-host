/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dns-host/client/service"
	"dns-host/pkg"
	"dns-host/pkg/color"
	"fmt"
	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var addDnsCmd = &cobra.Command{
	Use:   "dnsAdd",
	Short: "Добавление DNS",
	Long:  `Добавляет DNS в resolv.conf, установлена валидация для формата IP(является уникальным), hostname не должен привышать 255 байт'`,
	Run: func(cmd *cobra.Command, args []string) {
		hostname := cmd.Flag("server").Value.String()
		ip := cmd.Flag("ip").Value.String()
		if !pkg.Ip(ip).Valid() && !pkg.Domain(hostname).Valid() {
			print("Некорректные данные", color.CRed)
			return
		}

		res, err := service.AddDNS(hostname, ip)
		if err != nil && !res {
			print(fmt.Sprintf("Сервер %s IP %s не добавлен, ошибка %s", hostname, ip, err), color.CRed)
			return
		}
		print(fmt.Sprintf("Сервер %s IP %s добавлен", hostname, ip), color.CGreen)

	},
}

func init() {
	rootCmd.AddCommand(addDnsCmd)

	addDnsCmd.Flags().String("server", "server", "server")
	addDnsCmd.Flags().String("ip", "127.0.0.1", "IP")

}
