package cdn

import (
	"net/url"
)

type executor interface {
	sign(url *url.URL) (err error)
}
