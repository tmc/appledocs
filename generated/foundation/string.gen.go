// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [String] class.
var stringClass = _StringClass{objc.GetClass("NSString")}

type _StringClass struct {
	class objc.Class
}

// An interface definition for the [String] class.
type IString interface {
	objectivec.IObject
	StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters unsafe.Pointer) unsafe.Pointer
	StringByAddingPercentEscapesUsingEncoding(enc unsafe.Pointer) unsafe.Pointer
	StringByAppendingString(aString string) unsafe.Pointer
	StringByAppendingPathComponent(str string) unsafe.Pointer
	StringByAppendingPathComponentConformingToType(partialName string, contentType unsafe.Pointer) unsafe.Pointer
	StringByAppendingPathExtension(str string) unsafe.Pointer
	StringByAppendingPathExtensionForType(contentType unsafe.Pointer) unsafe.Pointer
	StringByApplyingTransformReverse(transform unsafe.Pointer, reverse bool) unsafe.Pointer
	BoundingRectWithSizeOptionsAttributes(size unsafe.Pointer, options unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer
	BoundingRectWithSizeOptionsAttributesContext(size unsafe.Pointer, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer
	CString() unsafe.Pointer
	CStringUsingEncoding(encoding unsafe.Pointer) unsafe.Pointer
	CStringLength() uint
	CanBeConvertedToEncoding(encoding unsafe.Pointer) bool
	CapitalizedStringWithLocale(locale unsafe.Pointer) unsafe.Pointer
	CaseInsensitiveCompare(string string) unsafe.Pointer
	CharacterAtIndex(index uint) unsafe.Pointer
	CommonPrefixWithStringOptions(str string, mask unsafe.Pointer) unsafe.Pointer
	Compare(string string) unsafe.Pointer
	CompareOptions(string string, mask unsafe.Pointer) unsafe.Pointer
	CompareOptionsRange(string string, mask unsafe.Pointer, rangeOfReceiverToCompare unsafe.Pointer) unsafe.Pointer
	CompareOptionsRangeLocale(string string, mask unsafe.Pointer, rangeOfReceiverToCompare unsafe.Pointer, locale objc.ID) unsafe.Pointer
	CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName string, flag bool, outputArray unsafe.Pointer, filterTypes unsafe.Pointer) uint
	ComponentsSeparatedByString(separator string) unsafe.Pointer
	ComponentsSeparatedByCharactersInSet(separator unsafe.Pointer) unsafe.Pointer
	ContainsString(str string) bool
	DataUsingEncoding(encoding unsafe.Pointer) unsafe.Pointer
	DataUsingEncodingAllowLossyConversion(encoding unsafe.Pointer, lossy bool) unsafe.Pointer
	DrawAtPointWithAttributes(point unsafe.Pointer, attrs unsafe.Pointer)
	DrawInRectWithAttributes(rect unsafe.Pointer, attrs unsafe.Pointer)
	DrawWithRectOptionsAttributes(rect unsafe.Pointer, options unsafe.Pointer, attributes unsafe.Pointer)
	DrawWithRectOptionsAttributesContext(rect unsafe.Pointer, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer)
	DrawAtPointForWidthWithFontFontSizeLineBreakModeBaselineAdjustment(point unsafe.Pointer, width float64, font unsafe.Pointer, fontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) unsafe.Pointer
	DrawAtPointForWidthWithFontLineBreakMode(point unsafe.Pointer, width float64, font unsafe.Pointer, lineBreakMode unsafe.Pointer) unsafe.Pointer
	DrawAtPointForWidthWithFontMinFontSizeActualFontSizeLineBreakModeBaselineAdjustment(point unsafe.Pointer, width float64, font unsafe.Pointer, minFontSize float64, actualFontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) unsafe.Pointer
	DrawAtPointWithFont(point unsafe.Pointer, font unsafe.Pointer) unsafe.Pointer
	DrawInRectWithFont(rect unsafe.Pointer, font unsafe.Pointer) unsafe.Pointer
	DrawInRectWithFontLineBreakMode(rect unsafe.Pointer, font unsafe.Pointer, lineBreakMode unsafe.Pointer) unsafe.Pointer
	DrawInRectWithFontLineBreakModeAlignment(rect unsafe.Pointer, font unsafe.Pointer, lineBreakMode unsafe.Pointer, alignment unsafe.Pointer) unsafe.Pointer
	EnumerateLinesUsingBlock(block unsafe.Pointer)
	EnumerateLinguisticTagsInRangeSchemeOptionsOrthographyUsingBlock(range_ unsafe.Pointer, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, block unsafe.Pointer)
	EnumerateSubstringsInRangeOptionsUsingBlock(range_ unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer)
	StringByFoldingWithOptionsLocale(options unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer
	GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint, usedBufferCount unsafe.Pointer, encoding unsafe.Pointer, options unsafe.Pointer, range_ unsafe.Pointer, leftover unsafe.Pointer) bool
	GetCString(bytes unsafe.Pointer)
	GetCStringMaxLength(bytes unsafe.Pointer, maxLength uint)
	GetCStringMaxLengthEncoding(buffer unsafe.Pointer, maxBufferCount uint, encoding unsafe.Pointer) bool
	GetCStringMaxLengthRangeRemainingRange(bytes unsafe.Pointer, maxLength uint, aRange unsafe.Pointer, leftoverRange unsafe.Pointer)
	GetCharacters(buffer unsafe.Pointer)
	GetCharactersRange(buffer unsafe.Pointer, range_ unsafe.Pointer)
	GetFileSystemRepresentationMaxLength(cname unsafe.Pointer, max uint) bool
	GetLineStartEndContentsEndForRange(startPtr unsafe.Pointer, lineEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ unsafe.Pointer)
	GetParagraphStartEndContentsEndForRange(startPtr unsafe.Pointer, parEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ unsafe.Pointer)
	HasPrefix(str string) bool
	HasSuffix(str string) bool
	IsEqualToString(aString string) bool
	LengthOfBytesUsingEncoding(enc unsafe.Pointer) uint
	LineRangeForRange(range_ unsafe.Pointer) unsafe.Pointer
	LinguisticTagsInRangeSchemeOptionsOrthographyTokenRanges(range_ unsafe.Pointer, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, tokenRanges unsafe.Pointer) unsafe.Pointer
	LocalizedCaseInsensitiveCompare(string string) unsafe.Pointer
	LocalizedCaseInsensitiveContainsString(str string) bool
	LocalizedCompare(string string) unsafe.Pointer
	LocalizedStandardCompare(string string) unsafe.Pointer
	LocalizedStandardContainsString(str string) bool
	LocalizedStandardRangeOfString(str string) unsafe.Pointer
	LossyCString() unsafe.Pointer
	LowercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer
	MaximumLengthOfBytesUsingEncoding(enc unsafe.Pointer) uint
	StringByPaddingToLengthWithStringStartingAtIndex(newLength uint, padString string, padIndex uint) unsafe.Pointer
	ParagraphRangeForRange(range_ unsafe.Pointer) unsafe.Pointer
	PropertyList() objc.ID
	PropertyListFromStringsFileFormat() unsafe.Pointer
	RangeOfString(searchString string) unsafe.Pointer
	RangeOfStringOptions(searchString string, mask unsafe.Pointer) unsafe.Pointer
	RangeOfStringOptionsRange(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch unsafe.Pointer) unsafe.Pointer
	RangeOfStringOptionsRangeLocale(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer
	RangeOfCharacterFromSet(searchSet unsafe.Pointer) unsafe.Pointer
	RangeOfCharacterFromSetOptions(searchSet unsafe.Pointer, mask unsafe.Pointer) unsafe.Pointer
	RangeOfCharacterFromSetOptionsRange(searchSet unsafe.Pointer, mask unsafe.Pointer, rangeOfReceiverToSearch unsafe.Pointer) unsafe.Pointer
	RangeOfComposedCharacterSequenceAtIndex(index uint) unsafe.Pointer
	RangeOfComposedCharacterSequencesForRange(range_ unsafe.Pointer) unsafe.Pointer
	StringByReplacingCharactersInRangeWithString(range_ unsafe.Pointer, replacement string) unsafe.Pointer
	StringByReplacingOccurrencesOfStringWithString(target string, replacement string) unsafe.Pointer
	StringByReplacingOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options unsafe.Pointer, searchRange unsafe.Pointer) unsafe.Pointer
	StringByReplacingPercentEscapesUsingEncoding(enc unsafe.Pointer) unsafe.Pointer
	SizeWithAttributes(attrs unsafe.Pointer) unsafe.Pointer
	SizeWithFont(font unsafe.Pointer) unsafe.Pointer
	SizeWithFontConstrainedToSize(font unsafe.Pointer, size unsafe.Pointer) unsafe.Pointer
	SizeWithFontConstrainedToSizeLineBreakMode(font unsafe.Pointer, size unsafe.Pointer, lineBreakMode unsafe.Pointer) unsafe.Pointer
	SizeWithFontForWidthLineBreakMode(font unsafe.Pointer, width float64, lineBreakMode unsafe.Pointer) unsafe.Pointer
	SizeWithFontMinFontSizeActualFontSizeForWidthLineBreakMode(font unsafe.Pointer, minFontSize float64, actualFontSize float64, width float64, lineBreakMode unsafe.Pointer) unsafe.Pointer
	Sr_sensorForDeletionRecordsFromSensor() unsafe.Pointer
	StringByAppendingFormat(format string) unsafe.Pointer
	StringsByAppendingPaths(paths unsafe.Pointer) unsafe.Pointer
	SubstringFromIndex(from uint) unsafe.Pointer
	SubstringToIndex(to uint) unsafe.Pointer
	SubstringWithRange(range_ unsafe.Pointer) unsafe.Pointer
	StringByTrimmingCharactersInSet(set unsafe.Pointer) unsafe.Pointer
	UppercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer
	VariantFittingPresentationWidth(width int) unsafe.Pointer
	WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool
	WriteToURLAtomicallyEncodingError(url unsafe.Pointer, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool
	WriteToFileAtomically(path string, useAuxiliaryFile bool) bool
	WriteToFileAtomicallyEncodingError(path string, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool
}

// A static, plain-text Unicode string object. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return stringClass.New()
}
// Returns an object initialized using the characters in a given C array, interpreted according to a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:encoding:)-20f9h
func NewStringWithCStringEncoding(nullTerminatedCString unsafe.Pointer, encoding unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCString:encoding:"), nullTerminatedCString, encoding)
	rv.Autorelease()
	return rv
}
// Initializes the receiver, a newly allocated object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:length:)
func NewStringWithCStringLength(bytes unsafe.Pointer, length uint) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCString:length:"), bytes, length)
	rv.Autorelease()
	return rv
}
// Returns an initialized object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(bytes:length:encoding:)
func NewStringWithBytesLengthEncoding(bytes unsafe.Pointer, len uint, encoding unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithBytes:length:encoding:"), bytes, len, encoding)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:deallocator:)
func NewStringWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len uint, encoding unsafe.Pointer, deallocator unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:deallocator:"), bytes, len, encoding, deallocator)
	rv.Autorelease()
	return rv
}
// Returns an initialized object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:freeWhenDone:)
func NewStringWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len uint, encoding unsafe.Pointer, freeBuffer bool) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:freeWhenDone:"), bytes, len, encoding, freeBuffer)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by reading data from the file at a given path using a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:encoding:)
func NewStringWithContentsOfFileEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfFile:encoding:error:"), path, enc, error)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale information. This method is meant to be called from within a variadic function, where the argument list will be available. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(format:locale:arguments:)
func NewStringWithFormatLocaleArguments(format string, locale objc.ID, argList unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:locale:arguments:"), format, locale, argList)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersArgumentsError(format string, validFormatSpecifiers string, argList unsafe.Pointer, error unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:arguments:error:"), format, validFormatSpecifiers, argList, error)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:locale:
func NewStringWithFormatLocale(format string, locale objc.ID) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:locale:"), format, locale)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:error:
func NewStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string, error unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:error:"), format, validFormatSpecifiers, error)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format string, validFormatSpecifiers string, locale objc.ID, argList unsafe.Pointer, error unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:"), format, validFormatSpecifiers, locale, argList, error)
	rv.Autorelease()
	return rv
}
// Returns an initialized object that contains a given number of characters from a given C array of UTF-16 code units. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(characters:length:)
func NewStringWithCharactersLength(characters unsafe.Pointer, length uint) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCharacters:length:"), characters, length)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:deallocator:)
func NewStringWithCharactersNoCopyLengthDeallocator(chars unsafe.Pointer, len uint, deallocator unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCharactersNoCopy:length:deallocator:"), chars, len, deallocator)
	rv.Autorelease()
	return rv
}
// Returns an initialized object that contains a given number of characters from a given C array of UTF-16 code units. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:freeWhenDone:)
func NewStringWithCharactersNoCopyLengthFreeWhenDone(characters unsafe.Pointer, length uint, freeBuffer bool) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCharactersNoCopy:length:freeWhenDone:"), characters, length, freeBuffer)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by reading data from the file at a given path and returns by reference the encoding used to interpret the characters. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:usedEncoding:)
func NewStringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfFile:usedEncoding:error:"), path, enc, error)
	rv.Autorelease()
	return rv
}
// Initializes the receiver, a newly allocated object, by reading data from the location named by a given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:)
func NewStringWithContentsOfURL(url unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted without any localization. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(format:arguments:)
func NewStringWithFormatArguments(format string, argList unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:arguments:"), format, argList)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:
func NewStringWithFormat(format string) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:"), format)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by reading data from a given URL and returns by reference the encoding used to interpret the data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:usedEncoding:)-2c72d
func NewStringWithContentsOfURLUsedEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfURL:usedEncoding:error:"), url, enc, error)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by copying the characters from another given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(string:)-210xa
func NewStringWithString(aString string) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithString:"), aString)
	rv.Autorelease()
	return rv
}
// Initializes the receiver, a newly allocated object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:)
func NewStringWithCString(bytes unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCString:"), bytes)
	rv.Autorelease()
	return rv
}
// Initializes the receiver, a newly allocated object, by reading data from the file named by . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:)
func NewStringWithContentsOfFile(path string) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfFile:"), path)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by converting given data into UTF-16 code units using a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(data:encoding:)
func NewStringWithDataEncoding(data unsafe.Pointer, encoding unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithData:encoding:"), data, encoding)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by reading data from a given URL interpreted using a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:encoding:)-715fw
func NewStringWithContentsOfURLEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfURL:encoding:error:"), url, enc, error)
	rv.Autorelease()
	return rv
}
// Initializes the receiver, a newly allocated object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CStringNoCopy:length:freeWhenDone:)
func NewStringWithCStringNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, freeBuffer bool) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCStringNoCopy:length:freeWhenDone:"), bytes, length, freeBuffer)
	rv.Autorelease()
	return rv
}
// Returns an object initialized by copying the characters from a given C array of UTF8-encoded bytes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(UTF8String:)-vg2b
func NewStringWithUTF8String(nullTerminatedCString unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithUTF8String:"), nullTerminatedCString)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(coder:)
func NewStringWithCoder(coder unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleError(format string, validFormatSpecifiers string, locale objc.ID, error unsafe.Pointer) String {
	instance := stringClass.Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:error:"), format, validFormatSpecifiers, locale, error)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormat(format string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:"), format)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormatFromTable(format string, table string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:fromTable:"), format, table)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:arguments:
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormatFromTableArguments(format string, table string, arguments unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:fromTable:arguments:"), format, table, arguments)
	return rv
}
// Returns a string containing the bytes in a given C array, interpreted according to a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:encoding:)-7auq8
func (sc _StringClass) StringWithCStringEncoding(cString unsafe.Pointer, enc unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithCString:encoding:"), cString, enc)
	return rv
}
// Returns a string created by copying the data from a given C array of UTF8-encoded bytes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(UTF8String:)-8bcy8
func (sc _StringClass) StringWithUTF8String(nullTerminatedCString unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithUTF8String:"), nullTerminatedCString)
	return rv
}
// Returns a string created by reading data from a given URL interpreted using a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:encoding:)-x6cv
func (sc _StringClass) StringWithContentsOfURLEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfURL:encoding:error:"), url, enc, error)
	return rv
}
// Returns a string created by reading data from a given URL and returns by reference the encoding used to interpret the data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:usedEncoding:)-9jrum
func (sc _StringClass) StringWithContentsOfURLUsedEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfURL:usedEncoding:error:"), url, enc, error)
	return rv
}
// Returns a human-readable string giving the name of a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedName(of:)
func (sc _StringClass) LocalizedNameOfStringEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedNameOfStringEncoding:"), encoding)
	return rv
}
// Returns a string created by using a given format string as a template into which the remaining argument values are substituted according to the current locale. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStringWithFormat:
func (sc _StringClass) LocalizedStringWithFormat(format string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedStringWithFormat:"), format)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStringWithValidatedFormat:validFormatSpecifiers:error:
func (sc _StringClass) LocalizedStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedStringWithValidatedFormat:validFormatSpecifiers:error:"), format, validFormatSpecifiers, error)
	return rv
}
// Returns a localized string intended for display in a notification alert. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedUserNotificationString(forKey:arguments:)
func (sc _StringClass) LocalizedUserNotificationStringForKeyArguments(key string, arguments unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedUserNotificationStringForKey:arguments:"), key, arguments)
	return rv
}
// Returns a string built from the strings in a given array by concatenating them with a path separator between each pair. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/path(withComponents:)
func (sc _StringClass) PathWithComponents(components unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("pathWithComponents:"), components)
	return rv
}
// Returns an empty string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string
func (sc _StringClass) String() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("string"))
	return rv
}
// Creates a new string using a given C-string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withCString:)
func (sc _StringClass) StringWithCString(bytes unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithCString:"), bytes)
	return rv
}
// Returns a string containing the characters in a given C-string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withCString:length:)
func (sc _StringClass) StringWithCStringLength(bytes unsafe.Pointer, length uint) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithCString:length:"), bytes, length)
	return rv
}
// Returns a string created by reading data from the file named by a given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withContentsOf:)
func (sc _StringClass) StringWithContentsOfURL(url unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithContentsOfURL:"), url)
	return rv
}
// Returns a string created by reading data from the file named by a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withContentsOfFile:)
func (sc _StringClass) StringWithContentsOfFile(path string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithContentsOfFile:"), path)
	return rv
}
// Returns the string encoding for the given data as detected by attempting to create a string according to the specified encoding options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringEncoding(for:encodingOptions:convertedString:usedLossyConversion:)
func (sc _StringClass) StringEncodingForDataEncodingOptionsConvertedStringUsedLossyConversion(data unsafe.Pointer, opts unsafe.Pointer, string string, usedLossyConversion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringEncodingForData:encodingOptions:convertedString:usedLossyConversion:"), data, opts, string, usedLossyConversion)
	return rv
}
// Returns a string containing a given number of characters taken from a given C array of UTF-16 code units. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithCharacters:length:
func (sc _StringClass) StringWithCharactersLength(characters unsafe.Pointer, length uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithCharacters:length:"), characters, length)
	return rv
}
// Returns a string created by reading data from the file at a given path interpreted using a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfFile:encoding:error:
func (sc _StringClass) StringWithContentsOfFileEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfFile:encoding:error:"), path, enc, error)
	return rv
}
// Returns a string created by reading data from the file at a given path and returns by reference the encoding used to interpret the file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfFile:usedEncoding:error:
func (sc _StringClass) StringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfFile:usedEncoding:error:"), path, enc, error)
	return rv
}
// Returns a string created by using a given format string as a template into which the remaining argument values are substituted. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithFormat:
func (sc _StringClass) StringWithFormat(format string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithFormat:"), format)
	return rv
}
// Returns a string created by copying the characters from another given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithString:
func (sc _StringClass) StringWithString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithString:"), string)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithValidatedFormat:validFormatSpecifiers:error:
func (sc _StringClass) StringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithValidatedFormat:validFormatSpecifiers:error:"), format, validFormatSpecifiers, error)
	return rv
}
// Returns a new string made from the receiver by replacing all characters not in the specified set with percent-encoded characters. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/addingPercentEncoding(withAllowedCharacters:)
func (s_ String) StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAddingPercentEncodingWithAllowedCharacters:"), allowedCharacters)
	return rv
}
// Returns a representation of the receiver using a given encoding to determine the percent escapes necessary to convert the receiver into a legal URL string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/addingPercentEscapes(using:)
func (s_ String) StringByAddingPercentEscapesUsingEncoding(enc unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAddingPercentEscapesUsingEncoding:"), enc)
	return rv
}
// Returns a new string made by appending a given string to the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appending(_:)
func (s_ String) StringByAppendingString(aString string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingString:"), aString)
	return rv
}
// Returns a new string made by appending to the receiver a given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathComponent(_:)
func (s_ String) StringByAppendingPathComponent(str string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingPathComponent:"), str)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathComponent(_:conformingTo:)
func (s_ String) StringByAppendingPathComponentConformingToType(partialName string, contentType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingPathComponent:conformingToType:"), partialName, contentType)
	return rv
}
// Returns a new string made by appending to the receiver an extension separator followed by a given extension. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathExtension(_:)
func (s_ String) StringByAppendingPathExtension(str string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingPathExtension:"), str)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathExtension(for:)
func (s_ String) StringByAppendingPathExtensionForType(contentType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingPathExtensionForType:"), contentType)
	return rv
}
// Returns a new string by applying a specified transform to the string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/applyingTransform(_:reverse:)
func (s_ String) StringByApplyingTransformReverse(transform unsafe.Pointer, reverse bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByApplyingTransform:reverse:"), transform, reverse)
	return rv
}
// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/boundingRect(with:options:attributes:)
func (s_ String) BoundingRectWithSizeOptionsAttributes(size unsafe.Pointer, options unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("boundingRectWithSize:options:attributes:"), size, options, attributes)
	return rv
}
// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/boundingRect(with:options:attributes:context:)
func (s_ String) BoundingRectWithSizeOptionsAttributesContext(size unsafe.Pointer, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("boundingRectWithSize:options:attributes:context:"), size, options, attributes, context)
	return rv
}
// Returns a representation of the receiver as a C string in the default C-string encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/cString()
func (s_ String) CString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("cString"))
	return rv
}
// Returns a representation of the string as a C string using a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/cString(using:)
func (s_ String) CStringUsingEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("cStringUsingEncoding:"), encoding)
	return rv
}
// Returns the length in char-sized units of the receiver’s C-string representation in the default C-string encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/cStringLength()
func (s_ String) CStringLength() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("cStringLength"))
	return rv
}
// Returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/canBeConverted(to:)
func (s_ String) CanBeConvertedToEncoding(encoding unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canBeConvertedToEncoding:"), encoding)
	return rv
}
// Returns a capitalized representation of the receiver using the specified locale. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/capitalized(with:)
func (s_ String) CapitalizedStringWithLocale(locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("capitalizedStringWithLocale:"), locale)
	return rv
}
// Returns the result of invoking with as the only option. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/caseInsensitiveCompare(_:)
func (s_ String) CaseInsensitiveCompare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("caseInsensitiveCompare:"), string)
	return rv
}
// Returns the character at a given UTF-16 code unit index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/character(at:)
func (s_ String) CharacterAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("characterAtIndex:"), index)
	return rv
}
// Returns a string containing characters the receiver and a given string have in common, starting from the beginning of each up to the first characters that aren’t equivalent. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/commonPrefix(with:options:)
func (s_ String) CommonPrefixWithStringOptions(str string, mask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("commonPrefixWithString:options:"), str, mask)
	return rv
}
// Returns the result of invoking with no options and the receiver’s full extent as the range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:)
func (s_ String) Compare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compare:"), string)
	return rv
}
// Compares the string with the specified string using the given options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:)
func (s_ String) CompareOptions(string string, mask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compare:options:"), string, mask)
	return rv
}
// Returns the result of invoking with a locale. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:range:)
func (s_ String) CompareOptionsRange(string string, mask unsafe.Pointer, rangeOfReceiverToCompare unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compare:options:range:"), string, mask, rangeOfReceiverToCompare)
	return rv
}
// Compares the string using the specified options and returns the lexical ordering for the range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:range:locale:)
func (s_ String) CompareOptionsRangeLocale(string string, mask unsafe.Pointer, rangeOfReceiverToCompare unsafe.Pointer, locale objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compare:options:range:locale:"), string, mask, rangeOfReceiverToCompare, locale)
	return rv
}
// Interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/completePath(into:caseSensitive:matchesInto:filterTypes:)
func (s_ String) CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName string, flag bool, outputArray unsafe.Pointer, filterTypes unsafe.Pointer) uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("completePathIntoString:caseSensitive:matchesIntoArray:filterTypes:"), outputName, flag, outputArray, filterTypes)
	return rv
}
// Returns an array containing substrings from the receiver that have been divided by a given separator. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/components(separatedBy:)-238fy
func (s_ String) ComponentsSeparatedByString(separator string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("componentsSeparatedByString:"), separator)
	return rv
}
// Returns an array containing substrings from the receiver that have been divided by characters in a given set. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/components(separatedBy:)-27x9g
func (s_ String) ComponentsSeparatedByCharactersInSet(separator unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("componentsSeparatedByCharactersInSet:"), separator)
	return rv
}
// Returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/contains(_:)
func (s_ String) ContainsString(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("containsString:"), str)
	return rv
}
// Returns an object containing a representation of the receiver encoded using a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/data(using:)
func (s_ String) DataUsingEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dataUsingEncoding:"), encoding)
	return rv
}
// Returns an object containing a representation of the receiver encoded using a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/data(using:allowLossyConversion:)
func (s_ String) DataUsingEncodingAllowLossyConversion(encoding unsafe.Pointer, lossy bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dataUsingEncoding:allowLossyConversion:"), encoding, lossy)
	return rv
}
// Draws the receiver with the font and other display characteristics of the given attributes, at the specified point in the current graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(at:withAttributes:)
func (s_ String) DrawAtPointWithAttributes(point unsafe.Pointer, attrs unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawAtPoint:withAttributes:"), point, attrs)
}
// Draws the attributed string inside the specified bounding rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(in:withAttributes:)
func (s_ String) DrawInRectWithAttributes(rect unsafe.Pointer, attrs unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawInRect:withAttributes:"), rect, attrs)
}
// Draws the receiver with the specified options and other display characteristics of the given attributes, within the specified rectangle in the current graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(with:options:attributes:)
func (s_ String) DrawWithRectOptionsAttributes(rect unsafe.Pointer, options unsafe.Pointer, attributes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawWithRect:options:attributes:"), rect, options, attributes)
}
// Draws the attributed string in the specified bounding rectangle using the provided options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(with:options:attributes:context:)
func (s_ String) DrawWithRectOptionsAttributesContext(rect unsafe.Pointer, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawWithRect:options:attributes:context:"), rect, options, attributes, context)
}
// Draws the string in a single line at the specified point in the current graphics context using the specified font and attributes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawAtPoint:forWidth:withFont:fontSize:lineBreakMode:baselineAdjustment:
func (s_ String) DrawAtPointForWidthWithFontFontSizeLineBreakModeBaselineAdjustment(point unsafe.Pointer, width float64, font unsafe.Pointer, fontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("drawAtPoint:forWidth:withFont:fontSize:lineBreakMode:baselineAdjustment:"), point, width, font, fontSize, lineBreakMode, baselineAdjustment)
	return rv
}
// Draws the string in a single line at the specified point in the current graphics context using the specified font and attributes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawAtPoint:forWidth:withFont:lineBreakMode:
func (s_ String) DrawAtPointForWidthWithFontLineBreakMode(point unsafe.Pointer, width float64, font unsafe.Pointer, lineBreakMode unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("drawAtPoint:forWidth:withFont:lineBreakMode:"), point, width, font, lineBreakMode)
	return rv
}
// Draws the string in a single line with the specified font and attributes, adjusting the font attributes as needed to render as much of the text as possible. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawAtPoint:forWidth:withFont:minFontSize:actualFontSize:lineBreakMode:baselineAdjustment:
func (s_ String) DrawAtPointForWidthWithFontMinFontSizeActualFontSizeLineBreakModeBaselineAdjustment(point unsafe.Pointer, width float64, font unsafe.Pointer, minFontSize float64, actualFontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("drawAtPoint:forWidth:withFont:minFontSize:actualFontSize:lineBreakMode:baselineAdjustment:"), point, width, font, minFontSize, actualFontSize, lineBreakMode, baselineAdjustment)
	return rv
}
// Draws the string in a single line at the specified point in the current graphics context using the specified font. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawAtPoint:withFont:
func (s_ String) DrawAtPointWithFont(point unsafe.Pointer, font unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("drawAtPoint:withFont:"), point, font)
	return rv
}
// Draws the string in the current graphics context using the specified bounding rectangle and font. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawInRect:withFont:
func (s_ String) DrawInRectWithFont(rect unsafe.Pointer, font unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("drawInRect:withFont:"), rect, font)
	return rv
}
// Draws the string in the current graphics context using the specified bounding rectangle, font, and attributes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawInRect:withFont:lineBreakMode:
func (s_ String) DrawInRectWithFontLineBreakMode(rect unsafe.Pointer, font unsafe.Pointer, lineBreakMode unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("drawInRect:withFont:lineBreakMode:"), rect, font, lineBreakMode)
	return rv
}
// Draws the string in the current graphics context using the specified bounding rectangle, font and attributes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawInRect:withFont:lineBreakMode:alignment:
func (s_ String) DrawInRectWithFontLineBreakModeAlignment(rect unsafe.Pointer, font unsafe.Pointer, lineBreakMode unsafe.Pointer, alignment unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("drawInRect:withFont:lineBreakMode:alignment:"), rect, font, lineBreakMode, alignment)
	return rv
}
// Enumerates all the lines in the string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/enumerateLines(_:)
func (s_ String) EnumerateLinesUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateLinesUsingBlock:"), block)
}
// Performs linguistic analysis on the specified string by enumerating the specific range of the string, providing the Block with the located tags. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/enumerateLinguisticTags(in:scheme:options:orthography:using:)
func (s_ String) EnumerateLinguisticTagsInRangeSchemeOptionsOrthographyUsingBlock(range_ unsafe.Pointer, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateLinguisticTagsInRange:scheme:options:orthography:usingBlock:"), range_, scheme, options, orthography, block)
}
// Enumerates the substrings of the specified type in the specified range of the string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/enumerateSubstrings(in:options:using:)
func (s_ String) EnumerateSubstringsInRangeOptionsUsingBlock(range_ unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateSubstringsInRange:options:usingBlock:"), range_, opts, block)
}
// Creates a string suitable for comparison by removing the specified character distinctions from a string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/folding(options:locale:)
func (s_ String) StringByFoldingWithOptionsLocale(options unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByFoldingWithOptions:locale:"), options, locale)
	return rv
}
// Gets a given range of characters as bytes in a specified encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getBytes(_:maxLength:usedLength:encoding:options:range:remaining:)
func (s_ String) GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint, usedBufferCount unsafe.Pointer, encoding unsafe.Pointer, options unsafe.Pointer, range_ unsafe.Pointer, leftover unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("getBytes:maxLength:usedLength:encoding:options:range:remainingRange:"), buffer, maxBufferCount, usedBufferCount, encoding, options, range_, leftover)
	return rv
}
// Invokes with as the maximum length, the receiver’s entire extent as the range, and for the remaining range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:)
func (s_ String) GetCString(bytes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCString:"), bytes)
}
// Invokes with as the maximum length in char-sized units, the receiver’s entire extent as the range, and for the remaining range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:maxLength:)
func (s_ String) GetCStringMaxLength(bytes unsafe.Pointer, maxLength uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCString:maxLength:"), bytes, maxLength)
}
// Converts the string to a given encoding and stores it in a buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:maxLength:encoding:)
func (s_ String) GetCStringMaxLengthEncoding(buffer unsafe.Pointer, maxBufferCount uint, encoding unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("getCString:maxLength:encoding:"), buffer, maxBufferCount, encoding)
	return rv
}
// Converts the receiver’s content to the default C-string encoding and stores them in a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:maxLength:range:remaining:)
func (s_ String) GetCStringMaxLengthRangeRemainingRange(bytes unsafe.Pointer, maxLength uint, aRange unsafe.Pointer, leftoverRange unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCString:maxLength:range:remainingRange:"), bytes, maxLength, aRange, leftoverRange)
}
// Copies all characters from the receiver into a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCharacters(_:)
func (s_ String) GetCharacters(buffer unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCharacters:"), buffer)
}
// Copies characters from a given range in the receiver into a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCharacters(_:range:)
func (s_ String) GetCharactersRange(buffer unsafe.Pointer, range_ unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCharacters:range:"), buffer, range_)
}
// Interprets the receiver as a system-independent path and fills a buffer with a C-string in a format and encoding suitable for use with file-system calls. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getFileSystemRepresentation(_:maxLength:)
func (s_ String) GetFileSystemRepresentationMaxLength(cname unsafe.Pointer, max uint) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("getFileSystemRepresentation:maxLength:"), cname, max)
	return rv
}
// Returns by reference the beginning of the first line and the end of the last line touched by the given range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getLineStart(_:end:contentsEnd:for:)
func (s_ String) GetLineStartEndContentsEndForRange(startPtr unsafe.Pointer, lineEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getLineStart:end:contentsEnd:forRange:"), startPtr, lineEndPtr, contentsEndPtr, range_)
}
// Returns by reference the beginning of the first paragraph and the end of the last paragraph touched by the given range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getParagraphStart(_:end:contentsEnd:for:)
func (s_ String) GetParagraphStartEndContentsEndForRange(startPtr unsafe.Pointer, parEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getParagraphStart:end:contentsEnd:forRange:"), startPtr, parEndPtr, contentsEndPtr, range_)
}
// Returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/hasPrefix(_:)
func (s_ String) HasPrefix(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasPrefix:"), str)
	return rv
}
// Returns a Boolean value that indicates whether a given string matches the ending characters of the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/hasSuffix(_:)
func (s_ String) HasSuffix(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasSuffix:"), str)
	return rv
}
// Returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/isEqual(to:)
func (s_ String) IsEqualToString(aString string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEqualToString:"), aString)
	return rv
}
// Returns the number of bytes required to store the receiver in a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lengthOfBytes(using:)
func (s_ String) LengthOfBytesUsingEncoding(enc unsafe.Pointer) uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("lengthOfBytesUsingEncoding:"), enc)
	return rv
}
// Returns the range of characters representing the line or lines containing a given range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lineRange(for:)
func (s_ String) LineRangeForRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lineRangeForRange:"), range_)
	return rv
}
// Returns an array of linguistic tags for the specified range and requested tags within the receiving string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/linguisticTags(in:scheme:options:orthography:tokenRanges:)
func (s_ String) LinguisticTagsInRangeSchemeOptionsOrthographyTokenRanges(range_ unsafe.Pointer, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, tokenRanges unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("linguisticTagsInRange:scheme:options:orthography:tokenRanges:"), range_, scheme, options, orthography, tokenRanges)
	return rv
}
// Compares the string with a given string using a case-insensitive, localized, comparison. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCaseInsensitiveCompare(_:)
func (s_ String) LocalizedCaseInsensitiveCompare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedCaseInsensitiveCompare:"), string)
	return rv
}
// Returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCaseInsensitiveContains(_:)
func (s_ String) LocalizedCaseInsensitiveContainsString(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("localizedCaseInsensitiveContainsString:"), str)
	return rv
}
// Compares the string and a given string using a localized comparison. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCompare(_:)
func (s_ String) LocalizedCompare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedCompare:"), string)
	return rv
}
// Compares strings as sorted by the Finder. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardCompare(_:)
func (s_ String) LocalizedStandardCompare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedStandardCompare:"), string)
	return rv
}
// Returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardContains(_:)
func (s_ String) LocalizedStandardContainsString(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("localizedStandardContainsString:"), str)
	return rv
}
// Finds and returns the range of the first occurrence of a given string within the string by performing a case and diacritic insensitive, locale-aware search. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardRange(of:)
func (s_ String) LocalizedStandardRangeOfString(str string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedStandardRangeOfString:"), str)
	return rv
}
// Returns a representation of the receiver as a C string in the default C-string encoding, possibly losing information in converting to that encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lossyCString()
func (s_ String) LossyCString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lossyCString"))
	return rv
}
// Returns a version of the string with all letters converted to lowercase, taking into account the specified locale. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lowercased(with:)
func (s_ String) LowercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lowercaseStringWithLocale:"), locale)
	return rv
}
// Returns the maximum number of bytes needed to store the receiver in a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/maximumLengthOfBytes(using:)
func (s_ String) MaximumLengthOfBytesUsingEncoding(enc unsafe.Pointer) uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("maximumLengthOfBytesUsingEncoding:"), enc)
	return rv
}
// Returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/padding(toLength:withPad:startingAt:)
func (s_ String) StringByPaddingToLengthWithStringStartingAtIndex(newLength uint, padString string, padIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByPaddingToLength:withString:startingAtIndex:"), newLength, padString, padIndex)
	return rv
}
// Returns the range of characters representing the paragraph or paragraphs containing a given range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/paragraphRange(for:)
func (s_ String) ParagraphRangeForRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("paragraphRangeForRange:"), range_)
	return rv
}
// Parses the receiver as a text representation of a property list, returning an , , , or object, according to the topmost element. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/propertyList()
func (s_ String) PropertyList() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("propertyList"))
	return rv
}
// Returns a dictionary object initialized with the keys and values found in the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/propertyListFromStringsFileFormat()
func (s_ String) PropertyListFromStringsFileFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("propertyListFromStringsFileFormat"))
	return rv
}
// Finds and returns the range of the first occurrence of a given string within the string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:)
func (s_ String) RangeOfString(searchString string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeOfString:"), searchString)
	return rv
}
// Finds and returns the range of the first occurrence of a given string within the string, subject to given options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:)
func (s_ String) RangeOfStringOptions(searchString string, mask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeOfString:options:"), searchString, mask)
	return rv
}
// Finds and returns the range of the first occurrence of a given string, within the given range of the string, subject to given options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:range:)
func (s_ String) RangeOfStringOptionsRange(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeOfString:options:range:"), searchString, mask, rangeOfReceiverToSearch)
	return rv
}
// Finds and returns the range of the first occurrence of a given string within a given range of the string, subject to given options, using the specified locale, if any. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:range:locale:)
func (s_ String) RangeOfStringOptionsRangeLocale(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeOfString:options:range:locale:"), searchString, mask, rangeOfReceiverToSearch, locale)
	return rv
}
// Finds and returns the range in the string of the first character from a given character set. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:)
func (s_ String) RangeOfCharacterFromSet(searchSet unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeOfCharacterFromSet:"), searchSet)
	return rv
}
// Finds and returns the range in the string of the first character, using given options, from a given character set. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:options:)
func (s_ String) RangeOfCharacterFromSetOptions(searchSet unsafe.Pointer, mask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeOfCharacterFromSet:options:"), searchSet, mask)
	return rv
}
// Finds and returns the range in the string of the first character from a given character set found in a given range with given options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:options:range:)
func (s_ String) RangeOfCharacterFromSetOptionsRange(searchSet unsafe.Pointer, mask unsafe.Pointer, rangeOfReceiverToSearch unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeOfCharacterFromSet:options:range:"), searchSet, mask, rangeOfReceiverToSearch)
	return rv
}
// Returns the range in the receiver of the composed character sequence located at a given index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfComposedCharacterSequence(at:)
func (s_ String) RangeOfComposedCharacterSequenceAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeOfComposedCharacterSequenceAtIndex:"), index)
	return rv
}
// Returns the range in the string of the composed character sequences for a given range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfComposedCharacterSequences(for:)
func (s_ String) RangeOfComposedCharacterSequencesForRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeOfComposedCharacterSequencesForRange:"), range_)
	return rv
}
// Returns a new string in which the characters in a specified range of the receiver are replaced by a given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingCharacters(in:with:)
func (s_ String) StringByReplacingCharactersInRangeWithString(range_ unsafe.Pointer, replacement string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByReplacingCharactersInRange:withString:"), range_, replacement)
	return rv
}
// Returns a new string in which all occurrences of a target string in the receiver are replaced by another given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingOccurrences(of:with:)
func (s_ String) StringByReplacingOccurrencesOfStringWithString(target string, replacement string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByReplacingOccurrencesOfString:withString:"), target, replacement)
	return rv
}
// Returns a new string in which all occurrences of a target string in a specified range of the receiver are replaced by another given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingOccurrences(of:with:options:range:)
func (s_ String) StringByReplacingOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options unsafe.Pointer, searchRange unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByReplacingOccurrencesOfString:withString:options:range:"), target, replacement, options, searchRange)
	return rv
}
// Returns a new string made by replacing in the receiver all percent escapes with the matching characters as determined by a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingPercentEscapes(using:)
func (s_ String) StringByReplacingPercentEscapesUsingEncoding(enc unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByReplacingPercentEscapesUsingEncoding:"), enc)
	return rv
}
// Returns the bounding box size the receiver occupies when drawn with the given attributes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/size(withAttributes:)
func (s_ String) SizeWithAttributes(attrs unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sizeWithAttributes:"), attrs)
	return rv
}
// Returns the size of the string if it were to be rendered with the specified font on a single line. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:
func (s_ String) SizeWithFont(font unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sizeWithFont:"), font)
	return rv
}
// Returns the size of the string if it were rendered and constrained to the specified size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:constrainedToSize:
func (s_ String) SizeWithFontConstrainedToSize(font unsafe.Pointer, size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sizeWithFont:constrainedToSize:"), font, size)
	return rv
}
// Returns the size of the string if it were rendered with the specified constraints. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:constrainedToSize:lineBreakMode:
func (s_ String) SizeWithFontConstrainedToSizeLineBreakMode(font unsafe.Pointer, size unsafe.Pointer, lineBreakMode unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sizeWithFont:constrainedToSize:lineBreakMode:"), font, size, lineBreakMode)
	return rv
}
// Returns the size of the string if it were to be rendered with the specified font and line attributes on a single line. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:forWidth:lineBreakMode:
func (s_ String) SizeWithFontForWidthLineBreakMode(font unsafe.Pointer, width float64, lineBreakMode unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sizeWithFont:forWidth:lineBreakMode:"), font, width, lineBreakMode)
	return rv
}
// Returns the size of the string if it were rendered with the specified constraints, including a variable font size, on a single line. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:minFontSize:actualFontSize:forWidth:lineBreakMode:
func (s_ String) SizeWithFontMinFontSizeActualFontSizeForWidthLineBreakMode(font unsafe.Pointer, minFontSize float64, actualFontSize float64, width float64, lineBreakMode unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sizeWithFont:minFontSize:actualFontSize:forWidth:lineBreakMode:"), font, minFontSize, actualFontSize, width, lineBreakMode)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sr_sensorForDeletionRecordsFromSensor()
func (s_ String) Sr_sensorForDeletionRecordsFromSensor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sr_sensorForDeletionRecordsFromSensor"))
	return rv
}
// Returns a string made by appending to the receiver a string constructed from a given format string and the following arguments. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringByAppendingFormat:
func (s_ String) StringByAppendingFormat(format string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingFormat:"), format)
	return rv
}
// Returns an array of strings made by separately appending to the receiver each string in a given array. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/strings(byAppendingPaths:)
func (s_ String) StringsByAppendingPaths(paths unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringsByAppendingPaths:"), paths)
	return rv
}
// Returns a new string containing the characters of the receiver from the one at a given index to the end. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/substring(from:)
func (s_ String) SubstringFromIndex(from uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("substringFromIndex:"), from)
	return rv
}
// Returns a new string containing the characters of the receiver up to, but not including, the one at a given index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/substring(to:)
func (s_ String) SubstringToIndex(to uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("substringToIndex:"), to)
	return rv
}
// Returns a string object containing the characters of the receiver that lie within a given range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/substring(with:)
func (s_ String) SubstringWithRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("substringWithRange:"), range_)
	return rv
}
// Returns a new string made by removing from both ends of the receiver characters contained in a given character set. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/trimmingCharacters(in:)
func (s_ String) StringByTrimmingCharactersInSet(set unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByTrimmingCharactersInSet:"), set)
	return rv
}
// Returns a version of the string with all letters converted to uppercase, taking into account the specified locale. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/uppercased(with:)
func (s_ String) UppercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("uppercaseStringWithLocale:"), locale)
	return rv
}
// Returns a string variation suitable for the specified presentation width. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/variantFittingPresentationWidth(_:)
func (s_ String) VariantFittingPresentationWidth(width int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("variantFittingPresentationWidth:"), width)
	return rv
}
// Writes the contents of the receiver to the location specified by a given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(to:atomically:)
func (s_ String) WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToURL:atomically:"), url, atomically)
	return rv
}
// Writes the contents of the receiver to the URL specified by using the specified encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(to:atomically:encoding:)
func (s_ String) WriteToURLAtomicallyEncodingError(url unsafe.Pointer, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToURL:atomically:encoding:error:"), url, useAuxiliaryFile, enc, error)
	return rv
}
// Writes the contents of the receiver to the file specified by a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(toFile:atomically:)
func (s_ String) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToFile:atomically:"), path, useAuxiliaryFile)
	return rv
}
// Writes the contents of the receiver to a file at a given path using a given encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(toFile:atomically:encoding:)
func (s_ String) WriteToFileAtomicallyEncodingError(path string, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToFile:atomically:encoding:error:"), path, useAuxiliaryFile, enc, error)
	return rv
}

