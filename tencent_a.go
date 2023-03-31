package cdn

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cryptor"
	"github.com/rs/xid"
)

var _ signer = (*tencentA)(nil)

type tencentA struct {
	pattern string
	key     string
}

func newTencentA(key string) *tencentA {
	return &tencentA{
		pattern: "%s%d0%s%s",
		key:     key,
	}
}

func (a *tencentA) sign(url *url.URL) (err error) {
	nowHex := time.Now().Unix()
	key := fmt.Sprintf(a.pattern, url.EscapedPath(), nowHex, xid.New().String(), a.key)
	query := url.Query()
	query.Add("sign", cryptor.New(key).Md5().Hex())
	url.RawQuery = query.Encode()

	return
}
