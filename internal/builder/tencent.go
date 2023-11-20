package builder

import (
	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cdn/internal/internal/param"
	"github.com/goexl/cdn/internal/internal/signer/a"
	"github.com/goexl/cdn/internal/internal/signer/b"
	"github.com/goexl/cdn/internal/internal/signer/c"
	"github.com/goexl/cdn/internal/internal/signer/d"
)

type Tencent struct {
	domain *Domain
	params *param.Domain
	signer internal.Signer
}

func NewTencent(domain *Domain, params *param.Domain) *Tencent {
	return &Tencent{
		domain: domain,
		params: params,
	}
}

func (t *Tencent) A(key string) *Tencent {
	t.signer = a.NewTencent(key)

	return t
}

func (t *Tencent) B(key string) *Tencent {
	t.signer = b.NewTencent(key)

	return t
}

func (t *Tencent) C(key string) *Tencent {
	t.signer = c.NewTencent(key)

	return t
}

func (t *Tencent) D(key string, sign string, timestamp string) *Tencent {
	t.signer = d.NewTencent(key, sign, timestamp)

	return t
}

func (t *Tencent) Build() (domain *Domain) {
	t.params.Signer = t.signer
	domain = t.domain

	return
}
