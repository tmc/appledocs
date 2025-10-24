// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterStartDiagnosticsParams */


/* debug [class_header]: Header for MTREnergyEVSEClusterStartDiagnosticsParams */
// The class instance for the [MTREnergyEVSEClusterStartDiagnosticsParams] class.
var (
	MTREnergyEVSEClusterStartDiagnosticsParamsClass     _MTREnergyEVSEClusterStartDiagnosticsParamsClass
	MTREnergyEVSEClusterStartDiagnosticsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterStartDiagnosticsParamsClass() _MTREnergyEVSEClusterStartDiagnosticsParamsClass {
	MTREnergyEVSEClusterStartDiagnosticsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterStartDiagnosticsParamsClass = _MTREnergyEVSEClusterStartDiagnosticsParamsClass{objc.GetClass("MTREnergyEVSEClusterStartDiagnosticsParams")}
	})
	return MTREnergyEVSEClusterStartDiagnosticsParamsClass
}

type _MTREnergyEVSEClusterStartDiagnosticsParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterStartDiagnosticsParams */
// An interface definition for the [MTREnergyEVSEClusterStartDiagnosticsParams] class.
type IMTREnergyEVSEClusterStartDiagnosticsParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterStartDiagnosticsParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterStartDiagnosticsParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterStartDiagnosticsParams */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterStartDiagnosticsParamsClass) Alloc() MTREnergyEVSEClusterStartDiagnosticsParams {
	rv := objc.Send[MTREnergyEVSEClusterStartDiagnosticsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterStartDiagnosticsParamsClass) New() MTREnergyEVSEClusterStartDiagnosticsParams {
	rv := objc.Send[MTREnergyEVSEClusterStartDiagnosticsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) Init() MTREnergyEVSEClusterStartDiagnosticsParams {
	rv := objc.Send[MTREnergyEVSEClusterStartDiagnosticsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) Autorelease() MTREnergyEVSEClusterStartDiagnosticsParams {
	rv := objc.Send[MTREnergyEVSEClusterStartDiagnosticsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterStartDiagnosticsParams creates a new MTREnergyEVSEClusterStartDiagnosticsParams instance.
func NewMTREnergyEVSEClusterStartDiagnosticsParams() MTREnergyEVSEClusterStartDiagnosticsParams {
	return getMTREnergyEVSEClusterStartDiagnosticsParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterStartDiagnosticsParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterStartDiagnosticsParams
type MTREnergyEVSEClusterStartDiagnosticsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterStartDiagnosticsParamsFrom constructs a [MTREnergyEVSEClusterStartDiagnosticsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterStartDiagnosticsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterStartDiagnosticsParams {
	return MTREnergyEVSEClusterStartDiagnosticsParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterStartDiagnosticsParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterStartDiagnosticsParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterStartDiagnosticsParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterStartDiagnosticsParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterStartDiagnosticsParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterStartDiagnosticsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterStartDiagnosticsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterstartdiagnosticsparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterstartdiagnosticsparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterStartDiagnosticsParams */



