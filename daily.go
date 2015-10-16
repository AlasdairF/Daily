package daily

import (
	"time"
)

const day = time.Hour * 24

type job struct {
	fn func()
	inProgress bool
}

var todo []job

func (j *job) finished() {
	j.inProgress = false
}

func (j *job) do() {
	if !j.inProgress {
		j.inProgress = true
		defer j.finished()
		j.fn()
	}
}

func Run(fn func(), now bool) {
	l := len(todo)
	newar := make([]job, l + 1)
	copy(newar, todo)
	newar[l] = job{fn, false}
	todo = newar
	if now {
		go todo[l].do()
	}
}

func daily() {
	for {
		time.Sleep(time.Duration(time.Now().UnixNano()) % day)
		for _, j := range todo {
			go j.do()
		}
	}
}

func init() {
	go daily()
}
