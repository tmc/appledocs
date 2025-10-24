// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalMeasurementClusterGetProfileInfoCommandParams] class.
var (
	MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass     _MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass
	MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClassOnce sync.Once
)

func getMTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass() _MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass {
	MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClassOnce.Do(func() {
		MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass = _MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass{objc.GetClass("MTRElectricalMeasurementClusterGetProfileInfoCommandParams")}
	})
	return MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass
}

type _MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalMeasurementClusterGetProfileInfoCommandParams] class.
type IMTRElectricalMeasurementClusterGetProfileInfoCommandParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetProfileInfoCommandParams
type MTRElectricalMeasurementClusterGetProfileInfoCommandParams struct {
	objectivec.Object
}

// MTRElectricalMeasurementClusterGetProfileInfoCommandParamsFrom constructs a [MTRElectricalMeasurementClusterGetProfileInfoCommandParams] from an unsafe.Pointer.
func MTRElectricalMeasurementClusterGetProfileInfoCommandParamsFrom(ptr unsafe.Pointer) MTRElectricalMeasurementClusterGetProfileInfoCommandParams {
	return MTRElectricalMeasurementClusterGetProfileInfoCommandParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass) Alloc() MTRElectricalMeasurementClusterGetProfileInfoCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetProfileInfoCommandParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass) New() MTRElectricalMeasurementClusterGetProfileInfoCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetProfileInfoCommandParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalMeasurementClusterGetProfileInfoCommandParams) Init() MTRElectricalMeasurementClusterGetProfileInfoCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetProfileInfoCommandParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalMeasurementClusterGetProfileInfoCommandParams) Autorelease() MTRElectricalMeasurementClusterGetProfileInfoCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetProfileInfoCommandParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalMeasurementClusterGetProfileInfoCommandParams creates a new MTRElectricalMeasurementClusterGetProfileInfoCommandParams instance.
func NewMTRElectricalMeasurementClusterGetProfileInfoCommandParams() MTRElectricalMeasurementClusterGetProfileInfoCommandParams {
	return getMTRElectricalMeasurementClusterGetProfileInfoCommandParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinfocommandparams/serversideprocessingtimeout
func (m_ MTRElectricalMeasurementClusterGetProfileInfoCommandParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinfocommandparams/serversideprocessingtimeout
func (m_ MTRElectricalMeasurementClusterGetProfileInfoCommandParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinfocommandparams/timedinvoketimeoutms
func (m_ MTRElectricalMeasurementClusterGetProfileInfoCommandParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinfocommandparams/timedinvoketimeoutms
func (m_ MTRElectricalMeasurementClusterGetProfileInfoCommandParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
