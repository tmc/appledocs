// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams] class.
var (
	MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass     _MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass
	MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClassOnce sync.Once
)

func getMTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass() _MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass {
	MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClassOnce.Do(func() {
		MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass = _MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass{objc.GetClass("MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams")}
	})
	return MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass
}

type _MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams] class.
type IMTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams
type MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams struct {
	objectivec.Object
}

// MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsFrom constructs a [MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams] from an unsafe.Pointer.
func MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsFrom(ptr unsafe.Pointer) MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams {
	return MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass) Alloc() MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass) New() MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams) Init() MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams) Autorelease() MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams creates a new MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams instance.
func NewMTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams() MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams {
	return getMTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropenbasiccommissioningwindowparams/commissioningtimeout
func (m_ MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams) CommissioningTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("commissioningTimeout"))
	return rv
}


// SetCommissioningTimeout sets the value of the commissioningTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropenbasiccommissioningwindowparams/commissioningtimeout
func (m_ MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams) SetCommissioningTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommissioningTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropenbasiccommissioningwindowparams/serversideprocessingtimeout
func (m_ MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropenbasiccommissioningwindowparams/serversideprocessingtimeout
func (m_ MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropenbasiccommissioningwindowparams/timedinvoketimeoutms
func (m_ MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropenbasiccommissioningwindowparams/timedinvoketimeoutms
func (m_ MTRAdministratorCommissioningClusterOpenBasicCommissioningWindowParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



