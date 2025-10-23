// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	AltersStateOfSelectedItem() bool /* primitive/slice/pointer. */
	SetAltersStateOfSelectedItem(value bool /* primitive/slice/pointer. */)
	AutoenablesItems() bool /* primitive/slice/pointer. */
	SetAutoenablesItems(value bool /* primitive/slice/pointer. */)
	IndexOfSelectedItem() int /* primitive/slice/pointer. */
	ItemArray() []MenuItem /* primitive/slice/pointer. */
	ItemTitles() []string /* primitive/slice/pointer. */
	LastItem() objc.IObject /* cross-framework: MenuItem */
	Menu() IMenu
	SetMenu(value IMenu)
	NumberOfItems() int /* primitive/slice/pointer. */
	PreferredEdge() int /* primitive/slice/pointer. */
	SetPreferredEdge(value int /* primitive/slice/pointer. */)
	PullsDown() bool /* primitive/slice/pointer. */
	SetPullsDown(value bool /* primitive/slice/pointer. */)
	SelectedItem() objc.IObject /* cross-framework: MenuItem */
	SelectedTag() int /* primitive/slice/pointer. */
	TitleOfSelectedItem() string /* primitive/slice/pointer. */
	UsesItemFromMenu() bool /* primitive/slice/pointer. */
	SetUsesItemFromMenu(value bool /* primitive/slice/pointer. */)
	Image() IImage
	SetImage(value IImage)
	// methods:
	AddItemWithTitle(title string /* primitive/slice/pointer. */)
	AddItemsWithTitles(itemTitles []string /* primitive/slice/pointer. */)
	IndexOfItem(item objc.IObject /* cross-framework MenuItem */) int /* primitive/slice/pointer. */
	IndexOfItemWithRepresentedObject(obj objectivec.IObject) int /* primitive/slice/pointer. */
	IndexOfItemWithTag(tag int /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */
	IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int /* primitive/slice/pointer. */
	IndexOfItemWithTitle(title string /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */
	InsertItemWithTitleAtIndex(title string /* primitive/slice/pointer. */, index int /* primitive/slice/pointer. */)
	ItemAtIndex(index int /* primitive/slice/pointer. */) objc.IObject /* cross-framework: MenuItem */
	ItemWithTitle(title string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: MenuItem */
	ItemTitleAtIndex(index int /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */
	RemoveAllItems()
	RemoveItemAtIndex(index int /* primitive/slice/pointer. */)
	RemoveItemWithTitle(title string /* primitive/slice/pointer. */)
	SelectItem(item objc.IObject /* cross-framework MenuItem */)
	SelectItemAtIndex(index int /* primitive/slice/pointer. */)
	SelectItemWithTag(tag int /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	SelectItemWithTitle(title string /* primitive/slice/pointer. */)
	SetTitle(string_ string /* primitive/slice/pointer. */)
	SynchronizeTitleAndSelectedItem()
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



// Returns an object initialized to the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/init(frame:pullsDown:)
func NewPopUpButtonWithFramePullsDown(buttonFrame coregraphics.CGRect, flag bool /* primitive/slice/pointer. */) PopUpButton {
	instance := getPopUpButtonClass().Alloc()
	rv := objc.Send[PopUpButton](instance.ID, objc.Sel("initWithFrame:pullsDown:"), buttonFrame, flag)
	rv.Autorelease()
	return rv
}



// Creates a standard pop-up button with a menu, target, and action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/popUpButtonWithMenu:target:action:
func (pc _PopUpButtonClass) PopUpButtonWithMenuTargetAction(menu IMenu, target objectivec.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("popUpButtonWithMenu:target:action:"), menu, target, action)
	return rv
}


// Creates a standard pull-down button with an image and menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullDownButtonWithImage:menu:
func (pc _PopUpButtonClass) PullDownButtonWithImageMenu(image IImage, menu IMenu) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pullDownButtonWithImage:menu:"), image, menu)
	return rv
}


// Creates a standard pull-down button with a title, image, and menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullDownButtonWithTitle:image:menu:
func (pc _PopUpButtonClass) PullDownButtonWithTitleImageMenu(title string /* primitive/slice/pointer. */, image IImage, menu IMenu) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pullDownButtonWithTitle:image:menu:"), objc.String(title), image, menu)
	return rv
}


// Creates a standard pull-down button with a title and menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullDownButtonWithTitle:menu:
func (pc _PopUpButtonClass) PullDownButtonWithTitleMenu(title string /* primitive/slice/pointer. */, menu IMenu) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pullDownButtonWithTitle:menu:"), objc.String(title), menu)
	return rv
}


// Adds an item with the specified title to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/addItem(withTitle:)
func (p_ PopUpButton) AddItemWithTitle(title string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addItemWithTitle:"), objc.String(title))
}


// Adds multiple items to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/addItems(withTitles:)
func (p_ PopUpButton) AddItemsWithTitles(itemTitles []string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addItemsWithTitles:"), itemTitles)
}


// Returns the index of the specified menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/index(of:)
func (p_ PopUpButton) IndexOfItem(item objc.IObject /* cross-framework MenuItem */) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItem:"), item)
	return rv
}


// Returns the index of the menu item that holds the specified represented object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withRepresentedObject:)
func (p_ PopUpButton) IndexOfItemWithRepresentedObject(obj objectivec.IObject) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithRepresentedObject:"), obj)
	return rv
}


// Returns the index of the menu item with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withTag:)
func (p_ PopUpButton) IndexOfItemWithTag(tag int /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}


// Returns the index of the menu item with the specified target and action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withTarget:andAction:)
func (p_ PopUpButton) IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}


// Returns the index of the item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withTitle:)
func (p_ PopUpButton) IndexOfItemWithTitle(title string /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTitle:"), objc.String(title))
	return rv
}


// Inserts an item at the specified position in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/insertItem(withTitle:at:)
func (p_ PopUpButton) InsertItemWithTitleAtIndex(title string /* primitive/slice/pointer. */, index int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("insertItemWithTitle:atIndex:"), objc.String(title), index)
}


// Returns the menu item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/item(at:)
func (p_ PopUpButton) ItemAtIndex(index int /* primitive/slice/pointer. */) objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}


// Returns the menu item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/item(withTitle:)
func (p_ PopUpButton) ItemWithTitle(title string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("itemWithTitle:"), objc.String(title))
	return rv
}


// Returns the title of the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemTitle(at:)
func (p_ PopUpButton) ItemTitleAtIndex(index int /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](p_.ID, objc.Sel("itemTitleAtIndex:"), index)
	return rv
}


// Removes all items in the receiver’s item menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/removeAllItems()
func (p_ PopUpButton) RemoveAllItems() {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAllItems"))
}


// Removes the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/removeItem(at:)
func (p_ PopUpButton) RemoveItemAtIndex(index int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeItemAtIndex:"), index)
}


// Removes the item with the specified title from the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/removeItem(withTitle:)
func (p_ PopUpButton) RemoveItemWithTitle(title string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeItemWithTitle:"), objc.String(title))
}


// Selects the specified menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/select(_:)
func (p_ PopUpButton) SelectItem(item objc.IObject /* cross-framework MenuItem */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectItem:"), item)
}


// Selects the item in the menu at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectItem(at:)
func (p_ PopUpButton) SelectItemAtIndex(index int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectItemAtIndex:"), index)
}


// Selects the menu item with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectItem(withTag:)
func (p_ PopUpButton) SelectItemWithTag(tag int /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("selectItemWithTag:"), tag)
	return rv
}


// Selects the item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectItem(withTitle:)
func (p_ PopUpButton) SelectItemWithTitle(title string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectItemWithTitle:"), objc.String(title))
}


// Sets the string displayed in the receiver when the user isn’t pressing the mouse button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/setTitle(_:)
func (p_ PopUpButton) SetTitle(string_ string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), objc.String(string_))
}


// Ensures that the item being displayed by the receiver agrees with the selected item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/synchronizeTitleAndSelectedItem()
func (p_ PopUpButton) SynchronizeTitleAndSelectedItem() {
	objc.Send[objc.ID](p_.ID, objc.Sel("synchronizeTitleAndSelectedItem"))
}


// When the value of this property is , the selected menu item’s is set to . When the value of this property is , the menu item’s is not changed. When this property changes, the of the currently selected item is updated appropriately. This property is ignored for pull-down buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/altersStateOfSelectedItem
func (p_ PopUpButton) AltersStateOfSelectedItem() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("altersStateOfSelectedItem"))
	return rv
}


// When the value of this property is , the selected menu item’s is set to . When the value of this property is , the menu item’s is not changed. When this property changes, the of the currently selected item is updated appropriately. This property is ignored for pull-down buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/altersStateOfSelectedItem
func (p_ PopUpButton) SetAltersStateOfSelectedItem(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAltersStateOfSelectedItem:"), value)
}


// A Boolean value indicating whether the button enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/autoenablesItems
func (p_ PopUpButton) AutoenablesItems() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("autoenablesItems"))
	return rv
}


// A Boolean value indicating whether the button enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/autoenablesItems
func (p_ PopUpButton) SetAutoenablesItems(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutoenablesItems:"), value)
}


// The index of the item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfSelectedItem
func (p_ PopUpButton) IndexOfSelectedItem() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}


// The array of menu item objects associated with the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemArray
func (p_ PopUpButton) ItemArray() []MenuItem /* primitive/slice/pointer. */ {
	rv := objc.Send[[]MenuItem](p_.ID, objc.Sel("itemArray"))
	return rv
}


// An array of strings corresponding to the titles of the items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemTitles
func (p_ PopUpButton) ItemTitles() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](p_.ID, objc.Sel("itemTitles"))
	return rv
}


// The last item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/lastItem
func (p_ PopUpButton) LastItem() objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("lastItem"))
	return rv
}


// The menu associated with the pop-up button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/menu
func (p_ PopUpButton) Menu() IMenu {
	rv := objc.Send[Menu](p_.ID, objc.Sel("menu"))
	return rv
}


// The menu associated with the pop-up button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/menu
func (p_ PopUpButton) SetMenu(value IMenu) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMenu:"), value)
}


// The number of items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/numberOfItems
func (p_ PopUpButton) NumberOfItems() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The edge of the button on which to display the menu when screen space is constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/preferredEdge
func (p_ PopUpButton) PreferredEdge() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("preferredEdge"))
	return rv
}


// The edge of the button on which to display the menu when screen space is constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/preferredEdge
func (p_ PopUpButton) SetPreferredEdge(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredEdge:"), value)
}


// A Boolean value indicating whether the button displays a pull-down or pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullsDown
func (p_ PopUpButton) PullsDown() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("pullsDown"))
	return rv
}


// A Boolean value indicating whether the button displays a pull-down or pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullsDown
func (p_ PopUpButton) SetPullsDown(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPullsDown:"), value)
}


// The menu item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectedItem
func (p_ PopUpButton) SelectedItem() objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("selectedItem"))
	return rv
}


// The tag of the menu item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectedTag
func (p_ PopUpButton) SelectedTag() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("selectedTag"))
	return rv
}


// The title of the item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/titleOfSelectedItem
func (p_ PopUpButton) TitleOfSelectedItem() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("titleOfSelectedItem"))
	return rv
}


// When is , a pull-down button uses the title of the first menu item and hides the first menu item. A pop-up button uses the title of the currently selected menu. The default value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/usesItemFromMenu
func (p_ PopUpButton) UsesItemFromMenu() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesItemFromMenu"))
	return rv
}


// When is , a pull-down button uses the title of the first menu item and hides the first menu item. A pop-up button uses the title of the currently selected menu. The default value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/usesItemFromMenu
func (p_ PopUpButton) SetUsesItemFromMenu(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesItemFromMenu:"), value)
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


