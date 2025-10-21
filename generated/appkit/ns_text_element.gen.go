// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextElement] class.
var (
	TextElementClass     _TextElementClass
	TextElementClassOnce sync.Once
)

func getTextElementClass() _TextElementClass {
	TextElementClassOnce.Do(func() {
		TextElementClass = _TextElementClass{objc.GetClass("NSTextElement")}
	})
	return TextElementClass
}

type _TextElementClass struct {
	class objc.Class
}

// An interface definition for the [TextElement] class.
type ITextElement interface {
	objectivec.IObject
}

// An abstract base class that represents the smallest units of text layout such as paragraphs or attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement
type TextElement struct {
	objectivec.Object
}

// TextElementFrom constructs a [TextElement] from an unsafe.Pointer.
//
// An abstract base class that represents the smallest units of text layout such as paragraphs or attachments.
func TextElementFrom(ptr unsafe.Pointer) TextElement {
	return TextElement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextElementClass) Alloc() TextElement {
	rv := objc.Send[TextElement](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextElementClass) New() TextElement {
	rv := objc.Send[TextElement](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextElement) Init() TextElement {
	rv := objc.Send[TextElement](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextElement) Autorelease() TextElement {
	rv := objc.Send[TextElement](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextElement creates a new TextElement instance.
func NewTextElement() TextElement {
	return getTextElementClass().New()
}


// An array of zero or more child text elements.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/childelements
func (t_ TextElement) ChildElements() NSTextElement {
	rv := objc.Send[NSTextElement](t_.ID, objc.Sel("childElements"))
	return rv
}


// SetChildElements sets the value of the childElements property.
// An array of zero or more child text elements.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/childelements
func (t_ TextElement) SetChildElements(value ITextElement) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChildElements:"), value)
}

// A range value that represents the range of the element inside the document.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/elementrange
func (t_ TextElement) ElementRange() NSTextRange {
	rv := objc.Send[NSTextRange](t_.ID, objc.Sel("elementRange"))
	return rv
}


// SetElementRange sets the value of the elementRange property.
// A range value that represents the range of the element inside the document.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/elementrange
func (t_ TextElement) SetElementRange(value ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setElementRange:"), value)
}

// A Boolean value that indicates whether this element is in the text layout.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/isrepresentedelement
func (t_ TextElement) IsRepresentedElement() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRepresentedElement"))
	return rv
}


// SetIsRepresentedElement sets the value of the isRepresentedElement property.
// A Boolean value that indicates whether this element is in the text layout.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/isrepresentedelement
func (t_ TextElement) SetIsRepresentedElement(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRepresentedElement:"), value)
}

// A value that represents the parent element if this text element is a child of an enclosing element.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/parent
func (t_ TextElement) Parent() NSTextElement {
	rv := objc.Send[NSTextElement](t_.ID, objc.Sel("parent"))
	return rv
}


// SetParent sets the value of the parent property.
// A value that represents the parent element if this text element is a child of an enclosing element.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/parent
func (t_ TextElement) SetParent(value ITextElement) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParent:"), value)
}

// The value that represents the current content manager.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/textcontentmanager
func (t_ TextElement) TextContentManager() NSTextContentManager {
	rv := objc.Send[NSTextContentManager](t_.ID, objc.Sel("textContentManager"))
	return rv
}


// SetTextContentManager sets the value of the textContentManager property.
// The value that represents the current content manager.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/textcontentmanager
func (t_ TextElement) SetTextContentManager(value ITextContentManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContentManager:"), value)
}



