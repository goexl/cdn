package c

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cryptor"
)

var _ internal.Signer = (*Tencent)(nil)

type Tencent struct {
	pattern string
	key     string
}

func NewTencent(key string) *Tencent {
	return &Tencent{
		pattern: "%s%s%s",
		key:     key,
	}
}

func (t *Tencent) Sign(url *url.URL) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 16)
	key := fmt.Sprintf(t.pattern, t.key, url.EscapedPath(), now)
	sign := cryptor.New(key).Md5().Hex()
	url.RawPath = fmt.Sprintf("/%s/%s%s", sign, now, url.EscapedPath())
	url.Path = fmt.Sprintf("/%s/%s%s", sign, now, url.Path)

	return
}
