// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */


/* debug [class_header]: Header for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */
// The class instance for the [MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams] class.
var (
	MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass     _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass
	MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass() _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass {
	MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass = _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams")}
	})
	return MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass
}

type _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */
// An interface definition for the [MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams] class.
type IMTRDeviceEnergyManagementModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass) Alloc() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass) New() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) Init() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) Autorelease() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementModeClusterChangeToModeResponseParams creates a new MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams instance.
func NewMTRDeviceEnergyManagementModeClusterChangeToModeResponseParams() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	return getMTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams
type MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsFrom constructs a [MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	return MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */

// Initialize an MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	instance := getMTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams/status
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams/status
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclusterchangetomoderesponseparams/statustext
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}/* debug [instance_properties/getter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclusterchangetomoderesponseparams/statustext
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}/* debug [instance_properties/setter]: statusText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams */


