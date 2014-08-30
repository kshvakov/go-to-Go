package gopinba

import (
	"fmt"
	"net"
	"os"
	//"code.google.com/p/goprotobuf/proto"
)

type pinba struct {
	pinbaServer string
	pinbaPort   int
}

func (pinba *pinba) Request() *request {

	hostname := "unknown"

	if value, err := os.Hostname(); err == nil {

		hostname = value
	}

	request := &request{
		hostname:   hostname,
		serverName: "unknown",
		scriptName: "unknown",
		timers:     make(map[int]*timer),
	}

	return request
}

func (pinba *pinba) Flush(request *request) {

	conn, _ := net.Dial("udp", fmt.Sprintf("%s:%d", pinba.pinbaServer, pinba.pinbaPort))

	conn.Write([]byte("hello"))
}
