package a

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cdn/internal/internal/constant"
	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
	"github.com/rs/xid"
)

var _ internal.Signer = (*Tencent)(nil)

type Tencent struct {
	pattern   string
	key       string
	signature string
}

func NewTencent(key string) *Tencent {
	return &Tencent{
		pattern:   "%s%d0%s%s",
		key:       key,
		signature: "sign",
	}
}

func (t *Tencent) Sign(url *url.URL, _ time.Duration) (err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(t.pattern, url.EscapedPath(), now, xid.New().String(), t.key)
	sb := gox.StringBuilder(url.RawQuery)
	if "" == url.RawQuery {
		sb.Append(constant.Question)
	} else {
		sb.Append(constant.And)
	}
	sb.Append(t.signature).Append(constant.Equal).Append(cryptor.New(key).Md5().Hex())
	url.RawQuery = sb.String()

	return
}
