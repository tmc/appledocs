// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SecureTextField] class.
var secureTextFieldClass = _SecureTextFieldClass{objc.GetClass("NSSecureTextField")}

type _SecureTextFieldClass struct {
	class objc.Class
}

// An interface definition for the [SecureTextField] class.
type ISecureTextField interface {
	ITextField
}

// A text field that hides the typed text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSecureTextField

type SecureTextField struct {
	TextField
}

// SecureTextFieldFrom constructs a [SecureTextField] from an unsafe.Pointer.
//
// A text field that hides the typed text.
func SecureTextFieldFrom(ptr unsafe.Pointer) SecureTextField {
	return SecureTextField{
		TextField: TextFieldFrom(ptr),
	}
}



