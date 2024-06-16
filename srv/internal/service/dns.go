package service

import (
	"context"
	"dns-host/pkg"
	"dns-host/pkg/cerror"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func NewDNSWorker(log *slog.Logger, pathResolve string) (*DNSWorker, error) {
	dns := DNSWorker{cacheDNS: map[pkg.Ip]pkg.Domain{}, pathResolve: pathResolve, log: log}
	_, err := dns.getAllDNS(context.Background())
	return &dns, err
}

type DNSWorker struct {
	log         *slog.Logger
	mut         sync.Mutex
	pathResolve string
	otherData   []string
	cacheDNS    map[pkg.Ip]pkg.Domain
}

func (d *DNSWorker) getAllDNS(ctx context.Context) (map[pkg.Ip]pkg.Domain, error) {
	if ctx.Err() != nil {
		return nil, cerror.ErrCancelled
	}

	if len(d.cacheDNS) > 0 {
		return d.cacheDNS, nil
	}

	d.otherData = []string{}

	file, err := os.ReadFile(d.pathResolve)
	if err != nil {
		return nil, err
	}

	d.mut.Lock()
	defer d.mut.Unlock()

	rows := strings.Split(string(file), "\n")
	for i := 0; i < len(rows); i++ {
		if len(rows[i]) == 0 {
			continue
		}
		if rows[i][0] == '#' { // пропускаем коментарии
			continue
		}

		dnsRow := strings.Fields(strings.TrimSpace(rows[i]))
		if len(dnsRow) != 2 && !pkg.Ip(dnsRow[0]).Valid() { // если это не данные о dns сервере записываем в слайс
			d.otherData = append(d.otherData, rows[i])
			continue
		}
		ip := pkg.Ip(dnsRow[0])
		domain := pkg.Domain(dnsRow[1])
		if ip.Valid() && domain.Valid() {
			d.cacheDNS[ip] = domain
		}
	}

	return d.cacheDNS, nil
}

func (d *DNSWorker) deleteDNS(ctx context.Context, name, ip string) error {
	if ctx.Err() != nil {
		return cerror.ErrCancelled
	}

	if ip != "" {
		d.mut.Lock()
		delete(d.cacheDNS, pkg.Ip(ip))
		d.mut.Unlock()
		err := d.rewriteDNS(ctx)
		if err != nil {
			return err
		}
		return nil
	}
	if name != "" {
		for k, v := range d.cacheDNS {
			if v == pkg.Domain(name) {
				d.mut.Lock()
				delete(d.cacheDNS, k)
				d.mut.Unlock()
			}
		}
		err := d.rewriteDNS(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	return cerror.ErrBadDNS
}

func (d *DNSWorker) addDNS(ctx context.Context, name, ip string) error {
	if ctx.Err() != nil {
		return cerror.ErrCancelled
	}
	err := func() error {
		d.mut.Lock()
		defer d.mut.Unlock()
		if _, ok := d.cacheDNS[pkg.Ip(ip)]; ok {
			return cerror.ErrRewrite
		}
		d.cacheDNS[pkg.Ip(ip)] = pkg.Domain(name)
		return nil
	}()

	if err != nil {
		return err
	}
	err = d.rewriteDNS(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DNSWorker) rewriteDNS(ctx context.Context) error {
	d.mut.Lock()
	defer d.mut.Unlock()

	err := d.writeDNS(ctx)
	if err != nil {
		return err
	}

	return d.restartManagerDNS()
}
func (d *DNSWorker) restartManagerDNS() error {
	cmd := exec.Command("systemctl restart NetworkManager")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (d *DNSWorker) writeDNS(ctx context.Context) error {
	file, err := os.OpenFile(d.pathResolve, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString("# File managed by dns-host service\n")
	for ip, domain := range d.cacheDNS {
		b.WriteString(fmt.Sprintf("%s \t %s\n", ip, domain))
	}

	for i := 0; i < len(d.otherData); i++ {
		b.WriteString(fmt.Sprintf("%s\n", d.otherData[i]))
	}
	_, err = file.WriteString(b.String())
	if err != nil {
		return err
	}

	return nil
}
