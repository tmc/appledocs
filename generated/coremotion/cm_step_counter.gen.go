// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StepCounter] class.
var (
	StepCounterClass     _StepCounterClass
	StepCounterClassOnce sync.Once
)

func getStepCounterClass() _StepCounterClass {
	StepCounterClassOnce.Do(func() {
		StepCounterClass = _StepCounterClass{objc.GetClass("CMStepCounter")}
	})
	return StepCounterClass
}

type _StepCounterClass struct {
	class objc.Class
}

// An interface definition for the [StepCounter] class.
type IStepCounter interface {
	objectivec.IObject
	QueryStepCountStartingFromToToQueueWithHandler(start foundation.IDate, end foundation.IDate, queue foundation.IOperationQueue, handler unsafe.Pointer)
	StartStepCountingUpdatesToQueueUpdateOnWithHandler(queue foundation.IOperationQueue, stepCounts int, handler unsafe.Pointer)
	StopStepCountingUpdates()
}

// The number of steps the user has taken with the device.
//
// Step information is gathered on devices with the appropriate built-in hardware and stored so that you can run queries to determine the user’s recent physical activity. You use this class to gather both current step data and any historical data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMStepCounter
type StepCounter struct {
	objectivec.Object
}

// StepCounterFrom constructs a [StepCounter] from an unsafe.Pointer.
//
// The number of steps the user has taken with the device.
func StepCounterFrom(ptr unsafe.Pointer) StepCounter {
	return StepCounter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StepCounterClass) Alloc() StepCounter {
	rv := objc.Send[StepCounter](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StepCounterClass) New() StepCounter {
	rv := objc.Send[StepCounter](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StepCounter) Init() StepCounter {
	rv := objc.Send[StepCounter](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StepCounter) Autorelease() StepCounter {
	rv := objc.Send[StepCounter](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStepCounter creates a new StepCounter instance.
func NewStepCounter() StepCounter {
	return getStepCounterClass().New()
}


// Returns a Boolean indicating whether step-counting support is available on the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMStepCounter/isStepCountingAvailable()
func (sc _StepCounterClass) IsStepCountingAvailable() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("isStepCountingAvailable"))
	return rv
}

// Gathers and returns historical step count data for the specified time period.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMStepCounter/queryStepCountStarting(from:to:to:withHandler:)
func (s_ StepCounter) QueryStepCountStartingFromToToQueueWithHandler(start foundation.IDate, end foundation.IDate, queue foundation.IOperationQueue, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("queryStepCountStartingFrom:to:toQueue:withHandler:"), start, end, queue, handler)
}

// Starts the delivery of current step-counting data to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMStepCounter/startStepCountingUpdates(to:updateOn:withHandler:)
func (s_ StepCounter) StartStepCountingUpdatesToQueueUpdateOnWithHandler(queue foundation.IOperationQueue, stepCounts int, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("startStepCountingUpdatesToQueue:updateOn:withHandler:"), queue, stepCounts, handler)
}

// Stops the delivery of step-counting updates to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMStepCounter/stopStepCountingUpdates()
func (s_ StepCounter) StopStepCountingUpdates() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stopStepCountingUpdates"))
}



