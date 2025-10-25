// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTimer */


/* debug [class_header]: Header for NSTimer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Timer */
// An interface definition for the [Timer] class.
type ITimer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Timer */
	// properties:
	FireDate() IDate
	SetFireDate(value IDate)
	Valid() bool
	TimeInterval() float64
	Tolerance() float64
	SetTolerance(value float64)
	UserInfo() objc.ID
	IsValid() bool
	SetIsValid(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Timer */
	// methods:
	Fire()
	Invalidate()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Timer */
// Alloc allocates a new instance without initialization.
func (tc _TimerClass) Alloc() Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Timer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Timer */

// Initializes a timer for the specified date and time interval with the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(fire:interval:repeats:block:)
func NewTimerWithFireDateIntervalRepeatsBlock(date IDate, interval float64, repeats bool, block unsafe.Pointer) Timer {
	instance := getTimerClass().Alloc()
	rv := objc.Send[Timer](instance.ID, objc.Sel("initWithFireDate:interval:repeats:block:"), date, interval, repeats, block)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTimerWithFireDateIntervalRepeatsBlock */


// Initializes a timer using the specified object and selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(fireAt:interval:target:selector:userInfo:repeats:)
func NewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats(date IDate, ti float64, t objc.IObject, s objc.SEL, ui objc.IObject, rep bool) Timer {
	instance := getTimerClass().Alloc()
	rv := objc.Send[Timer](instance.ID, objc.Sel("initWithFireDate:interval:target:selector:userInfo:repeats:"), date, ti, t, s, ui, rep)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats */


// Initializes a timer object with the specified invocation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:invocation:repeats:)
func NewTimerWithTimeIntervalInvocationRepeats(ti float64, invocation IInvocation, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}/* debug [class_init_methods/constructor]: NewTimerWithTimeIntervalInvocationRepeats */


// Initializes a timer object with the specified time interval and block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:repeats:block:)
func NewTimerWithTimeIntervalRepeatsBlock(interval float64, repeats bool, block unsafe.Pointer) Timer {
	rv := objc.Send[Timer](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}/* debug [class_init_methods/constructor]: NewTimerWithTimeIntervalRepeatsBlock */


// Initializes a timer object with the specified object and selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:target:selector:userInfo:repeats:)
func NewTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti float64, aTarget objc.IObject, aSelector objc.SEL, userInfo objc.IObject, yesOrNo bool) Timer {
	rv := objc.Send[Timer](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}/* debug [class_init_methods/constructor]: NewTimerWithTimeIntervalTargetSelectorUserInfoRepeats */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Timer */

// Initializes a timer object with the specified invocation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:invocation:repeats:)
func (tc _TimerClass) TimerWithTimeIntervalInvocationRepeats(ti float64, invocation IInvocation, yesOrNo bool) ITimer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("timerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimerWithTimeIntervalInvocationRepeats) */


// Initializes a timer object with the specified time interval and block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:repeats:block:)
func (tc _TimerClass) TimerWithTimeIntervalRepeatsBlock(interval float64, repeats bool, block unsafe.Pointer) ITimer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("timerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimerWithTimeIntervalRepeatsBlock) */


// Initializes a timer object with the specified object and selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:target:selector:userInfo:repeats:)
func (tc _TimerClass) TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti float64, aTarget objc.IObject, aSelector objc.SEL, userInfo objc.IObject, yesOrNo bool) ITimer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("timerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimerWithTimeIntervalTargetSelectorUserInfoRepeats) */


// Creates a new timer and schedules it on the current run loop in the default mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(timeInterval:invocation:repeats:)
func (tc _TimerClass) ScheduledTimerWithTimeIntervalInvocationRepeats(ti float64, invocation IInvocation, yesOrNo bool) ITimer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("scheduledTimerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ScheduledTimerWithTimeIntervalInvocationRepeats) */


// Creates a timer and schedules it on the current run loop in the default mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(timeInterval:target:selector:userInfo:repeats:)
func (tc _TimerClass) ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti float64, aTarget objc.IObject, aSelector objc.SEL, userInfo objc.IObject, yesOrNo bool) ITimer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("scheduledTimerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats) */


// Creates a timer and schedules it on the current run loop in the default mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(withTimeInterval:repeats:block:)
func (tc _TimerClass) ScheduledTimerWithTimeIntervalRepeatsBlock(interval float64, repeats bool, block unsafe.Pointer) ITimer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("scheduledTimerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ScheduledTimerWithTimeIntervalRepeatsBlock) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Timer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Timer */

// Causes the timer’s message to be sent to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/fire()
func (t_ Timer) Fire() {
	objc.Send[objc.ID](t_.ID, objc.Sel("fire"))
}/* debug [instance_methods/method]: Fire */


// Stops the timer from ever firing again and requests its removal from its run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/invalidate()
func (t_ Timer) Invalidate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidate"))
}/* debug [instance_methods/method]: Invalidate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Timer */

// The date at which the timer will fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/fireDate
func (t_ Timer) FireDate() IDate {
	rv := objc.Send[Date](t_.ID, objc.Sel("fireDate"))
	return rv
}/* debug [instance_properties/getter]: fireDate */


// The date at which the timer will fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/fireDate
func (t_ Timer) SetFireDate(value IDate) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFireDate:"), value)
}/* debug [instance_properties/setter]: fireDate */


// A Boolean value that indicates whether the timer is currently valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/isValid
func (t_ Timer) Valid() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("valid"))
	return rv
}/* debug [instance_properties/getter]: valid */


// The timer’s time interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/timeInterval
func (t_ Timer) TimeInterval() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("timeInterval"))
	return rv
}/* debug [instance_properties/getter]: timeInterval */


// The amount of time after the scheduled fire date that the timer may fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/tolerance
func (t_ Timer) Tolerance() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("tolerance"))
	return rv
}/* debug [instance_properties/getter]: tolerance */


// The amount of time after the scheduled fire date that the timer may fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/tolerance
func (t_ Timer) SetTolerance(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTolerance:"), value)
}/* debug [instance_properties/setter]: tolerance */


// The receiver’s object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Timer/userInfo
func (t_ Timer) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// A Boolean value that indicates whether the timer is currently valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/isvalid
func (t_ Timer) IsValid() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isValid"))
	return rv
}/* debug [instance_properties/getter]: isValid */


// A Boolean value that indicates whether the timer is currently valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/isvalid
func (t_ Timer) SetIsValid(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsValid:"), value)
}/* debug [instance_properties/setter]: isValid */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTimer */


