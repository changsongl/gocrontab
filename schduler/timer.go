package schduler

import (
	"github.com/influxdata/cron"
	"time"
)

type timer struct {
	c cron.Parsed
}

func newTimer(s string) (*timer, error){
	p, err := cron.ParseUTC(s)
	return &timer{p}, err
}

func (t *timer)getNextTime(time time.Time) (time.Time, error){
	return t.c.Next(time)
}
