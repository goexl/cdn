package cdn

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/goexl/cryptor"
)

var _ executor = (*tencentD)(nil)

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

func (d *tencentD) sign(url *url.URL) (err error) {
	nowHex := strconv.FormatInt(time.Now().Unix(), 16)
	key := fmt.Sprintf(d.pattern, d.key, url.EscapedPath(), nowHex)
	sign := cryptor.New(key).Md5().Hex()
	query := url.Query()
	query.Add(d.signature, sign)
	query.Add(d.timestamp, nowHex)
	url.RawQuery = query.Encode()

	return
}
