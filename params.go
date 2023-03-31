package cdn

type params struct {
	executors map[string]executor
}

func newParams() *params {
	return &params{
		executors: make(map[string]executor),
	}
}
