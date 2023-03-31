package cdn

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cryptor"
)

var _ executor = (*chuangcache)(nil)

type chuangcache struct {
	pattern string
	token   string
}

func newChuangcache(token string) *chuangcache {
	return &chuangcache{
		pattern: "%s%s%d",
		token:   token,
	}
}

func (c *chuangcache) sign(url *url.URL) (err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(c.pattern, c.token, url.Path, now)
	query := url.Query()
	query.Add("KEY1", cryptor.New(key).Md5().Hex())
	query.Add("KEY2", "timestamp")

	return
}
