// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	ItemArray() []MenuItem /* primitive/slice/pointer. */
	TitleOfSelectedItem() objc.IObject /* cross-framework: NSString */
	Image() IImage
	SetImage(value IImage)
	AltersStateOfSelectedItem() bool /* primitive/slice/pointer. */
	SetAltersStateOfSelectedItem(value bool /* primitive/slice/pointer. */)
	ArrowPosition() unsafe.Pointer
	SetArrowPosition(value unsafe.Pointer)
	AutoenablesItems() bool /* primitive/slice/pointer. */
	SetAutoenablesItems(value bool /* primitive/slice/pointer. */)
	IndexOfSelectedItem() int /* primitive/slice/pointer. */
	SetIndexOfSelectedItem(value int /* primitive/slice/pointer. */)
	ItemTitles() objc.IObject /* cross-framework: NSString */
	SetItemTitles(value objc.IObject /* cross-framework: NSString */)
	LastItem() objc.IObject /* cross-framework: MenuItem */
	SetLastItem(value objc.IObject /* cross-framework: MenuItem */)
	Menu() IMenu
	SetMenu(value IMenu)
	NumberOfItems() int /* primitive/slice/pointer. */
	SetNumberOfItems(value int /* primitive/slice/pointer. */)
	PreferredEdge() RectEdge /* not a class type */
	SetPreferredEdge(value RectEdge /* not a class type */)
	PullsDown() bool /* primitive/slice/pointer. */
	SetPullsDown(value bool /* primitive/slice/pointer. */)
	SelectedItem() objc.IObject /* cross-framework: MenuItem */
	SetSelectedItem(value objc.IObject /* cross-framework: MenuItem */)
	UsesItemFromMenu() bool /* primitive/slice/pointer. */
	SetUsesItemFromMenu(value bool /* primitive/slice/pointer. */)
	// methods:
	IndexOfItemWithTitle(title objc.IObject /* cross-framework NSString */) int /* primitive/slice/pointer. */
	InsertItemWithTitleAtIndex(title objc.IObject /* cross-framework NSString */, index int /* primitive/slice/pointer. */)
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

// Alloc allocates a new instance without initialization.
func (pc _PopUpButtonCellClass) Alloc() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns the index of the item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfItem(withTitle:)
func (p_ PopUpButtonCell) IndexOfItemWithTitle(title objc.IObject /* cross-framework NSString */) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTitle:"), title)
	return rv
}


// Inserts an item at the specified position in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/insertItem(withTitle:at:)
func (p_ PopUpButtonCell) InsertItemWithTitleAtIndex(title objc.IObject /* cross-framework NSString */, index int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("insertItemWithTitle:atIndex:"), title, index)
}


// An array of objects that represent the items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/itemArray
func (p_ PopUpButtonCell) ItemArray() []MenuItem /* primitive/slice/pointer. */ {
	rv := objc.Send[[]MenuItem](p_.ID, objc.Sel("itemArray"))
	return rv
}


// The title of the item last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/titleOfSelectedItem
func (p_ PopUpButtonCell) TitleOfSelectedItem() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("titleOfSelectedItem"))
	return rv
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


// A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/altersstateofselecteditem
func (p_ PopUpButtonCell) AltersStateOfSelectedItem() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("altersStateOfSelectedItem"))
	return rv
}


// A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/altersstateofselecteditem
func (p_ PopUpButtonCell) SetAltersStateOfSelectedItem(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAltersStateOfSelectedItem:"), value)
}


// The position of the arrow displayed on the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/arrowposition
func (p_ PopUpButtonCell) ArrowPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("arrowPosition"))
	return rv
}


// The position of the arrow displayed on the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/arrowposition
func (p_ PopUpButtonCell) SetArrowPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setArrowPosition:"), value)
}


// A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/autoenablesitems
func (p_ PopUpButtonCell) AutoenablesItems() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("autoenablesItems"))
	return rv
}


// A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/autoenablesitems
func (p_ PopUpButtonCell) SetAutoenablesItems(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutoenablesItems:"), value)
}


// The index of the item last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/indexofselecteditem
func (p_ PopUpButtonCell) IndexOfSelectedItem() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}


// The index of the item last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/indexofselecteditem
func (p_ PopUpButtonCell) SetIndexOfSelectedItem(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexOfSelectedItem:"), value)
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/itemtitles
func (p_ PopUpButtonCell) ItemTitles() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("itemTitles"))
	return rv
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/itemtitles
func (p_ PopUpButtonCell) SetItemTitles(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setItemTitles:"), value)
}


// The last item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/lastitem
func (p_ PopUpButtonCell) LastItem() objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("lastItem"))
	return rv
}


// The last item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/lastitem
func (p_ PopUpButtonCell) SetLastItem(value objc.IObject /* cross-framework: MenuItem */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLastItem:"), value)
}


// The pop-up button’s associated menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/menu
func (p_ PopUpButtonCell) Menu() IMenu {
	rv := objc.Send[Menu](p_.ID, objc.Sel("menu"))
	return rv
}


// The pop-up button’s associated menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/menu
func (p_ PopUpButtonCell) SetMenu(value IMenu) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMenu:"), value)
}


// The number of items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/numberofitems
func (p_ PopUpButtonCell) NumberOfItems() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The number of items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/numberofitems
func (p_ PopUpButtonCell) SetNumberOfItems(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfItems:"), value)
}


// The edge of the cell from which the menu should pop out when screen conditions are restrictive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/preferrededge
func (p_ PopUpButtonCell) PreferredEdge() RectEdge /* not a class type */ {
	rv := objc.Send[RectEdge](p_.ID, objc.Sel("preferredEdge"))
	return rv
}


// The edge of the cell from which the menu should pop out when screen conditions are restrictive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/preferrededge
func (p_ PopUpButtonCell) SetPreferredEdge(value RectEdge /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredEdge:"), value)
}


// A Boolean value that indicates the behavior of the button’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/pullsdown
func (p_ PopUpButtonCell) PullsDown() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("pullsDown"))
	return rv
}


// A Boolean value that indicates the behavior of the button’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/pullsdown
func (p_ PopUpButtonCell) SetPullsDown(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPullsDown:"), value)
}


// The menu item last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/selecteditem
func (p_ PopUpButtonCell) SelectedItem() objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("selectedItem"))
	return rv
}


// The menu item last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/selecteditem
func (p_ PopUpButtonCell) SetSelectedItem(value objc.IObject /* cross-framework: MenuItem */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectedItem:"), value)
}


// A Boolean value that indicates if the control uses an item from the menu for its own title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/usesitemfrommenu
func (p_ PopUpButtonCell) UsesItemFromMenu() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesItemFromMenu"))
	return rv
}


// A Boolean value that indicates if the control uses an item from the menu for its own title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/usesitemfrommenu
func (p_ PopUpButtonCell) SetUsesItemFromMenu(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesItemFromMenu:"), value)
}



