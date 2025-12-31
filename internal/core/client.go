package core

import (
	"net/url"
	"path"
	"time"

	"github.com/goexl/cdn/internal/internal/constant"
	"github.com/goexl/cdn/internal/internal/param"
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

// Client 客户端
type Client struct {
	params  *param.Cdn
	domains map[string]*param.Domain

	_ gox.Pointerized
}

func NewClient(params *param.Cdn) *Client {
	return &Client{
		params:  params,
		domains: make(map[string]*param.Domain),
	}
}

func (c *Client) Sign(from string, expired time.Duration) (signed *url.URL, err error) {
	if parsed, pe := url.Parse(from); nil != pe {
		err = pe
	} else if _domain, ee := c.lookupDomain(parsed.Host); nil != ee {
		err = ee
	} else if se := _domain.Sign(parsed, expired); nil != se {
		err = se
	} else {
		signed = parsed
	}

	return
}

func (c *Client) lookupDomain(host string) (domain *param.Domain, err error) {
	if cached, ok := c.domains[host]; ok {
		domain = cached
	} else {
		domain, err = c.matchDomain(host)
	}

	return
}

func (c *Client) matchDomain(host string) (domain *param.Domain, err error) {
	for pattern, value := range c.params.Domains {
		if matched, me := path.Match(pattern, host); nil == me && matched {
			domain = value
		}
		if nil != domain {
			break
		}
	}
	if nil == domain {
		domain = c.params.Domains[constant.Defaults]
	}
	if nil != domain {
		c.domains[host] = domain
	} else {
		err = exception.New().Message("没有匹配到域名").Field(field.New("domains", c.params.Domains)).Build()
	}

	return
}
