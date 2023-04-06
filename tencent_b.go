package cdn

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cryptor"
)

var _ signer = (*tencentB)(nil)

type tencentB struct {
	pattern string
	key     string
}

func newTencentB(key string) *tencentB {
	return &tencentB{
		pattern: "%s%s%s",
		key:     key,
	}
}

func (tb *tencentB) sign(url *url.URL) (err error) {
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf(tb.pattern, tb.key, now, url.EscapedPath())
	sign := cryptor.New(key).Md5().Hex()
	url.RawPath = fmt.Sprintf("%s/%s%s", now, sign, url.EscapedPath())
	url.Path = fmt.Sprintf("%s/%s%s", now, sign, url.Path)

	return
}
