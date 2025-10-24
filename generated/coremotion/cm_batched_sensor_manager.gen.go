// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMBatchedSensorManager */


/* debug [class_header]: Header for CMBatchedSensorManager */
// The class instance for the [BatchedSensorManager] class.
var (
	BatchedSensorManagerClass     _BatchedSensorManagerClass
	BatchedSensorManagerClassOnce sync.Once
)

func getBatchedSensorManagerClass() _BatchedSensorManagerClass {
	BatchedSensorManagerClassOnce.Do(func() {
		BatchedSensorManagerClass = _BatchedSensorManagerClass{objc.GetClass("CMBatchedSensorManager")}
	})
	return BatchedSensorManagerClass
}

type _BatchedSensorManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BatchedSensorManager */
// An interface definition for the [BatchedSensorManager] class.
type IBatchedSensorManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BatchedSensorManager */
	// properties:
	IsAccelerometerActive() bool
	SetIsAccelerometerActive(value bool)
	IsDeviceMotionActive() bool
	SetIsDeviceMotionActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BatchedSensorManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BatchedSensorManager */
// Alloc allocates a new instance without initialization.
func (bc _BatchedSensorManagerClass) Alloc() BatchedSensorManager {
	rv := objc.Send[BatchedSensorManager](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BatchedSensorManagerClass) New() BatchedSensorManager {
	rv := objc.Send[BatchedSensorManager](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BatchedSensorManager) Init() BatchedSensorManager {
	rv := objc.Send[BatchedSensorManager](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BatchedSensorManager) Autorelease() BatchedSensorManager {
	rv := objc.Send[BatchedSensorManager](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBatchedSensorManager creates a new BatchedSensorManager instance.
func NewBatchedSensorManager() BatchedSensorManager {
	return getBatchedSensorManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BatchedSensorManager */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager
type BatchedSensorManager struct {
	objectivec.Object
}

// BatchedSensorManagerFrom constructs a [BatchedSensorManager] from an unsafe.Pointer.
func BatchedSensorManagerFrom(ptr unsafe.Pointer) BatchedSensorManager {
	return BatchedSensorManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BatchedSensorManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BatchedSensorManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BatchedSensorManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/authorizationStatus
func (bc _BatchedSensorManagerClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(bc.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_properties_class/property]: authorizationStatus */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isAccelerometerSupported
func (bc _BatchedSensorManagerClass) AccelerometerSupported() bool {
	rv := objc.Send[bool](objc.ID(bc.class), objc.Sel("accelerometerSupported"))
	return rv
}/* debug [class_properties_class/property]: accelerometerSupported */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isDeviceMotionSupported
func (bc _BatchedSensorManagerClass) DeviceMotionSupported() bool {
	rv := objc.Send[bool](objc.ID(bc.class), objc.Sel("deviceMotionSupported"))
	return rv
}/* debug [class_properties_class/property]: deviceMotionSupported */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BatchedSensorManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BatchedSensorManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmbatchedsensormanager/isaccelerometeractive
func (b_ BatchedSensorManager) IsAccelerometerActive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isAccelerometerActive"))
	return rv
}/* debug [instance_properties/getter]: isAccelerometerActive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmbatchedsensormanager/isaccelerometeractive
func (b_ BatchedSensorManager) SetIsAccelerometerActive(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsAccelerometerActive:"), value)
}/* debug [instance_properties/setter]: isAccelerometerActive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmbatchedsensormanager/isdevicemotionactive
func (b_ BatchedSensorManager) IsDeviceMotionActive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isDeviceMotionActive"))
	return rv
}/* debug [instance_properties/getter]: isDeviceMotionActive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmbatchedsensormanager/isdevicemotionactive
func (b_ BatchedSensorManager) SetIsDeviceMotionActive(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsDeviceMotionActive:"), value)
}/* debug [instance_properties/setter]: isDeviceMotionActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMBatchedSensorManager */


