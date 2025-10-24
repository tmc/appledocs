// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRCertificateInfo] class.
var (
	MTRCertificateInfoClass     _MTRCertificateInfoClass
	MTRCertificateInfoClassOnce sync.Once
)

func getMTRCertificateInfoClass() _MTRCertificateInfoClass {
	MTRCertificateInfoClassOnce.Do(func() {
		MTRCertificateInfoClass = _MTRCertificateInfoClass{objc.GetClass("MTRCertificateInfo")}
	})
	return MTRCertificateInfoClass
}

type _MTRCertificateInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRCertificateInfo] class.
type IMTRCertificateInfo interface {
	objectivec.IObject
	// properties:
	Issuer() IMTRDistinguishedNameInfo
	SetIssuer(value IMTRDistinguishedNameInfo)
	NotAfter() objc.IObject /* cross-framework: Date */
	SetNotAfter(value objc.IObject /* cross-framework: Date */)
	NotBefore() objc.IObject /* cross-framework: Date */
	SetNotBefore(value objc.IObject /* cross-framework: Date */)
	PublicKeyData() objc.IObject /* cross-framework: Data */
	SetPublicKeyData(value objc.IObject /* cross-framework: Data */)
	Subject() IMTRDistinguishedNameInfo
	SetSubject(value IMTRDistinguishedNameInfo)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificateInfo
type MTRCertificateInfo struct {
	objectivec.Object
}

// MTRCertificateInfoFrom constructs a [MTRCertificateInfo] from an unsafe.Pointer.
func MTRCertificateInfoFrom(ptr unsafe.Pointer) MTRCertificateInfo {
	return MTRCertificateInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCertificateInfoClass) Alloc() MTRCertificateInfo {
	rv := objc.Send[MTRCertificateInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCertificateInfoClass) New() MTRCertificateInfo {
	rv := objc.Send[MTRCertificateInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCertificateInfo) Init() MTRCertificateInfo {
	rv := objc.Send[MTRCertificateInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCertificateInfo) Autorelease() MTRCertificateInfo {
	rv := objc.Send[MTRCertificateInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCertificateInfo creates a new MTRCertificateInfo instance.
func NewMTRCertificateInfo() MTRCertificateInfo {
	return getMTRCertificateInfoClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/issuer
func (m_ MTRCertificateInfo) Issuer() IMTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("issuer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/issuer
func (m_ MTRCertificateInfo) SetIssuer(value IMTRDistinguishedNameInfo) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIssuer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/notafter
func (m_ MTRCertificateInfo) NotAfter() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("notAfter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/notafter
func (m_ MTRCertificateInfo) SetNotAfter(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNotAfter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/notbefore
func (m_ MTRCertificateInfo) NotBefore() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("notBefore"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/notbefore
func (m_ MTRCertificateInfo) SetNotBefore(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNotBefore:"), value)
}


// Public key data for this certificate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/publickeydata
func (m_ MTRCertificateInfo) PublicKeyData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("publicKeyData"))
	return rv
}


// Public key data for this certificate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/publickeydata
func (m_ MTRCertificateInfo) SetPublicKeyData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPublicKeyData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/subject
func (m_ MTRCertificateInfo) Subject() IMTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("subject"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcertificateinfo/subject
func (m_ MTRCertificateInfo) SetSubject(value IMTRDistinguishedNameInfo) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubject:"), value)
}



