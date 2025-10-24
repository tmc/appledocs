// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextLineFragment */


/* debug [class_header]: Header for NSTextLineFragment */
// The class instance for the [TextLineFragment] class.
var (
	TextLineFragmentClass     _TextLineFragmentClass
	TextLineFragmentClassOnce sync.Once
)

func getTextLineFragmentClass() _TextLineFragmentClass {
	TextLineFragmentClassOnce.Do(func() {
		TextLineFragmentClass = _TextLineFragmentClass{objc.GetClass("NSTextLineFragment")}
	})
	return TextLineFragmentClass
}

type _TextLineFragmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextLineFragment */
// An interface definition for the [TextLineFragment] class.
type ITextLineFragment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextLineFragment */
	// properties:
	AttributedString() foundation.AttributedString
	CharacterRange() corefoundation.Range
	GlyphOrigin() corefoundation.CGPoint
	TypographicBounds() corefoundation.CGRect
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextLineFragment */
	// methods:
	CharacterIndexForPoint(point corefoundation.CGPoint) int
	DrawAtPointInContext(point corefoundation.CGPoint, context ContextRef /* not a class type */)
	FractionOfDistanceThroughGlyphForPoint(point corefoundation.CGPoint) float64
	LocationForCharacterAtIndex(index int) corefoundation.CGPoint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextLineFragment */
// Alloc allocates a new instance without initialization.
func (tc _TextLineFragmentClass) Alloc() TextLineFragment {
	rv := objc.Send[TextLineFragment](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextLineFragmentClass) New() TextLineFragment {
	rv := objc.Send[TextLineFragment](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextLineFragment) Init() TextLineFragment {
	rv := objc.Send[TextLineFragment](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextLineFragment) Autorelease() TextLineFragment {
	rv := objc.Send[TextLineFragment](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextLineFragment creates a new TextLineFragment instance.
func NewTextLineFragment() TextLineFragment {
	return getTextLineFragmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextLineFragment */
// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment.


// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment
type TextLineFragment struct {
	objectivec.Object
}

// TextLineFragmentFrom constructs a [TextLineFragment] from an unsafe.Pointer.
//
// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment.
func TextLineFragmentFrom(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextLineFragment */

// Creates a new line fragment from the attributed string for the range of characters you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/init(attributedString:range:)
func NewTextLineFragmentWithAttributedStringRange(attributedString foundation.AttributedString, range_ corefoundation.Range) TextLineFragment {
	instance := getTextLineFragmentClass().Alloc()
	rv := objc.Send[TextLineFragment](instance.ID, objc.Sel("initWithAttributedString:range:"), attributedString, range_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextLineFragmentWithAttributedStringRange */


// Creates a new line fragment with from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/init(coder:)
func NewTextLineFragmentWithCoder(aDecoder foundation.Coder) TextLineFragment {
	instance := getTextLineFragmentClass().Alloc()
	rv := objc.Send[TextLineFragment](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextLineFragmentWithCoder */


// Creates a new line fragment using the string, attributes, and range you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/init(string:attributes:range:)
func NewTextLineFragmentWithStringAttributesRange(string_ objc.IObject /* cross-framework: NSString */, attributes foundation.IDictionary, range_ corefoundation.Range) TextLineFragment {
	instance := getTextLineFragmentClass().Alloc()
	rv := objc.Send[TextLineFragment](instance.ID, objc.Sel("initWithString:attributes:range:"), string_, attributes, range_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextLineFragmentWithStringAttributesRange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextLineFragment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextLineFragment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextLineFragment */

// Returns character index for a point inside the line fragment coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/characterIndex(for:)
func (t_ TextLineFragment) CharacterIndexForPoint(point corefoundation.CGPoint) int {
	rv := objc.Send[int](t_.ID, objc.Sel("characterIndexForPoint:"), point)
	return rv
}/* debug [instance_methods/method]: CharacterIndexForPoint */


// Renders the line fragment contents at the rendering origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/draw(at:in:)
func (t_ TextLineFragment) DrawAtPointInContext(point corefoundation.CGPoint, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawAtPoint:inContext:"), point, context)
}/* debug [instance_methods/method]: DrawAtPointInContext */


// Returns character index for a point inside the line fragment coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/fractionOfDistanceThroughGlyph(for:)
func (t_ TextLineFragment) FractionOfDistanceThroughGlyphForPoint(point corefoundation.CGPoint) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("fractionOfDistanceThroughGlyphForPoint:"), point)
	return rv
}/* debug [instance_methods/method]: FractionOfDistanceThroughGlyphForPoint */


// Returns the location of the character at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/locationForCharacter(at:)
func (t_ TextLineFragment) LocationForCharacterAtIndex(index int) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t_.ID, objc.Sel("locationForCharacterAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: LocationForCharacterAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextLineFragment */

// The source attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/attributedString
func (t_ TextLineFragment) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}/* debug [instance_properties/getter]: attributedString */


// The string range for the source attributed string that corresponds to this line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/characterRange
func (t_ TextLineFragment) CharacterRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("characterRange"))
	return rv
}/* debug [instance_properties/getter]: characterRange */


// Rendering origin for the left-most glyph in the line fragment coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/glyphOrigin
func (t_ TextLineFragment) GlyphOrigin() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t_.ID, objc.Sel("glyphOrigin"))
	return rv
}/* debug [instance_properties/getter]: glyphOrigin */


// The typographic bounds that specifies the dimensions of the line fragment for laying out line fragments to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/typographicBounds
func (t_ TextLineFragment) TypographicBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("typographicBounds"))
	return rv
}/* debug [instance_properties/getter]: typographicBounds */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextLineFragment */


