// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MutableAttributedString] class.
var (
	MutableAttributedStringClass     _MutableAttributedStringClass
	MutableAttributedStringClassOnce sync.Once
)

func getMutableAttributedStringClass() _MutableAttributedStringClass {
	MutableAttributedStringClassOnce.Do(func() {
		MutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}
	})
	return MutableAttributedStringClass
}

type _MutableAttributedStringClass struct {
	class objc.Class
}

// An interface definition for the [MutableAttributedString] class.
type IMutableAttributedString interface {
	IAttributedString
	// properties:
	MutableString() MutableString /* not a class type */
	// methods:
	AddAttributeValueRange(name AttributedStringKey /* not a class type */, value objectivec.IObject, range_ Range /* not a class type */)
	AddAttributesRange(attrs IDictionary /* already interface */, range_ Range /* not a class type */)
	AppendAttributedString(attrString IAttributedString)
	AppendLocalizedFormat(format IAttributedString)
	ApplyFontTraitsRange(traitMask FontTraitMask /* not a class type */, range_ Range /* not a class type */)
	BeginEditing()
	DeleteCharactersInRange(range_ Range /* not a class type */)
	EndEditing()
	FixAttachmentAttributeInRange(range_ Range /* not a class type */)
	FixAttributesInRange(range_ Range /* not a class type */)
	FixFontAttributeInRange(range_ Range /* not a class type */)
	FixParagraphStyleAttributeInRange(range_ Range /* not a class type */)
	InsertAttributedStringAtIndex(attrString IAttributedString, loc uint /* primitive/slice/pointer */)
	ReadFromURLOptionsDocumentAttributesError(url IURL, opts IDictionary /* already interface */, dict IDictionary /* already interface */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	ReadFromDataOptionsDocumentAttributesError(data IData, opts IDictionary /* already interface */, dict IDictionary /* already interface */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	RemoveAttributeRange(name AttributedStringKey /* not a class type */, range_ Range /* not a class type */)
	ReplaceCharactersInRangeWithAttributedString(range_ Range /* not a class type */, attrString IAttributedString)
	ReplaceCharactersInRangeWithString(range_ Range /* not a class type */, str string /* primitive/slice/pointer */)
	SetAlignmentRange(alignment TextAlignment, range_ Range /* not a class type */)
	SetAttributedString(attrString IAttributedString)
	SetAttributesRange(attrs IDictionary /* already interface */, range_ Range /* not a class type */)
	SetBaseWritingDirectionRange(writingDirection WritingDirection, range_ Range /* not a class type */)
	SubscriptRange(range_ Range /* not a class type */)
	SuperscriptRange(range_ Range /* not a class type */)
	UnscriptRange(range_ Range /* not a class type */)
	UpdateAttachmentsFromPath(path string /* primitive/slice/pointer */)
}

// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text.
//
// The class declares additional methods for mutating the content of an attributed string. You can add and remove characters (raw strings) and attributes separately or together as attributed strings. See the class description for for more information about attributed strings. adds two primitive methods to those of . These primitive methods provide the basis for all the other methods in its class. The primitive method replaces a range of characters with those from a string, leaving all attribute information outside that range intact. The primitive method sets attributes and values for a given range of characters, replacing any previous attributes and values for that range. In macOS, AppKit also uses and its subclass to encapsulate the paragraph or ruler attributes used by the classes. Note that the default font for objects is Helvetica 12-point, which may differ from the macOS system font, so you may wish to create the string with non-default attributes suitable for your application using, for example, . is “toll-free bridged” with its Core Foundation counterpart, . See for more information.


// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString
type MutableAttributedString struct {
	AttributedString
}

// MutableAttributedStringFrom constructs a [MutableAttributedString] from an unsafe.Pointer.
//
// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text.
func MutableAttributedStringFrom(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{
		AttributedString: AttributedStringFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableAttributedStringClass) Alloc() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableAttributedStringClass) New() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableAttributedString) Init() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableAttributedString) Autorelease() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableAttributedString creates a new MutableAttributedString instance.
func NewMutableAttributedString() MutableAttributedString {
	return getMutableAttributedStringClass().New()
}



// Adds an attribute with the given name and value to the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/addAttribute(_:value:range:)
func (m_ MutableAttributedString) AddAttributeValueRange(name AttributedStringKey /* not a class type */, value objectivec.IObject, range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addAttribute:value:range:"), name, value, range_)
}


// Adds the given collection of attributes to the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/addAttributes(_:range:)
func (m_ MutableAttributedString) AddAttributesRange(attrs IDictionary /* already interface */, range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addAttributes:range:"), attrs, range_)
}


// Adds the characters and attributes of a given attributed string to the end of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/append(_:)
func (m_ MutableAttributedString) AppendAttributedString(attrString IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendAttributedString:"), attrString)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/appendLocalizedFormat:
func (m_ MutableAttributedString) AppendLocalizedFormat(format IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendLocalizedFormat:"), format)
}


// Applies the specified font-related attributes to characters in the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/applyFontTraits(_:range:)
func (m_ MutableAttributedString) ApplyFontTraitsRange(traitMask FontTraitMask /* not a class type */, range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("applyFontTraits:range:"), traitMask, range_)
}


// Begins the buffering of changes to the string’s characters and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/beginEditing()
func (m_ MutableAttributedString) BeginEditing() {
	objc.Send[objc.ID](m_.ID, objc.Sel("beginEditing"))
}


// Deletes the characters in the given range along with their associated attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/deleteCharacters(in:)
func (m_ MutableAttributedString) DeleteCharactersInRange(range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("deleteCharactersInRange:"), range_)
}


// Ends the buffering of changes to the string’s characters and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/endEditing()
func (m_ MutableAttributedString) EndEditing() {
	objc.Send[objc.ID](m_.ID, objc.Sel("endEditing"))
}


// Cleans up attachment attributes in the specified range and removes all attachment attributes assigned to characters except the designated attachment character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixAttachmentAttribute(in:)
func (m_ MutableAttributedString) FixAttachmentAttributeInRange(range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("fixAttachmentAttributeInRange:"), range_)
}


// Cleans up font, paragraph style, and attachment attributes within the given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixAttributes(in:)
func (m_ MutableAttributedString) FixAttributesInRange(range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("fixAttributesInRange:"), range_)
}


// Fixes the font attribute in the specified range and assigns default fonts where appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixFontAttribute(in:)
func (m_ MutableAttributedString) FixFontAttributeInRange(range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("fixFontAttributeInRange:"), range_)
}


// Fixes the paragraph style attributes in the specified range and assigns a paragraph style to all characters in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixParagraphStyleAttribute(in:)
func (m_ MutableAttributedString) FixParagraphStyleAttributeInRange(range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("fixParagraphStyleAttributeInRange:"), range_)
}


// Inserts the characters and attributes of the given attributed string into the receiver at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/insert(_:at:)
func (m_ MutableAttributedString) InsertAttributedStringAtIndex(attrString IAttributedString, loc uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertAttributedString:atIndex:"), attrString, loc)
}


// Sets the contents of attributed string using the contents of the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/read(from:options:documentAttributes:)-54wth
func (m_ MutableAttributedString) ReadFromURLOptionsDocumentAttributesError(url IURL, opts IDictionary /* already interface */, dict IDictionary /* already interface */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("readFromURL:options:documentAttributes:error:"), url, opts, dict, error_)
	return rv
}


// Sets the contents of the attributed string using the specified data object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/read(from:options:documentAttributes:)-5mbcx
func (m_ MutableAttributedString) ReadFromDataOptionsDocumentAttributesError(data IData, opts IDictionary /* already interface */, dict IDictionary /* already interface */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("readFromData:options:documentAttributes:error:"), data, opts, dict, error_)
	return rv
}


// Removes the named attribute from the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/removeAttribute(_:range:)
func (m_ MutableAttributedString) RemoveAttributeRange(name AttributedStringKey /* not a class type */, range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAttribute:range:"), name, range_)
}


// Replaces the characters and attributes in a given range with the characters and attributes of the given attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/replaceCharacters(in:with:)-1uaw7
func (m_ MutableAttributedString) ReplaceCharactersInRangeWithAttributedString(range_ Range /* not a class type */, attrString IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceCharactersInRange:withAttributedString:"), range_, attrString)
}


// Replaces the characters in the given range with the characters of the given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/replaceCharacters(in:with:)-6oq9r
func (m_ MutableAttributedString) ReplaceCharactersInRangeWithString(range_ Range /* not a class type */, str string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, objc.String(str))
}


// Sets the alignment characteristic of the paragraph style attribute for the specified range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAlignment(_:range:)
func (m_ MutableAttributedString) SetAlignmentRange(alignment TextAlignment, range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlignment:range:"), alignment, range_)
}


// Replaces the receiver’s entire contents with the characters and attributes of the given attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAttributedString(_:)
func (m_ MutableAttributedString) SetAttributedString(attrString IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributedString:"), attrString)
}


// Sets the attributes for the characters in the specified range to the specified attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAttributes(_:range:)
func (m_ MutableAttributedString) SetAttributesRange(attrs IDictionary /* already interface */, range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributes:range:"), attrs, range_)
}


// Sets the base writing direction for the characters to the specified direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setBaseWritingDirection(_:range:)
func (m_ MutableAttributedString) SetBaseWritingDirectionRange(writingDirection WritingDirection, range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBaseWritingDirection:range:"), writingDirection, range_)
}


// Decrements the value of the superscript attribute for characters in the specified range by one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/subscriptRange(_:)
func (m_ MutableAttributedString) SubscriptRange(range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscriptRange:"), range_)
}


// Increments the value of the superscript attribute for characters in the specified range by one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/superscriptRange(_:)
func (m_ MutableAttributedString) SuperscriptRange(range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("superscriptRange:"), range_)
}


// Removes the superscript attribute from the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/unscriptRange(_:)
func (m_ MutableAttributedString) UnscriptRange(range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unscriptRange:"), range_)
}


// Updates all attachments based on files contained in the RTFD file package at the specified file path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/updateAttachments(fromPath:)
func (m_ MutableAttributedString) UpdateAttachmentsFromPath(path string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("updateAttachmentsFromPath:"), objc.String(path))
}


// The character contents of the receiver as a mutable string object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/mutableString
func (m_ MutableAttributedString) MutableString() MutableString /* not a class type */ {
	rv := objc.Send[MutableString](m_.ID, objc.Sel("mutableString"))
	return rv
}



