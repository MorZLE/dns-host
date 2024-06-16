/*
Copyright © 2024 MorZle
*/
package main

import (
	"dns-host/client/cmd"
	"dns-host/client/internal"
)

func main() {
	srv := internal.InitGrpcClient("127.0.0.1:44044")
	cmd.InitGRPCinCLI(srv)
	cmd.Execute()
}
