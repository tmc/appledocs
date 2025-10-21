// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAdministratorCommissioningClusterOpenCommissioningWindowParams] class.
var (
	MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass     _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass
	MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClassOnce sync.Once
)

func getMTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass() _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass {
	MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClassOnce.Do(func() {
		MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass = _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass{objc.GetClass("MTRAdministratorCommissioningClusterOpenCommissioningWindowParams")}
	})
	return MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass
}

type _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAdministratorCommissioningClusterOpenCommissioningWindowParams] class.
type IMTRAdministratorCommissioningClusterOpenCommissioningWindowParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAdministratorCommissioningClusterOpenCommissioningWindowParams
type MTRAdministratorCommissioningClusterOpenCommissioningWindowParams struct {
	objectivec.Object
}

// MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsFrom constructs a [MTRAdministratorCommissioningClusterOpenCommissioningWindowParams] from an unsafe.Pointer.
func MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsFrom(ptr unsafe.Pointer) MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	return MTRAdministratorCommissioningClusterOpenCommissioningWindowParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass) Alloc() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass) New() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Init() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Autorelease() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAdministratorCommissioningClusterOpenCommissioningWindowParams creates a new MTRAdministratorCommissioningClusterOpenCommissioningWindowParams instance.
func NewMTRAdministratorCommissioningClusterOpenCommissioningWindowParams() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	return getMTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/discriminator
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Discriminator() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("discriminator"))
	return rv
}


// SetDiscriminator sets the value of the discriminator property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/discriminator
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetDiscriminator(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDiscriminator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/iterations
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Iterations() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("iterations"))
	return rv
}


// SetIterations sets the value of the iterations property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/iterations
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetIterations(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIterations:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/timedinvoketimeoutms
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/timedinvoketimeoutms
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/pakeverifier
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) PakeVerifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pakeVerifier"))
	return rv
}


// SetPakeVerifier sets the value of the pakeVerifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/pakeverifier
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetPakeVerifier(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPakeVerifier:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/salt
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Salt() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("salt"))
	return rv
}


// SetSalt sets the value of the salt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/salt
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetSalt(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSalt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/commissioningtimeout
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) CommissioningTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("commissioningTimeout"))
	return rv
}


// SetCommissioningTimeout sets the value of the commissioningTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/commissioningtimeout
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetCommissioningTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommissioningTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/pakepasscodeverifier
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) PakePasscodeVerifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pakePasscodeVerifier"))
	return rv
}


// SetPakePasscodeVerifier sets the value of the pakePasscodeVerifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/pakepasscodeverifier
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetPakePasscodeVerifier(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPakePasscodeVerifier:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/serversideprocessingtimeout
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/serversideprocessingtimeout
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



