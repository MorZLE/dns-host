package service

import (
	"context"
	pb "dns-host/gen/server"
)

func GetHostname() (string, error) {
	conn, closeConn := getConn()
	defer closeConn()

	resp, err := conn.GetHostname(context.Background(), &pb.GetHostnameRequest{})
	if err != nil {
		return "", err
	}
	return resp.Hostname, nil

}

func SetHostname(hostname string) error {
	conn, closeConn := getConn()
	defer closeConn()

	_, err := conn.SetHostname(context.Background(), &pb.SetHostnameRequest{Hostname: hostname})
	if err != nil {
		return err
	}
	return nil
}

func RestartHost() (bool, error) {
	conn, closeConn := getConn()
	defer closeConn()

	res, err := conn.RestartHost(context.Background(), &pb.RestartHostRequest{})
	if err != nil {
		return false, err
	}
	return res.Success, nil
}
