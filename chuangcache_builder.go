package cdn

type chuangcacheBuilder struct {
	builder *builder
	signers map[string]signer
}

func newChuangcacheBuilder(builder *builder) *chuangcacheBuilder {
	return &chuangcacheBuilder{
		builder: builder,
	}
}

func (cb *chuangcacheBuilder) Signer() *chuangcacheSignerBuilder {
	return newChuangcacheSignerBuilder(cb)
}

func (cb *chuangcacheBuilder) Build() (b *builder) {
	b = cb.builder
	for domain, _signer := range cb.signers {
		b.params.signers[domain] = _signer
	}

	return
}
