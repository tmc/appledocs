// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ComboBox] class.
type IComboBox interface {
	ITextField
	

	// properties:
	Completes() bool
	SetCompletes(value bool)
	DataSource() unsafe.Pointer
	SetDataSource(value unsafe.Pointer)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	IndexOfSelectedItem() int
	IntercellSpacing() corefoundation.CGSize
	SetIntercellSpacing(value corefoundation.CGSize)
	ButtonBordered() bool
	SetButtonBordered(value bool)
	ItemHeight() float64
	SetItemHeight(value float64)
	NumberOfItems() int
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
	ObjectValueOfSelectedItem() objc.ID
	ObjectValues() foundation.foundation.INSArray
	UsesDataSource() bool
	SetUsesDataSource(value bool)
	IsButtonBordered() bool
	SetIsButtonBordered(value bool)


	

	// methods:
	AddItemWithObjectValue(object objectivec.IObject)
	AddItemsWithObjectValues(objects foundation.foundation.INSArray)
	DeselectItemAtIndex(index int)
	IndexOfItemWithObjectValue(object objectivec.IObject) int
	InsertItemWithObjectValueAtIndex(object objectivec.IObject, index int)
	ItemObjectValueAtIndex(index int) objc.ID
	NoteNumberOfItemsChanged()
	ReloadData()
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItemWithObjectValue(object objectivec.IObject)
	ScrollItemAtIndexToTop(index int)
	ScrollItemAtIndexToVisible(index int)
	SelectItemAtIndex(index int)
	SelectItemWithObjectValue(object objectivec.IObject)


}





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




















// Adds an object to the end of the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/addItem(withObjectValue:)
func (c_ ComboBox) AddItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addItemWithObjectValue:"), object)
}


// Adds multiple objects to the end of the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/addItems(withObjectValues:)
func (c_ ComboBox) AddItemsWithObjectValues(objects foundation.foundation.INSArray) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addItemsWithObjectValues:"), objects)
}


// Deselects the pop-up list item at the specified index if it’s selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/deselectItem(at:)
func (c_ ComboBox) DeselectItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deselectItemAtIndex:"), index)
}


// Searches the receiver’s internal item list for the specified object and returns the lowest matching index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/indexOfItem(withObjectValue:)
func (c_ ComboBox) IndexOfItemWithObjectValue(object objectivec.IObject) int {
	rv := objc.Send[int](c_.ID, objc.Sel("indexOfItemWithObjectValue:"), object)
	return rv
}


// Inserts an object at the specified location in the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/insertItem(withObjectValue:at:)
func (c_ ComboBox) InsertItemWithObjectValueAtIndex(object objectivec.IObject, index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertItemWithObjectValue:atIndex:"), object, index)
}


// Returns the object located at the given index within the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/itemObjectValue(at:)
func (c_ ComboBox) ItemObjectValueAtIndex(index int) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("itemObjectValueAtIndex:"), index)
	return rv
}


// Informs the receiver that the number of items in its data source has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/noteNumberOfItemsChanged()
func (c_ ComboBox) NoteNumberOfItemsChanged() {
	objc.Send[objc.ID](c_.ID, objc.Sel("noteNumberOfItemsChanged"))
}


// Marks the receiver as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/reloadData()
func (c_ ComboBox) ReloadData() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadData"))
}


// Removes all items from the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/removeAllItems()
func (c_ ComboBox) RemoveAllItems() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllItems"))
}


// Removes the object at the specified location from the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/removeItem(at:)
func (c_ ComboBox) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeItemAtIndex:"), index)
}


// Removes all occurrences of the given object from the receiver’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/removeItem(withObjectValue:)
func (c_ ComboBox) RemoveItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeItemWithObjectValue:"), object)
}


// Scrolls the receiver’s pop-up list vertically so that the item at the specified index is as close to the top as possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/scrollItemAtIndexToTop(_:)
func (c_ ComboBox) ScrollItemAtIndexToTop(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollItemAtIndexToTop:"), index)
}


// Scrolls the receiver’s pop-up list vertically so that the item at the specified index is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/scrollItemAtIndexToVisible(_:)
func (c_ ComboBox) ScrollItemAtIndexToVisible(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollItemAtIndexToVisible:"), index)
}


// Selects the pop-up list row at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectItem(at:)
func (c_ ComboBox) SelectItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectItemAtIndex:"), index)
}


// Selects the first pop-up list item that corresponds to the given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectItem(withObjectValue:)
func (c_ ComboBox) SelectItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectItemWithObjectValue:"), object)
}







// A Boolean value indicating whether the combo box tries to complete what the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/completes
func (c_ ComboBox) Completes() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("completes"))
	return rv
}


// A Boolean value indicating whether the combo box tries to complete what the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/completes
func (c_ ComboBox) SetCompletes(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletes:"), value)
}


// The object that provides the item data for the combo box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/dataSource
func (c_ ComboBox) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dataSource"))
	return rv
}


// The object that provides the item data for the combo box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/dataSource
func (c_ ComboBox) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}


// A Boolean value indicating whether the combo box has a vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/hasVerticalScroller
func (c_ ComboBox) HasVerticalScroller() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasVerticalScroller"))
	return rv
}


// A Boolean value indicating whether the combo box has a vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/hasVerticalScroller
func (c_ ComboBox) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasVerticalScroller:"), value)
}


// The index of the last item selected from the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/indexOfSelectedItem
func (c_ ComboBox) IndexOfSelectedItem() int {
	rv := objc.Send[int](c_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}


// The horizontal and vertical spacing between cells in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/intercellSpacing
func (c_ ComboBox) IntercellSpacing() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("intercellSpacing"))
	return rv
}


// The horizontal and vertical spacing between cells in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/intercellSpacing
func (c_ ComboBox) SetIntercellSpacing(value corefoundation.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntercellSpacing:"), value)
}


// A Boolean value indicating whether the combo box displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/isButtonBordered
func (c_ ComboBox) ButtonBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("buttonBordered"))
	return rv
}


// A Boolean value indicating whether the combo box displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/isButtonBordered
func (c_ ComboBox) SetButtonBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setButtonBordered:"), value)
}


// The height of each item in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/itemHeight
func (c_ ComboBox) ItemHeight() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("itemHeight"))
	return rv
}


// The height of each item in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/itemHeight
func (c_ ComboBox) SetItemHeight(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemHeight:"), value)
}


// The total number of items in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/numberOfItems
func (c_ ComboBox) NumberOfItems() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The maximum number of visible items to display in the pop-up list at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/numberOfVisibleItems
func (c_ ComboBox) NumberOfVisibleItems() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfVisibleItems"))
	return rv
}


// The maximum number of visible items to display in the pop-up list at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/numberOfVisibleItems
func (c_ ComboBox) SetNumberOfVisibleItems(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfVisibleItems:"), value)
}


// The object corresponding to the last item selected from the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/objectValueOfSelectedItem
func (c_ ComboBox) ObjectValueOfSelectedItem() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectValueOfSelectedItem"))
	return rv
}


// An array of the items from the combo box’s internal list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/objectValues
func (c_ ComboBox) ObjectValues() foundation.foundation.INSArray {
	rv := objc.Send[foundation.NSArray](c_.ID, objc.Sel("objectValues"))
	return rv
}


// A Boolean value indicating whether the combo box retrieves its items from a data source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/usesDataSource
func (c_ ComboBox) UsesDataSource() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesDataSource"))
	return rv
}


// A Boolean value indicating whether the combo box retrieves its items from a data source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox/usesDataSource
func (c_ ComboBox) SetUsesDataSource(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesDataSource:"), value)
}


// A Boolean value indicating whether the combo box displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/isbuttonbordered
func (c_ ComboBox) IsButtonBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isButtonBordered"))
	return rv
}


// A Boolean value indicating whether the combo box displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/isbuttonbordered
func (c_ ComboBox) SetIsButtonBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsButtonBordered:"), value)
}








