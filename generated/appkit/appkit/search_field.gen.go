// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SearchField] class.
var SearchFieldClass objc.Class

func init() {
	SearchFieldClass = objc.GetClass("NSSearchField")
}

type SearchField struct {
	objc.ID
}

func SearchFieldFrom(ptr unsafe.Pointer) SearchField {
	return SearchField{
		ID: objc.ID(ptr),
	}
}




