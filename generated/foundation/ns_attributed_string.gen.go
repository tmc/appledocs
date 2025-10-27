// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	

	// properties:
	ContainsAttachments() bool
	Length() uint
	String() IString


	

	// methods:
	AttributeAtIndexEffectiveRange(attrName AttributedStringKey, location uint, range_ RangePointer) objc.ID
	AttributeAtIndexLongestEffectiveRangeInRange(attrName AttributedStringKey, location uint, range_ RangePointer, rangeLimit Range) objc.ID
	AttributedSubstringFromRange(range_ Range) IAttributedString
	AttributesAtIndexEffectiveRange(location uint, range_ RangePointer) IDictionary
	AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ RangePointer, rangeLimit Range) IDictionary
	BoundingRectWithSizeOptions(size corefoundation.CGSize, options StringDrawingOptions) corefoundation.CGRect
	BoundingRectWithSizeOptionsContext(size corefoundation.CGSize, options StringDrawingOptions, context objectivec.IObject) corefoundation.CGRect
	ContainsAttachmentsInRange(range_ Range) bool
	DataFromRangeDocumentAttributesError(range_ Range, dict IDictionary, error_ IError) IData
	DocFormatFromRangeDocumentAttributes(range_ Range, dict IDictionary) IData
	DoubleClickAtIndex(location uint) Range
	DrawAtPoint(point corefoundation.CGPoint)
	DrawInRect(rect corefoundation.CGRect)
	DrawWithRectOptions(rect corefoundation.CGRect, options StringDrawingOptions)
	DrawWithRectOptionsContext(rect corefoundation.CGRect, options StringDrawingOptions, context objectivec.IObject)
	EnumerateAttributeInRangeOptionsUsingBlock(attrName AttributedStringKey, enumerationRange Range, opts AttributedStringEnumerationOptions, block unsafe.Pointer)
	EnumerateAttributesInRangeOptionsUsingBlock(enumerationRange Range, opts AttributedStringEnumerationOptions, block unsafe.Pointer)
	FileWrapperFromRangeDocumentAttributesError(range_ Range, dict IDictionary, error_ IError) IFileWrapper
	FontAttributesInRange(range_ Range) IDictionary
	AttributedStringByInflectingString() IAttributedString
	IsEqualToAttributedString(other IAttributedString) bool
	ItemNumberInTextListAtIndex(list objectivec.IObject, location uint) int
	LineBreakBeforeIndexWithinRange(location uint, aRange Range) uint
	LineBreakByHyphenatingBeforeIndexWithinRange(location uint, aRange Range) uint
	NextWordFromIndexForward(location uint, isForward bool) uint
	PrefersRTFDInRange(range_ Range) bool
	RangeOfTextBlockAtIndex(block objectivec.IObject, location uint) Range
	RangeOfTextTableAtIndex(table objectivec.IObject, location uint) Range
	RangeOfTextListAtIndex(list objectivec.IObject, location uint) Range
	RTFFromRangeDocumentAttributes(range_ Range, dict IDictionary) IData
	RTFDFromRangeDocumentAttributes(range_ Range, dict IDictionary) IData
	RTFDFileWrapperFromRangeDocumentAttributes(range_ Range, dict IDictionary) IFileWrapper
	RulerAttributesInRange(range_ Range) IDictionary
	Size() corefoundation.CGSize


}





// Alloc allocates a new instance without initialization.
func (ac _AttributedStringClass) Alloc() AttributedString {
	rv := objc.Send[AttributedString](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A string of text that manages data, layout, and stylistic information for ranges of characters to support rendering.
//
// is a type you use to manage strings of stylized Unicode text. In addition to text, an attributed string contains key-value pairs known as that specify additional information to apply to ranges of characters within the string. Attributed strings support many different kinds of attributes, including: Rendering attributes that specify font, color, kern, ligature, and other details Attributes for attachments and adaptive image glyphs Semantic attributes such as link URLs or tool-tip information Language attributes to support automatic gender agreement and text layout Accessibility attributes that provide information for assistive technologies Attributes that summarize details of the Markdown import process Custom attributes you define for your app Use attributed strings anywhere you need styled text, or when you need to associate additional information with your text. Because is an immutable type, you specify all of the text and attributes for it at creation time and can’t change them later. You can create attributed strings directly from a string of characters and a dictionary of attributes. You can also create attributed strings from the contents of a file, including files that contain RTF, RTFD, HTML, Markdown, or other file formats. If you need to modify the contents of an attributed string later, use the type instead. If you create an without any font information, the string’s default font is Helvetica 12-point, which might differ from the default system font for the platform. To change the font, specify a font attribute at creation time.


// A string of text that manages data, layout, and stylistic information for ranges of characters to support rendering.
//
// [Full Topic]
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






// Creates an attributed string with an adaptive image glyph and applies the specified attributes to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(adaptiveImageGlyph:attributes:)
func NewAttributedStringWithAdaptiveImageGlyphAttributes(adaptiveImageGlyph objectivec.IObject, attributes IDictionary) AttributedString {
	rv := objc.Send[AttributedString](objc.ID(getAttributedStringClass().class), objc.Sel("attributedStringWithAdaptiveImageGlyph:attributes:"), adaptiveImageGlyph, attributes)
	return rv
}


// Creates an attributed string with an attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(attachment:)
func NewAttributedStringWithAttachment(attachment objectivec.IObject) AttributedString {
	rv := objc.Send[AttributedString](objc.ID(getAttributedStringClass().class), objc.Sel("attributedStringWithAttachment:"), attachment)
	return rv
}


// Creates an attributed string with an attachment and applies the specified attributes to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(attachment:attributes:)
func NewAttributedStringWithAttachmentAttributes(attachment objectivec.IObject, attributes IDictionary) AttributedString {
	rv := objc.Send[AttributedString](objc.ID(getAttributedStringClass().class), objc.Sel("attributedStringWithAttachment:attributes:"), attachment, attributes)
	return rv
}


// Creates a new attributed string from the contents of another attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(attributedString:)
func NewAttributedStringWithAttributedString(attrStr IAttributedString) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithAttributedString:"), attrStr)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from the contents of a specified URL that contains Markdown-formatted data using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithContentsOfMarkdownFileAtURL:options:baseURL:error:
func NewAttributedStringWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile IURL, options IAttributedStringMarkdownParsingOptions, baseURL IURL, error_ IError) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithContentsOfMarkdownFileAtURL:options:baseURL:error:"), markdownFile, options, baseURL, error_)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from the contents of the specified data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(data:options:documentAttributes:)
func NewAttributedStringWithDataOptionsDocumentAttributesError(data IData, options IDictionary, dict IDictionary, error_ IError) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithData:options:documentAttributes:error:"), data, options, dict, error_)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from Microsoft Word format data in the specified data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(docFormat:documentAttributes:)
func NewAttributedStringWithDocFormatDocumentAttributes(data IData, dict IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithDocFormat:documentAttributes:"), data, dict)
	rv.Autorelease()
	return rv
}


// Initializes a new attributed string object from the data at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(fileURL:options:documentAttributes:)
func NewAttributedStringWithFileURLOptionsDocumentAttributesError(url IURL, options IDictionary, dict IDictionary, error_ IError) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithFileURL:options:documentAttributes:error:"), url, options, dict, error_)
	rv.Autorelease()
	return rv
}


// Initializes an attributed string by substituting arguments into a specially formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:
func NewAttributedStringWithFormatOptionsLocale(format IAttributedString, options AttributedStringFormattingOptions, locale ILocale) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithFormat:options:locale:"), format, options, locale)
	rv.Autorelease()
	return rv
}


// Initializes an attributed string by substituting a list of function arguments into a specially formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:arguments:
func NewAttributedStringWithFormatOptionsLocaleArguments(format IAttributedString, options AttributedStringFormattingOptions, locale ILocale, arguments objectivec.IObject) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithFormat:options:locale:arguments:"), format, options, locale, arguments)
	rv.Autorelease()
	return rv
}


// Initializes an attributed string by substituting arguments into a specially formatted string and applying additional contextual information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:context:
func NewAttributedStringWithFormatOptionsLocaleContext(format IAttributedString, options AttributedStringFormattingOptions, locale ILocale, context IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithFormat:options:locale:context:"), format, options, locale, context)
	rv.Autorelease()
	return rv
}


// Initializes an attributed string by substituting a list of function arguments into a specially formatted string and applying additional contextual information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:context:arguments:
func NewAttributedStringWithFormatOptionsLocaleContextArguments(format IAttributedString, options AttributedStringFormattingOptions, locale ILocale, context IDictionary, arguments objectivec.IObject) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithFormat:options:locale:context:arguments:"), format, options, locale, context, arguments)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from the HTML in the specified data object and base URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(HTML:baseURL:documentAttributes:)
func NewAttributedStringWithHTMLBaseURLDocumentAttributes(data IData, base IURL, dict IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithHTML:baseURL:documentAttributes:"), data, base, dict)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from the HTML in the specified data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(HTML:documentAttributes:)
func NewAttributedStringWithHTMLDocumentAttributes(data IData, dict IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithHTML:documentAttributes:"), data, dict)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from the HTML in the specified data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(HTML:options:documentAttributes:)
func NewAttributedStringWithHTMLOptionsDocumentAttributes(data IData, options IDictionary, dict IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithHTML:options:documentAttributes:"), data, options, dict)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from Markdown-formatted data using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithMarkdown:options:baseURL:error:
func NewAttributedStringWithMarkdownOptionsBaseURLError(markdown IData, options IAttributedStringMarkdownParsingOptions, baseURL IURL, error_ IError) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithMarkdown:options:baseURL:error:"), markdown, options, baseURL, error_)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from a Markdown-formatted string using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithMarkdownString:options:baseURL:error:
func NewAttributedStringWithMarkdownStringOptionsBaseURLError(markdownString IString, options IAttributedStringMarkdownParsingOptions, baseURL IURL, error_ IError) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithMarkdownString:options:baseURL:error:"), markdownString, options, baseURL, error_)
	rv.Autorelease()
	return rv
}


// Initializes a new attribute string object from RTF or RTFD data in the file at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(path:documentAttributes:)
func NewAttributedStringWithPathDocumentAttributes(path IString, dict IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithPath:documentAttributes:"), path, dict)
	rv.Autorelease()
	return rv
}


// Creates an attributed string by decoding the stream of RTFD commands and data in the specified data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(RTFD:documentAttributes:)
func NewAttributedStringWithRTFDDocumentAttributes(data IData, dict IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithRTFD:documentAttributes:"), data, dict)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from the specified file wrapper that contains an RTFD document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(RTFDFileWrapper:documentAttributes:)
func NewAttributedStringWithRTFDFileWrapperDocumentAttributes(wrapper IFileWrapper, dict IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithRTFDFileWrapper:documentAttributes:"), wrapper, dict)
	rv.Autorelease()
	return rv
}


// Creates an attributed string by decoding the stream of RTF commands and data in the specified data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(RTF:documentAttributes:)
func NewAttributedStringWithRTFDocumentAttributes(data IData, dict IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithRTF:documentAttributes:"), data, dict)
	rv.Autorelease()
	return rv
}


// Creates an attributed string with the specified text and no attribute information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(string:)
func NewAttributedStringWithString(str IString) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithString:"), str)
	rv.Autorelease()
	return rv
}


// Creates an attributed string with the specified text and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(string:attributes:)
func NewAttributedStringWithStringAttributes(str IString, attrs IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithString:attributes:"), str, attrs)
	rv.Autorelease()
	return rv
}


// Initializes a new attributed string object from the data at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(URL:documentAttributes:)
func NewAttributedStringWithURLDocumentAttributes(url IURL, dict IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithURL:documentAttributes:"), url, dict)
	rv.Autorelease()
	return rv
}


// Creates an attributed string from the contents of the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(URL:options:documentAttributes:)
func NewAttributedStringWithURLOptionsDocumentAttributesError(url IURL, options IDictionary, dict IDictionary, error_ IError) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithURL:options:documentAttributes:error:"), url, options, dict, error_)
	rv.Autorelease()
	return rv
}







// Creates an attributed string with an adaptive image glyph and applies the specified attributes to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(adaptiveImageGlyph:attributes:)
func (ac _AttributedStringClass) AttributedStringWithAdaptiveImageGlyphAttributes(adaptiveImageGlyph objectivec.IObject, attributes IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("attributedStringWithAdaptiveImageGlyph:attributes:"), adaptiveImageGlyph, attributes)
	return rv
}


// Creates an attributed string with an attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(attachment:)
func (ac _AttributedStringClass) AttributedStringWithAttachment(attachment objectivec.IObject) IAttributedString {
	rv := objc.Send[AttributedString](objc.ID(ac.class), objc.Sel("attributedStringWithAttachment:"), attachment)
	return rv
}


// Creates an attributed string with an attachment and applies the specified attributes to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(attachment:attributes:)
func (ac _AttributedStringClass) AttributedStringWithAttachmentAttributes(attachment objectivec.IObject, attributes IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("attributedStringWithAttachment:attributes:"), attachment, attributes)
	return rv
}


// Creates an attributed string from the specified HTML data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/loadFromHTML(data:options:completionHandler:)
func (ac _AttributedStringClass) LoadFromHTMLWithDataOptionsCompletionHandler(data IData, options IDictionary, completionHandler AttributedStringCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("loadFromHTMLWithData:options:completionHandler:"), data, options, completionHandler)
}


// Creates an attributed string by converting the content of a local HTML file at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/loadFromHTML(fileURL:options:completionHandler:)
func (ac _AttributedStringClass) LoadFromHTMLWithFileURLOptionsCompletionHandler(fileURL IURL, options IDictionary, completionHandler AttributedStringCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("loadFromHTMLWithFileURL:options:completionHandler:"), fileURL, options, completionHandler)
}


// Creates an attributed string by converting the contents of the specified HTML URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/loadFromHTML(request:options:completionHandler:)
func (ac _AttributedStringClass) LoadFromHTMLWithRequestOptionsCompletionHandler(request IURLRequest, options IDictionary, completionHandler AttributedStringCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("loadFromHTMLWithRequest:options:completionHandler:"), request, options, completionHandler)
}


// Creates an attributed string from the specified HTML string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/loadFromHTML(string:options:completionHandler:)
func (ac _AttributedStringClass) LoadFromHTMLWithStringOptionsCompletionHandler(string_ IString, options IDictionary, completionHandler AttributedStringCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("loadFromHTMLWithString:options:completionHandler:"), string_, options, completionHandler)
}


// Creates an attributed string by substituting arguments into a specially formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/localizedAttributedStringWithFormat:
func (ac _AttributedStringClass) LocalizedAttributedStringWithFormat(format IAttributedString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("localizedAttributedStringWithFormat:"), format)
	return rv
}


// Creates an attributed string by substituting arguments into a specially formatted string and applying additional contextual information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/localizedAttributedStringWithFormat:context:
func (ac _AttributedStringClass) LocalizedAttributedStringWithFormatContext(format IAttributedString, context IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("localizedAttributedStringWithFormat:context:"), format, context)
	return rv
}


// Creates an attributed string by substituting a list of function arguments into a specially formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/localizedAttributedStringWithFormat:options:
func (ac _AttributedStringClass) LocalizedAttributedStringWithFormatOptions(format IAttributedString, options AttributedStringFormattingOptions) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("localizedAttributedStringWithFormat:options:"), format, options)
	return rv
}


// Creates an attributed string by substituting a list of function arguments into a specially formatted string and applying additional contextual information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/localizedAttributedStringWithFormat:options:context:
func (ac _AttributedStringClass) LocalizedAttributedStringWithFormatOptionsContext(format IAttributedString, options AttributedStringFormattingOptions, context IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("localizedAttributedStringWithFormat:options:context:"), format, options, context)
	return rv
}


// Returns an array of strings that represent file types that can be loaded as text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textFileTypes
func (ac _AttributedStringClass) TextFileTypes() IArray {
	rv := objc.Send[Array](objc.ID(ac.class), objc.Sel("textFileTypes"))
	return rv
}


// Returns an array of pasteboard types that can be loaded as text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textPasteboardTypes
func (ac _AttributedStringClass) TextPasteboardTypes() IArray {
	rv := objc.Send[Array](objc.ID(ac.class), objc.Sel("textPasteboardTypes"))
	return rv
}


// Returns an array of strings that represent file types that can be loaded as a text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textUnfilteredFileTypes
func (ac _AttributedStringClass) TextUnfilteredFileTypes() IArray {
	rv := objc.Send[Array](objc.ID(ac.class), objc.Sel("textUnfilteredFileTypes"))
	return rv
}


// Returns an array of pasteboard types that can be loaded as text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textUnfilteredPasteboardTypes
func (ac _AttributedStringClass) TextUnfilteredPasteboardTypes() IArray {
	rv := objc.Send[Array](objc.ID(ac.class), objc.Sel("textUnfilteredPasteboardTypes"))
	return rv
}







// An array of UTI strings that identify the file types that attributed strings support, either directly or through a user-installed filter service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textTypes
func (ac _AttributedStringClass) TextTypes() []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("textTypes"))
	return rv
}

// An array of UTI strings that identify the file types that attributed strings support directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textUnfilteredTypes
func (ac _AttributedStringClass) TextUnfilteredTypes() []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("textUnfilteredTypes"))
	return rv
}






// Returns the value for an attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attribute(_:at:effectiveRange:)
func (a_ AttributedString) AttributeAtIndexEffectiveRange(attrName AttributedStringKey, location uint, range_ RangePointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("attribute:atIndex:effectiveRange:"), attrName, location, range_)
	return rv
}


// Returns the value for the attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attribute(_:at:longestEffectiveRange:in:)
func (a_ AttributedString) AttributeAtIndexLongestEffectiveRangeInRange(attrName AttributedStringKey, location uint, range_ RangePointer, rangeLimit Range) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("attribute:atIndex:longestEffectiveRange:inRange:"), attrName, location, range_, rangeLimit)
	return rv
}


// Returns an attributed string consisting of the characters and attributes within the specified range in the attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributedSubstring(from:)
func (a_ AttributedString) AttributedSubstringFromRange(range_ Range) IAttributedString {
	rv := objc.Send[AttributedString](a_.ID, objc.Sel("attributedSubstringFromRange:"), range_)
	return rv
}


// Returns the attributes for the character at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributes(at:effectiveRange:)
func (a_ AttributedString) AttributesAtIndexEffectiveRange(location uint, range_ RangePointer) IDictionary {
	rv := objc.Send[Dictionary](a_.ID, objc.Sel("attributesAtIndex:effectiveRange:"), location, range_)
	return rv
}


// Returns the attributes for the character at the specified index and, by reference, the range where the attributes apply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributes(at:longestEffectiveRange:in:)
func (a_ AttributedString) AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ RangePointer, rangeLimit Range) IDictionary {
	rv := objc.Send[Dictionary](a_.ID, objc.Sel("attributesAtIndex:longestEffectiveRange:inRange:"), location, range_, rangeLimit)
	return rv
}


// Calculates and returns a bounding rectangle for the attributed string using the options specified within the specified rectangle in the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/boundingRect(with:options:)
func (a_ AttributedString) BoundingRectWithSizeOptions(size corefoundation.CGSize, options StringDrawingOptions) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](a_.ID, objc.Sel("boundingRectWithSize:options:"), size, options)
	return rv
}


// Returns the bounding rectangle necessary to draw the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/boundingRect(with:options:context:)
func (a_ AttributedString) BoundingRectWithSizeOptionsContext(size corefoundation.CGSize, options StringDrawingOptions, context objectivec.IObject) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](a_.ID, objc.Sel("boundingRectWithSize:options:context:"), size, options, context)
	return rv
}


// Returns a Boolean value that indicates if the attributed string contains an attachment in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/containsAttachments(in:)
func (a_ AttributedString) ContainsAttachmentsInRange(range_ Range) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("containsAttachmentsInRange:"), range_)
	return rv
}


// Returns a data object that contains a text stream corresponding to the characters and attributes within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/data(from:documentAttributes:)
func (a_ AttributedString) DataFromRangeDocumentAttributesError(range_ Range, dict IDictionary, error_ IError) IData {
	rv := objc.Send[Data](a_.ID, objc.Sel("dataFromRange:documentAttributes:error:"), range_, dict, error_)
	return rv
}


// Returns a data object that contains a Microsoft Word–format stream corresponding to the characters and attributes within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/docFormat(from:documentAttributes:)
func (a_ AttributedString) DocFormatFromRangeDocumentAttributes(range_ Range, dict IDictionary) IData {
	rv := objc.Send[Data](a_.ID, objc.Sel("docFormatFromRange:documentAttributes:"), range_, dict)
	return rv
}


// Returns the range of characters that form a word (or other linguistic unit) surrounding the specified index, taking language characteristics into account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/doubleClick(at:)
func (a_ AttributedString) DoubleClickAtIndex(location uint) Range {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("doubleClickAtIndex:"), location)
	return rv
}


// Draws the attributed string starting at the specified point in the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(at:)
func (a_ AttributedString) DrawAtPoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawAtPoint:"), point)
}


// Draws the attributed string inside the specified bounding rectangle in the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(in:)
func (a_ AttributedString) DrawInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawInRect:"), rect)
}


// Draws the attributed string with the specified options within the specified rectangle in the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(with:options:)
func (a_ AttributedString) DrawWithRectOptions(rect corefoundation.CGRect, options StringDrawingOptions) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawWithRect:options:"), rect, options)
}


// Draws the attributed string in the specified bounding rectangle using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(with:options:context:)
func (a_ AttributedString) DrawWithRectOptionsContext(rect corefoundation.CGRect, options StringDrawingOptions, context objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("drawWithRect:options:context:"), rect, options, context)
}


// Executes the specified closure or block for each range of a particular attribute in the attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/enumerateAttribute(_:in:options:using:)
func (a_ AttributedString) EnumerateAttributeInRangeOptionsUsingBlock(attrName AttributedStringKey, enumerationRange Range, opts AttributedStringEnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("enumerateAttribute:inRange:options:usingBlock:"), attrName, enumerationRange, opts, block)
}


// Executes the specified closure or block for each range of attributes in the attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/enumerateAttributes(in:options:using:)
func (a_ AttributedString) EnumerateAttributesInRangeOptionsUsingBlock(enumerationRange Range, opts AttributedStringEnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("enumerateAttributesInRange:options:usingBlock:"), enumerationRange, opts, block)
}


// Returns a file wrapper object that contains a text stream corresponding to the characters and attributes within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/fileWrapper(from:documentAttributes:)
func (a_ AttributedString) FileWrapperFromRangeDocumentAttributesError(range_ Range, dict IDictionary, error_ IError) IFileWrapper {
	rv := objc.Send[FileWrapper](a_.ID, objc.Sel("fileWrapperFromRange:documentAttributes:error:"), range_, dict, error_)
	return rv
}


// Returns the font attributes in effect for the character at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/fontAttributes(in:)
func (a_ AttributedString) FontAttributesInRange(range_ Range) IDictionary {
	rv := objc.Send[Dictionary](a_.ID, objc.Sel("fontAttributesInRange:"), range_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/inflecting()
func (a_ AttributedString) AttributedStringByInflectingString() IAttributedString {
	rv := objc.Send[AttributedString](a_.ID, objc.Sel("attributedStringByInflectingString"))
	return rv
}


// Returns a Boolean value that indicates whether the attributed string is equal to the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/isEqual(to:)
func (a_ AttributedString) IsEqualToAttributedString(other IAttributedString) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEqualToAttributedString:"), other)
	return rv
}


// Returns the index of the item at the specified location within the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/itemNumber(in:at:)
func (a_ AttributedString) ItemNumberInTextListAtIndex(list objectivec.IObject, location uint) int {
	rv := objc.Send[int](a_.ID, objc.Sel("itemNumberInTextList:atIndex:"), list, location)
	return rv
}


// Returns the appropriate line break when the character at the index doesn’t fit on the same line as the character at the beginning of the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/lineBreak(before:within:)
func (a_ AttributedString) LineBreakBeforeIndexWithinRange(location uint, aRange Range) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("lineBreakBeforeIndex:withinRange:"), location, aRange)
	return rv
}


// Returns the index of the closest character before the specified index, and within the specified range, that can fit on a new line by hyphenating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/lineBreakByHyphenating(before:within:)
func (a_ AttributedString) LineBreakByHyphenatingBeforeIndexWithinRange(location uint, aRange Range) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("lineBreakByHyphenatingBeforeIndex:withinRange:"), location, aRange)
	return rv
}


// Returns the index of the first character of the word after or before the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/nextWord(from:forward:)
func (a_ AttributedString) NextWordFromIndexForward(location uint, isForward bool) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("nextWordFromIndex:forward:"), location, isForward)
	return rv
}


// Returns a Boolean value that indicates whether the specified range of text prefers RTFD formatting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/prefersRTFD(in:)
func (a_ AttributedString) PrefersRTFDInRange(range_ Range) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prefersRTFDInRange:"), range_)
	return rv
}


// Returns the range of the individual text block that contains the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-1wrcp
func (a_ AttributedString) RangeOfTextBlockAtIndex(block objectivec.IObject, location uint) Range {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("rangeOfTextBlock:atIndex:"), block, location)
	return rv
}


// Returns the range of the specified text table that contains the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-3fevu
func (a_ AttributedString) RangeOfTextTableAtIndex(table objectivec.IObject, location uint) Range {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("rangeOfTextTable:atIndex:"), table, location)
	return rv
}


// Returns the range of the specified text list that contains the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-6um0x
func (a_ AttributedString) RangeOfTextListAtIndex(list objectivec.IObject, location uint) Range {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("rangeOfTextList:atIndex:"), list, location)
	return rv
}


// Returns a data object that contains an RTF stream corresponding to the characters and attributes within the specified range, omitting all attachment attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtf(from:documentAttributes:)
func (a_ AttributedString) RTFFromRangeDocumentAttributes(range_ Range, dict IDictionary) IData {
	rv := objc.Send[Data](a_.ID, objc.Sel("RTFFromRange:documentAttributes:"), range_, dict)
	return rv
}


// Returns a data object that contains an RTFD stream corresponding to the characters and attributes within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtfd(from:documentAttributes:)
func (a_ AttributedString) RTFDFromRangeDocumentAttributes(range_ Range, dict IDictionary) IData {
	rv := objc.Send[Data](a_.ID, objc.Sel("RTFDFromRange:documentAttributes:"), range_, dict)
	return rv
}


// Returns a file wrapper object that contains an RTFD document corresponding to the characters and attributes within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtfdFileWrapper(from:documentAttributes:)
func (a_ AttributedString) RTFDFileWrapperFromRangeDocumentAttributes(range_ Range, dict IDictionary) IFileWrapper {
	rv := objc.Send[FileWrapper](a_.ID, objc.Sel("RTFDFileWrapperFromRange:documentAttributes:"), range_, dict)
	return rv
}


// Returns the ruler (paragraph) attributes in effect for the characters within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/rulerAttributes(in:)
func (a_ AttributedString) RulerAttributesInRange(range_ Range) IDictionary {
	rv := objc.Send[Dictionary](a_.ID, objc.Sel("rulerAttributesInRange:"), range_)
	return rv
}


// Returns the size necessary to draw the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/size()
func (a_ AttributedString) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("size"))
	return rv
}







// A Boolean value that indicates whether the attribute string contains any attachment attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/containsAttachments
func (a_ AttributedString) ContainsAttachments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("containsAttachments"))
	return rv
}


// The length of the attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/length
func (a_ AttributedString) Length() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("length"))
	return rv
}


// The character contents of the attributed string as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (a_ AttributedString) String() IString {
	rv := objc.Send[String](a_.ID, objc.Sel("string"))
	return rv
}


// An array of UTI strings that identify the file types that attributed strings support, either directly or through a user-installed filter service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textTypes
func (a_ AttributedString) TextTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("textTypes"))
	return rv
}


// An array of UTI strings that identify the file types that attributed strings support directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/textUnfilteredTypes
func (a_ AttributedString) TextUnfilteredTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("textUnfilteredTypes"))
	return rv
}







