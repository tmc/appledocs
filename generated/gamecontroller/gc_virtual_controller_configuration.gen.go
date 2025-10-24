// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCVirtualControllerConfiguration */


/* debug [class_header]: Header for GCVirtualControllerConfiguration */
// The class instance for the [GCVirtualControllerConfiguration] class.
var (
	GCVirtualControllerConfigurationClass     _GCVirtualControllerConfigurationClass
	GCVirtualControllerConfigurationClassOnce sync.Once
)

func getGCVirtualControllerConfigurationClass() _GCVirtualControllerConfigurationClass {
	GCVirtualControllerConfigurationClassOnce.Do(func() {
		GCVirtualControllerConfigurationClass = _GCVirtualControllerConfigurationClass{objc.GetClass("GCVirtualControllerConfiguration")}
	})
	return GCVirtualControllerConfigurationClass
}

type _GCVirtualControllerConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCVirtualControllerConfiguration */
// An interface definition for the [GCVirtualControllerConfiguration] class.
type IGCVirtualControllerConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCVirtualControllerConfiguration */
	// properties:
	IsHidden() bool
	SetIsHidden(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCVirtualControllerConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCVirtualControllerConfiguration */
// Alloc allocates a new instance without initialization.
func (gc _GCVirtualControllerConfigurationClass) Alloc() GCVirtualControllerConfiguration {
	rv := objc.Send[GCVirtualControllerConfiguration](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCVirtualControllerConfigurationClass) New() GCVirtualControllerConfiguration {
	rv := objc.Send[GCVirtualControllerConfiguration](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCVirtualControllerConfiguration) Init() GCVirtualControllerConfiguration {
	rv := objc.Send[GCVirtualControllerConfiguration](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCVirtualControllerConfiguration) Autorelease() GCVirtualControllerConfiguration {
	rv := objc.Send[GCVirtualControllerConfiguration](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCVirtualControllerConfiguration creates a new GCVirtualControllerConfiguration instance.
func NewGCVirtualControllerConfiguration() GCVirtualControllerConfiguration {
	return getGCVirtualControllerConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCVirtualControllerConfiguration */
// The configuration of a virtual controller.
//
// You configure a virtual controller by specifying the input elements it contains. Then using the method, you can customize individual elements.


// The configuration of a virtual controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/Configuration
type GCVirtualControllerConfiguration struct {
	objectivec.Object
}

// GCVirtualControllerConfigurationFrom constructs a [GCVirtualControllerConfiguration] from an unsafe.Pointer.
//
// The configuration of a virtual controller.
func GCVirtualControllerConfigurationFrom(ptr unsafe.Pointer) GCVirtualControllerConfiguration {
	return GCVirtualControllerConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCVirtualControllerConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCVirtualControllerConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCVirtualControllerConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCVirtualControllerConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCVirtualControllerConfiguration */

// A Boolean value that indicates whether the system or the app presents the virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/configuration/ishidden
func (g_ GCVirtualControllerConfiguration) IsHidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// A Boolean value that indicates whether the system or the app presents the virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/configuration/ishidden
func (g_ GCVirtualControllerConfiguration) SetIsHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCVirtualControllerConfiguration */


