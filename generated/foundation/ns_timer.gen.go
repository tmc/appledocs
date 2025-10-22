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
	TimerClass     _TimerClass
	TimerClassOnce sync.Once
)

func getTimerClass() _TimerClass {
	TimerClassOnce.Do(func() {
		TimerClass = _TimerClass{objc.GetClass("NSTimer")}
	})
	return TimerClass
}

type _TimerClass struct {
	class objc.Class
}

// An interface definition for the [Timer] class.
type ITimer interface {
	objectivec.IObject
	Fire()
	Invalidate()
	FireDate() NSDate
	SetFireDate(value IDate)
	Valid() bool
	TimeInterval() TimeInterval
	Tolerance() TimeInterval
	SetTolerance(value ITimeInterval)
	UserInfo() objc.ID
	IsValid() bool
	SetIsValid(value bool)
}

// A timer that fires after a certain time interval has elapsed, sending a specified message to a target object.
//
// Timers work in conjunction with run loops. Run loops maintain strong references to their timers, so you don’t have to maintain your own strong reference to a timer after you have added it to a run loop. To use a timer effectively, you should be aware of how run loops operate. See for more information. A timer is not a real-time mechanism. If a timer’s firing time occurs during a long run loop callout or while the run loop is in a mode that isn’t monitoring the timer, the timer doesn’t fire until the next time the run loop checks the timer. Therefore, the actual time at which a timer fires can be significantly later. See also . is toll-free bridged with its Core Foundation counterpart, . See for more information.


// A timer that fires after a certain time interval has elapsed, sending a specified message to a target object.
//
// [Full Topic]
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



// Initializes a timer for the specified date and time interval with the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(fire:interval:repeats:block:)
func NewTimerWithFireDateIntervalRepeatsBlock(date IDate, interval TimeInterval, repeats bool, block unsafe.Pointer) Timer {
	instance := getTimerClass().Alloc()
	rv := objc.Send[Timer](instance.ID, objc.Sel("initWithFireDate:interval:repeats:block:"), date, interval, repeats, block)
	rv.Autorelease()
	return rv
}


// Initializes a timer using the specified object and selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(fireAt:interval:target:selector:userInfo:repeats:)
func NewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats(date IDate, ti TimeInterval, t objectivec.IObject, s objc.SEL, ui objectivec.IObject, rep bool) Timer {
	instance := getTimerClass().Alloc()
	rv := objc.Send[Timer](instance.ID, objc.Sel("initWithFireDate:interval:target:selector:userInfo:repeats:"), date, ti, t, s, ui, rep)
	rv.Autorelease()
	return rv
}


// Initializes a timer object with the specified invocation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:invocation:repeats:)
func NewTimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation IInvocation, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}


// Initializes a timer object with the specified time interval and block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:repeats:block:)
func NewTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block unsafe.Pointer) Timer {
	rv := objc.Send[Timer](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}


// Initializes a timer object with the specified object and selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:target:selector:userInfo:repeats:)
func NewTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objectivec.IObject, aSelector objc.SEL, userInfo objectivec.IObject, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}



// Initializes a timer object with the specified invocation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:invocation:repeats:)
func (tc _TimerClass) TimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation IInvocation, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("timerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}


// Initializes a timer object with the specified time interval and block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:repeats:block:)
func (tc _TimerClass) TimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block unsafe.Pointer) Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("timerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}


// Initializes a timer object with the specified object and selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:target:selector:userInfo:repeats:)
func (tc _TimerClass) TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objectivec.IObject, aSelector objc.SEL, userInfo objectivec.IObject, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("timerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}


// Creates a new timer and schedules it on the current run loop in the default mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(timeInterval:invocation:repeats:)
func (tc _TimerClass) ScheduledTimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation IInvocation, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("scheduledTimerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}


// Creates a timer and schedules it on the current run loop in the default mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(timeInterval:target:selector:userInfo:repeats:)
func (tc _TimerClass) ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objectivec.IObject, aSelector objc.SEL, userInfo objectivec.IObject, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("scheduledTimerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}


// Creates a timer and schedules it on the current run loop in the default mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(withTimeInterval:repeats:block:)
func (tc _TimerClass) ScheduledTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block unsafe.Pointer) Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("scheduledTimerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}


// Causes the timer’s message to be sent to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/fire()
func (t_ Timer) Fire() {
	objc.Send[objc.ID](t_.ID, objc.Sel("fire"))
}


// Stops the timer from ever firing again and requests its removal from its run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/invalidate()
func (t_ Timer) Invalidate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidate"))
}


// The date at which the timer will fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/fireDate
func (t_ Timer) FireDate() NSDate {
	rv := objc.Send[NSDate](t_.ID, objc.Sel("fireDate"))
	return rv
}


// The date at which the timer will fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/fireDate
func (t_ Timer) SetFireDate(value IDate) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFireDate:"), value)
}


// A Boolean value that indicates whether the timer is currently valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/isValid
func (t_ Timer) Valid() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("valid"))
	return rv
}


// The timer’s time interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/timeInterval
func (t_ Timer) TimeInterval() TimeInterval {
	rv := objc.Send[TimeInterval](t_.ID, objc.Sel("timeInterval"))
	return rv
}


// The amount of time after the scheduled fire date that the timer may fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/tolerance
func (t_ Timer) Tolerance() TimeInterval {
	rv := objc.Send[TimeInterval](t_.ID, objc.Sel("tolerance"))
	return rv
}


// The amount of time after the scheduled fire date that the timer may fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/tolerance
func (t_ Timer) SetTolerance(value ITimeInterval) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTolerance:"), value)
}


// The receiver’s object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/userInfo
func (t_ Timer) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("userInfo"))
	return rv
}


// A Boolean value that indicates whether the timer is currently valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/isvalid
func (t_ Timer) IsValid() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isValid"))
	return rv
}


// A Boolean value that indicates whether the timer is currently valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/isvalid
func (t_ Timer) SetIsValid(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsValid:"), value)
}


