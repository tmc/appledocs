// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextListElement */


/* debug [class_header]: Header for NSTextListElement */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextListElement */
// An interface definition for the [TextListElement] class.
type ITextListElement interface {
	ITextParagraph
	
/* debug [class_interface_properties]: Properties for TextListElement */
	// properties:
	AttributedString() foundation.AttributedString
	ChildElements() []TextListElement
	Contents() foundation.AttributedString
	MarkerAttributes() foundation.IDictionary
	ParentElement() ITextListElement
	TextList() ITextList
	Parent() ITextListElement
	SetParent(value ITextListElement)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextListElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextListElement */
// Alloc allocates a new instance without initialization.
func (tc _TextListElementClass) Alloc() TextListElement {
	rv := objc.Send[TextListElement](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextListElement */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextListElement */

// Creates a text list element with the list elements and nesting level you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(children:textList:nestingLevel:)
func NewTextListElementWithChildElementsTextListNestingLevel(children []TextListElement, textList ITextList, nestingLevel int) TextListElement {
	rv := objc.Send[TextListElement](objc.ID(getTextListElementClass().class), objc.Sel("textListElementWithChildElements:textList:nestingLevel:"), children, textList, nestingLevel)
	return rv
}/* debug [class_init_methods/constructor]: NewTextListElementWithChildElementsTextListNestingLevel */


// Creates a text list element with the list elements, nesting level, and marker attributes you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(contents:markerAttributes:textList:children:)
func NewTextListElementWithContentsMarkerAttributesTextListChildElements(contents foundation.AttributedString, markerAttributes foundation.IDictionary, textList ITextList, children []TextListElement) TextListElement {
	rv := objc.Send[TextListElement](objc.ID(getTextListElementClass().class), objc.Sel("textListElementWithContents:markerAttributes:textList:childElements:"), contents, markerAttributes, textList, children)
	return rv
}/* debug [class_init_methods/constructor]: NewTextListElementWithContentsMarkerAttributesTextListChildElements */


// Creates a text list element with the parent, list elements, nesting level, and marker attributes you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(parent:textList:contents:markerAttributes:children:)
func NewTextListElementWithParentElementTextListContentsMarkerAttributesChildElements(parent ITextListElement, textList ITextList, contents foundation.AttributedString, markerAttributes foundation.IDictionary, children []TextListElement) TextListElement {
	instance := getTextListElementClass().Alloc()
	rv := objc.Send[TextListElement](instance.ID, objc.Sel("initWithParentElement:textList:contents:markerAttributes:childElements:"), parent, textList, contents, markerAttributes, children)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextListElementWithParentElementTextListContentsMarkerAttributesChildElements */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextListElement */

// Creates a text list element with the list elements and nesting level you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(children:textList:nestingLevel:)
func (tc _TextListElementClass) TextListElementWithChildElementsTextListNestingLevel(children []TextListElement, textList ITextList, nestingLevel int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("textListElementWithChildElements:textList:nestingLevel:"), children, textList, nestingLevel)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TextListElementWithChildElementsTextListNestingLevel) */


// Creates a text list element with the list elements, nesting level, and marker attributes you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(contents:markerAttributes:textList:children:)
func (tc _TextListElementClass) TextListElementWithContentsMarkerAttributesTextListChildElements(contents foundation.AttributedString, markerAttributes foundation.IDictionary, textList ITextList, children []TextListElement) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("textListElementWithContents:markerAttributes:textList:childElements:"), contents, markerAttributes, textList, children)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TextListElementWithContentsMarkerAttributesTextListChildElements) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextListElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextListElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextListElement */

// An attributed string that represents the string the framework displays for this element taking into account markers and the indentation level of the list element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/attributedString
func (t_ TextListElement) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}/* debug [instance_properties/getter]: attributedString */


// An array that contains child text elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/childElements
func (t_ TextListElement) ChildElements() []TextListElement {
	rv := objc.Send[[]TextListElement](t_.ID, objc.Sel("childElements"))
	return rv
}/* debug [instance_properties/getter]: childElements */


// The text list element contents without markers and formatting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/contents
func (t_ TextListElement) Contents() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// A dictionary of attributed string keys and IDs that represent the list’s marker attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/markerAttributes
func (t_ TextListElement) MarkerAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("markerAttributes"))
	return rv
}/* debug [instance_properties/getter]: markerAttributes */


// A text list element that refers to the enclosing text list element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/parent
func (t_ TextListElement) ParentElement() ITextListElement {
	rv := objc.Send[TextListElement](t_.ID, objc.Sel("parentElement"))
	return rv
}/* debug [instance_properties/getter]: parentElement */


// The value that represents the text list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement/textList
func (t_ TextListElement) TextList() ITextList {
	rv := objc.Send[TextList](t_.ID, objc.Sel("textList"))
	return rv
}/* debug [instance_properties/getter]: textList */


// A text list element that refers to the enclosing text list element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/parent
func (t_ TextListElement) Parent() ITextListElement {
	rv := objc.Send[TextListElement](t_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// A text list element that refers to the enclosing text list element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistelement/parent
func (t_ TextListElement) SetParent(value ITextListElement) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextListElement */


