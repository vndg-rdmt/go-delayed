package delayed

import "time"

func (self *instance) expectedTime(now time.Time, interval time.Duration) time.Duration {
	return interval - ((now.Sub(self.epoch)) % interval)
}
