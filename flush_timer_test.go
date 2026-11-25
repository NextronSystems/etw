//go:build windows
// +build windows

package etw

import (
	"testing"
	"time"
)

func TestFlushTimerProperties(t *testing.T) {
	cases := []struct {
		interval    time.Duration
		timer       uint32
		logFileMode uint32
	}{
		{0, 0, 0},
		{time.Second, 1, 0},
		{3 * time.Second, 3, 0},
		{100 * time.Millisecond, 100, eventTraceUseMsFlushTimer},
		{1500 * time.Millisecond, 1500, eventTraceUseMsFlushTimer},
	}
	for _, c := range cases {
		timer, mode := flushTimerProperties(c.interval)
		if timer != c.timer || mode != c.logFileMode {
			t.Errorf("flushTimerProperties(%v) = (%d, %#x), want (%d, %#x)", c.interval, timer, mode, c.timer, c.logFileMode)
		}
	}
}
