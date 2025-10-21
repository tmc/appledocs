// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterCertificateChainRequestParams] class.
var (
	MTROperationalCredentialsClusterCertificateChainRequestParamsClass     _MTROperationalCredentialsClusterCertificateChainRequestParamsClass
	MTROperationalCredentialsClusterCertificateChainRequestParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterCertificateChainRequestParamsClass() _MTROperationalCredentialsClusterCertificateChainRequestParamsClass {
	MTROperationalCredentialsClusterCertificateChainRequestParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterCertificateChainRequestParamsClass = _MTROperationalCredentialsClusterCertificateChainRequestParamsClass{objc.GetClass("MTROperationalCredentialsClusterCertificateChainRequestParams")}
	})
	return MTROperationalCredentialsClusterCertificateChainRequestParamsClass
}

type _MTROperationalCredentialsClusterCertificateChainRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterCertificateChainRequestParams] class.
type IMTROperationalCredentialsClusterCertificateChainRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainRequestParams
type MTROperationalCredentialsClusterCertificateChainRequestParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterCertificateChainRequestParamsFrom constructs a [MTROperationalCredentialsClusterCertificateChainRequestParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterCertificateChainRequestParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterCertificateChainRequestParams {
	return MTROperationalCredentialsClusterCertificateChainRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterCertificateChainRequestParamsClass) Alloc() MTROperationalCredentialsClusterCertificateChainRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterCertificateChainRequestParamsClass) New() MTROperationalCredentialsClusterCertificateChainRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) Init() MTROperationalCredentialsClusterCertificateChainRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) Autorelease() MTROperationalCredentialsClusterCertificateChainRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterCertificateChainRequestParams creates a new MTROperationalCredentialsClusterCertificateChainRequestParams instance.
func NewMTROperationalCredentialsClusterCertificateChainRequestParams() MTROperationalCredentialsClusterCertificateChainRequestParams {
	return getMTROperationalCredentialsClusterCertificateChainRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainrequestparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainrequestparams/certificatetype
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) CertificateType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("certificateType"))
	return rv
}


// SetCertificateType sets the value of the certificateType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainrequestparams/certificatetype
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) SetCertificateType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificateType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainrequestparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



