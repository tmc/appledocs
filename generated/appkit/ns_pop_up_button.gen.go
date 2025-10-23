// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PopUpButton] class.
var (
	PopUpButtonClass     _PopUpButtonClass
	PopUpButtonClassOnce sync.Once
)

func getPopUpButtonClass() _PopUpButtonClass {
	PopUpButtonClassOnce.Do(func() {
		PopUpButtonClass = _PopUpButtonClass{objc.GetClass("NSPopUpButton")}
	})
	return PopUpButtonClass
}

type _PopUpButtonClass struct {
	class objc.Class
}

// An interface definition for the [PopUpButton] class.
type IPopUpButton interface {
	IButton
	Image() IImage
	SetImage(value IImage)
	AltersStateOfSelectedItem() bool
	SetAltersStateOfSelectedItem(value bool)
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	IndexOfSelectedItem() int
	SetIndexOfSelectedItem(value int)
	ItemArray() MenuItem
	SetItemArray(value MenuItem)
	ItemTitles() string
	SetItemTitles(value string)
	LastItem() MenuItem
	SetLastItem(value MenuItem)
	Menu() IMenu
	SetMenu(value IMenu)
	NumberOfItems() int
	SetNumberOfItems(value int)
	PreferredEdge() int
	SetPreferredEdge(value int)
	PullsDown() bool
	SetPullsDown(value bool)
	SelectedItem() MenuItem
	SetSelectedItem(value MenuItem)
	TitleOfSelectedItem() string
	SetTitleOfSelectedItem(value string)
	UsesItemFromMenu() bool
	SetUsesItemFromMenu(value bool)
	ItemAtIndex(index int) MenuItem
	ItemTitleAtIndex(index int) foundation.String
}

// A control for selecting an item from a list.
//
// An object uses an object to implement its user interface. Note that while a menu is tracking user input, programmatic changes to the menu, such as adding, removing, or changing items on the menu, is not reflected.


// A control for selecting an item from a list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton
type PopUpButton struct {
	Button
}

// PopUpButtonFrom constructs a [PopUpButton] from an unsafe.Pointer.
//
// A control for selecting an item from a list.
func PopUpButtonFrom(ptr unsafe.Pointer) PopUpButton {
	return PopUpButton{
		Button: ButtonFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PopUpButtonClass) Alloc() PopUpButton {
	rv := objc.Send[PopUpButton](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PopUpButtonClass) New() PopUpButton {
	rv := objc.Send[PopUpButton](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PopUpButton) Init() PopUpButton {
	rv := objc.Send[PopUpButton](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PopUpButton) Autorelease() PopUpButton {
	rv := objc.Send[PopUpButton](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPopUpButton creates a new PopUpButton instance.
func NewPopUpButton() PopUpButton {
	return getPopUpButtonClass().New()
}



// Returns the menu item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/item(at:)
func (p_ PopUpButton) ItemAtIndex(index int) MenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}


// Returns the title of the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemTitle(at:)
func (p_ PopUpButton) ItemTitleAtIndex(index int) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("itemTitleAtIndex:"), index)
	return rv
}


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/image
func (p_ PopUpButton) Image() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("image"))
	return rv
}


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/image
func (p_ PopUpButton) SetImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImage:"), value)
}


// When the value of this property is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/altersstateofselecteditem
func (p_ PopUpButton) AltersStateOfSelectedItem() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("altersStateOfSelectedItem"))
	return rv
}


// When the value of this property is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/altersstateofselecteditem
func (p_ PopUpButton) SetAltersStateOfSelectedItem(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAltersStateOfSelectedItem:"), value)
}


// A Boolean value indicating whether the button enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/autoenablesitems
func (p_ PopUpButton) AutoenablesItems() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("autoenablesItems"))
	return rv
}


// A Boolean value indicating whether the button enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/autoenablesitems
func (p_ PopUpButton) SetAutoenablesItems(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutoenablesItems:"), value)
}


// The index of the item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/indexofselecteditem
func (p_ PopUpButton) IndexOfSelectedItem() int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}


// The index of the item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/indexofselecteditem
func (p_ PopUpButton) SetIndexOfSelectedItem(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexOfSelectedItem:"), value)
}


// The array of menu item objects associated with the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/itemarray
func (p_ PopUpButton) ItemArray() MenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("itemArray"))
	return rv
}


// The array of menu item objects associated with the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/itemarray
func (p_ PopUpButton) SetItemArray(value MenuItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setItemArray:"), value)
}


// An array of strings corresponding to the titles of the items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/itemtitles
func (p_ PopUpButton) ItemTitles() string {
	rv := objc.Send[string](p_.ID, objc.Sel("itemTitles"))
	return rv
}


// An array of strings corresponding to the titles of the items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/itemtitles
func (p_ PopUpButton) SetItemTitles(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setItemTitles:"), objc.String(value))
}


// The last item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/lastitem
func (p_ PopUpButton) LastItem() MenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("lastItem"))
	return rv
}


// The last item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/lastitem
func (p_ PopUpButton) SetLastItem(value MenuItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLastItem:"), value)
}


// The menu associated with the pop-up button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/menu
func (p_ PopUpButton) Menu() IMenu {
	rv := objc.Send[Menu](p_.ID, objc.Sel("menu"))
	return rv
}


// The menu associated with the pop-up button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/menu
func (p_ PopUpButton) SetMenu(value IMenu) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMenu:"), value)
}


// The number of items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/numberofitems
func (p_ PopUpButton) NumberOfItems() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The number of items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/numberofitems
func (p_ PopUpButton) SetNumberOfItems(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfItems:"), value)
}


// The edge of the button on which to display the menu when screen space is constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/preferrededge
func (p_ PopUpButton) PreferredEdge() int {
	rv := objc.Send[int](p_.ID, objc.Sel("preferredEdge"))
	return rv
}


// The edge of the button on which to display the menu when screen space is constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/preferrededge
func (p_ PopUpButton) SetPreferredEdge(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredEdge:"), value)
}


// A Boolean value indicating whether the button displays a pull-down or pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/pullsdown
func (p_ PopUpButton) PullsDown() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pullsDown"))
	return rv
}


// A Boolean value indicating whether the button displays a pull-down or pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/pullsdown
func (p_ PopUpButton) SetPullsDown(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPullsDown:"), value)
}


// The menu item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/selecteditem
func (p_ PopUpButton) SelectedItem() MenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("selectedItem"))
	return rv
}


// The menu item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/selecteditem
func (p_ PopUpButton) SetSelectedItem(value MenuItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectedItem:"), value)
}


// The title of the item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/titleofselecteditem
func (p_ PopUpButton) TitleOfSelectedItem() string {
	rv := objc.Send[string](p_.ID, objc.Sel("titleOfSelectedItem"))
	return rv
}


// The title of the item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/titleofselecteditem
func (p_ PopUpButton) SetTitleOfSelectedItem(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitleOfSelectedItem:"), objc.String(value))
}


// When
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/usesitemfrommenu
func (p_ PopUpButton) UsesItemFromMenu() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesItemFromMenu"))
	return rv
}


// When
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/usesitemfrommenu
func (p_ PopUpButton) SetUsesItemFromMenu(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesItemFromMenu:"), value)
}



