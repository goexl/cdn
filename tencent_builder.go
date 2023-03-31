package cdn

type tencentBuilder struct {
	builder *builder
	signers map[string]signer
}

func newTencentBuilder(builder *builder) *tencentBuilder {
	return &tencentBuilder{
		builder: builder,
	}
}

func (tb *tencentBuilder) Signer() *tencentSignerBuilder {
	return newTencentSignerBuilder(tb)
}

func (tb *tencentBuilder) Build() (b *builder) {
	b = tb.builder
	for domain, _signer := range tb.signers {
		b.params.signers[domain] = _signer
	}

	return
}
