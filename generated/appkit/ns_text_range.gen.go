// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TextRange] class.
type ITextRange interface {
	objectivec.IObject
	// properties:
	EndLocation() objc.ID
	Empty() bool
	Location() objc.ID
	IsEmpty() bool
	SetIsEmpty(value bool)
	// methods:
	ContainsRange(textRange ITextRange) bool
	ContainsLocation(location objc.IObject) bool
	TextRangeByIntersectingWithTextRange(textRange ITextRange) unsafe.Pointer
	IntersectsWithTextRange(textRange ITextRange) bool
	IsEqualToTextRange(textRange ITextRange) bool
	TextRangeByFormingUnionWithTextRange(textRange ITextRange) unsafe.Pointer
}

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

// Alloc allocates a new instance without initialization.
func (tc _TextRangeClass) Alloc() TextRange {
	rv := objc.Send[TextRange](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a new text range at the location you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/init(location:)
func NewTextRangeWithLocation(location objc.IObject) TextRange {
	instance := getTextRangeClass().Alloc()
	rv := objc.Send[TextRange](instance.ID, objc.Sel("initWithLocation:"), location)
	rv.Autorelease()
	return rv
}


// Creates a new text range with the starting and ending locations you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/init(location:end:)
func NewTextRangeWithLocationEndLocation(location objc.IObject, endLocation objc.IObject) TextRange {
	instance := getTextRangeClass().Alloc()
	rv := objc.Send[TextRange](instance.ID, objc.Sel("initWithLocation:endLocation:"), location, endLocation)
	rv.Autorelease()
	return rv
}



// Determines if the text range you specify is in the current text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/contains(_:)-5j4y2
func (t_ TextRange) ContainsRange(textRange ITextRange) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("containsRange:"), textRange)
	return rv
}


// Determines if the text location you specify is in the current text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/contains(_:)-7hvi0
func (t_ TextRange) ContainsLocation(location objc.IObject) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("containsLocation:"), location)
	return rv
}


// Returns the range, if any, where two text ranges intersect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/intersection(_:)
func (t_ TextRange) TextRangeByIntersectingWithTextRange(textRange ITextRange) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textRangeByIntersectingWithTextRange:"), textRange)
	return rv
}


// Determines if two ranges intersect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/intersects(_:)
func (t_ TextRange) IntersectsWithTextRange(textRange ITextRange) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("intersectsWithTextRange:"), textRange)
	return rv
}


// Compares two text ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/isEqual(to:)
func (t_ TextRange) IsEqualToTextRange(textRange ITextRange) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEqualToTextRange:"), textRange)
	return rv
}


// Returns a new text range by forming the union with the text range you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/union(_:)
func (t_ TextRange) TextRangeByFormingUnionWithTextRange(textRange ITextRange) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textRangeByFormingUnionWithTextRange:"), textRange)
	return rv
}


// The ending location of the text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/endLocation
func (t_ TextRange) EndLocation() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("endLocation"))
	return rv
}


// Returns whether the text range is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/isEmpty
func (t_ TextRange) Empty() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("empty"))
	return rv
}


// The starting location of the text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange/location
func (t_ TextRange) Location() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("location"))
	return rv
}


// Returns whether the text range is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextrange/isempty
func (t_ TextRange) IsEmpty() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEmpty"))
	return rv
}


// Returns whether the text range is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextrange/isempty
func (t_ TextRange) SetIsEmpty(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEmpty:"), value)
}


