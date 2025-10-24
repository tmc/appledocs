// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */


/* debug [class_header]: Header for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */
// The class instance for the [MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams] class.
var (
	MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass     _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass
	MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass() _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass {
	MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass = _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass{objc.GetClass("MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams")}
	})
	return MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass
}

type _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */
// An interface definition for the [MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams] class.
type IMTRGeneralDiagnosticsClusterTimeSnapshotResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */
	// properties:
	PosixTimeMs() objc.IObject /* cross-framework: NSNumber */
	SetPosixTimeMs(value objc.IObject /* cross-framework: NSNumber */)
	SystemTimeMs() objc.IObject /* cross-framework: NSNumber */
	SetSystemTimeMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass) Alloc() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass) New() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) Init() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) Autorelease() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterTimeSnapshotResponseParams creates a new MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams instance.
func NewMTRGeneralDiagnosticsClusterTimeSnapshotResponseParams() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	return getMTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams
type MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsFrom constructs a [MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	return MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */

// Initialize an MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams/init(responseValue:)
func NewMTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	instance := getMTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass().Alloc()
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertimesnapshotresponseparams/posixtimems
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) PosixTimeMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("posixTimeMs"))
	return rv
}/* debug [instance_properties/getter]: posixTimeMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertimesnapshotresponseparams/posixtimems
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) SetPosixTimeMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPosixTimeMs:"), value)
}/* debug [instance_properties/setter]: posixTimeMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertimesnapshotresponseparams/systemtimems
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) SystemTimeMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("systemTimeMs"))
	return rv
}/* debug [instance_properties/getter]: systemTimeMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertimesnapshotresponseparams/systemtimems
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) SetSystemTimeMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemTimeMs:"), value)
}/* debug [instance_properties/setter]: systemTimeMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams */


