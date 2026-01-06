package a

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cdn/internal/internal/constant"
	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
)

var _ internal.Signer = (*Ks)(nil)

type Ks struct {
	pattern   string
	key       string
	signature string
}

func NewKs(key string) *Ks {
	return &Ks{
		pattern:   "%s%s%d",
		key:       key,
		signature: "sign",
	}
}

func (k *Ks) Sign(url *url.URL, _ time.Duration) (err error) {
	now := time.Now().Unix()
	key := cryptor.New(fmt.Sprintf(k.pattern, url.EscapedPath(), now)).Md5().Hex()
	sb := gox.StringBuilder(url.RawQuery)
	if "" == url.RawQuery {
		sb.Append(constant.Question)
	} else {
		sb.Append(constant.And)
	}
	sb.Append("t").Append(constant.Equal).Append(now).
		Append(constant.And).Append("k").Append(constant.Equal).Append(key)
	url.RawQuery = sb.String()

	return
}
