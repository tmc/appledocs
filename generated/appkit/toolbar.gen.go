// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Toolbar] class.
var (
	toolbarClass     _ToolbarClass
	toolbarClassOnce sync.Once
)

func getToolbarClass() _ToolbarClass {
	toolbarClassOnce.Do(func() {
		toolbarClass = _ToolbarClass{objc.GetClass("NSToolbar")}
	})
	return toolbarClass
}

type _ToolbarClass struct {
	class objc.Class
}

// An interface definition for the [Toolbar] class.
type IToolbar interface {
	objectivec.IObject
	InsertItemWithItemIdentifierAtIndex(itemIdentifier unsafe.Pointer, index int)
	RemoveItemAtIndex(index int)
	RemoveItemWithItemIdentifier(itemIdentifier unsafe.Pointer)
	RunCustomizationPalette(sender objc.ID)
	SetConfigurationFromDictionary(configDict unsafe.Pointer)
	ValidateVisibleItems()
}

// An object that manages the space above your app’s custom content and either below or integrated with the window’s title bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar
type Toolbar struct {
	objectivec.Object
}

// ToolbarFrom constructs a [Toolbar] from an unsafe.Pointer.
//
// An object that manages the space above your app’s custom content and either below or integrated with the window’s title bar.
func ToolbarFrom(ptr unsafe.Pointer) Toolbar {
	return Toolbar{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolbarClass) Alloc() Toolbar {
	rv := objc.Send[Toolbar](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolbarClass) New() Toolbar {
	rv := objc.Send[Toolbar](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Toolbar) Init() Toolbar {
	rv := objc.Send[Toolbar](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Toolbar) Autorelease() Toolbar {
	rv := objc.Send[Toolbar](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolbar creates a new Toolbar instance.
func NewToolbar() Toolbar {
	return getToolbarClass().New()
}


// Creates a newly allocated toolbar with the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/init(identifier:)
func NewToolbarWithIdentifier(identifier unsafe.Pointer) Toolbar {
	instance := getToolbarClass().Alloc()
	rv := objc.Send[Toolbar](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}


// Inserts an item into the toolbar at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/insertItem(withItemIdentifier:at:)
func (t_ Toolbar) InsertItemWithItemIdentifierAtIndex(itemIdentifier unsafe.Pointer, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertItemWithItemIdentifier:atIndex:"), itemIdentifier, index)
}

// Removes the item at the specified index in the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/removeItem(at:)
func (t_ Toolbar) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeItemAtIndex:"), index)
}

// Removes the item with matching in the receiving toolbar. If multiple items share the same identifier (as is the case with space items) all matching items will be removed. To remove only a single space item, use instead.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/removeItem(identifier:)
func (t_ Toolbar) RemoveItemWithItemIdentifier(itemIdentifier unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeItemWithItemIdentifier:"), itemIdentifier)
}

// Displays the toolbar’s customization palette and handles any user-initiated customizations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/runCustomizationPalette(_:)
func (t_ Toolbar) RunCustomizationPalette(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("runCustomizationPalette:"), sender)
}

// Specifies the new configuration details for the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/setConfiguration(_:)
func (t_ Toolbar) SetConfigurationFromDictionary(configDict unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setConfigurationFromDictionary:"), configDict)
}

// Validates the toolbar’s visible items during a window update.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/validateVisibleItems()
func (t_ Toolbar) ValidateVisibleItems() {
	objc.Send[objc.ID](t_.ID, objc.Sel("validateVisibleItems"))
}


