// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextElement */


/* debug [class_header]: Header for NSTextElement */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextElement */
// An interface definition for the [TextElement] class.
type ITextElement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextElement */
	// properties:
	ChildElements() []TextElement
	ElementRange() ITextRange
	SetElementRange(value ITextRange)
	IsRepresentedElement() bool
	ParentElement() ITextElement
	TextContentManager() ITextContentManager
	SetTextContentManager(value ITextContentManager)
	Parent() ITextElement
	SetParent(value ITextElement)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextElement */
// Alloc allocates a new instance without initialization.
func (tc _TextElementClass) Alloc() TextElement {
	rv := objc.Send[TextElement](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextElement */
// An abstract base class that represents the smallest units of text layout such as paragraphs or attachments.


// An abstract base class that represents the smallest units of text layout such as paragraphs or attachments.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextElement */

// Creates a new text element with the content manager you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement/init(textContentManager:)
func NewTextElementWithTextContentManager(textContentManager ITextContentManager) TextElement {
	instance := getTextElementClass().Alloc()
	rv := objc.Send[TextElement](instance.ID, objc.Sel("initWithTextContentManager:"), textContentManager)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextElementWithTextContentManager */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextElement */

// An array of zero or more child text elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement/childElements
func (t_ TextElement) ChildElements() []TextElement {
	rv := objc.Send[[]TextElement](t_.ID, objc.Sel("childElements"))
	return rv
}/* debug [instance_properties/getter]: childElements */


// A range value that represents the range of the element inside the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement/elementRange
func (t_ TextElement) ElementRange() ITextRange {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("elementRange"))
	return rv
}/* debug [instance_properties/getter]: elementRange */


// A range value that represents the range of the element inside the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement/elementRange
func (t_ TextElement) SetElementRange(value ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setElementRange:"), value)
}/* debug [instance_properties/setter]: elementRange */


// A Boolean value that indicates whether this element is in the text layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement/isRepresentedElement
func (t_ TextElement) IsRepresentedElement() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRepresentedElement"))
	return rv
}/* debug [instance_properties/getter]: isRepresentedElement */


// A value that represents the parent element if this text element is a child of an enclosing element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement/parent
func (t_ TextElement) ParentElement() ITextElement {
	rv := objc.Send[TextElement](t_.ID, objc.Sel("parentElement"))
	return rv
}/* debug [instance_properties/getter]: parentElement */


// The value that represents the current content manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement/textContentManager
func (t_ TextElement) TextContentManager() ITextContentManager {
	rv := objc.Send[TextContentManager](t_.ID, objc.Sel("textContentManager"))
	return rv
}/* debug [instance_properties/getter]: textContentManager */


// The value that represents the current content manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement/textContentManager
func (t_ TextElement) SetTextContentManager(value ITextContentManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContentManager:"), value)
}/* debug [instance_properties/setter]: textContentManager */


// A value that represents the parent element if this text element is a child of an enclosing element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/parent
func (t_ TextElement) Parent() ITextElement {
	rv := objc.Send[TextElement](t_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// A value that represents the parent element if this text element is a child of an enclosing element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextelement/parent
func (t_ TextElement) SetParent(value ITextElement) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextElement */


