package builder

import (
	"github.com/goexl/cdn/internal/core"
	"github.com/goexl/cdn/internal/internal/param"
)

type Cdn struct {
	params *param.Cdn
}

func New() *Cdn {
	return &Cdn{
		params: param.NewCdn(),
	}
}

func (c *Cdn) Domain() *Domain {
	return NewDomain(c)
}

func (c *Cdn) Build() *core.Client {
	return core.NewClient(c.params)
}
