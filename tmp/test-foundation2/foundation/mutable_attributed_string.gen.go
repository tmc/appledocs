// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var mutableAttributedStringClass _MutableAttributedStringClass

func init() {
	mutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}
}

type _MutableAttributedStringClass struct {
	class objc.Class
}

type MutableAttributedString struct {
	objc.ID
}

func MutableAttributedStringFrom(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{
		ID: objc.ID(ptr),
	}
}


// Adds the characters and attributes of a given attributed string to the end of the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/append(_:)
func (m_ MutableAttributedString) AppendAttributedString(attrString unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendAttributedString:"), attrString)
}


