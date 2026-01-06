package builder

import (
	"strings"

	"github.com/goexl/cdn/internal/internal/constant"
	"github.com/goexl/cdn/internal/internal/param"
)

type Domain struct {
	cdn      *Cdn
	params   *param.Domain
	patterns []string
}

func NewDomain(cdn *Cdn) *Domain {
	return &Domain{
		cdn:      cdn,
		params:   param.NewDomain(),
		patterns: make([]string, 0, 1),
	}
}

func (d *Domain) Host(host string) (tencent *Domain) {
	d.params.Host = host
	tencent = d

	return
}

func (d *Domain) Http() (tencent *Domain) {
	d.params.Scheme = constant.Http
	tencent = d

	return
}

func (d *Domain) Https() (tencent *Domain) {
	d.params.Scheme = constant.Https
	tencent = d

	return
}

func (d *Domain) Scheme(scheme string) (tencent *Domain) {
	switch strings.ToLower(scheme) {
	case constant.Http:
		d.params.Scheme = constant.Http
	case constant.Https:
		d.params.Scheme = constant.Https
	}
	tencent = d

	return
}

func (d *Domain) Default() (tencent *Domain) {
	d.patterns = append(d.patterns, constant.Defaults)
	tencent = d

	return
}

func (d *Domain) Pattern(pattern ...string) (tencent *Domain) {
	if 0 == len(pattern) {
		d.patterns = append(d.patterns, constant.Defaults)
	} else {
		d.patterns = append(d.patterns, pattern...)
	}
	tencent = d

	return
}

func (d *Domain) Tencent() *Tencent {
	return NewTencent(d, d.params)
}

func (d *Domain) Ks() *Ks {
	return NewKs(d, d.params)
}

func (d *Domain) Build() (cdn *Cdn) {
	for _, pattern := range d.patterns {
		d.cdn.params.Domains[pattern] = d.params
	}
	cdn = d.cdn

	return
}
