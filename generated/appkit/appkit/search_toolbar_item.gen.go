// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SearchToolbarItem] class.
var SearchToolbarItemClass objc.Class

func init() {
	SearchToolbarItemClass = objc.GetClass("NSSearchToolbarItem")
}

type SearchToolbarItem struct {
	objc.ID
}

func SearchToolbarItemFrom(ptr unsafe.Pointer) SearchToolbarItem {
	return SearchToolbarItem{
		ID: objc.ID(ptr),
	}
}




