// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

// Enum types and constants
// CFCalendarUnit - CFCalendarUnit constants are used to specify calendrical units, such as day or month, in various calendar calculations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit
type CalendarUnit uint

const (
	// kCFCalendarUnitMonth - Specifies the month unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/month
	kCFCalendarUnitMonth CalendarUnit = 0
)

// CFCharacterSetPredefinedSet - Defines a predefined character set.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet
type CharacterSetPredefinedSet uint

const (
	// kCFCharacterSetDecomposable - Canonically decomposable character set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/decomposable
	kCFCharacterSetDecomposable CharacterSetPredefinedSet = 0
	// kCFCharacterSetNewline - Newline character set ( ,  ,  , and  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/newline
	kCFCharacterSetNewline CharacterSetPredefinedSet = 0
)

// CFComparisonResult - Constants returned by comparison functions, indicating whether a value is equal to, less than, or greater than another value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult
type ComparisonResult uint

const (
	// kCFCompareGreaterThan - Returned by a comparison function if the first value is greater than the second value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult/compareGreaterThan
	kCFCompareGreaterThan ComparisonResult = 0
)

// CFDataSearchFlags - A 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags
type DataSearchFlags uint

// CFDateFormatterStyle - Data type for predefined date and time format styles.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle
type DateFormatterStyle uint

const (
	// kCFDateFormatterFullStyle - Specifies a full style with complete details, such as “Tuesday, April 12, 1952 AD” or “3:30:42pm PST”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/fullStyle
	kCFDateFormatterFullStyle DateFormatterStyle = 0
	// kCFDateFormatterLongStyle - Specifies a long style, typically with full text, such as “November 23, 1937” or “3:30:32pm”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/longStyle
	kCFDateFormatterLongStyle DateFormatterStyle = 0
	// kCFDateFormatterShortStyle - Specifies a short style, typically numeric only, such as “11/23/37” or “3:30pm”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/shortStyle
	kCFDateFormatterShortStyle DateFormatterStyle = 0
)

// CFFileSecurityClearOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions
type FileSecurityClearOptions uint

const (
	// kCFFileSecurityClearAccessControlList - Clear the access control list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/accessControlList
	kCFFileSecurityClearAccessControlList FileSecurityClearOptions = 0
	// kCFFileSecurityClearMode - Clear the file’s mode (POSIX permissions).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/mode
	kCFFileSecurityClearMode FileSecurityClearOptions = 0
	// kCFFileSecurityClearOwner - Clear the (POSIX) owner ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/owner
	kCFFileSecurityClearOwner FileSecurityClearOptions = 0
	// kCFFileSecurityClearOwnerUUID - Clear the owner UUID (for the access control list).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/ownerUUID
	kCFFileSecurityClearOwnerUUID FileSecurityClearOptions = 0
)

// CFGregorianUnitFlags - These option flags are used as a mask to indicate a specific set of fields in the CFGregorianDate or CFGregorianUnits structures.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags
type GregorianUnitFlags uint

// CFISO8601DateFormatOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions
type ISO8601DateFormatOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withColonSeparatorInTime
	kCFISO8601DateFormatWithColonSeparatorInTime ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withFractionalSeconds
	kCFISO8601DateFormatWithFractionalSeconds ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withWeekOfYear
	kCFISO8601DateFormatWithWeekOfYear ISO8601DateFormatOptions = 0
)

// CFLocaleLanguageDirection - These constants describe the text direction for a language. They are returned by the functions 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection
type LocaleLanguageDirection uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/unknown
	kCFLocaleLanguageDirectionUnknown LocaleLanguageDirection = 0
)

// CFNotificationSuspensionBehavior - Suspension flags that indicate how distributed notifications should be handled when the receiving application is in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior
type NotificationSuspensionBehavior uint

// CFNumberFormatterOptionFlags - Type for constants specifying how numbers should be parsed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterOptionFlags
type NumberFormatterOptionFlags uint

const (
	// kCFNumberFormatterParseIntegersOnly - Specifies that only integers should be parsed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterOptionFlags/parseIntegersOnly
	kCFNumberFormatterParseIntegersOnly NumberFormatterOptionFlags = 0
)

// CFNumberFormatterPadPosition - Type for constants specifying how numbers should be padded.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition
type NumberFormatterPadPosition uint

const (
	// kCFNumberFormatterPadAfterPrefix - Specifies the number of padding characters after the prefix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition/afterPrefix
	kCFNumberFormatterPadAfterPrefix NumberFormatterPadPosition = 0
	// kCFNumberFormatterPadAfterSuffix - Specifies the number of padding characters after the suffix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition/afterSuffix
	kCFNumberFormatterPadAfterSuffix NumberFormatterPadPosition = 0
	// kCFNumberFormatterPadBeforePrefix - Specifies the number of padding characters before the prefix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition/beforePrefix
	kCFNumberFormatterPadBeforePrefix NumberFormatterPadPosition = 0
	// kCFNumberFormatterPadBeforeSuffix - Specifies the number of padding characters before the suffix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition/beforeSuffix
	kCFNumberFormatterPadBeforeSuffix NumberFormatterPadPosition = 0
)

// CFNumberFormatterRoundingMode - These constants are used to specify how numbers should be rounded.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode
type NumberFormatterRoundingMode uint

const (
	// kCFNumberFormatterRoundCeiling - Round towards positive infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundCeiling
	kCFNumberFormatterRoundCeiling NumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundDown - Round towards zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundDown
	kCFNumberFormatterRoundDown NumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundFloor - Round towards negative infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundFloor
	kCFNumberFormatterRoundFloor NumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundHalfDown - Round towards the nearest integer, or towards zero if equidistant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundHalfDown
	kCFNumberFormatterRoundHalfDown NumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundHalfEven - Round towards the nearest integer, or towards an even number if equidistant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundHalfEven
	kCFNumberFormatterRoundHalfEven NumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundHalfUp - Round towards the nearest integer, or away from zero if equidistant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundHalfUp
	kCFNumberFormatterRoundHalfUp NumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundUp - Round away from zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundUp
	kCFNumberFormatterRoundUp NumberFormatterRoundingMode = 0
)

// CFNumberFormatterStyle - Type for constants specifying a formatter style.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle
type NumberFormatterStyle uint

const (
	// kCFNumberFormatterNoStyle - Specifies no style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/noStyle
	kCFNumberFormatterNoStyle NumberFormatterStyle = 0
	// kCFNumberFormatterPercentStyle - Specifies a percent style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/percentStyle
	kCFNumberFormatterPercentStyle NumberFormatterStyle = 0
)

// CFNumberType - Flags used by CFNumber to indicate the data type of a value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType
type NumberType uint

const (
	// kCFNumberCFIndexType - CFIndex value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/cfIndexType
	kCFNumberCFIndexType NumberType = 0
	// kCFNumberCGFloatType -  value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/cgFloatType
	kCFNumberCGFloatType NumberType = 0
	// kCFNumberCharType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/charType
	kCFNumberCharType NumberType = 0
	// kCFNumberDoubleType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/doubleType
	kCFNumberDoubleType NumberType = 0
	// kCFNumberFloat32Type - Thirty-two-bit real. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/float32Type
	kCFNumberFloat32Type NumberType = 0
	// kCFNumberFloat64Type - Sixty-four-bit real. The   data type is defined in   and conforms to the 64-bit IEEE 754 standard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/float64Type
	kCFNumberFloat64Type NumberType = 0
	// kCFNumberFloatType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/floatType
	kCFNumberFloatType NumberType = 0
	// kCFNumberIntType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/intType
	kCFNumberIntType NumberType = 0
	// kCFNumberLongLongType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/longLongType
	kCFNumberLongLongType NumberType = 0
	// kCFNumberLongType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/longType
	kCFNumberLongType NumberType = 0
	// kCFNumberMaxType - Same as  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/maxType
	kCFNumberMaxType NumberType = 0
	// kCFNumberNSIntegerType -  value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/nsIntegerType
	kCFNumberNSIntegerType NumberType = 0
	// kCFNumberSInt16Type - Sixteen-bit, signed integer. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/sInt16Type
	kCFNumberSInt16Type NumberType = 0
	// kCFNumberSInt32Type - Thirty-two-bit, signed integer. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/sInt32Type
	kCFNumberSInt32Type NumberType = 0
	// kCFNumberSInt64Type - Sixty-four-bit, signed integer. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/sInt64Type
	kCFNumberSInt64Type NumberType = 0
	// kCFNumberSInt8Type - Eight-bit, signed integer. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/sInt8Type
	kCFNumberSInt8Type NumberType = 0
	// kCFNumberShortType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/shortType
	kCFNumberShortType NumberType = 0
)

// CFPropertyListFormat - Specifies the format of a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListFormat
type PropertyListFormat uint

const (
	// kCFPropertyListBinaryFormat_v1_0 - Binary format version 1.0.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListFormat/binaryFormat_v1_0
	kCFPropertyListBinaryFormat_v1_0 PropertyListFormat = 0
	// kCFPropertyListOpenStepFormat - OpenStep format (use of this format is discouraged).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListFormat/openStepFormat
	kCFPropertyListOpenStepFormat PropertyListFormat = 0
)

// CFPropertyListMutabilityOptions - Type for flags that determine the degree of mutability of newly created property lists.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListMutabilityOptions
type PropertyListMutabilityOptions uint

const (
	// kCFPropertyListImmutable - Specifies that the property list should be immutable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListMutabilityOptions/kCFPropertyListImmutable
	kCFPropertyListImmutable PropertyListMutabilityOptions = 0
	// kCFPropertyListMutableContainers - Specifies that the property list should have mutable containers but immutable leaves.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListMutabilityOptions/mutableContainers
	kCFPropertyListMutableContainers PropertyListMutabilityOptions = 0
	// kCFPropertyListMutableContainersAndLeaves - Specifies that the property list should have mutable containers and mutable leaves.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListMutabilityOptions/mutableContainersAndLeaves
	kCFPropertyListMutableContainersAndLeaves PropertyListMutabilityOptions = 0
)

// CFRunLoopActivity - Run loop activity stages in which run loop observers can be scheduled.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity
type RunLoopActivity uint

const (
	// kCFRunLoopAfterWaiting - Inside the event processing loop after the run loop wakes up, but before processing the event that woke it up. This activity occurs only if the run loop did in fact go to sleep during the current loop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/afterWaiting
	kCFRunLoopAfterWaiting RunLoopActivity = 0
	// kCFRunLoopAllActivities - A combination of all the preceding stages.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/allActivities
	kCFRunLoopAllActivities RunLoopActivity = 0
	// kCFRunLoopBeforeSources - Inside the event processing loop before any sources are processed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/beforeSources
	kCFRunLoopBeforeSources RunLoopActivity = 0
	// kCFRunLoopBeforeTimers - Inside the event processing loop before any timers are processed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/beforeTimers
	kCFRunLoopBeforeTimers RunLoopActivity = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/beforeWaiting
	kCFRunLoopBeforeWaiting RunLoopActivity = 0
	// kCFRunLoopEntry - The entrance of the run loop, before entering the event processing loop. This activity occurs once for each call to   and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/entry
	kCFRunLoopEntry RunLoopActivity = 0
	// kCFRunLoopExit - The exit of the run loop, after exiting the event processing loop. This activity occurs once for each call to   and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/exit
	kCFRunLoopExit RunLoopActivity = 0
)

// CFRunLoopRunResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult
type RunLoopRunResult uint

const (
	// kCFRunLoopRunFinished - The running run loop mode has no sources or timers to process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult/finished
	kCFRunLoopRunFinished RunLoopRunResult = 0
	// kCFRunLoopRunHandledSource - A source has been processed. This value is returned only if the run loop was told to run only until a source was processed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult/handledSource
	kCFRunLoopRunHandledSource RunLoopRunResult = 0
	// kCFRunLoopRunStopped -  was called on the run loop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult/stopped
	kCFRunLoopRunStopped RunLoopRunResult = 0
	// kCFRunLoopRunTimedOut - The specified time interval for running the run loop has passed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult/timedOut
	kCFRunLoopRunTimedOut RunLoopRunResult = 0
)

// CFSocketCallBackType - Types of socket activity that can cause the callback function of a CFSocket object to be called.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType
type SocketCallBackType uint

const (
	// kCFSocketAcceptCallBack - New connections will be automatically accepted and the callback is called with the data argument being a pointer to a   of the child socket. This callback is usable only with listening sockets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/acceptCallBack
	kCFSocketAcceptCallBack SocketCallBackType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/connectCallBack
	kCFSocketConnectCallBack SocketCallBackType = 0
	// kCFSocketNoCallBack - No callback should be made for any activity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/kCFSocketNoCallBack
	kCFSocketNoCallBack SocketCallBackType = 0
	// kCFSocketReadCallBack - The callback is called when data is available to be read or a new connection is waiting to be accepted. The data is not automatically read; the callback must read the data itself.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/readCallBack
	kCFSocketReadCallBack SocketCallBackType = 0
	// kCFSocketWriteCallBack - The callback is called when the socket is writable. This callback type may be useful when large amounts of data are being sent rapidly over the socket and you want a notification when there is space in the kernel buffers for more data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/writeCallBack
	kCFSocketWriteCallBack SocketCallBackType = 0
)

// CFSocketError - Error codes for many CFSocket functions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketError
type SocketError uint

// CFStreamErrorDomain - Defines constants for values returned in the domain field of the 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamErrorDomain
type StreamErrorDomain uint

const (
	// kCFStreamErrorDomainPOSIX - The error code is an error code defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamErrorDomain/POSIX
	kCFStreamErrorDomainPOSIX StreamErrorDomain = 0
	// kCFStreamErrorDomainCustom - The error code is a custom error code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamErrorDomain/custom
	kCFStreamErrorDomainCustom StreamErrorDomain = 0
	// kCFStreamErrorDomainMacOSStatus - The error is an OSStatus value defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamErrorDomain/macOSStatus
	kCFStreamErrorDomainMacOSStatus StreamErrorDomain = 0
)

// CFStreamEventType - Defines constants for stream-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType
type StreamEventType uint

const (
	// kCFStreamEventCanAcceptBytes - The stream can accept bytes for writing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/canAcceptBytes
	kCFStreamEventCanAcceptBytes StreamEventType = 0
	// kCFStreamEventEndEncountered - The end of the stream has been reached.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/endEncountered
	kCFStreamEventEndEncountered StreamEventType = 0
	// kCFStreamEventErrorOccurred - An error has occurred on the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/errorOccurred
	kCFStreamEventErrorOccurred StreamEventType = 0
	// kCFStreamEventHasBytesAvailable - The stream has bytes to be read.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/hasBytesAvailable
	kCFStreamEventHasBytesAvailable StreamEventType = 0
	// kCFStreamEventNone - No event has occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/kCFStreamEventNone
	kCFStreamEventNone StreamEventType = 0
	// kCFStreamEventOpenCompleted - The open has completed successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/openCompleted
	kCFStreamEventOpenCompleted StreamEventType = 0
)

// CFStreamStatus - Constants that describe the status of a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus
type StreamStatus uint

const (
	// kCFStreamStatusAtEnd - There is no more data to read, or no more data can be written.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/atEnd
	kCFStreamStatusAtEnd StreamStatus = 0
	// kCFStreamStatusClosed - The stream is closed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/closed
	kCFStreamStatusClosed StreamStatus = 0
	// kCFStreamStatusError - An error occurred on the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/error
	kCFStreamStatusError StreamStatus = 0
	// kCFStreamStatusNotOpen - The stream is not open for reading or writing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/notOpen
	kCFStreamStatusNotOpen StreamStatus = 0
	// kCFStreamStatusOpen - The stream is open.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/open
	kCFStreamStatusOpen StreamStatus = 0
	// kCFStreamStatusOpening - The stream is being opened for reading or for writing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/opening
	kCFStreamStatusOpening StreamStatus = 0
	// kCFStreamStatusReading - The stream is being read from.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/reading
	kCFStreamStatusReading StreamStatus = 0
	// kCFStreamStatusWriting - The stream is being written to.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/writing
	kCFStreamStatusWriting StreamStatus = 0
)

// CFStringBuiltInEncodings - Encodings that are built-in on all platforms on which macOS runs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings
type StringBuiltInEncodings uint

// CFStringCompareFlags - A 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags
type StringCompareFlags uint

const (
	// kCFCompareAnchored - Performs searching only on characters at the beginning or end of the range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareAnchored
	kCFCompareAnchored StringCompareFlags = 0
	// kCFCompareBackwards - Specifies that the comparison should start at the last elements of the entities being compared (for example, strings or arrays).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareBackwards
	kCFCompareBackwards StringCompareFlags = 0
	// kCFCompareCaseInsensitive - Specifies that the comparison should ignore differences in case between alphabetical characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareCaseInsensitive
	kCFCompareCaseInsensitive StringCompareFlags = 0
	// kCFCompareDiacriticInsensitive - Specifies that the comparison should ignore diacritic markers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareDiacriticInsensitive
	kCFCompareDiacriticInsensitive StringCompareFlags = 0
	// kCFCompareForcedOrdering - Specifies that the comparison is forced to return either   or   if the strings are equivalent but not strictly equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareForcedOrdering
	kCFCompareForcedOrdering StringCompareFlags = 0
	// kCFCompareLocalized - Specifies that the comparison should take into account differences related to locale, such as the thousands separator character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareLocalized
	kCFCompareLocalized StringCompareFlags = 0
	// kCFCompareNonliteral - Specifies that loose equivalence is acceptable, especially as pertains to diacritical marks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareNonliteral
	kCFCompareNonliteral StringCompareFlags = 0
	// kCFCompareNumerically - Specifies that represented numeric values should be used as the basis for comparison and not the actual character values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareNumerically
	kCFCompareNumerically StringCompareFlags = 0
	// kCFCompareWidthInsensitive - Specifies that the comparison should ignore width differences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareWidthInsensitive
	kCFCompareWidthInsensitive StringCompareFlags = 0
)

// CFStringEncodings - Index type for constants used to specify external string encodings.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings
type StringEncodings uint

// CFStringNormalizationForm - Unicode normalization forms as described in Unicode Technical Report #15.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm
type StringNormalizationForm uint

const (
	// kCFStringNormalizationFormC - Canonical decomposition followed by canonical composition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm/C
	kCFStringNormalizationFormC StringNormalizationForm = 0
	// kCFStringNormalizationFormD - Canonical decomposition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm/D
	kCFStringNormalizationFormD StringNormalizationForm = 0
	// kCFStringNormalizationFormKC - Compatibility decomposition followed by canonical composition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm/KC
	kCFStringNormalizationFormKC StringNormalizationForm = 0
	// kCFStringNormalizationFormKD - Compatibility decomposition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm/KD
	kCFStringNormalizationFormKD StringNormalizationForm = 0
)

// CFStringTokenizerTokenType - Token types returned by 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType
type StringTokenizerTokenType uint

const (
	// kCFStringTokenizerTokenHasDerivedSubTokensMask - Compound token which may contain derived subtokens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/hasDerivedSubTokensMask
	kCFStringTokenizerTokenHasDerivedSubTokensMask StringTokenizerTokenType = 0
	// kCFStringTokenizerTokenHasHasNumbersMask - Appears to contain a number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/hasHasNumbersMask
	kCFStringTokenizerTokenHasHasNumbersMask StringTokenizerTokenType = 0
	// kCFStringTokenizerTokenHasNonLettersMask - Contains punctuation, symbols, and so on.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/hasNonLettersMask
	kCFStringTokenizerTokenHasNonLettersMask StringTokenizerTokenType = 0
	// kCFStringTokenizerTokenHasSubTokensMask - Compound token which may contain subtokens but with no derived subtokens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/hasSubTokensMask
	kCFStringTokenizerTokenHasSubTokensMask StringTokenizerTokenType = 0
	// kCFStringTokenizerTokenIsCJWordMask - Contains kana and/or ideographs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/isCJWordMask
	kCFStringTokenizerTokenIsCJWordMask StringTokenizerTokenType = 0
	// kCFStringTokenizerTokenNone - Has no token.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/kCFStringTokenizerTokenNone
	kCFStringTokenizerTokenNone StringTokenizerTokenType = 0
	// kCFStringTokenizerTokenNormal - Has a normal token.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/normal
	kCFStringTokenizerTokenNormal StringTokenizerTokenType = 0
)

// CFTimeZoneNameStyle - Index type for constants used to specify styles of time zone names.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle
type TimeZoneNameStyle uint

const (
	// kCFTimeZoneNameStyleDaylightSaving - Specifies the daylight saving name style; for example, “Central Daylight Time” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/daylightSaving
	kCFTimeZoneNameStyleDaylightSaving TimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleGeneric - Specifies the generic name style, which does not distinguish between daylight saving and standard time; for example, “Central Time” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/generic
	kCFTimeZoneNameStyleGeneric TimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleShortDaylightSaving - Specifies the short daylight saving name style; for example, “CDT” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/shortDaylightSaving
	kCFTimeZoneNameStyleShortDaylightSaving TimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleShortGeneric - Specifies the short generic name style, which does not distinguish between daylight saving and standard time; for example, “CT” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/shortGeneric
	kCFTimeZoneNameStyleShortGeneric TimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleShortStandard - Specifies the short standard name style; for example, “CST” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/shortStandard
	kCFTimeZoneNameStyleShortStandard TimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleStandard - Specifies the standard name style; for example, “Central Standard Time” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/standard
	kCFTimeZoneNameStyleStandard TimeZoneNameStyle = 0
)

// CFURLBookmarkCreationOptions - Type for bookmark data creation options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions
type URLBookmarkCreationOptions uint

const (
	// kCFURLBookmarkCreationPreferFileIDResolutionMask - Specifies that an alias created with the bookmark data prefers resolving with its embedded file ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/preferFileIDResolutionMask
	kCFURLBookmarkCreationPreferFileIDResolutionMask URLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess - When combined with the   option, specifies that you want to create a security-scoped bookmark that, when resolved, provides a security-scoped URL allowing read-only access to a file-system resource; for use in an app that adopts App Sandbox.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/securityScopeAllowOnlyReadAccess
	kCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess URLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationWithSecurityScope - Specifies that you want to create a security-scoped bookmark that, when resolved, provides a security-scoped URL allowing read/write access to a file-system resource; for use in an app that adopts App Sandbox.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/withSecurityScope
	kCFURLBookmarkCreationWithSecurityScope URLBookmarkCreationOptions = 0
)

// CFURLBookmarkResolutionOptions - Type for bookmark data resolution options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions
type URLBookmarkResolutionOptions uint

// CFURLComponentType - The types of components in a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType
type URLComponentType uint

// CFURLEnumeratorOptions - Options for controlling enumerator behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions
type URLEnumeratorOptions uint

const (
	// kCFURLEnumeratorDescendRecursively - The enumerator recurses into each subdirectory enumerated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/descendRecursively
	kCFURLEnumeratorDescendRecursively URLEnumeratorOptions = 0
	// kCFURLEnumeratorGenerateFileReferenceURLs - The enumerator generates file reference URLs instead of file path URLs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/generateFileReferenceURLs
	kCFURLEnumeratorGenerateFileReferenceURLs URLEnumeratorOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/generateRelativePathURLs
	kCFURLEnumeratorGenerateRelativePathURLs URLEnumeratorOptions = 0
	// kCFURLEnumeratorIncludeDirectoriesPostOrder - If provided along with the   option, the recursive enumerator returns a directory’s URL after returning the URLs of the directory’s descendents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/includeDirectoriesPostOrder
	kCFURLEnumeratorIncludeDirectoriesPostOrder URLEnumeratorOptions = 0
	// kCFURLEnumeratorIncludeDirectoriesPreOrder - If provided along with the   option, the recursive enumerator returns a directory’s URL before returning the URLs of the directory’s descendents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/includeDirectoriesPreOrder
	kCFURLEnumeratorIncludeDirectoriesPreOrder URLEnumeratorOptions = 0
	// kCFURLEnumeratorDefaultBehavior - The enumerator performs its default behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/kCFURLEnumeratorDefaultBehavior
	kCFURLEnumeratorDefaultBehavior URLEnumeratorOptions = 0
	// kCFURLEnumeratorSkipInvisibles - The enumerator skips “hidden” or “invisible” objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/skipInvisibles
	kCFURLEnumeratorSkipInvisibles URLEnumeratorOptions = 0
	// kCFURLEnumeratorSkipPackageContents - The enumerator skips package directory contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/skipPackageContents
	kCFURLEnumeratorSkipPackageContents URLEnumeratorOptions = 0
)

// CFURLEnumeratorResult - Result codes from the 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult
type URLEnumeratorResult uint

const (
	// kCFURLEnumeratorDirectoryPostOrderSuccess - The recursive post-order enumerator returned the URL for a directory after having returned the URLs for all of the directory’s descendents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult/directoryPostOrderSuccess
	kCFURLEnumeratorDirectoryPostOrderSuccess URLEnumeratorResult = 0
	// kCFURLEnumeratorEnd - The enumeration is complete.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult/end
	kCFURLEnumeratorEnd URLEnumeratorResult = 0
	// kCFURLEnumeratorError - An error occurred during enumeration. The   parameter of the function is populated with error information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult/error
	kCFURLEnumeratorError URLEnumeratorResult = 0
	// kCFURLEnumeratorSuccess - The enumerator was advanced successfully and returned a valid URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult/success
	kCFURLEnumeratorSuccess URLEnumeratorResult = 0
)

// CFURLError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError
type URLError uint

// CFURLPathStyle - Options you can use to determine how CFURL functions parse a file system path name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle
type URLPathStyle uint

// CFXMLEntityTypeCode - The entity type identification codes that the parser uses to describe XML entities.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode
type XMLEntityTypeCode uint

// CFXMLNodeTypeCode - The various XML data type identification codes that the parser uses to describe XML structures.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode
type XMLNodeTypeCode uint

const (
	// kCFXMLNodeTypeAttribute - Currently not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/attribute
	kCFXMLNodeTypeAttribute XMLNodeTypeCode = 0
	// kCFXMLNodeTypeAttributeListDeclaration - Indicates an attribute list declaration where the data string is the tag name and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/attributeListDeclaration
	kCFXMLNodeTypeAttributeListDeclaration XMLNodeTypeCode = 0
	// kCFXMLNodeTypeCDATASection - Indicates a CDATA section where the data string is the text of the CDATA and the additional information is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/cdataSection
	kCFXMLNodeTypeCDATASection XMLNodeTypeCode = 0
	// kCFXMLNodeTypeComment - Indicates a comment section where the data string is the text of the comment and the additional information is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/comment
	kCFXMLNodeTypeComment XMLNodeTypeCode = 0
	// kCFXMLNodeTypeDocument - Indicates a document where the data string is   and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/document
	kCFXMLNodeTypeDocument XMLNodeTypeCode = 0
	// kCFXMLNodeTypeDocumentFragment - Currently not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/documentFragment
	kCFXMLNodeTypeDocumentFragment XMLNodeTypeCode = 0
	// kCFXMLNodeTypeDocumentType - Indicates a document type where the data string is the name given to the top-level element and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/documentType
	kCFXMLNodeTypeDocumentType XMLNodeTypeCode = 0
	// kCFXMLNodeTypeElement - Indicates an element where the data string is the name of the tag and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/element
	kCFXMLNodeTypeElement XMLNodeTypeCode = 0
	// kCFXMLNodeTypeElementTypeDeclaration - Indicates an element type declaration where the data string is the tag name and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/elementTypeDeclaration
	kCFXMLNodeTypeElementTypeDeclaration XMLNodeTypeCode = 0
	// kCFXMLNodeTypeEntity - Indicates an entity where the data string is the name of the entity and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/entity
	kCFXMLNodeTypeEntity XMLNodeTypeCode = 0
	// kCFXMLNodeTypeEntityReference - Indicates an entity reference where the data string is the name of the referenced entity and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/entityReference
	kCFXMLNodeTypeEntityReference XMLNodeTypeCode = 0
	// kCFXMLNodeTypeNotation - Indicates a notation where the data string is the notation name and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/notation
	kCFXMLNodeTypeNotation XMLNodeTypeCode = 0
	// kCFXMLNodeTypeProcessingInstruction - Indicates a processing instruction where the data string is the name of the target and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/processingInstruction
	kCFXMLNodeTypeProcessingInstruction XMLNodeTypeCode = 0
	// kCFXMLNodeTypeText - Indicates a text section where the data string is the text’s contents and the additional information is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/text
	kCFXMLNodeTypeText XMLNodeTypeCode = 0
	// kCFXMLNodeTypeWhitespace - Indicates white space where the data string is the text of the white space and the additional information is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/whitespace
	kCFXMLNodeTypeWhitespace XMLNodeTypeCode = 0
)

// CFXMLParserOptions - Options you can use to control the parser’s treatment of an XML document.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions
type XMLParserOptions uint

// CFXMLParserStatusCode - The various status and error flags that can be returned by the parser.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode
type XMLParserStatusCode uint

const (
	// kCFXMLErrorElementlessDocument - Indicates a document containing no elements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorElementlessDocument
	kCFXMLErrorElementlessDocument XMLParserStatusCode = 0
	// kCFXMLErrorEncodingConversionFailure - Indicates an encoding conversion error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorEncodingConversionFailure
	kCFXMLErrorEncodingConversionFailure XMLParserStatusCode = 0
	// kCFXMLErrorMalformedCDSect - Indicates a malformed CDATA section.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedCDSect
	kCFXMLErrorMalformedCDSect XMLParserStatusCode = 0
	// kCFXMLErrorMalformedCharacterReference - Indicates a malformed character reference.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedCharacterReference
	kCFXMLErrorMalformedCharacterReference XMLParserStatusCode = 0
	// kCFXMLErrorMalformedCloseTag - Indicates a malformed close tag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedCloseTag
	kCFXMLErrorMalformedCloseTag XMLParserStatusCode = 0
	// kCFXMLErrorMalformedComment - Indicates a malformed comment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedComment
	kCFXMLErrorMalformedComment XMLParserStatusCode = 0
	// kCFXMLErrorMalformedDTD - Indicates a malformed DTD.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedDTD
	kCFXMLErrorMalformedDTD XMLParserStatusCode = 0
	// kCFXMLErrorMalformedDocument - Indicates a malformed document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedDocument
	kCFXMLErrorMalformedDocument XMLParserStatusCode = 0
	// kCFXMLErrorMalformedName - Indicates a malformed name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedName
	kCFXMLErrorMalformedName XMLParserStatusCode = 0
	// kCFXMLErrorMalformedParsedCharacterData - Indicates malformed character data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedParsedCharacterData
	kCFXMLErrorMalformedParsedCharacterData XMLParserStatusCode = 0
	// kCFXMLErrorMalformedProcessingInstruction - Indicates a malformed processing instruction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedProcessingInstruction
	kCFXMLErrorMalformedProcessingInstruction XMLParserStatusCode = 0
	// kCFXMLErrorMalformedStartTag - Indicates a malformed start tag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedStartTag
	kCFXMLErrorMalformedStartTag XMLParserStatusCode = 0
	// kCFXMLErrorNoData - Indicates a no data error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorNoData
	kCFXMLErrorNoData XMLParserStatusCode = 0
	// kCFXMLErrorUnexpectedEOF - Indicates an unexpected EOF occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorUnexpectedEOF
	kCFXMLErrorUnexpectedEOF XMLParserStatusCode = 0
	// kCFXMLErrorUnknownEncoding - Indicates an unknown encoding error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorUnknownEncoding
	kCFXMLErrorUnknownEncoding XMLParserStatusCode = 0
	// kCFXMLStatusParseSuccessful - Indicates the parser was successful.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/kCFXMLStatusParseSuccessful
	kCFXMLStatusParseSuccessful XMLParserStatusCode = 0
	// kCFXMLStatusParseInProgress - Indicates the parser is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/statusParseInProgress
	kCFXMLStatusParseInProgress XMLParserStatusCode = 0
	// kCFXMLStatusParseNotBegun - Indicates the parser has not begun.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/statusParseNotBegun
	kCFXMLStatusParseNotBegun XMLParserStatusCode = 0
)

// CGRectEdge enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge
type RectEdge uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge/maxXEdge
	RectMaxXEdge RectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge/maxYEdge
	RectMaxYEdge RectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge/minXEdge
	RectMinXEdge RectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge/minYEdge
	RectMinYEdge RectEdge = 0
)


