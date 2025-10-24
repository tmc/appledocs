// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextSelection] class.
var (
	TextSelectionClass     _TextSelectionClass
	TextSelectionClassOnce sync.Once
)

func getTextSelectionClass() _TextSelectionClass {
	TextSelectionClassOnce.Do(func() {
		TextSelectionClass = _TextSelectionClass{objc.GetClass("NSTextSelection")}
	})
	return TextSelectionClass
}

type _TextSelectionClass struct {
	class objc.Class
}

// An interface definition for the [TextSelection] class.
type ITextSelection interface {
	objectivec.IObject
	// properties:
	Affinity() TextSelectionAffinity
	AnchorPositionOffset() float64
	SetAnchorPositionOffset(value float64)
	Granularity() TextSelectionGranularity
	Logical() bool
	SetLogical(value bool)
	Transient() bool
	SecondarySelectionLocation() objc.ID
	SetSecondarySelectionLocation(value objc.ID)
	TextRanges() []TextRange
	TypingAttributes() foundation.IDictionary
	SetTypingAttributes(value foundation.IDictionary)
	IsLogical() bool
	SetIsLogical(value bool)
	IsTransient() bool
	SetIsTransient(value bool)
	// methods:
	TextSelectionWithTextRanges(textRanges []TextRange) ITextSelection
}

// A class that represents a single logical selection context that corresponds to an insertion point.


// A class that represents a single logical selection context that corresponds to an insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection
type TextSelection struct {
	objectivec.Object
}

// TextSelectionFrom constructs a [TextSelection] from an unsafe.Pointer.
//
// A class that represents a single logical selection context that corresponds to an insertion point.
func TextSelectionFrom(ptr unsafe.Pointer) TextSelection {
	return TextSelection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextSelectionClass) Alloc() TextSelection {
	rv := objc.Send[TextSelection](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextSelectionClass) New() TextSelection {
	rv := objc.Send[TextSelection](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextSelection) Init() TextSelection {
	rv := objc.Send[TextSelection](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextSelection) Autorelease() TextSelection {
	rv := objc.Send[TextSelection](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextSelection creates a new TextSelection instance.
func NewTextSelection() TextSelection {
	return getTextSelectionClass().New()
}



// Creates a test selection from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(coder:)
func NewTextSelectionWithCoder(coder foundation.Coder) TextSelection {
	instance := getTextSelectionClass().Alloc()
	rv := objc.Send[TextSelection](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Creates a new text selection with the location and selection affinity you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(_:affinity:)
func NewTextSelectionWithLocationAffinity(location objc.IObject, affinity TextSelectionAffinity) TextSelection {
	instance := getTextSelectionClass().Alloc()
	rv := objc.Send[TextSelection](instance.ID, objc.Sel("initWithLocation:affinity:"), location, affinity)
	rv.Autorelease()
	return rv
}


// Creates a new text selection with the range, selection affinity, and granularity you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(range:affinity:granularity:)
func NewTextSelectionWithRangeAffinityGranularity(range_ ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	instance := getTextSelectionClass().Alloc()
	rv := objc.Send[TextSelection](instance.ID, objc.Sel("initWithRange:affinity:granularity:"), range_, affinity, granularity)
	rv.Autorelease()
	return rv
}


// Creates a new text selection with the ranges, selection affinity, and granularity you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(_:affinity:granularity:)
func NewTextSelectionWithRangesAffinityGranularity(textRanges []TextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	instance := getTextSelectionClass().Alloc()
	rv := objc.Send[TextSelection](instance.ID, objc.Sel("initWithRanges:affinity:granularity:"), textRanges, affinity, granularity)
	rv.Autorelease()
	return rv
}



// Creates a subselection of the current text selection with the ranges you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/textSelection(_:)
func (t_ TextSelection) TextSelectionWithTextRanges(textRanges []TextRange) ITextSelection {
	rv := objc.Send[TextSelection](t_.ID, objc.Sel("textSelectionWithTextRanges:"), textRanges)
	return rv
}


// Returns the selection affinity of the text selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/affinity-swift.property
func (t_ TextSelection) Affinity() TextSelectionAffinity {
	rv := objc.Send[TextSelectionAffinity](t_.ID, objc.Sel("affinity"))
	return rv
}


// Represents the anchor position offset from the beginning of a line fragment in the visual order for the initial tap or click location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/anchorPositionOffset
func (t_ TextSelection) AnchorPositionOffset() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("anchorPositionOffset"))
	return rv
}


// Represents the anchor position offset from the beginning of a line fragment in the visual order for the initial tap or click location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/anchorPositionOffset
func (t_ TextSelection) SetAnchorPositionOffset(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAnchorPositionOffset:"), value)
}


// The granularity of the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/granularity-swift.property
func (t_ TextSelection) Granularity() TextSelectionGranularity {
	rv := objc.Send[TextSelectionGranularity](t_.ID, objc.Sel("granularity"))
	return rv
}


// A Boolean value that indicates whether the framework interprets the selection as logical or visual.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/isLogical
func (t_ TextSelection) Logical() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("logical"))
	return rv
}


// A Boolean value that indicates whether the framework interprets the selection as logical or visual.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/isLogical
func (t_ TextSelection) SetLogical(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLogical:"), value)
}


// A Boolean value that indicates transient text selection during drag handling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/isTransient
func (t_ TextSelection) Transient() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("transient"))
	return rv
}


// Specifies the secondary character location when user taps or clicks at a directional boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/secondarySelectionLocation
func (t_ TextSelection) SecondarySelectionLocation() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("secondarySelectionLocation"))
	return rv
}


// Specifies the secondary character location when user taps or clicks at a directional boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/secondarySelectionLocation
func (t_ TextSelection) SetSecondarySelectionLocation(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSecondarySelectionLocation:"), value)
}


// Represents an array of noncontiguous logical ranges in the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/textRanges
func (t_ TextSelection) TextRanges() []TextRange {
	rv := objc.Send[[]TextRange](t_.ID, objc.Sel("textRanges"))
	return rv
}


// The template attributes the framework uses for characters that replace the contents of this selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/typingAttributes
func (t_ TextSelection) TypingAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("typingAttributes"))
	return rv
}


// The template attributes the framework uses for characters that replace the contents of this selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/typingAttributes
func (t_ TextSelection) SetTypingAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypingAttributes:"), value)
}


// A Boolean value that indicates whether the framework interprets the selection as logical or visual.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/islogical
func (t_ TextSelection) IsLogical() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isLogical"))
	return rv
}


// A Boolean value that indicates whether the framework interprets the selection as logical or visual.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/islogical
func (t_ TextSelection) SetIsLogical(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsLogical:"), value)
}


// A Boolean value that indicates transient text selection during drag handling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/istransient
func (t_ TextSelection) IsTransient() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isTransient"))
	return rv
}


// A Boolean value that indicates transient text selection during drag handling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/istransient
func (t_ TextSelection) SetIsTransient(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsTransient:"), value)
}


