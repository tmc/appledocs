// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextFieldCell */


/* debug [class_header]: Header for NSTextFieldCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextFieldCell */
// An interface definition for the [TextFieldCell] class.
type ITextFieldCell interface {
	IActionCell
	
/* debug [class_interface_properties]: Properties for TextFieldCell */
	// properties:
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	PlaceholderString() objc.IObject /* cross-framework: NSString */
	SetPlaceholderString(value objc.IObject /* cross-framework: NSString */)
	TextColor() IColor
	SetTextColor(value IColor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextFieldCell */
	// methods:
	SetUpFieldEditorAttributes(textObj IText) IText
	SetWantsNotificationForMarkedText(flag bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextFieldCell */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextFieldCell */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextFieldCell */

// Initializes a text field cell that displays the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(textCell:)
func NewTextFieldCellTextCell(string_ objc.IObject /* cross-framework: NSString */) TextFieldCell {
	instance := getTextFieldCellClass().Alloc()
	rv := objc.Send[TextFieldCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextFieldCellTextCell */


// Initializes a text field cell from data in the provided unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(coder:)
func NewTextFieldCellWithCoder(coder foundation.Coder) TextFieldCell {
	instance := getTextFieldCellClass().Alloc()
	rv := objc.Send[TextFieldCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextFieldCellWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextFieldCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextFieldCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextFieldCell */

// Allows the cell to set up the field editor’s attributes before editing begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/setUpFieldEditorAttributes(_:)
func (t_ TextFieldCell) SetUpFieldEditorAttributes(textObj IText) IText {
	rv := objc.Send[Text](t_.ID, objc.Sel("setUpFieldEditorAttributes:"), textObj)
	return rv
}/* debug [instance_methods/method]: SetUpFieldEditorAttributes */


// Directs the cell’s associated field editor to post text change notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/setWantsNotificationForMarkedText(_:)
func (t_ TextFieldCell) SetWantsNotificationForMarkedText(flag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWantsNotificationForMarkedText:"), flag)
}/* debug [instance_methods/method]: SetWantsNotificationForMarkedText */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextFieldCell */

// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/allowedInputSourceLocales
func (t_ TextFieldCell) AllowedInputSourceLocales() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}/* debug [instance_properties/getter]: allowedInputSourceLocales */


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
}/* debug [instance_properties/setter]: allowedInputSourceLocales */


// The color of the cell’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/backgroundColor
func (t_ TextFieldCell) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The color of the cell’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/backgroundColor
func (t_ TextFieldCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The bezel style to use when drawing the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/bezelStyle
func (t_ TextFieldCell) BezelStyle() TextFieldBezelStyle {
	rv := objc.Send[TextFieldBezelStyle](t_.ID, objc.Sel("bezelStyle"))
	return rv
}/* debug [instance_properties/getter]: bezelStyle */


// The bezel style to use when drawing the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/bezelStyle
func (t_ TextFieldCell) SetBezelStyle(value TextFieldBezelStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezelStyle:"), value)
}/* debug [instance_properties/setter]: bezelStyle */


// A Boolean value that indicates whether the cell draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/drawsBackground
func (t_ TextFieldCell) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}/* debug [instance_properties/getter]: drawsBackground */


// A Boolean value that indicates whether the cell draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/drawsBackground
func (t_ TextFieldCell) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}/* debug [instance_properties/setter]: drawsBackground */


// The placeholder text for the cell, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderAttributedString
func (t_ TextFieldCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}/* debug [instance_properties/getter]: placeholderAttributedString */


// The placeholder text for the cell, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderAttributedString
func (t_ TextFieldCell) SetPlaceholderAttributedString(value foundation.AttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}/* debug [instance_properties/setter]: placeholderAttributedString */


// The placeholder text for the cell, specified as a plain text string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderString
func (t_ TextFieldCell) PlaceholderString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("placeholderString"))
	return rv
}/* debug [instance_properties/getter]: placeholderString */


// The placeholder text for the cell, specified as a plain text string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderString
func (t_ TextFieldCell) SetPlaceholderString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderString:"), value)
}/* debug [instance_properties/setter]: placeholderString */


// The color to use to draw the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/textColor
func (t_ TextFieldCell) TextColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("textColor"))
	return rv
}/* debug [instance_properties/getter]: textColor */


// The color to use to draw the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/textColor
func (t_ TextFieldCell) SetTextColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:"), value)
}/* debug [instance_properties/setter]: textColor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextFieldCell */


