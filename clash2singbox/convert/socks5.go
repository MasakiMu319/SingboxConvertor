package convert

import (
	"SingboxConvertor/clash2singbox/model/clash"
	"SingboxConvertor/clash2singbox/model/singbox"
)

func httpOpts(p *clash.Proxies, s *singbox.SingBoxOut) error {
	tls(p, s)
	p.Username = s.Username
	return nil
}

func socks5(p *clash.Proxies, s *singbox.SingBoxOut) error {
	tls(p, s)
	p.Username = s.Username
	if !p.Udp {
		s.Network = "tcp"
	}
	return nil
}
