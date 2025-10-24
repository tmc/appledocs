// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSTableCellView */


/* debug [class_header]: Header for NSTableCellView */
// The class instance for the [TableCellView] class.
var (
	TableCellViewClass     _TableCellViewClass
	TableCellViewClassOnce sync.Once
)

func getTableCellViewClass() _TableCellViewClass {
	TableCellViewClassOnce.Do(func() {
		TableCellViewClass = _TableCellViewClass{objc.GetClass("NSTableCellView")}
	})
	return TableCellViewClass
}

type _TableCellViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TableCellView */
// An interface definition for the [TableCellView] class.
type ITableCellView interface {
	IView
	
/* debug [class_interface_properties]: Properties for TableCellView */
	// properties:
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	DraggingImageComponents() []DraggingImageComponent
	ImageView() IImageView
	SetImageView(value IImageView)
	ObjectValue() objc.ID
	SetObjectValue(value objc.ID)
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	TextField() ITextField
	SetTextField(value ITextField)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TableCellView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TableCellView */
// Alloc allocates a new instance without initialization.
func (tc _TableCellViewClass) Alloc() TableCellView {
	rv := objc.Send[TableCellView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TableCellViewClass) New() TableCellView {
	rv := objc.Send[TableCellView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableCellView) Init() TableCellView {
	rv := objc.Send[TableCellView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableCellView) Autorelease() TableCellView {
	rv := objc.Send[TableCellView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableCellView creates a new TableCellView instance.
func NewTableCellView() TableCellView {
	return getTableCellViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TableCellView */
// A reusable container view shown for a particular cell in a table view that uses rows for content.
//
// The and properties are connected in Interface Builder. Additional properties can be added by subclassing and adding the required properties and connecting them programmatically or in Interface Builder. The is used when setting the value of the view cell by the method in the . If you use your own custom view cells that are not based on you should implement this property in order to be able to receive changes to cell values.


// A reusable container view shown for a particular cell in a table view that uses rows for content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView
type TableCellView struct {
	View
}

// TableCellViewFrom constructs a [TableCellView] from an unsafe.Pointer.
//
// A reusable container view shown for a particular cell in a table view that uses rows for content.
func TableCellViewFrom(ptr unsafe.Pointer) TableCellView {
	return TableCellView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TableCellView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TableCellView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TableCellView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TableCellView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TableCellView */

// This property is automatically set by the enclosing row view to let this view know what its background looks like.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/backgroundStyle
func (t_ TableCellView) BackgroundStyle() BackgroundStyle {
	rv := objc.Send[BackgroundStyle](t_.ID, objc.Sel("backgroundStyle"))
	return rv
}/* debug [instance_properties/getter]: backgroundStyle */


// This property is automatically set by the enclosing row view to let this view know what its background looks like.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/backgroundStyle
func (t_ TableCellView) SetBackgroundStyle(value BackgroundStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundStyle:"), value)
}/* debug [instance_properties/setter]: backgroundStyle */


// Returns dragging images for the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/draggingImageComponents
func (t_ TableCellView) DraggingImageComponents() []DraggingImageComponent {
	rv := objc.Send[[]DraggingImageComponent](t_.ID, objc.Sel("draggingImageComponents"))
	return rv
}/* debug [instance_properties/getter]: draggingImageComponents */


// Image displayed by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/imageView
func (t_ TableCellView) ImageView() IImageView {
	rv := objc.Send[ImageView](t_.ID, objc.Sel("imageView"))
	return rv
}/* debug [instance_properties/getter]: imageView */


// Image displayed by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/imageView
func (t_ TableCellView) SetImageView(value IImageView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImageView:"), value)
}/* debug [instance_properties/setter]: imageView */


// The object that represents the cell data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/objectValue
func (t_ TableCellView) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("objectValue"))
	return rv
}/* debug [instance_properties/getter]: objectValue */


// The object that represents the cell data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/objectValue
func (t_ TableCellView) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setObjectValue:"), value)
}/* debug [instance_properties/setter]: objectValue */


// Returns the row size style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/rowSizeStyle
func (t_ TableCellView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Send[TableViewRowSizeStyle](t_.ID, objc.Sel("rowSizeStyle"))
	return rv
}/* debug [instance_properties/getter]: rowSizeStyle */


// Returns the row size style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/rowSizeStyle
func (t_ TableCellView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowSizeStyle:"), value)
}/* debug [instance_properties/setter]: rowSizeStyle */


// Text displayed by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/textField
func (t_ TableCellView) TextField() ITextField {
	rv := objc.Send[TextField](t_.ID, objc.Sel("textField"))
	return rv
}/* debug [instance_properties/getter]: textField */


// Text displayed by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView/textField
func (t_ TableCellView) SetTextField(value ITextField) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextField:"), value)
}/* debug [instance_properties/setter]: textField */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTableCellView */



