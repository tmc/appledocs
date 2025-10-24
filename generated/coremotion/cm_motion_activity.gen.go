// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MotionActivity] class.
var (
	MotionActivityClass     _MotionActivityClass
	MotionActivityClassOnce sync.Once
)

func getMotionActivityClass() _MotionActivityClass {
	MotionActivityClassOnce.Do(func() {
		MotionActivityClass = _MotionActivityClass{objc.GetClass("CMMotionActivity")}
	})
	return MotionActivityClass
}

type _MotionActivityClass struct {
	class objc.Class
}

// An interface definition for the [MotionActivity] class.
type IMotionActivity interface {
	ILogItem
	// properties:
	Automotive() bool
	Confidence() MotionActivityConfidence
	Running() bool
	StartDate() objc.IObject /* cross-framework: NSDate */
	Stationary() bool
	Unknown() bool
	Walking() bool
	// methods:
}

// The data for a single motion update event.
//
// On devices that support motion, you can use a or object to request updates when the current type of motion changes. When a change occurs, the update information is packaged into a object and sent to your app. The motion-related properties of this class aren’t mutually exclusive. In other words, it’s possible for more than one of the motion-related properties to contain the value . For example, if the user was driving in a car and the car stopped at a red light, the update event associated with that change in motion would have both the and properties set to . It’s also possible for all of the properties to be set to when the device is in motion but the movement doesn’t correlate to walking, running, cycling, or automotive travel. You don’t create instances of this class yourself. The object creates them and sends them to the handler block you registered. For more information about how to initiate the delivery of motion activity updates to your app, see .


// The data for a single motion update event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivity
type MotionActivity struct {
	LogItem
}

// MotionActivityFrom constructs a [MotionActivity] from an unsafe.Pointer.
//
// The data for a single motion update event.
func MotionActivityFrom(ptr unsafe.Pointer) MotionActivity {
	return MotionActivity{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MotionActivityClass) Alloc() MotionActivity {
	rv := objc.Send[MotionActivity](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MotionActivityClass) New() MotionActivity {
	rv := objc.Send[MotionActivity](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MotionActivity) Init() MotionActivity {
	rv := objc.Send[MotionActivity](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MotionActivity) Autorelease() MotionActivity {
	rv := objc.Send[MotionActivity](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMotionActivity creates a new MotionActivity instance.
func NewMotionActivity() MotionActivity {
	return getMotionActivityClass().New()
}



// A Boolean indicating whether the device is in an automobile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivity/automotive
func (m_ MotionActivity) Automotive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("automotive"))
	return rv
}


// The confidence in the assessment of the motion type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivity/confidence
func (m_ MotionActivity) Confidence() MotionActivityConfidence {
	rv := objc.Send[MotionActivityConfidence](m_.ID, objc.Sel("confidence"))
	return rv
}


// A Boolean indicating whether the device is on a running person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivity/running
func (m_ MotionActivity) Running() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("running"))
	return rv
}


// The time at which the change in motion occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivity/startDate
func (m_ MotionActivity) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("startDate"))
	return rv
}


// A Boolean indicating whether the device is stationary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivity/stationary
func (m_ MotionActivity) Stationary() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("stationary"))
	return rv
}


// A Boolean indicating whether the type of motion is unknown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivity/unknown
func (m_ MotionActivity) Unknown() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("unknown"))
	return rv
}


// A Boolean indicating whether the device is on a walking person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivity/walking
func (m_ MotionActivity) Walking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("walking"))
	return rv
}


