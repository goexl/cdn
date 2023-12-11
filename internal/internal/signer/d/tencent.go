package d

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cdn/internal/internal/constant"
	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
)

var _ internal.Signer = (*Tencent)(nil)

type Tencent struct {
	pattern   string
	key       string
	signature string
	timestamp string
}

func NewTencent(key string, signature string, timestamp string) *Tencent {
	return &Tencent{
		pattern:   "%s%s%s",
		key:       key,
		signature: signature,
		timestamp: timestamp,
	}
}

func (t *Tencent) Sign(url *url.URL, _ time.Duration) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 16)
	key := fmt.Sprintf(t.pattern, t.key, url.EscapedPath(), now)
	sign := cryptor.New(key).Md5().Hex()
	sb := gox.StringBuilder(url.RawQuery)
	if "" == url.RawQuery {
		sb.Append(constant.Question)
	} else {
		sb.Append(constant.And)
	}
	sb.Append(t.signature).Append(constant.Equal).Append(sign)
	sb.Append(constant.And).Append(t.timestamp).Append(constant.Equal).Append(now)
	url.RawQuery = sb.String()

	return
}
