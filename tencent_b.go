package cdn

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cryptor"
)

var _ executor = (*tencentB)(nil)

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

func (b *tencentB) sign(url *url.URL) (err error) {
	nowHex := time.Now().Format("20060102150405")
	key := fmt.Sprintf(b.pattern, b.key, nowHex, url.EscapedPath())
	sign := cryptor.New(key).Md5().Hex()
	url.RawPath = fmt.Sprintf("%s/%s%s", nowHex, sign, url.EscapedPath())
	url.Path = fmt.Sprintf("%s/%s%s", nowHex, sign, url.Path)

	return
}
