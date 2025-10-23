// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableParagraphStyle] class.
var (
	MutableParagraphStyleClass     _MutableParagraphStyleClass
	MutableParagraphStyleClassOnce sync.Once
)

func getMutableParagraphStyleClass() _MutableParagraphStyleClass {
	MutableParagraphStyleClassOnce.Do(func() {
		MutableParagraphStyleClass = _MutableParagraphStyleClass{objc.GetClass("NSMutableParagraphStyle")}
	})
	return MutableParagraphStyleClass
}

type _MutableParagraphStyleClass struct {
	class objc.Class
}

// An interface definition for the [MutableParagraphStyle] class.
type IMutableParagraphStyle interface {
	IParagraphStyle
	// properties:
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	AllowsDefaultTighteningForTruncation() bool /* primitive/slice/pointer. */
	SetAllowsDefaultTighteningForTruncation(value bool /* primitive/slice/pointer. */)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	DefaultTabInterval() float64 /* primitive/slice/pointer. */
	SetDefaultTabInterval(value float64 /* primitive/slice/pointer. */)
	FirstLineHeadIndent() float64 /* primitive/slice/pointer. */
	SetFirstLineHeadIndent(value float64 /* primitive/slice/pointer. */)
	HeadIndent() float64 /* primitive/slice/pointer. */
	SetHeadIndent(value float64 /* primitive/slice/pointer. */)
	HeaderLevel() int /* primitive/slice/pointer. */
	SetHeaderLevel(value int /* primitive/slice/pointer. */)
	HyphenationFactor() float32 /* primitive/slice/pointer. */
	SetHyphenationFactor(value float32 /* primitive/slice/pointer. */)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	LineBreakStrategy() LineBreakStrategy
	SetLineBreakStrategy(value LineBreakStrategy)
	LineHeightMultiple() float64 /* primitive/slice/pointer. */
	SetLineHeightMultiple(value float64 /* primitive/slice/pointer. */)
	LineSpacing() float64 /* primitive/slice/pointer. */
	SetLineSpacing(value float64 /* primitive/slice/pointer. */)
	MaximumLineHeight() float64 /* primitive/slice/pointer. */
	SetMaximumLineHeight(value float64 /* primitive/slice/pointer. */)
	MinimumLineHeight() float64 /* primitive/slice/pointer. */
	SetMinimumLineHeight(value float64 /* primitive/slice/pointer. */)
	ParagraphSpacing() float64 /* primitive/slice/pointer. */
	SetParagraphSpacing(value float64 /* primitive/slice/pointer. */)
	ParagraphSpacingBefore() float64 /* primitive/slice/pointer. */
	SetParagraphSpacingBefore(value float64 /* primitive/slice/pointer. */)
	TabStops() []TextTab /* primitive/slice/pointer. */
	SetTabStops(value []TextTab /* primitive/slice/pointer. */)
	TailIndent() float64 /* primitive/slice/pointer. */
	SetTailIndent(value float64 /* primitive/slice/pointer. */)
	TextBlocks() []TextBlock /* primitive/slice/pointer. */
	SetTextBlocks(value []TextBlock /* primitive/slice/pointer. */)
	TextLists() []TextList /* primitive/slice/pointer. */
	SetTextLists(value []TextList /* primitive/slice/pointer. */)
	TighteningFactorForTruncation() float32 /* primitive/slice/pointer. */
	SetTighteningFactorForTruncation(value float32 /* primitive/slice/pointer. */)
	UsesDefaultHyphenation() bool /* primitive/slice/pointer. */
	SetUsesDefaultHyphenation(value bool /* primitive/slice/pointer. */)
	// methods:
	AddTabStop(anObject ITextTab)
	RemoveTabStop(anObject ITextTab)
	SetParagraphStyle(obj IParagraphStyle)
}

// An object for changing the values of the subattributes in a paragraph style attribute.
//
// The class adds methods to its superclass, , for changing the values of the subattributes in a paragraph style attribute. For more information, see and .


// An object for changing the values of the subattributes in a paragraph style attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle
type MutableParagraphStyle struct {
	ParagraphStyle
}

// MutableParagraphStyleFrom constructs a [MutableParagraphStyle] from an unsafe.Pointer.
//
// An object for changing the values of the subattributes in a paragraph style attribute.
func MutableParagraphStyleFrom(ptr unsafe.Pointer) MutableParagraphStyle {
	return MutableParagraphStyle{
		ParagraphStyle: ParagraphStyleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableParagraphStyleClass) Alloc() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableParagraphStyleClass) New() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableParagraphStyle) Init() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableParagraphStyle) Autorelease() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableParagraphStyle creates a new MutableParagraphStyle instance.
func NewMutableParagraphStyle() MutableParagraphStyle {
	return getMutableParagraphStyleClass().New()
}



// Adds the specified tab stop to the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/addTabStop(_:)
func (m_ MutableParagraphStyle) AddTabStop(anObject ITextTab) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addTabStop:"), anObject)
}


// Removes the first text tab with a location and type equal to the specified tab stop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/removeTabStop(_:)
func (m_ MutableParagraphStyle) RemoveTabStop(anObject ITextTab) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTabStop:"), anObject)
}


// Replaces the subattributes of the paragraph with those in the specified paragraph style object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/setParagraphStyle(_:)
func (m_ MutableParagraphStyle) SetParagraphStyle(obj IParagraphStyle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphStyle:"), obj)
}


// The text alignment of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/alignment
func (m_ MutableParagraphStyle) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](m_.ID, objc.Sel("alignment"))
	return rv
}


// The text alignment of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/alignment
func (m_ MutableParagraphStyle) SetAlignment(value TextAlignment) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlignment:"), value)
}


// A Boolean value that indicates whether the system tightens intercharacter spacing before truncating text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/allowsDefaultTighteningForTruncation
func (m_ MutableParagraphStyle) AllowsDefaultTighteningForTruncation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}


// A Boolean value that indicates whether the system tightens intercharacter spacing before truncating text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/allowsDefaultTighteningForTruncation
func (m_ MutableParagraphStyle) SetAllowsDefaultTighteningForTruncation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}


// The base writing direction for the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/baseWritingDirection
func (m_ MutableParagraphStyle) BaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](m_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// The base writing direction for the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/baseWritingDirection
func (m_ MutableParagraphStyle) SetBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBaseWritingDirection:"), value)
}


// A number used as the document’s default tab spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/defaultTabInterval
func (m_ MutableParagraphStyle) DefaultTabInterval() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("defaultTabInterval"))
	return rv
}


// A number used as the document’s default tab spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/defaultTabInterval
func (m_ MutableParagraphStyle) SetDefaultTabInterval(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultTabInterval:"), value)
}


// The indentation of the first line of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/firstLineHeadIndent
func (m_ MutableParagraphStyle) FirstLineHeadIndent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("firstLineHeadIndent"))
	return rv
}


// The indentation of the first line of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/firstLineHeadIndent
func (m_ MutableParagraphStyle) SetFirstLineHeadIndent(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFirstLineHeadIndent:"), value)
}


// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headIndent
func (m_ MutableParagraphStyle) HeadIndent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("headIndent"))
	return rv
}


// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headIndent
func (m_ MutableParagraphStyle) SetHeadIndent(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeadIndent:"), value)
}


// The paragraph’s header level for HTML generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headerLevel
func (m_ MutableParagraphStyle) HeaderLevel() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("headerLevel"))
	return rv
}


// The paragraph’s header level for HTML generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headerLevel
func (m_ MutableParagraphStyle) SetHeaderLevel(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeaderLevel:"), value)
}


// The paragraph’s threshold for hyphenation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/hyphenationFactor
func (m_ MutableParagraphStyle) HyphenationFactor() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](m_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// The paragraph’s threshold for hyphenation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/hyphenationFactor
func (m_ MutableParagraphStyle) SetHyphenationFactor(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHyphenationFactor:"), value)
}


// The mode for breaking lines in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakMode
func (m_ MutableParagraphStyle) LineBreakMode() LineBreakMode {
	rv := objc.Send[LineBreakMode](m_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// The mode for breaking lines in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakMode
func (m_ MutableParagraphStyle) SetLineBreakMode(value LineBreakMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineBreakMode:"), value)
}


// The strategies that the text system may use to break lines while laying out the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakStrategy
func (m_ MutableParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	rv := objc.Send[LineBreakStrategy](m_.ID, objc.Sel("lineBreakStrategy"))
	return rv
}


// The strategies that the text system may use to break lines while laying out the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakStrategy
func (m_ MutableParagraphStyle) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineBreakStrategy:"), value)
}


// The line height multiple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineHeightMultiple
func (m_ MutableParagraphStyle) LineHeightMultiple() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineHeightMultiple"))
	return rv
}


// The line height multiple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineHeightMultiple
func (m_ MutableParagraphStyle) SetLineHeightMultiple(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineHeightMultiple:"), value)
}


// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineSpacing
func (m_ MutableParagraphStyle) LineSpacing() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineSpacing"))
	return rv
}


// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineSpacing
func (m_ MutableParagraphStyle) SetLineSpacing(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineSpacing:"), value)
}


// The paragraph’s maximum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/maximumLineHeight
func (m_ MutableParagraphStyle) MaximumLineHeight() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("maximumLineHeight"))
	return rv
}


// The paragraph’s maximum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/maximumLineHeight
func (m_ MutableParagraphStyle) SetMaximumLineHeight(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumLineHeight:"), value)
}


// The paragraph’s minimum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/minimumLineHeight
func (m_ MutableParagraphStyle) MinimumLineHeight() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("minimumLineHeight"))
	return rv
}


// The paragraph’s minimum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/minimumLineHeight
func (m_ MutableParagraphStyle) SetMinimumLineHeight(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumLineHeight:"), value)
}


// The space after the end of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/paragraphSpacing
func (m_ MutableParagraphStyle) ParagraphSpacing() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("paragraphSpacing"))
	return rv
}


// The space after the end of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/paragraphSpacing
func (m_ MutableParagraphStyle) SetParagraphSpacing(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphSpacing:"), value)
}


// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/paragraphSpacingBefore
func (m_ MutableParagraphStyle) ParagraphSpacingBefore() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("paragraphSpacingBefore"))
	return rv
}


// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/paragraphSpacingBefore
func (m_ MutableParagraphStyle) SetParagraphSpacingBefore(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphSpacingBefore:"), value)
}


// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tabStops
func (m_ MutableParagraphStyle) TabStops() []TextTab /* primitive/slice/pointer. */ {
	rv := objc.Send[[]TextTab](m_.ID, objc.Sel("tabStops"))
	return rv
}


// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tabStops
func (m_ MutableParagraphStyle) SetTabStops(value []TextTab /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setTabStops:"), nsArray)
}


// The trailing indentation of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tailIndent
func (m_ MutableParagraphStyle) TailIndent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("tailIndent"))
	return rv
}


// The trailing indentation of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tailIndent
func (m_ MutableParagraphStyle) SetTailIndent(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTailIndent:"), value)
}


// The text blocks that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/textBlocks
func (m_ MutableParagraphStyle) TextBlocks() []TextBlock /* primitive/slice/pointer. */ {
	rv := objc.Send[[]TextBlock](m_.ID, objc.Sel("textBlocks"))
	return rv
}


// The text blocks that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/textBlocks
func (m_ MutableParagraphStyle) SetTextBlocks(value []TextBlock /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setTextBlocks:"), nsArray)
}


// The text lists that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/textLists
func (m_ MutableParagraphStyle) TextLists() []TextList /* primitive/slice/pointer. */ {
	rv := objc.Send[[]TextList](m_.ID, objc.Sel("textLists"))
	return rv
}


// The text lists that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/textLists
func (m_ MutableParagraphStyle) SetTextLists(value []TextList /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setTextLists:"), nsArray)
}


// The threshold for using tightening as an alternative to truncation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tighteningFactorForTruncation
func (m_ MutableParagraphStyle) TighteningFactorForTruncation() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](m_.ID, objc.Sel("tighteningFactorForTruncation"))
	return rv
}


// The threshold for using tightening as an alternative to truncation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tighteningFactorForTruncation
func (m_ MutableParagraphStyle) SetTighteningFactorForTruncation(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTighteningFactorForTruncation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/usesDefaultHyphenation
func (m_ MutableParagraphStyle) UsesDefaultHyphenation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/usesDefaultHyphenation
func (m_ MutableParagraphStyle) SetUsesDefaultHyphenation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUsesDefaultHyphenation:"), value)
}



