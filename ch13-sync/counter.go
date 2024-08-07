package counter

import "sync"

type SharedInt struct {
	mu  sync.Mutex
	val int
}

func (si *SharedInt) Inc() {
	si.SetVal(si.Value() + 1)
	//si.SetVal(si.val + 1)
}

func (si *SharedInt) SetVal(val int) {
	si.mu.Lock()
	defer si.mu.Unlock()
	si.val = val
}

func (si *SharedInt) Value() int {
	si.mu.Lock()
	defer si.mu.Unlock()
	return si.val
}

type Counter interface {
	Inc()
	Value() int
}

type CounterWithMutex struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *CounterWithMutex {
	return &CounterWithMutex{}
}

func (c *CounterWithMutex) Set(val int) {
	c.mu.Lock()
	defer c.mu.Unlock()

}

func (c *CounterWithMutex) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *CounterWithMutex) Value() int {
	return c.value
}

type CounterWithBufferedChan struct {
	ch  chan any
	val int
}

func NewCounterWithChan() *CounterWithBufferedChan {
	return &CounterWithBufferedChan{
		ch:  make(chan any, 1),
		val: 0,
	}
}
func (c *CounterWithBufferedChan) Inc() {
	c.ch <- struct{}{}
	defer func() { <-c.ch }()
	c.val++
}

func (c *CounterWithBufferedChan) Value() int {
	return c.val
}

type CounterWithUnbufferedChan struct {
	ch  chan any
	val int
}

func NewCounterWithUnbufferedChan() *CounterWithUnbufferedChan {

	c := &CounterWithUnbufferedChan{
		ch:  make(chan any),
		val: 0,
	}

	go c.incrementer()

	return c
}

func (c *CounterWithUnbufferedChan) incrementer() {
	for {
		<-c.ch
		c.val++
	}
}

func (c *CounterWithUnbufferedChan) Inc() {
	c.ch <- struct{}{}
}

func (c *CounterWithUnbufferedChan) Value() int {
	return c.val
}
