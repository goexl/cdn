package cdn

import (
	"github.com/goexl/cdn/internal/builder"
	"github.com/goexl/cdn/internal/core"
)

// Client 客户端
type Client = core.Client

func New() *builder.Cdn {
	return builder.New()
}
