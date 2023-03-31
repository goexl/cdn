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
	return d.signer.sign(url)
}
