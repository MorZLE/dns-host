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

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "setHost",
	Short: "Изменение хоста",
	Long:  `Установлена прокерка для хоста в 255 байт`,
	Run: func(cmd *cobra.Command, args []string) {
		hostname := cmd.Flag("hostname").Value.String()
		if !pkg.Domain(hostname).Valid() {
			print("Некорректные данные", color.CRed)
		}

		err := service.SetHostname(hostname)
		if err != nil {
			print(fmt.Sprintf("Сервер %s не изменен, ошибка %s", hostname, err), color.CRed)
			return
		}
		print(fmt.Sprintf("Сервер %s изменен", hostname), color.CGreen)
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)
	hostCmd.Flags().String("hostname", "localhost", "hostname")

}
