// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [FormCell] class.
type IFormCell interface {
	IActionCell
	// properties:
	TitleWidth() float64 /* primitive/slice/pointer. */
	SetTitleWidth(value float64 /* primitive/slice/pointer. */)
	AttributedTitle() objc.IObject /* cross-framework: AttributedString */
	SetAttributedTitle(value objc.IObject /* cross-framework: AttributedString */)
	IsOpaque() bool /* primitive/slice/pointer. */
	SetIsOpaque(value bool /* primitive/slice/pointer. */)
	PlaceholderAttributedString() objc.IObject /* cross-framework: AttributedString */
	SetPlaceholderAttributedString(value objc.IObject /* cross-framework: AttributedString */)
	PlaceholderString() string /* primitive/slice/pointer. */
	SetPlaceholderString(value string /* primitive/slice/pointer. */)
	PreferredTextFieldWidth() float64 /* primitive/slice/pointer. */
	SetPreferredTextFieldWidth(value float64 /* primitive/slice/pointer. */)
	Title() string /* primitive/slice/pointer. */
	SetTitle(value string /* primitive/slice/pointer. */)
	TitleAlignment() TextAlignment
	SetTitleAlignment(value TextAlignment)
	TitleBaseWritingDirection() WritingDirection
	SetTitleBaseWritingDirection(value WritingDirection)
	TitleFont() IFont
	SetTitleFont(value IFont)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (fc _FormCellClass) Alloc() FormCell {
	rv := objc.Send[FormCell](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The width of the title field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleWidth
func (f_ FormCell) TitleWidth() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](f_.ID, objc.Sel("titleWidth"))
	return rv
}


// The width of the title field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/titleWidth
func (f_ FormCell) SetTitleWidth(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleWidth:"), value)
}


// The title of the cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/attributedtitle
func (f_ FormCell) AttributedTitle() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[AttributedString](f_.ID, objc.Sel("attributedTitle"))
	return rv
}


// The title of the cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/attributedtitle
func (f_ FormCell) SetAttributedTitle(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAttributedTitle:"), value)
}


// A Boolean value indicating whether the title is empty and an opaque bezel is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/isopaque
func (f_ FormCell) IsOpaque() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isOpaque"))
	return rv
}


// A Boolean value indicating whether the title is empty and an opaque bezel is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/isopaque
func (f_ FormCell) SetIsOpaque(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsOpaque:"), value)
}


// The cell’s attributed placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/placeholderattributedstring
func (f_ FormCell) PlaceholderAttributedString() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[AttributedString](f_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// The cell’s attributed placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/placeholderattributedstring
func (f_ FormCell) SetPlaceholderAttributedString(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}


// The cell’s plain text placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/placeholderstring
func (f_ FormCell) PlaceholderString() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](f_.ID, objc.Sel("placeholderString"))
	return rv
}


// The cell’s plain text placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/placeholderstring
func (f_ FormCell) SetPlaceholderString(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}


// The preferred text field width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/preferredtextfieldwidth
func (f_ FormCell) PreferredTextFieldWidth() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](f_.ID, objc.Sel("preferredTextFieldWidth"))
	return rv
}


// The preferred text field width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/preferredtextfieldwidth
func (f_ FormCell) SetPreferredTextFieldWidth(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPreferredTextFieldWidth:"), value)
}


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/title
func (f_ FormCell) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](f_.ID, objc.Sel("title"))
	return rv
}


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/title
func (f_ FormCell) SetTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The alignment of the title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlealignment
func (f_ FormCell) TitleAlignment() TextAlignment {
	rv := objc.Send[TextAlignment](f_.ID, objc.Sel("titleAlignment"))
	return rv
}


// The alignment of the title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlealignment
func (f_ FormCell) SetTitleAlignment(value TextAlignment) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleAlignment:"), value)
}


// The default writing direction used to render the form cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlebasewritingdirection
func (f_ FormCell) TitleBaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](f_.ID, objc.Sel("titleBaseWritingDirection"))
	return rv
}


// The default writing direction used to render the form cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlebasewritingdirection
func (f_ FormCell) SetTitleBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleBaseWritingDirection:"), value)
}


// The font used to draw cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlefont
func (f_ FormCell) TitleFont() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("titleFont"))
	return rv
}


// The font used to draw cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlefont
func (f_ FormCell) SetTitleFont(value IFont) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleFont:"), value)
}



