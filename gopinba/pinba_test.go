package gopinba

import (
	"fmt"
	//"github.com/stretchr/testify/assert"
	"testing"
)

func TestPinba(t *testing.T) {
	pinba := New(&Options{})

	request := pinba.Request()
	request.SetScriptName("scriptName")

	var tmr *timer

	for i := 0; i < 20; i++ {

		tmr = request.TimerStart(&Tags{"tag": fmt.Sprintf("tag_%d", i)})

		if i == 4 {

			buf := []int{}

			for k := 0; k < 42; k++ {

				buf = append(buf, k)
			}
		}

		request.TimerStop(tmr)
	}

	info := request.GetInfo()

	for _, t := range info.timers {

		fmt.Printf("%#v\n", *t)
	}

	pinba.Flush(request)

	//assert.Equal(t, 1, 1)
}

func BenchmarkPinba(b *testing.B) {

	pinba := New(&Options{})

	var tmr *timer

	for i := 0; i < b.N; i++ {

		request := pinba.Request()

		for i := 0; i < 10; i++ {

			tmr = request.TimerStart(&Tags{"tag": "a"})

			request.TimerStop(tmr)
		}
	}
}
