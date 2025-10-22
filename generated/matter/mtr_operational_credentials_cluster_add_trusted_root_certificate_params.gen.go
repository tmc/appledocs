// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterAddTrustedRootCertificateParams] class.
var (
	MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass     _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass
	MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass() _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass {
	MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass = _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass{objc.GetClass("MTROperationalCredentialsClusterAddTrustedRootCertificateParams")}
	})
	return MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass
}

type _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterAddTrustedRootCertificateParams] class.
type IMTROperationalCredentialsClusterAddTrustedRootCertificateParams interface {
	objectivec.IObject
	RootCACertificate() foundation.Data
	SetRootCACertificate(value foundation.IData)
	RootCertificate() foundation.Data
	SetRootCertificate(value foundation.IData)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams
type MTROperationalCredentialsClusterAddTrustedRootCertificateParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterAddTrustedRootCertificateParamsFrom constructs a [MTROperationalCredentialsClusterAddTrustedRootCertificateParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterAddTrustedRootCertificateParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	return MTROperationalCredentialsClusterAddTrustedRootCertificateParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass) Alloc() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddTrustedRootCertificateParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass) New() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddTrustedRootCertificateParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) Init() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddTrustedRootCertificateParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) Autorelease() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddTrustedRootCertificateParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterAddTrustedRootCertificateParams creates a new MTROperationalCredentialsClusterAddTrustedRootCertificateParams instance.
func NewMTROperationalCredentialsClusterAddTrustedRootCertificateParams() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	return getMTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/rootcacertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) RootCACertificate() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootCACertificate"))
	return rv
}


// SetRootCACertificate sets the value of the rootCACertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/rootcacertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetRootCACertificate(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCACertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/rootcertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) RootCertificate() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootCertificate"))
	return rv
}


// SetRootCertificate sets the value of the rootCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/rootcertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetRootCertificate(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



