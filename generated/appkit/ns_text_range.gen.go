// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextRange */


/* debug [class_header]: Header for NSTextRange */
// The class instance for the [TextRange] class.
var (
	TextRangeClass     _TextRangeClass
	TextRangeClassOnce sync.Once
)

func getTextRangeClass() _TextRangeClass {
	TextRangeClassOnce.Do(func() {
		TextRangeClass = _TextRangeClass{objc.GetClass("NSTextRange")}
	})
	return TextRangeClass
}

type _TextRangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextRange */
// An interface definition for the [TextRange] class.
type ITextRange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextRange */
	// properties:
	EndLocation() unsafe.Pointer
	Empty() bool
	Location() unsafe.Pointer
	IsEmpty() bool
	SetIsEmpty(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextRange */
	// methods:
	ContainsRange(textRange ITextRange) bool
	ContainsLocation(location unsafe.Pointer) bool
	TextRangeByIntersectingWithTextRange(textRange ITextRange) objectivec.IObject
	IntersectsWithTextRange(textRange ITextRange) bool
	IsEqualToTextRange(textRange ITextRange) bool
	TextRangeByFormingUnionWithTextRange(textRange ITextRange) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextRange */
// Alloc allocates a new instance without initialization.
func (tc _TextRangeClass) Alloc() TextRange {
	rv := objc.Send[TextRange](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextRangeClass) New() TextRange {
	rv := objc.Send[TextRange](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextRange) Init() TextRange {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextRange) Autorelease() TextRange {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextRange creates a new TextRange instance.
func NewTextRange() TextRange {
	return getTextRangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextRange */
// A class that represents a contiguous range between two locations inside document contents.
//
// An consists of the starting and terminating locations. There the two basic properties: and , respectively. The terminating , , is directly following the last location in the range. For example, a location contains a range if is .


// A class that represents a contiguous range between two locations inside document contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange
type TextRange struct {
	objectivec.Object
}

// TextRangeFrom constructs a [TextRange] from an unsafe.Pointer.
//
// A class that represents a contiguous range between two locations inside document contents.
func TextRangeFrom(ptr unsafe.Pointer) TextRange {
	return TextRange{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextRange */

// Creates a new text range at the location you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/init(location:)
func NewTextRangeWithLocation(location unsafe.Pointer) TextRange {
	instance := getTextRangeClass().Alloc()
	rv := objc.Send[TextRange](instance.ID, objc.Sel("initWithLocation:"), location)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextRangeWithLocation */


// Creates a new text range with the starting and ending locations you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/init(location:end:)
func NewTextRangeWithLocationEndLocation(location unsafe.Pointer, endLocation unsafe.Pointer) TextRange {
	instance := getTextRangeClass().Alloc()
	rv := objc.Send[TextRange](instance.ID, objc.Sel("initWithLocation:endLocation:"), location, endLocation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextRangeWithLocationEndLocation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextRange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextRange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextRange */

// Determines if the text range you specify is in the current text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/contains(_:)-5j4y2
func (t_ TextRange) ContainsRange(textRange ITextRange) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("containsRange:"), textRange)
	return rv
}/* debug [instance_methods/method]: ContainsRange */


// Determines if the text location you specify is in the current text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/contains(_:)-7hvi0
func (t_ TextRange) ContainsLocation(location unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("containsLocation:"), location)
	return rv
}/* debug [instance_methods/method]: ContainsLocation */


// Returns the range, if any, where two text ranges intersect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/intersection(_:)
func (t_ TextRange) TextRangeByIntersectingWithTextRange(textRange ITextRange) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("textRangeByIntersectingWithTextRange:"), textRange)
	return rv
}/* debug [instance_methods/method]: TextRangeByIntersectingWithTextRange */


// Determines if two ranges intersect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/intersects(_:)
func (t_ TextRange) IntersectsWithTextRange(textRange ITextRange) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("intersectsWithTextRange:"), textRange)
	return rv
}/* debug [instance_methods/method]: IntersectsWithTextRange */


// Compares two text ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/isEqual(to:)
func (t_ TextRange) IsEqualToTextRange(textRange ITextRange) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEqualToTextRange:"), textRange)
	return rv
}/* debug [instance_methods/method]: IsEqualToTextRange */


// Returns a new text range by forming the union with the text range you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/union(_:)
func (t_ TextRange) TextRangeByFormingUnionWithTextRange(textRange ITextRange) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("textRangeByFormingUnionWithTextRange:"), textRange)
	return rv
}/* debug [instance_methods/method]: TextRangeByFormingUnionWithTextRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextRange */

// The ending location of the text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/endLocation
func (t_ TextRange) EndLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("endLocation"))
	return rv
}/* debug [instance_properties/getter]: endLocation */


// Returns whether the text range is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/isEmpty
func (t_ TextRange) Empty() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("empty"))
	return rv
}/* debug [instance_properties/getter]: empty */


// The starting location of the text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/location
func (t_ TextRange) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// Returns whether the text range is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextrange/isempty
func (t_ TextRange) IsEmpty() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEmpty"))
	return rv
}/* debug [instance_properties/getter]: isEmpty */


// Returns whether the text range is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextrange/isempty
func (t_ TextRange) SetIsEmpty(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEmpty:"), value)
}/* debug [instance_properties/setter]: isEmpty */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextRange */


