// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRSoftwareDiagnosticsClusterResetWatermarksParams */


/* debug [class_header]: Header for MTRSoftwareDiagnosticsClusterResetWatermarksParams */
// The class instance for the [MTRSoftwareDiagnosticsClusterResetWatermarksParams] class.
var (
	MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass     _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass
	MTRSoftwareDiagnosticsClusterResetWatermarksParamsClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterResetWatermarksParamsClass() _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass {
	MTRSoftwareDiagnosticsClusterResetWatermarksParamsClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass = _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass{objc.GetClass("MTRSoftwareDiagnosticsClusterResetWatermarksParams")}
	})
	return MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass
}

type _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSoftwareDiagnosticsClusterResetWatermarksParams */
// An interface definition for the [MTRSoftwareDiagnosticsClusterResetWatermarksParams] class.
type IMTRSoftwareDiagnosticsClusterResetWatermarksParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRSoftwareDiagnosticsClusterResetWatermarksParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSoftwareDiagnosticsClusterResetWatermarksParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSoftwareDiagnosticsClusterResetWatermarksParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass) Alloc() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterResetWatermarksParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass) New() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterResetWatermarksParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) Init() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterResetWatermarksParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) Autorelease() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterResetWatermarksParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterResetWatermarksParams creates a new MTRSoftwareDiagnosticsClusterResetWatermarksParams instance.
func NewMTRSoftwareDiagnosticsClusterResetWatermarksParams() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	return getMTRSoftwareDiagnosticsClusterResetWatermarksParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSoftwareDiagnosticsClusterResetWatermarksParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterResetWatermarksParams
type MTRSoftwareDiagnosticsClusterResetWatermarksParams struct {
	objectivec.Object
}

// MTRSoftwareDiagnosticsClusterResetWatermarksParamsFrom constructs a [MTRSoftwareDiagnosticsClusterResetWatermarksParams] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterResetWatermarksParamsFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	return MTRSoftwareDiagnosticsClusterResetWatermarksParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSoftwareDiagnosticsClusterResetWatermarksParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSoftwareDiagnosticsClusterResetWatermarksParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSoftwareDiagnosticsClusterResetWatermarksParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSoftwareDiagnosticsClusterResetWatermarksParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSoftwareDiagnosticsClusterResetWatermarksParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterResetWatermarksParams/serverSideProcessingTimeout
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterResetWatermarksParams/serverSideProcessingTimeout
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterResetWatermarksParams/timedInvokeTimeoutMs
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterResetWatermarksParams/timedInvokeTimeoutMs
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSoftwareDiagnosticsClusterResetWatermarksParams */



