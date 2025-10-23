// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	Opaque() bool
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	IsOpaque() bool
	SetIsOpaque(value bool)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
	PreferredTextFieldWidth() float64
	SetPreferredTextFieldWidth(value float64)
	Title() string
	SetTitle(value string)
	TitleAlignment() TextAlignment
	SetTitleAlignment(value ITextAlignment)
	TitleBaseWritingDirection() WritingDirection
	SetTitleBaseWritingDirection(value IWritingDirection)
	TitleFont() NSFont
	SetTitleFont(value IFont)
	TitleWidth() float64
	SetTitleWidth(value float64)
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/init(coder:)
func NewFormCellWithCoder(coder foundation.ICoder) FormCell {
	instance := getFormCellClass().Alloc()
	rv := objc.Send[FormCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// A Boolean value indicating whether the title is empty and an opaque bezel is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell/isOpaque
func (f_ FormCell) Opaque() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("opaque"))
	return rv
}


// The title of the cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/attributedtitle
func (f_ FormCell) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](f_.ID, objc.Sel("attributedTitle"))
	return rv
}


// The title of the cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/attributedtitle
func (f_ FormCell) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAttributedTitle:"), value)
}


// A Boolean value indicating whether the title is empty and an opaque bezel is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/isopaque
func (f_ FormCell) IsOpaque() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isOpaque"))
	return rv
}


// A Boolean value indicating whether the title is empty and an opaque bezel is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/isopaque
func (f_ FormCell) SetIsOpaque(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsOpaque:"), value)
}


// The cell’s attributed placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/placeholderattributedstring
func (f_ FormCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](f_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// The cell’s attributed placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/placeholderattributedstring
func (f_ FormCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}


// The cell’s plain text placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/placeholderstring
func (f_ FormCell) PlaceholderString() string {
	rv := objc.Send[string](f_.ID, objc.Sel("placeholderString"))
	return rv
}


// The cell’s plain text placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/placeholderstring
func (f_ FormCell) SetPlaceholderString(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}


// The preferred text field width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/preferredtextfieldwidth
func (f_ FormCell) PreferredTextFieldWidth() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("preferredTextFieldWidth"))
	return rv
}


// The preferred text field width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/preferredtextfieldwidth
func (f_ FormCell) SetPreferredTextFieldWidth(value float64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPreferredTextFieldWidth:"), value)
}


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/title
func (f_ FormCell) Title() string {
	rv := objc.Send[string](f_.ID, objc.Sel("title"))
	return rv
}


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/title
func (f_ FormCell) SetTitle(value string) {
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
func (f_ FormCell) SetTitleAlignment(value ITextAlignment) {
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
func (f_ FormCell) SetTitleBaseWritingDirection(value IWritingDirection) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleBaseWritingDirection:"), value)
}


// The font used to draw cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlefont
func (f_ FormCell) TitleFont() NSFont {
	rv := objc.Send[NSFont](f_.ID, objc.Sel("titleFont"))
	return rv
}


// The font used to draw cell’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlefont
func (f_ FormCell) SetTitleFont(value IFont) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleFont:"), value)
}


// The width of the title field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlewidth
func (f_ FormCell) TitleWidth() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("titleWidth"))
	return rv
}


// The width of the title field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/titlewidth
func (f_ FormCell) SetTitleWidth(value float64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTitleWidth:"), value)
}


