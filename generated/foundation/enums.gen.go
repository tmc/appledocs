// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
// NSAlignmentOptions - Values representing alignment operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions
type AlignmentOptions uint

// NSComparisonResult - Constants that indicate sort order.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ComparisonResult
type ComparisonResult uint

const (
	// OrderedAscending - The left operand is smaller than the right operand.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ComparisonResult/orderedAscending
	OrderedAscending ComparisonResult = 0
	// OrderedDescending - The left operand is greater than the right operand.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ComparisonResult/orderedDescending
	OrderedDescending ComparisonResult = 0
	// OrderedSame - The two operands are equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ComparisonResult/orderedSame
	OrderedSame ComparisonResult = 0
)

// NSDateFormatterStyle - The following constants specify predefined format styles for dates and times.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style
type DateFormatterStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/full
	DateFormatterFullStyle DateFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/long
	DateFormatterLongStyle DateFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/medium
	DateFormatterMediumStyle DateFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/none
	DateFormatterNoStyle DateFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/short
	DateFormatterShortStyle DateFormatterStyle = 0
)

// NSNotificationSuspensionBehavior - These constants specify the types of notification delivery suspension behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior
type NotificationSuspensionBehavior uint

const (
	// NotificationSuspensionBehaviorCoalesce - The server only queues the last notification of the specified name and object; earlier notifications are dropped. In cover methods for which suspension behavior is not an explicit argument,   is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior/coalesce
	NotificationSuspensionBehaviorCoalesce NotificationSuspensionBehavior = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior/deliverImmediately
	NotificationSuspensionBehaviorDeliverImmediately NotificationSuspensionBehavior = 4
	// NotificationSuspensionBehaviorDrop - The server doesn’t queue any notifications with this name and object until the notification center resumes notification delivery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior/drop
	NotificationSuspensionBehaviorDrop NotificationSuspensionBehavior = 1
	// NotificationSuspensionBehaviorHold - The server holds all matching notifications until the queue has been filled (queue size determined by the server), at which point the server may flush queued notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior/hold
	NotificationSuspensionBehaviorHold NotificationSuspensionBehavior = 3
)

// NSEnergyFormatterUnit - The units supported by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit
type EnergyFormatterUnit uint

const (
	EnergyFormatterUnitJoule EnergyFormatterUnit = 11
	EnergyFormatterUnitKilojoule EnergyFormatterUnit = 14
	NumberFormatterBehaviorDefault EnergyFormatterUnit = 0
	NumberFormatterBehavior10_0 EnergyFormatterUnit = 1000
	NumberFormatterBehavior10_4 EnergyFormatterUnit = 1040
)

// NSDirectoryEnumerationOptions - Options for enumerating the contents of directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions
type DirectoryEnumerationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/includesDirectoriesPostOrder
	DirectoryEnumerationIncludesDirectoriesPostOrder DirectoryEnumerationOptions = 2
	// DirectoryEnumerationSkipsHiddenFiles - An option to skip hidden files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsHiddenFiles
	DirectoryEnumerationSkipsHiddenFiles DirectoryEnumerationOptions = 1
	// DirectoryEnumerationSkipsPackageDescendants - An option to treat packages like files and not descend into their contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsPackageDescendants
	DirectoryEnumerationSkipsPackageDescendants DirectoryEnumerationOptions = 1
	// DirectoryEnumerationSkipsSubdirectoryDescendants - An option to perform a shallow enumeration that doesn’t descend into directories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsSubdirectoryDescendants
	DirectoryEnumerationSkipsSubdirectoryDescendants DirectoryEnumerationOptions = 1
)

// NSFileManagerItemReplacementOptions - Options for specifying the behavior of file replacement operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions
type FileManagerItemReplacementOptions uint

const (
	// FileManagerItemReplacementUsingNewMetadataOnly - Only metadata from the new item is used, and metadata from the original item isn’t preserved (default).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions/usingNewMetadataOnly
	FileManagerItemReplacementUsingNewMetadataOnly FileManagerItemReplacementOptions = 1
	// FileManagerItemReplacementWithoutDeletingBackupItem - The backup item remains in place after a successful replacement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions/withoutDeletingBackupItem
	FileManagerItemReplacementWithoutDeletingBackupItem FileManagerItemReplacementOptions = 1
)

// NSSearchPathDirectory - The location of significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
type SearchPathDirectory uint

const (
	// ApplicationDirectory - Supported applications ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationDirectory
	ApplicationDirectory SearchPathDirectory = 1
	// ApplicationScriptsDirectory - The user scripts folder for the calling application ( .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationScriptsDirectory
	ApplicationScriptsDirectory SearchPathDirectory = 23
	// ApplicationSupportDirectory - Application support files ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationSupportDirectory
	ApplicationSupportDirectory SearchPathDirectory = 14
	// DesktopDirectory - The user’s desktop directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/desktopDirectory
	DesktopDirectory SearchPathDirectory = 12
	// DocumentDirectory - Document directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentDirectory
	DocumentDirectory SearchPathDirectory = 9
	// DocumentationDirectory - Documentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentationDirectory
	DocumentationDirectory SearchPathDirectory = 8
	// InputMethodsDirectory - Input Methods  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/inputMethodsDirectory
	InputMethodsDirectory SearchPathDirectory = 16
	// ItemReplacementDirectory - The constant used to create a temporary directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/itemReplacementDirectory
	ItemReplacementDirectory SearchPathDirectory = 24
)

// NSSearchPathDomainMask - Domain constants specifying base locations to use when you search for significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask
type SearchPathDomainMask uint

const (
	// AllDomainsMask - All domains.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/allDomainsMask
	AllDomainsMask SearchPathDomainMask = 0
	// UserDomainMask - The user’s home directory—the place to install user’s personal items ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/userDomainMask
	UserDomainMask SearchPathDomainMask = 1
)

// NSURLRelationship - Constants indicating the relationship between a directory and an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship
type URLRelationship uint

const (
	URLRelationshipContains URLRelationship = 0
	URLRelationshipSame URLRelationship = 1
	FileManagerUnmountAllPartitionsAndEjectDisk URLRelationship = 1
	FileManagerUnmountWithoutUI URLRelationship = 1
	FileManagerSupportedSyncControlsPauseSync URLRelationship = 1
	FileManagerSupportedSyncControlsFailUploadOnConflict URLRelationship = 1
	FileManagerResumeSyncBehaviorPreserveLocalChanges URLRelationship = 0
	FileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict URLRelationship = 1
	FileManagerResumeSyncBehaviorDropLocalChanges URLRelationship = 2
	FileManagerUploadConflictPolicyDefault URLRelationship = 0
	FileManagerUploadConflictPolicyFailOnConflict URLRelationship = 1
	PointerFunctionsStrongMemory URLRelationship = 2
	PointerFunctionsZeroingWeakMemory URLRelationship = 3
	PointerFunctionsOpaqueMemory URLRelationship = 4
	PointerFunctionsMallocMemory URLRelationship = 5
	PointerFunctionsMachVirtualMemory URLRelationship = 6
	PointerFunctionsWeakMemory URLRelationship = 7
	PointerFunctionsObjectPersonality URLRelationship = 8
	PointerFunctionsOpaquePersonality URLRelationship = 9
	PointerFunctionsObjectPointerPersonality URLRelationship = 10
	PointerFunctionsCStringPersonality URLRelationship = 11
	PointerFunctionsStructPersonality URLRelationship = 12
	PointerFunctionsIntegerPersonality URLRelationship = 13
	PointerFunctionsCopyIn URLRelationship = 14
)

// NSFileManagerUnmountOptions - Options that specify the behavior of an unmount operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions
type FileManagerUnmountOptions uint

const (
	FileManagerUnmountAllPartitionsAndEjectDisk FileManagerUnmountOptions = 1
	FileManagerUnmountWithoutUI FileManagerUnmountOptions = 1
	FileManagerSupportedSyncControlsPauseSync FileManagerUnmountOptions = 1
	FileManagerSupportedSyncControlsFailUploadOnConflict FileManagerUnmountOptions = 1
	FileManagerResumeSyncBehaviorPreserveLocalChanges FileManagerUnmountOptions = 0
	FileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict FileManagerUnmountOptions = 1
	FileManagerResumeSyncBehaviorDropLocalChanges FileManagerUnmountOptions = 2
	FileManagerUploadConflictPolicyDefault FileManagerUnmountOptions = 0
	FileManagerUploadConflictPolicyFailOnConflict FileManagerUnmountOptions = 1
	PointerFunctionsStrongMemory FileManagerUnmountOptions = 2
	PointerFunctionsZeroingWeakMemory FileManagerUnmountOptions = 3
	PointerFunctionsOpaqueMemory FileManagerUnmountOptions = 4
	PointerFunctionsMallocMemory FileManagerUnmountOptions = 5
	PointerFunctionsMachVirtualMemory FileManagerUnmountOptions = 6
	PointerFunctionsWeakMemory FileManagerUnmountOptions = 7
	PointerFunctionsObjectPersonality FileManagerUnmountOptions = 8
	PointerFunctionsOpaquePersonality FileManagerUnmountOptions = 9
	PointerFunctionsObjectPointerPersonality FileManagerUnmountOptions = 10
	PointerFunctionsCStringPersonality FileManagerUnmountOptions = 11
	PointerFunctionsStructPersonality FileManagerUnmountOptions = 12
	PointerFunctionsIntegerPersonality FileManagerUnmountOptions = 13
	PointerFunctionsCopyIn FileManagerUnmountOptions = 14
)

// NSVolumeEnumerationOptions - Options for enumerating mounted volumes with the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions
type VolumeEnumerationOptions uint

const (
	VolumeEnumerationSkipHiddenVolumes VolumeEnumerationOptions = 1
	VolumeEnumerationProduceFileReferenceURLs VolumeEnumerationOptions = 1
	DirectoryEnumerationSkipsSubdirectoryDescendants VolumeEnumerationOptions = 1
	DirectoryEnumerationSkipsPackageDescendants VolumeEnumerationOptions = 1
	DirectoryEnumerationSkipsHiddenFiles VolumeEnumerationOptions = 1
	DirectoryEnumerationIncludesDirectoriesPostOrder VolumeEnumerationOptions = 2
	DirectoryEnumerationProducesRelativePathURLs VolumeEnumerationOptions = 3
	FileManagerItemReplacementUsingNewMetadataOnly VolumeEnumerationOptions = 1
	FileManagerItemReplacementWithoutDeletingBackupItem VolumeEnumerationOptions = 1
	URLRelationshipContains VolumeEnumerationOptions = 2
	URLRelationshipSame VolumeEnumerationOptions = 3
	FileManagerUnmountAllPartitionsAndEjectDisk VolumeEnumerationOptions = 1
	FileManagerUnmountWithoutUI VolumeEnumerationOptions = 1
	FileManagerSupportedSyncControlsPauseSync VolumeEnumerationOptions = 1
	FileManagerSupportedSyncControlsFailUploadOnConflict VolumeEnumerationOptions = 1
	FileManagerResumeSyncBehaviorPreserveLocalChanges VolumeEnumerationOptions = 0
	FileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict VolumeEnumerationOptions = 1
	FileManagerResumeSyncBehaviorDropLocalChanges VolumeEnumerationOptions = 2
	FileManagerUploadConflictPolicyDefault VolumeEnumerationOptions = 0
	FileManagerUploadConflictPolicyFailOnConflict VolumeEnumerationOptions = 1
	PointerFunctionsStrongMemory VolumeEnumerationOptions = 2
	PointerFunctionsZeroingWeakMemory VolumeEnumerationOptions = 3
	PointerFunctionsOpaqueMemory VolumeEnumerationOptions = 4
	PointerFunctionsMallocMemory VolumeEnumerationOptions = 5
	PointerFunctionsMachVirtualMemory VolumeEnumerationOptions = 6
	PointerFunctionsWeakMemory VolumeEnumerationOptions = 7
	PointerFunctionsObjectPersonality VolumeEnumerationOptions = 8
	PointerFunctionsOpaquePersonality VolumeEnumerationOptions = 9
	PointerFunctionsObjectPointerPersonality VolumeEnumerationOptions = 10
	PointerFunctionsCStringPersonality VolumeEnumerationOptions = 11
	PointerFunctionsStructPersonality VolumeEnumerationOptions = 12
	PointerFunctionsIntegerPersonality VolumeEnumerationOptions = 13
	PointerFunctionsCopyIn VolumeEnumerationOptions = 14
)

// NSFormattingUnitStyle - Specifies the width of the unit, determining the textual representation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/UnitStyle
type FormattingUnitStyle uint

const (
	FormattingUnitStyleShort FormattingUnitStyle = 1
	FormattingUnitStyleMedium FormattingUnitStyle = 2
	FormattingUnitStyleLong FormattingUnitStyle = 3
)

// NSInlinePresentationIntent - A type that defines presentation intent for runs of characters for traits like emphasis, strikethrough, and code voice.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent
type InlinePresentationIntent uint

const (
	// InlinePresentationIntentLineBreak - An intent that represents a line break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/lineBreak
	InlinePresentationIntentLineBreak InlinePresentationIntent = 1
	// InlinePresentationIntentSoftBreak - An intent that represents a soft line break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/softBreak
	InlinePresentationIntentSoftBreak InlinePresentationIntent = 1
	// InlinePresentationIntentStronglyEmphasized - An intent that represents a strongly emphasized presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/stronglyEmphasized
	InlinePresentationIntentStronglyEmphasized InlinePresentationIntent = 1
)

// NSAttributedStringEnumerationOptions - Options for enumerating attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/EnumerationOptions
type AttributedStringEnumerationOptions uint

// NSSpellingState - Constants for the spelling state attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/SpellingState
type SpellingState uint

// NSBinarySearchingOptions - Options for searches and insertions using 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions
type BinarySearchingOptions uint

const (
	// BinarySearchingFirstEqual - Specifies that the search should return the first object in the range that is equal to the given object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions/firstEqual
	BinarySearchingFirstEqual BinarySearchingOptions = 0
	// BinarySearchingInsertionIndex - Returns the index at which you should insert the object in order to maintain a sorted array.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions/insertionIndex
	BinarySearchingInsertionIndex BinarySearchingOptions = 0
	// BinarySearchingLastEqual - Specifies that the search should return the last object in the range that is equal to the given object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions/lastEqual
	BinarySearchingLastEqual BinarySearchingOptions = 0
)

// NSCalendarUnit - Calendrical units such as year, month, day and hour.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
type CalendarUnit uint

const (
	// CalendarCalendarUnit - Specifies the calendar of the calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSCalendarCalendarUnit
	CalendarCalendarUnit CalendarUnit = 22
	// DayCalendarUnit - Specifies the day unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSDayCalendarUnit
	DayCalendarUnit CalendarUnit = 11
	// HourCalendarUnit - Specifies the hour unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSHourCalendarUnit
	HourCalendarUnit CalendarUnit = 12
	// MinuteCalendarUnit - Specifies the minute unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSMinuteCalendarUnit
	MinuteCalendarUnit CalendarUnit = 13
	// MonthCalendarUnit - Specifies the month unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSMonthCalendarUnit
	MonthCalendarUnit CalendarUnit = 10
	// QuarterCalendarUnit - Specifies the quarter unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSQuarterCalendarUnit
	QuarterCalendarUnit CalendarUnit = 18
	// SecondCalendarUnit - Specifies the second unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSSecondCalendarUnit
	SecondCalendarUnit CalendarUnit = 14
	// WeekOfMonthCalendarUnit - Specifies the original week of a month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekOfMonthCalendarUnit
	WeekOfMonthCalendarUnit CalendarUnit = 19
	// WeekdayCalendarUnit - Specifies the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekdayCalendarUnit
	WeekdayCalendarUnit CalendarUnit = 16
	// YearCalendarUnit - Specifies the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSYearCalendarUnit
	YearCalendarUnit CalendarUnit = 9
	// CalendarUnitCalendar - Identifier for the calendar of a date components object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/calendar
	CalendarUnitCalendar CalendarUnit = 6
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/isLeapMonth
	CalendarUnitIsLeapMonth CalendarUnit = 0
	// CalendarUnitMinute - Identifier for the minute unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/minute
	CalendarUnitMinute CalendarUnit = 0
	// CalendarUnitNanosecond - Identifier for the nanosecond unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/nanosecond
	CalendarUnitNanosecond CalendarUnit = 4
	// CalendarUnitTimeZone - Identifier for the time zone of a date components object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/timeZone
	CalendarUnitTimeZone CalendarUnit = 7
	// CalendarUnitWeekOfMonth - Identifier for the week of the month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekOfMonth
	CalendarUnitWeekOfMonth CalendarUnit = 1
	// CalendarUnitWeekOfYear - Identifier for the week of the year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekOfYear
	CalendarUnitWeekOfYear CalendarUnit = 2
	// CalendarUnitWeekday - Identifier for the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekday
	CalendarUnitWeekday CalendarUnit = 0
	// CalendarUnitYearForWeekOfYear - Identifier for the week-counting year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/yearForWeekOfYear
	CalendarUnitYearForWeekOfYear CalendarUnit = 3
)

// NSCollectionChangeType - The type of change represented in computing the difference of an ordered collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCollectionChangeType
type CollectionChangeType uint

const (
	CollectionChangeInsert CollectionChangeType = 0
)

// NSCompoundPredicateType - Constants that describe the possible types of a compound predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType
type CompoundPredicateType uint

const (
	NotPredicateType CompoundPredicateType = 0
	AndPredicateType CompoundPredicateType = 1
	OrPredicateType CompoundPredicateType = 2
)

// NSDataBase64DecodingOptions - Options to modify the decoding algorithm used to decode Base64 encoded data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions
type DataBase64DecodingOptions uint

const (
	// DataBase64DecodingIgnoreUnknownCharacters - Modify the decoding algorithm so that it ignores unknown non-Base-64 bytes, including line ending characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions/ignoreUnknownCharacters
	DataBase64DecodingIgnoreUnknownCharacters DataBase64DecodingOptions = 1
)

// NSDataBase64EncodingOptions - Options for methods used to Base64 encode data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions
type DataBase64EncodingOptions uint

const (
	DataBase64Encoding64CharacterLineLength DataBase64EncodingOptions = 1
	DataBase64Encoding76CharacterLineLength DataBase64EncodingOptions = 1
	DataBase64EncodingEndLineWithCarriageReturn DataBase64EncodingOptions = 1
	DataBase64EncodingEndLineWithLineFeed DataBase64EncodingOptions = 1
	DataBase64DecodingIgnoreUnknownCharacters DataBase64EncodingOptions = 1
	DataCompressionAlgorithmLZFSE DataBase64EncodingOptions = 0
	DataCompressionAlgorithmLZ4 DataBase64EncodingOptions = 1
	DataCompressionAlgorithmLZMA DataBase64EncodingOptions = 2
	DataCompressionAlgorithmZlib DataBase64EncodingOptions = 3
)

// NSDataCompressionAlgorithm - An algorithm that indicates how to compress or decompress data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
type DataCompressionAlgorithm uint

const (
	// DataCompressionAlgorithmLZ4 - The LZ4 compression algorithm, recommended for fast compression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm/lz4
	DataCompressionAlgorithmLZ4 DataCompressionAlgorithm = 1
	// DataCompressionAlgorithmLZFSE - The LZFSE compression algorithm, recommended for use on Apple platforms.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm/lzfse
	DataCompressionAlgorithmLZFSE DataCompressionAlgorithm = 0
	// DataCompressionAlgorithmLZMA - The LZMA compression algorithm, recommended for high-compression ratio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm/lzma
	DataCompressionAlgorithmLZMA DataCompressionAlgorithm = 2
	// DataCompressionAlgorithmZlib - The zlib compression algorithm, recommended for cross-platform compression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm/zlib
	DataCompressionAlgorithmZlib DataCompressionAlgorithm = 3
)

// NSDataReadingOptions - Options for methods used to read data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions
type DataReadingOptions uint

const (
	// DataReadingMappedAlways - Hint to map the file in if possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/alwaysMapped
	DataReadingMappedAlways DataReadingOptions = 2
	// DataReadingMapped - Deprecated name for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/dataReadingMapped
	DataReadingMapped DataReadingOptions = 3
	// DataReadingMappedIfSafe - A hint indicating the file should be mapped into virtual memory, if possible and safe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/mappedIfSafe
	DataReadingMappedIfSafe DataReadingOptions = 1
	// MappedRead - Deprecated name for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/mappedRead
	MappedRead DataReadingOptions = 4
	// DataReadingUncached - A hint indicating the file should not be stored in the file-system caches.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/uncached
	DataReadingUncached DataReadingOptions = 1
	// UncachedRead - Deprecated name for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/uncachedRead
	UncachedRead DataReadingOptions = 5
)

// NSDataSearchOptions - Options for method used to search data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions
type DataSearchOptions uint

const (
	// DataSearchAnchored - Search is limited to start (or end, if searching backwards) of the data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions/anchored
	DataSearchAnchored DataSearchOptions = 1
	// DataSearchBackwards - Search from the end of the data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions/backwards
	DataSearchBackwards DataSearchOptions = 1
)

// NSDataWritingOptions - Options for methods used to write data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions
type DataWritingOptions uint

const (
	// DataWritingAtomic - An option to write data to an auxiliary file first and then replace the original file with the auxiliary file when the write completes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/atomic
	DataWritingAtomic DataWritingOptions = 1
	// AtomicWrite - An option that attempts to write data to an auxiliary file first and then exchange the files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/atomicWrite
	AtomicWrite DataWritingOptions = 9
	// DataWritingFileProtectionComplete - An option to make the file accessible only while the device is unlocked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/completeFileProtection
	DataWritingFileProtectionComplete DataWritingOptions = 4
	// DataWritingFileProtectionCompleteUnlessOpen - An option to allow the file to be accessible while the device is unlocked or the file is already open.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/completeFileProtectionUnlessOpen
	DataWritingFileProtectionCompleteUnlessOpen DataWritingOptions = 5
	// DataWritingFileProtectionCompleteUntilFirstUserAuthentication - An option to allow the file to be accessible after a user first unlocks the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/completeFileProtectionUntilFirstUserAuthentication
	DataWritingFileProtectionCompleteUntilFirstUserAuthentication DataWritingOptions = 6
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/completeFileProtectionWhenUserInactive
	DataWritingFileProtectionCompleteWhenUserInactive DataWritingOptions = 7
	// DataWritingFileProtectionMask - An option the system uses when determining the file protection options that the system assigns to the data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/fileProtectionMask
	DataWritingFileProtectionMask DataWritingOptions = 8
	// DataWritingFileProtectionNone - An option to not encrypt the file when writing it out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/noFileProtection
	DataWritingFileProtectionNone DataWritingOptions = 3
	// DataWritingWithoutOverwriting - An option that attempts to write data to a file and fails with an error if the destination file already exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/withoutOverwriting
	DataWritingWithoutOverwriting DataWritingOptions = 2
)

// NSEnumerationOptions - Options for block enumeration operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
type EnumerationOptions uint

const (
	// EnumerationConcurrent - Specifies that the Block enumeration should be concurrent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions/concurrent
	EnumerationConcurrent EnumerationOptions = 0
	// EnumerationReverse - Specifies that the enumeration should be performed in reverse.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions/reverse
	EnumerationReverse EnumerationOptions = 0
)

// NSFileCoordinatorWritingOptions - Options to use when changing the contents or attributes of a file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
type FileCoordinatorWritingOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forDeleting
	FileCoordinatorWritingForDeleting FileCoordinatorWritingOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forMerging
	FileCoordinatorWritingForMerging FileCoordinatorWritingOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forMoving
	FileCoordinatorWritingForMoving FileCoordinatorWritingOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forReplacing
	FileCoordinatorWritingForReplacing FileCoordinatorWritingOptions = 1
)

// NSFileManagerResumeSyncBehavior - The behaviors the file manager can apply to resolve conflicts when resuming a sync.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior
type FileManagerResumeSyncBehavior uint

const (
	FileManagerResumeSyncBehaviorPreserveLocalChanges FileManagerResumeSyncBehavior = 0
	FileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict FileManagerResumeSyncBehavior = 1
	FileManagerResumeSyncBehaviorDropLocalChanges FileManagerResumeSyncBehavior = 2
	FileManagerUploadConflictPolicyDefault FileManagerResumeSyncBehavior = 0
	FileManagerUploadConflictPolicyFailOnConflict FileManagerResumeSyncBehavior = 1
	PointerFunctionsStrongMemory FileManagerResumeSyncBehavior = 2
	PointerFunctionsZeroingWeakMemory FileManagerResumeSyncBehavior = 3
	PointerFunctionsOpaqueMemory FileManagerResumeSyncBehavior = 4
	PointerFunctionsMallocMemory FileManagerResumeSyncBehavior = 5
	PointerFunctionsMachVirtualMemory FileManagerResumeSyncBehavior = 6
	PointerFunctionsWeakMemory FileManagerResumeSyncBehavior = 7
	PointerFunctionsObjectPersonality FileManagerResumeSyncBehavior = 8
	PointerFunctionsOpaquePersonality FileManagerResumeSyncBehavior = 9
	PointerFunctionsObjectPointerPersonality FileManagerResumeSyncBehavior = 10
	PointerFunctionsCStringPersonality FileManagerResumeSyncBehavior = 11
	PointerFunctionsStructPersonality FileManagerResumeSyncBehavior = 12
	PointerFunctionsIntegerPersonality FileManagerResumeSyncBehavior = 13
	PointerFunctionsCopyIn FileManagerResumeSyncBehavior = 14
)

// NSFileManagerSupportedSyncControls - An option set of the sync controls available for an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls
type FileManagerSupportedSyncControls uint

const (
	FileManagerSupportedSyncControlsPauseSync FileManagerSupportedSyncControls = 1
	FileManagerSupportedSyncControlsFailUploadOnConflict FileManagerSupportedSyncControls = 1
	FileManagerResumeSyncBehaviorPreserveLocalChanges FileManagerSupportedSyncControls = 0
	FileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict FileManagerSupportedSyncControls = 1
	FileManagerResumeSyncBehaviorDropLocalChanges FileManagerSupportedSyncControls = 2
	FileManagerUploadConflictPolicyDefault FileManagerSupportedSyncControls = 0
	FileManagerUploadConflictPolicyFailOnConflict FileManagerSupportedSyncControls = 1
	PointerFunctionsStrongMemory FileManagerSupportedSyncControls = 2
	PointerFunctionsZeroingWeakMemory FileManagerSupportedSyncControls = 3
	PointerFunctionsOpaqueMemory FileManagerSupportedSyncControls = 4
	PointerFunctionsMallocMemory FileManagerSupportedSyncControls = 5
	PointerFunctionsMachVirtualMemory FileManagerSupportedSyncControls = 6
	PointerFunctionsWeakMemory FileManagerSupportedSyncControls = 7
	PointerFunctionsObjectPersonality FileManagerSupportedSyncControls = 8
	PointerFunctionsOpaquePersonality FileManagerSupportedSyncControls = 9
	PointerFunctionsObjectPointerPersonality FileManagerSupportedSyncControls = 10
	PointerFunctionsCStringPersonality FileManagerSupportedSyncControls = 11
	PointerFunctionsStructPersonality FileManagerSupportedSyncControls = 12
	PointerFunctionsIntegerPersonality FileManagerSupportedSyncControls = 13
	PointerFunctionsCopyIn FileManagerSupportedSyncControls = 14
)

// NSFileManagerUploadLocalVersionConflictPolicy - The policies the file manager can apply to resolve conflicts when uploading a local version of a file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy
type FileManagerUploadLocalVersionConflictPolicy uint

const (
	FileManagerUploadConflictPolicyDefault FileManagerUploadLocalVersionConflictPolicy = 0
	FileManagerUploadConflictPolicyFailOnConflict FileManagerUploadLocalVersionConflictPolicy = 1
	PointerFunctionsStrongMemory FileManagerUploadLocalVersionConflictPolicy = 2
	PointerFunctionsZeroingWeakMemory FileManagerUploadLocalVersionConflictPolicy = 3
	PointerFunctionsOpaqueMemory FileManagerUploadLocalVersionConflictPolicy = 4
	PointerFunctionsMallocMemory FileManagerUploadLocalVersionConflictPolicy = 5
	PointerFunctionsMachVirtualMemory FileManagerUploadLocalVersionConflictPolicy = 6
	PointerFunctionsWeakMemory FileManagerUploadLocalVersionConflictPolicy = 7
	PointerFunctionsObjectPersonality FileManagerUploadLocalVersionConflictPolicy = 8
	PointerFunctionsOpaquePersonality FileManagerUploadLocalVersionConflictPolicy = 9
	PointerFunctionsObjectPointerPersonality FileManagerUploadLocalVersionConflictPolicy = 10
	PointerFunctionsCStringPersonality FileManagerUploadLocalVersionConflictPolicy = 11
	PointerFunctionsStructPersonality FileManagerUploadLocalVersionConflictPolicy = 12
	PointerFunctionsIntegerPersonality FileManagerUploadLocalVersionConflictPolicy = 13
	PointerFunctionsCopyIn FileManagerUploadLocalVersionConflictPolicy = 14
)

// NSGrammaticalCase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase
type GrammaticalCase uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/adessive
	GrammaticalCaseAdessive GrammaticalCase = 7
)

// NSGrammaticalDefiniteness enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness
type GrammaticalDefiniteness uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness/definite
	GrammaticalDefinitenessDefinite GrammaticalDefiniteness = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness/indefinite
	GrammaticalDefinitenessIndefinite GrammaticalDefiniteness = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness/notSet
	GrammaticalDefinitenessNotSet GrammaticalDefiniteness = 0
)

// NSGrammaticalDetermination enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination
type GrammaticalDetermination uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination/dependent
	GrammaticalDeterminationDependent GrammaticalDetermination = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination/independent
	GrammaticalDeterminationIndependent GrammaticalDetermination = 1
)

// NSGrammaticalPerson enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson
type GrammaticalPerson uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/first
	GrammaticalPersonFirst GrammaticalPerson = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/notSet
	GrammaticalPersonNotSet GrammaticalPerson = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/third
	GrammaticalPersonThird GrammaticalPerson = 0
)

// NSGrammaticalPronounType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType
type GrammaticalPronounType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/personal
	GrammaticalPronounTypePersonal GrammaticalPronounType = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/reflexive
	GrammaticalPronounTypeReflexive GrammaticalPronounType = 2
)

// NSItemProviderErrorCode - The error codes that describe problems with consuming data from an item provider.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode
type ItemProviderErrorCode uint

const (
	ItemProviderUnavailableCoercionError ItemProviderErrorCode = 0
	CaseInsensitiveSearch ItemProviderErrorCode = 1
	LiteralSearch ItemProviderErrorCode = 2
	BackwardsSearch ItemProviderErrorCode = 4
	AnchoredSearch ItemProviderErrorCode = 8
	NumericSearch ItemProviderErrorCode = 64
	DiacriticInsensitiveSearch ItemProviderErrorCode = 65
	WidthInsensitiveSearch ItemProviderErrorCode = 66
	ForcedOrderingSearch ItemProviderErrorCode = 67
	RegularExpressionSearch ItemProviderErrorCode = 68
)

// NSItemProviderFileOptions - Data-access specifications that declare how to handle items.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions
type ItemProviderFileOptions uint

const (
	// ItemProviderFileOptionOpenInPlace - A data-access specification declaring that items should open in place, rather than being copied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions/openInPlace
	ItemProviderFileOptionOpenInPlace ItemProviderFileOptions = 1
)

// NSItemProviderRepresentationVisibility - Specifications that control which categories of processes can see an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility
type ItemProviderRepresentationVisibility uint

const (
	ItemProviderRepresentationVisibilityAll ItemProviderRepresentationVisibility = 0
	ItemProviderRepresentationVisibilityOwnProcess ItemProviderRepresentationVisibility = 3
	ItemProviderFileOptionOpenInPlace ItemProviderRepresentationVisibility = 1
	ItemProviderUnavailableCoercionError ItemProviderRepresentationVisibility = 2
	CaseInsensitiveSearch ItemProviderRepresentationVisibility = 1
	LiteralSearch ItemProviderRepresentationVisibility = 2
	BackwardsSearch ItemProviderRepresentationVisibility = 4
	AnchoredSearch ItemProviderRepresentationVisibility = 8
	NumericSearch ItemProviderRepresentationVisibility = 64
	DiacriticInsensitiveSearch ItemProviderRepresentationVisibility = 65
	WidthInsensitiveSearch ItemProviderRepresentationVisibility = 66
	ForcedOrderingSearch ItemProviderRepresentationVisibility = 67
	RegularExpressionSearch ItemProviderRepresentationVisibility = 68
)

// NSKeyValueChange - The kinds of changes that can be observed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueChange
type KeyValueChange uint

const (
	KeyValueChangeSetting KeyValueChange = 1
	KeyValueChangeInsertion KeyValueChange = 2
	KeyValueChangeRemoval KeyValueChange = 3
	KeyValueChangeReplacement KeyValueChange = 4
)

// NSKeyValueObservingOptions - The values that can be returned in a change dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions
type KeyValueObservingOptions uint

const (
	// KeyValueObservingOptionNew - Indicates that the change dictionary should provide the new attribute value, if applicable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions/new
	KeyValueObservingOptionNew KeyValueObservingOptions = 0
	// KeyValueObservingOptionOld - Indicates that the change dictionary should contain the old attribute value, if applicable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions/old
	KeyValueObservingOptionOld KeyValueObservingOptions = 0
	// KeyValueObservingOptionPrior - Whether separate notifications should be sent to the observer before and after each change, instead of a single notification after the change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions/prior
	KeyValueObservingOptionPrior KeyValueObservingOptions = 2
)

// NSKeyValueSetMutationKind enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind
type KeyValueSetMutationKind uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind/intersect
	KeyValueIntersectSetMutation KeyValueSetMutationKind = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind/minus
	KeyValueMinusSetMutation KeyValueSetMutationKind = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind/set
	KeyValueSetSetMutation KeyValueSetMutationKind = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind/union
	KeyValueUnionSetMutation KeyValueSetMutationKind = 1
)

// NSLinguisticTaggerOptions - Constants for linguistic tagger enumeration specifying which tokens to omit and whether to join names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options
type LinguisticTaggerOptions uint

const (
	LinguisticTaggerOmitWords LinguisticTaggerOptions = 1
	LinguisticTaggerOmitPunctuation LinguisticTaggerOptions = 1
	LinguisticTaggerOmitWhitespace LinguisticTaggerOptions = 1
	LinguisticTaggerOmitOther LinguisticTaggerOptions = 1
	LinguisticTaggerJoinNames LinguisticTaggerOptions = 1
)

// NSLinguisticTaggerUnit - Constants representing linguistic units.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTaggerUnit
type LinguisticTaggerUnit uint

const (
	LinguisticTaggerUnitWord LinguisticTaggerUnit = 0
	LinguisticTaggerUnitSentence LinguisticTaggerUnit = 1
	LinguisticTaggerUnitParagraph LinguisticTaggerUnit = 2
)

// NSMachPortOptions - Used to remove access rights to a mach port when the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options
type MachPortOptions uint

const (
	MachPortDeallocateNone MachPortOptions = 0
	WindowsNTOperatingSystem MachPortOptions = 1
	Windows95OperatingSystem MachPortOptions = 2
	SolarisOperatingSystem MachPortOptions = 3
	HPUXOperatingSystem MachPortOptions = 4
	MACHOperatingSystem MachPortOptions = 5
	SunOSOperatingSystem MachPortOptions = 6
	ActivityAnimationTrackingEnabled MachPortOptions = 7
	ActivityTrackingEnabled MachPortOptions = 8
	ActivityBackground MachPortOptions = 0
	ActivityLatencyCritical MachPortOptions = 0
	ActivityUserInteractive MachPortOptions = 1
	ProcessInfoThermalStateNominal MachPortOptions = 2
	ProcessInfoThermalStateFair MachPortOptions = 3
	ProcessInfoThermalStateSerious MachPortOptions = 4
	TextCheckingTypeOrthography MachPortOptions = 1
	TextCheckingTypeSpelling MachPortOptions = 1
	TextCheckingTypeGrammar MachPortOptions = 1
	TextCheckingTypeDate MachPortOptions = 1
	TextCheckingTypeAddress MachPortOptions = 1
	TextCheckingTypeLink MachPortOptions = 1
	TextCheckingTypeQuote MachPortOptions = 1
	TextCheckingTypeDash MachPortOptions = 1
	TextCheckingTypeReplacement MachPortOptions = 1
	TextCheckingTypeCorrection MachPortOptions = 1
	TextCheckingTypeRegularExpression MachPortOptions = 2
	TextCheckingTypePhoneNumber MachPortOptions = 3
	TextCheckingTypeTransitInformation MachPortOptions = 4
)

// NSOrderedCollectionDifferenceCalculationOptions - Constants that specify the options to use when creating an ordered collection difference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifferenceCalculationOptions
type OrderedCollectionDifferenceCalculationOptions uint

const (
	// OrderedCollectionDifferenceCalculationOmitInsertedObjects - An option that indicates that the difference should omit references to the insertions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifferenceCalculationOptions/omitInsertedObjects
	OrderedCollectionDifferenceCalculationOmitInsertedObjects OrderedCollectionDifferenceCalculationOptions = 0
	// OrderedCollectionDifferenceCalculationOmitRemovedObjects - An option that indicates that the difference should omit references to the removals.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifferenceCalculationOptions/omitRemovedObjects
	OrderedCollectionDifferenceCalculationOmitRemovedObjects OrderedCollectionDifferenceCalculationOptions = 0
)

// NSPointerFunctionsOptions - Defines the memory and personality options for an 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
type PointerFunctionsOptions uint

const (
	// PointerFunctionsCStringPersonality - Use a string hash and  ; C-string ‘ ’ style description.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/cStringPersonality
	PointerFunctionsCStringPersonality PointerFunctionsOptions = 9
	// PointerFunctionsCopyIn - Use the memory acquire function to allocate and copy items on input (see  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/copyIn
	PointerFunctionsCopyIn PointerFunctionsOptions = 12
	// PointerFunctionsIntegerPersonality - Use unshifted value as hash and equality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/integerPersonality
	PointerFunctionsIntegerPersonality PointerFunctionsOptions = 11
	// PointerFunctionsMachVirtualMemory - Use Mach memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/machVirtualMemory
	PointerFunctionsMachVirtualMemory PointerFunctionsOptions = 4
	// PointerFunctionsMallocMemory - Use   on removal,   on copy in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/mallocMemory
	PointerFunctionsMallocMemory PointerFunctionsOptions = 3
	// PointerFunctionsObjectPersonality - Use   and   methods for hashing and equality comparisons, use the   method for a description.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/objectPersonality
	PointerFunctionsObjectPersonality PointerFunctionsOptions = 6
	// PointerFunctionsObjectPointerPersonality - Use shifted pointer for the hash value and direct comparison to determine equality; use the   method for a description.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/objectPointerPersonality
	PointerFunctionsObjectPointerPersonality PointerFunctionsOptions = 8
	// PointerFunctionsOpaqueMemory - Take no action when pointers are deleted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/opaqueMemory
	PointerFunctionsOpaqueMemory PointerFunctionsOptions = 2
	// PointerFunctionsOpaquePersonality - Use shifted pointer for the hash value and direct comparison to determine equality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/opaquePersonality
	PointerFunctionsOpaquePersonality PointerFunctionsOptions = 7
	// PointerFunctionsStrongMemory - Use strong write-barriers to backing store; use garbage-collected memory on copy-in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/strongMemory
	PointerFunctionsStrongMemory PointerFunctionsOptions = 0
	// PointerFunctionsStructPersonality - Use a memory hash and   (using a size function that you must set—see  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/structPersonality
	PointerFunctionsStructPersonality PointerFunctionsOptions = 10
	// PointerFunctionsWeakMemory - Uses weak read and write barriers appropriate for ARC or GC. Using NSPointerFunctionsWeakMemory object references will turn to   on last release.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/weakMemory
	PointerFunctionsWeakMemory PointerFunctionsOptions = 5
	// PointerFunctionsZeroingWeakMemory - Use weak read and write barriers; use garbage-collected memory on copyIn.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctionsOptions/NSPointerFunctionsZeroingWeakMemory
	PointerFunctionsZeroingWeakMemory PointerFunctionsOptions = 1
)

// NSPresentationIntentKind - An enumeration of intended display styles for blocks of text like paragraphs, lists, and code blocks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind
type PresentationIntentKind uint

const (
	// PresentationIntentKindBlockQuote - A presentation style for a block quote.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindBlockQuote
	PresentationIntentKindBlockQuote PresentationIntentKind = 6
	// PresentationIntentKindCodeBlock - A presentation style for a block of code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindCodeBlock
	PresentationIntentKindCodeBlock PresentationIntentKind = 5
	// PresentationIntentKindListItem - A presentation style for a list of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindListItem
	PresentationIntentKindListItem PresentationIntentKind = 4
	// PresentationIntentKindOrderedList - A presentation style for an ordered list of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindOrderedList
	PresentationIntentKindOrderedList PresentationIntentKind = 2
	// PresentationIntentKindTableCell - A presentation style for a single cell of a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindTableCell
	PresentationIntentKindTableCell PresentationIntentKind = 11
	// PresentationIntentKindTableHeaderRow - A presentation style for the header row of a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindTableHeaderRow
	PresentationIntentKindTableHeaderRow PresentationIntentKind = 9
)

// NSPresentationIntentTableColumnAlignment - An enumeration of values for aligning the contents of table columns.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment
type PresentationIntentTableColumnAlignment uint

const (
	// PresentationIntentTableColumnAlignmentRight - A presentation style for columns with right-aligned text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment/NSPresentationIntentTableColumnAlignmentRight
	PresentationIntentTableColumnAlignmentRight PresentationIntentTableColumnAlignment = 2
)

// NSSaveOptions - The 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions
type SaveOptions uint

const (
	SaveOptionsYes SaveOptions = 0
	SaveOptionsNo SaveOptions = 1
)

// NSSortOptions - Options for block sorting operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortOptions
type SortOptions uint

const (
	// SortConcurrent - Specifies that the Block sort operation should be concurrent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortOptions/concurrent
	SortConcurrent SortOptions = 0
	// SortStable - Specifies that the sorted results should return compared items having equal value in the order they occurred originally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortOptions/stable
	SortStable SortOptions = 0
)

// NSStringCompareOptions - These values represent the options available to many of the string classes’ search and comparison methods.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions
type StringCompareOptions uint

const (
	// AnchoredSearch - Search is limited to start (or end, if  ) of source string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/anchored
	AnchoredSearch StringCompareOptions = 8
	// BackwardsSearch - Search from end of source string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/backwards
	BackwardsSearch StringCompareOptions = 4
	// CaseInsensitiveSearch - A case-insensitive search.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/caseInsensitive
	CaseInsensitiveSearch StringCompareOptions = 1
	// DiacriticInsensitiveSearch - Search ignores diacritic marks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/diacriticInsensitive
	DiacriticInsensitiveSearch StringCompareOptions = 65
	// ForcedOrderingSearch - Comparisons are forced to return either   or   if the strings are equivalent but not strictly equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/forcedOrdering
	ForcedOrderingSearch StringCompareOptions = 67
	// LiteralSearch - Exact character-by-character equivalence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/literal
	LiteralSearch StringCompareOptions = 2
	// NumericSearch - Numbers within strings are compared using numeric value, that is,   <   <  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/numeric
	NumericSearch StringCompareOptions = 64
	// RegularExpressionSearch - The search string is treated as an ICU-compatible regular expression. If set, no other options can apply except   and  . You can use this option only with the  … methods and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/regularExpression
	RegularExpressionSearch StringCompareOptions = 68
	// WidthInsensitiveSearch - Search ignores width differences in characters that have full-width and half-width forms, as occurs in East Asian character sets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/widthInsensitive
	WidthInsensitiveSearch StringCompareOptions = 66
)

// NSStringDrawingOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions
type StringDrawingOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/disableScreenFontSubstitution
	disableScreenFontSubstitution StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/oneShot
	oneShot StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/optionsResolvesNaturalAlignmentWithBaseWritingDirection
	optionsResolvesNaturalAlignmentWithBaseWritingDirection StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/truncatesLastVisibleLine
	truncatesLastVisibleLine StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/usesDeviceMetrics
	usesDeviceMetrics StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/usesFontLeading
	usesFontLeading StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/usesLineFragmentOrigin
	usesLineFragmentOrigin StringDrawingOptions = 0
)

// NSStringEncodingConversionOptions - Options for converting string encodings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions
type StringEncodingConversionOptions uint

const (
	// StringEncodingConversionAllowLossy - Allows lossy conversion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions/allowLossy
	StringEncodingConversionAllowLossy StringEncodingConversionOptions = 1
	// StringEncodingConversionExternalRepresentation - Specifies an external representation (with a byte-order mark, if necessary, to indicate endianness).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions/externalRepresentation
	StringEncodingConversionExternalRepresentation StringEncodingConversionOptions = 2
)

// NSStringEnumerationOptions - Constants to specify kinds of substrings and styles of enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions
type StringEnumerationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byCaretPositions
	StringEnumerationByCaretPositions StringEnumerationOptions = 5
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byComposedCharacterSequences
	StringEnumerationByComposedCharacterSequences StringEnumerationOptions = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byDeletionClusters
	StringEnumerationByDeletionClusters StringEnumerationOptions = 6
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byLines
	StringEnumerationByLines StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byParagraphs
	StringEnumerationByParagraphs StringEnumerationOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/bySentences
	StringEnumerationBySentences StringEnumerationOptions = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byWords
	StringEnumerationByWords StringEnumerationOptions = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/localized
	StringEnumerationLocalized StringEnumerationOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/reverse
	StringEnumerationReverse StringEnumerationOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/substringNotRequired
	StringEnumerationSubstringNotRequired StringEnumerationOptions = 1
)

// NSTextCheckingType - These constants specify the type of checking the methods should do. They are returned by 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
type TextCheckingType uint

const (
	// TextCheckingTypeAddress - Attempts to locate addresses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/address
	TextCheckingTypeAddress TextCheckingType = 1
	// TextCheckingTypeCorrection - Performs autocorrection on misspelled words.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/correction
	TextCheckingTypeCorrection TextCheckingType = 1
	// TextCheckingTypeDash - Replaces dashes with em-dashes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/dash
	TextCheckingTypeDash TextCheckingType = 1
	// TextCheckingTypeDate - Attempts to locate dates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/date
	TextCheckingTypeDate TextCheckingType = 1
	// TextCheckingTypeGrammar - Checks grammar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/grammar
	TextCheckingTypeGrammar TextCheckingType = 1
	// TextCheckingTypeLink - Attempts to locate URL links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/link
	TextCheckingTypeLink TextCheckingType = 1
	// TextCheckingTypeOrthography - Attempts to identify the language
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/orthography
	TextCheckingTypeOrthography TextCheckingType = 1
	// TextCheckingTypePhoneNumber - Matches a phone number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/phoneNumber
	TextCheckingTypePhoneNumber TextCheckingType = 3
	// TextCheckingTypeQuote - Replaces quotes with smart quotes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/quote
	TextCheckingTypeQuote TextCheckingType = 1
	// TextCheckingTypeRegularExpression - Matches a regular expression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/regularExpression
	TextCheckingTypeRegularExpression TextCheckingType = 2
	// TextCheckingTypeReplacement - Replaces characters such as (c) with the appropriate symbol (in this case ©).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/replacement
	TextCheckingTypeReplacement TextCheckingType = 1
	// TextCheckingTypeSpelling - Checks spelling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/spelling
	TextCheckingTypeSpelling TextCheckingType = 1
	// TextCheckingTypeTransitInformation - Matches a transit information, for example, flight information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/transitInformation
	TextCheckingTypeTransitInformation TextCheckingType = 4
)

// NSURLBookmarkCreationOptions - Options used when creating bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions
type URLBookmarkCreationOptions uint

const (
	// URLBookmarkCreationMinimalBookmark - Specifies that when creating a bookmark, it includes minimal information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/minimalBookmark
	URLBookmarkCreationMinimalBookmark URLBookmarkCreationOptions = 0
	// URLBookmarkCreationPreferFileIDResolution - Specifies that when creating a bookmark, upon resolution, its embedded file ID takes precedence over other sources of information (file system path, for example) when there’s a conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/preferFileIDResolution
	URLBookmarkCreationPreferFileIDResolution URLBookmarkCreationOptions = 0
	// URLBookmarkCreationSecurityScopeAllowOnlyReadAccess - Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read-only access to a file-system resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/securityScopeAllowOnlyReadAccess
	URLBookmarkCreationSecurityScopeAllowOnlyReadAccess URLBookmarkCreationOptions = 2
	// URLBookmarkCreationSuitableForBookmarkFile - Specifies that the bookmark data includes the required properties for creating Finder alias files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/suitableForBookmarkFile
	URLBookmarkCreationSuitableForBookmarkFile URLBookmarkCreationOptions = 0
	// URLBookmarkCreationWithSecurityScope - Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read/write access to a file-system resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/withSecurityScope
	URLBookmarkCreationWithSecurityScope URLBookmarkCreationOptions = 1
)

// NSURLBookmarkResolutionOptions - Options used when resolving bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions
type URLBookmarkResolutionOptions uint

const (
	// URLBookmarkResolutionWithoutImplicitStartAccessing - A property that specifies that resolution doesn’t implicitly start accessing the ephemeral security-scoped resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutImplicitStartAccessing
	URLBookmarkResolutionWithoutImplicitStartAccessing URLBookmarkResolutionOptions = 1
)

// NSURLErrorNetworkUnavailableReason - An enumeration of reasons why a task couldn’t satisfy networking constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLErrorNetworkUnavailableReason
type URLErrorNetworkUnavailableReason uint

const (
	// URLErrorNetworkUnavailableReasonCellular - A reason that indicates network is unavailable because the interface is cellular and cellular network is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLErrorNetworkUnavailableReason/NSURLErrorNetworkUnavailableReasonCellular
	URLErrorNetworkUnavailableReasonCellular URLErrorNetworkUnavailableReason = 0
	// URLErrorNetworkUnavailableReasonConstrained - A reason that indicates network is unavailable because the user enabled “Low Data Mode” in the Settings app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLErrorNetworkUnavailableReason/NSURLErrorNetworkUnavailableReasonConstrained
	URLErrorNetworkUnavailableReasonConstrained URLErrorNetworkUnavailableReason = 0
	// URLErrorNetworkUnavailableReasonExpensive - A reason that indicates network is unavailable because the system marked the interface as expensive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLErrorNetworkUnavailableReason/NSURLErrorNetworkUnavailableReasonExpensive
	URLErrorNetworkUnavailableReasonExpensive URLErrorNetworkUnavailableReason = 0
)

// NSURLRequestAttribution - The entities that can make a network request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/Attribution-swift.enum
type URLRequestAttribution uint

const (
	// URLRequestAttributionDeveloper - A developer-initiated network request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/Attribution-swift.enum/developer
	URLRequestAttributionDeveloper URLRequestAttribution = 0
	// URLRequestAttributionUser - The user explicitly directs the app to make a network request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/Attribution-swift.enum/user
	URLRequestAttributionUser URLRequestAttribution = 0
)

// NSURLRequestCachePolicy - The constants used to specify interaction with the cached responses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum
type URLRequestCachePolicy uint

const (
	// URLRequestReloadIgnoringCacheData - Replaced by  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/reloadIgnoringCacheData
	URLRequestReloadIgnoringCacheData URLRequestCachePolicy = 0
	// URLRequestReloadIgnoringLocalAndRemoteCacheData - Ignore local cache data, and instruct proxies and other intermediates to disregard their caches so far as the protocol allows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/reloadIgnoringLocalAndRemoteCacheData
	URLRequestReloadIgnoringLocalAndRemoteCacheData URLRequestCachePolicy = 0
	// URLRequestReloadIgnoringLocalCacheData - The URL load should be loaded only from the originating source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/reloadIgnoringLocalCacheData
	URLRequestReloadIgnoringLocalCacheData URLRequestCachePolicy = 0
	// URLRequestReloadRevalidatingCacheData - Use cache data if the origin source can validate it; otherwise, load from the origin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/reloadRevalidatingCacheData
	URLRequestReloadRevalidatingCacheData URLRequestCachePolicy = 0
	// URLRequestReturnCacheDataDontLoad - Use existing cache data, regardless or age or expiration date, and fail if no cached data is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/returnCacheDataDontLoad
	URLRequestReturnCacheDataDontLoad URLRequestCachePolicy = 0
	// URLRequestReturnCacheDataElseLoad - Use existing cache data, regardless or age or expiration date, loading from originating source only if there is no cached data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/returnCacheDataElseLoad
	URLRequestReturnCacheDataElseLoad URLRequestCachePolicy = 0
	// URLRequestUseProtocolCachePolicy - Use the caching logic defined in the protocol implementation, if any, for a particular URL load request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/useProtocolCachePolicy
	URLRequestUseProtocolCachePolicy URLRequestCachePolicy = 0
)

// NSURLRequestNetworkServiceType - Constants that specify how a request uses network resources.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum
type URLRequestNetworkServiceType uint

const (
	// URLNetworkServiceTypeAVStreaming - A service type for medium-delay tolerant, low-medium-loss tolerant, elastic flow, constant packet interval, and variable rate and size connections.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum/avStreaming
	URLNetworkServiceTypeAVStreaming URLRequestNetworkServiceType = 0
	// URLNetworkServiceTypeBackground - A service type for high-delay tolerant, high-loss tolerant, elastic flow, and variable size connections.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum/background
	URLNetworkServiceTypeBackground URLRequestNetworkServiceType = 0
	// URLNetworkServiceTypeCallSignaling - A service for low-loss tolerant, inelastic flow, jitter tolerant, short but bursty rate, and variable size connections.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum/callSignaling
	URLNetworkServiceTypeCallSignaling URLRequestNetworkServiceType = 0
	// URLNetworkServiceTypeDefault - A service type for standard network traffic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum/default
	URLNetworkServiceTypeDefault URLRequestNetworkServiceType = 0
	// URLNetworkServiceTypeResponsiveAV - A service type for low-delay tolerant, low-to-medium-loss tolerant, elastic flow, variable packet interval, rate, size responsive and time-sensitive connections.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum/responsiveAV
	URLNetworkServiceTypeResponsiveAV URLRequestNetworkServiceType = 0
	// URLNetworkServiceTypeResponsiveData - A service type for medium-delay tolerant, elastic and inelastic flow, bursty, and long-lived connections.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum/responsiveData
	URLNetworkServiceTypeResponsiveData URLRequestNetworkServiceType = 0
	// URLNetworkServiceTypeVideo - A service type for low-delay tolerant, very low-loss tolerant, inelastic flow, and constant packet rate connections.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum/video
	URLNetworkServiceTypeVideo URLRequestNetworkServiceType = 0
	// URLNetworkServiceTypeVoice - A service type for low-delay tolerant, very low-loss tolerant, inelastic flow, and constant packet rate connections.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum/voice
	URLNetworkServiceTypeVoice URLRequestNetworkServiceType = 0
	// URLNetworkServiceTypeVoIP - A service type for VoIP traffic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum/voip
	URLNetworkServiceTypeVoIP URLRequestNetworkServiceType = 0
)

// NSURLSessionWebSocketMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType
type URLSessionWebSocketMessageType uint

const (
	URLSessionWebSocketMessageTypeData URLSessionWebSocketMessageType = 0
	URLSessionWebSocketMessageTypeString URLSessionWebSocketMessageType = 1
	URLSessionWebSocketCloseCodeInvalid URLSessionWebSocketMessageType = 0
	URLSessionWebSocketCloseCodeNormalClosure URLSessionWebSocketMessageType = 1000
	URLSessionWebSocketCloseCodeGoingAway URLSessionWebSocketMessageType = 1001
	URLSessionWebSocketCloseCodeProtocolError URLSessionWebSocketMessageType = 1002
	URLSessionWebSocketCloseCodeUnsupportedData URLSessionWebSocketMessageType = 1003
	URLSessionWebSocketCloseCodeNoStatusReceived URLSessionWebSocketMessageType = 1005
	URLSessionWebSocketCloseCodeAbnormalClosure URLSessionWebSocketMessageType = 1006
	URLSessionWebSocketCloseCodeInvalidFramePayloadData URLSessionWebSocketMessageType = 1007
	URLSessionWebSocketCloseCodePolicyViolation URLSessionWebSocketMessageType = 1008
	URLSessionWebSocketCloseCodeMessageTooBig URLSessionWebSocketMessageType = 1009
	URLSessionWebSocketCloseCodeMandatoryExtensionMissing URLSessionWebSocketMessageType = 1010
	URLSessionWebSocketCloseCodeInternalServerError URLSessionWebSocketMessageType = 1011
	URLSessionWebSocketCloseCodeTLSHandshakeFailure URLSessionWebSocketMessageType = 1015
	URLSessionMultipathServiceTypeNone URLSessionWebSocketMessageType = 0
	URLSessionMultipathServiceTypeHandover URLSessionWebSocketMessageType = 1
	URLSessionMultipathServiceTypeInteractive URLSessionWebSocketMessageType = 2
	URLSessionMultipathServiceTypeAggregate URLSessionWebSocketMessageType = 3
	URLSessionDelayedRequestContinueLoading URLSessionWebSocketMessageType = 0
	URLSessionDelayedRequestUseNewRequest URLSessionWebSocketMessageType = 1
	URLSessionDelayedRequestCancel URLSessionWebSocketMessageType = 2
	URLSessionAuthChallengeUseCredential URLSessionWebSocketMessageType = 0
	URLSessionAuthChallengePerformDefaultHandling URLSessionWebSocketMessageType = 1
	URLSessionAuthChallengeCancelAuthenticationChallenge URLSessionWebSocketMessageType = 2
	URLSessionAuthChallengeRejectProtectionSpace URLSessionWebSocketMessageType = 3
	URLSessionResponseCancel URLSessionWebSocketMessageType = 0
	URLSessionResponseAllow URLSessionWebSocketMessageType = 1
	URLSessionResponseBecomeDownload URLSessionWebSocketMessageType = 2
	URLSessionResponseBecomeStream URLSessionWebSocketMessageType = 3
	URLSessionTaskMetricsResourceFetchTypeUnknown URLSessionWebSocketMessageType = 4
	URLSessionTaskMetricsResourceFetchTypeNetworkLoad URLSessionWebSocketMessageType = 5
	URLSessionTaskMetricsResourceFetchTypeServerPush URLSessionWebSocketMessageType = 6
	URLSessionTaskMetricsResourceFetchTypeLocalCache URLSessionWebSocketMessageType = 7
	URLSessionTaskMetricsDomainResolutionProtocolUnknown URLSessionWebSocketMessageType = 8
	URLSessionTaskMetricsDomainResolutionProtocolUDP URLSessionWebSocketMessageType = 9
	URLSessionTaskMetricsDomainResolutionProtocolTCP URLSessionWebSocketMessageType = 10
	URLSessionTaskMetricsDomainResolutionProtocolTLS URLSessionWebSocketMessageType = 11
	URLSessionTaskMetricsDomainResolutionProtocolHTTPS URLSessionWebSocketMessageType = 12
	AffineTransformComponents URLSessionWebSocketMessageType = 13
	AffineTransform URLSessionWebSocketMessageType = 14
	BackgroundActivityResultFinished URLSessionWebSocketMessageType = 1
	BackgroundActivityResultDeferred URLSessionWebSocketMessageType = 2
	NotificationSuspensionBehaviorDrop URLSessionWebSocketMessageType = 1
	NotificationSuspensionBehaviorCoalesce URLSessionWebSocketMessageType = 2
	NotificationSuspensionBehaviorHold URLSessionWebSocketMessageType = 3
	NotificationSuspensionBehaviorDeliverImmediately URLSessionWebSocketMessageType = 4
)

// NSUserNotificationActivationType - These constants describe how the user notification was activated.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/ActivationType-swift.enum
type UserNotificationActivationType uint

const (
	// UserNotificationActivationTypeReplied - The user replied to the notification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/ActivationType-swift.enum/replied
	UserNotificationActivationTypeReplied UserNotificationActivationType = 3
)

// NSWhoseSubelementIdentifier enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier
type WhoseSubelementIdentifier uint

const (
	// EverySubelement - Every element that meets the specifier test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/everySubelement
	EverySubelement WhoseSubelementIdentifier = 1
	// IndexSubelement - An element at a given index that meets the specifier test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/indexSubelement
	IndexSubelement WhoseSubelementIdentifier = 0
	// MiddleSubelement - The middle element that meets the specifier test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/middleSubelement
	MiddleSubelement WhoseSubelementIdentifier = 2
	// NoSubelement - No sub-element met the specifier test. Valid only for specifying the end sub-element.; that is, there is no end, so consider all elements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/noSubelement
	NoSubelement WhoseSubelementIdentifier = 4
	// RandomSubelement - Any element that meets the specifier test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/randomSubelement
	RandomSubelement WhoseSubelementIdentifier = 3
)

// NSXPCConnectionOptions - Options that you can pass to a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/Options
type XPCConnectionOptions uint

const (
	FileNoSuchFileError XPCConnectionOptions = 4
	FileLockingError XPCConnectionOptions = 255
	FileReadUnknownError XPCConnectionOptions = 256
	FileReadNoPermissionError XPCConnectionOptions = 257
	FileReadInvalidFileNameError XPCConnectionOptions = 258
	FileReadCorruptFileError XPCConnectionOptions = 259
	FileReadNoSuchFileError XPCConnectionOptions = 260
	FileReadInapplicableStringEncodingError XPCConnectionOptions = 261
	FileReadUnsupportedSchemeError XPCConnectionOptions = 262
	FileReadTooLargeError XPCConnectionOptions = 263
	FileReadUnknownStringEncodingError XPCConnectionOptions = 264
	FileWriteUnknownError XPCConnectionOptions = 512
	FileWriteNoPermissionError XPCConnectionOptions = 513
	FileWriteInvalidFileNameError XPCConnectionOptions = 514
	FileWriteFileExistsError XPCConnectionOptions = 515
	FileWriteInapplicableStringEncodingError XPCConnectionOptions = 517
	FileWriteUnsupportedSchemeError XPCConnectionOptions = 518
	FileWriteOutOfSpaceError XPCConnectionOptions = 640
	FileWriteVolumeReadOnlyError XPCConnectionOptions = 641
	FileManagerUnmountUnknownError XPCConnectionOptions = 642
	FileManagerUnmountBusyError XPCConnectionOptions = 643
	KeyValueValidationError XPCConnectionOptions = 1024
	FormattingError XPCConnectionOptions = 2048
	UserCancelledError XPCConnectionOptions = 3072
	FeatureUnsupportedError XPCConnectionOptions = 3073
	ExecutableNotLoadableError XPCConnectionOptions = 3074
	ExecutableArchitectureMismatchError XPCConnectionOptions = 3075
	ExecutableRuntimeMismatchError XPCConnectionOptions = 3076
	ExecutableLoadError XPCConnectionOptions = 3077
	ExecutableLinkError XPCConnectionOptions = 3078
	FileErrorMinimum XPCConnectionOptions = 0
	FileErrorMaximum XPCConnectionOptions = 1023
	ValidationErrorMinimum XPCConnectionOptions = 1024
	ValidationErrorMaximum XPCConnectionOptions = 2047
	ExecutableErrorMinimum XPCConnectionOptions = 2048
	ExecutableErrorMaximum XPCConnectionOptions = 2049
	FormattingErrorMinimum XPCConnectionOptions = 2048
	FormattingErrorMaximum XPCConnectionOptions = 2559
	PropertyListReadCorruptError XPCConnectionOptions = 2560
	PropertyListReadUnknownVersionError XPCConnectionOptions = 2561
	PropertyListReadStreamError XPCConnectionOptions = 2562
	PropertyListWriteStreamError XPCConnectionOptions = 2563
	PropertyListWriteInvalidError XPCConnectionOptions = 2564
	PropertyListErrorMinimum XPCConnectionOptions = 2565
	PropertyListErrorMaximum XPCConnectionOptions = 2566
	XPCConnectionInterrupted XPCConnectionOptions = 2567
	XPCConnectionInvalid XPCConnectionOptions = 2568
	XPCConnectionReplyInvalid XPCConnectionOptions = 2569
	XPCConnectionCodeSigningRequirementFailure XPCConnectionOptions = 2570
	XPCConnectionErrorMinimum XPCConnectionOptions = 2571
	XPCConnectionErrorMaximum XPCConnectionOptions = 2572
	UbiquitousFileUnavailableError XPCConnectionOptions = 2573
	UbiquitousFileNotUploadedDueToQuotaError XPCConnectionOptions = 2574
	UbiquitousFileUbiquityServerNotAvailable XPCConnectionOptions = 2575
	UbiquitousFileErrorMinimum XPCConnectionOptions = 2576
	UbiquitousFileErrorMaximum XPCConnectionOptions = 2577
	UserActivityHandoffFailedError XPCConnectionOptions = 2578
	UserActivityConnectionUnavailableError XPCConnectionOptions = 2579
	UserActivityRemoteApplicationTimedOutError XPCConnectionOptions = 2580
	UserActivityHandoffUserInfoTooLargeError XPCConnectionOptions = 2581
	UserActivityErrorMinimum XPCConnectionOptions = 2582
	UserActivityErrorMaximum XPCConnectionOptions = 2583
	CoderReadCorruptError XPCConnectionOptions = 2584
	CoderValueNotFoundError XPCConnectionOptions = 2585
	CoderInvalidValueError XPCConnectionOptions = 2586
	CoderErrorMinimum XPCConnectionOptions = 2587
	CoderErrorMaximum XPCConnectionOptions = 2588
	BundleErrorMinimum XPCConnectionOptions = 2589
	BundleErrorMaximum XPCConnectionOptions = 2590
	BundleOnDemandResourceOutOfSpaceError XPCConnectionOptions = 2591
	BundleOnDemandResourceExceededMaximumSizeError XPCConnectionOptions = 2592
	BundleOnDemandResourceInvalidTagError XPCConnectionOptions = 2593
	CloudSharingNetworkFailureError XPCConnectionOptions = 2594
	CloudSharingQuotaExceededError XPCConnectionOptions = 2595
	CloudSharingTooManyParticipantsError XPCConnectionOptions = 2596
	CloudSharingConflictError XPCConnectionOptions = 2597
	CloudSharingNoPermissionError XPCConnectionOptions = 2598
	CloudSharingOtherError XPCConnectionOptions = 2599
	CloudSharingErrorMinimum XPCConnectionOptions = 2600
	CloudSharingErrorMaximum XPCConnectionOptions = 2601
	CompressionFailedError XPCConnectionOptions = 2602
	DecompressionFailedError XPCConnectionOptions = 2603
	CompressionErrorMinimum XPCConnectionOptions = 2604
	CompressionErrorMaximum XPCConnectionOptions = 2605
)

// NSNetServicesError - These constants identify errors that can occur when accessing net services.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum
type NetServicesError uint

const (
	// NetServicesActivityInProgress - The net service cannot process the request at this time. No additional information about the network state is known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/activityInProgress
	NetServicesActivityInProgress NetServicesError = 0
	// NetServicesBadArgumentError - An invalid argument was used when creating the   object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/badArgumentError
	NetServicesBadArgumentError NetServicesError = 0
	// NetServicesCancelledError - The client canceled the action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/cancelledError
	NetServicesCancelledError NetServicesError = 0
	// NetServicesCollisionError - The service could not be published because the name is already in use. The name could be in use locally or on another system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/collisionError
	NetServicesCollisionError NetServicesError = 0
	// NetServicesInvalidError - The net service was improperly configured.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/invalidError
	NetServicesInvalidError NetServicesError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/missingRequiredConfigurationError
	NetServicesMissingRequiredConfigurationError NetServicesError = 0
	// NetServicesNotFoundError - The service could not be found on the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/notFoundError
	NetServicesNotFoundError NetServicesError = 0
	// NetServicesTimeoutError - The net service has timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/timeoutError
	NetServicesTimeoutError NetServicesError = 0
	// NetServicesUnknownError - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/unknownError
	NetServicesUnknownError NetServicesError = 0
)

// NSNetServiceOptions - These constants specify options for a network service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options
type NetServiceOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options/listenForConnections
	NetServiceListenForConnections NetServiceOptions = 2
	// NetServiceNoAutoRename - Specifies that the network service should not rename itself in the event of a name collision.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options/noAutoRename
	NetServiceNoAutoRename NetServiceOptions = 1
)

// NSNotificationCoalescing - The constants that specify how notifications are coalesced.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing
type NotificationCoalescing uint

const (
	// NotificationNoCoalescing - Do not coalesce notifications in the queue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing/none
	NotificationNoCoalescing NotificationCoalescing = 0
	// NotificationCoalescingOnName - Coalesce notifications with the same name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing/onName
	NotificationCoalescingOnName NotificationCoalescing = 1
	// NotificationCoalescingOnSender - Coalesce notifications with the same object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing/onSender
	NotificationCoalescingOnSender NotificationCoalescing = 2
)

// NSNumberFormatterBehavior - These constants specify the behavior of a number formatter. These constants are returned by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Behavior
type NumberFormatterBehavior uint

const (
	NumberFormatterBehaviorDefault NumberFormatterBehavior = 0
	NumberFormatterBehavior10_0 NumberFormatterBehavior = 1000
	NumberFormatterBehavior10_4 NumberFormatterBehavior = 1040
)

// NSNumberFormatterPadPosition - These constants are used to specify how numbers should be padded. These constants are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/PadPosition
type NumberFormatterPadPosition uint

// NSNumberFormatterRoundingMode - These constants are used to specify how numbers should be rounded. These constants are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum
type NumberFormatterRoundingMode uint

// NSNumberFormatterStyle - The predefined number format styles used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style
type NumberFormatterStyle uint

const (
	// NumberFormatterCurrencyStyle - A currency style format that uses the currency symbol defined by the number formatter locale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/currency
	NumberFormatterCurrencyStyle NumberFormatterStyle = 0
	// NumberFormatterCurrencyAccountingStyle - An accounting currency style format that uses the currency symbol defined by the number formatter locale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/currencyAccounting
	NumberFormatterCurrencyAccountingStyle NumberFormatterStyle = 3
	// NumberFormatterCurrencyISOCodeStyle - A currency style format that uses the ISO 4217 currency code defined by the number formatter locale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/currencyISOCode
	NumberFormatterCurrencyISOCodeStyle NumberFormatterStyle = 1
	// NumberFormatterCurrencyPluralStyle - A currency style format that uses the pluralized denomination defined by the number formatter locale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/currencyPlural
	NumberFormatterCurrencyPluralStyle NumberFormatterStyle = 2
	// NumberFormatterDecimalStyle - A decimal style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/decimal
	NumberFormatterDecimalStyle NumberFormatterStyle = 0
	// NumberFormatterNoStyle - An integer representation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/none
	NumberFormatterNoStyle NumberFormatterStyle = 0
	// NumberFormatterOrdinalStyle - An ordinal style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/ordinal
	NumberFormatterOrdinalStyle NumberFormatterStyle = 0
	// NumberFormatterPercentStyle - A percent style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/percent
	NumberFormatterPercentStyle NumberFormatterStyle = 0
	// NumberFormatterScientificStyle - A scientific style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/scientific
	NumberFormatterScientificStyle NumberFormatterStyle = 0
	// NumberFormatterSpellOutStyle - A style format in which numbers are spelled out in the language defined by the number formatter locale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/spellOut
	NumberFormatterSpellOutStyle NumberFormatterStyle = 0
)

// NSOperationQueuePriority - These constants let you prioritize the order in which operations execute.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum
type OperationQueuePriority uint

const (
	// OperationQueuePriorityLow - Operations receive low priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/low
	OperationQueuePriorityLow OperationQueuePriority = 0
	// OperationQueuePriorityNormal - Operations receive the normal priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/normal
	OperationQueuePriorityNormal OperationQueuePriority = 0
	// OperationQueuePriorityVeryHigh - Operations receive very high priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/veryHigh
	OperationQueuePriorityVeryHigh OperationQueuePriority = 8
	// OperationQueuePriorityVeryLow - Operations receive very low priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/veryLow
	OperationQueuePriorityVeryLow OperationQueuePriority = 0
)

// NSActivityOptions - Option flags used with 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions
type ActivityOptions uint

const (
	// ActivityAnimationTrackingEnabled - A flag to track the activity with an animation signpost interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/animationTrackingEnabled
	ActivityAnimationTrackingEnabled ActivityOptions = 0
	// ActivityAutomaticTerminationDisabled - A flag to prevent automatic termination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/automaticTerminationDisabled
	ActivityAutomaticTerminationDisabled ActivityOptions = 0
	// ActivityBackground - A flag to indicate the app has initiated some kind of work, but not as the direct result of user request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/background
	ActivityBackground ActivityOptions = 0
	// ActivityIdleDisplaySleepDisabled - A flag to require the screen to stay powered on.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/idleDisplaySleepDisabled
	ActivityIdleDisplaySleepDisabled ActivityOptions = 0
	// ActivityIdleSystemSleepDisabled - A flag to prevent idle sleep.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/idleSystemSleepDisabled
	ActivityIdleSystemSleepDisabled ActivityOptions = 0
	// ActivityLatencyCritical - A flag to indicate the activity requires the highest amount of timer and I/O precision available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/latencyCritical
	ActivityLatencyCritical ActivityOptions = 0
	// ActivitySuddenTerminationDisabled - A flag to prevent sudden termination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/suddenTerminationDisabled
	ActivitySuddenTerminationDisabled ActivityOptions = 0
	// ActivityTrackingEnabled - A flag to track the activity with a signpost interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/trackingEnabled
	ActivityTrackingEnabled ActivityOptions = 1
	// ActivityUserInitiated - A flag to indicate the app is performing a user-requested action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/userInitiated
	ActivityUserInitiated ActivityOptions = 0
	// ActivityUserInitiatedAllowingIdleSystemSleep - A flag to indicate the app is performing a user-requested action, but that the system can sleep on idle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/userInitiatedAllowingIdleSystemSleep
	ActivityUserInitiatedAllowingIdleSystemSleep ActivityOptions = 0
	// ActivityUserInteractive - A flag to indicate the app is responding to user interaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/userInteractive
	ActivityUserInteractive ActivityOptions = 1
)

// NSProcessInfoThermalState - Values used to indicate the system’s thermal state.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum
type ProcessInfoThermalState uint

const (
	// ProcessInfoThermalStateCritical - The thermal state is significantly impacting the performance of the system and the device needs to cool down.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum/critical
	ProcessInfoThermalStateCritical ProcessInfoThermalState = 0
	// ProcessInfoThermalStateFair - The thermal state is slightly elevated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum/fair
	ProcessInfoThermalStateFair ProcessInfoThermalState = 1
	// ProcessInfoThermalStateNominal - The thermal state is within normal limits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum/nominal
	ProcessInfoThermalStateNominal ProcessInfoThermalState = 0
	// ProcessInfoThermalStateSerious - The thermal state is high.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum/serious
	ProcessInfoThermalStateSerious ProcessInfoThermalState = 2
)

// NSQualityOfService - Constants that indicate the nature and importance of work to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService
type QualityOfService uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/background
	QualityOfServiceBackground QualityOfService = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/default
	QualityOfServiceDefault QualityOfService = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/userInitiated
	QualityOfServiceUserInitiated QualityOfService = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/userInteractive
	QualityOfServiceUserInteractive QualityOfService = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/utility
	QualityOfServiceUtility QualityOfService = 0
)

// NSStreamEvent - Describes the constants that may be sent to the delegate as a bit field in the second parameter of 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Event
type StreamEvent uint

const (
	StreamEventNone StreamEvent = 0
	StreamEventOpenCompleted StreamEvent = 1
	StreamEventHasBytesAvailable StreamEvent = 1
	StreamEventHasSpaceAvailable StreamEvent = 1
	StreamEventErrorOccurred StreamEvent = 1
	StreamEventEndEncountered StreamEvent = 1
)

// NSStreamStatus - The type declared for the constants listed in 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status
type StreamStatus uint

const (
	// StreamStatusAtEnd - There is no more data to read, or no more data can be written to the stream. When this status is returned, the stream is in a “non-blocking” mode and no data are available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/atEnd
	StreamStatusAtEnd StreamStatus = 5
	// StreamStatusClosed - The stream is closed (  has been called on it).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/closed
	StreamStatusClosed StreamStatus = 6
	// StreamStatusError - The remote end of the connection can’t be contacted, or the connection has been severed for some other reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/error
	StreamStatusError StreamStatus = 7
	// StreamStatusNotOpen - The stream is not open for reading or writing. This status is returned before the underlying call to open a stream but after it’s been created.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/notOpen
	StreamStatusNotOpen StreamStatus = 0
	// StreamStatusOpen - The stream is open, but no reading or writing is occurring.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/open
	StreamStatusOpen StreamStatus = 2
	// StreamStatusOpening - The stream is in the process of being opened for reading or for writing. For network streams, this status might include the time after the stream was opened, but while network DNS resolution is happening.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/opening
	StreamStatusOpening StreamStatus = 1
	// StreamStatusReading - Data is being read from the stream. This status would be returned if code on another thread were to call   on the stream while a   call ( ) was in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/reading
	StreamStatusReading StreamStatus = 3
	// StreamStatusWriting - Data is being written to the stream. This status would be returned if code on another thread were to call   on the stream while a   call ( ) was in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/writing
	StreamStatusWriting StreamStatus = 4
)

// NSURLCredentialPersistence - Constants that specify how long the credential will be kept.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum
type URLCredentialPersistence uint

const (
	URLCredentialPersistenceNone URLCredentialPersistence = 0
	URLCredentialPersistenceForSession URLCredentialPersistence = 1
	URLCredentialPersistencePermanent URLCredentialPersistence = 2
	URLCredentialPersistenceSynchronizable URLCredentialPersistence = 3
)

// NSURLSessionAuthChallengeDisposition - Constants passed by session or task delegates to the provided continuation block in response to an authentication challenge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition
type URLSessionAuthChallengeDisposition uint

const (
	// URLSessionAuthChallengeCancelAuthenticationChallenge - Cancel the entire request. The provided credential parameter is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/cancelAuthenticationChallenge
	URLSessionAuthChallengeCancelAuthenticationChallenge URLSessionAuthChallengeDisposition = 2
	// URLSessionAuthChallengePerformDefaultHandling - Use the default handling for the challenge as though this delegate method were not implemented. The provided credential parameter is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/performDefaultHandling
	URLSessionAuthChallengePerformDefaultHandling URLSessionAuthChallengeDisposition = 1
	// URLSessionAuthChallengeRejectProtectionSpace - Reject this challenge, and call the authentication delegate method again with the next authentication protection space. The provided credential parameter is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/rejectProtectionSpace
	URLSessionAuthChallengeRejectProtectionSpace URLSessionAuthChallengeDisposition = 3
	// URLSessionAuthChallengeUseCredential - Use the specified credential, which may be  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/useCredential
	URLSessionAuthChallengeUseCredential URLSessionAuthChallengeDisposition = 0
)

// NSURLSessionDelayedRequestDisposition - The action to take on a delayed URL session task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition
type URLSessionDelayedRequestDisposition uint

const (
	// URLSessionDelayedRequestCancel - A disposition indicating that the task should be canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/cancel
	URLSessionDelayedRequestCancel URLSessionDelayedRequestDisposition = 2
	// URLSessionDelayedRequestContinueLoading - A disposition indicating that the task should proceed with the original request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/continueLoading
	URLSessionDelayedRequestContinueLoading URLSessionDelayedRequestDisposition = 0
	// URLSessionDelayedRequestUseNewRequest - A disposition indicating that the task should use a new request to perform the network load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/useNewRequest
	URLSessionDelayedRequestUseNewRequest URLSessionDelayedRequestDisposition = 1
)

// NSURLSessionMultipathServiceType - Constants that specify the type of service that Multipath TCP uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum
type URLSessionMultipathServiceType uint

// NSURLSessionTaskState - Constants for determining the current state of a task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum
type URLSessionTaskState uint

const (
	// URLSessionTaskStateCanceling - The task has received a   message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum/canceling
	URLSessionTaskStateCanceling URLSessionTaskState = 2
	// URLSessionTaskStateCompleted - The task has completed (without being canceled), and the task’s delegate receives no further callbacks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum/completed
	URLSessionTaskStateCompleted URLSessionTaskState = 3
	// URLSessionTaskStateRunning - The task is currently being serviced by the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum/running
	URLSessionTaskStateRunning URLSessionTaskState = 0
	// URLSessionTaskStateSuspended - The task was suspended by the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum/suspended
	URLSessionTaskStateSuspended URLSessionTaskState = 1
)

// NSURLSessionWebSocketCloseCode - A code that indicates why a WebSocket connection closed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/CloseCode-swift.enum
type URLSessionWebSocketCloseCode uint

const (
	// URLSessionWebSocketCloseCodeInvalid - A code that indicates the connection is still open.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/CloseCode-swift.enum/invalid
	URLSessionWebSocketCloseCodeInvalid URLSessionWebSocketCloseCode = 0
	// URLSessionWebSocketCloseCodeInvalidFramePayloadData - A code that indicates the server terminated the connection because it received data inconsistent with the message’s type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/CloseCode-swift.enum/invalidFramePayloadData
	URLSessionWebSocketCloseCodeInvalidFramePayloadData URLSessionWebSocketCloseCode = 0
	// URLSessionWebSocketCloseCodeMessageTooBig - A code that indicates an endpoint is terminating the connection because it received a message too big for it to process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/CloseCode-swift.enum/messageTooBig
	URLSessionWebSocketCloseCodeMessageTooBig URLSessionWebSocketCloseCode = 0
	// URLSessionWebSocketCloseCodeTLSHandshakeFailure - A reserved code that indicates the connection closed due to the failure to perform a TLS handshake.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/CloseCode-swift.enum/tlsHandshakeFailure
	URLSessionWebSocketCloseCodeTLSHandshakeFailure URLSessionWebSocketCloseCode = 0
)

// NSXMLNodeKind enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum
type XMLNodeKind uint

const (
	// XMLAttributeKind - Specifies an attribute node
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/attribute
	XMLAttributeKind XMLNodeKind = 3
	// XMLAttributeDeclarationKind - Specifies an attribute-list declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/attributeDeclaration
	XMLAttributeDeclarationKind XMLNodeKind = 10
	// XMLCommentKind - Specifies a comment node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/comment
	XMLCommentKind XMLNodeKind = 6
	// XMLElementKind - Specifies an element node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/element
	XMLElementKind XMLNodeKind = 2
	// XMLEntityDeclarationKind - Specifies an entity-declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/entityDeclaration
	XMLEntityDeclarationKind XMLNodeKind = 9
	// XMLInvalidKind - Indicates a node object created without a valid kind being specified (as returned by the   method).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/invalid
	XMLInvalidKind XMLNodeKind = 0
	// XMLNamespaceKind - Specifies a namespace node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/namespace
	XMLNamespaceKind XMLNodeKind = 4
	// XMLNotationDeclarationKind - Specifies a notation declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/notationDeclaration
	XMLNotationDeclarationKind XMLNodeKind = 0
	// XMLProcessingInstructionKind - Specifies a processing-instruction node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/processingInstruction
	XMLProcessingInstructionKind XMLNodeKind = 5
)

// NSXMLNodeOptions - These constants are input and output options for all 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options
type XMLNodeOptions uint

const (
	// XMLNodeOptionsNone - No options are requested for this input or output action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXMLNodeOptions/NSXMLNodeOptionsNone
	XMLNodeOptionsNone XMLNodeOptions = 0
)

// NSXMLParserExternalEntityResolvingPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/ExternalEntityResolvingPolicy-swift.enum
type XMLParserExternalEntityResolvingPolicy uint

const (
	XMLParserResolveExternalEntitiesNever XMLParserExternalEntityResolvingPolicy = 0
	XMLParserResolveExternalEntitiesNoNetwork XMLParserExternalEntityResolvingPolicy = 1
	XMLParserResolveExternalEntitiesSameOriginOnly XMLParserExternalEntityResolvingPolicy = 2
)


