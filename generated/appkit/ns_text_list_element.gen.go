// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TextListElement] class.
var (
	TextListElementClass     _TextListElementClass
	TextListElementClassOnce sync.Once
)

func getTextListElementClass() _TextListElementClass {
	TextListElementClassOnce.Do(func() {
		TextListElementClass = _TextListElementClass{objc.GetClass("NSTextListElement")}
	})
	return TextListElementClass
}

type _TextListElementClass struct {
	class objc.Class
}

// An interface definition for the [TextListElement] class.
type ITextListElement interface {
	ITextParagraph
}

// A class that represents a text list node.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement
type TextListElement struct {
	TextParagraph
}

// TextListElementFrom constructs a [TextListElement] from an unsafe.Pointer.
//
// A class that represents a text list node.
func TextListElementFrom(ptr unsafe.Pointer) TextListElement {
	return TextListElement{
		TextParagraph: TextParagraphFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextListElementClass) Alloc() TextListElement {
	rv := objc.Send[TextListElement](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextListElementClass) New() TextListElement {
	rv := objc.Send[TextListElement](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextListElement) Init() TextListElement {
	rv := objc.Send[TextListElement](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextListElement) Autorelease() TextListElement {
	rv := objc.Send[TextListElement](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextListElement creates a new TextListElement instance.
func NewTextListElement() TextListElement {
	return getTextListElementClass().New()
}


// An attributed string that represents the string the framework displays for this element taking into account markers and the indentation level of the list element.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/attributedstring
func (t_ TextListElement) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}


// SetAttributedString sets the value of the attributedString property.
// An attributed string that represents the string the framework displays for this element taking into account markers and the indentation level of the list element.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/attributedstring
func (t_ TextListElement) SetAttributedString(value foundation.IAttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributedString:"), value)
}

// An array that contains child text elements.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/childelements
func (t_ TextListElement) ChildElements() NSTextListElement {
	rv := objc.Send[NSTextListElement](t_.ID, objc.Sel("childElements"))
	return rv
}


// SetChildElements sets the value of the childElements property.
// An array that contains child text elements.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/childelements
func (t_ TextListElement) SetChildElements(value ITextListElement) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChildElements:"), value)
}

// The text list element contents without markers and formatting.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/contents
func (t_ TextListElement) Contents() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("contents"))
	return rv
}


// SetContents sets the value of the contents property.
// The text list element contents without markers and formatting.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/contents
func (t_ TextListElement) SetContents(value foundation.IAttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContents:"), value)
}

// A dictionary of attributed string keys and IDs that represent the list’s marker attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/markerattributes
func (t_ TextListElement) MarkerAttributes() coreml.Key {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("markerAttributes"))
	return rv
}


// SetMarkerAttributes sets the value of the markerAttributes property.
// A dictionary of attributed string keys and IDs that represent the list’s marker attributes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/markerattributes
func (t_ TextListElement) SetMarkerAttributes(value coreml.IKey) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMarkerAttributes:"), value)
}

// A text list element that refers to the enclosing text list element.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/parent
func (t_ TextListElement) Parent() NSTextListElement {
	rv := objc.Send[NSTextListElement](t_.ID, objc.Sel("parent"))
	return rv
}


// SetParent sets the value of the parent property.
// A text list element that refers to the enclosing text list element.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/parent
func (t_ TextListElement) SetParent(value ITextListElement) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParent:"), value)
}

// The value that represents the text list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/textlist
func (t_ TextListElement) TextList() NSTextList {
	rv := objc.Send[NSTextList](t_.ID, objc.Sel("textList"))
	return rv
}


// SetTextList sets the value of the textList property.
// The value that represents the text list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/textlist
func (t_ TextListElement) SetTextList(value ITextList) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextList:"), value)
}



