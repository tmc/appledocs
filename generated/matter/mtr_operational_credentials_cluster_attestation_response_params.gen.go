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
	// properties:
	AttestationElements() objc.IObject /* cross-framework: Data */
	SetAttestationElements(value objc.IObject /* cross-framework: Data */)
	AttestationSignature() objc.IObject /* cross-framework: Data */
	SetAttestationSignature(value objc.IObject /* cross-framework: Data */)
	Signature() objc.IObject /* cross-framework: Data */
	SetSignature(value objc.IObject /* cross-framework: Data */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/attestationelements
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) AttestationElements() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationElements"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/attestationelements
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetAttestationElements(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationElements:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/attestationsignature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) AttestationSignature() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationSignature"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/attestationsignature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetAttestationSignature(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationSignature:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/signature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) Signature() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("signature"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/signature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetSignature(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSignature:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



