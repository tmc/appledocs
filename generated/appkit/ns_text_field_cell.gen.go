// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	SetWantsNotificationForMarkedText(flag bool)
}

// An object that enhances the text display capabilities of a cell.
//
// The class adds to the text display capabilities of the class by allowing you to set the color of both the text and its background. You can also specify whether the cell draws its background at all. All of the methods declared by this class are also declared by the class, which uses objects to draw and edit text. The cover methods call the corresponding methods. Placeholder strings, set using the or property, appear in the text field cell if the actual string is or an empty string. They’re drawn in gray on the cell and aren’t archived in the “pre-10.2” nib format.
//
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


// Directs the cell’s associated field editor to post text change notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/setWantsNotificationForMarkedText(_:)
func (t_ TextFieldCell) SetWantsNotificationForMarkedText(flag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWantsNotificationForMarkedText:"), flag)
}

// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/allowedInputSourceLocales
func (t_ TextFieldCell) AllowedInputSourceLocales() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}


// SetAllowedInputSourceLocales sets the value of the allowedInputSourceLocales property.
// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/allowedInputSourceLocales
func (t_ TextFieldCell) SetAllowedInputSourceLocales(value []string) {
	// Convert Go slice to NSArray
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/backgroundColor
func (t_ TextFieldCell) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The color of the cell’s background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/backgroundColor
func (t_ TextFieldCell) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The bezel style to use when drawing the text field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/bezelStyle
func (t_ TextFieldCell) BezelStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("bezelStyle"))
	return rv
}


// SetBezelStyle sets the value of the bezelStyle property.
// The bezel style to use when drawing the text field.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/bezelStyle
func (t_ TextFieldCell) SetBezelStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezelStyle:"), value)
}

// The placeholder text for the cell, specified as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderAttributedString
func (t_ TextFieldCell) PlaceholderAttributedString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// SetPlaceholderAttributedString sets the value of the placeholderAttributedString property.
// The placeholder text for the cell, specified as an attributed string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderAttributedString
func (t_ TextFieldCell) SetPlaceholderAttributedString(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}

// The placeholder text for the cell, specified as a plain text string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderString
func (t_ TextFieldCell) PlaceholderString() string {
	rv := objc.Send[string](t_.ID, objc.Sel("placeholderString"))
	return rv
}


// SetPlaceholderString sets the value of the placeholderString property.
// The placeholder text for the cell, specified as a plain text string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderString
func (t_ TextFieldCell) SetPlaceholderString(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}

// The color to use to draw the cell’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/textColor
func (t_ TextFieldCell) TextColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textColor"))
	return rv
}


// SetTextColor sets the value of the textColor property.
// The color to use to draw the cell’s text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/textColor
func (t_ TextFieldCell) SetTextColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:"), value)
}



