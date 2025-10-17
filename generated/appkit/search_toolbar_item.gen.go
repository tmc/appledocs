// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SearchToolbarItem] class.
var searchToolbarItemClass = _SearchToolbarItemClass{objc.GetClass("NSSearchToolbarItem")}

type _SearchToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [SearchToolbarItem] class.
type ISearchToolbarItem interface {
	IToolbarItem
}

// A toolbar item that contains a search field optimized for performing text-based searches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem

type SearchToolbarItem struct {
	ToolbarItem
}

// SearchToolbarItemFrom constructs a [SearchToolbarItem] from an unsafe.Pointer.
//
// A toolbar item that contains a search field optimized for performing text-based searches.
func SearchToolbarItemFrom(ptr unsafe.Pointer) SearchToolbarItem {
	return SearchToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}



