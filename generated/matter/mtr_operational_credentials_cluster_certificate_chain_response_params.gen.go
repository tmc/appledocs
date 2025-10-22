// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterCertificateChainResponseParams] class.
var (
	MTROperationalCredentialsClusterCertificateChainResponseParamsClass     _MTROperationalCredentialsClusterCertificateChainResponseParamsClass
	MTROperationalCredentialsClusterCertificateChainResponseParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterCertificateChainResponseParamsClass() _MTROperationalCredentialsClusterCertificateChainResponseParamsClass {
	MTROperationalCredentialsClusterCertificateChainResponseParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterCertificateChainResponseParamsClass = _MTROperationalCredentialsClusterCertificateChainResponseParamsClass{objc.GetClass("MTROperationalCredentialsClusterCertificateChainResponseParams")}
	})
	return MTROperationalCredentialsClusterCertificateChainResponseParamsClass
}

type _MTROperationalCredentialsClusterCertificateChainResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterCertificateChainResponseParams] class.
type IMTROperationalCredentialsClusterCertificateChainResponseParams interface {
	objectivec.IObject
	Certificate() foundation.Data
	SetCertificate(value foundation.IData)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainResponseParams
type MTROperationalCredentialsClusterCertificateChainResponseParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterCertificateChainResponseParamsFrom constructs a [MTROperationalCredentialsClusterCertificateChainResponseParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterCertificateChainResponseParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterCertificateChainResponseParams {
	return MTROperationalCredentialsClusterCertificateChainResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterCertificateChainResponseParamsClass) Alloc() MTROperationalCredentialsClusterCertificateChainResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterCertificateChainResponseParamsClass) New() MTROperationalCredentialsClusterCertificateChainResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) Init() MTROperationalCredentialsClusterCertificateChainResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) Autorelease() MTROperationalCredentialsClusterCertificateChainResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterCertificateChainResponseParams creates a new MTROperationalCredentialsClusterCertificateChainResponseParams instance.
func NewMTROperationalCredentialsClusterCertificateChainResponseParams() MTROperationalCredentialsClusterCertificateChainResponseParams {
	return getMTROperationalCredentialsClusterCertificateChainResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainresponseparams/certificate
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) Certificate() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("certificate"))
	return rv
}


// SetCertificate sets the value of the certificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainresponseparams/certificate
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) SetCertificate(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclustercertificatechainresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



