// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRDeviceAttestationInfo] class.
type IMTRDeviceAttestationInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo
type MTRDeviceAttestationInfo struct {
	objectivec.Object
}

// MTRDeviceAttestationInfoFrom constructs a [MTRDeviceAttestationInfo] from an unsafe.Pointer.
func MTRDeviceAttestationInfoFrom(ptr unsafe.Pointer) MTRDeviceAttestationInfo {
	return MTRDeviceAttestationInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceAttestationInfoClass) Alloc() MTRDeviceAttestationInfo {
	rv := objc.Send[MTRDeviceAttestationInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/certificationdeclaration
func (m_ MTRDeviceAttestationInfo) CertificationDeclaration() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("certificationDeclaration"))
	return rv
}


// SetCertificationDeclaration sets the value of the certificationDeclaration property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/certificationdeclaration
func (m_ MTRDeviceAttestationInfo) SetCertificationDeclaration(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificationDeclaration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/challenge
func (m_ MTRDeviceAttestationInfo) Challenge() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("challenge"))
	return rv
}


// SetChallenge sets the value of the challenge property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/challenge
func (m_ MTRDeviceAttestationInfo) SetChallenge(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChallenge:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/deviceattestationcertificate
func (m_ MTRDeviceAttestationInfo) DeviceAttestationCertificate() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("deviceAttestationCertificate"))
	return rv
}


// SetDeviceAttestationCertificate sets the value of the deviceAttestationCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/deviceattestationcertificate
func (m_ MTRDeviceAttestationInfo) SetDeviceAttestationCertificate(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceAttestationCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/elementssignature
func (m_ MTRDeviceAttestationInfo) ElementsSignature() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("elementsSignature"))
	return rv
}


// SetElementsSignature sets the value of the elementsSignature property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/elementssignature
func (m_ MTRDeviceAttestationInfo) SetElementsSignature(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElementsSignature:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/elementstlv
func (m_ MTRDeviceAttestationInfo) ElementsTLV() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("elementsTLV"))
	return rv
}


// SetElementsTLV sets the value of the elementsTLV property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/elementstlv
func (m_ MTRDeviceAttestationInfo) SetElementsTLV(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElementsTLV:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/firmwareinfo
func (m_ MTRDeviceAttestationInfo) FirmwareInfo() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("firmwareInfo"))
	return rv
}


// SetFirmwareInfo sets the value of the firmwareInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/firmwareinfo
func (m_ MTRDeviceAttestationInfo) SetFirmwareInfo(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFirmwareInfo:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/nonce
func (m_ MTRDeviceAttestationInfo) Nonce() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("nonce"))
	return rv
}


// SetNonce sets the value of the nonce property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/nonce
func (m_ MTRDeviceAttestationInfo) SetNonce(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNonce:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/productattestationintermediatecertificate
func (m_ MTRDeviceAttestationInfo) ProductAttestationIntermediateCertificate() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("productAttestationIntermediateCertificate"))
	return rv
}


// SetProductAttestationIntermediateCertificate sets the value of the productAttestationIntermediateCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/productattestationintermediatecertificate
func (m_ MTRDeviceAttestationInfo) SetProductAttestationIntermediateCertificate(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductAttestationIntermediateCertificate:"), value)
}



