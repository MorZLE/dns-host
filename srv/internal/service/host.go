package service

import (
	"context"
	"dns-host/srv/model/cerror"
	"fmt"
	"os"
	"os/exec"
)

func getHostname(ctx context.Context) (string, error) {
	return os.Hostname()
}

func setHostname(ctx context.Context, newHost string) error {
	const op = "service.Host.SetHostname"
	if newHost == "" {
		return cerror.ErrBadHostname
	}

	cmd := exec.Command(fmt.Sprintf("hostnamectl set-hostname %s", newHost))
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("systemctl restart systemd-hostnamed")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
