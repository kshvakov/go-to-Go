package gopinba

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPinba(t *testing.T) {
	pinba := New(&Options{})

	request := pinba.Request()
	request.SetScriptName("scriptName")

	var timerID int

	for i := 0; i < 20; i++ {

		timerID = request.TimerStart(&Tags{"tag": fmt.Sprintf("tag_%d", i)})

		if i == 4 {

			buf := []int{}

			for k := 0; k < 42; k++ {

				buf = append(buf, k)
			}
		}

		fmt.Println(timerID)

		request.TimerStop(timerID)
	}

	request.TimerStop(42)

	fmt.Println(request.GetInfo())

	pinba.Flush(request)

	assert.Equal(t, 1, 1)
}

func BenchmarkPinba(b *testing.B) {

	pinba := New(&Options{})

	request := pinba.Request()

	var timerID int

	for i := 0; i < b.N; i++ {

		timerID = request.TimerStart(&Tags{"tag": "a"})

		request.TimerStop(timerID)
	}

	pinba.Flush(request)
}
