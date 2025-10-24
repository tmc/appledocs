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
	// properties:
	RootCACertificate() objc.IObject /* cross-framework: Data */
	SetRootCACertificate(value objc.IObject /* cross-framework: Data */)
	RootCertificate() objc.IObject /* cross-framework: Data */
	SetRootCertificate(value objc.IObject /* cross-framework: Data */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/rootcacertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) RootCACertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootCACertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/rootcacertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetRootCACertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCACertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/rootcertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) RootCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/rootcertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetRootCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddtrustedrootcertificateparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



