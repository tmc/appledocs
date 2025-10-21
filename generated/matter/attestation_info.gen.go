// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AttestationInfo] class.
type IAttestationInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo
type AttestationInfo struct {
	objectivec.Object
}

// AttestationInfoFrom constructs a [AttestationInfo] from an unsafe.Pointer.
func AttestationInfoFrom(ptr unsafe.Pointer) AttestationInfo {
	return AttestationInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AttestationInfoClass) Alloc() AttestationInfo {
	rv := objc.Send[AttestationInfo](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/certificationdeclaration
func (a_ AttestationInfo) CertificationDeclaration() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("certificationDeclaration"))
	return rv
}


// SetCertificationDeclaration sets the value of the certificationDeclaration property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/certificationdeclaration
func (a_ AttestationInfo) SetCertificationDeclaration(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCertificationDeclaration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/challenge
func (a_ AttestationInfo) Challenge() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("challenge"))
	return rv
}


// SetChallenge sets the value of the challenge property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/challenge
func (a_ AttestationInfo) SetChallenge(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChallenge:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/dac
func (a_ AttestationInfo) Dac() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("dac"))
	return rv
}


// SetDac sets the value of the dac property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/dac
func (a_ AttestationInfo) SetDac(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDac:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/elements
func (a_ AttestationInfo) Elements() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("elements"))
	return rv
}


// SetElements sets the value of the elements property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/elements
func (a_ AttestationInfo) SetElements(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setElements:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/elementssignature
func (a_ AttestationInfo) ElementsSignature() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("elementsSignature"))
	return rv
}


// SetElementsSignature sets the value of the elementsSignature property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/elementssignature
func (a_ AttestationInfo) SetElementsSignature(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setElementsSignature:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/firmwareinfo
func (a_ AttestationInfo) FirmwareInfo() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("firmwareInfo"))
	return rv
}


// SetFirmwareInfo sets the value of the firmwareInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/firmwareinfo
func (a_ AttestationInfo) SetFirmwareInfo(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFirmwareInfo:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/nonce
func (a_ AttestationInfo) Nonce() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("nonce"))
	return rv
}


// SetNonce sets the value of the nonce property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/nonce
func (a_ AttestationInfo) SetNonce(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNonce:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/pai
func (a_ AttestationInfo) Pai() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("pai"))
	return rv
}


// SetPai sets the value of the pai property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/pai
func (a_ AttestationInfo) SetPai(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPai:"), value)
}



