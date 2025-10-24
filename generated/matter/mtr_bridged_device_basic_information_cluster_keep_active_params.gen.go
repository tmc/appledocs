// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRBridgedDeviceBasicInformationClusterKeepActiveParams */


/* debug [class_header]: Header for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */
// The class instance for the [MTRBridgedDeviceBasicInformationClusterKeepActiveParams] class.
var (
	MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass     _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass
	MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClassOnce sync.Once
)

func getMTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass() _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass {
	MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClassOnce.Do(func() {
		MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass = _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass{objc.GetClass("MTRBridgedDeviceBasicInformationClusterKeepActiveParams")}
	})
	return MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass
}

type _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */
// An interface definition for the [MTRBridgedDeviceBasicInformationClusterKeepActiveParams] class.
type IMTRBridgedDeviceBasicInformationClusterKeepActiveParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StayActiveDuration() objc.IObject /* cross-framework: NSNumber */
	SetStayActiveDuration(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass) Alloc() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterKeepActiveParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass) New() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterKeepActiveParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) Init() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterKeepActiveParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) Autorelease() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterKeepActiveParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicInformationClusterKeepActiveParams creates a new MTRBridgedDeviceBasicInformationClusterKeepActiveParams instance.
func NewMTRBridgedDeviceBasicInformationClusterKeepActiveParams() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	return getMTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams
type MTRBridgedDeviceBasicInformationClusterKeepActiveParams struct {
	objectivec.Object
}

// MTRBridgedDeviceBasicInformationClusterKeepActiveParamsFrom constructs a [MTRBridgedDeviceBasicInformationClusterKeepActiveParams] from an unsafe.Pointer.
func MTRBridgedDeviceBasicInformationClusterKeepActiveParamsFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	return MTRBridgedDeviceBasicInformationClusterKeepActiveParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBridgedDeviceBasicInformationClusterKeepActiveParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBridgedDeviceBasicInformationClusterKeepActiveParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/serverSideProcessingTimeout
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/serverSideProcessingTimeout
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterkeepactiveparams/stayactiveduration
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) StayActiveDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stayActiveDuration"))
	return rv
}/* debug [instance_properties/getter]: stayActiveDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterkeepactiveparams/stayactiveduration
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetStayActiveDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStayActiveDuration:"), value)
}/* debug [instance_properties/setter]: stayActiveDuration */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterkeepactiveparams/timedinvoketimeoutms
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterkeepactiveparams/timedinvoketimeoutms
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterkeepactiveparams/timeoutms
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) TimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterkeepactiveparams/timeoutms
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBridgedDeviceBasicInformationClusterKeepActiveParams */



