// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TextFieldCell] class.
var (
	TextFieldCellClass     _TextFieldCellClass
	TextFieldCellClassOnce sync.Once
)

func getTextFieldCellClass() _TextFieldCellClass {
	TextFieldCellClassOnce.Do(func() {
		TextFieldCellClass = _TextFieldCellClass{objc.GetClass("NSTextFieldCell")}
	})
	return TextFieldCellClass
}

type _TextFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [TextFieldCell] class.
type ITextFieldCell interface {
	IActionCell
	// properties:
	PlaceholderString() objc.IObject /* cross-framework: NSString */
	SetPlaceholderString(value objc.IObject /* cross-framework: NSString */)
	AllowedInputSourceLocales() objc.IObject /* cross-framework: NSString */
	SetAllowedInputSourceLocales(value objc.IObject /* cross-framework: NSString */)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BezelStyle() unsafe.Pointer
	SetBezelStyle(value unsafe.Pointer)
	DrawsBackground() bool /* primitive/slice/pointer. */
	SetDrawsBackground(value bool /* primitive/slice/pointer. */)
	PlaceholderAttributedString() objc.IObject /* cross-framework: AttributedString */
	SetPlaceholderAttributedString(value objc.IObject /* cross-framework: AttributedString */)
	TextColor() IColor
	SetTextColor(value IColor)
	// methods:
}

// An object that enhances the text display capabilities of a cell.
//
// The class adds to the text display capabilities of the class by allowing you to set the color of both the text and its background. You can also specify whether the cell draws its background at all. All of the methods declared by this class are also declared by the class, which uses objects to draw and edit text. The cover methods call the corresponding methods. Placeholder strings, set using the or property, appear in the text field cell if the actual string is or an empty string. They’re drawn in gray on the cell and aren’t archived in the “pre-10.2” nib format.


// An object that enhances the text display capabilities of a cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell
type TextFieldCell struct {
	ActionCell
}

// TextFieldCellFrom constructs a [TextFieldCell] from an unsafe.Pointer.
//
// An object that enhances the text display capabilities of a cell.
func TextFieldCellFrom(ptr unsafe.Pointer) TextFieldCell {
	return TextFieldCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextFieldCellClass) Alloc() TextFieldCell {
	rv := objc.Send[TextFieldCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextFieldCellClass) New() TextFieldCell {
	rv := objc.Send[TextFieldCell](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextFieldCell) Init() TextFieldCell {
	rv := objc.Send[TextFieldCell](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextFieldCell) Autorelease() TextFieldCell {
	rv := objc.Send[TextFieldCell](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextFieldCell creates a new TextFieldCell instance.
func NewTextFieldCell() TextFieldCell {
	return getTextFieldCellClass().New()
}



// The placeholder text for the cell, specified as a plain text string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderString
func (t_ TextFieldCell) PlaceholderString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("placeholderString"))
	return rv
}


// The placeholder text for the cell, specified as a plain text string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderString
func (t_ TextFieldCell) SetPlaceholderString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderString:"), value)
}


// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/allowedinputsourcelocales
func (t_ TextFieldCell) AllowedInputSourceLocales() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}


// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/allowedinputsourcelocales
func (t_ TextFieldCell) SetAllowedInputSourceLocales(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedInputSourceLocales:"), value)
}


// The color of the cell’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/backgroundcolor
func (t_ TextFieldCell) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The color of the cell’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/backgroundcolor
func (t_ TextFieldCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The bezel style to use when drawing the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/bezelstyle
func (t_ TextFieldCell) BezelStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("bezelStyle"))
	return rv
}


// The bezel style to use when drawing the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/bezelstyle
func (t_ TextFieldCell) SetBezelStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezelStyle:"), value)
}


// A Boolean value that indicates whether the cell draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/drawsbackground
func (t_ TextFieldCell) DrawsBackground() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that indicates whether the cell draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/drawsbackground
func (t_ TextFieldCell) SetDrawsBackground(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The placeholder text for the cell, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/placeholderattributedstring
func (t_ TextFieldCell) PlaceholderAttributedString() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[AttributedString](t_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// The placeholder text for the cell, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/placeholderattributedstring
func (t_ TextFieldCell) SetPlaceholderAttributedString(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}


// The color to use to draw the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/textcolor
func (t_ TextFieldCell) TextColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("textColor"))
	return rv
}


// The color to use to draw the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/textcolor
func (t_ TextFieldCell) SetTextColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:"), value)
}



