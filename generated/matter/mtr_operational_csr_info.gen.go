// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCSRInfo] class.
var (
	MTROperationalCSRInfoClass     _MTROperationalCSRInfoClass
	MTROperationalCSRInfoClassOnce sync.Once
)

func getMTROperationalCSRInfoClass() _MTROperationalCSRInfoClass {
	MTROperationalCSRInfoClassOnce.Do(func() {
		MTROperationalCSRInfoClass = _MTROperationalCSRInfoClass{objc.GetClass("MTROperationalCSRInfo")}
	})
	return MTROperationalCSRInfoClass
}

type _MTROperationalCSRInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCSRInfo] class.
type IMTROperationalCSRInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo
type MTROperationalCSRInfo struct {
	objectivec.Object
}

// MTROperationalCSRInfoFrom constructs a [MTROperationalCSRInfo] from an unsafe.Pointer.
func MTROperationalCSRInfoFrom(ptr unsafe.Pointer) MTROperationalCSRInfo {
	return MTROperationalCSRInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCSRInfoClass) Alloc() MTROperationalCSRInfo {
	rv := objc.Send[MTROperationalCSRInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCSRInfoClass) New() MTROperationalCSRInfo {
	rv := objc.Send[MTROperationalCSRInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCSRInfo) Init() MTROperationalCSRInfo {
	rv := objc.Send[MTROperationalCSRInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCSRInfo) Autorelease() MTROperationalCSRInfo {
	rv := objc.Send[MTROperationalCSRInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCSRInfo creates a new MTROperationalCSRInfo instance.
func NewMTROperationalCSRInfo() MTROperationalCSRInfo {
	return getMTROperationalCSRInfoClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/attestationsignature
func (m_ MTROperationalCSRInfo) AttestationSignature() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationSignature"))
	return rv
}


// SetAttestationSignature sets the value of the attestationSignature property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/attestationsignature
func (m_ MTROperationalCSRInfo) SetAttestationSignature(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationSignature:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csr
func (m_ MTROperationalCSRInfo) Csr() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("csr"))
	return rv
}


// SetCsr sets the value of the csr property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csr
func (m_ MTROperationalCSRInfo) SetCsr(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsr:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csrelementstlv
func (m_ MTROperationalCSRInfo) CsrElementsTLV() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("csrElementsTLV"))
	return rv
}


// SetCsrElementsTLV sets the value of the csrElementsTLV property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csrelementstlv
func (m_ MTROperationalCSRInfo) SetCsrElementsTLV(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrElementsTLV:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csrnonce
func (m_ MTROperationalCSRInfo) CsrNonce() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("csrNonce"))
	return rv
}


// SetCsrNonce sets the value of the csrNonce property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csrnonce
func (m_ MTROperationalCSRInfo) SetCsrNonce(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrNonce:"), value)
}



