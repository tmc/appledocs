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
	HeaderLevel() int
	SetHeaderLevel(value int)
	LineBreakStrategy() unsafe.Pointer
	SetLineBreakStrategy(value unsafe.Pointer)
	MaximumLineHeight() float64
	SetMaximumLineHeight(value float64)
	MinimumLineHeight() float64
	SetMinimumLineHeight(value float64)
	Alignment() unsafe.Pointer
	SetAlignment(value unsafe.Pointer)
	AllowsDefaultTighteningForTruncation() bool
	SetAllowsDefaultTighteningForTruncation(value bool)
	BaseWritingDirection() unsafe.Pointer
	SetBaseWritingDirection(value unsafe.Pointer)
	DefaultTabInterval() float64
	SetDefaultTabInterval(value float64)
	FirstLineHeadIndent() float64
	SetFirstLineHeadIndent(value float64)
	HeadIndent() float64
	SetHeadIndent(value float64)
	HyphenationFactor() float32
	SetHyphenationFactor(value float32)
	LineBreakMode() unsafe.Pointer
	SetLineBreakMode(value unsafe.Pointer)
	LineHeightMultiple() float64
	SetLineHeightMultiple(value float64)
	LineSpacing() float64
	SetLineSpacing(value float64)
	ParagraphSpacing() float64
	SetParagraphSpacing(value float64)
	ParagraphSpacingBefore() float64
	SetParagraphSpacingBefore(value float64)
	TabStops() TextTab
	SetTabStops(value TextTab)
	TailIndent() float64
	SetTailIndent(value float64)
	TextBlocks() TextBlock
	SetTextBlocks(value TextBlock)
	TextLists() TextList
	SetTextLists(value TextList)
	TighteningFactorForTruncation() float32
	SetTighteningFactorForTruncation(value float32)
	UsesDefaultHyphenation() bool
	SetUsesDefaultHyphenation(value bool)
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



// The paragraph’s header level for HTML generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headerLevel
func (m_ MutableParagraphStyle) HeaderLevel() int {
	rv := objc.Send[int](m_.ID, objc.Sel("headerLevel"))
	return rv
}


// The paragraph’s header level for HTML generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headerLevel
func (m_ MutableParagraphStyle) SetHeaderLevel(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeaderLevel:"), value)
}


// The strategies that the text system may use to break lines while laying out the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakStrategy
func (m_ MutableParagraphStyle) LineBreakStrategy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("lineBreakStrategy"))
	return rv
}


// The strategies that the text system may use to break lines while laying out the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakStrategy
func (m_ MutableParagraphStyle) SetLineBreakStrategy(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineBreakStrategy:"), value)
}


// The paragraph’s maximum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/maximumLineHeight
func (m_ MutableParagraphStyle) MaximumLineHeight() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("maximumLineHeight"))
	return rv
}


// The paragraph’s maximum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/maximumLineHeight
func (m_ MutableParagraphStyle) SetMaximumLineHeight(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumLineHeight:"), value)
}


// The paragraph’s minimum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/minimumLineHeight
func (m_ MutableParagraphStyle) MinimumLineHeight() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("minimumLineHeight"))
	return rv
}


// The paragraph’s minimum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/minimumLineHeight
func (m_ MutableParagraphStyle) SetMinimumLineHeight(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumLineHeight:"), value)
}


// The text alignment of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/alignment
func (m_ MutableParagraphStyle) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("alignment"))
	return rv
}


// The text alignment of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/alignment
func (m_ MutableParagraphStyle) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlignment:"), value)
}


// A Boolean value that indicates whether the system tightens intercharacter spacing before truncating text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/allowsdefaulttighteningfortruncation
func (m_ MutableParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}


// A Boolean value that indicates whether the system tightens intercharacter spacing before truncating text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/allowsdefaulttighteningfortruncation
func (m_ MutableParagraphStyle) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}


// The base writing direction for the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/basewritingdirection
func (m_ MutableParagraphStyle) BaseWritingDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// The base writing direction for the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/basewritingdirection
func (m_ MutableParagraphStyle) SetBaseWritingDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBaseWritingDirection:"), value)
}


// A number used as the document’s default tab spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/defaulttabinterval
func (m_ MutableParagraphStyle) DefaultTabInterval() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("defaultTabInterval"))
	return rv
}


// A number used as the document’s default tab spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/defaulttabinterval
func (m_ MutableParagraphStyle) SetDefaultTabInterval(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultTabInterval:"), value)
}


// The indentation of the first line of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/firstlineheadindent
func (m_ MutableParagraphStyle) FirstLineHeadIndent() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("firstLineHeadIndent"))
	return rv
}


// The indentation of the first line of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/firstlineheadindent
func (m_ MutableParagraphStyle) SetFirstLineHeadIndent(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFirstLineHeadIndent:"), value)
}


// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/headindent
func (m_ MutableParagraphStyle) HeadIndent() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("headIndent"))
	return rv
}


// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/headindent
func (m_ MutableParagraphStyle) SetHeadIndent(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeadIndent:"), value)
}


// The paragraph’s threshold for hyphenation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/hyphenationfactor
func (m_ MutableParagraphStyle) HyphenationFactor() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// The paragraph’s threshold for hyphenation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/hyphenationfactor
func (m_ MutableParagraphStyle) SetHyphenationFactor(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHyphenationFactor:"), value)
}


// The mode for breaking lines in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linebreakmode
func (m_ MutableParagraphStyle) LineBreakMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// The mode for breaking lines in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linebreakmode
func (m_ MutableParagraphStyle) SetLineBreakMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineBreakMode:"), value)
}


// The line height multiple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/lineheightmultiple
func (m_ MutableParagraphStyle) LineHeightMultiple() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineHeightMultiple"))
	return rv
}


// The line height multiple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/lineheightmultiple
func (m_ MutableParagraphStyle) SetLineHeightMultiple(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineHeightMultiple:"), value)
}


// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linespacing
func (m_ MutableParagraphStyle) LineSpacing() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineSpacing"))
	return rv
}


// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linespacing
func (m_ MutableParagraphStyle) SetLineSpacing(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineSpacing:"), value)
}


// The space after the end of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/paragraphspacing
func (m_ MutableParagraphStyle) ParagraphSpacing() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("paragraphSpacing"))
	return rv
}


// The space after the end of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/paragraphspacing
func (m_ MutableParagraphStyle) SetParagraphSpacing(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphSpacing:"), value)
}


// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/paragraphspacingbefore
func (m_ MutableParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("paragraphSpacingBefore"))
	return rv
}


// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/paragraphspacingbefore
func (m_ MutableParagraphStyle) SetParagraphSpacingBefore(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphSpacingBefore:"), value)
}


// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tabstops
func (m_ MutableParagraphStyle) TabStops() TextTab {
	rv := objc.Send[TextTab](m_.ID, objc.Sel("tabStops"))
	return rv
}


// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tabstops
func (m_ MutableParagraphStyle) SetTabStops(value TextTab) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTabStops:"), value)
}


// The trailing indentation of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tailindent
func (m_ MutableParagraphStyle) TailIndent() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("tailIndent"))
	return rv
}


// The trailing indentation of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tailindent
func (m_ MutableParagraphStyle) SetTailIndent(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTailIndent:"), value)
}


// The text blocks that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textblocks
func (m_ MutableParagraphStyle) TextBlocks() TextBlock {
	rv := objc.Send[TextBlock](m_.ID, objc.Sel("textBlocks"))
	return rv
}


// The text blocks that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textblocks
func (m_ MutableParagraphStyle) SetTextBlocks(value TextBlock) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTextBlocks:"), value)
}


// The text lists that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textlists
func (m_ MutableParagraphStyle) TextLists() TextList {
	rv := objc.Send[TextList](m_.ID, objc.Sel("textLists"))
	return rv
}


// The text lists that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textlists
func (m_ MutableParagraphStyle) SetTextLists(value TextList) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTextLists:"), value)
}


// The threshold for using tightening as an alternative to truncation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tighteningfactorfortruncation
func (m_ MutableParagraphStyle) TighteningFactorForTruncation() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("tighteningFactorForTruncation"))
	return rv
}


// The threshold for using tightening as an alternative to truncation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tighteningfactorfortruncation
func (m_ MutableParagraphStyle) SetTighteningFactorForTruncation(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTighteningFactorForTruncation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/usesdefaulthyphenation
func (m_ MutableParagraphStyle) UsesDefaultHyphenation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/usesdefaulthyphenation
func (m_ MutableParagraphStyle) SetUsesDefaultHyphenation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUsesDefaultHyphenation:"), value)
}



