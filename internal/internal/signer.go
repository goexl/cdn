package internal

import (
	"net/url"
	"time"
)

type Signer interface {
	Sign(url *url.URL, expired time.Duration) (err error)
}
