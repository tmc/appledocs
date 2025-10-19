// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SecureTextFieldCell] class.
var (
	secureTextFieldCellClass     _SecureTextFieldCellClass
	secureTextFieldCellClassOnce sync.Once
)

func getSecureTextFieldCellClass() _SecureTextFieldCellClass {
	secureTextFieldCellClassOnce.Do(func() {
		secureTextFieldCellClass = _SecureTextFieldCellClass{objc.GetClass("NSSecureTextFieldCell")}
	})
	return secureTextFieldCellClass
}

type _SecureTextFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [SecureTextFieldCell] class.
type ISecureTextFieldCell interface {
	ITextFieldCell
}

// A text field whose value is hidden from the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSecureTextFieldCell
type SecureTextFieldCell struct {
	TextFieldCell
}

// SecureTextFieldCellFrom constructs a [SecureTextFieldCell] from an unsafe.Pointer.
//
// A text field whose value is hidden from the user.
func SecureTextFieldCellFrom(ptr unsafe.Pointer) SecureTextFieldCell {
	return SecureTextFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SecureTextFieldCellClass) Alloc() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SecureTextFieldCellClass) New() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SecureTextFieldCell) Init() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SecureTextFieldCell) Autorelease() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSecureTextFieldCell creates a new SecureTextFieldCell instance.
func NewSecureTextFieldCell() SecureTextFieldCell {
	return getSecureTextFieldCellClass().New()
}




