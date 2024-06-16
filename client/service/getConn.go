package service

import (
	pb "dns-host/gen/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type closeConnect func() error

var addr string

func SetServerAddr(hostname string) {
	addr = hostname
}

func getConn() (pb.ServiceDNSClient, closeConnect) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewServiceDNSClient(conn), conn.Close
}
