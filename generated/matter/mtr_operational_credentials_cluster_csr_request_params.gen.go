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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/csrnonce
func (m_ MTROperationalCredentialsClusterCSRRequestParams) CsrNonce() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("csrNonce"))
	return rv
}


// SetCsrNonce sets the value of the csrNonce property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/csrnonce
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetCsrNonce(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrNonce:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/isforupdatenoc
func (m_ MTROperationalCredentialsClusterCSRRequestParams) IsForUpdateNOC() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("isForUpdateNOC"))
	return rv
}


// SetIsForUpdateNOC sets the value of the isForUpdateNOC property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/isforupdatenoc
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetIsForUpdateNOC(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsForUpdateNOC:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterCSRRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCSRRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercsrrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



