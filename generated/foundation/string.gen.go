// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [String] class.
var (
	stringClass     _StringClass
	stringClassOnce sync.Once
)

func getStringClass() _StringClass {
	stringClassOnce.Do(func() {
		stringClass = _StringClass{objc.GetClass("NSString")}
	})
	return stringClass
}

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
	BoundingRectWithSizeOptionsAttributes(size Size, options unsafe.Pointer, attributes unsafe.Pointer) Rect
	BoundingRectWithSizeOptionsAttributesContext(size coregraphics.CGSize, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer) coregraphics.CGRect
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
	CompareOptionsRange(string string, mask unsafe.Pointer, rangeOfReceiverToCompare Range) unsafe.Pointer
	CompareOptionsRangeLocale(string string, mask unsafe.Pointer, rangeOfReceiverToCompare Range, locale objc.ID) unsafe.Pointer
	CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName string, flag bool, outputArray unsafe.Pointer, filterTypes unsafe.Pointer) uint
	ComponentsSeparatedByString(separator string) []string
	ComponentsSeparatedByCharactersInSet(separator unsafe.Pointer) []string
	ContainsString(str string) bool
	DataUsingEncoding(encoding unsafe.Pointer) unsafe.Pointer
	DataUsingEncodingAllowLossyConversion(encoding unsafe.Pointer, lossy bool) unsafe.Pointer
	DrawAtPointWithAttributes(point coregraphics.CGPoint, attrs unsafe.Pointer)
	DrawInRectWithAttributes(rect coregraphics.CGRect, attrs unsafe.Pointer)
	DrawWithRectOptionsAttributes(rect Rect, options unsafe.Pointer, attributes unsafe.Pointer)
	DrawWithRectOptionsAttributesContext(rect coregraphics.CGRect, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer)
	DrawAtPointForWidthWithFontFontSizeLineBreakModeBaselineAdjustment(point coregraphics.CGPoint, width float64, font unsafe.Pointer, fontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) coregraphics.CGSize
	DrawAtPointForWidthWithFontLineBreakMode(point coregraphics.CGPoint, width float64, font unsafe.Pointer, lineBreakMode unsafe.Pointer) coregraphics.CGSize
	DrawAtPointForWidthWithFontMinFontSizeActualFontSizeLineBreakModeBaselineAdjustment(point coregraphics.CGPoint, width float64, font unsafe.Pointer, minFontSize float64, actualFontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) coregraphics.CGSize
	DrawAtPointWithFont(point coregraphics.CGPoint, font unsafe.Pointer) coregraphics.CGSize
	DrawInRectWithFont(rect coregraphics.CGRect, font unsafe.Pointer) coregraphics.CGSize
	DrawInRectWithFontLineBreakMode(rect coregraphics.CGRect, font unsafe.Pointer, lineBreakMode unsafe.Pointer) coregraphics.CGSize
	DrawInRectWithFontLineBreakModeAlignment(rect coregraphics.CGRect, font unsafe.Pointer, lineBreakMode unsafe.Pointer, alignment unsafe.Pointer) coregraphics.CGSize
	EnumerateLinesUsingBlock(block unsafe.Pointer)
	EnumerateLinguisticTagsInRangeSchemeOptionsOrthographyUsingBlock(range_ Range, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, block unsafe.Pointer)
	EnumerateSubstringsInRangeOptionsUsingBlock(range_ Range, opts unsafe.Pointer, block unsafe.Pointer)
	StringByFoldingWithOptionsLocale(options unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer
	GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint, usedBufferCount unsafe.Pointer, encoding unsafe.Pointer, options unsafe.Pointer, range_ Range, leftover unsafe.Pointer) bool
	GetCString(bytes unsafe.Pointer)
	GetCStringMaxLength(bytes unsafe.Pointer, maxLength uint)
	GetCStringMaxLengthEncoding(buffer unsafe.Pointer, maxBufferCount uint, encoding unsafe.Pointer) bool
	GetCStringMaxLengthRangeRemainingRange(bytes unsafe.Pointer, maxLength uint, aRange Range, leftoverRange unsafe.Pointer)
	GetCharacters(buffer unsafe.Pointer)
	GetCharactersRange(buffer unsafe.Pointer, range_ Range)
	GetFileSystemRepresentationMaxLength(cname unsafe.Pointer, max uint) bool
	GetLineStartEndContentsEndForRange(startPtr unsafe.Pointer, lineEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ Range)
	GetParagraphStartEndContentsEndForRange(startPtr unsafe.Pointer, parEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ Range)
	HasPrefix(str string) bool
	HasSuffix(str string) bool
	IsEqualToString(aString string) bool
	LengthOfBytesUsingEncoding(enc unsafe.Pointer) uint
	LineRangeForRange(range_ Range) Range
	LinguisticTagsInRangeSchemeOptionsOrthographyTokenRanges(range_ Range, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, tokenRanges unsafe.Pointer) []string
	LocalizedCaseInsensitiveCompare(string string) unsafe.Pointer
	LocalizedCaseInsensitiveContainsString(str string) bool
	LocalizedCompare(string string) unsafe.Pointer
	LocalizedStandardCompare(string string) unsafe.Pointer
	LocalizedStandardContainsString(str string) bool
	LocalizedStandardRangeOfString(str string) Range
	LossyCString() unsafe.Pointer
	LowercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer
	MaximumLengthOfBytesUsingEncoding(enc unsafe.Pointer) uint
	StringByPaddingToLengthWithStringStartingAtIndex(newLength uint, padString string, padIndex uint) unsafe.Pointer
	ParagraphRangeForRange(range_ Range) Range
	PropertyList() objc.ID
	PropertyListFromStringsFileFormat() unsafe.Pointer
	RangeOfString(searchString string) Range
	RangeOfStringOptions(searchString string, mask unsafe.Pointer) Range
	RangeOfStringOptionsRange(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch Range) Range
	RangeOfStringOptionsRangeLocale(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch Range, locale unsafe.Pointer) Range
	RangeOfCharacterFromSet(searchSet unsafe.Pointer) Range
	RangeOfCharacterFromSetOptions(searchSet unsafe.Pointer, mask unsafe.Pointer) Range
	RangeOfCharacterFromSetOptionsRange(searchSet unsafe.Pointer, mask unsafe.Pointer, rangeOfReceiverToSearch Range) Range
	RangeOfComposedCharacterSequenceAtIndex(index uint) Range
	RangeOfComposedCharacterSequencesForRange(range_ Range) Range
	StringByReplacingCharactersInRangeWithString(range_ Range, replacement string) unsafe.Pointer
	StringByReplacingOccurrencesOfStringWithString(target string, replacement string) unsafe.Pointer
	StringByReplacingOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options unsafe.Pointer, searchRange Range) unsafe.Pointer
	StringByReplacingPercentEscapesUsingEncoding(enc unsafe.Pointer) unsafe.Pointer
	SizeWithAttributes(attrs unsafe.Pointer) coregraphics.CGSize
	SizeWithFont(font unsafe.Pointer) coregraphics.CGSize
	SizeWithFontConstrainedToSize(font unsafe.Pointer, size coregraphics.CGSize) coregraphics.CGSize
	SizeWithFontConstrainedToSizeLineBreakMode(font unsafe.Pointer, size coregraphics.CGSize, lineBreakMode unsafe.Pointer) coregraphics.CGSize
	SizeWithFontForWidthLineBreakMode(font unsafe.Pointer, width float64, lineBreakMode unsafe.Pointer) coregraphics.CGSize
	SizeWithFontMinFontSizeActualFontSizeForWidthLineBreakMode(font unsafe.Pointer, minFontSize float64, actualFontSize float64, width float64, lineBreakMode unsafe.Pointer) coregraphics.CGSize
	Sr_sensorForDeletionRecordsFromSensor() unsafe.Pointer
	StringByAppendingFormat(format string) unsafe.Pointer
	StringsByAppendingPaths(paths unsafe.Pointer) []string
	SubstringFromIndex(from uint) unsafe.Pointer
	SubstringToIndex(to uint) unsafe.Pointer
	SubstringWithRange(range_ Range) unsafe.Pointer
	StringByTrimmingCharactersInSet(set unsafe.Pointer) unsafe.Pointer
	UppercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer
	VariantFittingPresentationWidth(width int) unsafe.Pointer
	WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool
	WriteToURLAtomicallyEncodingError(url unsafe.Pointer, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool
	WriteToFileAtomically(path string, useAuxiliaryFile bool) bool
	WriteToFileAtomicallyEncodingError(path string, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool
}

// A static, plain-text Unicode string object.
//
// You can use this type in Swift when you need reference semantics or other Foundation-specific behavior. The class and its mutable subclass, , provide an extensive set of APIs for working with strings, including methods for comparing, searching, and modifying strings. objects are used throughout Foundation and other Cocoa frameworks, serving as the basis for all textual and linguistic functionality on the platform. is with its Core Foundation counterpart, . See for more information.
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


// Initializes the receiver, a newly allocated object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:length:)
func NewStringWithCStringLength(bytes unsafe.Pointer, length uint) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCString:length:"), bytes, length)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted without any localization.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(format:arguments:)
func NewStringWithFormatArguments(format string, argList unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:arguments:"), objc.String(format), argList)
	rv.Autorelease()
	return rv
}

// Initializes the receiver, a newly allocated object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:)
func NewStringWithCString(bytes unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCString:"), bytes)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by reading data from the file at a given path and returns by reference the encoding used to interpret the characters.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:usedEncoding:)
func NewStringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfFile:usedEncoding:error:"), objc.String(path), enc, error)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by reading data from a given URL and returns by reference the encoding used to interpret the data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:usedEncoding:)-2c72d
func NewStringWithContentsOfURLUsedEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfURL:usedEncoding:error:"), url, enc, error)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by copying the characters from a given C array of UTF8-encoded bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(UTF8String:)-vg2b
func NewStringWithUTF8String(nullTerminatedCString unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithUTF8String:"), nullTerminatedCString)
	rv.Autorelease()
	return rv
}

// Returns an initialized object that contains a given number of characters from a given C array of UTF-16 code units.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:freeWhenDone:)
func NewStringWithCharactersNoCopyLengthFreeWhenDone(characters unsafe.Pointer, length uint, freeBuffer bool) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCharactersNoCopy:length:freeWhenDone:"), characters, length, freeBuffer)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by converting given data into UTF-16 code units using a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(data:encoding:)
func NewStringWithDataEncoding(data unsafe.Pointer, encoding unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithData:encoding:"), data, encoding)
	rv.Autorelease()
	return rv
}

// Initializes the receiver, a newly allocated object, by reading data from the file named by .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:)
func NewStringWithContentsOfFile(path string) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	rv.Autorelease()
	return rv
}

// Returns an object initialized using the characters in a given C array, interpreted according to a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:encoding:)-20f9h
func NewStringWithCStringEncoding(nullTerminatedCString unsafe.Pointer, encoding unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCString:encoding:"), nullTerminatedCString, encoding)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:error:
func NewStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string, error unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), error)
	rv.Autorelease()
	return rv
}

// Returns an initialized object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:freeWhenDone:)
func NewStringWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len uint, encoding unsafe.Pointer, freeBuffer bool) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:freeWhenDone:"), bytes, len, encoding, freeBuffer)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:
func NewStringWithFormat(format string) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:"), objc.String(format))
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format string, validFormatSpecifiers string, locale objc.ID, argList unsafe.Pointer, error unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, argList, error)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersArgumentsError(format string, validFormatSpecifiers string, argList unsafe.Pointer, error unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), argList, error)
	rv.Autorelease()
	return rv
}

// Returns an initialized object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(bytes:length:encoding:)
func NewStringWithBytesLengthEncoding(bytes unsafe.Pointer, len uint, encoding unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithBytes:length:encoding:"), bytes, len, encoding)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:deallocator:)
func NewStringWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len uint, encoding unsafe.Pointer, deallocator unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:deallocator:"), bytes, len, encoding, deallocator)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:deallocator:)
func NewStringWithCharactersNoCopyLengthDeallocator(chars unsafe.Pointer, len uint, deallocator unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCharactersNoCopy:length:deallocator:"), chars, len, deallocator)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(coder:)
func NewStringWithCoder(coder unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by copying the characters from another given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(string:)-210xa
func NewStringWithString(aString string) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithString:"), objc.String(aString))
	rv.Autorelease()
	return rv
}

// Initializes the receiver, a newly allocated object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CStringNoCopy:length:freeWhenDone:)
func NewStringWithCStringNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, freeBuffer bool) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCStringNoCopy:length:freeWhenDone:"), bytes, length, freeBuffer)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by reading data from the file at a given path using a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:encoding:)
func NewStringWithContentsOfFileEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfFile:encoding:error:"), objc.String(path), enc, error)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:locale:
func NewStringWithFormatLocale(format string, locale objc.ID) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:locale:"), objc.String(format), locale)
	rv.Autorelease()
	return rv
}

// Returns an initialized object that contains a given number of characters from a given C array of UTF-16 code units.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(characters:length:)
func NewStringWithCharactersLength(characters unsafe.Pointer, length uint) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithCharacters:length:"), characters, length)
	rv.Autorelease()
	return rv
}

// Initializes the receiver, a newly allocated object, by reading data from the location named by a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:)
func NewStringWithContentsOfURL(url unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by reading data from a given URL interpreted using a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:encoding:)-715fw
func NewStringWithContentsOfURLEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithContentsOfURL:encoding:error:"), url, enc, error)
	rv.Autorelease()
	return rv
}

// Returns an object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale information. This method is meant to be called from within a variadic function, where the argument list will be available.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(format:locale:arguments:)
func NewStringWithFormatLocaleArguments(format string, locale objc.ID, argList unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithFormat:locale:arguments:"), objc.String(format), locale, argList)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleError(format string, validFormatSpecifiers string, locale objc.ID, error unsafe.Pointer) String {
	instance := getStringClass().Alloc()
	rv := objc.Send[String](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, error)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormat(format string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:"), objc.String(format))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormatFromTable(format string, table string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:fromTable:"), objc.String(format), objc.String(table))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:arguments:
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormatFromTableArguments(format string, table string, arguments unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:fromTable:arguments:"), objc.String(format), objc.String(table), arguments)
	return rv
}

// Returns a string containing the bytes in a given C array, interpreted according to a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(CString:encoding:)-7auq8
func (sc _StringClass) StringWithCStringEncoding(cString unsafe.Pointer, enc unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithCString:encoding:"), cString, enc)
	return rv
}

// Returns a string created by copying the data from a given C array of UTF8-encoded bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(UTF8String:)-8bcy8
func (sc _StringClass) StringWithUTF8String(nullTerminatedCString unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithUTF8String:"), nullTerminatedCString)
	return rv
}

// Returns a string created by reading data from a given URL interpreted using a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:encoding:)-x6cv
func (sc _StringClass) StringWithContentsOfURLEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfURL:encoding:error:"), url, enc, error)
	return rv
}

// Returns a string created by reading data from a given URL and returns by reference the encoding used to interpret the data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:usedEncoding:)-9jrum
func (sc _StringClass) StringWithContentsOfURLUsedEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfURL:usedEncoding:error:"), url, enc, error)
	return rv
}

// Returns a human-readable string giving the name of a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedName(of:)
func (sc _StringClass) LocalizedNameOfStringEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedNameOfStringEncoding:"), encoding)
	return rv
}

// Returns a string created by using a given format string as a template into which the remaining argument values are substituted according to the current locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStringWithFormat:
func (sc _StringClass) LocalizedStringWithFormat(format string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedStringWithFormat:"), objc.String(format))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStringWithValidatedFormat:validFormatSpecifiers:error:
func (sc _StringClass) LocalizedStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedStringWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), error)
	return rv
}

// Returns a localized string intended for display in a notification alert.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedUserNotificationString(forKey:arguments:)
func (sc _StringClass) LocalizedUserNotificationStringForKeyArguments(key string, arguments objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("localizedUserNotificationStringForKey:arguments:"), objc.String(key), arguments)
	return rv
}

// Returns a string built from the strings in a given array by concatenating them with a path separator between each pair.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/path(withComponents:)
func (sc _StringClass) PathWithComponents(components unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("pathWithComponents:"), components)
	return rv
}

// Returns an empty string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string
func (sc _StringClass) String() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("string"))
	return rv
}

// Creates a new string using a given C-string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withCString:)
func (sc _StringClass) StringWithCString(bytes unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithCString:"), bytes)
	return rv
}

// Returns a string containing the characters in a given C-string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withCString:length:)
func (sc _StringClass) StringWithCStringLength(bytes unsafe.Pointer, length uint) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithCString:length:"), bytes, length)
	return rv
}

// Returns a string created by reading data from the file named by a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withContentsOf:)
func (sc _StringClass) StringWithContentsOfURL(url unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithContentsOfURL:"), url)
	return rv
}

// Returns a string created by reading data from the file named by a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/string(withContentsOfFile:)
func (sc _StringClass) StringWithContentsOfFile(path string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("stringWithContentsOfFile:"), objc.String(path))
	return rv
}

// Returns the string encoding for the given data as detected by attempting to create a string according to the specified encoding options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringEncoding(for:encodingOptions:convertedString:usedLossyConversion:)
func (sc _StringClass) StringEncodingForDataEncodingOptionsConvertedStringUsedLossyConversion(data unsafe.Pointer, opts unsafe.Pointer, string string, usedLossyConversion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringEncodingForData:encodingOptions:convertedString:usedLossyConversion:"), data, opts, objc.String(string), usedLossyConversion)
	return rv
}

// Returns a string containing a given number of characters taken from a given C array of UTF-16 code units.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithCharacters:length:
func (sc _StringClass) StringWithCharactersLength(characters unsafe.Pointer, length uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithCharacters:length:"), characters, length)
	return rv
}

// Returns a string created by reading data from the file at a given path interpreted using a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfFile:encoding:error:
func (sc _StringClass) StringWithContentsOfFileEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfFile:encoding:error:"), objc.String(path), enc, error)
	return rv
}

// Returns a string created by reading data from the file at a given path and returns by reference the encoding used to interpret the file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfFile:usedEncoding:error:
func (sc _StringClass) StringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithContentsOfFile:usedEncoding:error:"), objc.String(path), enc, error)
	return rv
}

// Returns a string created by using a given format string as a template into which the remaining argument values are substituted.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithFormat:
func (sc _StringClass) StringWithFormat(format string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithFormat:"), objc.String(format))
	return rv
}

// Returns a string created by copying the characters from another given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithString:
func (sc _StringClass) StringWithString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithString:"), objc.String(string))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringWithValidatedFormat:validFormatSpecifiers:error:
func (sc _StringClass) StringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stringWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), error)
	return rv
}

// Returns a new string made from the receiver by replacing all characters not in the specified set with percent-encoded characters.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/addingPercentEncoding(withAllowedCharacters:)
func (s_ String) StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAddingPercentEncodingWithAllowedCharacters:"), allowedCharacters)
	return rv
}

// Returns a representation of the receiver using a given encoding to determine the percent escapes necessary to convert the receiver into a legal URL string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/addingPercentEscapes(using:)
func (s_ String) StringByAddingPercentEscapesUsingEncoding(enc unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAddingPercentEscapesUsingEncoding:"), enc)
	return rv
}

// Returns a new string made by appending a given string to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appending(_:)
func (s_ String) StringByAppendingString(aString string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingString:"), objc.String(aString))
	return rv
}

// Returns a new string made by appending to the receiver a given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathComponent(_:)
func (s_ String) StringByAppendingPathComponent(str string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingPathComponent:"), objc.String(str))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathComponent(_:conformingTo:)
func (s_ String) StringByAppendingPathComponentConformingToType(partialName string, contentType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingPathComponent:conformingToType:"), objc.String(partialName), contentType)
	return rv
}

// Returns a new string made by appending to the receiver an extension separator followed by a given extension.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathExtension(_:)
func (s_ String) StringByAppendingPathExtension(str string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingPathExtension:"), objc.String(str))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/appendingPathExtension(for:)
func (s_ String) StringByAppendingPathExtensionForType(contentType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingPathExtensionForType:"), contentType)
	return rv
}

// Returns a new string by applying a specified transform to the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/applyingTransform(_:reverse:)
func (s_ String) StringByApplyingTransformReverse(transform unsafe.Pointer, reverse bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByApplyingTransform:reverse:"), transform, reverse)
	return rv
}

// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/boundingRect(with:options:attributes:)
func (s_ String) BoundingRectWithSizeOptionsAttributes(size Size, options unsafe.Pointer, attributes unsafe.Pointer) Rect {
	rv := objc.Send[Rect](s_.ID, objc.Sel("boundingRectWithSize:options:attributes:"), size, options, attributes)
	return rv
}

// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/boundingRect(with:options:attributes:context:)
func (s_ String) BoundingRectWithSizeOptionsAttributesContext(size coregraphics.CGSize, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("boundingRectWithSize:options:attributes:context:"), size, options, attributes, context)
	return rv
}

// Returns a representation of the receiver as a C string in the default C-string encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/cString()
func (s_ String) CString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("cString"))
	return rv
}

// Returns a representation of the string as a C string using a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/cString(using:)
func (s_ String) CStringUsingEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("cStringUsingEncoding:"), encoding)
	return rv
}

// Returns the length in char-sized units of the receiver’s C-string representation in the default C-string encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/cStringLength()
func (s_ String) CStringLength() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("cStringLength"))
	return rv
}

// Returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/canBeConverted(to:)
func (s_ String) CanBeConvertedToEncoding(encoding unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canBeConvertedToEncoding:"), encoding)
	return rv
}

// Returns a capitalized representation of the receiver using the specified locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/capitalized(with:)
func (s_ String) CapitalizedStringWithLocale(locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("capitalizedStringWithLocale:"), locale)
	return rv
}

// Returns the result of invoking with as the only option.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/caseInsensitiveCompare(_:)
func (s_ String) CaseInsensitiveCompare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("caseInsensitiveCompare:"), objc.String(string))
	return rv
}

// Returns the character at a given UTF-16 code unit index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/character(at:)
func (s_ String) CharacterAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("characterAtIndex:"), index)
	return rv
}

// Returns a string containing characters the receiver and a given string have in common, starting from the beginning of each up to the first characters that aren’t equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/commonPrefix(with:options:)
func (s_ String) CommonPrefixWithStringOptions(str string, mask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("commonPrefixWithString:options:"), objc.String(str), mask)
	return rv
}

// Returns the result of invoking with no options and the receiver’s full extent as the range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:)
func (s_ String) Compare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compare:"), objc.String(string))
	return rv
}

// Compares the string with the specified string using the given options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:)
func (s_ String) CompareOptions(string string, mask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compare:options:"), objc.String(string), mask)
	return rv
}

// Returns the result of invoking with a locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:range:)
func (s_ String) CompareOptionsRange(string string, mask unsafe.Pointer, rangeOfReceiverToCompare Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compare:options:range:"), objc.String(string), mask, rangeOfReceiverToCompare)
	return rv
}

// Compares the string using the specified options and returns the lexical ordering for the range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:range:locale:)
func (s_ String) CompareOptionsRangeLocale(string string, mask unsafe.Pointer, rangeOfReceiverToCompare Range, locale objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compare:options:range:locale:"), objc.String(string), mask, rangeOfReceiverToCompare, locale)
	return rv
}

// Interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/completePath(into:caseSensitive:matchesInto:filterTypes:)
func (s_ String) CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName string, flag bool, outputArray unsafe.Pointer, filterTypes unsafe.Pointer) uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("completePathIntoString:caseSensitive:matchesIntoArray:filterTypes:"), objc.String(outputName), flag, outputArray, filterTypes)
	return rv
}

// Returns an array containing substrings from the receiver that have been divided by a given separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/components(separatedBy:)-238fy
func (s_ String) ComponentsSeparatedByString(separator string) []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("componentsSeparatedByString:"), objc.String(separator))
	return rv
}

// Returns an array containing substrings from the receiver that have been divided by characters in a given set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/components(separatedBy:)-27x9g
func (s_ String) ComponentsSeparatedByCharactersInSet(separator unsafe.Pointer) []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("componentsSeparatedByCharactersInSet:"), separator)
	return rv
}

// Returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/contains(_:)
func (s_ String) ContainsString(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("containsString:"), objc.String(str))
	return rv
}

// Returns an object containing a representation of the receiver encoded using a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/data(using:)
func (s_ String) DataUsingEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dataUsingEncoding:"), encoding)
	return rv
}

// Returns an object containing a representation of the receiver encoded using a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/data(using:allowLossyConversion:)
func (s_ String) DataUsingEncodingAllowLossyConversion(encoding unsafe.Pointer, lossy bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dataUsingEncoding:allowLossyConversion:"), encoding, lossy)
	return rv
}

// Draws the receiver with the font and other display characteristics of the given attributes, at the specified point in the current graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(at:withAttributes:)
func (s_ String) DrawAtPointWithAttributes(point coregraphics.CGPoint, attrs unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawAtPoint:withAttributes:"), point, attrs)
}

// Draws the attributed string inside the specified bounding rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(in:withAttributes:)
func (s_ String) DrawInRectWithAttributes(rect coregraphics.CGRect, attrs unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawInRect:withAttributes:"), rect, attrs)
}

// Draws the receiver with the specified options and other display characteristics of the given attributes, within the specified rectangle in the current graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(with:options:attributes:)
func (s_ String) DrawWithRectOptionsAttributes(rect Rect, options unsafe.Pointer, attributes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawWithRect:options:attributes:"), rect, options, attributes)
}

// Draws the attributed string in the specified bounding rectangle using the provided options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/draw(with:options:attributes:context:)
func (s_ String) DrawWithRectOptionsAttributesContext(rect coregraphics.CGRect, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawWithRect:options:attributes:context:"), rect, options, attributes, context)
}

// Draws the string in a single line at the specified point in the current graphics context using the specified font and attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawAtPoint:forWidth:withFont:fontSize:lineBreakMode:baselineAdjustment:
func (s_ String) DrawAtPointForWidthWithFontFontSizeLineBreakModeBaselineAdjustment(point coregraphics.CGPoint, width float64, font unsafe.Pointer, fontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("drawAtPoint:forWidth:withFont:fontSize:lineBreakMode:baselineAdjustment:"), point, width, font, fontSize, lineBreakMode, baselineAdjustment)
	return rv
}

// Draws the string in a single line at the specified point in the current graphics context using the specified font and attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawAtPoint:forWidth:withFont:lineBreakMode:
func (s_ String) DrawAtPointForWidthWithFontLineBreakMode(point coregraphics.CGPoint, width float64, font unsafe.Pointer, lineBreakMode unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("drawAtPoint:forWidth:withFont:lineBreakMode:"), point, width, font, lineBreakMode)
	return rv
}

// Draws the string in a single line with the specified font and attributes, adjusting the font attributes as needed to render as much of the text as possible.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawAtPoint:forWidth:withFont:minFontSize:actualFontSize:lineBreakMode:baselineAdjustment:
func (s_ String) DrawAtPointForWidthWithFontMinFontSizeActualFontSizeLineBreakModeBaselineAdjustment(point coregraphics.CGPoint, width float64, font unsafe.Pointer, minFontSize float64, actualFontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("drawAtPoint:forWidth:withFont:minFontSize:actualFontSize:lineBreakMode:baselineAdjustment:"), point, width, font, minFontSize, actualFontSize, lineBreakMode, baselineAdjustment)
	return rv
}

// Draws the string in a single line at the specified point in the current graphics context using the specified font.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawAtPoint:withFont:
func (s_ String) DrawAtPointWithFont(point coregraphics.CGPoint, font unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("drawAtPoint:withFont:"), point, font)
	return rv
}

// Draws the string in the current graphics context using the specified bounding rectangle and font.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawInRect:withFont:
func (s_ String) DrawInRectWithFont(rect coregraphics.CGRect, font unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("drawInRect:withFont:"), rect, font)
	return rv
}

// Draws the string in the current graphics context using the specified bounding rectangle, font, and attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawInRect:withFont:lineBreakMode:
func (s_ String) DrawInRectWithFontLineBreakMode(rect coregraphics.CGRect, font unsafe.Pointer, lineBreakMode unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("drawInRect:withFont:lineBreakMode:"), rect, font, lineBreakMode)
	return rv
}

// Draws the string in the current graphics context using the specified bounding rectangle, font and attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/drawInRect:withFont:lineBreakMode:alignment:
func (s_ String) DrawInRectWithFontLineBreakModeAlignment(rect coregraphics.CGRect, font unsafe.Pointer, lineBreakMode unsafe.Pointer, alignment unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("drawInRect:withFont:lineBreakMode:alignment:"), rect, font, lineBreakMode, alignment)
	return rv
}

// Enumerates all the lines in the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/enumerateLines(_:)
func (s_ String) EnumerateLinesUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateLinesUsingBlock:"), block)
}

// Performs linguistic analysis on the specified string by enumerating the specific range of the string, providing the Block with the located tags.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/enumerateLinguisticTags(in:scheme:options:orthography:using:)
func (s_ String) EnumerateLinguisticTagsInRangeSchemeOptionsOrthographyUsingBlock(range_ Range, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateLinguisticTagsInRange:scheme:options:orthography:usingBlock:"), range_, scheme, options, orthography, block)
}

// Enumerates the substrings of the specified type in the specified range of the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/enumerateSubstrings(in:options:using:)
func (s_ String) EnumerateSubstringsInRangeOptionsUsingBlock(range_ Range, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateSubstringsInRange:options:usingBlock:"), range_, opts, block)
}

// Creates a string suitable for comparison by removing the specified character distinctions from a string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/folding(options:locale:)
func (s_ String) StringByFoldingWithOptionsLocale(options unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByFoldingWithOptions:locale:"), options, locale)
	return rv
}

// Gets a given range of characters as bytes in a specified encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getBytes(_:maxLength:usedLength:encoding:options:range:remaining:)
func (s_ String) GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint, usedBufferCount unsafe.Pointer, encoding unsafe.Pointer, options unsafe.Pointer, range_ Range, leftover unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("getBytes:maxLength:usedLength:encoding:options:range:remainingRange:"), buffer, maxBufferCount, usedBufferCount, encoding, options, range_, leftover)
	return rv
}

// Invokes with as the maximum length, the receiver’s entire extent as the range, and for the remaining range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:)
func (s_ String) GetCString(bytes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCString:"), bytes)
}

// Invokes with as the maximum length in char-sized units, the receiver’s entire extent as the range, and for the remaining range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:maxLength:)
func (s_ String) GetCStringMaxLength(bytes unsafe.Pointer, maxLength uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCString:maxLength:"), bytes, maxLength)
}

// Converts the string to a given encoding and stores it in a buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:maxLength:encoding:)
func (s_ String) GetCStringMaxLengthEncoding(buffer unsafe.Pointer, maxBufferCount uint, encoding unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("getCString:maxLength:encoding:"), buffer, maxBufferCount, encoding)
	return rv
}

// Converts the receiver’s content to the default C-string encoding and stores them in a given buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:maxLength:range:remaining:)
func (s_ String) GetCStringMaxLengthRangeRemainingRange(bytes unsafe.Pointer, maxLength uint, aRange Range, leftoverRange unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCString:maxLength:range:remainingRange:"), bytes, maxLength, aRange, leftoverRange)
}

// Copies all characters from the receiver into a given buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCharacters(_:)
func (s_ String) GetCharacters(buffer unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCharacters:"), buffer)
}

// Copies characters from a given range in the receiver into a given buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getCharacters(_:range:)
func (s_ String) GetCharactersRange(buffer unsafe.Pointer, range_ Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getCharacters:range:"), buffer, range_)
}

// Interprets the receiver as a system-independent path and fills a buffer with a C-string in a format and encoding suitable for use with file-system calls.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getFileSystemRepresentation(_:maxLength:)
func (s_ String) GetFileSystemRepresentationMaxLength(cname unsafe.Pointer, max uint) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("getFileSystemRepresentation:maxLength:"), cname, max)
	return rv
}

// Returns by reference the beginning of the first line and the end of the last line touched by the given range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getLineStart(_:end:contentsEnd:for:)
func (s_ String) GetLineStartEndContentsEndForRange(startPtr unsafe.Pointer, lineEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getLineStart:end:contentsEnd:forRange:"), startPtr, lineEndPtr, contentsEndPtr, range_)
}

// Returns by reference the beginning of the first paragraph and the end of the last paragraph touched by the given range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/getParagraphStart(_:end:contentsEnd:for:)
func (s_ String) GetParagraphStartEndContentsEndForRange(startPtr unsafe.Pointer, parEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getParagraphStart:end:contentsEnd:forRange:"), startPtr, parEndPtr, contentsEndPtr, range_)
}

// Returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/hasPrefix(_:)
func (s_ String) HasPrefix(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasPrefix:"), objc.String(str))
	return rv
}

// Returns a Boolean value that indicates whether a given string matches the ending characters of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/hasSuffix(_:)
func (s_ String) HasSuffix(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasSuffix:"), objc.String(str))
	return rv
}

// Returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/isEqual(to:)
func (s_ String) IsEqualToString(aString string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEqualToString:"), objc.String(aString))
	return rv
}

// Returns the number of bytes required to store the receiver in a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lengthOfBytes(using:)
func (s_ String) LengthOfBytesUsingEncoding(enc unsafe.Pointer) uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("lengthOfBytesUsingEncoding:"), enc)
	return rv
}

// Returns the range of characters representing the line or lines containing a given range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lineRange(for:)
func (s_ String) LineRangeForRange(range_ Range) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("lineRangeForRange:"), range_)
	return rv
}

// Returns an array of linguistic tags for the specified range and requested tags within the receiving string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/linguisticTags(in:scheme:options:orthography:tokenRanges:)
func (s_ String) LinguisticTagsInRangeSchemeOptionsOrthographyTokenRanges(range_ Range, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, tokenRanges unsafe.Pointer) []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("linguisticTagsInRange:scheme:options:orthography:tokenRanges:"), range_, scheme, options, orthography, tokenRanges)
	return rv
}

// Compares the string with a given string using a case-insensitive, localized, comparison.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCaseInsensitiveCompare(_:)
func (s_ String) LocalizedCaseInsensitiveCompare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedCaseInsensitiveCompare:"), objc.String(string))
	return rv
}

// Returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCaseInsensitiveContains(_:)
func (s_ String) LocalizedCaseInsensitiveContainsString(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("localizedCaseInsensitiveContainsString:"), objc.String(str))
	return rv
}

// Compares the string and a given string using a localized comparison.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCompare(_:)
func (s_ String) LocalizedCompare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedCompare:"), objc.String(string))
	return rv
}

// Compares strings as sorted by the Finder.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardCompare(_:)
func (s_ String) LocalizedStandardCompare(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedStandardCompare:"), objc.String(string))
	return rv
}

// Returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardContains(_:)
func (s_ String) LocalizedStandardContainsString(str string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("localizedStandardContainsString:"), objc.String(str))
	return rv
}

// Finds and returns the range of the first occurrence of a given string within the string by performing a case and diacritic insensitive, locale-aware search.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardRange(of:)
func (s_ String) LocalizedStandardRangeOfString(str string) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("localizedStandardRangeOfString:"), objc.String(str))
	return rv
}

// Returns a representation of the receiver as a C string in the default C-string encoding, possibly losing information in converting to that encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lossyCString()
func (s_ String) LossyCString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lossyCString"))
	return rv
}

// Returns a version of the string with all letters converted to lowercase, taking into account the specified locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lowercased(with:)
func (s_ String) LowercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lowercaseStringWithLocale:"), locale)
	return rv
}

// Returns the maximum number of bytes needed to store the receiver in a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/maximumLengthOfBytes(using:)
func (s_ String) MaximumLengthOfBytesUsingEncoding(enc unsafe.Pointer) uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("maximumLengthOfBytesUsingEncoding:"), enc)
	return rv
}

// Returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/padding(toLength:withPad:startingAt:)
func (s_ String) StringByPaddingToLengthWithStringStartingAtIndex(newLength uint, padString string, padIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByPaddingToLength:withString:startingAtIndex:"), newLength, objc.String(padString), padIndex)
	return rv
}

// Returns the range of characters representing the paragraph or paragraphs containing a given range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/paragraphRange(for:)
func (s_ String) ParagraphRangeForRange(range_ Range) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("paragraphRangeForRange:"), range_)
	return rv
}

// Parses the receiver as a text representation of a property list, returning an , , , or object, according to the topmost element.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/propertyList()
func (s_ String) PropertyList() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("propertyList"))
	return rv
}

// Returns a dictionary object initialized with the keys and values found in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/propertyListFromStringsFileFormat()
func (s_ String) PropertyListFromStringsFileFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("propertyListFromStringsFileFormat"))
	return rv
}

// Finds and returns the range of the first occurrence of a given string within the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:)
func (s_ String) RangeOfString(searchString string) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfString:"), objc.String(searchString))
	return rv
}

// Finds and returns the range of the first occurrence of a given string within the string, subject to given options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:)
func (s_ String) RangeOfStringOptions(searchString string, mask unsafe.Pointer) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfString:options:"), objc.String(searchString), mask)
	return rv
}

// Finds and returns the range of the first occurrence of a given string, within the given range of the string, subject to given options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:range:)
func (s_ String) RangeOfStringOptionsRange(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch Range) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfString:options:range:"), objc.String(searchString), mask, rangeOfReceiverToSearch)
	return rv
}

// Finds and returns the range of the first occurrence of a given string within a given range of the string, subject to given options, using the specified locale, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:range:locale:)
func (s_ String) RangeOfStringOptionsRangeLocale(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch Range, locale unsafe.Pointer) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfString:options:range:locale:"), objc.String(searchString), mask, rangeOfReceiverToSearch, locale)
	return rv
}

// Finds and returns the range in the string of the first character from a given character set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:)
func (s_ String) RangeOfCharacterFromSet(searchSet unsafe.Pointer) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfCharacterFromSet:"), searchSet)
	return rv
}

// Finds and returns the range in the string of the first character, using given options, from a given character set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:options:)
func (s_ String) RangeOfCharacterFromSetOptions(searchSet unsafe.Pointer, mask unsafe.Pointer) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfCharacterFromSet:options:"), searchSet, mask)
	return rv
}

// Finds and returns the range in the string of the first character from a given character set found in a given range with given options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:options:range:)
func (s_ String) RangeOfCharacterFromSetOptionsRange(searchSet unsafe.Pointer, mask unsafe.Pointer, rangeOfReceiverToSearch Range) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfCharacterFromSet:options:range:"), searchSet, mask, rangeOfReceiverToSearch)
	return rv
}

// Returns the range in the receiver of the composed character sequence located at a given index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfComposedCharacterSequence(at:)
func (s_ String) RangeOfComposedCharacterSequenceAtIndex(index uint) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfComposedCharacterSequenceAtIndex:"), index)
	return rv
}

// Returns the range in the string of the composed character sequences for a given range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/rangeOfComposedCharacterSequences(for:)
func (s_ String) RangeOfComposedCharacterSequencesForRange(range_ Range) Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("rangeOfComposedCharacterSequencesForRange:"), range_)
	return rv
}

// Returns a new string in which the characters in a specified range of the receiver are replaced by a given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingCharacters(in:with:)
func (s_ String) StringByReplacingCharactersInRangeWithString(range_ Range, replacement string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByReplacingCharactersInRange:withString:"), range_, objc.String(replacement))
	return rv
}

// Returns a new string in which all occurrences of a target string in the receiver are replaced by another given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingOccurrences(of:with:)
func (s_ String) StringByReplacingOccurrencesOfStringWithString(target string, replacement string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByReplacingOccurrencesOfString:withString:"), objc.String(target), objc.String(replacement))
	return rv
}

// Returns a new string in which all occurrences of a target string in a specified range of the receiver are replaced by another given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingOccurrences(of:with:options:range:)
func (s_ String) StringByReplacingOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options unsafe.Pointer, searchRange Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByReplacingOccurrencesOfString:withString:options:range:"), objc.String(target), objc.String(replacement), options, searchRange)
	return rv
}

// Returns a new string made by replacing in the receiver all percent escapes with the matching characters as determined by a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/replacingPercentEscapes(using:)
func (s_ String) StringByReplacingPercentEscapesUsingEncoding(enc unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByReplacingPercentEscapesUsingEncoding:"), enc)
	return rv
}

// Returns the bounding box size the receiver occupies when drawn with the given attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/size(withAttributes:)
func (s_ String) SizeWithAttributes(attrs unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("sizeWithAttributes:"), attrs)
	return rv
}

// Returns the size of the string if it were to be rendered with the specified font on a single line.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:
func (s_ String) SizeWithFont(font unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("sizeWithFont:"), font)
	return rv
}

// Returns the size of the string if it were rendered and constrained to the specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:constrainedToSize:
func (s_ String) SizeWithFontConstrainedToSize(font unsafe.Pointer, size coregraphics.CGSize) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("sizeWithFont:constrainedToSize:"), font, size)
	return rv
}

// Returns the size of the string if it were rendered with the specified constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:constrainedToSize:lineBreakMode:
func (s_ String) SizeWithFontConstrainedToSizeLineBreakMode(font unsafe.Pointer, size coregraphics.CGSize, lineBreakMode unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("sizeWithFont:constrainedToSize:lineBreakMode:"), font, size, lineBreakMode)
	return rv
}

// Returns the size of the string if it were to be rendered with the specified font and line attributes on a single line.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:forWidth:lineBreakMode:
func (s_ String) SizeWithFontForWidthLineBreakMode(font unsafe.Pointer, width float64, lineBreakMode unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("sizeWithFont:forWidth:lineBreakMode:"), font, width, lineBreakMode)
	return rv
}

// Returns the size of the string if it were rendered with the specified constraints, including a variable font size, on a single line.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sizeWithFont:minFontSize:actualFontSize:forWidth:lineBreakMode:
func (s_ String) SizeWithFontMinFontSizeActualFontSizeForWidthLineBreakMode(font unsafe.Pointer, minFontSize float64, actualFontSize float64, width float64, lineBreakMode unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("sizeWithFont:minFontSize:actualFontSize:forWidth:lineBreakMode:"), font, minFontSize, actualFontSize, width, lineBreakMode)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/sr_sensorForDeletionRecordsFromSensor()
func (s_ String) Sr_sensorForDeletionRecordsFromSensor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sr_sensorForDeletionRecordsFromSensor"))
	return rv
}

// Returns a string made by appending to the receiver a string constructed from a given format string and the following arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/stringByAppendingFormat:
func (s_ String) StringByAppendingFormat(format string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAppendingFormat:"), objc.String(format))
	return rv
}

// Returns an array of strings made by separately appending to the receiver each string in a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/strings(byAppendingPaths:)
func (s_ String) StringsByAppendingPaths(paths unsafe.Pointer) []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("stringsByAppendingPaths:"), paths)
	return rv
}

// Returns a new string containing the characters of the receiver from the one at a given index to the end.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/substring(from:)
func (s_ String) SubstringFromIndex(from uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("substringFromIndex:"), from)
	return rv
}

// Returns a new string containing the characters of the receiver up to, but not including, the one at a given index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/substring(to:)
func (s_ String) SubstringToIndex(to uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("substringToIndex:"), to)
	return rv
}

// Returns a string object containing the characters of the receiver that lie within a given range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/substring(with:)
func (s_ String) SubstringWithRange(range_ Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("substringWithRange:"), range_)
	return rv
}

// Returns a new string made by removing from both ends of the receiver characters contained in a given character set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/trimmingCharacters(in:)
func (s_ String) StringByTrimmingCharactersInSet(set unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByTrimmingCharactersInSet:"), set)
	return rv
}

// Returns a version of the string with all letters converted to uppercase, taking into account the specified locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/uppercased(with:)
func (s_ String) UppercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("uppercaseStringWithLocale:"), locale)
	return rv
}

// Returns a string variation suitable for the specified presentation width.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/variantFittingPresentationWidth(_:)
func (s_ String) VariantFittingPresentationWidth(width int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("variantFittingPresentationWidth:"), width)
	return rv
}

// Writes the contents of the receiver to the location specified by a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(to:atomically:)
func (s_ String) WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToURL:atomically:"), url, atomically)
	return rv
}

// Writes the contents of the receiver to the URL specified by using the specified encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(to:atomically:encoding:)
func (s_ String) WriteToURLAtomicallyEncodingError(url unsafe.Pointer, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToURL:atomically:encoding:error:"), url, useAuxiliaryFile, enc, error)
	return rv
}

// Writes the contents of the receiver to the file specified by a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(toFile:atomically:)
func (s_ String) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToFile:atomically:"), objc.String(path), useAuxiliaryFile)
	return rv
}

// Writes the contents of the receiver to a file at a given path using a given encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/write(toFile:atomically:encoding:)
func (s_ String) WriteToFileAtomicallyEncodingError(path string, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("writeToFile:atomically:encoding:error:"), objc.String(path), useAuxiliaryFile, enc, error)
	return rv
}

// A new string that replaces the current home directory portion of the current path with a tilde ( ) character.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/abbreviatingWithTildeInPath
func (s_ String) StringByAbbreviatingWithTildeInPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByAbbreviatingWithTildeInPath"))
	return rv
}


// The Boolean value of the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/boolValue
func (s_ String) BoolValue() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("boolValue"))
	return rv
}


// A capitalized representation of the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/capitalized
func (s_ String) CapitalizedString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("capitalizedString"))
	return rv
}


// A string made by normalizing the string’s contents using the Unicode Normalization Form D.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/decomposedStringWithCanonicalMapping
func (s_ String) DecomposedStringWithCanonicalMapping() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("decomposedStringWithCanonicalMapping"))
	return rv
}


// A string made by normalizing the receiver’s contents using the Unicode Normalization Form KD.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/decomposedStringWithCompatibilityMapping
func (s_ String) DecomposedStringWithCompatibilityMapping() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("decomposedStringWithCompatibilityMapping"))
	return rv
}


// A new string made by deleting the last path component from the receiver, along with any final path separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deletingLastPathComponent
func (s_ String) StringByDeletingLastPathComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByDeletingLastPathComponent"))
	return rv
}


// A new string made by deleting the extension (if any, and only the last) from the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/deletingPathExtension
func (s_ String) StringByDeletingPathExtension() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByDeletingPathExtension"))
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/description
func (s_ String) Description() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("description"))
	return rv
}


// The floating-point value of the string as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/doubleValue
func (s_ String) DoubleValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("doubleValue"))
	return rv
}


// A new string made by expanding the initial component of the receiver to its full path value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/expandingTildeInPath
func (s_ String) StringByExpandingTildeInPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByExpandingTildeInPath"))
	return rv
}


// The fastest encoding to which the receiver may be converted without loss of information.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/fastestEncoding
func (s_ String) FastestEncoding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("fastestEncoding"))
	return rv
}


// A file system-specific representation of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/fileSystemRepresentation
func (s_ String) FileSystemRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("fileSystemRepresentation"))
	return rv
}


// The floating-point value of the string as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/floatValue
func (s_ String) FloatValue() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("floatValue"))
	return rv
}


// An unsigned integer that can be used as a hash table address.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/hash
func (s_ String) Hash() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("hash"))
	return rv
}


// The integer value of the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/intValue
func (s_ String) IntValue() int {
	rv := objc.Send[int](s_.ID, objc.Sel("intValue"))
	return rv
}


// The value of the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/integerValue
func (s_ String) IntegerValue() int {
	rv := objc.Send[int](s_.ID, objc.Sel("integerValue"))
	return rv
}


// A Boolean value that indicates whether the receiver represents an absolute path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/isAbsolutePath
func (s_ String) AbsolutePath() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("absolutePath"))
	return rv
}


// The last path component of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lastPathComponent
func (s_ String) LastPathComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lastPathComponent"))
	return rv
}


// The number of UTF-16 code units in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/length
func (s_ String) Length() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("length"))
	return rv
}


// Returns a capitalized representation of the receiver using the current locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedCapitalized
func (s_ String) LocalizedCapitalizedString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedCapitalizedString"))
	return rv
}


// Returns a version of the string with all letters converted to lowercase, taking into account the current locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedLowercase
func (s_ String) LocalizedLowercaseString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedLowercaseString"))
	return rv
}


// Returns a version of the string with all letters converted to uppercase, taking into account the current locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/localizedUppercase
func (s_ String) LocalizedUppercaseString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("localizedUppercaseString"))
	return rv
}


// The value of the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/longLongValue
func (s_ String) LongLongValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("longLongValue"))
	return rv
}


// A lowercase representation of the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/lowercased
func (s_ String) LowercaseString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lowercaseString"))
	return rv
}


// The file-system path components of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/pathComponents
func (s_ String) PathComponents() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("pathComponents"))
	return rv
}


// The path extension, if any, of the string as interpreted as a path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/pathExtension
func (s_ String) PathExtension() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("pathExtension"))
	return rv
}


// A string made by normalizing the string’s contents using the Unicode Normalization Form C.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/precomposedStringWithCanonicalMapping
func (s_ String) PrecomposedStringWithCanonicalMapping() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("precomposedStringWithCanonicalMapping"))
	return rv
}


// A string made by normalizing the receiver’s contents using the Unicode Normalization Form KC.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/precomposedStringWithCompatibilityMapping
func (s_ String) PrecomposedStringWithCompatibilityMapping() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("precomposedStringWithCompatibilityMapping"))
	return rv
}


// Returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/removingPercentEncoding
func (s_ String) StringByRemovingPercentEncoding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByRemovingPercentEncoding"))
	return rv
}


// A new string made from the receiver by resolving all symbolic links and standardizing path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/resolvingSymlinksInPath
func (s_ String) StringByResolvingSymlinksInPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByResolvingSymlinksInPath"))
	return rv
}


// The smallest encoding to which the receiver can be converted without loss of information.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/smallestEncoding
func (s_ String) SmallestEncoding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("smallestEncoding"))
	return rv
}


// A new string made by removing extraneous path components from the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/standardizingPath
func (s_ String) StringByStandardizingPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stringByStandardizingPath"))
	return rv
}


// An uppercase representation of the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/uppercased
func (s_ String) UppercaseString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("uppercaseString"))
	return rv
}


// A null-terminated UTF8 representation of the string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/utf8String
func (s_ String) UTF8String() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("UTF8String"))
	return rv
}



