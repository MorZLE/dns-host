package service

import (
	"context"
	"dns-host/pkg/cerror"
	"fmt"
	"os"
	"os/exec"
)

func getHostname(ctx context.Context) (string, error) {
	if ctx.Err() != nil {
		return "", cerror.ErrCancelled
	}
	ctx.Done()

	return os.Hostname()
}

func setHostname(ctx context.Context, newHost string) error {
	if ctx.Err() != nil {
		return cerror.ErrCancelled
	}
	if newHost == "" {
		return cerror.ErrBadHostname
	}
	cmd := exec.Command(fmt.Sprintf("sudo hostnamectl set-hostname %s", newHost))
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
