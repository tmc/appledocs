// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextParagraph */


/* debug [class_header]: Header for NSTextParagraph */
// The class instance for the [TextParagraph] class.
var (
	TextParagraphClass     _TextParagraphClass
	TextParagraphClassOnce sync.Once
)

func getTextParagraphClass() _TextParagraphClass {
	TextParagraphClassOnce.Do(func() {
		TextParagraphClass = _TextParagraphClass{objc.GetClass("NSTextParagraph")}
	})
	return TextParagraphClass
}

type _TextParagraphClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextParagraph */
// An interface definition for the [TextParagraph] class.
type ITextParagraph interface {
	ITextElement
	
/* debug [class_interface_properties]: Properties for TextParagraph */
	// properties:
	AttributedString() foundation.AttributedString
	ParagraphContentRange() ITextRange
	ParagraphSeparatorRange() ITextRange
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextParagraph */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextParagraph */
// Alloc allocates a new instance without initialization.
func (tc _TextParagraphClass) Alloc() TextParagraph {
	rv := objc.Send[TextParagraph](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextParagraphClass) New() TextParagraph {
	rv := objc.Send[TextParagraph](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextParagraph) Init() TextParagraph {
	rv := objc.Send[TextParagraph](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextParagraph) Autorelease() TextParagraph {
	rv := objc.Send[TextParagraph](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextParagraph creates a new TextParagraph instance.
func NewTextParagraph() TextParagraph {
	return getTextParagraphClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextParagraph */
// A class that represents a single paragraph backed by an attributed string as the contents.


// A class that represents a single paragraph backed by an attributed string as the contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextParagraph
type TextParagraph struct {
	TextElement
}

// TextParagraphFrom constructs a [TextParagraph] from an unsafe.Pointer.
//
// A class that represents a single paragraph backed by an attributed string as the contents.
func TextParagraphFrom(ptr unsafe.Pointer) TextParagraph {
	return TextParagraph{
		TextElement: TextElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextParagraph */

// Creates a new paragraph with the attributed string you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextParagraph/init(attributedString:)
func NewTextParagraphWithAttributedString(attributedString foundation.AttributedString) TextParagraph {
	instance := getTextParagraphClass().Alloc()
	rv := objc.Send[TextParagraph](instance.ID, objc.Sel("initWithAttributedString:"), attributedString)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextParagraphWithAttributedString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextParagraph */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextParagraph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextParagraph */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextParagraph */

// Returns the source attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextParagraph/attributedString
func (t_ TextParagraph) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}/* debug [instance_properties/getter]: attributedString */


// Returns the range of the paragraph in the containing text’s attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextParagraph/paragraphContentRange
func (t_ TextParagraph) ParagraphContentRange() ITextRange {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("paragraphContentRange"))
	return rv
}/* debug [instance_properties/getter]: paragraphContentRange */


// Returns the range of the paragraph separator in the containing text’s attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextParagraph/paragraphSeparatorRange
func (t_ TextParagraph) ParagraphSeparatorRange() ITextRange {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("paragraphSeparatorRange"))
	return rv
}/* debug [instance_properties/getter]: paragraphSeparatorRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextParagraph */


