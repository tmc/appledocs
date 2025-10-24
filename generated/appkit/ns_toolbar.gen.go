// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSToolbar */


/* debug [class_header]: Header for NSToolbar */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Toolbar */
// An interface definition for the [Toolbar] class.
type IToolbar interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Toolbar */
	// properties:
	AllowsDisplayModeCustomization() bool
	SetAllowsDisplayModeCustomization(value bool)
	AllowsExtensionItems() bool
	SetAllowsExtensionItems(value bool)
	AllowsUserCustomization() bool
	SetAllowsUserCustomization(value bool)
	AutosavesConfiguration() bool
	SetAutosavesConfiguration(value bool)
	CenteredItemIdentifier() ToolbarItemIdentifier /* typedef */
	SetCenteredItemIdentifier(value ToolbarItemIdentifier /* typedef */)
	CenteredItemIdentifiers() unsafe.Pointer
	SetCenteredItemIdentifiers(value unsafe.Pointer)
	ConfigurationDictionary() foundation.IDictionary
	CustomizationPaletteIsRunning() bool
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisplayMode() ToolbarDisplayMode
	SetDisplayMode(value ToolbarDisplayMode)
	FullScreenAccessoryView() IView
	SetFullScreenAccessoryView(value IView)
	FullScreenAccessoryViewMaxHeight() float64
	SetFullScreenAccessoryViewMaxHeight(value float64)
	FullScreenAccessoryViewMinHeight() float64
	SetFullScreenAccessoryViewMinHeight(value float64)
	Identifier() ToolbarIdentifier /* typedef */
	Visible() bool
	SetVisible(value bool)
	ItemIdentifiers() []string
	SetItemIdentifiers(value []string)
	Items() []ToolbarItem
	SelectedItemIdentifier() ToolbarItemIdentifier /* typedef */
	SetSelectedItemIdentifier(value ToolbarItemIdentifier /* typedef */)
	ShowsBaselineSeparator() bool
	SetShowsBaselineSeparator(value bool)
	SizeMode() ToolbarSizeMode
	SetSizeMode(value ToolbarSizeMode)
	VisibleItems() []ToolbarItem
	Configuration() objc.IObject /* cross-framework: NSString */
	SetConfiguration(value objc.IObject /* cross-framework: NSString */)
	IsVisible() bool
	SetIsVisible(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Toolbar */
	// methods:
	InsertItemWithItemIdentifierAtIndex(itemIdentifier ToolbarItemIdentifier /* typedef */, index int)
	RemoveItemAtIndex(index int)
	RemoveItemWithItemIdentifier(itemIdentifier ToolbarItemIdentifier /* typedef */)
	RunCustomizationPalette(sender objc.IObject)
	ValidateVisibleItems()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Toolbar */
// Alloc allocates a new instance without initialization.
func (tc _ToolbarClass) Alloc() Toolbar {
	rv := objc.Send[Toolbar](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Toolbar */
// An object that manages the space above your app’s custom content and either below or integrated with the window’s title bar.
//
// An object manages the controls and views that apply to the main window’s content area. Toolbars provide convenient access to the commands and features people use most often. Toolbars are also user-configurable and support the display of an interactive customization palette. Create and configure your toolbar programmatically or using Interface Builder. Add items to the toolbar that correspond to the commands you want to feature in your window. Each item has a corresponding object, which you use to make changes. Each toolbar manages a unique set of items, but you can synchronize the items and state of multiple toolbars by assigning the same value to their properties. For more information about how to use toolbars, see .


// An object that manages the space above your app’s custom content and either below or integrated with the window’s title bar.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Toolbar */

// Creates a newly allocated toolbar with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/init(identifier:)
func NewToolbarWithIdentifier(identifier ToolbarIdentifier /* typedef */) Toolbar {
	instance := getToolbarClass().Alloc()
	rv := objc.Send[Toolbar](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewToolbarWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Toolbar */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Toolbar */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Toolbar */

// Inserts an item into the toolbar at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/insertItem(withItemIdentifier:at:)
func (t_ Toolbar) InsertItemWithItemIdentifierAtIndex(itemIdentifier ToolbarItemIdentifier /* typedef */, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertItemWithItemIdentifier:atIndex:"), itemIdentifier, index)
}/* debug [instance_methods/method]: InsertItemWithItemIdentifierAtIndex */


// Removes the item at the specified index in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/removeItem(at:)
func (t_ Toolbar) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeItemAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveItemAtIndex */


// Removes the item with matching in the receiving toolbar. If multiple items share the same identifier (as is the case with space items) all matching items will be removed. To remove only a single space item, use instead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/removeItem(identifier:)
func (t_ Toolbar) RemoveItemWithItemIdentifier(itemIdentifier ToolbarItemIdentifier /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeItemWithItemIdentifier:"), itemIdentifier)
}/* debug [instance_methods/method]: RemoveItemWithItemIdentifier */


// Displays the toolbar’s customization palette and handles any user-initiated customizations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/runCustomizationPalette(_:)
func (t_ Toolbar) RunCustomizationPalette(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("runCustomizationPalette:"), sender)
}/* debug [instance_methods/method]: RunCustomizationPalette */


// Validates the toolbar’s visible items during a window update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/validateVisibleItems()
func (t_ Toolbar) ValidateVisibleItems() {
	objc.Send[objc.ID](t_.ID, objc.Sel("validateVisibleItems"))
}/* debug [instance_methods/method]: ValidateVisibleItems */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Toolbar */

// Whether or not the user is allowed to change display modes at run time. This functionality is independent of customizing the order of the items themselves. Only disable when the functionality or legibility of your toolbar could not be improved by another display mode. The user’s selection will be persisted using the toolbar’s when is enabled. The default is YES for apps linked on macOS 15.0 and above.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsDisplayModeCustomization
func (t_ Toolbar) AllowsDisplayModeCustomization() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDisplayModeCustomization"))
	return rv
}/* debug [instance_properties/getter]: allowsDisplayModeCustomization */


// Whether or not the user is allowed to change display modes at run time. This functionality is independent of customizing the order of the items themselves. Only disable when the functionality or legibility of your toolbar could not be improved by another display mode. The user’s selection will be persisted using the toolbar’s when is enabled. The default is YES for apps linked on macOS 15.0 and above.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsDisplayModeCustomization
func (t_ Toolbar) SetAllowsDisplayModeCustomization(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDisplayModeCustomization:"), value)
}/* debug [instance_properties/setter]: allowsDisplayModeCustomization */


// A Boolean value that indicates whether the toolbar can add items for Action extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsExtensionItems
func (t_ Toolbar) AllowsExtensionItems() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsExtensionItems"))
	return rv
}/* debug [instance_properties/getter]: allowsExtensionItems */


// A Boolean value that indicates whether the toolbar can add items for Action extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsExtensionItems
func (t_ Toolbar) SetAllowsExtensionItems(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsExtensionItems:"), value)
}/* debug [instance_properties/setter]: allowsExtensionItems */


// A Boolean value that indicates whether users can modify the contents of the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsUserCustomization
func (t_ Toolbar) AllowsUserCustomization() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsUserCustomization"))
	return rv
}/* debug [instance_properties/getter]: allowsUserCustomization */


// A Boolean value that indicates whether users can modify the contents of the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsUserCustomization
func (t_ Toolbar) SetAllowsUserCustomization(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsUserCustomization:"), value)
}/* debug [instance_properties/setter]: allowsUserCustomization */


// A Boolean value that indicates whether the toolbar autosaves its configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/autosavesConfiguration
func (t_ Toolbar) AutosavesConfiguration() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autosavesConfiguration"))
	return rv
}/* debug [instance_properties/getter]: autosavesConfiguration */


// A Boolean value that indicates whether the toolbar autosaves its configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/autosavesConfiguration
func (t_ Toolbar) SetAutosavesConfiguration(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutosavesConfiguration:"), value)
}/* debug [instance_properties/setter]: autosavesConfiguration */


// The item to display in the center of the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifier
func (t_ Toolbar) CenteredItemIdentifier() ToolbarItemIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("centeredItemIdentifier"))
	return rv
}/* debug [instance_properties/getter]: centeredItemIdentifier */


// The item to display in the center of the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifier
func (t_ Toolbar) SetCenteredItemIdentifier(value ToolbarItemIdentifier /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCenteredItemIdentifier:"), value)
}/* debug [instance_properties/setter]: centeredItemIdentifier */


// The set of custom items to display in the center of the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifiers
func (t_ Toolbar) CenteredItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("centeredItemIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: centeredItemIdentifiers */


// The set of custom items to display in the center of the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifiers
func (t_ Toolbar) SetCenteredItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCenteredItemIdentifiers:"), value)
}/* debug [instance_properties/setter]: centeredItemIdentifiers */


// A dictionary containing the current configuration details for the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/configuration
func (t_ Toolbar) ConfigurationDictionary() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("configurationDictionary"))
	return rv
}/* debug [instance_properties/getter]: configurationDictionary */


// A Boolean value that indicates whether the toolbar’s customization palette is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/customizationPaletteIsRunning
func (t_ Toolbar) CustomizationPaletteIsRunning() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("customizationPaletteIsRunning"))
	return rv
}/* debug [instance_properties/getter]: customizationPaletteIsRunning */


// The object you use to customize the toolbar contents and configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/delegate
func (t_ Toolbar) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The object you use to customize the toolbar contents and configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/delegate
func (t_ Toolbar) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A value that indicates whether the toolbar displays items using a name, icon, or combination of elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/displayMode-swift.property
func (t_ Toolbar) DisplayMode() ToolbarDisplayMode {
	rv := objc.Send[ToolbarDisplayMode](t_.ID, objc.Sel("displayMode"))
	return rv
}/* debug [instance_properties/getter]: displayMode */


// A value that indicates whether the toolbar displays items using a name, icon, or combination of elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/displayMode-swift.property
func (t_ Toolbar) SetDisplayMode(value ToolbarDisplayMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisplayMode:"), value)
}/* debug [instance_properties/setter]: displayMode */


// The toolbar’s full screen accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryView
func (t_ Toolbar) FullScreenAccessoryView() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("fullScreenAccessoryView"))
	return rv
}/* debug [instance_properties/getter]: fullScreenAccessoryView */


// The toolbar’s full screen accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryView
func (t_ Toolbar) SetFullScreenAccessoryView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenAccessoryView:"), value)
}/* debug [instance_properties/setter]: fullScreenAccessoryView */


// The maximum height of the toolbar’s full screen accessory view, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMaxHeight
func (t_ Toolbar) FullScreenAccessoryViewMaxHeight() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("fullScreenAccessoryViewMaxHeight"))
	return rv
}/* debug [instance_properties/getter]: fullScreenAccessoryViewMaxHeight */


// The maximum height of the toolbar’s full screen accessory view, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMaxHeight
func (t_ Toolbar) SetFullScreenAccessoryViewMaxHeight(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenAccessoryViewMaxHeight:"), value)
}/* debug [instance_properties/setter]: fullScreenAccessoryViewMaxHeight */


// The minimum height of the toolbar’s full screen accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMinHeight
func (t_ Toolbar) FullScreenAccessoryViewMinHeight() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("fullScreenAccessoryViewMinHeight"))
	return rv
}/* debug [instance_properties/getter]: fullScreenAccessoryViewMinHeight */


// The minimum height of the toolbar’s full screen accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryViewMinHeight
func (t_ Toolbar) SetFullScreenAccessoryViewMinHeight(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenAccessoryViewMinHeight:"), value)
}/* debug [instance_properties/setter]: fullScreenAccessoryViewMinHeight */


// The value you use to identify the toolbar in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/identifier-swift.property
func (t_ Toolbar) Identifier() ToolbarIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that indicates whether the toolbar is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/isVisible
func (t_ Toolbar) Visible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("visible"))
	return rv
}/* debug [instance_properties/getter]: visible */


// A Boolean value that indicates whether the toolbar is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/isVisible
func (t_ Toolbar) SetVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisible:"), value)
}/* debug [instance_properties/setter]: visible */


// An array of itemIdentifiers that represent the current items in the toolbar. Setting this property will set the current items in the toolbar by diffing against items that already exist. Use this with great caution if is enabled as it will override any customizations the user has made. This property is key value observable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/itemIdentifiers
func (t_ Toolbar) ItemIdentifiers() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("itemIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: itemIdentifiers */


// An array of itemIdentifiers that represent the current items in the toolbar. Setting this property will set the current items in the toolbar by diffing against items that already exist. Use this with great caution if is enabled as it will override any customizations the user has made. This property is key value observable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/itemIdentifiers
func (t_ Toolbar) SetItemIdentifiers(value []string) {
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
}/* debug [instance_properties/setter]: itemIdentifiers */


// An array containing the toolbar’s current items, in order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/items
func (t_ Toolbar) Items() []ToolbarItem {
	rv := objc.Send[[]ToolbarItem](t_.ID, objc.Sel("items"))
	return rv
}/* debug [instance_properties/getter]: items */


// The identifier of the toolbar’s currently selected item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/selectedItemIdentifier
func (t_ Toolbar) SelectedItemIdentifier() ToolbarItemIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("selectedItemIdentifier"))
	return rv
}/* debug [instance_properties/getter]: selectedItemIdentifier */


// The identifier of the toolbar’s currently selected item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/selectedItemIdentifier
func (t_ Toolbar) SetSelectedItemIdentifier(value ToolbarItemIdentifier /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedItemIdentifier:"), value)
}/* debug [instance_properties/setter]: selectedItemIdentifier */


// A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/showsBaselineSeparator
func (t_ Toolbar) ShowsBaselineSeparator() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("showsBaselineSeparator"))
	return rv
}/* debug [instance_properties/getter]: showsBaselineSeparator */


// A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/showsBaselineSeparator
func (t_ Toolbar) SetShowsBaselineSeparator(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShowsBaselineSeparator:"), value)
}/* debug [instance_properties/setter]: showsBaselineSeparator */


// The toolbar’s size mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/sizeMode-swift.property
func (t_ Toolbar) SizeMode() ToolbarSizeMode {
	rv := objc.Send[ToolbarSizeMode](t_.ID, objc.Sel("sizeMode"))
	return rv
}/* debug [instance_properties/getter]: sizeMode */


// The toolbar’s size mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/sizeMode-swift.property
func (t_ Toolbar) SetSizeMode(value ToolbarSizeMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSizeMode:"), value)
}/* debug [instance_properties/setter]: sizeMode */


// An array containing the toolbar’s currently visible items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/visibleItems
func (t_ Toolbar) VisibleItems() []ToolbarItem {
	rv := objc.Send[[]ToolbarItem](t_.ID, objc.Sel("visibleItems"))
	return rv
}/* debug [instance_properties/getter]: visibleItems */


// A dictionary containing the current configuration details for the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/configuration
func (t_ Toolbar) Configuration() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// A dictionary containing the current configuration details for the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/configuration
func (t_ Toolbar) SetConfiguration(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */


// A Boolean value that indicates whether the toolbar is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/isvisible
func (t_ Toolbar) IsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVisible"))
	return rv
}/* debug [instance_properties/getter]: isVisible */


// A Boolean value that indicates whether the toolbar is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/isvisible
func (t_ Toolbar) SetIsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsVisible:"), value)
}/* debug [instance_properties/setter]: isVisible */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSToolbar */


