package queue

type QueueJob struct  {
	jobFunc QueueFunc
	jobParams []interface{}
}

func (job QueueJob) Execute()  {
	job.jobFunc(job.jobParams...)
}
