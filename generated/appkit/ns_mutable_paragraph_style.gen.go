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
}

// An object for changing the values of the subattributes in a paragraph style attribute.
//
// The class adds methods to its superclass, , for changing the values of the subattributes in a paragraph style attribute. For more information, see and .
//
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


// The text alignment of the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/alignment
func (m_ MutableParagraphStyle) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](m_.ID, objc.Sel("alignment"))
	return rv
}


// SetAlignment sets the value of the alignment property.
// The text alignment of the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/alignment
func (m_ MutableParagraphStyle) SetAlignment(value TextAlignment) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlignment:"), value)
}

// A Boolean value that indicates whether the system tightens intercharacter spacing before truncating text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/allowsdefaulttighteningfortruncation
func (m_ MutableParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}


// SetAllowsDefaultTighteningForTruncation sets the value of the allowsDefaultTighteningForTruncation property.
// A Boolean value that indicates whether the system tightens intercharacter spacing before truncating text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/allowsdefaulttighteningfortruncation
func (m_ MutableParagraphStyle) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}

// The base writing direction for the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/basewritingdirection
func (m_ MutableParagraphStyle) BaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](m_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// SetBaseWritingDirection sets the value of the baseWritingDirection property.
// The base writing direction for the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/basewritingdirection
func (m_ MutableParagraphStyle) SetBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBaseWritingDirection:"), value)
}

// A number used as the document’s default tab spacing.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/defaulttabinterval
func (m_ MutableParagraphStyle) DefaultTabInterval() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("defaultTabInterval"))
	return rv
}


// SetDefaultTabInterval sets the value of the defaultTabInterval property.
// A number used as the document’s default tab spacing.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/defaulttabinterval
func (m_ MutableParagraphStyle) SetDefaultTabInterval(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultTabInterval:"), value)
}

// The indentation of the first line of the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/firstlineheadindent
func (m_ MutableParagraphStyle) FirstLineHeadIndent() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("firstLineHeadIndent"))
	return rv
}


// SetFirstLineHeadIndent sets the value of the firstLineHeadIndent property.
// The indentation of the first line of the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/firstlineheadindent
func (m_ MutableParagraphStyle) SetFirstLineHeadIndent(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFirstLineHeadIndent:"), value)
}

// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/headindent
func (m_ MutableParagraphStyle) HeadIndent() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("headIndent"))
	return rv
}


// SetHeadIndent sets the value of the headIndent property.
// The indentation of the paragraph’s lines other than the first.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/headindent
func (m_ MutableParagraphStyle) SetHeadIndent(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeadIndent:"), value)
}

// The paragraph’s header level for HTML generation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/headerlevel
func (m_ MutableParagraphStyle) HeaderLevel() int {
	rv := objc.Send[int](m_.ID, objc.Sel("headerLevel"))
	return rv
}


// SetHeaderLevel sets the value of the headerLevel property.
// The paragraph’s header level for HTML generation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/headerlevel
func (m_ MutableParagraphStyle) SetHeaderLevel(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeaderLevel:"), value)
}

// The paragraph’s threshold for hyphenation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/hyphenationfactor
func (m_ MutableParagraphStyle) HyphenationFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// SetHyphenationFactor sets the value of the hyphenationFactor property.
// The paragraph’s threshold for hyphenation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/hyphenationfactor
func (m_ MutableParagraphStyle) SetHyphenationFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHyphenationFactor:"), value)
}

// The mode for breaking lines in the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linebreakmode
func (m_ MutableParagraphStyle) LineBreakMode() LineBreakMode {
	rv := objc.Send[LineBreakMode](m_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// SetLineBreakMode sets the value of the lineBreakMode property.
// The mode for breaking lines in the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linebreakmode
func (m_ MutableParagraphStyle) SetLineBreakMode(value LineBreakMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineBreakMode:"), value)
}

// The strategies that the text system may use to break lines while laying out the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linebreakstrategy
func (m_ MutableParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	rv := objc.Send[LineBreakStrategy](m_.ID, objc.Sel("lineBreakStrategy"))
	return rv
}


// SetLineBreakStrategy sets the value of the lineBreakStrategy property.
// The strategies that the text system may use to break lines while laying out the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linebreakstrategy
func (m_ MutableParagraphStyle) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineBreakStrategy:"), value)
}

// The line height multiple.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/lineheightmultiple
func (m_ MutableParagraphStyle) LineHeightMultiple() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineHeightMultiple"))
	return rv
}


// SetLineHeightMultiple sets the value of the lineHeightMultiple property.
// The line height multiple.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/lineheightmultiple
func (m_ MutableParagraphStyle) SetLineHeightMultiple(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineHeightMultiple:"), value)
}

// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linespacing
func (m_ MutableParagraphStyle) LineSpacing() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineSpacing"))
	return rv
}


// SetLineSpacing sets the value of the lineSpacing property.
// The distance in points between the bottom of one line fragment and the top of the next.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/linespacing
func (m_ MutableParagraphStyle) SetLineSpacing(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineSpacing:"), value)
}

// The paragraph’s maximum line height.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/maximumlineheight
func (m_ MutableParagraphStyle) MaximumLineHeight() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("maximumLineHeight"))
	return rv
}


// SetMaximumLineHeight sets the value of the maximumLineHeight property.
// The paragraph’s maximum line height.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/maximumlineheight
func (m_ MutableParagraphStyle) SetMaximumLineHeight(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumLineHeight:"), value)
}

// The paragraph’s minimum line height.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/minimumlineheight
func (m_ MutableParagraphStyle) MinimumLineHeight() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("minimumLineHeight"))
	return rv
}


// SetMinimumLineHeight sets the value of the minimumLineHeight property.
// The paragraph’s minimum line height.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/minimumlineheight
func (m_ MutableParagraphStyle) SetMinimumLineHeight(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumLineHeight:"), value)
}

// The space after the end of the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/paragraphspacing
func (m_ MutableParagraphStyle) ParagraphSpacing() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("paragraphSpacing"))
	return rv
}


// SetParagraphSpacing sets the value of the paragraphSpacing property.
// The space after the end of the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/paragraphspacing
func (m_ MutableParagraphStyle) SetParagraphSpacing(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphSpacing:"), value)
}

// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/paragraphspacingbefore
func (m_ MutableParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("paragraphSpacingBefore"))
	return rv
}


// SetParagraphSpacingBefore sets the value of the paragraphSpacingBefore property.
// The distance between the paragraph’s top and the beginning of its text content.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/paragraphspacingbefore
func (m_ MutableParagraphStyle) SetParagraphSpacingBefore(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphSpacingBefore:"), value)
}

// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tabstops
func (m_ MutableParagraphStyle) TabStops() NSTextTab {
	rv := objc.Send[NSTextTab](m_.ID, objc.Sel("tabStops"))
	return rv
}


// SetTabStops sets the value of the tabStops property.
// The text tab objects that represent the paragraph’s tab stops.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tabstops
func (m_ MutableParagraphStyle) SetTabStops(value ITextTab) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTabStops:"), value)
}

// The trailing indentation of the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tailindent
func (m_ MutableParagraphStyle) TailIndent() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("tailIndent"))
	return rv
}


// SetTailIndent sets the value of the tailIndent property.
// The trailing indentation of the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tailindent
func (m_ MutableParagraphStyle) SetTailIndent(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTailIndent:"), value)
}

// The text blocks that contain the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textblocks
func (m_ MutableParagraphStyle) TextBlocks() NSTextBlock {
	rv := objc.Send[NSTextBlock](m_.ID, objc.Sel("textBlocks"))
	return rv
}


// SetTextBlocks sets the value of the textBlocks property.
// The text blocks that contain the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textblocks
func (m_ MutableParagraphStyle) SetTextBlocks(value ITextBlock) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTextBlocks:"), value)
}

// The text lists that contain the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textlists
func (m_ MutableParagraphStyle) TextLists() NSTextList {
	rv := objc.Send[NSTextList](m_.ID, objc.Sel("textLists"))
	return rv
}


// SetTextLists sets the value of the textLists property.
// The text lists that contain the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textlists
func (m_ MutableParagraphStyle) SetTextLists(value ITextList) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTextLists:"), value)
}

// The threshold for using tightening as an alternative to truncation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tighteningfactorfortruncation
func (m_ MutableParagraphStyle) TighteningFactorForTruncation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("tighteningFactorForTruncation"))
	return rv
}


// SetTighteningFactorForTruncation sets the value of the tighteningFactorForTruncation property.
// The threshold for using tightening as an alternative to truncation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/tighteningfactorfortruncation
func (m_ MutableParagraphStyle) SetTighteningFactorForTruncation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTighteningFactorForTruncation:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/usesdefaulthyphenation
func (m_ MutableParagraphStyle) UsesDefaultHyphenation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}


// SetUsesDefaultHyphenation sets the value of the usesDefaultHyphenation property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/usesdefaulthyphenation
func (m_ MutableParagraphStyle) SetUsesDefaultHyphenation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUsesDefaultHyphenation:"), value)
}



