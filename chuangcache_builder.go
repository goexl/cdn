package cdn

type chuangcacheBuilder struct {
	builder  *builder
	domain   string
	executor signer
}

func newChuangcacheBuilder(builder *builder) *chuangcacheBuilder {
	return &chuangcacheBuilder{
		builder: builder,
	}
}

func (cb *chuangcacheBuilder) Default() *chuangcacheBuilder {
	return cb.Domain(defaults)
}

func (cb *chuangcacheBuilder) Any() *chuangcacheBuilder {
	return cb.Domain(defaults)
}

func (cb *chuangcacheBuilder) Domain(domain string) *chuangcacheBuilder {
	cb.domain = domain

	return cb
}

func (cb *chuangcacheBuilder) Token(token string) *chuangcacheBuilder {
	cb.executor = newChuangcache(token)

	return cb
}

func (cb *chuangcacheBuilder) Build() (b *builder) {
	b = cb.builder
	b.params.signers[cb.domain] = cb.executor

	return
}
