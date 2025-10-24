// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSComboBox */


/* debug [class_header]: Header for NSComboBox */
// The class instance for the [ComboBox] class.
var (
	ComboBoxClass     _ComboBoxClass
	ComboBoxClassOnce sync.Once
)

func getComboBoxClass() _ComboBoxClass {
	ComboBoxClassOnce.Do(func() {
		ComboBoxClass = _ComboBoxClass{objc.GetClass("NSComboBox")}
	})
	return ComboBoxClass
}

type _ComboBoxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComboBox */
// An interface definition for the [ComboBox] class.
type IComboBox interface {
	ITextField
	
/* debug [class_interface_properties]: Properties for ComboBox */
	// properties:
	Completes() bool
	SetCompletes(value bool)
	DataSource() unsafe.Pointer
	SetDataSource(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	IndexOfSelectedItem() int
	IntercellSpacing() Size /* not a class type */
	SetIntercellSpacing(value Size /* not a class type */)
	ButtonBordered() bool
	SetButtonBordered(value bool)
	ItemHeight() float64
	SetItemHeight(value float64)
	NumberOfItems() int
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
	ObjectValueOfSelectedItem() objc.ID
	ObjectValues() objc.IObject /* cross-framework: NSArray */
	UsesDataSource() bool
	SetUsesDataSource(value bool)
	IsButtonBordered() bool
	SetIsButtonBordered(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComboBox */
	// methods:
	AddItemWithObjectValue(object objc.IObject)
	AddItemsWithObjectValues(objects objc.IObject /* cross-framework: NSArray */)
	DeselectItemAtIndex(index int)
	IndexOfItemWithObjectValue(object objc.IObject) int
	InsertItemWithObjectValueAtIndex(object objc.IObject, index int)
	ItemObjectValueAtIndex(index int) objc.ID
	NoteNumberOfItemsChanged()
	ReloadData()
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItemWithObjectValue(object objc.IObject)
	ScrollItemAtIndexToTop(index int)
	ScrollItemAtIndexToVisible(index int)
	SelectItemAtIndex(index int)
	SelectItemWithObjectValue(object objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComboBox */
// Alloc allocates a new instance without initialization.
func (cc _ComboBoxClass) Alloc() ComboBox {
	rv := objc.Send[ComboBox](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComboBoxClass) New() ComboBox {
	rv := objc.Send[ComboBox](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComboBox) Init() ComboBox {
	rv := objc.Send[ComboBox](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComboBox) Autorelease() ComboBox {
	rv := objc.Send[ComboBox](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComboBox creates a new ComboBox instance.
func NewComboBox() ComboBox {
	return getComboBoxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComboBox */
// A view that displays a list of values in a pop-up menu where the user selects a value or types in a custom value.
//
// A combo box combines the behavior of an object with an object. A combo box displays a list of values from a pop-up list, but also provides a means for users to type in custom values. For example, here’s a combo box in its initial state. Clicking in the text portion of the control allows the user to edit the current value. When the user clicks the down arrow at the right side of the text field, the pop-up list appears. The class uses to implement its user interface. Also see the protocol, which declares the methods that uses to access the contents of its data source object.


// A view that displays a list of values in a pop-up menu where the user selects a value or types in a custom value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox
type ComboBox struct {
	TextField
}

// ComboBoxFrom constructs a [ComboBox] from an unsafe.Pointer.
//
// A view that displays a list of values in a pop-up menu where the user selects a value or types in a custom value.
func ComboBoxFrom(ptr unsafe.Pointer) ComboBox {
	return ComboBox{
		TextField: TextFieldFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComboBox *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComboBox */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComboBox */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComboBox */

// Adds an object to the end of the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/addItem(withObjectValue:)
func (c_ ComboBox) AddItemWithObjectValue(object objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addItemWithObjectValue:"), object)
}/* debug [instance_methods/method]: AddItemWithObjectValue */


// Adds multiple objects to the end of the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/addItems(withObjectValues:)
func (c_ ComboBox) AddItemsWithObjectValues(objects objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addItemsWithObjectValues:"), objects)
}/* debug [instance_methods/method]: AddItemsWithObjectValues */


// Deselects the pop-up list item at the specified index if it’s selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/deselectItem(at:)
func (c_ ComboBox) DeselectItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deselectItemAtIndex:"), index)
}/* debug [instance_methods/method]: DeselectItemAtIndex */


// Searches the receiver’s internal item list for the specified object and returns the lowest matching index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/indexOfItem(withObjectValue:)
func (c_ ComboBox) IndexOfItemWithObjectValue(object objc.IObject) int {
	rv := objc.Send[int](c_.ID, objc.Sel("indexOfItemWithObjectValue:"), object)
	return rv
}/* debug [instance_methods/method]: IndexOfItemWithObjectValue */


// Inserts an object at the specified location in the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/insertItem(withObjectValue:at:)
func (c_ ComboBox) InsertItemWithObjectValueAtIndex(object objc.IObject, index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertItemWithObjectValue:atIndex:"), object, index)
}/* debug [instance_methods/method]: InsertItemWithObjectValueAtIndex */


// Returns the object located at the given index within the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/itemObjectValue(at:)
func (c_ ComboBox) ItemObjectValueAtIndex(index int) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("itemObjectValueAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ItemObjectValueAtIndex */


// Informs the receiver that the number of items in its data source has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/noteNumberOfItemsChanged()
func (c_ ComboBox) NoteNumberOfItemsChanged() {
	objc.Send[objc.ID](c_.ID, objc.Sel("noteNumberOfItemsChanged"))
}/* debug [instance_methods/method]: NoteNumberOfItemsChanged */


// Marks the receiver as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/reloadData()
func (c_ ComboBox) ReloadData() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadData"))
}/* debug [instance_methods/method]: ReloadData */


// Removes all items from the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/removeAllItems()
func (c_ ComboBox) RemoveAllItems() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllItems"))
}/* debug [instance_methods/method]: RemoveAllItems */


// Removes the object at the specified location from the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/removeItem(at:)
func (c_ ComboBox) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeItemAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveItemAtIndex */


// Removes all occurrences of the given object from the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/removeItem(withObjectValue:)
func (c_ ComboBox) RemoveItemWithObjectValue(object objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeItemWithObjectValue:"), object)
}/* debug [instance_methods/method]: RemoveItemWithObjectValue */


// Scrolls the receiver’s pop-up list vertically so that the item at the specified index is as close to the top as possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/scrollItemAtIndexToTop(_:)
func (c_ ComboBox) ScrollItemAtIndexToTop(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollItemAtIndexToTop:"), index)
}/* debug [instance_methods/method]: ScrollItemAtIndexToTop */


// Scrolls the receiver’s pop-up list vertically so that the item at the specified index is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/scrollItemAtIndexToVisible(_:)
func (c_ ComboBox) ScrollItemAtIndexToVisible(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollItemAtIndexToVisible:"), index)
}/* debug [instance_methods/method]: ScrollItemAtIndexToVisible */


// Selects the pop-up list row at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectItem(at:)
func (c_ ComboBox) SelectItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectItemAtIndex:"), index)
}/* debug [instance_methods/method]: SelectItemAtIndex */


// Selects the first pop-up list item that corresponds to the given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectItem(withObjectValue:)
func (c_ ComboBox) SelectItemWithObjectValue(object objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectItemWithObjectValue:"), object)
}/* debug [instance_methods/method]: SelectItemWithObjectValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComboBox */

// A Boolean value indicating whether the combo box tries to complete what the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/completes
func (c_ ComboBox) Completes() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("completes"))
	return rv
}/* debug [instance_properties/getter]: completes */


// A Boolean value indicating whether the combo box tries to complete what the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/completes
func (c_ ComboBox) SetCompletes(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletes:"), value)
}/* debug [instance_properties/setter]: completes */


// The object that provides the item data for the combo box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/dataSource
func (c_ ComboBox) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// The object that provides the item data for the combo box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/dataSource
func (c_ ComboBox) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/delegate
func (c_ ComboBox) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/delegate
func (c_ ComboBox) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value indicating whether the combo box has a vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/hasVerticalScroller
func (c_ ComboBox) HasVerticalScroller() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasVerticalScroller"))
	return rv
}/* debug [instance_properties/getter]: hasVerticalScroller */


// A Boolean value indicating whether the combo box has a vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/hasVerticalScroller
func (c_ ComboBox) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasVerticalScroller:"), value)
}/* debug [instance_properties/setter]: hasVerticalScroller */


// The index of the last item selected from the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/indexOfSelectedItem
func (c_ ComboBox) IndexOfSelectedItem() int {
	rv := objc.Send[int](c_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}/* debug [instance_properties/getter]: indexOfSelectedItem */


// The horizontal and vertical spacing between cells in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/intercellSpacing
func (c_ ComboBox) IntercellSpacing() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("intercellSpacing"))
	return rv
}/* debug [instance_properties/getter]: intercellSpacing */


// The horizontal and vertical spacing between cells in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/intercellSpacing
func (c_ ComboBox) SetIntercellSpacing(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntercellSpacing:"), value)
}/* debug [instance_properties/setter]: intercellSpacing */


// A Boolean value indicating whether the combo box displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/isButtonBordered
func (c_ ComboBox) ButtonBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("buttonBordered"))
	return rv
}/* debug [instance_properties/getter]: buttonBordered */


// A Boolean value indicating whether the combo box displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/isButtonBordered
func (c_ ComboBox) SetButtonBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setButtonBordered:"), value)
}/* debug [instance_properties/setter]: buttonBordered */


// The height of each item in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/itemHeight
func (c_ ComboBox) ItemHeight() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("itemHeight"))
	return rv
}/* debug [instance_properties/getter]: itemHeight */


// The height of each item in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/itemHeight
func (c_ ComboBox) SetItemHeight(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemHeight:"), value)
}/* debug [instance_properties/setter]: itemHeight */


// The total number of items in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/numberOfItems
func (c_ ComboBox) NumberOfItems() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfItems"))
	return rv
}/* debug [instance_properties/getter]: numberOfItems */


// The maximum number of visible items to display in the pop-up list at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/numberOfVisibleItems
func (c_ ComboBox) NumberOfVisibleItems() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfVisibleItems"))
	return rv
}/* debug [instance_properties/getter]: numberOfVisibleItems */


// The maximum number of visible items to display in the pop-up list at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/numberOfVisibleItems
func (c_ ComboBox) SetNumberOfVisibleItems(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfVisibleItems:"), value)
}/* debug [instance_properties/setter]: numberOfVisibleItems */


// The object corresponding to the last item selected from the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/objectValueOfSelectedItem
func (c_ ComboBox) ObjectValueOfSelectedItem() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectValueOfSelectedItem"))
	return rv
}/* debug [instance_properties/getter]: objectValueOfSelectedItem */


// An array of the items from the combo box’s internal list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/objectValues
func (c_ ComboBox) ObjectValues() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](c_.ID, objc.Sel("objectValues"))
	return rv
}/* debug [instance_properties/getter]: objectValues */


// A Boolean value indicating whether the combo box retrieves its items from a data source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/usesDataSource
func (c_ ComboBox) UsesDataSource() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesDataSource"))
	return rv
}/* debug [instance_properties/getter]: usesDataSource */


// A Boolean value indicating whether the combo box retrieves its items from a data source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/usesDataSource
func (c_ ComboBox) SetUsesDataSource(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesDataSource:"), value)
}/* debug [instance_properties/setter]: usesDataSource */


// A Boolean value indicating whether the combo box displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/isbuttonbordered
func (c_ ComboBox) IsButtonBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isButtonBordered"))
	return rv
}/* debug [instance_properties/getter]: isButtonBordered */


// A Boolean value indicating whether the combo box displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/isbuttonbordered
func (c_ ComboBox) SetIsButtonBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsButtonBordered:"), value)
}/* debug [instance_properties/setter]: isButtonBordered */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSComboBox */



