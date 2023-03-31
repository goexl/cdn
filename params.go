package cdn

type params struct {
	domains map[string]*domain
}

func newParams() *params {
	return &params{
		domains: make(map[string]*domain),
	}
}
