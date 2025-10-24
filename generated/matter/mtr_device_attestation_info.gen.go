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
	// properties:
	CertificationDeclaration() objc.IObject /* cross-framework: Data */
	SetCertificationDeclaration(value objc.IObject /* cross-framework: Data */)
	Challenge() objc.IObject /* cross-framework: Data */
	SetChallenge(value objc.IObject /* cross-framework: Data */)
	DeviceAttestationCertificate() objc.IObject /* cross-framework: Data */
	SetDeviceAttestationCertificate(value objc.IObject /* cross-framework: Data */)
	ElementsSignature() objc.IObject /* cross-framework: Data */
	SetElementsSignature(value objc.IObject /* cross-framework: Data */)
	ElementsTLV() objc.IObject /* cross-framework: Data */
	SetElementsTLV(value objc.IObject /* cross-framework: Data */)
	FirmwareInfo() objc.IObject /* cross-framework: Data */
	SetFirmwareInfo(value objc.IObject /* cross-framework: Data */)
	Nonce() objc.IObject /* cross-framework: Data */
	SetNonce(value objc.IObject /* cross-framework: Data */)
	ProductAttestationIntermediateCertificate() objc.IObject /* cross-framework: Data */
	SetProductAttestationIntermediateCertificate(value objc.IObject /* cross-framework: Data */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/certificationdeclaration
func (m_ MTRDeviceAttestationInfo) CertificationDeclaration() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("certificationDeclaration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/certificationdeclaration
func (m_ MTRDeviceAttestationInfo) SetCertificationDeclaration(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificationDeclaration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/challenge
func (m_ MTRDeviceAttestationInfo) Challenge() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("challenge"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/challenge
func (m_ MTRDeviceAttestationInfo) SetChallenge(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChallenge:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/deviceattestationcertificate
func (m_ MTRDeviceAttestationInfo) DeviceAttestationCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("deviceAttestationCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/deviceattestationcertificate
func (m_ MTRDeviceAttestationInfo) SetDeviceAttestationCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceAttestationCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/elementssignature
func (m_ MTRDeviceAttestationInfo) ElementsSignature() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("elementsSignature"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/elementssignature
func (m_ MTRDeviceAttestationInfo) SetElementsSignature(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElementsSignature:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/elementstlv
func (m_ MTRDeviceAttestationInfo) ElementsTLV() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("elementsTLV"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/elementstlv
func (m_ MTRDeviceAttestationInfo) SetElementsTLV(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElementsTLV:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/firmwareinfo
func (m_ MTRDeviceAttestationInfo) FirmwareInfo() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("firmwareInfo"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/firmwareinfo
func (m_ MTRDeviceAttestationInfo) SetFirmwareInfo(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFirmwareInfo:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/nonce
func (m_ MTRDeviceAttestationInfo) Nonce() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("nonce"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/nonce
func (m_ MTRDeviceAttestationInfo) SetNonce(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNonce:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/productattestationintermediatecertificate
func (m_ MTRDeviceAttestationInfo) ProductAttestationIntermediateCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("productAttestationIntermediateCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceattestationinfo/productattestationintermediatecertificate
func (m_ MTRDeviceAttestationInfo) SetProductAttestationIntermediateCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductAttestationIntermediateCertificate:"), value)
}



