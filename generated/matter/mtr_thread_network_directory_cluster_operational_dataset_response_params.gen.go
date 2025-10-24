// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */


/* debug [class_header]: Header for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */
// The class instance for the [MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams] class.
var (
	MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass     _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass
	MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass() _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass {
	MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass = _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass{objc.GetClass("MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams")}
	})
	return MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass
}

type _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */
// An interface definition for the [MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams] class.
type IMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */
	// properties:
	OperationalDataset() objc.IObject /* cross-framework: NSData */
	SetOperationalDataset(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass) Alloc() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass) New() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams) Init() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams) Autorelease() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams creates a new MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams instance.
func NewMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	return getMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams
type MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsFrom constructs a [MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	return MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams/operationalDataset
func (m_ MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams) OperationalDataset() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("operationalDataset"))
	return rv
}/* debug [instance_properties/getter]: operationalDataset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams/operationalDataset
func (m_ MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams) SetOperationalDataset(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalDataset:"), value)
}/* debug [instance_properties/setter]: operationalDataset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams */



