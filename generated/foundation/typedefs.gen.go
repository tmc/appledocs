// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Type aliases and typedefs
// FileAttributeKey - Keys in dictionaries used to get and set file attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileAttributeKey
// NSFileAttributeKey is a string typedef
type FileAttributeKey = string
// FileAttributeType - Values representing a file’s type attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileAttributeType
// NSFileAttributeType is a string typedef
type FileAttributeType = string
// FileProtectionType - Protection level values that can be associated with a file attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileProtectionType
// NSFileProtectionType is a string typedef
type FileProtectionType = string
// AttributedStringKey - The attributes you apply to ranges of characters in an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
// NSAttributedStringKey is a string typedef
type AttributedStringKey = string
// CalendarIdentifier - The supported calendar types.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier
// NSCalendarIdentifier is a string typedef
type CalendarIdentifier = string
// URLResourceKey - Keys that apply to file system URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLResourceKey
// NSURLResourceKey is a string typedef
type URLResourceKey = string
// ExceptionName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExceptionName
// NSExceptionName is a string typedef
type ExceptionName = string
// HashTableOptions - Components in a bit-field to specify the behavior of elements in an   object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTableOptions
type HashTableOptions uint
// LinguisticTag - A token, lexical class, name, lemma, language, or script returned by a linguistic tagger for natural language text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTag
// NSLinguisticTag is a string typedef
type LinguisticTag = string
// LinguisticTagScheme - Constants for the tag schemes specified when initializing a linguistic tagger.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagScheme
// NSLinguisticTagScheme is a string typedef
type LinguisticTagScheme = string
// LocaleKey - The keys used to access components of a locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/Key
// NSLocaleKey is a string typedef
type LocaleKey = string
// MapTableOptions - Constants used as components in a bitfield to specify the behavior of elements (keys and values) in an   object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTableOptions
type MapTableOptions uint
// SKIPPED: NSPoint - conflicts with existing enum "Point"
// PointArray - Type indicating a parameter is array of   structures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointArray
// NSPointArray has base type: NSPoint *
type PointArray uintptr
// PointPointer - Type indicating a parameter is a pointer to an   structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointPointer
// NSPointPointer has base type: NSPoint *
type PointPointer uintptr
// RangePointer - Type indicating a parameter is a pointer to an   structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangePointer
// NSRangePointer has base type: NSRange *
type RangePointer uintptr
// SKIPPED: NSRect - conflicts with existing enum "Rect"
// RectArray - Type indicating a parameter is array of   structures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectArray
// NSRectArray has base type: NSRect *
type RectArray uintptr
// RectPointer - Type indicating a parameter is a pointer to an   structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectPointer
// NSRectPointer has base type: NSRect *
type RectPointer uintptr
// SKIPPED: NSSize - conflicts with existing enum "Size"
// SizeArray - Type indicating a parameter is an array of   structures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSizeArray
// NSSizeArray has base type: NSSize *
type SizeArray uintptr
// SizePointer - Type indicating parameter is a pointer to an   structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSizePointer
// NSSizePointer has base type: NSSize *
type SizePointer uintptr
// TextCheckingKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey
// NSTextCheckingKey is a string typedef
type TextCheckingKey = string
// TextCheckingTypes - Defines the types of checking that are available. These values can be combined using the C-bitwise OR operator. The system supports its own internal types, and the user can extend those types by subclassing   and adding their own custom types.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingTypes
// NSTextCheckingTypes has base type: uint64_t
type TextCheckingTypes uintptr
// URLBookmarkFileCreationOptions - Options used when creating file bookmark data
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkFileCreationOptions
type URLBookmarkFileCreationOptions uint
// Integer - Describes an integer.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSInteger
// NSInteger has base type: long
type Integer uintptr
// DataRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFData
// CFDataRef has base type: const struct __CFData *
type DataRef uintptr
// URLRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURL
// CFURLRef has base type: const struct __CFURL *
type URLRef uintptr
// UUIDRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUID
// CFUUIDRef has base type: const struct __CFUUID *
type UUIDRef uintptr
// URLFileProtectionType - Protection-level values for a URL resource key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLFileProtection
// NSURLFileProtectionType is a string typedef
type URLFileProtectionType = string
// URLFileResourceType - Possible values for the type of file resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLFileResourceType
// NSURLFileResourceType is a string typedef
type URLFileResourceType = string
// URLThumbnailDictionaryItem - Possible keys for the   dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLThumbnailDictionaryItem
// NSURLThumbnailDictionaryItem is a string typedef
type URLThumbnailDictionaryItem = string
// URLUbiquitousItemDownloadingStatus - Values that describe the iCloud storage state of a file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLUbiquitousItemDownloadingStatus
// NSURLUbiquitousItemDownloadingStatus is a string typedef
type URLUbiquitousItemDownloadingStatus = string
// URLUbiquitousSharedItemPermissions - The key for the permissions of a shared item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLUbiquitousSharedItemPermissions
// NSURLUbiquitousSharedItemPermissions is a string typedef
type URLUbiquitousSharedItemPermissions = string
// URLUbiquitousSharedItemRole - The key for the role of a shared item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLUbiquitousSharedItemRole
// NSURLUbiquitousSharedItemRole is a string typedef
type URLUbiquitousSharedItemRole = string
// TimeInterval - A number of seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/TimeInterval
// NSTimeInterval has base type: double
type TimeInterval uintptr
// unichar - Type for UTF-16 code units.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/unichar
// unichar has base type: unsigned short
type unichar uintptr

