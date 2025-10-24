// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterSetActivePresetRequestParams */


/* debug [class_header]: Header for MTRThermostatClusterSetActivePresetRequestParams */
// The class instance for the [MTRThermostatClusterSetActivePresetRequestParams] class.
var (
	MTRThermostatClusterSetActivePresetRequestParamsClass     _MTRThermostatClusterSetActivePresetRequestParamsClass
	MTRThermostatClusterSetActivePresetRequestParamsClassOnce sync.Once
)

func getMTRThermostatClusterSetActivePresetRequestParamsClass() _MTRThermostatClusterSetActivePresetRequestParamsClass {
	MTRThermostatClusterSetActivePresetRequestParamsClassOnce.Do(func() {
		MTRThermostatClusterSetActivePresetRequestParamsClass = _MTRThermostatClusterSetActivePresetRequestParamsClass{objc.GetClass("MTRThermostatClusterSetActivePresetRequestParams")}
	})
	return MTRThermostatClusterSetActivePresetRequestParamsClass
}

type _MTRThermostatClusterSetActivePresetRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterSetActivePresetRequestParams */
// An interface definition for the [MTRThermostatClusterSetActivePresetRequestParams] class.
type IMTRThermostatClusterSetActivePresetRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterSetActivePresetRequestParams */
	// properties:
	PresetHandle() objc.IObject /* cross-framework: NSData */
	SetPresetHandle(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterSetActivePresetRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterSetActivePresetRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetActivePresetRequestParamsClass) Alloc() MTRThermostatClusterSetActivePresetRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActivePresetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterSetActivePresetRequestParamsClass) New() MTRThermostatClusterSetActivePresetRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActivePresetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterSetActivePresetRequestParams) Init() MTRThermostatClusterSetActivePresetRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActivePresetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterSetActivePresetRequestParams) Autorelease() MTRThermostatClusterSetActivePresetRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActivePresetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterSetActivePresetRequestParams creates a new MTRThermostatClusterSetActivePresetRequestParams instance.
func NewMTRThermostatClusterSetActivePresetRequestParams() MTRThermostatClusterSetActivePresetRequestParams {
	return getMTRThermostatClusterSetActivePresetRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterSetActivePresetRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams
type MTRThermostatClusterSetActivePresetRequestParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetActivePresetRequestParamsFrom constructs a [MTRThermostatClusterSetActivePresetRequestParams] from an unsafe.Pointer.
func MTRThermostatClusterSetActivePresetRequestParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetActivePresetRequestParams {
	return MTRThermostatClusterSetActivePresetRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterSetActivePresetRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterSetActivePresetRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterSetActivePresetRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterSetActivePresetRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterSetActivePresetRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/presetHandle
func (m_ MTRThermostatClusterSetActivePresetRequestParams) PresetHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("presetHandle"))
	return rv
}/* debug [instance_properties/getter]: presetHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/presetHandle
func (m_ MTRThermostatClusterSetActivePresetRequestParams) SetPresetHandle(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}/* debug [instance_properties/setter]: presetHandle */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetActivePresetRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetActivePresetRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetactivepresetrequestparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetActivePresetRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetactivepresetrequestparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetActivePresetRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterSetActivePresetRequestParams */



