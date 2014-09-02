package gopinba

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInSlice(t *testing.T) {

	slice := []string{"a", "b", "c", "d", "e"}

	position, exist := inSlice("c", slice)

	assert.True(t, exist)
	assert.Equal(t, 2, position)

	position, exist = inSlice("z", slice)

	assert.False(t, exist)
	assert.Equal(t, position, -1)
}

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

	tmr = request.TimerStart(&Tags{"tag": "tag_1", "operation": "select"})
	request.TimerStop(tmr)

	result := pinba.Flush(request)

	fmt.Println(result)

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
