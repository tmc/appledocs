//go:build darwin && ios

// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Toolbar


// Inserts an item into the toolbar at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/insertItem(withItemIdentifier:at:)
func (t_ Toolbar) InsertItemWithItemIdentifierAtIndex(itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertItemWithItemIdentifier:atIndex:"), itemIdentifier, index)
}

// Removes the item at the specified index in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/removeItem(at:)
func (t_ Toolbar) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeItemAtIndex:"), index)
}

// Displays the toolbar’s customization palette and handles any user-initiated customizations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/runCustomizationPalette(_:)
func (t_ Toolbar) RunCustomizationPalette(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("runCustomizationPalette:"), sender)
}

// iOS-only properties

// A Boolean value that indicates whether users can modify the contents of the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsUserCustomization
func (t_ Toolbar) AllowsUserCustomization() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsUserCustomization"))
	return rv
}
func (t_ Toolbar) SetAllowsUserCustomization(value bool) {
	t_.ID.Send(objc.RegisterName("setAllowsUserCustomization:"), value)
}

// A Boolean value that indicates whether the toolbar autosaves its configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/autosavesConfiguration
func (t_ Toolbar) AutosavesConfiguration() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autosavesConfiguration"))
	return rv
}
func (t_ Toolbar) SetAutosavesConfiguration(value bool) {
	t_.ID.Send(objc.RegisterName("setAutosavesConfiguration:"), value)
}

// A Boolean value that indicates whether the toolbar’s customization palette is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/customizationPaletteIsRunning
func (t_ Toolbar) CustomizationPaletteIsRunning() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("customizationPaletteIsRunning"))
	return rv
}

// The object you use to customize the toolbar contents and configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/delegate
func (t_ Toolbar) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}
func (t_ Toolbar) SetDelegate(value objc.ID) {
	t_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// A value that indicates whether the toolbar displays items using a name, icon, or combination of elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/displayMode-swift.property
func (t_ Toolbar) DisplayMode() ToolbarDisplayMode {
	rv := objc.Send[ToolbarDisplayMode](t_.ID, objc.Sel("displayMode"))
	return rv
}
func (t_ Toolbar) SetDisplayMode(value ToolbarDisplayMode) {
	t_.ID.Send(objc.RegisterName("setDisplayMode:"), value)
}

// The value you use to identify the toolbar in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/identifier-swift.property
func (t_ Toolbar) Identifier() objc.IObject /* cross-framework: ToolbarIdentifier */ {
	rv := objc.Send[ToolbarIdentifier](t_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value that indicates whether the toolbar is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/isVisible
func (t_ Toolbar) Visible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("visible"))
	return rv
}
func (t_ Toolbar) SetVisible(value bool) {
	t_.ID.Send(objc.RegisterName("setVisible:"), value)
}

// An array containing the toolbar’s current items, in order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/items
func (t_ Toolbar) Items() []IToolbarItem {
	rv := objc.Send[[]ToolbarItem](t_.ID, objc.Sel("items"))
	return rv
}

// The identifier of the toolbar’s currently selected item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/selectedItemIdentifier
func (t_ Toolbar) SelectedItemIdentifier() objc.IObject /* cross-framework: ToolbarItemIdentifier */ {
	rv := objc.Send[ToolbarItemIdentifier](t_.ID, objc.Sel("selectedItemIdentifier"))
	return rv
}
func (t_ Toolbar) SetSelectedItemIdentifier(value objc.IObject /* cross-framework: ToolbarItemIdentifier */) {
	t_.ID.Send(objc.RegisterName("setSelectedItemIdentifier:"), value)
}

// An array containing the toolbar’s currently visible items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/visibleItems
func (t_ Toolbar) VisibleItems() []IToolbarItem {
	rv := objc.Send[[]ToolbarItem](t_.ID, objc.Sel("visibleItems"))
	return rv
}




