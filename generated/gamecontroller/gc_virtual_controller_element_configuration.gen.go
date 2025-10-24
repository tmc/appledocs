// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCVirtualControllerElementConfiguration */


/* debug [class_header]: Header for GCVirtualControllerElementConfiguration */
// The class instance for the [GCVirtualControllerElementConfiguration] class.
var (
	GCVirtualControllerElementConfigurationClass     _GCVirtualControllerElementConfigurationClass
	GCVirtualControllerElementConfigurationClassOnce sync.Once
)

func getGCVirtualControllerElementConfigurationClass() _GCVirtualControllerElementConfigurationClass {
	GCVirtualControllerElementConfigurationClassOnce.Do(func() {
		GCVirtualControllerElementConfigurationClass = _GCVirtualControllerElementConfigurationClass{objc.GetClass("GCVirtualControllerElementConfiguration")}
	})
	return GCVirtualControllerElementConfigurationClass
}

type _GCVirtualControllerElementConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCVirtualControllerElementConfiguration */
// An interface definition for the [GCVirtualControllerElementConfiguration] class.
type IGCVirtualControllerElementConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCVirtualControllerElementConfiguration */
	// properties:
	IsHidden() bool
	SetIsHidden(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCVirtualControllerElementConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCVirtualControllerElementConfiguration */
// Alloc allocates a new instance without initialization.
func (gc _GCVirtualControllerElementConfigurationClass) Alloc() GCVirtualControllerElementConfiguration {
	rv := objc.Send[GCVirtualControllerElementConfiguration](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCVirtualControllerElementConfigurationClass) New() GCVirtualControllerElementConfiguration {
	rv := objc.Send[GCVirtualControllerElementConfiguration](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCVirtualControllerElementConfiguration) Init() GCVirtualControllerElementConfiguration {
	rv := objc.Send[GCVirtualControllerElementConfiguration](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCVirtualControllerElementConfiguration) Autorelease() GCVirtualControllerElementConfiguration {
	rv := objc.Send[GCVirtualControllerElementConfiguration](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCVirtualControllerElementConfiguration creates a new GCVirtualControllerElementConfiguration instance.
func NewGCVirtualControllerElementConfiguration() GCVirtualControllerElementConfiguration {
	return getGCVirtualControllerElementConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCVirtualControllerElementConfiguration */
// The properties of a virtual controller’s element that you can customize.


// The properties of a virtual controller’s element that you can customize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/ElementConfiguration
type GCVirtualControllerElementConfiguration struct {
	objectivec.Object
}

// GCVirtualControllerElementConfigurationFrom constructs a [GCVirtualControllerElementConfiguration] from an unsafe.Pointer.
//
// The properties of a virtual controller’s element that you can customize.
func GCVirtualControllerElementConfigurationFrom(ptr unsafe.Pointer) GCVirtualControllerElementConfiguration {
	return GCVirtualControllerElementConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCVirtualControllerElementConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCVirtualControllerElementConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCVirtualControllerElementConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCVirtualControllerElementConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCVirtualControllerElementConfiguration */

// A Boolean value that determines whether the virtual controller hides the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/elementconfiguration/ishidden
func (g_ GCVirtualControllerElementConfiguration) IsHidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// A Boolean value that determines whether the virtual controller hides the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/elementconfiguration/ishidden
func (g_ GCVirtualControllerElementConfiguration) SetIsHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCVirtualControllerElementConfiguration */


