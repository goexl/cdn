package cdn

type chuangcacheSignerBuilder struct {
	builder *chuangcacheBuilder
	domain  string
	signer  signer
}

func newChuangcacheSignerBuilder(builder *chuangcacheBuilder) *chuangcacheSignerBuilder {
	return &chuangcacheSignerBuilder{
		builder: builder,
	}
}

func (csb *chuangcacheSignerBuilder) Default() *chuangcacheSignerBuilder {
	return csb.Domain(defaults)
}

func (csb *chuangcacheSignerBuilder) Any() *chuangcacheSignerBuilder {
	return csb.Domain(defaults)
}

func (csb *chuangcacheSignerBuilder) Domain(domain string) *chuangcacheSignerBuilder {
	csb.domain = domain

	return csb
}

func (csb *chuangcacheSignerBuilder) A(token string) *chuangcacheSignerBuilder {
	csb.signer = newChuangcacheA(token)

	return csb
}

func (csb *chuangcacheSignerBuilder) Build() (cb *chuangcacheBuilder) {
	cb = csb.builder
	cb.signers[csb.domain] = csb.signer

	return
}
