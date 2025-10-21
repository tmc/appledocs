// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	AddItem(newItem IMenuItem)
	AddItemWithTitleActionKeyEquivalent(string_ string, selector objc.SEL, charCode string) MenuItem
	AttachedMenu() Menu
	CancelTracking()
	CancelTrackingWithoutAnimation()
	ContextMenuRepresentation() objc.ID
	HelpRequested(eventPtr IEvent)
	IndexOfItem(item IMenuItem) int
	IndexOfItemWithRepresentedObject(object objectivec.IObject) int
	IndexOfItemWithSubmenu(submenu IMenu) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int
	IndexOfItemWithTitle(title string) int
	InsertItemAtIndex(newItem IMenuItem, index int)
	InsertItemWithTitleActionKeyEquivalentAtIndex(string_ string, selector objc.SEL, charCode string, index int) MenuItem
	IsAttached() bool
	ItemAtIndex(index int) MenuItem
	ItemWithTag(tag int) MenuItem
	ItemWithTitle(title string) MenuItem
	ItemChanged(item IMenuItem)
	LocationForSubmenu(submenu IMenu) coregraphics.CGPoint
	MenuRepresentation() objc.ID
	PerformActionForItemAtIndex(index int)
	PerformKeyEquivalent(event IEvent) bool
	PopUpMenuPositioningItemAtLocationInView(item IMenuItem, location coregraphics.CGPoint, view IView) bool
	RemoveAllItems()
	RemoveItem(item IMenuItem)
	RemoveItemAtIndex(index int)
	SetContextMenuRepresentation(menuRep objectivec.IObject)
	SetMenuRepresentation(menuRep objectivec.IObject)
	SetSubmenuForItem(menu IMenu, item IMenuItem)
	SetTearOffMenuRepresentation(menuRep objectivec.IObject)
	SizeToFit()
	SubmenuAction(sender objectivec.IObject)
	TearOffMenuRepresentation() objc.ID
	Update()
}

// An object that manages an app’s menus.
//
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


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/init(coder:)
func NewMenuWithCoder(coder foundation.ICoder) Menu {
	instance := getMenuClass().Alloc()
	rv := objc.Send[Menu](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Initializes and returns a menu having the specified title and with autoenabling of menu items turned on.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/init(title:)
func NewMenuWithTitle(title string) Menu {
	instance := getMenuClass().Alloc()
	rv := objc.Send[Menu](instance.ID, objc.Sel("initWithTitle:"), objc.String(title))
	rv.Autorelease()
	return rv
}


// Returns a Boolean value that indicates whether the menu bar is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/menuBarVisible()
func (mc _MenuClass) MenuBarVisible() bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("menuBarVisible"))
	return rv
}

// Returns the zone from which objects should be allocated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/menuZone()
func (mc _MenuClass) MenuZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("menuZone"))
	return rv
}

// Creates a palette style menu displaying user-selectable color tags.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/paletteMenuWithColors:titles:selectionHandler:
func (mc _MenuClass) PaletteMenuWithColorsTitlesSelectionHandler(colors []Color, itemTitles []string, onSelectionChange unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("paletteMenuWithColors:titles:selectionHandler:"), colors, itemTitles, onSelectionChange)
	return rv
}

// Creates a palette style menu displaying user-selectable color tags that tint using the specified array of colors.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/paletteMenuWithColors:titles:templateImage:selectionHandler:
func (mc _MenuClass) PaletteMenuWithColorsTitlesTemplateImageSelectionHandler(colors []Color, itemTitles []string, image IImage, onSelectionChange unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("paletteMenuWithColors:titles:templateImage:selectionHandler:"), colors, itemTitles, image, onSelectionChange)
	return rv
}

// Displays a contextual menu over a view for an event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (mc _MenuClass) PopUpContextMenuWithEventForView(menu IMenu, event IEvent, view IView) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:"), menu, event, view)
}

// Displays a contextual menu over a view for an event using a specified font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:with:)
func (mc _MenuClass) PopUpContextMenuWithEventForViewWithFont(menu IMenu, event IEvent, view IView, font IFont) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:withFont:"), menu, event, view, font)
}

// Sets whether the menu bar is visible and selectable by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/setMenuBarVisible(_:)
func (mc _MenuClass) SetMenuBarVisible(visible bool) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("setMenuBarVisible:"), visible)
}

// Sets the zone from which objects should be allocated
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/setMenuZone:
func (mc _MenuClass) SetMenuZone(zone unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("setMenuZone:"), zone)
}

// Adds a menu item to the end of the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/addItem(_:)
func (m_ Menu) AddItem(newItem IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addItem:"), newItem)
}

// Creates a new menu item and adds it to the end of the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/addItem(withTitle:action:keyEquivalent:)
func (m_ Menu) AddItemWithTitleActionKeyEquivalent(string_ string, selector objc.SEL, charCode string) MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("addItemWithTitle:action:keyEquivalent:"), objc.String(string_), selector, objc.String(charCode))
	return rv
}

// Returns the menu currently attached to the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/attachedMenu
func (m_ Menu) AttachedMenu() Menu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("attachedMenu"))
	return rv
}

// Dismisses the menu and ends all menu tracking.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/cancelTracking()
func (m_ Menu) CancelTracking() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelTracking"))
}

// Dismisses the menu and ends all menu tracking without displaying the associated animation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/cancelTrackingWithoutAnimation()
func (m_ Menu) CancelTrackingWithoutAnimation() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelTrackingWithoutAnimation"))
}

// Deprecated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/contextMenuRepresentation
func (m_ Menu) ContextMenuRepresentation() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("contextMenuRepresentation"))
	return rv
}

// Overridden by subclasses to implement specialized context-sensitive help behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/helpRequested(with:)
func (m_ Menu) HelpRequested(eventPtr IEvent) {
	objc.Send[objc.ID](m_.ID, objc.Sel("helpRequested:"), eventPtr)
}

// Returns the index identifying the location of a specified menu item in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/index(of:)
func (m_ Menu) IndexOfItem(item IMenuItem) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItem:"), item)
	return rv
}

// Returns the index of the first menu item in the menu that has a given represented object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withRepresentedObject:)
func (m_ Menu) IndexOfItemWithRepresentedObject(object objectivec.IObject) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithRepresentedObject:"), object)
	return rv
}

// Returns the index of the menu item in the menu with the given submenu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withSubmenu:)
func (m_ Menu) IndexOfItemWithSubmenu(submenu IMenu) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithSubmenu:"), submenu)
	return rv
}

// Returns the index of the first menu item in the menu identified by a tag.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTag:)
func (m_ Menu) IndexOfItemWithTag(tag int) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}

// Returns the index of the first menu item in the menu that has a specified action and target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTarget:andAction:)
func (m_ Menu) IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}

// Returns the index of the first menu item in the menu that has a specified title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTitle:)
func (m_ Menu) IndexOfItemWithTitle(title string) int {
	rv := objc.Send[int](m_.ID, objc.Sel("indexOfItemWithTitle:"), objc.String(title))
	return rv
}

// Inserts a menu item into the menu at a specific location.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/insertItem(_:at:)
func (m_ Menu) InsertItemAtIndex(newItem IMenuItem, index int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertItem:atIndex:"), newItem, index)
}

// Creates and adds a menu item at a specified location in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/insertItem(withTitle:action:keyEquivalent:at:)
func (m_ Menu) InsertItemWithTitleActionKeyEquivalentAtIndex(string_ string, selector objc.SEL, charCode string, index int) MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("insertItemWithTitle:action:keyEquivalent:atIndex:"), objc.String(string_), selector, objc.String(charCode), index)
	return rv
}

// Returns a Boolean value that indicates whether the menu is currently attached to another menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/isAttached
func (m_ Menu) IsAttached() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAttached"))
	return rv
}

// Returns the menu item at a specific location of the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/item(at:)
func (m_ Menu) ItemAtIndex(index int) MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}

// Returns the first menu item in the menu with the specified tag.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/item(withTag:)
func (m_ Menu) ItemWithTag(tag int) MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("itemWithTag:"), tag)
	return rv
}

// Returns the first menu item in the menu with a specified title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/item(withTitle:)
func (m_ Menu) ItemWithTitle(title string) MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("itemWithTitle:"), objc.String(title))
	return rv
}

// Invoked when a menu item is modified visually (for example, its title changes).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/itemChanged(_:)
func (m_ Menu) ItemChanged(item IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("itemChanged:"), item)
}

// Returns the location in screen coordinates where the given submenu is displayed when opened as a submenu of the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/locationForSubmenu:
func (m_ Menu) LocationForSubmenu(submenu IMenu) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](m_.ID, objc.Sel("locationForSubmenu:"), submenu)
	return rv
}

// Deprecated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/menuRepresentation
func (m_ Menu) MenuRepresentation() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("menuRepresentation"))
	return rv
}

// Causes the application to send the action message of a specified menu item to its target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/performActionForItem(at:)
func (m_ Menu) PerformActionForItemAtIndex(index int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("performActionForItemAtIndex:"), index)
}

// Performs the action for the menu item that corresponds to the given key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/performKeyEquivalent(with:)
func (m_ Menu) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("performKeyEquivalent:"), event)
	return rv
}

// Pops up the menu at the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUp(positioning:at:in:)
func (m_ Menu) PopUpMenuPositioningItemAtLocationInView(item IMenuItem, location coregraphics.CGPoint, view IView) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("popUpMenuPositioningItem:atLocation:inView:"), item, location, view)
	return rv
}

// Removes all the menu items in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/removeAllItems()
func (m_ Menu) RemoveAllItems() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllItems"))
}

// Removes a menu item from the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/removeItem(_:)
func (m_ Menu) RemoveItem(item IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeItem:"), item)
}

// Removes the menu item at a specified location in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/removeItem(at:)
func (m_ Menu) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeItemAtIndex:"), index)
}

// Deprecated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/setContextMenuRepresentation:
func (m_ Menu) SetContextMenuRepresentation(menuRep objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContextMenuRepresentation:"), menuRep)
}

// Deprecated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/setMenuRepresentation:
func (m_ Menu) SetMenuRepresentation(menuRep objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuRepresentation:"), menuRep)
}

// Assigns a menu to be a submenu of the menu controlled by a given menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/setSubmenu(_:for:)
func (m_ Menu) SetSubmenuForItem(menu IMenu, item IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubmenu:forItem:"), menu, item)
}

// Deprecated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/setTearOffMenuRepresentation:
func (m_ Menu) SetTearOffMenuRepresentation(menuRep objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTearOffMenuRepresentation:"), menuRep)
}

// Resizes the menu to exactly fit its items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/sizeToFit
func (m_ Menu) SizeToFit() {
	objc.Send[objc.ID](m_.ID, objc.Sel("sizeToFit"))
}

// The action method assigned to menu items that open submenus.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/submenuAction(_:)
func (m_ Menu) SubmenuAction(sender objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("submenuAction:"), sender)
}

// Deprecated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/tearOffMenuRepresentation
func (m_ Menu) TearOffMenuRepresentation() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("tearOffMenuRepresentation"))
	return rv
}

// Enables or disables the menu items of the menu based on the NSMenuValidation informal protocol and sizes the menu to fit its current menu items if necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/update()
func (m_ Menu) Update() {
	objc.Send[objc.ID](m_.ID, objc.Sel("update"))
}

// Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/allowsContextMenuPlugIns
func (m_ Menu) AllowsContextMenuPlugIns() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsContextMenuPlugIns"))
	return rv
}


// SetAllowsContextMenuPlugIns sets the value of the allowsContextMenuPlugIns property.
// Indicates whether the pop-up menu allows appending of contextual menu plug-in items.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/allowsContextMenuPlugIns
func (m_ Menu) SetAllowsContextMenuPlugIns(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsContextMenuPlugIns:"), value)
}

// Indicates whether the menu automatically enables and disables its menu items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/autoenablesItems
func (m_ Menu) AutoenablesItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("autoenablesItems"))
	return rv
}


// SetAutoenablesItems sets the value of the autoenablesItems property.
// Indicates whether the menu automatically enables and disables its menu items.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/autoenablesItems
func (m_ Menu) SetAutoenablesItems(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutoenablesItems:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/automaticallyInsertsWritingToolsItems
func (m_ Menu) AutomaticallyInsertsWritingToolsItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("automaticallyInsertsWritingToolsItems"))
	return rv
}


// SetAutomaticallyInsertsWritingToolsItems sets the value of the automaticallyInsertsWritingToolsItems property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/automaticallyInsertsWritingToolsItems
func (m_ Menu) SetAutomaticallyInsertsWritingToolsItems(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutomaticallyInsertsWritingToolsItems:"), value)
}

// The delegate of the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/delegate
func (m_ Menu) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate of the menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/delegate
func (m_ Menu) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}

// The font of the menu and its submenus.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/font
func (m_ Menu) Font() NSFont {
	rv := objc.Send[NSFont](m_.ID, objc.Sel("font"))
	return rv
}


// SetFont sets the value of the font property.
// The font of the menu and its submenus.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/font
func (m_ Menu) SetFont(value IFont) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFont:"), value)
}

// Indicates the currently highlighted item in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/highlightedItem
func (m_ Menu) HighlightedItem() NSMenuItem {
	rv := objc.Send[NSMenuItem](m_.ID, objc.Sel("highlightedItem"))
	return rv
}

// Indicates whether the menu is offscreen or attached to another menu (or if it’s the main menu).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/isTornOff
func (m_ Menu) TornOff() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("tornOff"))
	return rv
}

// An array containing the menu items in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/items
func (m_ Menu) ItemArray() []MenuItem {
	rv := objc.Send[[]MenuItem](m_.ID, objc.Sel("itemArray"))
	return rv
}


// SetItemArray sets the value of the itemArray property.
// An array containing the menu items in the menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/items
func (m_ Menu) SetItemArray(value []MenuItem) {
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

// The menu bar height for the main menu in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/menuBarHeight
func (m_ Menu) MenuBarHeight() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("menuBarHeight"))
	return rv
}

// Indicates whether messages are sent to the application’s windows each time the menu changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/menuChangedMessagesEnabled
func (m_ Menu) MenuChangedMessagesEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("menuChangedMessagesEnabled"))
	return rv
}


// SetMenuChangedMessagesEnabled sets the value of the menuChangedMessagesEnabled property.
// Indicates whether messages are sent to the application’s windows each time the menu changes.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/menuChangedMessagesEnabled
func (m_ Menu) SetMenuChangedMessagesEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuChangedMessagesEnabled:"), value)
}

// The minimum width of the menu in screen coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/minimumWidth
func (m_ Menu) MinimumWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("minimumWidth"))
	return rv
}


// SetMinimumWidth sets the value of the minimumWidth property.
// The minimum width of the menu in screen coordinates.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/minimumWidth
func (m_ Menu) SetMinimumWidth(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumWidth:"), value)
}

// The number of menu items in the menu, including separator items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/numberOfItems
func (m_ Menu) NumberOfItems() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfItems"))
	return rv
}

// The presentation style of the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/presentationStyle-swift.property
func (m_ Menu) PresentationStyle() MenuPresentationStyle {
	rv := objc.Send[MenuPresentationStyle](m_.ID, objc.Sel("presentationStyle"))
	return rv
}


// SetPresentationStyle sets the value of the presentationStyle property.
// The presentation style of the menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/presentationStyle-swift.property
func (m_ Menu) SetPresentationStyle(value MenuPresentationStyle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentationStyle:"), value)
}

// The available properties for the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/propertiesToUpdate
func (m_ Menu) PropertiesToUpdate() MenuProperties {
	rv := objc.Send[MenuProperties](m_.ID, objc.Sel("propertiesToUpdate"))
	return rv
}

// The menu items that are currently selected.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/selectedItems
func (m_ Menu) SelectedItems() []MenuItem {
	rv := objc.Send[[]MenuItem](m_.ID, objc.Sel("selectedItems"))
	return rv
}


// SetSelectedItems sets the value of the selectedItems property.
// The menu items that are currently selected.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/selectedItems
func (m_ Menu) SetSelectedItems(value []MenuItem) {
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
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedItems:"), nsArray)
}

// The selection mode of the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/selectionMode-swift.property
func (m_ Menu) SelectionMode() MenuSelectionMode {
	rv := objc.Send[MenuSelectionMode](m_.ID, objc.Sel("selectionMode"))
	return rv
}


// SetSelectionMode sets the value of the selectionMode property.
// The selection mode of the menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/selectionMode-swift.property
func (m_ Menu) SetSelectionMode(value MenuSelectionMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectionMode:"), value)
}

// Indicates whether the menu displays the state column.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/showsStateColumn
func (m_ Menu) ShowsStateColumn() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsStateColumn"))
	return rv
}


// SetShowsStateColumn sets the value of the showsStateColumn property.
// Indicates whether the menu displays the state column.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/showsStateColumn
func (m_ Menu) SetShowsStateColumn(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsStateColumn:"), value)
}

// The size of the menu in screen coordinates
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/size
func (m_ Menu) Size() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](m_.ID, objc.Sel("size"))
	return rv
}

// The parent menu that contains the menu as a submenu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/supermenu
func (m_ Menu) Supermenu() NSMenu {
	rv := objc.Send[NSMenu](m_.ID, objc.Sel("supermenu"))
	return rv
}


// SetSupermenu sets the value of the supermenu property.
// The parent menu that contains the menu as a submenu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/supermenu
func (m_ Menu) SetSupermenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupermenu:"), value)
}

// The title of the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/title
func (m_ Menu) Title() string {
	rv := objc.Send[string](m_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/title
func (m_ Menu) SetTitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// Configures the layout direction of menu items in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/userInterfaceLayoutDirection
func (m_ Menu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](m_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// SetUserInterfaceLayoutDirection sets the value of the userInterfaceLayoutDirection property.
// Configures the layout direction of menu items in the menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/userInterfaceLayoutDirection
func (m_ Menu) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}

// Indicates whether the menu is offscreen or attached to another menu (or if it’s the main menu).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/istornoff
func (m_ Menu) IsTornOff() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isTornOff"))
	return rv
}


// SetIsTornOff sets the value of the isTornOff property.
// Indicates whether the menu is offscreen or attached to another menu (or if it’s the main menu).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/istornoff
func (m_ Menu) SetIsTornOff(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsTornOff:"), value)
}

// An array containing the menu items in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/items
func (m_ Menu) Items() NSMenuItem {
	rv := objc.Send[NSMenuItem](m_.ID, objc.Sel("items"))
	return rv
}


// SetItems sets the value of the items property.
// An array containing the menu items in the menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/items
func (m_ Menu) SetItems(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setItems:"), value)
}


