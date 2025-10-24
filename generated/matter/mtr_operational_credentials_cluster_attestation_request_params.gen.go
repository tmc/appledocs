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
	// properties:
	AttestationNonce() objc.IObject /* cross-framework: Data */
	SetAttestationNonce(value objc.IObject /* cross-framework: Data */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/attestationnonce
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) AttestationNonce() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationNonce"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/attestationnonce
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) SetAttestationNonce(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationNonce:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterattestationrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



