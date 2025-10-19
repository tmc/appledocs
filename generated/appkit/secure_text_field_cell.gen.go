// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SecureTextFieldCell] class.
var secureTextFieldCellClass = _SecureTextFieldCellClass{objc.GetClass("NSSecureTextFieldCell")}

type _SecureTextFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [SecureTextFieldCell] class.
type ISecureTextFieldCell interface {
	ITextFieldCell
}

// A text field whose value is hidden from the user. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return secureTextFieldCellClass.New()
}




