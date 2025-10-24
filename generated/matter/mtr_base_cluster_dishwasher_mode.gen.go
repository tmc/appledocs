// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterDishwasherMode */


/* debug [class_header]: Header for MTRBaseClusterDishwasherMode */
// The class instance for the [MTRBaseClusterDishwasherMode] class.
var (
	MTRBaseClusterDishwasherModeClass     _MTRBaseClusterDishwasherModeClass
	MTRBaseClusterDishwasherModeClassOnce sync.Once
)

func getMTRBaseClusterDishwasherModeClass() _MTRBaseClusterDishwasherModeClass {
	MTRBaseClusterDishwasherModeClassOnce.Do(func() {
		MTRBaseClusterDishwasherModeClass = _MTRBaseClusterDishwasherModeClass{objc.GetClass("MTRBaseClusterDishwasherMode")}
	})
	return MTRBaseClusterDishwasherModeClass
}

type _MTRBaseClusterDishwasherModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterDishwasherMode */
// An interface definition for the [MTRBaseClusterDishwasherMode] class.
type IMTRBaseClusterDishwasherMode interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterDishwasherMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterDishwasherMode */
	// methods:
	ChangeToModeWithParamsCompletion(params IMTRDishwasherModeClusterChangeToModeParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterDishwasherMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDishwasherModeClass) Alloc() MTRBaseClusterDishwasherMode {
	rv := objc.Send[MTRBaseClusterDishwasherMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterDishwasherModeClass) New() MTRBaseClusterDishwasherMode {
	rv := objc.Send[MTRBaseClusterDishwasherMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDishwasherMode) Init() MTRBaseClusterDishwasherMode {
	rv := objc.Send[MTRBaseClusterDishwasherMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDishwasherMode) Autorelease() MTRBaseClusterDishwasherMode {
	rv := objc.Send[MTRBaseClusterDishwasherMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDishwasherMode creates a new MTRBaseClusterDishwasherMode instance.
func NewMTRBaseClusterDishwasherMode() MTRBaseClusterDishwasherMode {
	return getMTRBaseClusterDishwasherModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterDishwasherMode */
// Cluster Dishwasher Mode
//
// Attributes and commands for selecting a mode from a list of supported options.


// Cluster Dishwasher Mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherMode
type MTRBaseClusterDishwasherMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDishwasherModeFrom constructs a [MTRBaseClusterDishwasherMode] from an unsafe.Pointer.
//
// Cluster Dishwasher Mode
func MTRBaseClusterDishwasherModeFrom(ptr unsafe.Pointer) MTRBaseClusterDishwasherMode {
	return MTRBaseClusterDishwasherMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterDishwasherMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterDishwasherMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterDishwasherMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterDishwasherMode */

// Command ChangeToMode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherMode/changeToMode(with:completion:)
func (m_ MTRBaseClusterDishwasherMode) ChangeToModeWithParamsCompletion(params IMTRDishwasherModeClusterChangeToModeParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ChangeToModeWithParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterDishwasherMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterDishwasherMode */



