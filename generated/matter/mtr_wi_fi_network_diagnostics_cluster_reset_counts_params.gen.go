// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWiFiNetworkDiagnosticsClusterResetCountsParams */


/* debug [class_header]: Header for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */
// The class instance for the [MTRWiFiNetworkDiagnosticsClusterResetCountsParams] class.
var (
	MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass     _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass
	MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClassOnce sync.Once
)

func getMTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass() _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass {
	MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClassOnce.Do(func() {
		MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass = _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass{objc.GetClass("MTRWiFiNetworkDiagnosticsClusterResetCountsParams")}
	})
	return MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass
}

type _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */
// An interface definition for the [MTRWiFiNetworkDiagnosticsClusterResetCountsParams] class.
type IMTRWiFiNetworkDiagnosticsClusterResetCountsParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass) Alloc() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass) New() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWiFiNetworkDiagnosticsClusterResetCountsParams) Init() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWiFiNetworkDiagnosticsClusterResetCountsParams) Autorelease() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWiFiNetworkDiagnosticsClusterResetCountsParams creates a new MTRWiFiNetworkDiagnosticsClusterResetCountsParams instance.
func NewMTRWiFiNetworkDiagnosticsClusterResetCountsParams() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	return getMTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterResetCountsParams
type MTRWiFiNetworkDiagnosticsClusterResetCountsParams struct {
	objectivec.Object
}

// MTRWiFiNetworkDiagnosticsClusterResetCountsParamsFrom constructs a [MTRWiFiNetworkDiagnosticsClusterResetCountsParams] from an unsafe.Pointer.
func MTRWiFiNetworkDiagnosticsClusterResetCountsParamsFrom(ptr unsafe.Pointer) MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	return MTRWiFiNetworkDiagnosticsClusterResetCountsParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWiFiNetworkDiagnosticsClusterResetCountsParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWiFiNetworkDiagnosticsClusterResetCountsParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterResetCountsParams/serverSideProcessingTimeout
func (m_ MTRWiFiNetworkDiagnosticsClusterResetCountsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterResetCountsParams/serverSideProcessingTimeout
func (m_ MTRWiFiNetworkDiagnosticsClusterResetCountsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterResetCountsParams/timedInvokeTimeoutMs
func (m_ MTRWiFiNetworkDiagnosticsClusterResetCountsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterResetCountsParams/timedInvokeTimeoutMs
func (m_ MTRWiFiNetworkDiagnosticsClusterResetCountsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWiFiNetworkDiagnosticsClusterResetCountsParams */



