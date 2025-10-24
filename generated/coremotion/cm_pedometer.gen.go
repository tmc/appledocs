// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMPedometer */


/* debug [class_header]: Header for CMPedometer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Pedometer */
// An interface definition for the [Pedometer] class.
type IPedometer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Pedometer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Pedometer */
	// methods:
	QueryPedometerDataFromDateToDateWithHandler(start objc.IObject /* cross-framework: NSDate */, end objc.IObject /* cross-framework: NSDate */, handler PedometerHandler /* not a class type */)
	StartPedometerUpdatesFromDateWithHandler(start objc.IObject /* cross-framework: NSDate */, handler PedometerHandler /* not a class type */)
	StopPedometerUpdates()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Pedometer */
// Alloc allocates a new instance without initialization.
func (pc _PedometerClass) Alloc() Pedometer {
	rv := objc.Send[Pedometer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Pedometer */
// An object for fetching the system-generated live walking data.
//
// You use a pedometer object to retrieve step counts and other information about the distance traveled and the number of floors ascended or descended. The pedometer object manages a cache of historic data that you can query or you can ask for live updates as the data is processed. To use a pedometer object, create an instance of this class and call the appropriate methods. Use the method to retrieve data that has already been gathered. To get live updates, use the method to start the delivery of events to the handler you provide.


// An object for fetching the system-generated live walking data.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Pedometer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Pedometer */

// Returns a value indicating whether the app is authorized to gather pedometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/authorizationStatus()
func (pc _PedometerClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(pc.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationStatus) */


// Returns a Boolean value indicating whether cadence information is available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isCadenceAvailable()
func (pc _PedometerClass) IsCadenceAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isCadenceAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsCadenceAvailable) */


// Returns a Boolean value indicating whether distance estimation is available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isDistanceAvailable()
func (pc _PedometerClass) IsDistanceAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isDistanceAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsDistanceAvailable) */


// Returns a Boolean value indicating whether floor counting is available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isFloorCountingAvailable()
func (pc _PedometerClass) IsFloorCountingAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isFloorCountingAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsFloorCountingAvailable) */


// Returns a Boolean value indicating whether pace information is available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isPaceAvailable()
func (pc _PedometerClass) IsPaceAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isPaceAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsPaceAvailable) */


// Returns a Boolean value indicating whether pedometer events are available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isPedometerEventTrackingAvailable()
func (pc _PedometerClass) IsPedometerEventTrackingAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isPedometerEventTrackingAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsPedometerEventTrackingAvailable) */


// Returns a Boolean value indicating whether step counting is available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/isStepCountingAvailable()
func (pc _PedometerClass) IsStepCountingAvailable() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isStepCountingAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsStepCountingAvailable) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Pedometer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Pedometer */

// Retrieves the data between the specified start and end dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/queryPedometerData(from:to:withHandler:)
func (p_ Pedometer) QueryPedometerDataFromDateToDateWithHandler(start objc.IObject /* cross-framework: NSDate */, end objc.IObject /* cross-framework: NSDate */, handler PedometerHandler /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("queryPedometerDataFromDate:toDate:withHandler:"), start, end, handler)
}/* debug [instance_methods/method]: QueryPedometerDataFromDateToDateWithHandler */


// Starts the delivery of recent pedestrian-related data to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/startUpdates(from:withHandler:)
func (p_ Pedometer) StartPedometerUpdatesFromDateWithHandler(start objc.IObject /* cross-framework: NSDate */, handler PedometerHandler /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startPedometerUpdatesFromDate:withHandler:"), start, handler)
}/* debug [instance_methods/method]: StartPedometerUpdatesFromDateWithHandler */


// Stops the delivery of recent pedestrian data updates to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/stopUpdates()
func (p_ Pedometer) StopPedometerUpdates() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopPedometerUpdates"))
}/* debug [instance_methods/method]: StopPedometerUpdates */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Pedometer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMPedometer */


