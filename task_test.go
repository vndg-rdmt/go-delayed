package delayed

import (
	"sync"
	"testing"
	"time"
)

func TestSetInterval(t *testing.T) {
	d := New(DefaultConfig())

	wg := sync.WaitGroup{}
	wg.Add(10)

	const interval = time.Second * 2
	spent := time.Now().Unix()

	d.SetInterval(interval, func() {

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
