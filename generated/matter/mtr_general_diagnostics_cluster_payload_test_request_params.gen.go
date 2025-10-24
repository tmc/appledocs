// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRGeneralDiagnosticsClusterPayloadTestRequestParams */


/* debug [class_header]: Header for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */
// The class instance for the [MTRGeneralDiagnosticsClusterPayloadTestRequestParams] class.
var (
	MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass     _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass
	MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass() _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass {
	MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass = _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass{objc.GetClass("MTRGeneralDiagnosticsClusterPayloadTestRequestParams")}
	})
	return MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass
}

type _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */
// An interface definition for the [MTRGeneralDiagnosticsClusterPayloadTestRequestParams] class.
type IMTRGeneralDiagnosticsClusterPayloadTestRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */
	// properties:
	EnableKey() objc.IObject /* cross-framework: NSData */
	SetEnableKey(value objc.IObject /* cross-framework: NSData */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Count() objc.IObject /* cross-framework: NSNumber */
	SetCount(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass) Alloc() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass) New() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) Init() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) Autorelease() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterPayloadTestRequestParams creates a new MTRGeneralDiagnosticsClusterPayloadTestRequestParams instance.
func NewMTRGeneralDiagnosticsClusterPayloadTestRequestParams() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	return getMTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams
type MTRGeneralDiagnosticsClusterPayloadTestRequestParams struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterPayloadTestRequestParamsFrom constructs a [MTRGeneralDiagnosticsClusterPayloadTestRequestParams] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterPayloadTestRequestParamsFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	return MTRGeneralDiagnosticsClusterPayloadTestRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRGeneralDiagnosticsClusterPayloadTestRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRGeneralDiagnosticsClusterPayloadTestRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/enableKey
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) EnableKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("enableKey"))
	return rv
}/* debug [instance_properties/getter]: enableKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/enableKey
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetEnableKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnableKey:"), value)
}/* debug [instance_properties/setter]: enableKey */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/timedInvokeTimeoutMs
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/timedInvokeTimeoutMs
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterpayloadtestrequestparams/count
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) Count() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterpayloadtestrequestparams/count
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCount:"), value)
}/* debug [instance_properties/setter]: count */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterpayloadtestrequestparams/serversideprocessingtimeout
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterpayloadtestrequestparams/serversideprocessingtimeout
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterpayloadtestrequestparams/value
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterpayloadtestrequestparams/value
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRGeneralDiagnosticsClusterPayloadTestRequestParams */



