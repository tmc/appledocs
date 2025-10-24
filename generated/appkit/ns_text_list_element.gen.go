// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	AttributedString() foundation.AttributedString
	ChildElements() []TextListElement
	Contents() foundation.AttributedString
	MarkerAttributes() foundation.IDictionary
	ParentElement() ITextListElement
	TextList() ITextList
	Parent() ITextListElement
	SetParent(value ITextListElement)
	// methods:
}

// A class that represents a text list node.


// A class that represents a text list node.
//
// [Full Topic]
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



// Creates a text list element with the list elements and nesting level you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(children:textList:nestingLevel:)
func NewTextListElementWithChildElementsTextListNestingLevel(children []TextListElement, textList ITextList, nestingLevel int) TextListElement {
	rv := objc.Send[TextListElement](objc.ID(getTextListElementClass().class), objc.Sel("textListElementWithChildElements:textList:nestingLevel:"), children, textList, nestingLevel)
	return rv
}


// Creates a text list element with the list elements, nesting level, and marker attributes you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(contents:markerAttributes:textList:children:)
func NewTextListElementWithContentsMarkerAttributesTextListChildElements(contents foundation.AttributedString, markerAttributes foundation.IDictionary, textList ITextList, children []TextListElement) TextListElement {
	rv := objc.Send[TextListElement](objc.ID(getTextListElementClass().class), objc.Sel("textListElementWithContents:markerAttributes:textList:childElements:"), contents, markerAttributes, textList, children)
	return rv
}


// Creates a text list element with the parent, list elements, nesting level, and marker attributes you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(parent:textList:contents:markerAttributes:children:)
func NewTextListElementWithParentElementTextListContentsMarkerAttributesChildElements(parent ITextListElement, textList ITextList, contents foundation.AttributedString, markerAttributes foundation.IDictionary, children []TextListElement) TextListElement {
	instance := getTextListElementClass().Alloc()
	rv := objc.Send[TextListElement](instance.ID, objc.Sel("initWithParentElement:textList:contents:markerAttributes:childElements:"), parent, textList, contents, markerAttributes, children)
	rv.Autorelease()
	return rv
}



// Creates a text list element with the list elements and nesting level you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(children:textList:nestingLevel:)
func (tc _TextListElementClass) TextListElementWithChildElementsTextListNestingLevel(children []TextListElement, textList ITextList, nestingLevel int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("textListElementWithChildElements:textList:nestingLevel:"), children, textList, nestingLevel)
	return rv
}


// Creates a text list element with the list elements, nesting level, and marker attributes you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(contents:markerAttributes:textList:children:)
func (tc _TextListElementClass) TextListElementWithContentsMarkerAttributesTextListChildElements(contents foundation.AttributedString, markerAttributes foundation.IDictionary, textList ITextList, children []TextListElement) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("textListElementWithContents:markerAttributes:textList:childElements:"), contents, markerAttributes, textList, children)
	return rv
}


// An attributed string that represents the string the framework displays for this element taking into account markers and the indentation level of the list element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/attributedString
func (t_ TextListElement) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}


// An array that contains child text elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/childElements
func (t_ TextListElement) ChildElements() []TextListElement {
	rv := objc.Send[[]TextListElement](t_.ID, objc.Sel("childElements"))
	return rv
}


// The text list element contents without markers and formatting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/contents
func (t_ TextListElement) Contents() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("contents"))
	return rv
}


// A dictionary of attributed string keys and IDs that represent the list’s marker attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/markerAttributes
func (t_ TextListElement) MarkerAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("markerAttributes"))
	return rv
}


// A text list element that refers to the enclosing text list element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/parent
func (t_ TextListElement) ParentElement() ITextListElement {
	rv := objc.Send[TextListElement](t_.ID, objc.Sel("parentElement"))
	return rv
}


// The value that represents the text list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/textList
func (t_ TextListElement) TextList() ITextList {
	rv := objc.Send[TextList](t_.ID, objc.Sel("textList"))
	return rv
}


// A text list element that refers to the enclosing text list element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/parent
func (t_ TextListElement) Parent() ITextListElement {
	rv := objc.Send[TextListElement](t_.ID, objc.Sel("parent"))
	return rv
}


// A text list element that refers to the enclosing text list element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/parent
func (t_ TextListElement) SetParent(value ITextListElement) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParent:"), value)
}


