// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
// NSDirectoryEnumerationOptions - Options for enumerating the contents of directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions
type NSDirectoryEnumerationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/includesDirectoriesPostOrder
	NSDirectoryEnumerationIncludesDirectoriesPostOrder NSDirectoryEnumerationOptions = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/producesRelativePathURLs
	NSDirectoryEnumerationProducesRelativePathURLs NSDirectoryEnumerationOptions = 3
	// NSDirectoryEnumerationSkipsHiddenFiles - An option to skip hidden files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsHiddenFiles
	NSDirectoryEnumerationSkipsHiddenFiles NSDirectoryEnumerationOptions = 1
	// NSDirectoryEnumerationSkipsPackageDescendants - An option to treat packages like files and not descend into their contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsPackageDescendants
	NSDirectoryEnumerationSkipsPackageDescendants NSDirectoryEnumerationOptions = 1
	// NSDirectoryEnumerationSkipsSubdirectoryDescendants - An option to perform a shallow enumeration that doesn’t descend into directories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions/skipsSubdirectoryDescendants
	NSDirectoryEnumerationSkipsSubdirectoryDescendants NSDirectoryEnumerationOptions = 1
)

// NSFileManagerItemReplacementOptions - Options for specifying the behavior of file replacement operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions
type NSFileManagerItemReplacementOptions uint

const (
	// NSFileManagerItemReplacementUsingNewMetadataOnly - Only metadata from the new item is used, and metadata from the original item isn’t preserved (default).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions/usingNewMetadataOnly
	NSFileManagerItemReplacementUsingNewMetadataOnly NSFileManagerItemReplacementOptions = 1
	// NSFileManagerItemReplacementWithoutDeletingBackupItem - The backup item remains in place after a successful replacement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions/withoutDeletingBackupItem
	NSFileManagerItemReplacementWithoutDeletingBackupItem NSFileManagerItemReplacementOptions = 1
)

// NSSearchPathDirectory - The location of significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
type NSSearchPathDirectory uint

const (
	// NSAdminApplicationDirectory - System and network administration applications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/adminApplicationDirectory
	NSAdminApplicationDirectory NSSearchPathDirectory = 4
	// NSAllApplicationsDirectory - All directories where applications can be stored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/allApplicationsDirectory
	NSAllApplicationsDirectory NSSearchPathDirectory = 100
	// NSAllLibrariesDirectory - All directories where resources can be stored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/allLibrariesDirectory
	NSAllLibrariesDirectory NSSearchPathDirectory = 101
	// NSApplicationDirectory - Supported applications ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationDirectory
	NSApplicationDirectory NSSearchPathDirectory = 1
	// NSApplicationScriptsDirectory - The user scripts folder for the calling application ( .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationScriptsDirectory
	NSApplicationScriptsDirectory NSSearchPathDirectory = 23
	// NSApplicationSupportDirectory - Application support files ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/applicationSupportDirectory
	NSApplicationSupportDirectory NSSearchPathDirectory = 14
	// NSAutosavedInformationDirectory - The user’s autosaved documents ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/autosavedInformationDirectory
	NSAutosavedInformationDirectory NSSearchPathDirectory = 11
	// NSCachesDirectory - Discardable cache files ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/cachesDirectory
	NSCachesDirectory NSSearchPathDirectory = 13
	// NSCoreServiceDirectory - Core services ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/coreServiceDirectory
	NSCoreServiceDirectory NSSearchPathDirectory = 10
	// NSDemoApplicationDirectory - Unsupported applications and demonstration versions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/demoApplicationDirectory
	NSDemoApplicationDirectory NSSearchPathDirectory = 2
	// NSDesktopDirectory - The user’s desktop directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/desktopDirectory
	NSDesktopDirectory NSSearchPathDirectory = 12
	// NSDeveloperApplicationDirectory - Developer applications ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/developerApplicationDirectory
	NSDeveloperApplicationDirectory NSSearchPathDirectory = 3
	// NSDeveloperDirectory - Developer resources ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/developerDirectory
	NSDeveloperDirectory NSSearchPathDirectory = 6
	// NSDocumentDirectory - Document directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentDirectory
	NSDocumentDirectory NSSearchPathDirectory = 9
	// NSDocumentationDirectory - Documentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentationDirectory
	NSDocumentationDirectory NSSearchPathDirectory = 8
	// NSDownloadsDirectory - The user’s downloads directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/downloadsDirectory
	NSDownloadsDirectory NSSearchPathDirectory = 15
	// NSInputMethodsDirectory - Input Methods  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/inputMethodsDirectory
	NSInputMethodsDirectory NSSearchPathDirectory = 16
	// NSItemReplacementDirectory - The constant used to create a temporary directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/itemReplacementDirectory
	NSItemReplacementDirectory NSSearchPathDirectory = 24
	// NSLibraryDirectory - Various user-visible documentation, support, and configuration files ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/libraryDirectory
	NSLibraryDirectory NSSearchPathDirectory = 5
	// NSMoviesDirectory - The user’s Movies directory  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/moviesDirectory
	NSMoviesDirectory NSSearchPathDirectory = 17
	// NSMusicDirectory - The user’s Music directory ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/musicDirectory
	NSMusicDirectory NSSearchPathDirectory = 18
	// NSPicturesDirectory - The user’s Pictures directory ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/picturesDirectory
	NSPicturesDirectory NSSearchPathDirectory = 19
	// NSPreferencePanesDirectory - The PreferencePanes directory for use with System Preferences ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/preferencePanesDirectory
	NSPreferencePanesDirectory NSSearchPathDirectory = 22
	// NSPrinterDescriptionDirectory - The system’s PPDs directory ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/printerDescriptionDirectory
	NSPrinterDescriptionDirectory NSSearchPathDirectory = 20
	// NSSharedPublicDirectory - The user’s Public sharing directory ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/sharedPublicDirectory
	NSSharedPublicDirectory NSSearchPathDirectory = 21
	// NSTrashDirectory - The trash directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/trashDirectory
	NSTrashDirectory NSSearchPathDirectory = 102
	// NSUserDirectory - User home directories ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/userDirectory
	NSUserDirectory NSSearchPathDirectory = 7
)

// NSSearchPathDomainMask - Domain constants specifying base locations to use when you search for significant directories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask
type NSSearchPathDomainMask uint

const (
	// NSAllDomainsMask - All domains.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/allDomainsMask
	NSAllDomainsMask NSSearchPathDomainMask = 0
	// NSLocalDomainMask - The place to install items available to everyone on this machine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/localDomainMask
	NSLocalDomainMask NSSearchPathDomainMask = 2
	// NSNetworkDomainMask - The place to install items available on the network ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/networkDomainMask
	NSNetworkDomainMask NSSearchPathDomainMask = 4
	// NSSystemDomainMask - A directory for system files provided by Apple ( ) .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/systemDomainMask
	NSSystemDomainMask NSSearchPathDomainMask = 8
	// NSUserDomainMask - The user’s home directory—the place to install user’s personal items ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask/userDomainMask
	NSUserDomainMask NSSearchPathDomainMask = 1
)

// NSURLRelationship - Constants indicating the relationship between a directory and an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship
type NSURLRelationship uint

const (
	// NSURLRelationshipContains - The directory contains the specified item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship/contains
	NSURLRelationshipContains NSURLRelationship = 0
	// NSURLRelationshipOther - The directory does not contain the item and is not the same as the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship/other
	NSURLRelationshipOther NSURLRelationship = 2
	// NSURLRelationshipSame - The directory and the item are the same. This relationship occurs when the value of the   is the same for the directory and item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship/same
	NSURLRelationshipSame NSURLRelationship = 1
)

// NSFileManagerUnmountOptions - Options that specify the behavior of an unmount operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions
type NSFileManagerUnmountOptions uint

const (
	// NSFileManagerUnmountAllPartitionsAndEjectDisk - Specifies that all partitions on an unmountable disk should be unmounted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions/allPartitionsAndEjectDisk
	NSFileManagerUnmountAllPartitionsAndEjectDisk NSFileManagerUnmountOptions = 1
	// NSFileManagerUnmountWithoutUI - Specifies that no UI should accompany the unmount operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions/withoutUI
	NSFileManagerUnmountWithoutUI NSFileManagerUnmountOptions = 1
)

// NSVolumeEnumerationOptions - Options for enumerating mounted volumes with the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions
type NSVolumeEnumerationOptions uint

const (
	// NSVolumeEnumerationProduceFileReferenceURLs - The enumeration produces file reference URLs rather than path-based URLs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions/produceFileReferenceURLs
	NSVolumeEnumerationProduceFileReferenceURLs NSVolumeEnumerationOptions = 1
	// NSVolumeEnumerationSkipHiddenVolumes - The enumeration skips hidden volumes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions/skipHiddenVolumes
	NSVolumeEnumerationSkipHiddenVolumes NSVolumeEnumerationOptions = 1
)

// NSInlinePresentationIntent - A type that defines presentation intent for runs of characters for traits like emphasis, strikethrough, and code voice.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent
type NSInlinePresentationIntent uint

const (
	// NSInlinePresentationIntentEmphasized - An intent that represents an emphasized presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/emphasized
	NSInlinePresentationIntentEmphasized NSInlinePresentationIntent = 1
	// NSInlinePresentationIntentInlineHTML - An intent that represents an inline HTML presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/inlineHTML
	NSInlinePresentationIntentInlineHTML NSInlinePresentationIntent = 1
	// NSInlinePresentationIntentStronglyEmphasized - An intent that represents a strongly emphasized presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/stronglyEmphasized
	NSInlinePresentationIntentStronglyEmphasized NSInlinePresentationIntent = 1
)

// NSAttributedStringEnumerationOptions - Options for enumerating attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/EnumerationOptions
type NSAttributedStringEnumerationOptions uint

// NSSpellingState - Constants for the spelling state attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/SpellingState
type NSSpellingState uint

// NSAttributedStringFormattingOptions - Options to use when creating an attributed string from a format string and variable list of arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingOptions
type NSAttributedStringFormattingOptions uint

const (
	// NSAttributedStringFormattingApplyReplacementIndexAttribute - An option to apply to the replaced portions of text in a format string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingOptions/NSAttributedStringFormattingApplyReplacementIndexAttribute
	NSAttributedStringFormattingApplyReplacementIndexAttribute NSAttributedStringFormattingOptions = 1
	// NSAttributedStringFormattingInsertArgumentAttributesWithoutMerging - An option to replace the attributes in a substituted string with those of the provided attributed string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingOptions/NSAttributedStringFormattingInsertArgumentAttributesWithoutMerging
	NSAttributedStringFormattingInsertArgumentAttributesWithoutMerging NSAttributedStringFormattingOptions = 0
)

// NSAttributedStringMarkdownParsingFailurePolicy - A type that represents policies for handling parsing failures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingFailurePolicy
type NSAttributedStringMarkdownParsingFailurePolicy uint

const (
	// NSAttributedStringMarkdownParsingFailureReturnError - A policy to return an error from the initializer if parsing fails.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingFailurePolicy/NSAttributedStringMarkdownParsingFailureReturnError
	NSAttributedStringMarkdownParsingFailureReturnError NSAttributedStringMarkdownParsingFailurePolicy = 0
	// NSAttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible - A policy to return a partially parsed string, if possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingFailurePolicy/NSAttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible
	NSAttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible NSAttributedStringMarkdownParsingFailurePolicy = 1
)

// NSBackgroundActivityResult - These constants indicate whether background activity has been completed successfully or whether additional processing should be deferred until a more optimal time.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/Result
type NSBackgroundActivityResult uint

const (
	// NSBackgroundActivityResultDeferred - System conditions have changed since the time the activity began executing, and deferral of additional work is recommended.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/Result/deferred
	NSBackgroundActivityResultDeferred NSBackgroundActivityResult = 2
	// NSBackgroundActivityResultFinished - The activity has finished executing. If the activity repeats, the next invocation is scheduled by the system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/Result/finished
	NSBackgroundActivityResultFinished NSBackgroundActivityResult = 1
)

// NSCalendarOptions - The options for arithmetic operations involving calendars.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
type NSCalendarOptions uint

const (
	// NSCalendarMatchFirst - Specifies that, if there are two or more matching times, the operation should return the first occurrence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchFirst
	NSCalendarMatchFirst NSCalendarOptions = 5
	// NSCalendarMatchLast - Specifies that, if there are two or more matching times, the operation should return the last occurrence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchLast
	NSCalendarMatchLast NSCalendarOptions = 6
	// NSCalendarMatchNextTime - Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given   object, this method uses the   existing value of the missing unit and   preserve the lower units’ values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchNextTime
	NSCalendarMatchNextTime NSCalendarOptions = 4
	// NSCalendarMatchNextTimePreservingSmallerUnits - Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given   object, this method uses the   existing value of the missing unit and preserves the lower units’ values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchNextTimePreservingSmallerUnits
	NSCalendarMatchNextTimePreservingSmallerUnits NSCalendarOptions = 3
	// NSCalendarMatchPreviousTimePreservingSmallerUnits - Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given   object, this method uses the   existing value of the missing unit and preserves the lower units’ values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchPreviousTimePreservingSmallerUnits
	NSCalendarMatchPreviousTimePreservingSmallerUnits NSCalendarOptions = 2
	// NSCalendarMatchStrictly - Specifies that the operation should travel as far forward or backward as necessary looking for a match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/matchStrictly
	NSCalendarMatchStrictly NSCalendarOptions = 0
	// NSCalendarSearchBackwards - Specifies that the operation should travel backwards to find the previous match before the given date.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/searchBackwards
	NSCalendarSearchBackwards NSCalendarOptions = 1
	// NSCalendarWrapComponents - Specifies that the components specified for an   object should be incremented and wrap around to zero/one on overflow, but should not cause higher units to be incremented.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options/wrapComponents
	NSCalendarWrapComponents NSCalendarOptions = 0
)

// NSCalendarUnit - Calendrical units such as year, month, day and hour.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
type NSCalendarUnit uint

const (
	// NSCalendarCalendarUnit - Specifies the calendar of the calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSCalendarCalendarUnit
	NSCalendarCalendarUnit NSCalendarUnit = 22
	// NSDayCalendarUnit - Specifies the day unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSDayCalendarUnit
	NSDayCalendarUnit NSCalendarUnit = 11
	// NSEraCalendarUnit - Specifies the era unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSEraCalendarUnit
	NSEraCalendarUnit NSCalendarUnit = 8
	// NSHourCalendarUnit - Specifies the hour unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSHourCalendarUnit
	NSHourCalendarUnit NSCalendarUnit = 12
	// NSMinuteCalendarUnit - Specifies the minute unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSMinuteCalendarUnit
	NSMinuteCalendarUnit NSCalendarUnit = 13
	// NSMonthCalendarUnit - Specifies the month unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSMonthCalendarUnit
	NSMonthCalendarUnit NSCalendarUnit = 10
	// NSQuarterCalendarUnit - Specifies the quarter unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSQuarterCalendarUnit
	NSQuarterCalendarUnit NSCalendarUnit = 18
	// NSSecondCalendarUnit - Specifies the second unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSSecondCalendarUnit
	NSSecondCalendarUnit NSCalendarUnit = 14
	// NSTimeZoneCalendarUnit - Specifies the time zone of the calendar as an  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSTimeZoneCalendarUnit
	NSTimeZoneCalendarUnit NSCalendarUnit = 23
	// NSWeekCalendarUnit - Specifies the week unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekCalendarUnit
	NSWeekCalendarUnit NSCalendarUnit = 15
	// NSWeekOfMonthCalendarUnit - Specifies the original week of a month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekOfMonthCalendarUnit
	NSWeekOfMonthCalendarUnit NSCalendarUnit = 19
	// NSWeekOfYearCalendarUnit - Specifies the original week of the year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekOfYearCalendarUnit
	NSWeekOfYearCalendarUnit NSCalendarUnit = 20
	// NSWeekdayCalendarUnit - Specifies the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekdayCalendarUnit
	NSWeekdayCalendarUnit NSCalendarUnit = 16
	// NSWeekdayOrdinalCalendarUnit - Specifies the ordinal weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekdayOrdinalCalendarUnit
	NSWeekdayOrdinalCalendarUnit NSCalendarUnit = 17
	// NSYearCalendarUnit - Specifies the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSYearCalendarUnit
	NSYearCalendarUnit NSCalendarUnit = 9
	// NSYearForWeekOfYearCalendarUnit - Specifies the year when the calendar is being interpreted as a week-based calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSYearForWeekOfYearCalendarUnit
	NSYearForWeekOfYearCalendarUnit NSCalendarUnit = 21
	// NSCalendarUnitCalendar - Identifier for the calendar of a date components object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/calendar
	NSCalendarUnitCalendar NSCalendarUnit = 6
	// NSCalendarUnitDay - Identifier for the day unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/day
	NSCalendarUnitDay NSCalendarUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/dayOfYear
	NSCalendarUnitDayOfYear NSCalendarUnit = 5
	// NSCalendarUnitEra - Identifier for the era unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/era
	NSCalendarUnitEra NSCalendarUnit = 0
	// NSCalendarUnitHour - Identifier for the hour unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/hour
	NSCalendarUnitHour NSCalendarUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/isLeapMonth
	NSCalendarUnitIsLeapMonth NSCalendarUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/isRepeatedDay
	NSCalendarUnitIsRepeatedDay NSCalendarUnit = 0
	// NSCalendarUnitMinute - Identifier for the minute unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/minute
	NSCalendarUnitMinute NSCalendarUnit = 0
	// NSCalendarUnitMonth - Identifier for the month unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/month
	NSCalendarUnitMonth NSCalendarUnit = 0
	// NSCalendarUnitNanosecond - Identifier for the nanosecond unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/nanosecond
	NSCalendarUnitNanosecond NSCalendarUnit = 4
	// NSCalendarUnitQuarter - Identifier for the quarter of the calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/quarter
	NSCalendarUnitQuarter NSCalendarUnit = 0
	// NSCalendarUnitSecond - Identifier for the second unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/second
	NSCalendarUnitSecond NSCalendarUnit = 0
	// NSCalendarUnitTimeZone - Identifier for the time zone of a date components object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/timeZone
	NSCalendarUnitTimeZone NSCalendarUnit = 7
	// NSCalendarUnitWeekOfMonth - Identifier for the week of the month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekOfMonth
	NSCalendarUnitWeekOfMonth NSCalendarUnit = 1
	// NSCalendarUnitWeekOfYear - Identifier for the week of the year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekOfYear
	NSCalendarUnitWeekOfYear NSCalendarUnit = 2
	// NSCalendarUnitWeekday - Identifier for the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekday
	NSCalendarUnitWeekday NSCalendarUnit = 0
	// NSCalendarUnitWeekdayOrdinal - Identifier for the ordinal weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekdayOrdinal
	NSCalendarUnitWeekdayOrdinal NSCalendarUnit = 0
	// NSCalendarUnitYear - Identifier for the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/year
	NSCalendarUnitYear NSCalendarUnit = 0
	// NSCalendarUnitYearForWeekOfYear - Identifier for the week-counting year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/yearForWeekOfYear
	NSCalendarUnitYearForWeekOfYear NSCalendarUnit = 3
)

// NSDataBase64DecodingOptions - Options to modify the decoding algorithm used to decode Base64 encoded data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions
type NSDataBase64DecodingOptions uint

const (
	NSDataBase64DecodingIgnoreUnknownCharacters NSDataBase64DecodingOptions = 1
)

// NSDataBase64EncodingOptions - Options for methods used to Base64 encode data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions
type NSDataBase64EncodingOptions uint

const (
	NSDataBase64Encoding64CharacterLineLength NSDataBase64EncodingOptions = 1
	NSDataBase64Encoding76CharacterLineLength NSDataBase64EncodingOptions = 1
	NSDataBase64EncodingEndLineWithCarriageReturn NSDataBase64EncodingOptions = 1
	NSDataBase64EncodingEndLineWithLineFeed NSDataBase64EncodingOptions = 1
)

// NSDataCompressionAlgorithm - An algorithm that indicates how to compress or decompress data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
type NSDataCompressionAlgorithm uint

const (
	NSDataCompressionAlgorithmLZFSE NSDataCompressionAlgorithm = 0
	NSDataCompressionAlgorithmLZ4 NSDataCompressionAlgorithm = 1
	NSDataCompressionAlgorithmLZMA NSDataCompressionAlgorithm = 2
	NSDataCompressionAlgorithmZlib NSDataCompressionAlgorithm = 3
)

// NSDataReadingOptions - Options for methods used to read data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions
type NSDataReadingOptions uint

const (
	NSDataReadingMappedIfSafe NSDataReadingOptions = 1
	NSDataReadingUncached NSDataReadingOptions = 1
	NSDataReadingMappedAlways NSDataReadingOptions = 2
	NSDataReadingMapped NSDataReadingOptions = 3
	NSMappedRead NSDataReadingOptions = 4
	NSUncachedRead NSDataReadingOptions = 5
)

// NSDataSearchOptions - Options for method used to search data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions
type NSDataSearchOptions uint

const (
	NSDataSearchBackwards NSDataSearchOptions = 1
	NSDataSearchAnchored NSDataSearchOptions = 1
)

// NSDataWritingOptions - Options for methods used to write data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions
type NSDataWritingOptions uint

const (
	NSDataWritingAtomic NSDataWritingOptions = 1
	NSDataWritingWithoutOverwriting NSDataWritingOptions = 2
	NSDataWritingFileProtectionNone NSDataWritingOptions = 3
	NSDataWritingFileProtectionComplete NSDataWritingOptions = 4
	NSDataWritingFileProtectionCompleteUnlessOpen NSDataWritingOptions = 5
	NSDataWritingFileProtectionCompleteUntilFirstUserAuthentication NSDataWritingOptions = 6
	NSDataWritingFileProtectionCompleteWhenUserInactive NSDataWritingOptions = 7
	NSDataWritingFileProtectionMask NSDataWritingOptions = 8
	NSAtomicWrite NSDataWritingOptions = 9
)

// NSFileCoordinatorReadingOptions - Options to use when reading the contents or attributes of a file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions
type NSFileCoordinatorReadingOptions uint

const (
	// NSFileCoordinatorReadingForUploading - Specify this content when reading an item for the purpose of uploading its contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/forUploading
	NSFileCoordinatorReadingForUploading NSFileCoordinatorReadingOptions = 3
	// NSFileCoordinatorReadingImmediatelyAvailableMetadataOnly - Specify this constant if you want to read an item’s metadata without triggering a download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/immediatelyAvailableMetadataOnly
	NSFileCoordinatorReadingImmediatelyAvailableMetadataOnly NSFileCoordinatorReadingOptions = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/resolvesSymbolicLink
	NSFileCoordinatorReadingResolvesSymbolicLink NSFileCoordinatorReadingOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions/withoutChanges
	NSFileCoordinatorReadingWithoutChanges NSFileCoordinatorReadingOptions = 1
)

// NSFileCoordinatorWritingOptions - Options to use when changing the contents or attributes of a file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
type NSFileCoordinatorWritingOptions uint

const (
	// NSFileCoordinatorWritingContentIndependentMetadataOnly - Select this option when writing to change the file’s metadata only and not its contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/contentIndependentMetadataOnly
	NSFileCoordinatorWritingContentIndependentMetadataOnly NSFileCoordinatorWritingOptions = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forDeleting
	NSFileCoordinatorWritingForDeleting NSFileCoordinatorWritingOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forMoving
	NSFileCoordinatorWritingForMoving NSFileCoordinatorWritingOptions = 1
)

// NSFileManagerResumeSyncBehavior - The behaviors the file manager can apply to resolve conflicts when resuming a sync.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior
type NSFileManagerResumeSyncBehavior uint

const (
	// NSFileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict - Resumes sync by first uploading the local version of the file, failing if the provider detects a conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior/afterUploadWithFailOnConflict
	NSFileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict NSFileManagerResumeSyncBehavior = 1
	// NSFileManagerResumeSyncBehaviorDropLocalChanges - Resumes synchronizing by overwriting any local changes with the remote version of the file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior/dropLocalChanges
	NSFileManagerResumeSyncBehaviorDropLocalChanges NSFileManagerResumeSyncBehavior = 2
	// NSFileManagerResumeSyncBehaviorPreserveLocalChanges - Resumes synchronizing by uploading the local version of the file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior/preserveLocalChanges
	NSFileManagerResumeSyncBehaviorPreserveLocalChanges NSFileManagerResumeSyncBehavior = 0
)

// NSFileManagerSupportedSyncControls - An option set of the sync controls available for an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls
type NSFileManagerSupportedSyncControls uint

const (
	// NSFileManagerSupportedSyncControlsFailUploadOnConflict - The file provider supports failing an upload if the local and server versions conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls/failUploadOnConflict
	NSFileManagerSupportedSyncControlsFailUploadOnConflict NSFileManagerSupportedSyncControls = 1
	// NSFileManagerSupportedSyncControlsPauseSync - The file provider supports pausing the sync on the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls/pauseSync
	NSFileManagerSupportedSyncControlsPauseSync NSFileManagerSupportedSyncControls = 1
)

// NSFileManagerUploadLocalVersionConflictPolicy - The policies the file manager can apply to resolve conflicts when uploading a local version of a file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy
type NSFileManagerUploadLocalVersionConflictPolicy uint

const (
	// NSFileManagerUploadConflictPolicyDefault - Resolves the conflict using the policy defined by the file provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy/conflictPolicyDefault
	NSFileManagerUploadConflictPolicyDefault NSFileManagerUploadLocalVersionConflictPolicy = 0
	// NSFileManagerUploadConflictPolicyFailOnConflict - Resolves the conflict by causing the upload to fail.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy/conflictPolicyFailOnConflict
	NSFileManagerUploadConflictPolicyFailOnConflict NSFileManagerUploadLocalVersionConflictPolicy = 1
)

// NSFileVersionAddingOptions - Options for adding a new file version.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/AddingOptions
type NSFileVersionAddingOptions uint

const (
	NSFileVersionAddingByMoving NSFileVersionAddingOptions = 1
)

// NSFileVersionReplacingOptions - Options for replacing a file version.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/ReplacingOptions
type NSFileVersionReplacingOptions uint

const (
	// NSFileVersionReplacingByMoving - An option to perform replacing by moving a file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/ReplacingOptions/byMoving
	NSFileVersionReplacingByMoving NSFileVersionReplacingOptions = 1
)

// NSGrammaticalCase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase
type NSGrammaticalCase uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/ablative
	NSGrammaticalCaseAblative NSGrammaticalCase = 6
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/accusative
	NSGrammaticalCaseAccusative NSGrammaticalCase = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/adessive
	NSGrammaticalCaseAdessive NSGrammaticalCase = 7
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/allative
	NSGrammaticalCaseAllative NSGrammaticalCase = 8
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/dative
	NSGrammaticalCaseDative NSGrammaticalCase = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/elative
	NSGrammaticalCaseElative NSGrammaticalCase = 9
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/essive
	NSGrammaticalCaseEssive NSGrammaticalCase = 11
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/genitive
	NSGrammaticalCaseGenitive NSGrammaticalCase = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/illative
	NSGrammaticalCaseIllative NSGrammaticalCase = 10
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/inessive
	NSGrammaticalCaseInessive NSGrammaticalCase = 12
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/locative
	NSGrammaticalCaseLocative NSGrammaticalCase = 13
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/nominative
	NSGrammaticalCaseNominative NSGrammaticalCase = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/notSet
	NSGrammaticalCaseNotSet NSGrammaticalCase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/prepositional
	NSGrammaticalCasePrepositional NSGrammaticalCase = 5
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase/translative
	NSGrammaticalCaseTranslative NSGrammaticalCase = 14
)

// NSGrammaticalPerson enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson
type NSGrammaticalPerson uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/first
	NSGrammaticalPersonFirst NSGrammaticalPerson = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/notSet
	NSGrammaticalPersonNotSet NSGrammaticalPerson = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/second
	NSGrammaticalPersonSecond NSGrammaticalPerson = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson/third
	NSGrammaticalPersonThird NSGrammaticalPerson = 3
)

// NSGrammaticalPronounType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType
type NSGrammaticalPronounType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/notSet
	NSGrammaticalPronounTypeNotSet NSGrammaticalPronounType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/personal
	NSGrammaticalPronounTypePersonal NSGrammaticalPronounType = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/possessive
	NSGrammaticalPronounTypePossessive NSGrammaticalPronounType = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType/reflexive
	NSGrammaticalPronounTypeReflexive NSGrammaticalPronounType = 2
)

// NSItemProviderErrorCode - The error codes that describe problems with consuming data from an item provider.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode
type NSItemProviderErrorCode uint

const (
	// NSItemProviderItemUnavailableError - An error code indicating that the requested data was unavailable from an item provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode/itemUnavailableError
	NSItemProviderItemUnavailableError NSItemProviderErrorCode = -1000
	// NSItemProviderUnavailableCoercionError - An error code indicating that the requested data type coercion is unavailable from an item provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode/unavailableCoercionError
	NSItemProviderUnavailableCoercionError NSItemProviderErrorCode = -1099
	// NSItemProviderUnexpectedValueClassError - An error code indicating that type coercion to the requested class failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode/unexpectedValueClassError
	NSItemProviderUnexpectedValueClassError NSItemProviderErrorCode = -1100
	// NSItemProviderUnknownError - An error code indicating an unknown error with consuming data from an item provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode/unknownError
	NSItemProviderUnknownError NSItemProviderErrorCode = -1
)

// UIPreferredPresentationStyle - The presentation styles that determine how a view shows an item provider’s data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/PreferredPresentationStyle-swift.enum
type UIPreferredPresentationStyle uint

// NSItemProviderFileOptions - Data-access specifications that declare how to handle items.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions
type NSItemProviderFileOptions uint

const (
	// NSItemProviderFileOptionOpenInPlace - A data-access specification declaring that items should open in place, rather than being copied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions/openInPlace
	NSItemProviderFileOptionOpenInPlace NSItemProviderFileOptions = 1
)

// NSItemProviderRepresentationVisibility - Specifications that control which categories of processes can see an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility
type NSItemProviderRepresentationVisibility uint

const (
	// NSItemProviderRepresentationVisibilityAll - A representation visibility specification conferring item visibility to all processes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility/all
	NSItemProviderRepresentationVisibilityAll NSItemProviderRepresentationVisibility = 0
	// NSItemProviderRepresentationVisibilityGroup - A representation visibility specification confining item visibility to the app’s app group.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility/group
	NSItemProviderRepresentationVisibilityGroup NSItemProviderRepresentationVisibility = 2
	// NSItemProviderRepresentationVisibilityOwnProcess - A representation visibility specification confining item visibility to the app that is the source of the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility/ownProcess
	NSItemProviderRepresentationVisibilityOwnProcess NSItemProviderRepresentationVisibility = 3
	// NSItemProviderRepresentationVisibilityTeam - A representation visibility specification confining item visibility to processes created by the app’s development team.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility/team
	NSItemProviderRepresentationVisibilityTeam NSItemProviderRepresentationVisibility = 1
)

// NSLinguisticTaggerOptions - Constants for linguistic tagger enumeration specifying which tokens to omit and whether to join names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options
type NSLinguisticTaggerOptions uint

const (
	NSLinguisticTaggerOmitWords NSLinguisticTaggerOptions = 1
	NSLinguisticTaggerOmitPunctuation NSLinguisticTaggerOptions = 1
	NSLinguisticTaggerOmitWhitespace NSLinguisticTaggerOptions = 1
	NSLinguisticTaggerOmitOther NSLinguisticTaggerOptions = 1
	NSLinguisticTaggerJoinNames NSLinguisticTaggerOptions = 1
)

// NSMachPortOptions - Used to remove access rights to a mach port when the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options
type NSMachPortOptions uint

const (
	// NSMachPortDeallocateReceiveRight - Remove a receive right when the   object is invalidated or destroyed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options/deallocateReceiveRight
	NSMachPortDeallocateReceiveRight NSMachPortOptions = 0
	// NSMachPortDeallocateSendRight - Deallocate a send right when the   object is invalidated or destroyed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options/deallocateSendRight
	NSMachPortDeallocateSendRight NSMachPortOptions = 0
	// NSMachPortDeallocateNone - Do not remove any send or receive rights.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPortOptions/NSMachPortDeallocateNone
	NSMachPortDeallocateNone NSMachPortOptions = 0
)

// NSPointerFunctionsOptions - Defines the memory and personality options for an 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
type NSPointerFunctionsOptions uint

const (
	// NSPointerFunctionsCStringPersonality - Use a string hash and  ; C-string ‘ ’ style description.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/cStringPersonality
	NSPointerFunctionsCStringPersonality NSPointerFunctionsOptions = 9
	// NSPointerFunctionsCopyIn - Use the memory acquire function to allocate and copy items on input (see  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/copyIn
	NSPointerFunctionsCopyIn NSPointerFunctionsOptions = 12
	// NSPointerFunctionsIntegerPersonality - Use unshifted value as hash and equality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/integerPersonality
	NSPointerFunctionsIntegerPersonality NSPointerFunctionsOptions = 11
	// NSPointerFunctionsMachVirtualMemory - Use Mach memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/machVirtualMemory
	NSPointerFunctionsMachVirtualMemory NSPointerFunctionsOptions = 4
	// NSPointerFunctionsMallocMemory - Use   on removal,   on copy in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/mallocMemory
	NSPointerFunctionsMallocMemory NSPointerFunctionsOptions = 3
	// NSPointerFunctionsObjectPersonality - Use   and   methods for hashing and equality comparisons, use the   method for a description.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/objectPersonality
	NSPointerFunctionsObjectPersonality NSPointerFunctionsOptions = 6
	// NSPointerFunctionsObjectPointerPersonality - Use shifted pointer for the hash value and direct comparison to determine equality; use the   method for a description.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/objectPointerPersonality
	NSPointerFunctionsObjectPointerPersonality NSPointerFunctionsOptions = 8
	// NSPointerFunctionsOpaqueMemory - Take no action when pointers are deleted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/opaqueMemory
	NSPointerFunctionsOpaqueMemory NSPointerFunctionsOptions = 2
	// NSPointerFunctionsOpaquePersonality - Use shifted pointer for the hash value and direct comparison to determine equality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/opaquePersonality
	NSPointerFunctionsOpaquePersonality NSPointerFunctionsOptions = 7
	// NSPointerFunctionsStrongMemory - Use strong write-barriers to backing store; use garbage-collected memory on copy-in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/strongMemory
	NSPointerFunctionsStrongMemory NSPointerFunctionsOptions = 0
	// NSPointerFunctionsStructPersonality - Use a memory hash and   (using a size function that you must set—see  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/structPersonality
	NSPointerFunctionsStructPersonality NSPointerFunctionsOptions = 10
	// NSPointerFunctionsWeakMemory - Uses weak read and write barriers appropriate for ARC or GC. Using NSPointerFunctionsWeakMemory object references will turn to   on last release.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/weakMemory
	NSPointerFunctionsWeakMemory NSPointerFunctionsOptions = 5
	// NSPointerFunctionsZeroingWeakMemory - Use weak read and write barriers; use garbage-collected memory on copyIn.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctionsOptions/NSPointerFunctionsZeroingWeakMemory
	NSPointerFunctionsZeroingWeakMemory NSPointerFunctionsOptions = 1
)

// NSSaveOptions - The 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions
type NSSaveOptions uint

const (
	// NSSaveOptionsAsk - Indicates the user should be asked before saving any modified documents on closing. When no option is specified, this is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions/ask
	NSSaveOptionsAsk NSSaveOptions = 2
	// NSSaveOptionsNo - Indicates a modified document should not be saved on closing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions/no
	NSSaveOptionsNo NSSaveOptions = 1
	// NSSaveOptionsYes - Indicates a modified document should be saved on closing without asking the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSaveOptions/yes
	NSSaveOptionsYes NSSaveOptions = 0
)

// NSStringCompareOptions - These values represent the options available to many of the string classes’ search and comparison methods.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions
type NSStringCompareOptions uint

const (
	// NSAnchoredSearch - Search is limited to start (or end, if  ) of source string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/anchored
	NSAnchoredSearch NSStringCompareOptions = 8
	// NSBackwardsSearch - Search from end of source string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/backwards
	NSBackwardsSearch NSStringCompareOptions = 4
	// NSCaseInsensitiveSearch - A case-insensitive search.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/caseInsensitive
	NSCaseInsensitiveSearch NSStringCompareOptions = 1
	// NSDiacriticInsensitiveSearch - Search ignores diacritic marks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/diacriticInsensitive
	NSDiacriticInsensitiveSearch NSStringCompareOptions = 65
	// NSForcedOrderingSearch - Comparisons are forced to return either   or   if the strings are equivalent but not strictly equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/forcedOrdering
	NSForcedOrderingSearch NSStringCompareOptions = 67
	// NSLiteralSearch - Exact character-by-character equivalence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/literal
	NSLiteralSearch NSStringCompareOptions = 2
	// NSNumericSearch - Numbers within strings are compared using numeric value, that is,   <   <  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/numeric
	NSNumericSearch NSStringCompareOptions = 64
	// NSRegularExpressionSearch - The search string is treated as an ICU-compatible regular expression. If set, no other options can apply except   and  . You can use this option only with the  … methods and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/regularExpression
	NSRegularExpressionSearch NSStringCompareOptions = 68
	// NSWidthInsensitiveSearch - Search ignores width differences in characters that have full-width and half-width forms, as occurs in East Asian character sets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/widthInsensitive
	NSWidthInsensitiveSearch NSStringCompareOptions = 66
)

// NSStringDrawingOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions
type NSStringDrawingOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/disableScreenFontSubstitution
	disableScreenFontSubstitution NSStringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/oneShot
	oneShot NSStringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/optionsResolvesNaturalAlignmentWithBaseWritingDirection
	optionsResolvesNaturalAlignmentWithBaseWritingDirection NSStringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/truncatesLastVisibleLine
	truncatesLastVisibleLine NSStringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/usesDeviceMetrics
	usesDeviceMetrics NSStringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/usesFontLeading
	usesFontLeading NSStringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/usesLineFragmentOrigin
	usesLineFragmentOrigin NSStringDrawingOptions = 0
)

// NSStringEncodingConversionOptions - Options for converting string encodings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions
type NSStringEncodingConversionOptions uint

const (
	// NSStringEncodingConversionAllowLossy - Allows lossy conversion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions/allowLossy
	NSStringEncodingConversionAllowLossy NSStringEncodingConversionOptions = 1
	// NSStringEncodingConversionExternalRepresentation - Specifies an external representation (with a byte-order mark, if necessary, to indicate endianness).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions/externalRepresentation
	NSStringEncodingConversionExternalRepresentation NSStringEncodingConversionOptions = 2
)

// NSStringEnumerationOptions - Constants to specify kinds of substrings and styles of enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions
type NSStringEnumerationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byCaretPositions
	NSStringEnumerationByCaretPositions NSStringEnumerationOptions = 5
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byComposedCharacterSequences
	NSStringEnumerationByComposedCharacterSequences NSStringEnumerationOptions = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byDeletionClusters
	NSStringEnumerationByDeletionClusters NSStringEnumerationOptions = 6
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byLines
	NSStringEnumerationByLines NSStringEnumerationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byParagraphs
	NSStringEnumerationByParagraphs NSStringEnumerationOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/bySentences
	NSStringEnumerationBySentences NSStringEnumerationOptions = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byWords
	NSStringEnumerationByWords NSStringEnumerationOptions = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/localized
	NSStringEnumerationLocalized NSStringEnumerationOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/reverse
	NSStringEnumerationReverse NSStringEnumerationOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/substringNotRequired
	NSStringEnumerationSubstringNotRequired NSStringEnumerationOptions = 1
)

// NSTimeZoneNameStyle - Constants you use to specify a style when presenting time zone names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle
type NSTimeZoneNameStyle uint

const (
	// NSTimeZoneNameStyleDaylightSaving - Specifies a daylight saving name style. For example, “Central Daylight Time” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/daylightSaving
	NSTimeZoneNameStyleDaylightSaving NSTimeZoneNameStyle = 2
	// NSTimeZoneNameStyleGeneric - Specifies a generic name style. For example, “Central Time” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/generic
	NSTimeZoneNameStyleGeneric NSTimeZoneNameStyle = 4
	// NSTimeZoneNameStyleShortDaylightSaving - Specifies a short daylight saving name style.  For example, “CDT” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/shortDaylightSaving
	NSTimeZoneNameStyleShortDaylightSaving NSTimeZoneNameStyle = 3
	// NSTimeZoneNameStyleShortGeneric - Specifies a generic time zone name. For example, “CT” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/shortGeneric
	NSTimeZoneNameStyleShortGeneric NSTimeZoneNameStyle = 5
	// NSTimeZoneNameStyleShortStandard - Specifies a short name style. For example, “CST” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/shortStandard
	NSTimeZoneNameStyleShortStandard NSTimeZoneNameStyle = 1
	// NSTimeZoneNameStyleStandard - Specifies a standard name style. For example, “Central Standard Time” for Central Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle/standard
	NSTimeZoneNameStyleStandard NSTimeZoneNameStyle = 0
)

// NSURLBookmarkCreationOptions - Options used when creating bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions
type NSURLBookmarkCreationOptions uint

const (
	// NSURLBookmarkCreationMinimalBookmark - Specifies that when creating a bookmark, it includes minimal information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/minimalBookmark
	NSURLBookmarkCreationMinimalBookmark NSURLBookmarkCreationOptions = 0
	// NSURLBookmarkCreationPreferFileIDResolution - Specifies that when creating a bookmark, upon resolution, its embedded file ID takes precedence over other sources of information (file system path, for example) when there’s a conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/preferFileIDResolution
	NSURLBookmarkCreationPreferFileIDResolution NSURLBookmarkCreationOptions = 0
	// NSURLBookmarkCreationSecurityScopeAllowOnlyReadAccess - Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read-only access to a file-system resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/securityScopeAllowOnlyReadAccess
	NSURLBookmarkCreationSecurityScopeAllowOnlyReadAccess NSURLBookmarkCreationOptions = 2
	// NSURLBookmarkCreationSuitableForBookmarkFile - Specifies that the bookmark data includes the required properties for creating Finder alias files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/suitableForBookmarkFile
	NSURLBookmarkCreationSuitableForBookmarkFile NSURLBookmarkCreationOptions = 0
	// NSURLBookmarkCreationWithSecurityScope - Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read/write access to a file-system resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/withSecurityScope
	NSURLBookmarkCreationWithSecurityScope NSURLBookmarkCreationOptions = 1
	// NSURLBookmarkCreationWithoutImplicitSecurityScope - Prevents inclusion of a bookmark’s implicit ephemeral security scope, when creating one without security scope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/withoutImplicitSecurityScope
	NSURLBookmarkCreationWithoutImplicitSecurityScope NSURLBookmarkCreationOptions = 3
)

// NSURLBookmarkResolutionOptions - Options used when resolving bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions
type NSURLBookmarkResolutionOptions uint

const (
	// NSURLBookmarkResolutionWithSecurityScope - Specifies that the security scope, applied to the bookmark when it was created, should be used during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withSecurityScope
	NSURLBookmarkResolutionWithSecurityScope NSURLBookmarkResolutionOptions = 0
	// NSURLBookmarkResolutionWithoutImplicitStartAccessing - A property that specifies that resolution doesn’t implicitly start accessing the ephemeral security-scoped resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutImplicitStartAccessing
	NSURLBookmarkResolutionWithoutImplicitStartAccessing NSURLBookmarkResolutionOptions = 1
	// NSURLBookmarkResolutionWithoutMounting - Specifies that no volume should be mounted during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutMounting
	NSURLBookmarkResolutionWithoutMounting NSURLBookmarkResolutionOptions = 0
	// NSURLBookmarkResolutionWithoutUI - Specifies that no UI feedback should accompany resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withoutUI
	NSURLBookmarkResolutionWithoutUI NSURLBookmarkResolutionOptions = 0
)

// NSURLErrorNetworkUnavailableReason - An enumeration of reasons why a task couldn’t satisfy networking constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLErrorNetworkUnavailableReason
type NSURLErrorNetworkUnavailableReason uint

// NSURLRequestCachePolicy - The constants used to specify interaction with the cached responses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum
type NSURLRequestCachePolicy uint

const (
	// NSURLRequestReloadIgnoringLocalAndRemoteCacheData - Ignore local cache data, and instruct proxies and other intermediates to disregard their caches so far as the protocol allows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum/reloadIgnoringLocalAndRemoteCacheData
	NSURLRequestReloadIgnoringLocalAndRemoteCacheData NSURLRequestCachePolicy = 0
)

// NSURLRequestNetworkServiceType - Constants that specify how a request uses network resources.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum
type NSURLRequestNetworkServiceType uint

// NSURLSessionWebSocketMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType
type NSURLSessionWebSocketMessageType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType/NSURLSessionWebSocketMessageTypeData
	NSURLSessionWebSocketMessageTypeData NSURLSessionWebSocketMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType/NSURLSessionWebSocketMessageTypeString
	NSURLSessionWebSocketMessageTypeString NSURLSessionWebSocketMessageType = 1
)

// NSXPCConnectionOptions - Options that you can pass to a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/Options
type NSXPCConnectionOptions uint

// NSNetServicesError - These constants identify errors that can occur when accessing net services.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum
type NSNetServicesError uint

const (
	// NSNetServicesActivityInProgress - The net service cannot process the request at this time. No additional information about the network state is known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/activityInProgress
	NSNetServicesActivityInProgress NSNetServicesError = -72003
	// NSNetServicesBadArgumentError - An invalid argument was used when creating the   object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/badArgumentError
	NSNetServicesBadArgumentError NSNetServicesError = -72004
	// NSNetServicesCancelledError - The client canceled the action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/cancelledError
	NSNetServicesCancelledError NSNetServicesError = -72005
	// NSNetServicesCollisionError - The service could not be published because the name is already in use. The name could be in use locally or on another system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/collisionError
	NSNetServicesCollisionError NSNetServicesError = -72001
	// NSNetServicesInvalidError - The net service was improperly configured.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/invalidError
	NSNetServicesInvalidError NSNetServicesError = -72006
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/missingRequiredConfigurationError
	NSNetServicesMissingRequiredConfigurationError NSNetServicesError = -72006
	// NSNetServicesNotFoundError - The service could not be found on the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/notFoundError
	NSNetServicesNotFoundError NSNetServicesError = -72002
	// NSNetServicesTimeoutError - The net service has timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/timeoutError
	NSNetServicesTimeoutError NSNetServicesError = -72007
	// NSNetServicesUnknownError - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum/unknownError
	NSNetServicesUnknownError NSNetServicesError = -72000
)

// NSNetServiceOptions - These constants specify options for a network service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options
type NSNetServiceOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options/listenForConnections
	NSNetServiceListenForConnections NSNetServiceOptions = 2
	// NSNetServiceNoAutoRename - Specifies that the network service should not rename itself in the event of a name collision.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/Options/noAutoRename
	NSNetServiceNoAutoRename NSNetServiceOptions = 1
)

// NSNotificationCoalescing - The constants that specify how notifications are coalesced.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing
type NSNotificationCoalescing uint

const (
	// NSNotificationNoCoalescing - Do not coalesce notifications in the queue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing/none
	NSNotificationNoCoalescing NSNotificationCoalescing = 0
)

// NSPostingStyle - The constants that specify when notifications are posted.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle
type NSPostingStyle uint

const (
	// NSPostASAP - The notification is posted at the end of the current notification callout or timer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle/asap
	NSPostASAP NSPostingStyle = 2
	// NSPostNow - The notification is posted immediately after coalescing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle/now
	NSPostNow NSPostingStyle = 3
	// NSPostWhenIdle - The notification is posted when the run loop is idle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle/whenIdle
	NSPostWhenIdle NSPostingStyle = 1
)

// NSOperationQueuePriority - These constants let you prioritize the order in which operations execute.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum
type NSOperationQueuePriority uint

const (
	// NSOperationQueuePriorityHigh - Operations receive high priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/high
	NSOperationQueuePriorityHigh NSOperationQueuePriority = 4
	// NSOperationQueuePriorityLow - Operations receive low priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/low
	NSOperationQueuePriorityLow NSOperationQueuePriority = -4
	// NSOperationQueuePriorityNormal - Operations receive the normal priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/normal
	NSOperationQueuePriorityNormal NSOperationQueuePriority = 0
	// NSOperationQueuePriorityVeryHigh - Operations receive very high priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/veryHigh
	NSOperationQueuePriorityVeryHigh NSOperationQueuePriority = 8
	// NSOperationQueuePriorityVeryLow - Operations receive very low priority for execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum/veryLow
	NSOperationQueuePriorityVeryLow NSOperationQueuePriority = -8
)

// NSActivityOptions - Option flags used with 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions
type NSActivityOptions uint

const (
	// NSActivityAnimationTrackingEnabled - A flag to track the activity with an animation signpost interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/animationTrackingEnabled
	NSActivityAnimationTrackingEnabled NSActivityOptions = 0
	// NSActivityBackground - A flag to indicate the app has initiated some kind of work, but not as the direct result of user request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/background
	NSActivityBackground NSActivityOptions = 0
	// NSActivityLatencyCritical - A flag to indicate the activity requires the highest amount of timer and I/O precision available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/latencyCritical
	NSActivityLatencyCritical NSActivityOptions = 0
	// NSActivitySuddenTerminationDisabled - A flag to prevent sudden termination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/suddenTerminationDisabled
	NSActivitySuddenTerminationDisabled NSActivityOptions = 0
	// NSActivityUserInitiated - A flag to indicate the app is performing a user-requested action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/userInitiated
	NSActivityUserInitiated NSActivityOptions = 0
	// NSActivityUserInitiatedAllowingIdleSystemSleep - A flag to indicate the app is performing a user-requested action, but that the system can sleep on idle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/userInitiatedAllowingIdleSystemSleep
	NSActivityUserInitiatedAllowingIdleSystemSleep NSActivityOptions = 0
	// NSActivityUserInteractive - A flag to indicate the app is responding to user interaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/userInteractive
	NSActivityUserInteractive NSActivityOptions = 1
)

// NSProcessInfoThermalState - Values used to indicate the system’s thermal state.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum
type NSProcessInfoThermalState uint

const (
	NSProcessInfoThermalStateNominal NSProcessInfoThermalState = 0
	NSProcessInfoThermalStateFair NSProcessInfoThermalState = 1
	NSProcessInfoThermalStateSerious NSProcessInfoThermalState = 2
	NSProcessInfoThermalStateCritical NSProcessInfoThermalState = 3
)

// NSQualityOfService - Constants that indicate the nature and importance of work to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService
type NSQualityOfService uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/background
	NSQualityOfServiceBackground NSQualityOfService = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/default
	NSQualityOfServiceDefault NSQualityOfService = -1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/userInitiated
	NSQualityOfServiceUserInitiated NSQualityOfService = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/userInteractive
	NSQualityOfServiceUserInteractive NSQualityOfService = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/QualityOfService/utility
	NSQualityOfServiceUtility NSQualityOfService = 0
)

// NSURLSessionAuthChallengeDisposition - Constants passed by session or task delegates to the provided continuation block in response to an authentication challenge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition
type NSURLSessionAuthChallengeDisposition uint

const (
	// NSURLSessionAuthChallengeRejectProtectionSpace - Reject this challenge, and call the authentication delegate method again with the next authentication protection space. The provided credential parameter is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/rejectProtectionSpace
	NSURLSessionAuthChallengeRejectProtectionSpace NSURLSessionAuthChallengeDisposition = 3
)

// NSURLSessionDelayedRequestDisposition - The action to take on a delayed URL session task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition
type NSURLSessionDelayedRequestDisposition uint

const (
	// NSURLSessionDelayedRequestCancel - A disposition indicating that the task should be canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/cancel
	NSURLSessionDelayedRequestCancel NSURLSessionDelayedRequestDisposition = 2
	// NSURLSessionDelayedRequestContinueLoading - A disposition indicating that the task should proceed with the original request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/continueLoading
	NSURLSessionDelayedRequestContinueLoading NSURLSessionDelayedRequestDisposition = 0
	// NSURLSessionDelayedRequestUseNewRequest - A disposition indicating that the task should use a new request to perform the network load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/useNewRequest
	NSURLSessionDelayedRequestUseNewRequest NSURLSessionDelayedRequestDisposition = 1
)

// NSURLSessionResponseDisposition - Constants indicating how a data or upload session should proceed after receiving the initial headers.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition
type NSURLSessionResponseDisposition uint

const (
	// NSURLSessionResponseAllow - Allow the load operation to continue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/allow
	NSURLSessionResponseAllow NSURLSessionResponseDisposition = 1
	// NSURLSessionResponseBecomeDownload - Convert the response for this request to use a  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/becomeDownload
	NSURLSessionResponseBecomeDownload NSURLSessionResponseDisposition = 2
	// NSURLSessionResponseBecomeStream - Convert the response for this request to use a  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/becomeStream
	NSURLSessionResponseBecomeStream NSURLSessionResponseDisposition = 3
	// NSURLSessionResponseCancel - Cancel the load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/cancel
	NSURLSessionResponseCancel NSURLSessionResponseDisposition = 0
)

// NSURLSessionMultipathServiceType - Constants that specify the type of service that Multipath TCP uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum
type NSURLSessionMultipathServiceType uint

const (
	// NSURLSessionMultipathServiceTypeAggregate - A service that aggregates the capacities of other Multipath options in an attempt to increase throughput and minimize latency.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/aggregate
	NSURLSessionMultipathServiceTypeAggregate NSURLSessionMultipathServiceType = 0
	// NSURLSessionMultipathServiceTypeHandover - A Multipath TCP service that provides seamless handover between Wi-Fi and cellular in order to preserve the connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/handover
	NSURLSessionMultipathServiceTypeHandover NSURLSessionMultipathServiceType = 0
	// NSURLSessionMultipathServiceTypeInteractive - A service whereby Multipath TCP attempts to use the lowest-latency interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/interactive
	NSURLSessionMultipathServiceTypeInteractive NSURLSessionMultipathServiceType = 0
	// NSURLSessionMultipathServiceTypeNone - The default service type indicating that Multipath TCP should not be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum/none
	NSURLSessionMultipathServiceTypeNone NSURLSessionMultipathServiceType = 0
)

// NSURLSessionTaskState - Constants for determining the current state of a task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum
type NSURLSessionTaskState uint

const (
	NSURLSessionTaskStateRunning NSURLSessionTaskState = 0
	NSURLSessionTaskStateSuspended NSURLSessionTaskState = 1
	NSURLSessionTaskStateCanceling NSURLSessionTaskState = 2
	NSURLSessionTaskStateCompleted NSURLSessionTaskState = 3
)

// NSXMLDTDNodeKind - The type defined for the constants that specify the kind and subkind of DTD declaration represented by an 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum
type NSXMLDTDNodeKind uint

const (
	// NSXMLElementDeclarationAnyKind - Identifies an   element declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/anyDeclaration
	NSXMLElementDeclarationAnyKind NSXMLDTDNodeKind = 18
	// NSXMLAttributeCDATAKind - Identifies an attribute-list declaration with a   (character data) value type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/cdataAttribute
	NSXMLAttributeCDATAKind NSXMLDTDNodeKind = 6
	// NSXMLElementDeclarationElementKind - Identifies a declaration of an element with child elements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/elementDeclaration
	NSXMLElementDeclarationElementKind NSXMLDTDNodeKind = 20
	// NSXMLElementDeclarationEmptyKind - Identifies a declaration ( ) of an empty element.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/emptyDeclaration
	NSXMLElementDeclarationEmptyKind NSXMLDTDNodeKind = 17
	// NSXMLAttributeEntitiesKind - Identifies an attribute-list declaration with an   value type (refers to multiple unparsed entities declared elsewhere in document).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/entitiesAttribute
	NSXMLAttributeEntitiesKind NSXMLDTDNodeKind = 11
	// NSXMLAttributeEntityKind - Identifies an attribute-list declaration with an   value type (refers to unparsed entity declared in document).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/entityAttribute
	NSXMLAttributeEntityKind NSXMLDTDNodeKind = 10
	// NSXMLAttributeEnumerationKind - Identifies an attribute-list declaration with an enumeration value type (list of all possible values).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/enumerationAttribute
	NSXMLAttributeEnumerationKind NSXMLDTDNodeKind = 14
	// NSXMLEntityGeneralKind - Identifies a general entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/general
	NSXMLEntityGeneralKind NSXMLDTDNodeKind = 1
	// NSXMLAttributeIDKind - Identifies an attribute-list declaration with an   value type (per-document unique element name).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/idAttribute
	NSXMLAttributeIDKind NSXMLDTDNodeKind = 7
	// NSXMLAttributeIDRefKind - Identifies an attribute-list declaration with an   value type (refers to element   type).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/idRefAttribute
	NSXMLAttributeIDRefKind NSXMLDTDNodeKind = 8
	// NSXMLAttributeIDRefsKind - Identifies an attribute-list declaration with an   value type (refers to multiple elements of   type).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/idRefsAttribute
	NSXMLAttributeIDRefsKind NSXMLDTDNodeKind = 9
	// NSXMLElementDeclarationMixedKind - Identifies a declaration of an element with mixed content ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/mixedDeclaration
	NSXMLElementDeclarationMixedKind NSXMLDTDNodeKind = 19
	// NSXMLAttributeNMTokenKind - Identifies an attribute-list declaration with a   value type (name token).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/nmTokenAttribute
	NSXMLAttributeNMTokenKind NSXMLDTDNodeKind = 12
	// NSXMLAttributeNMTokensKind - Identifies an attribute-list declaration with a   value type (multiple name tokens)
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/nmTokensAttribute
	NSXMLAttributeNMTokensKind NSXMLDTDNodeKind = 13
	// NSXMLAttributeNotationKind - Identifies an attribute-list declaration with a   value type (name of declared notation).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/notationAttribute
	NSXMLAttributeNotationKind NSXMLDTDNodeKind = 15
	// NSXMLEntityParameterKind - Identifies a parameter entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/parameter
	NSXMLEntityParameterKind NSXMLDTDNodeKind = 4
	// NSXMLEntityParsedKind - Identifies a parsed entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/parsed
	NSXMLEntityParsedKind NSXMLDTDNodeKind = 2
	// NSXMLEntityPredefined - Identifies a predefined entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/predefined
	NSXMLEntityPredefined NSXMLDTDNodeKind = 5
	// NSXMLElementDeclarationUndefinedKind - Identifies an undefined element declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/undefinedDeclaration
	NSXMLElementDeclarationUndefinedKind NSXMLDTDNodeKind = 16
	// NSXMLEntityUnparsedKind - Identifies an unparsed entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/unparsed
	NSXMLEntityUnparsedKind NSXMLDTDNodeKind = 3
)

// NSXMLNodeKind enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum
type NSXMLNodeKind uint

const (
	// NSXMLDTDKind - Specifies a document-type declaration (DTD) node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/DTDKind
	NSXMLDTDKind NSXMLNodeKind = 8
	// NSXMLAttributeKind - Specifies an attribute node
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/attribute
	NSXMLAttributeKind NSXMLNodeKind = 3
	// NSXMLAttributeDeclarationKind - Specifies an attribute-list declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/attributeDeclaration
	NSXMLAttributeDeclarationKind NSXMLNodeKind = 10
	// NSXMLCommentKind - Specifies a comment node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/comment
	NSXMLCommentKind NSXMLNodeKind = 6
	// NSXMLDocumentKind - Specifies a document node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/document
	NSXMLDocumentKind NSXMLNodeKind = 1
	// NSXMLElementKind - Specifies an element node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/element
	NSXMLElementKind NSXMLNodeKind = 2
	// NSXMLElementDeclarationKind - Specifies an element declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/elementDeclaration
	NSXMLElementDeclarationKind NSXMLNodeKind = 11
	// NSXMLEntityDeclarationKind - Specifies an entity-declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/entityDeclaration
	NSXMLEntityDeclarationKind NSXMLNodeKind = 9
	// NSXMLInvalidKind - Indicates a node object created without a valid kind being specified (as returned by the   method).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/invalid
	NSXMLInvalidKind NSXMLNodeKind = 0
	// NSXMLNamespaceKind - Specifies a namespace node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/namespace
	NSXMLNamespaceKind NSXMLNodeKind = 4
	// NSXMLNotationDeclarationKind - Specifies a notation declaration node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/notationDeclaration
	NSXMLNotationDeclarationKind NSXMLNodeKind = 12
	// NSXMLProcessingInstructionKind - Specifies a processing-instruction node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/processingInstruction
	NSXMLProcessingInstructionKind NSXMLNodeKind = 5
	// NSXMLTextKind - Specifies a text node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/text
	NSXMLTextKind NSXMLNodeKind = 7
)

// NSXMLNodeOptions - These constants are input and output options for all 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options
type NSXMLNodeOptions uint

const (
	// NSXMLNodePreserveDTD - Specifies that declarations in a DTD should be preserved until it the DTD is modified. For example, parameter entities are by default expanded; with this option, they are written out as they originally occur in the DTD.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodePreserveDTD
	NSXMLNodePreserveDTD NSXMLNodeOptions = 1
	// NSXMLNodePreserveQuotes - Specifies that the quoting style used in the input XML (single or double quotes) be preserved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodePreserveQuotes
	NSXMLNodePreserveQuotes NSXMLNodeOptions = 0
	// NSXMLNodeUseSingleQuotes - Requests that NSXML use single quotes for the value of an attribute or namespace node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodeUseSingleQuotes
	NSXMLNodeUseSingleQuotes NSXMLNodeOptions = 1
)


