package cdn

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
	"github.com/rs/xid"
)

var _ signer = (*tencentA)(nil)

type tencentA struct {
	pattern   string
	key       string
	signature string
}

func newTencentA(key string) *tencentA {
	return &tencentA{
		pattern:   "%s%d0%s%s",
		key:       key,
		signature: "sign",
	}
}

func (ta *tencentA) sign(url *url.URL) (err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(ta.pattern, url.EscapedPath(), now, xid.New().String(), ta.key)
	sb := gox.StringBuilder(url.RawQuery)
	if "" == url.RawQuery {
		sb.Append(question)
	} else {
		sb.Append(and)
	}
	sb.Append(ta.signature).Append(equal).Append(cryptor.New(key).Md5().Hex())
	url.RawQuery = sb.String()

	return
}
