// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterCSRRequestParams] class.
var (
	MTROperationalCredentialsClusterCSRRequestParamsClass     _MTROperationalCredentialsClusterCSRRequestParamsClass
	MTROperationalCredentialsClusterCSRRequestParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterCSRRequestParamsClass() _MTROperationalCredentialsClusterCSRRequestParamsClass {
	MTROperationalCredentialsClusterCSRRequestParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterCSRRequestParamsClass = _MTROperationalCredentialsClusterCSRRequestParamsClass{objc.GetClass("MTROperationalCredentialsClusterCSRRequestParams")}
	})
	return MTROperationalCredentialsClusterCSRRequestParamsClass
}

type _MTROperationalCredentialsClusterCSRRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterCSRRequestParams] class.
type IMTROperationalCredentialsClusterCSRRequestParams interface {
	objectivec.IObject
	// properties:
	CsrNonce() objc.IObject /* cross-framework: Data */
	SetCsrNonce(value objc.IObject /* cross-framework: Data */)
	IsForUpdateNOC() objc.IObject /* cross-framework: NSNumber */
	SetIsForUpdateNOC(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams
type MTROperationalCredentialsClusterCSRRequestParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterCSRRequestParamsFrom constructs a [MTROperationalCredentialsClusterCSRRequestParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterCSRRequestParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterCSRRequestParams {
	return MTROperationalCredentialsClusterCSRRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterCSRRequestParamsClass) Alloc() MTROperationalCredentialsClusterCSRRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterCSRRequestParamsClass) New() MTROperationalCredentialsClusterCSRRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterCSRRequestParams) Init() MTROperationalCredentialsClusterCSRRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterCSRRequestParams) Autorelease() MTROperationalCredentialsClusterCSRRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterCSRRequestParams creates a new MTROperationalCredentialsClusterCSRRequestParams instance.
func NewMTROperationalCredentialsClusterCSRRequestParams() MTROperationalCredentialsClusterCSRRequestParams {
	return getMTROperationalCredentialsClusterCSRRequestParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/csrnonce
func (m_ MTROperationalCredentialsClusterCSRRequestParams) CsrNonce() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("csrNonce"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/csrnonce
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetCsrNonce(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrNonce:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/isforupdatenoc
func (m_ MTROperationalCredentialsClusterCSRRequestParams) IsForUpdateNOC() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("isForUpdateNOC"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/isforupdatenoc
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetIsForUpdateNOC(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsForUpdateNOC:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterCSRRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCSRRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



