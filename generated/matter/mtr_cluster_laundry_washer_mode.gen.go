// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterLaundryWasherMode */


/* debug [class_header]: Header for MTRClusterLaundryWasherMode */
// The class instance for the [MTRClusterLaundryWasherMode] class.
var (
	MTRClusterLaundryWasherModeClass     _MTRClusterLaundryWasherModeClass
	MTRClusterLaundryWasherModeClassOnce sync.Once
)

func getMTRClusterLaundryWasherModeClass() _MTRClusterLaundryWasherModeClass {
	MTRClusterLaundryWasherModeClassOnce.Do(func() {
		MTRClusterLaundryWasherModeClass = _MTRClusterLaundryWasherModeClass{objc.GetClass("MTRClusterLaundryWasherMode")}
	})
	return MTRClusterLaundryWasherModeClass
}

type _MTRClusterLaundryWasherModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterLaundryWasherMode */
// An interface definition for the [MTRClusterLaundryWasherMode] class.
type IMTRClusterLaundryWasherMode interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterLaundryWasherMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterLaundryWasherMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterLaundryWasherMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterLaundryWasherModeClass) Alloc() MTRClusterLaundryWasherMode {
	rv := objc.Send[MTRClusterLaundryWasherMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterLaundryWasherModeClass) New() MTRClusterLaundryWasherMode {
	rv := objc.Send[MTRClusterLaundryWasherMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterLaundryWasherMode) Init() MTRClusterLaundryWasherMode {
	rv := objc.Send[MTRClusterLaundryWasherMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterLaundryWasherMode) Autorelease() MTRClusterLaundryWasherMode {
	rv := objc.Send[MTRClusterLaundryWasherMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterLaundryWasherMode creates a new MTRClusterLaundryWasherMode instance.
func NewMTRClusterLaundryWasherMode() MTRClusterLaundryWasherMode {
	return getMTRClusterLaundryWasherModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterLaundryWasherMode */
// Cluster Laundry Washer Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Laundry Washer Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode
type MTRClusterLaundryWasherMode struct {
	MTRGenericCluster
}

// MTRClusterLaundryWasherModeFrom constructs a [MTRClusterLaundryWasherMode] from an unsafe.Pointer.
//
// Cluster Laundry Washer Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterLaundryWasherModeFrom(ptr unsafe.Pointer) MTRClusterLaundryWasherMode {
	return MTRClusterLaundryWasherMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterLaundryWasherMode */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/init(device:endpointID:queue:)
func NewMTRClusterLaundryWasherModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterLaundryWasherMode {
	instance := getMTRClusterLaundryWasherModeClass().Alloc()
	rv := objc.Send[MTRClusterLaundryWasherMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterLaundryWasherModeWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterLaundryWasherMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterLaundryWasherMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterLaundryWasherMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterLaundryWasherMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterLaundryWasherMode */


