// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Timer] class.
var TimerClass objc.Class

func init() {
	TimerClass = objc.GetClass("NSTimer")
}

type Timer struct {
	objc.ID
}

func TimerFrom(ptr unsafe.Pointer) Timer {
	return Timer{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc Timer) Alloc() Timer {
	ret := objc.ID(TimerClass).Send(objc.RegisterName("alloc"))
	return Timer{ret}
}

// Init initializes the instance.
func (t_ Timer) Init() Timer {
	ret := t_.ID.Send(objc.RegisterName("init"))
	return Timer{ret}
}
// Initializes a timer for the specified date and time interval with the specified block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/init(fire:interval:repeats:block:)
func NewTimerWithFireDateIntervalRepeatsBlock(date unsafe.Pointer, interval TimeInterval, repeats bool, block unsafe.Pointer) Timer {
	instance := Timer{}.Alloc()
	sel := objc.RegisterName("initWithFireDate:interval:repeats:block:")
	ret := instance.ID.Send(sel, date, interval, repeats, block)
	instance = Timer{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a timer using the specified object and selector. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/init(fireAt:interval:target:selector:userInfo:repeats:)
func NewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats(date unsafe.Pointer, ti TimeInterval, t objc.ID, s objc.SEL, ui objc.ID, rep bool) Timer {
	instance := Timer{}.Alloc()
	sel := objc.RegisterName("initWithFireDate:interval:target:selector:userInfo:repeats:")
	ret := instance.ID.Send(sel, date, ti, t, s, ui, rep)
	instance = Timer{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Initializes a timer object with the specified invocation object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/init(timeInterval:invocation:repeats:)
func (tc Timer) TimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation unsafe.Pointer, yesOrNo bool) unsafe.Pointer {
	sel := objc.RegisterName("timerWithTimeInterval:invocation:repeats:")
	ret := objc.ID(TimerClass).Send(sel, ti, invocation, yesOrNo)
	return unsafe.Pointer(ret)
}
// Initializes a timer object with the specified time interval and block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/init(timeInterval:repeats:block:)
func (tc Timer) TimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("timerWithTimeInterval:repeats:block:")
	ret := objc.ID(TimerClass).Send(sel, interval, repeats, block)
	return unsafe.Pointer(ret)
}
// Initializes a timer object with the specified object and selector. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/init(timeInterval:target:selector:userInfo:repeats:)
func (tc Timer) TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.ID, aSelector objc.SEL, userInfo objc.ID, yesOrNo bool) unsafe.Pointer {
	sel := objc.RegisterName("timerWithTimeInterval:target:selector:userInfo:repeats:")
	ret := objc.ID(TimerClass).Send(sel, ti, aTarget, aSelector, userInfo, yesOrNo)
	return unsafe.Pointer(ret)
}
// Creates a new timer and schedules it on the current run loop in the default mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/scheduledTimer(timeInterval:invocation:repeats:)
func (tc Timer) ScheduledTimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation unsafe.Pointer, yesOrNo bool) unsafe.Pointer {
	sel := objc.RegisterName("scheduledTimerWithTimeInterval:invocation:repeats:")
	ret := objc.ID(TimerClass).Send(sel, ti, invocation, yesOrNo)
	return unsafe.Pointer(ret)
}
// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/scheduledTimer(timeInterval:target:selector:userInfo:repeats:)
func (tc Timer) ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.ID, aSelector objc.SEL, userInfo objc.ID, yesOrNo bool) unsafe.Pointer {
	sel := objc.RegisterName("scheduledTimerWithTimeInterval:target:selector:userInfo:repeats:")
	ret := objc.ID(TimerClass).Send(sel, ti, aTarget, aSelector, userInfo, yesOrNo)
	return unsafe.Pointer(ret)
}
// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/scheduledTimer(withTimeInterval:repeats:block:)
func (tc Timer) ScheduledTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("scheduledTimerWithTimeInterval:repeats:block:")
	ret := objc.ID(TimerClass).Send(sel, interval, repeats, block)
	return unsafe.Pointer(ret)
}
// Causes the timer’s message to be sent to its target. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/fire()
func (t_ Timer) Fire() {
	sel := objc.RegisterName("fire")
	t_.ID.Send(sel)
}
// Stops the timer from ever firing again and requests its removal from its run loop. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Timer/invalidate()
func (t_ Timer) Invalidate() {
	sel := objc.RegisterName("invalidate")
	t_.ID.Send(sel)
}

