// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterCSRResponseParams] class.
var (
	MTROperationalCredentialsClusterCSRResponseParamsClass     _MTROperationalCredentialsClusterCSRResponseParamsClass
	MTROperationalCredentialsClusterCSRResponseParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterCSRResponseParamsClass() _MTROperationalCredentialsClusterCSRResponseParamsClass {
	MTROperationalCredentialsClusterCSRResponseParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterCSRResponseParamsClass = _MTROperationalCredentialsClusterCSRResponseParamsClass{objc.GetClass("MTROperationalCredentialsClusterCSRResponseParams")}
	})
	return MTROperationalCredentialsClusterCSRResponseParamsClass
}

type _MTROperationalCredentialsClusterCSRResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterCSRResponseParams] class.
type IMTROperationalCredentialsClusterCSRResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRResponseParams
type MTROperationalCredentialsClusterCSRResponseParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterCSRResponseParamsFrom constructs a [MTROperationalCredentialsClusterCSRResponseParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterCSRResponseParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterCSRResponseParams {
	return MTROperationalCredentialsClusterCSRResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterCSRResponseParamsClass) Alloc() MTROperationalCredentialsClusterCSRResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterCSRResponseParamsClass) New() MTROperationalCredentialsClusterCSRResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterCSRResponseParams) Init() MTROperationalCredentialsClusterCSRResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterCSRResponseParams) Autorelease() MTROperationalCredentialsClusterCSRResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterCSRResponseParams creates a new MTROperationalCredentialsClusterCSRResponseParams instance.
func NewMTROperationalCredentialsClusterCSRResponseParams() MTROperationalCredentialsClusterCSRResponseParams {
	return getMTROperationalCredentialsClusterCSRResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrresponseparams/attestationsignature
func (m_ MTROperationalCredentialsClusterCSRResponseParams) AttestationSignature() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationSignature"))
	return rv
}


// SetAttestationSignature sets the value of the attestationSignature property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrresponseparams/attestationsignature
func (m_ MTROperationalCredentialsClusterCSRResponseParams) SetAttestationSignature(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationSignature:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrresponseparams/nocsrelements
func (m_ MTROperationalCredentialsClusterCSRResponseParams) NocsrElements() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("nocsrElements"))
	return rv
}


// SetNocsrElements sets the value of the nocsrElements property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrresponseparams/nocsrelements
func (m_ MTROperationalCredentialsClusterCSRResponseParams) SetNocsrElements(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocsrElements:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCSRResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCSRResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



