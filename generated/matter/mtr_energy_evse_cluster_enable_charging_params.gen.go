// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterEnableChargingParams */


/* debug [class_header]: Header for MTREnergyEVSEClusterEnableChargingParams */
// The class instance for the [MTREnergyEVSEClusterEnableChargingParams] class.
var (
	MTREnergyEVSEClusterEnableChargingParamsClass     _MTREnergyEVSEClusterEnableChargingParamsClass
	MTREnergyEVSEClusterEnableChargingParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterEnableChargingParamsClass() _MTREnergyEVSEClusterEnableChargingParamsClass {
	MTREnergyEVSEClusterEnableChargingParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterEnableChargingParamsClass = _MTREnergyEVSEClusterEnableChargingParamsClass{objc.GetClass("MTREnergyEVSEClusterEnableChargingParams")}
	})
	return MTREnergyEVSEClusterEnableChargingParamsClass
}

type _MTREnergyEVSEClusterEnableChargingParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterEnableChargingParams */
// An interface definition for the [MTREnergyEVSEClusterEnableChargingParams] class.
type IMTREnergyEVSEClusterEnableChargingParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterEnableChargingParams */
	// properties:
	ChargingEnabledUntil() objc.IObject /* cross-framework: NSNumber */
	SetChargingEnabledUntil(value objc.IObject /* cross-framework: NSNumber */)
	MaximumChargeCurrent() objc.IObject /* cross-framework: NSNumber */
	SetMaximumChargeCurrent(value objc.IObject /* cross-framework: NSNumber */)
	MinimumChargeCurrent() objc.IObject /* cross-framework: NSNumber */
	SetMinimumChargeCurrent(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterEnableChargingParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterEnableChargingParams */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEnableChargingParamsClass) Alloc() MTREnergyEVSEClusterEnableChargingParams {
	rv := objc.Send[MTREnergyEVSEClusterEnableChargingParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterEnableChargingParamsClass) New() MTREnergyEVSEClusterEnableChargingParams {
	rv := objc.Send[MTREnergyEVSEClusterEnableChargingParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEnableChargingParams) Init() MTREnergyEVSEClusterEnableChargingParams {
	rv := objc.Send[MTREnergyEVSEClusterEnableChargingParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEnableChargingParams) Autorelease() MTREnergyEVSEClusterEnableChargingParams {
	rv := objc.Send[MTREnergyEVSEClusterEnableChargingParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEnableChargingParams creates a new MTREnergyEVSEClusterEnableChargingParams instance.
func NewMTREnergyEVSEClusterEnableChargingParams() MTREnergyEVSEClusterEnableChargingParams {
	return getMTREnergyEVSEClusterEnableChargingParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterEnableChargingParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams
type MTREnergyEVSEClusterEnableChargingParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEnableChargingParamsFrom constructs a [MTREnergyEVSEClusterEnableChargingParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterEnableChargingParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEnableChargingParams {
	return MTREnergyEVSEClusterEnableChargingParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterEnableChargingParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterEnableChargingParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterEnableChargingParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterEnableChargingParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterEnableChargingParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/chargingEnabledUntil
func (m_ MTREnergyEVSEClusterEnableChargingParams) ChargingEnabledUntil() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("chargingEnabledUntil"))
	return rv
}/* debug [instance_properties/getter]: chargingEnabledUntil */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/chargingEnabledUntil
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetChargingEnabledUntil(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChargingEnabledUntil:"), value)
}/* debug [instance_properties/setter]: chargingEnabledUntil */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenablechargingparams/maximumchargecurrent
func (m_ MTREnergyEVSEClusterEnableChargingParams) MaximumChargeCurrent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maximumChargeCurrent"))
	return rv
}/* debug [instance_properties/getter]: maximumChargeCurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenablechargingparams/maximumchargecurrent
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetMaximumChargeCurrent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumChargeCurrent:"), value)
}/* debug [instance_properties/setter]: maximumChargeCurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenablechargingparams/minimumchargecurrent
func (m_ MTREnergyEVSEClusterEnableChargingParams) MinimumChargeCurrent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minimumChargeCurrent"))
	return rv
}/* debug [instance_properties/getter]: minimumChargeCurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenablechargingparams/minimumchargecurrent
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetMinimumChargeCurrent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumChargeCurrent:"), value)
}/* debug [instance_properties/setter]: minimumChargeCurrent */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenablechargingparams/serversideprocessingtimeout
func (m_ MTREnergyEVSEClusterEnableChargingParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenablechargingparams/serversideprocessingtimeout
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenablechargingparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterEnableChargingParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenablechargingparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterEnableChargingParams */



