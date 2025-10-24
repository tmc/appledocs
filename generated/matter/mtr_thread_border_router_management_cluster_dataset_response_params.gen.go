// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadBorderRouterManagementClusterDatasetResponseParams */


/* debug [class_header]: Header for MTRThreadBorderRouterManagementClusterDatasetResponseParams */
// The class instance for the [MTRThreadBorderRouterManagementClusterDatasetResponseParams] class.
var (
	MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass     _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass
	MTRThreadBorderRouterManagementClusterDatasetResponseParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterDatasetResponseParamsClass() _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass {
	MTRThreadBorderRouterManagementClusterDatasetResponseParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass = _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterDatasetResponseParams")}
	})
	return MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass
}

type _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadBorderRouterManagementClusterDatasetResponseParams */
// An interface definition for the [MTRThreadBorderRouterManagementClusterDatasetResponseParams] class.
type IMTRThreadBorderRouterManagementClusterDatasetResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadBorderRouterManagementClusterDatasetResponseParams */
	// properties:
	Dataset() foundation.Data
	SetDataset(value foundation.Data)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadBorderRouterManagementClusterDatasetResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadBorderRouterManagementClusterDatasetResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass) Alloc() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass) New() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterDatasetResponseParams) Init() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterDatasetResponseParams) Autorelease() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterDatasetResponseParams creates a new MTRThreadBorderRouterManagementClusterDatasetResponseParams instance.
func NewMTRThreadBorderRouterManagementClusterDatasetResponseParams() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	return getMTRThreadBorderRouterManagementClusterDatasetResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadBorderRouterManagementClusterDatasetResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterDatasetResponseParams
type MTRThreadBorderRouterManagementClusterDatasetResponseParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterDatasetResponseParamsFrom constructs a [MTRThreadBorderRouterManagementClusterDatasetResponseParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterDatasetResponseParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	return MTRThreadBorderRouterManagementClusterDatasetResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadBorderRouterManagementClusterDatasetResponseParams */

// Initialize an MTRThreadBorderRouterManagementClusterDatasetResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterDatasetResponseParams/init(responseValue:)
func NewMTRThreadBorderRouterManagementClusterDatasetResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	instance := getMTRThreadBorderRouterManagementClusterDatasetResponseParamsClass().Alloc()
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRThreadBorderRouterManagementClusterDatasetResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadBorderRouterManagementClusterDatasetResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadBorderRouterManagementClusterDatasetResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadBorderRouterManagementClusterDatasetResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadBorderRouterManagementClusterDatasetResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclusterdatasetresponseparams/dataset
func (m_ MTRThreadBorderRouterManagementClusterDatasetResponseParams) Dataset() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("dataset"))
	return rv
}/* debug [instance_properties/getter]: dataset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclusterdatasetresponseparams/dataset
func (m_ MTRThreadBorderRouterManagementClusterDatasetResponseParams) SetDataset(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataset:"), value)
}/* debug [instance_properties/setter]: dataset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadBorderRouterManagementClusterDatasetResponseParams */


