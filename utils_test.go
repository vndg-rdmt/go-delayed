package delayed

import (
	"testing"
	"time"
)

func newDummyDaemon(now time.Time) *instance {
	return New(Config{Epoch: now}).(*instance)
}

type testCaseExpectedTime struct {
	epoch    time.Time
	now      time.Time
	interval time.Duration
	expected time.Duration
}

func testsExpectedTime(now time.Time) []*testCaseExpectedTime {
	return []*testCaseExpectedTime{
		{
			epoch:    now,
			now:      now.Add(time.Hour),
			interval: time.Hour,
			expected: time.Hour,
		},
		{
			epoch:    now,
			now:      now.Add(time.Hour - time.Minute),
			interval: time.Hour,
			expected: time.Minute,
		},
		{
			epoch:    now,
			now:      now.Add(time.Minute * 30),
			interval: time.Hour,
			expected: time.Minute * 30,
		},
	}
}

func TestExpectedTime(t *testing.T) {
	now := time.Now()

	for i, testCase := range testsExpectedTime(now) {
		d := newDummyDaemon(now)

		if d.expectedTime(testCase.now, testCase.interval) == testCase.expected {
			t.Logf("[test %d] passed", i)
		} else {
			t.Errorf("[test %d] failed (got %d, want %d)", i, d.expectedTime(testCase.now, testCase.interval), testCase.expected)
		}
	}
}
