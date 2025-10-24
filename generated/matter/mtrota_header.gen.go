// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTAHeader */


/* debug [class_header]: Header for MTROTAHeader */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTAHeader */
// An interface definition for the [MTROTAHeader] class.
type IMTROTAHeader interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTAHeader */
	// properties:
	ImageDigest() objc.IObject /* cross-framework: NSData */
	SetImageDigest(value objc.IObject /* cross-framework: NSData */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROTAHeader */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTAHeader */
// Alloc allocates a new instance without initialization.
func (mc _MTROTAHeaderClass) Alloc() MTROTAHeader {
	rv := objc.Send[MTROTAHeader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTAHeader */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader
type MTROTAHeader struct {
	objectivec.Object
}

// MTROTAHeaderFrom constructs a [MTROTAHeader] from an unsafe.Pointer.
func MTROTAHeaderFrom(ptr unsafe.Pointer) MTROTAHeader {
	return MTROTAHeader{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTAHeader */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/init(data:)
func NewMTROTAHeaderWithData(data objc.IObject /* cross-framework: NSData */) MTROTAHeader {
	instance := getMTROTAHeaderClass().Alloc()
	rv := objc.Send[MTROTAHeader](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROTAHeaderWithData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTAHeader */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTAHeader */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTAHeader */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTAHeader */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/imageDigest
func (m_ MTROTAHeader) ImageDigest() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("imageDigest"))
	return rv
}/* debug [instance_properties/getter]: imageDigest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/imageDigest
func (m_ MTROTAHeader) SetImageDigest(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageDigest:"), value)
}/* debug [instance_properties/setter]: imageDigest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/imageDigestType
func (m_ MTROTAHeader) ImageDigestType() MTROTAImageDigestType {
	rv := objc.Send[MTROTAImageDigestType](m_.ID, objc.Sel("imageDigestType"))
	return rv
}/* debug [instance_properties/getter]: imageDigestType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/imageDigestType
func (m_ MTROTAHeader) SetImageDigestType(value MTROTAImageDigestType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageDigestType:"), value)
}/* debug [instance_properties/setter]: imageDigestType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/maxApplicableVersion
func (m_ MTROTAHeader) MaxApplicableVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxApplicableVersion"))
	return rv
}/* debug [instance_properties/getter]: maxApplicableVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/maxApplicableVersion
func (m_ MTROTAHeader) SetMaxApplicableVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxApplicableVersion:"), value)
}/* debug [instance_properties/setter]: maxApplicableVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/minApplicableVersion
func (m_ MTROTAHeader) MinApplicableVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minApplicableVersion"))
	return rv
}/* debug [instance_properties/getter]: minApplicableVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/minApplicableVersion
func (m_ MTROTAHeader) SetMinApplicableVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinApplicableVersion:"), value)
}/* debug [instance_properties/setter]: minApplicableVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/payloadSize
func (m_ MTROTAHeader) PayloadSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("payloadSize"))
	return rv
}/* debug [instance_properties/getter]: payloadSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/payloadSize
func (m_ MTROTAHeader) SetPayloadSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayloadSize:"), value)
}/* debug [instance_properties/setter]: payloadSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/productID
func (m_ MTROTAHeader) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}/* debug [instance_properties/getter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/productID
func (m_ MTROTAHeader) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}/* debug [instance_properties/setter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/releaseNotesURL
func (m_ MTROTAHeader) ReleaseNotesURL() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("releaseNotesURL"))
	return rv
}/* debug [instance_properties/getter]: releaseNotesURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/releaseNotesURL
func (m_ MTROTAHeader) SetReleaseNotesURL(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReleaseNotesURL:"), value)
}/* debug [instance_properties/setter]: releaseNotesURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/softwareVersion
func (m_ MTROTAHeader) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/softwareVersion
func (m_ MTROTAHeader) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/softwareVersionString
func (m_ MTROTAHeader) SoftwareVersionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("softwareVersionString"))
	return rv
}/* debug [instance_properties/getter]: softwareVersionString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/softwareVersionString
func (m_ MTROTAHeader) SetSoftwareVersionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersionString:"), value)
}/* debug [instance_properties/setter]: softwareVersionString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/vendorID
func (m_ MTROTAHeader) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader/vendorID
func (m_ MTROTAHeader) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}/* debug [instance_properties/setter]: vendorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTAHeader */


