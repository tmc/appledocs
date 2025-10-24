// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMFallDetectionManager */


/* debug [class_header]: Header for CMFallDetectionManager */
// The class instance for the [FallDetectionManager] class.
var (
	FallDetectionManagerClass     _FallDetectionManagerClass
	FallDetectionManagerClassOnce sync.Once
)

func getFallDetectionManagerClass() _FallDetectionManagerClass {
	FallDetectionManagerClassOnce.Do(func() {
		FallDetectionManagerClass = _FallDetectionManagerClass{objc.GetClass("CMFallDetectionManager")}
	})
	return FallDetectionManagerClass
}

type _FallDetectionManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FallDetectionManager */
// An interface definition for the [FallDetectionManager] class.
type IFallDetectionManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FallDetectionManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FallDetectionManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FallDetectionManager */
// Alloc allocates a new instance without initialization.
func (fc _FallDetectionManagerClass) Alloc() FallDetectionManager {
	rv := objc.Send[FallDetectionManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FallDetectionManagerClass) New() FallDetectionManager {
	rv := objc.Send[FallDetectionManager](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FallDetectionManager) Init() FallDetectionManager {
	rv := objc.Send[FallDetectionManager](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FallDetectionManager) Autorelease() FallDetectionManager {
	rv := objc.Send[FallDetectionManager](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFallDetectionManager creates a new FallDetectionManager instance.
func NewFallDetectionManager() FallDetectionManager {
	return getFallDetectionManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FallDetectionManager */
// An object for managing fall detection events.
//
// In Series 4 and later, Apple Watch can detect when a wearer falls, and contact emergency services if necessary. Using the , your app can request the user’s authorization, and set up a delegate to receive notifications about these . For more information, see . requires an entitlement from Apple. To apply for the entitlement, see . This entitlement allows the app to run in the background without requiring any additional capabilities. However, you can add capabilities for other background modes, as needed by your app. There are two approaches to detecting falls in your app. You can either query for samples in HealthKit, or you can use Core Motion’s .


// An object for managing fall detection events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager
type FallDetectionManager struct {
	objectivec.Object
}

// FallDetectionManagerFrom constructs a [FallDetectionManager] from an unsafe.Pointer.
//
// An object for managing fall detection events.
func FallDetectionManagerFrom(ptr unsafe.Pointer) FallDetectionManager {
	return FallDetectionManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FallDetectionManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FallDetectionManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FallDetectionManager */

// A Boolean value that indicates whether the current device supports fall detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/isAvailable
func (fc _FallDetectionManagerClass) Available() bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("available"))
	return rv
}/* debug [class_properties_class/property]: available */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FallDetectionManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FallDetectionManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMFallDetectionManager */


