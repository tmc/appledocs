// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterLaundryDryerControls */


/* debug [class_header]: Header for MTRClusterLaundryDryerControls */
// The class instance for the [MTRClusterLaundryDryerControls] class.
var (
	MTRClusterLaundryDryerControlsClass     _MTRClusterLaundryDryerControlsClass
	MTRClusterLaundryDryerControlsClassOnce sync.Once
)

func getMTRClusterLaundryDryerControlsClass() _MTRClusterLaundryDryerControlsClass {
	MTRClusterLaundryDryerControlsClassOnce.Do(func() {
		MTRClusterLaundryDryerControlsClass = _MTRClusterLaundryDryerControlsClass{objc.GetClass("MTRClusterLaundryDryerControls")}
	})
	return MTRClusterLaundryDryerControlsClass
}

type _MTRClusterLaundryDryerControlsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterLaundryDryerControls */
// An interface definition for the [MTRClusterLaundryDryerControls] class.
type IMTRClusterLaundryDryerControls interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterLaundryDryerControls */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterLaundryDryerControls */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterLaundryDryerControls */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterLaundryDryerControlsClass) Alloc() MTRClusterLaundryDryerControls {
	rv := objc.Send[MTRClusterLaundryDryerControls](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterLaundryDryerControlsClass) New() MTRClusterLaundryDryerControls {
	rv := objc.Send[MTRClusterLaundryDryerControls](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterLaundryDryerControls) Init() MTRClusterLaundryDryerControls {
	rv := objc.Send[MTRClusterLaundryDryerControls](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterLaundryDryerControls) Autorelease() MTRClusterLaundryDryerControls {
	rv := objc.Send[MTRClusterLaundryDryerControls](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterLaundryDryerControls creates a new MTRClusterLaundryDryerControls instance.
func NewMTRClusterLaundryDryerControls() MTRClusterLaundryDryerControls {
	return getMTRClusterLaundryDryerControlsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterLaundryDryerControls */
// Cluster Laundry Dryer Controls This cluster provides a way to access options associated with the operation of a laundry dryer device type.


// Cluster Laundry Dryer Controls This cluster provides a way to access options associated with the operation of a laundry dryer device type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls
type MTRClusterLaundryDryerControls struct {
	MTRGenericCluster
}

// MTRClusterLaundryDryerControlsFrom constructs a [MTRClusterLaundryDryerControls] from an unsafe.Pointer.
//
// Cluster Laundry Dryer Controls This cluster provides a way to access options associated with the operation of a laundry dryer device type.
func MTRClusterLaundryDryerControlsFrom(ptr unsafe.Pointer) MTRClusterLaundryDryerControls {
	return MTRClusterLaundryDryerControls{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterLaundryDryerControls */

// The queue is currently unused, but may be used in the future for calling completions for command invocations if commands are added to this cluster.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/init(device:endpointID:queue:)
func NewMTRClusterLaundryDryerControlsWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterLaundryDryerControls {
	instance := getMTRClusterLaundryDryerControlsClass().Alloc()
	rv := objc.Send[MTRClusterLaundryDryerControls](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterLaundryDryerControlsWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterLaundryDryerControls */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterLaundryDryerControls */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterLaundryDryerControls */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterLaundryDryerControls */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterLaundryDryerControls */


