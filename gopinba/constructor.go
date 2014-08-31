package gopinba

import (
	"os"
)

func New(options *Options) *pinba {

	hostname := "unknown"

	if value, err := os.Hostname(); err == nil {

		hostname = value
	}

	pinba := &pinba{
		hostname:    hostname,
		serverName:  "unknown",
		pinbaServer: "127.0.0.1",
		pinbaPort:   30002,
	}

	return pinba
}
