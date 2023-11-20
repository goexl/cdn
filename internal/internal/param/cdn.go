package param

type Cdn struct {
	Domains map[string]*Domain
}

func NewCdn() *Cdn {
	return &Cdn{
		Domains: make(map[string]*Domain),
	}
}
