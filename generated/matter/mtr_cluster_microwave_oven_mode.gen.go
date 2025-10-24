// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterMicrowaveOvenMode */


/* debug [class_header]: Header for MTRClusterMicrowaveOvenMode */
// The class instance for the [MTRClusterMicrowaveOvenMode] class.
var (
	MTRClusterMicrowaveOvenModeClass     _MTRClusterMicrowaveOvenModeClass
	MTRClusterMicrowaveOvenModeClassOnce sync.Once
)

func getMTRClusterMicrowaveOvenModeClass() _MTRClusterMicrowaveOvenModeClass {
	MTRClusterMicrowaveOvenModeClassOnce.Do(func() {
		MTRClusterMicrowaveOvenModeClass = _MTRClusterMicrowaveOvenModeClass{objc.GetClass("MTRClusterMicrowaveOvenMode")}
	})
	return MTRClusterMicrowaveOvenModeClass
}

type _MTRClusterMicrowaveOvenModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterMicrowaveOvenMode */
// An interface definition for the [MTRClusterMicrowaveOvenMode] class.
type IMTRClusterMicrowaveOvenMode interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterMicrowaveOvenMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterMicrowaveOvenMode */
	// methods:
	ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterMicrowaveOvenMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterMicrowaveOvenModeClass) Alloc() MTRClusterMicrowaveOvenMode {
	rv := objc.Send[MTRClusterMicrowaveOvenMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterMicrowaveOvenModeClass) New() MTRClusterMicrowaveOvenMode {
	rv := objc.Send[MTRClusterMicrowaveOvenMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterMicrowaveOvenMode) Init() MTRClusterMicrowaveOvenMode {
	rv := objc.Send[MTRClusterMicrowaveOvenMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterMicrowaveOvenMode) Autorelease() MTRClusterMicrowaveOvenMode {
	rv := objc.Send[MTRClusterMicrowaveOvenMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterMicrowaveOvenMode creates a new MTRClusterMicrowaveOvenMode instance.
func NewMTRClusterMicrowaveOvenMode() MTRClusterMicrowaveOvenMode {
	return getMTRClusterMicrowaveOvenModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterMicrowaveOvenMode */
// Cluster Microwave Oven Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Microwave Oven Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode
type MTRClusterMicrowaveOvenMode struct {
	MTRGenericCluster
}

// MTRClusterMicrowaveOvenModeFrom constructs a [MTRClusterMicrowaveOvenMode] from an unsafe.Pointer.
//
// Cluster Microwave Oven Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterMicrowaveOvenModeFrom(ptr unsafe.Pointer) MTRClusterMicrowaveOvenMode {
	return MTRClusterMicrowaveOvenMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterMicrowaveOvenMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterMicrowaveOvenMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterMicrowaveOvenMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterMicrowaveOvenMode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode/readAttributeCurrentMode(with:)
func (m_ MTRClusterMicrowaveOvenMode) ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeCurrentModeWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeCurrentModeWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterMicrowaveOvenMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterMicrowaveOvenMode */



