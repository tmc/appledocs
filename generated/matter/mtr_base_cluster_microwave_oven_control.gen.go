// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterMicrowaveOvenControl */


/* debug [class_header]: Header for MTRBaseClusterMicrowaveOvenControl */
// The class instance for the [MTRBaseClusterMicrowaveOvenControl] class.
var (
	MTRBaseClusterMicrowaveOvenControlClass     _MTRBaseClusterMicrowaveOvenControlClass
	MTRBaseClusterMicrowaveOvenControlClassOnce sync.Once
)

func getMTRBaseClusterMicrowaveOvenControlClass() _MTRBaseClusterMicrowaveOvenControlClass {
	MTRBaseClusterMicrowaveOvenControlClassOnce.Do(func() {
		MTRBaseClusterMicrowaveOvenControlClass = _MTRBaseClusterMicrowaveOvenControlClass{objc.GetClass("MTRBaseClusterMicrowaveOvenControl")}
	})
	return MTRBaseClusterMicrowaveOvenControlClass
}

type _MTRBaseClusterMicrowaveOvenControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterMicrowaveOvenControl */
// An interface definition for the [MTRBaseClusterMicrowaveOvenControl] class.
type IMTRBaseClusterMicrowaveOvenControl interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterMicrowaveOvenControl */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterMicrowaveOvenControl */
	// methods:
	AddMoreTimeWithParamsCompletion(params IMTRMicrowaveOvenControlClusterAddMoreTimeParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterMicrowaveOvenControl */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterMicrowaveOvenControlClass) Alloc() MTRBaseClusterMicrowaveOvenControl {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterMicrowaveOvenControlClass) New() MTRBaseClusterMicrowaveOvenControl {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterMicrowaveOvenControl) Init() MTRBaseClusterMicrowaveOvenControl {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterMicrowaveOvenControl) Autorelease() MTRBaseClusterMicrowaveOvenControl {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterMicrowaveOvenControl creates a new MTRBaseClusterMicrowaveOvenControl instance.
func NewMTRBaseClusterMicrowaveOvenControl() MTRBaseClusterMicrowaveOvenControl {
	return getMTRBaseClusterMicrowaveOvenControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterMicrowaveOvenControl */
// Cluster Microwave Oven Control
//
// Attributes and commands for configuring the microwave oven control, and reporting cooking stats.


// Cluster Microwave Oven Control
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl
type MTRBaseClusterMicrowaveOvenControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterMicrowaveOvenControlFrom constructs a [MTRBaseClusterMicrowaveOvenControl] from an unsafe.Pointer.
//
// Cluster Microwave Oven Control
func MTRBaseClusterMicrowaveOvenControlFrom(ptr unsafe.Pointer) MTRBaseClusterMicrowaveOvenControl {
	return MTRBaseClusterMicrowaveOvenControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterMicrowaveOvenControl *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterMicrowaveOvenControl */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterMicrowaveOvenControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterMicrowaveOvenControl */

// Command AddMoreTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/addMoreTime(with:completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) AddMoreTimeWithParamsCompletion(params IMTRMicrowaveOvenControlClusterAddMoreTimeParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addMoreTimeWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: AddMoreTimeWithParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterMicrowaveOvenControl */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterMicrowaveOvenControl */



