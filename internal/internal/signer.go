package internal

import (
	"net/url"
)

type Signer interface {
	Sign(url *url.URL) (err error)
}
