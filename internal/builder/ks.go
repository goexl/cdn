package builder

import (
	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cdn/internal/internal/param"
	"github.com/goexl/cdn/internal/internal/signer/a"
	"github.com/goexl/cdn/internal/internal/signer/b"
)

type Ks struct {
	domain *Domain
	params *param.Domain
	signer internal.Signer
}

func NewKs(domain *Domain, params *param.Domain) *Ks {
	return &Ks{
		domain: domain,
		params: params,
	}
}

func (k *Ks) A(key string) (ks *Ks) {
	k.signer = a.NewKs(key)
	ks = k

	return
}

func (k *Ks) B(key string) (ks *Ks) {
	k.signer = b.NewKs(key)
	ks = k

	return
}

func (k *Ks) Build() (domain *Domain) {
	k.params.Signer = k.signer
	domain = k.domain

	return
}
