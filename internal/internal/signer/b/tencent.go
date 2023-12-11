package b

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cryptor"
)

var _ internal.Signer = (*Tencent)(nil)

type Tencent struct {
	pattern string
	key     string
}

func NewTencent(key string) *Tencent {
	return &Tencent{
		pattern: "%s%s%s",
		key:     key,
	}
}

func (tb *Tencent) Sign(url *url.URL, _ time.Duration) (err error) {
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf(tb.pattern, tb.key, now, url.EscapedPath())
	sign := cryptor.New(key).Md5().Hex()
	url.RawPath = fmt.Sprintf("%s/%s%s", now, sign, url.EscapedPath())
	url.Path = fmt.Sprintf("%s/%s%s", now, sign, url.Path)

	return
}
