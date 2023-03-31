package cdn

type tencentBuilder struct {
	builder  *builder
	domain   *domain
	signer   signer
	patterns []string
}

func newTencentBuilder(builder *builder) *tencentBuilder {
	return &tencentBuilder{
		builder:  builder,
		domain:   newDomain(),
		patterns: make([]string, 0, 1),
	}
}

func (tb *tencentBuilder) Host(host string) *tencentBuilder {
	tb.domain.host = host

	return tb
}

func (tb *tencentBuilder) Http() *tencentBuilder {
	tb.domain.scheme = http

	return tb
}

func (tb *tencentBuilder) Https() *tencentBuilder {
	tb.domain.scheme = https

	return tb
}

func (tb *tencentBuilder) Default() *tencentBuilder {
	tb.patterns = append(tb.patterns, defaults)

	return tb
}

func (tb *tencentBuilder) Pattern(pattern ...string) *tencentBuilder {
	if 0 == len(pattern) {
		tb.patterns = append(tb.patterns, defaults)
	} else {
		tb.patterns = append(tb.patterns, pattern...)
	}

	return tb
}

func (tb *tencentBuilder) Signer() *tencentSignerBuilder {
	return newTencentSignerBuilder(tb)
}

func (tb *tencentBuilder) Build() (b *builder) {
	b = tb.builder
	for _, pattern := range tb.patterns {
		b.params.domains[pattern] = tb.domain
	}

	return
}
