// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterSetpointRaiseLowerParams] class.
var (
	MTRThermostatClusterSetpointRaiseLowerParamsClass     _MTRThermostatClusterSetpointRaiseLowerParamsClass
	MTRThermostatClusterSetpointRaiseLowerParamsClassOnce sync.Once
)

func getMTRThermostatClusterSetpointRaiseLowerParamsClass() _MTRThermostatClusterSetpointRaiseLowerParamsClass {
	MTRThermostatClusterSetpointRaiseLowerParamsClassOnce.Do(func() {
		MTRThermostatClusterSetpointRaiseLowerParamsClass = _MTRThermostatClusterSetpointRaiseLowerParamsClass{objc.GetClass("MTRThermostatClusterSetpointRaiseLowerParams")}
	})
	return MTRThermostatClusterSetpointRaiseLowerParamsClass
}

type _MTRThermostatClusterSetpointRaiseLowerParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterSetpointRaiseLowerParams] class.
type IMTRThermostatClusterSetpointRaiseLowerParams interface {
	objectivec.IObject
	// properties:
	Amount() objc.IObject /* cross-framework: NSNumber */
	SetAmount(value objc.IObject /* cross-framework: NSNumber */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams
type MTRThermostatClusterSetpointRaiseLowerParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetpointRaiseLowerParamsFrom constructs a [MTRThermostatClusterSetpointRaiseLowerParams] from an unsafe.Pointer.
func MTRThermostatClusterSetpointRaiseLowerParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetpointRaiseLowerParams {
	return MTRThermostatClusterSetpointRaiseLowerParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetpointRaiseLowerParamsClass) Alloc() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterSetpointRaiseLowerParamsClass) New() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Init() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Autorelease() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterSetpointRaiseLowerParams creates a new MTRThermostatClusterSetpointRaiseLowerParams instance.
func NewMTRThermostatClusterSetpointRaiseLowerParams() MTRThermostatClusterSetpointRaiseLowerParams {
	return getMTRThermostatClusterSetpointRaiseLowerParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/amount
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Amount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("amount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/amount
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetAmount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAmount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/mode
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/mode
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



