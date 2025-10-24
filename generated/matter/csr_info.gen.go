// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CSRInfo] class.
var (
	CSRInfoClass     _CSRInfoClass
	CSRInfoClassOnce sync.Once
)

func getCSRInfoClass() _CSRInfoClass {
	CSRInfoClassOnce.Do(func() {
		CSRInfoClass = _CSRInfoClass{objc.GetClass("CSRInfo")}
	})
	return CSRInfoClass
}

type _CSRInfoClass struct {
	class objc.Class
}

// An interface definition for the [CSRInfo] class.
type ICSRInfo interface {
	objectivec.IObject
	// properties:
	Csr() objc.IObject /* cross-framework: Data */
	SetCsr(value objc.IObject /* cross-framework: Data */)
	Elements() objc.IObject /* cross-framework: Data */
	SetElements(value objc.IObject /* cross-framework: Data */)
	ElementsSignature() objc.IObject /* cross-framework: Data */
	SetElementsSignature(value objc.IObject /* cross-framework: Data */)
	Nonce() objc.IObject /* cross-framework: Data */
	SetNonce(value objc.IObject /* cross-framework: Data */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo
type CSRInfo struct {
	objectivec.Object
}

// CSRInfoFrom constructs a [CSRInfo] from an unsafe.Pointer.
func CSRInfoFrom(ptr unsafe.Pointer) CSRInfo {
	return CSRInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CSRInfoClass) Alloc() CSRInfo {
	rv := objc.Send[CSRInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSRInfoClass) New() CSRInfo {
	rv := objc.Send[CSRInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSRInfo) Init() CSRInfo {
	rv := objc.Send[CSRInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSRInfo) Autorelease() CSRInfo {
	rv := objc.Send[CSRInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSRInfo creates a new CSRInfo instance.
func NewCSRInfo() CSRInfo {
	return getCSRInfoClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/csr
func (c_ CSRInfo) Csr() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("csr"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/csr
func (c_ CSRInfo) SetCsr(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCsr:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/elements
func (c_ CSRInfo) Elements() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("elements"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/elements
func (c_ CSRInfo) SetElements(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setElements:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/elementssignature
func (c_ CSRInfo) ElementsSignature() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("elementsSignature"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/elementssignature
func (c_ CSRInfo) SetElementsSignature(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setElementsSignature:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/nonce
func (c_ CSRInfo) Nonce() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("nonce"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/nonce
func (c_ CSRInfo) SetNonce(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNonce:"), value)
}



