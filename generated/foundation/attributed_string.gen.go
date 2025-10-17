// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AttributedString] class.
var AttributedStringClass objc.Class

func init() {
	AttributedStringClass = objc.GetClass("NSAttributedString")
}

type AttributedString struct {
	objc.ID
}

func AttributedStringFrom(ptr unsafe.Pointer) AttributedString {
	return AttributedString{
		ID: objc.ID(ptr),
	}
}


// Returns the value for an attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/attribute(_:at:effectiveRange:)
func (a_ AttributedString) AttributeAtIndexEffectiveRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("attribute:atIndex:effectiveRange:")
	ret := a_.ID.Send(sel, attrName, location, range_)
	return ret
}
// Returns the value for the attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/attribute(_:at:longestEffectiveRange:in:)
func (a_ AttributedString) AttributeAtIndexLongestEffectiveRangeInRange(attrName unsafe.Pointer, location uint, range_ unsafe.Pointer, rangeLimit Range) objc.ID {
	sel := objc.RegisterName("attribute:atIndex:longestEffectiveRange:inRange:")
	ret := a_.ID.Send(sel, attrName, location, range_, rangeLimit)
	return ret
}
// Returns an attributed string consisting of the characters and attributes within the specified range in the attributed string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/attributedSubstring(from:)
func (a_ AttributedString) AttributedSubstringFromRange(range_ Range) unsafe.Pointer {
	sel := objc.RegisterName("attributedSubstringFromRange:")
	ret := a_.ID.Send(sel, range_)
	return unsafe.Pointer(ret)
}
// Returns the attributes for the character at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/attributes(at:effectiveRange:)
func (a_ AttributedString) AttributesAtIndexEffectiveRange(location uint, range_ unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("attributesAtIndex:effectiveRange:")
	ret := a_.ID.Send(sel, location, range_)
	return unsafe.Pointer(ret)
}
// Returns the attributes for the character at the specified index and, by reference, the range where the attributes apply. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/attributes(at:longestEffectiveRange:in:)
func (a_ AttributedString) AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ unsafe.Pointer, rangeLimit Range) unsafe.Pointer {
	sel := objc.RegisterName("attributesAtIndex:longestEffectiveRange:inRange:")
	ret := a_.ID.Send(sel, location, range_, rangeLimit)
	return unsafe.Pointer(ret)
}
// Returns the bounding rectangle necessary to draw the string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/boundingRect(with:options:context:)
func (a_ AttributedString) BoundingRectWithSizeOptionsContext(size Size, options unsafe.Pointer, context unsafe.Pointer) Rect {
	sel := objc.RegisterName("boundingRectWithSize:options:context:")
	ret := a_.ID.Send(sel, size, options, context)
	return Rect(ret)
}
// Returns a Boolean value that indicates if the attributed string contains an attachment in the specified range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/containsAttachments(in:)
func (a_ AttributedString) ContainsAttachmentsInRange(range_ Range) bool {
	sel := objc.RegisterName("containsAttachmentsInRange:")
	ret := a_.ID.Send(sel, range_)
	return ret != 0
}
// Returns a data object that contains a text stream corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/data(from:documentAttributes:)
func (a_ AttributedString) DataFromRangeDocumentAttributesError(range_ Range, dict unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataFromRange:documentAttributes:error:")
	ret := a_.ID.Send(sel, range_, dict, error)
	return unsafe.Pointer(ret)
}
// Returns a data object that contains a Microsoft Word–format stream corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/docFormat(from:documentAttributes:)
func (a_ AttributedString) DocFormatFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("docFormatFromRange:documentAttributes:")
	ret := a_.ID.Send(sel, range_, dict)
	return unsafe.Pointer(ret)
}
// Returns the range of characters that form a word (or other linguistic unit) surrounding the specified index, taking language characteristics into account. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/doubleClick(at:)
func (a_ AttributedString) DoubleClickAtIndex(location uint) Range {
	sel := objc.RegisterName("doubleClickAtIndex:")
	ret := a_.ID.Send(sel, location)
	return Range(ret)
}
// Draws the attributed string starting at the specified point in the current graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/draw(at:)
func (a_ AttributedString) DrawAtPoint(point Point) {
	sel := objc.RegisterName("drawAtPoint:")
	a_.ID.Send(sel, point)
}
// Draws the attributed string inside the specified bounding rectangle in the current graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/draw(in:)
func (a_ AttributedString) DrawInRect(rect Rect) {
	sel := objc.RegisterName("drawInRect:")
	a_.ID.Send(sel, rect)
}
// Draws the attributed string in the specified bounding rectangle using the provided options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/draw(with:options:context:)
func (a_ AttributedString) DrawWithRectOptionsContext(rect Rect, options unsafe.Pointer, context unsafe.Pointer) {
	sel := objc.RegisterName("drawWithRect:options:context:")
	a_.ID.Send(sel, rect, options, context)
}
// Executes the specified closure or block for each range of a particular attribute in the attributed string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/enumerateAttribute(_:in:options:using:)
func (a_ AttributedString) EnumerateAttributeInRangeOptionsUsingBlock(attrName unsafe.Pointer, enumerationRange Range, opts unsafe.Pointer, block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateAttribute:inRange:options:usingBlock:")
	a_.ID.Send(sel, attrName, enumerationRange, opts, block)
}
// Executes the specified closure or block for each range of attributes in the attributed string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/enumerateAttributes(in:options:using:)
func (a_ AttributedString) EnumerateAttributesInRangeOptionsUsingBlock(enumerationRange Range, opts unsafe.Pointer, block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateAttributesInRange:options:usingBlock:")
	a_.ID.Send(sel, enumerationRange, opts, block)
}
// Returns a file wrapper object that contains a text stream corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/fileWrapper(from:documentAttributes:)
func (a_ AttributedString) FileWrapperFromRangeDocumentAttributesError(range_ Range, dict unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fileWrapperFromRange:documentAttributes:error:")
	ret := a_.ID.Send(sel, range_, dict, error)
	return unsafe.Pointer(ret)
}
// Returns the font attributes in effect for the character at the specified location. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/fontAttributes(in:)
func (a_ AttributedString) FontAttributesInRange(range_ Range) unsafe.Pointer {
	sel := objc.RegisterName("fontAttributesInRange:")
	ret := a_.ID.Send(sel, range_)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/inflecting()
func (a_ AttributedString) AttributedStringByInflectingString() unsafe.Pointer {
	sel := objc.RegisterName("attributedStringByInflectingString")
	ret := a_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value that indicates whether the attributed string is equal to the specified string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/isEqual(to:)
func (a_ AttributedString) IsEqualToAttributedString(other unsafe.Pointer) bool {
	sel := objc.RegisterName("isEqualToAttributedString:")
	ret := a_.ID.Send(sel, other)
	return ret != 0
}
// Returns the index of the item at the specified location within the list. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/itemNumber(in:at:)
func (a_ AttributedString) ItemNumberInTextListAtIndex(list unsafe.Pointer, location uint) int {
	sel := objc.RegisterName("itemNumberInTextList:atIndex:")
	ret := a_.ID.Send(sel, list, location)
	return int(ret)
}
// Returns the appropriate line break when the character at the index doesn’t fit on the same line as the character at the beginning of the range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/lineBreak(before:within:)
func (a_ AttributedString) LineBreakBeforeIndexWithinRange(location uint, aRange Range) uint {
	sel := objc.RegisterName("lineBreakBeforeIndex:withinRange:")
	ret := a_.ID.Send(sel, location, aRange)
	return uint(ret)
}
// Returns the index of the closest character before the specified index, and within the specified range, that can fit on a new line by hyphenating. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/lineBreakByHyphenating(before:within:)
func (a_ AttributedString) LineBreakByHyphenatingBeforeIndexWithinRange(location uint, aRange Range) uint {
	sel := objc.RegisterName("lineBreakByHyphenatingBeforeIndex:withinRange:")
	ret := a_.ID.Send(sel, location, aRange)
	return uint(ret)
}
// Returns the index of the first character of the word after or before the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/nextWord(from:forward:)
func (a_ AttributedString) NextWordFromIndexForward(location uint, isForward bool) uint {
	sel := objc.RegisterName("nextWordFromIndex:forward:")
	ret := a_.ID.Send(sel, location, isForward)
	return uint(ret)
}
// Returns a Boolean value that indicates whether the specified range of text prefers RTFD formatting. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/prefersRTFD(in:)
func (a_ AttributedString) PrefersRTFDInRange(range_ Range) bool {
	sel := objc.RegisterName("prefersRTFDInRange:")
	ret := a_.ID.Send(sel, range_)
	return ret != 0
}
// Returns the range of the individual text block that contains the specified location. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/range(of:at:)-1wrcp
func (a_ AttributedString) RangeOfTextBlockAtIndex(block unsafe.Pointer, location uint) Range {
	sel := objc.RegisterName("rangeOfTextBlock:atIndex:")
	ret := a_.ID.Send(sel, block, location)
	return Range(ret)
}
// Returns the range of the specified text table that contains the specified location. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/range(of:at:)-3fevu
func (a_ AttributedString) RangeOfTextTableAtIndex(table unsafe.Pointer, location uint) Range {
	sel := objc.RegisterName("rangeOfTextTable:atIndex:")
	ret := a_.ID.Send(sel, table, location)
	return Range(ret)
}
// Returns the range of the specified text list that contains the specified location. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/range(of:at:)-6um0x
func (a_ AttributedString) RangeOfTextListAtIndex(list unsafe.Pointer, location uint) Range {
	sel := objc.RegisterName("rangeOfTextList:atIndex:")
	ret := a_.ID.Send(sel, list, location)
	return Range(ret)
}
// Returns a data object that contains an RTF stream corresponding to the characters and attributes within the specified range, omitting all attachment attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/rtf(from:documentAttributes:)
func (a_ AttributedString) RTFFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("RTFFromRange:documentAttributes:")
	ret := a_.ID.Send(sel, range_, dict)
	return unsafe.Pointer(ret)
}
// Returns a data object that contains an RTFD stream corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/rtfd(from:documentAttributes:)
func (a_ AttributedString) RTFDFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("RTFDFromRange:documentAttributes:")
	ret := a_.ID.Send(sel, range_, dict)
	return unsafe.Pointer(ret)
}
// Returns a file wrapper object that contains an RTFD document corresponding to the characters and attributes within the specified range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/rtfdFileWrapper(from:documentAttributes:)
func (a_ AttributedString) RTFDFileWrapperFromRangeDocumentAttributes(range_ Range, dict unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("RTFDFileWrapperFromRange:documentAttributes:")
	ret := a_.ID.Send(sel, range_, dict)
	return unsafe.Pointer(ret)
}
// Returns the ruler (paragraph) attributes in effect for the characters within the specified range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/rulerAttributes(in:)
func (a_ AttributedString) RulerAttributesInRange(range_ Range) unsafe.Pointer {
	sel := objc.RegisterName("rulerAttributesInRange:")
	ret := a_.ID.Send(sel, range_)
	return unsafe.Pointer(ret)
}
// Returns the size necessary to draw the string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAttributedString/size()
func (a_ AttributedString) Size() Size {
	sel := objc.RegisterName("size")
	ret := a_.ID.Send(sel)
	return Size(ret)
}

