// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAdministratorCommissioningClusterRevokeCommissioningParams] class.
var (
	MTRAdministratorCommissioningClusterRevokeCommissioningParamsClass     _MTRAdministratorCommissioningClusterRevokeCommissioningParamsClass
	MTRAdministratorCommissioningClusterRevokeCommissioningParamsClassOnce sync.Once
)

func getMTRAdministratorCommissioningClusterRevokeCommissioningParamsClass() _MTRAdministratorCommissioningClusterRevokeCommissioningParamsClass {
	MTRAdministratorCommissioningClusterRevokeCommissioningParamsClassOnce.Do(func() {
		MTRAdministratorCommissioningClusterRevokeCommissioningParamsClass = _MTRAdministratorCommissioningClusterRevokeCommissioningParamsClass{objc.GetClass("MTRAdministratorCommissioningClusterRevokeCommissioningParams")}
	})
	return MTRAdministratorCommissioningClusterRevokeCommissioningParamsClass
}

type _MTRAdministratorCommissioningClusterRevokeCommissioningParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAdministratorCommissioningClusterRevokeCommissioningParams] class.
type IMTRAdministratorCommissioningClusterRevokeCommissioningParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAdministratorCommissioningClusterRevokeCommissioningParams
type MTRAdministratorCommissioningClusterRevokeCommissioningParams struct {
	objectivec.Object
}

// MTRAdministratorCommissioningClusterRevokeCommissioningParamsFrom constructs a [MTRAdministratorCommissioningClusterRevokeCommissioningParams] from an unsafe.Pointer.
func MTRAdministratorCommissioningClusterRevokeCommissioningParamsFrom(ptr unsafe.Pointer) MTRAdministratorCommissioningClusterRevokeCommissioningParams {
	return MTRAdministratorCommissioningClusterRevokeCommissioningParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAdministratorCommissioningClusterRevokeCommissioningParamsClass) Alloc() MTRAdministratorCommissioningClusterRevokeCommissioningParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterRevokeCommissioningParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAdministratorCommissioningClusterRevokeCommissioningParamsClass) New() MTRAdministratorCommissioningClusterRevokeCommissioningParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterRevokeCommissioningParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAdministratorCommissioningClusterRevokeCommissioningParams) Init() MTRAdministratorCommissioningClusterRevokeCommissioningParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterRevokeCommissioningParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAdministratorCommissioningClusterRevokeCommissioningParams) Autorelease() MTRAdministratorCommissioningClusterRevokeCommissioningParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterRevokeCommissioningParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAdministratorCommissioningClusterRevokeCommissioningParams creates a new MTRAdministratorCommissioningClusterRevokeCommissioningParams instance.
func NewMTRAdministratorCommissioningClusterRevokeCommissioningParams() MTRAdministratorCommissioningClusterRevokeCommissioningParams {
	return getMTRAdministratorCommissioningClusterRevokeCommissioningParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusterrevokecommissioningparams/serversideprocessingtimeout
func (m_ MTRAdministratorCommissioningClusterRevokeCommissioningParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusterrevokecommissioningparams/serversideprocessingtimeout
func (m_ MTRAdministratorCommissioningClusterRevokeCommissioningParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusterrevokecommissioningparams/timedinvoketimeoutms
func (m_ MTRAdministratorCommissioningClusterRevokeCommissioningParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusterrevokecommissioningparams/timedinvoketimeoutms
func (m_ MTRAdministratorCommissioningClusterRevokeCommissioningParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



