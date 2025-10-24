// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceAttestationInfo */


/* debug [class_header]: Header for MTRDeviceAttestationInfo */
// The class instance for the [MTRDeviceAttestationInfo] class.
var (
	MTRDeviceAttestationInfoClass     _MTRDeviceAttestationInfoClass
	MTRDeviceAttestationInfoClassOnce sync.Once
)

func getMTRDeviceAttestationInfoClass() _MTRDeviceAttestationInfoClass {
	MTRDeviceAttestationInfoClassOnce.Do(func() {
		MTRDeviceAttestationInfoClass = _MTRDeviceAttestationInfoClass{objc.GetClass("MTRDeviceAttestationInfo")}
	})
	return MTRDeviceAttestationInfoClass
}

type _MTRDeviceAttestationInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceAttestationInfo */
// An interface definition for the [MTRDeviceAttestationInfo] class.
type IMTRDeviceAttestationInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceAttestationInfo */
	// properties:
	CertificationDeclaration() objc.IObject /* cross-framework: NSData */
	Challenge() objc.IObject /* cross-framework: NSData */
	DeviceAttestationCertificate() unsafe.Pointer
	ElementsSignature() objc.IObject /* cross-framework: NSData */
	ElementsTLV() unsafe.Pointer
	FirmwareInfo() objc.IObject /* cross-framework: NSData */
	Nonce() objc.IObject /* cross-framework: NSData */
	ProductAttestationIntermediateCertificate() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceAttestationInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceAttestationInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceAttestationInfoClass) Alloc() MTRDeviceAttestationInfo {
	rv := objc.Send[MTRDeviceAttestationInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceAttestationInfoClass) New() MTRDeviceAttestationInfo {
	rv := objc.Send[MTRDeviceAttestationInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceAttestationInfo) Init() MTRDeviceAttestationInfo {
	rv := objc.Send[MTRDeviceAttestationInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceAttestationInfo) Autorelease() MTRDeviceAttestationInfo {
	rv := objc.Send[MTRDeviceAttestationInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceAttestationInfo creates a new MTRDeviceAttestationInfo instance.
func NewMTRDeviceAttestationInfo() MTRDeviceAttestationInfo {
	return getMTRDeviceAttestationInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceAttestationInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo
type MTRDeviceAttestationInfo struct {
	objectivec.Object
}

// MTRDeviceAttestationInfoFrom constructs a [MTRDeviceAttestationInfo] from an unsafe.Pointer.
func MTRDeviceAttestationInfoFrom(ptr unsafe.Pointer) MTRDeviceAttestationInfo {
	return MTRDeviceAttestationInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceAttestationInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo/init(deviceAttestationChallenge:nonce:elementsTLV:elementsSignature:deviceAttestationCertificate:productAttestationIntermediateCertificate:certificationDeclaration:firmwareInfo:)
func NewMTRDeviceAttestationInfoWithDeviceAttestationChallengeNonceElementsTLVElementsSignatureDeviceAttestationCertificateProductAttestationIntermediateCertificateCertificationDeclarationFirmwareInfo(challenge objc.IObject /* cross-framework: NSData */, nonce objc.IObject /* cross-framework: NSData */, elementsTLV unsafe.Pointer, elementsSignature objc.IObject /* cross-framework: NSData */, deviceAttestationCertificate unsafe.Pointer, processAttestationIntermediateCertificate unsafe.Pointer, certificationDeclaration objc.IObject /* cross-framework: NSData */, firmwareInfo objc.IObject /* cross-framework: NSData */) MTRDeviceAttestationInfo {
	instance := getMTRDeviceAttestationInfoClass().Alloc()
	rv := objc.Send[MTRDeviceAttestationInfo](instance.ID, objc.Sel("initWithDeviceAttestationChallenge:nonce:elementsTLV:elementsSignature:deviceAttestationCertificate:productAttestationIntermediateCertificate:certificationDeclaration:firmwareInfo:"), challenge, nonce, elementsTLV, elementsSignature, deviceAttestationCertificate, processAttestationIntermediateCertificate, certificationDeclaration, firmwareInfo)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceAttestationInfoWithDeviceAttestationChallengeNonceElementsTLVElementsSignatureDeviceAttestationCertificateProductAttestationIntermediateCertificateCertificationDeclarationFirmwareInfo */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceAttestationInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceAttestationInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceAttestationInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceAttestationInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo/certificationDeclaration
func (m_ MTRDeviceAttestationInfo) CertificationDeclaration() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("certificationDeclaration"))
	return rv
}/* debug [instance_properties/getter]: certificationDeclaration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo/challenge
func (m_ MTRDeviceAttestationInfo) Challenge() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("challenge"))
	return rv
}/* debug [instance_properties/getter]: challenge */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo/deviceAttestationCertificate
func (m_ MTRDeviceAttestationInfo) DeviceAttestationCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deviceAttestationCertificate"))
	return rv
}/* debug [instance_properties/getter]: deviceAttestationCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo/elementsSignature
func (m_ MTRDeviceAttestationInfo) ElementsSignature() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("elementsSignature"))
	return rv
}/* debug [instance_properties/getter]: elementsSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo/elementsTLV
func (m_ MTRDeviceAttestationInfo) ElementsTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("elementsTLV"))
	return rv
}/* debug [instance_properties/getter]: elementsTLV */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo/firmwareInfo
func (m_ MTRDeviceAttestationInfo) FirmwareInfo() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("firmwareInfo"))
	return rv
}/* debug [instance_properties/getter]: firmwareInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo/nonce
func (m_ MTRDeviceAttestationInfo) Nonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("nonce"))
	return rv
}/* debug [instance_properties/getter]: nonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo/productAttestationIntermediateCertificate
func (m_ MTRDeviceAttestationInfo) ProductAttestationIntermediateCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("productAttestationIntermediateCertificate"))
	return rv
}/* debug [instance_properties/getter]: productAttestationIntermediateCertificate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceAttestationInfo */


