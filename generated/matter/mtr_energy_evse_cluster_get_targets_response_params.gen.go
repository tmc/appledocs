// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterGetTargetsResponseParams */


/* debug [class_header]: Header for MTREnergyEVSEClusterGetTargetsResponseParams */
// The class instance for the [MTREnergyEVSEClusterGetTargetsResponseParams] class.
var (
	MTREnergyEVSEClusterGetTargetsResponseParamsClass     _MTREnergyEVSEClusterGetTargetsResponseParamsClass
	MTREnergyEVSEClusterGetTargetsResponseParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterGetTargetsResponseParamsClass() _MTREnergyEVSEClusterGetTargetsResponseParamsClass {
	MTREnergyEVSEClusterGetTargetsResponseParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterGetTargetsResponseParamsClass = _MTREnergyEVSEClusterGetTargetsResponseParamsClass{objc.GetClass("MTREnergyEVSEClusterGetTargetsResponseParams")}
	})
	return MTREnergyEVSEClusterGetTargetsResponseParamsClass
}

type _MTREnergyEVSEClusterGetTargetsResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterGetTargetsResponseParams */
// An interface definition for the [MTREnergyEVSEClusterGetTargetsResponseParams] class.
type IMTREnergyEVSEClusterGetTargetsResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterGetTargetsResponseParams */
	// properties:
	ChargingTargetSchedules() objc.IObject /* cross-framework: NSArray */
	SetChargingTargetSchedules(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterGetTargetsResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterGetTargetsResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterGetTargetsResponseParamsClass) Alloc() MTREnergyEVSEClusterGetTargetsResponseParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterGetTargetsResponseParamsClass) New() MTREnergyEVSEClusterGetTargetsResponseParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterGetTargetsResponseParams) Init() MTREnergyEVSEClusterGetTargetsResponseParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterGetTargetsResponseParams) Autorelease() MTREnergyEVSEClusterGetTargetsResponseParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterGetTargetsResponseParams creates a new MTREnergyEVSEClusterGetTargetsResponseParams instance.
func NewMTREnergyEVSEClusterGetTargetsResponseParams() MTREnergyEVSEClusterGetTargetsResponseParams {
	return getMTREnergyEVSEClusterGetTargetsResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterGetTargetsResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsResponseParams
type MTREnergyEVSEClusterGetTargetsResponseParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterGetTargetsResponseParamsFrom constructs a [MTREnergyEVSEClusterGetTargetsResponseParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterGetTargetsResponseParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterGetTargetsResponseParams {
	return MTREnergyEVSEClusterGetTargetsResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterGetTargetsResponseParams */

// Initialize an MTREnergyEVSEClusterGetTargetsResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsResponseParams/init(responseValue:)
func NewMTREnergyEVSEClusterGetTargetsResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTREnergyEVSEClusterGetTargetsResponseParams {
	instance := getMTREnergyEVSEClusterGetTargetsResponseParamsClass().Alloc()
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTREnergyEVSEClusterGetTargetsResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterGetTargetsResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterGetTargetsResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterGetTargetsResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterGetTargetsResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsResponseParams/chargingTargetSchedules
func (m_ MTREnergyEVSEClusterGetTargetsResponseParams) ChargingTargetSchedules() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("chargingTargetSchedules"))
	return rv
}/* debug [instance_properties/getter]: chargingTargetSchedules */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsResponseParams/chargingTargetSchedules
func (m_ MTREnergyEVSEClusterGetTargetsResponseParams) SetChargingTargetSchedules(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChargingTargetSchedules:"), value)
}/* debug [instance_properties/setter]: chargingTargetSchedules */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterGetTargetsResponseParams */


