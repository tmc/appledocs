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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/vendorid
func (m_ MTROTAHeader) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/vendorid
func (m_ MTROTAHeader) SetVendorID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/imagedigesttype
func (m_ MTROTAHeader) ImageDigestType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("imageDigestType"))
	return rv
}


// SetImageDigestType sets the value of the imageDigestType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/imagedigesttype
func (m_ MTROTAHeader) SetImageDigestType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageDigestType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/maxapplicableversion
func (m_ MTROTAHeader) MaxApplicableVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("maxApplicableVersion"))
	return rv
}


// SetMaxApplicableVersion sets the value of the maxApplicableVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/maxapplicableversion
func (m_ MTROTAHeader) SetMaxApplicableVersion(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxApplicableVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/minapplicableversion
func (m_ MTROTAHeader) MinApplicableVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minApplicableVersion"))
	return rv
}


// SetMinApplicableVersion sets the value of the minApplicableVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/minapplicableversion
func (m_ MTROTAHeader) SetMinApplicableVersion(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinApplicableVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/imagedigest
func (m_ MTROTAHeader) ImageDigest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("imageDigest"))
	return rv
}


// SetImageDigest sets the value of the imageDigest property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/imagedigest
func (m_ MTROTAHeader) SetImageDigest(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageDigest:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/releasenotesurl
func (m_ MTROTAHeader) ReleaseNotesURL() string {
	rv := objc.Send[string](m_.ID, objc.Sel("releaseNotesURL"))
	return rv
}


// SetReleaseNotesURL sets the value of the releaseNotesURL property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/releasenotesurl
func (m_ MTROTAHeader) SetReleaseNotesURL(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReleaseNotesURL:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/productid
func (m_ MTROTAHeader) ProductID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("productID"))
	return rv
}


// SetProductID sets the value of the productID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/productid
func (m_ MTROTAHeader) SetProductID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/softwareversionstring
func (m_ MTROTAHeader) SoftwareVersionString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("softwareVersionString"))
	return rv
}


// SetSoftwareVersionString sets the value of the softwareVersionString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/softwareversionstring
func (m_ MTROTAHeader) SetSoftwareVersionString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersionString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/softwareversion
func (m_ MTROTAHeader) SoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// SetSoftwareVersion sets the value of the softwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/softwareversion
func (m_ MTROTAHeader) SetSoftwareVersion(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/payloadsize
func (m_ MTROTAHeader) PayloadSize() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("payloadSize"))
	return rv
}


// SetPayloadSize sets the value of the payloadSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotaheader/payloadsize
func (m_ MTROTAHeader) SetPayloadSize(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayloadSize:"), value)
}



