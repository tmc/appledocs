// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Alignment() TextAlignment
	AllowsDefaultTighteningForTruncation() bool
	BaseWritingDirection() WritingDirection
	DefaultTabInterval() float64
	FirstLineHeadIndent() float64
	HeadIndent() float64
	HeaderLevel() int
	HyphenationFactor() float32
	LineBreakMode() LineBreakMode
	LineBreakStrategy() LineBreakStrategy
	LineHeightMultiple() float64
	LineSpacing() float64
	MaximumLineHeight() float64
	MinimumLineHeight() float64
	ParagraphSpacing() float64
	ParagraphSpacingBefore() float64
	TabStops() []TextTab
	TailIndent() float64
	TextBlocks() []TextBlock
	TextLists() []TextList
	TighteningFactorForTruncation() float32
	UsesDefaultHyphenation() bool
	// methods:
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



// Returns the default writing direction for the specified language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/defaultWritingDirection(forLanguage:)
func (pc _ParagraphStyleClass) DefaultWritingDirectionForLanguage(languageName objc.IObject /* cross-framework: NSString */) WritingDirection {
	rv := objc.Send[WritingDirection](objc.ID(pc.class), objc.Sel("defaultWritingDirectionForLanguage:"), languageName)
	return rv
}


// The default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/default
func (pc _ParagraphStyleClass) DefaultParagraphStyle() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](objc.ID(pc.class), objc.Sel("defaultParagraphStyle"))
	return rv
}

// The text alignment of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/alignment
func (p_ ParagraphStyle) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](p_.ID, objc.Sel("alignment"))
	return rv
}


// A Boolean value that indicates whether the system tightens character spacing before truncating text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/allowsDefaultTighteningForTruncation
func (p_ ParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}


// The base writing direction for the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/baseWritingDirection
func (p_ ParagraphStyle) BaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](p_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// The default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/default
func (p_ ParagraphStyle) DefaultParagraphStyle() IParagraphStyle {
	rv := objc.Send[ParagraphStyle](p_.ID, objc.Sel("defaultParagraphStyle"))
	return rv
}


// The documentwide default tab interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/defaultTabInterval
func (p_ ParagraphStyle) DefaultTabInterval() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("defaultTabInterval"))
	return rv
}


// The indentation of the first line of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/firstLineHeadIndent
func (p_ ParagraphStyle) FirstLineHeadIndent() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("firstLineHeadIndent"))
	return rv
}


// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/headIndent
func (p_ ParagraphStyle) HeadIndent() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("headIndent"))
	return rv
}


// The paragraph’s header level for HTML generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/headerLevel
func (p_ ParagraphStyle) HeaderLevel() int {
	rv := objc.Send[int](p_.ID, objc.Sel("headerLevel"))
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


// The line height multiple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineHeightMultiple
func (p_ ParagraphStyle) LineHeightMultiple() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineHeightMultiple"))
	return rv
}


// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineSpacing
func (p_ ParagraphStyle) LineSpacing() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineSpacing"))
	return rv
}


// The paragraph’s maximum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/maximumLineHeight
func (p_ ParagraphStyle) MaximumLineHeight() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("maximumLineHeight"))
	return rv
}


// The paragraph’s minimum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/minimumLineHeight
func (p_ ParagraphStyle) MinimumLineHeight() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minimumLineHeight"))
	return rv
}


// Distance between the bottom of this paragraph and top of next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/paragraphSpacing
func (p_ ParagraphStyle) ParagraphSpacing() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("paragraphSpacing"))
	return rv
}


// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/paragraphSpacingBefore
func (p_ ParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("paragraphSpacingBefore"))
	return rv
}


// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/tabStops
func (p_ ParagraphStyle) TabStops() []TextTab {
	rv := objc.Send[[]TextTab](p_.ID, objc.Sel("tabStops"))
	return rv
}


// The trailing indentation of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/tailIndent
func (p_ ParagraphStyle) TailIndent() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("tailIndent"))
	return rv
}


// The text blocks that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/textBlocks
func (p_ ParagraphStyle) TextBlocks() []TextBlock {
	rv := objc.Send[[]TextBlock](p_.ID, objc.Sel("textBlocks"))
	return rv
}


// The text lists that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/textLists
func (p_ ParagraphStyle) TextLists() []TextList {
	rv := objc.Send[[]TextList](p_.ID, objc.Sel("textLists"))
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



