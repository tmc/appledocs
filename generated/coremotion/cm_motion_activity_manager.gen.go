// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMMotionActivityManager */


/* debug [class_header]: Header for CMMotionActivityManager */
// The class instance for the [MotionActivityManager] class.
var (
	MotionActivityManagerClass     _MotionActivityManagerClass
	MotionActivityManagerClassOnce sync.Once
)

func getMotionActivityManagerClass() _MotionActivityManagerClass {
	MotionActivityManagerClassOnce.Do(func() {
		MotionActivityManagerClass = _MotionActivityManagerClass{objc.GetClass("CMMotionActivityManager")}
	})
	return MotionActivityManagerClass
}

type _MotionActivityManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MotionActivityManager */
// An interface definition for the [MotionActivityManager] class.
type IMotionActivityManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MotionActivityManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MotionActivityManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MotionActivityManager */
// Alloc allocates a new instance without initialization.
func (mc _MotionActivityManagerClass) Alloc() MotionActivityManager {
	rv := objc.Send[MotionActivityManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MotionActivityManagerClass) New() MotionActivityManager {
	rv := objc.Send[MotionActivityManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MotionActivityManager) Init() MotionActivityManager {
	rv := objc.Send[MotionActivityManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MotionActivityManager) Autorelease() MotionActivityManager {
	rv := objc.Send[MotionActivityManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMotionActivityManager creates a new MotionActivityManager instance.
func NewMotionActivityManager() MotionActivityManager {
	return getMotionActivityManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MotionActivityManager */
// An object that manages access to the motion data stored by the device.
//
// Motion data reflects whether the user is walking, running, in a vehicle, or stationary for periods of time. Using this class, you can ask for notifications when the current type of motion changes or you can gather past motion change data. For example, a navigation app might look for changes in the current type of motion and offer different directions for each.


// An object that manages access to the motion data stored by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager
type MotionActivityManager struct {
	objectivec.Object
}

// MotionActivityManagerFrom constructs a [MotionActivityManager] from an unsafe.Pointer.
//
// An object that manages access to the motion data stored by the device.
func MotionActivityManagerFrom(ptr unsafe.Pointer) MotionActivityManager {
	return MotionActivityManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MotionActivityManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MotionActivityManager */

// Returns a value indicating whether the app is authorized to retrieve stored motion data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/authorizationStatus()
func (mc _MotionActivityManagerClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(mc.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationStatus) */


// Returns a Boolean indicating whether motion data is available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/isActivityAvailable()
func (mc _MotionActivityManagerClass) IsActivityAvailable() bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("isActivityAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsActivityAvailable) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MotionActivityManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MotionActivityManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MotionActivityManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMMotionActivityManager */


