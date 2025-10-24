// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
// ByteCountFormatterCountStyle - Specifies display of file or storage byte counts. The display style is platform specific.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/CountStyle-swift.enum
type ByteCountFormatterCountStyle uint

const (
	ByteCountFormatterCountStyleFile ByteCountFormatterCountStyle = 0
	ByteCountFormatterCountStyleMemory ByteCountFormatterCountStyle = 1
	ByteCountFormatterCountStyleDecimal ByteCountFormatterCountStyle = 2
	ByteCountFormatterCountStyleBinary ByteCountFormatterCountStyle = 3
)

// ByteCountFormatterUnits - Specifies the units appropriate for the formatter to display. Specifying any units explicitly causes just those units to be used in showing the number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units
type ByteCountFormatterUnits uint

const (
	ByteCountFormatterUseDefault ByteCountFormatterUnits = 0
	ByteCountFormatterUseBytes ByteCountFormatterUnits = 1
	ByteCountFormatterUseKB ByteCountFormatterUnits = 2
	ByteCountFormatterUseMB ByteCountFormatterUnits = 4
	ByteCountFormatterUseGB ByteCountFormatterUnits = 8
	ByteCountFormatterUseTB ByteCountFormatterUnits = 16
	ByteCountFormatterUsePB ByteCountFormatterUnits = 32
	ByteCountFormatterUseEB ByteCountFormatterUnits = 64
	ByteCountFormatterUseZB ByteCountFormatterUnits = 128
	ByteCountFormatterUseYBOrHigher ByteCountFormatterUnits = 255
	ByteCountFormatterUseAll ByteCountFormatterUnits = 65535
)

// ComparisonResult - Constants that indicate sort order.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ComparisonResult
type ComparisonResult int

const (
	// OrderedAscending - The left operand is smaller than the right operand.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ComparisonResult/orderedAscending
	OrderedAscending ComparisonResult = -1
	// OrderedDescending - The left operand is greater than the right operand.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ComparisonResult/orderedDescending
	OrderedDescending ComparisonResult = 1
	// OrderedSame - The two operands are equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ComparisonResult/orderedSame
	OrderedSame ComparisonResult = 0
)

// DateComponentsFormatterUnitsStyle - Constants for specifying how to represent quantities of time.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/UnitsStyle-swift.enum
type DateComponentsFormatterUnitsStyle uint

const (
	DateComponentsFormatterUnitsStylePositional DateComponentsFormatterUnitsStyle = 0
	DateComponentsFormatterUnitsStyleAbbreviated DateComponentsFormatterUnitsStyle = 1
	DateComponentsFormatterUnitsStyleShort DateComponentsFormatterUnitsStyle = 2
	DateComponentsFormatterUnitsStyleFull DateComponentsFormatterUnitsStyle = 3
	DateComponentsFormatterUnitsStyleSpellOut DateComponentsFormatterUnitsStyle = 4
	DateComponentsFormatterUnitsStyleBrief DateComponentsFormatterUnitsStyle = 5
)

// DateComponentsFormatterZeroFormattingBehavior - Formatting constants for when values contain zeroes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/ZeroFormattingBehavior-swift.struct
type DateComponentsFormatterZeroFormattingBehavior uint

const (
	DateComponentsFormatterZeroFormattingBehaviorDefault DateComponentsFormatterZeroFormattingBehavior = 1
	DateComponentsFormatterZeroFormattingBehaviorDropLeading DateComponentsFormatterZeroFormattingBehavior = 2
	DateComponentsFormatterZeroFormattingBehaviorDropMiddle DateComponentsFormatterZeroFormattingBehavior = 4
	DateComponentsFormatterZeroFormattingBehaviorDropTrailing DateComponentsFormatterZeroFormattingBehavior = 8
	DateComponentsFormatterZeroFormattingBehaviorPad DateComponentsFormatterZeroFormattingBehavior = 65536
)

// DateFormatterBehavior - Constants that specify the behavior 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Behavior
type DateFormatterBehavior uint

const (
	DateFormatterBehaviorDefault DateFormatterBehavior = 0
	DateFormatterBehavior10_0 DateFormatterBehavior = 1000
	DateFormatterBehavior10_4 DateFormatterBehavior = 1040
)

// DateFormatterStyle - The following constants specify predefined format styles for dates and times.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style
type DateFormatterStyle uint

const (
	// DateFormatterLongStyle - Specifies a long style, typically with full text, such as “November 23, 1937” or “3:30:32 PM PST”. Equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/long
	DateFormatterLongStyle DateFormatterStyle = 0
	// DateFormatterNoStyle - Specifies no style. Equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/none
	DateFormatterNoStyle DateFormatterStyle = 0
	// DateFormatterShortStyle - Specifies a short style, typically numeric only, such as “11/23/37” or “3:30 PM”. Equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/short
	DateFormatterShortStyle DateFormatterStyle = 0
)

// DateIntervalFormatterStyle - Formatting styles for individual date and time values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/Style
type DateIntervalFormatterStyle uint

const (
	DateIntervalFormatterNoStyle DateIntervalFormatterStyle = 0
	DateIntervalFormatterShortStyle DateIntervalFormatterStyle = 1
	DateIntervalFormatterMediumStyle DateIntervalFormatterStyle = 2
	DateIntervalFormatterLongStyle DateIntervalFormatterStyle = 3
	DateIntervalFormatterFullStyle DateIntervalFormatterStyle = 4
)

// DistributedNotificationOptions - These constants specify the behavior of notifications posted using the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/Options
type DistributedNotificationOptions uint

const (
	DistributedNotificationDeliverImmediately DistributedNotificationOptions = 1
	DistributedNotificationPostToAllSessions DistributedNotificationOptions = 2
)

// NotificationSuspensionBehavior - These constants specify the types of notification delivery suspension behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior
type NotificationSuspensionBehavior uint

const (
	NotificationSuspensionBehaviorDrop NotificationSuspensionBehavior = 1
	NotificationSuspensionBehaviorCoalesce NotificationSuspensionBehavior = 2
	NotificationSuspensionBehaviorHold NotificationSuspensionBehavior = 3
	NotificationSuspensionBehaviorDeliverImmediately NotificationSuspensionBehavior = 4
)

// EnergyFormatterUnit - The units supported by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit
type EnergyFormatterUnit uint

const (
	EnergyFormatterUnitJoule EnergyFormatterUnit = 11
	EnergyFormatterUnitKilojoule EnergyFormatterUnit = 14
	EnergyFormatterUnitCalorie EnergyFormatterUnit = 1792
	EnergyFormatterUnitKilocalorie EnergyFormatterUnit = 1792
)

// DirectoryEnumerationOptions - Options for enumerating the contents of directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions
type DirectoryEnumerationOptions uint

const (
	DirectoryEnumerationSkipsSubdirectoryDescendants DirectoryEnumerationOptions = 1
	DirectoryEnumerationSkipsPackageDescendants DirectoryEnumerationOptions = 2
	DirectoryEnumerationSkipsHiddenFiles DirectoryEnumerationOptions = 4
	DirectoryEnumerationIncludesDirectoriesPostOrder DirectoryEnumerationOptions = 5
	DirectoryEnumerationProducesRelativePathURLs DirectoryEnumerationOptions = 6
)

// FileManagerItemReplacementOptions - Options for specifying the behavior of file replacement operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions
type FileManagerItemReplacementOptions uint

const (
	FileManagerItemReplacementUsingNewMetadataOnly FileManagerItemReplacementOptions = 1
	FileManagerItemReplacementWithoutDeletingBackupItem FileManagerItemReplacementOptions = 2
)

// SearchPathDirectory - The location of significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
type SearchPathDirectory uint

const (
	// ApplicationSupportDirectory - Application support files ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationSupportDirectory
	ApplicationSupportDirectory SearchPathDirectory = 14
	// CachesDirectory - Discardable cache files ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/cachesDirectory
	CachesDirectory SearchPathDirectory = 13
	// DocumentDirectory - Document directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentDirectory
	DocumentDirectory SearchPathDirectory = 9
)

// SearchPathDomainMask - Domain constants specifying base locations to use when you search for significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask
type SearchPathDomainMask uint

const (
	UserDomainMask SearchPathDomainMask = 1
	LocalDomainMask SearchPathDomainMask = 2
	NetworkDomainMask SearchPathDomainMask = 4
	SystemDomainMask SearchPathDomainMask = 8
	AllDomainsMask SearchPathDomainMask = 65535
)

// URLRelationship - Constants indicating the relationship between a directory and an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship
type URLRelationship uint

const (
	URLRelationshipContains URLRelationship = 0
	URLRelationshipSame URLRelationship = 1
	URLRelationshipOther URLRelationship = 2
)

// FileManagerUnmountOptions - Options that specify the behavior of an unmount operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions
type FileManagerUnmountOptions uint

const (
	FileManagerUnmountAllPartitionsAndEjectDisk FileManagerUnmountOptions = 1
	FileManagerUnmountWithoutUI FileManagerUnmountOptions = 2
)

// VolumeEnumerationOptions - Options for enumerating mounted volumes with the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions
type VolumeEnumerationOptions uint

const (
	VolumeEnumerationSkipHiddenVolumes VolumeEnumerationOptions = 2
	VolumeEnumerationProduceFileReferenceURLs VolumeEnumerationOptions = 4
)

// FileWrapperReadingOptions - Reading options that can be set by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/ReadingOptions
type FileWrapperReadingOptions uint

const (
	FileWrapperReadingImmediate FileWrapperReadingOptions = 1
	FileWrapperReadingWithoutMapping FileWrapperReadingOptions = 2
)

// FileWrapperWritingOptions - Writing options that can be set by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/WritingOptions
type FileWrapperWritingOptions uint

const (
	FileWrapperWritingAtomic FileWrapperWritingOptions = 1
	FileWrapperWritingWithNameUpdating FileWrapperWritingOptions = 2
)

// FormattingContext - The formatting context for a formatter.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/Context
type FormattingContext uint

const (
	FormattingContextUnknown FormattingContext = 0
	FormattingContextDynamic FormattingContext = 1
	FormattingContextStandalone FormattingContext = 2
	FormattingContextListItem FormattingContext = 3
	FormattingContextBeginningOfSentence FormattingContext = 4
	FormattingContextMiddleOfSentence FormattingContext = 5
)

// FormattingUnitStyle - Specifies the width of the unit, determining the textual representation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/UnitStyle
type FormattingUnitStyle uint

const (
	FormattingUnitStyleShort FormattingUnitStyle = 1
	FormattingUnitStyleMedium FormattingUnitStyle = 2
	FormattingUnitStyleLong FormattingUnitStyle = 3
)

// HTTPCookieAcceptPolicy - Cookie acceptance policies implemented by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy
type HTTPCookieAcceptPolicy uint

const (
	HTTPCookieAcceptPolicyAlways HTTPCookieAcceptPolicy = 0
	HTTPCookieAcceptPolicyNever HTTPCookieAcceptPolicy = 1
	HTTPCookieAcceptPolicyOnlyFromMainDocumentDomain HTTPCookieAcceptPolicy = 2
)

// ISO8601DateFormatOptions - Options used to generate and parse ISO 8601 date representations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options
type ISO8601DateFormatOptions uint

const (
	ISO8601DateFormatWithYear ISO8601DateFormatOptions = 0
	ISO8601DateFormatWithMonth ISO8601DateFormatOptions = 1
	ISO8601DateFormatWithWeekOfYear ISO8601DateFormatOptions = 2
	ISO8601DateFormatWithDay ISO8601DateFormatOptions = 3
	ISO8601DateFormatWithTime ISO8601DateFormatOptions = 4
	ISO8601DateFormatWithTimeZone ISO8601DateFormatOptions = 5
	ISO8601DateFormatWithSpaceBetweenDateAndTime ISO8601DateFormatOptions = 6
	ISO8601DateFormatWithDashSeparatorInDate ISO8601DateFormatOptions = 7
	ISO8601DateFormatWithColonSeparatorInTime ISO8601DateFormatOptions = 8
	ISO8601DateFormatWithColonSeparatorInTimeZone ISO8601DateFormatOptions = 9
	ISO8601DateFormatWithFractionalSeconds ISO8601DateFormatOptions = 10
	ISO8601DateFormatWithFullDate ISO8601DateFormatOptions = 11
	ISO8601DateFormatWithFullTime ISO8601DateFormatOptions = 12
	ISO8601DateFormatWithInternetDateTime ISO8601DateFormatOptions = 13
)

// InlinePresentationIntent - A type that defines presentation intent for runs of characters for traits like emphasis, strikethrough, and code voice.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent
type InlinePresentationIntent uint

const (
	InlinePresentationIntentEmphasized InlinePresentationIntent = 1
	InlinePresentationIntentStronglyEmphasized InlinePresentationIntent = 2
	InlinePresentationIntentCode InlinePresentationIntent = 4
	InlinePresentationIntentStrikethrough InlinePresentationIntent = 32
	InlinePresentationIntentSoftBreak InlinePresentationIntent = 64
	InlinePresentationIntentLineBreak InlinePresentationIntent = 128
	InlinePresentationIntentInlineHTML InlinePresentationIntent = 256
	InlinePresentationIntentBlockHTML InlinePresentationIntent = 512
)

// JSONReadingOptions - Options used when creating Foundation objects from JSON data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions
type JSONReadingOptions uint

const (
	JSONReadingMutableContainers JSONReadingOptions = 1
	JSONReadingMutableLeaves JSONReadingOptions = 2
	JSONReadingFragmentsAllowed JSONReadingOptions = 4
	JSONReadingJSON5Allowed JSONReadingOptions = 5
	JSONReadingTopLevelDictionaryAssumed JSONReadingOptions = 6
	JSONReadingAllowFragments JSONReadingOptions = 7
)

// JSONWritingOptions - Options for writing JSON data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/WritingOptions
type JSONWritingOptions uint

const (
	// JSONWritingFragmentsAllowed - Specifies that the parser should allow top-level objects that aren’t arrays or dictionaries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/WritingOptions/fragmentsAllowed
	JSONWritingFragmentsAllowed JSONWritingOptions = 4
)

// LengthFormatterUnit - The units supported by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit
type LengthFormatterUnit uint

const (
	LengthFormatterUnitMillimeter LengthFormatterUnit = 8
	LengthFormatterUnitCentimeter LengthFormatterUnit = 9
	LengthFormatterUnitMeter LengthFormatterUnit = 11
	LengthFormatterUnitKilometer LengthFormatterUnit = 14
	LengthFormatterUnitInch LengthFormatterUnit = 1280
	LengthFormatterUnitFoot LengthFormatterUnit = 1280
	LengthFormatterUnitYard LengthFormatterUnit = 1280
	LengthFormatterUnitMile LengthFormatterUnit = 1280
)

// MassFormatterUnit - The units supported by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/Unit
type MassFormatterUnit uint

const (
	MassFormatterUnitGram MassFormatterUnit = 11
	MassFormatterUnitKilogram MassFormatterUnit = 14
	MassFormatterUnitOunce MassFormatterUnit = 1536
	MassFormatterUnitPound MassFormatterUnit = 1536
	MassFormatterUnitStone MassFormatterUnit = 1536
)

// MeasurementFormatterUnitOptions - Measurement formatter options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/UnitOptions-swift.struct
type MeasurementFormatterUnitOptions uint

const (
	MeasurementFormatterUnitOptionsProvidedUnit MeasurementFormatterUnitOptions = 1
	MeasurementFormatterUnitOptionsNaturalScale MeasurementFormatterUnitOptions = 2
	MeasurementFormatterUnitOptionsTemperatureWithoutUnit MeasurementFormatterUnitOptions = 4
)

// AppleEventSendOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/SendOptions
type AppleEventSendOptions uint

// AttributedStringEnumerationOptions - Options for enumerating attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/EnumerationOptions
type AttributedStringEnumerationOptions uint

const (
	AttributedStringEnumerationReverse AttributedStringEnumerationOptions = 2
	AttributedStringEnumerationLongestEffectiveRangeNotRequired AttributedStringEnumerationOptions = 1048576
)

// SpellingState - Constants for the spelling state attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/SpellingState
type SpellingState uint

// BackgroundActivityResult - These constants indicate whether background activity has been completed successfully or whether additional processing should be deferred until a more optimal time.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/Result
type BackgroundActivityResult uint

const (
	// BackgroundActivityResultDeferred - System conditions have changed since the time the activity began executing, and deferral of additional work is recommended.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/Result/deferred
	BackgroundActivityResultDeferred BackgroundActivityResult = 2
	// BackgroundActivityResultFinished - The activity has finished executing. If the activity repeats, the next invocation is scheduled by the system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/Result/finished
	BackgroundActivityResultFinished BackgroundActivityResult = 1
)

// BinarySearchingOptions - Options for searches and insertions using 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions
type BinarySearchingOptions uint

const (
	// BinarySearchingFirstEqual - Specifies that the search should return the first object in the range that is equal to the given object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions/firstEqual
	BinarySearchingFirstEqual BinarySearchingOptions = 256
	// BinarySearchingInsertionIndex - Returns the index at which you should insert the object in order to maintain a sorted array.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions/insertionIndex
	BinarySearchingInsertionIndex BinarySearchingOptions = 1024
	// BinarySearchingLastEqual - Specifies that the search should return the last object in the range that is equal to the given object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions/lastEqual
	BinarySearchingLastEqual BinarySearchingOptions = 512
)

// CalendarOptions - The options for arithmetic operations involving calendars.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
type CalendarOptions uint

const (
	CalendarWrapComponents CalendarOptions = 1
	CalendarMatchStrictly CalendarOptions = 2
	CalendarSearchBackwards CalendarOptions = 3
	CalendarMatchPreviousTimePreservingSmallerUnits CalendarOptions = 4
	CalendarMatchNextTimePreservingSmallerUnits CalendarOptions = 5
	CalendarMatchNextTime CalendarOptions = 6
	CalendarMatchFirst CalendarOptions = 7
	CalendarMatchLast CalendarOptions = 8
)

// CalendarUnit - Calendrical units such as year, month, day and hour.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
type CalendarUnit uint

const (
	CalendarUnitQuarter CalendarUnit = 0
	CalendarUnitWeekOfMonth CalendarUnit = 1
	CalendarUnitWeekOfYear CalendarUnit = 2
	CalendarUnitYearForWeekOfYear CalendarUnit = 3
	CalendarUnitNanosecond CalendarUnit = 4
	CalendarUnitDayOfYear CalendarUnit = 5
	CalendarUnitCalendar CalendarUnit = 6
	CalendarUnitTimeZone CalendarUnit = 7
	EraCalendarUnit CalendarUnit = 8
	YearCalendarUnit CalendarUnit = 9
	MonthCalendarUnit CalendarUnit = 10
	DayCalendarUnit CalendarUnit = 11
	HourCalendarUnit CalendarUnit = 12
	MinuteCalendarUnit CalendarUnit = 13
	SecondCalendarUnit CalendarUnit = 14
	WeekCalendarUnit CalendarUnit = 15
	WeekdayCalendarUnit CalendarUnit = 16
	WeekdayOrdinalCalendarUnit CalendarUnit = 17
	QuarterCalendarUnit CalendarUnit = 18
	WeekOfMonthCalendarUnit CalendarUnit = 19
	WeekOfYearCalendarUnit CalendarUnit = 20
	YearForWeekOfYearCalendarUnit CalendarUnit = 21
	CalendarCalendarUnit CalendarUnit = 22
	TimeZoneCalendarUnit CalendarUnit = 23
)

// DecodingFailurePolicy - Policies describing the action the coder should take when encountering decode failures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/DecodingFailurePolicy-swift.enum
type DecodingFailurePolicy uint

const (
	// DecodingFailurePolicyRaiseException - A failure policy that directs the coder to raise an exception.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/DecodingFailurePolicy-swift.enum/raiseException
	DecodingFailurePolicyRaiseException DecodingFailurePolicy = 0
	// DecodingFailurePolicySetErrorAndReturn - A failure policy that directs the coder to capture the failure as an error object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/DecodingFailurePolicy-swift.enum/setErrorAndReturn
	DecodingFailurePolicySetErrorAndReturn DecodingFailurePolicy = 1
)

// ComparisonPredicateModifier - Constants that describe the possible types of modifier for a comparison predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier
type ComparisonPredicateModifier uint

const (
	DirectPredicateModifier ComparisonPredicateModifier = 0
	AllPredicateModifier ComparisonPredicateModifier = 1
	AnyPredicateModifier ComparisonPredicateModifier = 2
)

// PredicateOperatorType - Defines the type of comparison for a comparison predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator
type PredicateOperatorType uint

const (
	LessThanPredicateOperatorType PredicateOperatorType = 0
	LessThanOrEqualToPredicateOperatorType PredicateOperatorType = 1
	GreaterThanPredicateOperatorType PredicateOperatorType = 2
	GreaterThanOrEqualToPredicateOperatorType PredicateOperatorType = 3
	EqualToPredicateOperatorType PredicateOperatorType = 4
	NotEqualToPredicateOperatorType PredicateOperatorType = 5
	MatchesPredicateOperatorType PredicateOperatorType = 6
	LikePredicateOperatorType PredicateOperatorType = 7
	BeginsWithPredicateOperatorType PredicateOperatorType = 8
	EndsWithPredicateOperatorType PredicateOperatorType = 9
	InPredicateOperatorType PredicateOperatorType = 10
	CustomSelectorPredicateOperatorType PredicateOperatorType = 11
	ContainsPredicateOperatorType PredicateOperatorType = 12
	BetweenPredicateOperatorType PredicateOperatorType = 13
)

// ComparisonPredicateOptions - Constants that describe the possible types of string comparison for comparison predicates.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
type ComparisonPredicateOptions uint

const (
	CaseInsensitivePredicateOption ComparisonPredicateOptions = 1
	DiacriticInsensitivePredicateOption ComparisonPredicateOptions = 2
	NormalizedPredicateOption ComparisonPredicateOptions = 3
)

// CompoundPredicateType - Constants that describe the possible types of a compound predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType
type CompoundPredicateType uint

const (
	NotPredicateType CompoundPredicateType = 0
	AndPredicateType CompoundPredicateType = 1
	OrPredicateType CompoundPredicateType = 2
)

// DataBase64DecodingOptions - Options to modify the decoding algorithm used to decode Base64 encoded data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions
type DataBase64DecodingOptions uint

const (
	DataBase64DecodingIgnoreUnknownCharacters DataBase64DecodingOptions = 1
)

// DataBase64EncodingOptions - Options for methods used to Base64 encode data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions
type DataBase64EncodingOptions uint

const (
	DataBase64Encoding64CharacterLineLength DataBase64EncodingOptions = 1
	DataBase64Encoding76CharacterLineLength DataBase64EncodingOptions = 2
	DataBase64EncodingEndLineWithCarriageReturn DataBase64EncodingOptions = 16
	DataBase64EncodingEndLineWithLineFeed DataBase64EncodingOptions = 32
)

// DataCompressionAlgorithm - An algorithm that indicates how to compress or decompress data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
type DataCompressionAlgorithm uint

const (
	DataCompressionAlgorithmLZFSE DataCompressionAlgorithm = 0
	DataCompressionAlgorithmLZ4 DataCompressionAlgorithm = 1
	DataCompressionAlgorithmLZMA DataCompressionAlgorithm = 2
	DataCompressionAlgorithmZlib DataCompressionAlgorithm = 3
)

// DataReadingOptions - Options for methods used to read data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions
type DataReadingOptions uint

const (
	DataReadingMappedIfSafe DataReadingOptions = 1
	DataReadingUncached DataReadingOptions = 2
	DataReadingMappedAlways DataReadingOptions = 3
	DataReadingMapped DataReadingOptions = 4
	MappedRead DataReadingOptions = 5
	UncachedRead DataReadingOptions = 6
)

// DataSearchOptions - Options for method used to search data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions
type DataSearchOptions uint

const (
	DataSearchBackwards DataSearchOptions = 1
	DataSearchAnchored DataSearchOptions = 2
)

// DataWritingOptions - Options for methods used to write data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions
type DataWritingOptions uint

const (
	DataWritingAtomic DataWritingOptions = 1
	DataWritingWithoutOverwriting DataWritingOptions = 2
	DataWritingFileProtectionNone DataWritingOptions = 3
	DataWritingFileProtectionComplete DataWritingOptions = 4
	DataWritingFileProtectionCompleteUnlessOpen DataWritingOptions = 5
	DataWritingFileProtectionCompleteUntilFirstUserAuthentication DataWritingOptions = 6
	DataWritingFileProtectionCompleteWhenUserInactive DataWritingOptions = 7
	DataWritingFileProtectionMask DataWritingOptions = 8
	AtomicWrite DataWritingOptions = 9
)

// CalculationError - Calculation error constants used to describe an error in 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError
type CalculationError uint

const (
	CalculationNoError CalculationError = 0
	CalculationLossOfPrecision CalculationError = 1
	CalculationUnderflow CalculationError = 2
	CalculationOverflow CalculationError = 3
	CalculationDivideByZero CalculationError = 4
)

// RoundingMode - These constants specify rounding behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/RoundingMode
type RoundingMode uint

const (
	RoundPlain RoundingMode = 0
	RoundDown RoundingMode = 1
	RoundUp RoundingMode = 2
	RoundBankers RoundingMode = 3
)

// EnumerationOptions - Options for block enumeration operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
type EnumerationOptions uint

const (
	// EnumerationConcurrent - Specifies that the Block enumeration should be concurrent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions/concurrent
	EnumerationConcurrent EnumerationOptions = 1
	// EnumerationReverse - Specifies that the enumeration should be performed in reverse.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions/reverse
	EnumerationReverse EnumerationOptions = 2
)

// ExpressionType - Defines the possible types of an expression.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum
type ExpressionType uint

const (
	// AggregateExpressionType - An expression that defines an aggregate of   objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/aggregate
	AggregateExpressionType ExpressionType = 9
	// IntersectSetExpressionType - An expression that creates an intersection of the results of two nested expressions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/intersectSet
	IntersectSetExpressionType ExpressionType = 6
	// MinusSetExpressionType - An expression that combines two nested expression results by set subtraction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/minusSet
	MinusSetExpressionType ExpressionType = 7
	// SubqueryExpressionType - An expression that filters a collection using a subpredicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/subquery
	SubqueryExpressionType ExpressionType = 8
	// UnionSetExpressionType - An expression that creates a union of the results of two nested expressions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/unionSet
	UnionSetExpressionType ExpressionType = 5
)

// FileCoordinatorReadingOptions - Options to use when reading the contents or attributes of a file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions
type FileCoordinatorReadingOptions uint

const (
	FileCoordinatorReadingWithoutChanges FileCoordinatorReadingOptions = 1
	FileCoordinatorReadingResolvesSymbolicLink FileCoordinatorReadingOptions = 2
	FileCoordinatorReadingImmediatelyAvailableMetadataOnly FileCoordinatorReadingOptions = 3
	FileCoordinatorReadingForUploading FileCoordinatorReadingOptions = 4
)

// FileCoordinatorWritingOptions - Options to use when changing the contents or attributes of a file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
type FileCoordinatorWritingOptions uint

const (
	FileCoordinatorWritingForDeleting FileCoordinatorWritingOptions = 1
	FileCoordinatorWritingForMoving FileCoordinatorWritingOptions = 2
	FileCoordinatorWritingForMerging FileCoordinatorWritingOptions = 4
	FileCoordinatorWritingForReplacing FileCoordinatorWritingOptions = 8
	FileCoordinatorWritingContentIndependentMetadataOnly FileCoordinatorWritingOptions = 9
)

// FileManagerResumeSyncBehavior - The behaviors the file manager can apply to resolve conflicts when resuming a sync.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior
type FileManagerResumeSyncBehavior uint

const (
	FileManagerResumeSyncBehaviorPreserveLocalChanges FileManagerResumeSyncBehavior = 0
	FileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict FileManagerResumeSyncBehavior = 1
	FileManagerResumeSyncBehaviorDropLocalChanges FileManagerResumeSyncBehavior = 2
)

// FileManagerSupportedSyncControls - An option set of the sync controls available for an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls
type FileManagerSupportedSyncControls uint

const (
	FileManagerSupportedSyncControlsPauseSync FileManagerSupportedSyncControls = 1
	FileManagerSupportedSyncControlsFailUploadOnConflict FileManagerSupportedSyncControls = 2
)

// FileManagerUploadLocalVersionConflictPolicy - The policies the file manager can apply to resolve conflicts when uploading a local version of a file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy
type FileManagerUploadLocalVersionConflictPolicy uint

const (
	FileManagerUploadConflictPolicyDefault FileManagerUploadLocalVersionConflictPolicy = 0
	FileManagerUploadConflictPolicyFailOnConflict FileManagerUploadLocalVersionConflictPolicy = 1
)

// FileVersionAddingOptions - Options for adding a new file version.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/AddingOptions
type FileVersionAddingOptions uint

const (
	FileVersionAddingByMoving FileVersionAddingOptions = 1
)

// FileVersionReplacingOptions - Options for replacing a file version.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/ReplacingOptions
type FileVersionReplacingOptions uint

const (
	FileVersionReplacingByMoving FileVersionReplacingOptions = 1
)

// GrammaticalCase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase
type GrammaticalCase uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/ablative
	GrammaticalCaseAblative GrammaticalCase = 6
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/accusative
	GrammaticalCaseAccusative GrammaticalCase = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/adessive
	GrammaticalCaseAdessive GrammaticalCase = 7
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/allative
	GrammaticalCaseAllative GrammaticalCase = 8
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/dative
	GrammaticalCaseDative GrammaticalCase = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/elative
	GrammaticalCaseElative GrammaticalCase = 9
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/essive
	GrammaticalCaseEssive GrammaticalCase = 11
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/genitive
	GrammaticalCaseGenitive GrammaticalCase = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/illative
	GrammaticalCaseIllative GrammaticalCase = 10
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/inessive
	GrammaticalCaseInessive GrammaticalCase = 12
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/locative
	GrammaticalCaseLocative GrammaticalCase = 13
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/nominative
	GrammaticalCaseNominative GrammaticalCase = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/notSet
	GrammaticalCaseNotSet GrammaticalCase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/prepositional
	GrammaticalCasePrepositional GrammaticalCase = 5
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/translative
	GrammaticalCaseTranslative GrammaticalCase = 14
)

// GrammaticalDefiniteness enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness
type GrammaticalDefiniteness uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness/definite
	GrammaticalDefinitenessDefinite GrammaticalDefiniteness = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness/indefinite
	GrammaticalDefinitenessIndefinite GrammaticalDefiniteness = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness/notSet
	GrammaticalDefinitenessNotSet GrammaticalDefiniteness = 0
)

// GrammaticalDetermination enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination
type GrammaticalDetermination uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination/dependent
	GrammaticalDeterminationDependent GrammaticalDetermination = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination/independent
	GrammaticalDeterminationIndependent GrammaticalDetermination = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination/notSet
	GrammaticalDeterminationNotSet GrammaticalDetermination = 0
)

// GrammaticalGender - A representation of grammatical gender, used for inflecting strings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalGender
type GrammaticalGender uint

const (
	GrammaticalGenderNotSet GrammaticalGender = 0
	GrammaticalGenderFeminine GrammaticalGender = 1
	GrammaticalGenderMasculine GrammaticalGender = 2
	GrammaticalGenderNeuter GrammaticalGender = 3
)

// GrammaticalNumber - A representation of grammatical number, used for inflecting strings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber
type GrammaticalNumber uint

const (
	GrammaticalNumberNotSet GrammaticalNumber = 0
	GrammaticalNumberSingular GrammaticalNumber = 1
	GrammaticalNumberZero GrammaticalNumber = 2
	GrammaticalNumberPlural GrammaticalNumber = 3
	GrammaticalNumberPluralTwo GrammaticalNumber = 4
	GrammaticalNumberPluralFew GrammaticalNumber = 5
	GrammaticalNumberPluralMany GrammaticalNumber = 6
)

// GrammaticalPartOfSpeech - A representation of grammatical parts of speech, used for inflecting strings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech
type GrammaticalPartOfSpeech uint

const (
	GrammaticalPartOfSpeechNotSet GrammaticalPartOfSpeech = 0
	GrammaticalPartOfSpeechDeterminer GrammaticalPartOfSpeech = 1
	GrammaticalPartOfSpeechPronoun GrammaticalPartOfSpeech = 2
	GrammaticalPartOfSpeechLetter GrammaticalPartOfSpeech = 3
	GrammaticalPartOfSpeechAdverb GrammaticalPartOfSpeech = 4
	GrammaticalPartOfSpeechParticle GrammaticalPartOfSpeech = 5
	GrammaticalPartOfSpeechAdjective GrammaticalPartOfSpeech = 6
	GrammaticalPartOfSpeechAdposition GrammaticalPartOfSpeech = 7
	GrammaticalPartOfSpeechVerb GrammaticalPartOfSpeech = 8
	GrammaticalPartOfSpeechNoun GrammaticalPartOfSpeech = 9
	GrammaticalPartOfSpeechConjunction GrammaticalPartOfSpeech = 10
	GrammaticalPartOfSpeechNumeral GrammaticalPartOfSpeech = 11
	GrammaticalPartOfSpeechInterjection GrammaticalPartOfSpeech = 12
	GrammaticalPartOfSpeechPreposition GrammaticalPartOfSpeech = 13
	GrammaticalPartOfSpeechAbbreviation GrammaticalPartOfSpeech = 14
)

// GrammaticalPerson enum type
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
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/second
	GrammaticalPersonSecond GrammaticalPerson = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/third
	GrammaticalPersonThird GrammaticalPerson = 3
)

// GrammaticalPronounType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType
type GrammaticalPronounType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/notSet
	GrammaticalPronounTypeNotSet GrammaticalPronounType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/personal
	GrammaticalPronounTypePersonal GrammaticalPronounType = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/possessive
	GrammaticalPronounTypePossessive GrammaticalPronounType = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/reflexive
	GrammaticalPronounTypeReflexive GrammaticalPronounType = 2
)

// KeyValueChange - The kinds of changes that can be observed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueChange
type KeyValueChange uint

const (
	// KeyValueChangeInsertion - Indicates that an object has been inserted into the to-many relationship that is being observed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueChange/insertion
	KeyValueChangeInsertion KeyValueChange = 2
	// KeyValueChangeRemoval - Indicates that an object has been removed from the to-many relationship that is being observed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueChange/removal
	KeyValueChangeRemoval KeyValueChange = 3
	// KeyValueChangeReplacement - Indicates that an object has been replaced in the to-many relationship that is being observed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueChange/replacement
	KeyValueChangeReplacement KeyValueChange = 4
	// KeyValueChangeSetting - Indicates that the value of the observed key path was set to a new value. This change can occur when observing an attribute of an object, as well as properties that specify to-one and to-many relationships.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueChange/setting
	KeyValueChangeSetting KeyValueChange = 1
)

// KeyValueObservingOptions - The values that can be returned in a change dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions
type KeyValueObservingOptions uint

const (
	// KeyValueObservingOptionInitial - If specified, a notification should be sent to the observer immediately, before the observer registration method even returns.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions/initial
	KeyValueObservingOptionInitial KeyValueObservingOptions = 3
	// KeyValueObservingOptionNew - Indicates that the change dictionary should provide the new attribute value, if applicable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions/new
	KeyValueObservingOptionNew KeyValueObservingOptions = 1
	// KeyValueObservingOptionOld - Indicates that the change dictionary should contain the old attribute value, if applicable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions/old
	KeyValueObservingOptionOld KeyValueObservingOptions = 2
	// KeyValueObservingOptionPrior - Whether separate notifications should be sent to the observer before and after each change, instead of a single notification after the change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions/prior
	KeyValueObservingOptionPrior KeyValueObservingOptions = 4
)

// KeyValueSetMutationKind enum type
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

// LinguisticTaggerOptions - Constants for linguistic tagger enumeration specifying which tokens to omit and whether to join names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options
type LinguisticTaggerOptions uint

const (
	LinguisticTaggerOmitWords LinguisticTaggerOptions = 1
	LinguisticTaggerOmitPunctuation LinguisticTaggerOptions = 2
	LinguisticTaggerOmitWhitespace LinguisticTaggerOptions = 4
	LinguisticTaggerOmitOther LinguisticTaggerOptions = 8
	LinguisticTaggerJoinNames LinguisticTaggerOptions = 16
)

// LinguisticTaggerUnit - Constants representing linguistic units.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTaggerUnit
type LinguisticTaggerUnit uint

const (
	LinguisticTaggerUnitWord LinguisticTaggerUnit = 0
	LinguisticTaggerUnitSentence LinguisticTaggerUnit = 1
	LinguisticTaggerUnitParagraph LinguisticTaggerUnit = 2
	LinguisticTaggerUnitDocument LinguisticTaggerUnit = 3
)

// LocaleLanguageDirection - The directions that a language may take across a page of text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection
type LocaleLanguageDirection uint

// OrderedCollectionDifferenceCalculationOptions - Constants that specify the options to use when creating an ordered collection difference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifferenceCalculationOptions
type OrderedCollectionDifferenceCalculationOptions uint

const (
	// OrderedCollectionDifferenceCalculationInferMoves - An option that identifies insertions or removals as moves.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifferenceCalculationOptions/inferMoves
	OrderedCollectionDifferenceCalculationInferMoves OrderedCollectionDifferenceCalculationOptions = 4
	// OrderedCollectionDifferenceCalculationOmitInsertedObjects - An option that indicates that the difference should omit references to the insertions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifferenceCalculationOptions/omitInsertedObjects
	OrderedCollectionDifferenceCalculationOmitInsertedObjects OrderedCollectionDifferenceCalculationOptions = 1
	// OrderedCollectionDifferenceCalculationOmitRemovedObjects - An option that indicates that the difference should omit references to the removals.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifferenceCalculationOptions/omitRemovedObjects
	OrderedCollectionDifferenceCalculationOmitRemovedObjects OrderedCollectionDifferenceCalculationOptions = 2
)

// PointerFunctionsOptions - Defines the memory and personality options for an 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
type PointerFunctionsOptions uint

const (
	PointerFunctionsStrongMemory PointerFunctionsOptions = 0
	PointerFunctionsZeroingWeakMemory PointerFunctionsOptions = 1
	PointerFunctionsOpaqueMemory PointerFunctionsOptions = 2
	PointerFunctionsMallocMemory PointerFunctionsOptions = 3
	PointerFunctionsMachVirtualMemory PointerFunctionsOptions = 4
	PointerFunctionsWeakMemory PointerFunctionsOptions = 5
	PointerFunctionsObjectPersonality PointerFunctionsOptions = 6
	PointerFunctionsOpaquePersonality PointerFunctionsOptions = 7
	PointerFunctionsObjectPointerPersonality PointerFunctionsOptions = 8
	PointerFunctionsCStringPersonality PointerFunctionsOptions = 9
	PointerFunctionsStructPersonality PointerFunctionsOptions = 10
	PointerFunctionsIntegerPersonality PointerFunctionsOptions = 11
	PointerFunctionsCopyIn PointerFunctionsOptions = 12
)

// InsertionPosition - The following constants are defined by 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/InsertionPosition
type InsertionPosition uint

const (
	PositionAfter InsertionPosition = 0
	PositionBefore InsertionPosition = 1
	PositionBeginning InsertionPosition = 2
	PositionEnd InsertionPosition = 3
	PositionReplace InsertionPosition = 4
)

// PresentationIntentKind - An enumeration of intended display styles for blocks of text like paragraphs, lists, and code blocks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind
type PresentationIntentKind int

const (
	// PresentationIntentKindBlockQuote - A presentation style for a block quote.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindBlockQuote
	PresentationIntentKindBlockQuote PresentationIntentKind = 6
	// PresentationIntentKindCodeBlock - A presentation style for a block of code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindCodeBlock
	PresentationIntentKindCodeBlock PresentationIntentKind = 5
	// PresentationIntentKindHeader - A presentation style for a section header.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindHeader
	PresentationIntentKindHeader PresentationIntentKind = 1
	// PresentationIntentKindListItem - A presentation style for a list of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindListItem
	PresentationIntentKindListItem PresentationIntentKind = 4
	// PresentationIntentKindOrderedList - A presentation style for an ordered list of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindOrderedList
	PresentationIntentKindOrderedList PresentationIntentKind = 2
	// PresentationIntentKindParagraph - A presentation style for a paragraph of text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindParagraph
	PresentationIntentKindParagraph PresentationIntentKind = 0
	// PresentationIntentKindTable - A presentation style for a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindTable
	PresentationIntentKindTable PresentationIntentKind = 8
	// PresentationIntentKindTableCell - A presentation style for a single cell of a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindTableCell
	PresentationIntentKindTableCell PresentationIntentKind = 11
	// PresentationIntentKindTableHeaderRow - A presentation style for the header row of a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindTableHeaderRow
	PresentationIntentKindTableHeaderRow PresentationIntentKind = 9
	// PresentationIntentKindTableRow - A presentation style for a row of a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindTableRow
	PresentationIntentKindTableRow PresentationIntentKind = 10
	// PresentationIntentKindThematicBreak - A presentation style for a horizontal rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindThematicBreak
	PresentationIntentKindThematicBreak PresentationIntentKind = 7
	// PresentationIntentKindUnorderedList - A presentation style for an unordered list of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindUnorderedList
	PresentationIntentKindUnorderedList PresentationIntentKind = 3
)

// PresentationIntentTableColumnAlignment - An enumeration of values for aligning the contents of table columns.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment
type PresentationIntentTableColumnAlignment int

const (
	PresentationIntentTableColumnAlignmentLeft PresentationIntentTableColumnAlignment = 0
	PresentationIntentTableColumnAlignmentCenter PresentationIntentTableColumnAlignment = 1
	PresentationIntentTableColumnAlignmentRight PresentationIntentTableColumnAlignment = 2
)

// RectEdge enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge
type RectEdge uint

const (
	MinXEdge RectEdge = 0
	MinYEdge RectEdge = 1
	MaxXEdge RectEdge = 2
	MaxYEdge RectEdge = 3
)

// MatchingFlags - Set by the Block as the matching progresses, completes, or fails. Used by the method 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags
type MatchingFlags uint

const (
	MatchingProgress MatchingFlags = 1
	MatchingCompleted MatchingFlags = 2
	MatchingHitEnd MatchingFlags = 4
	MatchingRequiredEnd MatchingFlags = 8
	MatchingInternalError MatchingFlags = 16
)

// MatchingOptions - The matching options constants specify the reporting, completion and matching rules to the expression matching methods. These constants are used by all methods that search for, or replace values, using a regular expression.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
type MatchingOptions uint

const (
	// MatchingReportCompletion - Call the Block once after the completion of any matching. This option has no effect for methods other than  . See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/reportCompletion
	MatchingReportCompletion MatchingOptions = 2
	// MatchingReportProgress - Call the Block periodically during long-running match operations. This option has no effect for methods other than  . See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/reportProgress
	MatchingReportProgress MatchingOptions = 1
)

// RegularExpressionOptions - These constants define the regular expression options. These constants are used by the property 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
type RegularExpressionOptions uint

const (
	// RegularExpressionAnchorsMatchLines - Allow   and   to match the start and end of lines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/anchorsMatchLines
	RegularExpressionAnchorsMatchLines RegularExpressionOptions = 16
	// RegularExpressionCaseInsensitive - Match letters in the pattern independent of case.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/caseInsensitive
	RegularExpressionCaseInsensitive RegularExpressionOptions = 1
	// RegularExpressionDotMatchesLineSeparators - Allow   to match any character, including line separators.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/dotMatchesLineSeparators
	RegularExpressionDotMatchesLineSeparators RegularExpressionOptions = 8
	// RegularExpressionUseUnicodeWordBoundaries - Use Unicode   to specify word boundaries (otherwise, traditional regular expression word boundaries are used).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/useUnicodeWordBoundaries
	RegularExpressionUseUnicodeWordBoundaries RegularExpressionOptions = 64
)

// RelativePosition - These constants are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier/RelativePosition-swift.enum
type RelativePosition uint

const (
	RelativeAfter RelativePosition = 0
	RelativeBefore RelativePosition = 1
)

// SaveOptions - The 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions
type SaveOptions uint

const (
	SaveOptionsYes SaveOptions = 0
	SaveOptionsNo SaveOptions = 1
	SaveOptionsAsk SaveOptions = 2
)

// SortOptions - Options for block sorting operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortOptions
type SortOptions uint

const (
	// SortConcurrent - Specifies that the Block sort operation should be concurrent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortOptions/concurrent
	SortConcurrent SortOptions = 1
	// SortStable - Specifies that the sorted results should return compared items having equal value in the order they occurred originally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortOptions/stable
	SortStable SortOptions = 16
)

// TestComparisonOperation - These are passed to  
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSpecifierTest/TestComparisonOperation
type TestComparisonOperation uint

const (
	EqualToComparison TestComparisonOperation = 0
	LessThanOrEqualToComparison TestComparisonOperation = 1
	LessThanComparison TestComparisonOperation = 2
	GreaterThanOrEqualToComparison TestComparisonOperation = 3
	GreaterThanComparison TestComparisonOperation = 4
	BeginsWithComparison TestComparisonOperation = 5
	EndsWithComparison TestComparisonOperation = 6
	ContainsComparison TestComparisonOperation = 7
)

// StringCompareOptions - These values represent the options available to many of the string classes’ search and comparison methods.
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

// StringDrawingOptions enum type
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

// StringEncodingConversionOptions - Options for converting string encodings.
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

// StringEnumerationOptions - Constants to specify kinds of substrings and styles of enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions
type StringEnumerationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byCaretPositions
	StringEnumerationByCaretPositions StringEnumerationOptions = 5
	// StringEnumerationByComposedCharacterSequences - Enumerates by composed character sequences. Equivalent to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byComposedCharacterSequences
	StringEnumerationByComposedCharacterSequences StringEnumerationOptions = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byDeletionClusters
	StringEnumerationByDeletionClusters StringEnumerationOptions = 6
	// StringEnumerationByLines - Enumerates by lines. Equivalent to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byLines
	StringEnumerationByLines StringEnumerationOptions = 0
	// StringEnumerationByParagraphs - Enumerates by paragraphs. Equivalent to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byParagraphs
	StringEnumerationByParagraphs StringEnumerationOptions = 1
	// StringEnumerationBySentences - Enumerates by sentences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/bySentences
	StringEnumerationBySentences StringEnumerationOptions = 4
	// StringEnumerationByWords - Enumerates by words.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byWords
	StringEnumerationByWords StringEnumerationOptions = 3
	// StringEnumerationLocalized - Causes the enumeration to occur using the current locale. This does not make a difference in line, paragraph, or composed character sequence enumeration, but it may for words or sentences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/localized
	StringEnumerationLocalized StringEnumerationOptions = 1024
	// StringEnumerationReverse - Causes enumeration to occur from the end of the specified range to the start.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/reverse
	StringEnumerationReverse StringEnumerationOptions = 256
	// StringEnumerationSubstringNotRequired - A way to indicate that the block does not need substring, in which case   will be passed. This is simply a performance shortcut.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/substringNotRequired
	StringEnumerationSubstringNotRequired StringEnumerationOptions = 512
)

// TextCheckingType - These constants specify the type of checking the methods should do. They are returned by 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
type TextCheckingType uint

const (
	// TextCheckingTypeDate - Attempts to locate dates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/date
	TextCheckingTypeDate TextCheckingType = 8
	// TextCheckingTypeLink - Attempts to locate URL links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/link
	TextCheckingTypeLink TextCheckingType = 32
	// TextCheckingTypeRegularExpression - Matches a regular expression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/regularExpression
	TextCheckingTypeRegularExpression TextCheckingType = 513
)

// TimeZoneNameStyle - Constants you use to specify a style when presenting time zone names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle
type TimeZoneNameStyle uint

const (
	TimeZoneNameStyleStandard TimeZoneNameStyle = 0
	TimeZoneNameStyleShortStandard TimeZoneNameStyle = 1
	TimeZoneNameStyleDaylightSaving TimeZoneNameStyle = 2
	TimeZoneNameStyleShortDaylightSaving TimeZoneNameStyle = 3
	TimeZoneNameStyleGeneric TimeZoneNameStyle = 4
	TimeZoneNameStyleShortGeneric TimeZoneNameStyle = 5
)

// URLErrorNetworkUnavailableReason - An enumeration of reasons why a task couldn’t satisfy networking constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLErrorNetworkUnavailableReason
type URLErrorNetworkUnavailableReason int

const (
	URLErrorNetworkUnavailableReasonCellular URLErrorNetworkUnavailableReason = 0
	URLErrorNetworkUnavailableReasonExpensive URLErrorNetworkUnavailableReason = 1
	URLErrorNetworkUnavailableReasonConstrained URLErrorNetworkUnavailableReason = 2
)

// URLRequestAttribution - The entities that can make a network request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/Attribution-swift.enum
type URLRequestAttribution uint

const (
	URLRequestAttributionDeveloper URLRequestAttribution = 0
	URLRequestAttributionUser URLRequestAttribution = 1
)

// URLRequestCachePolicy - The constants used to specify interaction with the cached responses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum
type URLRequestCachePolicy uint

const (
	// URLRequestReloadIgnoringLocalCacheData - The URL load should be loaded only from the originating source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/reloadIgnoringLocalCacheData
	URLRequestReloadIgnoringLocalCacheData URLRequestCachePolicy = 1
	// URLRequestReturnCacheDataDontLoad - Use existing cache data, regardless or age or expiration date, and fail if no cached data is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/returnCacheDataDontLoad
	URLRequestReturnCacheDataDontLoad URLRequestCachePolicy = 3
	// URLRequestReturnCacheDataElseLoad - Use existing cache data, regardless or age or expiration date, loading from originating source only if there is no cached data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/returnCacheDataElseLoad
	URLRequestReturnCacheDataElseLoad URLRequestCachePolicy = 2
	// URLRequestUseProtocolCachePolicy - Use the caching logic defined in the protocol implementation, if any, for a particular URL load request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/useProtocolCachePolicy
	URLRequestUseProtocolCachePolicy URLRequestCachePolicy = 0
)

// URLRequestNetworkServiceType - Constants that specify how a request uses network resources.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum
type URLRequestNetworkServiceType uint

const (
	URLNetworkServiceTypeDefault URLRequestNetworkServiceType = 0
	URLNetworkServiceTypeVoIP URLRequestNetworkServiceType = 1
	URLNetworkServiceTypeVideo URLRequestNetworkServiceType = 2
	URLNetworkServiceTypeBackground URLRequestNetworkServiceType = 3
	URLNetworkServiceTypeVoice URLRequestNetworkServiceType = 4
	URLNetworkServiceTypeResponsiveData URLRequestNetworkServiceType = 6
	URLNetworkServiceTypeAVStreaming URLRequestNetworkServiceType = 7
	URLNetworkServiceTypeResponsiveAV URLRequestNetworkServiceType = 8
	URLNetworkServiceTypeCallSignaling URLRequestNetworkServiceType = 9
)

// URLSessionWebSocketMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType
type URLSessionWebSocketMessageType int

const (
	URLSessionWebSocketMessageTypeData URLSessionWebSocketMessageType = 0
	URLSessionWebSocketMessageTypeString URLSessionWebSocketMessageType = 1
)

// UserNotificationActivationType - These constants describe how the user notification was activated.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/ActivationType-swift.enum
type UserNotificationActivationType uint

const (
	UserNotificationActivationTypeNone UserNotificationActivationType = 0
	UserNotificationActivationTypeContentsClicked UserNotificationActivationType = 1
	UserNotificationActivationTypeActionButtonClicked UserNotificationActivationType = 2
	UserNotificationActivationTypeReplied UserNotificationActivationType = 3
	UserNotificationActivationTypeAdditionalActionClicked UserNotificationActivationType = 4
)

// WhoseSubelementIdentifier enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier
type WhoseSubelementIdentifier uint

const (
	IndexSubelement WhoseSubelementIdentifier = 0
	EverySubelement WhoseSubelementIdentifier = 1
	MiddleSubelement WhoseSubelementIdentifier = 2
	RandomSubelement WhoseSubelementIdentifier = 3
	NoSubelement WhoseSubelementIdentifier = 4
)

// XPCConnectionOptions - Options that you can pass to a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/Options
type XPCConnectionOptions uint

const (
	XPCConnectionPrivileged XPCConnectionOptions = 4096
)

// NetServicesError - These constants identify errors that can occur when accessing net services.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum
type NetServicesError int

const (
	NetServicesUnknownError NetServicesError = -72000
	NetServicesCollisionError NetServicesError = -72001
	NetServicesNotFoundError NetServicesError = -72002
	NetServicesActivityInProgress NetServicesError = -72003
	NetServicesBadArgumentError NetServicesError = -72004
	NetServicesCancelledError NetServicesError = -72005
	NetServicesInvalidError NetServicesError = -72006
	NetServicesTimeoutError NetServicesError = -72007
	NetServicesMissingRequiredConfigurationError NetServicesError = -72006
)

// NetServiceOptions - These constants specify options for a network service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options
type NetServiceOptions uint

const (
	NetServiceNoAutoRename NetServiceOptions = 1
	NetServiceListenForConnections NetServiceOptions = 2
)

// NotificationCoalescing - The constants that specify how notifications are coalesced.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing
type NotificationCoalescing uint

const (
	NotificationNoCoalescing NotificationCoalescing = 0
	NotificationCoalescingOnName NotificationCoalescing = 1
	NotificationCoalescingOnSender NotificationCoalescing = 2
)

// PostingStyle - The constants that specify when notifications are posted.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle
type PostingStyle uint

const (
	PostWhenIdle PostingStyle = 1
	PostASAP PostingStyle = 2
	PostNow PostingStyle = 3
)

// NumberFormatterBehavior - These constants specify the behavior of a number formatter. These constants are returned by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Behavior
type NumberFormatterBehavior uint

const (
	NumberFormatterBehaviorDefault NumberFormatterBehavior = 0
	NumberFormatterBehavior10_0 NumberFormatterBehavior = 1000
	NumberFormatterBehavior10_4 NumberFormatterBehavior = 1040
)

// NumberFormatterPadPosition - These constants are used to specify how numbers should be padded. These constants are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/PadPosition
type NumberFormatterPadPosition uint

// NumberFormatterRoundingMode - These constants are used to specify how numbers should be rounded. These constants are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum
type NumberFormatterRoundingMode uint

// NumberFormatterStyle - The predefined number format styles used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style
type NumberFormatterStyle uint

const (
	NumberFormatterOrdinalStyle NumberFormatterStyle = 0
	NumberFormatterCurrencyISOCodeStyle NumberFormatterStyle = 1
	NumberFormatterCurrencyPluralStyle NumberFormatterStyle = 2
	NumberFormatterCurrencyAccountingStyle NumberFormatterStyle = 3
)

// OperationQueuePriority - These constants let you prioritize the order in which operations execute.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum
type OperationQueuePriority int

const (
	OperationQueuePriorityVeryLow OperationQueuePriority = -8
	OperationQueuePriorityLow OperationQueuePriority = -4
	OperationQueuePriorityNormal OperationQueuePriority = 0
	OperationQueuePriorityHigh OperationQueuePriority = 4
	OperationQueuePriorityVeryHigh OperationQueuePriority = 8
)

// PersonNameComponentsFormatterOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Options
type PersonNameComponentsFormatterOptions uint

const (
	PersonNameComponentsFormatterPhonetic PersonNameComponentsFormatterOptions = 2
)

// PersonNameComponentsFormatterStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Style-swift.enum
type PersonNameComponentsFormatterStyle uint

const (
	// PersonNameComponentsFormatterStyleAbbreviated - The maximally abbreviated form of a name. See “ ” for details about its specific behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Style-swift.enum/abbreviated
	PersonNameComponentsFormatterStyleAbbreviated PersonNameComponentsFormatterStyle = 4
	// PersonNameComponentsFormatterStyleDefault - The form with minimally necessary features for differentiation in a casual setting. See “ ” for details about its specific behavior. Equivalent to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Style-swift.enum/default
	PersonNameComponentsFormatterStyleDefault PersonNameComponentsFormatterStyle = 0
	// PersonNameComponentsFormatterStyleLong - The fully qualified form complete with all known components. See “ ” for details about its specific behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Style-swift.enum/long
	PersonNameComponentsFormatterStyleLong PersonNameComponentsFormatterStyle = 3
	// PersonNameComponentsFormatterStyleMedium - Equivalent to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Style-swift.enum/medium
	PersonNameComponentsFormatterStyleMedium PersonNameComponentsFormatterStyle = 2
	// PersonNameComponentsFormatterStyleShort - The shortened form appropriate for display in space-constrained settings, contingent on user preferences and language defaults. See “ ” for details about its specific behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Style-swift.enum/short
	PersonNameComponentsFormatterStyleShort PersonNameComponentsFormatterStyle = 1
)

// TaskTerminationReason - Constants that specify the termination reason values that the system returns.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/TerminationReason-swift.enum
type TaskTerminationReason uint

const (
	TaskTerminationReasonExit TaskTerminationReason = 1
	TaskTerminationReasonUncaughtSignal TaskTerminationReason = 2
)

// ActivityOptions - Option flags used with 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions
type ActivityOptions uint

const (
	// ActivityLatencyCritical - A flag to indicate the activity requires the highest amount of timer and I/O precision available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/latencyCritical
	ActivityLatencyCritical ActivityOptions = 0
	// ActivityUserInitiated - A flag to indicate the app is performing a user-requested action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/userInitiated
	ActivityUserInitiated ActivityOptions = 0
)

// ProcessInfoThermalState - Values used to indicate the system’s thermal state.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum
type ProcessInfoThermalState uint

const (
	ProcessInfoThermalStateNominal ProcessInfoThermalState = 0
	ProcessInfoThermalStateFair ProcessInfoThermalState = 1
	ProcessInfoThermalStateSerious ProcessInfoThermalState = 2
	ProcessInfoThermalStateCritical ProcessInfoThermalState = 3
)

// PropertyListMutabilityOptions - These constants specify mutability options in property lists.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/MutabilityOptions
type PropertyListMutabilityOptions uint

// PropertyListFormat - These constants are used to specify a property list serialization format.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat
type PropertyListFormat uint

// QualityOfService - Constants that indicate the nature and importance of work to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService
type QualityOfService int

const (
	QualityOfServiceUserInteractive QualityOfService = 33
	QualityOfServiceUserInitiated QualityOfService = 25
	QualityOfServiceUtility QualityOfService = 17
	QualityOfServiceBackground QualityOfService = 9
	QualityOfServiceDefault QualityOfService = -1
)

// RelativeDateTimeFormatterStyle - A type that represents the style to use when formatting relative dates, such as “1 week ago” or “last week”.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/DateTimeStyle-swift.enum
type RelativeDateTimeFormatterStyle uint

const (
	RelativeDateTimeFormatterStyleNumeric RelativeDateTimeFormatterStyle = 0
	RelativeDateTimeFormatterStyleNamed RelativeDateTimeFormatterStyle = 1
)

// RelativeDateTimeFormatterUnitsStyle - A type that represents the style to use when formatting the units of relative dates.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/UnitsStyle-swift.enum
type RelativeDateTimeFormatterUnitsStyle uint

const (
	RelativeDateTimeFormatterUnitsStyleFull RelativeDateTimeFormatterUnitsStyle = 0
	RelativeDateTimeFormatterUnitsStyleSpellOut RelativeDateTimeFormatterUnitsStyle = 1
	RelativeDateTimeFormatterUnitsStyleShort RelativeDateTimeFormatterUnitsStyle = 2
	RelativeDateTimeFormatterUnitsStyleAbbreviated RelativeDateTimeFormatterUnitsStyle = 3
)

// StreamEvent - Describes the constants that may be sent to the delegate as a bit field in the second parameter of 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Event
type StreamEvent uint

const (
	// StreamEventErrorOccurred - An error has occurred on the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Event/errorOccurred
	StreamEventErrorOccurred StreamEvent = 8
	// StreamEventHasSpaceAvailable - The stream can accept bytes for writing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Event/hasSpaceAvailable
	StreamEventHasSpaceAvailable StreamEvent = 4
)

// StreamStatus - The type declared for the constants listed in 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/Status
type StreamStatus uint

const (
	StreamStatusNotOpen StreamStatus = 0
	StreamStatusOpening StreamStatus = 1
	StreamStatusOpen StreamStatus = 2
	StreamStatusReading StreamStatus = 3
	StreamStatusWriting StreamStatus = 4
	StreamStatusAtEnd StreamStatus = 5
	StreamStatusClosed StreamStatus = 6
	StreamStatusError StreamStatus = 7
)

// URLCacheStoragePolicy - These constants specify the caching strategy used by an 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/StoragePolicy
type URLCacheStoragePolicy uint

const (
	URLCacheStorageAllowed URLCacheStoragePolicy = 0
	URLCacheStorageAllowedInMemoryOnly URLCacheStoragePolicy = 1
	URLCacheStorageNotAllowed URLCacheStoragePolicy = 2
)

// URLCredentialPersistence - Constants that specify how long the credential will be kept.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum
type URLCredentialPersistence uint

const (
	// URLCredentialPersistenceForSession - The credential should be stored only for this session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum/forSession
	URLCredentialPersistenceForSession URLCredentialPersistence = 1
)

// URLSessionAuthChallengeDisposition - Constants passed by session or task delegates to the provided continuation block in response to an authentication challenge.
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
	// URLSessionAuthChallengeUseCredential - Use the specified credential, which may be  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/useCredential
	URLSessionAuthChallengeUseCredential URLSessionAuthChallengeDisposition = 0
)

// URLSessionMultipathServiceType - Constants that specify the type of service that Multipath TCP uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum
type URLSessionMultipathServiceType uint

const (
	URLSessionMultipathServiceTypeNone URLSessionMultipathServiceType = 0
	URLSessionMultipathServiceTypeHandover URLSessionMultipathServiceType = 1
	URLSessionMultipathServiceTypeInteractive URLSessionMultipathServiceType = 2
	URLSessionMultipathServiceTypeAggregate URLSessionMultipathServiceType = 3
)

// URLSessionTaskState - Constants for determining the current state of a task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum
type URLSessionTaskState uint

const (
	URLSessionTaskStateRunning URLSessionTaskState = 0
	URLSessionTaskStateSuspended URLSessionTaskState = 1
	URLSessionTaskStateCanceling URLSessionTaskState = 2
	URLSessionTaskStateCompleted URLSessionTaskState = 3
)


