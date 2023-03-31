package cdn

var _ = New

// New 创建构建器
func New() *builder {
	return newBuilder()
}
