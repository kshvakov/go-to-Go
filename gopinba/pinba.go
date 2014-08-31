package gopinba

import (
	"fmt"
	"net"
	"time"
	//"code.google.com/p/goprotobuf/proto"
)

type pinba struct {
	pinbaServer string
	pinbaPort   int
	hostname    string
	serverName  string
}

func (pinba *pinba) Request() *request {

	request := &request{
		timeStart:  time.Now(),
		scriptName: "unknown",
		//timers:     make([]*timer, 0, 10),
	}

	return request
}

func (pinba *pinba) Flush(request *request) {

	conn, _ := net.Dial("udp", fmt.Sprintf("%s:%d", pinba.pinbaServer, pinba.pinbaPort))

	conn.Write([]byte("hello"))
}
