// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Timer] class.
var (
	timerClass     _TimerClass
	timerClassOnce sync.Once
)

func getTimerClass() _TimerClass {
	timerClassOnce.Do(func() {
		timerClass = _TimerClass{objc.GetClass("NSTimer")}
	})
	return timerClass
}

type _TimerClass struct {
	class objc.Class
}

// An interface definition for the [Timer] class.
type ITimer interface {
	objectivec.IObject
	Fire()
	Invalidate()
}

// A timer that fires after a certain time interval has elapsed, sending a specified message to a target object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer
type Timer struct {
	objectivec.Object
}

// TimerFrom constructs a [Timer] from an unsafe.Pointer.
//
// A timer that fires after a certain time interval has elapsed, sending a specified message to a target object.
func TimerFrom(ptr unsafe.Pointer) Timer {
	return Timer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TimerClass) Alloc() Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TimerClass) New() Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Timer) Init() Timer {
	rv := objc.Send[Timer](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Timer) Autorelease() Timer {
	rv := objc.Send[Timer](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTimer creates a new Timer instance.
func NewTimer() Timer {
	return getTimerClass().New()
}


// Initializes a timer object with the specified time interval and block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:repeats:block:)
func NewTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block unsafe.Pointer) Timer {
	rv := objc.Send[Timer](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}
// Initializes a timer object with the specified object and selector. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:target:selector:userInfo:repeats:)
func NewTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.ID, aSelector objc.SEL, userInfo objc.ID, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}
// Initializes a timer for the specified date and time interval with the specified block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(fire:interval:repeats:block:)
func NewTimerWithFireDateIntervalRepeatsBlock(date unsafe.Pointer, interval TimeInterval, repeats bool, block unsafe.Pointer) Timer {
	instance := getTimerClass().Alloc()
	rv := objc.Send[Timer](instance.ID, objc.Sel("initWithFireDate:interval:repeats:block:"), date, interval, repeats, block)
	rv.Autorelease()
	return rv
}
// Initializes a timer using the specified object and selector. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(fireAt:interval:target:selector:userInfo:repeats:)
func NewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats(date unsafe.Pointer, ti TimeInterval, t objc.ID, s objc.SEL, ui objc.ID, rep bool) Timer {
	instance := getTimerClass().Alloc()
	rv := objc.Send[Timer](instance.ID, objc.Sel("initWithFireDate:interval:target:selector:userInfo:repeats:"), date, ti, t, s, ui, rep)
	rv.Autorelease()
	return rv
}
// Initializes a timer object with the specified invocation object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:invocation:repeats:)
func NewTimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation unsafe.Pointer, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}


// Initializes a timer object with the specified invocation object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:invocation:repeats:)
func (tc _TimerClass) TimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation unsafe.Pointer, yesOrNo bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("timerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}
// Initializes a timer object with the specified time interval and block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:repeats:block:)
func (tc _TimerClass) TimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("timerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}
// Initializes a timer object with the specified object and selector. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:target:selector:userInfo:repeats:)
func (tc _TimerClass) TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.ID, aSelector objc.SEL, userInfo objc.ID, yesOrNo bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("timerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}
// Creates a new timer and schedules it on the current run loop in the default mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(timeInterval:invocation:repeats:)
func (tc _TimerClass) ScheduledTimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation unsafe.Pointer, yesOrNo bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("scheduledTimerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}
// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(timeInterval:target:selector:userInfo:repeats:)
func (tc _TimerClass) ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.ID, aSelector objc.SEL, userInfo objc.ID, yesOrNo bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("scheduledTimerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}
// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(withTimeInterval:repeats:block:)
func (tc _TimerClass) ScheduledTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("scheduledTimerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}
// Causes the timer’s message to be sent to its target. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/fire()
func (t_ Timer) Fire() {
	objc.Send[objc.ID](t_.ID, objc.Sel("fire"))
}
// Stops the timer from ever firing again and requests its removal from its run loop. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/invalidate()
func (t_ Timer) Invalidate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidate"))
}

