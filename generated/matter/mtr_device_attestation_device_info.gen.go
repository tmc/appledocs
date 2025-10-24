// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceAttestationDeviceInfo */


/* debug [class_header]: Header for MTRDeviceAttestationDeviceInfo */
// The class instance for the [MTRDeviceAttestationDeviceInfo] class.
var (
	MTRDeviceAttestationDeviceInfoClass     _MTRDeviceAttestationDeviceInfoClass
	MTRDeviceAttestationDeviceInfoClassOnce sync.Once
)

func getMTRDeviceAttestationDeviceInfoClass() _MTRDeviceAttestationDeviceInfoClass {
	MTRDeviceAttestationDeviceInfoClassOnce.Do(func() {
		MTRDeviceAttestationDeviceInfoClass = _MTRDeviceAttestationDeviceInfoClass{objc.GetClass("MTRDeviceAttestationDeviceInfo")}
	})
	return MTRDeviceAttestationDeviceInfoClass
}

type _MTRDeviceAttestationDeviceInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceAttestationDeviceInfo */
// An interface definition for the [MTRDeviceAttestationDeviceInfo] class.
type IMTRDeviceAttestationDeviceInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceAttestationDeviceInfo */
	// properties:
	AttestationChallenge() objc.IObject /* cross-framework: NSData */
	AttestationNonce() objc.IObject /* cross-framework: NSData */
	BasicInformationProductID() objc.IObject /* cross-framework: NSNumber */
	BasicInformationVendorID() objc.IObject /* cross-framework: NSNumber */
	CertificateDeclaration() objc.IObject /* cross-framework: NSData */
	CertificationDeclaration() objc.IObject /* cross-framework: NSData */
	DacCertificate() unsafe.Pointer
	DacPAICertificate() unsafe.Pointer
	ElementsSignature() objc.IObject /* cross-framework: NSData */
	ElementsTLV() unsafe.Pointer
	ProductID() objc.IObject /* cross-framework: NSNumber */
	VendorID() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceAttestationDeviceInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceAttestationDeviceInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceAttestationDeviceInfoClass) Alloc() MTRDeviceAttestationDeviceInfo {
	rv := objc.Send[MTRDeviceAttestationDeviceInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceAttestationDeviceInfoClass) New() MTRDeviceAttestationDeviceInfo {
	rv := objc.Send[MTRDeviceAttestationDeviceInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceAttestationDeviceInfo) Init() MTRDeviceAttestationDeviceInfo {
	rv := objc.Send[MTRDeviceAttestationDeviceInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceAttestationDeviceInfo) Autorelease() MTRDeviceAttestationDeviceInfo {
	rv := objc.Send[MTRDeviceAttestationDeviceInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceAttestationDeviceInfo creates a new MTRDeviceAttestationDeviceInfo instance.
func NewMTRDeviceAttestationDeviceInfo() MTRDeviceAttestationDeviceInfo {
	return getMTRDeviceAttestationDeviceInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceAttestationDeviceInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo
type MTRDeviceAttestationDeviceInfo struct {
	objectivec.Object
}

// MTRDeviceAttestationDeviceInfoFrom constructs a [MTRDeviceAttestationDeviceInfo] from an unsafe.Pointer.
func MTRDeviceAttestationDeviceInfoFrom(ptr unsafe.Pointer) MTRDeviceAttestationDeviceInfo {
	return MTRDeviceAttestationDeviceInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceAttestationDeviceInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceAttestationDeviceInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceAttestationDeviceInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceAttestationDeviceInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceAttestationDeviceInfo */

// The attestation challenge from the secure session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/attestationChallenge
func (m_ MTRDeviceAttestationDeviceInfo) AttestationChallenge() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("attestationChallenge"))
	return rv
}/* debug [instance_properties/getter]: attestationChallenge */


// The attestation nonce from the AttestationRequest command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/attestationNonce
func (m_ MTRDeviceAttestationDeviceInfo) AttestationNonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("attestationNonce"))
	return rv
}/* debug [instance_properties/getter]: attestationNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/basicInformationProductID
func (m_ MTRDeviceAttestationDeviceInfo) BasicInformationProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("basicInformationProductID"))
	return rv
}/* debug [instance_properties/getter]: basicInformationProductID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/basicInformationVendorID
func (m_ MTRDeviceAttestationDeviceInfo) BasicInformationVendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("basicInformationVendorID"))
	return rv
}/* debug [instance_properties/getter]: basicInformationVendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/certificateDeclaration
func (m_ MTRDeviceAttestationDeviceInfo) CertificateDeclaration() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("certificateDeclaration"))
	return rv
}/* debug [instance_properties/getter]: certificateDeclaration */


// The certification declaration of the device, if available. This is a DER-encoded string representing a CMS-formatted certification declaration. May be nil only if attestation verification failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/certificationDeclaration
func (m_ MTRDeviceAttestationDeviceInfo) CertificationDeclaration() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("certificationDeclaration"))
	return rv
}/* debug [instance_properties/getter]: certificationDeclaration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/dacCertificate
func (m_ MTRDeviceAttestationDeviceInfo) DacCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dacCertificate"))
	return rv
}/* debug [instance_properties/getter]: dacCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/dacPAICertificate
func (m_ MTRDeviceAttestationDeviceInfo) DacPAICertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dacPAICertificate"))
	return rv
}/* debug [instance_properties/getter]: dacPAICertificate */


// A signature, using the device attestation private key of the device that sent the attestation information, over the concatenation of elementsTLV and attestationChallenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/elementsSignature
func (m_ MTRDeviceAttestationDeviceInfo) ElementsSignature() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("elementsSignature"))
	return rv
}/* debug [instance_properties/getter]: elementsSignature */


// The TLV-encoded attestation_elements_message that was used to find the certificationDeclaration (possibly unsuccessfully).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/elementsTLV
func (m_ MTRDeviceAttestationDeviceInfo) ElementsTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("elementsTLV"))
	return rv
}/* debug [instance_properties/getter]: elementsTLV */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/productID
func (m_ MTRDeviceAttestationDeviceInfo) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}/* debug [instance_properties/getter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo/vendorID
func (m_ MTRDeviceAttestationDeviceInfo) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceAttestationDeviceInfo */



