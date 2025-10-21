// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// The class defines the visual appearance of pop-up buttons that display pop-up or pull-down menus. Pop-up menus present the user with a set of choices, much the way radio buttons do, but using much less space. Pull-down menus also provide a set of choices but present the information in a slightly different way, usually to provide a set of commands from which the user can choose.
//
// The class implements the user interface for the class. Changes made to a menu (such as adding, removing, or changing the items) are not apparent while the menu is being displayed or interacted with.
//
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


// The image displayed by the cell, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/image
func (p_ PopUpButtonCell) Image() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The image displayed by the cell, if any.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/image
func (p_ PopUpButtonCell) SetImage(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImage:"), value)
}

// A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/altersstateofselecteditem
func (p_ PopUpButtonCell) AltersStateOfSelectedItem() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("altersStateOfSelectedItem"))
	return rv
}


// SetAltersStateOfSelectedItem sets the value of the altersStateOfSelectedItem property.
// A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/altersstateofselecteditem
func (p_ PopUpButtonCell) SetAltersStateOfSelectedItem(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAltersStateOfSelectedItem:"), value)
}

// The position of the arrow displayed on the button.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/arrowposition
func (p_ PopUpButtonCell) ArrowPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("arrowPosition"))
	return rv
}


// SetArrowPosition sets the value of the arrowPosition property.
// The position of the arrow displayed on the button.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/arrowposition
func (p_ PopUpButtonCell) SetArrowPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setArrowPosition:"), value)
}

// A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/autoenablesitems
func (p_ PopUpButtonCell) AutoenablesItems() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("autoenablesItems"))
	return rv
}


// SetAutoenablesItems sets the value of the autoenablesItems property.
// A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/autoenablesitems
func (p_ PopUpButtonCell) SetAutoenablesItems(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutoenablesItems:"), value)
}

// The index of the item last selected by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/indexofselecteditem
func (p_ PopUpButtonCell) IndexOfSelectedItem() int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}


// SetIndexOfSelectedItem sets the value of the indexOfSelectedItem property.
// The index of the item last selected by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/indexofselecteditem
func (p_ PopUpButtonCell) SetIndexOfSelectedItem(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexOfSelectedItem:"), value)
}

// An array of
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/itemarray
func (p_ PopUpButtonCell) ItemArray() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("itemArray"))
	return rv
}


// SetItemArray sets the value of the itemArray property.
// An array of

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/itemarray
func (p_ PopUpButtonCell) SetItemArray(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setItemArray:"), value)
}

// An array of
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/itemtitles
func (p_ PopUpButtonCell) ItemTitles() string {
	rv := objc.Send[string](p_.ID, objc.Sel("itemTitles"))
	return rv
}


// SetItemTitles sets the value of the itemTitles property.
// An array of

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/itemtitles
func (p_ PopUpButtonCell) SetItemTitles(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setItemTitles:"), objc.String(value))
}

// The last item in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/lastitem
func (p_ PopUpButtonCell) LastItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("lastItem"))
	return rv
}


// SetLastItem sets the value of the lastItem property.
// The last item in the menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/lastitem
func (p_ PopUpButtonCell) SetLastItem(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLastItem:"), value)
}

// The pop-up button’s associated menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/menu
func (p_ PopUpButtonCell) Menu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("menu"))
	return rv
}


// SetMenu sets the value of the menu property.
// The pop-up button’s associated menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/menu
func (p_ PopUpButtonCell) SetMenu(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMenu:"), value)
}

// The number of items in the menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/numberofitems
func (p_ PopUpButtonCell) NumberOfItems() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfItems"))
	return rv
}


// SetNumberOfItems sets the value of the numberOfItems property.
// The number of items in the menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/numberofitems
func (p_ PopUpButtonCell) SetNumberOfItems(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfItems:"), value)
}

// The edge of the cell from which the menu should pop out when screen conditions are restrictive.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/preferrededge
func (p_ PopUpButtonCell) PreferredEdge() int {
	rv := objc.Send[int](p_.ID, objc.Sel("preferredEdge"))
	return rv
}


// SetPreferredEdge sets the value of the preferredEdge property.
// The edge of the cell from which the menu should pop out when screen conditions are restrictive.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/preferrededge
func (p_ PopUpButtonCell) SetPreferredEdge(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredEdge:"), value)
}

// A Boolean value that indicates the behavior of the button’s menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/pullsdown
func (p_ PopUpButtonCell) PullsDown() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pullsDown"))
	return rv
}


// SetPullsDown sets the value of the pullsDown property.
// A Boolean value that indicates the behavior of the button’s menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/pullsdown
func (p_ PopUpButtonCell) SetPullsDown(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPullsDown:"), value)
}

// The menu item last selected by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/selecteditem
func (p_ PopUpButtonCell) SelectedItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectedItem"))
	return rv
}


// SetSelectedItem sets the value of the selectedItem property.
// The menu item last selected by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/selecteditem
func (p_ PopUpButtonCell) SetSelectedItem(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectedItem:"), value)
}

// The title of the item last selected by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/titleofselecteditem
func (p_ PopUpButtonCell) TitleOfSelectedItem() string {
	rv := objc.Send[string](p_.ID, objc.Sel("titleOfSelectedItem"))
	return rv
}


// SetTitleOfSelectedItem sets the value of the titleOfSelectedItem property.
// The title of the item last selected by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/titleofselecteditem
func (p_ PopUpButtonCell) SetTitleOfSelectedItem(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitleOfSelectedItem:"), objc.String(value))
}

// A Boolean value that indicates if the control uses an item from the menu for its own title.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/usesitemfrommenu
func (p_ PopUpButtonCell) UsesItemFromMenu() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesItemFromMenu"))
	return rv
}


// SetUsesItemFromMenu sets the value of the usesItemFromMenu property.
// A Boolean value that indicates if the control uses an item from the menu for its own title.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/usesitemfrommenu
func (p_ PopUpButtonCell) SetUsesItemFromMenu(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesItemFromMenu:"), value)
}



