// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterDishwasherMode */


/* debug [class_header]: Header for MTRClusterDishwasherMode */
// The class instance for the [MTRClusterDishwasherMode] class.
var (
	MTRClusterDishwasherModeClass     _MTRClusterDishwasherModeClass
	MTRClusterDishwasherModeClassOnce sync.Once
)

func getMTRClusterDishwasherModeClass() _MTRClusterDishwasherModeClass {
	MTRClusterDishwasherModeClassOnce.Do(func() {
		MTRClusterDishwasherModeClass = _MTRClusterDishwasherModeClass{objc.GetClass("MTRClusterDishwasherMode")}
	})
	return MTRClusterDishwasherModeClass
}

type _MTRClusterDishwasherModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterDishwasherMode */
// An interface definition for the [MTRClusterDishwasherMode] class.
type IMTRClusterDishwasherMode interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterDishwasherMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterDishwasherMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterDishwasherMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDishwasherModeClass) Alloc() MTRClusterDishwasherMode {
	rv := objc.Send[MTRClusterDishwasherMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterDishwasherModeClass) New() MTRClusterDishwasherMode {
	rv := objc.Send[MTRClusterDishwasherMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDishwasherMode) Init() MTRClusterDishwasherMode {
	rv := objc.Send[MTRClusterDishwasherMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDishwasherMode) Autorelease() MTRClusterDishwasherMode {
	rv := objc.Send[MTRClusterDishwasherMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDishwasherMode creates a new MTRClusterDishwasherMode instance.
func NewMTRClusterDishwasherMode() MTRClusterDishwasherMode {
	return getMTRClusterDishwasherModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterDishwasherMode */
// Cluster Dishwasher Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Dishwasher Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode
type MTRClusterDishwasherMode struct {
	MTRGenericCluster
}

// MTRClusterDishwasherModeFrom constructs a [MTRClusterDishwasherMode] from an unsafe.Pointer.
//
// Cluster Dishwasher Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterDishwasherModeFrom(ptr unsafe.Pointer) MTRClusterDishwasherMode {
	return MTRClusterDishwasherMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterDishwasherMode */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/init(device:endpointID:queue:)
func NewMTRClusterDishwasherModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterDishwasherMode {
	instance := getMTRClusterDishwasherModeClass().Alloc()
	rv := objc.Send[MTRClusterDishwasherMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterDishwasherModeWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterDishwasherMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterDishwasherMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterDishwasherMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterDishwasherMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterDishwasherMode */


