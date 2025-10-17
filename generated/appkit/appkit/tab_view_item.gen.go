// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TabViewItem] class.
var TabViewItemClass objc.Class

func init() {
	TabViewItemClass = objc.GetClass("NSTabViewItem")
}

type TabViewItem struct {
	objc.ID
}

func TabViewItemFrom(ptr unsafe.Pointer) TabViewItem {
	return TabViewItem{
		ID: objc.ID(ptr),
	}
}




