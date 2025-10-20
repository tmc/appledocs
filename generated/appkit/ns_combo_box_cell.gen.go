// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	AddItemsWithObjectValues(objects objc.ID)
	InsertItemWithObjectValueAtIndex(object objc.ID, index int)
	ItemObjectValueAtIndex(index int) objc.ID
	RemoveAllItems()
	RemoveItemAtIndex(index int)
}

// The user interface of a combo box.
//
// is a subclass of used to implement the user interface of “combo boxes” (see for information on how combo boxes look and work). The subclass of uses a single , and essentially all of the class’s methods simply invoke the corresponding method. Also see the protocol, which declares the methods that an object uses to access the contents of its data source object.
//
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

// Alloc allocates a new instance without initialization.
func (cc _ComboBoxCellClass) Alloc() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Adds multiple objects to the internal item list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/addItems(withObjectValues:)
func (c_ ComboBoxCell) AddItemsWithObjectValues(objects objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addItemsWithObjectValues:"), objects)
}

// Inserts an object at the specified location in the internal item list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/insertItem(withObjectValue:at:)
func (c_ ComboBoxCell) InsertItemWithObjectValueAtIndex(object objc.ID, index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertItemWithObjectValue:atIndex:"), object, index)
}

// Returns the object located at the specified location in the internal item list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/itemObjectValue(at:)
func (c_ ComboBoxCell) ItemObjectValueAtIndex(index int) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("itemObjectValueAtIndex:"), index)
	return rv
}

// Removes all items from the combo box’s internal item list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/removeAllItems()
func (c_ ComboBoxCell) RemoveAllItems() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllItems"))
}

// Removes the object at the specified location from the combo box’s internal item list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/removeItem(at:)
func (c_ ComboBoxCell) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeItemAtIndex:"), index)
}

// The object that provides the data displayed in the combo box’s pop-up list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/dataSource
func (c_ ComboBoxCell) DataSource() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}


// SetDataSource sets the value of the dataSource property.
// The object that provides the data displayed in the combo box’s pop-up list.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/dataSource
func (c_ ComboBoxCell) SetDataSource(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}
// The index of the last item selected from the pop-up list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/indexOfSelectedItem
func (c_ ComboBoxCell) IndexOfSelectedItem() int {
	rv := objc.Send[int](c_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}

// A Boolean value that indicates whether the combo box button displays a border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/isButtonBordered
func (c_ ComboBoxCell) ButtonBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("buttonBordered"))
	return rv
}


// SetButtonBordered sets the value of the buttonBordered property.
// A Boolean value that indicates whether the combo box button displays a border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/isButtonBordered
func (c_ ComboBoxCell) SetButtonBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setButtonBordered:"), value)
}
// The maximum number of items visible in the pop-up list at any one time.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/numberOfVisibleItems
func (c_ ComboBoxCell) NumberOfVisibleItems() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfVisibleItems"))
	return rv
}


// SetNumberOfVisibleItems sets the value of the numberOfVisibleItems property.
// The maximum number of items visible in the pop-up list at any one time.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/numberOfVisibleItems
func (c_ ComboBoxCell) SetNumberOfVisibleItems(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfVisibleItems:"), value)
}


