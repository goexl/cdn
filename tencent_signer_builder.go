package cdn

type tencentSignerBuilder struct {
	builder *tencentBuilder
	domain  string
	signer  signer
}

func newTencentSignerBuilder(builder *tencentBuilder) *tencentSignerBuilder {
	return &tencentSignerBuilder{
		builder: builder,
	}
}

func (tsb *tencentSignerBuilder) Default() *tencentSignerBuilder {
	return tsb.Domain(defaults)
}

func (tsb *tencentSignerBuilder) Any() *tencentSignerBuilder {
	return tsb.Domain(defaults)
}

func (tsb *tencentSignerBuilder) Domain(domain string) *tencentSignerBuilder {
	tsb.domain = domain

	return tsb
}

func (tsb *tencentSignerBuilder) A(key string) *tencentSignerBuilder {
	tsb.signer = newTencentA(key)

	return tsb
}

func (tsb *tencentSignerBuilder) B(key string) *tencentSignerBuilder {
	tsb.signer = newTencentB(key)

	return tsb
}

func (tsb *tencentSignerBuilder) C(key string) *tencentSignerBuilder {
	tsb.signer = newTencentC(key)

	return tsb
}

func (tsb *tencentSignerBuilder) D(key string, sign string, timestamp string) *tencentSignerBuilder {
	tsb.signer = newTencentD(key, sign, timestamp)

	return tsb
}

func (tsb *tencentSignerBuilder) Build() (tb *tencentBuilder) {
	tb = tsb.builder
	tb.domain.signer = tsb.signer

	return
}
