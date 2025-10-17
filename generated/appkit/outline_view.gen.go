// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OutlineView] class.
var outlineViewClass = _OutlineViewClass{objc.GetClass("NSOutlineView")}

type _OutlineViewClass struct {
	class objc.Class
}

// An interface definition for the [OutlineView] class.
type IOutlineView interface {
	ITableView
}

// A view that uses a row-and-column format to display hierarchical data like directories and files that can be expanded and collapsed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView

type OutlineView struct {
	TableView
}

// OutlineViewFrom constructs a [OutlineView] from an unsafe.Pointer.
//
// A view that uses a row-and-column format to display hierarchical data like directories and files that can be expanded and collapsed.
func OutlineViewFrom(ptr unsafe.Pointer) OutlineView {
	return OutlineView{
		TableView: TableViewFrom(ptr),
	}
}



