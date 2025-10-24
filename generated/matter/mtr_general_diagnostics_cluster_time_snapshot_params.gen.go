// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRGeneralDiagnosticsClusterTimeSnapshotParams */


/* debug [class_header]: Header for MTRGeneralDiagnosticsClusterTimeSnapshotParams */
// The class instance for the [MTRGeneralDiagnosticsClusterTimeSnapshotParams] class.
var (
	MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass     _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass
	MTRGeneralDiagnosticsClusterTimeSnapshotParamsClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterTimeSnapshotParamsClass() _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass {
	MTRGeneralDiagnosticsClusterTimeSnapshotParamsClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass = _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass{objc.GetClass("MTRGeneralDiagnosticsClusterTimeSnapshotParams")}
	})
	return MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass
}

type _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRGeneralDiagnosticsClusterTimeSnapshotParams */
// An interface definition for the [MTRGeneralDiagnosticsClusterTimeSnapshotParams] class.
type IMTRGeneralDiagnosticsClusterTimeSnapshotParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRGeneralDiagnosticsClusterTimeSnapshotParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRGeneralDiagnosticsClusterTimeSnapshotParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRGeneralDiagnosticsClusterTimeSnapshotParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass) Alloc() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass) New() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) Init() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) Autorelease() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterTimeSnapshotParams creates a new MTRGeneralDiagnosticsClusterTimeSnapshotParams instance.
func NewMTRGeneralDiagnosticsClusterTimeSnapshotParams() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	return getMTRGeneralDiagnosticsClusterTimeSnapshotParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRGeneralDiagnosticsClusterTimeSnapshotParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotParams
type MTRGeneralDiagnosticsClusterTimeSnapshotParams struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterTimeSnapshotParamsFrom constructs a [MTRGeneralDiagnosticsClusterTimeSnapshotParams] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterTimeSnapshotParamsFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	return MTRGeneralDiagnosticsClusterTimeSnapshotParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRGeneralDiagnosticsClusterTimeSnapshotParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRGeneralDiagnosticsClusterTimeSnapshotParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRGeneralDiagnosticsClusterTimeSnapshotParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRGeneralDiagnosticsClusterTimeSnapshotParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRGeneralDiagnosticsClusterTimeSnapshotParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotParams/serverSideProcessingTimeout
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotParams/serverSideProcessingTimeout
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertimesnapshotparams/timedinvoketimeoutms
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertimesnapshotparams/timedinvoketimeoutms
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRGeneralDiagnosticsClusterTimeSnapshotParams */



