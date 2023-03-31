package cdn

import (
	"net/url"
)

type signer interface {
	sign(url *url.URL) (err error)
}
