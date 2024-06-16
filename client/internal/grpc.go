package internal

import (
	"context"
	pb "dns-host/gen/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func InitGrpcClient(addr string) ClientDNSHost {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	cli := ClientDNSHost{conn: pb.NewServiceDNSClient(conn)}

	return cli
}

type ClientDNSHost struct {
	conn pb.ServiceDNSClient
}

func (c *ClientDNSHost) GetHostname() (string, error) {
	resp, err := c.conn.GetHostname(context.Background(), &pb.GetHostnameRequest{})
	if err != nil {
		return "", err
	}
	return resp.Hostname, nil

}

func (c *ClientDNSHost) SetHostname(hostname string) error {
	_, err := c.conn.SetHostname(context.Background(), &pb.SetHostnameRequest{Hostname: hostname})
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientDNSHost) GetAllDNS() ([]*pb.Dns, error) {
	resp, err := c.conn.GetAllDNS(context.Background(), &pb.GetAllDNSRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Items, nil

}

func (c *ClientDNSHost) AddDNS(nameserver, ip string) (bool, error) {
	resp, err := c.conn.AddDNS(context.Background(), &pb.AddDNSRequest{NameServer: nameserver, Ip: ip})
	if err != nil {
		return nil, err
	}
	return resp.Success, nil

}

func (c *ClientDNSHost) DeleteDNS(nameserver, ip string) (bool, error) {
	resp, err := c.conn.DeleteDNS(context.Background(), &pb.DeleteDNSRequest{NameServer: nameserver, Ip: ip})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}
