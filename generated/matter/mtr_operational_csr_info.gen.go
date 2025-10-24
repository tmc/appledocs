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
	// properties:
	AttestationSignature() objc.IObject /* cross-framework: Data */
	SetAttestationSignature(value objc.IObject /* cross-framework: Data */)
	Csr() objc.IObject /* cross-framework: Data */
	SetCsr(value objc.IObject /* cross-framework: Data */)
	CsrElementsTLV() objc.IObject /* cross-framework: Data */
	SetCsrElementsTLV(value objc.IObject /* cross-framework: Data */)
	CsrNonce() objc.IObject /* cross-framework: Data */
	SetCsrNonce(value objc.IObject /* cross-framework: Data */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/attestationsignature
func (m_ MTROperationalCSRInfo) AttestationSignature() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationSignature"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/attestationsignature
func (m_ MTROperationalCSRInfo) SetAttestationSignature(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationSignature:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csr
func (m_ MTROperationalCSRInfo) Csr() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("csr"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csr
func (m_ MTROperationalCSRInfo) SetCsr(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsr:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csrelementstlv
func (m_ MTROperationalCSRInfo) CsrElementsTLV() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("csrElementsTLV"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csrelementstlv
func (m_ MTROperationalCSRInfo) SetCsrElementsTLV(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrElementsTLV:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csrnonce
func (m_ MTROperationalCSRInfo) CsrNonce() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("csrNonce"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcsrinfo/csrnonce
func (m_ MTROperationalCSRInfo) SetCsrNonce(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrNonce:"), value)
}



