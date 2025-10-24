// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterPowerTopology */


/* debug [class_header]: Header for MTRBaseClusterPowerTopology */
// The class instance for the [MTRBaseClusterPowerTopology] class.
var (
	MTRBaseClusterPowerTopologyClass     _MTRBaseClusterPowerTopologyClass
	MTRBaseClusterPowerTopologyClassOnce sync.Once
)

func getMTRBaseClusterPowerTopologyClass() _MTRBaseClusterPowerTopologyClass {
	MTRBaseClusterPowerTopologyClassOnce.Do(func() {
		MTRBaseClusterPowerTopologyClass = _MTRBaseClusterPowerTopologyClass{objc.GetClass("MTRBaseClusterPowerTopology")}
	})
	return MTRBaseClusterPowerTopologyClass
}

type _MTRBaseClusterPowerTopologyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterPowerTopology */
// An interface definition for the [MTRBaseClusterPowerTopology] class.
type IMTRBaseClusterPowerTopology interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterPowerTopology */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterPowerTopology */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterPowerTopology */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterPowerTopologyClass) Alloc() MTRBaseClusterPowerTopology {
	rv := objc.Send[MTRBaseClusterPowerTopology](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterPowerTopologyClass) New() MTRBaseClusterPowerTopology {
	rv := objc.Send[MTRBaseClusterPowerTopology](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterPowerTopology) Init() MTRBaseClusterPowerTopology {
	rv := objc.Send[MTRBaseClusterPowerTopology](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterPowerTopology) Autorelease() MTRBaseClusterPowerTopology {
	rv := objc.Send[MTRBaseClusterPowerTopology](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterPowerTopology creates a new MTRBaseClusterPowerTopology instance.
func NewMTRBaseClusterPowerTopology() MTRBaseClusterPowerTopology {
	return getMTRBaseClusterPowerTopologyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterPowerTopology */
// Cluster Power Topology
//
// The Power Topology Cluster provides a mechanism for expressing how power is flowing between endpoints.


// Cluster Power Topology
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology
type MTRBaseClusterPowerTopology struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterPowerTopologyFrom constructs a [MTRBaseClusterPowerTopology] from an unsafe.Pointer.
//
// Cluster Power Topology
func MTRBaseClusterPowerTopologyFrom(ptr unsafe.Pointer) MTRBaseClusterPowerTopology {
	return MTRBaseClusterPowerTopology{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterPowerTopology */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/init(device:endpointID:queue:)
func NewMTRBaseClusterPowerTopologyWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterPowerTopology {
	instance := getMTRBaseClusterPowerTopologyClass().Alloc()
	rv := objc.Send[MTRBaseClusterPowerTopology](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterPowerTopologyWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterPowerTopology */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterPowerTopology */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterPowerTopology */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterPowerTopology */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterPowerTopology */


