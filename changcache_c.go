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

func (cc *chuangcacheC) sign(from *url.URL) (err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(cc.pattern, cc.token, from.Path, now)
	sb := gox.StringBuilder(from.RawQuery)
	if "" == from.RawQuery {
		sb.Append(question)
	} else {
		sb.Append(and)
	}
	sb.Append(cc.signature).Append(equal).Append(cryptor.New(key).Md5().Hex())
	sb.Append(and).Append(cc.timestamp).Append(equal).Append(now)
	from.RawQuery = sb.String()

	return
}
