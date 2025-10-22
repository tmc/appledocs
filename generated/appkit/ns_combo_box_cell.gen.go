// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	AddItemsWithObjectValues(objects objectivec.IObject)
	InsertItemWithObjectValueAtIndex(object objectivec.IObject, index int)
	ItemObjectValueAtIndex(index int) objc.ID
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	DataSource() objc.ID
	SetDataSource(value objc.ID)
	IndexOfSelectedItem() int
	ButtonBordered() bool
	SetButtonBordered(value bool)
	ItemHeight() float64
	SetItemHeight(value float64)
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
	Completes() bool
	SetCompletes(value bool)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	IntercellSpacing() coregraphics.CGSize
	SetIntercellSpacing(value coregraphics.CGSize)
	IsButtonBordered() bool
	SetIsButtonBordered(value bool)
	NumberOfItems() int
	SetNumberOfItems(value int)
	ObjectValueOfSelectedItem() unsafe.Pointer
	SetObjectValueOfSelectedItem(value unsafe.Pointer)
	ObjectValues() unsafe.Pointer
	SetObjectValues(value unsafe.Pointer)
	UsesDataSource() bool
	SetUsesDataSource(value bool)
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
func (c_ ComboBoxCell) AddItemsWithObjectValues(objects objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addItemsWithObjectValues:"), objects)
}

// Inserts an object at the specified location in the internal item list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/insertItem(withObjectValue:at:)
func (c_ ComboBoxCell) InsertItemWithObjectValueAtIndex(object objectivec.IObject, index int) {
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

// The height of each item in the combo box’s pop-up list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/itemHeight
func (c_ ComboBoxCell) ItemHeight() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("itemHeight"))
	return rv
}


// SetItemHeight sets the value of the itemHeight property.
// The height of each item in the combo box’s pop-up list.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/itemHeight
func (c_ ComboBoxCell) SetItemHeight(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemHeight:"), value)
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

// A Boolean value that indicates if the combo box tries to complete text entered by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/completes
func (c_ ComboBoxCell) Completes() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("completes"))
	return rv
}


// SetCompletes sets the value of the completes property.
// A Boolean value that indicates if the combo box tries to complete text entered by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/completes
func (c_ ComboBoxCell) SetCompletes(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletes:"), value)
}

// A Boolean value that indicates if the combo box displays a vertical scroller.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/hasverticalscroller
func (c_ ComboBoxCell) HasVerticalScroller() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasVerticalScroller"))
	return rv
}


// SetHasVerticalScroller sets the value of the hasVerticalScroller property.
// A Boolean value that indicates if the combo box displays a vertical scroller.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/hasverticalscroller
func (c_ ComboBoxCell) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasVerticalScroller:"), value)
}

// The spacing between cells in the combo box’s pop-up list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/intercellspacing
func (c_ ComboBoxCell) IntercellSpacing() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("intercellSpacing"))
	return rv
}


// SetIntercellSpacing sets the value of the intercellSpacing property.
// The spacing between cells in the combo box’s pop-up list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/intercellspacing
func (c_ ComboBoxCell) SetIntercellSpacing(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntercellSpacing:"), value)
}

// A Boolean value that indicates whether the combo box button displays a border.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/isbuttonbordered
func (c_ ComboBoxCell) IsButtonBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isButtonBordered"))
	return rv
}


// SetIsButtonBordered sets the value of the isButtonBordered property.
// A Boolean value that indicates whether the combo box button displays a border.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/isbuttonbordered
func (c_ ComboBoxCell) SetIsButtonBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsButtonBordered:"), value)
}

// The total number of items in the pop-up list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/numberofitems
func (c_ ComboBoxCell) NumberOfItems() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfItems"))
	return rv
}


// SetNumberOfItems sets the value of the numberOfItems property.
// The total number of items in the pop-up list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/numberofitems
func (c_ ComboBoxCell) SetNumberOfItems(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfItems:"), value)
}

// The object corresponding to the last item selected from the pop-up list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/objectvalueofselecteditem
func (c_ ComboBoxCell) ObjectValueOfSelectedItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("objectValueOfSelectedItem"))
	return rv
}


// SetObjectValueOfSelectedItem sets the value of the objectValueOfSelectedItem property.
// The object corresponding to the last item selected from the pop-up list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/objectvalueofselecteditem
func (c_ ComboBoxCell) SetObjectValueOfSelectedItem(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectValueOfSelectedItem:"), value)
}

// The combo box’s internal item list in an array.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/objectvalues
func (c_ ComboBoxCell) ObjectValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("objectValues"))
	return rv
}


// SetObjectValues sets the value of the objectValues property.
// The combo box’s internal item list in an array.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/objectvalues
func (c_ ComboBoxCell) SetObjectValues(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectValues:"), value)
}

// A Boolean value that indicates if the combo box uses an external data source to populate its pop-up list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/usesdatasource
func (c_ ComboBoxCell) UsesDataSource() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesDataSource"))
	return rv
}


// SetUsesDataSource sets the value of the usesDataSource property.
// A Boolean value that indicates if the combo box uses an external data source to populate its pop-up list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/usesdatasource
func (c_ ComboBoxCell) SetUsesDataSource(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesDataSource:"), value)
}



