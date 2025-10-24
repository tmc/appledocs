// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFormCell */


/* debug [class_header]: Header for NSFormCell */
// The class instance for the [FormCell] class.
var (
	FormCellClass     _FormCellClass
	FormCellClassOnce sync.Once
)

func getFormCellClass() _FormCellClass {
	FormCellClassOnce.Do(func() {
		FormCellClass = _FormCellClass{objc.GetClass("NSFormCell")}
	})
	return FormCellClass
}

type _FormCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FormCell */
// An interface definition for the [FormCell] class.
type IFormCell interface {
	IActionCell
	
/* debug [class_interface_properties]: Properties for FormCell */
	// properties:
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	Opaque() bool
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	PlaceholderString() objc.IObject /* cross-framework: NSString */
	SetPlaceholderString(value objc.IObject /* cross-framework: NSString */)
	PreferredTextFieldWidth() float64
	SetPreferredTextFieldWidth(value float64)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	TitleAlignment() TextAlignment
	SetTitleAlignment(value TextAlignment)
	TitleBaseWritingDirection() WritingDirection
	SetTitleBaseWritingDirection(value WritingDirection)
	TitleFont() IFont
	SetTitleFont(value IFont)
	TitleWidth() float64
	SetTitleWidth(value float64)
	IsOpaque() bool
	SetIsOpaque(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FormCell */
	// methods:
	TitleWidth(size Size /* not a class type */) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FormCell */
// Alloc allocates a new instance without initialization.
func (fc _FormCellClass) Alloc() FormCell {
	rv := objc.Send[FormCell](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FormCellClass) New() FormCell {
	rv := objc.Send[FormCell](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FormCell) Init() FormCell {
	rv := objc.Send[FormCell](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FormCell) Autorelease() FormCell {
	rv := objc.Send[FormCell](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFormCell creates a new FormCell instance.
func NewFormCell() FormCell {
	return getFormCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FormCell */
// The class is used to implement text entry fields in a form. The left part of an object contains a title. The right part contains an editable text entry field.
//
// An object implements the user interface of an object.


// The class is used to implement text entry fields in a form. The left part of an object contains a title. The right part contains an editable text entry field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell
type FormCell struct {
	ActionCell
}

// FormCellFrom constructs a [FormCell] from an unsafe.Pointer.
//
// The class is used to implement text entry fields in a form. The left part of an object contains a title. The right part contains an editable text entry field.
func FormCellFrom(ptr unsafe.Pointer) FormCell {
	return FormCell{
		ActionCell: ActionCellFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FormCell */

// Returns an object initialized with the specified title string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/init(textCell:)
func NewFormCellTextCell(string_ objc.IObject /* cross-framework: NSString */) FormCell {
	instance := getFormCellClass().Alloc()
	rv := objc.Send[FormCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFormCellTextCell */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/init(coder:)
func NewFormCellWithCoder(coder foundation.Coder) FormCell {
	instance := getFormCellClass().Alloc()
	rv := objc.Send[FormCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFormCellWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FormCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FormCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FormCell */

// Returns the width of the title field constrained to the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleWidth(_:)
func (f_ FormCell) TitleWidth(size Size /* not a class type */) float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("titleWidth:"), size)
	return rv
}/* debug [instance_methods/method]: TitleWidth */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FormCell */

// The title of the cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/attributedTitle
func (f_ FormCell) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](f_.ID, objc.Sel("attributedTitle"))
	return rv
}/* debug [instance_properties/getter]: attributedTitle */


// The title of the cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/attributedTitle
func (f_ FormCell) SetAttributedTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAttributedTitle:"), value)
}/* debug [instance_properties/setter]: attributedTitle */


// A Boolean value indicating whether the title is empty and an opaque bezel is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/isOpaque
func (f_ FormCell) Opaque() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("opaque"))
	return rv
}/* debug [instance_properties/getter]: opaque */


// The cell’s attributed placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/placeholderAttributedString
func (f_ FormCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](f_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}/* debug [instance_properties/getter]: placeholderAttributedString */


// The cell’s attributed placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/placeholderAttributedString
func (f_ FormCell) SetPlaceholderAttributedString(value foundation.AttributedString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}/* debug [instance_properties/setter]: placeholderAttributedString */


// The cell’s plain text placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/placeholderString
func (f_ FormCell) PlaceholderString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("placeholderString"))
	return rv
}/* debug [instance_properties/getter]: placeholderString */


// The cell’s plain text placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/placeholderString
func (f_ FormCell) SetPlaceholderString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPlaceholderString:"), value)
}/* debug [instance_properties/setter]: placeholderString */


// The preferred text field width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/preferredTextFieldWidth
func (f_ FormCell) PreferredTextFieldWidth() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("preferredTextFieldWidth"))
	return rv
}/* debug [instance_properties/getter]: preferredTextFieldWidth */


// The preferred text field width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/preferredTextFieldWidth
func (f_ FormCell) SetPreferredTextFieldWidth(value float64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPreferredTextFieldWidth:"), value)
}/* debug [instance_properties/setter]: preferredTextFieldWidth */


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/title
func (f_ FormCell) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/title
func (f_ FormCell) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The alignment of the title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleAlignment
func (f_ FormCell) TitleAlignment() TextAlignment {
	rv := objc.Send[TextAlignment](f_.ID, objc.Sel("titleAlignment"))
	return rv
}/* debug [instance_properties/getter]: titleAlignment */


// The alignment of the title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleAlignment
func (f_ FormCell) SetTitleAlignment(value TextAlignment) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleAlignment:"), value)
}/* debug [instance_properties/setter]: titleAlignment */


// The default writing direction used to render the form cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleBaseWritingDirection
func (f_ FormCell) TitleBaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](f_.ID, objc.Sel("titleBaseWritingDirection"))
	return rv
}/* debug [instance_properties/getter]: titleBaseWritingDirection */


// The default writing direction used to render the form cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleBaseWritingDirection
func (f_ FormCell) SetTitleBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleBaseWritingDirection:"), value)
}/* debug [instance_properties/setter]: titleBaseWritingDirection */


// The font used to draw cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleFont
func (f_ FormCell) TitleFont() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("titleFont"))
	return rv
}/* debug [instance_properties/getter]: titleFont */


// The font used to draw cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleFont
func (f_ FormCell) SetTitleFont(value IFont) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleFont:"), value)
}/* debug [instance_properties/setter]: titleFont */


// The width of the title field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleWidth
func (f_ FormCell) TitleWidth() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("titleWidth"))
	return rv
}/* debug [instance_properties/getter]: titleWidth */


// The width of the title field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleWidth
func (f_ FormCell) SetTitleWidth(value float64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleWidth:"), value)
}/* debug [instance_properties/setter]: titleWidth */


// A Boolean value indicating whether the title is empty and an opaque bezel is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/isopaque
func (f_ FormCell) IsOpaque() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isOpaque"))
	return rv
}/* debug [instance_properties/getter]: isOpaque */


// A Boolean value indicating whether the title is empty and an opaque bezel is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/isopaque
func (f_ FormCell) SetIsOpaque(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsOpaque:"), value)
}/* debug [instance_properties/setter]: isOpaque */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFormCell */


