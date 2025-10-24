// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDiagnosticsClusterResetCountsParams */


/* debug [class_header]: Header for MTRThreadNetworkDiagnosticsClusterResetCountsParams */
// The class instance for the [MTRThreadNetworkDiagnosticsClusterResetCountsParams] class.
var (
	MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass     _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass
	MTRThreadNetworkDiagnosticsClusterResetCountsParamsClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterResetCountsParamsClass() _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass {
	MTRThreadNetworkDiagnosticsClusterResetCountsParamsClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass = _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterResetCountsParams")}
	})
	return MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass
}

type _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDiagnosticsClusterResetCountsParams */
// An interface definition for the [MTRThreadNetworkDiagnosticsClusterResetCountsParams] class.
type IMTRThreadNetworkDiagnosticsClusterResetCountsParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDiagnosticsClusterResetCountsParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDiagnosticsClusterResetCountsParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDiagnosticsClusterResetCountsParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass) Alloc() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass) New() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterResetCountsParams) Init() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterResetCountsParams) Autorelease() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterResetCountsParams creates a new MTRThreadNetworkDiagnosticsClusterResetCountsParams instance.
func NewMTRThreadNetworkDiagnosticsClusterResetCountsParams() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	return getMTRThreadNetworkDiagnosticsClusterResetCountsParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDiagnosticsClusterResetCountsParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterResetCountsParams
type MTRThreadNetworkDiagnosticsClusterResetCountsParams struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterResetCountsParamsFrom constructs a [MTRThreadNetworkDiagnosticsClusterResetCountsParams] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterResetCountsParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	return MTRThreadNetworkDiagnosticsClusterResetCountsParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDiagnosticsClusterResetCountsParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDiagnosticsClusterResetCountsParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDiagnosticsClusterResetCountsParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDiagnosticsClusterResetCountsParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDiagnosticsClusterResetCountsParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterResetCountsParams/serverSideProcessingTimeout
func (m_ MTRThreadNetworkDiagnosticsClusterResetCountsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterResetCountsParams/serverSideProcessingTimeout
func (m_ MTRThreadNetworkDiagnosticsClusterResetCountsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterResetCountsParams/timedInvokeTimeoutMs
func (m_ MTRThreadNetworkDiagnosticsClusterResetCountsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterResetCountsParams/timedInvokeTimeoutMs
func (m_ MTRThreadNetworkDiagnosticsClusterResetCountsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDiagnosticsClusterResetCountsParams */



