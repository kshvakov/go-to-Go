package gopinba

func New(options *Options) *pinba {

	pinba := &pinba{pinbaServer: "127.0.0.1", pinbaPort: 30002}

	return pinba
}
