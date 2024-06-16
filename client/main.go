/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"dns-host/client/cmd"
	"dns-host/client/config"
	"dns-host/client/service"
)

func main() {
	cfg := config.NewConfig()
	service.SetServerAddr(cfg.GRPC.Host)
	cmd.Execute()
}
