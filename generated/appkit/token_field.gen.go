// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TokenField] class.
var tokenFieldClass = _TokenFieldClass{objc.GetClass("NSTokenField")}

type _TokenFieldClass struct {
	class objc.Class
}

// A text field that converts text into visually distinct tokens. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField

type TokenField struct {
	TextField
}

// TokenFieldFrom constructs a [TokenField] from an unsafe.Pointer.
//
// A text field that converts text into visually distinct tokens.
func TokenFieldFrom(ptr unsafe.Pointer) TokenField {
	return TokenField{
		TextField: TextFieldFrom(ptr),
	}
}



