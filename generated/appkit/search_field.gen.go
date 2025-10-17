// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SearchField] class.
var searchFieldClass = _SearchFieldClass{objc.GetClass("NSSearchField")}

type _SearchFieldClass struct {
	class objc.Class
}

// A text field optimized for performing text-based searches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField

type SearchField struct {
	TextField
}

// SearchFieldFrom constructs a [SearchField] from an unsafe.Pointer.
//
// A text field optimized for performing text-based searches.
func SearchFieldFrom(ptr unsafe.Pointer) SearchField {
	return SearchField{
		TextField: TextFieldFrom(ptr),
	}
}



