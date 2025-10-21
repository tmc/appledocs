// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Pedometer] class.
var (
	PedometerClass     _PedometerClass
	PedometerClassOnce sync.Once
)

func getPedometerClass() _PedometerClass {
	PedometerClassOnce.Do(func() {
		PedometerClass = _PedometerClass{objc.GetClass("CMPedometer")}
	})
	return PedometerClass
}

type _PedometerClass struct {
	class objc.Class
}

// An interface definition for the [Pedometer] class.
type IPedometer interface {
	objectivec.IObject
	QueryPedometerDataFromDateToDateWithHandler(start foundation.IDate, end foundation.IDate, handler unsafe.Pointer)
	StartPedometerEventUpdatesWithHandler(handler unsafe.Pointer)
	StartPedometerUpdatesFromDateWithHandler(start foundation.IDate, handler unsafe.Pointer)
	StopPedometerEventUpdates()
	StopPedometerUpdates()
}

// An object for fetching the system-generated live walking data.
//
// You use a pedometer object to retrieve step counts and other information about the distance traveled and the number of floors ascended or descended. The pedometer object manages a cache of historic data that you can query or you can ask for live updates as the data is processed. To use a pedometer object, create an instance of this class and call the appropriate methods. Use the method to retrieve data that has already been gathered. To get live updates, use the method to start the delivery of events to the handler you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer
type Pedometer struct {
	objectivec.Object
}

// PedometerFrom constructs a [Pedometer] from an unsafe.Pointer.
//
// An object for fetching the system-generated live walking data.
func PedometerFrom(ptr unsafe.Pointer) Pedometer {
	return Pedometer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PedometerClass) Alloc() Pedometer {
	rv := objc.Send[Pedometer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PedometerClass) New() Pedometer {
	rv := objc.Send[Pedometer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Pedometer) Init() Pedometer {
	rv := objc.Send[Pedometer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Pedometer) Autorelease() Pedometer {
	rv := objc.Send[Pedometer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPedometer creates a new Pedometer instance.
func NewPedometer() Pedometer {
	return getPedometerClass().New()
}


// Returns a value indicating whether the app is authorized to gather pedometer data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/authorizationStatus()
func (pc _PedometerClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(pc.class), objc.Sel("authorizationStatus"))
	return rv
}

// Returns a Boolean value indicating whether cadence information is available on the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isCadenceAvailable()
func (pc _PedometerClass) IsCadenceAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isCadenceAvailable"))
	return rv
}

// Returns a Boolean value indicating whether distance estimation is available on the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isDistanceAvailable()
func (pc _PedometerClass) IsDistanceAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isDistanceAvailable"))
	return rv
}

// Returns a Boolean value indicating whether floor counting is available on the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isFloorCountingAvailable()
func (pc _PedometerClass) IsFloorCountingAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isFloorCountingAvailable"))
	return rv
}

// Returns a Boolean value indicating whether pace information is available on the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isPaceAvailable()
func (pc _PedometerClass) IsPaceAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isPaceAvailable"))
	return rv
}

// Returns a Boolean value indicating whether pedometer events are available on the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isPedometerEventTrackingAvailable()
func (pc _PedometerClass) IsPedometerEventTrackingAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isPedometerEventTrackingAvailable"))
	return rv
}

// Returns a Boolean value indicating whether step counting is available on the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isStepCountingAvailable()
func (pc _PedometerClass) IsStepCountingAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isStepCountingAvailable"))
	return rv
}

// Retrieves the data between the specified start and end dates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/queryPedometerData(from:to:withHandler:)
func (p_ Pedometer) QueryPedometerDataFromDateToDateWithHandler(start foundation.IDate, end foundation.IDate, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("queryPedometerDataFromDate:toDate:withHandler:"), start, end, handler)
}

// Starts the delivery of pedometer events to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/startEventUpdates(handler:)
func (p_ Pedometer) StartPedometerEventUpdatesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startPedometerEventUpdatesWithHandler:"), handler)
}

// Starts the delivery of recent pedestrian-related data to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/startUpdates(from:withHandler:)
func (p_ Pedometer) StartPedometerUpdatesFromDateWithHandler(start foundation.IDate, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startPedometerUpdatesFromDate:withHandler:"), start, handler)
}

// Stops the delivery of pedometer events to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/stopEventUpdates()
func (p_ Pedometer) StopPedometerEventUpdates() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopPedometerEventUpdates"))
}

// Stops the delivery of recent pedestrian data updates to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/stopUpdates()
func (p_ Pedometer) StopPedometerUpdates() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopPedometerUpdates"))
}



