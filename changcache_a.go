package cdn

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cryptor"
)

var _ signer = (*chuangcacheA)(nil)

type chuangcacheA struct {
	pattern string
	token   string
}

func newChuangcacheA(token string) *chuangcacheA {
	return &chuangcacheA{
		pattern: "%s%s%d",
		token:   token,
	}
}

func (ca *chuangcacheA) sign(url *url.URL) (err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(ca.pattern, ca.token, url.Path, now)
	query := url.Query()
	query.Add("KEY1", cryptor.New(key).Md5().Hex())
	query.Add("KEY2", "timestamp")

	return
}
