package gocrontab

import (
	"github.com/ChangsongLiQD/gocrontab/schduler"
)

type cron struct {
	s []schduler.Scheduler
}

type Crontab interface {
	AddJob(s string, f func()) error
	Start()
}

func New() Crontab{
	return &cron{[]schduler.Scheduler{}}
}

func (c *cron)AddJob(s string, f func()) error{
	sch, err := schduler.New(s, f)
	if err != nil{
		return err
	}

	c.s = append(c.s, sch)
	return nil
}

func (c *cron)Start(){
	for _, s := range c.s{
		s.Start()
	}
}
