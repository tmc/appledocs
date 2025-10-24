// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterLaundryDryerControls */


/* debug [class_header]: Header for MTRBaseClusterLaundryDryerControls */
// The class instance for the [MTRBaseClusterLaundryDryerControls] class.
var (
	MTRBaseClusterLaundryDryerControlsClass     _MTRBaseClusterLaundryDryerControlsClass
	MTRBaseClusterLaundryDryerControlsClassOnce sync.Once
)

func getMTRBaseClusterLaundryDryerControlsClass() _MTRBaseClusterLaundryDryerControlsClass {
	MTRBaseClusterLaundryDryerControlsClassOnce.Do(func() {
		MTRBaseClusterLaundryDryerControlsClass = _MTRBaseClusterLaundryDryerControlsClass{objc.GetClass("MTRBaseClusterLaundryDryerControls")}
	})
	return MTRBaseClusterLaundryDryerControlsClass
}

type _MTRBaseClusterLaundryDryerControlsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterLaundryDryerControls */
// An interface definition for the [MTRBaseClusterLaundryDryerControls] class.
type IMTRBaseClusterLaundryDryerControls interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterLaundryDryerControls */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterLaundryDryerControls */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterLaundryDryerControls */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterLaundryDryerControlsClass) Alloc() MTRBaseClusterLaundryDryerControls {
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterLaundryDryerControlsClass) New() MTRBaseClusterLaundryDryerControls {
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterLaundryDryerControls) Init() MTRBaseClusterLaundryDryerControls {
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterLaundryDryerControls) Autorelease() MTRBaseClusterLaundryDryerControls {
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterLaundryDryerControls creates a new MTRBaseClusterLaundryDryerControls instance.
func NewMTRBaseClusterLaundryDryerControls() MTRBaseClusterLaundryDryerControls {
	return getMTRBaseClusterLaundryDryerControlsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterLaundryDryerControls */
// Cluster Laundry Dryer Controls
//
// This cluster provides a way to access options associated with the operation of a laundry dryer device type.


// Cluster Laundry Dryer Controls
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls
type MTRBaseClusterLaundryDryerControls struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterLaundryDryerControlsFrom constructs a [MTRBaseClusterLaundryDryerControls] from an unsafe.Pointer.
//
// Cluster Laundry Dryer Controls
func MTRBaseClusterLaundryDryerControlsFrom(ptr unsafe.Pointer) MTRBaseClusterLaundryDryerControls {
	return MTRBaseClusterLaundryDryerControls{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterLaundryDryerControls */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/init(device:endpointID:queue:)
func NewMTRBaseClusterLaundryDryerControlsWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterLaundryDryerControls {
	instance := getMTRBaseClusterLaundryDryerControlsClass().Alloc()
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterLaundryDryerControlsWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterLaundryDryerControls */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterLaundryDryerControls */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterLaundryDryerControls */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterLaundryDryerControls */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterLaundryDryerControls */


