package cdn

type builder struct {
	params *params
}

func newBuilder() *builder {
	return &builder{
		params: newParams(),
	}
}

func (b *builder) Tencent() *tencentBuilder {
	return newTencentBuilder(b)
}

func (b *builder) Chuangcache() *chuangcacheBuilder {
	return newChuangcacheBuilder(b)
}

func (b *builder) Build() *Encoder {
	return newEncoder(b.params)
}
