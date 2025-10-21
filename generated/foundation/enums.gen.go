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
	NotificationSuspensionBehaviorCoalesce NotificationSuspensionBehavior = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior/deliverImmediately
	NotificationSuspensionBehaviorDeliverImmediately NotificationSuspensionBehavior = 0
	// NotificationSuspensionBehaviorDrop - The server doesn’t queue any notifications with this name and object until the notification center resumes notification delivery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior/drop
	NotificationSuspensionBehaviorDrop NotificationSuspensionBehavior = 0
	// NotificationSuspensionBehaviorHold - The server holds all matching notifications until the queue has been filled (queue size determined by the server), at which point the server may flush queued notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior/hold
	NotificationSuspensionBehaviorHold NotificationSuspensionBehavior = 0
)

// NSEnergyFormatterUnit - The units supported by the
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit
type EnergyFormatterUnit uint

// NSDirectoryEnumerationOptions - Options for enumerating the contents of directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions
type DirectoryEnumerationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/includesDirectoriesPostOrder
	DirectoryEnumerationIncludesDirectoriesPostOrder DirectoryEnumerationOptions = 0
	// DirectoryEnumerationSkipsHiddenFiles - An option to skip hidden files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsHiddenFiles
	DirectoryEnumerationSkipsHiddenFiles DirectoryEnumerationOptions = 0
	// DirectoryEnumerationSkipsPackageDescendants - An option to treat packages like files and not descend into their contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsPackageDescendants
	DirectoryEnumerationSkipsPackageDescendants DirectoryEnumerationOptions = 0
	// DirectoryEnumerationSkipsSubdirectoryDescendants - An option to perform a shallow enumeration that doesn’t descend into directories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsSubdirectoryDescendants
	DirectoryEnumerationSkipsSubdirectoryDescendants DirectoryEnumerationOptions = 0
)

// NSFileManagerItemReplacementOptions - Options for specifying the behavior of file replacement operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions
type FileManagerItemReplacementOptions uint

const (
	// FileManagerItemReplacementUsingNewMetadataOnly - Only metadata from the new item is used, and metadata from the original item isn’t preserved (default).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions/usingNewMetadataOnly
	FileManagerItemReplacementUsingNewMetadataOnly FileManagerItemReplacementOptions = 0
	// FileManagerItemReplacementWithoutDeletingBackupItem - The backup item remains in place after a successful replacement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions/withoutDeletingBackupItem
	FileManagerItemReplacementWithoutDeletingBackupItem FileManagerItemReplacementOptions = 0
)

// NSSearchPathDirectory - The location of significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
type SearchPathDirectory uint

const (
	// ApplicationDirectory - Supported applications ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationDirectory
	ApplicationDirectory SearchPathDirectory = 0
	// ApplicationScriptsDirectory - The user scripts folder for the calling application ( .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationScriptsDirectory
	ApplicationScriptsDirectory SearchPathDirectory = 0
	// ApplicationSupportDirectory - Application support files ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationSupportDirectory
	ApplicationSupportDirectory SearchPathDirectory = 0
	// DesktopDirectory - The user’s desktop directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/desktopDirectory
	DesktopDirectory SearchPathDirectory = 0
	// DocumentDirectory - Document directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentDirectory
	DocumentDirectory SearchPathDirectory = 0
	// DocumentationDirectory - Documentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentationDirectory
	DocumentationDirectory SearchPathDirectory = 0
	// InputMethodsDirectory - Input Methods  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/inputMethodsDirectory
	InputMethodsDirectory SearchPathDirectory = 0
	// ItemReplacementDirectory - The constant used to create a temporary directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/itemReplacementDirectory
	ItemReplacementDirectory SearchPathDirectory = 0
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
	UserDomainMask SearchPathDomainMask = 0
)

// NSURLRelationship - Constants indicating the relationship between a directory and an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship
type URLRelationship uint

// NSFileManagerUnmountOptions - Options that specify the behavior of an unmount operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions
type FileManagerUnmountOptions uint

// NSVolumeEnumerationOptions - Options for enumerating mounted volumes with the
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions
type VolumeEnumerationOptions uint

// NSFormattingUnitStyle - Specifies the width of the unit, determining the textual representation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/UnitStyle
type FormattingUnitStyle uint

// NSInlinePresentationIntent - A type that defines presentation intent for runs of characters for traits like emphasis, strikethrough, and code voice.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent
type InlinePresentationIntent uint

const (
	// InlinePresentationIntentLineBreak - An intent that represents a line break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/lineBreak
	InlinePresentationIntentLineBreak InlinePresentationIntent = 0
	// InlinePresentationIntentSoftBreak - An intent that represents a soft line break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/softBreak
	InlinePresentationIntentSoftBreak InlinePresentationIntent = 0
	// InlinePresentationIntentStronglyEmphasized - An intent that represents a strongly emphasized presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/stronglyEmphasized
	InlinePresentationIntentStronglyEmphasized InlinePresentationIntent = 0
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
	CalendarCalendarUnit CalendarUnit = 0
	// DayCalendarUnit - Specifies the day unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSDayCalendarUnit
	DayCalendarUnit CalendarUnit = 0
	// HourCalendarUnit - Specifies the hour unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSHourCalendarUnit
	HourCalendarUnit CalendarUnit = 0
	// MinuteCalendarUnit - Specifies the minute unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSMinuteCalendarUnit
	MinuteCalendarUnit CalendarUnit = 0
	// MonthCalendarUnit - Specifies the month unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSMonthCalendarUnit
	MonthCalendarUnit CalendarUnit = 0
	// QuarterCalendarUnit - Specifies the quarter unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSQuarterCalendarUnit
	QuarterCalendarUnit CalendarUnit = 0
	// SecondCalendarUnit - Specifies the second unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSSecondCalendarUnit
	SecondCalendarUnit CalendarUnit = 0
	// WeekOfMonthCalendarUnit - Specifies the original week of a month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekOfMonthCalendarUnit
	WeekOfMonthCalendarUnit CalendarUnit = 0
	// WeekdayCalendarUnit - Specifies the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekdayCalendarUnit
	WeekdayCalendarUnit CalendarUnit = 0
	// YearCalendarUnit - Specifies the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSYearCalendarUnit
	YearCalendarUnit CalendarUnit = 0
	// CalendarUnitCalendar - Identifier for the calendar of a date components object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/calendar
	CalendarUnitCalendar CalendarUnit = 0
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
	CalendarUnitNanosecond CalendarUnit = 0
	// CalendarUnitTimeZone - Identifier for the time zone of a date components object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/timeZone
	CalendarUnitTimeZone CalendarUnit = 0
	// CalendarUnitWeekOfMonth - Identifier for the week of the month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekOfMonth
	CalendarUnitWeekOfMonth CalendarUnit = 0
	// CalendarUnitWeekOfYear - Identifier for the week of the year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekOfYear
	CalendarUnitWeekOfYear CalendarUnit = 0
	// CalendarUnitWeekday - Identifier for the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekday
	CalendarUnitWeekday CalendarUnit = 0
	// CalendarUnitYearForWeekOfYear - Identifier for the week-counting year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/yearForWeekOfYear
	CalendarUnitYearForWeekOfYear CalendarUnit = 0
)

// NSCollectionChangeType - The type of change represented in computing the difference of an ordered collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCollectionChangeType
type CollectionChangeType uint

// NSCompoundPredicateType - Constants that describe the possible types of a compound predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType
type CompoundPredicateType uint

// NSDataBase64DecodingOptions - Options to modify the decoding algorithm used to decode Base64 encoded data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions
type DataBase64DecodingOptions uint

const (
	// DataBase64DecodingIgnoreUnknownCharacters - Modify the decoding algorithm so that it ignores unknown non-Base-64 bytes, including line ending characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions/ignoreUnknownCharacters
	DataBase64DecodingIgnoreUnknownCharacters DataBase64DecodingOptions = 0
)

// NSDataBase64EncodingOptions - Options for methods used to Base64 encode data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions
type DataBase64EncodingOptions uint

// NSDataCompressionAlgorithm - An algorithm that indicates how to compress or decompress data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
type DataCompressionAlgorithm uint

const (
	// DataCompressionAlgorithmLZ4 - The LZ4 compression algorithm, recommended for fast compression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm/lz4
	DataCompressionAlgorithmLZ4 DataCompressionAlgorithm = 0
	// DataCompressionAlgorithmLZFSE - The LZFSE compression algorithm, recommended for use on Apple platforms.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm/lzfse
	DataCompressionAlgorithmLZFSE DataCompressionAlgorithm = 0
	// DataCompressionAlgorithmLZMA - The LZMA compression algorithm, recommended for high-compression ratio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm/lzma
	DataCompressionAlgorithmLZMA DataCompressionAlgorithm = 0
	// DataCompressionAlgorithmZlib - The zlib compression algorithm, recommended for cross-platform compression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm/zlib
	DataCompressionAlgorithmZlib DataCompressionAlgorithm = 0
)

// NSDataReadingOptions - Options for methods used to read data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions
type DataReadingOptions uint

const (
	// DataReadingMappedAlways - Hint to map the file in if possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/alwaysMapped
	DataReadingMappedAlways DataReadingOptions = 0
	// DataReadingMapped - Deprecated name for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/dataReadingMapped
	DataReadingMapped DataReadingOptions = 0
	// DataReadingMappedIfSafe - A hint indicating the file should be mapped into virtual memory, if possible and safe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/mappedIfSafe
	DataReadingMappedIfSafe DataReadingOptions = 0
	// MappedRead - Deprecated name for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/mappedRead
	MappedRead DataReadingOptions = 0
	// DataReadingUncached - A hint indicating the file should not be stored in the file-system caches.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/uncached
	DataReadingUncached DataReadingOptions = 0
	// UncachedRead - Deprecated name for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/uncachedRead
	UncachedRead DataReadingOptions = 0
)

// NSDataSearchOptions - Options for method used to search data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions
type DataSearchOptions uint

const (
	// DataSearchAnchored - Search is limited to start (or end, if searching backwards) of the data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions/anchored
	DataSearchAnchored DataSearchOptions = 0
	// DataSearchBackwards - Search from the end of the data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions/backwards
	DataSearchBackwards DataSearchOptions = 0
)

// NSDataWritingOptions - Options for methods used to write data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions
type DataWritingOptions uint

const (
	// DataWritingAtomic - An option to write data to an auxiliary file first and then replace the original file with the auxiliary file when the write completes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/atomic
	DataWritingAtomic DataWritingOptions = 0
	// AtomicWrite - An option that attempts to write data to an auxiliary file first and then exchange the files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/atomicWrite
	AtomicWrite DataWritingOptions = 0
	// DataWritingFileProtectionComplete - An option to make the file accessible only while the device is unlocked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/completeFileProtection
	DataWritingFileProtectionComplete DataWritingOptions = 0
	// DataWritingFileProtectionCompleteUnlessOpen - An option to allow the file to be accessible while the device is unlocked or the file is already open.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/completeFileProtectionUnlessOpen
	DataWritingFileProtectionCompleteUnlessOpen DataWritingOptions = 0
	// DataWritingFileProtectionCompleteUntilFirstUserAuthentication - An option to allow the file to be accessible after a user first unlocks the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/completeFileProtectionUntilFirstUserAuthentication
	DataWritingFileProtectionCompleteUntilFirstUserAuthentication DataWritingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/completeFileProtectionWhenUserInactive
	DataWritingFileProtectionCompleteWhenUserInactive DataWritingOptions = 0
	// DataWritingFileProtectionMask - An option the system uses when determining the file protection options that the system assigns to the data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/fileProtectionMask
	DataWritingFileProtectionMask DataWritingOptions = 0
	// DataWritingFileProtectionNone - An option to not encrypt the file when writing it out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/noFileProtection
	DataWritingFileProtectionNone DataWritingOptions = 0
	// DataWritingWithoutOverwriting - An option that attempts to write data to a file and fails with an error if the destination file already exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions/withoutOverwriting
	DataWritingWithoutOverwriting DataWritingOptions = 0
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
	FileCoordinatorWritingForDeleting FileCoordinatorWritingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forMerging
	FileCoordinatorWritingForMerging FileCoordinatorWritingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forMoving
	FileCoordinatorWritingForMoving FileCoordinatorWritingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forReplacing
	FileCoordinatorWritingForReplacing FileCoordinatorWritingOptions = 0
)

// NSFileManagerResumeSyncBehavior - The behaviors the file manager can apply to resolve conflicts when resuming a sync.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior
type FileManagerResumeSyncBehavior uint

// NSFileManagerSupportedSyncControls - An option set of the sync controls available for an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls
type FileManagerSupportedSyncControls uint

// NSFileManagerUploadLocalVersionConflictPolicy - The policies the file manager can apply to resolve conflicts when uploading a local version of a file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy
type FileManagerUploadLocalVersionConflictPolicy uint

// NSGrammaticalCase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase
type GrammaticalCase uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/adessive
	GrammaticalCaseAdessive GrammaticalCase = 0
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
	GrammaticalDefinitenessIndefinite GrammaticalDefiniteness = 0
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
	GrammaticalDeterminationIndependent GrammaticalDetermination = 0
)

// NSGrammaticalPerson enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson
type GrammaticalPerson uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/first
	GrammaticalPersonFirst GrammaticalPerson = 0
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
	GrammaticalPronounTypePersonal GrammaticalPronounType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/reflexive
	GrammaticalPronounTypeReflexive GrammaticalPronounType = 0
)

// NSItemProviderErrorCode - The error codes that describe problems with consuming data from an item provider.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode
type ItemProviderErrorCode uint

// NSItemProviderFileOptions - Data-access specifications that declare how to handle items.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions
type ItemProviderFileOptions uint

const (
	// ItemProviderFileOptionOpenInPlace - A data-access specification declaring that items should open in place, rather than being copied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions/openInPlace
	ItemProviderFileOptionOpenInPlace ItemProviderFileOptions = 0
)

// NSItemProviderRepresentationVisibility - Specifications that control which categories of processes can see an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility
type ItemProviderRepresentationVisibility uint

// NSKeyValueChange - The kinds of changes that can be observed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueChange
type KeyValueChange uint

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
	KeyValueObservingOptionPrior KeyValueObservingOptions = 0
)

// NSKeyValueSetMutationKind enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind
type KeyValueSetMutationKind uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind/intersect
	KeyValueIntersectSetMutation KeyValueSetMutationKind = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind/minus
	KeyValueMinusSetMutation KeyValueSetMutationKind = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind/set
	KeyValueSetSetMutation KeyValueSetMutationKind = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind/union
	KeyValueUnionSetMutation KeyValueSetMutationKind = 0
)

// NSLinguisticTaggerOptions - Constants for linguistic tagger enumeration specifying which tokens to omit and whether to join names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options
type LinguisticTaggerOptions uint

// NSMachPortOptions - Used to remove access rights to a mach port when the
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options
type MachPortOptions uint

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
	PointerFunctionsCStringPersonality PointerFunctionsOptions = 0
	// PointerFunctionsCopyIn - Use the memory acquire function to allocate and copy items on input (see  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/copyIn
	PointerFunctionsCopyIn PointerFunctionsOptions = 0
	// PointerFunctionsIntegerPersonality - Use unshifted value as hash and equality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/integerPersonality
	PointerFunctionsIntegerPersonality PointerFunctionsOptions = 0
	// PointerFunctionsMachVirtualMemory - Use Mach memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/machVirtualMemory
	PointerFunctionsMachVirtualMemory PointerFunctionsOptions = 0
	// PointerFunctionsMallocMemory - Use   on removal,   on copy in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/mallocMemory
	PointerFunctionsMallocMemory PointerFunctionsOptions = 0
	// PointerFunctionsObjectPersonality - Use   and   methods for hashing and equality comparisons, use the   method for a description.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/objectPersonality
	PointerFunctionsObjectPersonality PointerFunctionsOptions = 0
	// PointerFunctionsObjectPointerPersonality - Use shifted pointer for the hash value and direct comparison to determine equality; use the   method for a description.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/objectPointerPersonality
	PointerFunctionsObjectPointerPersonality PointerFunctionsOptions = 0
	// PointerFunctionsOpaqueMemory - Take no action when pointers are deleted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/opaqueMemory
	PointerFunctionsOpaqueMemory PointerFunctionsOptions = 0
	// PointerFunctionsOpaquePersonality - Use shifted pointer for the hash value and direct comparison to determine equality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/opaquePersonality
	PointerFunctionsOpaquePersonality PointerFunctionsOptions = 0
	// PointerFunctionsStrongMemory - Use strong write-barriers to backing store; use garbage-collected memory on copy-in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/strongMemory
	PointerFunctionsStrongMemory PointerFunctionsOptions = 0
	// PointerFunctionsStructPersonality - Use a memory hash and   (using a size function that you must set—see  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/structPersonality
	PointerFunctionsStructPersonality PointerFunctionsOptions = 0
	// PointerFunctionsWeakMemory - Uses weak read and write barriers appropriate for ARC or GC. Using NSPointerFunctionsWeakMemory object references will turn to   on last release.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/weakMemory
	PointerFunctionsWeakMemory PointerFunctionsOptions = 0
	// PointerFunctionsZeroingWeakMemory - Use weak read and write barriers; use garbage-collected memory on copyIn.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctionsOptions/NSPointerFunctionsZeroingWeakMemory
	PointerFunctionsZeroingWeakMemory PointerFunctionsOptions = 0
)

// NSPresentationIntentKind - An enumeration of intended display styles for blocks of text like paragraphs, lists, and code blocks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind
type PresentationIntentKind uint

const (
	// PresentationIntentKindBlockQuote - A presentation style for a block quote.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindBlockQuote
	PresentationIntentKindBlockQuote PresentationIntentKind = 0
	// PresentationIntentKindCodeBlock - A presentation style for a block of code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindCodeBlock
	PresentationIntentKindCodeBlock PresentationIntentKind = 0
	// PresentationIntentKindListItem - A presentation style for a list of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindListItem
	PresentationIntentKindListItem PresentationIntentKind = 0
	// PresentationIntentKindOrderedList - A presentation style for an ordered list of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindOrderedList
	PresentationIntentKindOrderedList PresentationIntentKind = 0
	// PresentationIntentKindTableCell - A presentation style for a single cell of a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindTableCell
	PresentationIntentKindTableCell PresentationIntentKind = 0
	// PresentationIntentKindTableHeaderRow - A presentation style for the header row of a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindTableHeaderRow
	PresentationIntentKindTableHeaderRow PresentationIntentKind = 0
)

// NSPresentationIntentTableColumnAlignment - An enumeration of values for aligning the contents of table columns.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment
type PresentationIntentTableColumnAlignment uint

const (
	// PresentationIntentTableColumnAlignmentRight - A presentation style for columns with right-aligned text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment/NSPresentationIntentTableColumnAlignmentRight
	PresentationIntentTableColumnAlignmentRight PresentationIntentTableColumnAlignment = 0
)

// NSSaveOptions - The
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions
type SaveOptions uint

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
	AnchoredSearch StringCompareOptions = 0
	// BackwardsSearch - Search from end of source string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/backwards
	BackwardsSearch StringCompareOptions = 0
	// CaseInsensitiveSearch - A case-insensitive search.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/caseInsensitive
	CaseInsensitiveSearch StringCompareOptions = 0
	// DiacriticInsensitiveSearch - Search ignores diacritic marks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/diacriticInsensitive
	DiacriticInsensitiveSearch StringCompareOptions = 0
	// ForcedOrderingSearch - Comparisons are forced to return either   or   if the strings are equivalent but not strictly equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/forcedOrdering
	ForcedOrderingSearch StringCompareOptions = 0
	// LiteralSearch - Exact character-by-character equivalence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/literal
	LiteralSearch StringCompareOptions = 0
	// NumericSearch - Numbers within strings are compared using numeric value, that is,   <   <  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/numeric
	NumericSearch StringCompareOptions = 0
	// RegularExpressionSearch - The search string is treated as an ICU-compatible regular expression. If set, no other options can apply except   and  . You can use this option only with the  … methods and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/regularExpression
	RegularExpressionSearch StringCompareOptions = 0
	// WidthInsensitiveSearch - Search ignores width differences in characters that have full-width and half-width forms, as occurs in East Asian character sets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/widthInsensitive
	WidthInsensitiveSearch StringCompareOptions = 0
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
	StringEncodingConversionAllowLossy StringEncodingConversionOptions = 0
	// StringEncodingConversionExternalRepresentation - Specifies an external representation (with a byte-order mark, if necessary, to indicate endianness).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions/externalRepresentation
	StringEncodingConversionExternalRepresentation StringEncodingConversionOptions = 0
)

// NSStringEnumerationOptions - Constants to specify kinds of substrings and styles of enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions
type StringEnumerationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byCaretPositions
	StringEnumerationByCaretPositions StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byComposedCharacterSequences
	StringEnumerationByComposedCharacterSequences StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byDeletionClusters
	StringEnumerationByDeletionClusters StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byLines
	StringEnumerationByLines StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byParagraphs
	StringEnumerationByParagraphs StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/bySentences
	StringEnumerationBySentences StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byWords
	StringEnumerationByWords StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/localized
	StringEnumerationLocalized StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/reverse
	StringEnumerationReverse StringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/substringNotRequired
	StringEnumerationSubstringNotRequired StringEnumerationOptions = 0
)

// NSTextCheckingType - These constants specify the type of checking the methods should do. They are returned by
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
type TextCheckingType uint

const (
	// TextCheckingTypeAddress - Attempts to locate addresses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/address
	TextCheckingTypeAddress TextCheckingType = 0
	// TextCheckingTypeCorrection - Performs autocorrection on misspelled words.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/correction
	TextCheckingTypeCorrection TextCheckingType = 0
	// TextCheckingTypeDash - Replaces dashes with em-dashes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/dash
	TextCheckingTypeDash TextCheckingType = 0
	// TextCheckingTypeDate - Attempts to locate dates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/date
	TextCheckingTypeDate TextCheckingType = 0
	// TextCheckingTypeGrammar - Checks grammar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/grammar
	TextCheckingTypeGrammar TextCheckingType = 0
	// TextCheckingTypeLink - Attempts to locate URL links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/link
	TextCheckingTypeLink TextCheckingType = 0
	// TextCheckingTypeOrthography - Attempts to identify the language
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/orthography
	TextCheckingTypeOrthography TextCheckingType = 0
	// TextCheckingTypePhoneNumber - Matches a phone number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/phoneNumber
	TextCheckingTypePhoneNumber TextCheckingType = 0
	// TextCheckingTypeQuote - Replaces quotes with smart quotes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/quote
	TextCheckingTypeQuote TextCheckingType = 0
	// TextCheckingTypeRegularExpression - Matches a regular expression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/regularExpression
	TextCheckingTypeRegularExpression TextCheckingType = 0
	// TextCheckingTypeReplacement - Replaces characters such as (c) with the appropriate symbol (in this case ©).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/replacement
	TextCheckingTypeReplacement TextCheckingType = 0
	// TextCheckingTypeSpelling - Checks spelling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/spelling
	TextCheckingTypeSpelling TextCheckingType = 0
	// TextCheckingTypeTransitInformation - Matches a transit information, for example, flight information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/transitInformation
	TextCheckingTypeTransitInformation TextCheckingType = 0
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
	URLBookmarkCreationSecurityScopeAllowOnlyReadAccess URLBookmarkCreationOptions = 0
	// URLBookmarkCreationSuitableForBookmarkFile - Specifies that the bookmark data includes the required properties for creating Finder alias files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/suitableForBookmarkFile
	URLBookmarkCreationSuitableForBookmarkFile URLBookmarkCreationOptions = 0
	// URLBookmarkCreationWithSecurityScope - Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read/write access to a file-system resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/withSecurityScope
	URLBookmarkCreationWithSecurityScope URLBookmarkCreationOptions = 0
)

// NSURLBookmarkResolutionOptions - Options used when resolving bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions
type URLBookmarkResolutionOptions uint

const (
	// URLBookmarkResolutionWithoutImplicitStartAccessing - A property that specifies that resolution doesn’t implicitly start accessing the ephemeral security-scoped resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutImplicitStartAccessing
	URLBookmarkResolutionWithoutImplicitStartAccessing URLBookmarkResolutionOptions = 0
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

// NSUserNotificationActivationType - These constants describe how the user notification was activated.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/ActivationType-swift.enum
type UserNotificationActivationType uint

const (
	// UserNotificationActivationTypeReplied - The user replied to the notification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/ActivationType-swift.enum/replied
	UserNotificationActivationTypeReplied UserNotificationActivationType = 0
)

// NSWhoseSubelementIdentifier enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier
type WhoseSubelementIdentifier uint

const (
	// EverySubelement - Every element that meets the specifier test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/everySubelement
	EverySubelement WhoseSubelementIdentifier = 0
	// IndexSubelement - An element at a given index that meets the specifier test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/indexSubelement
	IndexSubelement WhoseSubelementIdentifier = 0
	// MiddleSubelement - The middle element that meets the specifier test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/middleSubelement
	MiddleSubelement WhoseSubelementIdentifier = 0
	// NoSubelement - No sub-element met the specifier test. Valid only for specifying the end sub-element.; that is, there is no end, so consider all elements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/noSubelement
	NoSubelement WhoseSubelementIdentifier = 0
	// RandomSubelement - Any element that meets the specifier test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier/randomSubelement
	RandomSubelement WhoseSubelementIdentifier = 0
)

// NSXPCConnectionOptions - Options that you can pass to a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/Options
type XPCConnectionOptions uint

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
	NetServiceListenForConnections NetServiceOptions = 0
	// NetServiceNoAutoRename - Specifies that the network service should not rename itself in the event of a name collision.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options/noAutoRename
	NetServiceNoAutoRename NetServiceOptions = 0
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
	NotificationCoalescingOnName NotificationCoalescing = 0
	// NotificationCoalescingOnSender - Coalesce notifications with the same object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing/onSender
	NotificationCoalescingOnSender NotificationCoalescing = 0
)

// NSNumberFormatterBehavior - These constants specify the behavior of a number formatter. These constants are returned by the
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Behavior
type NumberFormatterBehavior uint

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
	NumberFormatterCurrencyAccountingStyle NumberFormatterStyle = 0
	// NumberFormatterCurrencyISOCodeStyle - A currency style format that uses the ISO 4217 currency code defined by the number formatter locale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/currencyISOCode
	NumberFormatterCurrencyISOCodeStyle NumberFormatterStyle = 0
	// NumberFormatterCurrencyPluralStyle - A currency style format that uses the pluralized denomination defined by the number formatter locale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style/currencyPlural
	NumberFormatterCurrencyPluralStyle NumberFormatterStyle = 0
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
	OperationQueuePriorityVeryHigh OperationQueuePriority = 0
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
	ActivityTrackingEnabled ActivityOptions = 0
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
	ActivityUserInteractive ActivityOptions = 0
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
	ProcessInfoThermalStateFair ProcessInfoThermalState = 0
	// ProcessInfoThermalStateNominal - The thermal state is within normal limits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum/nominal
	ProcessInfoThermalStateNominal ProcessInfoThermalState = 0
	// ProcessInfoThermalStateSerious - The thermal state is high.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum/serious
	ProcessInfoThermalStateSerious ProcessInfoThermalState = 0
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

// NSStreamStatus - The type declared for the constants listed in
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status
type StreamStatus uint

const (
	// StreamStatusAtEnd - There is no more data to read, or no more data can be written to the stream. When this status is returned, the stream is in a “non-blocking” mode and no data are available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/atEnd
	StreamStatusAtEnd StreamStatus = 0
	// StreamStatusClosed - The stream is closed (  has been called on it).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/closed
	StreamStatusClosed StreamStatus = 0
	// StreamStatusError - The remote end of the connection can’t be contacted, or the connection has been severed for some other reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/error
	StreamStatusError StreamStatus = 0
	// StreamStatusNotOpen - The stream is not open for reading or writing. This status is returned before the underlying call to open a stream but after it’s been created.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/notOpen
	StreamStatusNotOpen StreamStatus = 0
	// StreamStatusOpen - The stream is open, but no reading or writing is occurring.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/open
	StreamStatusOpen StreamStatus = 0
	// StreamStatusOpening - The stream is in the process of being opened for reading or for writing. For network streams, this status might include the time after the stream was opened, but while network DNS resolution is happening.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/opening
	StreamStatusOpening StreamStatus = 0
	// StreamStatusReading - Data is being read from the stream. This status would be returned if code on another thread were to call   on the stream while a   call ( ) was in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/reading
	StreamStatusReading StreamStatus = 0
	// StreamStatusWriting - Data is being written to the stream. This status would be returned if code on another thread were to call   on the stream while a   call ( ) was in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status/writing
	StreamStatusWriting StreamStatus = 0
)

// NSURLCredentialPersistence - Constants that specify how long the credential will be kept.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum
type URLCredentialPersistence uint

// NSURLSessionAuthChallengeDisposition - Constants passed by session or task delegates to the provided continuation block in response to an authentication challenge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition
type URLSessionAuthChallengeDisposition uint

const (
	// URLSessionAuthChallengeCancelAuthenticationChallenge - Cancel the entire request. The provided credential parameter is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/cancelAuthenticationChallenge
	URLSessionAuthChallengeCancelAuthenticationChallenge URLSessionAuthChallengeDisposition = 0
	// URLSessionAuthChallengePerformDefaultHandling - Use the default handling for the challenge as though this delegate method were not implemented. The provided credential parameter is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/performDefaultHandling
	URLSessionAuthChallengePerformDefaultHandling URLSessionAuthChallengeDisposition = 0
	// URLSessionAuthChallengeRejectProtectionSpace - Reject this challenge, and call the authentication delegate method again with the next authentication protection space. The provided credential parameter is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/rejectProtectionSpace
	URLSessionAuthChallengeRejectProtectionSpace URLSessionAuthChallengeDisposition = 0
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
	URLSessionDelayedRequestCancel URLSessionDelayedRequestDisposition = 0
	// URLSessionDelayedRequestContinueLoading - A disposition indicating that the task should proceed with the original request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/continueLoading
	URLSessionDelayedRequestContinueLoading URLSessionDelayedRequestDisposition = 0
	// URLSessionDelayedRequestUseNewRequest - A disposition indicating that the task should use a new request to perform the network load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/useNewRequest
	URLSessionDelayedRequestUseNewRequest URLSessionDelayedRequestDisposition = 0
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
	URLSessionTaskStateCanceling URLSessionTaskState = 0
	// URLSessionTaskStateCompleted - The task has completed (without being canceled), and the task’s delegate receives no further callbacks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum/completed
	URLSessionTaskStateCompleted URLSessionTaskState = 0
	// URLSessionTaskStateRunning - The task is currently being serviced by the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum/running
	URLSessionTaskStateRunning URLSessionTaskState = 0
	// URLSessionTaskStateSuspended - The task was suspended by the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum/suspended
	URLSessionTaskStateSuspended URLSessionTaskState = 0
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
	XMLAttributeKind XMLNodeKind = 0
	// XMLAttributeDeclarationKind - Specifies an attribute-list declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/attributeDeclaration
	XMLAttributeDeclarationKind XMLNodeKind = 0
	// XMLCommentKind - Specifies a comment node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/comment
	XMLCommentKind XMLNodeKind = 0
	// XMLElementKind - Specifies an element node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/element
	XMLElementKind XMLNodeKind = 0
	// XMLEntityDeclarationKind - Specifies an entity-declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/entityDeclaration
	XMLEntityDeclarationKind XMLNodeKind = 0
	// XMLInvalidKind - Indicates a node object created without a valid kind being specified (as returned by the   method).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/invalid
	XMLInvalidKind XMLNodeKind = 0
	// XMLNamespaceKind - Specifies a namespace node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/namespace
	XMLNamespaceKind XMLNodeKind = 0
	// XMLNotationDeclarationKind - Specifies a notation declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/notationDeclaration
	XMLNotationDeclarationKind XMLNodeKind = 0
	// XMLProcessingInstructionKind - Specifies a processing-instruction node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/processingInstruction
	XMLProcessingInstructionKind XMLNodeKind = 0
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
