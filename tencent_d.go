package cdn

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
)

var _ signer = (*tencentD)(nil)

type tencentD struct {
	pattern   string
	key       string
	signature string
	timestamp string
}

func newTencentD(key string, signature string, timestamp string) *tencentD {
	return &tencentD{
		pattern:   "%s%s%s",
		key:       key,
		signature: signature,
		timestamp: timestamp,
	}
}

func (td *tencentD) sign(url *url.URL) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 16)
	key := fmt.Sprintf(td.pattern, td.key, url.EscapedPath(), now)
	sign := cryptor.New(key).Md5().Hex()
	sb := gox.StringBuilder(url.RawQuery)
	if "" == url.RawQuery {
		sb.Append(question)
	} else {
		sb.Append(and)
	}
	sb.Append(td.signature).Append(equal).Append(sign)
	sb.Append(and).Append(td.timestamp).Append(equal).Append(now)
	url.RawQuery = sb.String()

	return
}
