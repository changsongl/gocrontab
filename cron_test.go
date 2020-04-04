package gocrontab

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	s := New()
	s.AddJob("0 * * * * *", func(){fmt.Println(0)})
	s.AddJob("30 * * * * *", func(){fmt.Println(30)})

	s.Start()
	time.Sleep(100000 * time.Second)
}