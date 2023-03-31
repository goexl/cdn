package cdn

import (
	"net/url"
	"path"

	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

// Encoder 编码器
type Encoder struct {
	params  *params
	signers map[string]signer
	_       gox.CannotCopy
}

func newEncoder(params *params) *Encoder {
	return &Encoder{
		params: params,
	}
}

func (e *Encoder) Encode(from string) (encoded *url.URL, err error) {
	if parsed, pe := url.Parse(from); nil != pe {
		err = pe
	} else if _executor, ee := e.signer(parsed.Host); nil != ee {
		err = ee
	} else {
		err = _executor.sign(parsed)
	}

	return
}

func (e *Encoder) signer(host string) (signer signer, err error) {
	if cached, ok := e.signers[host]; ok {
		signer = cached
	} else {
		signer, err = e.match(host)
	}

	return
}

func (e *Encoder) match(host string) (signer signer, err error) {
	for domain, value := range e.params.signers {
		if matched, me := path.Match(domain, host); nil == me && matched {
			signer = value
		}
		if nil != signer {
			break
		}
	}
	if nil == signer {
		signer = e.params.signers[defaults]
	}
	if nil != signer {
		e.signers[host] = signer
	} else {
		err = exc.NewField("没有匹配到域名", field.New("domains", e.params.signers))
	}

	return
}
