package cdn

import (
	"path"

	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

// CDN 客户端
type CDN struct {
	params  *params
	domains map[string]*domain
	_       gox.CannotCopy
}

func newCDN(params *params) *CDN {
	return &CDN{
		params:  params,
		domains: make(map[string]*domain),
	}
}

func (c *CDN) lookupDomain(host string) (domain *domain, err error) {
	if cached, ok := c.domains[host]; ok {
		domain = cached
	} else {
		domain, err = c.matchDomain(host)
	}

	return
}

func (c *CDN) matchDomain(host string) (domain *domain, err error) {
	for pattern, value := range c.params.domains {
		if matched, me := path.Match(pattern, host); nil == me && matched {
			domain = value
		}
		if nil != domain {
			break
		}
	}
	if nil == domain {
		domain = c.params.domains[defaults]
	}
	if nil != domain {
		c.domains[host] = domain
	} else {
		err = exc.NewField("没有匹配到域名", field.New("domains", c.params.domains))
	}

	return
}
