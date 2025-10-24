// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRGeneralDiagnosticsClusterPayloadTestResponseParams */


/* debug [class_header]: Header for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */
// The class instance for the [MTRGeneralDiagnosticsClusterPayloadTestResponseParams] class.
var (
	MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass     _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass
	MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass() _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass {
	MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass = _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass{objc.GetClass("MTRGeneralDiagnosticsClusterPayloadTestResponseParams")}
	})
	return MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass
}

type _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */
// An interface definition for the [MTRGeneralDiagnosticsClusterPayloadTestResponseParams] class.
type IMTRGeneralDiagnosticsClusterPayloadTestResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */
	// properties:
	Payload() foundation.Data
	SetPayload(value foundation.Data)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass) Alloc() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass) New() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterPayloadTestResponseParams) Init() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterPayloadTestResponseParams) Autorelease() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterPayloadTestResponseParams creates a new MTRGeneralDiagnosticsClusterPayloadTestResponseParams instance.
func NewMTRGeneralDiagnosticsClusterPayloadTestResponseParams() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	return getMTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestResponseParams
type MTRGeneralDiagnosticsClusterPayloadTestResponseParams struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterPayloadTestResponseParamsFrom constructs a [MTRGeneralDiagnosticsClusterPayloadTestResponseParams] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterPayloadTestResponseParamsFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	return MTRGeneralDiagnosticsClusterPayloadTestResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */

// Initialize an MTRGeneralDiagnosticsClusterPayloadTestResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestResponseParams/init(responseValue:)
func NewMTRGeneralDiagnosticsClusterPayloadTestResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	instance := getMTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass().Alloc()
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRGeneralDiagnosticsClusterPayloadTestResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRGeneralDiagnosticsClusterPayloadTestResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterpayloadtestresponseparams/payload
func (m_ MTRGeneralDiagnosticsClusterPayloadTestResponseParams) Payload() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("payload"))
	return rv
}/* debug [instance_properties/getter]: payload */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterpayloadtestresponseparams/payload
func (m_ MTRGeneralDiagnosticsClusterPayloadTestResponseParams) SetPayload(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayload:"), value)
}/* debug [instance_properties/setter]: payload */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRGeneralDiagnosticsClusterPayloadTestResponseParams */


