package param

import (
	"net/url"

	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cdn/internal/internal/constant"
)

type Domain struct {
	Host   string
	Scheme string
	Signer internal.Signer
}

func NewDomain() *Domain {
	return &Domain{
		Scheme: constant.Https,
	}
}

func (d *Domain) Sign(url *url.URL) (err error) {
	if se := d.Signer.Sign(url); nil != se {
		err = se
	} else {
		url.Host = d.Host
		url.Scheme = d.Scheme
	}

	return
}
