
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Toolbar] class.
var ToolbarClass _ToolbarClass

func init() {
	ToolbarClass = _ToolbarClass{objc.GetClass("NSToolbar")}
}

type _ToolbarClass struct {
	objc.Class
}

// An interface definition for the [Toolbar] class.
type IToolbar interface {
	ID() objc.ID
	InsertItemWithItemIdentifierAtIndex(itemIdentifier unsafe.Pointer, index int)
	RemoveItemAtIndex(index int)
	RemoveItemWithItemIdentifier(itemIdentifier unsafe.Pointer)
	RunCustomizationPalette(sender objc.ID)
	SetConfigurationFromDictionary(configDict unsafe.Pointer)
	ValidateVisibleItems()
}

type Toolbar struct {
	id objc.ID
}

func ToolbarFrom(ptr unsafe.Pointer) Toolbar {
	return Toolbar{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ Toolbar) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _ToolbarClass) Alloc() Toolbar {
	rv := objc.Send[Toolbar](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _ToolbarClass) New() Toolbar {
	rv := objc.Send[Toolbar](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewToolbar creates and returns a new initialized instance.
func NewToolbar() Toolbar {
	return ToolbarClass.New()
}

// Init initializes the instance.
func (t_ Toolbar) Init() Toolbar {
	rv := objc.Send[Toolbar](t_.ID(), selInit)
	return rv
}
// Inserts an item into the toolbar at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/insertItem(withItemIdentifier:at:)
func (t_ Toolbar) InsertItemWithItemIdentifierAtIndex(itemIdentifier unsafe.Pointer, index int) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("insertItemWithItemIdentifier:atIndex:"), itemIdentifier, index)
}
// Removes the item at the specified index in the toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/removeItem(at:)
func (t_ Toolbar) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("removeItemAtIndex:"), index)
}
// Removes the item with matching   in the receiving toolbar. If multiple items share the same identifier (as is the case with space items) all matching items will be removed. To remove only a single space item, use   instead. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/removeItem(identifier:)
func (t_ Toolbar) RemoveItemWithItemIdentifier(itemIdentifier unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("removeItemWithItemIdentifier:"), itemIdentifier)
}
// Displays the toolbar’s customization palette and handles any user-initiated customizations. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/runCustomizationPalette(_:)
func (t_ Toolbar) RunCustomizationPalette(sender objc.ID) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("runCustomizationPalette:"), sender)
}
// Specifies the new configuration details for the toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/setConfiguration(_:)
func (t_ Toolbar) SetConfigurationFromDictionary(configDict unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setConfigurationFromDictionary:"), configDict)
}
// Validates the toolbar’s visible items during a window update. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/validateVisibleItems()
func (t_ Toolbar) ValidateVisibleItems() {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("validateVisibleItems"))
}
// Whether or not the user is allowed to change display modes at run time.   This functionality is independent of customizing the order of the items themselves.   Only disable when the functionality or legibility of your toolbar could not be improved by another display mode.   The user’s selection will be persisted using the toolbar’s   when   is enabled.   The default is YES for apps linked on macOS 15.0 and above. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/allowsDisplayModeCustomization
func (t_ Toolbar) AllowsDisplayModeCustomization() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("allowsDisplayModeCustomization"))
	return rv
}
// SetAllowsDisplayModeCustomization sets the value of the allowsDisplayModeCustomization property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/allowsDisplayModeCustomization
func (t_ Toolbar) SetAllowsDisplayModeCustomization(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setAllowsDisplayModeCustomization:"), value)
}
// A Boolean value that indicates whether the toolbar can add items for Action extensions. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/allowsExtensionItems
func (t_ Toolbar) AllowsExtensionItems() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("allowsExtensionItems"))
	return rv
}
// SetAllowsExtensionItems sets the value of the allowsExtensionItems property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/allowsExtensionItems
func (t_ Toolbar) SetAllowsExtensionItems(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setAllowsExtensionItems:"), value)
}
// A Boolean value that indicates whether users can modify the contents of the toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/allowsUserCustomization
func (t_ Toolbar) AllowsUserCustomization() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("allowsUserCustomization"))
	return rv
}
// SetAllowsUserCustomization sets the value of the allowsUserCustomization property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/allowsUserCustomization
func (t_ Toolbar) SetAllowsUserCustomization(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setAllowsUserCustomization:"), value)
}
// A Boolean value that indicates whether the toolbar autosaves its configuration. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/autosavesConfiguration
func (t_ Toolbar) AutosavesConfiguration() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("autosavesConfiguration"))
	return rv
}
// SetAutosavesConfiguration sets the value of the autosavesConfiguration property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/autosavesConfiguration
func (t_ Toolbar) SetAutosavesConfiguration(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setAutosavesConfiguration:"), value)
}
// The item to display in the center of the toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/centeredItemIdentifier
func (t_ Toolbar) CenteredItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("centeredItemIdentifier"))
	return rv
}
// SetCenteredItemIdentifier sets the value of the centeredItemIdentifier property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/centeredItemIdentifier
func (t_ Toolbar) SetCenteredItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setCenteredItemIdentifier:"), value)
}
// The set of custom items to display in the center of the toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/centeredItemIdentifiers
func (t_ Toolbar) CenteredItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("centeredItemIdentifiers"))
	return rv
}
// SetCenteredItemIdentifiers sets the value of the centeredItemIdentifiers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/centeredItemIdentifiers
func (t_ Toolbar) SetCenteredItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setCenteredItemIdentifiers:"), value)
}
// A dictionary containing the current configuration details for the toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/configuration
func (t_ Toolbar) ConfigurationDictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("configurationDictionary"))
	return rv
}
// A Boolean value that indicates whether the toolbar’s customization palette is in use. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/customizationPaletteIsRunning
func (t_ Toolbar) CustomizationPaletteIsRunning() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("customizationPaletteIsRunning"))
	return rv
}
// The object you use to customize the toolbar contents and configuration. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/delegate
func (t_ Toolbar) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/delegate
func (t_ Toolbar) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setDelegate:"), value)
}
// A value that indicates whether the toolbar displays items using a name, icon, or combination of elements. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/displayMode-swift.property
func (t_ Toolbar) DisplayMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("displayMode"))
	return rv
}
// SetDisplayMode sets the value of the displayMode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/displayMode-swift.property
func (t_ Toolbar) SetDisplayMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setDisplayMode:"), value)
}
// The toolbar’s full screen accessory view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/fullScreenAccessoryView
func (t_ Toolbar) FullScreenAccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("fullScreenAccessoryView"))
	return rv
}
// SetFullScreenAccessoryView sets the value of the fullScreenAccessoryView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/fullScreenAccessoryView
func (t_ Toolbar) SetFullScreenAccessoryView(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setFullScreenAccessoryView:"), value)
}
// The maximum height of the toolbar’s full screen accessory view, in points. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMaxHeight
func (t_ Toolbar) FullScreenAccessoryViewMaxHeight() float64 {
	rv := objc.Send[float64](t_.ID(), objc.RegisterName("fullScreenAccessoryViewMaxHeight"))
	return rv
}
// SetFullScreenAccessoryViewMaxHeight sets the value of the fullScreenAccessoryViewMaxHeight property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMaxHeight
func (t_ Toolbar) SetFullScreenAccessoryViewMaxHeight(value float64) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setFullScreenAccessoryViewMaxHeight:"), value)
}
// The minimum height of the toolbar’s full screen accessory view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMinHeight
func (t_ Toolbar) FullScreenAccessoryViewMinHeight() float64 {
	rv := objc.Send[float64](t_.ID(), objc.RegisterName("fullScreenAccessoryViewMinHeight"))
	return rv
}
// SetFullScreenAccessoryViewMinHeight sets the value of the fullScreenAccessoryViewMinHeight property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMinHeight
func (t_ Toolbar) SetFullScreenAccessoryViewMinHeight(value float64) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setFullScreenAccessoryViewMinHeight:"), value)
}
// The value you use to identify the toolbar in your app. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/identifier-swift.property
func (t_ Toolbar) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("identifier"))
	return rv
}
// A Boolean value that indicates whether the toolbar is visible. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/isVisible
func (t_ Toolbar) Visible() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("visible"))
	return rv
}
// SetVisible sets the value of the visible property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/isVisible
func (t_ Toolbar) SetVisible(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setVisible:"), value)
}
// An array of itemIdentifiers that represent the current items in the toolbar.   Setting this property will set the current items in the toolbar by diffing against items that already exist.   Use this with great caution if   is enabled as it will override any customizations the user has made.   This property is key value observable. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/itemIdentifiers
func (t_ Toolbar) ItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("itemIdentifiers"))
	return rv
}
// SetItemIdentifiers sets the value of the itemIdentifiers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/itemIdentifiers
func (t_ Toolbar) SetItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setItemIdentifiers:"), value)
}
// An array containing the toolbar’s current items, in order. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/items
func (t_ Toolbar) Items() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("items"))
	return rv
}
// The identifier of the toolbar’s currently selected item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/selectedItemIdentifier
func (t_ Toolbar) SelectedItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("selectedItemIdentifier"))
	return rv
}
// SetSelectedItemIdentifier sets the value of the selectedItemIdentifier property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/selectedItemIdentifier
func (t_ Toolbar) SetSelectedItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setSelectedItemIdentifier:"), value)
}
// A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/showsBaselineSeparator
func (t_ Toolbar) ShowsBaselineSeparator() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("showsBaselineSeparator"))
	return rv
}
// SetShowsBaselineSeparator sets the value of the showsBaselineSeparator property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/showsBaselineSeparator
func (t_ Toolbar) SetShowsBaselineSeparator(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setShowsBaselineSeparator:"), value)
}
// The toolbar’s size mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/sizeMode-swift.property
func (t_ Toolbar) SizeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("sizeMode"))
	return rv
}
// SetSizeMode sets the value of the sizeMode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/sizeMode-swift.property
func (t_ Toolbar) SetSizeMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setSizeMode:"), value)
}
// An array containing the toolbar’s currently visible items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbar/visibleItems
func (t_ Toolbar) VisibleItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("visibleItems"))
	return rv
}
