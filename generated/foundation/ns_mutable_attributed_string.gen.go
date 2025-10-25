// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMutableAttributedString */


/* debug [class_header]: Header for NSMutableAttributedString */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableAttributedString */
// An interface definition for the [MutableAttributedString] class.
type IMutableAttributedString interface {
	IAttributedString
	
/* debug [class_interface_properties]: Properties for MutableAttributedString */
	// properties:
	MutableString() IMutableString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableAttributedString */
	// methods:
	AddAttributeValueRange(name AttributedStringKey, value objc.IObject, range_ objc.IObject /* cross-framework: Range */)
	AddAttributesRange(attrs IDictionary, range_ objc.IObject /* cross-framework: Range */)
	AppendAttributedString(attrString IAttributedString)
	AppendLocalizedFormat(format IAttributedString)
	ApplyFontTraitsRange(traitMask FontTraitMask /* not a class type */, range_ objc.IObject /* cross-framework: Range */)
	BeginEditing()
	DeleteCharactersInRange(range_ objc.IObject /* cross-framework: Range */)
	EndEditing()
	FixAttachmentAttributeInRange(range_ objc.IObject /* cross-framework: Range */)
	FixAttributesInRange(range_ objc.IObject /* cross-framework: Range */)
	FixFontAttributeInRange(range_ objc.IObject /* cross-framework: Range */)
	FixParagraphStyleAttributeInRange(range_ objc.IObject /* cross-framework: Range */)
	InsertAttributedStringAtIndex(attrString IAttributedString, loc uint)
	ReadFromURLOptionsDocumentAttributesError(url IURL, opts IDictionary, dict IDictionary, error_ IError) bool
	ReadFromDataOptionsDocumentAttributesError(data IData, opts IDictionary, dict IDictionary, error_ IError) bool
	RemoveAttributeRange(name AttributedStringKey, range_ objc.IObject /* cross-framework: Range */)
	ReplaceCharactersInRangeWithAttributedString(range_ objc.IObject /* cross-framework: Range */, attrString IAttributedString)
	ReplaceCharactersInRangeWithString(range_ objc.IObject /* cross-framework: Range */, str IString)
	SetAlignmentRange(alignment TextAlignment /* not a class type */, range_ objc.IObject /* cross-framework: Range */)
	SetAttributedString(attrString IAttributedString)
	SetAttributesRange(attrs IDictionary, range_ objc.IObject /* cross-framework: Range */)
	SetBaseWritingDirectionRange(writingDirection WritingDirection /* not a class type */, range_ objc.IObject /* cross-framework: Range */)
	SubscriptRange(range_ objc.IObject /* cross-framework: Range */)
	SuperscriptRange(range_ objc.IObject /* cross-framework: Range */)
	UnscriptRange(range_ objc.IObject /* cross-framework: Range */)
	UpdateAttachmentsFromPath(path IString)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableAttributedString */
// Alloc allocates a new instance without initialization.
func (mc _MutableAttributedStringClass) Alloc() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableAttributedString */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableAttributedString *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableAttributedString */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableAttributedString */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableAttributedString */

// Adds an attribute with the given name and value to the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/addAttribute(_:value:range:)
func (m_ MutableAttributedString) AddAttributeValueRange(name AttributedStringKey, value objc.IObject, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addAttribute:value:range:"), name, value, range_)
}/* debug [instance_methods/method]: AddAttributeValueRange */


// Adds the given collection of attributes to the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/addAttributes(_:range:)
func (m_ MutableAttributedString) AddAttributesRange(attrs IDictionary, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addAttributes:range:"), attrs, range_)
}/* debug [instance_methods/method]: AddAttributesRange */


// Adds the characters and attributes of a given attributed string to the end of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/append(_:)
func (m_ MutableAttributedString) AppendAttributedString(attrString IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendAttributedString:"), attrString)
}/* debug [instance_methods/method]: AppendAttributedString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/appendLocalizedFormat:
func (m_ MutableAttributedString) AppendLocalizedFormat(format IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendLocalizedFormat:"), format)
}/* debug [instance_methods/method]: AppendLocalizedFormat */


// Applies the specified font-related attributes to characters in the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/applyFontTraits(_:range:)
func (m_ MutableAttributedString) ApplyFontTraitsRange(traitMask FontTraitMask /* not a class type */, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("applyFontTraits:range:"), traitMask, range_)
}/* debug [instance_methods/method]: ApplyFontTraitsRange */


// Begins the buffering of changes to the string’s characters and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/beginEditing()
func (m_ MutableAttributedString) BeginEditing() {
	objc.Send[objc.ID](m_.ID, objc.Sel("beginEditing"))
}/* debug [instance_methods/method]: BeginEditing */


// Deletes the characters in the given range along with their associated attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/deleteCharacters(in:)
func (m_ MutableAttributedString) DeleteCharactersInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("deleteCharactersInRange:"), range_)
}/* debug [instance_methods/method]: DeleteCharactersInRange */


// Ends the buffering of changes to the string’s characters and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/endEditing()
func (m_ MutableAttributedString) EndEditing() {
	objc.Send[objc.ID](m_.ID, objc.Sel("endEditing"))
}/* debug [instance_methods/method]: EndEditing */


// Cleans up attachment attributes in the specified range and removes all attachment attributes assigned to characters except the designated attachment character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixAttachmentAttribute(in:)
func (m_ MutableAttributedString) FixAttachmentAttributeInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("fixAttachmentAttributeInRange:"), range_)
}/* debug [instance_methods/method]: FixAttachmentAttributeInRange */


// Cleans up font, paragraph style, and attachment attributes within the given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixAttributes(in:)
func (m_ MutableAttributedString) FixAttributesInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("fixAttributesInRange:"), range_)
}/* debug [instance_methods/method]: FixAttributesInRange */


// Fixes the font attribute in the specified range and assigns default fonts where appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixFontAttribute(in:)
func (m_ MutableAttributedString) FixFontAttributeInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("fixFontAttributeInRange:"), range_)
}/* debug [instance_methods/method]: FixFontAttributeInRange */


// Fixes the paragraph style attributes in the specified range and assigns a paragraph style to all characters in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixParagraphStyleAttribute(in:)
func (m_ MutableAttributedString) FixParagraphStyleAttributeInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("fixParagraphStyleAttributeInRange:"), range_)
}/* debug [instance_methods/method]: FixParagraphStyleAttributeInRange */


// Inserts the characters and attributes of the given attributed string into the receiver at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/insert(_:at:)
func (m_ MutableAttributedString) InsertAttributedStringAtIndex(attrString IAttributedString, loc uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertAttributedString:atIndex:"), attrString, loc)
}/* debug [instance_methods/method]: InsertAttributedStringAtIndex */


// Sets the contents of attributed string using the contents of the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/read(from:options:documentAttributes:)-54wth
func (m_ MutableAttributedString) ReadFromURLOptionsDocumentAttributesError(url IURL, opts IDictionary, dict IDictionary, error_ IError) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readFromURL:options:documentAttributes:error:"), url, opts, dict, error_)
	return rv
}/* debug [instance_methods/method]: ReadFromURLOptionsDocumentAttributesError */


// Sets the contents of the attributed string using the specified data object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/read(from:options:documentAttributes:)-5mbcx
func (m_ MutableAttributedString) ReadFromDataOptionsDocumentAttributesError(data IData, opts IDictionary, dict IDictionary, error_ IError) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readFromData:options:documentAttributes:error:"), data, opts, dict, error_)
	return rv
}/* debug [instance_methods/method]: ReadFromDataOptionsDocumentAttributesError */


// Removes the named attribute from the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/removeAttribute(_:range:)
func (m_ MutableAttributedString) RemoveAttributeRange(name AttributedStringKey, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAttribute:range:"), name, range_)
}/* debug [instance_methods/method]: RemoveAttributeRange */


// Replaces the characters and attributes in a given range with the characters and attributes of the given attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/replaceCharacters(in:with:)-1uaw7
func (m_ MutableAttributedString) ReplaceCharactersInRangeWithAttributedString(range_ objc.IObject /* cross-framework: Range */, attrString IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceCharactersInRange:withAttributedString:"), range_, attrString)
}/* debug [instance_methods/method]: ReplaceCharactersInRangeWithAttributedString */


// Replaces the characters in the given range with the characters of the given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/replaceCharacters(in:with:)-6oq9r
func (m_ MutableAttributedString) ReplaceCharactersInRangeWithString(range_ objc.IObject /* cross-framework: Range */, str IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, str)
}/* debug [instance_methods/method]: ReplaceCharactersInRangeWithString */


// Sets the alignment characteristic of the paragraph style attribute for the specified range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAlignment(_:range:)
func (m_ MutableAttributedString) SetAlignmentRange(alignment TextAlignment /* not a class type */, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlignment:range:"), alignment, range_)
}/* debug [instance_methods/method]: SetAlignmentRange */


// Replaces the receiver’s entire contents with the characters and attributes of the given attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAttributedString(_:)
func (m_ MutableAttributedString) SetAttributedString(attrString IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributedString:"), attrString)
}/* debug [instance_methods/method]: SetAttributedString */


// Sets the attributes for the characters in the specified range to the specified attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAttributes(_:range:)
func (m_ MutableAttributedString) SetAttributesRange(attrs IDictionary, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributes:range:"), attrs, range_)
}/* debug [instance_methods/method]: SetAttributesRange */


// Sets the base writing direction for the characters to the specified direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setBaseWritingDirection(_:range:)
func (m_ MutableAttributedString) SetBaseWritingDirectionRange(writingDirection WritingDirection /* not a class type */, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBaseWritingDirection:range:"), writingDirection, range_)
}/* debug [instance_methods/method]: SetBaseWritingDirectionRange */


// Decrements the value of the superscript attribute for characters in the specified range by one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/subscriptRange(_:)
func (m_ MutableAttributedString) SubscriptRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscriptRange:"), range_)
}/* debug [instance_methods/method]: SubscriptRange */


// Increments the value of the superscript attribute for characters in the specified range by one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/superscriptRange(_:)
func (m_ MutableAttributedString) SuperscriptRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("superscriptRange:"), range_)
}/* debug [instance_methods/method]: SuperscriptRange */


// Removes the superscript attribute from the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/unscriptRange(_:)
func (m_ MutableAttributedString) UnscriptRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unscriptRange:"), range_)
}/* debug [instance_methods/method]: UnscriptRange */


// Updates all attachments based on files contained in the RTFD file package at the specified file path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/updateAttachments(fromPath:)
func (m_ MutableAttributedString) UpdateAttachmentsFromPath(path IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("updateAttachmentsFromPath:"), path)
}/* debug [instance_methods/method]: UpdateAttachmentsFromPath */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableAttributedString */

// The character contents of the receiver as a mutable string object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/mutableString
func (m_ MutableAttributedString) MutableString() IMutableString {
	rv := objc.Send[MutableString](m_.ID, objc.Sel("mutableString"))
	return rv
}/* debug [instance_properties/getter]: mutableString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableAttributedString */


