package delayed

import (
	"time"
)

type instance struct {
	epoch time.Time
}

// Schedules new task to be executed on interval, starting
// from the Daemon epoch.
func (self *instance) SetInterval(interval time.Duration, d func()) *time.Timer {

	// Create new task
	newTask := &task{
		time:         interval,
		exec:         d,
		parent:       self,
		plannerTimer: nil,
	}

	// Delay interval execution
	defer newTask.delayIntervalExecution()
	return newTask.plannerTimer
}

// Schedules new task to be executed after timeout, starting
// from the Daemon epoch.
func (self *instance) SetTimeout(timeout time.Duration, d func()) *time.Timer {

	// Create new task
	newTask := &task{
		time:         timeout,
		exec:         d,
		parent:       self,
		plannerTimer: nil,
	}

	// Delay interval execution
	defer newTask.delayExecution()
	return newTask.plannerTimer
}
