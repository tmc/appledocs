// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AttestationInfo */


/* debug [class_header]: Header for AttestationInfo */
// The class instance for the [AttestationInfo] class.
var (
	AttestationInfoClass     _AttestationInfoClass
	AttestationInfoClassOnce sync.Once
)

func getAttestationInfoClass() _AttestationInfoClass {
	AttestationInfoClassOnce.Do(func() {
		AttestationInfoClass = _AttestationInfoClass{objc.GetClass("AttestationInfo")}
	})
	return AttestationInfoClass
}

type _AttestationInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AttestationInfo */
// An interface definition for the [AttestationInfo] class.
type IAttestationInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AttestationInfo */
	// properties:
	CertificationDeclaration() objc.IObject /* cross-framework: NSData */
	SetCertificationDeclaration(value objc.IObject /* cross-framework: NSData */)
	Challenge() objc.IObject /* cross-framework: NSData */
	SetChallenge(value objc.IObject /* cross-framework: NSData */)
	Dac() objc.IObject /* cross-framework: NSData */
	SetDac(value objc.IObject /* cross-framework: NSData */)
	Elements() objc.IObject /* cross-framework: NSData */
	SetElements(value objc.IObject /* cross-framework: NSData */)
	ElementsSignature() objc.IObject /* cross-framework: NSData */
	SetElementsSignature(value objc.IObject /* cross-framework: NSData */)
	FirmwareInfo() objc.IObject /* cross-framework: NSData */
	SetFirmwareInfo(value objc.IObject /* cross-framework: NSData */)
	Nonce() objc.IObject /* cross-framework: NSData */
	SetNonce(value objc.IObject /* cross-framework: NSData */)
	Pai() objc.IObject /* cross-framework: NSData */
	SetPai(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AttestationInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AttestationInfo */
// Alloc allocates a new instance without initialization.
func (ac _AttestationInfoClass) Alloc() AttestationInfo {
	rv := objc.Send[AttestationInfo](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AttestationInfoClass) New() AttestationInfo {
	rv := objc.Send[AttestationInfo](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttestationInfo) Init() AttestationInfo {
	rv := objc.Send[AttestationInfo](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttestationInfo) Autorelease() AttestationInfo {
	rv := objc.Send[AttestationInfo](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttestationInfo creates a new AttestationInfo instance.
func NewAttestationInfo() AttestationInfo {
	return getAttestationInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AttestationInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo
type AttestationInfo struct {
	objectivec.Object
}

// AttestationInfoFrom constructs a [AttestationInfo] from an unsafe.Pointer.
func AttestationInfoFrom(ptr unsafe.Pointer) AttestationInfo {
	return AttestationInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AttestationInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/init(challenge:nonce:elements:elementsSignature:dac:pai:certificationDeclaration:firmwareInfo:)
func NewAttestationInfoWithChallengeNonceElementsElementsSignatureDacPaiCertificationDeclarationFirmwareInfo(challenge objc.IObject /* cross-framework: NSData */, nonce objc.IObject /* cross-framework: NSData */, elements objc.IObject /* cross-framework: NSData */, elementsSignature objc.IObject /* cross-framework: NSData */, dac objc.IObject /* cross-framework: NSData */, pai objc.IObject /* cross-framework: NSData */, certificationDeclaration objc.IObject /* cross-framework: NSData */, firmwareInfo objc.IObject /* cross-framework: NSData */) AttestationInfo {
	instance := getAttestationInfoClass().Alloc()
	rv := objc.Send[AttestationInfo](instance.ID, objc.Sel("initWithChallenge:nonce:elements:elementsSignature:dac:pai:certificationDeclaration:firmwareInfo:"), challenge, nonce, elements, elementsSignature, dac, pai, certificationDeclaration, firmwareInfo)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAttestationInfoWithChallengeNonceElementsElementsSignatureDacPaiCertificationDeclarationFirmwareInfo */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AttestationInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AttestationInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AttestationInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AttestationInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/certificationDeclaration
func (a_ AttestationInfo) CertificationDeclaration() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("certificationDeclaration"))
	return rv
}/* debug [instance_properties/getter]: certificationDeclaration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/certificationDeclaration
func (a_ AttestationInfo) SetCertificationDeclaration(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCertificationDeclaration:"), value)
}/* debug [instance_properties/setter]: certificationDeclaration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/challenge
func (a_ AttestationInfo) Challenge() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("challenge"))
	return rv
}/* debug [instance_properties/getter]: challenge */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/challenge
func (a_ AttestationInfo) SetChallenge(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChallenge:"), value)
}/* debug [instance_properties/setter]: challenge */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/dac
func (a_ AttestationInfo) Dac() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("dac"))
	return rv
}/* debug [instance_properties/getter]: dac */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/dac
func (a_ AttestationInfo) SetDac(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDac:"), value)
}/* debug [instance_properties/setter]: dac */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/elements
func (a_ AttestationInfo) Elements() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("elements"))
	return rv
}/* debug [instance_properties/getter]: elements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/elements
func (a_ AttestationInfo) SetElements(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setElements:"), value)
}/* debug [instance_properties/setter]: elements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/elementsSignature
func (a_ AttestationInfo) ElementsSignature() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("elementsSignature"))
	return rv
}/* debug [instance_properties/getter]: elementsSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/elementsSignature
func (a_ AttestationInfo) SetElementsSignature(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setElementsSignature:"), value)
}/* debug [instance_properties/setter]: elementsSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/firmwareInfo
func (a_ AttestationInfo) FirmwareInfo() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("firmwareInfo"))
	return rv
}/* debug [instance_properties/getter]: firmwareInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/firmwareInfo
func (a_ AttestationInfo) SetFirmwareInfo(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFirmwareInfo:"), value)
}/* debug [instance_properties/setter]: firmwareInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/nonce
func (a_ AttestationInfo) Nonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("nonce"))
	return rv
}/* debug [instance_properties/getter]: nonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/nonce
func (a_ AttestationInfo) SetNonce(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNonce:"), value)
}/* debug [instance_properties/setter]: nonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/pai
func (a_ AttestationInfo) Pai() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("pai"))
	return rv
}/* debug [instance_properties/getter]: pai */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo/pai
func (a_ AttestationInfo) SetPai(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPai:"), value)
}/* debug [instance_properties/setter]: pai */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AttestationInfo */


