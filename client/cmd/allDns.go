package cmd

import (
	"dns-host/client/service"
	"dns-host/pkg/color"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"

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
		t := table.NewWriter()
		t.AppendHeader(table.Row{"#", "server", "ip"})

		for i, v := range dns {
			t.AppendRow(table.Row{i + 1, v.NameServer, v.Ip})
		}
		fmt.Println(t.Render())
	},
}

func init() {
	rootCmd.AddCommand(allDnsCmd)

}
