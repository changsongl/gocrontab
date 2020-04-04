package schduler

import (
	"time"
)

type scheduler struct {
	count int
	c chan int
	f func()
	start chan int
	stop chan int
	t *timer
}

type Closer interface {
	Close()
}

type Scheduler interface {
	Start() Closer
}

func New(sch string, f func()) (*scheduler, error){
	c := make(chan int, 1)
	strt := make(chan int, 1)
	stp := make(chan int, 1)
	t, err := newTimer(sch)
	if err != nil{
		return nil, err
	}

	return &scheduler{0, c, f, strt, stp, t}, nil
}

func (s *scheduler)GetCloser() Closer{
	return NewCloser(s.stop)
}

func (s *scheduler)Start()Closer{
	go s.consume()

	go func(){
		now := time.Now()
		for{
			nt, err :=s.produce(now)
			if err != nil{
				s.GetCloser().Close()
				return
			}
			now = nt
		}
	}()

	return nil
}

func (s *scheduler)produce(t time.Time) (time.Time, error){
	nt, err := s.t.getNextTime(t)
	if err != nil{
		return nt, err
	}
	now := time.Now().Unix()
	next := nt.Unix()
	st := next - now


	if st > 0{
		time.Sleep(time.Second * time.Duration(st))
	}

	s.count += 1
	s.c <- s.count

	return nt, nil
}

func (s *scheduler)consume(){
	for range s.c{
		go s.f()
	}
}



