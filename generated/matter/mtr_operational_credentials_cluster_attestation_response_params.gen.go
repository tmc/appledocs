// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterAttestationResponseParams] class.
var (
	MTROperationalCredentialsClusterAttestationResponseParamsClass     _MTROperationalCredentialsClusterAttestationResponseParamsClass
	MTROperationalCredentialsClusterAttestationResponseParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterAttestationResponseParamsClass() _MTROperationalCredentialsClusterAttestationResponseParamsClass {
	MTROperationalCredentialsClusterAttestationResponseParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterAttestationResponseParamsClass = _MTROperationalCredentialsClusterAttestationResponseParamsClass{objc.GetClass("MTROperationalCredentialsClusterAttestationResponseParams")}
	})
	return MTROperationalCredentialsClusterAttestationResponseParamsClass
}

type _MTROperationalCredentialsClusterAttestationResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterAttestationResponseParams] class.
type IMTROperationalCredentialsClusterAttestationResponseParams interface {
	objectivec.IObject
	AttestationElements() foundation.Data
	SetAttestationElements(value foundation.IData)
	AttestationSignature() foundation.Data
	SetAttestationSignature(value foundation.IData)
	Signature() foundation.Data
	SetSignature(value foundation.IData)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams
type MTROperationalCredentialsClusterAttestationResponseParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterAttestationResponseParamsFrom constructs a [MTROperationalCredentialsClusterAttestationResponseParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterAttestationResponseParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterAttestationResponseParams {
	return MTROperationalCredentialsClusterAttestationResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterAttestationResponseParamsClass) Alloc() MTROperationalCredentialsClusterAttestationResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterAttestationResponseParamsClass) New() MTROperationalCredentialsClusterAttestationResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) Init() MTROperationalCredentialsClusterAttestationResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) Autorelease() MTROperationalCredentialsClusterAttestationResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterAttestationResponseParams creates a new MTROperationalCredentialsClusterAttestationResponseParams instance.
func NewMTROperationalCredentialsClusterAttestationResponseParams() MTROperationalCredentialsClusterAttestationResponseParams {
	return getMTROperationalCredentialsClusterAttestationResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/attestationelements
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) AttestationElements() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationElements"))
	return rv
}


// SetAttestationElements sets the value of the attestationElements property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/attestationelements
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetAttestationElements(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationElements:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/attestationsignature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) AttestationSignature() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationSignature"))
	return rv
}


// SetAttestationSignature sets the value of the attestationSignature property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/attestationsignature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetAttestationSignature(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationSignature:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/signature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) Signature() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("signature"))
	return rv
}


// SetSignature sets the value of the signature property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/signature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetSignature(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSignature:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



