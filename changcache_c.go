package cdn

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
)

var _ signer = (*chuangcacheC)(nil)

type chuangcacheC struct {
	pattern   string
	token     string
	signature string
	timestamp string
}

func newChuangcacheC(token string) *chuangcacheC {
	return &chuangcacheC{
		pattern:   "%s%s%d",
		token:     token,
		signature: "sign",
		timestamp: "t",
	}
}

func (cc *chuangcacheC) sign(url *url.URL) (err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(cc.pattern, cc.token, url.Path, now)
	sign := cryptor.New(key).Md5().Hex()
	sb := gox.StringBuilder(url.RawQuery)
	if "" == url.RawQuery {
		sb.Append(question)
	} else {
		sb.Append(and)
	}
	sb.Append(cc.signature).Append(equal).Append(sign)
	sb.Append(and).Append(cc.timestamp).Append(equal).Append(now)
	url.RawQuery = sb.String()

	return
}
