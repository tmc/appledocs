// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AttributedString] class.
var (
	attributedStringClass     _AttributedStringClass
	attributedStringClassOnce sync.Once
)

func getAttributedStringClass() _AttributedStringClass {
	attributedStringClassOnce.Do(func() {
		attributedStringClass = _AttributedStringClass{objc.GetClass("NSAttributedString")}
	})
	return attributedStringClass
}

type _AttributedStringClass struct {
	class objc.Class
}

// An interface definition for the [AttributedString] class.
type IAttributedString interface {
	objectivec.IObject
	AttributeAtIndexEffectiveRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer) objc.ID
	AttributeAtIndexLongestEffectiveRangeInRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer, rangeLimit unsafe.Pointer) objc.ID
	AttributedSubstringFromRange(range_ unsafe.Pointer) unsafe.Pointer
	AttributesAtIndexEffectiveRange(location uint, range_ unsafe.Pointer) unsafe.Pointer
	AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ unsafe.Pointer, rangeLimit unsafe.Pointer) unsafe.Pointer
	BoundingRectWithSizeOptionsContext(size unsafe.Pointer, options unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer
	ContainsAttachmentsInRange(range_ unsafe.Pointer) bool
	DataFromRangeDocumentAttributesError(range_ unsafe.Pointer, dict unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	DocFormatFromRangeDocumentAttributes(range_ unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer
	DoubleClickAtIndex(location uint) unsafe.Pointer
	DrawAtPoint(point unsafe.Pointer)
	DrawInRect(rect unsafe.Pointer)
	DrawWithRectOptionsContext(rect unsafe.Pointer, options unsafe.Pointer, context unsafe.Pointer)
	EnumerateAttributeInRangeOptionsUsingBlock(attrName unsafe.Pointer, enumerationRange unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer)
	EnumerateAttributesInRangeOptionsUsingBlock(enumerationRange unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer)
	FileWrapperFromRangeDocumentAttributesError(range_ unsafe.Pointer, dict unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	FontAttributesInRange(range_ unsafe.Pointer) unsafe.Pointer
	AttributedStringByInflectingString() unsafe.Pointer
	IsEqualToAttributedString(other unsafe.Pointer) bool
	ItemNumberInTextListAtIndex(list unsafe.Pointer, location uint) int
	LineBreakBeforeIndexWithinRange(location uint, aRange unsafe.Pointer) uint
	LineBreakByHyphenatingBeforeIndexWithinRange(location uint, aRange unsafe.Pointer) uint
	NextWordFromIndexForward(location uint, isForward bool) uint
	PrefersRTFDInRange(range_ unsafe.Pointer) bool
	RangeOfTextBlockAtIndex(block unsafe.Pointer, location uint) unsafe.Pointer
	RangeOfTextTableAtIndex(table unsafe.Pointer, location uint) unsafe.Pointer
	RangeOfTextListAtIndex(list unsafe.Pointer, location uint) unsafe.Pointer
	RTFFromRangeDocumentAttributes(range_ unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer
	RTFDFromRangeDocumentAttributes(range_ unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer
	RTFDFileWrapperFromRangeDocumentAttributes(range_ unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer
	RulerAttributesInRange(range_ unsafe.Pointer) unsafe.Pointer
	Size() unsafe.Pointer
}

// A string of text that manages data, layout, and stylistic information for ranges of characters to support rendering. [Full Topic]
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


// Returns the value for an attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attribute(_:at:effectiveRange:)
func (a_ AttributedString) AttributeAtIndexEffectiveRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("attribute:atIndex:effectiveRange:"), attrName, location, range_)
	return rv
}
// Returns the value for the attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attribute(_:at:longestEffectiveRange:in:)
func (a_ AttributedString) AttributeAtIndexLongestEffectiveRangeInRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer, rangeLimit unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("attribute:atIndex:longestEffectiveRange:inRange:"), attrName, location, range_, rangeLimit)
	return rv
}
// Returns an attributed string consisting of the characters and attributes within the specified range in the attributed string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributedSubstring(from:)
func (a_ AttributedString) AttributedSubstringFromRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributedSubstringFromRange:"), range_)
	return rv
}
// Returns the attributes for the character at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributes(at:effectiveRange:)
func (a_ AttributedString) AttributesAtIndexEffectiveRange(location uint, range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributesAtIndex:effectiveRange:"), location, range_)
	return rv
}
// Returns the attributes for the character at the specified index and, by reference, the range where the attributes apply. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributes(at:longestEffectiveRange:in:)
func (a_ AttributedString) AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ unsafe.Pointer, rangeLimit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributesAtIndex:longestEffectiveRange:inRange:"), location, range_, rangeLimit)
	return rv
}
// Returns the bounding rectangle necessary to draw the string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/boundingRect(with:options:context:)
func (a_ AttributedString) BoundingRectWithSizeOptionsContext(size unsafe.Pointer, options unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("boundingRectWithSize:options:context:"), size, options, context)
	return rv
}
// Returns a Boolean value that indicates if the attributed string contains an attachment in the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/containsAttachments(in:)
func (a_ AttributedString) ContainsAttachmentsInRange(range_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("containsAttachmentsInRange:"), range_)
	return rv
}
// Returns a data object that contains a text stream corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/data(from:documentAttributes:)
func (a_ AttributedString) DataFromRangeDocumentAttributesError(range_ unsafe.Pointer, dict unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("dataFromRange:documentAttributes:error:"), range_, dict, error)
	return rv
}
// Returns a data object that contains a Microsoft Word–format stream corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/docFormat(from:documentAttributes:)
func (a_ AttributedString) DocFormatFromRangeDocumentAttributes(range_ unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("docFormatFromRange:documentAttributes:"), range_, dict)
	return rv
}
// Returns the range of characters that form a word (or other linguistic unit) surrounding the specified index, taking language characteristics into account. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/doubleClick(at:)
func (a_ AttributedString) DoubleClickAtIndex(location uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("doubleClickAtIndex:"), location)
	return rv
}
// Draws the attributed string starting at the specified point in the current graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(at:)
func (a_ AttributedString) DrawAtPoint(point unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawAtPoint:"), point)
}
// Draws the attributed string inside the specified bounding rectangle in the current graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(in:)
func (a_ AttributedString) DrawInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawInRect:"), rect)
}
// Draws the attributed string in the specified bounding rectangle using the provided options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(with:options:context:)
func (a_ AttributedString) DrawWithRectOptionsContext(rect unsafe.Pointer, options unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawWithRect:options:context:"), rect, options, context)
}
// Executes the specified closure or block for each range of a particular attribute in the attributed string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/enumerateAttribute(_:in:options:using:)
func (a_ AttributedString) EnumerateAttributeInRangeOptionsUsingBlock(attrName unsafe.Pointer, enumerationRange unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("enumerateAttribute:inRange:options:usingBlock:"), attrName, enumerationRange, opts, block)
}
// Executes the specified closure or block for each range of attributes in the attributed string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/enumerateAttributes(in:options:using:)
func (a_ AttributedString) EnumerateAttributesInRangeOptionsUsingBlock(enumerationRange unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("enumerateAttributesInRange:options:usingBlock:"), enumerationRange, opts, block)
}
// Returns a file wrapper object that contains a text stream corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/fileWrapper(from:documentAttributes:)
func (a_ AttributedString) FileWrapperFromRangeDocumentAttributesError(range_ unsafe.Pointer, dict unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("fileWrapperFromRange:documentAttributes:error:"), range_, dict, error)
	return rv
}
// Returns the font attributes in effect for the character at the specified location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/fontAttributes(in:)
func (a_ AttributedString) FontAttributesInRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("fontAttributesInRange:"), range_)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/inflecting()
func (a_ AttributedString) AttributedStringByInflectingString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributedStringByInflectingString"))
	return rv
}
// Returns a Boolean value that indicates whether the attributed string is equal to the specified string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/isEqual(to:)
func (a_ AttributedString) IsEqualToAttributedString(other unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEqualToAttributedString:"), other)
	return rv
}
// Returns the index of the item at the specified location within the list. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/itemNumber(in:at:)
func (a_ AttributedString) ItemNumberInTextListAtIndex(list unsafe.Pointer, location uint) int {
	rv := objc.Send[int](a_.ID, objc.Sel("itemNumberInTextList:atIndex:"), list, location)
	return rv
}
// Returns the appropriate line break when the character at the index doesn’t fit on the same line as the character at the beginning of the range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/lineBreak(before:within:)
func (a_ AttributedString) LineBreakBeforeIndexWithinRange(location uint, aRange unsafe.Pointer) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("lineBreakBeforeIndex:withinRange:"), location, aRange)
	return rv
}
// Returns the index of the closest character before the specified index, and within the specified range, that can fit on a new line by hyphenating. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/lineBreakByHyphenating(before:within:)
func (a_ AttributedString) LineBreakByHyphenatingBeforeIndexWithinRange(location uint, aRange unsafe.Pointer) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("lineBreakByHyphenatingBeforeIndex:withinRange:"), location, aRange)
	return rv
}
// Returns the index of the first character of the word after or before the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/nextWord(from:forward:)
func (a_ AttributedString) NextWordFromIndexForward(location uint, isForward bool) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("nextWordFromIndex:forward:"), location, isForward)
	return rv
}
// Returns a Boolean value that indicates whether the specified range of text prefers RTFD formatting. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/prefersRTFD(in:)
func (a_ AttributedString) PrefersRTFDInRange(range_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prefersRTFDInRange:"), range_)
	return rv
}
// Returns the range of the individual text block that contains the specified location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-1wrcp
func (a_ AttributedString) RangeOfTextBlockAtIndex(block unsafe.Pointer, location uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("rangeOfTextBlock:atIndex:"), block, location)
	return rv
}
// Returns the range of the specified text table that contains the specified location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-3fevu
func (a_ AttributedString) RangeOfTextTableAtIndex(table unsafe.Pointer, location uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("rangeOfTextTable:atIndex:"), table, location)
	return rv
}
// Returns the range of the specified text list that contains the specified location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-6um0x
func (a_ AttributedString) RangeOfTextListAtIndex(list unsafe.Pointer, location uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("rangeOfTextList:atIndex:"), list, location)
	return rv
}
// Returns a data object that contains an RTF stream corresponding to the characters and attributes within the specified range, omitting all attachment attributes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtf(from:documentAttributes:)
func (a_ AttributedString) RTFFromRangeDocumentAttributes(range_ unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("RTFFromRange:documentAttributes:"), range_, dict)
	return rv
}
// Returns a data object that contains an RTFD stream corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtfd(from:documentAttributes:)
func (a_ AttributedString) RTFDFromRangeDocumentAttributes(range_ unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("RTFDFromRange:documentAttributes:"), range_, dict)
	return rv
}
// Returns a file wrapper object that contains an RTFD document corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtfdFileWrapper(from:documentAttributes:)
func (a_ AttributedString) RTFDFileWrapperFromRangeDocumentAttributes(range_ unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("RTFDFileWrapperFromRange:documentAttributes:"), range_, dict)
	return rv
}
// Returns the ruler (paragraph) attributes in effect for the characters within the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rulerAttributes(in:)
func (a_ AttributedString) RulerAttributesInRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("rulerAttributesInRange:"), range_)
	return rv
}
// Returns the size necessary to draw the string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/size()
func (a_ AttributedString) Size() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("size"))
	return rv
}


