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
}

//
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
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/attestationchallenge
func (m_ MTRDeviceAttestationDeviceInfo) AttestationChallenge() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("attestationChallenge"))
	return rv
}


// SetAttestationChallenge sets the value of the attestationChallenge property.
// The attestation challenge from the secure session.

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/attestationchallenge
func (m_ MTRDeviceAttestationDeviceInfo) SetAttestationChallenge(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationChallenge:"), value)
}

// The attestation nonce from the AttestationRequest command.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/attestationnonce
func (m_ MTRDeviceAttestationDeviceInfo) AttestationNonce() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("attestationNonce"))
	return rv
}


// SetAttestationNonce sets the value of the attestationNonce property.
// The attestation nonce from the AttestationRequest command.

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/attestationnonce
func (m_ MTRDeviceAttestationDeviceInfo) SetAttestationNonce(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationNonce:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/daccertificate
func (m_ MTRDeviceAttestationDeviceInfo) DacCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dacCertificate"))
	return rv
}


// SetDacCertificate sets the value of the dacCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/daccertificate
func (m_ MTRDeviceAttestationDeviceInfo) SetDacCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDacCertificate:"), value)
}

// The certification declaration of the device, if available. This is a DER-encoded string
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/certificationdeclaration
func (m_ MTRDeviceAttestationDeviceInfo) CertificationDeclaration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("certificationDeclaration"))
	return rv
}


// SetCertificationDeclaration sets the value of the certificationDeclaration property.
// The certification declaration of the device, if available. This is a DER-encoded string

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/certificationdeclaration
func (m_ MTRDeviceAttestationDeviceInfo) SetCertificationDeclaration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificationDeclaration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/vendorid
func (m_ MTRDeviceAttestationDeviceInfo) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/vendorid
func (m_ MTRDeviceAttestationDeviceInfo) SetVendorID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}

// A signature, using the device attestation private key of the device that sent
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/elementssignature
func (m_ MTRDeviceAttestationDeviceInfo) ElementsSignature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("elementsSignature"))
	return rv
}


// SetElementsSignature sets the value of the elementsSignature property.
// A signature, using the device attestation private key of the device that sent

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/elementssignature
func (m_ MTRDeviceAttestationDeviceInfo) SetElementsSignature(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElementsSignature:"), value)
}

// The TLV-encoded attestation_elements_message that was used to find the
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/elementstlv
func (m_ MTRDeviceAttestationDeviceInfo) ElementsTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("elementsTLV"))
	return rv
}


// SetElementsTLV sets the value of the elementsTLV property.
// The TLV-encoded attestation_elements_message that was used to find the

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/elementstlv
func (m_ MTRDeviceAttestationDeviceInfo) SetElementsTLV(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElementsTLV:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/dacpaicertificate
func (m_ MTRDeviceAttestationDeviceInfo) DacPAICertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dacPAICertificate"))
	return rv
}


// SetDacPAICertificate sets the value of the dacPAICertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/dacpaicertificate
func (m_ MTRDeviceAttestationDeviceInfo) SetDacPAICertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDacPAICertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/basicinformationvendorid
func (m_ MTRDeviceAttestationDeviceInfo) BasicInformationVendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("basicInformationVendorID"))
	return rv
}


// SetBasicInformationVendorID sets the value of the basicInformationVendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/basicinformationvendorid
func (m_ MTRDeviceAttestationDeviceInfo) SetBasicInformationVendorID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBasicInformationVendorID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/certificatedeclaration
func (m_ MTRDeviceAttestationDeviceInfo) CertificateDeclaration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("certificateDeclaration"))
	return rv
}


// SetCertificateDeclaration sets the value of the certificateDeclaration property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/certificatedeclaration
func (m_ MTRDeviceAttestationDeviceInfo) SetCertificateDeclaration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificateDeclaration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/productid
func (m_ MTRDeviceAttestationDeviceInfo) ProductID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("productID"))
	return rv
}


// SetProductID sets the value of the productID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/productid
func (m_ MTRDeviceAttestationDeviceInfo) SetProductID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/basicinformationproductid
func (m_ MTRDeviceAttestationDeviceInfo) BasicInformationProductID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("basicInformationProductID"))
	return rv
}


// SetBasicInformationProductID sets the value of the basicInformationProductID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationdeviceinfo/basicinformationproductid
func (m_ MTRDeviceAttestationDeviceInfo) SetBasicInformationProductID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBasicInformationProductID:"), value)
}



