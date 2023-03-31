package cdn

import (
	"net/url"
	"path"

	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

// CDN 客户端
type CDN struct {
	params  *params
	signers map[string]signer
	_       gox.CannotCopy
}

func newCDN(params *params) *CDN {
	return &CDN{
		params: params,
	}
}

func (c *CDN) Sign(from string) (encoded *url.URL, err error) {
	if parsed, pe := url.Parse(from); nil != pe {
		err = pe
	} else if sign, ee := c.lookupSigner(parsed.Host); nil != ee {
		err = ee
	} else {
		err = sign.sign(parsed)
	}

	return
}

func (c *CDN) lookupSigner(host string) (signer signer, err error) {
	if cached, ok := c.signers[host]; ok {
		signer = cached
	} else {
		signer, err = c.matchSigner(host)
	}

	return
}

func (c *CDN) matchSigner(host string) (signer signer, err error) {
	for domain, value := range c.params.signers {
		if matched, me := path.Match(domain, host); nil == me && matched {
			signer = value
		}
		if nil != signer {
			break
		}
	}
	if nil == signer {
		signer = c.params.signers[defaults]
	}
	if nil != signer {
		c.signers[host] = signer
	} else {
		err = exc.NewField("没有匹配到域名", field.New("domains", c.params.signers))
	}

	return
}
