// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterPowerTopology */


/* debug [class_header]: Header for MTRClusterPowerTopology */
// The class instance for the [MTRClusterPowerTopology] class.
var (
	MTRClusterPowerTopologyClass     _MTRClusterPowerTopologyClass
	MTRClusterPowerTopologyClassOnce sync.Once
)

func getMTRClusterPowerTopologyClass() _MTRClusterPowerTopologyClass {
	MTRClusterPowerTopologyClassOnce.Do(func() {
		MTRClusterPowerTopologyClass = _MTRClusterPowerTopologyClass{objc.GetClass("MTRClusterPowerTopology")}
	})
	return MTRClusterPowerTopologyClass
}

type _MTRClusterPowerTopologyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterPowerTopology */
// An interface definition for the [MTRClusterPowerTopology] class.
type IMTRClusterPowerTopology interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterPowerTopology */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterPowerTopology */
	// methods:
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterPowerTopology */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPowerTopologyClass) Alloc() MTRClusterPowerTopology {
	rv := objc.Send[MTRClusterPowerTopology](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterPowerTopologyClass) New() MTRClusterPowerTopology {
	rv := objc.Send[MTRClusterPowerTopology](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPowerTopology) Init() MTRClusterPowerTopology {
	rv := objc.Send[MTRClusterPowerTopology](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPowerTopology) Autorelease() MTRClusterPowerTopology {
	rv := objc.Send[MTRClusterPowerTopology](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPowerTopology creates a new MTRClusterPowerTopology instance.
func NewMTRClusterPowerTopology() MTRClusterPowerTopology {
	return getMTRClusterPowerTopologyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterPowerTopology */
// Cluster Power Topology The Power Topology Cluster provides a mechanism for expressing how power is flowing between endpoints.


// Cluster Power Topology The Power Topology Cluster provides a mechanism for expressing how power is flowing between endpoints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology
type MTRClusterPowerTopology struct {
	MTRGenericCluster
}

// MTRClusterPowerTopologyFrom constructs a [MTRClusterPowerTopology] from an unsafe.Pointer.
//
// Cluster Power Topology The Power Topology Cluster provides a mechanism for expressing how power is flowing between endpoints.
func MTRClusterPowerTopologyFrom(ptr unsafe.Pointer) MTRClusterPowerTopology {
	return MTRClusterPowerTopology{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterPowerTopology *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterPowerTopology */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterPowerTopology */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterPowerTopology */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology/readAttributeClusterRevision(with:)
func (m_ MTRClusterPowerTopology) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterPowerTopology */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterPowerTopology */



