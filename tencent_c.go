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

func (tc *tencentC) sign(url *url.URL) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 16)
	key := fmt.Sprintf(tc.pattern, tc.key, url.EscapedPath(), now)
	sign := cryptor.New(key).Md5().Hex()
	url.RawPath = fmt.Sprintf("/%s/%s%s", sign, now, url.EscapedPath())
	url.Path = fmt.Sprintf("/%s/%s%s", sign, now, url.Path)

	return
}
