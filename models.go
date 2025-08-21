package cron_manager

import "github.com/go-co-op/gocron/v2"

type CronManager struct {
	OriginalCron interface{}
	Jobs         map[string]string // id/name
	crons
}

//	type job struct {
//		Id   string
//		Name string
//	}
type crons struct {
	gocoopV2 gocron.Scheduler
}

type MainPage struct {
	TotalJobs int
	Running   int
	Stopped   int
	Jobs      []JobPage `json:"jobs"`
}

type JobPage struct {
	Id      string
	Name    string
	LastRun string
	NextRun string
}

//type CronRegisters
