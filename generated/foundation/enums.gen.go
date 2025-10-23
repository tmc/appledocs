// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
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
	// FormattingContextBeginningOfSentence - The formatting context for the beginning of a sentence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/Context/beginningOfSentence
	FormattingContextBeginningOfSentence FormattingContext = 4
	// FormattingContextDynamic - A formatting context determined automatically at runtime.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/Context/dynamic
	FormattingContextDynamic FormattingContext = 1
	// FormattingContextListItem - The formatting context for a list or menu item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/Context/listItem
	FormattingContextListItem FormattingContext = 3
	// FormattingContextMiddleOfSentence - The formatting context for the middle of a sentence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/Context/middleOfSentence
	FormattingContextMiddleOfSentence FormattingContext = 5
	// FormattingContextStandalone - The formatting context for stand-alone usage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/Context/standalone
	FormattingContextStandalone FormattingContext = 2
	// FormattingContextUnknown - An unknown formatting context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/Context/unknown
	FormattingContextUnknown FormattingContext = 0
)

// AttributedStringEnumerationOptions - Options for enumerating attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/EnumerationOptions
type AttributedStringEnumerationOptions uint

const (
	// AttributedStringEnumerationLongestEffectiveRangeNotRequired - If   option is supplied, then the longest effective range computation is not performed; the blocks may be invoked with consecutive attribute runs that have the same value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/EnumerationOptions/longestEffectiveRangeNotRequired
	AttributedStringEnumerationLongestEffectiveRangeNotRequired AttributedStringEnumerationOptions = 1048576
	// AttributedStringEnumerationReverse - Causes the enumeration to occur in reverse.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/EnumerationOptions/reverse
	AttributedStringEnumerationReverse AttributedStringEnumerationOptions = 2
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

// AttributedStringMarkdownInterpretedSyntax - A type that represents the syntax for intepreting a Markdown string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownInterpretedSyntax
type AttributedStringMarkdownInterpretedSyntax int

const (
	// AttributedStringMarkdownInterpretedSyntaxFull - A syntax value that interprets the full Markdown syntax and produces all relevant attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownInterpretedSyntax/NSAttributedStringMarkdownInterpretedSyntaxFull
	AttributedStringMarkdownInterpretedSyntaxFull AttributedStringMarkdownInterpretedSyntax = 0
	// AttributedStringMarkdownInterpretedSyntaxInlineOnly - A syntax value that parses all Markdown text, but interprets only attributes that apply to inline spans.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownInterpretedSyntax/NSAttributedStringMarkdownInterpretedSyntaxInlineOnly
	AttributedStringMarkdownInterpretedSyntaxInlineOnly AttributedStringMarkdownInterpretedSyntax = 1
	// AttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace - A syntax value that parses all Markdown text, but interprets only attributes that apply to inline spans, perserving white space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownInterpretedSyntax/NSAttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace
	AttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace AttributedStringMarkdownInterpretedSyntax = 2
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
	// WeekdayCalendarUnit - Specifies the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekdayCalendarUnit
	WeekdayCalendarUnit CalendarUnit = 16
	// WeekdayOrdinalCalendarUnit - Specifies the ordinal weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekdayOrdinalCalendarUnit
	WeekdayOrdinalCalendarUnit CalendarUnit = 17
	// WeekOfMonthCalendarUnit - Specifies the original week of a month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekOfMonthCalendarUnit
	WeekOfMonthCalendarUnit CalendarUnit = 19
	// WeekOfYearCalendarUnit - Specifies the original week of the year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSWeekOfYearCalendarUnit
	WeekOfYearCalendarUnit CalendarUnit = 20
	// YearCalendarUnit - Specifies the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSYearCalendarUnit
	YearCalendarUnit CalendarUnit = 9
	// YearForWeekOfYearCalendarUnit - Specifies the year when the calendar is being interpreted as a week-based calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/NSYearForWeekOfYearCalendarUnit
	YearForWeekOfYearCalendarUnit CalendarUnit = 21
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
	// CalendarUnitWeekday - Identifier for the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekday
	CalendarUnitWeekday CalendarUnit = 0
	// CalendarUnitWeekdayOrdinal - Identifier for the ordinal weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekdayOrdinal
	CalendarUnitWeekdayOrdinal CalendarUnit = 0
	// CalendarUnitWeekOfMonth - Identifier for the week of the month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekOfMonth
	CalendarUnitWeekOfMonth CalendarUnit = 1
	// CalendarUnitWeekOfYear - Identifier for the week of the year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/weekOfYear
	CalendarUnitWeekOfYear CalendarUnit = 2
	// CalendarUnitYear - Identifier for the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/year
	CalendarUnitYear CalendarUnit = 0
	// CalendarUnitYearForWeekOfYear - Identifier for the week-counting year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit/yearForWeekOfYear
	CalendarUnitYearForWeekOfYear CalendarUnit = 3
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
	// FileCoordinatorWritingForMerging - When this constant is specified, the file coordinator calls the   method of relevant file presenters to give them a chance to save their changes before your code makes its changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forMerging
	FileCoordinatorWritingForMerging FileCoordinatorWritingOptions = 4
	// FileCoordinatorWritingForMoving - When specified for a directory item, the file coordinator waits for already running read and write operations of the directory’s contents, which were themselves initiated through a file coordinator, to finish before moving the directory. Queued, but not executing, read and write operations on the directory’s contents wait until the move operation finishes. This option has no effect on files. You can safely use it when moving file-system items without checking to see whether those items are files or directories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forMoving
	FileCoordinatorWritingForMoving FileCoordinatorWritingOptions = 2
	// FileCoordinatorWritingForReplacing - Specifies whether the act of writing to the file involves actually replacing the file with a different file (or directory). If the current file coordinator is waiting for another object to move or rename the file, this option treats the operation as the creation of a new file (instead of as the replacement of the old file); otherwise, this constant causes the same behavior as the   constant. Use this method when the moving or creating an item should replace any item currently stored at that location. To avoid a race condition, use it regardless of whether an item is actually in the way before the writing begins. Do not use this method when simply updating the contents of the existing file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions/forReplacing
	FileCoordinatorWritingForReplacing FileCoordinatorWritingOptions = 8
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

// GrammaticalDefiniteness enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness
type GrammaticalDefiniteness uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness/indefinite
	GrammaticalDefinitenessIndefinite GrammaticalDefiniteness = 1
)

// ItemProviderRepresentationVisibility - Specifications that control which categories of processes can see an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility
type ItemProviderRepresentationVisibility uint

const (
	// ItemProviderRepresentationVisibilityGroup - A representation visibility specification confining item visibility to the app’s app group.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility/group
	ItemProviderRepresentationVisibilityGroup ItemProviderRepresentationVisibility = 2
)

// PresentationIntentKind - An enumeration of intended display styles for blocks of text like paragraphs, lists, and code blocks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind
type PresentationIntentKind int

const (
	// PresentationIntentKindListItem - A presentation style for a list of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindListItem
	PresentationIntentKindListItem PresentationIntentKind = 4
)

// PresentationIntentTableColumnAlignment - An enumeration of values for aligning the contents of table columns.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment
type PresentationIntentTableColumnAlignment int

const (
	// PresentationIntentTableColumnAlignmentCenter - A presentation style for columns with center-aligned text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment/NSPresentationIntentTableColumnAlignmentCenter
	PresentationIntentTableColumnAlignmentCenter PresentationIntentTableColumnAlignment = 1
	// PresentationIntentTableColumnAlignmentLeft - A presentation style for columns with left-aligned text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment/NSPresentationIntentTableColumnAlignmentLeft
	PresentationIntentTableColumnAlignmentLeft PresentationIntentTableColumnAlignment = 0
	// PresentationIntentTableColumnAlignmentRight - A presentation style for columns with right-aligned text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment/NSPresentationIntentTableColumnAlignmentRight
	PresentationIntentTableColumnAlignmentRight PresentationIntentTableColumnAlignment = 2
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
	// StringDrawingOptionsResolvesNaturalAlignmentWithBaseWritingDirection - Specifies the behavior for resolving   to the visual alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions/optionsResolvesNaturalAlignmentWithBaseWritingDirection
	StringDrawingOptionsResolvesNaturalAlignmentWithBaseWritingDirection StringDrawingOptions = 0
	// StringDrawingTruncatesLastVisibleLine - Truncates and adds the ellipsis character to the last visible line if the text doesn’t fit into the specified bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions/truncatesLastVisibleLine
	StringDrawingTruncatesLastVisibleLine StringDrawingOptions = 0
	// StringDrawingUsesDeviceMetrics - Uses image glyph bounds instead of typographic bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions/usesDeviceMetrics
	StringDrawingUsesDeviceMetrics StringDrawingOptions = 0
	// StringDrawingUsesFontLeading - Uses the font leading for calculating line heights.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions/usesFontLeading
	StringDrawingUsesFontLeading StringDrawingOptions = 0
	// StringDrawingUsesLineFragmentOrigin - Uses the line fragment origin instead of the baseline origin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions/usesLineFragmentOrigin
	StringDrawingUsesLineFragmentOrigin StringDrawingOptions = 0
)

// XMLParserError - The following error codes are defined by 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/ErrorCode
type XMLParserError uint

const (
	XMLParserInternalError XMLParserError = 1
	XMLParserOutOfMemoryError XMLParserError = 2
	XMLParserDocumentStartError XMLParserError = 3
	XMLParserEmptyDocumentError XMLParserError = 4
	XMLParserPrematureDocumentEndError XMLParserError = 5
	XMLParserInvalidHexCharacterRefError XMLParserError = 6
	XMLParserInvalidDecimalCharacterRefError XMLParserError = 7
	XMLParserInvalidCharacterRefError XMLParserError = 8
	XMLParserInvalidCharacterError XMLParserError = 9
	XMLParserCharacterRefAtEOFError XMLParserError = 10
	XMLParserCharacterRefInPrologError XMLParserError = 11
	XMLParserCharacterRefInEpilogError XMLParserError = 12
	XMLParserCharacterRefInDTDError XMLParserError = 13
	XMLParserEntityRefAtEOFError XMLParserError = 14
	XMLParserEntityRefInPrologError XMLParserError = 15
	XMLParserEntityRefInEpilogError XMLParserError = 16
	XMLParserEntityRefInDTDError XMLParserError = 17
	XMLParserParsedEntityRefAtEOFError XMLParserError = 18
	XMLParserParsedEntityRefInPrologError XMLParserError = 19
	XMLParserParsedEntityRefInEpilogError XMLParserError = 20
	XMLParserParsedEntityRefInInternalSubsetError XMLParserError = 21
	XMLParserEntityReferenceWithoutNameError XMLParserError = 22
	XMLParserEntityReferenceMissingSemiError XMLParserError = 23
	XMLParserParsedEntityRefNoNameError XMLParserError = 24
	XMLParserParsedEntityRefMissingSemiError XMLParserError = 25
	XMLParserUndeclaredEntityError XMLParserError = 26
	XMLParserUnparsedEntityError XMLParserError = 28
	XMLParserEntityIsExternalError XMLParserError = 29
	XMLParserEntityIsParameterError XMLParserError = 30
	XMLParserUnknownEncodingError XMLParserError = 31
	XMLParserEncodingNotSupportedError XMLParserError = 32
	XMLParserStringNotStartedError XMLParserError = 33
	XMLParserStringNotClosedError XMLParserError = 34
	XMLParserNamespaceDeclarationError XMLParserError = 35
	XMLParserEntityNotStartedError XMLParserError = 36
	XMLParserEntityNotFinishedError XMLParserError = 37
	XMLParserLessThanSymbolInAttributeError XMLParserError = 38
	XMLParserAttributeNotStartedError XMLParserError = 39
	XMLParserAttributeNotFinishedError XMLParserError = 40
	XMLParserAttributeHasNoValueError XMLParserError = 41
	XMLParserAttributeRedefinedError XMLParserError = 42
	XMLParserLiteralNotStartedError XMLParserError = 43
	XMLParserLiteralNotFinishedError XMLParserError = 44
	XMLParserCommentNotFinishedError XMLParserError = 45
	XMLParserProcessingInstructionNotStartedError XMLParserError = 46
	XMLParserProcessingInstructionNotFinishedError XMLParserError = 47
	XMLParserNotationNotStartedError XMLParserError = 48
	XMLParserNotationNotFinishedError XMLParserError = 49
	XMLParserAttributeListNotStartedError XMLParserError = 50
	XMLParserAttributeListNotFinishedError XMLParserError = 51
	XMLParserMixedContentDeclNotStartedError XMLParserError = 52
	XMLParserMixedContentDeclNotFinishedError XMLParserError = 53
	XMLParserElementContentDeclNotStartedError XMLParserError = 54
	XMLParserElementContentDeclNotFinishedError XMLParserError = 55
	XMLParserXMLDeclNotStartedError XMLParserError = 56
	XMLParserXMLDeclNotFinishedError XMLParserError = 57
	XMLParserConditionalSectionNotStartedError XMLParserError = 58
	XMLParserConditionalSectionNotFinishedError XMLParserError = 59
	XMLParserExternalSubsetNotFinishedError XMLParserError = 60
	XMLParserDOCTYPEDeclNotFinishedError XMLParserError = 61
	XMLParserMisplacedCDATAEndStringError XMLParserError = 62
	XMLParserCDATANotFinishedError XMLParserError = 63
	XMLParserMisplacedXMLDeclarationError XMLParserError = 64
	XMLParserSpaceRequiredError XMLParserError = 65
	XMLParserSeparatorRequiredError XMLParserError = 66
	XMLParserNMTOKENRequiredError XMLParserError = 67
	XMLParserNAMERequiredError XMLParserError = 68
	XMLParserPCDATARequiredError XMLParserError = 69
	XMLParserURIRequiredError XMLParserError = 70
	XMLParserPublicIdentifierRequiredError XMLParserError = 71
	XMLParserLTRequiredError XMLParserError = 72
	XMLParserGTRequiredError XMLParserError = 73
	XMLParserLTSlashRequiredError XMLParserError = 74
	XMLParserEqualExpectedError XMLParserError = 75
	XMLParserTagNameMismatchError XMLParserError = 76
	XMLParserUnfinishedTagError XMLParserError = 77
	XMLParserStandaloneValueError XMLParserError = 78
	XMLParserInvalidEncodingNameError XMLParserError = 79
	XMLParserCommentContainsDoubleHyphenError XMLParserError = 80
	XMLParserInvalidEncodingError XMLParserError = 81
	XMLParserExternalStandaloneEntityError XMLParserError = 82
	XMLParserInvalidConditionalSectionError XMLParserError = 83
	XMLParserEntityValueRequiredError XMLParserError = 84
	XMLParserNotWellBalancedError XMLParserError = 85
	XMLParserExtraContentError XMLParserError = 86
	XMLParserInvalidCharacterInEntityError XMLParserError = 87
	XMLParserParsedEntityRefInInternalError XMLParserError = 88
	XMLParserEntityRefLoopError XMLParserError = 89
	XMLParserEntityBoundaryError XMLParserError = 90
	XMLParserInvalidURIError XMLParserError = 91
	XMLParserURIFragmentError XMLParserError = 92
	XMLParserNoDTDError XMLParserError = 94
	XMLParserDelegateAbortedParseError XMLParserError = 512
)

// XMLParserExternalEntityResolvingPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/ExternalEntityResolvingPolicy-swift.enum
type XMLParserExternalEntityResolvingPolicy uint

const (
	XMLParserResolveExternalEntitiesNever XMLParserExternalEntityResolvingPolicy = 0
	XMLParserResolveExternalEntitiesNoNetwork XMLParserExternalEntityResolvingPolicy = 1
	XMLParserResolveExternalEntitiesSameOriginOnly XMLParserExternalEntityResolvingPolicy = 2
	XMLParserResolveExternalEntitiesAlways XMLParserExternalEntityResolvingPolicy = 3
)

// AlignmentOptions - Values representing alignment operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions
type AlignmentOptions uint

const (
	// AlignAllEdgesInward - Aligns all edges inward. This is the same as  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignAllEdgesInward
	AlignAllEdgesInward AlignmentOptions = 0
	// AlignAllEdgesNearest - Aligns all edges to the nearest value. This is the same as  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignAllEdgesNearest
	AlignAllEdgesNearest AlignmentOptions = 0
	// AlignAllEdgesOutward - Aligns all edges outwards. This is the same as  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignAllEdgesOutward
	AlignAllEdgesOutward AlignmentOptions = 0
	// AlignHeightInward - Specifies that alignment of the height should be to the nearest inward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignHeightInward
	AlignHeightInward AlignmentOptions = 32
	// AlignHeightNearest - Specifies that alignment of the height should be to the nearest integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignHeightNearest
	AlignHeightNearest AlignmentOptions = 2097152
	// AlignHeightOutward - Specifies that alignment of the height should be to the nearest outward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignHeightOutward
	AlignHeightOutward AlignmentOptions = 8192
	// AlignMaxXInward - Specifies that alignment of the maximum X coordinate should be to the nearest inward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMaxXInward
	AlignMaxXInward AlignmentOptions = 4
	// AlignMaxXNearest - Specifies that alignment of the maximum X coordinate should be to the nearest integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMaxXNearest
	AlignMaxXNearest AlignmentOptions = 262144
	// AlignMaxXOutward - Specifies that alignment of the maximum X coordinate should be to the nearest outward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMaxXOutward
	AlignMaxXOutward AlignmentOptions = 1024
	// AlignMaxYInward - Specifies that alignment of the maximum X coordinate should be to the nearest inward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMaxYInward
	AlignMaxYInward AlignmentOptions = 8
	// AlignMaxYNearest - Specifies that alignment of the maximum Y coordinate should be to the nearest integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMaxYNearest
	AlignMaxYNearest AlignmentOptions = 524288
	// AlignMaxYOutward - Specifies that alignment of the maximum Y coordinate should be to the nearest outward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMaxYOutward
	AlignMaxYOutward AlignmentOptions = 2048
	// AlignMinXInward - Specifies that alignment of the minimum X coordinate should be to the nearest inward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMinXInward
	AlignMinXInward AlignmentOptions = 1
	// AlignMinXNearest - Specifies that alignment of the minimum X coordinate should be to the nearest integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMinXNearest
	AlignMinXNearest AlignmentOptions = 65536
	// AlignMinXOutward - Specifies that alignment of the minimum X coordinate should be to the nearest outward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMinXOutward
	AlignMinXOutward AlignmentOptions = 256
	// AlignMinYInward - Specifies that alignment of the minimum Y coordinate should be to the nearest inward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMinYInward
	AlignMinYInward AlignmentOptions = 2
	// AlignMinYNearest - Specifies that alignment of the minimum Y coordinate should be to the nearest integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMinYNearest
	AlignMinYNearest AlignmentOptions = 131072
	// AlignMinYOutward - Specifies that alignment of the minimum Y coordinate should be to the nearest outward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignMinYOutward
	AlignMinYOutward AlignmentOptions = 512
	// AlignRectFlipped - This option should be included  if the rectangle is in a flipped coordinate system. This allows 0.5 to be treated in a visually consistent way.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignRectFlipped
	AlignRectFlipped AlignmentOptions = -9223372036854775808
	// AlignWidthInward - Specifies that alignment of the width should be to the nearest inward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignWidthInward
	AlignWidthInward AlignmentOptions = 16
	// AlignWidthNearest - Specifies that alignment of the width should be to the nearest integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignWidthNearest
	AlignWidthNearest AlignmentOptions = 1048576
	// AlignWidthOutward - Specifies that alignment of the width should be to the nearest outward integral value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/AlignmentOptions/alignWidthOutward
	AlignWidthOutward AlignmentOptions = 4096
)

// DateComponentsFormatterUnitsStyle - Constants for specifying how to represent quantities of time.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/UnitsStyle-swift.enum
type DateComponentsFormatterUnitsStyle uint

const (
	// DateComponentsFormatterUnitsStyleAbbreviated - A style that uses the most abbreviated spelling for units of time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/UnitsStyle-swift.enum/abbreviated
	DateComponentsFormatterUnitsStyleAbbreviated DateComponentsFormatterUnitsStyle = 1
	// DateComponentsFormatterUnitsStyleBrief - A style that uses a shortened spelling for units of time that is shorter than  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/UnitsStyle-swift.enum/brief
	DateComponentsFormatterUnitsStyleBrief DateComponentsFormatterUnitsStyle = 5
	// DateComponentsFormatterUnitsStyleFull - A style that spells out the units of time, but not the quantities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/UnitsStyle-swift.enum/full
	DateComponentsFormatterUnitsStyleFull DateComponentsFormatterUnitsStyle = 3
	// DateComponentsFormatterUnitsStylePositional - A style that uses the position of a unit of time to identify its value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/UnitsStyle-swift.enum/positional
	DateComponentsFormatterUnitsStylePositional DateComponentsFormatterUnitsStyle = 0
	// DateComponentsFormatterUnitsStyleShort - A style that uses a shortened spelling for units.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/UnitsStyle-swift.enum/short
	DateComponentsFormatterUnitsStyleShort DateComponentsFormatterUnitsStyle = 2
	// DateComponentsFormatterUnitsStyleSpellOut - A style that spells out the units and quantities of time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/UnitsStyle-swift.enum/spellOut
	DateComponentsFormatterUnitsStyleSpellOut DateComponentsFormatterUnitsStyle = 4
)

// DateComponentsFormatterZeroFormattingBehavior - Formatting constants for when values contain zeroes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/ZeroFormattingBehavior-swift.struct
type DateComponentsFormatterZeroFormattingBehavior uint

const (
	// DateComponentsFormatterZeroFormattingBehaviorNone - No formatting behavior. This behavior prevents the dropping of zero values or adding of zeroes for padding. For example, with hours, minutes, and seconds displayed, the abbreviated value for one hour and 10 seconds is “1h 0m 10s”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponentsFormatterZeroFormattingBehavior/NSDateComponentsFormatterZeroFormattingBehaviorNone
	DateComponentsFormatterZeroFormattingBehaviorNone DateComponentsFormatterZeroFormattingBehavior = 0
	// DateComponentsFormatterZeroFormattingBehaviorDefault - The default formatting behavior. When using positional units, this behavior drops leading zeroes but pads middle and trailing values with zeros as needed. For example, with hours, minutes, and seconds displayed, the value for one hour and 10 seconds is “1:00:10”. For all other unit styles, this behavior drops all units whose values are 0. For example, when days, hours, minutes, and seconds are allowed, the abbreviated version of one hour and 10 seconds is displayed as “1h 10s”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/ZeroFormattingBehavior-swift.struct/default
	DateComponentsFormatterZeroFormattingBehaviorDefault DateComponentsFormatterZeroFormattingBehavior = 1
	// DateComponentsFormatterZeroFormattingBehaviorDropAll - The drop all zero units behavior. This behavior drops all units whose values are 0. For example, when days, hours, minutes, and seconds are allowed, the abbreviated version of one hour is displayed as “1h”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/ZeroFormattingBehavior-swift.struct/dropAll
	DateComponentsFormatterZeroFormattingBehaviorDropAll DateComponentsFormatterZeroFormattingBehavior = 0
	// DateComponentsFormatterZeroFormattingBehaviorDropLeading - The drop leading zeroes formatting behavior. Units whose values are 0 are dropped starting at the beginning of the sequence. Units continue to be dropped until a non-zero value is encountered. For example, when days, hours, minutes, and seconds are allowed, the abbreviated version of ten minutes is displayed as “10m 0s”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/ZeroFormattingBehavior-swift.struct/dropLeading
	DateComponentsFormatterZeroFormattingBehaviorDropLeading DateComponentsFormatterZeroFormattingBehavior = 2
	// DateComponentsFormatterZeroFormattingBehaviorDropMiddle - The drop middle zero units behavior. Units whose values are 0 are dropped from anywhere in the middle of a sequence. For example, when days, hours, minutes, and seconds are allowed, the abbreviated version of one hour, zero minutes, and five seconds is displayed as “0d 1h 5s”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/ZeroFormattingBehavior-swift.struct/dropMiddle
	DateComponentsFormatterZeroFormattingBehaviorDropMiddle DateComponentsFormatterZeroFormattingBehavior = 4
	// DateComponentsFormatterZeroFormattingBehaviorDropTrailing - The drop trailing zero units behavior. Units whose value is 0 are dropped starting at the end of the sequence. For example, when days, hours, minutes, and seconds are allowed, the abbreviated version of one hour is displayed as “0d 1h”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/ZeroFormattingBehavior-swift.struct/dropTrailing
	DateComponentsFormatterZeroFormattingBehaviorDropTrailing DateComponentsFormatterZeroFormattingBehavior = 8
	// DateComponentsFormatterZeroFormattingBehaviorPad - The add padding zeroes behavior. This behavior pads values with zeroes as appropriate. For example, consider the value of one hour formatted using the positional and abbreviated unit styles. When days, hours, minutes, and seconds are allowed, the value is displayed as “0d 1:00:00” using the positional style, and as “0d 1h 0m 0s” using the abbreviated style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/ZeroFormattingBehavior-swift.struct/pad
	DateComponentsFormatterZeroFormattingBehaviorPad DateComponentsFormatterZeroFormattingBehavior = 65536
)

// DateFormatterBehavior - Constants that specify the behavior 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Behavior
type DateFormatterBehavior uint

const (
	// DateFormatterBehavior10_0 - Specifies formatting behavior equivalent to that in OS X 10.0.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Behavior/behavior10_0
	DateFormatterBehavior10_0 DateFormatterBehavior = 1000
	// DateFormatterBehavior10_4 - Specifies formatting behavior equivalent for OS X 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Behavior/behavior10_4
	DateFormatterBehavior10_4 DateFormatterBehavior = 1040
	// DateFormatterBehaviorDefault - Specifies default formatting behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Behavior/default
	DateFormatterBehaviorDefault DateFormatterBehavior = 0
)

// DateFormatterStyle - The following constants specify predefined format styles for dates and times.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style
type DateFormatterStyle uint

const (
	// DateFormatterFullStyle - Specifies a full style with complete details, such as “Tuesday, April 12, 1952 AD” or “3:30:42 PM Pacific Standard Time”. Equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/full
	DateFormatterFullStyle DateFormatterStyle = 0
	// DateFormatterLongStyle - Specifies a long style, typically with full text, such as “November 23, 1937” or “3:30:32 PM PST”. Equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/long
	DateFormatterLongStyle DateFormatterStyle = 0
	// DateFormatterMediumStyle - Specifies a medium style, typically with abbreviated text, such as “Nov 23, 1937” or “3:30:32 PM”. Equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style/medium
	DateFormatterMediumStyle DateFormatterStyle = 0
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
	// DateIntervalFormatterFullStyle - A fully spelled out date or time format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/Style/full
	DateIntervalFormatterFullStyle DateIntervalFormatterStyle = 4
	// DateIntervalFormatterLongStyle - A long length date or time format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/Style/long
	DateIntervalFormatterLongStyle DateIntervalFormatterStyle = 3
	// DateIntervalFormatterMediumStyle - A medium length date or time format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/Style/medium
	DateIntervalFormatterMediumStyle DateIntervalFormatterStyle = 2
	// DateIntervalFormatterNoStyle - No information for the date or time. Use this style when you do not want to include date or time information in the resulting string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/Style/none
	DateIntervalFormatterNoStyle DateIntervalFormatterStyle = 0
	// DateIntervalFormatterShortStyle - An abbreviated date or time format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/Style/short
	DateIntervalFormatterShortStyle DateIntervalFormatterStyle = 1
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

// VolumeEnumerationOptions - Options for enumerating mounted volumes with the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions
type VolumeEnumerationOptions uint

const (
	VolumeEnumerationSkipHiddenVolumes VolumeEnumerationOptions = 2
	VolumeEnumerationProduceFileReferenceURLs VolumeEnumerationOptions = 4
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
	// DocumentationDirectory - Documentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentationDirectory
	DocumentationDirectory SearchPathDirectory = 8
	// DocumentDirectory - Document directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory/documentDirectory
	DocumentDirectory SearchPathDirectory = 9
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

// InlinePresentationIntent - A type that defines presentation intent for runs of characters for traits like emphasis, strikethrough, and code voice.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent
type InlinePresentationIntent uint

const (
	// InlinePresentationIntentBlockHTML - An intent that represents a block HTML presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/blockHTML
	InlinePresentationIntentBlockHTML InlinePresentationIntent = 512
	// InlinePresentationIntentCode - An intent that represents a code voice presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/code
	InlinePresentationIntentCode InlinePresentationIntent = 4
	// InlinePresentationIntentEmphasized - An intent that represents an emphasized presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/emphasized
	InlinePresentationIntentEmphasized InlinePresentationIntent = 1
	// InlinePresentationIntentInlineHTML - An intent that represents an inline HTML presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/inlineHTML
	InlinePresentationIntentInlineHTML InlinePresentationIntent = 256
	// InlinePresentationIntentLineBreak - An intent that represents a line break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/lineBreak
	InlinePresentationIntentLineBreak InlinePresentationIntent = 128
	// InlinePresentationIntentSoftBreak - An intent that represents a soft line break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/softBreak
	InlinePresentationIntentSoftBreak InlinePresentationIntent = 64
	// InlinePresentationIntentStrikethrough - An intent that represents a strikethrough presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/strikethrough
	InlinePresentationIntentStrikethrough InlinePresentationIntent = 32
	// InlinePresentationIntentStronglyEmphasized - An intent that represents a strongly emphasized presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent/stronglyEmphasized
	InlinePresentationIntentStronglyEmphasized InlinePresentationIntent = 2
)

// ISO8601DateFormatOptions - Options used to generate and parse ISO 8601 date representations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options
type ISO8601DateFormatOptions uint

const (
	// ISO8601DateFormatWithColonSeparatorInTime - The date representation uses the colon separator ( ) in the time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withColonSeparatorInTime
	ISO8601DateFormatWithColonSeparatorInTime ISO8601DateFormatOptions = 8
	// ISO8601DateFormatWithColonSeparatorInTimeZone - The date representation uses the colon separator ( ) in the time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withColonSeparatorInTimeZone
	ISO8601DateFormatWithColonSeparatorInTimeZone ISO8601DateFormatOptions = 9
	// ISO8601DateFormatWithDashSeparatorInDate - The date representation uses the dash separator ( ) in the date.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withDashSeparatorInDate
	ISO8601DateFormatWithDashSeparatorInDate ISO8601DateFormatOptions = 7
	// ISO8601DateFormatWithDay - The date representation includes the day. The format for day is inferred based on provided options: If   is specified,   is used. If   is specified,   is used. Otherwise,   is used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withDay
	ISO8601DateFormatWithDay ISO8601DateFormatOptions = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withFractionalSeconds
	ISO8601DateFormatWithFractionalSeconds ISO8601DateFormatOptions = 10
	// ISO8601DateFormatWithFullDate - The date representation includes the year, month, and day. Equivalent to specifying  ,  , and 
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withFullDate
	ISO8601DateFormatWithFullDate ISO8601DateFormatOptions = 11
	// ISO8601DateFormatWithFullTime - The date representation includes the hour, minute, and second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withFullTime
	ISO8601DateFormatWithFullTime ISO8601DateFormatOptions = 12
	// ISO8601DateFormatWithInternetDateTime - The format used for internet date times, according to the   standard. Equivalent to specifying  ,  ,  ,  , and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withInternetDateTime
	ISO8601DateFormatWithInternetDateTime ISO8601DateFormatOptions = 13
	// ISO8601DateFormatWithMonth - The date representation includes the month. The format for month is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withMonth
	ISO8601DateFormatWithMonth ISO8601DateFormatOptions = 1
	// ISO8601DateFormatWithSpaceBetweenDateAndTime - The date representation uses a space ( ) instead of   between the date and time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withSpaceBetweenDateAndTime
	ISO8601DateFormatWithSpaceBetweenDateAndTime ISO8601DateFormatOptions = 6
	// ISO8601DateFormatWithTime - The date representation includes the time. The format for time is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withTime
	ISO8601DateFormatWithTime ISO8601DateFormatOptions = 4
	// ISO8601DateFormatWithTimeZone - The date representation includes the timezone. The format for timezone is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withTimeZone
	ISO8601DateFormatWithTimeZone ISO8601DateFormatOptions = 5
	// ISO8601DateFormatWithWeekOfYear - The date representation includes the week of the year. The format for week of year is  , including the   prefix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withWeekOfYear
	ISO8601DateFormatWithWeekOfYear ISO8601DateFormatOptions = 2
	// ISO8601DateFormatWithYear - The date representation includes the year. The format for year is inferred based on the other specified options. If   is specified,   is used. Otherwise,   is used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options/withYear
	ISO8601DateFormatWithYear ISO8601DateFormatOptions = 0
)

// CollectionChangeType - The type of change represented in computing the difference of an ordered collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCollectionChangeType
type CollectionChangeType uint

const (
	// CollectionChangeInsert - A change type that represents the insertion of an object into an ordered collection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCollectionChangeType/insert
	CollectionChangeInsert CollectionChangeType = 0
	// CollectionChangeRemove - A change type that represents the removal of an object from an ordered collection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCollectionChangeType/remove
	CollectionChangeRemove CollectionChangeType = 1
)

// DataBase64DecodingOptions - Options to modify the decoding algorithm used to decode Base64 encoded data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions
type DataBase64DecodingOptions uint

const (
	// DataBase64DecodingIgnoreUnknownCharacters - Modify the decoding algorithm so that it ignores unknown non-Base-64 bytes, including line ending characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions/ignoreUnknownCharacters
	DataBase64DecodingIgnoreUnknownCharacters DataBase64DecodingOptions = 1
)

// DataBase64EncodingOptions - Options for methods used to Base64 encode data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions
type DataBase64EncodingOptions uint

const (
	// DataBase64EncodingEndLineWithCarriageReturn - When a maximum line length is set, specify that the line ending to insert should include a carriage return.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions/endLineWithCarriageReturn
	DataBase64EncodingEndLineWithCarriageReturn DataBase64EncodingOptions = 16
	// DataBase64EncodingEndLineWithLineFeed - When a maximum line length is set, specify that the line ending to insert should include a line feed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions/endLineWithLineFeed
	DataBase64EncodingEndLineWithLineFeed DataBase64EncodingOptions = 32
	// DataBase64Encoding64CharacterLineLength - Set the maximum line length to 64 characters, after which a line ending is inserted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions/lineLength64Characters
	DataBase64Encoding64CharacterLineLength DataBase64EncodingOptions = 1
	// DataBase64Encoding76CharacterLineLength - Set the maximum line length to 76 characters, after which a line ending is inserted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions/lineLength76Characters
	DataBase64Encoding76CharacterLineLength DataBase64EncodingOptions = 2
)

// DataCompressionAlgorithm - An algorithm that indicates how to compress or decompress data.
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

// DataReadingOptions - Options for methods used to read data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions
type DataReadingOptions uint

const (
	// DataReadingMappedAlways - Hint to map the file in if possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/alwaysMapped
	DataReadingMappedAlways DataReadingOptions = 3
	// DataReadingMapped - Deprecated name for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/dataReadingMapped
	DataReadingMapped DataReadingOptions = 4
	// DataReadingMappedIfSafe - A hint indicating the file should be mapped into virtual memory, if possible and safe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/mappedIfSafe
	DataReadingMappedIfSafe DataReadingOptions = 1
	// MappedRead - Deprecated name for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/mappedRead
	MappedRead DataReadingOptions = 5
	// DataReadingUncached - A hint indicating the file should not be stored in the file-system caches.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/uncached
	DataReadingUncached DataReadingOptions = 2
	// UncachedRead - Deprecated name for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions/uncachedRead
	UncachedRead DataReadingOptions = 6
)

// DataSearchOptions - Options for method used to search data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions
type DataSearchOptions uint

const (
	// DataSearchAnchored - Search is limited to start (or end, if searching backwards) of the data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions/anchored
	DataSearchAnchored DataSearchOptions = 2
	// DataSearchBackwards - Search from the end of the data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions/backwards
	DataSearchBackwards DataSearchOptions = 1
)

// DataWritingOptions - Options for methods used to write data objects.
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

// CalculationError - Calculation error constants used to describe an error in 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError
type CalculationError uint

const (
	// CalculationDivideByZero - The caller tried to divide by  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError/divideByZero
	CalculationDivideByZero CalculationError = 4
	// CalculationLossOfPrecision - The number can’t be represented in 38 significant digits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError/lossOfPrecision
	CalculationLossOfPrecision CalculationError = 1
	// CalculationNoError - No error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError/noError
	CalculationNoError CalculationError = 0
	// CalculationOverflow - The number is too large to represent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError/overflow
	CalculationOverflow CalculationError = 3
	// CalculationUnderflow - The number is too small to represent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError/underflow
	CalculationUnderflow CalculationError = 2
)

// RoundingMode - These constants specify rounding behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/RoundingMode
type RoundingMode uint

const (
	// RoundBankers - Round to the closest possible return value; when halfway between two possibilities, return the possibility whose last digit is even.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/RoundingMode/bankers
	RoundBankers RoundingMode = 3
	// RoundDown - Round return values down.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/RoundingMode/down
	RoundDown RoundingMode = 1
	// RoundPlain - Round to the closest possible return value; when caught halfway between two positive numbers, round up; when caught between two negative numbers, round down.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/RoundingMode/plain
	RoundPlain RoundingMode = 0
	// RoundUp - Round return values up.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/RoundingMode/up
	RoundUp RoundingMode = 2
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

// LinguisticTaggerOptions - Constants for linguistic tagger enumeration specifying which tokens to omit and whether to join names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options
type LinguisticTaggerOptions uint

const (
	// LinguisticTaggerJoinNames - Typically, multiple-word names will be returned as multiple tokens, following the standard tokenization practice of the tagger.  If this option is set, then multiple-word names will be joined together and returned as a single token.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options/joinNames
	LinguisticTaggerJoinNames LinguisticTaggerOptions = 16
	// LinguisticTaggerOmitOther - Omit tokens of type   (non-linguistic items, such as symbols).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options/omitOther
	LinguisticTaggerOmitOther LinguisticTaggerOptions = 8
	// LinguisticTaggerOmitPunctuation - Omit tokens of type   (all punctuation).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options/omitPunctuation
	LinguisticTaggerOmitPunctuation LinguisticTaggerOptions = 2
	// LinguisticTaggerOmitWhitespace - Omit tokens of type   (whitespace of all sorts).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options/omitWhitespace
	LinguisticTaggerOmitWhitespace LinguisticTaggerOptions = 4
	// LinguisticTaggerOmitWords - Omit tokens of type   (items considered to be words).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options/omitWords
	LinguisticTaggerOmitWords LinguisticTaggerOptions = 1
)

// LinguisticTaggerUnit - Constants representing linguistic units.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTaggerUnit
type LinguisticTaggerUnit uint

const (
	// LinguisticTaggerUnitDocument - The document in its entirety.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTaggerUnit/document
	LinguisticTaggerUnitDocument LinguisticTaggerUnit = 3
	// LinguisticTaggerUnitParagraph - An individual paragraph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTaggerUnit/paragraph
	LinguisticTaggerUnitParagraph LinguisticTaggerUnit = 2
	// LinguisticTaggerUnitSentence - An individual sentence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTaggerUnit/sentence
	LinguisticTaggerUnitSentence LinguisticTaggerUnit = 1
	// LinguisticTaggerUnitWord - An individual word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTaggerUnit/word
	LinguisticTaggerUnitWord LinguisticTaggerUnit = 0
)

// LocaleLanguageDirection - The directions that a language may take across a page of text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection
type LocaleLanguageDirection uint

const (
	// LocaleLanguageDirectionBottomToTop - The language direction is from bottom to top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection/bottomToTop
	LocaleLanguageDirectionBottomToTop LocaleLanguageDirection = 0
	// LocaleLanguageDirectionLeftToRight - The language direction is from left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection/leftToRight
	LocaleLanguageDirectionLeftToRight LocaleLanguageDirection = 0
	// LocaleLanguageDirectionRightToLeft - The language direction is from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection/rightToLeft
	LocaleLanguageDirectionRightToLeft LocaleLanguageDirection = 0
	// LocaleLanguageDirectionTopToBottom - The language direction is from top to bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection/topToBottom
	LocaleLanguageDirectionTopToBottom LocaleLanguageDirection = 0
	// LocaleLanguageDirectionUnknown - The direction of the language is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection/unknown
	LocaleLanguageDirectionUnknown LocaleLanguageDirection = 0
)

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
	// PointerFunctionsZeroingWeakMemory - Use weak read and write barriers; use garbage-collected memory on copyIn.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctionsOptions/NSPointerFunctionsZeroingWeakMemory
	PointerFunctionsZeroingWeakMemory PointerFunctionsOptions = 1
	// PointerFunctionsCopyIn - Use the memory acquire function to allocate and copy items on input (see  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/copyIn
	PointerFunctionsCopyIn PointerFunctionsOptions = 12
	// PointerFunctionsCStringPersonality - Use a string hash and  ; C-string ‘ ’ style description.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options/cStringPersonality
	PointerFunctionsCStringPersonality PointerFunctionsOptions = 9
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
)

// RectEdge enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge
type RectEdge uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMaxXEdge
	MaxXEdge RectEdge = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMaxYEdge
	MaxYEdge RectEdge = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMinXEdge
	MinXEdge RectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMinYEdge
	MinYEdge RectEdge = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge/maxX
	RectEdgeMaxX RectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge/maxY
	RectEdgeMaxY RectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge/minX
	RectEdgeMinX RectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectEdge/minY
	RectEdgeMinY RectEdge = 0
)

// MatchingFlags - Set by the Block as the matching progresses, completes, or fails. Used by the method 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags
type MatchingFlags uint

const (
	// MatchingCompleted - Set when the Block is called after matching has completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags/completed
	MatchingCompleted MatchingFlags = 2
	// MatchingHitEnd - Set when the current match operation reached the end of the search range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags/hitEnd
	MatchingHitEnd MatchingFlags = 4
	// MatchingInternalError - Set when matching failed due to an internal error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags/internalError
	MatchingInternalError MatchingFlags = 16
	// MatchingProgress - Set when the Block is called to report progress during a long-running match operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags/progress
	MatchingProgress MatchingFlags = 1
	// MatchingRequiredEnd - Set when the current match depended on the location of the end of the search range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags/requiredEnd
	MatchingRequiredEnd MatchingFlags = 8
)

// MatchingOptions - The matching options constants specify the reporting, completion and matching rules to the expression matching methods. These constants are used by all methods that search for, or replace values, using a regular expression.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
type MatchingOptions uint

const (
	// MatchingAnchored - Specifies that matches are limited to those at the start of the search range. See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/anchored
	MatchingAnchored MatchingOptions = 4
	// MatchingReportCompletion - Call the Block once after the completion of any matching. This option has no effect for methods other than  . See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/reportCompletion
	MatchingReportCompletion MatchingOptions = 2
	// MatchingReportProgress - Call the Block periodically during long-running match operations. This option has no effect for methods other than  . See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/reportProgress
	MatchingReportProgress MatchingOptions = 1
	// MatchingWithoutAnchoringBounds - Specifies that   and   will not automatically match the beginning and end of the search range, but will still match the beginning and end of the entire string. This constant has no effect if the search range contains the entire string. See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/withoutAnchoringBounds
	MatchingWithoutAnchoringBounds MatchingOptions = 16
	// MatchingWithTransparentBounds - Specifies that matching may examine parts of the string beyond the bounds of the search range, for purposes such as word boundary detection, lookahead, etc. This constant has no effect if the search range contains the entire string. See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/withTransparentBounds
	MatchingWithTransparentBounds MatchingOptions = 8
)

// RegularExpressionOptions - These constants define the regular expression options. These constants are used by the property 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
type RegularExpressionOptions uint

const (
	// RegularExpressionAllowCommentsAndWhitespace - Ignore whitespace and #-prefixed comments in the pattern.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/allowCommentsAndWhitespace
	RegularExpressionAllowCommentsAndWhitespace RegularExpressionOptions = 2
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
	// RegularExpressionIgnoreMetacharacters - Treat the entire pattern as a literal string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/ignoreMetacharacters
	RegularExpressionIgnoreMetacharacters RegularExpressionOptions = 4
	// RegularExpressionUseUnicodeWordBoundaries - Use Unicode   to specify word boundaries (otherwise, traditional regular expression word boundaries are used).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/useUnicodeWordBoundaries
	RegularExpressionUseUnicodeWordBoundaries RegularExpressionOptions = 64
	// RegularExpressionUseUnixLineSeparators - Treat only   as a line separator (otherwise, all standard line separators are used).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/useUnixLineSeparators
	RegularExpressionUseUnixLineSeparators RegularExpressionOptions = 32
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

// TextCheckingType - These constants specify the type of checking the methods should do. They are returned by 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
type TextCheckingType uint

const (
	// TextCheckingTypeAddress - Attempts to locate addresses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/address
	TextCheckingTypeAddress TextCheckingType = 16
	// TextCheckingTypeCorrection - Performs autocorrection on misspelled words.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/correction
	TextCheckingTypeCorrection TextCheckingType = 512
	// TextCheckingTypeDash - Replaces dashes with em-dashes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/dash
	TextCheckingTypeDash TextCheckingType = 128
	// TextCheckingTypeDate - Attempts to locate dates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/date
	TextCheckingTypeDate TextCheckingType = 8
	// TextCheckingTypeGrammar - Checks grammar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/grammar
	TextCheckingTypeGrammar TextCheckingType = 4
	// TextCheckingTypeLink - Attempts to locate URL links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/link
	TextCheckingTypeLink TextCheckingType = 32
	// TextCheckingTypeOrthography - Attempts to identify the language
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/orthography
	TextCheckingTypeOrthography TextCheckingType = 1
	// TextCheckingTypePhoneNumber - Matches a phone number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/phoneNumber
	TextCheckingTypePhoneNumber TextCheckingType = 514
	// TextCheckingTypeQuote - Replaces quotes with smart quotes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/quote
	TextCheckingTypeQuote TextCheckingType = 64
	// TextCheckingTypeRegularExpression - Matches a regular expression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/regularExpression
	TextCheckingTypeRegularExpression TextCheckingType = 513
	// TextCheckingTypeReplacement - Replaces characters such as (c) with the appropriate symbol (in this case ©).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/replacement
	TextCheckingTypeReplacement TextCheckingType = 256
	// TextCheckingTypeSpelling - Checks spelling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/spelling
	TextCheckingTypeSpelling TextCheckingType = 2
	// TextCheckingTypeTransitInformation - Matches a transit information, for example, flight information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/transitInformation
	TextCheckingTypeTransitInformation TextCheckingType = 515
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
	// URLBookmarkCreationWithoutImplicitSecurityScope - Prevents inclusion of a bookmark’s implicit ephemeral security scope, when creating one without security scope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/withoutImplicitSecurityScope
	URLBookmarkCreationWithoutImplicitSecurityScope URLBookmarkCreationOptions = 1027
	// URLBookmarkCreationWithSecurityScope - Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read/write access to a file-system resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions/withSecurityScope
	URLBookmarkCreationWithSecurityScope URLBookmarkCreationOptions = 1025
)

// URLBookmarkResolutionOptions - Options used when resolving bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions
type URLBookmarkResolutionOptions uint

const (
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
	// URLBookmarkResolutionWithSecurityScope - Specifies that the security scope, applied to the bookmark when it was created, should be used during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions/withSecurityScope
	URLBookmarkResolutionWithSecurityScope URLBookmarkResolutionOptions = 513
)

// NumberFormatterBehavior - These constants specify the behavior of a number formatter. These constants are returned by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Behavior
type NumberFormatterBehavior uint

const (
	// NumberFormatterBehavior10_0 - The number-formatter behavior as it existed prior to macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Behavior/behavior10_0
	NumberFormatterBehavior10_0 NumberFormatterBehavior = 1000
	// NumberFormatterBehavior10_4 - The number-formatter behavior since macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Behavior/behavior10_4
	NumberFormatterBehavior10_4 NumberFormatterBehavior = 1040
	// NumberFormatterBehaviorDefault - The number-formatter behavior set as the default for new instances. You can set the default formatter behavior with the class method  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Behavior/default
	NumberFormatterBehaviorDefault NumberFormatterBehavior = 0
)

// NumberFormatterPadPosition - These constants are used to specify how numbers should be padded. These constants are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/PadPosition
type NumberFormatterPadPosition uint

const (
	// NumberFormatterPadAfterPrefix - Specifies that the padding should occur after the prefix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/PadPosition/afterPrefix
	NumberFormatterPadAfterPrefix NumberFormatterPadPosition = 0
	// NumberFormatterPadAfterSuffix - Specifies that the padding should occur after the suffix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/PadPosition/afterSuffix
	NumberFormatterPadAfterSuffix NumberFormatterPadPosition = 0
	// NumberFormatterPadBeforePrefix - Specifies that the padding should occur before the prefix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/PadPosition/beforePrefix
	NumberFormatterPadBeforePrefix NumberFormatterPadPosition = 0
	// NumberFormatterPadBeforeSuffix - Specifies that the padding should occur before the suffix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/PadPosition/beforeSuffix
	NumberFormatterPadBeforeSuffix NumberFormatterPadPosition = 0
)

// NumberFormatterRoundingMode - These constants are used to specify how numbers should be rounded. These constants are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum
type NumberFormatterRoundingMode uint

const (
	// NumberFormatterRoundCeiling - Round towards positive infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum/ceiling
	NumberFormatterRoundCeiling NumberFormatterRoundingMode = 0
	// NumberFormatterRoundDown - Round towards zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum/down
	NumberFormatterRoundDown NumberFormatterRoundingMode = 0
	// NumberFormatterRoundFloor - Round towards negative infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum/floor
	NumberFormatterRoundFloor NumberFormatterRoundingMode = 0
	// NumberFormatterRoundHalfDown - Round towards the nearest integer, or towards zero if equidistant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum/halfDown
	NumberFormatterRoundHalfDown NumberFormatterRoundingMode = 0
	// NumberFormatterRoundHalfEven - Round towards the nearest integer, or towards an even number if equidistant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum/halfEven
	NumberFormatterRoundHalfEven NumberFormatterRoundingMode = 0
	// NumberFormatterRoundHalfUp - Round towards the nearest integer, or away from zero if equidistant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum/halfUp
	NumberFormatterRoundHalfUp NumberFormatterRoundingMode = 0
	// NumberFormatterRoundUp - Round away from zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum/up
	NumberFormatterRoundUp NumberFormatterRoundingMode = 0
)

// NumberFormatterStyle - The predefined number format styles used by the 
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

// AXFeatureOverrideSessionOptions - Options indicating which Accessibility features will be turned on or off when an override session is held by your app.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options
type AXFeatureOverrideSessionOptions uint

// AXFeatureOverrideSessionError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionError-swift.struct/Code
type AXFeatureOverrideSessionError uint

// APActivationPayloadErrorCode - Error codes that an App Clip activation payload returns.
//
// [Full Topic]: https://developer.apple.com/documentation/AppClip/APActivationPayloadError/Code
type APActivationPayloadErrorCode uint

// AccessibilityAnnotationPosition - Constants that specify the position where the annotation applies.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition
type AccessibilityAnnotationPosition uint

// BezierPathElement - Constants that specify basic path element commands.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType
type BezierPathElement uint

const (
	// BezierPathElementClosePath - Marks the end of the current subpath at the specified point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/closePath
	BezierPathElementClosePath BezierPathElement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/cubicCurveTo
	BezierPathElementCubicCurveTo BezierPathElement = 0
	// BezierPathElementCurveTo - Creates a curved line segment from the current point to the specified endpoint using two control points to define the curve.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/curveTo
	BezierPathElementCurveTo BezierPathElement = 0
	// BezierPathElementLineTo - Creates a straight line from the current drawing point to the specified point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/lineTo
	BezierPathElementLineTo BezierPathElement = 0
	// BezierPathElementMoveTo - Moves the path object’s current drawing point to the specified point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/moveTo
	BezierPathElementMoveTo BezierPathElement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/quadraticCurveTo
	BezierPathElementQuadraticCurveTo BezierPathElement = 0
)

// LineCapStyle - Constants that specify the shape of endpoints for an open path when it is stroked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum
type LineCapStyle uint

const (
	// LineCapStyleButt - Specifies a butt line cap style for endpoints for an open path when stroked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum/butt
	LineCapStyleButt LineCapStyle = 0
	// LineCapStyleRound - Specifies a round line cap style for endpoints for an open path when stroked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum/round
	LineCapStyleRound LineCapStyle = 0
	// LineCapStyleSquare - Specifies a square line cap style for endpoints for an open path when stroked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum/square
	LineCapStyleSquare LineCapStyle = 0
)

// LineJoinStyle - Constants that specify the shape of the joins between connected segments of a stroked path.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum
type LineJoinStyle uint

const (
	// LineJoinStyleBevel - Specifies a bevel line shape of the joints between connected segments of a stroked path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum/bevel
	LineJoinStyleBevel LineJoinStyle = 0
	// LineJoinStyleMiter - Specifies a miter line shape of the joints between connected segments of a stroked path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum/miter
	LineJoinStyleMiter LineJoinStyle = 0
	// LineJoinStyleRound - Specifies a round line shape of the joints between connected segments of a stroked path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum/round
	LineJoinStyleRound LineJoinStyle = 0
)

// WindingRule - Constants that specify the winding rule a Bézier path uses.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/WindingRule-swift.enum
type WindingRule uint

const (
	// WindingRuleEvenOdd - Specifies the even-odd winding rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/WindingRule-swift.enum/evenOdd
	WindingRuleEvenOdd WindingRule = 0
	// WindingRuleNonZero - Specifies the non-zero winding rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/WindingRule-swift.enum/nonZero
	WindingRuleNonZero WindingRule = 0
)

// BorderType - These constants specify the type of a view’s border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType
type BorderType uint

const (
	// BezelBorder - A concave border that makes the view look sunken.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/bezelBorder
	BezelBorder BorderType = 0
	// GrooveBorder - A thin border that looks etched around the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/grooveBorder
	GrooveBorder BorderType = 0
	// LineBorder - A black line border around the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/lineBorder
	LineBorder BorderType = 0
	// NoBorder - No border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/noBorder
	NoBorder BorderType = 0
)

// CharacterCollection - Values that map character identifiers to glyphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection
type CharacterCollection uint

// ColorType - Constants that indicate the color’s type, and which methods may be called on the color object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType
type ColorType uint

// ColorSystemEffect - Constants for user interactions that change the appearance of a view or control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect
type ColorSystemEffect uint

// SaveOperationType - Constants for specifying the type of document-save operation to perform.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType
type SaveOperationType uint

const (
	// SaveAsOperation - An operation that writes the document’s contents to a new location and updates the document to point to that location
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveAsOperation
	SaveAsOperation SaveOperationType = 0
	// SaveOperation - An operation that overwrites a document’s file or file package with the document’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveOperation
	SaveOperation SaveOperationType = 0
)

// EventMask - Constants that you use to filter out specific event types from the stream of incoming events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask
type EventMask uint

// FocusRingPlacement - Constants that indicate how the system draws the focus ring.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement
type FocusRingPlacement uint

// FocusRingType - Constants that describe the style of the focus ring.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType
type FocusRingType uint

// FontCollectionOptions - Constants that support font collection management.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollectionOptions
type FontCollectionOptions uint

// FontRenderingMode - The font rendering mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode
type FontRenderingMode uint

// ImageLoadStatus - Status values for incremental image loading.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus
type ImageLoadStatus uint

// GlyphProperty - Glyph properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty
type GlyphProperty uint

// TextLayoutOrientation - Constants that describe the text layout orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation
type TextLayoutOrientation uint

const (
	// TextLayoutOrientationHorizontal - Lines render horizontally, each line following the previous from top to bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation/horizontal
	TextLayoutOrientationHorizontal TextLayoutOrientation = 0
	// TextLayoutOrientationVertical - Lines render vertically, each line following the previous from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation/vertical
	TextLayoutOrientationVertical TextLayoutOrientation = 0
)

// TypesetterBehavior - Constants that determine the layout manager’s behavior during layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum
type TypesetterBehavior uint

// LineBreakMode - Constants that specify what happens when a line is too long for a container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode
type LineBreakMode uint

// LineMovementDirection - The direction in which a line moves.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection
type LineMovementDirection uint

// LineSweepDirection - Values that describe the progression of text on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection
type LineSweepDirection uint

// LineBreakStrategy - Constants that specify how the text system breaks lines while laying out paragraphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct
type LineBreakStrategy uint

// TextTabType - Constants that specify the type of tab stop.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType
type TextTabType uint

// PasteboardAccessBehavior - A value indicating pasteboard access behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum
type PasteboardAccessBehavior uint

const (
	// PasteboardAccessBehaviorAlwaysAllow - The system will automatically allow all pasteboard access, without notifying the user.  The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/alwaysAllow
	PasteboardAccessBehaviorAlwaysAllow PasteboardAccessBehavior = 0
	// PasteboardAccessBehaviorAlwaysDeny - The system will automatically deny all pasteboard access, without notifying the user. However, access that is both user originated and paste related will always be allowed, and will not result in a notification. The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/alwaysDeny
	PasteboardAccessBehaviorAlwaysDeny PasteboardAccessBehavior = 0
	// PasteboardAccessBehaviorAsk - The system will notify the user and ask for permission before granting pasteboard access. However, access that is both user originated and paste related will always be allowed, and will not result in a notification. The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/ask
	PasteboardAccessBehaviorAsk PasteboardAccessBehavior = 0
	// PasteboardAccessBehaviorDefault - The default behavior for the General pasteboard is to ask upon programmatic access. All other pasteboards default to always allow access.   If an app has never triggered a pasteboard access alert, its General pasteboard will report   behavior. Such an app is not shown in the corresponding System Settings pane.   Once programmatic pasteboard access triggers the first pasteboard access alert, the state automatically changes to  . At this point the app starts being shown in System Settings, where the user can toggle the behavior between  ,  , and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/default
	PasteboardAccessBehaviorDefault PasteboardAccessBehavior = 0
)

// PasteboardContentsOptions - Options for preparing the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ContentsOptions
type PasteboardContentsOptions uint

// PasteboardReadingOptions - Options that specify how to interpret data on the pasteboard when initializing pasteboard data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
type PasteboardReadingOptions uint

const (
	// PasteboardReadingAsData - An option to read data from the pasteboard as-is and return it as a data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asData
	PasteboardReadingAsData PasteboardReadingOptions = 0
	// PasteboardReadingAsKeyedArchive - An option to read data from the pasteboard and use it to initialize the object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asKeyedArchive
	PasteboardReadingAsKeyedArchive PasteboardReadingOptions = 0
	// PasteboardReadingAsPropertyList - An option to read data from the pasteboard and unserialize it as a property list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asPropertyList
	PasteboardReadingAsPropertyList PasteboardReadingOptions = 0
	// PasteboardReadingAsString - An option to read data from the pasteboard and convert it to a string object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asString
	PasteboardReadingAsString PasteboardReadingOptions = 0
)

// PasteboardWritingOptions - Type to specify options for writing to a pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions
type PasteboardWritingOptions uint

const (
	// PasteboardWritingPromised - Data for a type with this option is promised, not immediately written.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions/promised
	PasteboardWritingPromised PasteboardWritingOptions = 0
)

// CorrectionIndicatorType - Constants that allow an app to specify the correction indicator type displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionIndicatorType
type CorrectionIndicatorType uint

// CorrectionResponse - The correction response passed to the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse
type CorrectionResponse uint

// TextAlignment - Constants that specify text alignment.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment
type TextAlignment uint

// TextInsertionIndicatorAutomaticModeOptions - Options that affect the automatic display mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/AutomaticModeOptions-swift.struct
type TextInsertionIndicatorAutomaticModeOptions uint

// TextInsertionIndicatorDisplayMode - Constants that determine how to display the system text cursor in a custom text UI.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum
type TextInsertionIndicatorDisplayMode uint

// TitlebarSeparatorStyle - Styles that determine the type of separator displayed between the title bar and content of a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle
type TitlebarSeparatorStyle uint

// WindowAnimationBehavior - Constants that control the automatic window animation behavior windows use when ordering to the front or out of view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum
type WindowAnimationBehavior uint

// BackingStoreType - Constants that specify how the window device buffers the drawing done in a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
type BackingStoreType uint

// WindowButton - Constants that provide a way to access standard title bar buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType
type WindowButton uint

// WindowCollectionBehavior - Window collection behaviors related to Mission Control, Spaces, and Stage Manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct
type WindowCollectionBehavior uint

// WindowDepth - A type that represents the depth, or amount of memory, for a single pixel in a window or screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth
type WindowDepth uint

// WindowNumberListOptions - Options to use when retrieving window numbers from the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions
type WindowNumberListOptions uint

// WindowOcclusionState - Specifies whether the window is occluded.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct
type WindowOcclusionState uint

// WindowOrderingMode - Constants that let you specify how a window is ordered relative to another window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode
type WindowOrderingMode uint

// SelectionDirection - Constants that specify the direction a window is currently using to change the key view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection
type SelectionDirection uint

// WindowSharingType - Constants that represent the access levels other processes can have to a window’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum
type WindowSharingType uint

// WindowStyleMask - Constants that specify the style of a window, and that you can combine with the C bitwise OR operator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
type WindowStyleMask uint

// WindowTabbingMode - The preferred tabbing behavior of a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum
type WindowTabbingMode uint

// WindowTitleVisibility - Specifies the appearance of the window’s title bar area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum
type WindowTitleVisibility uint

// WindowToolbarStyle - Styles that determine the appearance and location of the toolbar in relation to the title bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum
type WindowToolbarStyle uint

// WindowUserTabbingPreference - A value that indicates the user’s preference for window tabbing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum
type WindowUserTabbingPreference uint

// WorkspaceAuthorizationType - The types of privileged file operations that can be authorized by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType
type WorkspaceAuthorizationType uint

const (
	// WorkspaceAuthorizationTypeCreateSymbolicLink - Authorization for the app to create a symbolic link.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType/createSymbolicLink
	WorkspaceAuthorizationTypeCreateSymbolicLink WorkspaceAuthorizationType = 0
	// WorkspaceAuthorizationTypeReplaceFile - Authorization for the app to perform an atomic file write without changing the target file’s permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType/replaceFile
	WorkspaceAuthorizationTypeReplaceFile WorkspaceAuthorizationType = 0
	// WorkspaceAuthorizationTypeSetAttributes - Authorization for the app to change specific file attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType/setAttributes
	WorkspaceAuthorizationTypeSetAttributes WorkspaceAuthorizationType = 0
)

// WorkspaceIconCreationOptions - Constants that describe options for creating icons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions
type WorkspaceIconCreationOptions uint

const (
	// Exclude10_4ElementsIconCreationOption - An option to suppress generation of the new higher resolution icon representations that are supported in macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions/exclude10_4ElementsIconCreationOption
	Exclude10_4ElementsIconCreationOption WorkspaceIconCreationOptions = 0
	// ExcludeQuickDrawElementsIconCreationOption - An option to suppress generation of the QuickDraw format icon representations that are used in macOS 10.0 through macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions/excludeQuickDrawElementsIconCreationOption
	ExcludeQuickDrawElementsIconCreationOption WorkspaceIconCreationOptions = 0
)

// WorkspaceLaunchOptions - Constants specifying how you want to launch an app
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions
type WorkspaceLaunchOptions uint

// WritingDirection - Constants that specify the writing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection
type WritingDirection uint

const (
	// WritingDirectionLeftToRight - The writing direction is left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/leftToRight
	WritingDirectionLeftToRight WritingDirection = 0
	// WritingDirectionNatural - The writing direction of the current script that the system determines using the Unicode Bidi Algorithm rules P2 and P3.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/natural
	WritingDirectionNatural WritingDirection = 0
	// WritingDirectionRightToLeft - The writing direction is right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/rightToLeft
	WritingDirectionRightToLeft WritingDirection = 0
)

// FetchRequestResultType - Constants that specify the possible result types a fetch request can return.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType
type FetchRequestResultType uint

// FetchedResultsChangeType - Constants that specify the possible types of changes that are reported.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType
type FetchedResultsChangeType uint

// CFCalendarUnit - CFCalendarUnit constants are used to specify calendrical units, such as day or month, in various calendar calculations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit
type CFCalendarUnit uint

const (
	// kCFCalendarUnitDay - Specifies the day unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/day
	kCFCalendarUnitDay CFCalendarUnit = 0
	// kCFCalendarUnitEra - Specifies the era unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/era
	kCFCalendarUnitEra CFCalendarUnit = 0
	// kCFCalendarUnitHour - Specifies the hour unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/hour
	kCFCalendarUnitHour CFCalendarUnit = 0
	// kCFCalendarUnitMinute - Specifies the minute unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/minute
	kCFCalendarUnitMinute CFCalendarUnit = 0
	// kCFCalendarUnitMonth - Specifies the month unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/month
	kCFCalendarUnitMonth CFCalendarUnit = 0
	// kCFCalendarUnitQuarter - Specifies the quarter-year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/quarter
	kCFCalendarUnitQuarter CFCalendarUnit = 0
	// kCFCalendarUnitSecond - Specifies the second unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/second
	kCFCalendarUnitSecond CFCalendarUnit = 0
	// kCFCalendarUnitWeek - Specifies the week unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/week
	kCFCalendarUnitWeek CFCalendarUnit = 0
	// kCFCalendarUnitWeekOfMonth - Specifies the original week of a month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/weekOfMonth
	kCFCalendarUnitWeekOfMonth CFCalendarUnit = 0
	// kCFCalendarUnitWeekOfYear - Specifies the original week of the year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/weekOfYear
	kCFCalendarUnitWeekOfYear CFCalendarUnit = 0
	// kCFCalendarUnitWeekday - Specifies the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/weekday
	kCFCalendarUnitWeekday CFCalendarUnit = 0
	// kCFCalendarUnitWeekdayOrdinal - Specifies the ordinal weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/weekdayOrdinal
	kCFCalendarUnitWeekdayOrdinal CFCalendarUnit = 0
	// kCFCalendarUnitYear - Specifies the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/year
	kCFCalendarUnitYear CFCalendarUnit = 0
	// kCFCalendarUnitYearForWeekOfYear - Specifies the relative year for a week within a year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/yearForWeekOfYear
	kCFCalendarUnitYearForWeekOfYear CFCalendarUnit = 0
)

// CFCharacterSetPredefinedSet - Defines a predefined character set.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet
type CFCharacterSetPredefinedSet uint

// CFDataSearchFlags - A 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags
type CFDataSearchFlags uint

const (
	// kCFDataSearchAnchored - Performs searching only on bytes at the beginning or, if   is also specified, at the end of the search range. No match at the beginning or end means nothing is found, even if a matching sequence of bytes occurs elsewhere in the data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags/anchored
	kCFDataSearchAnchored CFDataSearchFlags = 0
)

// CFDateFormatterStyle - Data type for predefined date and time format styles.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle
type CFDateFormatterStyle uint

// CFLocaleLanguageDirection - These constants describe the text direction for a language. They are returned by the functions 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection
type CFLocaleLanguageDirection uint

// CFRunLoopRunResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult
type CFRunLoopRunResult uint

const (
	// kCFRunLoopRunFinished - The running run loop mode has no sources or timers to process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult/finished
	kCFRunLoopRunFinished CFRunLoopRunResult = 0
	// kCFRunLoopRunHandledSource - A source has been processed. This value is returned only if the run loop was told to run only until a source was processed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult/handledSource
	kCFRunLoopRunHandledSource CFRunLoopRunResult = 0
	// kCFRunLoopRunStopped -  was called on the run loop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult/stopped
	kCFRunLoopRunStopped CFRunLoopRunResult = 0
	// kCFRunLoopRunTimedOut - The specified time interval for running the run loop has passed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult/timedOut
	kCFRunLoopRunTimedOut CFRunLoopRunResult = 0
)

// CFStringBuiltInEncodings - Encodings that are built-in on all platforms on which macOS runs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings
type CFStringBuiltInEncodings uint

const (
	// kCFStringEncodingASCII - An encoding constant that identifies the ASCII encoding (decimal values 0 through 127).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/ASCII
	kCFStringEncodingASCII CFStringBuiltInEncodings = 0
	// kCFStringEncodingUTF8 - An encoding constant that identifies the UTF 8 encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF8
	kCFStringEncodingUTF8 CFStringBuiltInEncodings = 0
	// kCFStringEncodingISOLatin1 - An encoding constant that identifies the ISO Latin 1 encoding (ISO 8859-1)
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/isoLatin1
	kCFStringEncodingISOLatin1 CFStringBuiltInEncodings = 0
	// kCFStringEncodingMacRoman - An encoding constant that identifies the Mac Roman encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/macRoman
	kCFStringEncodingMacRoman CFStringBuiltInEncodings = 0
	// kCFStringEncodingNextStepLatin - An encoding constant that identifies the NextStep/OpenStep encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/nextStepLatin
	kCFStringEncodingNextStepLatin CFStringBuiltInEncodings = 0
	// kCFStringEncodingWindowsLatin1 - An encoding constant that identifies the Windows Latin 1 encoding (ANSI codepage 1252).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/windowsLatin1
	kCFStringEncodingWindowsLatin1 CFStringBuiltInEncodings = 0
)

// CFStringCompareFlags - A 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags
type CFStringCompareFlags uint

// CFStringEncodings - Index type for constants used to specify external string encodings.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings
type CFStringEncodings uint

// CFTimeZoneNameStyle - Index type for constants used to specify styles of time zone names.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle
type CFTimeZoneNameStyle uint

// CFURLBookmarkCreationOptions - Type for bookmark data creation options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions
type CFURLBookmarkCreationOptions uint

const (
	// kCFURLBookmarkCreationMinimalBookmarkMask - Specifies that an alias created with the bookmark data be created with minimal information, which may make it smaller but still able to resolve in certain ways.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/minimalBookmarkMask
	kCFURLBookmarkCreationMinimalBookmarkMask CFURLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationPreferFileIDResolutionMask - Specifies that an alias created with the bookmark data prefers resolving with its embedded file ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/preferFileIDResolutionMask
	kCFURLBookmarkCreationPreferFileIDResolutionMask CFURLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess - When combined with the   option, specifies that you want to create a security-scoped bookmark that, when resolved, provides a security-scoped URL allowing read-only access to a file-system resource; for use in an app that adopts App Sandbox.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/securityScopeAllowOnlyReadAccess
	kCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess CFURLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationSuitableForBookmarkFile - Specifies that the bookmark data include properties required to create Finder alias files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/suitableForBookmarkFile
	kCFURLBookmarkCreationSuitableForBookmarkFile CFURLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationWithSecurityScope - Specifies that you want to create a security-scoped bookmark that, when resolved, provides a security-scoped URL allowing read/write access to a file-system resource; for use in an app that adopts App Sandbox.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/withSecurityScope
	kCFURLBookmarkCreationWithSecurityScope CFURLBookmarkCreationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/withoutImplicitSecurityScope
	kCFURLBookmarkCreationWithoutImplicitSecurityScope CFURLBookmarkCreationOptions = 0
)

// CFURLBookmarkResolutionOptions - Type for bookmark data resolution options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions
type CFURLBookmarkResolutionOptions uint

const (
	// kCFURLBookmarkResolutionWithSecurityScope - Specifies that the security scope, applied to the bookmark when it was created, should be used during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfurlBookmarkResolutionWithSecurityScope
	kCFURLBookmarkResolutionWithSecurityScope CFURLBookmarkResolutionOptions = 0
)

// CFURLComponentType - The types of components in a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType
type CFURLComponentType uint

const (
	// kCFURLComponentHost - The URL’s host.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/host
	kCFURLComponentHost CFURLComponentType = 0
	// kCFURLComponentPath - The URL’s path component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/path
	kCFURLComponentPath CFURLComponentType = 0
)

// CFURLError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError
type CFURLError uint

// CFURLPathStyle - Options you can use to determine how CFURL functions parse a file system path name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle
type CFURLPathStyle uint

const (
	// kCFURLHFSPathStyle - Indicates a HFS style path name. Components are colon delimited. A leading colon indicates a relative path, otherwise the first path component denotes the volume.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle/cfurlhfsPathStyle
	kCFURLHFSPathStyle CFURLPathStyle = 0
)

// CGRectEdge enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge
type CGRectEdge uint

// LSAcceptanceFlags - The specification that determines whether an app can accept (open) an item.
//
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lsacceptanceflags
type LSAcceptanceFlags uint

// LSHandlerOptions - The specification that controls the selection of handlers.
//
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lshandleroptions
type LSHandlerOptions uint

// LSItemInfoFlags - The specification that provides information about an item.
//
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lsiteminfoflags
type LSItemInfoFlags uint

// LSLaunchFlags - The specification for launching an app.
//
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lslaunchflags
type LSLaunchFlags uint

// LSRequestedInfo - The specification that controls which information to obtain about an item.
//
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lsrequestedinfo
type LSRequestedInfo uint

// LSRolesMask - The specification that sets the desired role or roles for an app to claim for an item or a family of items.
//
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lsrolesmask
type LSRolesMask uint

// DirectionalRectEdge - Constants that specify an edge or a set of edges, taking the user interface layout direction into account.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSDirectionalRectEdge
type DirectionalRectEdge uint

// RectAlignment - Constants that specify alignment to an edge or a set of edges depending on the user interface layout direction.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSRectAlignment
type RectAlignment uint

// TextLayoutManagerSegmentOptions - Values that describe where and how the framework extends segments of a selection.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextLayoutManager/SegmentOptions
type TextLayoutManagerSegmentOptions uint

// TextLayoutManagerSegmentType - Values that describe the rendering of selection boundaries.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextLayoutManager/SegmentType
type TextLayoutManagerSegmentType uint

// TextSelectionNavigationDirection - Values that describe the direction of a selection.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextSelectionNavigation/Direction
type TextSelectionNavigationDirection uint

// UIGuidedAccessErrorCode - Error codes for Guided Access.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/GuidedAccessError/Code
type UIGuidedAccessErrorCode uint

// UIAccessibilityHearingDeviceEar - Constants that specify how a person is using a hearing device.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/HearingDeviceEar
type UIAccessibilityHearingDeviceEar uint

// UIAccessibilityContainerType - Constants that indicate the type of content in a data-based container.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibilityContainerType
type UIAccessibilityContainerType uint

// UIApplicationCategory - Constants that describe the types of apps in the system.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplication/Category
type UIApplicationCategory uint

const (
	// UIApplicationCategoryWebBrowser - The app is a web browser.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplication/Category/webBrowser
	UIApplicationCategoryWebBrowser UIApplicationCategory = 0
)

// UIApplicationState - Constants that indicate the running states of an app.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplication/State
type UIApplicationState uint

const (
	// UIApplicationStateActive - The app is running in the foreground and currently receiving events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplication/State/active
	UIApplicationStateActive UIApplicationState = 0
	// UIApplicationStateBackground - The app is running in the background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplication/State/background
	UIApplicationStateBackground UIApplicationState = 0
	// UIApplicationStateInactive - The app is running in the foreground but isn’t receiving events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplication/State/inactive
	UIApplicationStateInactive UIApplicationState = 0
)

// UIApplicationCategoryDefaultStatus - The default status of an application for some category.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplicationCategoryDefaultStatus
type UIApplicationCategoryDefaultStatus int

const (
	// UIApplicationCategoryDefaultStatusIsDefault - The application is the default for the category.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplicationCategoryDefaultStatus/UIApplicationCategoryDefaultStatusIsDefault
	UIApplicationCategoryDefaultStatusIsDefault UIApplicationCategoryDefaultStatus = 0
)

// UIAxis - A structure that specifies the layout axes.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAxis
type UIAxis uint

// UIBackgroundRefreshStatus - Constants that indicate whether background execution is enabled for the app.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIBackgroundRefreshStatus
type UIBackgroundRefreshStatus uint

const (
	// UIBackgroundRefreshStatusAvailable - Background updates are available for the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIBackgroundRefreshStatus/available
	UIBackgroundRefreshStatusAvailable UIBackgroundRefreshStatus = 0
	// UIBackgroundRefreshStatusRestricted - Background updates are unavailable and the user cannot enable them again.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIBackgroundRefreshStatus/restricted
	UIBackgroundRefreshStatusRestricted UIBackgroundRefreshStatus = 0
)

// UIBaselineAdjustment - Vertical adjustment options.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIBaselineAdjustment
type UIBaselineAdjustment uint

// UIColorProminence - A type that indicates the prominence of a color in the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIColor/Prominence-swift.enum
type UIColorProminence uint

const (
	// UIColorProminencePrimary - A color with a primary prominence, the most prominent in the interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIColor/Prominence-swift.enum/primary
	UIColorProminencePrimary UIColorProminence = 0
	// UIColorProminenceQuaternary - A color with a quaternary prominence, the least prominent in the interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIColor/Prominence-swift.enum/quaternary
	UIColorProminenceQuaternary UIColorProminence = 0
	// UIColorProminenceSecondary - A color with a secondary prominence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIColor/Prominence-swift.enum/secondary
	UIColorProminenceSecondary UIColorProminence = 0
	// UIColorProminenceTertiary - A color with a tertiary prominence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIColor/Prominence-swift.enum/tertiary
	UIColorProminenceTertiary UIColorProminence = 0
)

// UIControlEvents - Constants describing the types of events possible for controls.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIControl/Event
type UIControlEvents uint

const (
	// UIControlEventTouchDown - A touch-down event in the control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIControl/Event/touchDown
	UIControlEventTouchDown UIControlEvents = 0
)

// UIControlState - Constants describing the state of a control.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIControl/State-swift.struct
type UIControlState uint

// UIDocumentChangeKind - Constants that specify the kind of change to a document.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/ChangeKind
type UIDocumentChangeKind uint

const (
	// UIDocumentChangeCleared - The document is cleared of outstanding changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/ChangeKind/cleared
	UIDocumentChangeCleared UIDocumentChangeKind = 0
	// UIDocumentChangeDone - A change has been made to the document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/ChangeKind/done
	UIDocumentChangeDone UIDocumentChangeKind = 0
	// UIDocumentChangeRedone - An undone change to the document has been redone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/ChangeKind/redone
	UIDocumentChangeRedone UIDocumentChangeKind = 0
	// UIDocumentChangeUndone - A change to the document has been undone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/ChangeKind/undone
	UIDocumentChangeUndone UIDocumentChangeKind = 0
)

// UIDocumentSaveOperation - Constants that specify the type of save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/SaveOperation
type UIDocumentSaveOperation uint

const (
	// UIDocumentSaveForCreating - The document is being saved for the first time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/SaveOperation/forCreating
	UIDocumentSaveForCreating UIDocumentSaveOperation = 0
	// UIDocumentSaveForOverwriting - The document is being saved by overwriting the current version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/SaveOperation/forOverwriting
	UIDocumentSaveForOverwriting UIDocumentSaveOperation = 0
)

// UIDocumentState - Constants that specify the document state.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/State
type UIDocumentState uint

const (
	// UIDocumentStateClosed - There was an error in reading the document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/State/closed
	UIDocumentStateClosed UIDocumentState = 0
	// UIDocumentStateEditingDisabled - The document is busy and it isn’t currently safe for user edits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/State/editingDisabled
	UIDocumentStateEditingDisabled UIDocumentState = 0
	// UIDocumentStateInConflict - Conflicts exist for the document file located at the file URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/State/inConflict
	UIDocumentStateInConflict UIDocumentState = 0
	// UIDocumentStateNormal - The document is open, editing is enabled, and there are no conflicts or errors associated with it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/State/normal
	UIDocumentStateNormal UIDocumentState = 0
	// UIDocumentStateProgressAvailable - The document is being downloaded or uploaded and progress information is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/State/progressAvailable
	UIDocumentStateProgressAvailable UIDocumentState = 0
	// UIDocumentStateSavingError - There was an error in saving or reverting the document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/State/savingError
	UIDocumentStateSavingError UIDocumentState = 0
)

// UIDocumentPickerMode - Modes that define the type of file transfer operation that the document picker uses.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocumentPickerMode
type UIDocumentPickerMode uint

const (
	// UIDocumentPickerModeExportToService - The document picker exports a local file to a destination outside the app’s sandbox.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocumentPickerMode/exportToService
	UIDocumentPickerModeExportToService UIDocumentPickerMode = 0
	// UIDocumentPickerModeOpen - The document picker opens an external file outside the app’s sandbox.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocumentPickerMode/open
	UIDocumentPickerModeOpen UIDocumentPickerMode = 0
)

// UIFontDescriptorSymbolicTraits - Constants that describe the stylistic aspects of a font.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIFontDescriptor/SymbolicTraits-swift.struct
type UIFontDescriptorSymbolicTraits uint

const (
	// UIFontDescriptorTraitExpanded - The font’s characters have an expanded width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIFontDescriptor/SymbolicTraits-swift.struct/traitExpanded
	UIFontDescriptorTraitExpanded UIFontDescriptorSymbolicTraits = 0
	// UIFontDescriptorTraitLooseLeading - The font uses a leading value that’s greater than the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIFontDescriptor/SymbolicTraits-swift.struct/traitLooseLeading
	UIFontDescriptorTraitLooseLeading UIFontDescriptorSymbolicTraits = 0
	// UIFontDescriptorTraitMonoSpace - The font’s characters all have the same width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIFontDescriptor/SymbolicTraits-swift.struct/traitMonoSpace
	UIFontDescriptorTraitMonoSpace UIFontDescriptorSymbolicTraits = 0
)

// UIGuidedAccessAccessibilityFeature - Constants that describe accessibility features for Guided Access.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIGuidedAccessAccessibilityFeature
type UIGuidedAccessAccessibilityFeature uint

// UIInterfaceOrientation - Constants that specify the orientation of the app’s user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIInterfaceOrientation
type UIInterfaceOrientation uint

// UIInterfaceOrientationMask - Constants that specify a view controller’s supported interface orientations.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIInterfaceOrientationMask
type UIInterfaceOrientationMask uint

// UILineBreakMode - Options for wrapping and truncating text.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UILineBreakMode
type UILineBreakMode int

// UISceneCollectionJoinBehavior - A set of behaviors that specify how a new scene joins a scene collection.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UISceneCollectionJoinBehavior
type UISceneCollectionJoinBehavior uint

// UISceneErrorCode - Error codes for issues with scenes.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UISceneError/Code
type UISceneErrorCode uint

// UITableViewScrollPosition - The position in the table view (top, middle, bottom) to scroll a specified row to.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UITableView/ScrollPosition
type UITableViewScrollPosition uint

// UITableViewSelfSizingInvalidation - Constants that describe modes for invalidating the size of self-sizing table view cells.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UITableView/SelfSizingInvalidation-swift.enum
type UITableViewSelfSizingInvalidation uint

// UITableViewSeparatorInsetReference - Constants that indicate how to interpret the separator inset value of a table view.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UITableView/SeparatorInsetReference-swift.enum
type UITableViewSeparatorInsetReference uint

// UITableViewStyle - Constants for the table view styles.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UITableView/Style-swift.enum
type UITableViewStyle uint

// UITableViewCellSeparatorStyle - The style for cells to use as separators.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UITableViewCell/SeparatorStyle
type UITableViewCellSeparatorStyle uint

// UITableViewContentHuggingElements - Constants that determine which types of items in a table view tightly hug their content.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UITableViewContentHuggingElements
type UITableViewContentHuggingElements uint

// UITextAlignment - Options for aligning text horizontally.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UITextAlignment
type UITextAlignment int

// UITextItemContentType - Constants that describe and capture the type of content a text item represents along with a specific related value.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UITextItemContentType
type UITextItemContentType int

// UIUserInterfaceLayoutDirection - Constants that specify the directional flow of the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIUserInterfaceLayoutDirection
type UIUserInterfaceLayoutDirection uint

const (
	// UIUserInterfaceLayoutDirectionLeftToRight - The layout direction is left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIUserInterfaceLayoutDirection/leftToRight
	UIUserInterfaceLayoutDirectionLeftToRight UIUserInterfaceLayoutDirection = 0
)

// TextScalingType - Constants that specify the text scaling.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextScalingType
type TextScalingType uint

const (
	// TextScalingiOS - Font sizes throughout the document appear visually similar to how they would render in iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextScalingType/iOS
	TextScalingiOS TextScalingType = 0
	// TextScalingStandard - Font sizes throughout the document appear visually similar to how they would render in macOS and non-Apple platforms.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextScalingType/standard
	TextScalingStandard TextScalingType = 0
)

// TextWritingDirection - Options for specifying text-writing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextWritingDirection
type TextWritingDirection uint

const (
	// TextWritingDirectionEmbedding - Text is embedded in text with another writing direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextWritingDirection/embedding
	TextWritingDirectionEmbedding TextWritingDirection = 0
	// TextWritingDirectionOverride - Enables character types with inherent directionality to be overridden when required for special cases, such as for part numbers made of mixed English, digits, and Hebrew letters to be written from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextWritingDirection/override
	TextWritingDirectionOverride TextWritingDirection = 0
)

// UnderlineStyle - Constants for the underline style and strikethrough style attribute keys.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSUnderlineStyle
type UnderlineStyle uint

const (
	// UnderlineStyleByWord - Draw the line only beneath or through words, not whitespace.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/byWord
	UnderlineStyleByWord UnderlineStyle = 0
	// UnderlineStyleNone - Don’t draw a line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSUnderlineStyle/NSUnderlineStyleNone
	UnderlineStyleNone UnderlineStyle = 0
	// UnderlineStyleDouble - Draw a double line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSUnderlineStyle/double
	UnderlineStyleDouble UnderlineStyle = 0
	// UnderlineStylePatternDash - Draw a line of dashes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSUnderlineStyle/patternDash
	UnderlineStylePatternDash UnderlineStyle = 0
	// UnderlineStylePatternDashDot - Draw a line of alternating dashes and dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSUnderlineStyle/patternDashDot
	UnderlineStylePatternDashDot UnderlineStyle = 0
	// UnderlineStylePatternDashDotDot - Draw a line of alternating dashes and two dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSUnderlineStyle/patternDashDotDot
	UnderlineStylePatternDashDotDot UnderlineStyle = 0
	// UnderlineStylePatternDot - Draw a line of dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSUnderlineStyle/patternDot
	UnderlineStylePatternDot UnderlineStyle = 0
	// UnderlineStyleThick - Draw a thick line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSUnderlineStyle/thick
	UnderlineStyleThick UnderlineStyle = 0
)

// WritingDirectionFormatType - Constants for the writing direction attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSWritingDirectionFormatType
type WritingDirectionFormatType uint

const (
	// WritingDirectionEmbedding - Text is embedded in text with another writing direction. For example, an English quotation in the middle of an Arabic sentence could be marked as being embedded left-to-right text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirectionFormatType/embedding
	WritingDirectionEmbedding WritingDirectionFormatType = 0
)


