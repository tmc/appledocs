// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterMicrowaveOvenMode */


/* debug [class_header]: Header for MTRBaseClusterMicrowaveOvenMode */
// The class instance for the [MTRBaseClusterMicrowaveOvenMode] class.
var (
	MTRBaseClusterMicrowaveOvenModeClass     _MTRBaseClusterMicrowaveOvenModeClass
	MTRBaseClusterMicrowaveOvenModeClassOnce sync.Once
)

func getMTRBaseClusterMicrowaveOvenModeClass() _MTRBaseClusterMicrowaveOvenModeClass {
	MTRBaseClusterMicrowaveOvenModeClassOnce.Do(func() {
		MTRBaseClusterMicrowaveOvenModeClass = _MTRBaseClusterMicrowaveOvenModeClass{objc.GetClass("MTRBaseClusterMicrowaveOvenMode")}
	})
	return MTRBaseClusterMicrowaveOvenModeClass
}

type _MTRBaseClusterMicrowaveOvenModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterMicrowaveOvenMode */
// An interface definition for the [MTRBaseClusterMicrowaveOvenMode] class.
type IMTRBaseClusterMicrowaveOvenMode interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterMicrowaveOvenMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterMicrowaveOvenMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterMicrowaveOvenMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterMicrowaveOvenModeClass) Alloc() MTRBaseClusterMicrowaveOvenMode {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterMicrowaveOvenModeClass) New() MTRBaseClusterMicrowaveOvenMode {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterMicrowaveOvenMode) Init() MTRBaseClusterMicrowaveOvenMode {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterMicrowaveOvenMode) Autorelease() MTRBaseClusterMicrowaveOvenMode {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterMicrowaveOvenMode creates a new MTRBaseClusterMicrowaveOvenMode instance.
func NewMTRBaseClusterMicrowaveOvenMode() MTRBaseClusterMicrowaveOvenMode {
	return getMTRBaseClusterMicrowaveOvenModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterMicrowaveOvenMode */
// Cluster Microwave Oven Mode
//
// Attributes and commands for selecting a mode from a list of supported options.


// Cluster Microwave Oven Mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode
type MTRBaseClusterMicrowaveOvenMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterMicrowaveOvenModeFrom constructs a [MTRBaseClusterMicrowaveOvenMode] from an unsafe.Pointer.
//
// Cluster Microwave Oven Mode
func MTRBaseClusterMicrowaveOvenModeFrom(ptr unsafe.Pointer) MTRBaseClusterMicrowaveOvenMode {
	return MTRBaseClusterMicrowaveOvenMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterMicrowaveOvenMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterMicrowaveOvenMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterMicrowaveOvenMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterMicrowaveOvenMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterMicrowaveOvenMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterMicrowaveOvenMode */



