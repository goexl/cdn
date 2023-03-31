package cdn

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/goexl/cryptor"
)

var _ signer = (*tencentC)(nil)

type tencentC struct {
	pattern string
	key     string
}

func newTencentC(key string) *tencentC {
	return &tencentC{
		pattern: "%s%s%s",
		key:     key,
	}
}

func (c *tencentC) sign(url *url.URL) (err error) {
	nowHex := strconv.FormatInt(time.Now().Unix(), 16)
	key := fmt.Sprintf(c.pattern, c.key, url.EscapedPath(), nowHex)
	sign := cryptor.New(key).Md5().Hex()
	url.RawPath = fmt.Sprintf("/%s/%s%s", sign, nowHex, url.EscapedPath())
	url.Path = fmt.Sprintf("/%s/%s%s", sign, nowHex, url.Path)

	return
}
