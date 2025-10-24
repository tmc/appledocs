// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPopUpButton */


/* debug [class_header]: Header for NSPopUpButton */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PopUpButton */
// An interface definition for the [PopUpButton] class.
type IPopUpButton interface {
	IButton
	
/* debug [class_interface_properties]: Properties for PopUpButton */
	// properties:
	AltersStateOfSelectedItem() bool
	SetAltersStateOfSelectedItem(value bool)
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
	SelectedTag() int
	TitleOfSelectedItem() objc.IObject /* cross-framework: NSString */
	UsesItemFromMenu() bool
	SetUsesItemFromMenu(value bool)
	Image() IImage
	SetImage(value IImage)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PopUpButton */
	// methods:
	AddItemWithTitle(title objc.IObject /* cross-framework: NSString */)
	AddItemsWithTitles(itemTitles []string)
	IndexOfItem(item IMenuItem) int
	IndexOfItemWithRepresentedObject(obj objc.IObject) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.SEL) int
	IndexOfItemWithTitle(title objc.IObject /* cross-framework: NSString */) int
	InsertItemWithTitleAtIndex(title objc.IObject /* cross-framework: NSString */, index int)
	ItemAtIndex(index int) IMenuItem
	ItemWithTitle(title objc.IObject /* cross-framework: NSString */) IMenuItem
	ItemTitleAtIndex(index int) foundation.String
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItemWithTitle(title objc.IObject /* cross-framework: NSString */)
	SelectItem(item IMenuItem)
	SelectItemAtIndex(index int)
	SelectItemWithTag(tag int) bool
	SelectItemWithTitle(title objc.IObject /* cross-framework: NSString */)
	SetTitle(string_ objc.IObject /* cross-framework: NSString */)
	SynchronizeTitleAndSelectedItem()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PopUpButton */
// Alloc allocates a new instance without initialization.
func (pc _PopUpButtonClass) Alloc() PopUpButton {
	rv := objc.Send[PopUpButton](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PopUpButton */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PopUpButton */

// Returns an object initialized to the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/init(frame:pullsDown:)
func NewPopUpButtonWithFramePullsDown(buttonFrame Rect /* not a class type */, flag bool) PopUpButton {
	instance := getPopUpButtonClass().Alloc()
	rv := objc.Send[PopUpButton](instance.ID, objc.Sel("initWithFrame:pullsDown:"), buttonFrame, flag)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPopUpButtonWithFramePullsDown */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PopUpButton */

// Creates a standard pop-up button with a menu, target, and action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/popUpButtonWithMenu:target:action:
func (pc _PopUpButtonClass) PopUpButtonWithMenuTargetAction(menu IMenu, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("popUpButtonWithMenu:target:action:"), menu, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PopUpButtonWithMenuTargetAction) */


// Creates a standard pull-down button with an image and menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullDownButtonWithImage:menu:
func (pc _PopUpButtonClass) PullDownButtonWithImageMenu(image IImage, menu IMenu) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("pullDownButtonWithImage:menu:"), image, menu)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PullDownButtonWithImageMenu) */


// Creates a standard pull-down button with a title, image, and menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullDownButtonWithTitle:image:menu:
func (pc _PopUpButtonClass) PullDownButtonWithTitleImageMenu(title objc.IObject /* cross-framework: NSString */, image IImage, menu IMenu) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("pullDownButtonWithTitle:image:menu:"), title, image, menu)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PullDownButtonWithTitleImageMenu) */


// Creates a standard pull-down button with a title and menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullDownButtonWithTitle:menu:
func (pc _PopUpButtonClass) PullDownButtonWithTitleMenu(title objc.IObject /* cross-framework: NSString */, menu IMenu) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("pullDownButtonWithTitle:menu:"), title, menu)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PullDownButtonWithTitleMenu) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PopUpButton */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PopUpButton */

// Adds an item with the specified title to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/addItem(withTitle:)
func (p_ PopUpButton) AddItemWithTitle(title objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addItemWithTitle:"), title)
}/* debug [instance_methods/method]: AddItemWithTitle */


// Adds multiple items to the end of the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/addItems(withTitles:)
func (p_ PopUpButton) AddItemsWithTitles(itemTitles []string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addItemsWithTitles:"), itemTitles)
}/* debug [instance_methods/method]: AddItemsWithTitles */


// Returns the index of the specified menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/index(of:)
func (p_ PopUpButton) IndexOfItem(item IMenuItem) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItem:"), item)
	return rv
}/* debug [instance_methods/method]: IndexOfItem */


// Returns the index of the menu item that holds the specified represented object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withRepresentedObject:)
func (p_ PopUpButton) IndexOfItemWithRepresentedObject(obj objc.IObject) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithRepresentedObject:"), obj)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithRepresentedObject */


// Returns the index of the menu item with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withTag:)
func (p_ PopUpButton) IndexOfItemWithTag(tag int) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithTag */


// Returns the index of the menu item with the specified target and action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withTarget:andAction:)
func (p_ PopUpButton) IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.SEL) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithTargetAndAction */


// Returns the index of the item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withTitle:)
func (p_ PopUpButton) IndexOfItemWithTitle(title objc.IObject /* cross-framework: NSString */) int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfItemWithTitle:"), title)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithTitle */


// Inserts an item at the specified position in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/insertItem(withTitle:at:)
func (p_ PopUpButton) InsertItemWithTitleAtIndex(title objc.IObject /* cross-framework: NSString */, index int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("insertItemWithTitle:atIndex:"), title, index)
}/* debug [instance_methods/method]: InsertItemWithTitleAtIndex */


// Returns the menu item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/item(at:)
func (p_ PopUpButton) ItemAtIndex(index int) IMenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ItemAtIndex */


// Returns the menu item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/item(withTitle:)
func (p_ PopUpButton) ItemWithTitle(title objc.IObject /* cross-framework: NSString */) IMenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("itemWithTitle:"), title)
	return rv
}/* debug [instance_methods/method]: ItemWithTitle */


// Returns the title of the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemTitle(at:)
func (p_ PopUpButton) ItemTitleAtIndex(index int) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("itemTitleAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ItemTitleAtIndex */


// Removes all items in the receiver’s item menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/removeAllItems()
func (p_ PopUpButton) RemoveAllItems() {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAllItems"))
}/* debug [instance_methods/method]: RemoveAllItems */


// Removes the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/removeItem(at:)
func (p_ PopUpButton) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeItemAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveItemAtIndex */


// Removes the item with the specified title from the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/removeItem(withTitle:)
func (p_ PopUpButton) RemoveItemWithTitle(title objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeItemWithTitle:"), title)
}/* debug [instance_methods/method]: RemoveItemWithTitle */


// Selects the specified menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/select(_:)
func (p_ PopUpButton) SelectItem(item IMenuItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectItem:"), item)
}/* debug [instance_methods/method]: SelectItem */


// Selects the item in the menu at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectItem(at:)
func (p_ PopUpButton) SelectItemAtIndex(index int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectItemAtIndex:"), index)
}/* debug [instance_methods/method]: SelectItemAtIndex */


// Selects the menu item with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectItem(withTag:)
func (p_ PopUpButton) SelectItemWithTag(tag int) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("selectItemWithTag:"), tag)
	return rv
}/* debug [instance_methods/method]: SelectItemWithTag */


// Selects the item with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectItem(withTitle:)
func (p_ PopUpButton) SelectItemWithTitle(title objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectItemWithTitle:"), title)
}/* debug [instance_methods/method]: SelectItemWithTitle */


// Sets the string displayed in the receiver when the user isn’t pressing the mouse button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/setTitle(_:)
func (p_ PopUpButton) SetTitle(string_ objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), string_)
}/* debug [instance_methods/method]: SetTitle */


// Ensures that the item being displayed by the receiver agrees with the selected item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/synchronizeTitleAndSelectedItem()
func (p_ PopUpButton) SynchronizeTitleAndSelectedItem() {
	objc.Send[objc.ID](p_.ID, objc.Sel("synchronizeTitleAndSelectedItem"))
}/* debug [instance_methods/method]: SynchronizeTitleAndSelectedItem */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PopUpButton */

// When the value of this property is , the selected menu item’s is set to . When the value of this property is , the menu item’s is not changed. When this property changes, the of the currently selected item is updated appropriately. This property is ignored for pull-down buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/altersStateOfSelectedItem
func (p_ PopUpButton) AltersStateOfSelectedItem() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("altersStateOfSelectedItem"))
	return rv
}/* debug [instance_properties/getter]: altersStateOfSelectedItem */


// When the value of this property is , the selected menu item’s is set to . When the value of this property is , the menu item’s is not changed. When this property changes, the of the currently selected item is updated appropriately. This property is ignored for pull-down buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/altersStateOfSelectedItem
func (p_ PopUpButton) SetAltersStateOfSelectedItem(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAltersStateOfSelectedItem:"), value)
}/* debug [instance_properties/setter]: altersStateOfSelectedItem */


// A Boolean value indicating whether the button enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/autoenablesItems
func (p_ PopUpButton) AutoenablesItems() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("autoenablesItems"))
	return rv
}/* debug [instance_properties/getter]: autoenablesItems */


// A Boolean value indicating whether the button enables and disables its items every time a user event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/autoenablesItems
func (p_ PopUpButton) SetAutoenablesItems(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutoenablesItems:"), value)
}/* debug [instance_properties/setter]: autoenablesItems */


// The index of the item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfSelectedItem
func (p_ PopUpButton) IndexOfSelectedItem() int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}/* debug [instance_properties/getter]: indexOfSelectedItem */


// The array of menu item objects associated with the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemArray
func (p_ PopUpButton) ItemArray() []MenuItem {
	rv := objc.Send[[]MenuItem](p_.ID, objc.Sel("itemArray"))
	return rv
}/* debug [instance_properties/getter]: itemArray */


// An array of strings corresponding to the titles of the items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemTitles
func (p_ PopUpButton) ItemTitles() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("itemTitles"))
	return rv
}/* debug [instance_properties/getter]: itemTitles */


// The last item in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/lastItem
func (p_ PopUpButton) LastItem() IMenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("lastItem"))
	return rv
}/* debug [instance_properties/getter]: lastItem */


// The menu associated with the pop-up button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/menu
func (p_ PopUpButton) Menu() IMenu {
	rv := objc.Send[Menu](p_.ID, objc.Sel("menu"))
	return rv
}/* debug [instance_properties/getter]: menu */


// The menu associated with the pop-up button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/menu
func (p_ PopUpButton) SetMenu(value IMenu) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMenu:"), value)
}/* debug [instance_properties/setter]: menu */


// The number of items in the menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/numberOfItems
func (p_ PopUpButton) NumberOfItems() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfItems"))
	return rv
}/* debug [instance_properties/getter]: numberOfItems */


// The edge of the button on which to display the menu when screen space is constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/preferredEdge
func (p_ PopUpButton) PreferredEdge() RectEdge /* not a class type */ {
	rv := objc.Send[RectEdge](p_.ID, objc.Sel("preferredEdge"))
	return rv
}/* debug [instance_properties/getter]: preferredEdge */


// The edge of the button on which to display the menu when screen space is constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/preferredEdge
func (p_ PopUpButton) SetPreferredEdge(value RectEdge /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredEdge:"), value)
}/* debug [instance_properties/setter]: preferredEdge */


// A Boolean value indicating whether the button displays a pull-down or pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullsDown
func (p_ PopUpButton) PullsDown() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pullsDown"))
	return rv
}/* debug [instance_properties/getter]: pullsDown */


// A Boolean value indicating whether the button displays a pull-down or pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullsDown
func (p_ PopUpButton) SetPullsDown(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPullsDown:"), value)
}/* debug [instance_properties/setter]: pullsDown */


// The menu item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectedItem
func (p_ PopUpButton) SelectedItem() IMenuItem {
	rv := objc.Send[MenuItem](p_.ID, objc.Sel("selectedItem"))
	return rv
}/* debug [instance_properties/getter]: selectedItem */


// The tag of the menu item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectedTag
func (p_ PopUpButton) SelectedTag() int {
	rv := objc.Send[int](p_.ID, objc.Sel("selectedTag"))
	return rv
}/* debug [instance_properties/getter]: selectedTag */


// The title of the item that was last selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/titleOfSelectedItem
func (p_ PopUpButton) TitleOfSelectedItem() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("titleOfSelectedItem"))
	return rv
}/* debug [instance_properties/getter]: titleOfSelectedItem */


// When is , a pull-down button uses the title of the first menu item and hides the first menu item. A pop-up button uses the title of the currently selected menu. The default value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/usesItemFromMenu
func (p_ PopUpButton) UsesItemFromMenu() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesItemFromMenu"))
	return rv
}/* debug [instance_properties/getter]: usesItemFromMenu */


// When is , a pull-down button uses the title of the first menu item and hides the first menu item. A pop-up button uses the title of the currently selected menu. The default value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/usesItemFromMenu
func (p_ PopUpButton) SetUsesItemFromMenu(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesItemFromMenu:"), value)
}/* debug [instance_properties/setter]: usesItemFromMenu */


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/image
func (p_ PopUpButton) Image() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/image
func (p_ PopUpButton) SetImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPopUpButton */


