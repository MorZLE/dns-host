package service

import (
	"context"
	pb "dns-host/gen/server"
)

func GetAllDNS() ([]*pb.Dns, error) {
	conn, closeConn := getConn()
	defer closeConn()

	resp, err := conn.GetAllDNS(context.Background(), &pb.GetAllDNSRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Items, nil

}

func AddDNS(nameserver, ip string) (bool, error) {
	conn, closeConn := getConn()
	defer closeConn()

	resp, err := conn.AddDNS(context.Background(), &pb.AddDNSRequest{NameServer: nameserver, Ip: ip})
	if err != nil {
		return false, err
	}
	return resp.Success, nil

}

func DeleteDNS(nameserver, ip string) (bool, error) {
	conn, closeConn := getConn()
	defer closeConn()

	resp, err := conn.DeleteDNS(context.Background(), &pb.DeleteDNSRequest{NameServer: nameserver, Ip: ip})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func RestartDNS() (bool, error) {
	conn, closeConn := getConn()
	defer closeConn()

	res, err := conn.RestartDNS(context.Background(), &pb.RestartDNSRequest{})
	if err != nil {
		return false, err
	}

	return res.Success, nil
}
