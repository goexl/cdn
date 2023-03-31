package cdn

import (
	"github.com/goexl/gox"
)

// CDN 客户端
type CDN struct {
	params  *params
	signers map[string]signer
	_       gox.CannotCopy
}

func newCDN(params *params) *CDN {
	return &CDN{
		params: params,
	}
}
