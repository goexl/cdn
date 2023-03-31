package cdn

type params struct {
	signers map[string]signer
}

func newParams() *params {
	return &params{
		signers: make(map[string]signer),
	}
}
