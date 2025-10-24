// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXFeatureOverrideSessionManager */


/* debug [class_header]: Header for AXFeatureOverrideSessionManager */
// The class instance for the [AXFeatureOverrideSessionManager] class.
var (
	AXFeatureOverrideSessionManagerClass     _AXFeatureOverrideSessionManagerClass
	AXFeatureOverrideSessionManagerClassOnce sync.Once
)

func getAXFeatureOverrideSessionManagerClass() _AXFeatureOverrideSessionManagerClass {
	AXFeatureOverrideSessionManagerClassOnce.Do(func() {
		AXFeatureOverrideSessionManagerClass = _AXFeatureOverrideSessionManagerClass{objc.GetClass("AXFeatureOverrideSessionManager")}
	})
	return AXFeatureOverrideSessionManagerClass
}

type _AXFeatureOverrideSessionManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXFeatureOverrideSessionManager */
// An interface definition for the [AXFeatureOverrideSessionManager] class.
type IAXFeatureOverrideSessionManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXFeatureOverrideSessionManager */
	// properties:
	AXFeatureOverrideSessionErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXFeatureOverrideSessionManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXFeatureOverrideSessionManager */
// Alloc allocates a new instance without initialization.
func (ac _AXFeatureOverrideSessionManagerClass) Alloc() AXFeatureOverrideSessionManager {
	rv := objc.Send[AXFeatureOverrideSessionManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXFeatureOverrideSessionManagerClass) New() AXFeatureOverrideSessionManager {
	rv := objc.Send[AXFeatureOverrideSessionManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXFeatureOverrideSessionManager) Init() AXFeatureOverrideSessionManager {
	rv := objc.Send[AXFeatureOverrideSessionManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXFeatureOverrideSessionManager) Autorelease() AXFeatureOverrideSessionManager {
	rv := objc.Send[AXFeatureOverrideSessionManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXFeatureOverrideSessionManager creates a new AXFeatureOverrideSessionManager instance.
func NewAXFeatureOverrideSessionManager() AXFeatureOverrideSessionManager {
	return getAXFeatureOverrideSessionManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXFeatureOverrideSessionManager */
// A manager class to begin and end accessibility feature override sessions. Multiple override sessions are reconciled by combining the requests, preferring feature enablement. Ending all sessions restores the prior state of Accessibility feature enablement. Your app must be entitled with com.apple.developer.accessibility.merchant-api-control.


// A manager class to begin and end accessibility feature override sessions. Multiple override sessions are reconciled by combining the requests, preferring feature enablement. Ending all sessions restores the prior state of Accessibility feature enablement. Your app must be entitled with com.apple.developer.accessibility.merchant-api-control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionManager
type AXFeatureOverrideSessionManager struct {
	objectivec.Object
}

// AXFeatureOverrideSessionManagerFrom constructs a [AXFeatureOverrideSessionManager] from an unsafe.Pointer.
//
// A manager class to begin and end accessibility feature override sessions. Multiple override sessions are reconciled by combining the requests, preferring feature enablement. Ending all sessions restores the prior state of Accessibility feature enablement. Your app must be entitled with com.apple.developer.accessibility.merchant-api-control.
func AXFeatureOverrideSessionManagerFrom(ptr unsafe.Pointer) AXFeatureOverrideSessionManager {
	return AXFeatureOverrideSessionManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXFeatureOverrideSessionManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXFeatureOverrideSessionManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXFeatureOverrideSessionManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionManager/sharedInstance
func (ac _AXFeatureOverrideSessionManagerClass) SharedInstance() AXFeatureOverrideSessionManager {
	rv := objc.Send[AXFeatureOverrideSessionManager](objc.ID(ac.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_properties_class/property]: sharedInstance */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXFeatureOverrideSessionManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXFeatureOverrideSessionManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axfeatureoverridesessionerrordomain
func (a_ AXFeatureOverrideSessionManager) AXFeatureOverrideSessionErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AXFeatureOverrideSessionErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: AXFeatureOverrideSessionErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXFeatureOverrideSessionManager */


