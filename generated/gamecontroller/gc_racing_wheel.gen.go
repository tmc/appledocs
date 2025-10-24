// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCRacingWheel */


/* debug [class_header]: Header for GCRacingWheel */
// The class instance for the [GCRacingWheel] class.
var (
	GCRacingWheelClass     _GCRacingWheelClass
	GCRacingWheelClassOnce sync.Once
)

func getGCRacingWheelClass() _GCRacingWheelClass {
	GCRacingWheelClassOnce.Do(func() {
		GCRacingWheelClass = _GCRacingWheelClass{objc.GetClass("GCRacingWheel")}
	})
	return GCRacingWheelClass
}

type _GCRacingWheelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCRacingWheel */
// An interface definition for the [GCRacingWheel] class.
type IGCRacingWheel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCRacingWheel */
	// properties:
	Acquired() bool
	Snapshot() bool
	WheelInput() IGCRacingWheelInput
	IsAcquired() bool
	SetIsAcquired(value bool)
	IsSnapshot() bool
	SetIsSnapshot(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCRacingWheel */
	// methods:
	AcquireDeviceWithError(error_ unsafe.Pointer) bool
	Capture() IGCRacingWheel
	RelinquishDevice()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCRacingWheel */
// Alloc allocates a new instance without initialization.
func (gc _GCRacingWheelClass) Alloc() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCRacingWheelClass) New() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCRacingWheel) Init() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCRacingWheel) Autorelease() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCRacingWheel creates a new GCRacingWheel instance.
func NewGCRacingWheel() GCRacingWheel {
	return getGCRacingWheelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCRacingWheel */
// An object that represents a physical racing wheel controller connected to a device.


// An object that represents a physical racing wheel controller connected to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel
type GCRacingWheel struct {
	objectivec.Object
}

// GCRacingWheelFrom constructs a [GCRacingWheel] from an unsafe.Pointer.
//
// An object that represents a physical racing wheel controller connected to a device.
func GCRacingWheelFrom(ptr unsafe.Pointer) GCRacingWheel {
	return GCRacingWheel{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCRacingWheel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCRacingWheel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCRacingWheel */

// The racing wheels connected to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/connectedRacingWheels
func (gc _GCRacingWheelClass) ConnectedRacingWheels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("connectedRacingWheels"))
	return rv
}/* debug [class_properties_class/property]: connectedRacingWheels */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCRacingWheel */

// Starts receiving events from the racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/acquireDevice()
func (g_ GCRacingWheel) AcquireDeviceWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("acquireDeviceWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: AcquireDeviceWithError */


// Returns a snapshot of the racing wheel with its current element values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/capture()
func (g_ GCRacingWheel) Capture() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](g_.ID, objc.Sel("capture"))
	return rv
}/* debug [instance_methods/method]: Capture */


// Stops receiving events from the racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/relinquishDevice()
func (g_ GCRacingWheel) RelinquishDevice() {
	objc.Send[objc.ID](g_.ID, objc.Sel("relinquishDevice"))
}/* debug [instance_methods/method]: RelinquishDevice */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCRacingWheel */

// The racing wheels connected to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/connectedRacingWheels
func (g_ GCRacingWheel) ConnectedRacingWheels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("connectedRacingWheels"))
	return rv
}/* debug [instance_properties/getter]: connectedRacingWheels */


// A Boolean value that indicates whether the racing wheel sends events to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/isAcquired
func (g_ GCRacingWheel) Acquired() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("acquired"))
	return rv
}/* debug [instance_properties/getter]: acquired */


// A Boolean value that indicates whether the object is a snapshot of a racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/isSnapshot
func (g_ GCRacingWheel) Snapshot() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("snapshot"))
	return rv
}/* debug [instance_properties/getter]: snapshot */


// The physical input profile for the racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/wheelInput
func (g_ GCRacingWheel) WheelInput() IGCRacingWheelInput {
	rv := objc.Send[GCRacingWheelInput](g_.ID, objc.Sel("wheelInput"))
	return rv
}/* debug [instance_properties/getter]: wheelInput */


// A Boolean value that indicates whether the racing wheel sends events to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/isacquired
func (g_ GCRacingWheel) IsAcquired() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isAcquired"))
	return rv
}/* debug [instance_properties/getter]: isAcquired */


// A Boolean value that indicates whether the racing wheel sends events to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/isacquired
func (g_ GCRacingWheel) SetIsAcquired(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsAcquired:"), value)
}/* debug [instance_properties/setter]: isAcquired */


// A Boolean value that indicates whether the object is a snapshot of a racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/issnapshot
func (g_ GCRacingWheel) IsSnapshot() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isSnapshot"))
	return rv
}/* debug [instance_properties/getter]: isSnapshot */


// A Boolean value that indicates whether the object is a snapshot of a racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/issnapshot
func (g_ GCRacingWheel) SetIsSnapshot(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsSnapshot:"), value)
}/* debug [instance_properties/setter]: isSnapshot */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCRacingWheel */



