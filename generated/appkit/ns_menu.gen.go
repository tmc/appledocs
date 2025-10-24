// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMenu */


/* debug [class_header]: Header for NSMenu */
// The class instance for the [Menu] class.
var (
	MenuClass     _MenuClass
	MenuClassOnce sync.Once
)

func getMenuClass() _MenuClass {
	MenuClassOnce.Do(func() {
		MenuClass = _MenuClass{objc.GetClass("NSMenu")}
	})
	return MenuClass
}

type _MenuClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Menu */
// An interface definition for the [Menu] class.
type IMenu interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Menu */
	// properties:
	ItemArray() []MenuItem
	SetItemArray(value []MenuItem)
	NumberOfItems() int
	AllowsContextMenuPlugIns() bool
	SetAllowsContextMenuPlugIns(value bool)
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	AutomaticallyInsertsWritingToolsItems() bool
	SetAutomaticallyInsertsWritingToolsItems(value bool)
	Delegate() objc.IObject /* cross-framework: MenuDelegate */
	SetDelegate(value objc.IObject /* cross-framework: MenuDelegate */)
	Font() IFont
	SetFont(value IFont)
	HighlightedItem() IMenuItem
	SetHighlightedItem(value IMenuItem)
	IsTornOff() bool
	SetIsTornOff(value bool)
	Items() IMenuItem
	SetItems(value IMenuItem)
	MenuBarHeight() float64
	SetMenuBarHeight(value float64)
	MenuChangedMessagesEnabled() bool
	SetMenuChangedMessagesEnabled(value bool)
	MinimumWidth() float64
	SetMinimumWidth(value float64)
	PresentationStyle() objectivec.IObject
	SetPresentationStyle(value objectivec.IObject)
	PropertiesToUpdate() objectivec.IObject
	SetPropertiesToUpdate(value objectivec.IObject)
	SelectedItems() IMenuItem
	SetSelectedItems(value IMenuItem)
	SelectionMode() objectivec.IObject
	SetSelectionMode(value objectivec.IObject)
	ShowsStateColumn() bool
	SetShowsStateColumn(value bool)
	Size() Size /* not a class type */
	SetSize(value Size /* not a class type */)
	Supermenu() IMenu
	SetSupermenu(value IMenu)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Menu */
	// methods:
	AddItem(newItem IMenuItem)
	AddItemWithTitleActionKeyEquivalent(string_ objc.IObject /* cross-framework: NSString */, selector objc.SEL, charCode objc.IObject /* cross-framework: NSString */) IMenuItem
	IndexOfItem(item IMenuItem) int
	IndexOfItemWithRepresentedObject(object objc.IObject) int
	IndexOfItemWithSubmenu(submenu IMenu) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.SEL) int
	IndexOfItemWithTitle(title objc.IObject /* cross-framework: NSString */) int
	InsertItemAtIndex(newItem IMenuItem, index int)
	InsertItemWithTitleActionKeyEquivalentAtIndex(string_ objc.IObject /* cross-framework: NSString */, selector objc.SEL, charCode objc.IObject /* cross-framework: NSString */, index int) IMenuItem
	ItemAtIndex(index int) IMenuItem
	ItemWithTag(tag int) IMenuItem
	ItemWithTitle(title objc.IObject /* cross-framework: NSString */) IMenuItem
	ItemChanged(item IMenuItem)
	RemoveAllItems()
	RemoveItem(item IMenuItem)
	RemoveItemAtIndex(index int)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Menu */
// Alloc allocates a new instance without initialization.
func (mc _MenuClass) Alloc() Menu {
	rv := objc.Send[Menu](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MenuClass) New() Menu {
	rv := objc.Send[Menu](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Menu) Init() Menu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Menu) Autorelease() Menu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenu creates a new Menu instance.
func NewMenu() Menu {
	return getMenuClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Menu */
// An object that manages an app’s menus.


// An object that manages an app’s menus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu
type Menu struct {
	objectivec.Object
}

// MenuFrom constructs a [Menu] from an unsafe.Pointer.
//
// An object that manages an app’s menus.
func MenuFrom(ptr unsafe.Pointer) Menu {
	return Menu{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Menu *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Menu */

// Displays a contextual menu over a view for an event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (mc _MenuClass) PopUpContextMenuWithEventForView(menu IMenu, event IEvent, view IView) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:"), menu, event, view)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PopUpContextMenuWithEventForView) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Menu */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Menu */

// Adds a menu item to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/addItem(_:)
func (m_ Menu) AddItem(newItem IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addItem:"), newItem)
}/* debug [instance_methods/method]: AddItem */


// Creates a new menu item and adds it to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/addItem(withTitle:action:keyEquivalent:)
func (m_ Menu) AddItemWithTitleActionKeyEquivalent(string_ objc.IObject /* cross-framework: NSString */, selector objc.SEL, charCode objc.IObject /* cross-framework: NSString */) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("addItemWithTitle:action:keyEquivalent:"), string_, selector, charCode)
	return rv
}/* debug [instance_methods/method]: AddItemWithTitleActionKeyEquivalent */


// Returns the index identifying the location of a specified menu item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/index(of:)
func (m_ Menu) IndexOfItem(item IMenuItem) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItem:"), item)
	return rv
}/* debug [instance_methods/method]: IndexOfItem */


// Returns the index of the first menu item in the menu that has a given represented object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withRepresentedObject:)
func (m_ Menu) IndexOfItemWithRepresentedObject(object objc.IObject) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithRepresentedObject:"), object)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithRepresentedObject */


// Returns the index of the menu item in the menu with the given submenu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withSubmenu:)
func (m_ Menu) IndexOfItemWithSubmenu(submenu IMenu) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithSubmenu:"), submenu)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithSubmenu */


// Returns the index of the first menu item in the menu identified by a tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTag:)
func (m_ Menu) IndexOfItemWithTag(tag int) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithTag */


// Returns the index of the first menu item in the menu that has a specified action and target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTarget:andAction:)
func (m_ Menu) IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.SEL) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithTargetAndAction */


// Returns the index of the first menu item in the menu that has a specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTitle:)
func (m_ Menu) IndexOfItemWithTitle(title objc.IObject /* cross-framework: NSString */) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithTitle:"), title)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithTitle */


// Inserts a menu item into the menu at a specific location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/insertItem(_:at:)
func (m_ Menu) InsertItemAtIndex(newItem IMenuItem, index int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertItem:atIndex:"), newItem, index)
}/* debug [instance_methods/method]: InsertItemAtIndex */


// Creates and adds a menu item at a specified location in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/insertItem(withTitle:action:keyEquivalent:at:)
func (m_ Menu) InsertItemWithTitleActionKeyEquivalentAtIndex(string_ objc.IObject /* cross-framework: NSString */, selector objc.SEL, charCode objc.IObject /* cross-framework: NSString */, index int) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("insertItemWithTitle:action:keyEquivalent:atIndex:"), string_, selector, charCode, index)
	return rv
}/* debug [instance_methods/method]: InsertItemWithTitleActionKeyEquivalentAtIndex */


// Returns the menu item at a specific location of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/item(at:)
func (m_ Menu) ItemAtIndex(index int) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ItemAtIndex */


// Returns the first menu item in the menu with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/item(withTag:)
func (m_ Menu) ItemWithTag(tag int) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("itemWithTag:"), tag)
	return rv
}/* debug [instance_methods/method]: ItemWithTag */


// Returns the first menu item in the menu with a specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/item(withTitle:)
func (m_ Menu) ItemWithTitle(title objc.IObject /* cross-framework: NSString */) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("itemWithTitle:"), title)
	return rv
}/* debug [instance_methods/method]: ItemWithTitle */


// Invoked when a menu item is modified visually (for example, its title changes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/itemChanged(_:)
func (m_ Menu) ItemChanged(item IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("itemChanged:"), item)
}/* debug [instance_methods/method]: ItemChanged */


// Removes all the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/removeAllItems()
func (m_ Menu) RemoveAllItems() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllItems"))
}/* debug [instance_methods/method]: RemoveAllItems */


// Removes a menu item from the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/removeItem(_:)
func (m_ Menu) RemoveItem(item IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeItem:"), item)
}/* debug [instance_methods/method]: RemoveItem */


// Removes the menu item at a specified location in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/removeItem(at:)
func (m_ Menu) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeItemAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveItemAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Menu */

// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/items
func (m_ Menu) ItemArray() []MenuItem {
	rv := objc.Send[[]MenuItem](m_.ID, objc.Sel("itemArray"))
	return rv
}/* debug [instance_properties/getter]: itemArray */


// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/items
func (m_ Menu) SetItemArray(value []MenuItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setItemArray:"), nsArray)
}/* debug [instance_properties/setter]: itemArray */


// The number of menu items in the menu, including separator items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/numberOfItems
func (m_ Menu) NumberOfItems() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfItems"))
	return rv
}/* debug [instance_properties/getter]: numberOfItems */


// Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/allowscontextmenuplugins
func (m_ Menu) AllowsContextMenuPlugIns() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsContextMenuPlugIns"))
	return rv
}/* debug [instance_properties/getter]: allowsContextMenuPlugIns */


// Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/allowscontextmenuplugins
func (m_ Menu) SetAllowsContextMenuPlugIns(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsContextMenuPlugIns:"), value)
}/* debug [instance_properties/setter]: allowsContextMenuPlugIns */


// Indicates whether the menu automatically enables and disables its menu items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/autoenablesitems
func (m_ Menu) AutoenablesItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("autoenablesItems"))
	return rv
}/* debug [instance_properties/getter]: autoenablesItems */


// Indicates whether the menu automatically enables and disables its menu items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/autoenablesitems
func (m_ Menu) SetAutoenablesItems(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutoenablesItems:"), value)
}/* debug [instance_properties/setter]: autoenablesItems */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/automaticallyinsertswritingtoolsitems
func (m_ Menu) AutomaticallyInsertsWritingToolsItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("automaticallyInsertsWritingToolsItems"))
	return rv
}/* debug [instance_properties/getter]: automaticallyInsertsWritingToolsItems */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/automaticallyinsertswritingtoolsitems
func (m_ Menu) SetAutomaticallyInsertsWritingToolsItems(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutomaticallyInsertsWritingToolsItems:"), value)
}/* debug [instance_properties/setter]: automaticallyInsertsWritingToolsItems */


// The delegate of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/delegate
func (m_ Menu) Delegate() objc.IObject /* cross-framework: MenuDelegate */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/delegate
func (m_ Menu) SetDelegate(value objc.IObject /* cross-framework: MenuDelegate */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The font of the menu and its submenus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/font
func (m_ Menu) Font() IFont {
	rv := objc.Send[Font](m_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font of the menu and its submenus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/font
func (m_ Menu) SetFont(value IFont) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// Indicates the currently highlighted item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/highlighteditem
func (m_ Menu) HighlightedItem() IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("highlightedItem"))
	return rv
}/* debug [instance_properties/getter]: highlightedItem */


// Indicates the currently highlighted item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/highlighteditem
func (m_ Menu) SetHighlightedItem(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHighlightedItem:"), value)
}/* debug [instance_properties/setter]: highlightedItem */


// Indicates whether the menu is offscreen or attached to another menu (or if it’s the main menu).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/istornoff
func (m_ Menu) IsTornOff() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isTornOff"))
	return rv
}/* debug [instance_properties/getter]: isTornOff */


// Indicates whether the menu is offscreen or attached to another menu (or if it’s the main menu).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/istornoff
func (m_ Menu) SetIsTornOff(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsTornOff:"), value)
}/* debug [instance_properties/setter]: isTornOff */


// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/items
func (m_ Menu) Items() IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("items"))
	return rv
}/* debug [instance_properties/getter]: items */


// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/items
func (m_ Menu) SetItems(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setItems:"), value)
}/* debug [instance_properties/setter]: items */


// The menu bar height for the main menu in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/menubarheight
func (m_ Menu) MenuBarHeight() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("menuBarHeight"))
	return rv
}/* debug [instance_properties/getter]: menuBarHeight */


// The menu bar height for the main menu in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/menubarheight
func (m_ Menu) SetMenuBarHeight(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuBarHeight:"), value)
}/* debug [instance_properties/setter]: menuBarHeight */


// Indicates whether messages are sent to the application’s windows each time the menu changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/menuchangedmessagesenabled
func (m_ Menu) MenuChangedMessagesEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("menuChangedMessagesEnabled"))
	return rv
}/* debug [instance_properties/getter]: menuChangedMessagesEnabled */


// Indicates whether messages are sent to the application’s windows each time the menu changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/menuchangedmessagesenabled
func (m_ Menu) SetMenuChangedMessagesEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuChangedMessagesEnabled:"), value)
}/* debug [instance_properties/setter]: menuChangedMessagesEnabled */


// The minimum width of the menu in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/minimumwidth
func (m_ Menu) MinimumWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("minimumWidth"))
	return rv
}/* debug [instance_properties/getter]: minimumWidth */


// The minimum width of the menu in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/minimumwidth
func (m_ Menu) SetMinimumWidth(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumWidth:"), value)
}/* debug [instance_properties/setter]: minimumWidth */


// The presentation style of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/presentationstyle-swift.property
func (m_ Menu) PresentationStyle() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("presentationStyle"))
	return rv
}/* debug [instance_properties/getter]: presentationStyle */


// The presentation style of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/presentationstyle-swift.property
func (m_ Menu) SetPresentationStyle(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentationStyle:"), value)
}/* debug [instance_properties/setter]: presentationStyle */


// The available properties for the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/propertiestoupdate
func (m_ Menu) PropertiesToUpdate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("propertiesToUpdate"))
	return rv
}/* debug [instance_properties/getter]: propertiesToUpdate */


// The available properties for the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/propertiestoupdate
func (m_ Menu) SetPropertiesToUpdate(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPropertiesToUpdate:"), value)
}/* debug [instance_properties/setter]: propertiesToUpdate */


// The menu items that are currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selecteditems
func (m_ Menu) SelectedItems() IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("selectedItems"))
	return rv
}/* debug [instance_properties/getter]: selectedItems */


// The menu items that are currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selecteditems
func (m_ Menu) SetSelectedItems(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedItems:"), value)
}/* debug [instance_properties/setter]: selectedItems */


// The selection mode of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selectionmode-swift.property
func (m_ Menu) SelectionMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("selectionMode"))
	return rv
}/* debug [instance_properties/getter]: selectionMode */


// The selection mode of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selectionmode-swift.property
func (m_ Menu) SetSelectionMode(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectionMode:"), value)
}/* debug [instance_properties/setter]: selectionMode */


// Indicates whether the menu displays the state column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/showsstatecolumn
func (m_ Menu) ShowsStateColumn() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsStateColumn"))
	return rv
}/* debug [instance_properties/getter]: showsStateColumn */


// Indicates whether the menu displays the state column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/showsstatecolumn
func (m_ Menu) SetShowsStateColumn(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsStateColumn:"), value)
}/* debug [instance_properties/setter]: showsStateColumn */


// The size of the menu in screen coordinates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/size
func (m_ Menu) Size() Size /* not a class type */ {
	rv := objc.Send[Size](m_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The size of the menu in screen coordinates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/size
func (m_ Menu) SetSize(value Size /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */


// The parent menu that contains the menu as a submenu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/supermenu
func (m_ Menu) Supermenu() IMenu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("supermenu"))
	return rv
}/* debug [instance_properties/getter]: supermenu */


// The parent menu that contains the menu as a submenu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/supermenu
func (m_ Menu) SetSupermenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupermenu:"), value)
}/* debug [instance_properties/setter]: supermenu */


// The title of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/title
func (m_ Menu) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/title
func (m_ Menu) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// Configures the layout direction of menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/userinterfacelayoutdirection
func (m_ Menu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](m_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}/* debug [instance_properties/getter]: userInterfaceLayoutDirection */


// Configures the layout direction of menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/userinterfacelayoutdirection
func (m_ Menu) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}/* debug [instance_properties/setter]: userInterfaceLayoutDirection */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMenu */



