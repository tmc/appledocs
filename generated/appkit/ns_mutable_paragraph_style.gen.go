// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMutableParagraphStyle */


/* debug [class_header]: Header for NSMutableParagraphStyle */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableParagraphStyle */
// An interface definition for the [MutableParagraphStyle] class.
type IMutableParagraphStyle interface {
	IParagraphStyle
	
/* debug [class_interface_properties]: Properties for MutableParagraphStyle */
	// properties:
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	AllowsDefaultTighteningForTruncation() bool
	SetAllowsDefaultTighteningForTruncation(value bool)
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
	HyphenationFactor() float32
	SetHyphenationFactor(value float32)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	LineBreakStrategy() LineBreakStrategy
	SetLineBreakStrategy(value LineBreakStrategy)
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
	TabStops() []TextTab
	SetTabStops(value []TextTab)
	TailIndent() float64
	SetTailIndent(value float64)
	TextBlocks() []TextBlock
	SetTextBlocks(value []TextBlock)
	TextLists() []TextList
	SetTextLists(value []TextList)
	TighteningFactorForTruncation() float32
	SetTighteningFactorForTruncation(value float32)
	UsesDefaultHyphenation() bool
	SetUsesDefaultHyphenation(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableParagraphStyle */
	// methods:
	AddTabStop(anObject ITextTab)
	RemoveTabStop(anObject ITextTab)
	SetParagraphStyle(obj IParagraphStyle)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableParagraphStyle */
// Alloc allocates a new instance without initialization.
func (mc _MutableParagraphStyleClass) Alloc() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableParagraphStyle */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableParagraphStyle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableParagraphStyle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableParagraphStyle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableParagraphStyle */

// Adds the specified tab stop to the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/addTabStop(_:)
func (m_ MutableParagraphStyle) AddTabStop(anObject ITextTab) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addTabStop:"), anObject)
}/* debug [instance_methods/method]: AddTabStop */


// Removes the first text tab with a location and type equal to the specified tab stop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/removeTabStop(_:)
func (m_ MutableParagraphStyle) RemoveTabStop(anObject ITextTab) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTabStop:"), anObject)
}/* debug [instance_methods/method]: RemoveTabStop */


// Replaces the subattributes of the paragraph with those in the specified paragraph style object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/setParagraphStyle(_:)
func (m_ MutableParagraphStyle) SetParagraphStyle(obj IParagraphStyle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphStyle:"), obj)
}/* debug [instance_methods/method]: SetParagraphStyle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableParagraphStyle */

// The text alignment of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/alignment
func (m_ MutableParagraphStyle) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](m_.ID, objc.Sel("alignment"))
	return rv
}/* debug [instance_properties/getter]: alignment */


// The text alignment of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/alignment
func (m_ MutableParagraphStyle) SetAlignment(value TextAlignment) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlignment:"), value)
}/* debug [instance_properties/setter]: alignment */


// A Boolean value that indicates whether the system tightens intercharacter spacing before truncating text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/allowsDefaultTighteningForTruncation
func (m_ MutableParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}/* debug [instance_properties/getter]: allowsDefaultTighteningForTruncation */


// A Boolean value that indicates whether the system tightens intercharacter spacing before truncating text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/allowsDefaultTighteningForTruncation
func (m_ MutableParagraphStyle) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}/* debug [instance_properties/setter]: allowsDefaultTighteningForTruncation */


// The base writing direction for the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/baseWritingDirection
func (m_ MutableParagraphStyle) BaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](m_.ID, objc.Sel("baseWritingDirection"))
	return rv
}/* debug [instance_properties/getter]: baseWritingDirection */


// The base writing direction for the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/baseWritingDirection
func (m_ MutableParagraphStyle) SetBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBaseWritingDirection:"), value)
}/* debug [instance_properties/setter]: baseWritingDirection */


// A number used as the document’s default tab spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/defaultTabInterval
func (m_ MutableParagraphStyle) DefaultTabInterval() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("defaultTabInterval"))
	return rv
}/* debug [instance_properties/getter]: defaultTabInterval */


// A number used as the document’s default tab spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/defaultTabInterval
func (m_ MutableParagraphStyle) SetDefaultTabInterval(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultTabInterval:"), value)
}/* debug [instance_properties/setter]: defaultTabInterval */


// The indentation of the first line of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/firstLineHeadIndent
func (m_ MutableParagraphStyle) FirstLineHeadIndent() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("firstLineHeadIndent"))
	return rv
}/* debug [instance_properties/getter]: firstLineHeadIndent */


// The indentation of the first line of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/firstLineHeadIndent
func (m_ MutableParagraphStyle) SetFirstLineHeadIndent(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFirstLineHeadIndent:"), value)
}/* debug [instance_properties/setter]: firstLineHeadIndent */


// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headIndent
func (m_ MutableParagraphStyle) HeadIndent() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("headIndent"))
	return rv
}/* debug [instance_properties/getter]: headIndent */


// The indentation of the paragraph’s lines other than the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headIndent
func (m_ MutableParagraphStyle) SetHeadIndent(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeadIndent:"), value)
}/* debug [instance_properties/setter]: headIndent */


// The paragraph’s header level for HTML generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headerLevel
func (m_ MutableParagraphStyle) HeaderLevel() int {
	rv := objc.Send[int](m_.ID, objc.Sel("headerLevel"))
	return rv
}/* debug [instance_properties/getter]: headerLevel */


// The paragraph’s header level for HTML generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/headerLevel
func (m_ MutableParagraphStyle) SetHeaderLevel(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeaderLevel:"), value)
}/* debug [instance_properties/setter]: headerLevel */


// The paragraph’s threshold for hyphenation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/hyphenationFactor
func (m_ MutableParagraphStyle) HyphenationFactor() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("hyphenationFactor"))
	return rv
}/* debug [instance_properties/getter]: hyphenationFactor */


// The paragraph’s threshold for hyphenation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/hyphenationFactor
func (m_ MutableParagraphStyle) SetHyphenationFactor(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHyphenationFactor:"), value)
}/* debug [instance_properties/setter]: hyphenationFactor */


// The mode for breaking lines in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakMode
func (m_ MutableParagraphStyle) LineBreakMode() LineBreakMode {
	rv := objc.Send[LineBreakMode](m_.ID, objc.Sel("lineBreakMode"))
	return rv
}/* debug [instance_properties/getter]: lineBreakMode */


// The mode for breaking lines in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakMode
func (m_ MutableParagraphStyle) SetLineBreakMode(value LineBreakMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineBreakMode:"), value)
}/* debug [instance_properties/setter]: lineBreakMode */


// The strategies that the text system may use to break lines while laying out the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakStrategy
func (m_ MutableParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	rv := objc.Send[LineBreakStrategy](m_.ID, objc.Sel("lineBreakStrategy"))
	return rv
}/* debug [instance_properties/getter]: lineBreakStrategy */


// The strategies that the text system may use to break lines while laying out the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineBreakStrategy
func (m_ MutableParagraphStyle) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineBreakStrategy:"), value)
}/* debug [instance_properties/setter]: lineBreakStrategy */


// The line height multiple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineHeightMultiple
func (m_ MutableParagraphStyle) LineHeightMultiple() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineHeightMultiple"))
	return rv
}/* debug [instance_properties/getter]: lineHeightMultiple */


// The line height multiple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineHeightMultiple
func (m_ MutableParagraphStyle) SetLineHeightMultiple(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineHeightMultiple:"), value)
}/* debug [instance_properties/setter]: lineHeightMultiple */


// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineSpacing
func (m_ MutableParagraphStyle) LineSpacing() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineSpacing"))
	return rv
}/* debug [instance_properties/getter]: lineSpacing */


// The distance in points between the bottom of one line fragment and the top of the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/lineSpacing
func (m_ MutableParagraphStyle) SetLineSpacing(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineSpacing:"), value)
}/* debug [instance_properties/setter]: lineSpacing */


// The paragraph’s maximum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/maximumLineHeight
func (m_ MutableParagraphStyle) MaximumLineHeight() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("maximumLineHeight"))
	return rv
}/* debug [instance_properties/getter]: maximumLineHeight */


// The paragraph’s maximum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/maximumLineHeight
func (m_ MutableParagraphStyle) SetMaximumLineHeight(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumLineHeight:"), value)
}/* debug [instance_properties/setter]: maximumLineHeight */


// The paragraph’s minimum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/minimumLineHeight
func (m_ MutableParagraphStyle) MinimumLineHeight() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("minimumLineHeight"))
	return rv
}/* debug [instance_properties/getter]: minimumLineHeight */


// The paragraph’s minimum line height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/minimumLineHeight
func (m_ MutableParagraphStyle) SetMinimumLineHeight(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumLineHeight:"), value)
}/* debug [instance_properties/setter]: minimumLineHeight */


// The space after the end of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/paragraphSpacing
func (m_ MutableParagraphStyle) ParagraphSpacing() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("paragraphSpacing"))
	return rv
}/* debug [instance_properties/getter]: paragraphSpacing */


// The space after the end of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/paragraphSpacing
func (m_ MutableParagraphStyle) SetParagraphSpacing(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphSpacing:"), value)
}/* debug [instance_properties/setter]: paragraphSpacing */


// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/paragraphSpacingBefore
func (m_ MutableParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("paragraphSpacingBefore"))
	return rv
}/* debug [instance_properties/getter]: paragraphSpacingBefore */


// The distance between the paragraph’s top and the beginning of its text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/paragraphSpacingBefore
func (m_ MutableParagraphStyle) SetParagraphSpacingBefore(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParagraphSpacingBefore:"), value)
}/* debug [instance_properties/setter]: paragraphSpacingBefore */


// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tabStops
func (m_ MutableParagraphStyle) TabStops() []TextTab {
	rv := objc.Send[[]TextTab](m_.ID, objc.Sel("tabStops"))
	return rv
}/* debug [instance_properties/getter]: tabStops */


// The text tab objects that represent the paragraph’s tab stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tabStops
func (m_ MutableParagraphStyle) SetTabStops(value []TextTab) {
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
}/* debug [instance_properties/setter]: tabStops */


// The trailing indentation of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tailIndent
func (m_ MutableParagraphStyle) TailIndent() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("tailIndent"))
	return rv
}/* debug [instance_properties/getter]: tailIndent */


// The trailing indentation of the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tailIndent
func (m_ MutableParagraphStyle) SetTailIndent(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTailIndent:"), value)
}/* debug [instance_properties/setter]: tailIndent */


// The text blocks that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/textBlocks
func (m_ MutableParagraphStyle) TextBlocks() []TextBlock {
	rv := objc.Send[[]TextBlock](m_.ID, objc.Sel("textBlocks"))
	return rv
}/* debug [instance_properties/getter]: textBlocks */


// The text blocks that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/textBlocks
func (m_ MutableParagraphStyle) SetTextBlocks(value []TextBlock) {
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
}/* debug [instance_properties/setter]: textBlocks */


// The text lists that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/textLists
func (m_ MutableParagraphStyle) TextLists() []TextList {
	rv := objc.Send[[]TextList](m_.ID, objc.Sel("textLists"))
	return rv
}/* debug [instance_properties/getter]: textLists */


// The text lists that contain the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/textLists
func (m_ MutableParagraphStyle) SetTextLists(value []TextList) {
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
}/* debug [instance_properties/setter]: textLists */


// The threshold for using tightening as an alternative to truncation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tighteningFactorForTruncation
func (m_ MutableParagraphStyle) TighteningFactorForTruncation() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("tighteningFactorForTruncation"))
	return rv
}/* debug [instance_properties/getter]: tighteningFactorForTruncation */


// The threshold for using tightening as an alternative to truncation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/tighteningFactorForTruncation
func (m_ MutableParagraphStyle) SetTighteningFactorForTruncation(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTighteningFactorForTruncation:"), value)
}/* debug [instance_properties/setter]: tighteningFactorForTruncation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/usesDefaultHyphenation
func (m_ MutableParagraphStyle) UsesDefaultHyphenation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}/* debug [instance_properties/getter]: usesDefaultHyphenation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/usesDefaultHyphenation
func (m_ MutableParagraphStyle) SetUsesDefaultHyphenation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUsesDefaultHyphenation:"), value)
}/* debug [instance_properties/setter]: usesDefaultHyphenation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableParagraphStyle */



