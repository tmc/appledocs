// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Type aliases and typedefs
// DistributedNotificationCenterType - This constant specifies the notification center type.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/CenterType
// NSDistributedNotificationCenterType is a string typedef
type DistributedNotificationCenterType = string
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
// AppleEventManagerSuspensionID - Identifies an Apple event whose handling has been suspended. Can be used to resume handling of the Apple event.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/SuspensionID
// NSAppleEventManagerSuspensionID has base type: const struct __NSAppleEventManagerSuspension *
type AppleEventManagerSuspensionID uintptr
// AttributedStringKey - The attributes you apply to ranges of characters in an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
// NSAttributedStringKey is a string typedef
type AttributedStringKey = string
// AttributedStringFormattingContextKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingContextKey
// NSAttributedStringFormattingContextKey is a string typedef
type AttributedStringFormattingContextKey = string
// CalendarIdentifier - The supported calendar types.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier
// NSCalendarIdentifier is a string typedef
type CalendarIdentifier = string
// ErrorUserInfoKey - These keys may exist in the user info dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/UserInfoKey
// NSErrorUserInfoKey is a string typedef
type ErrorUserInfoKey = string
// ErrorDomain type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSErrorDomain
// NSErrorDomain is a string typedef
type ErrorDomain = string
// FileProviderServiceName - The name used to identify a File Provider service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileProviderServiceName
// NSFileProviderServiceName is a string typedef
type FileProviderServiceName = string
// HashTableOptions - Components in a bit-field to specify the behavior of elements in an   object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTableOptions
type HashTableOptions uint
// KeyValueChangeKey - The keys that can appear in the change dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueChangeKey
// NSKeyValueChangeKey is a string typedef
type KeyValueChangeKey = string
// KeyValueOperator - These constants define the available array operators. See   for more information.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator
// NSKeyValueOperator is a string typedef
type KeyValueOperator = string
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
// SKIPPED: NSRect - conflicts with existing enum "Rect"
// SKIPPED: NSSize - conflicts with existing enum "Size"
// TimeInterval - A number of seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/TimeInterval
// NSTimeInterval has base type: double
type TimeInterval uintptr
// URLBookmarkFileCreationOptions - Options used when creating file bookmark data
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkFileCreationOptions
type URLBookmarkFileCreationOptions uint
// UserActivityPersistentIdentifier - The type that defines a persistent identifier value for a user activity.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivityPersistentIdentifier
// NSUserActivityPersistentIdentifier is a string typedef
type UserActivityPersistentIdentifier = string
// ProgressFileOperationKind - The kind of file operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/FileOperationKind-swift.struct
// NSProgressFileOperationKind is a string typedef
type ProgressFileOperationKind = string
// SKIPPED: (^ - invalid Go identifier "(^"
// Original type: void (^(^)(NSProgress *))(void) NSProgressPublishingHandler
// ProgressKind - An object that represents the kind of progress.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProgressKind
// NSProgressKind is a string typedef
type ProgressKind = string
// ProgressUserInfoKey - Keys for the user info dictionary that affect the autogenerated localized additional description string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey
// NSProgressUserInfoKey is a string typedef
type ProgressUserInfoKey = string
// RunLoopMode - Modes that a run loop operates in.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/Mode
// NSRunLoopMode is a string typedef
type RunLoopMode = string
// SocketNativeHandle - Type for the platform-specific native socket handle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketNativeHandle
type SocketNativeHandle int32
// StringEncodingDetectionOptionsKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey
// NSStringEncodingDetectionOptionsKey is a string typedef
type StringEncodingDetectionOptionsKey = string
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
// URLResourceKey - Keys that apply to file system URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLResourceKey
// NSURLResourceKey is a string typedef
type URLResourceKey = string
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
// UndoManagerUserInfoKey - An extensible namespace for undo and redo user info keys.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/UserInfoKey
// NSUndoManagerUserInfoKey is a string typedef
type UndoManagerUserInfoKey = string
// unichar - Type for UTF-16 code units.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/unichar
// unichar has base type: unsigned short
type unichar uintptr

