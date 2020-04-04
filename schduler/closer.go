package schduler

type closer struct {
	stop chan int
}

func NewCloser(c chan int) *closer{
	return &closer{c}
}

func (c *closer) Close(){
	c.stop <- 1
}
