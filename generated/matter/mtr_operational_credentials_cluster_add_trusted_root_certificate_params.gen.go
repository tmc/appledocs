// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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




