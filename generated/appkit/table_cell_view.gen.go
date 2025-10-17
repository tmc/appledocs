
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableCellView] class.
var TableCellViewClass _TableCellViewClass

func init() {
	TableCellViewClass = _TableCellViewClass{objc.GetClass("NSTableCellView")}
}

type _TableCellViewClass struct {
	objc.Class
}

// An interface definition for the [TableCellView] class.
type ITableCellView interface {
	ID() objc.ID
}

type TableCellView struct {
	id objc.ID
}

func TableCellViewFrom(ptr unsafe.Pointer) TableCellView {
	return TableCellView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TableCellView) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TableCellViewClass) Alloc() TableCellView {
	rv := objc.Send[TableCellView](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TableCellViewClass) New() TableCellView {
	rv := objc.Send[TableCellView](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTableCellView creates and returns a new initialized instance.
func NewTableCellView() TableCellView {
	return TableCellViewClass.New()
}

// Init initializes the instance.
func (t_ TableCellView) Init() TableCellView {
	rv := objc.Send[TableCellView](t_.ID(), selInit)
	return rv
}
// This property is automatically set by the enclosing row view to let this view know what its background looks like. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/backgroundStyle
func (t_ TableCellView) BackgroundStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("backgroundStyle"))
	return rv
}
// SetBackgroundStyle sets the value of the backgroundStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/backgroundStyle
func (t_ TableCellView) SetBackgroundStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setBackgroundStyle:"), value)
}
// Returns dragging images for the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/draggingImageComponents
func (t_ TableCellView) DraggingImageComponents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("draggingImageComponents"))
	return rv
}
// Image displayed by the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/imageView
func (t_ TableCellView) ImageView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("imageView"))
	return rv
}
// SetImageView sets the value of the imageView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/imageView
func (t_ TableCellView) SetImageView(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setImageView:"), value)
}
// The object that represents the cell data. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/objectValue
func (t_ TableCellView) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](t_.ID(), objc.RegisterName("objectValue"))
	return rv
}
// SetObjectValue sets the value of the objectValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/objectValue
func (t_ TableCellView) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setObjectValue:"), value)
}
// Returns the row size style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/rowSizeStyle
func (t_ TableCellView) RowSizeStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("rowSizeStyle"))
	return rv
}
// SetRowSizeStyle sets the value of the rowSizeStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/rowSizeStyle
func (t_ TableCellView) SetRowSizeStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setRowSizeStyle:"), value)
}
// Text displayed by the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/textField
func (t_ TableCellView) TextField() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("textField"))
	return rv
}
// SetTextField sets the value of the textField property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableCellView/textField
func (t_ TableCellView) SetTextField(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setTextField:"), value)
}
