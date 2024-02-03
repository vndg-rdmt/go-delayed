package delayed

import "time"

type Scheduler interface {
	SetInterval(interval time.Duration, d func()) *PlannerTimer
	SetTimeout(timeout time.Duration, d func()) *PlannerTimer
}

type Config struct {
	Epoch time.Time `json:"epoch" yaml:"epoch"`
}

// Default config with epoch equal current time
func DefaultConfig() Config {
	return Config{
		Epoch: time.Now(),
	}
}

// Creates new scheduler.
// Requires epoch, which equals to a starting point,
// all tasks are pontentially must be started from.
func New(c Config) Scheduler {
	return &instance{
		epoch: c.Epoch,
	}
}
