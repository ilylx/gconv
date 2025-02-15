// Package gtimer implements Hierarchical Timing Wheel for interval/delayed jobs
// running and management.
//
// This package is designed for management for millions of timing jobs. The differences
// between gtimer and gcron are as follows:
//  1. package gcron is implemented based on package gtimer.
//  2. gtimer is designed for high performance and for millions of timing jobs.
//  3. gcron supports configuration pattern grammar like linux crontab, which is more manually
//     readable.
//  4. gtimer's benchmark OP is measured in nanoseconds, and gcron's benchmark OP is measured
//     in microseconds.
//
// ALSO VERY NOTE the common delay of the timer: https://github.com/golang/go/issues/14410
package gtimer

import (
	"fmt"
	"github.com/ilylx/gconv/internal/cmdenv"
	"math"
	"time"
)

const (
	StatusReady           = 0             // Job is ready for running.
	StatusRunning         = 1             // Job is already running.
	StatusStopped         = 2             // Job is stopped.
	StatusReset           = 3             // Job is reset.
	StatusClosed          = -1            // Job is closed and waiting to be deleted.
	gPanicExit            = "exit"        // Internal usage for custom job exit function with panic.
	gDefaultTimes         = math.MaxInt32 // Default limit running times, a big number.
	gDefaultSlotNumber    = 10            // Default slot number.
	gDefaultWheelInterval = 50            // Default wheel interval.
	gDefaultWheelLevel    = 6             // Default wheel level.
	gCmdenvKey            = "gf.gtimer"   // Configuration key for command argument or environment.
)

var (
	defaultSlots    = cmdenv.Get(fmt.Sprintf("%s.slots", gCmdenvKey), gDefaultSlotNumber).Int()
	defaultLevel    = cmdenv.Get(fmt.Sprintf("%s.level", gCmdenvKey), gDefaultWheelLevel).Int()
	defaultInterval = cmdenv.Get(fmt.Sprintf("%s.interval", gCmdenvKey), gDefaultWheelInterval).Duration() * time.Millisecond
	defaultTimer    = New(defaultSlots, defaultInterval, defaultLevel)
)

// SetTimeout runs the job once after duration of <delay>.
// It is like the one in javascript.
func SetTimeout(delay time.Duration, job JobFunc) {
	AddOnce(delay, job)
}

// SetInterval runs the job every duration of <delay>.
// It is like the one in javascript.
func SetInterval(interval time.Duration, job JobFunc) {
	Add(interval, job)
}

// Add adds a timing job to the default timer, which runs in interval of <interval>.
func Add(interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.Add(interval, job)
}

// AddEntry adds a timing job to the default timer with detailed parameters.
//
// The parameter <interval> specifies the running interval of the job.
//
// The parameter <singleton> specifies whether the job running in singleton mode.
// There's only one of the same job is allowed running when its a singleton mode job.
//
// The parameter <times> specifies limit for the job running times, which means the job
// exits if its run times exceeds the <times>.
//
// The parameter <status> specifies the job status when it's firstly added to the timer.
func AddEntry(interval time.Duration, job JobFunc, singleton bool, times int, status int) *Entry {
	return defaultTimer.AddEntry(interval, job, singleton, times, status)
}

// AddSingleton is a convenience function for add singleton mode job.
func AddSingleton(interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.AddSingleton(interval, job)
}

// AddOnce is a convenience function for adding a job which only runs once and then exits.
func AddOnce(interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.AddOnce(interval, job)
}

// AddTimes is a convenience function for adding a job which is limited running times.
func AddTimes(interval time.Duration, times int, job JobFunc) *Entry {
	return defaultTimer.AddTimes(interval, times, job)
}

// DelayAdd adds a timing job after delay of <interval> duration.
// Also see Add.
func DelayAdd(delay time.Duration, interval time.Duration, job JobFunc) {
	defaultTimer.DelayAdd(delay, interval, job)
}

// DelayAddEntry adds a timing job after delay of <interval> duration.
// Also see AddEntry.
func DelayAddEntry(delay time.Duration, interval time.Duration, job JobFunc, singleton bool, times int, status int) {
	defaultTimer.DelayAddEntry(delay, interval, job, singleton, times, status)
}

// DelayAddSingleton adds a timing job after delay of <interval> duration.
// Also see AddSingleton.
func DelayAddSingleton(delay time.Duration, interval time.Duration, job JobFunc) {
	defaultTimer.DelayAddSingleton(delay, interval, job)
}

// DelayAddOnce adds a timing job after delay of <interval> duration.
// Also see AddOnce.
func DelayAddOnce(delay time.Duration, interval time.Duration, job JobFunc) {
	defaultTimer.DelayAddOnce(delay, interval, job)
}

// DelayAddTimes adds a timing job after delay of <interval> duration.
// Also see AddTimes.
func DelayAddTimes(delay time.Duration, interval time.Duration, times int, job JobFunc) {
	defaultTimer.DelayAddTimes(delay, interval, times, job)
}

// Exit is used in timing job internally, which exits and marks it closed from timer.
// The timing job will be automatically removed from timer later. It uses "panic-recover"
// mechanism internally implementing this feature, which is designed for simplification
// and convenience.
func Exit() {
	panic(gPanicExit)
}
