// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [issuerTextField] class.
var (
	IssuerTextFieldClass     _issuerTextFieldClass
	IssuerTextFieldClassOnce sync.Once
)

func getissuerTextFieldClass() _issuerTextFieldClass {
	IssuerTextFieldClassOnce.Do(func() {
		IssuerTextFieldClass = _issuerTextFieldClass{objc.GetClass("issuerTextField")}
	})
	return IssuerTextFieldClass
}

type _issuerTextFieldClass struct {
	class objc.Class
}

// An interface definition for the [issuerTextField] class.
type IissuerTextField interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityTableCellView/issuerTextField-c.ivar
type issuerTextField struct {
	objectivec.Object
}

// issuerTextFieldFrom constructs a [issuerTextField] from an unsafe.Pointer.
func issuerTextFieldFrom(ptr unsafe.Pointer) issuerTextField {
	return issuerTextField{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _issuerTextFieldClass) Alloc() issuerTextField {
	rv := objc.Send[issuerTextField](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _issuerTextFieldClass) New() issuerTextField {
	rv := objc.Send[issuerTextField](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ issuerTextField) Init() issuerTextField {
	rv := objc.Send[issuerTextField](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ issuerTextField) Autorelease() issuerTextField {
	rv := objc.Send[issuerTextField](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewissuerTextField creates a new issuerTextField instance.
func NewissuerTextField() issuerTextField {
	return getissuerTextFieldClass().New()
}




