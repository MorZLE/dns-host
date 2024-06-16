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

// dnsRestartCmd represents the dnsRestart command
var dnsRestartCmd = &cobra.Command{
	Use:   "dnsRestart",
	Short: "Перезапускает сервис NetworkManager",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ok, err := service.RestartDNS()
		if err != nil && !ok {
			print(fmt.Sprintf("Сервис DNS не перезапущен, ошибка %s", err), color.CRed)
			return
		}
		print("Сервис DNS перезапущен", color.CGreen)
	},
}

func init() {
	rootCmd.AddCommand(dnsRestartCmd)
}
