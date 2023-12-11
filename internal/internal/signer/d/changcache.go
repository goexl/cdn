package d

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cdn/internal/internal/constant"
	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
)

var _ internal.Signer = (*Chuangcache)(nil)

type Chuangcache struct {
	pattern   string
	token     string
	signature string
	timestamp string
}

func NewChuangcache(token string) *Chuangcache {
	return &Chuangcache{
		pattern:   "%s%s%d",
		token:     token,
		signature: "sign",
		timestamp: "t",
	}
}

func (c *Chuangcache) Sign(url *url.URL, expired time.Duration) (err error) {
	now := time.Now().Add(expired).Unix()
	key := fmt.Sprintf(c.pattern, c.token, url.Path, now)
	sign := cryptor.New(key).Md5().Hex()
	sb := gox.StringBuilder(url.RawQuery)
	if "" == url.RawQuery {
		sb.Append(constant.Question)
	} else {
		sb.Append(constant.And)
	}
	sb.Append(c.signature).Append(constant.Equal).Append(sign)
	sb.Append(constant.And).Append(c.timestamp).Append(constant.Equal).Append(now)
	url.RawQuery = sb.String()

	return
}
