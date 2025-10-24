// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSHapticFeedbackManager */


/* debug [class_header]: Header for NSHapticFeedbackManager */
// The class instance for the [HapticFeedbackManager] class.
var (
	HapticFeedbackManagerClass     _HapticFeedbackManagerClass
	HapticFeedbackManagerClassOnce sync.Once
)

func getHapticFeedbackManagerClass() _HapticFeedbackManagerClass {
	HapticFeedbackManagerClassOnce.Do(func() {
		HapticFeedbackManagerClass = _HapticFeedbackManagerClass{objc.GetClass("NSHapticFeedbackManager")}
	})
	return HapticFeedbackManagerClass
}

type _HapticFeedbackManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HapticFeedbackManager */
// An interface definition for the [HapticFeedbackManager] class.
type IHapticFeedbackManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HapticFeedbackManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HapticFeedbackManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HapticFeedbackManager */
// Alloc allocates a new instance without initialization.
func (hc _HapticFeedbackManagerClass) Alloc() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HapticFeedbackManagerClass) New() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HapticFeedbackManager) Init() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HapticFeedbackManager) Autorelease() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHapticFeedbackManager creates a new HapticFeedbackManager instance.
func NewHapticFeedbackManager() HapticFeedbackManager {
	return getHapticFeedbackManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HapticFeedbackManager */
// An object that provides access to the haptic feedback management attributes on a system with a Force Touch trackpad.


// An object that provides access to the haptic feedback management attributes on a system with a Force Touch trackpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager
type HapticFeedbackManager struct {
	objectivec.Object
}

// HapticFeedbackManagerFrom constructs a [HapticFeedbackManager] from an unsafe.Pointer.
//
// An object that provides access to the haptic feedback management attributes on a system with a Force Touch trackpad.
func HapticFeedbackManagerFrom(ptr unsafe.Pointer) HapticFeedbackManager {
	return HapticFeedbackManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HapticFeedbackManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HapticFeedbackManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HapticFeedbackManager */

// Requests a haptic feedback performer object that is based on the current input device, accessibility settings, and user preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/defaultPerformer
func (hc _HapticFeedbackManagerClass) DefaultPerformer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("defaultPerformer"))
	return rv
}/* debug [class_properties_class/property]: defaultPerformer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HapticFeedbackManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HapticFeedbackManager */

// Requests a haptic feedback performer object that is based on the current input device, accessibility settings, and user preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/defaultPerformer
func (h_ HapticFeedbackManager) DefaultPerformer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("defaultPerformer"))
	return rv
}/* debug [instance_properties/getter]: defaultPerformer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSHapticFeedbackManager */



