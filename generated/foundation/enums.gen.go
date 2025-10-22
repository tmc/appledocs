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
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/producesRelativePathURLs
	DirectoryEnumerationProducesRelativePathURLs DirectoryEnumerationOptions = 3
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

// NSSearchPathDomainMask - Domain constants specifying base locations to use when you search for significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask
type SearchPathDomainMask uint

const (
	// AllDomainsMask - All domains.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/allDomainsMask
	AllDomainsMask SearchPathDomainMask = 0
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

// NSURLRelationship - Constants indicating the relationship between a directory and an item.
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

// NSFileManagerUnmountOptions - Options that specify the behavior of an unmount operation.
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
	FileManagerUnmountWithoutUI FileManagerUnmountOptions = 1
)

// NSVolumeEnumerationOptions - Options for enumerating mounted volumes with the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions
type VolumeEnumerationOptions uint

const (
	// VolumeEnumerationProduceFileReferenceURLs - The enumeration produces file reference URLs rather than path-based URLs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions/produceFileReferenceURLs
	VolumeEnumerationProduceFileReferenceURLs VolumeEnumerationOptions = 1
	// VolumeEnumerationSkipHiddenVolumes - The enumeration skips hidden volumes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions/skipHiddenVolumes
	VolumeEnumerationSkipHiddenVolumes VolumeEnumerationOptions = 1
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
	// InlinePresentationIntentEmphasized - An intent that represents an emphasized presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/emphasized
	InlinePresentationIntentEmphasized InlinePresentationIntent = 1
	// InlinePresentationIntentInlineHTML - An intent that represents an inline HTML presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/inlineHTML
	InlinePresentationIntentInlineHTML InlinePresentationIntent = 1
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

// NSAttributedStringFormattingOptions - Options to use when creating an attributed string from a format string and variable list of arguments.
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

// NSAttributedStringMarkdownParsingFailurePolicy - A type that represents policies for handling parsing failures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingFailurePolicy
type AttributedStringMarkdownParsingFailurePolicy uint

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

// NSBackgroundActivityResult - These constants indicate whether background activity has been completed successfully or whether additional processing should be deferred until a more optimal time.
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

// NSCalendarOptions - The options for arithmetic operations involving calendars.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
type CalendarOptions uint

const (
	// CalendarMatchFirst - Specifies that, if there are two or more matching times, the operation should return the first occurrence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchFirst
	CalendarMatchFirst CalendarOptions = 5
	// CalendarMatchLast - Specifies that, if there are two or more matching times, the operation should return the last occurrence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchLast
	CalendarMatchLast CalendarOptions = 6
	// CalendarMatchNextTime - Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given   object, this method uses the   existing value of the missing unit and   preserve the lower units’ values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchNextTime
	CalendarMatchNextTime CalendarOptions = 4
	// CalendarMatchNextTimePreservingSmallerUnits - Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given   object, this method uses the   existing value of the missing unit and preserves the lower units’ values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchNextTimePreservingSmallerUnits
	CalendarMatchNextTimePreservingSmallerUnits CalendarOptions = 3
	// CalendarMatchPreviousTimePreservingSmallerUnits - Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given   object, this method uses the   existing value of the missing unit and preserves the lower units’ values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchPreviousTimePreservingSmallerUnits
	CalendarMatchPreviousTimePreservingSmallerUnits CalendarOptions = 2
	// CalendarMatchStrictly - Specifies that the operation should travel as far forward or backward as necessary looking for a match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchStrictly
	CalendarMatchStrictly CalendarOptions = 0
	// CalendarSearchBackwards - Specifies that the operation should travel backwards to find the previous match before the given date.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/searchBackwards
	CalendarSearchBackwards CalendarOptions = 1
	// CalendarWrapComponents - Specifies that the components specified for an   object should be incremented and wrap around to zero/one on overflow, but should not cause higher units to be incremented.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/wrapComponents
	CalendarWrapComponents CalendarOptions = 0
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

// NSCollectionChangeType - The type of change represented in computing the difference of an ordered collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCollectionChangeType
type CollectionChangeType uint

const (
	CollectionChangeInsert CollectionChangeType = 0
	CollectionChangeRemove CollectionChangeType = 1
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

// NSFileCoordinatorReadingOptions - Options to use when reading the contents or attributes of a file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions
type FileCoordinatorReadingOptions uint

const (
	// FileCoordinatorReadingForUploading - Specify this content when reading an item for the purpose of uploading its contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/forUploading
	FileCoordinatorReadingForUploading FileCoordinatorReadingOptions = 3
	// FileCoordinatorReadingImmediatelyAvailableMetadataOnly - Specify this constant if you want to read an item’s metadata without triggering a download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/immediatelyAvailableMetadataOnly
	FileCoordinatorReadingImmediatelyAvailableMetadataOnly FileCoordinatorReadingOptions = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/resolvesSymbolicLink
	FileCoordinatorReadingResolvesSymbolicLink FileCoordinatorReadingOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/withoutChanges
	FileCoordinatorReadingWithoutChanges FileCoordinatorReadingOptions = 1
)

// NSFileCoordinatorWritingOptions - Options to use when changing the contents or attributes of a file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
type FileCoordinatorWritingOptions uint

const (
	// FileCoordinatorWritingContentIndependentMetadataOnly - Select this option when writing to change the file’s metadata only and not its contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/contentIndependentMetadataOnly
	FileCoordinatorWritingContentIndependentMetadataOnly FileCoordinatorWritingOptions = 2
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

// NSFileManagerSupportedSyncControls - An option set of the sync controls available for an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls
type FileManagerSupportedSyncControls uint

const (
	// FileManagerSupportedSyncControlsFailUploadOnConflict - The file provider supports failing an upload if the local and server versions conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls/failUploadOnConflict
	FileManagerSupportedSyncControlsFailUploadOnConflict FileManagerSupportedSyncControls = 1
	// FileManagerSupportedSyncControlsPauseSync - The file provider supports pausing the sync on the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls/pauseSync
	FileManagerSupportedSyncControlsPauseSync FileManagerSupportedSyncControls = 1
)

// NSFileManagerUploadLocalVersionConflictPolicy - The policies the file manager can apply to resolve conflicts when uploading a local version of a file.
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

// NSFileVersionReplacingOptions - Options for replacing a file version.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/ReplacingOptions
type FileVersionReplacingOptions uint

const (
	FileVersionReplacingByMoving FileVersionReplacingOptions = 1
)

// NSGrammaticalCase enum type
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

// NSGrammaticalDefiniteness enum type
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

// NSGrammaticalDetermination enum type
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
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/second
	GrammaticalPersonSecond GrammaticalPerson = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/third
	GrammaticalPersonThird GrammaticalPerson = 3
)

// NSGrammaticalPronounType enum type
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

// NSItemProviderErrorCode - The error codes that describe problems with consuming data from an item provider.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode
type ItemProviderErrorCode uint

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
type PreferredPresentationStyle uint

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
	LinguisticTaggerUnitDocument LinguisticTaggerUnit = 3
)

// NSMachPortOptions - Used to remove access rights to a mach port when the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options
type MachPortOptions uint

const (
	// MachPortDeallocateReceiveRight - Remove a receive right when the   object is invalidated or destroyed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options/deallocateReceiveRight
	MachPortDeallocateReceiveRight MachPortOptions = 0
	// MachPortDeallocateSendRight - Deallocate a send right when the   object is invalidated or destroyed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options/deallocateSendRight
	MachPortDeallocateSendRight MachPortOptions = 0
	// MachPortDeallocateNone - Do not remove any send or receive rights.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPortOptions/NSMachPortDeallocateNone
	MachPortDeallocateNone MachPortOptions = 0
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
	TextCheckingTypePhoneNumber TextCheckingType = 2
	// TextCheckingTypeQuote - Replaces quotes with smart quotes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/quote
	TextCheckingTypeQuote TextCheckingType = 0
	// TextCheckingTypeRegularExpression - Matches a regular expression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/regularExpression
	TextCheckingTypeRegularExpression TextCheckingType = 1
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
	TextCheckingTypeTransitInformation TextCheckingType = 3
)

// NSTimeZoneNameStyle - Constants you use to specify a style when presenting time zone names.
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
	// URLBookmarkCreationWithoutImplicitSecurityScope - Prevents inclusion of a bookmark’s implicit ephemeral security scope, when creating one without security scope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/withoutImplicitSecurityScope
	URLBookmarkCreationWithoutImplicitSecurityScope URLBookmarkCreationOptions = 3
)

// NSURLBookmarkResolutionOptions - Options used when resolving bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions
type URLBookmarkResolutionOptions uint

const (
	// URLBookmarkResolutionWithSecurityScope - Specifies that the security scope, applied to the bookmark when it was created, should be used during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withSecurityScope
	URLBookmarkResolutionWithSecurityScope URLBookmarkResolutionOptions = 0
	// URLBookmarkResolutionWithoutImplicitStartAccessing - A property that specifies that resolution doesn’t implicitly start accessing the ephemeral security-scoped resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutImplicitStartAccessing
	URLBookmarkResolutionWithoutImplicitStartAccessing URLBookmarkResolutionOptions = 1
	// URLBookmarkResolutionWithoutMounting - Specifies that no volume should be mounted during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutMounting
	URLBookmarkResolutionWithoutMounting URLBookmarkResolutionOptions = 0
	// URLBookmarkResolutionWithoutUI - Specifies that no UI feedback should accompany resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutUI
	URLBookmarkResolutionWithoutUI URLBookmarkResolutionOptions = 0
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
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType/NSURLSessionWebSocketMessageTypeData
	URLSessionWebSocketMessageTypeData URLSessionWebSocketMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType/NSURLSessionWebSocketMessageTypeString
	URLSessionWebSocketMessageTypeString URLSessionWebSocketMessageType = 1
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

// NSPostingStyle - The constants that specify when notifications are posted.
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
	ProcessInfoThermalStateCritical ProcessInfoThermalState = 3
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
	QualityOfServiceDefault QualityOfService = -1
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

// NSURLSessionResponseDisposition - Constants indicating how a data or upload session should proceed after receiving the initial headers.
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

// NSURLSessionMultipathServiceType - Constants that specify the type of service that Multipath TCP uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum
type URLSessionMultipathServiceType uint

const (
	// URLSessionMultipathServiceTypeAggregate - A service that aggregates the capacities of other Multipath options in an attempt to increase throughput and minimize latency.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/aggregate
	URLSessionMultipathServiceTypeAggregate URLSessionMultipathServiceType = 0
	// URLSessionMultipathServiceTypeHandover - A Multipath TCP service that provides seamless handover between Wi-Fi and cellular in order to preserve the connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/handover
	URLSessionMultipathServiceTypeHandover URLSessionMultipathServiceType = 0
	// URLSessionMultipathServiceTypeInteractive - A service whereby Multipath TCP attempts to use the lowest-latency interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/interactive
	URLSessionMultipathServiceTypeInteractive URLSessionMultipathServiceType = 0
	// URLSessionMultipathServiceTypeNone - The default service type indicating that Multipath TCP should not be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/none
	URLSessionMultipathServiceTypeNone URLSessionMultipathServiceType = 0
)

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

// NSXMLDTDNodeKind - The type defined for the constants that specify the kind and subkind of DTD declaration represented by an 
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

// NSXMLNodeKind enum type
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

// NSXMLNodeOptions - These constants are input and output options for all 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options
type XMLNodeOptions uint

const (
	// XMLNodeOptionsNone - No options are requested for this input or output action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXMLNodeOptions/NSXMLNodeOptionsNone
	XMLNodeOptionsNone XMLNodeOptions = 0
	// XMLNodePreserveDTD - Specifies that declarations in a DTD should be preserved until it the DTD is modified. For example, parameter entities are by default expanded; with this option, they are written out as they originally occur in the DTD.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodePreserveDTD
	XMLNodePreserveDTD XMLNodeOptions = 1
	// XMLNodePreserveQuotes - Specifies that the quoting style used in the input XML (single or double quotes) be preserved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodePreserveQuotes
	XMLNodePreserveQuotes XMLNodeOptions = 0
	// XMLNodeUseSingleQuotes - Requests that NSXML use single quotes for the value of an attribute or namespace node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodeUseSingleQuotes
	XMLNodeUseSingleQuotes XMLNodeOptions = 1
)

// NSXMLParserExternalEntityResolvingPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/ExternalEntityResolvingPolicy-swift.enum
type XMLParserExternalEntityResolvingPolicy uint

const (
	XMLParserResolveExternalEntitiesNever XMLParserExternalEntityResolvingPolicy = 0
	XMLParserResolveExternalEntitiesNoNetwork XMLParserExternalEntityResolvingPolicy = 1
	XMLParserResolveExternalEntitiesSameOriginOnly XMLParserExternalEntityResolvingPolicy = 2
	XMLParserResolveExternalEntitiesAlways XMLParserExternalEntityResolvingPolicy = 3
)


