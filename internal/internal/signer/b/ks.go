package b

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goexl/cdn/internal/internal"
	"github.com/goexl/cryptor"
)

var _ internal.Signer = (*Ks)(nil)

type Ks struct {
	pattern   string
	key       string
	signature string
}

func NewKs(key string) *Ks {
	return &Ks{
		pattern:   "%s%s%d",
		key:       key,
		signature: "sign",
	}
}

func (k *Ks) Sign(url *url.URL, _ time.Duration) (err error) {
	now := time.Now().Unix()
	hash := fmt.Sprintf(k.pattern, k.key, url.EscapedPath(), now)
	key := cryptor.New(hash).Md5().Hex()
	url.RawPath = fmt.Sprintf("%s/%d%s", key, now, url.EscapedPath())
	url.Path = fmt.Sprintf("%s/%d%s", key, now, url.Path)

	return
}
