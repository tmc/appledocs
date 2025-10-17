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



