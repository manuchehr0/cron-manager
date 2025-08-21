package cron_manager

import (
	"errors"
	"github.com/go-co-op/gocron/v2"
	"log"
	"time"
)

var (
	manager *CronManager
)

func DefineCron(cron interface{}) (*CronManager, error) {
	manager = &CronManager{OriginalCron: cron, Jobs: make(map[string]string)}

	// check if cron is supported
	switch c := cron.(type) {
	case gocron.Scheduler:
		manager.crons.gocoopV2 = c
	default:
		return nil, errors.New("not supported Cron: see full supported crons in https://github.....")
	}
	log.Println("successfully mounted the cron")
	return manager, nil
}

//	func (m *CronManager) Register(jobIds ...string) {
//		m.Jobs = append(m.Jobs, jobIds...)
//	}
func (m *CronManager) Register(id, name string) {
	// todo do it using map
	//m.Jobs = append(m.Jobs, job{
	//	Id:   id,
	//	Name: name,
	//})
	m.Jobs[id] = name
}
func (m *CronManager) runJobNow(id string) error {
	switch m.OriginalCron.(type) {
	case gocron.Scheduler:
		for _, j := range m.gocoopV2.Jobs() {
			if _, ok := m.Jobs[j.ID().String()]; ok && j.ID().String() == id {
				log.Println("id:", id, ok)
				if err := j.RunNow(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m *CronManager) retrieveJobsData() *MainPage {
	// todo add check if manager is empty return page that says manager is empty
	res := new(MainPage)
	switch m.OriginalCron.(type) {
	case gocron.Scheduler:
		for _, j := range m.gocoopV2.Jobs() {
			if jobName, ok := m.Jobs[j.ID().String()]; ok {
				var lastRunStr, nextRunStr string
				lastRun, err := j.LastRun()
				if err != nil {
					lastRunStr = "unknown"
				}
				nextRun, err := j.NextRun()
				if err != nil {
					nextRunStr = "unknown"
				}
				lastRunStr = lastRun.Format(time.DateTime)
				nextRunStr = nextRun.Format(time.DateTime)
				res.TotalJobs++
				res.Jobs = append(res.Jobs, JobPage{
					Id:      j.ID().String(),
					Name:    jobName,
					LastRun: lastRunStr,
					NextRun: nextRunStr,
				})
			}
		}
	}
	return res
}

// 	if Contains(m.Jobs, j.ID().String()) {
//
