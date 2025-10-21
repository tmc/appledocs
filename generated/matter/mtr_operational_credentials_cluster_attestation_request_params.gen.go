// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterAttestationRequestParams] class.
var (
	MTROperationalCredentialsClusterAttestationRequestParamsClass     _MTROperationalCredentialsClusterAttestationRequestParamsClass
	MTROperationalCredentialsClusterAttestationRequestParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterAttestationRequestParamsClass() _MTROperationalCredentialsClusterAttestationRequestParamsClass {
	MTROperationalCredentialsClusterAttestationRequestParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterAttestationRequestParamsClass = _MTROperationalCredentialsClusterAttestationRequestParamsClass{objc.GetClass("MTROperationalCredentialsClusterAttestationRequestParams")}
	})
	return MTROperationalCredentialsClusterAttestationRequestParamsClass
}

type _MTROperationalCredentialsClusterAttestationRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterAttestationRequestParams] class.
type IMTROperationalCredentialsClusterAttestationRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationRequestParams
type MTROperationalCredentialsClusterAttestationRequestParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterAttestationRequestParamsFrom constructs a [MTROperationalCredentialsClusterAttestationRequestParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterAttestationRequestParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterAttestationRequestParams {
	return MTROperationalCredentialsClusterAttestationRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterAttestationRequestParamsClass) Alloc() MTROperationalCredentialsClusterAttestationRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterAttestationRequestParamsClass) New() MTROperationalCredentialsClusterAttestationRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) Init() MTROperationalCredentialsClusterAttestationRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) Autorelease() MTROperationalCredentialsClusterAttestationRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterAttestationRequestParams creates a new MTROperationalCredentialsClusterAttestationRequestParams instance.
func NewMTROperationalCredentialsClusterAttestationRequestParams() MTROperationalCredentialsClusterAttestationRequestParams {
	return getMTROperationalCredentialsClusterAttestationRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/attestationnonce
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) AttestationNonce() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("attestationNonce"))
	return rv
}


// SetAttestationNonce sets the value of the attestationNonce property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/attestationnonce
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) SetAttestationNonce(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationNonce:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



