package cdn

type chuangcacheBuilder struct {
	builder  *builder
	domain   *domain
	signer   signer
	patterns []string
}

func newChuangcacheBuilder(builder *builder) *chuangcacheBuilder {
	return &chuangcacheBuilder{
		builder:  builder,
		domain:   newDomain(),
		patterns: make([]string, 0, 1),
	}
}

func (cb *chuangcacheBuilder) Host(host string) *chuangcacheBuilder {
	cb.domain.host = host

	return cb
}

func (cb *chuangcacheBuilder) Http() *chuangcacheBuilder {
	cb.domain.scheme = http

	return cb
}

func (cb *chuangcacheBuilder) Https() *chuangcacheBuilder {
	cb.domain.scheme = https

	return cb
}

func (cb *chuangcacheBuilder) Default() *chuangcacheBuilder {
	cb.patterns = append(cb.patterns, defaults)

	return cb
}

func (cb *chuangcacheBuilder) Pattern(pattern ...string) *chuangcacheBuilder {
	if 0 == len(pattern) {
		cb.patterns = append(cb.patterns, defaults)
	} else {
		cb.patterns = append(cb.patterns, pattern...)
	}

	return cb
}

func (cb *chuangcacheBuilder) Signer() *chuangcacheSignerBuilder {
	return newChuangcacheSignerBuilder(cb)
}

func (cb *chuangcacheBuilder) Build() (b *builder) {
	b = cb.builder
	for _, pattern := range cb.patterns {
		b.params.domains[pattern] = cb.domain
	}

	return
}
