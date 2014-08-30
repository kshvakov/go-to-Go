package gopinba

import (
	//"fmt"
	"runtime"
	"time"
)

var memStats runtime.MemStats

type request struct {
	timerID    int
	hostname   string
	serverName string
	scriptName string

	timers map[int]*timer
}

func (request *request) SetServerName(serverName string) {

	request.serverName = serverName
}

func (request *request) SetScriptName(scriptName string) {

	request.scriptName = scriptName
}

func (request *request) TimerStart(tags *Tags) int {

	request.timerID++

	timerID := request.timerID

	runtime.ReadMemStats(&memStats)

	request.timers[timerID] = &timer{
		started:     true,
		timeStart:   time.Now(),
		tags:        tags,
		memoryUsage: memStats.TotalAlloc,
	}

	return timerID
}

func (request *request) TimerStop(timerID int) {

	if timer, found := request.timers[timerID]; found {

		runtime.ReadMemStats(&memStats)
		timer.started = false
		timer.timeEnd = time.Now()
		timer.memoryUsage = memStats.TotalAlloc - timer.memoryUsage
	}
}

func (request *request) GetInfo() *request {
	return request
}
