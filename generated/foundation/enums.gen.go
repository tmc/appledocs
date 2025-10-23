// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
// AlignmentOptions - Values representing alignment operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions
type AlignmentOptions uint

const (
	AlignMinXInward AlignmentOptions = 1
	AlignMinYInward AlignmentOptions = 2
	AlignMaxXInward AlignmentOptions = 4
	AlignMaxYInward AlignmentOptions = 8
	AlignWidthInward AlignmentOptions = 16
	AlignHeightInward AlignmentOptions = 32
	AlignMinXOutward AlignmentOptions = 256
	AlignMinYOutward AlignmentOptions = 512
	AlignMaxXOutward AlignmentOptions = 1024
	AlignMaxYOutward AlignmentOptions = 2048
	AlignWidthOutward AlignmentOptions = 4096
	AlignHeightOutward AlignmentOptions = 8192
	AlignMinXNearest AlignmentOptions = 65536
	AlignMinYNearest AlignmentOptions = 131072
	AlignMaxXNearest AlignmentOptions = 262144
	AlignMaxYNearest AlignmentOptions = 524288
	AlignWidthNearest AlignmentOptions = 1048576
	AlignHeightNearest AlignmentOptions = 2097152
	AlignRectFlipped AlignmentOptions = -9223372036854775808
)

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
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/includesDirectoriesPostOrder
	DirectoryEnumerationIncludesDirectoriesPostOrder DirectoryEnumerationOptions = 5
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/producesRelativePathURLs
	DirectoryEnumerationProducesRelativePathURLs DirectoryEnumerationOptions = 6
	// DirectoryEnumerationSkipsHiddenFiles - An option to skip hidden files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsHiddenFiles
	DirectoryEnumerationSkipsHiddenFiles DirectoryEnumerationOptions = 4
	// DirectoryEnumerationSkipsPackageDescendants - An option to treat packages like files and not descend into their contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsPackageDescendants
	DirectoryEnumerationSkipsPackageDescendants DirectoryEnumerationOptions = 2
	// DirectoryEnumerationSkipsSubdirectoryDescendants - An option to perform a shallow enumeration that doesn’t descend into directories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsSubdirectoryDescendants
	DirectoryEnumerationSkipsSubdirectoryDescendants DirectoryEnumerationOptions = 1
)

// FileManagerItemReplacementOptions - Options for specifying the behavior of file replacement operations.
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
	FileManagerItemReplacementWithoutDeletingBackupItem FileManagerItemReplacementOptions = 2
)

// SearchPathDirectory - The location of significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
type SearchPathDirectory uint

const (
	// AdminApplicationDirectory - System and network administration applications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/adminApplicationDirectory
	AdminApplicationDirectory SearchPathDirectory = 4
	// AllApplicationsDirectory - All directories where applications can be stored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/allApplicationsDirectory
	AllApplicationsDirectory SearchPathDirectory = 100
	// AllLibrariesDirectory - All directories where resources can be stored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/allLibrariesDirectory
	AllLibrariesDirectory SearchPathDirectory = 101
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
	// AutosavedInformationDirectory - The user’s autosaved documents ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/autosavedInformationDirectory
	AutosavedInformationDirectory SearchPathDirectory = 11
	// CachesDirectory - Discardable cache files ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/cachesDirectory
	CachesDirectory SearchPathDirectory = 13
	// CoreServiceDirectory - Core services ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/coreServiceDirectory
	CoreServiceDirectory SearchPathDirectory = 10
	// DemoApplicationDirectory - Unsupported applications and demonstration versions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/demoApplicationDirectory
	DemoApplicationDirectory SearchPathDirectory = 2
	// DesktopDirectory - The user’s desktop directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/desktopDirectory
	DesktopDirectory SearchPathDirectory = 12
	// DeveloperApplicationDirectory - Developer applications ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/developerApplicationDirectory
	DeveloperApplicationDirectory SearchPathDirectory = 3
	// DeveloperDirectory - Developer resources ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/developerDirectory
	DeveloperDirectory SearchPathDirectory = 6
	// DocumentDirectory - Document directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentDirectory
	DocumentDirectory SearchPathDirectory = 9
	// DocumentationDirectory - Documentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentationDirectory
	DocumentationDirectory SearchPathDirectory = 8
	// DownloadsDirectory - The user’s downloads directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/downloadsDirectory
	DownloadsDirectory SearchPathDirectory = 15
	// InputMethodsDirectory - Input Methods  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/inputMethodsDirectory
	InputMethodsDirectory SearchPathDirectory = 16
	// ItemReplacementDirectory - The constant used to create a temporary directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/itemReplacementDirectory
	ItemReplacementDirectory SearchPathDirectory = 24
	// LibraryDirectory - Various user-visible documentation, support, and configuration files ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/libraryDirectory
	LibraryDirectory SearchPathDirectory = 5
	// MoviesDirectory - The user’s Movies directory  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/moviesDirectory
	MoviesDirectory SearchPathDirectory = 17
	// MusicDirectory - The user’s Music directory ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/musicDirectory
	MusicDirectory SearchPathDirectory = 18
	// PicturesDirectory - The user’s Pictures directory ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/picturesDirectory
	PicturesDirectory SearchPathDirectory = 19
	// PreferencePanesDirectory - The PreferencePanes directory for use with System Preferences ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/preferencePanesDirectory
	PreferencePanesDirectory SearchPathDirectory = 22
	// PrinterDescriptionDirectory - The system’s PPDs directory ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/printerDescriptionDirectory
	PrinterDescriptionDirectory SearchPathDirectory = 20
	// SharedPublicDirectory - The user’s Public sharing directory ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/sharedPublicDirectory
	SharedPublicDirectory SearchPathDirectory = 21
	// TrashDirectory - The trash directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/trashDirectory
	TrashDirectory SearchPathDirectory = 102
	// UserDirectory - User home directories ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/userDirectory
	UserDirectory SearchPathDirectory = 7
)

// SearchPathDomainMask - Domain constants specifying base locations to use when you search for significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask
type SearchPathDomainMask uint

const (
	// AllDomainsMask - All domains.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/allDomainsMask
	AllDomainsMask SearchPathDomainMask = 65535
	// LocalDomainMask - The place to install items available to everyone on this machine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/localDomainMask
	LocalDomainMask SearchPathDomainMask = 2
	// NetworkDomainMask - The place to install items available on the network ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/networkDomainMask
	NetworkDomainMask SearchPathDomainMask = 4
	// SystemDomainMask - A directory for system files provided by Apple ( ) .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/systemDomainMask
	SystemDomainMask SearchPathDomainMask = 8
	// UserDomainMask - The user’s home directory—the place to install user’s personal items ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/userDomainMask
	UserDomainMask SearchPathDomainMask = 1
)

// URLRelationship - Constants indicating the relationship between a directory and an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship
type URLRelationship uint

const (
	// URLRelationshipContains - The directory contains the specified item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship/contains
	URLRelationshipContains URLRelationship = 0
	// URLRelationshipOther - The directory does not contain the item and is not the same as the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship/other
	URLRelationshipOther URLRelationship = 2
	// URLRelationshipSame - The directory and the item are the same. This relationship occurs when the value of the   is the same for the directory and item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship/same
	URLRelationshipSame URLRelationship = 1
)

// FileManagerUnmountOptions - Options that specify the behavior of an unmount operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions
type FileManagerUnmountOptions uint

const (
	// FileManagerUnmountAllPartitionsAndEjectDisk - Specifies that all partitions on an unmountable disk should be unmounted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions/allPartitionsAndEjectDisk
	FileManagerUnmountAllPartitionsAndEjectDisk FileManagerUnmountOptions = 1
	// FileManagerUnmountWithoutUI - Specifies that no UI should accompany the unmount operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions/withoutUI
	FileManagerUnmountWithoutUI FileManagerUnmountOptions = 2
)

// VolumeEnumerationOptions - Options for enumerating mounted volumes with the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions
type VolumeEnumerationOptions uint

const (
	// VolumeEnumerationProduceFileReferenceURLs - The enumeration produces file reference URLs rather than path-based URLs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions/produceFileReferenceURLs
	VolumeEnumerationProduceFileReferenceURLs VolumeEnumerationOptions = 4
	// VolumeEnumerationSkipHiddenVolumes - The enumeration skips hidden volumes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions/skipHiddenVolumes
	VolumeEnumerationSkipHiddenVolumes VolumeEnumerationOptions = 2
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
	// InlinePresentationIntentEmphasized - An intent that represents an emphasized presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/emphasized
	InlinePresentationIntentEmphasized InlinePresentationIntent = 1
	// InlinePresentationIntentInlineHTML - An intent that represents an inline HTML presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/inlineHTML
	InlinePresentationIntentInlineHTML InlinePresentationIntent = 256
	// InlinePresentationIntentStronglyEmphasized - An intent that represents a strongly emphasized presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/stronglyEmphasized
	InlinePresentationIntentStronglyEmphasized InlinePresentationIntent = 2
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

// AttributedStringFormattingOptions - Options to use when creating an attributed string from a format string and variable list of arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingOptions
type AttributedStringFormattingOptions uint

const (
	// AttributedStringFormattingApplyReplacementIndexAttribute - An option to apply to the replaced portions of text in a format string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingOptions/NSAttributedStringFormattingApplyReplacementIndexAttribute
	AttributedStringFormattingApplyReplacementIndexAttribute AttributedStringFormattingOptions = 1
	// AttributedStringFormattingInsertArgumentAttributesWithoutMerging - An option to replace the attributes in a substituted string with those of the provided attributed string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingOptions/NSAttributedStringFormattingInsertArgumentAttributesWithoutMerging
	AttributedStringFormattingInsertArgumentAttributesWithoutMerging AttributedStringFormattingOptions = 0
)

// AttributedStringMarkdownParsingFailurePolicy - A type that represents policies for handling parsing failures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingFailurePolicy
type AttributedStringMarkdownParsingFailurePolicy int

const (
	// AttributedStringMarkdownParsingFailureReturnError - A policy to return an error from the initializer if parsing fails.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingFailurePolicy/NSAttributedStringMarkdownParsingFailureReturnError
	AttributedStringMarkdownParsingFailureReturnError AttributedStringMarkdownParsingFailurePolicy = 0
	// AttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible - A policy to return a partially parsed string, if possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingFailurePolicy/NSAttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible
	AttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible AttributedStringMarkdownParsingFailurePolicy = 1
)

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
	BinarySearchingFirstEqual BinarySearchingOptions = 256
	BinarySearchingLastEqual BinarySearchingOptions = 512
	BinarySearchingInsertionIndex BinarySearchingOptions = 1024
)

// CalendarOptions - The options for arithmetic operations involving calendars.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
type CalendarOptions uint

const (
	// CalendarMatchFirst - Specifies that, if there are two or more matching times, the operation should return the first occurrence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchFirst
	CalendarMatchFirst CalendarOptions = 7
	// CalendarMatchLast - Specifies that, if there are two or more matching times, the operation should return the last occurrence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchLast
	CalendarMatchLast CalendarOptions = 8
	// CalendarMatchNextTime - Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given   object, this method uses the   existing value of the missing unit and   preserve the lower units’ values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchNextTime
	CalendarMatchNextTime CalendarOptions = 6
	// CalendarMatchNextTimePreservingSmallerUnits - Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given   object, this method uses the   existing value of the missing unit and preserves the lower units’ values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchNextTimePreservingSmallerUnits
	CalendarMatchNextTimePreservingSmallerUnits CalendarOptions = 5
	// CalendarMatchPreviousTimePreservingSmallerUnits - Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given   object, this method uses the   existing value of the missing unit and preserves the lower units’ values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchPreviousTimePreservingSmallerUnits
	CalendarMatchPreviousTimePreservingSmallerUnits CalendarOptions = 4
	// CalendarMatchStrictly - Specifies that the operation should travel as far forward or backward as necessary looking for a match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchStrictly
	CalendarMatchStrictly CalendarOptions = 2
	// CalendarSearchBackwards - Specifies that the operation should travel backwards to find the previous match before the given date.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/searchBackwards
	CalendarSearchBackwards CalendarOptions = 3
	// CalendarWrapComponents - Specifies that the components specified for an   object should be incremented and wrap around to zero/one on overflow, but should not cause higher units to be incremented.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/wrapComponents
	CalendarWrapComponents CalendarOptions = 1
)

// CalendarUnit - Calendrical units such as year, month, day and hour.
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
	// EraCalendarUnit - Specifies the era unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSEraCalendarUnit
	EraCalendarUnit CalendarUnit = 8
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
	// TimeZoneCalendarUnit - Specifies the time zone of the calendar as an  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSTimeZoneCalendarUnit
	TimeZoneCalendarUnit CalendarUnit = 23
	// WeekCalendarUnit - Specifies the week unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekCalendarUnit
	WeekCalendarUnit CalendarUnit = 15
	// WeekOfMonthCalendarUnit - Specifies the original week of a month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekOfMonthCalendarUnit
	WeekOfMonthCalendarUnit CalendarUnit = 19
	// WeekOfYearCalendarUnit - Specifies the original week of the year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekOfYearCalendarUnit
	WeekOfYearCalendarUnit CalendarUnit = 20
	// WeekdayCalendarUnit - Specifies the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekdayCalendarUnit
	WeekdayCalendarUnit CalendarUnit = 16
	// WeekdayOrdinalCalendarUnit - Specifies the ordinal weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekdayOrdinalCalendarUnit
	WeekdayOrdinalCalendarUnit CalendarUnit = 17
	// YearCalendarUnit - Specifies the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSYearCalendarUnit
	YearCalendarUnit CalendarUnit = 9
	// YearForWeekOfYearCalendarUnit - Specifies the year when the calendar is being interpreted as a week-based calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSYearForWeekOfYearCalendarUnit
	YearForWeekOfYearCalendarUnit CalendarUnit = 21
	// CalendarUnitCalendar - Identifier for the calendar of a date components object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/calendar
	CalendarUnitCalendar CalendarUnit = 6
	// CalendarUnitDay - Identifier for the day unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/day
	CalendarUnitDay CalendarUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/dayOfYear
	CalendarUnitDayOfYear CalendarUnit = 5
	// CalendarUnitEra - Identifier for the era unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/era
	CalendarUnitEra CalendarUnit = 0
	// CalendarUnitHour - Identifier for the hour unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/hour
	CalendarUnitHour CalendarUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/isLeapMonth
	CalendarUnitIsLeapMonth CalendarUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/isRepeatedDay
	CalendarUnitIsRepeatedDay CalendarUnit = 0
	// CalendarUnitMinute - Identifier for the minute unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/minute
	CalendarUnitMinute CalendarUnit = 0
	// CalendarUnitMonth - Identifier for the month unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/month
	CalendarUnitMonth CalendarUnit = 0
	// CalendarUnitNanosecond - Identifier for the nanosecond unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/nanosecond
	CalendarUnitNanosecond CalendarUnit = 4
	// CalendarUnitQuarter - Identifier for the quarter of the calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/quarter
	CalendarUnitQuarter CalendarUnit = 0
	// CalendarUnitSecond - Identifier for the second unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/second
	CalendarUnitSecond CalendarUnit = 0
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
	// CalendarUnitWeekdayOrdinal - Identifier for the ordinal weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekdayOrdinal
	CalendarUnitWeekdayOrdinal CalendarUnit = 0
	// CalendarUnitYear - Identifier for the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/year
	CalendarUnitYear CalendarUnit = 0
	// CalendarUnitYearForWeekOfYear - Identifier for the week-counting year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/yearForWeekOfYear
	CalendarUnitYearForWeekOfYear CalendarUnit = 3
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
	// FileCoordinatorReadingForUploading - Specify this content when reading an item for the purpose of uploading its contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/forUploading
	FileCoordinatorReadingForUploading FileCoordinatorReadingOptions = 4
	// FileCoordinatorReadingImmediatelyAvailableMetadataOnly - Specify this constant if you want to read an item’s metadata without triggering a download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/immediatelyAvailableMetadataOnly
	FileCoordinatorReadingImmediatelyAvailableMetadataOnly FileCoordinatorReadingOptions = 3
	// FileCoordinatorReadingResolvesSymbolicLink - Specify this constant if you want an item that might be a symbolic link to resolve to the file pointed to by that link (instead of to the link itself). When you use this option, the system provides the resolved URL to the accessor block in place of the original URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/resolvesSymbolicLink
	FileCoordinatorReadingResolvesSymbolicLink FileCoordinatorReadingOptions = 2
	// FileCoordinatorReadingWithoutChanges - Specify this constant if your code does not need other objects to save changes first. If you do   specify this constant, the   method of relevant file presenters is called before your code reads the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/withoutChanges
	FileCoordinatorReadingWithoutChanges FileCoordinatorReadingOptions = 1
)

// FileCoordinatorWritingOptions - Options to use when changing the contents or attributes of a file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
type FileCoordinatorWritingOptions uint

const (
	// FileCoordinatorWritingContentIndependentMetadataOnly - Select this option when writing to change the file’s metadata only and not its contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/contentIndependentMetadataOnly
	FileCoordinatorWritingContentIndependentMetadataOnly FileCoordinatorWritingOptions = 9
	// FileCoordinatorWritingForDeleting - When this constant is specified, the file coordinator calls the   or   method of relevant file presenters to give them a chance to make adjustments before the item is deleted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forDeleting
	FileCoordinatorWritingForDeleting FileCoordinatorWritingOptions = 1
	// FileCoordinatorWritingForMoving - When specified for a directory item, the file coordinator waits for already running read and write operations of the directory’s contents, which were themselves initiated through a file coordinator, to finish before moving the directory. Queued, but not executing, read and write operations on the directory’s contents wait until the move operation finishes. This option has no effect on files. You can safely use it when moving file-system items without checking to see whether those items are files or directories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forMoving
	FileCoordinatorWritingForMoving FileCoordinatorWritingOptions = 2
)

// FileManagerResumeSyncBehavior - The behaviors the file manager can apply to resolve conflicts when resuming a sync.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior
type FileManagerResumeSyncBehavior uint

const (
	// FileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict - Resumes sync by first uploading the local version of the file, failing if the provider detects a conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior/afterUploadWithFailOnConflict
	FileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict FileManagerResumeSyncBehavior = 1
	// FileManagerResumeSyncBehaviorDropLocalChanges - Resumes synchronizing by overwriting any local changes with the remote version of the file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior/dropLocalChanges
	FileManagerResumeSyncBehaviorDropLocalChanges FileManagerResumeSyncBehavior = 2
	// FileManagerResumeSyncBehaviorPreserveLocalChanges - Resumes synchronizing by uploading the local version of the file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior/preserveLocalChanges
	FileManagerResumeSyncBehaviorPreserveLocalChanges FileManagerResumeSyncBehavior = 0
)

// FileManagerSupportedSyncControls - An option set of the sync controls available for an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls
type FileManagerSupportedSyncControls uint

const (
	// FileManagerSupportedSyncControlsFailUploadOnConflict - The file provider supports failing an upload if the local and server versions conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls/failUploadOnConflict
	FileManagerSupportedSyncControlsFailUploadOnConflict FileManagerSupportedSyncControls = 2
	// FileManagerSupportedSyncControlsPauseSync - The file provider supports pausing the sync on the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls/pauseSync
	FileManagerSupportedSyncControlsPauseSync FileManagerSupportedSyncControls = 1
)

// FileManagerUploadLocalVersionConflictPolicy - The policies the file manager can apply to resolve conflicts when uploading a local version of a file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy
type FileManagerUploadLocalVersionConflictPolicy uint

const (
	// FileManagerUploadConflictPolicyDefault - Resolves the conflict using the policy defined by the file provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy/conflictPolicyDefault
	FileManagerUploadConflictPolicyDefault FileManagerUploadLocalVersionConflictPolicy = 0
	// FileManagerUploadConflictPolicyFailOnConflict - Resolves the conflict by causing the upload to fail.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy/conflictPolicyFailOnConflict
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
	// FileVersionReplacingByMoving - An option to perform replacing by moving a file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/ReplacingOptions/byMoving
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
	GrammaticalDefinitenessNotSet GrammaticalDefiniteness = 0
	GrammaticalDefinitenessIndefinite GrammaticalDefiniteness = 1
	GrammaticalDefinitenessDefinite GrammaticalDefiniteness = 2
)

// GrammaticalDetermination enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination
type GrammaticalDetermination uint

const (
	GrammaticalDeterminationNotSet GrammaticalDetermination = 0
	GrammaticalDeterminationIndependent GrammaticalDetermination = 1
	GrammaticalDeterminationDependent GrammaticalDetermination = 2
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

// ItemProviderErrorCode - The error codes that describe problems with consuming data from an item provider.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode
type ItemProviderErrorCode int

const (
	// ItemProviderItemUnavailableError - An error code indicating that the requested data was unavailable from an item provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode/itemUnavailableError
	ItemProviderItemUnavailableError ItemProviderErrorCode = -1000
	// ItemProviderUnavailableCoercionError - An error code indicating that the requested data type coercion is unavailable from an item provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode/unavailableCoercionError
	ItemProviderUnavailableCoercionError ItemProviderErrorCode = -1099
	// ItemProviderUnexpectedValueClassError - An error code indicating that type coercion to the requested class failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode/unexpectedValueClassError
	ItemProviderUnexpectedValueClassError ItemProviderErrorCode = -1100
	// ItemProviderUnknownError - An error code indicating an unknown error with consuming data from an item provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode/unknownError
	ItemProviderUnknownError ItemProviderErrorCode = -1
)

// UIPreferredPresentationStyle - The presentation styles that determine how a view shows an item provider’s data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/PreferredPresentationStyle-swift.enum
type UIPreferredPresentationStyle uint

// ItemProviderFileOptions - Data-access specifications that declare how to handle items.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions
type ItemProviderFileOptions uint

const (
	// ItemProviderFileOptionOpenInPlace - A data-access specification declaring that items should open in place, rather than being copied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions/openInPlace
	ItemProviderFileOptionOpenInPlace ItemProviderFileOptions = 1
)

// ItemProviderRepresentationVisibility - Specifications that control which categories of processes can see an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility
type ItemProviderRepresentationVisibility uint

const (
	// ItemProviderRepresentationVisibilityAll - A representation visibility specification conferring item visibility to all processes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility/all
	ItemProviderRepresentationVisibilityAll ItemProviderRepresentationVisibility = 0
	// ItemProviderRepresentationVisibilityGroup - A representation visibility specification confining item visibility to the app’s app group.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility/group
	ItemProviderRepresentationVisibilityGroup ItemProviderRepresentationVisibility = 2
	// ItemProviderRepresentationVisibilityOwnProcess - A representation visibility specification confining item visibility to the app that is the source of the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility/ownProcess
	ItemProviderRepresentationVisibilityOwnProcess ItemProviderRepresentationVisibility = 3
	// ItemProviderRepresentationVisibilityTeam - A representation visibility specification confining item visibility to processes created by the app’s development team.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility/team
	ItemProviderRepresentationVisibilityTeam ItemProviderRepresentationVisibility = 1
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
	KeyValueUnionSetMutation KeyValueSetMutationKind = 1
	KeyValueMinusSetMutation KeyValueSetMutationKind = 2
	KeyValueIntersectSetMutation KeyValueSetMutationKind = 3
	KeyValueSetSetMutation KeyValueSetMutationKind = 4
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

// LocaleLanguageDirection - The directions that a language may take across a page of text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection
type LocaleLanguageDirection uint

// MachPortOptions - Used to remove access rights to a mach port when the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options
type MachPortOptions uint

const (
	// MachPortDeallocateReceiveRight - Remove a receive right when the   object is invalidated or destroyed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options/deallocateReceiveRight
	MachPortDeallocateReceiveRight MachPortOptions = 2
	// MachPortDeallocateSendRight - Deallocate a send right when the   object is invalidated or destroyed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options/deallocateSendRight
	MachPortDeallocateSendRight MachPortOptions = 1
	// MachPortDeallocateNone - Do not remove any send or receive rights.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPortOptions/NSMachPortDeallocateNone
	MachPortDeallocateNone MachPortOptions = 0
)

// OrderedCollectionDifferenceCalculationOptions - Constants that specify the options to use when creating an ordered collection difference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifferenceCalculationOptions
type OrderedCollectionDifferenceCalculationOptions uint

const (
	OrderedCollectionDifferenceCalculationOmitInsertedObjects OrderedCollectionDifferenceCalculationOptions = 1
	OrderedCollectionDifferenceCalculationOmitRemovedObjects OrderedCollectionDifferenceCalculationOptions = 2
	OrderedCollectionDifferenceCalculationInferMoves OrderedCollectionDifferenceCalculationOptions = 4
)

// PointerFunctionsOptions - Defines the memory and personality options for an 
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
	PresentationIntentKindParagraph PresentationIntentKind = 0
	PresentationIntentKindHeader PresentationIntentKind = 1
	PresentationIntentKindOrderedList PresentationIntentKind = 2
	PresentationIntentKindUnorderedList PresentationIntentKind = 3
	PresentationIntentKindListItem PresentationIntentKind = 4
	PresentationIntentKindCodeBlock PresentationIntentKind = 5
	PresentationIntentKindBlockQuote PresentationIntentKind = 6
	PresentationIntentKindThematicBreak PresentationIntentKind = 7
	PresentationIntentKindTable PresentationIntentKind = 8
	PresentationIntentKindTableHeaderRow PresentationIntentKind = 9
	PresentationIntentKindTableRow PresentationIntentKind = 10
	PresentationIntentKindTableCell PresentationIntentKind = 11
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
	// SaveOptionsAsk - Indicates the user should be asked before saving any modified documents on closing. When no option is specified, this is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions/ask
	SaveOptionsAsk SaveOptions = 2
	// SaveOptionsNo - Indicates a modified document should not be saved on closing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions/no
	SaveOptionsNo SaveOptions = 1
	// SaveOptionsYes - Indicates a modified document should be saved on closing without asking the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions/yes
	SaveOptionsYes SaveOptions = 0
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

// TimeZoneNameStyle - Constants you use to specify a style when presenting time zone names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle
type TimeZoneNameStyle uint

const (
	// TimeZoneNameStyleDaylightSaving - Specifies a daylight saving name style. For example, “Central Daylight Time” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/daylightSaving
	TimeZoneNameStyleDaylightSaving TimeZoneNameStyle = 2
	// TimeZoneNameStyleGeneric - Specifies a generic name style. For example, “Central Time” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/generic
	TimeZoneNameStyleGeneric TimeZoneNameStyle = 4
	// TimeZoneNameStyleShortDaylightSaving - Specifies a short daylight saving name style.  For example, “CDT” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/shortDaylightSaving
	TimeZoneNameStyleShortDaylightSaving TimeZoneNameStyle = 3
	// TimeZoneNameStyleShortGeneric - Specifies a generic time zone name. For example, “CT” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/shortGeneric
	TimeZoneNameStyleShortGeneric TimeZoneNameStyle = 5
	// TimeZoneNameStyleShortStandard - Specifies a short name style. For example, “CST” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/shortStandard
	TimeZoneNameStyleShortStandard TimeZoneNameStyle = 1
	// TimeZoneNameStyleStandard - Specifies a standard name style. For example, “Central Standard Time” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/standard
	TimeZoneNameStyleStandard TimeZoneNameStyle = 0
)

// URLBookmarkCreationOptions - Options used when creating bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions
type URLBookmarkCreationOptions uint

const (
	// URLBookmarkCreationMinimalBookmark - Specifies that when creating a bookmark, it includes minimal information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/minimalBookmark
	URLBookmarkCreationMinimalBookmark URLBookmarkCreationOptions = 512
	// URLBookmarkCreationPreferFileIDResolution - Specifies that when creating a bookmark, upon resolution, its embedded file ID takes precedence over other sources of information (file system path, for example) when there’s a conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/preferFileIDResolution
	URLBookmarkCreationPreferFileIDResolution URLBookmarkCreationOptions = 0
	// URLBookmarkCreationSecurityScopeAllowOnlyReadAccess - Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read-only access to a file-system resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/securityScopeAllowOnlyReadAccess
	URLBookmarkCreationSecurityScopeAllowOnlyReadAccess URLBookmarkCreationOptions = 1026
	// URLBookmarkCreationSuitableForBookmarkFile - Specifies that the bookmark data includes the required properties for creating Finder alias files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/suitableForBookmarkFile
	URLBookmarkCreationSuitableForBookmarkFile URLBookmarkCreationOptions = 1024
	// URLBookmarkCreationWithSecurityScope - Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read/write access to a file-system resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/withSecurityScope
	URLBookmarkCreationWithSecurityScope URLBookmarkCreationOptions = 1025
	// URLBookmarkCreationWithoutImplicitSecurityScope - Prevents inclusion of a bookmark’s implicit ephemeral security scope, when creating one without security scope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/withoutImplicitSecurityScope
	URLBookmarkCreationWithoutImplicitSecurityScope URLBookmarkCreationOptions = 1027
)

// URLBookmarkResolutionOptions - Options used when resolving bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions
type URLBookmarkResolutionOptions uint

const (
	// URLBookmarkResolutionWithSecurityScope - Specifies that the security scope, applied to the bookmark when it was created, should be used during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withSecurityScope
	URLBookmarkResolutionWithSecurityScope URLBookmarkResolutionOptions = 513
	// URLBookmarkResolutionWithoutImplicitStartAccessing - A property that specifies that resolution doesn’t implicitly start accessing the ephemeral security-scoped resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutImplicitStartAccessing
	URLBookmarkResolutionWithoutImplicitStartAccessing URLBookmarkResolutionOptions = 514
	// URLBookmarkResolutionWithoutMounting - Specifies that no volume should be mounted during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutMounting
	URLBookmarkResolutionWithoutMounting URLBookmarkResolutionOptions = 512
	// URLBookmarkResolutionWithoutUI - Specifies that no UI feedback should accompany resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutUI
	URLBookmarkResolutionWithoutUI URLBookmarkResolutionOptions = 256
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

// URLRequestCachePolicy - The constants used to specify interaction with the cached responses.
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
	URLRequestReloadIgnoringLocalAndRemoteCacheData URLRequestCachePolicy = 4
	// URLRequestReloadIgnoringLocalCacheData - The URL load should be loaded only from the originating source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/reloadIgnoringLocalCacheData
	URLRequestReloadIgnoringLocalCacheData URLRequestCachePolicy = 1
	// URLRequestReloadRevalidatingCacheData - Use cache data if the origin source can validate it; otherwise, load from the origin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/reloadRevalidatingCacheData
	URLRequestReloadRevalidatingCacheData URLRequestCachePolicy = 5
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
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType/NSURLSessionWebSocketMessageTypeData
	URLSessionWebSocketMessageTypeData URLSessionWebSocketMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType/NSURLSessionWebSocketMessageTypeString
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
	// NetServicesActivityInProgress - The net service cannot process the request at this time. No additional information about the network state is known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/activityInProgress
	NetServicesActivityInProgress NetServicesError = -72003
	// NetServicesBadArgumentError - An invalid argument was used when creating the   object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/badArgumentError
	NetServicesBadArgumentError NetServicesError = -72004
	// NetServicesCancelledError - The client canceled the action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/cancelledError
	NetServicesCancelledError NetServicesError = -72005
	// NetServicesCollisionError - The service could not be published because the name is already in use. The name could be in use locally or on another system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/collisionError
	NetServicesCollisionError NetServicesError = -72001
	// NetServicesInvalidError - The net service was improperly configured.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/invalidError
	NetServicesInvalidError NetServicesError = -72006
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/missingRequiredConfigurationError
	NetServicesMissingRequiredConfigurationError NetServicesError = -72006
	// NetServicesNotFoundError - The service could not be found on the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/notFoundError
	NetServicesNotFoundError NetServicesError = -72002
	// NetServicesTimeoutError - The net service has timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/timeoutError
	NetServicesTimeoutError NetServicesError = -72007
	// NetServicesUnknownError - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/unknownError
	NetServicesUnknownError NetServicesError = -72000
)

// NetServiceOptions - These constants specify options for a network service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options
type NetServiceOptions uint

const (
	// NetServiceListenForConnections - Specifies that a TCP listener should be started for both IPv4 and IPv6 on the port specified by this service. If the listening port can’t be opened, the service calls its delegate’s   method to report the error. The listener supports only TCP connections. If the service’s type does not end with  , publication fails with  . Whenever a client connects to the listening socket, the service calls its delegate’s   method with a pair of   objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options/listenForConnections
	NetServiceListenForConnections NetServiceOptions = 2
	// NetServiceNoAutoRename - Specifies that the network service should not rename itself in the event of a name collision.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options/noAutoRename
	NetServiceNoAutoRename NetServiceOptions = 1
)

// NotificationCoalescing - The constants that specify how notifications are coalesced.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing
type NotificationCoalescing uint

const (
	// NotificationNoCoalescing - Do not coalesce notifications in the queue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing/none
	NotificationNoCoalescing NotificationCoalescing = 0
)

// PostingStyle - The constants that specify when notifications are posted.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle
type PostingStyle uint

const (
	// PostASAP - The notification is posted at the end of the current notification callout or timer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle/asap
	PostASAP PostingStyle = 2
	// PostNow - The notification is posted immediately after coalescing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle/now
	PostNow PostingStyle = 3
	// PostWhenIdle - The notification is posted when the run loop is idle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle/whenIdle
	PostWhenIdle PostingStyle = 1
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
	// OperationQueuePriorityHigh - Operations receive high priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/high
	OperationQueuePriorityHigh OperationQueuePriority = 4
	// OperationQueuePriorityLow - Operations receive low priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/low
	OperationQueuePriorityLow OperationQueuePriority = -4
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
	OperationQueuePriorityVeryLow OperationQueuePriority = -8
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
	// ActivityAnimationTrackingEnabled - A flag to track the activity with an animation signpost interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/animationTrackingEnabled
	ActivityAnimationTrackingEnabled ActivityOptions = 32769
	// ActivityBackground - A flag to indicate the app has initiated some kind of work, but not as the direct result of user request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/background
	ActivityBackground ActivityOptions = 0
	// ActivityLatencyCritical - A flag to indicate the activity requires the highest amount of timer and I/O precision available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/latencyCritical
	ActivityLatencyCritical ActivityOptions = 0
	// ActivitySuddenTerminationDisabled - A flag to prevent sudden termination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/suddenTerminationDisabled
	ActivitySuddenTerminationDisabled ActivityOptions = 16384
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
	ActivityUserInteractive ActivityOptions = 32771
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

// QualityOfService - Constants that indicate the nature and importance of work to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService
type QualityOfService int

const (
	// QualityOfServiceBackground - Used for work that is not user initiated or visible. In general, a user is unaware that this work is even happening. For example, pre-fetching content, search indexing, backups, or syncing of data with external systems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/background
	QualityOfServiceBackground QualityOfService = 9
	// QualityOfServiceDefault - Indicates no explicit quality of service information. Whenever possible, an appropriate quality of service is determined from available sources. Otherwise, some quality of service level between   and   is used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/default
	QualityOfServiceDefault QualityOfService = -1
	// QualityOfServiceUserInitiated - Used for performing work that has been explicitly requested by the user, and for which results must be immediately presented in order to allow for further user interaction. For example, loading an email after a user has selected it in a message list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/userInitiated
	QualityOfServiceUserInitiated QualityOfService = 25
	// QualityOfServiceUserInteractive - Used for work directly involved in providing an interactive UI. For example, processing control events or drawing to the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/userInteractive
	QualityOfServiceUserInteractive QualityOfService = 33
	// QualityOfServiceUtility - Used for performing work which the user is unlikely to be immediately waiting for the results. This work may have been requested by the user or initiated automatically, and often operates at user-visible timescales using a non-modal progress indicator. For example, periodic content updates or bulk file operations, such as media import.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/utility
	QualityOfServiceUtility QualityOfService = 17
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

// URLCredentialPersistence - Constants that specify how long the credential will be kept.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum
type URLCredentialPersistence uint

const (
	URLCredentialPersistenceNone URLCredentialPersistence = 0
	URLCredentialPersistenceForSession URLCredentialPersistence = 1
	URLCredentialPersistencePermanent URLCredentialPersistence = 2
	URLCredentialPersistenceSynchronizable URLCredentialPersistence = 3
)

// URLSessionAuthChallengeDisposition - Constants passed by session or task delegates to the provided continuation block in response to an authentication challenge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition
type URLSessionAuthChallengeDisposition uint

const (
	// URLSessionAuthChallengeRejectProtectionSpace - Reject this challenge, and call the authentication delegate method again with the next authentication protection space. The provided credential parameter is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/rejectProtectionSpace
	URLSessionAuthChallengeRejectProtectionSpace URLSessionAuthChallengeDisposition = 3
)

// URLSessionDelayedRequestDisposition - The action to take on a delayed URL session task.
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

// URLSessionResponseDisposition - Constants indicating how a data or upload session should proceed after receiving the initial headers.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition
type URLSessionResponseDisposition uint

const (
	// URLSessionResponseAllow - Allow the load operation to continue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/allow
	URLSessionResponseAllow URLSessionResponseDisposition = 1
	// URLSessionResponseBecomeDownload - Convert the response for this request to use a  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/becomeDownload
	URLSessionResponseBecomeDownload URLSessionResponseDisposition = 2
	// URLSessionResponseBecomeStream - Convert the response for this request to use a  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/becomeStream
	URLSessionResponseBecomeStream URLSessionResponseDisposition = 3
	// URLSessionResponseCancel - Cancel the load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/cancel
	URLSessionResponseCancel URLSessionResponseDisposition = 0
)

// URLSessionMultipathServiceType - Constants that specify the type of service that Multipath TCP uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum
type URLSessionMultipathServiceType uint

const (
	// URLSessionMultipathServiceTypeAggregate - A service that aggregates the capacities of other Multipath options in an attempt to increase throughput and minimize latency.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/aggregate
	URLSessionMultipathServiceTypeAggregate URLSessionMultipathServiceType = 3
	// URLSessionMultipathServiceTypeHandover - A Multipath TCP service that provides seamless handover between Wi-Fi and cellular in order to preserve the connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/handover
	URLSessionMultipathServiceTypeHandover URLSessionMultipathServiceType = 1
	// URLSessionMultipathServiceTypeInteractive - A service whereby Multipath TCP attempts to use the lowest-latency interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/interactive
	URLSessionMultipathServiceTypeInteractive URLSessionMultipathServiceType = 2
	// URLSessionMultipathServiceTypeNone - The default service type indicating that Multipath TCP should not be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/none
	URLSessionMultipathServiceTypeNone URLSessionMultipathServiceType = 0
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

// XMLDTDNodeKind - The type defined for the constants that specify the kind and subkind of DTD declaration represented by an 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum
type XMLDTDNodeKind uint

const (
	// XMLElementDeclarationAnyKind - Identifies an   element declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/anyDeclaration
	XMLElementDeclarationAnyKind XMLDTDNodeKind = 18
	// XMLAttributeCDATAKind - Identifies an attribute-list declaration with a   (character data) value type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/cdataAttribute
	XMLAttributeCDATAKind XMLDTDNodeKind = 6
	// XMLElementDeclarationElementKind - Identifies a declaration of an element with child elements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/elementDeclaration
	XMLElementDeclarationElementKind XMLDTDNodeKind = 20
	// XMLElementDeclarationEmptyKind - Identifies a declaration ( ) of an empty element.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/emptyDeclaration
	XMLElementDeclarationEmptyKind XMLDTDNodeKind = 17
	// XMLAttributeEntitiesKind - Identifies an attribute-list declaration with an   value type (refers to multiple unparsed entities declared elsewhere in document).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/entitiesAttribute
	XMLAttributeEntitiesKind XMLDTDNodeKind = 11
	// XMLAttributeEntityKind - Identifies an attribute-list declaration with an   value type (refers to unparsed entity declared in document).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/entityAttribute
	XMLAttributeEntityKind XMLDTDNodeKind = 10
	// XMLAttributeEnumerationKind - Identifies an attribute-list declaration with an enumeration value type (list of all possible values).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/enumerationAttribute
	XMLAttributeEnumerationKind XMLDTDNodeKind = 14
	// XMLEntityGeneralKind - Identifies a general entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/general
	XMLEntityGeneralKind XMLDTDNodeKind = 1
	// XMLAttributeIDKind - Identifies an attribute-list declaration with an   value type (per-document unique element name).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/idAttribute
	XMLAttributeIDKind XMLDTDNodeKind = 7
	// XMLAttributeIDRefKind - Identifies an attribute-list declaration with an   value type (refers to element   type).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/idRefAttribute
	XMLAttributeIDRefKind XMLDTDNodeKind = 8
	// XMLAttributeIDRefsKind - Identifies an attribute-list declaration with an   value type (refers to multiple elements of   type).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/idRefsAttribute
	XMLAttributeIDRefsKind XMLDTDNodeKind = 9
	// XMLElementDeclarationMixedKind - Identifies a declaration of an element with mixed content ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/mixedDeclaration
	XMLElementDeclarationMixedKind XMLDTDNodeKind = 19
	// XMLAttributeNMTokenKind - Identifies an attribute-list declaration with a   value type (name token).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/nmTokenAttribute
	XMLAttributeNMTokenKind XMLDTDNodeKind = 12
	// XMLAttributeNMTokensKind - Identifies an attribute-list declaration with a   value type (multiple name tokens)
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/nmTokensAttribute
	XMLAttributeNMTokensKind XMLDTDNodeKind = 13
	// XMLAttributeNotationKind - Identifies an attribute-list declaration with a   value type (name of declared notation).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/notationAttribute
	XMLAttributeNotationKind XMLDTDNodeKind = 15
	// XMLEntityParameterKind - Identifies a parameter entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/parameter
	XMLEntityParameterKind XMLDTDNodeKind = 4
	// XMLEntityParsedKind - Identifies a parsed entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/parsed
	XMLEntityParsedKind XMLDTDNodeKind = 2
	// XMLEntityPredefined - Identifies a predefined entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/predefined
	XMLEntityPredefined XMLDTDNodeKind = 5
	// XMLElementDeclarationUndefinedKind - Identifies an undefined element declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/undefinedDeclaration
	XMLElementDeclarationUndefinedKind XMLDTDNodeKind = 16
	// XMLEntityUnparsedKind - Identifies an unparsed entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/unparsed
	XMLEntityUnparsedKind XMLDTDNodeKind = 3
)

// XMLDocumentContentKind - Type used to define the kind of document content.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/ContentKind
type XMLDocumentContentKind uint

const (
	XMLDocumentXMLKind XMLDocumentContentKind = 0
	XMLDocumentXHTMLKind XMLDocumentContentKind = 1
	XMLDocumentHTMLKind XMLDocumentContentKind = 2
	XMLDocumentTextKind XMLDocumentContentKind = 3
)

// XMLNodeKind enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum
type XMLNodeKind uint

const (
	// XMLDTDKind - Specifies a document-type declaration (DTD) node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/DTDKind
	XMLDTDKind XMLNodeKind = 8
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
	// XMLDocumentKind - Specifies a document node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/document
	XMLDocumentKind XMLNodeKind = 1
	// XMLElementKind - Specifies an element node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/element
	XMLElementKind XMLNodeKind = 2
	// XMLElementDeclarationKind - Specifies an element declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/elementDeclaration
	XMLElementDeclarationKind XMLNodeKind = 11
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
	XMLNotationDeclarationKind XMLNodeKind = 12
	// XMLProcessingInstructionKind - Specifies a processing-instruction node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/processingInstruction
	XMLProcessingInstructionKind XMLNodeKind = 5
	// XMLTextKind - Specifies a text node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/text
	XMLTextKind XMLNodeKind = 7
)

// XMLNodeOptions - These constants are input and output options for all 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options
type XMLNodeOptions uint

const (
	// XMLNodePreserveDTD - Specifies that declarations in a DTD should be preserved until it the DTD is modified. For example, parameter entities are by default expanded; with this option, they are written out as they originally occur in the DTD.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodePreserveDTD
	XMLNodePreserveDTD XMLNodeOptions = 67108864
	// XMLNodePreserveQuotes - Specifies that the quoting style used in the input XML (single or double quotes) be preserved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodePreserveQuotes
	XMLNodePreserveQuotes XMLNodeOptions = 0
	// XMLNodeUseSingleQuotes - Requests that NSXML use single quotes for the value of an attribute or namespace node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodeUseSingleQuotes
	XMLNodeUseSingleQuotes XMLNodeOptions = 8
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


