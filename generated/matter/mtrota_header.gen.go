// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTAHeader] class.
var (
	MTROTAHeaderClass     _MTROTAHeaderClass
	MTROTAHeaderClassOnce sync.Once
)

func getMTROTAHeaderClass() _MTROTAHeaderClass {
	MTROTAHeaderClassOnce.Do(func() {
		MTROTAHeaderClass = _MTROTAHeaderClass{objc.GetClass("MTROTAHeader")}
	})
	return MTROTAHeaderClass
}

type _MTROTAHeaderClass struct {
	class objc.Class
}

// An interface definition for the [MTROTAHeader] class.
type IMTROTAHeader interface {
	objectivec.IObject
	// properties:
	ImageDigest() objc.IObject /* cross-framework: Data */
	SetImageDigest(value objc.IObject /* cross-framework: Data */)
	ImageDigestType() MTROTAImageDigestType
	SetImageDigestType(value MTROTAImageDigestType)
	MaxApplicableVersion() objc.IObject /* cross-framework: NSNumber */
	SetMaxApplicableVersion(value objc.IObject /* cross-framework: NSNumber */)
	MinApplicableVersion() objc.IObject /* cross-framework: NSNumber */
	SetMinApplicableVersion(value objc.IObject /* cross-framework: NSNumber */)
	PayloadSize() objc.IObject /* cross-framework: NSNumber */
	SetPayloadSize(value objc.IObject /* cross-framework: NSNumber */)
	ProductID() objc.IObject /* cross-framework: NSNumber */
	SetProductID(value objc.IObject /* cross-framework: NSNumber */)
	ReleaseNotesURL() objc.IObject /* cross-framework: NSString */
	SetReleaseNotesURL(value objc.IObject /* cross-framework: NSString */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersionString() objc.IObject /* cross-framework: NSString */
	SetSoftwareVersionString(value objc.IObject /* cross-framework: NSString */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader
type MTROTAHeader struct {
	objectivec.Object
}

// MTROTAHeaderFrom constructs a [MTROTAHeader] from an unsafe.Pointer.
func MTROTAHeaderFrom(ptr unsafe.Pointer) MTROTAHeader {
	return MTROTAHeader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTAHeaderClass) Alloc() MTROTAHeader {
	rv := objc.Send[MTROTAHeader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTAHeaderClass) New() MTROTAHeader {
	rv := objc.Send[MTROTAHeader](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTAHeader) Init() MTROTAHeader {
	rv := objc.Send[MTROTAHeader](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTAHeader) Autorelease() MTROTAHeader {
	rv := objc.Send[MTROTAHeader](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTAHeader creates a new MTROTAHeader instance.
func NewMTROTAHeader() MTROTAHeader {
	return getMTROTAHeaderClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/imagedigest
func (m_ MTROTAHeader) ImageDigest() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("imageDigest"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/imagedigest
func (m_ MTROTAHeader) SetImageDigest(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageDigest:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/imagedigesttype
func (m_ MTROTAHeader) ImageDigestType() MTROTAImageDigestType {
	rv := objc.Send[MTROTAImageDigestType](m_.ID, objc.Sel("imageDigestType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/imagedigesttype
func (m_ MTROTAHeader) SetImageDigestType(value MTROTAImageDigestType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageDigestType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/maxapplicableversion
func (m_ MTROTAHeader) MaxApplicableVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxApplicableVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/maxapplicableversion
func (m_ MTROTAHeader) SetMaxApplicableVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxApplicableVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/minapplicableversion
func (m_ MTROTAHeader) MinApplicableVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minApplicableVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/minapplicableversion
func (m_ MTROTAHeader) SetMinApplicableVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinApplicableVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/payloadsize
func (m_ MTROTAHeader) PayloadSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("payloadSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/payloadsize
func (m_ MTROTAHeader) SetPayloadSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayloadSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/productid
func (m_ MTROTAHeader) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/productid
func (m_ MTROTAHeader) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/releasenotesurl
func (m_ MTROTAHeader) ReleaseNotesURL() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("releaseNotesURL"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/releasenotesurl
func (m_ MTROTAHeader) SetReleaseNotesURL(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReleaseNotesURL:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/softwareversion
func (m_ MTROTAHeader) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/softwareversion
func (m_ MTROTAHeader) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/softwareversionstring
func (m_ MTROTAHeader) SoftwareVersionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("softwareVersionString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/softwareversionstring
func (m_ MTROTAHeader) SetSoftwareVersionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersionString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/vendorid
func (m_ MTROTAHeader) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/vendorid
func (m_ MTROTAHeader) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}



