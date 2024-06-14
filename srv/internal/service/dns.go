package service

import (
	"context"
	"dns-host/srv/model"
	"dns-host/srv/model/cerror"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"
)

func NewDns(log *slog.Logger, pathResolve string) (*Dns, error) {
	dns := Dns{cacheDNS: map[model.Ip]model.Domain{}, pathResolve: pathResolve, log: log}
	_, err := dns.getAllDNS(context.Background())
	return &dns, err
}

type Dns struct {
	log         *slog.Logger
	mut         sync.Mutex
	pathResolve string
	otherData   []string
	cacheDNS    map[model.Ip]model.Domain
}

func (d *Dns) getAllDNS(ctx context.Context) (map[model.Ip]model.Domain, error) {
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
		if len(dnsRow) != 2 && !model.Ip(dnsRow[0]).Valid() { // если это не данные о dns сервере записываем в слайс
			d.otherData = append(d.otherData, rows[i])
			continue
		}
		ip := model.Ip(dnsRow[0])
		domain := model.Domain(dnsRow[1])
		if ip.Valid() && domain.Valid() {
			d.cacheDNS[ip] = domain
		}
	}

	return d.cacheDNS, nil
}

func (d *Dns) setDNS(ctx context.Context, name, ip string) error {
	d.mut.Lock()
	d.cacheDNS[model.Ip(ip)] = model.Domain(name)
	d.mut.Unlock()

	return d.rewriteDNS(ctx)
}

func (d *Dns) deleteDNS(ctx context.Context, name, ip string) error {
	if ip != "" {
		d.mut.Lock()
		delete(d.cacheDNS, model.Ip(ip))
		d.mut.Unlock()
		err := d.rewriteDNS(ctx)
		if err != nil {
			return err
		}
		return nil
	}
	if name != "" {
		for k, v := range d.cacheDNS {
			if v == model.Domain(name) {
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

func (d *Dns) rewriteDNS(ctx context.Context) error {
	d.mut.Lock()
	defer d.mut.Unlock()

	err := d.writeDNS(ctx)
	if err != nil {
		return err
	}
	//	cmd := exec.Command("systemctl restart NetworkManager")
	//	err = cmd.Run()
	//	if err != nil {
	//		return err
	//	}

	return nil
}

func (d *Dns) writeDNS(ctx context.Context) error {
	file, err := os.OpenFile(d.pathResolve, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString("# File managed by dns-host\n")
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
