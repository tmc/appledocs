// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Toolbar] class.
var ToolbarClass objc.Class

func init() {
	ToolbarClass = objc.GetClass("NSToolbar")
}

type Toolbar struct {
	objc.ID
}

func ToolbarFrom(ptr unsafe.Pointer) Toolbar {
	return Toolbar{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc Toolbar) Alloc() Toolbar {
	ret := objc.ID(ToolbarClass).Send(objc.RegisterName("alloc"))
	return Toolbar{ret}
}

// Init initializes the instance.
func (t_ Toolbar) Init() Toolbar {
	ret := t_.ID.Send(objc.RegisterName("init"))
	return Toolbar{ret}
}
// Creates a new toolbar with an empty identifier string. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/init()
func NewToolbar() Toolbar {
	instance := Toolbar{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Creates a newly allocated toolbar with the specified identifier. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/init(identifier:)
func NewToolbarWithIdentifier(identifier unsafe.Pointer) Toolbar {
	instance := Toolbar{}.Alloc()
	sel := objc.RegisterName("initWithIdentifier:")
	ret := instance.ID.Send(sel, identifier)
	instance = Toolbar{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Inserts an item into the toolbar at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/insertItem(withItemIdentifier:at:)
func (t_ Toolbar) InsertItemWithItemIdentifierAtIndex(itemIdentifier unsafe.Pointer, index int) {
	sel := objc.RegisterName("insertItemWithItemIdentifier:atIndex:")
	t_.ID.Send(sel, itemIdentifier, index)
}
// Removes the item at the specified index in the toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/removeItem(at:)
func (t_ Toolbar) RemoveItemAtIndex(index int) {
	sel := objc.RegisterName("removeItemAtIndex:")
	t_.ID.Send(sel, index)
}
// Removes the item with matching   in the receiving toolbar. If multiple items share the same identifier (as is the case with space items) all matching items will be removed. To remove only a single space item, use   instead. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/removeItem(identifier:)
func (t_ Toolbar) RemoveItemWithItemIdentifier(itemIdentifier unsafe.Pointer) {
	sel := objc.RegisterName("removeItemWithItemIdentifier:")
	t_.ID.Send(sel, itemIdentifier)
}
// Displays the toolbar’s customization palette and handles any user-initiated customizations. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/runCustomizationPalette(_:)
func (t_ Toolbar) RunCustomizationPalette(sender objc.ID) {
	sel := objc.RegisterName("runCustomizationPalette:")
	t_.ID.Send(sel, sender)
}
// Specifies the new configuration details for the toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/setConfiguration(_:)
func (t_ Toolbar) SetConfigurationFromDictionary(configDict unsafe.Pointer) {
	sel := objc.RegisterName("setConfigurationFromDictionary:")
	t_.ID.Send(sel, configDict)
}
// Validates the toolbar’s visible items during a window update. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/validateVisibleItems()
func (t_ Toolbar) ValidateVisibleItems() {
	sel := objc.RegisterName("validateVisibleItems")
	t_.ID.Send(sel)
}

