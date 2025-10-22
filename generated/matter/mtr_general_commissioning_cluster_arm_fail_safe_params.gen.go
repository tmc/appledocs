// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralCommissioningClusterArmFailSafeParams] class.
var (
	MTRGeneralCommissioningClusterArmFailSafeParamsClass     _MTRGeneralCommissioningClusterArmFailSafeParamsClass
	MTRGeneralCommissioningClusterArmFailSafeParamsClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterArmFailSafeParamsClass() _MTRGeneralCommissioningClusterArmFailSafeParamsClass {
	MTRGeneralCommissioningClusterArmFailSafeParamsClassOnce.Do(func() {
		MTRGeneralCommissioningClusterArmFailSafeParamsClass = _MTRGeneralCommissioningClusterArmFailSafeParamsClass{objc.GetClass("MTRGeneralCommissioningClusterArmFailSafeParams")}
	})
	return MTRGeneralCommissioningClusterArmFailSafeParamsClass
}

type _MTRGeneralCommissioningClusterArmFailSafeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterArmFailSafeParams] class.
type IMTRGeneralCommissioningClusterArmFailSafeParams interface {
	objectivec.IObject
	Breadcrumb() foundation.Number
	SetBreadcrumb(value foundation.INumber)
	ExpiryLengthSeconds() foundation.Number
	SetExpiryLengthSeconds(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterArmFailSafeParams
type MTRGeneralCommissioningClusterArmFailSafeParams struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterArmFailSafeParamsFrom constructs a [MTRGeneralCommissioningClusterArmFailSafeParams] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterArmFailSafeParamsFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterArmFailSafeParams {
	return MTRGeneralCommissioningClusterArmFailSafeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterArmFailSafeParamsClass) Alloc() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterArmFailSafeParamsClass) New() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) Init() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) Autorelease() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterArmFailSafeParams creates a new MTRGeneralCommissioningClusterArmFailSafeParams instance.
func NewMTRGeneralCommissioningClusterArmFailSafeParams() MTRGeneralCommissioningClusterArmFailSafeParams {
	return getMTRGeneralCommissioningClusterArmFailSafeParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/breadcrumb
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) Breadcrumb() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("breadcrumb"))
	return rv
}


// SetBreadcrumb sets the value of the breadcrumb property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/breadcrumb
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) SetBreadcrumb(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/expirylengthseconds
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) ExpiryLengthSeconds() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("expiryLengthSeconds"))
	return rv
}


// SetExpiryLengthSeconds sets the value of the expiryLengthSeconds property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/expirylengthseconds
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) SetExpiryLengthSeconds(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExpiryLengthSeconds:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/serversideprocessingtimeout
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/serversideprocessingtimeout
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



