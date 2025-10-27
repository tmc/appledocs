// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	PlaceholderAttributedString() foundation.foundation.INSAttributedString
	SetPlaceholderAttributedString(value foundation.foundation.INSAttributedString)
	PlaceholderString() foundation.foundation.INSString
	SetPlaceholderString(value foundation.foundation.INSString)
	TextColor() IColor
	SetTextColor(value IColor)


	

	// methods:
	SetUpFieldEditorAttributes(textObj IText) IText
	SetWantsNotificationForMarkedText(flag bool)


}





// Alloc allocates a new instance without initialization.
func (tc _TextFieldCellClass) Alloc() TextFieldCell {
	rv := objc.Send[TextFieldCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Initializes a text field cell that displays the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(textCell:)
func NewTextFieldCellTextCell(string_ foundation.foundation.INSString) TextFieldCell {
	instance := getTextFieldCellClass().Alloc()
	rv := objc.Send[TextFieldCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}


// Initializes a text field cell from data in the provided unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(coder:)
func NewTextFieldCellWithCoder(coder foundation.foundation.INSCoder) TextFieldCell {
	instance := getTextFieldCellClass().Alloc()
	rv := objc.Send[TextFieldCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

















// Allows the cell to set up the field editor’s attributes before editing begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/setUpFieldEditorAttributes(_:)
func (t_ TextFieldCell) SetUpFieldEditorAttributes(textObj IText) IText {
	rv := objc.Send[Text](t_.ID, objc.Sel("setUpFieldEditorAttributes:"), textObj)
	return rv
}


// Directs the cell’s associated field editor to post text change notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/setWantsNotificationForMarkedText(_:)
func (t_ TextFieldCell) SetWantsNotificationForMarkedText(flag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWantsNotificationForMarkedText:"), flag)
}







// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/allowedInputSourceLocales
func (t_ TextFieldCell) AllowedInputSourceLocales() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}


// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/allowedInputSourceLocales
func (t_ TextFieldCell) SetAllowedInputSourceLocales(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedInputSourceLocales:"), nsArray)
}


// The color of the cell’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/backgroundColor
func (t_ TextFieldCell) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The color of the cell’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/backgroundColor
func (t_ TextFieldCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The bezel style to use when drawing the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/bezelStyle
func (t_ TextFieldCell) BezelStyle() TextFieldBezelStyle {
	rv := objc.Send[TextFieldBezelStyle](t_.ID, objc.Sel("bezelStyle"))
	return rv
}


// The bezel style to use when drawing the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/bezelStyle
func (t_ TextFieldCell) SetBezelStyle(value TextFieldBezelStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezelStyle:"), value)
}


// A Boolean value that indicates whether the cell draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/drawsBackground
func (t_ TextFieldCell) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that indicates whether the cell draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/drawsBackground
func (t_ TextFieldCell) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The placeholder text for the cell, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderAttributedString
func (t_ TextFieldCell) PlaceholderAttributedString() foundation.foundation.INSAttributedString {
	rv := objc.Send[foundation.NSAttributedString](t_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// The placeholder text for the cell, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderAttributedString
func (t_ TextFieldCell) SetPlaceholderAttributedString(value foundation.foundation.INSAttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}


// The placeholder text for the cell, specified as a plain text string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderString
func (t_ TextFieldCell) PlaceholderString() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("placeholderString"))
	return rv
}


// The placeholder text for the cell, specified as a plain text string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderString
func (t_ TextFieldCell) SetPlaceholderString(value foundation.foundation.INSString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderString:"), value)
}


// The color to use to draw the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/textColor
func (t_ TextFieldCell) TextColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("textColor"))
	return rv
}


// The color to use to draw the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/textColor
func (t_ TextFieldCell) SetTextColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:"), value)
}







