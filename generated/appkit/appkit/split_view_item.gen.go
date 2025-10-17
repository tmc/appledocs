// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SplitViewItem] class.
var SplitViewItemClass objc.Class

func init() {
	SplitViewItemClass = objc.GetClass("NSSplitViewItem")
}

type SplitViewItem struct {
	objc.ID
}

func SplitViewItemFrom(ptr unsafe.Pointer) SplitViewItem {
	return SplitViewItem{
		ID: objc.ID(ptr),
	}
}


//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func (sc SplitViewItem) InspectorWithViewController(viewController unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("inspectorWithViewController:")
	ret := objc.ID(SplitViewItemClass).Send(sel, viewController)
	return unsafe.Pointer(ret)
}


