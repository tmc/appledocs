// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Altimeter] class.
var (
	AltimeterClass     _AltimeterClass
	AltimeterClassOnce sync.Once
)

func getAltimeterClass() _AltimeterClass {
	AltimeterClassOnce.Do(func() {
		AltimeterClass = _AltimeterClass{objc.GetClass("CMAltimeter")}
	})
	return AltimeterClass
}

type _AltimeterClass struct {
	class objc.Class
}

// An interface definition for the [Altimeter] class.
type IAltimeter interface {
	objectivec.IObject
	StartAbsoluteAltitudeUpdatesToQueueWithHandler(queue foundation.OperationQueue, handler unsafe.Pointer)
	StartRelativeAltitudeUpdatesToQueueWithHandler(queue foundation.OperationQueue, handler unsafe.Pointer)
	StopAbsoluteAltitudeUpdates()
	StopRelativeAltitudeUpdates()
}

// An object that initiates the delivery of altitude-related changes.
//
// Altitude events report changes in both the relative and absolute altitude. For example, a hiking app could use this object to track the user’s elevation change over the course of a hike, or to report their current absolute altitude during the hike. Because altitude events may not be available on all devices, always call the method before starting relative altitude updates, and call before starting absolute altitude updates. After checking the availability of altitude data, call the method to start receiving relative altitude data, or call the method for absolute altitude data. Core Motion generates events at regular intervals (regardless of whether the data has changed) and delivers them to the block you specified. When you no longer need the event data, call the or methods respectively.


// An object that initiates the delivery of altitude-related changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter
type Altimeter struct {
	objectivec.Object
}

// AltimeterFrom constructs a [Altimeter] from an unsafe.Pointer.
//
// An object that initiates the delivery of altitude-related changes.
func AltimeterFrom(ptr unsafe.Pointer) Altimeter {
	return Altimeter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AltimeterClass) Alloc() Altimeter {
	rv := objc.Send[Altimeter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AltimeterClass) New() Altimeter {
	rv := objc.Send[Altimeter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Altimeter) Init() Altimeter {
	rv := objc.Send[Altimeter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Altimeter) Autorelease() Altimeter {
	rv := objc.Send[Altimeter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAltimeter creates a new Altimeter instance.
func NewAltimeter() Altimeter {
	return getAltimeterClass().New()
}



// Returns a value indicating whether the app is authorized to retrieve altimeter data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/authorizationStatus()
func (ac _AltimeterClass) AuthorizationStatus() CMAuthorizationStatus {
	rv := objc.Send[CMAuthorizationStatus](objc.ID(ac.class), objc.Sel("authorizationStatus"))
	return rv
}


// Returns a Boolean value indicating whether the current device reports changes in the absolute altitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/isAbsoluteAltitudeAvailable()
func (ac _AltimeterClass) IsAbsoluteAltitudeAvailable() bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("isAbsoluteAltitudeAvailable"))
	return rv
}


// Returns a Boolean value indicating whether the current device supports generating data for relative altitude changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/isRelativeAltitudeAvailable()
func (ac _AltimeterClass) IsRelativeAltitudeAvailable() bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("isRelativeAltitudeAvailable"))
	return rv
}


// Starts the delivery of absolute altitude data to the specified handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/startAbsoluteAltitudeUpdates(to:withHandler:)
func (a_ Altimeter) StartAbsoluteAltitudeUpdatesToQueueWithHandler(queue foundation.OperationQueue, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("startAbsoluteAltitudeUpdatesToQueue:withHandler:"), queue, handler)
}


// Starts the delivery of relative altitude data to the specified handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/startRelativeAltitudeUpdates(to:withHandler:)
func (a_ Altimeter) StartRelativeAltitudeUpdatesToQueueWithHandler(queue foundation.OperationQueue, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("startRelativeAltitudeUpdatesToQueue:withHandler:"), queue, handler)
}


// Stops the delivery of absolute altitude data for this altimeter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/stopAbsoluteAltitudeUpdates()
func (a_ Altimeter) StopAbsoluteAltitudeUpdates() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopAbsoluteAltitudeUpdates"))
}


// Stops the delivery of relative altitude data for the altimeter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/stopRelativeAltitudeUpdates()
func (a_ Altimeter) StopRelativeAltitudeUpdates() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopRelativeAltitudeUpdates"))
}



