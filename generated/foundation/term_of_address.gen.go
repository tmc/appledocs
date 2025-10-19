// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TermOfAddress] class.
var (
	termOfAddressClass     _TermOfAddressClass
	termOfAddressClassOnce sync.Once
)

func getTermOfAddressClass() _TermOfAddressClass {
	termOfAddressClassOnce.Do(func() {
		termOfAddressClass = _TermOfAddressClass{objc.GetClass("NSTermOfAddress")}
	})
	return termOfAddressClass
}

type _TermOfAddressClass struct {
	class objc.Class
}

// An interface definition for the [TermOfAddress] class.
type ITermOfAddress interface {
	objectivec.IObject
}

// The type for representing grammatical gender in localized text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress
type TermOfAddress struct {
	objectivec.Object
}

// TermOfAddressFrom constructs a [TermOfAddress] from an unsafe.Pointer.
//
// The type for representing grammatical gender in localized text.
func TermOfAddressFrom(ptr unsafe.Pointer) TermOfAddress {
	return TermOfAddress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TermOfAddressClass) Alloc() TermOfAddress {
	rv := objc.Send[TermOfAddress](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TermOfAddressClass) New() TermOfAddress {
	rv := objc.Send[TermOfAddress](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TermOfAddress) Init() TermOfAddress {
	rv := objc.Send[TermOfAddress](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TermOfAddress) Autorelease() TermOfAddress {
	rv := objc.Send[TermOfAddress](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTermOfAddress creates a new TermOfAddress instance.
func NewTermOfAddress() TermOfAddress {
	return getTermOfAddressClass().New()
}




