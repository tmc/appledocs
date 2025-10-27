// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PopUpButtonCell] class.
var (
	PopUpButtonCellClass     _PopUpButtonCellClass
	PopUpButtonCellClassOnce sync.Once
)

func getPopUpButtonCellClass() _PopUpButtonCellClass {
	PopUpButtonCellClassOnce.Do(func() {
		PopUpButtonCellClass = _PopUpButtonCellClass{objc.GetClass("NSPopUpButtonCell")}
	})
	return PopUpButtonCellClass
}

type _PopUpButtonCellClass struct {
	class objc.Class
}





// An interface definition for the [PopUpButtonCell] class.
type IPopUpButtonCell interface {
	IMenuItemCell
	

	// properties:
	AltersStateOfSelectedItem() bool
	SetAltersStateOfSelectedItem(value bool)
	ArrowPosition() PopUpArrowPosition
	SetArrowPosition(value PopUpArrowPosition)
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	IndexOfSelectedItem() int
	ItemArray() []MenuItem
	ItemTitles() []string
	LastItem() IMenuItem
	Menu() IMenu
	SetMenu(value IMenu)
	NumberOfItems() int
	PreferredEdge() RectEdge /* not a class type */
	SetPreferredEdge(value RectEdge /* not a class type */)
	PullsDown() bool
	SetPullsDown(value bool)
	SelectedItem() IMenuItem
	TitleOfSelectedItem() foundation.foundation.INSString
	UsesItemFromMenu() bool
	SetUsesItemFromMenu(value bool)
	Image() IImage
	SetImage(value IImage)


	

	// methods:
	AddItemWithTitle(title foundation.foundation.INSString)
	AddItemsWithTitles(itemTitles []string)
	AttachPopUpWithFrameInView(cellFrame corefoundation.CGRect, controlView IView)
	DismissPopUp()
	IndexOfItem(item IMenuItem) int
	IndexOfItemWithRepresentedObject(obj objectivec.IObject) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int
	IndexOfItemWithTitle(title foundation.foundation.INSString) int
	InsertItemWithTitleAtIndex(title foundation.foundation.INSString, index int)
	ItemAtIndex(index int) IMenuItem
	ItemWithTitle(title foundation.foundation.INSString) IMenuItem
	ItemTitleAtIndex(index int) foundation.String
	PerformClickWithFrameInView(frame corefoundation.CGRect, controlView IView)
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItemWithTitle(title foundation.foundation.INSString)
	SelectItem(item IMenuItem)
	SelectItemAtIndex(index int)
	SelectItemWithTag(tag int) bool
	SelectItemWithTitle(title foundation.foundation.INSString)
	SetTitle(string_ foundation.foundation.INSString)
	SynchronizeTitleAndSelectedItem()


}





// Alloc allocates a new instance without initialization.
func (pc _PopUpButtonCellClass) Alloc() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PopUpButtonCellClass) New() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PopUpButtonCell) Init() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PopUpButtonCell) Autorelease() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPopUpButtonCell creates a new PopUpButtonCell instance.
func NewPopUpButtonCell() PopUpButtonCell {
	return getPopUpButtonCellClass().New()
}





// The class defines the visual appearance of pop-up buttons that display pop-up or pull-down menus. Pop-up menus present the user with a set of choices, much the way radio buttons do, but using much less space. Pull-down menus also provide a set of choices but present the information in a slightly different way, usually to provide a set of commands from which the user can choose.
//
// The class implements the user interface for the class. Changes made to a menu (such as adding, removing, or changing the items) are not apparent while the menu is being displayed or interacted with.


// The class defines the visual appearance of pop-up buttons that display pop-up or pull-down menus. Pop-up menus present the user with a set of choices, much the way radio buttons do, but using much less space. Pull-down menus also provide a set of choices but present the information in a slightly different way, usually to provide a set of commands from which the user can choose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell
type PopUpButtonCell struct {
	MenuItemCell
}

// PopUpButtonCellFrom constructs a [PopUpButtonCell] from an unsafe.Pointer.
//
// The class defines the visual appearance of pop-up buttons that display pop-up or pull-down menus. Pop-up menus present the user with a set of choices, much the way radio buttons do, but using much less space. Pull-down menus also provide a set of choices but present the information in a slightly different way, usually to provide a set of commands from which the user can choose.
func PopUpButtonCellFrom(ptr unsafe.Pointer) PopUpButtonCell {
	return PopUpButtonCell{
		MenuItemCell: MenuItemCellFrom(ptr),
	}
}






// Returns an object initialized with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/init(textCell:pullsDown:)
func NewPopUpButtonCellTextCellPullsDown(stringValue foundation.foundation.INSString, pullDown bool) PopUpButtonCell {
	instance := getPopUpButtonCellClass().Alloc()
	rv := objc.Send[PopUpButtonCell](instance.ID, objc.Sel("initTextCell:pullsDown:"), stringValue, pullDown)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/init(coder:)
func NewPopUpButtonCellWithCoder(coder foundation.foundation.INSCoder) PopUpButtonCell {
	instance := getPopUpButtonCellClass().Alloc()
	rv := objc.Send[PopUpButtonCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

















// Adds an item with the specified title to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/addItem(withTitle:)
func (p_ PopUpButtonCell) AddItemWithTitle(title foundation.foundation.INSString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addItemWithTitle:"), title)
}


// Adds multiple items to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/addItems(withTitles:)
func (p_ PopUpButtonCell) AddItemsWithTitles(itemTitles []string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addItemsWithTitles:"), itemTitles)
}


// Sets up the receiver to display a menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/attachPopUp(withFrame:in:)
func (p_ PopUpButtonCell) AttachPopUpWithFrameInView(cellFrame corefoundation.CGRect, controlView IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("attachPopUpWithFrame:inView:"), cellFrame, controlView)
}


// Dismisses the pop-up button’s menu by ordering its window out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/dismissPopUp()
func (p_ PopUpButtonCell) DismissPopUp() {
	objc.Send[objc.ID](p_.ID, objc.Sel("dismissPopUp"))
}


// Returns the index of the specified menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/index(of:)
func (p_ PopUpButtonCell) IndexOfItem(item IMenuItem) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItem:"), item)
	return rv
}


// Returns the index of the menu item that holds the specified represented object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfItem(withRepresentedObject:)
func (p_ PopUpButtonCell) IndexOfItemWithRepresentedObject(obj objectivec.IObject) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithRepresentedObject:"), obj)
	return rv
}


// Returns the index of the menu item with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfItem(withTag:)
func (p_ PopUpButtonCell) IndexOfItemWithTag(tag int) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}


// Returns the index of the menu item with the specified target and action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfItem(withTarget:andAction:)
func (p_ PopUpButtonCell) IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}


// Returns the index of the item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfItem(withTitle:)
func (p_ PopUpButtonCell) IndexOfItemWithTitle(title foundation.foundation.INSString) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTitle:"), title)
	return rv
}


// Inserts an item at the specified position in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/insertItem(withTitle:at:)
func (p_ PopUpButtonCell) InsertItemWithTitleAtIndex(title foundation.foundation.INSString, index int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("insertItemWithTitle:atIndex:"), title, index)
}


// Returns the menu item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/item(at:)
func (p_ PopUpButtonCell) ItemAtIndex(index int) IMenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}


// Returns the menu item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/item(withTitle:)
func (p_ PopUpButtonCell) ItemWithTitle(title foundation.foundation.INSString) IMenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("itemWithTitle:"), title)
	return rv
}


// Returns the title of the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/itemTitle(at:)
func (p_ PopUpButtonCell) ItemTitleAtIndex(index int) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("itemTitleAtIndex:"), index)
	return rv
}


// Displays the receiver’s menu and track mouse events in it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/performClick(withFrame:in:)
func (p_ PopUpButtonCell) PerformClickWithFrameInView(frame corefoundation.CGRect, controlView IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performClickWithFrame:inView:"), frame, controlView)
}


// Removes all items in the receiver’s item menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/removeAllItems()
func (p_ PopUpButtonCell) RemoveAllItems() {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAllItems"))
}


// Removes the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/removeItem(at:)
func (p_ PopUpButtonCell) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeItemAtIndex:"), index)
}


// Removes the item with the specified title from the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/removeItem(withTitle:)
func (p_ PopUpButtonCell) RemoveItemWithTitle(title foundation.foundation.INSString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeItemWithTitle:"), title)
}


// Selects the specified menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/select(_:)
func (p_ PopUpButtonCell) SelectItem(item IMenuItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectItem:"), item)
}


// Selects the item in the menu at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/selectItem(at:)
func (p_ PopUpButtonCell) SelectItemAtIndex(index int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectItemAtIndex:"), index)
}


// Selects the menu item with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/selectItem(withTag:)
func (p_ PopUpButtonCell) SelectItemWithTag(tag int) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("selectItemWithTag:"), tag)
	return rv
}


// Selects the item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/selectItem(withTitle:)
func (p_ PopUpButtonCell) SelectItemWithTitle(title foundation.foundation.INSString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectItemWithTitle:"), title)
}


// Sets the string displayed in the receiver when the user isn’t pressing the mouse button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/setTitle(_:)
func (p_ PopUpButtonCell) SetTitle(string_ foundation.foundation.INSString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), string_)
}


// Synchronizes the pop-up button’s displayed item with the currently selected menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/synchronizeTitleAndSelectedItem()
func (p_ PopUpButtonCell) SynchronizeTitleAndSelectedItem() {
	objc.Send[objc.ID](p_.ID, objc.Sel("synchronizeTitleAndSelectedItem"))
}







// A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/altersStateOfSelectedItem
func (p_ PopUpButtonCell) AltersStateOfSelectedItem() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("altersStateOfSelectedItem"))
	return rv
}


// A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/altersStateOfSelectedItem
func (p_ PopUpButtonCell) SetAltersStateOfSelectedItem(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAltersStateOfSelectedItem:"), value)
}


// The position of the arrow displayed on the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/arrowPosition
func (p_ PopUpButtonCell) ArrowPosition() PopUpArrowPosition {
	rv := objc.Send[PopUpArrowPosition](p_.ID, objc.Sel("arrowPosition"))
	return rv
}


// The position of the arrow displayed on the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/arrowPosition
func (p_ PopUpButtonCell) SetArrowPosition(value PopUpArrowPosition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setArrowPosition:"), value)
}


// A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/autoenablesItems
func (p_ PopUpButtonCell) AutoenablesItems() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("autoenablesItems"))
	return rv
}


// A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/autoenablesItems
func (p_ PopUpButtonCell) SetAutoenablesItems(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutoenablesItems:"), value)
}


// The index of the item last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfSelectedItem
func (p_ PopUpButtonCell) IndexOfSelectedItem() int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}


// An array of objects that represent the items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/itemArray
func (p_ PopUpButtonCell) ItemArray() []MenuItem {
	rv := objc.Send[[]MenuItem](p_.ID, objc.Sel("itemArray"))
	return rv
}


// An array of objects containing the titles of every item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/itemTitles
func (p_ PopUpButtonCell) ItemTitles() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("itemTitles"))
	return rv
}


// The last item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/lastItem
func (p_ PopUpButtonCell) LastItem() IMenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("lastItem"))
	return rv
}


// The pop-up button’s associated menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/menu
func (p_ PopUpButtonCell) Menu() IMenu {
	rv := objc.Send[Menu](p_.ID, objc.Sel("menu"))
	return rv
}


// The pop-up button’s associated menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/menu
func (p_ PopUpButtonCell) SetMenu(value IMenu) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMenu:"), value)
}


// The number of items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/numberOfItems
func (p_ PopUpButtonCell) NumberOfItems() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The edge of the cell from which the menu should pop out when screen conditions are restrictive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/preferredEdge
func (p_ PopUpButtonCell) PreferredEdge() RectEdge /* not a class type */ {
	rv := objc.Send[RectEdge](p_.ID, objc.Sel("preferredEdge"))
	return rv
}


// The edge of the cell from which the menu should pop out when screen conditions are restrictive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/preferredEdge
func (p_ PopUpButtonCell) SetPreferredEdge(value RectEdge /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredEdge:"), value)
}


// A Boolean value that indicates the behavior of the button’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/pullsDown
func (p_ PopUpButtonCell) PullsDown() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pullsDown"))
	return rv
}


// A Boolean value that indicates the behavior of the button’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/pullsDown
func (p_ PopUpButtonCell) SetPullsDown(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPullsDown:"), value)
}


// The menu item last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/selectedItem
func (p_ PopUpButtonCell) SelectedItem() IMenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("selectedItem"))
	return rv
}


// The title of the item last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/titleOfSelectedItem
func (p_ PopUpButtonCell) TitleOfSelectedItem() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("titleOfSelectedItem"))
	return rv
}


// A Boolean value that indicates if the control uses an item from the menu for its own title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/usesItemFromMenu
func (p_ PopUpButtonCell) UsesItemFromMenu() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesItemFromMenu"))
	return rv
}


// A Boolean value that indicates if the control uses an item from the menu for its own title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/usesItemFromMenu
func (p_ PopUpButtonCell) SetUsesItemFromMenu(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesItemFromMenu:"), value)
}


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/image
func (p_ PopUpButtonCell) Image() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("image"))
	return rv
}


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/image
func (p_ PopUpButtonCell) SetImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImage:"), value)
}







