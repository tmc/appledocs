// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterOvenMode */


/* debug [class_header]: Header for MTRClusterOvenMode */
// The class instance for the [MTRClusterOvenMode] class.
var (
	MTRClusterOvenModeClass     _MTRClusterOvenModeClass
	MTRClusterOvenModeClassOnce sync.Once
)

func getMTRClusterOvenModeClass() _MTRClusterOvenModeClass {
	MTRClusterOvenModeClassOnce.Do(func() {
		MTRClusterOvenModeClass = _MTRClusterOvenModeClass{objc.GetClass("MTRClusterOvenMode")}
	})
	return MTRClusterOvenModeClass
}

type _MTRClusterOvenModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterOvenMode */
// An interface definition for the [MTRClusterOvenMode] class.
type IMTRClusterOvenMode interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterOvenMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterOvenMode */
	// methods:
	ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterOvenMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOvenModeClass) Alloc() MTRClusterOvenMode {
	rv := objc.Send[MTRClusterOvenMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterOvenModeClass) New() MTRClusterOvenMode {
	rv := objc.Send[MTRClusterOvenMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOvenMode) Init() MTRClusterOvenMode {
	rv := objc.Send[MTRClusterOvenMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOvenMode) Autorelease() MTRClusterOvenMode {
	rv := objc.Send[MTRClusterOvenMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOvenMode creates a new MTRClusterOvenMode instance.
func NewMTRClusterOvenMode() MTRClusterOvenMode {
	return getMTRClusterOvenModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterOvenMode */
// Cluster Oven Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Oven Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode
type MTRClusterOvenMode struct {
	MTRGenericCluster
}

// MTRClusterOvenModeFrom constructs a [MTRClusterOvenMode] from an unsafe.Pointer.
//
// Cluster Oven Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterOvenModeFrom(ptr unsafe.Pointer) MTRClusterOvenMode {
	return MTRClusterOvenMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterOvenMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterOvenMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterOvenMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterOvenMode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/readAttributeSupportedModes(with:)
func (m_ MTRClusterOvenMode) ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedModesWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeSupportedModesWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterOvenMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterOvenMode */



