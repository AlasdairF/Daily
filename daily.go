package daily

import (
	"time"
)

const (
	day = time.Hour * 24
	halfday = time.Hour * 12
)

type job struct {
	fn func()
	inProgress bool
	wait int64
}

var todo []job

func (j *job) finished() {
	j.inProgress = false
}

func (j *job) do() {
	time.Sleep(j.wait)
	if !j.inProgress {
		j.inProgress = true
		defer j.finished()
		j.fn()
	}
}

func (j *job) doNow() {
	j.inProgress = true
	defer j.finished()
	j.fn()
}

func Run(fn func(), secondsPastMidnight int64, now bool) {
	l := len(todo)
	newar := make([]job, l + 1)
	copy(newar, todo)
	newar[l] = job{fn, false, (secondsPastMidnight % 86400) * time.Second}
	todo = newar
	if now {
		go newar[l].doNow()
	}
}

func daily() {
	for {
		time.Sleep(time.Now().UnixNano() % day)
		for _, j := range todo {
			go j.do()
		}
		time.Sleep(halfday) // just to make sure it doesn't try to do it again if no nanoseconds have passed
	}
}

func init() {
	go daily()
}
