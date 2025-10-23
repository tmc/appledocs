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
// ErrorDomain type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSErrorDomain
// NSErrorDomain is a string typedef
type ErrorDomain = string
// UncaughtExceptionHandler type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUncaughtExceptionHandler
// NSUncaughtExceptionHandler has base type: void (NSException *)
type UncaughtExceptionHandler uintptr
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
// NotificationName - A structure that defines the name of a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct
// NSNotificationName is a string typedef
type NotificationName = string
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
// AppKitVersion - Constants for determining which version of AppKit is available.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppKitVersion
// NSAppKitVersion has base type: double
type AppKitVersion uintptr
// PasteboardName - Constants that represent the standard pasteboard names.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/Name-swift.struct
// NSPasteboardName is a string typedef
type PasteboardName = string
// PasteboardType - The supported pasteboard types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType
// NSPasteboardType is a string typedef
type PasteboardType = string
// PasteboardReadingOptionKey - Options for reading pasteboard data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptionKey
// NSPasteboardReadingOptionKey is a string typedef
type PasteboardReadingOptionKey = string
// WorkspaceDesktopImageOptionKey - Keys that indicate how to display a new desktop image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/DesktopImageOptionKey
// NSWorkspaceDesktopImageOptionKey is a string typedef
type WorkspaceDesktopImageOptionKey = string
// Integer - Describes an integer.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSInteger
// NSInteger has base type: long
type Integer uintptr
// objc_exception_handler type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_handler
// objc_exception_handler has base type: void (*)(id, void *)
type objc_exception_handler uintptr
// objc_exception_matcher type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_matcher
// objc_exception_matcher has base type: int (*)(Class, id)
type objc_exception_matcher uintptr
// objc_exception_preprocessor type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_preprocessor
// objc_exception_preprocessor has base type: id (*)(id)
type objc_exception_preprocessor uintptr
// objc_func_loadImage type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_func_loadImage
// objc_func_loadImage has base type: void (*)(const struct mach_header *)
type objc_func_loadImage uintptr
// objc_hook_getClass type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_hook_getClass
// objc_hook_getClass has base type: _Bool (*)(const char *, Class *)
type objc_hook_getClass uintptr
// objc_hook_getImageName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_hook_getImageName
// objc_hook_getImageName has base type: _Bool (*)(Class, const char **)
type objc_hook_getImageName uintptr
// objc_hook_lazyClassNamer type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_hook_lazyClassNamer
// objc_hook_lazyClassNamer has base type: const char *(*)(Class)
type objc_hook_lazyClassNamer uintptr
// objc_objectptr_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_objectptr_t
// objc_objectptr_t has base type: const void *
type objc_objectptr_t uintptr
// objc_uncaught_exception_handler type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_uncaught_exception_handler
// objc_uncaught_exception_handler has base type: void (*)(id)
type objc_uncaught_exception_handler uintptr
// objc_zone_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_zone_t
// objc_zone_t has base type: struct _malloc_zone_t *
type objc_zone_t uintptr
// AllocatorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocator
// CFAllocatorRef has base type: const struct __CFAllocator *
type AllocatorRef uintptr
// ArrayRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArray
// CFArrayRef has base type: const struct __CFArray *
type ArrayRef uintptr
// AttributedStringRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedString
// CFAttributedStringRef has base type: const struct __CFAttributedString *
type AttributedStringRef uintptr
// BagRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBag
// CFBagRef has base type: const struct __CFBag *
type BagRef uintptr
// BinaryHeapRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeap
// CFBinaryHeapRef has base type: struct __CFBinaryHeap *
type BinaryHeapRef uintptr
// BitVectorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVector
// CFBitVectorRef has base type: const struct __CFBitVector *
type BitVectorRef uintptr
// BooleanRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBoolean
// CFBooleanRef has base type: const struct __CFBoolean *
type BooleanRef uintptr
// BundleRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundle
// CFBundleRef has base type: struct __CFBundle *
type BundleRef uintptr
// CalendarRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendar
// CFCalendarRef has base type: struct __CFCalendar *
type CalendarRef uintptr
// CharacterSetRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSet
// CFCharacterSetRef has base type: const struct __CFCharacterSet *
type CharacterSetRef uintptr
// DataRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFData
// CFDataRef has base type: const struct __CFData *
type DataRef uintptr
// DateRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDate
// CFDateRef has base type: const struct __CFDate *
type DateRef uintptr
// DateFormatterRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatter
// CFDateFormatterRef has base type: struct __CFDateFormatter *
type DateFormatterRef uintptr
// DictionaryRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionary
// CFDictionaryRef has base type: const struct __CFDictionary *
type DictionaryRef uintptr
// ErrorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFError
// CFErrorRef has base type: struct __CFError *
type ErrorRef uintptr
// FileDescriptorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptor
// CFFileDescriptorRef has base type: struct __CFFileDescriptor *
type FileDescriptorRef uintptr
// LocaleRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocale
// CFLocaleRef has base type: const struct __CFLocale *
type LocaleRef uintptr
// MutableAttributedStringRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableAttributedString
// CFMutableAttributedStringRef has base type: struct __CFAttributedString *
type MutableAttributedStringRef uintptr
// MutableCharacterSetRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableCharacterSet
// CFMutableCharacterSetRef has base type: struct __CFCharacterSet *
type MutableCharacterSetRef uintptr
// MutableDataRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableData
// CFMutableDataRef has base type: struct __CFData *
type MutableDataRef uintptr
// NullRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNull
// CFNullRef has base type: const struct __CFNull *
type NullRef uintptr
// OptionFlags - A bitfield used for passing special allocation and other requests into Core Foundation functions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFOptionFlags
// CFOptionFlags has base type: unsigned long
type OptionFlags uintptr
// StringRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFString
// CFStringRef has base type: const struct __CFString *
type StringRef uintptr
// TimeZoneRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZone
// CFTimeZoneRef has base type: const struct __CFTimeZone *
type TimeZoneRef uintptr
// URLRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURL
// CFURLRef has base type: const struct __CFURL *
type URLRef uintptr
// URLBookmarkFileCreationOptions - Type for bookmark file creation options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkFileCreationOptions
// CFURLBookmarkFileCreationOptions has base type: CFOptionFlags
type URLBookmarkFileCreationOptions uintptr
// UUIDRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUID
// CFUUIDRef has base type: const struct __CFUUID *
type UUIDRef uintptr
// WriteStreamRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStream
// CFWriteStreamRef has base type: struct __CFWriteStream *
type WriteStreamRef uintptr
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
// ApplicationOpenExternalURLOptionsKey - Options for opening a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplication/OpenExternalURLOptionsKey
// UIApplicationOpenExternalURLOptionsKey is a string typedef
type ApplicationOpenExternalURLOptionsKey = string
// BackgroundTaskIdentifier - A unique token that identifies a request to run in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIBackgroundTaskIdentifier
type BackgroundTaskIdentifier uint
// ContentSizeCategory - Constants that indicate the preferred size of your content.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIContentSizeCategory
// UIContentSizeCategory is a string typedef
type ContentSizeCategory = string
// DocumentCreationIntent - An app intent that creates new documents for your app.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/CreationIntent
// UIDocumentCreationIntent is a string typedef
type DocumentCreationIntent = string

