// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

// Enum types and constants
// OSLogEntryStoreCategory - A classification of how the entry was to be stored and rotated at the point when it was created.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum
type OSLogEntryStoreCategory uint

const (
// OSLogEntryStoreCategoryLongTerm1 - The entry was tagged with a hint indicating the system should try to preserve it for approximately 1 day.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum/longTerm1
OSLogEntryStoreCategoryLongTerm1 OSLogEntryStoreCategory = 0
// OSLogEntryStoreCategoryLongTerm14 - The entry was tagged with a hint indicating the system should try to preserve it for approximately 14 days.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum/longTerm14
OSLogEntryStoreCategoryLongTerm14 OSLogEntryStoreCategory = 0
// OSLogEntryStoreCategoryLongTerm3 - The entry was tagged with a hint indicating the system should try to preserve it for approximately 3 days.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum/longTerm3
OSLogEntryStoreCategoryLongTerm3 OSLogEntryStoreCategory = 0
// OSLogEntryStoreCategoryLongTerm30 - The entry was tagged with a hint indicating the system should try to preserve it for approximately 30 days.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum/longTerm30
OSLogEntryStoreCategoryLongTerm30 OSLogEntryStoreCategory = 0
// OSLogEntryStoreCategoryLongTerm7 - The entry was tagged with a hint indicating the system should try to preserve it for approximately 7 days.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum/longTerm7
OSLogEntryStoreCategoryLongTerm7 OSLogEntryStoreCategory = 0
// OSLogEntryStoreCategoryLongTermAuto - The entry was tagged with a hint indicating the system should try to preserve it based on the amount of space available.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum/longTermAuto
OSLogEntryStoreCategoryLongTermAuto OSLogEntryStoreCategory = 0
// OSLogEntryStoreCategoryMetadata - This entry was generated as information about the other entries or about the sequence of entries as a whole.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum/metadata
OSLogEntryStoreCategoryMetadata OSLogEntryStoreCategory = 0
// OSLogEntryStoreCategoryShortTerm - This entry was not intended to be long-lived, and was captured in the ring buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum/shortTerm
OSLogEntryStoreCategoryShortTerm OSLogEntryStoreCategory = 0
// OSLogEntryStoreCategoryUndefined - This entry’s purpose is unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum/undefined
OSLogEntryStoreCategoryUndefined OSLogEntryStoreCategory = 0
)

// OSLogEntryLogLevel - The log level at which the entry was generated.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntryLog/Level-swift.enum
type OSLogEntryLogLevel uint

const (
// OSLogEntryLogLevelDebug - A log level that captures diagnostic information.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntryLog/Level-swift.enum/debug
OSLogEntryLogLevelDebug OSLogEntryLogLevel = 0
// OSLogEntryLogLevelError - The log level that captures errors.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntryLog/Level-swift.enum/error
OSLogEntryLogLevelError OSLogEntryLogLevel = 0
// OSLogEntryLogLevelFault - The log level that captures fault information.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntryLog/Level-swift.enum/fault
OSLogEntryLogLevelFault OSLogEntryLogLevel = 0
// OSLogEntryLogLevelInfo - The log level that captures additional information.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntryLog/Level-swift.enum/info
OSLogEntryLogLevelInfo OSLogEntryLogLevel = 0
// OSLogEntryLogLevelNotice - The log level that captures notifications.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntryLog/Level-swift.enum/notice
OSLogEntryLogLevelNotice OSLogEntryLogLevel = 0
)

// OSLogEntrySignpostType - The available signpost types.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/SignpostType-swift.enum
type OSLogEntrySignpostType uint

const (
// OSLogEntrySignpostTypeEvent - The signpost marks an event.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/SignpostType-swift.enum/event
OSLogEntrySignpostTypeEvent OSLogEntrySignpostType = 0
// OSLogEntrySignpostTypeIntervalBegin - The signpost marks the start of a time interval.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/SignpostType-swift.enum/intervalBegin
OSLogEntrySignpostTypeIntervalBegin OSLogEntrySignpostType = 0
// OSLogEntrySignpostTypeIntervalEnd - The signpost marks the end of a time interval.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/SignpostType-swift.enum/intervalEnd
OSLogEntrySignpostTypeIntervalEnd OSLogEntrySignpostType = 0
// OSLogEntrySignpostTypeUndefined - The signpost does not have a type.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/SignpostType-swift.enum/undefined
OSLogEntrySignpostTypeUndefined OSLogEntrySignpostType = 0
)

// OSLogEnumeratorOptions - Option to control the direction of the iteration.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEnumerator/Options
type OSLogEnumeratorOptions uint

const (
// OSLogEnumeratorReverse - Tells the framework to iterate backwards.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEnumerator/Options/reverse
OSLogEnumeratorReverse OSLogEnumeratorOptions = 0
)

// OSLogMessageComponentArgumentCategory - The data type corresponding to the argument provided in a message payload.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/ArgumentCategory-swift.enum
type OSLogMessageComponentArgumentCategory uint

const (
// OSLogMessageComponentArgumentCategoryData - The argument is an   object.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/ArgumentCategory-swift.enum/data
OSLogMessageComponentArgumentCategoryData OSLogMessageComponentArgumentCategory = 0
// OSLogMessageComponentArgumentCategoryDouble - The argument is a double.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/ArgumentCategory-swift.enum/double
OSLogMessageComponentArgumentCategoryDouble OSLogMessageComponentArgumentCategory = 0
// OSLogMessageComponentArgumentCategoryInt64 - The argument is a 64-bit signed integer.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/ArgumentCategory-swift.enum/int64
OSLogMessageComponentArgumentCategoryInt64 OSLogMessageComponentArgumentCategory = 0
// OSLogMessageComponentArgumentCategoryString - The argument is a string.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/ArgumentCategory-swift.enum/string
OSLogMessageComponentArgumentCategoryString OSLogMessageComponentArgumentCategory = 0
// OSLogMessageComponentArgumentCategoryUInt64 - The argument is a 64-bit unsigned integer.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/ArgumentCategory-swift.enum/uInt64
OSLogMessageComponentArgumentCategoryUInt64 OSLogMessageComponentArgumentCategory = 0
// OSLogMessageComponentArgumentCategoryUndefined - The argument’s type is not defined.
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/ArgumentCategory-swift.enum/undefined
OSLogMessageComponentArgumentCategoryUndefined OSLogMessageComponentArgumentCategory = 0
)

// OSLogStoreScope enum type
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/Scope
type OSLogStoreScope uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/Scope/currentProcessIdentifier
OSLogStoreCurrentProcessIdentifier OSLogStoreScope = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/Scope/system
OSLogStoreSystem OSLogStoreScope = 0
)


