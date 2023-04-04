package cdn

import (
	"net/url"
)

type domain struct {
	host   string
	scheme string
	signer signer
}

func newDomain() *domain {
	return &domain{
		scheme: https,
	}
}

func (d *domain) sign(url *url.URL) (err error) {
	if se := d.signer.sign(url); nil != se {
		err = se
	} else {
		url.Host = d.host
		url.Scheme = d.scheme
	}

	return
}
