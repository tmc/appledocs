//go:build darwin && ios

// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ToolbarItem


// Validates the toolbar item’s menu and its ability to perfrom its action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/validate()
func (t_ ToolbarItem) Validate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("validate"))
}

// iOS-only properties

// The action method to call when someone clicks on the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/action
func (t_ ToolbarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](t_.ID, objc.Sel("action"))
	return rv
}
func (t_ ToolbarItem) SetAction(value objc.SEL) {
	t_.ID.Send(objc.RegisterName("setAction:"), value)
}

// A Boolean value that indicates whether the toolbar automatically validates the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/autovalidates
func (t_ ToolbarItem) Autovalidates() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autovalidates"))
	return rv
}
func (t_ ToolbarItem) SetAutovalidates(value bool) {
	t_.ID.Send(objc.RegisterName("setAutovalidates:"), value)
}

// The image to display for the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/image
func (t_ ToolbarItem) Image() IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}
func (t_ ToolbarItem) SetImage(value IImage) {
	t_.ID.Send(objc.RegisterName("setImage:"), value)
}

// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isEnabled
func (t_ ToolbarItem) Enabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("enabled"))
	return rv
}
func (t_ ToolbarItem) SetEnabled(value bool) {
	t_.ID.Send(objc.RegisterName("setEnabled:"), value)
}

// The value you use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/itemIdentifier
func (t_ ToolbarItem) ItemIdentifier() objc.IObject /* cross-framework: ToolbarItemIdentifier */ {
	rv := objc.Send[ToolbarItemIdentifier](t_.ID, objc.Sel("itemIdentifier"))
	return rv
}

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

// The label that appears for this item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/label
func (t_ ToolbarItem) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("label"))
	return rv
}
func (t_ ToolbarItem) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	t_.ID.Send(objc.RegisterName("setLabel:"), value)
}

// The label that appears when the toolbar item is in the customization palette.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/paletteLabel
func (t_ ToolbarItem) PaletteLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("paletteLabel"))
	return rv
}
func (t_ ToolbarItem) SetPaletteLabel(value objc.IObject /* cross-framework: NSString */) {
	t_.ID.Send(objc.RegisterName("setPaletteLabel:"), value)
}

// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/tag
func (t_ ToolbarItem) Tag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("tag"))
	return rv
}
func (t_ ToolbarItem) SetTag(value int) {
	t_.ID.Send(objc.RegisterName("setTag:"), value)
}

// The object that defines the action method the toolbar item calls when clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/target
func (t_ ToolbarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("target"))
	return rv
}
func (t_ ToolbarItem) SetTarget(value objc.ID) {
	t_.ID.Send(objc.RegisterName("setTarget:"), value)
}

// The tooltip to display when someone hovers over the item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolTip
func (t_ ToolbarItem) ToolTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("toolTip"))
	return rv
}
func (t_ ToolbarItem) SetToolTip(value objc.IObject /* cross-framework: NSString */) {
	t_.ID.Send(objc.RegisterName("setToolTip:"), value)
}

// The toolbar that currently includes the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolbar
func (t_ ToolbarItem) Toolbar() IToolbar {
	rv := objc.Send[Toolbar](t_.ID, objc.Sel("toolbar"))
	return rv
}

// The display priority associated with the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/visibilityPriority-swift.property
func (t_ ToolbarItem) VisibilityPriority() objc.IObject /* cross-framework: ToolbarItemVisibilityPriority */ {
	rv := objc.Send[ToolbarItemVisibilityPriority](t_.ID, objc.Sel("visibilityPriority"))
	return rv
}
func (t_ ToolbarItem) SetVisibilityPriority(value objc.IObject /* cross-framework: ToolbarItemVisibilityPriority */) {
	t_.ID.Send(objc.RegisterName("setVisibilityPriority:"), value)
}




