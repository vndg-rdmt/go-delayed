package delayed

import (
	"sync"
	"testing"
	"time"
)

type nowmx struct {
	spent int64
	mx    sync.RWMutex
}

func (self *nowmx) get() int64 {
	self.mx.RLock()
	defer self.mx.RUnlock()
	return self.spent
}

func (self *nowmx) set(n int64) {
	self.mx.Lock()
	defer self.mx.Unlock()
	self.spent = n
}

func TestSetInterval(t *testing.T) {
	d := New(DefaultConfig())

	wg := sync.WaitGroup{}
	wg.Add(5)

	const interval = time.Second * 2
	n := nowmx{
		spent: time.Now().Unix(),
	}

	var c uint8
	p := d.SetInterval(interval, func() {
		if c == 10 {
			return
		}
		defer func() {
			c++
		}()

		now := time.Now().Unix()

		if (now - n.get()) == int64(interval.Seconds()) {
			t.Logf("[passed]")

		} else {
			t.Errorf("[failed]")
		}

		n.set(now)
		wg.Done()
	})

	wg.Wait()
	p.get().Stop()
}

func TestSetTimeout(t *testing.T) {
	d := New(DefaultConfig())

	wg := sync.WaitGroup{}
	wg.Add(1)

	const interval = time.Second * 1
	spent := time.Now().Unix()

	d.SetTimeout(interval, func() {

		now := time.Now().Unix()

		if (now - spent) == int64(interval.Seconds()) {
			t.Logf("[passed]")

		} else {
			t.Errorf("[failed]")
		}

		spent = now
		wg.Done()
	})

	wg.Wait()
}
