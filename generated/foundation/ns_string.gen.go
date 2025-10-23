// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [String] class.
var (
	StringClass     _StringClass
	StringClassOnce sync.Once
)

func getStringClass() _StringClass {
	StringClassOnce.Do(func() {
		StringClass = _StringClass{objc.GetClass("NSString")}
	})
	return StringClass
}

type _StringClass struct {
	class objc.Class
}

// An interface definition for the [String] class.
type IString interface {
	objectivec.IObject
	// properties:
	StringByAbbreviatingWithTildeInPath() IString
	BoolValue() bool /* primitive/slice/pointer. */
	CapitalizedString() IString
	DecomposedStringWithCanonicalMapping() IString
	DecomposedStringWithCompatibilityMapping() IString
	StringByDeletingLastPathComponent() IString
	StringByDeletingPathExtension() IString
	Description() IString
	DoubleValue() float64 /* primitive/slice/pointer. */
	StringByExpandingTildeInPath() IString
	FastestEncoding() StringEncoding /* not a class type */
	FileSystemRepresentation() unsafe.Pointer
	FloatValue() float32 /* primitive/slice/pointer. */
	Hash() uint /* primitive/slice/pointer. */
	IntValue() int /* primitive/slice/pointer. */
	IntegerValue() int /* primitive/slice/pointer. */
	AbsolutePath() bool /* primitive/slice/pointer. */
	LastPathComponent() IString
	Length() uint /* primitive/slice/pointer. */
	LocalizedCapitalizedString() IString
	LocalizedLowercaseString() IString
	LocalizedUppercaseString() IString
	LongLongValue() unsafe.Pointer
	LowercaseString() IString
	PathComponents() []string /* primitive/slice/pointer. */
	PathExtension() IString
	PrecomposedStringWithCanonicalMapping() IString
	PrecomposedStringWithCompatibilityMapping() IString
	StringByRemovingPercentEncoding() IString
	StringByResolvingSymlinksInPath() IString
	SmallestEncoding() StringEncoding /* not a class type */
	StringByStandardizingPath() IString
	UppercaseString() IString
	UTF8String() unsafe.Pointer
	AbbreviatingWithTildeInPath() IString
	SetAbbreviatingWithTildeInPath(value IString)
	Capitalized() IString
	SetCapitalized(value IString)
	CustomPlaygroundQuickLook() unsafe.Pointer
	SetCustomPlaygroundQuickLook(value unsafe.Pointer)
	DeletingLastPathComponent() IString
	SetDeletingLastPathComponent(value IString)
	DeletingPathExtension() IString
	SetDeletingPathExtension(value IString)
	ExpandingTildeInPath() IString
	SetExpandingTildeInPath(value IString)
	IsAbsolutePath() bool /* primitive/slice/pointer. */
	SetIsAbsolutePath(value bool /* primitive/slice/pointer. */)
	LocalizedCapitalized() IString
	SetLocalizedCapitalized(value IString)
	LocalizedLowercase() IString
	SetLocalizedLowercase(value IString)
	LocalizedUppercase() IString
	SetLocalizedUppercase(value IString)
	Lowercased() IString
	SetLowercased(value IString)
	RemovingPercentEncoding() IString
	SetRemovingPercentEncoding(value IString)
	ResolvingSymlinksInPath() IString
	SetResolvingSymlinksInPath(value IString)
	StandardizingPath() IString
	SetStandardizingPath(value IString)
	Uppercased() IString
	SetUppercased(value IString)
	// methods:
	StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters CharacterSet /* not a class type */) IString
	StringByAppendingString(aString IString) IString
	StringByAppendingPathComponent(str IString) IString
	StringByAppendingPathComponentConformingToType(partialName IString, contentType objc.IObject /* cross-framework UTType */) IString
	StringByAppendingPathExtension(str IString) IString
	StringByAppendingPathExtensionForType(contentType objc.IObject /* cross-framework UTType */) IString
	StringByApplyingTransformReverse(transform StringTransform /* not a class type */, reverse bool /* primitive/slice/pointer. */) IString
	BoundingRectWithSizeOptionsAttributes(size objc.IObject /* cross-framework Size */, options StringDrawingOptions, attributes IDictionary /* already interface */) objc.IObject /* cross-framework: Rect */
	BoundingRectWithSizeOptionsAttributesContext(size objc.IObject /* cross-framework Size */, options StringDrawingOptions, attributes IDictionary /* already interface */, context objc.IObject /* cross-framework StringDrawingContext */) objc.IObject /* cross-framework: Rect */
	CStringUsingEncoding(encoding StringEncoding /* not a class type */) unsafe.Pointer
	CanBeConvertedToEncoding(encoding StringEncoding /* not a class type */) bool /* primitive/slice/pointer. */
	CapitalizedStringWithLocale(locale Locale /* not a class type */) IString
	CaseInsensitiveCompare(string_ IString) ComparisonResult /* not a class type */
	CharacterAtIndex(index uint /* primitive/slice/pointer. */) unsafe.Pointer
	CommonPrefixWithStringOptions(str IString, mask StringCompareOptions) IString
	Compare(string_ IString) ComparisonResult /* not a class type */
	CompareOptions(string_ IString, mask StringCompareOptions) ComparisonResult /* not a class type */
	CompareOptionsRange(string_ IString, mask StringCompareOptions, rangeOfReceiverToCompare Range /* not a class type */) ComparisonResult /* not a class type */
	CompareOptionsRangeLocale(string_ IString, mask StringCompareOptions, rangeOfReceiverToCompare Range /* not a class type */, locale objectivec.IObject) ComparisonResult /* not a class type */
	CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName IString, flag bool /* primitive/slice/pointer. */, outputArray []string /* primitive/slice/pointer. */, filterTypes []string /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */
	ComponentsSeparatedByString(separator IString) []string /* primitive/slice/pointer. */
	ComponentsSeparatedByCharactersInSet(separator CharacterSet /* not a class type */) []string /* primitive/slice/pointer. */
	ContainsString(str IString) bool /* primitive/slice/pointer. */
	DataUsingEncoding(encoding StringEncoding /* not a class type */) IData
	DataUsingEncodingAllowLossyConversion(encoding StringEncoding /* not a class type */, lossy bool /* primitive/slice/pointer. */) IData
	DrawAtPointWithAttributes(point objc.IObject /* cross-framework Point */, attrs IDictionary /* already interface */)
	DrawInRectWithAttributes(rect objc.IObject /* cross-framework Rect */, attrs IDictionary /* already interface */)
	DrawWithRectOptionsAttributes(rect objc.IObject /* cross-framework Rect */, options StringDrawingOptions, attributes IDictionary /* already interface */)
	DrawWithRectOptionsAttributesContext(rect objc.IObject /* cross-framework Rect */, options StringDrawingOptions, attributes IDictionary /* already interface */, context objc.IObject /* cross-framework StringDrawingContext */)
	EnumerateLinesUsingBlock(block unsafe.Pointer)
	EnumerateSubstringsInRangeOptionsUsingBlock(range_ Range /* not a class type */, opts StringEnumerationOptions, block unsafe.Pointer)
	StringByFoldingWithOptionsLocale(options StringCompareOptions, locale Locale /* not a class type */) IString
	GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint /* primitive/slice/pointer. */, usedBufferCount UInteger /* not a class type */, encoding StringEncoding /* not a class type */, options StringEncodingConversionOptions, range_ Range /* not a class type */, leftover RangePointer /* not a class type */) bool /* primitive/slice/pointer. */
	GetCStringMaxLengthEncoding(buffer unsafe.Pointer, maxBufferCount uint /* primitive/slice/pointer. */, encoding StringEncoding /* not a class type */) bool /* primitive/slice/pointer. */
	GetCharacters(buffer unsafe.Pointer)
	GetCharactersRange(buffer unsafe.Pointer, range_ Range /* not a class type */)
	GetFileSystemRepresentationMaxLength(cname unsafe.Pointer, max uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	GetLineStartEndContentsEndForRange(startPtr UInteger /* not a class type */, lineEndPtr UInteger /* not a class type */, contentsEndPtr UInteger /* not a class type */, range_ Range /* not a class type */)
	GetParagraphStartEndContentsEndForRange(startPtr UInteger /* not a class type */, parEndPtr UInteger /* not a class type */, contentsEndPtr UInteger /* not a class type */, range_ Range /* not a class type */)
	HasPrefix(str IString) bool /* primitive/slice/pointer. */
	HasSuffix(str IString) bool /* primitive/slice/pointer. */
	IsEqualToString(aString IString) bool /* primitive/slice/pointer. */
	LengthOfBytesUsingEncoding(enc StringEncoding /* not a class type */) uint /* primitive/slice/pointer. */
	LineRangeForRange(range_ Range /* not a class type */) Range /* not a class type */
	LocalizedCaseInsensitiveCompare(string_ IString) ComparisonResult /* not a class type */
	LocalizedCaseInsensitiveContainsString(str IString) bool /* primitive/slice/pointer. */
	LocalizedCompare(string_ IString) ComparisonResult /* not a class type */
	LocalizedStandardCompare(string_ IString) ComparisonResult /* not a class type */
	LocalizedStandardContainsString(str IString) bool /* primitive/slice/pointer. */
	LocalizedStandardRangeOfString(str IString) Range /* not a class type */
	LowercaseStringWithLocale(locale Locale /* not a class type */) IString
	MaximumLengthOfBytesUsingEncoding(enc StringEncoding /* not a class type */) uint /* primitive/slice/pointer. */
	StringByPaddingToLengthWithStringStartingAtIndex(newLength uint /* primitive/slice/pointer. */, padString IString, padIndex uint /* primitive/slice/pointer. */) IString
	ParagraphRangeForRange(range_ Range /* not a class type */) Range /* not a class type */
	PropertyList() objc.ID
	PropertyListFromStringsFileFormat() IDictionary
	RangeOfString(searchString IString) Range /* not a class type */
	RangeOfStringOptions(searchString IString, mask StringCompareOptions) Range /* not a class type */
	RangeOfStringOptionsRange(searchString IString, mask StringCompareOptions, rangeOfReceiverToSearch Range /* not a class type */) Range /* not a class type */
	RangeOfStringOptionsRangeLocale(searchString IString, mask StringCompareOptions, rangeOfReceiverToSearch Range /* not a class type */, locale Locale /* not a class type */) Range /* not a class type */
	RangeOfCharacterFromSet(searchSet CharacterSet /* not a class type */) Range /* not a class type */
	RangeOfCharacterFromSetOptions(searchSet CharacterSet /* not a class type */, mask StringCompareOptions) Range /* not a class type */
	RangeOfCharacterFromSetOptionsRange(searchSet CharacterSet /* not a class type */, mask StringCompareOptions, rangeOfReceiverToSearch Range /* not a class type */) Range /* not a class type */
	RangeOfComposedCharacterSequenceAtIndex(index uint /* primitive/slice/pointer. */) Range /* not a class type */
	RangeOfComposedCharacterSequencesForRange(range_ Range /* not a class type */) Range /* not a class type */
	StringByReplacingCharactersInRangeWithString(range_ Range /* not a class type */, replacement IString) IString
	StringByReplacingOccurrencesOfStringWithString(target IString, replacement IString) IString
	StringByReplacingOccurrencesOfStringWithStringOptionsRange(target IString, replacement IString, options StringCompareOptions, searchRange Range /* not a class type */) IString
	SizeWithAttributes(attrs IDictionary /* already interface */) objc.IObject /* cross-framework: Size */
	Sr_sensorForDeletionRecordsFromSensor() unsafe.Pointer
	StringByAppendingFormat(format IString) IString
	StringsByAppendingPaths(paths []string /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */
	SubstringFromIndex(from uint /* primitive/slice/pointer. */) IString
	SubstringToIndex(to uint /* primitive/slice/pointer. */) IString
	SubstringWithRange(range_ Range /* not a class type */) IString
	StringByTrimmingCharactersInSet(set CharacterSet /* not a class type */) IString
	UppercaseStringWithLocale(locale Locale /* not a class type */) IString
	VariantFittingPresentationWidth(width int /* primitive/slice/pointer. */) IString
	WriteToURLAtomicallyEncodingError(url objc.IObject /* cross-framework NSURL */, useAuxiliaryFile bool /* primitive/slice/pointer. */, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	WriteToFileAtomicallyEncodingError(path IString, useAuxiliaryFile bool /* primitive/slice/pointer. */, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
}

// A static, plain-text Unicode string object.
//
// You can use this type in Swift when you need reference semantics or other Foundation-specific behavior. The class and its mutable subclass, , provide an extensive set of APIs for working with strings, including methods for comparing, searching, and modifying strings. objects are used throughout Foundation and other Cocoa frameworks, serving as the basis for all textual and linguistic functionality on the platform. is with its Core Foundation counterpart, . See for more information.


// A static, plain-text Unicode string object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString
type String struct {
	objectivec.Object
}

// StringFrom constructs a [String] from an unsafe.Pointer.
//
// A static, plain-text Unicode string object.
func StringFrom(ptr unsafe.Pointer) String {
	return String{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StringClass) Alloc() String {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StringClass) New() String {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ String) Init() String {
	rv := objc.Send[String](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ String) Autorelease() String {
	rv := objc.Send[String](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewString creates a new String instance.
func NewString() String {
	return getStringClass().New()
}



// Returns an initialized object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(bytes:length:encoding:)
func NewStringWithBytesLengthEncoding(bytes unsafe.Pointer, len_ uint /* primitive/slice/pointer. */, encoding StringEncoding /* not a class type */) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithBytes:length:encoding:"), bytes, len_, encoding)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:deallocator:)
func NewStringWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len_ uint /* primitive/slice/pointer. */, encoding StringEncoding /* not a class type */, deallocator unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:deallocator:"), bytes, len_, encoding, deallocator)
	rv.Autorelease()
	return rv
}


// Returns an initialized object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:freeWhenDone:)
func NewStringWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len_ uint /* primitive/slice/pointer. */, encoding StringEncoding /* not a class type */, freeBuffer bool /* primitive/slice/pointer. */) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:freeWhenDone:"), bytes, len_, encoding, freeBuffer)
	rv.Autorelease()
	return rv
}


// Initializes the receiver, a newly allocated object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:)
func NewStringWithCString(bytes unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCString:"), bytes)
	rv.Autorelease()
	return rv
}


// Returns an object initialized using the characters in a given C array, interpreted according to a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:encoding:)-20f9h
func NewStringWithCStringEncoding(nullTerminatedCString unsafe.Pointer, encoding StringEncoding /* not a class type */) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCString:encoding:"), nullTerminatedCString, encoding)
	rv.Autorelease()
	return rv
}


// Initializes the receiver, a newly allocated object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:length:)
func NewStringWithCStringLength(bytes unsafe.Pointer, length uint /* primitive/slice/pointer. */) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCString:length:"), bytes, length)
	rv.Autorelease()
	return rv
}


// Initializes the receiver, a newly allocated object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CStringNoCopy:length:freeWhenDone:)
func NewStringWithCStringNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint /* primitive/slice/pointer. */, freeBuffer bool /* primitive/slice/pointer. */) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCStringNoCopy:length:freeWhenDone:"), bytes, length, freeBuffer)
	rv.Autorelease()
	return rv
}


// Returns an initialized object that contains a given number of characters from a given C array of UTF-16 code units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(characters:length:)
func NewStringWithCharactersLength(characters unsafe.Pointer, length uint /* primitive/slice/pointer. */) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCharacters:length:"), characters, length)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:deallocator:)
func NewStringWithCharactersNoCopyLengthDeallocator(chars unsafe.Pointer, len_ uint /* primitive/slice/pointer. */, deallocator unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCharactersNoCopy:length:deallocator:"), chars, len_, deallocator)
	rv.Autorelease()
	return rv
}


// Returns an initialized object that contains a given number of characters from a given C array of UTF-16 code units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:freeWhenDone:)
func NewStringWithCharactersNoCopyLengthFreeWhenDone(characters unsafe.Pointer, length uint /* primitive/slice/pointer. */, freeBuffer bool /* primitive/slice/pointer. */) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCharactersNoCopy:length:freeWhenDone:"), characters, length, freeBuffer)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(coder:)
func NewStringWithCoder(coder ICoder) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes the receiver, a newly allocated object, by reading data from the file named by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:)
func NewStringWithContentsOfFile(path IString) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfFile:"), path)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by reading data from the file at a given path using a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:encoding:)
func NewStringWithContentsOfFileEncodingError(path IString, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfFile:encoding:error:"), path, enc, error_)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by reading data from the file at a given path and returns by reference the encoding used to interpret the characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:usedEncoding:)
func NewStringWithContentsOfFileUsedEncodingError(path IString, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfFile:usedEncoding:error:"), path, enc, error_)
	rv.Autorelease()
	return rv
}


// Initializes the receiver, a newly allocated object, by reading data from the location named by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:)
func NewStringWithContentsOfURL(url objc.IObject /* cross-framework NSURL */) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by reading data from a given URL interpreted using a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:encoding:)-715fw
func NewStringWithContentsOfURLEncodingError(url objc.IObject /* cross-framework NSURL */, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfURL:encoding:error:"), url, enc, error_)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by reading data from a given URL and returns by reference the encoding used to interpret the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:usedEncoding:)-2c72d
func NewStringWithContentsOfURLUsedEncodingError(url objc.IObject /* cross-framework NSURL */, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfURL:usedEncoding:error:"), url, enc, error_)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by converting given data into UTF-16 code units using a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(data:encoding:)
func NewStringWithDataEncoding(data IData, encoding StringEncoding /* not a class type */) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithData:encoding:"), data, encoding)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:
func NewStringWithFormat(format IString) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:"), format)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted without any localization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(format:arguments:)
func NewStringWithFormatArguments(format IString, argList unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:arguments:"), format, argList)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:locale:
func NewStringWithFormatLocale(format IString, locale objectivec.IObject) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:locale:"), format, locale)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale information. This method is meant to be called from within a variadic function, where the argument list will be available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(format:locale:arguments:)
func NewStringWithFormatLocaleArguments(format IString, locale objectivec.IObject, argList unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:locale:arguments:"), format, locale, argList)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by copying the characters from another given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(string:)-210xa
func NewStringWithString(aString IString) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithString:"), aString)
	rv.Autorelease()
	return rv
}


// Returns an object initialized by copying the characters from a given C array of UTF8-encoded bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(UTF8String:)-vg2b
func NewStringWithUTF8String(nullTerminatedCString unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithUTF8String:"), nullTerminatedCString)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersArgumentsError(format IString, validFormatSpecifiers IString, argList unsafe.Pointer, error_ unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:arguments:error:"), format, validFormatSpecifiers, argList, error_)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:error:
func NewStringWithValidatedFormatValidFormatSpecifiersError(format IString, validFormatSpecifiers IString, error_ unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:error:"), format, validFormatSpecifiers, error_)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format IString, validFormatSpecifiers IString, locale objectivec.IObject, argList unsafe.Pointer, error_ unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:"), format, validFormatSpecifiers, locale, argList, error_)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleError(format IString, validFormatSpecifiers IString, locale objectivec.IObject, error_ unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:error:"), format, validFormatSpecifiers, locale, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormat(format IString) IString {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:"), format)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormatFromTable(format IString, table IString) IString {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:fromTable:"), format, table)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:arguments:
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormatFromTableArguments(format IString, table IString, arguments unsafe.Pointer) IString {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:fromTable:arguments:"), format, table, arguments)
	return rv
}


// Returns a string containing the bytes in a given C array, interpreted according to a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:encoding:)-7auq8
func (sc _StringClass) StringWithCStringEncoding(cString unsafe.Pointer, enc StringEncoding /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithCString:encoding:"), cString, enc)
	return rv
}


// Returns a string created by copying the data from a given C array of UTF8-encoded bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(UTF8String:)-8bcy8
func (sc _StringClass) StringWithUTF8String(nullTerminatedCString unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithUTF8String:"), nullTerminatedCString)
	return rv
}


// Returns a string created by reading data from a given URL interpreted using a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:encoding:)-x6cv
func (sc _StringClass) StringWithContentsOfURLEncodingError(url objc.IObject /* cross-framework NSURL */, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfURL:encoding:error:"), url, enc, error_)
	return rv
}


// Returns a string created by reading data from a given URL and returns by reference the encoding used to interpret the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:usedEncoding:)-9jrum
func (sc _StringClass) StringWithContentsOfURLUsedEncodingError(url objc.IObject /* cross-framework NSURL */, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfURL:usedEncoding:error:"), url, enc, error_)
	return rv
}


// Returns a human-readable string giving the name of a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedName(of:)
func (sc _StringClass) LocalizedNameOfStringEncoding(encoding StringEncoding /* not a class type */) IString {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("localizedNameOfStringEncoding:"), encoding)
	return rv
}


// Returns a string created by using a given format string as a template into which the remaining argument values are substituted according to the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStringWithFormat:
func (sc _StringClass) LocalizedStringWithFormat(format IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedStringWithFormat:"), format)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStringWithValidatedFormat:validFormatSpecifiers:error:
func (sc _StringClass) LocalizedStringWithValidatedFormatValidFormatSpecifiersError(format IString, validFormatSpecifiers IString, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedStringWithValidatedFormat:validFormatSpecifiers:error:"), format, validFormatSpecifiers, error_)
	return rv
}


// Returns a localized string intended for display in a notification alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedUserNotificationString(forKey:arguments:)
func (sc _StringClass) LocalizedUserNotificationStringForKeyArguments(key IString, arguments IArray) IString {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("localizedUserNotificationStringForKey:arguments:"), key, arguments)
	return rv
}


// Returns a string built from the strings in a given array by concatenating them with a path separator between each pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/path(withComponents:)
func (sc _StringClass) PathWithComponents(components []string /* primitive/slice/pointer. */) IString {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("pathWithComponents:"), components)
	return rv
}


// Creates a new string using a given C-string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withCString:)
func (sc _StringClass) StringWithCString(bytes unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithCString:"), bytes)
	return rv
}


// Returns a string containing the characters in a given C-string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withCString:length:)
func (sc _StringClass) StringWithCStringLength(bytes unsafe.Pointer, length uint /* primitive/slice/pointer. */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithCString:length:"), bytes, length)
	return rv
}


// Returns a string created by reading data from the file named by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withContentsOf:)
func (sc _StringClass) StringWithContentsOfURL(url objc.IObject /* cross-framework NSURL */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithContentsOfURL:"), url)
	return rv
}


// Returns a string created by reading data from the file named by a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withContentsOfFile:)
func (sc _StringClass) StringWithContentsOfFile(path IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithContentsOfFile:"), path)
	return rv
}


// Returns an empty string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string
func (sc _StringClass) String() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("string"))
	return rv
}


// Returns the string encoding for the given data as detected by attempting to create a string according to the specified encoding options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringEncoding(for:encodingOptions:convertedString:usedLossyConversion:)
func (sc _StringClass) StringEncodingForDataEncodingOptionsConvertedStringUsedLossyConversion(data IData, opts IDictionary /* already interface */, string_ IString, usedLossyConversion unsafe.Pointer) StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](objc.ID(sc.class), objc.Sel("stringEncodingForData:encodingOptions:convertedString:usedLossyConversion:"), data, opts, string_, usedLossyConversion)
	return rv
}


// Returns a string containing a given number of characters taken from a given C array of UTF-16 code units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithCharacters:length:
func (sc _StringClass) StringWithCharactersLength(characters unsafe.Pointer, length uint /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithCharacters:length:"), characters, length)
	return rv
}


// Returns a string created by reading data from the file at a given path interpreted using a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfFile:encoding:error:
func (sc _StringClass) StringWithContentsOfFileEncodingError(path IString, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfFile:encoding:error:"), path, enc, error_)
	return rv
}


// Returns a string created by reading data from the file at a given path and returns by reference the encoding used to interpret the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfFile:usedEncoding:error:
func (sc _StringClass) StringWithContentsOfFileUsedEncodingError(path IString, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfFile:usedEncoding:error:"), path, enc, error_)
	return rv
}


// Returns a string created by using a given format string as a template into which the remaining argument values are substituted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithFormat:
func (sc _StringClass) StringWithFormat(format IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithFormat:"), format)
	return rv
}


// Returns a string created by copying the characters from another given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithString:
func (sc _StringClass) StringWithString(string_ IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithString:"), string_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithValidatedFormat:validFormatSpecifiers:error:
func (sc _StringClass) StringWithValidatedFormatValidFormatSpecifiersError(format IString, validFormatSpecifiers IString, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithValidatedFormat:validFormatSpecifiers:error:"), format, validFormatSpecifiers, error_)
	return rv
}


// Returns a zero-terminated list of the encodings string objects support in the application’s environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/availableStringEncodings
func (sc _StringClass) AvailableStringEncodings() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](objc.ID(sc.class), objc.Sel("availableStringEncodings"))
	return rv
}

// Returns the C-string encoding assumed for any method accepting a C string as an argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/defaultCStringEncoding
func (sc _StringClass) DefaultCStringEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](objc.ID(sc.class), objc.Sel("defaultCStringEncoding"))
	return rv
}

// Returns a new string made from the receiver by replacing all characters not in the specified set with percent-encoded characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/addingPercentEncoding(withAllowedCharacters:)
func (s_ String) StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters CharacterSet /* not a class type */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByAddingPercentEncodingWithAllowedCharacters:"), allowedCharacters)
	return rv
}


// Returns a new string made by appending a given string to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appending(_:)
func (s_ String) StringByAppendingString(aString IString) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByAppendingString:"), aString)
	return rv
}


// Returns a new string made by appending to the receiver a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathComponent(_:)
func (s_ String) StringByAppendingPathComponent(str IString) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByAppendingPathComponent:"), str)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathComponent(_:conformingTo:)
func (s_ String) StringByAppendingPathComponentConformingToType(partialName IString, contentType objc.IObject /* cross-framework UTType */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByAppendingPathComponent:conformingToType:"), partialName, contentType)
	return rv
}


// Returns a new string made by appending to the receiver an extension separator followed by a given extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathExtension(_:)
func (s_ String) StringByAppendingPathExtension(str IString) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByAppendingPathExtension:"), str)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathExtension(for:)
func (s_ String) StringByAppendingPathExtensionForType(contentType objc.IObject /* cross-framework UTType */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByAppendingPathExtensionForType:"), contentType)
	return rv
}


// Returns a new string by applying a specified transform to the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/applyingTransform(_:reverse:)
func (s_ String) StringByApplyingTransformReverse(transform StringTransform /* not a class type */, reverse bool /* primitive/slice/pointer. */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByApplyingTransform:reverse:"), transform, reverse)
	return rv
}


// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/boundingRect(with:options:attributes:)
func (s_ String) BoundingRectWithSizeOptionsAttributes(size objc.IObject /* cross-framework Size */, options StringDrawingOptions, attributes IDictionary /* already interface */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("boundingRectWithSize:options:attributes:"), size, options, attributes)
	return rv
}


// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/boundingRect(with:options:attributes:context:)
func (s_ String) BoundingRectWithSizeOptionsAttributesContext(size objc.IObject /* cross-framework Size */, options StringDrawingOptions, attributes IDictionary /* already interface */, context objc.IObject /* cross-framework StringDrawingContext */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("boundingRectWithSize:options:attributes:context:"), size, options, attributes, context)
	return rv
}


// Returns a representation of the string as a C string using a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/cString(using:)
func (s_ String) CStringUsingEncoding(encoding StringEncoding /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("cStringUsingEncoding:"), encoding)
	return rv
}


// Returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/canBeConverted(to:)
func (s_ String) CanBeConvertedToEncoding(encoding StringEncoding /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("canBeConvertedToEncoding:"), encoding)
	return rv
}


// Returns a capitalized representation of the receiver using the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/capitalized(with:)
func (s_ String) CapitalizedStringWithLocale(locale Locale /* not a class type */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("capitalizedStringWithLocale:"), locale)
	return rv
}


// Returns the result of invoking with as the only option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/caseInsensitiveCompare(_:)
func (s_ String) CaseInsensitiveCompare(string_ IString) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("caseInsensitiveCompare:"), string_)
	return rv
}


// Returns the character at a given UTF-16 code unit index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/character(at:)
func (s_ String) CharacterAtIndex(index uint /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("characterAtIndex:"), index)
	return rv
}


// Returns a string containing characters the receiver and a given string have in common, starting from the beginning of each up to the first characters that aren’t equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/commonPrefix(with:options:)
func (s_ String) CommonPrefixWithStringOptions(str IString, mask StringCompareOptions) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("commonPrefixWithString:options:"), str, mask)
	return rv
}


// Returns the result of invoking with no options and the receiver’s full extent as the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:)
func (s_ String) Compare(string_ IString) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("compare:"), string_)
	return rv
}


// Compares the string with the specified string using the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:)
func (s_ String) CompareOptions(string_ IString, mask StringCompareOptions) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("compare:options:"), string_, mask)
	return rv
}


// Returns the result of invoking with a locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:range:)
func (s_ String) CompareOptionsRange(string_ IString, mask StringCompareOptions, rangeOfReceiverToCompare Range /* not a class type */) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("compare:options:range:"), string_, mask, rangeOfReceiverToCompare)
	return rv
}


// Compares the string using the specified options and returns the lexical ordering for the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:range:locale:)
func (s_ String) CompareOptionsRangeLocale(string_ IString, mask StringCompareOptions, rangeOfReceiverToCompare Range /* not a class type */, locale objectivec.IObject) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("compare:options:range:locale:"), string_, mask, rangeOfReceiverToCompare, locale)
	return rv
}


// Interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/completePath(into:caseSensitive:matchesInto:filterTypes:)
func (s_ String) CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName IString, flag bool /* primitive/slice/pointer. */, outputArray []string /* primitive/slice/pointer. */, filterTypes []string /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](s_.ID, objc.Sel("completePathIntoString:caseSensitive:matchesIntoArray:filterTypes:"), outputName, flag, outputArray, filterTypes)
	return rv
}


// Returns an array containing substrings from the receiver that have been divided by a given separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/components(separatedBy:)-238fy
func (s_ String) ComponentsSeparatedByString(separator IString) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("componentsSeparatedByString:"), separator)
	return rv
}


// Returns an array containing substrings from the receiver that have been divided by characters in a given set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/components(separatedBy:)-27x9g
func (s_ String) ComponentsSeparatedByCharactersInSet(separator CharacterSet /* not a class type */) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("componentsSeparatedByCharactersInSet:"), separator)
	return rv
}


// Returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/contains(_:)
func (s_ String) ContainsString(str IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("containsString:"), str)
	return rv
}


// Returns an object containing a representation of the receiver encoded using a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/data(using:)
func (s_ String) DataUsingEncoding(encoding StringEncoding /* not a class type */) IData {
	rv := objc.Send[Data](s_.ID, objc.Sel("dataUsingEncoding:"), encoding)
	return rv
}


// Returns an object containing a representation of the receiver encoded using a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/data(using:allowLossyConversion:)
func (s_ String) DataUsingEncodingAllowLossyConversion(encoding StringEncoding /* not a class type */, lossy bool /* primitive/slice/pointer. */) IData {
	rv := objc.Send[Data](s_.ID, objc.Sel("dataUsingEncoding:allowLossyConversion:"), encoding, lossy)
	return rv
}


// Draws the receiver with the font and other display characteristics of the given attributes, at the specified point in the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(at:withAttributes:)
func (s_ String) DrawAtPointWithAttributes(point objc.IObject /* cross-framework Point */, attrs IDictionary /* already interface */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawAtPoint:withAttributes:"), point, attrs)
}


// Draws the attributed string inside the specified bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(in:withAttributes:)
func (s_ String) DrawInRectWithAttributes(rect objc.IObject /* cross-framework Rect */, attrs IDictionary /* already interface */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawInRect:withAttributes:"), rect, attrs)
}


// Draws the receiver with the specified options and other display characteristics of the given attributes, within the specified rectangle in the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(with:options:attributes:)
func (s_ String) DrawWithRectOptionsAttributes(rect objc.IObject /* cross-framework Rect */, options StringDrawingOptions, attributes IDictionary /* already interface */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawWithRect:options:attributes:"), rect, options, attributes)
}


// Draws the attributed string in the specified bounding rectangle using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(with:options:attributes:context:)
func (s_ String) DrawWithRectOptionsAttributesContext(rect objc.IObject /* cross-framework Rect */, options StringDrawingOptions, attributes IDictionary /* already interface */, context objc.IObject /* cross-framework StringDrawingContext */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawWithRect:options:attributes:context:"), rect, options, attributes, context)
}


// Enumerates all the lines in the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/enumerateLines(_:)
func (s_ String) EnumerateLinesUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateLinesUsingBlock:"), block)
}


// Enumerates the substrings of the specified type in the specified range of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/enumerateSubstrings(in:options:using:)
func (s_ String) EnumerateSubstringsInRangeOptionsUsingBlock(range_ Range /* not a class type */, opts StringEnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateSubstringsInRange:options:usingBlock:"), range_, opts, block)
}


// Creates a string suitable for comparison by removing the specified character distinctions from a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/folding(options:locale:)
func (s_ String) StringByFoldingWithOptionsLocale(options StringCompareOptions, locale Locale /* not a class type */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByFoldingWithOptions:locale:"), options, locale)
	return rv
}


// Gets a given range of characters as bytes in a specified encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getBytes(_:maxLength:usedLength:encoding:options:range:remaining:)
func (s_ String) GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint /* primitive/slice/pointer. */, usedBufferCount UInteger /* not a class type */, encoding StringEncoding /* not a class type */, options StringEncodingConversionOptions, range_ Range /* not a class type */, leftover RangePointer /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("getBytes:maxLength:usedLength:encoding:options:range:remainingRange:"), buffer, maxBufferCount, usedBufferCount, encoding, options, range_, leftover)
	return rv
}


// Converts the string to a given encoding and stores it in a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:maxLength:encoding:)
func (s_ String) GetCStringMaxLengthEncoding(buffer unsafe.Pointer, maxBufferCount uint /* primitive/slice/pointer. */, encoding StringEncoding /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("getCString:maxLength:encoding:"), buffer, maxBufferCount, encoding)
	return rv
}


// Copies all characters from the receiver into a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCharacters(_:)
func (s_ String) GetCharacters(buffer unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCharacters:"), buffer)
}


// Copies characters from a given range in the receiver into a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCharacters(_:range:)
func (s_ String) GetCharactersRange(buffer unsafe.Pointer, range_ Range /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCharacters:range:"), buffer, range_)
}


// Interprets the receiver as a system-independent path and fills a buffer with a C-string in a format and encoding suitable for use with file-system calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getFileSystemRepresentation(_:maxLength:)
func (s_ String) GetFileSystemRepresentationMaxLength(cname unsafe.Pointer, max uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("getFileSystemRepresentation:maxLength:"), cname, max)
	return rv
}


// Returns by reference the beginning of the first line and the end of the last line touched by the given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getLineStart(_:end:contentsEnd:for:)
func (s_ String) GetLineStartEndContentsEndForRange(startPtr UInteger /* not a class type */, lineEndPtr UInteger /* not a class type */, contentsEndPtr UInteger /* not a class type */, range_ Range /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getLineStart:end:contentsEnd:forRange:"), startPtr, lineEndPtr, contentsEndPtr, range_)
}


// Returns by reference the beginning of the first paragraph and the end of the last paragraph touched by the given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getParagraphStart(_:end:contentsEnd:for:)
func (s_ String) GetParagraphStartEndContentsEndForRange(startPtr UInteger /* not a class type */, parEndPtr UInteger /* not a class type */, contentsEndPtr UInteger /* not a class type */, range_ Range /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getParagraphStart:end:contentsEnd:forRange:"), startPtr, parEndPtr, contentsEndPtr, range_)
}


// Returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/hasPrefix(_:)
func (s_ String) HasPrefix(str IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasPrefix:"), str)
	return rv
}


// Returns a Boolean value that indicates whether a given string matches the ending characters of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/hasSuffix(_:)
func (s_ String) HasSuffix(str IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasSuffix:"), str)
	return rv
}


// Returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/isEqual(to:)
func (s_ String) IsEqualToString(aString IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEqualToString:"), aString)
	return rv
}


// Returns the number of bytes required to store the receiver in a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lengthOfBytes(using:)
func (s_ String) LengthOfBytesUsingEncoding(enc StringEncoding /* not a class type */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](s_.ID, objc.Sel("lengthOfBytesUsingEncoding:"), enc)
	return rv
}


// Returns the range of characters representing the line or lines containing a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lineRange(for:)
func (s_ String) LineRangeForRange(range_ Range /* not a class type */) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("lineRangeForRange:"), range_)
	return rv
}


// Compares the string with a given string using a case-insensitive, localized, comparison.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCaseInsensitiveCompare(_:)
func (s_ String) LocalizedCaseInsensitiveCompare(string_ IString) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("localizedCaseInsensitiveCompare:"), string_)
	return rv
}


// Returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCaseInsensitiveContains(_:)
func (s_ String) LocalizedCaseInsensitiveContainsString(str IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("localizedCaseInsensitiveContainsString:"), str)
	return rv
}


// Compares the string and a given string using a localized comparison.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCompare(_:)
func (s_ String) LocalizedCompare(string_ IString) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("localizedCompare:"), string_)
	return rv
}


// Compares strings as sorted by the Finder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardCompare(_:)
func (s_ String) LocalizedStandardCompare(string_ IString) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("localizedStandardCompare:"), string_)
	return rv
}


// Returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardContains(_:)
func (s_ String) LocalizedStandardContainsString(str IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("localizedStandardContainsString:"), str)
	return rv
}


// Finds and returns the range of the first occurrence of a given string within the string by performing a case and diacritic insensitive, locale-aware search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardRange(of:)
func (s_ String) LocalizedStandardRangeOfString(str IString) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("localizedStandardRangeOfString:"), str)
	return rv
}


// Returns a version of the string with all letters converted to lowercase, taking into account the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lowercased(with:)
func (s_ String) LowercaseStringWithLocale(locale Locale /* not a class type */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("lowercaseStringWithLocale:"), locale)
	return rv
}


// Returns the maximum number of bytes needed to store the receiver in a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/maximumLengthOfBytes(using:)
func (s_ String) MaximumLengthOfBytesUsingEncoding(enc StringEncoding /* not a class type */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](s_.ID, objc.Sel("maximumLengthOfBytesUsingEncoding:"), enc)
	return rv
}


// Returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/padding(toLength:withPad:startingAt:)
func (s_ String) StringByPaddingToLengthWithStringStartingAtIndex(newLength uint /* primitive/slice/pointer. */, padString IString, padIndex uint /* primitive/slice/pointer. */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByPaddingToLength:withString:startingAtIndex:"), newLength, padString, padIndex)
	return rv
}


// Returns the range of characters representing the paragraph or paragraphs containing a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/paragraphRange(for:)
func (s_ String) ParagraphRangeForRange(range_ Range /* not a class type */) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("paragraphRangeForRange:"), range_)
	return rv
}


// Parses the receiver as a text representation of a property list, returning an , , , or object, according to the topmost element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/propertyList()
func (s_ String) PropertyList() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("propertyList"))
	return rv
}


// Returns a dictionary object initialized with the keys and values found in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/propertyListFromStringsFileFormat()
func (s_ String) PropertyListFromStringsFileFormat() IDictionary {
	rv := objc.Send[Dictionary](s_.ID, objc.Sel("propertyListFromStringsFileFormat"))
	return rv
}


// Finds and returns the range of the first occurrence of a given string within the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:)
func (s_ String) RangeOfString(searchString IString) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfString:"), searchString)
	return rv
}


// Finds and returns the range of the first occurrence of a given string within the string, subject to given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:)
func (s_ String) RangeOfStringOptions(searchString IString, mask StringCompareOptions) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfString:options:"), searchString, mask)
	return rv
}


// Finds and returns the range of the first occurrence of a given string, within the given range of the string, subject to given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:range:)
func (s_ String) RangeOfStringOptionsRange(searchString IString, mask StringCompareOptions, rangeOfReceiverToSearch Range /* not a class type */) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfString:options:range:"), searchString, mask, rangeOfReceiverToSearch)
	return rv
}


// Finds and returns the range of the first occurrence of a given string within a given range of the string, subject to given options, using the specified locale, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:range:locale:)
func (s_ String) RangeOfStringOptionsRangeLocale(searchString IString, mask StringCompareOptions, rangeOfReceiverToSearch Range /* not a class type */, locale Locale /* not a class type */) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfString:options:range:locale:"), searchString, mask, rangeOfReceiverToSearch, locale)
	return rv
}


// Finds and returns the range in the string of the first character from a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:)
func (s_ String) RangeOfCharacterFromSet(searchSet CharacterSet /* not a class type */) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfCharacterFromSet:"), searchSet)
	return rv
}


// Finds and returns the range in the string of the first character, using given options, from a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:options:)
func (s_ String) RangeOfCharacterFromSetOptions(searchSet CharacterSet /* not a class type */, mask StringCompareOptions) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfCharacterFromSet:options:"), searchSet, mask)
	return rv
}


// Finds and returns the range in the string of the first character from a given character set found in a given range with given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:options:range:)
func (s_ String) RangeOfCharacterFromSetOptionsRange(searchSet CharacterSet /* not a class type */, mask StringCompareOptions, rangeOfReceiverToSearch Range /* not a class type */) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfCharacterFromSet:options:range:"), searchSet, mask, rangeOfReceiverToSearch)
	return rv
}


// Returns the range in the receiver of the composed character sequence located at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfComposedCharacterSequence(at:)
func (s_ String) RangeOfComposedCharacterSequenceAtIndex(index uint /* primitive/slice/pointer. */) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfComposedCharacterSequenceAtIndex:"), index)
	return rv
}


// Returns the range in the string of the composed character sequences for a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfComposedCharacterSequences(for:)
func (s_ String) RangeOfComposedCharacterSequencesForRange(range_ Range /* not a class type */) Range /* not a class type */ {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfComposedCharacterSequencesForRange:"), range_)
	return rv
}


// Returns a new string in which the characters in a specified range of the receiver are replaced by a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingCharacters(in:with:)
func (s_ String) StringByReplacingCharactersInRangeWithString(range_ Range /* not a class type */, replacement IString) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByReplacingCharactersInRange:withString:"), range_, replacement)
	return rv
}


// Returns a new string in which all occurrences of a target string in the receiver are replaced by another given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingOccurrences(of:with:)
func (s_ String) StringByReplacingOccurrencesOfStringWithString(target IString, replacement IString) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByReplacingOccurrencesOfString:withString:"), target, replacement)
	return rv
}


// Returns a new string in which all occurrences of a target string in a specified range of the receiver are replaced by another given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingOccurrences(of:with:options:range:)
func (s_ String) StringByReplacingOccurrencesOfStringWithStringOptionsRange(target IString, replacement IString, options StringCompareOptions, searchRange Range /* not a class type */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByReplacingOccurrencesOfString:withString:options:range:"), target, replacement, options, searchRange)
	return rv
}


// Returns the bounding box size the receiver occupies when drawn with the given attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/size(withAttributes:)
func (s_ String) SizeWithAttributes(attrs IDictionary /* already interface */) objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](s_.ID, objc.Sel("sizeWithAttributes:"), attrs)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sr_sensorForDeletionRecordsFromSensor()
func (s_ String) Sr_sensorForDeletionRecordsFromSensor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sr_sensorForDeletionRecordsFromSensor"))
	return rv
}


// Returns a string made by appending to the receiver a string constructed from a given format string and the following arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringByAppendingFormat:
func (s_ String) StringByAppendingFormat(format IString) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByAppendingFormat:"), format)
	return rv
}


// Returns an array of strings made by separately appending to the receiver each string in a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/strings(byAppendingPaths:)
func (s_ String) StringsByAppendingPaths(paths []string /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("stringsByAppendingPaths:"), paths)
	return rv
}


// Returns a new string containing the characters of the receiver from the one at a given index to the end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/substring(from:)
func (s_ String) SubstringFromIndex(from uint /* primitive/slice/pointer. */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("substringFromIndex:"), from)
	return rv
}


// Returns a new string containing the characters of the receiver up to, but not including, the one at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/substring(to:)
func (s_ String) SubstringToIndex(to uint /* primitive/slice/pointer. */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("substringToIndex:"), to)
	return rv
}


// Returns a string object containing the characters of the receiver that lie within a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/substring(with:)
func (s_ String) SubstringWithRange(range_ Range /* not a class type */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("substringWithRange:"), range_)
	return rv
}


// Returns a new string made by removing from both ends of the receiver characters contained in a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/trimmingCharacters(in:)
func (s_ String) StringByTrimmingCharactersInSet(set CharacterSet /* not a class type */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByTrimmingCharactersInSet:"), set)
	return rv
}


// Returns a version of the string with all letters converted to uppercase, taking into account the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/uppercased(with:)
func (s_ String) UppercaseStringWithLocale(locale Locale /* not a class type */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("uppercaseStringWithLocale:"), locale)
	return rv
}


// Returns a string variation suitable for the specified presentation width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/variantFittingPresentationWidth(_:)
func (s_ String) VariantFittingPresentationWidth(width int /* primitive/slice/pointer. */) IString {
	rv := objc.Send[String](s_.ID, objc.Sel("variantFittingPresentationWidth:"), width)
	return rv
}


// Writes the contents of the receiver to the URL specified by using the specified encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(to:atomically:encoding:)
func (s_ String) WriteToURLAtomicallyEncodingError(url objc.IObject /* cross-framework NSURL */, useAuxiliaryFile bool /* primitive/slice/pointer. */, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToURL:atomically:encoding:error:"), url, useAuxiliaryFile, enc, error_)
	return rv
}


// Writes the contents of the receiver to a file at a given path using a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(toFile:atomically:encoding:)
func (s_ String) WriteToFileAtomicallyEncodingError(path IString, useAuxiliaryFile bool /* primitive/slice/pointer. */, enc StringEncoding /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToFile:atomically:encoding:error:"), path, useAuxiliaryFile, enc, error_)
	return rv
}


// A new string that replaces the current home directory portion of the current path with a tilde ( ) character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/abbreviatingWithTildeInPath
func (s_ String) StringByAbbreviatingWithTildeInPath() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByAbbreviatingWithTildeInPath"))
	return rv
}


// Returns a zero-terminated list of the encodings string objects support in the application’s environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/availableStringEncodings
func (s_ String) AvailableStringEncodings() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](s_.ID, objc.Sel("availableStringEncodings"))
	return rv
}


// The Boolean value of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/boolValue
func (s_ String) BoolValue() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("boolValue"))
	return rv
}


// A capitalized representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/capitalized
func (s_ String) CapitalizedString() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("capitalizedString"))
	return rv
}


// A string made by normalizing the string’s contents using the Unicode Normalization Form D.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/decomposedStringWithCanonicalMapping
func (s_ String) DecomposedStringWithCanonicalMapping() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("decomposedStringWithCanonicalMapping"))
	return rv
}


// A string made by normalizing the receiver’s contents using the Unicode Normalization Form KD.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/decomposedStringWithCompatibilityMapping
func (s_ String) DecomposedStringWithCompatibilityMapping() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("decomposedStringWithCompatibilityMapping"))
	return rv
}


// Returns the C-string encoding assumed for any method accepting a C string as an argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/defaultCStringEncoding
func (s_ String) DefaultCStringEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](s_.ID, objc.Sel("defaultCStringEncoding"))
	return rv
}


// A new string made by deleting the last path component from the receiver, along with any final path separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deletingLastPathComponent
func (s_ String) StringByDeletingLastPathComponent() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByDeletingLastPathComponent"))
	return rv
}


// A new string made by deleting the extension (if any, and only the last) from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deletingPathExtension
func (s_ String) StringByDeletingPathExtension() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByDeletingPathExtension"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/description
func (s_ String) Description() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("description"))
	return rv
}


// The floating-point value of the string as a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/doubleValue
func (s_ String) DoubleValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("doubleValue"))
	return rv
}


// A new string made by expanding the initial component of the receiver to its full path value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/expandingTildeInPath
func (s_ String) StringByExpandingTildeInPath() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByExpandingTildeInPath"))
	return rv
}


// The fastest encoding to which the receiver may be converted without loss of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/fastestEncoding
func (s_ String) FastestEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](s_.ID, objc.Sel("fastestEncoding"))
	return rv
}


// A file system-specific representation of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/fileSystemRepresentation
func (s_ String) FileSystemRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("fileSystemRepresentation"))
	return rv
}


// The floating-point value of the string as a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/floatValue
func (s_ String) FloatValue() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("floatValue"))
	return rv
}


// An unsigned integer that can be used as a hash table address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/hash
func (s_ String) Hash() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](s_.ID, objc.Sel("hash"))
	return rv
}


// The integer value of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/intValue
func (s_ String) IntValue() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("intValue"))
	return rv
}


// The value of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/integerValue
func (s_ String) IntegerValue() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("integerValue"))
	return rv
}


// A Boolean value that indicates whether the receiver represents an absolute path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/isAbsolutePath
func (s_ String) AbsolutePath() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("absolutePath"))
	return rv
}


// The last path component of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lastPathComponent
func (s_ String) LastPathComponent() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("lastPathComponent"))
	return rv
}


// The number of UTF-16 code units in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/length
func (s_ String) Length() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](s_.ID, objc.Sel("length"))
	return rv
}


// Returns a capitalized representation of the receiver using the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCapitalized
func (s_ String) LocalizedCapitalizedString() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("localizedCapitalizedString"))
	return rv
}


// Returns a version of the string with all letters converted to lowercase, taking into account the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedLowercase
func (s_ String) LocalizedLowercaseString() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("localizedLowercaseString"))
	return rv
}


// Returns a version of the string with all letters converted to uppercase, taking into account the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedUppercase
func (s_ String) LocalizedUppercaseString() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("localizedUppercaseString"))
	return rv
}


// The value of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/longLongValue
func (s_ String) LongLongValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("longLongValue"))
	return rv
}


// A lowercase representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lowercased
func (s_ String) LowercaseString() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("lowercaseString"))
	return rv
}


// The file-system path components of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/pathComponents
func (s_ String) PathComponents() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("pathComponents"))
	return rv
}


// The path extension, if any, of the string as interpreted as a path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/pathExtension
func (s_ String) PathExtension() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("pathExtension"))
	return rv
}


// A string made by normalizing the string’s contents using the Unicode Normalization Form C.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/precomposedStringWithCanonicalMapping
func (s_ String) PrecomposedStringWithCanonicalMapping() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("precomposedStringWithCanonicalMapping"))
	return rv
}


// A string made by normalizing the receiver’s contents using the Unicode Normalization Form KC.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/precomposedStringWithCompatibilityMapping
func (s_ String) PrecomposedStringWithCompatibilityMapping() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("precomposedStringWithCompatibilityMapping"))
	return rv
}


// Returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/removingPercentEncoding
func (s_ String) StringByRemovingPercentEncoding() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByRemovingPercentEncoding"))
	return rv
}


// A new string made from the receiver by resolving all symbolic links and standardizing path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/resolvingSymlinksInPath
func (s_ String) StringByResolvingSymlinksInPath() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByResolvingSymlinksInPath"))
	return rv
}


// The smallest encoding to which the receiver can be converted without loss of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/smallestEncoding
func (s_ String) SmallestEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](s_.ID, objc.Sel("smallestEncoding"))
	return rv
}


// A new string made by removing extraneous path components from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/standardizingPath
func (s_ String) StringByStandardizingPath() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("stringByStandardizingPath"))
	return rv
}


// An uppercase representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/uppercased
func (s_ String) UppercaseString() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("uppercaseString"))
	return rv
}


// A null-terminated UTF8 representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/utf8String
func (s_ String) UTF8String() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("UTF8String"))
	return rv
}


// A new string that replaces the current home directory portion of the current path with a tilde (
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/abbreviatingwithtildeinpath
func (s_ String) AbbreviatingWithTildeInPath() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("abbreviatingWithTildeInPath"))
	return rv
}


// A new string that replaces the current home directory portion of the current path with a tilde (
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/abbreviatingwithtildeinpath
func (s_ String) SetAbbreviatingWithTildeInPath(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAbbreviatingWithTildeInPath:"), value)
}


// A capitalized representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/capitalized
func (s_ String) Capitalized() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("capitalized"))
	return rv
}


// A capitalized representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/capitalized
func (s_ String) SetCapitalized(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCapitalized:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/customplaygroundquicklook
func (s_ String) CustomPlaygroundQuickLook() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("customPlaygroundQuickLook"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/customplaygroundquicklook
func (s_ String) SetCustomPlaygroundQuickLook(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomPlaygroundQuickLook:"), value)
}


// A new string made by deleting the last path component from the receiver, along with any final path separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/deletinglastpathcomponent
func (s_ String) DeletingLastPathComponent() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("deletingLastPathComponent"))
	return rv
}


// A new string made by deleting the last path component from the receiver, along with any final path separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/deletinglastpathcomponent
func (s_ String) SetDeletingLastPathComponent(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDeletingLastPathComponent:"), value)
}


// A new string made by deleting the extension (if any, and only the last) from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/deletingpathextension
func (s_ String) DeletingPathExtension() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("deletingPathExtension"))
	return rv
}


// A new string made by deleting the extension (if any, and only the last) from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/deletingpathextension
func (s_ String) SetDeletingPathExtension(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDeletingPathExtension:"), value)
}


// A new string made by expanding the initial component of the receiver to its full path value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/expandingtildeinpath
func (s_ String) ExpandingTildeInPath() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("expandingTildeInPath"))
	return rv
}


// A new string made by expanding the initial component of the receiver to its full path value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/expandingtildeinpath
func (s_ String) SetExpandingTildeInPath(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setExpandingTildeInPath:"), value)
}


// A Boolean value that indicates whether the receiver represents an absolute path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/isabsolutepath
func (s_ String) IsAbsolutePath() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAbsolutePath"))
	return rv
}


// A Boolean value that indicates whether the receiver represents an absolute path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/isabsolutepath
func (s_ String) SetIsAbsolutePath(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAbsolutePath:"), value)
}


// Returns a capitalized representation of the receiver using the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/localizedcapitalized
func (s_ String) LocalizedCapitalized() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("localizedCapitalized"))
	return rv
}


// Returns a capitalized representation of the receiver using the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/localizedcapitalized
func (s_ String) SetLocalizedCapitalized(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLocalizedCapitalized:"), value)
}


// Returns a version of the string with all letters converted to lowercase, taking into account the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/localizedlowercase
func (s_ String) LocalizedLowercase() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("localizedLowercase"))
	return rv
}


// Returns a version of the string with all letters converted to lowercase, taking into account the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/localizedlowercase
func (s_ String) SetLocalizedLowercase(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLocalizedLowercase:"), value)
}


// Returns a version of the string with all letters converted to uppercase, taking into account the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/localizeduppercase
func (s_ String) LocalizedUppercase() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("localizedUppercase"))
	return rv
}


// Returns a version of the string with all letters converted to uppercase, taking into account the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/localizeduppercase
func (s_ String) SetLocalizedUppercase(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLocalizedUppercase:"), value)
}


// A lowercase representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/lowercased
func (s_ String) Lowercased() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("lowercased"))
	return rv
}


// A lowercase representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/lowercased
func (s_ String) SetLowercased(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLowercased:"), value)
}


// Returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/removingpercentencoding
func (s_ String) RemovingPercentEncoding() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("removingPercentEncoding"))
	return rv
}


// Returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/removingpercentencoding
func (s_ String) SetRemovingPercentEncoding(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRemovingPercentEncoding:"), value)
}


// A new string made from the receiver by resolving all symbolic links and standardizing path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/resolvingsymlinksinpath
func (s_ String) ResolvingSymlinksInPath() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("resolvingSymlinksInPath"))
	return rv
}


// A new string made from the receiver by resolving all symbolic links and standardizing path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/resolvingsymlinksinpath
func (s_ String) SetResolvingSymlinksInPath(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResolvingSymlinksInPath:"), value)
}


// A new string made by removing extraneous path components from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/standardizingpath
func (s_ String) StandardizingPath() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("standardizingPath"))
	return rv
}


// A new string made by removing extraneous path components from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/standardizingpath
func (s_ String) SetStandardizingPath(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStandardizingPath:"), value)
}


// An uppercase representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/uppercased
func (s_ String) Uppercased() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("uppercased"))
	return rv
}


// An uppercase representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/uppercased
func (s_ String) SetUppercased(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUppercased:"), value)
}


