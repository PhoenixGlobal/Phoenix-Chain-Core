package fetcher

import (
	"sync"
	"time"

	"Phoenix-Chain-Core/consensus/pbft/types"
)

var (
	arriveTimeout = 500 * time.Millisecond
)

// SetArriveTimeout set timeout.
func SetArriveTimeout(duration time.Duration) {
	arriveTimeout = duration
}

// MatchFunc is a function that judges the matching of messages.
type MatchFunc func(types.Message) bool

// ExecutorFunc defines the execution function.
type ExecutorFunc func(types.Message)

// ExpireFunc defines the timeout execution function.
type ExpireFunc func()

type task struct {
	id string
	// Specify whether the message matches the task
	match MatchFunc
	// Callback executed function
	executor ExecutorFunc

	// Timeout callback function
	expire ExpireFunc
	// Task addition time
	time time.Time
}

// Fetcher manages the logic associated with fetch.
type Fetcher struct {
	lock    sync.Mutex
	newTask chan *task
	quit    chan struct{}
	tasks   map[string]*task
}

// NewFetcher returns a new pointer to the Fetcher.
func NewFetcher() *Fetcher {
	fetcher := &Fetcher{
		newTask: make(chan *task, 1),
		tasks:   make(map[string]*task),
		quit:    make(chan struct{}),
	}
	return fetcher
}

// Start turns on for Fetch.
func (f *Fetcher) Start() {
	go f.loop()
}

// Stop turns off for Fetch.
func (f *Fetcher) Stop() {
	close(f.quit)
}

// AddTask adds a fetcher task.
func (f *Fetcher) AddTask(id string, match MatchFunc, executor ExecutorFunc, expire ExpireFunc) {
	select {
	case <-f.quit:
	case f.newTask <- &task{id: id, match: match, executor: executor, expire: expire, time: time.Now()}:
	}
}

// MatchTask matching task.
func (f *Fetcher) MatchTask(id string, message types.Message) bool {
	f.lock.Lock()
	defer f.lock.Unlock()
	if t, ok := f.tasks[id]; ok {
		if t.match(message) {
			go t.executor(message)
			delete(f.tasks, id)
			return true
		}
	}
	return false
}

// Len returns the number of existing tasks.
func (f *Fetcher) Len() int {
	f.lock.Lock()
	defer f.lock.Unlock()
	return len(f.tasks)
}

// The main logic of fetcher, listening to tasks that require
// fetcher and continuous processing.Simultaneously delete expired tasks.
func (f *Fetcher) loop() {
	fetchTimer := time.NewTimer(0)
	for {
		select {
		case task := <-f.newTask:
			f.lock.Lock()
			if len(f.tasks) == 0 {
				fetchTimer.Reset(arriveTimeout)
			}
			f.tasks[task.id] = task
			f.lock.Unlock()

		case <-fetchTimer.C:
			f.lock.Lock()
			for id, task := range f.tasks {
				if time.Since(task.time) > arriveTimeout {
					if task.expire != nil {
						task.expire()
					}
					delete(f.tasks, id)
				}
			}
			if len(f.tasks) == 0 {
				fetchTimer.Stop()
			} else {
				fetchTimer.Reset(arriveTimeout)
			}
			f.lock.Unlock()
		case <-f.quit:
			f.lock.Lock()
			f.tasks = make(map[string]*task)
			fetchTimer.Stop()
			f.lock.Unlock()
			return
		}
	}
}
