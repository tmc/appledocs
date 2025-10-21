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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/csr
func (c_ CSRInfo) Csr() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("csr"))
	return rv
}


// SetCsr sets the value of the csr property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/csr
func (c_ CSRInfo) SetCsr(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCsr:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/elements
func (c_ CSRInfo) Elements() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("elements"))
	return rv
}


// SetElements sets the value of the elements property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/elements
func (c_ CSRInfo) SetElements(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setElements:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/elementssignature
func (c_ CSRInfo) ElementsSignature() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("elementsSignature"))
	return rv
}


// SetElementsSignature sets the value of the elementsSignature property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/elementssignature
func (c_ CSRInfo) SetElementsSignature(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setElementsSignature:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/nonce
func (c_ CSRInfo) Nonce() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("nonce"))
	return rv
}


// SetNonce sets the value of the nonce property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/csrinfo/nonce
func (c_ CSRInfo) SetNonce(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNonce:"), value)
}



