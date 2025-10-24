// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterCommissionerControl */


/* debug [class_header]: Header for MTRBaseClusterCommissionerControl */
// The class instance for the [MTRBaseClusterCommissionerControl] class.
var (
	MTRBaseClusterCommissionerControlClass     _MTRBaseClusterCommissionerControlClass
	MTRBaseClusterCommissionerControlClassOnce sync.Once
)

func getMTRBaseClusterCommissionerControlClass() _MTRBaseClusterCommissionerControlClass {
	MTRBaseClusterCommissionerControlClassOnce.Do(func() {
		MTRBaseClusterCommissionerControlClass = _MTRBaseClusterCommissionerControlClass{objc.GetClass("MTRBaseClusterCommissionerControl")}
	})
	return MTRBaseClusterCommissionerControlClass
}

type _MTRBaseClusterCommissionerControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterCommissionerControl */
// An interface definition for the [MTRBaseClusterCommissionerControl] class.
type IMTRBaseClusterCommissionerControl interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterCommissionerControl */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterCommissionerControl */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterCommissionerControl */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterCommissionerControlClass) Alloc() MTRBaseClusterCommissionerControl {
	rv := objc.Send[MTRBaseClusterCommissionerControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterCommissionerControlClass) New() MTRBaseClusterCommissionerControl {
	rv := objc.Send[MTRBaseClusterCommissionerControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterCommissionerControl) Init() MTRBaseClusterCommissionerControl {
	rv := objc.Send[MTRBaseClusterCommissionerControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterCommissionerControl) Autorelease() MTRBaseClusterCommissionerControl {
	rv := objc.Send[MTRBaseClusterCommissionerControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterCommissionerControl creates a new MTRBaseClusterCommissionerControl instance.
func NewMTRBaseClusterCommissionerControl() MTRBaseClusterCommissionerControl {
	return getMTRBaseClusterCommissionerControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterCommissionerControl */
// Cluster Commissioner Control
//
// Supports the ability for clients to request the commissioning of themselves or other nodes onto a fabric which the cluster server can commission onto.


// Cluster Commissioner Control
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterCommissionerControl
type MTRBaseClusterCommissionerControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterCommissionerControlFrom constructs a [MTRBaseClusterCommissionerControl] from an unsafe.Pointer.
//
// Cluster Commissioner Control
func MTRBaseClusterCommissionerControlFrom(ptr unsafe.Pointer) MTRBaseClusterCommissionerControl {
	return MTRBaseClusterCommissionerControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterCommissionerControl */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterCommissionerControl/init(device:endpointID:queue:)
func NewMTRBaseClusterCommissionerControlWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterCommissionerControl {
	instance := getMTRBaseClusterCommissionerControlClass().Alloc()
	rv := objc.Send[MTRBaseClusterCommissionerControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterCommissionerControlWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterCommissionerControl */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterCommissionerControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterCommissionerControl */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterCommissionerControl */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterCommissionerControl */


