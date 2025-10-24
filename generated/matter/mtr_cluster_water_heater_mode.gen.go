// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterWaterHeaterMode */


/* debug [class_header]: Header for MTRClusterWaterHeaterMode */
// The class instance for the [MTRClusterWaterHeaterMode] class.
var (
	MTRClusterWaterHeaterModeClass     _MTRClusterWaterHeaterModeClass
	MTRClusterWaterHeaterModeClassOnce sync.Once
)

func getMTRClusterWaterHeaterModeClass() _MTRClusterWaterHeaterModeClass {
	MTRClusterWaterHeaterModeClassOnce.Do(func() {
		MTRClusterWaterHeaterModeClass = _MTRClusterWaterHeaterModeClass{objc.GetClass("MTRClusterWaterHeaterMode")}
	})
	return MTRClusterWaterHeaterModeClass
}

type _MTRClusterWaterHeaterModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterWaterHeaterMode */
// An interface definition for the [MTRClusterWaterHeaterMode] class.
type IMTRClusterWaterHeaterMode interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterWaterHeaterMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterWaterHeaterMode */
	// methods:
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterWaterHeaterMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterWaterHeaterModeClass) Alloc() MTRClusterWaterHeaterMode {
	rv := objc.Send[MTRClusterWaterHeaterMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterWaterHeaterModeClass) New() MTRClusterWaterHeaterMode {
	rv := objc.Send[MTRClusterWaterHeaterMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterWaterHeaterMode) Init() MTRClusterWaterHeaterMode {
	rv := objc.Send[MTRClusterWaterHeaterMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterWaterHeaterMode) Autorelease() MTRClusterWaterHeaterMode {
	rv := objc.Send[MTRClusterWaterHeaterMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterWaterHeaterMode creates a new MTRClusterWaterHeaterMode instance.
func NewMTRClusterWaterHeaterMode() MTRClusterWaterHeaterMode {
	return getMTRClusterWaterHeaterModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterWaterHeaterMode */
// Cluster Water Heater Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Water Heater Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode
type MTRClusterWaterHeaterMode struct {
	MTRGenericCluster
}

// MTRClusterWaterHeaterModeFrom constructs a [MTRClusterWaterHeaterMode] from an unsafe.Pointer.
//
// Cluster Water Heater Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterWaterHeaterModeFrom(ptr unsafe.Pointer) MTRClusterWaterHeaterMode {
	return MTRClusterWaterHeaterMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterWaterHeaterMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterWaterHeaterMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterWaterHeaterMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterWaterHeaterMode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterWaterHeaterMode) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterWaterHeaterMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterWaterHeaterMode */



