// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [Menu] class.
type IMenu interface {
	objectivec.IObject
	

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
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)
	Supermenu() IMenu
	SetSupermenu(value IMenu)
	Title() foundation.foundation.INSString
	SetTitle(value foundation.foundation.INSString)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)


	

	// methods:
	AddItem(newItem IMenuItem)
	AddItemWithTitleActionKeyEquivalent(string_ foundation.foundation.INSString, selector objc.SEL, charCode foundation.foundation.INSString) IMenuItem
	IndexOfItem(item IMenuItem) int
	IndexOfItemWithRepresentedObject(object objectivec.IObject) int
	IndexOfItemWithSubmenu(submenu IMenu) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int
	IndexOfItemWithTitle(title foundation.foundation.INSString) int
	InsertItemAtIndex(newItem IMenuItem, index int)
	InsertItemWithTitleActionKeyEquivalentAtIndex(string_ foundation.foundation.INSString, selector objc.SEL, charCode foundation.foundation.INSString, index int) IMenuItem
	ItemAtIndex(index int) IMenuItem
	ItemWithTag(tag int) IMenuItem
	ItemWithTitle(title foundation.foundation.INSString) IMenuItem
	ItemChanged(item IMenuItem)
	RemoveAllItems()
	RemoveItem(item IMenuItem)
	RemoveItemAtIndex(index int)


}





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










// Displays a contextual menu over a view for an event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (mc _MenuClass) PopUpContextMenuWithEventForView(menu IMenu, event IEvent, view IView) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:"), menu, event, view)
}












// Adds a menu item to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/addItem(_:)
func (m_ Menu) AddItem(newItem IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addItem:"), newItem)
}


// Creates a new menu item and adds it to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/addItem(withTitle:action:keyEquivalent:)
func (m_ Menu) AddItemWithTitleActionKeyEquivalent(string_ foundation.foundation.INSString, selector objc.SEL, charCode foundation.foundation.INSString) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("addItemWithTitle:action:keyEquivalent:"), string_, selector, charCode)
	return rv
}


// Returns the index identifying the location of a specified menu item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/index(of:)
func (m_ Menu) IndexOfItem(item IMenuItem) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItem:"), item)
	return rv
}


// Returns the index of the first menu item in the menu that has a given represented object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withRepresentedObject:)
func (m_ Menu) IndexOfItemWithRepresentedObject(object objectivec.IObject) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithRepresentedObject:"), object)
	return rv
}


// Returns the index of the menu item in the menu with the given submenu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withSubmenu:)
func (m_ Menu) IndexOfItemWithSubmenu(submenu IMenu) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithSubmenu:"), submenu)
	return rv
}


// Returns the index of the first menu item in the menu identified by a tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTag:)
func (m_ Menu) IndexOfItemWithTag(tag int) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}


// Returns the index of the first menu item in the menu that has a specified action and target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTarget:andAction:)
func (m_ Menu) IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}


// Returns the index of the first menu item in the menu that has a specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTitle:)
func (m_ Menu) IndexOfItemWithTitle(title foundation.foundation.INSString) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithTitle:"), title)
	return rv
}


// Inserts a menu item into the menu at a specific location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/insertItem(_:at:)
func (m_ Menu) InsertItemAtIndex(newItem IMenuItem, index int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertItem:atIndex:"), newItem, index)
}


// Creates and adds a menu item at a specified location in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/insertItem(withTitle:action:keyEquivalent:at:)
func (m_ Menu) InsertItemWithTitleActionKeyEquivalentAtIndex(string_ foundation.foundation.INSString, selector objc.SEL, charCode foundation.foundation.INSString, index int) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("insertItemWithTitle:action:keyEquivalent:atIndex:"), string_, selector, charCode, index)
	return rv
}


// Returns the menu item at a specific location of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/item(at:)
func (m_ Menu) ItemAtIndex(index int) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}


// Returns the first menu item in the menu with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/item(withTag:)
func (m_ Menu) ItemWithTag(tag int) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("itemWithTag:"), tag)
	return rv
}


// Returns the first menu item in the menu with a specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/item(withTitle:)
func (m_ Menu) ItemWithTitle(title foundation.foundation.INSString) IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("itemWithTitle:"), title)
	return rv
}


// Invoked when a menu item is modified visually (for example, its title changes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/itemChanged(_:)
func (m_ Menu) ItemChanged(item IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("itemChanged:"), item)
}


// Removes all the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/removeAllItems()
func (m_ Menu) RemoveAllItems() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllItems"))
}


// Removes a menu item from the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/removeItem(_:)
func (m_ Menu) RemoveItem(item IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeItem:"), item)
}


// Removes the menu item at a specified location in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/removeItem(at:)
func (m_ Menu) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeItemAtIndex:"), index)
}







// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/items
func (m_ Menu) ItemArray() []MenuItem {
	rv := objc.Send[[]MenuItem](m_.ID, objc.Sel("itemArray"))
	return rv
}


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
}


// The number of menu items in the menu, including separator items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/numberOfItems
func (m_ Menu) NumberOfItems() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfItems"))
	return rv
}


// Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/allowscontextmenuplugins
func (m_ Menu) AllowsContextMenuPlugIns() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsContextMenuPlugIns"))
	return rv
}


// Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/allowscontextmenuplugins
func (m_ Menu) SetAllowsContextMenuPlugIns(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsContextMenuPlugIns:"), value)
}


// Indicates whether the menu automatically enables and disables its menu items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/autoenablesitems
func (m_ Menu) AutoenablesItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("autoenablesItems"))
	return rv
}


// Indicates whether the menu automatically enables and disables its menu items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/autoenablesitems
func (m_ Menu) SetAutoenablesItems(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutoenablesItems:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/automaticallyinsertswritingtoolsitems
func (m_ Menu) AutomaticallyInsertsWritingToolsItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("automaticallyInsertsWritingToolsItems"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/automaticallyinsertswritingtoolsitems
func (m_ Menu) SetAutomaticallyInsertsWritingToolsItems(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutomaticallyInsertsWritingToolsItems:"), value)
}


// The font of the menu and its submenus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/font
func (m_ Menu) Font() IFont {
	rv := objc.Send[Font](m_.ID, objc.Sel("font"))
	return rv
}


// The font of the menu and its submenus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/font
func (m_ Menu) SetFont(value IFont) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFont:"), value)
}


// Indicates the currently highlighted item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/highlighteditem
func (m_ Menu) HighlightedItem() IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("highlightedItem"))
	return rv
}


// Indicates the currently highlighted item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/highlighteditem
func (m_ Menu) SetHighlightedItem(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHighlightedItem:"), value)
}


// Indicates whether the menu is offscreen or attached to another menu (or if it’s the main menu).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/istornoff
func (m_ Menu) IsTornOff() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isTornOff"))
	return rv
}


// Indicates whether the menu is offscreen or attached to another menu (or if it’s the main menu).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/istornoff
func (m_ Menu) SetIsTornOff(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsTornOff:"), value)
}


// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/items
func (m_ Menu) Items() IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("items"))
	return rv
}


// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/items
func (m_ Menu) SetItems(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setItems:"), value)
}


// The menu bar height for the main menu in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/menubarheight
func (m_ Menu) MenuBarHeight() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("menuBarHeight"))
	return rv
}


// The menu bar height for the main menu in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/menubarheight
func (m_ Menu) SetMenuBarHeight(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuBarHeight:"), value)
}


// Indicates whether messages are sent to the application’s windows each time the menu changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/menuchangedmessagesenabled
func (m_ Menu) MenuChangedMessagesEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("menuChangedMessagesEnabled"))
	return rv
}


// Indicates whether messages are sent to the application’s windows each time the menu changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/menuchangedmessagesenabled
func (m_ Menu) SetMenuChangedMessagesEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuChangedMessagesEnabled:"), value)
}


// The minimum width of the menu in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/minimumwidth
func (m_ Menu) MinimumWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("minimumWidth"))
	return rv
}


// The minimum width of the menu in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/minimumwidth
func (m_ Menu) SetMinimumWidth(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumWidth:"), value)
}


// The presentation style of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/presentationstyle-swift.property
func (m_ Menu) PresentationStyle() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("presentationStyle"))
	return rv
}


// The presentation style of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/presentationstyle-swift.property
func (m_ Menu) SetPresentationStyle(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentationStyle:"), value)
}


// The available properties for the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/propertiestoupdate
func (m_ Menu) PropertiesToUpdate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("propertiesToUpdate"))
	return rv
}


// The available properties for the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/propertiestoupdate
func (m_ Menu) SetPropertiesToUpdate(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPropertiesToUpdate:"), value)
}


// The menu items that are currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selecteditems
func (m_ Menu) SelectedItems() IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("selectedItems"))
	return rv
}


// The menu items that are currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selecteditems
func (m_ Menu) SetSelectedItems(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedItems:"), value)
}


// The selection mode of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selectionmode-swift.property
func (m_ Menu) SelectionMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("selectionMode"))
	return rv
}


// The selection mode of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selectionmode-swift.property
func (m_ Menu) SetSelectionMode(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectionMode:"), value)
}


// Indicates whether the menu displays the state column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/showsstatecolumn
func (m_ Menu) ShowsStateColumn() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsStateColumn"))
	return rv
}


// Indicates whether the menu displays the state column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/showsstatecolumn
func (m_ Menu) SetShowsStateColumn(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsStateColumn:"), value)
}


// The size of the menu in screen coordinates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/size
func (m_ Menu) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("size"))
	return rv
}


// The size of the menu in screen coordinates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/size
func (m_ Menu) SetSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}


// The parent menu that contains the menu as a submenu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/supermenu
func (m_ Menu) Supermenu() IMenu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("supermenu"))
	return rv
}


// The parent menu that contains the menu as a submenu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/supermenu
func (m_ Menu) SetSupermenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupermenu:"), value)
}


// The title of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/title
func (m_ Menu) Title() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}


// The title of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/title
func (m_ Menu) SetTitle(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}


// Configures the layout direction of menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/userinterfacelayoutdirection
func (m_ Menu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](m_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// Configures the layout direction of menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/userinterfacelayoutdirection
func (m_ Menu) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}








