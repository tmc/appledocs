// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ParagraphStyle] class.
var (
	ParagraphStyleClass     _ParagraphStyleClass
	ParagraphStyleClassOnce sync.Once
)

func getParagraphStyleClass() _ParagraphStyleClass {
	ParagraphStyleClassOnce.Do(func() {
		ParagraphStyleClass = _ParagraphStyleClass{objc.GetClass("NSParagraphStyle")}
	})
	return ParagraphStyleClass
}

type _ParagraphStyleClass struct {
	class objc.Class
}

// An interface definition for the [ParagraphStyle] class.
type IParagraphStyle interface {
	objectivec.IObject
	AllowsDefaultTighteningForTruncation() bool
	HyphenationFactor() float32
	LineBreakMode() LineBreakMode
	LineBreakStrategy() LineBreakStrategy
	TighteningFactorForTruncation() float32
	UsesDefaultHyphenation() bool
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	DefaultTabInterval() float64
	SetDefaultTabInterval(value float64)
	FirstLineHeadIndent() float64
	SetFirstLineHeadIndent(value float64)
	HeadIndent() float64
	SetHeadIndent(value float64)
	HeaderLevel() int
	SetHeaderLevel(value int)
	LineHeightMultiple() float64
	SetLineHeightMultiple(value float64)
	LineSpacing() float64
	SetLineSpacing(value float64)
	MaximumLineHeight() float64
	SetMaximumLineHeight(value float64)
	MinimumLineHeight() float64
	SetMinimumLineHeight(value float64)
	ParagraphSpacing() float64
	SetParagraphSpacing(value float64)
	ParagraphSpacingBefore() float64
	SetParagraphSpacingBefore(value float64)
	TabStops() NSTextTab
	SetTabStops(value ITextTab)
	TailIndent() float64
	SetTailIndent(value float64)
	TextBlocks() NSTextBlock
	SetTextBlocks(value ITextBlock)
	TextLists() NSTextList
	SetTextLists(value ITextList)
}

// The paragraph or ruler attributes for an attributed string.
//
// An object stores formatting information for a paragraph of text. The formatting information includes the amount of space between lines, indentations for lines of text, line heights, tab-stop positions, and more. Apply paragraph styles to the text of an attributed string by adding the attribute and setting its value to an instance of this class. The text-rendering system uses the paragraph style information in an attributed string to lay out and render the text. The class manages an immutable set of style information, but you can create an when you want to modify the style information before applying it to your text.


// The paragraph or ruler attributes for an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle

type ParagraphStyle struct {
	objectivec.Object
}

// ParagraphStyleFrom constructs a [ParagraphStyle] from an unsafe.Pointer.
//
// The paragraph or ruler attributes for an attributed string.
func ParagraphStyleFrom(ptr unsafe.Pointer) ParagraphStyle {
	return ParagraphStyle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ParagraphStyleClass) Alloc() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ParagraphStyleClass) New() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParagraphStyle) Init() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParagraphStyle) Autorelease() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParagraphStyle creates a new ParagraphStyle instance.
func NewParagraphStyle() ParagraphStyle {
	return getParagraphStyleClass().New()
}



// A Boolean value that indicates whether the system tightens character spacing before truncating text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/allowsDefaultTighteningForTruncation

func (p_ ParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}


// The paragraph’s threshold for hyphenation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/hyphenationFactor

func (p_ ParagraphStyle) HyphenationFactor() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// The mode for breaking lines in the paragraph that don’t fit within a container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineBreakMode

func (p_ ParagraphStyle) LineBreakMode() LineBreakMode {
	rv := objc.Send[LineBreakMode](p_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// The strategy for breaking lines while laying out paragraphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineBreakStrategy-swift.property

func (p_ ParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	rv := objc.Send[LineBreakStrategy](p_.ID, objc.Sel("lineBreakStrategy"))
	return rv
}


// The threshold for using tightening as an alternative to truncation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/tighteningFactorForTruncation

func (p_ ParagraphStyle) TighteningFactorForTruncation() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("tighteningFactorForTruncation"))
	return rv
}


// A Boolean value that indicates whether the paragraph style uses the system hyphenation settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/usesDefaultHyphenation

func (p_ ParagraphStyle) UsesDefaultHyphenation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}


// The text alignment of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/alignment

func (p_ ParagraphStyle) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](p_.ID, objc.Sel("alignment"))
	return rv
}


// The text alignment of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/alignment

func (p_ ParagraphStyle) SetAlignment(value TextAlignment) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignment:"), value)
}


// The base writing direction for the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/basewritingdirection

func (p_ ParagraphStyle) BaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](p_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// The base writing direction for the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/basewritingdirection

func (p_ ParagraphStyle) SetBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBaseWritingDirection:"), value)
}


// The documentwide default tab interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/defaulttabinterval

func (p_ ParagraphStyle) DefaultTabInterval() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("defaultTabInterval"))
	return rv
}


// The documentwide default tab interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/defaulttabinterval

func (p_ ParagraphStyle) SetDefaultTabInterval(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultTabInterval:"), value)
}


// The indentation of the first line of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/firstlineheadindent

func (p_ ParagraphStyle) FirstLineHeadIndent() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("firstLineHeadIndent"))
	return rv
}


// The indentation of the first line of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/firstlineheadindent

func (p_ ParagraphStyle) SetFirstLineHeadIndent(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFirstLineHeadIndent:"), value)
}


// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/headindent

func (p_ ParagraphStyle) HeadIndent() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("headIndent"))
	return rv
}


// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/headindent

func (p_ ParagraphStyle) SetHeadIndent(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHeadIndent:"), value)
}


// The paragraph’s header level for HTML generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/headerlevel

func (p_ ParagraphStyle) HeaderLevel() int {
	rv := objc.Send[int](p_.ID, objc.Sel("headerLevel"))
	return rv
}


// The paragraph’s header level for HTML generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/headerlevel

func (p_ ParagraphStyle) SetHeaderLevel(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHeaderLevel:"), value)
}


// The line height multiple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/lineheightmultiple

func (p_ ParagraphStyle) LineHeightMultiple() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineHeightMultiple"))
	return rv
}


// The line height multiple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/lineheightmultiple

func (p_ ParagraphStyle) SetLineHeightMultiple(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineHeightMultiple:"), value)
}


// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/linespacing

func (p_ ParagraphStyle) LineSpacing() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineSpacing"))
	return rv
}


// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/linespacing

func (p_ ParagraphStyle) SetLineSpacing(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineSpacing:"), value)
}


// The paragraph’s maximum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/maximumlineheight

func (p_ ParagraphStyle) MaximumLineHeight() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("maximumLineHeight"))
	return rv
}


// The paragraph’s maximum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/maximumlineheight

func (p_ ParagraphStyle) SetMaximumLineHeight(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaximumLineHeight:"), value)
}


// The paragraph’s minimum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/minimumlineheight

func (p_ ParagraphStyle) MinimumLineHeight() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minimumLineHeight"))
	return rv
}


// The paragraph’s minimum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/minimumlineheight

func (p_ ParagraphStyle) SetMinimumLineHeight(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMinimumLineHeight:"), value)
}


// Distance between the bottom of this paragraph and top of next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/paragraphspacing

func (p_ ParagraphStyle) ParagraphSpacing() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("paragraphSpacing"))
	return rv
}


// Distance between the bottom of this paragraph and top of next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/paragraphspacing

func (p_ ParagraphStyle) SetParagraphSpacing(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParagraphSpacing:"), value)
}


// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/paragraphspacingbefore

func (p_ ParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("paragraphSpacingBefore"))
	return rv
}


// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/paragraphspacingbefore

func (p_ ParagraphStyle) SetParagraphSpacingBefore(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParagraphSpacingBefore:"), value)
}


// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/tabstops

func (p_ ParagraphStyle) TabStops() NSTextTab {
	rv := objc.Send[NSTextTab](p_.ID, objc.Sel("tabStops"))
	return rv
}


// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/tabstops

func (p_ ParagraphStyle) SetTabStops(value ITextTab) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTabStops:"), value)
}


// The trailing indentation of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/tailindent

func (p_ ParagraphStyle) TailIndent() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("tailIndent"))
	return rv
}


// The trailing indentation of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/tailindent

func (p_ ParagraphStyle) SetTailIndent(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTailIndent:"), value)
}


// The text blocks that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/textblocks

func (p_ ParagraphStyle) TextBlocks() NSTextBlock {
	rv := objc.Send[NSTextBlock](p_.ID, objc.Sel("textBlocks"))
	return rv
}


// The text blocks that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/textblocks

func (p_ ParagraphStyle) SetTextBlocks(value ITextBlock) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTextBlocks:"), value)
}


// The text lists that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/textlists

func (p_ ParagraphStyle) TextLists() NSTextList {
	rv := objc.Send[NSTextList](p_.ID, objc.Sel("textLists"))
	return rv
}


// The text lists that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/textlists

func (p_ ParagraphStyle) SetTextLists(value ITextList) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTextLists:"), value)
}



