// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterLaundryWasherControls */


/* debug [class_header]: Header for MTRBaseClusterLaundryWasherControls */
// The class instance for the [MTRBaseClusterLaundryWasherControls] class.
var (
	MTRBaseClusterLaundryWasherControlsClass     _MTRBaseClusterLaundryWasherControlsClass
	MTRBaseClusterLaundryWasherControlsClassOnce sync.Once
)

func getMTRBaseClusterLaundryWasherControlsClass() _MTRBaseClusterLaundryWasherControlsClass {
	MTRBaseClusterLaundryWasherControlsClassOnce.Do(func() {
		MTRBaseClusterLaundryWasherControlsClass = _MTRBaseClusterLaundryWasherControlsClass{objc.GetClass("MTRBaseClusterLaundryWasherControls")}
	})
	return MTRBaseClusterLaundryWasherControlsClass
}

type _MTRBaseClusterLaundryWasherControlsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterLaundryWasherControls */
// An interface definition for the [MTRBaseClusterLaundryWasherControls] class.
type IMTRBaseClusterLaundryWasherControls interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterLaundryWasherControls */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterLaundryWasherControls */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterLaundryWasherControls */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterLaundryWasherControlsClass) Alloc() MTRBaseClusterLaundryWasherControls {
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterLaundryWasherControlsClass) New() MTRBaseClusterLaundryWasherControls {
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterLaundryWasherControls) Init() MTRBaseClusterLaundryWasherControls {
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterLaundryWasherControls) Autorelease() MTRBaseClusterLaundryWasherControls {
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterLaundryWasherControls creates a new MTRBaseClusterLaundryWasherControls instance.
func NewMTRBaseClusterLaundryWasherControls() MTRBaseClusterLaundryWasherControls {
	return getMTRBaseClusterLaundryWasherControlsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterLaundryWasherControls */
// Cluster Laundry Washer Controls
//
// This cluster supports remotely monitoring and controlling the different types of functionality available to a washing device, such as a washing machine.


// Cluster Laundry Washer Controls
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls
type MTRBaseClusterLaundryWasherControls struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterLaundryWasherControlsFrom constructs a [MTRBaseClusterLaundryWasherControls] from an unsafe.Pointer.
//
// Cluster Laundry Washer Controls
func MTRBaseClusterLaundryWasherControlsFrom(ptr unsafe.Pointer) MTRBaseClusterLaundryWasherControls {
	return MTRBaseClusterLaundryWasherControls{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterLaundryWasherControls */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/init(device:endpointID:queue:)
func NewMTRBaseClusterLaundryWasherControlsWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterLaundryWasherControls {
	instance := getMTRBaseClusterLaundryWasherControlsClass().Alloc()
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterLaundryWasherControlsWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterLaundryWasherControls */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterLaundryWasherControls */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterLaundryWasherControls */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterLaundryWasherControls */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterLaundryWasherControls */


