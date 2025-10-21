// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AttributedString] class.
var (
	AttributedStringClass     _AttributedStringClass
	AttributedStringClassOnce sync.Once
)

func getAttributedStringClass() _AttributedStringClass {
	AttributedStringClassOnce.Do(func() {
		AttributedStringClass = _AttributedStringClass{objc.GetClass("NSAttributedString")}
	})
	return AttributedStringClass
}

type _AttributedStringClass struct {
	class objc.Class
}

// An interface definition for the [AttributedString] class.
type IAttributedString interface {
	objectivec.IObject
	AttributeAtIndexEffectiveRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer) objc.ID
	AttributeAtIndexLongestEffectiveRangeInRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer, rangeLimit Range) objc.ID
	AttributedSubstringFromRange(range_ Range) unsafe.Pointer
	AttributesAtIndexEffectiveRange(location uint, range_ unsafe.Pointer) unsafe.Pointer
	AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ unsafe.Pointer, rangeLimit Range) unsafe.Pointer
	BoundingRectWithSizeOptionsContext(size coregraphics.CGSize, options unsafe.Pointer, context unsafe.Pointer) coregraphics.CGRect
	ContainsAttachmentsInRange(range_ Range) bool
	DataFromRangeDocumentAttributesError(range_ Range, dict unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	DocFormatFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer
	DoubleClickAtIndex(location uint) Range
	DrawAtPoint(point coregraphics.CGPoint)
	DrawInRect(rect coregraphics.CGRect)
	DrawWithRectOptionsContext(rect coregraphics.CGRect, options unsafe.Pointer, context unsafe.Pointer)
	EnumerateAttributeInRangeOptionsUsingBlock(attrName unsafe.Pointer, enumerationRange Range, opts unsafe.Pointer, block unsafe.Pointer)
	EnumerateAttributesInRangeOptionsUsingBlock(enumerationRange Range, opts unsafe.Pointer, block unsafe.Pointer)
	FileWrapperFromRangeDocumentAttributesError(range_ Range, dict unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	FontAttributesInRange(range_ Range) unsafe.Pointer
	AttributedStringByInflectingString() unsafe.Pointer
	IsEqualToAttributedString(other unsafe.Pointer) bool
	ItemNumberInTextListAtIndex(list unsafe.Pointer, location uint) int
	LineBreakBeforeIndexWithinRange(location uint, aRange Range) uint
	LineBreakByHyphenatingBeforeIndexWithinRange(location uint, aRange Range) uint
	NextWordFromIndexForward(location uint, isForward bool) uint
	PrefersRTFDInRange(range_ Range) bool
	RangeOfTextBlockAtIndex(block unsafe.Pointer, location uint) Range
	RangeOfTextTableAtIndex(table unsafe.Pointer, location uint) Range
	RangeOfTextListAtIndex(list unsafe.Pointer, location uint) Range
	RTFFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer
	RTFDFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer
	RTFDFileWrapperFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer
	RulerAttributesInRange(range_ Range) unsafe.Pointer
	Size() coregraphics.CGSize
}

// A string of text that manages data, layout, and stylistic information for ranges of characters to support rendering.
//
// is a type you use to manage strings of stylized Unicode text. In addition to text, an attributed string contains key-value pairs known as that specify additional information to apply to ranges of characters within the string. Attributed strings support many different kinds of attributes, including: Rendering attributes that specify font, color, kern, ligature, and other details Attributes for attachments and adaptive image glyphs Semantic attributes such as link URLs or tool-tip information Language attributes to support automatic gender agreement and text layout Accessibility attributes that provide information for assistive technologies Attributes that summarize details of the Markdown import process Custom attributes you define for your app Use attributed strings anywhere you need styled text, or when you need to associate additional information with your text. Because is an immutable type, you specify all of the text and attributes for it at creation time and can’t change them later. You can create attributed strings directly from a string of characters and a dictionary of attributes. You can also create attributed strings from the contents of a file, including files that contain RTF, RTFD, HTML, Markdown, or other file formats. If you need to modify the contents of an attributed string later, use the type instead. If you create an without any font information, the string’s default font is Helvetica 12-point, which might differ from the default system font for the platform. To change the font, specify a font attribute at creation time.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString
type AttributedString struct {
	objectivec.Object
}

// AttributedStringFrom constructs a [AttributedString] from an unsafe.Pointer.
//
// A string of text that manages data, layout, and stylistic information for ranges of characters to support rendering.
func AttributedStringFrom(ptr unsafe.Pointer) AttributedString {
	return AttributedString{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AttributedStringClass) Alloc() AttributedString {
	rv := objc.Send[AttributedString](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AttributedStringClass) New() AttributedString {
	rv := objc.Send[AttributedString](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttributedString) Init() AttributedString {
	rv := objc.Send[AttributedString](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttributedString) Autorelease() AttributedString {
	rv := objc.Send[AttributedString](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttributedString creates a new AttributedString instance.
func NewAttributedString() AttributedString {
	return getAttributedStringClass().New()
}

// An array of UTI strings that identify the file types that attributed strings support, either directly or through a user-installed filter service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textTypes
func (ac _AttributedStringClass) TextTypes() []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("textTypes"))
	return rv
}

// An array of UTI strings that identify the file types that attributed strings support directly.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textUnfilteredTypes
func (ac _AttributedStringClass) TextUnfilteredTypes() []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("textUnfilteredTypes"))
	return rv
}

// Returns the value for an attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attribute(_:at:effectiveRange:)
func (a_ AttributedString) AttributeAtIndexEffectiveRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("attribute:atIndex:effectiveRange:"), attrName, location, range_)
	return rv
}

// Returns the value for the attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attribute(_:at:longestEffectiveRange:in:)
func (a_ AttributedString) AttributeAtIndexLongestEffectiveRangeInRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer, rangeLimit Range) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("attribute:atIndex:longestEffectiveRange:inRange:"), attrName, location, range_, rangeLimit)
	return rv
}

// Returns an attributed string consisting of the characters and attributes within the specified range in the attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributedSubstring(from:)
func (a_ AttributedString) AttributedSubstringFromRange(range_ Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributedSubstringFromRange:"), range_)
	return rv
}

// Returns the attributes for the character at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributes(at:effectiveRange:)
func (a_ AttributedString) AttributesAtIndexEffectiveRange(location uint, range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributesAtIndex:effectiveRange:"), location, range_)
	return rv
}

// Returns the attributes for the character at the specified index and, by reference, the range where the attributes apply.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributes(at:longestEffectiveRange:in:)
func (a_ AttributedString) AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ unsafe.Pointer, rangeLimit Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributesAtIndex:longestEffectiveRange:inRange:"), location, range_, rangeLimit)
	return rv
}

// Returns the bounding rectangle necessary to draw the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/boundingRect(with:options:context:)
func (a_ AttributedString) BoundingRectWithSizeOptionsContext(size coregraphics.CGSize, options unsafe.Pointer, context unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](a_.ID, objc.Sel("boundingRectWithSize:options:context:"), size, options, context)
	return rv
}

// Returns a Boolean value that indicates if the attributed string contains an attachment in the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/containsAttachments(in:)
func (a_ AttributedString) ContainsAttachmentsInRange(range_ Range) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("containsAttachmentsInRange:"), range_)
	return rv
}

// Returns a data object that contains a text stream corresponding to the characters and attributes within the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/data(from:documentAttributes:)
func (a_ AttributedString) DataFromRangeDocumentAttributesError(range_ Range, dict unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("dataFromRange:documentAttributes:error:"), range_, dict, error_)
	return rv
}

// Returns a data object that contains a Microsoft Word–format stream corresponding to the characters and attributes within the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/docFormat(from:documentAttributes:)
func (a_ AttributedString) DocFormatFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("docFormatFromRange:documentAttributes:"), range_, dict)
	return rv
}

// Returns the range of characters that form a word (or other linguistic unit) surrounding the specified index, taking language characteristics into account.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/doubleClick(at:)
func (a_ AttributedString) DoubleClickAtIndex(location uint) Range {
	rv := objc.Send[Range](a_.ID, objc.Sel("doubleClickAtIndex:"), location)
	return rv
}

// Draws the attributed string starting at the specified point in the current graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(at:)
func (a_ AttributedString) DrawAtPoint(point coregraphics.CGPoint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawAtPoint:"), point)
}

// Draws the attributed string inside the specified bounding rectangle in the current graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(in:)
func (a_ AttributedString) DrawInRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawInRect:"), rect)
}

// Draws the attributed string in the specified bounding rectangle using the provided options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(with:options:context:)
func (a_ AttributedString) DrawWithRectOptionsContext(rect coregraphics.CGRect, options unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawWithRect:options:context:"), rect, options, context)
}

// Executes the specified closure or block for each range of a particular attribute in the attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/enumerateAttribute(_:in:options:using:)
func (a_ AttributedString) EnumerateAttributeInRangeOptionsUsingBlock(attrName unsafe.Pointer, enumerationRange Range, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("enumerateAttribute:inRange:options:usingBlock:"), attrName, enumerationRange, opts, block)
}

// Executes the specified closure or block for each range of attributes in the attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/enumerateAttributes(in:options:using:)
func (a_ AttributedString) EnumerateAttributesInRangeOptionsUsingBlock(enumerationRange Range, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("enumerateAttributesInRange:options:usingBlock:"), enumerationRange, opts, block)
}

// Returns a file wrapper object that contains a text stream corresponding to the characters and attributes within the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/fileWrapper(from:documentAttributes:)
func (a_ AttributedString) FileWrapperFromRangeDocumentAttributesError(range_ Range, dict unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("fileWrapperFromRange:documentAttributes:error:"), range_, dict, error_)
	return rv
}

// Returns the font attributes in effect for the character at the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/fontAttributes(in:)
func (a_ AttributedString) FontAttributesInRange(range_ Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("fontAttributesInRange:"), range_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/inflecting()
func (a_ AttributedString) AttributedStringByInflectingString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributedStringByInflectingString"))
	return rv
}

// Returns a Boolean value that indicates whether the attributed string is equal to the specified string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/isEqual(to:)
func (a_ AttributedString) IsEqualToAttributedString(other unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEqualToAttributedString:"), other)
	return rv
}

// Returns the index of the item at the specified location within the list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/itemNumber(in:at:)
func (a_ AttributedString) ItemNumberInTextListAtIndex(list unsafe.Pointer, location uint) int {
	rv := objc.Send[int](a_.ID, objc.Sel("itemNumberInTextList:atIndex:"), list, location)
	return rv
}

// Returns the appropriate line break when the character at the index doesn’t fit on the same line as the character at the beginning of the range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/lineBreak(before:within:)
func (a_ AttributedString) LineBreakBeforeIndexWithinRange(location uint, aRange Range) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("lineBreakBeforeIndex:withinRange:"), location, aRange)
	return rv
}

// Returns the index of the closest character before the specified index, and within the specified range, that can fit on a new line by hyphenating.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/lineBreakByHyphenating(before:within:)
func (a_ AttributedString) LineBreakByHyphenatingBeforeIndexWithinRange(location uint, aRange Range) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("lineBreakByHyphenatingBeforeIndex:withinRange:"), location, aRange)
	return rv
}

// Returns the index of the first character of the word after or before the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/nextWord(from:forward:)
func (a_ AttributedString) NextWordFromIndexForward(location uint, isForward bool) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("nextWordFromIndex:forward:"), location, isForward)
	return rv
}

// Returns a Boolean value that indicates whether the specified range of text prefers RTFD formatting.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/prefersRTFD(in:)
func (a_ AttributedString) PrefersRTFDInRange(range_ Range) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prefersRTFDInRange:"), range_)
	return rv
}

// Returns the range of the individual text block that contains the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-1wrcp
func (a_ AttributedString) RangeOfTextBlockAtIndex(block unsafe.Pointer, location uint) Range {
	rv := objc.Send[Range](a_.ID, objc.Sel("rangeOfTextBlock:atIndex:"), block, location)
	return rv
}

// Returns the range of the specified text table that contains the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-3fevu
func (a_ AttributedString) RangeOfTextTableAtIndex(table unsafe.Pointer, location uint) Range {
	rv := objc.Send[Range](a_.ID, objc.Sel("rangeOfTextTable:atIndex:"), table, location)
	return rv
}

// Returns the range of the specified text list that contains the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-6um0x
func (a_ AttributedString) RangeOfTextListAtIndex(list unsafe.Pointer, location uint) Range {
	rv := objc.Send[Range](a_.ID, objc.Sel("rangeOfTextList:atIndex:"), list, location)
	return rv
}

// Returns a data object that contains an RTF stream corresponding to the characters and attributes within the specified range, omitting all attachment attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtf(from:documentAttributes:)
func (a_ AttributedString) RTFFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("RTFFromRange:documentAttributes:"), range_, dict)
	return rv
}

// Returns a data object that contains an RTFD stream corresponding to the characters and attributes within the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtfd(from:documentAttributes:)
func (a_ AttributedString) RTFDFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("RTFDFromRange:documentAttributes:"), range_, dict)
	return rv
}

// Returns a file wrapper object that contains an RTFD document corresponding to the characters and attributes within the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtfdFileWrapper(from:documentAttributes:)
func (a_ AttributedString) RTFDFileWrapperFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("RTFDFileWrapperFromRange:documentAttributes:"), range_, dict)
	return rv
}

// Returns the ruler (paragraph) attributes in effect for the characters within the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rulerAttributes(in:)
func (a_ AttributedString) RulerAttributesInRange(range_ Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("rulerAttributesInRange:"), range_)
	return rv
}

// Returns the size necessary to draw the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/size()
func (a_ AttributedString) Size() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](a_.ID, objc.Sel("size"))
	return rv
}

// The length of the attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/length
func (a_ AttributedString) Length() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("length"))
	return rv
}

// The character contents of the attributed string as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (a_ AttributedString) String() string {
	rv := objc.Send[string](a_.ID, objc.Sel("string"))
	return rv
}

// An array of UTI strings that identify the file types that attributed strings support, either directly or through a user-installed filter service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textTypes
func (a_ AttributedString) TextTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("textTypes"))
	return rv
}

// An array of UTI strings that identify the file types that attributed strings support directly.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textUnfilteredTypes
func (a_ AttributedString) TextUnfilteredTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("textUnfilteredTypes"))
	return rv
}
