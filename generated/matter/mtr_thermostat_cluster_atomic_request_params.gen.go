// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterAtomicRequestParams */


/* debug [class_header]: Header for MTRThermostatClusterAtomicRequestParams */
// The class instance for the [MTRThermostatClusterAtomicRequestParams] class.
var (
	MTRThermostatClusterAtomicRequestParamsClass     _MTRThermostatClusterAtomicRequestParamsClass
	MTRThermostatClusterAtomicRequestParamsClassOnce sync.Once
)

func getMTRThermostatClusterAtomicRequestParamsClass() _MTRThermostatClusterAtomicRequestParamsClass {
	MTRThermostatClusterAtomicRequestParamsClassOnce.Do(func() {
		MTRThermostatClusterAtomicRequestParamsClass = _MTRThermostatClusterAtomicRequestParamsClass{objc.GetClass("MTRThermostatClusterAtomicRequestParams")}
	})
	return MTRThermostatClusterAtomicRequestParamsClass
}

type _MTRThermostatClusterAtomicRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterAtomicRequestParams */
// An interface definition for the [MTRThermostatClusterAtomicRequestParams] class.
type IMTRThermostatClusterAtomicRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterAtomicRequestParams */
	// properties:
	AttributeRequests() objc.IObject /* cross-framework: NSArray */
	SetAttributeRequests(value objc.IObject /* cross-framework: NSArray */)
	RequestType() objc.IObject /* cross-framework: NSNumber */
	SetRequestType(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Timeout() objc.IObject /* cross-framework: NSNumber */
	SetTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterAtomicRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterAtomicRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterAtomicRequestParamsClass) Alloc() MTRThermostatClusterAtomicRequestParams {
	rv := objc.Send[MTRThermostatClusterAtomicRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterAtomicRequestParamsClass) New() MTRThermostatClusterAtomicRequestParams {
	rv := objc.Send[MTRThermostatClusterAtomicRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterAtomicRequestParams) Init() MTRThermostatClusterAtomicRequestParams {
	rv := objc.Send[MTRThermostatClusterAtomicRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterAtomicRequestParams) Autorelease() MTRThermostatClusterAtomicRequestParams {
	rv := objc.Send[MTRThermostatClusterAtomicRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterAtomicRequestParams creates a new MTRThermostatClusterAtomicRequestParams instance.
func NewMTRThermostatClusterAtomicRequestParams() MTRThermostatClusterAtomicRequestParams {
	return getMTRThermostatClusterAtomicRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterAtomicRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams
type MTRThermostatClusterAtomicRequestParams struct {
	objectivec.Object
}

// MTRThermostatClusterAtomicRequestParamsFrom constructs a [MTRThermostatClusterAtomicRequestParams] from an unsafe.Pointer.
func MTRThermostatClusterAtomicRequestParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterAtomicRequestParams {
	return MTRThermostatClusterAtomicRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterAtomicRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterAtomicRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterAtomicRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterAtomicRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterAtomicRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/attributeRequests
func (m_ MTRThermostatClusterAtomicRequestParams) AttributeRequests() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("attributeRequests"))
	return rv
}/* debug [instance_properties/getter]: attributeRequests */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/attributeRequests
func (m_ MTRThermostatClusterAtomicRequestParams) SetAttributeRequests(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeRequests:"), value)
}/* debug [instance_properties/setter]: attributeRequests */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicrequestparams/requesttype
func (m_ MTRThermostatClusterAtomicRequestParams) RequestType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestType"))
	return rv
}/* debug [instance_properties/getter]: requestType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicrequestparams/requesttype
func (m_ MTRThermostatClusterAtomicRequestParams) SetRequestType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestType:"), value)
}/* debug [instance_properties/setter]: requestType */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicrequestparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterAtomicRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicrequestparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterAtomicRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicrequestparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterAtomicRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicrequestparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterAtomicRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicrequestparams/timeout
func (m_ MTRThermostatClusterAtomicRequestParams) Timeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeout"))
	return rv
}/* debug [instance_properties/getter]: timeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicrequestparams/timeout
func (m_ MTRThermostatClusterAtomicRequestParams) SetTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeout:"), value)
}/* debug [instance_properties/setter]: timeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterAtomicRequestParams */



