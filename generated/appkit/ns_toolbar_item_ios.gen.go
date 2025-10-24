//go:build darwin && ios

// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ToolbarItem


// iOS-only properties

// The menu item to use for the toolbar item is in the overflow menu in a Mac app built with Mac Catalyst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/itemMenuFormRepresentation
func (t_ ToolbarItem) ItemMenuFormRepresentation() MenuElement /* not a class type */ {
	rv := objc.Send[MenuElement](t_.ID, objc.Sel("itemMenuFormRepresentation"))
	return rv
}
func (t_ ToolbarItem) SetItemMenuFormRepresentation(value MenuElement /* not a class type */) {
	t_.ID.Send(objc.RegisterName("setItemMenuFormRepresentation:"), value)
}




