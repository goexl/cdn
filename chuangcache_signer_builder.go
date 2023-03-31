package cdn

type chuangcacheSignerBuilder struct {
	builder *chuangcacheBuilder
	signer  signer
}

func newChuangcacheSignerBuilder(builder *chuangcacheBuilder) *chuangcacheSignerBuilder {
	return &chuangcacheSignerBuilder{
		builder: builder,
	}
}

func (csb *chuangcacheSignerBuilder) A(token string) *chuangcacheSignerBuilder {
	csb.signer = newChuangcacheA(token)

	return csb
}

func (csb *chuangcacheSignerBuilder) Build() (cb *chuangcacheBuilder) {
	cb = csb.builder
	cb.signer = csb.signer

	return
}
