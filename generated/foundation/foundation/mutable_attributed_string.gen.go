// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableAttributedString] class.
var MutableAttributedStringClass objc.Class

func init() {
	MutableAttributedStringClass = objc.GetClass("NSMutableAttributedString")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableAttributedString/append(_:)
func (m_ MutableAttributedString) AppendAttributedString(attrString unsafe.Pointer) {
	sel := objc.RegisterName("appendAttributedString:")
	m_.ID.Send(sel, attrString)
}


