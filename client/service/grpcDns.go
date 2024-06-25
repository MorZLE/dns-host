package service

import (
	"context"
	pb "dns-host/gen/server"
)

// GetAllDNS возвращает все dns из resolv.conf
func GetAllDNS() ([]*pb.Dns, error) {
	conn, closeConn := getConn()
	defer closeConn()

	resp, err := conn.GetAllDNS(context.Background(), &pb.GetAllDNSRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Items, nil

}

// AddDNS добавляет dns в resolv.conf
func AddDNS(nameserver, ip string) error {
	conn, closeConn := getConn()
	defer closeConn()

	_, err := conn.AddDNS(context.Background(), &pb.AddDNSRequest{NameServer: nameserver, Ip: ip})
	if err != nil {
		return err
	}
	return nil

}

// DeleteDNS удаляет dns из resolv.conf
func DeleteDNS(nameserver, ip string) error {
	conn, closeConn := getConn()
	defer closeConn()

	_, err := conn.DeleteDNS(context.Background(), &pb.DeleteDNSRequest{NameServer: nameserver, Ip: ip})
	if err != nil {
		return err
	}
	return nil
}
