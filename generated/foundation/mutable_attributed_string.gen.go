// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableAttributedString] class.
var mutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}

type _MutableAttributedStringClass struct {
	class objc.Class
}

// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString

type MutableAttributedString struct {
	AttributedString
}

// MutableAttributedStringFrom constructs a [MutableAttributedString] from an unsafe.Pointer.
//
// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text.
func MutableAttributedStringFrom(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{
		AttributedString: AttributedStringFrom(ptr),
	}
}

// Adds the characters and attributes of a given attributed string to the end of the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/append(_:)
func (m_ MutableAttributedString) AppendAttributedString(attrString unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendAttributedString:"), attrString)
}


