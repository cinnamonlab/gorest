package queue

import (
	"github.com/cinnamonlab/WorkerPool"
)

type Queue struct {
	
}

// NewLauncher returns a new Launcher.
func NewQueue() *Queue {
	return &Queue{}
}

type QueueFunc func(input ...interface{})

var sharedQueue *Queue

func init() {
	sharedQueue = NewQueue()
}

func Start(numberProcess int)  {
	sharedQueue.Start(numberProcess)
}

func (queue *Queue) Start(numberProcess int)  {
	workerpool.Start(numberProcess)
}

func NewTask(function QueueFunc,params ...interface{})  {
	queueTask := QueueJob{
		jobFunc:function,
		jobParams : params,
	}
	workerpool.AddNewTask(queueTask)
}