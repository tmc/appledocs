// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterCommissionerControl */


/* debug [class_header]: Header for MTRClusterCommissionerControl */
// The class instance for the [MTRClusterCommissionerControl] class.
var (
	MTRClusterCommissionerControlClass     _MTRClusterCommissionerControlClass
	MTRClusterCommissionerControlClassOnce sync.Once
)

func getMTRClusterCommissionerControlClass() _MTRClusterCommissionerControlClass {
	MTRClusterCommissionerControlClassOnce.Do(func() {
		MTRClusterCommissionerControlClass = _MTRClusterCommissionerControlClass{objc.GetClass("MTRClusterCommissionerControl")}
	})
	return MTRClusterCommissionerControlClass
}

type _MTRClusterCommissionerControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterCommissionerControl */
// An interface definition for the [MTRClusterCommissionerControl] class.
type IMTRClusterCommissionerControl interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterCommissionerControl */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterCommissionerControl */
	// methods:
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterCommissionerControl */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterCommissionerControlClass) Alloc() MTRClusterCommissionerControl {
	rv := objc.Send[MTRClusterCommissionerControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterCommissionerControlClass) New() MTRClusterCommissionerControl {
	rv := objc.Send[MTRClusterCommissionerControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterCommissionerControl) Init() MTRClusterCommissionerControl {
	rv := objc.Send[MTRClusterCommissionerControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterCommissionerControl) Autorelease() MTRClusterCommissionerControl {
	rv := objc.Send[MTRClusterCommissionerControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterCommissionerControl creates a new MTRClusterCommissionerControl instance.
func NewMTRClusterCommissionerControl() MTRClusterCommissionerControl {
	return getMTRClusterCommissionerControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterCommissionerControl */
// Cluster Commissioner Control Supports the ability for clients to request the commissioning of themselves or other nodes onto a fabric which the cluster server can commission onto.


// Cluster Commissioner Control Supports the ability for clients to request the commissioning of themselves or other nodes onto a fabric which the cluster server can commission onto.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl
type MTRClusterCommissionerControl struct {
	MTRGenericCluster
}

// MTRClusterCommissionerControlFrom constructs a [MTRClusterCommissionerControl] from an unsafe.Pointer.
//
// Cluster Commissioner Control Supports the ability for clients to request the commissioning of themselves or other nodes onto a fabric which the cluster server can commission onto.
func MTRClusterCommissionerControlFrom(ptr unsafe.Pointer) MTRClusterCommissionerControl {
	return MTRClusterCommissionerControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterCommissionerControl *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterCommissionerControl */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterCommissionerControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterCommissionerControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/readAttributeFeatureMap(with:)
func (m_ MTRClusterCommissionerControl) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterCommissionerControl */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterCommissionerControl */



