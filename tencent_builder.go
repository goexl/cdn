package cdn

type tencentBuilder struct {
	builder  *builder
	domain   string
	executor signer
}

func newTencentBuilder(builder *builder) *tencentBuilder {
	return &tencentBuilder{
		builder: builder,
	}
}

func (tb *tencentBuilder) Default() *tencentBuilder {
	return tb.Domain(defaults)
}

func (tb *tencentBuilder) Any() *tencentBuilder {
	return tb.Domain(defaults)
}

func (tb *tencentBuilder) Domain(domain string) *tencentBuilder {
	tb.domain = domain

	return tb
}

func (tb *tencentBuilder) A(key string) *tencentBuilder {
	tb.executor = newTencentA(key)

	return tb
}

func (tb *tencentBuilder) B(key string) *tencentBuilder {
	tb.executor = newTencentB(key)

	return tb
}

func (tb *tencentBuilder) C(key string) *tencentBuilder {
	tb.executor = newTencentC(key)

	return tb
}

func (tb *tencentBuilder) D(key string, sign string, timestamp string) *tencentBuilder {
	tb.executor = newTencentD(key, sign, timestamp)

	return tb
}

func (tb *tencentBuilder) Build() (b *builder) {
	b = tb.builder
	b.params.signers[tb.domain] = tb.executor

	return
}
