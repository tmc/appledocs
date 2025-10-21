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
	ToolbarClass     _ToolbarClass
	ToolbarClassOnce sync.Once
)

func getToolbarClass() _ToolbarClass {
	ToolbarClassOnce.Do(func() {
		ToolbarClass = _ToolbarClass{objc.GetClass("NSToolbar")}
	})
	return ToolbarClass
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
// An object manages the controls and views that apply to the main window’s content area. Toolbars provide convenient access to the commands and features people use most often. Toolbars are also user-configurable and support the display of an interactive customization palette. Create and configure your toolbar programmatically or using Interface Builder. Add items to the toolbar that correspond to the commands you want to feature in your window. Each item has a corresponding object, which you use to make changes. Each toolbar manages a unique set of items, but you can synchronize the items and state of multiple toolbars by assigning the same value to their properties. For more information about how to use toolbars, see .
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

// Whether or not the user is allowed to change display modes at run time. This functionality is independent of customizing the order of the items themselves. Only disable when the functionality or legibility of your toolbar could not be improved by another display mode. The user’s selection will be persisted using the toolbar’s when is enabled. The default is YES for apps linked on macOS 15.0 and above.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsDisplayModeCustomization
func (t_ Toolbar) AllowsDisplayModeCustomization() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDisplayModeCustomization"))
	return rv
}

// SetAllowsDisplayModeCustomization sets the value of the allowsDisplayModeCustomization property.
// Whether or not the user is allowed to change display modes at run time. This functionality is independent of customizing the order of the items themselves. Only disable when the functionality or legibility of your toolbar could not be improved by another display mode. The user’s selection will be persisted using the toolbar’s when is enabled. The default is YES for apps linked on macOS 15.0 and above.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsDisplayModeCustomization
func (t_ Toolbar) SetAllowsDisplayModeCustomization(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDisplayModeCustomization:"), value)
}

// A Boolean value that indicates whether the toolbar can add items for Action extensions.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsExtensionItems
func (t_ Toolbar) AllowsExtensionItems() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsExtensionItems"))
	return rv
}

// SetAllowsExtensionItems sets the value of the allowsExtensionItems property.
// A Boolean value that indicates whether the toolbar can add items for Action extensions.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsExtensionItems
func (t_ Toolbar) SetAllowsExtensionItems(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsExtensionItems:"), value)
}

// A Boolean value that indicates whether users can modify the contents of the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsUserCustomization
func (t_ Toolbar) AllowsUserCustomization() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsUserCustomization"))
	return rv
}

// SetAllowsUserCustomization sets the value of the allowsUserCustomization property.
// A Boolean value that indicates whether users can modify the contents of the toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsUserCustomization
func (t_ Toolbar) SetAllowsUserCustomization(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsUserCustomization:"), value)
}

// A Boolean value that indicates whether the toolbar autosaves its configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/autosavesConfiguration
func (t_ Toolbar) AutosavesConfiguration() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autosavesConfiguration"))
	return rv
}

// SetAutosavesConfiguration sets the value of the autosavesConfiguration property.
// A Boolean value that indicates whether the toolbar autosaves its configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/autosavesConfiguration
func (t_ Toolbar) SetAutosavesConfiguration(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutosavesConfiguration:"), value)
}

// The item to display in the center of the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifier
func (t_ Toolbar) CenteredItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("centeredItemIdentifier"))
	return rv
}

// SetCenteredItemIdentifier sets the value of the centeredItemIdentifier property.
// The item to display in the center of the toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifier
func (t_ Toolbar) SetCenteredItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCenteredItemIdentifier:"), value)
}

// The set of custom items to display in the center of the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifiers
func (t_ Toolbar) CenteredItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("centeredItemIdentifiers"))
	return rv
}

// SetCenteredItemIdentifiers sets the value of the centeredItemIdentifiers property.
// The set of custom items to display in the center of the toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifiers
func (t_ Toolbar) SetCenteredItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCenteredItemIdentifiers:"), value)
}

// A dictionary containing the current configuration details for the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/configuration
func (t_ Toolbar) ConfigurationDictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("configurationDictionary"))
	return rv
}

// A Boolean value that indicates whether the toolbar’s customization palette is in use.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/customizationPaletteIsRunning
func (t_ Toolbar) CustomizationPaletteIsRunning() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("customizationPaletteIsRunning"))
	return rv
}

// The object you use to customize the toolbar contents and configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/delegate
func (t_ Toolbar) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}

// SetDelegate sets the value of the delegate property.
// The object you use to customize the toolbar contents and configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/delegate
func (t_ Toolbar) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// A value that indicates whether the toolbar displays items using a name, icon, or combination of elements.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/displayMode-swift.property
func (t_ Toolbar) DisplayMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("displayMode"))
	return rv
}

// SetDisplayMode sets the value of the displayMode property.
// A value that indicates whether the toolbar displays items using a name, icon, or combination of elements.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/displayMode-swift.property
func (t_ Toolbar) SetDisplayMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisplayMode:"), value)
}

// The toolbar’s full screen accessory view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryView
func (t_ Toolbar) FullScreenAccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("fullScreenAccessoryView"))
	return rv
}

// SetFullScreenAccessoryView sets the value of the fullScreenAccessoryView property.
// The toolbar’s full screen accessory view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryView
func (t_ Toolbar) SetFullScreenAccessoryView(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenAccessoryView:"), value)
}

// The maximum height of the toolbar’s full screen accessory view, in points.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMaxHeight
func (t_ Toolbar) FullScreenAccessoryViewMaxHeight() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("fullScreenAccessoryViewMaxHeight"))
	return rv
}

// SetFullScreenAccessoryViewMaxHeight sets the value of the fullScreenAccessoryViewMaxHeight property.
// The maximum height of the toolbar’s full screen accessory view, in points.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMaxHeight
func (t_ Toolbar) SetFullScreenAccessoryViewMaxHeight(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenAccessoryViewMaxHeight:"), value)
}

// The minimum height of the toolbar’s full screen accessory view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMinHeight
func (t_ Toolbar) FullScreenAccessoryViewMinHeight() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("fullScreenAccessoryViewMinHeight"))
	return rv
}

// SetFullScreenAccessoryViewMinHeight sets the value of the fullScreenAccessoryViewMinHeight property.
// The minimum height of the toolbar’s full screen accessory view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMinHeight
func (t_ Toolbar) SetFullScreenAccessoryViewMinHeight(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenAccessoryViewMinHeight:"), value)
}

// The value you use to identify the toolbar in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/identifier-swift.property
func (t_ Toolbar) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value that indicates whether the toolbar is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/isVisible
func (t_ Toolbar) Visible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("visible"))
	return rv
}

// SetVisible sets the value of the visible property.
// A Boolean value that indicates whether the toolbar is visible.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/isVisible
func (t_ Toolbar) SetVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisible:"), value)
}

// An array of itemIdentifiers that represent the current items in the toolbar. Setting this property will set the current items in the toolbar by diffing against items that already exist. Use this with great caution if is enabled as it will override any customizations the user has made. This property is key value observable.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/itemIdentifiers
func (t_ Toolbar) ItemIdentifiers() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("itemIdentifiers"))
	return rv
}

// SetItemIdentifiers sets the value of the itemIdentifiers property.
// An array of itemIdentifiers that represent the current items in the toolbar. Setting this property will set the current items in the toolbar by diffing against items that already exist. Use this with great caution if is enabled as it will override any customizations the user has made. This property is key value observable.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/itemIdentifiers
func (t_ Toolbar) SetItemIdentifiers(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setItemIdentifiers:"), nsArray)
}

// An array containing the toolbar’s current items, in order.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/items
func (t_ Toolbar) Items() []ToolbarItem {
	rv := objc.Send[[]ToolbarItem](t_.ID, objc.Sel("items"))
	return rv
}

// The identifier of the toolbar’s currently selected item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/selectedItemIdentifier
func (t_ Toolbar) SelectedItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("selectedItemIdentifier"))
	return rv
}

// SetSelectedItemIdentifier sets the value of the selectedItemIdentifier property.
// The identifier of the toolbar’s currently selected item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/selectedItemIdentifier
func (t_ Toolbar) SetSelectedItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedItemIdentifier:"), value)
}

// A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/showsBaselineSeparator
func (t_ Toolbar) ShowsBaselineSeparator() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("showsBaselineSeparator"))
	return rv
}

// SetShowsBaselineSeparator sets the value of the showsBaselineSeparator property.
// A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/showsBaselineSeparator
func (t_ Toolbar) SetShowsBaselineSeparator(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShowsBaselineSeparator:"), value)
}

// The toolbar’s size mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/sizeMode-swift.property
func (t_ Toolbar) SizeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("sizeMode"))
	return rv
}

// SetSizeMode sets the value of the sizeMode property.
// The toolbar’s size mode.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/sizeMode-swift.property
func (t_ Toolbar) SetSizeMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSizeMode:"), value)
}

// An array containing the toolbar’s currently visible items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/visibleItems
func (t_ Toolbar) VisibleItems() []ToolbarItem {
	rv := objc.Send[[]ToolbarItem](t_.ID, objc.Sel("visibleItems"))
	return rv
}
