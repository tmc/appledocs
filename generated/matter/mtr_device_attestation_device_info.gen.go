// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRDeviceAttestationDeviceInfo] class.
type IMTRDeviceAttestationDeviceInfo interface {
	objectivec.IObject
	// properties:
	AttestationChallenge() objc.IObject /* cross-framework: Data */
	SetAttestationChallenge(value objc.IObject /* cross-framework: Data */)
	AttestationNonce() objc.IObject /* cross-framework: Data */
	SetAttestationNonce(value objc.IObject /* cross-framework: Data */)
	BasicInformationProductID() objc.IObject /* cross-framework: NSNumber */
	SetBasicInformationProductID(value objc.IObject /* cross-framework: NSNumber */)
	BasicInformationVendorID() objc.IObject /* cross-framework: NSNumber */
	SetBasicInformationVendorID(value objc.IObject /* cross-framework: NSNumber */)
	CertificateDeclaration() objc.IObject /* cross-framework: Data */
	SetCertificateDeclaration(value objc.IObject /* cross-framework: Data */)
	CertificationDeclaration() objc.IObject /* cross-framework: Data */
	SetCertificationDeclaration(value objc.IObject /* cross-framework: Data */)
	DacCertificate() objc.IObject /* cross-framework: Data */
	SetDacCertificate(value objc.IObject /* cross-framework: Data */)
	DacPAICertificate() objc.IObject /* cross-framework: Data */
	SetDacPAICertificate(value objc.IObject /* cross-framework: Data */)
	ElementsSignature() objc.IObject /* cross-framework: Data */
	SetElementsSignature(value objc.IObject /* cross-framework: Data */)
	ElementsTLV() objc.IObject /* cross-framework: Data */
	SetElementsTLV(value objc.IObject /* cross-framework: Data */)
	ProductID() objc.IObject /* cross-framework: NSNumber */
	SetProductID(value objc.IObject /* cross-framework: NSNumber */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo
type MTRDeviceAttestationDeviceInfo struct {
	objectivec.Object
}

// MTRDeviceAttestationDeviceInfoFrom constructs a [MTRDeviceAttestationDeviceInfo] from an unsafe.Pointer.
func MTRDeviceAttestationDeviceInfoFrom(ptr unsafe.Pointer) MTRDeviceAttestationDeviceInfo {
	return MTRDeviceAttestationDeviceInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceAttestationDeviceInfoClass) Alloc() MTRDeviceAttestationDeviceInfo {
	rv := objc.Send[MTRDeviceAttestationDeviceInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The attestation challenge from the secure session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/attestationchallenge
func (m_ MTRDeviceAttestationDeviceInfo) AttestationChallenge() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationChallenge"))
	return rv
}


// The attestation challenge from the secure session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/attestationchallenge
func (m_ MTRDeviceAttestationDeviceInfo) SetAttestationChallenge(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationChallenge:"), value)
}


// The attestation nonce from the AttestationRequest command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/attestationnonce
func (m_ MTRDeviceAttestationDeviceInfo) AttestationNonce() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationNonce"))
	return rv
}


// The attestation nonce from the AttestationRequest command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/attestationnonce
func (m_ MTRDeviceAttestationDeviceInfo) SetAttestationNonce(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationNonce:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/basicinformationproductid
func (m_ MTRDeviceAttestationDeviceInfo) BasicInformationProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("basicInformationProductID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/basicinformationproductid
func (m_ MTRDeviceAttestationDeviceInfo) SetBasicInformationProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBasicInformationProductID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/basicinformationvendorid
func (m_ MTRDeviceAttestationDeviceInfo) BasicInformationVendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("basicInformationVendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/basicinformationvendorid
func (m_ MTRDeviceAttestationDeviceInfo) SetBasicInformationVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBasicInformationVendorID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/certificatedeclaration
func (m_ MTRDeviceAttestationDeviceInfo) CertificateDeclaration() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("certificateDeclaration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/certificatedeclaration
func (m_ MTRDeviceAttestationDeviceInfo) SetCertificateDeclaration(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificateDeclaration:"), value)
}


// The certification declaration of the device, if available. This is a DER-encoded string
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/certificationdeclaration
func (m_ MTRDeviceAttestationDeviceInfo) CertificationDeclaration() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("certificationDeclaration"))
	return rv
}


// The certification declaration of the device, if available. This is a DER-encoded string
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/certificationdeclaration
func (m_ MTRDeviceAttestationDeviceInfo) SetCertificationDeclaration(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificationDeclaration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/daccertificate
func (m_ MTRDeviceAttestationDeviceInfo) DacCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("dacCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/daccertificate
func (m_ MTRDeviceAttestationDeviceInfo) SetDacCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDacCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/dacpaicertificate
func (m_ MTRDeviceAttestationDeviceInfo) DacPAICertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("dacPAICertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/dacpaicertificate
func (m_ MTRDeviceAttestationDeviceInfo) SetDacPAICertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDacPAICertificate:"), value)
}


// A signature, using the device attestation private key of the device that sent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/elementssignature
func (m_ MTRDeviceAttestationDeviceInfo) ElementsSignature() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("elementsSignature"))
	return rv
}


// A signature, using the device attestation private key of the device that sent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/elementssignature
func (m_ MTRDeviceAttestationDeviceInfo) SetElementsSignature(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElementsSignature:"), value)
}


// The TLV-encoded attestation_elements_message that was used to find the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/elementstlv
func (m_ MTRDeviceAttestationDeviceInfo) ElementsTLV() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("elementsTLV"))
	return rv
}


// The TLV-encoded attestation_elements_message that was used to find the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/elementstlv
func (m_ MTRDeviceAttestationDeviceInfo) SetElementsTLV(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElementsTLV:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/productid
func (m_ MTRDeviceAttestationDeviceInfo) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/productid
func (m_ MTRDeviceAttestationDeviceInfo) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/vendorid
func (m_ MTRDeviceAttestationDeviceInfo) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/vendorid
func (m_ MTRDeviceAttestationDeviceInfo) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}



