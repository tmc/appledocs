// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [String] class.
var StringClass objc.Class

func init() {
	StringClass = objc.GetClass("NSString")
}

type String struct {
	objc.ID
}

func StringFrom(ptr unsafe.Pointer) String {
	return String{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc String) Alloc() String {
	ret := objc.ID(StringClass).Send(objc.RegisterName("alloc"))
	return String{ret}
}

// Init initializes the instance.
func (s_ String) Init() String {
	ret := s_.ID.Send(objc.RegisterName("init"))
	return String{ret}
}
// Returns an initialized   object that contains no characters. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init()
func NewString() String {
	instance := String{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes the receiver, a newly allocated   object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(CString:)
func NewStringWithCString(bytes unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithCString:")
	ret := instance.ID.Send(sel, bytes)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized using the characters in a given C array, interpreted according to a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(CString:encoding:)-20f9h
func NewStringWithCStringEncoding(nullTerminatedCString unsafe.Pointer, encoding unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithCString:encoding:")
	ret := instance.ID.Send(sel, nullTerminatedCString, encoding)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes the receiver, a newly allocated   object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(CString:length:)
func NewStringWithCStringLength(bytes unsafe.Pointer, length uint) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithCString:length:")
	ret := instance.ID.Send(sel, bytes, length)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes the receiver, a newly allocated   object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(CStringNoCopy:length:freeWhenDone:)
func NewStringWithCStringNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, freeBuffer bool) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithCStringNoCopy:length:freeWhenDone:")
	ret := instance.ID.Send(sel, bytes, length, freeBuffer)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by copying the characters from a given C array of UTF8-encoded bytes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(UTF8String:)-vg2b
func NewStringWithUTF8String(nullTerminatedCString unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithUTF8String:")
	ret := instance.ID.Send(sel, nullTerminatedCString)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an initialized   object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(bytes:length:encoding:)
func NewStringWithBytesLengthEncoding(bytes unsafe.Pointer, len uint, encoding unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithBytes:length:encoding:")
	ret := instance.ID.Send(sel, bytes, len, encoding)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:deallocator:)
func NewStringWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len uint, encoding unsafe.Pointer, deallocator unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithBytesNoCopy:length:encoding:deallocator:")
	ret := instance.ID.Send(sel, bytes, len, encoding, deallocator)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an initialized   object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:freeWhenDone:)
func NewStringWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len uint, encoding unsafe.Pointer, freeBuffer bool) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithBytesNoCopy:length:encoding:freeWhenDone:")
	ret := instance.ID.Send(sel, bytes, len, encoding, freeBuffer)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an initialized   object that contains a given number of characters from a given C array of UTF-16 code units. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(characters:length:)
func NewStringWithCharactersLength(characters unsafe.Pointer, length uint) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithCharacters:length:")
	ret := instance.ID.Send(sel, characters, length)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(charactersNoCopy:length:deallocator:)
func NewStringWithCharactersNoCopyLengthDeallocator(chars unsafe.Pointer, len uint, deallocator unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithCharactersNoCopy:length:deallocator:")
	ret := instance.ID.Send(sel, chars, len, deallocator)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an initialized   object that contains a given number of characters from a given C array of UTF-16 code units. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(charactersNoCopy:length:freeWhenDone:)
func NewStringWithCharactersNoCopyLengthFreeWhenDone(characters unsafe.Pointer, length uint, freeBuffer bool) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithCharactersNoCopy:length:freeWhenDone:")
	ret := instance.ID.Send(sel, characters, length, freeBuffer)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(coder:)
func NewStringWithCoder(coder unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes the receiver, a newly allocated   object, by reading data from the file named by  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(contentsOfFile:)
func NewStringWithContentsOfFile(path string) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfFile:")
	ret := instance.ID.Send(sel, path)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by reading data from the file at a given path using a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(contentsOfFile:encoding:)
func NewStringWithContentsOfFileEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfFile:encoding:error:")
	ret := instance.ID.Send(sel, path, enc, error)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by reading data from the file at a given path and returns by reference the encoding used to interpret the characters. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(contentsOfFile:usedEncoding:)
func NewStringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfFile:usedEncoding:error:")
	ret := instance.ID.Send(sel, path, enc, error)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes the receiver, a newly allocated   object, by reading data from the location named by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(contentsOfURL:)
func NewStringWithContentsOfURL(url unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfURL:")
	ret := instance.ID.Send(sel, url)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by reading data from a given URL interpreted using a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(contentsOfURL:encoding:)-715fw
func NewStringWithContentsOfURLEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfURL:encoding:error:")
	ret := instance.ID.Send(sel, url, enc, error)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by reading data from a given URL and returns by reference the encoding used to interpret the data. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(contentsOfURL:usedEncoding:)-2c72d
func NewStringWithContentsOfURLUsedEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfURL:usedEncoding:error:")
	ret := instance.ID.Send(sel, url, enc, error)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by converting given data into UTF-16 code units using a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(data:encoding:)
func NewStringWithDataEncoding(data unsafe.Pointer, encoding unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithData:encoding:")
	ret := instance.ID.Send(sel, data, encoding)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted without any localization. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(format:arguments:)
func NewStringWithFormatArguments(format string, argList unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithFormat:arguments:")
	ret := instance.ID.Send(sel, format, argList)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale information. This method is meant to be called from within a variadic function, where the argument list will be available. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(format:locale:arguments:)
func NewStringWithFormatLocaleArguments(format string, locale objc.ID, argList unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithFormat:locale:arguments:")
	ret := instance.ID.Send(sel, format, locale, argList)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by copying the characters from another given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(string:)-210xa
func NewStringWithString(aString string) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithString:")
	ret := instance.ID.Send(sel, aString)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/initWithFormat:
func NewStringWithFormat(format string) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithFormat:")
	ret := instance.ID.Send(sel, format)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/initWithFormat:locale:
func NewStringWithFormatLocale(format string, locale objc.ID) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithFormat:locale:")
	ret := instance.ID.Send(sel, format, locale)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersArgumentsError(format string, validFormatSpecifiers string, argList unsafe.Pointer, error unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithValidatedFormat:validFormatSpecifiers:arguments:error:")
	ret := instance.ID.Send(sel, format, validFormatSpecifiers, argList, error)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:error:
func NewStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string, error unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithValidatedFormat:validFormatSpecifiers:error:")
	ret := instance.ID.Send(sel, format, validFormatSpecifiers, error)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format string, validFormatSpecifiers string, locale objc.ID, argList unsafe.Pointer, error unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:")
	ret := instance.ID.Send(sel, format, validFormatSpecifiers, locale, argList, error)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleError(format string, validFormatSpecifiers string, locale objc.ID, error unsafe.Pointer) String {
	instance := String{}.Alloc()
	sel := objc.RegisterName("initWithValidatedFormat:validFormatSpecifiers:locale:error:")
	ret := instance.ID.Send(sel, format, validFormatSpecifiers, locale, error)
	instance = String{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:
func (sc String) DeferredLocalizedIntentsStringWithFormat(format string) unsafe.Pointer {
	sel := objc.RegisterName("deferredLocalizedIntentsStringWithFormat:")
	ret := objc.ID(StringClass).Send(sel, format)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:
func (sc String) DeferredLocalizedIntentsStringWithFormatFromTable(format string, table string) unsafe.Pointer {
	sel := objc.RegisterName("deferredLocalizedIntentsStringWithFormat:fromTable:")
	ret := objc.ID(StringClass).Send(sel, format, table)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:arguments:
func (sc String) DeferredLocalizedIntentsStringWithFormatFromTableArguments(format string, table string, arguments unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("deferredLocalizedIntentsStringWithFormat:fromTable:arguments:")
	ret := objc.ID(StringClass).Send(sel, format, table, arguments)
	return unsafe.Pointer(ret)
}
// Returns a string containing the bytes in a given C array, interpreted according to a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(CString:encoding:)-7auq8
func (sc String) StringWithCStringEncoding(cString unsafe.Pointer, enc unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringWithCString:encoding:")
	ret := objc.ID(StringClass).Send(sel, cString, enc)
	return unsafe.Pointer(ret)
}
// Returns a string created by copying the data from a given C array of UTF8-encoded bytes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(UTF8String:)-8bcy8
func (sc String) StringWithUTF8String(nullTerminatedCString unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringWithUTF8String:")
	ret := objc.ID(StringClass).Send(sel, nullTerminatedCString)
	return unsafe.Pointer(ret)
}
// Returns a string created by reading data from a given URL interpreted using a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(contentsOfURL:encoding:)-x6cv
func (sc String) StringWithContentsOfURLEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringWithContentsOfURL:encoding:error:")
	ret := objc.ID(StringClass).Send(sel, url, enc, error)
	return unsafe.Pointer(ret)
}
// Returns a string created by reading data from a given URL and returns by reference the encoding used to interpret the data. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/init(contentsOfURL:usedEncoding:)-9jrum
func (sc String) StringWithContentsOfURLUsedEncodingError(url unsafe.Pointer, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringWithContentsOfURL:usedEncoding:error:")
	ret := objc.ID(StringClass).Send(sel, url, enc, error)
	return unsafe.Pointer(ret)
}
// Returns a human-readable string giving the name of a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedName(of:)
func (sc String) LocalizedNameOfStringEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("localizedNameOfStringEncoding:")
	ret := objc.ID(StringClass).Send(sel, encoding)
	return unsafe.Pointer(ret)
}
// Returns a string created by using a given format string as a template into which the remaining argument values are substituted according to the current locale. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedStringWithFormat:
func (sc String) LocalizedStringWithFormat(format string) unsafe.Pointer {
	sel := objc.RegisterName("localizedStringWithFormat:")
	ret := objc.ID(StringClass).Send(sel, format)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedStringWithValidatedFormat:validFormatSpecifiers:error:
func (sc String) LocalizedStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("localizedStringWithValidatedFormat:validFormatSpecifiers:error:")
	ret := objc.ID(StringClass).Send(sel, format, validFormatSpecifiers, error)
	return unsafe.Pointer(ret)
}
// Returns a localized string intended for display in a notification alert. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedUserNotificationString(forKey:arguments:)
func (sc String) LocalizedUserNotificationStringForKeyArguments(key string, arguments unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("localizedUserNotificationStringForKey:arguments:")
	ret := objc.ID(StringClass).Send(sel, key, arguments)
	return unsafe.Pointer(ret)
}
// Returns a string built from the strings in a given array by concatenating them with a path separator between each pair. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/path(withComponents:)
func (sc String) PathWithComponents(components unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("pathWithComponents:")
	ret := objc.ID(StringClass).Send(sel, components)
	return unsafe.Pointer(ret)
}
// Returns an empty string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/string
func (sc String) String() unsafe.Pointer {
	sel := objc.RegisterName("string")
	ret := objc.ID(StringClass).Send(sel)
	return unsafe.Pointer(ret)
}
// Creates a new string using a given C-string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/string(withCString:)
func (sc String) StringWithCString(bytes unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("stringWithCString:")
	ret := objc.ID(StringClass).Send(sel, bytes)
	return objc.ID(ret)
}
// Returns a string containing the characters in a given C-string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/string(withCString:length:)
func (sc String) StringWithCStringLength(bytes unsafe.Pointer, length uint) objc.ID {
	sel := objc.RegisterName("stringWithCString:length:")
	ret := objc.ID(StringClass).Send(sel, bytes, length)
	return objc.ID(ret)
}
// Returns a string created by reading data from the file named by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/string(withContentsOf:)
func (sc String) StringWithContentsOfURL(url unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("stringWithContentsOfURL:")
	ret := objc.ID(StringClass).Send(sel, url)
	return objc.ID(ret)
}
// Returns a string created by reading data from the file named by a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/string(withContentsOfFile:)
func (sc String) StringWithContentsOfFile(path string) objc.ID {
	sel := objc.RegisterName("stringWithContentsOfFile:")
	ret := objc.ID(StringClass).Send(sel, path)
	return objc.ID(ret)
}
// Returns the string encoding for the given data as detected by attempting to create a string according to the specified encoding options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/stringEncoding(for:encodingOptions:convertedString:usedLossyConversion:)
func (sc String) StringEncodingForDataEncodingOptionsConvertedStringUsedLossyConversion(data unsafe.Pointer, opts unsafe.Pointer, string string, usedLossyConversion unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringEncodingForData:encodingOptions:convertedString:usedLossyConversion:")
	ret := objc.ID(StringClass).Send(sel, data, opts, string, usedLossyConversion)
	return unsafe.Pointer(ret)
}
// Returns a string containing a given number of characters taken from a given C array of UTF-16 code units. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/stringWithCharacters:length:
func (sc String) StringWithCharactersLength(characters unsafe.Pointer, length uint) unsafe.Pointer {
	sel := objc.RegisterName("stringWithCharacters:length:")
	ret := objc.ID(StringClass).Send(sel, characters, length)
	return unsafe.Pointer(ret)
}
// Returns a string created by reading data from the file at a given path interpreted using a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/stringWithContentsOfFile:encoding:error:
func (sc String) StringWithContentsOfFileEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringWithContentsOfFile:encoding:error:")
	ret := objc.ID(StringClass).Send(sel, path, enc, error)
	return unsafe.Pointer(ret)
}
// Returns a string created by reading data from the file at a given path and returns by reference the encoding used to interpret the file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/stringWithContentsOfFile:usedEncoding:error:
func (sc String) StringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringWithContentsOfFile:usedEncoding:error:")
	ret := objc.ID(StringClass).Send(sel, path, enc, error)
	return unsafe.Pointer(ret)
}
// Returns a string created by using a given format string as a template into which the remaining argument values are substituted. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/stringWithFormat:
func (sc String) StringWithFormat(format string) unsafe.Pointer {
	sel := objc.RegisterName("stringWithFormat:")
	ret := objc.ID(StringClass).Send(sel, format)
	return unsafe.Pointer(ret)
}
// Returns a string created by copying the characters from another given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/stringWithString:
func (sc String) StringWithString(string string) unsafe.Pointer {
	sel := objc.RegisterName("stringWithString:")
	ret := objc.ID(StringClass).Send(sel, string)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/stringWithValidatedFormat:validFormatSpecifiers:error:
func (sc String) StringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringWithValidatedFormat:validFormatSpecifiers:error:")
	ret := objc.ID(StringClass).Send(sel, format, validFormatSpecifiers, error)
	return unsafe.Pointer(ret)
}
// Returns a new string made from the receiver by replacing all characters not in the specified set with percent-encoded characters. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/addingPercentEncoding(withAllowedCharacters:)
func (s_ String) StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringByAddingPercentEncodingWithAllowedCharacters:")
	ret := s_.ID.Send(sel, allowedCharacters)
	return unsafe.Pointer(ret)
}
// Returns a representation of the receiver using a given encoding to determine the percent escapes necessary to convert the receiver into a legal URL string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/addingPercentEscapes(using:)
func (s_ String) StringByAddingPercentEscapesUsingEncoding(enc unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringByAddingPercentEscapesUsingEncoding:")
	ret := s_.ID.Send(sel, enc)
	return unsafe.Pointer(ret)
}
// Returns a new string made by appending a given string to the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/appending(_:)
func (s_ String) StringByAppendingString(aString string) unsafe.Pointer {
	sel := objc.RegisterName("stringByAppendingString:")
	ret := s_.ID.Send(sel, aString)
	return unsafe.Pointer(ret)
}
// Returns a new string made by appending to the receiver a given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/appendingPathComponent(_:)
func (s_ String) StringByAppendingPathComponent(str string) unsafe.Pointer {
	sel := objc.RegisterName("stringByAppendingPathComponent:")
	ret := s_.ID.Send(sel, str)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/appendingPathComponent(_:conformingTo:)
func (s_ String) StringByAppendingPathComponentConformingToType(partialName string, contentType unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringByAppendingPathComponent:conformingToType:")
	ret := s_.ID.Send(sel, partialName, contentType)
	return unsafe.Pointer(ret)
}
// Returns a new string made by appending to the receiver an extension separator followed by a given extension. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/appendingPathExtension(_:)
func (s_ String) StringByAppendingPathExtension(str string) unsafe.Pointer {
	sel := objc.RegisterName("stringByAppendingPathExtension:")
	ret := s_.ID.Send(sel, str)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/appendingPathExtension(for:)
func (s_ String) StringByAppendingPathExtensionForType(contentType unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringByAppendingPathExtensionForType:")
	ret := s_.ID.Send(sel, contentType)
	return unsafe.Pointer(ret)
}
// Returns a new string by applying a specified transform to the string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/applyingTransform(_:reverse:)
func (s_ String) StringByApplyingTransformReverse(transform unsafe.Pointer, reverse bool) unsafe.Pointer {
	sel := objc.RegisterName("stringByApplyingTransform:reverse:")
	ret := s_.ID.Send(sel, transform, reverse)
	return unsafe.Pointer(ret)
}
// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/boundingRect(with:options:attributes:)
func (s_ String) BoundingRectWithSizeOptionsAttributes(size Size, options unsafe.Pointer, attributes unsafe.Pointer) Rect {
	sel := objc.RegisterName("boundingRectWithSize:options:attributes:")
	ret := s_.ID.Send(sel, size, options, attributes)
	return Rect(ret)
}
// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/boundingRect(with:options:attributes:context:)
func (s_ String) BoundingRectWithSizeOptionsAttributesContext(size Size, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer) Rect {
	sel := objc.RegisterName("boundingRectWithSize:options:attributes:context:")
	ret := s_.ID.Send(sel, size, options, attributes, context)
	return Rect(ret)
}
// Returns a representation of the receiver as a C string in the default C-string encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/cString()
func (s_ String) CString() unsafe.Pointer {
	sel := objc.RegisterName("cString")
	ret := s_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns a representation of the string as a C string using a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/cString(using:)
func (s_ String) CStringUsingEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("cStringUsingEncoding:")
	ret := s_.ID.Send(sel, encoding)
	return unsafe.Pointer(ret)
}
// Returns the length in char-sized units of the receiver’s C-string representation in the default C-string encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/cStringLength()
func (s_ String) CStringLength() uint {
	sel := objc.RegisterName("cStringLength")
	ret := s_.ID.Send(sel)
	return uint(ret)
}
// Returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/canBeConverted(to:)
func (s_ String) CanBeConvertedToEncoding(encoding unsafe.Pointer) bool {
	sel := objc.RegisterName("canBeConvertedToEncoding:")
	ret := s_.ID.Send(sel, encoding)
	return ret != 0
}
// Returns a capitalized representation of the receiver using the specified locale. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/capitalized(with:)
func (s_ String) CapitalizedStringWithLocale(locale unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("capitalizedStringWithLocale:")
	ret := s_.ID.Send(sel, locale)
	return unsafe.Pointer(ret)
}
// Returns the result of invoking   with   as the only option. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/caseInsensitiveCompare(_:)
func (s_ String) CaseInsensitiveCompare(string string) unsafe.Pointer {
	sel := objc.RegisterName("caseInsensitiveCompare:")
	ret := s_.ID.Send(sel, string)
	return unsafe.Pointer(ret)
}
// Returns the character at a given UTF-16 code unit index. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/character(at:)
func (s_ String) CharacterAtIndex(index uint) unsafe.Pointer {
	sel := objc.RegisterName("characterAtIndex:")
	ret := s_.ID.Send(sel, index)
	return unsafe.Pointer(ret)
}
// Returns a string containing characters the receiver and a given string have in common, starting from the beginning of each up to the first characters that aren’t equivalent. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/commonPrefix(with:options:)
func (s_ String) CommonPrefixWithStringOptions(str string, mask unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("commonPrefixWithString:options:")
	ret := s_.ID.Send(sel, str, mask)
	return unsafe.Pointer(ret)
}
// Returns the result of invoking   with no options and the receiver’s full extent as the range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:)
func (s_ String) Compare(string string) unsafe.Pointer {
	sel := objc.RegisterName("compare:")
	ret := s_.ID.Send(sel, string)
	return unsafe.Pointer(ret)
}
// Compares the string with the specified string using the given options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:)
func (s_ String) CompareOptions(string string, mask unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("compare:options:")
	ret := s_.ID.Send(sel, string, mask)
	return unsafe.Pointer(ret)
}
// Returns the result of invoking   with a   locale. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:range:)
func (s_ String) CompareOptionsRange(string string, mask unsafe.Pointer, rangeOfReceiverToCompare Range) unsafe.Pointer {
	sel := objc.RegisterName("compare:options:range:")
	ret := s_.ID.Send(sel, string, mask, rangeOfReceiverToCompare)
	return unsafe.Pointer(ret)
}
// Compares the string using the specified options and returns the lexical ordering for the range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:range:locale:)
func (s_ String) CompareOptionsRangeLocale(string string, mask unsafe.Pointer, rangeOfReceiverToCompare Range, locale objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("compare:options:range:locale:")
	ret := s_.ID.Send(sel, string, mask, rangeOfReceiverToCompare, locale)
	return unsafe.Pointer(ret)
}
// Interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/completePath(into:caseSensitive:matchesInto:filterTypes:)
func (s_ String) CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName string, flag bool, outputArray unsafe.Pointer, filterTypes unsafe.Pointer) uint {
	sel := objc.RegisterName("completePathIntoString:caseSensitive:matchesIntoArray:filterTypes:")
	ret := s_.ID.Send(sel, outputName, flag, outputArray, filterTypes)
	return uint(ret)
}
// Returns an array containing substrings from the receiver that have been divided by a given separator. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/components(separatedBy:)-238fy
func (s_ String) ComponentsSeparatedByString(separator string) unsafe.Pointer {
	sel := objc.RegisterName("componentsSeparatedByString:")
	ret := s_.ID.Send(sel, separator)
	return unsafe.Pointer(ret)
}
// Returns an array containing substrings from the receiver that have been divided by characters in a given set. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/components(separatedBy:)-27x9g
func (s_ String) ComponentsSeparatedByCharactersInSet(separator unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("componentsSeparatedByCharactersInSet:")
	ret := s_.ID.Send(sel, separator)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/contains(_:)
func (s_ String) ContainsString(str string) bool {
	sel := objc.RegisterName("containsString:")
	ret := s_.ID.Send(sel, str)
	return ret != 0
}
// Returns an   object containing a representation of the receiver encoded using a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/data(using:)
func (s_ String) DataUsingEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataUsingEncoding:")
	ret := s_.ID.Send(sel, encoding)
	return unsafe.Pointer(ret)
}
// Returns an   object containing a representation of the receiver encoded using a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/data(using:allowLossyConversion:)
func (s_ String) DataUsingEncodingAllowLossyConversion(encoding unsafe.Pointer, lossy bool) unsafe.Pointer {
	sel := objc.RegisterName("dataUsingEncoding:allowLossyConversion:")
	ret := s_.ID.Send(sel, encoding, lossy)
	return unsafe.Pointer(ret)
}
// Draws the receiver with the font and other display characteristics of the given attributes, at the specified point in the current graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/draw(at:withAttributes:)
func (s_ String) DrawAtPointWithAttributes(point Point, attrs unsafe.Pointer) {
	sel := objc.RegisterName("drawAtPoint:withAttributes:")
	s_.ID.Send(sel, point, attrs)
}
// Draws the attributed string inside the specified bounding rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/draw(in:withAttributes:)
func (s_ String) DrawInRectWithAttributes(rect Rect, attrs unsafe.Pointer) {
	sel := objc.RegisterName("drawInRect:withAttributes:")
	s_.ID.Send(sel, rect, attrs)
}
// Draws the receiver with the specified options and other display characteristics of the given attributes, within the specified rectangle in the current graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/draw(with:options:attributes:)
func (s_ String) DrawWithRectOptionsAttributes(rect Rect, options unsafe.Pointer, attributes unsafe.Pointer) {
	sel := objc.RegisterName("drawWithRect:options:attributes:")
	s_.ID.Send(sel, rect, options, attributes)
}
// Draws the attributed string in the specified bounding rectangle using the provided options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/draw(with:options:attributes:context:)
func (s_ String) DrawWithRectOptionsAttributesContext(rect Rect, options unsafe.Pointer, attributes unsafe.Pointer, context unsafe.Pointer) {
	sel := objc.RegisterName("drawWithRect:options:attributes:context:")
	s_.ID.Send(sel, rect, options, attributes, context)
}
// Draws the string in a single line at the specified point in the current graphics context using the specified font and attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/drawAtPoint:forWidth:withFont:fontSize:lineBreakMode:baselineAdjustment:
func (s_ String) DrawAtPointForWidthWithFontFontSizeLineBreakModeBaselineAdjustment(point Point, width float64, font unsafe.Pointer, fontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) Size {
	sel := objc.RegisterName("drawAtPoint:forWidth:withFont:fontSize:lineBreakMode:baselineAdjustment:")
	ret := s_.ID.Send(sel, point, width, font, fontSize, lineBreakMode, baselineAdjustment)
	return Size(ret)
}
// Draws the string in a single line at the specified point in the current graphics context using the specified font and attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/drawAtPoint:forWidth:withFont:lineBreakMode:
func (s_ String) DrawAtPointForWidthWithFontLineBreakMode(point Point, width float64, font unsafe.Pointer, lineBreakMode unsafe.Pointer) Size {
	sel := objc.RegisterName("drawAtPoint:forWidth:withFont:lineBreakMode:")
	ret := s_.ID.Send(sel, point, width, font, lineBreakMode)
	return Size(ret)
}
// Draws the string in a single line with the specified font and attributes, adjusting the font attributes as needed to render as much of the text as possible. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/drawAtPoint:forWidth:withFont:minFontSize:actualFontSize:lineBreakMode:baselineAdjustment:
func (s_ String) DrawAtPointForWidthWithFontMinFontSizeActualFontSizeLineBreakModeBaselineAdjustment(point Point, width float64, font unsafe.Pointer, minFontSize float64, actualFontSize float64, lineBreakMode unsafe.Pointer, baselineAdjustment unsafe.Pointer) Size {
	sel := objc.RegisterName("drawAtPoint:forWidth:withFont:minFontSize:actualFontSize:lineBreakMode:baselineAdjustment:")
	ret := s_.ID.Send(sel, point, width, font, minFontSize, actualFontSize, lineBreakMode, baselineAdjustment)
	return Size(ret)
}
// Draws the string in a single line at the specified point in the current graphics context using the specified font. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/drawAtPoint:withFont:
func (s_ String) DrawAtPointWithFont(point Point, font unsafe.Pointer) Size {
	sel := objc.RegisterName("drawAtPoint:withFont:")
	ret := s_.ID.Send(sel, point, font)
	return Size(ret)
}
// Draws the string in the current graphics context using the specified bounding rectangle and font. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/drawInRect:withFont:
func (s_ String) DrawInRectWithFont(rect Rect, font unsafe.Pointer) Size {
	sel := objc.RegisterName("drawInRect:withFont:")
	ret := s_.ID.Send(sel, rect, font)
	return Size(ret)
}
// Draws the string in the current graphics context using the specified bounding rectangle, font, and attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/drawInRect:withFont:lineBreakMode:
func (s_ String) DrawInRectWithFontLineBreakMode(rect Rect, font unsafe.Pointer, lineBreakMode unsafe.Pointer) Size {
	sel := objc.RegisterName("drawInRect:withFont:lineBreakMode:")
	ret := s_.ID.Send(sel, rect, font, lineBreakMode)
	return Size(ret)
}
// Draws the string in the current graphics context using the specified bounding rectangle, font and attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/drawInRect:withFont:lineBreakMode:alignment:
func (s_ String) DrawInRectWithFontLineBreakModeAlignment(rect Rect, font unsafe.Pointer, lineBreakMode unsafe.Pointer, alignment unsafe.Pointer) Size {
	sel := objc.RegisterName("drawInRect:withFont:lineBreakMode:alignment:")
	ret := s_.ID.Send(sel, rect, font, lineBreakMode, alignment)
	return Size(ret)
}
// Enumerates all the lines in the string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/enumerateLines(_:)
func (s_ String) EnumerateLinesUsingBlock(block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateLinesUsingBlock:")
	s_.ID.Send(sel, block)
}
// Performs linguistic analysis on the specified string by enumerating the specific range of the string, providing the Block with the located tags. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/enumerateLinguisticTags(in:scheme:options:orthography:using:)
func (s_ String) EnumerateLinguisticTagsInRangeSchemeOptionsOrthographyUsingBlock(range_ Range, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateLinguisticTagsInRange:scheme:options:orthography:usingBlock:")
	s_.ID.Send(sel, range_, scheme, options, orthography, block)
}
// Enumerates the substrings of the specified type in the specified range of the string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/enumerateSubstrings(in:options:using:)
func (s_ String) EnumerateSubstringsInRangeOptionsUsingBlock(range_ Range, opts unsafe.Pointer, block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateSubstringsInRange:options:usingBlock:")
	s_.ID.Send(sel, range_, opts, block)
}
// Creates a string suitable for comparison by removing the specified character distinctions from a string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/folding(options:locale:)
func (s_ String) StringByFoldingWithOptionsLocale(options unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringByFoldingWithOptions:locale:")
	ret := s_.ID.Send(sel, options, locale)
	return unsafe.Pointer(ret)
}
// Gets a given range of characters as bytes in a specified encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getBytes(_:maxLength:usedLength:encoding:options:range:remaining:)
func (s_ String) GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint, usedBufferCount unsafe.Pointer, encoding unsafe.Pointer, options unsafe.Pointer, range_ Range, leftover unsafe.Pointer) bool {
	sel := objc.RegisterName("getBytes:maxLength:usedLength:encoding:options:range:remainingRange:")
	ret := s_.ID.Send(sel, buffer, maxBufferCount, usedBufferCount, encoding, options, range_, leftover)
	return ret != 0
}
// Invokes   with   as the maximum length, the receiver’s entire extent as the range, and   for the remaining range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getCString(_:)
func (s_ String) GetCString(bytes unsafe.Pointer) {
	sel := objc.RegisterName("getCString:")
	s_.ID.Send(sel, bytes)
}
// Invokes   with   as the maximum length in char-sized units, the receiver’s entire extent as the range, and   for the remaining range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getCString(_:maxLength:)
func (s_ String) GetCStringMaxLength(bytes unsafe.Pointer, maxLength uint) {
	sel := objc.RegisterName("getCString:maxLength:")
	s_.ID.Send(sel, bytes, maxLength)
}
// Converts the string to a given encoding and stores it in a buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getCString(_:maxLength:encoding:)
func (s_ String) GetCStringMaxLengthEncoding(buffer unsafe.Pointer, maxBufferCount uint, encoding unsafe.Pointer) bool {
	sel := objc.RegisterName("getCString:maxLength:encoding:")
	ret := s_.ID.Send(sel, buffer, maxBufferCount, encoding)
	return ret != 0
}
// Converts the receiver’s content to the default C-string encoding and stores them in a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getCString(_:maxLength:range:remaining:)
func (s_ String) GetCStringMaxLengthRangeRemainingRange(bytes unsafe.Pointer, maxLength uint, aRange Range, leftoverRange unsafe.Pointer) {
	sel := objc.RegisterName("getCString:maxLength:range:remainingRange:")
	s_.ID.Send(sel, bytes, maxLength, aRange, leftoverRange)
}
// Copies all characters from the receiver into a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getCharacters(_:)
func (s_ String) GetCharacters(buffer unsafe.Pointer) {
	sel := objc.RegisterName("getCharacters:")
	s_.ID.Send(sel, buffer)
}
// Copies characters from a given range in the receiver into a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getCharacters(_:range:)
func (s_ String) GetCharactersRange(buffer unsafe.Pointer, range_ Range) {
	sel := objc.RegisterName("getCharacters:range:")
	s_.ID.Send(sel, buffer, range_)
}
// Interprets the receiver as a system-independent path and fills a buffer with a C-string in a format and encoding suitable for use with file-system calls. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getFileSystemRepresentation(_:maxLength:)
func (s_ String) GetFileSystemRepresentationMaxLength(cname unsafe.Pointer, max uint) bool {
	sel := objc.RegisterName("getFileSystemRepresentation:maxLength:")
	ret := s_.ID.Send(sel, cname, max)
	return ret != 0
}
// Returns by reference the beginning of the first line and the end of the last line touched by the given range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getLineStart(_:end:contentsEnd:for:)
func (s_ String) GetLineStartEndContentsEndForRange(startPtr unsafe.Pointer, lineEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ Range) {
	sel := objc.RegisterName("getLineStart:end:contentsEnd:forRange:")
	s_.ID.Send(sel, startPtr, lineEndPtr, contentsEndPtr, range_)
}
// Returns by reference the beginning of the first paragraph and the end of the last paragraph touched by the given range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/getParagraphStart(_:end:contentsEnd:for:)
func (s_ String) GetParagraphStartEndContentsEndForRange(startPtr unsafe.Pointer, parEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ Range) {
	sel := objc.RegisterName("getParagraphStart:end:contentsEnd:forRange:")
	s_.ID.Send(sel, startPtr, parEndPtr, contentsEndPtr, range_)
}
// Returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/hasPrefix(_:)
func (s_ String) HasPrefix(str string) bool {
	sel := objc.RegisterName("hasPrefix:")
	ret := s_.ID.Send(sel, str)
	return ret != 0
}
// Returns a Boolean value that indicates whether a given string matches the ending characters of the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/hasSuffix(_:)
func (s_ String) HasSuffix(str string) bool {
	sel := objc.RegisterName("hasSuffix:")
	ret := s_.ID.Send(sel, str)
	return ret != 0
}
// Returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/isEqual(to:)
func (s_ String) IsEqualToString(aString string) bool {
	sel := objc.RegisterName("isEqualToString:")
	ret := s_.ID.Send(sel, aString)
	return ret != 0
}
// Returns the number of bytes required to store the receiver in a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/lengthOfBytes(using:)
func (s_ String) LengthOfBytesUsingEncoding(enc unsafe.Pointer) uint {
	sel := objc.RegisterName("lengthOfBytesUsingEncoding:")
	ret := s_.ID.Send(sel, enc)
	return uint(ret)
}
// Returns the range of characters representing the line or lines containing a given range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/lineRange(for:)
func (s_ String) LineRangeForRange(range_ Range) Range {
	sel := objc.RegisterName("lineRangeForRange:")
	ret := s_.ID.Send(sel, range_)
	return Range(ret)
}
// Returns an array of linguistic tags for the specified range and requested tags within the receiving string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/linguisticTags(in:scheme:options:orthography:tokenRanges:)
func (s_ String) LinguisticTagsInRangeSchemeOptionsOrthographyTokenRanges(range_ Range, scheme unsafe.Pointer, options unsafe.Pointer, orthography unsafe.Pointer, tokenRanges unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("linguisticTagsInRange:scheme:options:orthography:tokenRanges:")
	ret := s_.ID.Send(sel, range_, scheme, options, orthography, tokenRanges)
	return unsafe.Pointer(ret)
}
// Compares the string with a given string using a case-insensitive, localized, comparison. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedCaseInsensitiveCompare(_:)
func (s_ String) LocalizedCaseInsensitiveCompare(string string) unsafe.Pointer {
	sel := objc.RegisterName("localizedCaseInsensitiveCompare:")
	ret := s_.ID.Send(sel, string)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedCaseInsensitiveContains(_:)
func (s_ String) LocalizedCaseInsensitiveContainsString(str string) bool {
	sel := objc.RegisterName("localizedCaseInsensitiveContainsString:")
	ret := s_.ID.Send(sel, str)
	return ret != 0
}
// Compares the string and a given string using a localized comparison. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedCompare(_:)
func (s_ String) LocalizedCompare(string string) unsafe.Pointer {
	sel := objc.RegisterName("localizedCompare:")
	ret := s_.ID.Send(sel, string)
	return unsafe.Pointer(ret)
}
// Compares strings as sorted by the Finder. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedStandardCompare(_:)
func (s_ String) LocalizedStandardCompare(string string) unsafe.Pointer {
	sel := objc.RegisterName("localizedStandardCompare:")
	ret := s_.ID.Send(sel, string)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedStandardContains(_:)
func (s_ String) LocalizedStandardContainsString(str string) bool {
	sel := objc.RegisterName("localizedStandardContainsString:")
	ret := s_.ID.Send(sel, str)
	return ret != 0
}
// Finds and returns the range of the first occurrence of a given string within the string by performing a case and diacritic insensitive, locale-aware search. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/localizedStandardRange(of:)
func (s_ String) LocalizedStandardRangeOfString(str string) Range {
	sel := objc.RegisterName("localizedStandardRangeOfString:")
	ret := s_.ID.Send(sel, str)
	return Range(ret)
}
// Returns a representation of the receiver as a C string in the default C-string encoding, possibly losing information in converting to that encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/lossyCString()
func (s_ String) LossyCString() unsafe.Pointer {
	sel := objc.RegisterName("lossyCString")
	ret := s_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns a version of the string with all letters converted to lowercase, taking into account the specified locale. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/lowercased(with:)
func (s_ String) LowercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("lowercaseStringWithLocale:")
	ret := s_.ID.Send(sel, locale)
	return unsafe.Pointer(ret)
}
// Returns the maximum number of bytes needed to store the receiver in a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/maximumLengthOfBytes(using:)
func (s_ String) MaximumLengthOfBytesUsingEncoding(enc unsafe.Pointer) uint {
	sel := objc.RegisterName("maximumLengthOfBytesUsingEncoding:")
	ret := s_.ID.Send(sel, enc)
	return uint(ret)
}
// Returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/padding(toLength:withPad:startingAt:)
func (s_ String) StringByPaddingToLengthWithStringStartingAtIndex(newLength uint, padString string, padIndex uint) unsafe.Pointer {
	sel := objc.RegisterName("stringByPaddingToLength:withString:startingAtIndex:")
	ret := s_.ID.Send(sel, newLength, padString, padIndex)
	return unsafe.Pointer(ret)
}
// Returns the range of characters representing the paragraph or paragraphs containing a given range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/paragraphRange(for:)
func (s_ String) ParagraphRangeForRange(range_ Range) Range {
	sel := objc.RegisterName("paragraphRangeForRange:")
	ret := s_.ID.Send(sel, range_)
	return Range(ret)
}
// Parses the receiver as a text representation of a property list, returning an  ,  ,  , or   object, according to the topmost element. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/propertyList()
func (s_ String) PropertyList() objc.ID {
	sel := objc.RegisterName("propertyList")
	ret := s_.ID.Send(sel)
	return ret
}
// Returns a dictionary object initialized with the keys and values found in the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/propertyListFromStringsFileFormat()
func (s_ String) PropertyListFromStringsFileFormat() unsafe.Pointer {
	sel := objc.RegisterName("propertyListFromStringsFileFormat")
	ret := s_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Finds and returns the range of the first occurrence of a given string within the string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/range(of:)
func (s_ String) RangeOfString(searchString string) Range {
	sel := objc.RegisterName("rangeOfString:")
	ret := s_.ID.Send(sel, searchString)
	return Range(ret)
}
// Finds and returns the range of the first occurrence of a given string within the string, subject to given options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/range(of:options:)
func (s_ String) RangeOfStringOptions(searchString string, mask unsafe.Pointer) Range {
	sel := objc.RegisterName("rangeOfString:options:")
	ret := s_.ID.Send(sel, searchString, mask)
	return Range(ret)
}
// Finds and returns the range of the first occurrence of a given string, within the given range of the string, subject to given options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/range(of:options:range:)
func (s_ String) RangeOfStringOptionsRange(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch Range) Range {
	sel := objc.RegisterName("rangeOfString:options:range:")
	ret := s_.ID.Send(sel, searchString, mask, rangeOfReceiverToSearch)
	return Range(ret)
}
// Finds and returns the range of the first occurrence of a given string within a given range of the string, subject to given options, using the specified locale, if any. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/range(of:options:range:locale:)
func (s_ String) RangeOfStringOptionsRangeLocale(searchString string, mask unsafe.Pointer, rangeOfReceiverToSearch Range, locale unsafe.Pointer) Range {
	sel := objc.RegisterName("rangeOfString:options:range:locale:")
	ret := s_.ID.Send(sel, searchString, mask, rangeOfReceiverToSearch, locale)
	return Range(ret)
}
// Finds and returns the range in the string of the first character from a given character set. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/rangeOfCharacter(from:)
func (s_ String) RangeOfCharacterFromSet(searchSet unsafe.Pointer) Range {
	sel := objc.RegisterName("rangeOfCharacterFromSet:")
	ret := s_.ID.Send(sel, searchSet)
	return Range(ret)
}
// Finds and returns the range in the string of the first character, using given options, from a given character set. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/rangeOfCharacter(from:options:)
func (s_ String) RangeOfCharacterFromSetOptions(searchSet unsafe.Pointer, mask unsafe.Pointer) Range {
	sel := objc.RegisterName("rangeOfCharacterFromSet:options:")
	ret := s_.ID.Send(sel, searchSet, mask)
	return Range(ret)
}
// Finds and returns the range in the string of the first character from a given character set found in a given range with given options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/rangeOfCharacter(from:options:range:)
func (s_ String) RangeOfCharacterFromSetOptionsRange(searchSet unsafe.Pointer, mask unsafe.Pointer, rangeOfReceiverToSearch Range) Range {
	sel := objc.RegisterName("rangeOfCharacterFromSet:options:range:")
	ret := s_.ID.Send(sel, searchSet, mask, rangeOfReceiverToSearch)
	return Range(ret)
}
// Returns the range in the receiver of the composed character sequence located at a given index. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/rangeOfComposedCharacterSequence(at:)
func (s_ String) RangeOfComposedCharacterSequenceAtIndex(index uint) Range {
	sel := objc.RegisterName("rangeOfComposedCharacterSequenceAtIndex:")
	ret := s_.ID.Send(sel, index)
	return Range(ret)
}
// Returns the range in the string of the composed character sequences for a given range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/rangeOfComposedCharacterSequences(for:)
func (s_ String) RangeOfComposedCharacterSequencesForRange(range_ Range) Range {
	sel := objc.RegisterName("rangeOfComposedCharacterSequencesForRange:")
	ret := s_.ID.Send(sel, range_)
	return Range(ret)
}
// Returns a new string in which the characters in a specified range of the receiver are replaced by a given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/replacingCharacters(in:with:)
func (s_ String) StringByReplacingCharactersInRangeWithString(range_ Range, replacement string) unsafe.Pointer {
	sel := objc.RegisterName("stringByReplacingCharactersInRange:withString:")
	ret := s_.ID.Send(sel, range_, replacement)
	return unsafe.Pointer(ret)
}
// Returns a new string in which all occurrences of a target string in the receiver are replaced by another given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/replacingOccurrences(of:with:)
func (s_ String) StringByReplacingOccurrencesOfStringWithString(target string, replacement string) unsafe.Pointer {
	sel := objc.RegisterName("stringByReplacingOccurrencesOfString:withString:")
	ret := s_.ID.Send(sel, target, replacement)
	return unsafe.Pointer(ret)
}
// Returns a new string in which all occurrences of a target string in a specified range of the receiver are replaced by another given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/replacingOccurrences(of:with:options:range:)
func (s_ String) StringByReplacingOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options unsafe.Pointer, searchRange Range) unsafe.Pointer {
	sel := objc.RegisterName("stringByReplacingOccurrencesOfString:withString:options:range:")
	ret := s_.ID.Send(sel, target, replacement, options, searchRange)
	return unsafe.Pointer(ret)
}
// Returns a new string made by replacing in the receiver all percent escapes with the matching characters as determined by a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/replacingPercentEscapes(using:)
func (s_ String) StringByReplacingPercentEscapesUsingEncoding(enc unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringByReplacingPercentEscapesUsingEncoding:")
	ret := s_.ID.Send(sel, enc)
	return unsafe.Pointer(ret)
}
// Returns the bounding box size the receiver occupies when drawn with the given attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/size(withAttributes:)
func (s_ String) SizeWithAttributes(attrs unsafe.Pointer) Size {
	sel := objc.RegisterName("sizeWithAttributes:")
	ret := s_.ID.Send(sel, attrs)
	return Size(ret)
}
// Returns the size of the string if it were to be rendered with the specified font on a single line. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/sizeWithFont:
func (s_ String) SizeWithFont(font unsafe.Pointer) Size {
	sel := objc.RegisterName("sizeWithFont:")
	ret := s_.ID.Send(sel, font)
	return Size(ret)
}
// Returns the size of the string if it were rendered and constrained to the specified size. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/sizeWithFont:constrainedToSize:
func (s_ String) SizeWithFontConstrainedToSize(font unsafe.Pointer, size Size) Size {
	sel := objc.RegisterName("sizeWithFont:constrainedToSize:")
	ret := s_.ID.Send(sel, font, size)
	return Size(ret)
}
// Returns the size of the string if it were rendered with the specified constraints. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/sizeWithFont:constrainedToSize:lineBreakMode:
func (s_ String) SizeWithFontConstrainedToSizeLineBreakMode(font unsafe.Pointer, size Size, lineBreakMode unsafe.Pointer) Size {
	sel := objc.RegisterName("sizeWithFont:constrainedToSize:lineBreakMode:")
	ret := s_.ID.Send(sel, font, size, lineBreakMode)
	return Size(ret)
}
// Returns the size of the string if it were to be rendered with the specified font and line attributes on a single line. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/sizeWithFont:forWidth:lineBreakMode:
func (s_ String) SizeWithFontForWidthLineBreakMode(font unsafe.Pointer, width float64, lineBreakMode unsafe.Pointer) Size {
	sel := objc.RegisterName("sizeWithFont:forWidth:lineBreakMode:")
	ret := s_.ID.Send(sel, font, width, lineBreakMode)
	return Size(ret)
}
// Returns the size of the string if it were rendered with the specified constraints, including a variable font size, on a single line. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/sizeWithFont:minFontSize:actualFontSize:forWidth:lineBreakMode:
func (s_ String) SizeWithFontMinFontSizeActualFontSizeForWidthLineBreakMode(font unsafe.Pointer, minFontSize float64, actualFontSize float64, width float64, lineBreakMode unsafe.Pointer) Size {
	sel := objc.RegisterName("sizeWithFont:minFontSize:actualFontSize:forWidth:lineBreakMode:")
	ret := s_.ID.Send(sel, font, minFontSize, actualFontSize, width, lineBreakMode)
	return Size(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/sr_sensorForDeletionRecordsFromSensor()
func (s_ String) Sr_sensorForDeletionRecordsFromSensor() unsafe.Pointer {
	sel := objc.RegisterName("sr_sensorForDeletionRecordsFromSensor")
	ret := s_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns a string made by appending to the receiver a string constructed from a given format string and the following arguments. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/stringByAppendingFormat:
func (s_ String) StringByAppendingFormat(format string) unsafe.Pointer {
	sel := objc.RegisterName("stringByAppendingFormat:")
	ret := s_.ID.Send(sel, format)
	return unsafe.Pointer(ret)
}
// Returns an array of strings made by separately appending to the receiver each string in a given array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/strings(byAppendingPaths:)
func (s_ String) StringsByAppendingPaths(paths unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringsByAppendingPaths:")
	ret := s_.ID.Send(sel, paths)
	return unsafe.Pointer(ret)
}
// Returns a new string containing the characters of the receiver from the one at a given index to the end. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/substring(from:)
func (s_ String) SubstringFromIndex(from uint) unsafe.Pointer {
	sel := objc.RegisterName("substringFromIndex:")
	ret := s_.ID.Send(sel, from)
	return unsafe.Pointer(ret)
}
// Returns a new string containing the characters of the receiver up to, but not including, the one at a given index. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/substring(to:)
func (s_ String) SubstringToIndex(to uint) unsafe.Pointer {
	sel := objc.RegisterName("substringToIndex:")
	ret := s_.ID.Send(sel, to)
	return unsafe.Pointer(ret)
}
// Returns a string object containing the characters of the receiver that lie within a given range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/substring(with:)
func (s_ String) SubstringWithRange(range_ Range) unsafe.Pointer {
	sel := objc.RegisterName("substringWithRange:")
	ret := s_.ID.Send(sel, range_)
	return unsafe.Pointer(ret)
}
// Returns a new string made by removing from both ends of the receiver characters contained in a given character set. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/trimmingCharacters(in:)
func (s_ String) StringByTrimmingCharactersInSet(set unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stringByTrimmingCharactersInSet:")
	ret := s_.ID.Send(sel, set)
	return unsafe.Pointer(ret)
}
// Returns a version of the string with all letters converted to uppercase, taking into account the specified locale. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/uppercased(with:)
func (s_ String) UppercaseStringWithLocale(locale unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("uppercaseStringWithLocale:")
	ret := s_.ID.Send(sel, locale)
	return unsafe.Pointer(ret)
}
// Returns a string variation suitable for the specified presentation width. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/variantFittingPresentationWidth(_:)
func (s_ String) VariantFittingPresentationWidth(width int) unsafe.Pointer {
	sel := objc.RegisterName("variantFittingPresentationWidth:")
	ret := s_.ID.Send(sel, width)
	return unsafe.Pointer(ret)
}
// Writes the contents of the receiver to the location specified by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/write(to:atomically:)
func (s_ String) WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool {
	sel := objc.RegisterName("writeToURL:atomically:")
	ret := s_.ID.Send(sel, url, atomically)
	return ret != 0
}
// Writes the contents of the receiver to the URL specified by   using the specified encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/write(to:atomically:encoding:)
func (s_ String) WriteToURLAtomicallyEncodingError(url unsafe.Pointer, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("writeToURL:atomically:encoding:error:")
	ret := s_.ID.Send(sel, url, useAuxiliaryFile, enc, error)
	return ret != 0
}
// Writes the contents of the receiver to the file specified by a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/write(toFile:atomically:)
func (s_ String) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	sel := objc.RegisterName("writeToFile:atomically:")
	ret := s_.ID.Send(sel, path, useAuxiliaryFile)
	return ret != 0
}
// Writes the contents of the receiver to a file at a given path using a given encoding. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSString/write(toFile:atomically:encoding:)
func (s_ String) WriteToFileAtomicallyEncodingError(path string, useAuxiliaryFile bool, enc unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("writeToFile:atomically:encoding:error:")
	ret := s_.ID.Send(sel, path, useAuxiliaryFile, enc, error)
	return ret != 0
}

