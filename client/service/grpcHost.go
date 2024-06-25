package service

import (
	"context"
	pb "dns-host/gen/server"
)

// GetHostname возвращает имя хоста
func GetHostname() (string, error) {
	conn, closeConn := getConn()
	defer closeConn()

	resp, err := conn.GetHostname(context.Background(), &pb.GetHostnameRequest{})
	if err != nil {
		return "", err
	}
	return resp.Hostname, nil

}

// SetHostname устанавливает имя хоста
func SetHostname(hostname string) error {
	conn, closeConn := getConn()
	defer closeConn()

	_, err := conn.SetHostname(context.Background(), &pb.SetHostnameRequest{Hostname: hostname})
	if err != nil {
		return err
	}
	return nil
}
