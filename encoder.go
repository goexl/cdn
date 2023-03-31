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
	params    *params
	executors map[string]executor
	_         gox.CannotCopy
}

func newEncoder(params *params) *Encoder {
	return &Encoder{
		params: params,
	}
}

func (e *Encoder) Encode(from string) (encoded *url.URL, err error) {
	if parsed, pe := url.Parse(from); nil != pe {
		err = pe
	} else if _executor, ee := e.executor(parsed.Host); nil != ee {
		err = ee
	} else {
		err = _executor.sign(parsed)
	}

	return
}

func (e *Encoder) executor(host string) (executor executor, err error) {
	if cached, ok := e.executors[host]; ok {
		executor = cached
	} else {
		executor, err = e.match(host)
	}

	return
}

func (e *Encoder) match(host string) (executor executor, err error) {
	for domain, value := range e.params.executors {
		if matched, me := path.Match(domain, host); nil == me && matched {
			executor = value
		}
		if nil != executor {
			break
		}
	}
	if nil == executor {
		executor = e.params.executors[defaults]
	}
	if nil != executor {
		e.executors[host] = executor
	} else {
		err = exc.NewField("没有匹配到域名", field.New("domains", e.params.executors))
	}

	return
}
