// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SecureTextField] class.
var SecureTextFieldClass objc.Class

func init() {
	SecureTextFieldClass = objc.GetClass("NSSecureTextField")
}

type SecureTextField struct {
	objc.ID
}

func SecureTextFieldFrom(ptr unsafe.Pointer) SecureTextField {
	return SecureTextField{
		ID: objc.ID(ptr),
	}
}



