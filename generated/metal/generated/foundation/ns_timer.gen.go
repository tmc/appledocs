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
	// properties:
	FireDate() IDate
	SetFireDate(value IDate)
	IsValid() bool
	SetIsValid(value bool)
	TimeInterval() float64
	SetTimeInterval(value float64)
	Tolerance() float64
	SetTolerance(value float64)
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	// methods:
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



// The date at which the timer will fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/firedate
func (t_ Timer) FireDate() IDate {
	rv := objc.Send[Date](t_.ID, objc.Sel("fireDate"))
	return rv
}


// The date at which the timer will fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/firedate
func (t_ Timer) SetFireDate(value IDate) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFireDate:"), value)
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


// The timer’s time interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/timeinterval
func (t_ Timer) TimeInterval() float64 {
	rv := objc.Send[TimeInterval](t_.ID, objc.Sel("timeInterval"))
	return rv
}


// The timer’s time interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/timeinterval
func (t_ Timer) SetTimeInterval(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTimeInterval:"), value)
}


// The amount of time after the scheduled fire date that the timer may fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/tolerance
func (t_ Timer) Tolerance() float64 {
	rv := objc.Send[TimeInterval](t_.ID, objc.Sel("tolerance"))
	return rv
}


// The amount of time after the scheduled fire date that the timer may fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/tolerance
func (t_ Timer) SetTolerance(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTolerance:"), value)
}


// The receiver’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/userinfo
func (t_ Timer) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("userInfo"))
	return rv
}


// The receiver’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/timer/userinfo
func (t_ Timer) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUserInfo:"), value)
}



