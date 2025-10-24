// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMMovementDisorderManager */


/* debug [class_header]: Header for CMMovementDisorderManager */
// The class instance for the [MovementDisorderManager] class.
var (
	MovementDisorderManagerClass     _MovementDisorderManagerClass
	MovementDisorderManagerClassOnce sync.Once
)

func getMovementDisorderManagerClass() _MovementDisorderManagerClass {
	MovementDisorderManagerClassOnce.Do(func() {
		MovementDisorderManagerClass = _MovementDisorderManagerClass{objc.GetClass("CMMovementDisorderManager")}
	})
	return MovementDisorderManagerClass
}

type _MovementDisorderManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MovementDisorderManager */
// An interface definition for the [MovementDisorderManager] class.
type IMovementDisorderManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MovementDisorderManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MovementDisorderManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MovementDisorderManager */
// Alloc allocates a new instance without initialization.
func (mc _MovementDisorderManagerClass) Alloc() MovementDisorderManager {
	rv := objc.Send[MovementDisorderManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MovementDisorderManagerClass) New() MovementDisorderManager {
	rv := objc.Send[MovementDisorderManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovementDisorderManager) Init() MovementDisorderManager {
	rv := objc.Send[MovementDisorderManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovementDisorderManager) Autorelease() MovementDisorderManager {
	rv := objc.Send[MovementDisorderManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovementDisorderManager creates a new MovementDisorderManager instance.
func NewMovementDisorderManager() MovementDisorderManager {
	return getMovementDisorderManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MovementDisorderManager */
// A manager for recording and querying movement disorder data.
//
// Use to measure a resting Parkinsonian tremor in the 3-7 Hz range and choreiform dyskinetic symptoms. When collecting data, the user should wear Apple Watch on their most affected arm. requires an entitlement from Apple. To apply for the entitlement, see .


// A manager for recording and querying movement disorder data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager
type MovementDisorderManager struct {
	objectivec.Object
}

// MovementDisorderManagerFrom constructs a [MovementDisorderManager] from an unsafe.Pointer.
//
// A manager for recording and querying movement disorder data.
func MovementDisorderManagerFrom(ptr unsafe.Pointer) MovementDisorderManager {
	return MovementDisorderManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MovementDisorderManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MovementDisorderManager */

// A value indicating whether the user has authorized the app to monitor and query for movement disorder data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/authorizationStatus()
func (mc _MovementDisorderManagerClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(mc.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationStatus) */


// A Boolean value indicating whether the current device supports the movement disorder manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/isAvailable()
func (mc _MovementDisorderManagerClass) IsAvailable() bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("isAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsAvailable) */


// Returns a string that describes the movement disorder algorithm’s current version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/version()
func (mc _MovementDisorderManagerClass) Version() foundation.String {
	rv := objc.Send[foundation.String](objc.ID(mc.class), objc.Sel("version"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Version) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MovementDisorderManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MovementDisorderManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MovementDisorderManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMMovementDisorderManager */


