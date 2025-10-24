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
	// properties:
	CertificationDeclaration() objc.IObject /* cross-framework: Data */
	SetCertificationDeclaration(value objc.IObject /* cross-framework: Data */)
	Challenge() objc.IObject /* cross-framework: Data */
	SetChallenge(value objc.IObject /* cross-framework: Data */)
	Dac() objc.IObject /* cross-framework: Data */
	SetDac(value objc.IObject /* cross-framework: Data */)
	Elements() objc.IObject /* cross-framework: Data */
	SetElements(value objc.IObject /* cross-framework: Data */)
	ElementsSignature() objc.IObject /* cross-framework: Data */
	SetElementsSignature(value objc.IObject /* cross-framework: Data */)
	FirmwareInfo() objc.IObject /* cross-framework: Data */
	SetFirmwareInfo(value objc.IObject /* cross-framework: Data */)
	Nonce() objc.IObject /* cross-framework: Data */
	SetNonce(value objc.IObject /* cross-framework: Data */)
	Pai() objc.IObject /* cross-framework: Data */
	SetPai(value objc.IObject /* cross-framework: Data */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/certificationdeclaration
func (a_ AttestationInfo) CertificationDeclaration() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("certificationDeclaration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/certificationdeclaration
func (a_ AttestationInfo) SetCertificationDeclaration(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCertificationDeclaration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/challenge
func (a_ AttestationInfo) Challenge() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("challenge"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/challenge
func (a_ AttestationInfo) SetChallenge(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChallenge:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/dac
func (a_ AttestationInfo) Dac() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("dac"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/dac
func (a_ AttestationInfo) SetDac(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDac:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/elements
func (a_ AttestationInfo) Elements() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("elements"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/elements
func (a_ AttestationInfo) SetElements(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setElements:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/elementssignature
func (a_ AttestationInfo) ElementsSignature() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("elementsSignature"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/elementssignature
func (a_ AttestationInfo) SetElementsSignature(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setElementsSignature:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/firmwareinfo
func (a_ AttestationInfo) FirmwareInfo() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("firmwareInfo"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/firmwareinfo
func (a_ AttestationInfo) SetFirmwareInfo(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFirmwareInfo:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/nonce
func (a_ AttestationInfo) Nonce() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("nonce"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/nonce
func (a_ AttestationInfo) SetNonce(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNonce:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/pai
func (a_ AttestationInfo) Pai() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("pai"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/attestationinfo/pai
func (a_ AttestationInfo) SetPai(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPai:"), value)
}



