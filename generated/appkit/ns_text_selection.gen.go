// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
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
	Affinity() unsafe.Pointer
	SetAffinity(value unsafe.Pointer)
	AnchorPositionOffset() float64
	SetAnchorPositionOffset(value float64)
	Granularity() unsafe.Pointer
	SetGranularity(value unsafe.Pointer)
	IsLogical() bool
	SetIsLogical(value bool)
	IsTransient() bool
	SetIsTransient(value bool)
	SecondarySelectionLocation() TextLocation /* not a class type */
	SetSecondarySelectionLocation(value TextLocation /* not a class type */)
	TextRanges() objc.IObject /* cross-framework: TextRange */
	SetTextRanges(value objc.IObject /* cross-framework: TextRange */)
	TypingAttributes() objc.IObject /* cross-framework: Key */
	SetTypingAttributes(value objc.IObject /* cross-framework: Key */)
	// methods:
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



// Returns the selection affinity of the text selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/affinity-swift.property
func (t_ TextSelection) Affinity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("affinity"))
	return rv
}


// Returns the selection affinity of the text selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/affinity-swift.property
func (t_ TextSelection) SetAffinity(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAffinity:"), value)
}


// Represents the anchor position offset from the beginning of a line fragment in the visual order for the initial tap or click location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/anchorpositionoffset
func (t_ TextSelection) AnchorPositionOffset() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("anchorPositionOffset"))
	return rv
}


// Represents the anchor position offset from the beginning of a line fragment in the visual order for the initial tap or click location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/anchorpositionoffset
func (t_ TextSelection) SetAnchorPositionOffset(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAnchorPositionOffset:"), value)
}


// The granularity of the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/granularity-swift.property
func (t_ TextSelection) Granularity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("granularity"))
	return rv
}


// The granularity of the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/granularity-swift.property
func (t_ TextSelection) SetGranularity(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGranularity:"), value)
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


// Specifies the secondary character location when user taps or clicks at a directional boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/secondaryselectionlocation
func (t_ TextSelection) SecondarySelectionLocation() TextLocation /* not a class type */ {
	rv := objc.Send[TextLocation](t_.ID, objc.Sel("secondarySelectionLocation"))
	return rv
}


// Specifies the secondary character location when user taps or clicks at a directional boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/secondaryselectionlocation
func (t_ TextSelection) SetSecondarySelectionLocation(value TextLocation /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSecondarySelectionLocation:"), value)
}


// Represents an array of noncontiguous logical ranges in the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/textranges
func (t_ TextSelection) TextRanges() objc.IObject /* cross-framework: TextRange */ {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("textRanges"))
	return rv
}


// Represents an array of noncontiguous logical ranges in the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/textranges
func (t_ TextSelection) SetTextRanges(value objc.IObject /* cross-framework: TextRange */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextRanges:"), value)
}


// The template attributes the framework uses for characters that replace the contents of this selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/typingattributes
func (t_ TextSelection) TypingAttributes() objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("typingAttributes"))
	return rv
}


// The template attributes the framework uses for characters that replace the contents of this selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselection/typingattributes
func (t_ TextSelection) SetTypingAttributes(value objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypingAttributes:"), value)
}



