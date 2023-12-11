package builder

import (
	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cdn/internal/internal/param"
	"github.com/goexl/cdn/internal/internal/signer/d"
)

type Chuangcache struct {
	domain *Domain
	params *param.Domain
	signer internal.Signer
}

func NewChuangcache(builder *Domain, params *param.Domain) *Chuangcache {
	return &Chuangcache{
		domain: builder,
		params: params,
	}
}

func (c *Chuangcache) C(token string) (cache *Chuangcache) {
	c.signer = d.NewChuangcache(token)
	cache = c

	return
}

func (c *Chuangcache) Build() (domain *Domain) {
	c.params.Signer = c.signer
	domain = c.domain

	return
}
