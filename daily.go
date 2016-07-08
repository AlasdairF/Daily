package daily

import (
	"time"
	"log"
)

const (
	day = time.Hour * 24
)

var (
	todo []job
	logging bool
)

type job struct {
	fn func()
	inProgress bool
	wait time.Duration
	name string
}

func (j *job) finished() {
	j.inProgress = false
	if logging {
		log.Println(`Daily: Finished`, j.name)
	}
}

func (j *job) do() {
	time.Sleep(j.wait)
	if j.inProgress {
		if logging {
			log.Println(`Daily: Already in progress`, j.name)
		}
	} else {
		j.inProgress = true
		if logging {
			log.Println(`Daily: Running`, j.name)
		}
		defer j.finished()
		j.fn()
	}
}

func (j *job) doNow() {
	j.inProgress = true
	if logging {
		log.Println(`Daily: Running`, j.name)
	}
	defer j.finished()
	j.fn()
}

func EnableLogging() {
	logging = true
}

func Run(name string, fn func(), secondsPastMidnight time.Duration, now bool) {
	l := len(todo)
	newar := make([]job, l + 1)
	copy(newar, todo)
	newar[l] = job{fn:fn, inProgress:false, wait:secondsPastMidnight * time.Second, name:name}
	todo = newar
	if now {
		go newar[l].doNow()
	}
}

func daily() {
	for {
		time.Sleep(day - (time.Duration(time.Now().UnixNano()) % day))
		for _, j := range todo {
			go j.do()
		}
		time.Sleep(time.Minute) // just to make sure it doesn't try to do it again if no nanoseconds have passed
	}
}

func init() {
	go daily()
}
