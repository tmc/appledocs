// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SearchFieldCell] class.
var searchFieldCellClass = _SearchFieldCellClass{objc.GetClass("NSSearchFieldCell")}

type _SearchFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [SearchFieldCell] class.
type ISearchFieldCell interface {
	ITextFieldCell
}

// The programmatic interface for text fields that are used for text-based searches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell

type SearchFieldCell struct {
	TextFieldCell
}

// SearchFieldCellFrom constructs a [SearchFieldCell] from an unsafe.Pointer.
//
// The programmatic interface for text fields that are used for text-based searches.
func SearchFieldCellFrom(ptr unsafe.Pointer) SearchFieldCell {
	return SearchFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}



