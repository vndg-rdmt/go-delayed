package delayed

import "time"

// Scheduler task
type task struct {

	// Time, which describes task's execution delay
	// in the context of planning call.
	// It can be timeout, interval and etc.
	time time.Duration

	// Call, which is delayed by the task.
	exec func()

	// Parent daemon proccess, which resources task using to delay
	// self task execution properly.
	parent *instance

	// plannerTimer controller to manipulate delayed execution.
	plannerTimer *time.Timer
}

// Executes task and delays new execution.
func (self *task) execAndDelayNew() {
	self.exec()
	self.delayIntervalExecution()
}

// Delays execution to the next self.time stop.
func (self *task) delayExecution() {
	self.plannerTimer = time.AfterFunc(
		self.parent.expectedTime(time.Now(), self.time), self.exec,
	)
}

// Delays execution and planning of the next call of
// the task to the next self.time stop.
func (self *task) delayIntervalExecution() {
	self.plannerTimer = time.AfterFunc(
		self.parent.expectedTime(time.Now(), self.time), self.execAndDelayNew,
	)
}
