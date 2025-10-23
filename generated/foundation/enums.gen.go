// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
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

// FileManagerSupportedSyncControls - An option set of the sync controls available for an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls
type FileManagerSupportedSyncControls uint

const (
	FileManagerSupportedSyncControlsPauseSync FileManagerSupportedSyncControls = 1
	FileManagerSupportedSyncControlsFailUploadOnConflict FileManagerSupportedSyncControls = 2
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

// PresentationIntentTableColumnAlignment - An enumeration of values for aligning the contents of table columns.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment
type PresentationIntentTableColumnAlignment int

const (
	PresentationIntentTableColumnAlignmentLeft PresentationIntentTableColumnAlignment = 0
	PresentationIntentTableColumnAlignmentCenter PresentationIntentTableColumnAlignment = 1
	PresentationIntentTableColumnAlignmentRight PresentationIntentTableColumnAlignment = 2
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
	DirectoryEnumerationSkipsSubdirectoryDescendants DirectoryEnumerationOptions = 1
	DirectoryEnumerationSkipsPackageDescendants DirectoryEnumerationOptions = 2
	DirectoryEnumerationSkipsHiddenFiles DirectoryEnumerationOptions = 4
	DirectoryEnumerationIncludesDirectoriesPostOrder DirectoryEnumerationOptions = 5
	DirectoryEnumerationProducesRelativePathURLs DirectoryEnumerationOptions = 6
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

// CGRectEdge enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge
type CGRectEdge uint

// TextScalingType - Constants that specify the text scaling.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextScalingType
type TextScalingType uint

// TextWritingDirection - Options for specifying text-writing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextWritingDirection
type TextWritingDirection uint

// UnderlineStyle - Constants for the underline style and strikethrough style attribute keys.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSUnderlineStyle
type UnderlineStyle uint

// WritingDirectionFormatType - Constants for the writing direction attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSWritingDirectionFormatType
type WritingDirectionFormatType uint


