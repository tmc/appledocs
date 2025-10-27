// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ComboBoxCell] class.
var (
	ComboBoxCellClass     _ComboBoxCellClass
	ComboBoxCellClassOnce sync.Once
)

func getComboBoxCellClass() _ComboBoxCellClass {
	ComboBoxCellClassOnce.Do(func() {
		ComboBoxCellClass = _ComboBoxCellClass{objc.GetClass("NSComboBoxCell")}
	})
	return ComboBoxCellClass
}

type _ComboBoxCellClass struct {
	class objc.Class
}





// An interface definition for the [ComboBoxCell] class.
type IComboBoxCell interface {
	ITextFieldCell
	

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
	CompletedString(string_ foundation.foundation.INSString) foundation.String
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
func (cc _ComboBoxCellClass) Alloc() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComboBoxCellClass) New() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComboBoxCell) Init() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComboBoxCell) Autorelease() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComboBoxCell creates a new ComboBoxCell instance.
func NewComboBoxCell() ComboBoxCell {
	return getComboBoxCellClass().New()
}





// The user interface of a combo box.
//
// is a subclass of used to implement the user interface of “combo boxes” (see for information on how combo boxes look and work). The subclass of uses a single , and essentially all of the class’s methods simply invoke the corresponding method. Also see the protocol, which declares the methods that an object uses to access the contents of its data source object.


// The user interface of a combo box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell
type ComboBoxCell struct {
	TextFieldCell
}

// ComboBoxCellFrom constructs a [ComboBoxCell] from an unsafe.Pointer.
//
// The user interface of a combo box.
func ComboBoxCellFrom(ptr unsafe.Pointer) ComboBoxCell {
	return ComboBoxCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}




















// Adds the specified object to the internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/addItem(withObjectValue:)
func (c_ ComboBoxCell) AddItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addItemWithObjectValue:"), object)
}


// Adds multiple objects to the internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/addItems(withObjectValues:)
func (c_ ComboBoxCell) AddItemsWithObjectValues(objects foundation.foundation.INSArray) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addItemsWithObjectValues:"), objects)
}


// Returns a string from the combo box’s pop-up list that starts with the given substring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/completedString(_:)
func (c_ ComboBoxCell) CompletedString(string_ foundation.foundation.INSString) foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("completedString:"), string_)
	return rv
}


// Deselects the pop-up list item at the given index if it’s selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/deselectItem(at:)
func (c_ ComboBoxCell) DeselectItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deselectItemAtIndex:"), index)
}


// Searches the combo box’s internal item list for the given object and returns the matching index number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/indexOfItem(withObjectValue:)
func (c_ ComboBoxCell) IndexOfItemWithObjectValue(object objectivec.IObject) int {
	rv := objc.Send[int](c_.ID, objc.Sel("indexOfItemWithObjectValue:"), object)
	return rv
}


// Inserts an object at the specified location in the internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/insertItem(withObjectValue:at:)
func (c_ ComboBoxCell) InsertItemWithObjectValueAtIndex(object objectivec.IObject, index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertItemWithObjectValue:atIndex:"), object, index)
}


// Returns the object located at the specified location in the internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/itemObjectValue(at:)
func (c_ ComboBoxCell) ItemObjectValueAtIndex(index int) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("itemObjectValueAtIndex:"), index)
	return rv
}


// Informs the combo box that the number of items in its data source has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/noteNumberOfItemsChanged()
func (c_ ComboBoxCell) NoteNumberOfItemsChanged() {
	objc.Send[objc.ID](c_.ID, objc.Sel("noteNumberOfItemsChanged"))
}


// Marks the combo box as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/reloadData()
func (c_ ComboBoxCell) ReloadData() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadData"))
}


// Removes all items from the combo box’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/removeAllItems()
func (c_ ComboBoxCell) RemoveAllItems() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllItems"))
}


// Removes the object at the specified location from the combo box’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/removeItem(at:)
func (c_ ComboBoxCell) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeItemAtIndex:"), index)
}


// Removes all occurrences of the specified object from the combo box’s internal item list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/removeItem(withObjectValue:)
func (c_ ComboBoxCell) RemoveItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeItemWithObjectValue:"), object)
}


// Scrolls the combo box’s pop-up list vertically so that the item at the given index is as close to the top as possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/scrollItemAtIndexToTop(_:)
func (c_ ComboBoxCell) ScrollItemAtIndexToTop(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollItemAtIndexToTop:"), index)
}


// Scrolls the combo box’s pop-up list vertically so that the item at the given index is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/scrollItemAtIndexToVisible(_:)
func (c_ ComboBoxCell) ScrollItemAtIndexToVisible(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollItemAtIndexToVisible:"), index)
}


// Selects the pop-up list row at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/selectItem(at:)
func (c_ ComboBoxCell) SelectItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectItemAtIndex:"), index)
}


// Selects the first pop-up list item that corresponds to the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/selectItem(withObjectValue:)
func (c_ ComboBoxCell) SelectItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectItemWithObjectValue:"), object)
}







// A Boolean value that indicates if the combo box tries to complete text entered by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/completes
func (c_ ComboBoxCell) Completes() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("completes"))
	return rv
}


// A Boolean value that indicates if the combo box tries to complete text entered by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/completes
func (c_ ComboBoxCell) SetCompletes(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletes:"), value)
}


// The object that provides the data displayed in the combo box’s pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/dataSource
func (c_ ComboBoxCell) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dataSource"))
	return rv
}


// The object that provides the data displayed in the combo box’s pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/dataSource
func (c_ ComboBoxCell) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}


// A Boolean value that indicates if the combo box displays a vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/hasVerticalScroller
func (c_ ComboBoxCell) HasVerticalScroller() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasVerticalScroller"))
	return rv
}


// A Boolean value that indicates if the combo box displays a vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/hasVerticalScroller
func (c_ ComboBoxCell) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasVerticalScroller:"), value)
}


// The index of the last item selected from the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/indexOfSelectedItem
func (c_ ComboBoxCell) IndexOfSelectedItem() int {
	rv := objc.Send[int](c_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}


// The spacing between cells in the combo box’s pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/intercellSpacing
func (c_ ComboBoxCell) IntercellSpacing() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("intercellSpacing"))
	return rv
}


// The spacing between cells in the combo box’s pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/intercellSpacing
func (c_ ComboBoxCell) SetIntercellSpacing(value corefoundation.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntercellSpacing:"), value)
}


// A Boolean value that indicates whether the combo box button displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/isButtonBordered
func (c_ ComboBoxCell) ButtonBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("buttonBordered"))
	return rv
}


// A Boolean value that indicates whether the combo box button displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/isButtonBordered
func (c_ ComboBoxCell) SetButtonBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setButtonBordered:"), value)
}


// The height of each item in the combo box’s pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/itemHeight
func (c_ ComboBoxCell) ItemHeight() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("itemHeight"))
	return rv
}


// The height of each item in the combo box’s pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/itemHeight
func (c_ ComboBoxCell) SetItemHeight(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemHeight:"), value)
}


// The total number of items in the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/numberOfItems
func (c_ ComboBoxCell) NumberOfItems() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The maximum number of items visible in the pop-up list at any one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/numberOfVisibleItems
func (c_ ComboBoxCell) NumberOfVisibleItems() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfVisibleItems"))
	return rv
}


// The maximum number of items visible in the pop-up list at any one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/numberOfVisibleItems
func (c_ ComboBoxCell) SetNumberOfVisibleItems(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfVisibleItems:"), value)
}


// The object corresponding to the last item selected from the pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/objectValueOfSelectedItem
func (c_ ComboBoxCell) ObjectValueOfSelectedItem() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectValueOfSelectedItem"))
	return rv
}


// The combo box’s internal item list in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/objectValues
func (c_ ComboBoxCell) ObjectValues() foundation.foundation.INSArray {
	rv := objc.Send[foundation.NSArray](c_.ID, objc.Sel("objectValues"))
	return rv
}


// A Boolean value that indicates if the combo box uses an external data source to populate its pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/usesDataSource
func (c_ ComboBoxCell) UsesDataSource() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesDataSource"))
	return rv
}


// A Boolean value that indicates if the combo box uses an external data source to populate its pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/usesDataSource
func (c_ ComboBoxCell) SetUsesDataSource(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesDataSource:"), value)
}


// A Boolean value that indicates whether the combo box button displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/isbuttonbordered
func (c_ ComboBoxCell) IsButtonBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isButtonBordered"))
	return rv
}


// A Boolean value that indicates whether the combo box button displays a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/isbuttonbordered
func (c_ ComboBoxCell) SetIsButtonBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsButtonBordered:"), value)
}








