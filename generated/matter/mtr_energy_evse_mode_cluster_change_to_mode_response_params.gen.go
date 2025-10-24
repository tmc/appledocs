// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEModeClusterChangeToModeResponseParams */


/* debug [class_header]: Header for MTREnergyEVSEModeClusterChangeToModeResponseParams */
// The class instance for the [MTREnergyEVSEModeClusterChangeToModeResponseParams] class.
var (
	MTREnergyEVSEModeClusterChangeToModeResponseParamsClass     _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass
	MTREnergyEVSEModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTREnergyEVSEModeClusterChangeToModeResponseParamsClass() _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass {
	MTREnergyEVSEModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTREnergyEVSEModeClusterChangeToModeResponseParamsClass = _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTREnergyEVSEModeClusterChangeToModeResponseParams")}
	})
	return MTREnergyEVSEModeClusterChangeToModeResponseParamsClass
}

type _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEModeClusterChangeToModeResponseParams */
// An interface definition for the [MTREnergyEVSEModeClusterChangeToModeResponseParams] class.
type IMTREnergyEVSEModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEModeClusterChangeToModeResponseParams */
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEModeClusterChangeToModeResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEModeClusterChangeToModeResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass) Alloc() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass) New() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) Init() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) Autorelease() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEModeClusterChangeToModeResponseParams creates a new MTREnergyEVSEModeClusterChangeToModeResponseParams instance.
func NewMTREnergyEVSEModeClusterChangeToModeResponseParams() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	return getMTREnergyEVSEModeClusterChangeToModeResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEModeClusterChangeToModeResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams
type MTREnergyEVSEModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTREnergyEVSEModeClusterChangeToModeResponseParamsFrom constructs a [MTREnergyEVSEModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTREnergyEVSEModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEModeClusterChangeToModeResponseParams {
	return MTREnergyEVSEModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEModeClusterChangeToModeResponseParams */

// Initialize an MTREnergyEVSEModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTREnergyEVSEModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTREnergyEVSEModeClusterChangeToModeResponseParams {
	instance := getMTREnergyEVSEModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTREnergyEVSEModeClusterChangeToModeResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEModeClusterChangeToModeResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEModeClusterChangeToModeResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEModeClusterChangeToModeResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEModeClusterChangeToModeResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevsemodeclusterchangetomoderesponseparams/status
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevsemodeclusterchangetomoderesponseparams/status
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevsemodeclusterchangetomoderesponseparams/statustext
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}/* debug [instance_properties/getter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevsemodeclusterchangetomoderesponseparams/statustext
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}/* debug [instance_properties/setter]: statusText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEModeClusterChangeToModeResponseParams */


