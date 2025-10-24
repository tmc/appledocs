// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
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
	AutomaticallyInsertsWritingToolsItems() bool
	SetAutomaticallyInsertsWritingToolsItems(value bool)
	ItemArray() []objc.IObject /* cross-framework: MenuItem */
	SetItemArray(value []objc.IObject /* cross-framework: MenuItem */)
	MenuChangedMessagesEnabled() bool
	SetMenuChangedMessagesEnabled(value bool)
	AllowsContextMenuPlugIns() bool
	SetAllowsContextMenuPlugIns(value bool)
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	Delegate() MenuDelegate /* not a class type */
	SetDelegate(value MenuDelegate /* not a class type */)
	Font() IFont
	SetFont(value IFont)
	HighlightedItem() objc.IObject /* cross-framework: MenuItem */
	SetHighlightedItem(value objc.IObject /* cross-framework: MenuItem */)
	IsTornOff() bool
	SetIsTornOff(value bool)
	Items() objc.IObject /* cross-framework: MenuItem */
	SetItems(value objc.IObject /* cross-framework: MenuItem */)
	MenuBarHeight() float64
	SetMenuBarHeight(value float64)
	MinimumWidth() float64
	SetMinimumWidth(value float64)
	NumberOfItems() int
	SetNumberOfItems(value int)
	PresentationStyle() unsafe.Pointer
	SetPresentationStyle(value unsafe.Pointer)
	PropertiesToUpdate() unsafe.Pointer
	SetPropertiesToUpdate(value unsafe.Pointer)
	SelectedItems() objc.IObject /* cross-framework: MenuItem */
	SetSelectedItems(value objc.IObject /* cross-framework: MenuItem */)
	SelectionMode() unsafe.Pointer
	SetSelectionMode(value unsafe.Pointer)
	ShowsStateColumn() bool
	SetShowsStateColumn(value bool)
	Size() objc.IObject /* cross-framework: Size */
	SetSize(value objc.IObject /* cross-framework: Size */)
	Supermenu() IMenu
	SetSupermenu(value IMenu)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	// methods:
	PopUpMenuPositioningItemAtLocationInView(item objc.IObject /* cross-framework: MenuItem */, location objc.IObject /* cross-framework: Point */, view IView) bool
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

// Alloc allocates a new instance without initialization.
func (mc _MenuClass) Alloc() Menu {
	rv := objc.Send[Menu](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Displays a contextual menu over a view for an event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (mc _MenuClass) PopUpContextMenuWithEventForView(menu IMenu, event IEvent, view IView) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:"), menu, event, view)
}


// Displays a contextual menu over a view for an event using a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:with:)
func (mc _MenuClass) PopUpContextMenuWithEventForViewWithFont(menu IMenu, event IEvent, view IView, font IFont) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:withFont:"), menu, event, view, font)
}


// Pops up the menu at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUp(positioning:at:in:)
func (m_ Menu) PopUpMenuPositioningItemAtLocationInView(item objc.IObject /* cross-framework: MenuItem */, location objc.IObject /* cross-framework: Point */, view IView) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("popUpMenuPositioningItem:atLocation:inView:"), item, location, view)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/automaticallyInsertsWritingToolsItems
func (m_ Menu) AutomaticallyInsertsWritingToolsItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("automaticallyInsertsWritingToolsItems"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/automaticallyInsertsWritingToolsItems
func (m_ Menu) SetAutomaticallyInsertsWritingToolsItems(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutomaticallyInsertsWritingToolsItems:"), value)
}


// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/items
func (m_ Menu) ItemArray() []objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[[]MenuItem](m_.ID, objc.Sel("itemArray"))
	return rv
}


// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/items
func (m_ Menu) SetItemArray(value []objc.IObject /* cross-framework: MenuItem */) {
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
	objc.Send[objc.ID](m_.ID, objc.Sel("setItemArray:"), nsArray)
}


// Indicates whether messages are sent to the application’s windows each time the menu changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/menuChangedMessagesEnabled
func (m_ Menu) MenuChangedMessagesEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("menuChangedMessagesEnabled"))
	return rv
}


// Indicates whether messages are sent to the application’s windows each time the menu changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/menuChangedMessagesEnabled
func (m_ Menu) SetMenuChangedMessagesEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuChangedMessagesEnabled:"), value)
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


// The delegate of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/delegate
func (m_ Menu) Delegate() MenuDelegate /* not a class type */ {
	rv := objc.Send[MenuDelegate](m_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/delegate
func (m_ Menu) SetDelegate(value MenuDelegate /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
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
func (m_ Menu) HighlightedItem() objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("highlightedItem"))
	return rv
}


// Indicates the currently highlighted item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/highlighteditem
func (m_ Menu) SetHighlightedItem(value objc.IObject /* cross-framework: MenuItem */) {
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
func (m_ Menu) Items() objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("items"))
	return rv
}


// An array containing the menu items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/items
func (m_ Menu) SetItems(value objc.IObject /* cross-framework: MenuItem */) {
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


// The number of menu items in the menu, including separator items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/numberofitems
func (m_ Menu) NumberOfItems() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The number of menu items in the menu, including separator items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/numberofitems
func (m_ Menu) SetNumberOfItems(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfItems:"), value)
}


// The presentation style of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/presentationstyle-swift.property
func (m_ Menu) PresentationStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("presentationStyle"))
	return rv
}


// The presentation style of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/presentationstyle-swift.property
func (m_ Menu) SetPresentationStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentationStyle:"), value)
}


// The available properties for the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/propertiestoupdate
func (m_ Menu) PropertiesToUpdate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("propertiesToUpdate"))
	return rv
}


// The available properties for the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/propertiestoupdate
func (m_ Menu) SetPropertiesToUpdate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPropertiesToUpdate:"), value)
}


// The menu items that are currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selecteditems
func (m_ Menu) SelectedItems() objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("selectedItems"))
	return rv
}


// The menu items that are currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selecteditems
func (m_ Menu) SetSelectedItems(value objc.IObject /* cross-framework: MenuItem */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedItems:"), value)
}


// The selection mode of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selectionmode-swift.property
func (m_ Menu) SelectionMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("selectionMode"))
	return rv
}


// The selection mode of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/selectionmode-swift.property
func (m_ Menu) SetSelectionMode(value unsafe.Pointer) {
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
func (m_ Menu) Size() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](m_.ID, objc.Sel("size"))
	return rv
}


// The size of the menu in screen coordinates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/size
func (m_ Menu) SetSize(value objc.IObject /* cross-framework: Size */) {
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
func (m_ Menu) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}


// The title of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/title
func (m_ Menu) SetTitle(value objc.IObject /* cross-framework: NSString */) {
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



