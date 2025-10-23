// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

// Enum types and constants
// CalendarUnit - CFCalendarUnit constants are used to specify calendrical units, such as day or month, in various calendar calculations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit
type CalendarUnit uint

const (
	// kCFCalendarUnitDay - Specifies the day unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/day
	kCFCalendarUnitDay CalendarUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/dayOfYear
	kCFCalendarUnitDayOfYear CalendarUnit = 0
	// kCFCalendarUnitEra - Specifies the era unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/era
	kCFCalendarUnitEra CalendarUnit = 0
	// kCFCalendarUnitHour - Specifies the hour unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/hour
	kCFCalendarUnitHour CalendarUnit = 0
	// kCFCalendarUnitMinute - Specifies the minute unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/minute
	kCFCalendarUnitMinute CalendarUnit = 0
	// kCFCalendarUnitMonth - Specifies the month unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/month
	kCFCalendarUnitMonth CalendarUnit = 0
	// kCFCalendarUnitQuarter - Specifies the quarter-year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/quarter
	kCFCalendarUnitQuarter CalendarUnit = 0
	// kCFCalendarUnitSecond - Specifies the second unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/second
	kCFCalendarUnitSecond CalendarUnit = 0
	// kCFCalendarUnitWeek - Specifies the week unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/week
	kCFCalendarUnitWeek CalendarUnit = 0
	// kCFCalendarUnitWeekOfMonth - Specifies the original week of a month calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/weekOfMonth
	kCFCalendarUnitWeekOfMonth CalendarUnit = 0
	// kCFCalendarUnitWeekOfYear - Specifies the original week of the year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/weekOfYear
	kCFCalendarUnitWeekOfYear CalendarUnit = 0
	// kCFCalendarUnitWeekday - Specifies the weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/weekday
	kCFCalendarUnitWeekday CalendarUnit = 0
	// kCFCalendarUnitWeekdayOrdinal - Specifies the ordinal weekday unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/weekdayOrdinal
	kCFCalendarUnitWeekdayOrdinal CalendarUnit = 0
	// kCFCalendarUnitYear - Specifies the year unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/year
	kCFCalendarUnitYear CalendarUnit = 0
	// kCFCalendarUnitYearForWeekOfYear - Specifies the relative year for a week within a year calendar unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/yearForWeekOfYear
	kCFCalendarUnitYearForWeekOfYear CalendarUnit = 0
)

// CharacterSetPredefinedSet - Defines a predefined character set.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet
type CharacterSetPredefinedSet uint

const (
	// kCFCharacterSetAlphaNumeric - Alpha Numeric character set (Unicode General Category L*, M*, & N*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/alphaNumeric
	kCFCharacterSetAlphaNumeric CharacterSetPredefinedSet = 0
	// kCFCharacterSetCapitalizedLetter - Titlecase character set (Unicode General Category Lt).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/capitalizedLetter
	kCFCharacterSetCapitalizedLetter CharacterSetPredefinedSet = 0
	// kCFCharacterSetControl - Control character set (Unicode General Category Cc and Cf).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/control
	kCFCharacterSetControl CharacterSetPredefinedSet = 0
	// kCFCharacterSetDecimalDigit - Decimal digit character set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/decimalDigit
	kCFCharacterSetDecimalDigit CharacterSetPredefinedSet = 0
	// kCFCharacterSetDecomposable - Canonically decomposable character set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/decomposable
	kCFCharacterSetDecomposable CharacterSetPredefinedSet = 0
	// kCFCharacterSetIllegal - Illegal character set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/illegal
	kCFCharacterSetIllegal CharacterSetPredefinedSet = 0
	// kCFCharacterSetLetter - Letter character set (Unicode General Category L* & M*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/letter
	kCFCharacterSetLetter CharacterSetPredefinedSet = 0
	// kCFCharacterSetLowercaseLetter - Lowercase character set (Unicode General Category Ll).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/lowercaseLetter
	kCFCharacterSetLowercaseLetter CharacterSetPredefinedSet = 0
	// kCFCharacterSetNewline - Newline character set ( ,  ,  , and  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/newline
	kCFCharacterSetNewline CharacterSetPredefinedSet = 0
	// kCFCharacterSetNonBase - Non-base character set (Unicode General Category M*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/nonBase
	kCFCharacterSetNonBase CharacterSetPredefinedSet = 0
	// kCFCharacterSetPunctuation - Punctuation character set (Unicode General Category P*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/punctuation
	kCFCharacterSetPunctuation CharacterSetPredefinedSet = 0
	// kCFCharacterSetSymbol - Symbol character set (Unicode General Category S*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/symbol
	kCFCharacterSetSymbol CharacterSetPredefinedSet = 0
	// kCFCharacterSetUppercaseLetter - Uppercase character set (Unicode General Category Lu and Lt).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/uppercaseLetter
	kCFCharacterSetUppercaseLetter CharacterSetPredefinedSet = 0
	// kCFCharacterSetWhitespace - Whitespace character set (Unicode General Category Zs and U0009 CHARACTER TABULATION).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/whitespace
	kCFCharacterSetWhitespace CharacterSetPredefinedSet = 0
	// kCFCharacterSetWhitespaceAndNewline - Whitespace and Newline character set (Unicode General Category Z*,  , and  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/whitespaceAndNewline
	kCFCharacterSetWhitespaceAndNewline CharacterSetPredefinedSet = 0
)

// ComparisonResult - Constants returned by comparison functions, indicating whether a value is equal to, less than, or greater than another value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult
type ComparisonResult uint

const (
	// kCFCompareEqualTo - Returned by a comparison function if the first value is equal to the second value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult/compareEqualTo
	kCFCompareEqualTo ComparisonResult = 0
	// kCFCompareGreaterThan - Returned by a comparison function if the first value is greater than the second value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult/compareGreaterThan
	kCFCompareGreaterThan ComparisonResult = 0
	// kCFCompareLessThan - Returned by a comparison function if the first value is less than the second value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult/compareLessThan
	kCFCompareLessThan ComparisonResult = 0
)

// DataSearchFlags - A 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags
type DataSearchFlags uint

const (
	// kCFDataSearchAnchored - Performs searching only on bytes at the beginning or, if   is also specified, at the end of the search range. No match at the beginning or end means nothing is found, even if a matching sequence of bytes occurs elsewhere in the data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags/anchored
	kCFDataSearchAnchored DataSearchFlags = 0
	// kCFDataSearchBackwards - Performs searching from the end of the range toward the beginning.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags/backwards
	kCFDataSearchBackwards DataSearchFlags = 0
)

// DateFormatterStyle - Data type for predefined date and time format styles.
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
	// kCFDateFormatterMediumStyle - Specifies a medium style, typically with abbreviated text, such as “Nov 23, 1937”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/mediumStyle
	kCFDateFormatterMediumStyle DateFormatterStyle = 0
	// kCFDateFormatterNoStyle - Specifies no output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/noStyle
	kCFDateFormatterNoStyle DateFormatterStyle = 0
	// kCFDateFormatterShortStyle - Specifies a short style, typically numeric only, such as “11/23/37” or “3:30pm”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/shortStyle
	kCFDateFormatterShortStyle DateFormatterStyle = 0
)

// FileSecurityClearOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions
type FileSecurityClearOptions uint

const (
	// kCFFileSecurityClearAccessControlList - Clear the access control list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/accessControlList
	kCFFileSecurityClearAccessControlList FileSecurityClearOptions = 0
	// kCFFileSecurityClearGroup - Clear the (POSIX) group ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/group
	kCFFileSecurityClearGroup FileSecurityClearOptions = 0
	// kCFFileSecurityClearGroupUUID - Clear the group UUID (for the access control list).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/groupUUID
	kCFFileSecurityClearGroupUUID FileSecurityClearOptions = 0
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

// GregorianUnitFlags - These option flags are used as a mask to indicate a specific set of fields in the CFGregorianDate or CFGregorianUnits structures.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags
type GregorianUnitFlags uint

const (
	// kCFGregorianAllUnits - Specifies all fields.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/allUnits
	kCFGregorianAllUnits GregorianUnitFlags = 0
	// kCFGregorianUnitsDays - Specifies the day field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsDays
	kCFGregorianUnitsDays GregorianUnitFlags = 0
	// kCFGregorianUnitsHours - Specifies the hours field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsHours
	kCFGregorianUnitsHours GregorianUnitFlags = 0
	// kCFGregorianUnitsMinutes - Specifies the minutes field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsMinutes
	kCFGregorianUnitsMinutes GregorianUnitFlags = 0
	// kCFGregorianUnitsMonths - Specifies the month field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsMonths
	kCFGregorianUnitsMonths GregorianUnitFlags = 0
	// kCFGregorianUnitsSeconds - Specifies the seconds field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsSeconds
	kCFGregorianUnitsSeconds GregorianUnitFlags = 0
	// kCFGregorianUnitsYears - Specifies the year field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsYears
	kCFGregorianUnitsYears GregorianUnitFlags = 0
)

// ISO8601DateFormatOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions
type ISO8601DateFormatOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withColonSeparatorInTime
	kCFISO8601DateFormatWithColonSeparatorInTime ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withColonSeparatorInTimeZone
	kCFISO8601DateFormatWithColonSeparatorInTimeZone ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withDashSeparatorInDate
	kCFISO8601DateFormatWithDashSeparatorInDate ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withDay
	kCFISO8601DateFormatWithDay ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withFractionalSeconds
	kCFISO8601DateFormatWithFractionalSeconds ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withFullDate
	kCFISO8601DateFormatWithFullDate ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withFullTime
	kCFISO8601DateFormatWithFullTime ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withInternetDateTime
	kCFISO8601DateFormatWithInternetDateTime ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withMonth
	kCFISO8601DateFormatWithMonth ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withSpaceBetweenDateAndTime
	kCFISO8601DateFormatWithSpaceBetweenDateAndTime ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withTime
	kCFISO8601DateFormatWithTime ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withTimeZone
	kCFISO8601DateFormatWithTimeZone ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withWeekOfYear
	kCFISO8601DateFormatWithWeekOfYear ISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withYear
	kCFISO8601DateFormatWithYear ISO8601DateFormatOptions = 0
)

// LocaleLanguageDirection - These constants describe the text direction for a language. They are returned by the functions 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection
type LocaleLanguageDirection uint

const (
	// kCFLocaleLanguageDirectionBottomToTop - The language direction is from bottom to top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/bottomToTop
	kCFLocaleLanguageDirectionBottomToTop LocaleLanguageDirection = 0
	// kCFLocaleLanguageDirectionLeftToRight - The language direction is from left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/leftToRight
	kCFLocaleLanguageDirectionLeftToRight LocaleLanguageDirection = 0
	// kCFLocaleLanguageDirectionRightToLeft - The language direction is from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/rightToLeft
	kCFLocaleLanguageDirectionRightToLeft LocaleLanguageDirection = 0
	// kCFLocaleLanguageDirectionTopToBottom - The language direction is from top to bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/topToBottom
	kCFLocaleLanguageDirectionTopToBottom LocaleLanguageDirection = 0
	// kCFLocaleLanguageDirectionUnknown - The direction of the language is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/unknown
	kCFLocaleLanguageDirectionUnknown LocaleLanguageDirection = 0
)

// NotificationSuspensionBehavior - Suspension flags that indicate how distributed notifications should be handled when the receiving application is in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior
type NotificationSuspensionBehavior uint

const (
	// NotificationSuspensionBehaviorCoalesce - The server will only queue the last notification of the specified name and object; earlier notifications are dropped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior/coalesce
	NotificationSuspensionBehaviorCoalesce NotificationSuspensionBehavior = 0
	// NotificationSuspensionBehaviorDeliverImmediately - The server will deliver notifications of the specified name and object whether or not the application is in the background. When a notification with this suspension behavior is matched, it has the effect of first flushing any queued notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior/deliverImmediately
	NotificationSuspensionBehaviorDeliverImmediately NotificationSuspensionBehavior = 0
	// NotificationSuspensionBehaviorDrop - The server will not queue any notifications of the specified name and object while the receiving application is in the background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior/drop
	NotificationSuspensionBehaviorDrop NotificationSuspensionBehavior = 0
	// NotificationSuspensionBehaviorHold - The server will hold all matching notifications until the queue has been filled (queue size determined by the server) at which point the server may flush queued notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior/hold
	NotificationSuspensionBehaviorHold NotificationSuspensionBehavior = 0
)

// NumberFormatterOptionFlags - Type for constants specifying how numbers should be parsed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterOptionFlags
type NumberFormatterOptionFlags uint

const (
	// kCFNumberFormatterParseIntegersOnly - Specifies that only integers should be parsed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterOptionFlags/parseIntegersOnly
	kCFNumberFormatterParseIntegersOnly NumberFormatterOptionFlags = 0
)

// NumberFormatterPadPosition - Type for constants specifying how numbers should be padded.
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

// NumberFormatterRoundingMode - These constants are used to specify how numbers should be rounded.
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

// NumberFormatterStyle - Type for constants specifying a formatter style.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle
type NumberFormatterStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/currencyAccountingStyle
	kCFNumberFormatterCurrencyAccountingStyle NumberFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/currencyISOCodeStyle
	kCFNumberFormatterCurrencyISOCodeStyle NumberFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/currencyPluralStyle
	kCFNumberFormatterCurrencyPluralStyle NumberFormatterStyle = 0
	// kCFNumberFormatterCurrencyStyle - Specifies a currency style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/currencyStyle
	kCFNumberFormatterCurrencyStyle NumberFormatterStyle = 0
	// kCFNumberFormatterDecimalStyle - Specifies a decimal style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/decimalStyle
	kCFNumberFormatterDecimalStyle NumberFormatterStyle = 0
	// kCFNumberFormatterNoStyle - Specifies no style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/noStyle
	kCFNumberFormatterNoStyle NumberFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/ordinalStyle
	kCFNumberFormatterOrdinalStyle NumberFormatterStyle = 0
	// kCFNumberFormatterPercentStyle - Specifies a percent style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/percentStyle
	kCFNumberFormatterPercentStyle NumberFormatterStyle = 0
	// kCFNumberFormatterScientificStyle - Specifies a scientific style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/scientificStyle
	kCFNumberFormatterScientificStyle NumberFormatterStyle = 0
	// kCFNumberFormatterSpellOutStyle - Specifies a spelled out format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/spellOutStyle
	kCFNumberFormatterSpellOutStyle NumberFormatterStyle = 0
)

// NumberType - Flags used by CFNumber to indicate the data type of a value.
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

// PropertyListFormat - Specifies the format of a property list.
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
	// kCFPropertyListXMLFormat_v1_0 - XML format version 1.0.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListFormat/xmlFormat_v1_0
	kCFPropertyListXMLFormat_v1_0 PropertyListFormat = 0
)

// PropertyListMutabilityOptions - Type for flags that determine the degree of mutability of newly created property lists.
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

// RunLoopActivity - Run loop activity stages in which run loop observers can be scheduled.
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
	// kCFRunLoopBeforeWaiting - Inside the event processing loop before the run loop sleeps, waiting for a source or timer to fire. This activity does not occur if   is called with a timeout of 0 seconds. It also does not occur in a particular iteration of the event processing loop if a version 0 source fires.
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

// RunLoopRunResult enum type
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

// SocketCallBackType - Types of socket activity that can cause the callback function of a CFSocket object to be called.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType
type SocketCallBackType uint

const (
	// kCFSocketAcceptCallBack - New connections will be automatically accepted and the callback is called with the data argument being a pointer to a   of the child socket. This callback is usable only with listening sockets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/acceptCallBack
	kCFSocketAcceptCallBack SocketCallBackType = 0
	// kCFSocketConnectCallBack - If a connection attempt is made in the background by calling   or   with a negative timeout value, this callback type is made when the connect finishes. In this case the data argument is either   or a pointer to an   error code, if the connect failed. This callback will never be sent more than once for a given socket.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/connectCallBack
	kCFSocketConnectCallBack SocketCallBackType = 0
	// kCFSocketDataCallBack - Incoming data will be read in chunks in the background and the callback is called with the data argument being a CFData object containing the read data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/dataCallBack
	kCFSocketDataCallBack SocketCallBackType = 0
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

// SocketError - Error codes for many CFSocket functions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketError
type SocketError uint

const (
	// kCFSocketError - The socket operation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketError/error
	kCFSocketError SocketError = 0
	// kCFSocketSuccess - The socket operation succeeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketError/success
	kCFSocketSuccess SocketError = 0
	// kCFSocketTimeout - The socket operation timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketError/timeout
	kCFSocketTimeout SocketError = 0
)

// StreamErrorDomain - Defines constants for values returned in the domain field of the 
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

// StreamEventType - Defines constants for stream-related events.
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

// StreamStatus - Constants that describe the status of a stream.
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

// StringBuiltInEncodings - Encodings that are built-in on all platforms on which macOS runs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings
type StringBuiltInEncodings uint

const (
	// kCFStringEncodingASCII - An encoding constant that identifies the ASCII encoding (decimal values 0 through 127).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/ASCII
	kCFStringEncodingASCII StringBuiltInEncodings = 0
	// kCFStringEncodingUTF16 - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF16Format encoding (alias of kCFStringEncodingUnicode).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF16
	kCFStringEncodingUTF16 StringBuiltInEncodings = 0
	// kCFStringEncodingUTF16BE - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF16BEFormat encoding. This constant specifies big-endian byte order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF16BE
	kCFStringEncodingUTF16BE StringBuiltInEncodings = 0
	// kCFStringEncodingUTF16LE - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF16LEFormat encoding. This constant specifies little-endian byte order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF16LE
	kCFStringEncodingUTF16LE StringBuiltInEncodings = 0
	// kCFStringEncodingUTF32 - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF32Format encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF32
	kCFStringEncodingUTF32 StringBuiltInEncodings = 0
	// kCFStringEncodingUTF32BE - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF32BEFormat encoding. This constant specifies big-endian byte order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF32BE
	kCFStringEncodingUTF32BE StringBuiltInEncodings = 0
	// kCFStringEncodingUTF32LE - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF32LEFormat encoding. This constant specifies little-endian byte order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF32LE
	kCFStringEncodingUTF32LE StringBuiltInEncodings = 0
	// kCFStringEncodingUTF8 - An encoding constant that identifies the UTF 8 encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF8
	kCFStringEncodingUTF8 StringBuiltInEncodings = 0
	// kCFStringEncodingISOLatin1 - An encoding constant that identifies the ISO Latin 1 encoding (ISO 8859-1)
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/isoLatin1
	kCFStringEncodingISOLatin1 StringBuiltInEncodings = 0
	// kCFStringEncodingMacRoman - An encoding constant that identifies the Mac Roman encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/macRoman
	kCFStringEncodingMacRoman StringBuiltInEncodings = 0
	// kCFStringEncodingNextStepLatin - An encoding constant that identifies the NextStep/OpenStep encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/nextStepLatin
	kCFStringEncodingNextStepLatin StringBuiltInEncodings = 0
	// kCFStringEncodingNonLossyASCII - An encoding constant that identifies non-lossy ASCII encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/nonLossyASCII
	kCFStringEncodingNonLossyASCII StringBuiltInEncodings = 0
	// kCFStringEncodingUnicode - An encoding constant that identifies the Unicode encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/unicode
	kCFStringEncodingUnicode StringBuiltInEncodings = 0
	// kCFStringEncodingWindowsLatin1 - An encoding constant that identifies the Windows Latin 1 encoding (ANSI codepage 1252).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/windowsLatin1
	kCFStringEncodingWindowsLatin1 StringBuiltInEncodings = 0
)

// StringCompareFlags - A 
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

// StringEncodings - Index type for constants used to specify external string encodings.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings
type StringEncodings uint

const (
	// kCFStringEncodingANSEL - ANSEL (ANSI Z39.47).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ANSEL
	kCFStringEncodingANSEL StringEncodings = 0
	// kCFStringEncodingCNS_11643_92_P1 - CNS 11643-1992 plane 1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/CNS_11643_92_P1
	kCFStringEncodingCNS_11643_92_P1 StringEncodings = 0
	// kCFStringEncodingCNS_11643_92_P2 - CNS 11643-1992 plane 2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/CNS_11643_92_P2
	kCFStringEncodingCNS_11643_92_P2 StringEncodings = 0
	// kCFStringEncodingCNS_11643_92_P3 - CNS 11643-1992 plane 3 (was plane 14 in 1986 version).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/CNS_11643_92_P3
	kCFStringEncodingCNS_11643_92_P3 StringEncodings = 0
	// kCFStringEncodingEBCDIC_CP037 - code page 037, extended EBCDIC (Latin-1 set) for US, Canada.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EBCDIC_CP037
	kCFStringEncodingEBCDIC_CP037 StringEncodings = 0
	// kCFStringEncodingEBCDIC_US - basic EBCDIC-US
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EBCDIC_US
	kCFStringEncodingEBCDIC_US StringEncodings = 0
	// kCFStringEncodingEUC_CN - ISO 646, GB 2312-80.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EUC_CN
	kCFStringEncodingEUC_CN StringEncodings = 0
	// kCFStringEncodingEUC_JP - ISO 646, 1-byte katakana, JIS 208, JIS 212.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EUC_JP
	kCFStringEncodingEUC_JP StringEncodings = 0
	// kCFStringEncodingEUC_KR - ISO 646, KS C 5601-1987.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EUC_KR
	kCFStringEncodingEUC_KR StringEncodings = 0
	// kCFStringEncodingEUC_TW - ISO 646, CNS 11643-1992 Planes 1-16.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EUC_TW
	kCFStringEncodingEUC_TW StringEncodings = 0
	// kCFStringEncodingGBK_95 - Annex to GB 13000-93; for Windows 95.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/GBK_95
	kCFStringEncodingGBK_95 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/GB_18030_2000
	kCFStringEncodingGB_18030_2000 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/GB_2312_80
	kCFStringEncodingGB_2312_80 StringEncodings = 0
	// kCFStringEncodingHZ_GB_2312 - HZ (RFC 1842, for Chinese mail & news).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/HZ_GB_2312
	kCFStringEncodingHZ_GB_2312 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_CN
	kCFStringEncodingISO_2022_CN StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_CN_EXT
	kCFStringEncodingISO_2022_CN_EXT StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_JP
	kCFStringEncodingISO_2022_JP StringEncodings = 0
	// kCFStringEncodingISO_2022_JP_1 - RFC 2237.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_JP_1
	kCFStringEncodingISO_2022_JP_1 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_JP_2
	kCFStringEncodingISO_2022_JP_2 StringEncodings = 0
	// kCFStringEncodingISO_2022_JP_3 - JIS X0213.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_JP_3
	kCFStringEncodingISO_2022_JP_3 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_KR
	kCFStringEncodingISO_2022_KR StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_C6226_78
	kCFStringEncodingJIS_C6226_78 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_X0201_76
	kCFStringEncodingJIS_X0201_76 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_X0208_83
	kCFStringEncodingJIS_X0208_83 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_X0208_90
	kCFStringEncodingJIS_X0208_90 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_X0212_90
	kCFStringEncodingJIS_X0212_90 StringEncodings = 0
	// kCFStringEncodingKOI8_R - Russian internet standard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/KOI8_R
	kCFStringEncodingKOI8_R StringEncodings = 0
	// kCFStringEncodingKOI8_U - RFC 2319, Ukrainian.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/KOI8_U
	kCFStringEncodingKOI8_U StringEncodings = 0
	// kCFStringEncodingKSC_5601_87 - Same as KSC 5601-92 without Johab annex.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/KSC_5601_87
	kCFStringEncodingKSC_5601_87 StringEncodings = 0
	// kCFStringEncodingUTF7 - kTextEncodingUnicodeDefault + kUnicodeUTF7Format RFC2152.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/UTF7
	kCFStringEncodingUTF7 StringEncodings = 0
	// kCFStringEncodingUTF7_IMAP - UTF-7 (IMAP folder variant) RFC3501.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/UTF7_IMAP
	kCFStringEncodingUTF7_IMAP StringEncodings = 0
	// kCFStringEncodingVISCII - RFC 1456, Vietnamese.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/VISCII
	kCFStringEncodingVISCII StringEncodings = 0
	// kCFStringEncodingBig5 - Big-5 (has variants)
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/big5
	kCFStringEncodingBig5 StringEncodings = 0
	// kCFStringEncodingBig5_E - Taiwan Big-5E standard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/big5_E
	kCFStringEncodingBig5_E StringEncodings = 0
	// kCFStringEncodingBig5_HKSCS_1999 - Big-5 with Hong Kong special char set supplement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/big5_HKSCS_1999
	kCFStringEncodingBig5_HKSCS_1999 StringEncodings = 0
	// kCFStringEncodingDOSArabic - Code page 864.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosArabic
	kCFStringEncodingDOSArabic StringEncodings = 0
	// kCFStringEncodingDOSBalticRim - Code page 775.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosBalticRim
	kCFStringEncodingDOSBalticRim StringEncodings = 0
	// kCFStringEncodingDOSCanadianFrench - Code page 863.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosCanadianFrench
	kCFStringEncodingDOSCanadianFrench StringEncodings = 0
	// kCFStringEncodingDOSChineseSimplif - Code page 936, also for Windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosChineseSimplif
	kCFStringEncodingDOSChineseSimplif StringEncodings = 0
	// kCFStringEncodingDOSChineseTrad - Code page 950, also for Windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosChineseTrad
	kCFStringEncodingDOSChineseTrad StringEncodings = 0
	// kCFStringEncodingDOSCyrillic - Code page 855, IBM Cyrillic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosCyrillic
	kCFStringEncodingDOSCyrillic StringEncodings = 0
	// kCFStringEncodingDOSGreek - Code page 737 (formerly code page 437G).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosGreek
	kCFStringEncodingDOSGreek StringEncodings = 0
	// kCFStringEncodingDOSGreek1 - Code page 851.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosGreek1
	kCFStringEncodingDOSGreek1 StringEncodings = 0
	// kCFStringEncodingDOSGreek2 - Code page 869, IBM Modern Greek.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosGreek2
	kCFStringEncodingDOSGreek2 StringEncodings = 0
	// kCFStringEncodingDOSHebrew - Code page 862.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosHebrew
	kCFStringEncodingDOSHebrew StringEncodings = 0
	// kCFStringEncodingDOSIcelandic - Code page 861.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosIcelandic
	kCFStringEncodingDOSIcelandic StringEncodings = 0
	// kCFStringEncodingDOSJapanese - Code page 932, also for Windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosJapanese
	kCFStringEncodingDOSJapanese StringEncodings = 0
	// kCFStringEncodingDOSKorean - Code page 949, also for Windows; Unified Hangul Code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosKorean
	kCFStringEncodingDOSKorean StringEncodings = 0
	// kCFStringEncodingDOSLatin1 - Code page 850, “Multilingual”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosLatin1
	kCFStringEncodingDOSLatin1 StringEncodings = 0
	// kCFStringEncodingDOSLatin2 - Code page 852, Slavic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosLatin2
	kCFStringEncodingDOSLatin2 StringEncodings = 0
	// kCFStringEncodingDOSLatinUS - Code page 437.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosLatinUS
	kCFStringEncodingDOSLatinUS StringEncodings = 0
	// kCFStringEncodingDOSNordic - Code page 865.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosNordic
	kCFStringEncodingDOSNordic StringEncodings = 0
	// kCFStringEncodingDOSPortuguese - Code page 860.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosPortuguese
	kCFStringEncodingDOSPortuguese StringEncodings = 0
	// kCFStringEncodingDOSRussian - Code page 866.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosRussian
	kCFStringEncodingDOSRussian StringEncodings = 0
	// kCFStringEncodingDOSThai - Code page 874, also for Windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosThai
	kCFStringEncodingDOSThai StringEncodings = 0
	// kCFStringEncodingDOSTurkish - Code page 857, IBM Turkish.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosTurkish
	kCFStringEncodingDOSTurkish StringEncodings = 0
	// kCFStringEncodingISOLatin10 - ISO 8859-16.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin10
	kCFStringEncodingISOLatin10 StringEncodings = 0
	// kCFStringEncodingISOLatin2 - ISO 8859-2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin2
	kCFStringEncodingISOLatin2 StringEncodings = 0
	// kCFStringEncodingISOLatin3 - ISO 8859-3.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin3
	kCFStringEncodingISOLatin3 StringEncodings = 0
	// kCFStringEncodingISOLatin4 - ISO 8859-4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin4
	kCFStringEncodingISOLatin4 StringEncodings = 0
	// kCFStringEncodingISOLatin5 - ISO 8859-9.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin5
	kCFStringEncodingISOLatin5 StringEncodings = 0
	// kCFStringEncodingISOLatin6 - ISO 8859-10.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin6
	kCFStringEncodingISOLatin6 StringEncodings = 0
	// kCFStringEncodingISOLatin7 - ISO 8859-13.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin7
	kCFStringEncodingISOLatin7 StringEncodings = 0
	// kCFStringEncodingISOLatin8 - ISO 8859-14.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin8
	kCFStringEncodingISOLatin8 StringEncodings = 0
	// kCFStringEncodingISOLatin9 - ISO 8859-15.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin9
	kCFStringEncodingISOLatin9 StringEncodings = 0
	// kCFStringEncodingISOLatinArabic - ISO 8859-6, =ASMO 708, =DOS CP 708.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinArabic
	kCFStringEncodingISOLatinArabic StringEncodings = 0
	// kCFStringEncodingISOLatinCyrillic - ISO 8859-5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinCyrillic
	kCFStringEncodingISOLatinCyrillic StringEncodings = 0
	// kCFStringEncodingISOLatinGreek - ISO 8859-7.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinGreek
	kCFStringEncodingISOLatinGreek StringEncodings = 0
	// kCFStringEncodingISOLatinHebrew - ISO 8859-8.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinHebrew
	kCFStringEncodingISOLatinHebrew StringEncodings = 0
	// kCFStringEncodingISOLatinThai - ISO 8859-11.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinThai
	kCFStringEncodingISOLatinThai StringEncodings = 0
	// kCFStringEncodingKSC_5601_92_Johab - KSC 5601-92 Johab annex.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ksc_5601_92_Johab
	kCFStringEncodingKSC_5601_92_Johab StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macArabic
	kCFStringEncodingMacArabic StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macArmenian
	kCFStringEncodingMacArmenian StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macBengali
	kCFStringEncodingMacBengali StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macBurmese
	kCFStringEncodingMacBurmese StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macCeltic
	kCFStringEncodingMacCeltic StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macCentralEurRoman
	kCFStringEncodingMacCentralEurRoman StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macChineseSimp
	kCFStringEncodingMacChineseSimp StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macChineseTrad
	kCFStringEncodingMacChineseTrad StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macCroatian
	kCFStringEncodingMacCroatian StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macCyrillic
	kCFStringEncodingMacCyrillic StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macDevanagari
	kCFStringEncodingMacDevanagari StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macDingbats
	kCFStringEncodingMacDingbats StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macEthiopic
	kCFStringEncodingMacEthiopic StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macExtArabic
	kCFStringEncodingMacExtArabic StringEncodings = 0
	// kCFStringEncodingMacFarsi - Like MacArabic but uses Farsi digits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macFarsi
	kCFStringEncodingMacFarsi StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGaelic
	kCFStringEncodingMacGaelic StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGeorgian
	kCFStringEncodingMacGeorgian StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGreek
	kCFStringEncodingMacGreek StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGujarati
	kCFStringEncodingMacGujarati StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGurmukhi
	kCFStringEncodingMacGurmukhi StringEncodings = 0
	// kCFStringEncodingMacHFS - Meta-value, should never appear in a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macHFS
	kCFStringEncodingMacHFS StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macHebrew
	kCFStringEncodingMacHebrew StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macIcelandic
	kCFStringEncodingMacIcelandic StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macInuit
	kCFStringEncodingMacInuit StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macJapanese
	kCFStringEncodingMacJapanese StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macKannada
	kCFStringEncodingMacKannada StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macKhmer
	kCFStringEncodingMacKhmer StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macKorean
	kCFStringEncodingMacKorean StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macLaotian
	kCFStringEncodingMacLaotian StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macMalayalam
	kCFStringEncodingMacMalayalam StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macMongolian
	kCFStringEncodingMacMongolian StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macOriya
	kCFStringEncodingMacOriya StringEncodings = 0
	// kCFStringEncodingMacRomanLatin1 - Mac OS Roman permuted to align with ISO Latin-1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macRomanLatin1
	kCFStringEncodingMacRomanLatin1 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macRomanian
	kCFStringEncodingMacRomanian StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macSinhalese
	kCFStringEncodingMacSinhalese StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macSymbol
	kCFStringEncodingMacSymbol StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macTamil
	kCFStringEncodingMacTamil StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macTelugu
	kCFStringEncodingMacTelugu StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macThai
	kCFStringEncodingMacThai StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macTibetan
	kCFStringEncodingMacTibetan StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macTurkish
	kCFStringEncodingMacTurkish StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macUkrainian
	kCFStringEncodingMacUkrainian StringEncodings = 0
	// kCFStringEncodingMacVT100 - VT100102 font from Comm Toolbox: Latin-1 repertoire + box drawing etc.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macVT100
	kCFStringEncodingMacVT100 StringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macVietnamese
	kCFStringEncodingMacVietnamese StringEncodings = 0
	// kCFStringEncodingNextStepJapanese - NextStep Japanese encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/nextStepJapanese
	kCFStringEncodingNextStepJapanese StringEncodings = 0
	// kCFStringEncodingShiftJIS - Plain Shift-JIS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/shiftJIS
	kCFStringEncodingShiftJIS StringEncodings = 0
	// kCFStringEncodingShiftJIS_X0213 - Shift-JIS format encoding of JIS X0213 planes 1 and 2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/shiftJIS_X0213
	kCFStringEncodingShiftJIS_X0213 StringEncodings = 0
	// kCFStringEncodingShiftJIS_X0213_00 - Shift-JIS format encoding of JIS X0213 planes 1 and 2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/shiftJIS_X0213_00
	kCFStringEncodingShiftJIS_X0213_00 StringEncodings = 0
	// kCFStringEncodingShiftJIS_X0213_MenKuTen - JIS X0213 in plane-row-column notation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/shiftJIS_X0213_MenKuTen
	kCFStringEncodingShiftJIS_X0213_MenKuTen StringEncodings = 0
	// kCFStringEncodingWindowsArabic - Code page 1256.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsArabic
	kCFStringEncodingWindowsArabic StringEncodings = 0
	// kCFStringEncodingWindowsBalticRim - Code page 1257.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsBalticRim
	kCFStringEncodingWindowsBalticRim StringEncodings = 0
	// kCFStringEncodingWindowsCyrillic - Code page 1251, Slavic Cyrillic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsCyrillic
	kCFStringEncodingWindowsCyrillic StringEncodings = 0
	// kCFStringEncodingWindowsGreek - Code page 1253.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsGreek
	kCFStringEncodingWindowsGreek StringEncodings = 0
	// kCFStringEncodingWindowsHebrew - Code page 1255.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsHebrew
	kCFStringEncodingWindowsHebrew StringEncodings = 0
	// kCFStringEncodingWindowsKoreanJohab - Code page 1361, for Windows NT.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsKoreanJohab
	kCFStringEncodingWindowsKoreanJohab StringEncodings = 0
	// kCFStringEncodingWindowsLatin2 - Code page 1250, Central Europe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsLatin2
	kCFStringEncodingWindowsLatin2 StringEncodings = 0
	// kCFStringEncodingWindowsLatin5 - Code page 1254, Turkish.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsLatin5
	kCFStringEncodingWindowsLatin5 StringEncodings = 0
	// kCFStringEncodingWindowsVietnamese - Code page 1258.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsVietnamese
	kCFStringEncodingWindowsVietnamese StringEncodings = 0
)

// StringNormalizationForm - Unicode normalization forms as described in Unicode Technical Report #15.
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

// StringTokenizerTokenType - Token types returned by 
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

// TimeZoneNameStyle - Index type for constants used to specify styles of time zone names.
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

// URLBookmarkCreationOptions - Type for bookmark data creation options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions
type URLBookmarkCreationOptions uint

const (
	// kCFURLBookmarkCreationMinimalBookmarkMask - Specifies that an alias created with the bookmark data be created with minimal information, which may make it smaller but still able to resolve in certain ways.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/minimalBookmarkMask
	kCFURLBookmarkCreationMinimalBookmarkMask URLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationPreferFileIDResolutionMask - Specifies that an alias created with the bookmark data prefers resolving with its embedded file ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/preferFileIDResolutionMask
	kCFURLBookmarkCreationPreferFileIDResolutionMask URLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess - When combined with the   option, specifies that you want to create a security-scoped bookmark that, when resolved, provides a security-scoped URL allowing read-only access to a file-system resource; for use in an app that adopts App Sandbox.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/securityScopeAllowOnlyReadAccess
	kCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess URLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationSuitableForBookmarkFile - Specifies that the bookmark data include properties required to create Finder alias files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/suitableForBookmarkFile
	kCFURLBookmarkCreationSuitableForBookmarkFile URLBookmarkCreationOptions = 0
	// kCFURLBookmarkCreationWithSecurityScope - Specifies that you want to create a security-scoped bookmark that, when resolved, provides a security-scoped URL allowing read/write access to a file-system resource; for use in an app that adopts App Sandbox.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/withSecurityScope
	kCFURLBookmarkCreationWithSecurityScope URLBookmarkCreationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions/withoutImplicitSecurityScope
	kCFURLBookmarkCreationWithoutImplicitSecurityScope URLBookmarkCreationOptions = 0
)

// URLBookmarkResolutionOptions - Type for bookmark data resolution options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions
type URLBookmarkResolutionOptions uint

const (
	// kCFBookmarkResolutionWithoutMountingMask - Specifies that no volume should be mounted during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfBookmarkResolutionWithoutMountingMask
	kCFBookmarkResolutionWithoutMountingMask URLBookmarkResolutionOptions = 0
	// kCFBookmarkResolutionWithoutUIMask - Specifies that no UI feedback accompany resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfBookmarkResolutionWithoutUIMask
	kCFBookmarkResolutionWithoutUIMask URLBookmarkResolutionOptions = 0
	// kCFURLBookmarkResolutionWithSecurityScope - Specifies that the security scope, applied to the bookmark when it was created, should be used during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfurlBookmarkResolutionWithSecurityScope
	kCFURLBookmarkResolutionWithSecurityScope URLBookmarkResolutionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfurlBookmarkResolutionWithoutImplicitStartAccessing
	kCFURLBookmarkResolutionWithoutImplicitStartAccessing URLBookmarkResolutionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfurlBookmarkResolutionWithoutMountingMask
	kCFURLBookmarkResolutionWithoutMountingMask URLBookmarkResolutionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfurlBookmarkResolutionWithoutUIMask
	kCFURLBookmarkResolutionWithoutUIMask URLBookmarkResolutionOptions = 0
)

// URLComponentType - The types of components in a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType
type URLComponentType uint

const (
	// kCFURLComponentFragment - The URL’s fragment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/fragment
	kCFURLComponentFragment URLComponentType = 0
	// kCFURLComponentHost - The URL’s host.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/host
	kCFURLComponentHost URLComponentType = 0
	// kCFURLComponentNetLocation - The URL’s network location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/netLocation
	kCFURLComponentNetLocation URLComponentType = 0
	// kCFURLComponentParameterString - The URL’s parameter string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/parameterString
	kCFURLComponentParameterString URLComponentType = 0
	// kCFURLComponentPassword - The user’s password.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/password
	kCFURLComponentPassword URLComponentType = 0
	// kCFURLComponentPath - The URL’s path component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/path
	kCFURLComponentPath URLComponentType = 0
	// kCFURLComponentPort - The URL’s port.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/port
	kCFURLComponentPort URLComponentType = 0
	// kCFURLComponentQuery - The URL’s query.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/query
	kCFURLComponentQuery URLComponentType = 0
	// kCFURLComponentResourceSpecifier - The URL’s resource specifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/resourceSpecifier
	kCFURLComponentResourceSpecifier URLComponentType = 0
	// kCFURLComponentScheme - The URL’s scheme.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/scheme
	kCFURLComponentScheme URLComponentType = 0
	// kCFURLComponentUser - The URL’s user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/user
	kCFURLComponentUser URLComponentType = 0
	// kCFURLComponentUserInfo - The user’s information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/userInfo
	kCFURLComponentUserInfo URLComponentType = 0
)

// URLEnumeratorOptions - Options for controlling enumerator behavior.
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

// URLEnumeratorResult - Result codes from the 
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

// URLError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError
type URLError uint

const (
	// kCFURLImproperArgumentsError - Indicates one or more arguments are improper.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/improperArgumentsError
	kCFURLImproperArgumentsError URLError = 0
	// kCFURLPropertyKeyUnavailableError - Indicates a property key was unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/propertyKeyUnavailableError
	kCFURLPropertyKeyUnavailableError URLError = 0
	// kCFURLRemoteHostUnavailableError - Indicates a remote host is unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/remoteHostUnavailableError
	kCFURLRemoteHostUnavailableError URLError = 0
	// kCFURLResourceAccessViolationError - Indicates an error in accessing a resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/resourceAccessViolationError
	kCFURLResourceAccessViolationError URLError = 0
	// kCFURLResourceNotFoundError - Indicates a resource was not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/resourceNotFoundError
	kCFURLResourceNotFoundError URLError = 0
	// kCFURLTimeoutError - Indicates a timeout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/timeoutError
	kCFURLTimeoutError URLError = 0
	// kCFURLUnknownError - Indicates an unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/unknownError
	kCFURLUnknownError URLError = 0
	// kCFURLUnknownPropertyKeyError - Indicates a property key is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/unknownPropertyKeyError
	kCFURLUnknownPropertyKeyError URLError = 0
	// kCFURLUnknownSchemeError - Indicates that the scheme is not recognized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/unknownSchemeError
	kCFURLUnknownSchemeError URLError = 0
)

// URLPathStyle - Options you can use to determine how CFURL functions parse a file system path name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle
type URLPathStyle uint

const (
	// kCFURLWindowsPathStyle - Indicates a Windows style path name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle/cfurlWindowsPathStyle
	kCFURLWindowsPathStyle URLPathStyle = 0
	// kCFURLHFSPathStyle - Indicates a HFS style path name. Components are colon delimited. A leading colon indicates a relative path, otherwise the first path component denotes the volume.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle/cfurlhfsPathStyle
	kCFURLHFSPathStyle URLPathStyle = 0
	// kCFURLPOSIXPathStyle - Indicates a POSIX style path name. Components are slash delimited. A leading slash indicates an absolute path; a trailing slash is not significant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle/cfurlposixPathStyle
	kCFURLPOSIXPathStyle URLPathStyle = 0
)

// XMLEntityTypeCode - The entity type identification codes that the parser uses to describe XML entities.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode
type XMLEntityTypeCode uint

const (
	// kCFXMLEntityTypeCharacter - Indicates a character entity type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/character
	kCFXMLEntityTypeCharacter XMLEntityTypeCode = 0
	// kCFXMLEntityTypeParameter - Implies a parsed, internal entity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/parameter
	kCFXMLEntityTypeParameter XMLEntityTypeCode = 0
	// kCFXMLEntityTypeParsedExternal - Indicates a parsed, external entity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/parsedExternal
	kCFXMLEntityTypeParsedExternal XMLEntityTypeCode = 0
	// kCFXMLEntityTypeParsedInternal - Indicates a parsed, internal entity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/parsedInternal
	kCFXMLEntityTypeParsedInternal XMLEntityTypeCode = 0
	// kCFXMLEntityTypeUnparsed - Indicates an unparsed entity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/unparsed
	kCFXMLEntityTypeUnparsed XMLEntityTypeCode = 0
)

// XMLNodeTypeCode - The various XML data type identification codes that the parser uses to describe XML structures.
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

// XMLParserOptions - Options you can use to control the parser’s treatment of an XML document.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions
type XMLParserOptions uint

const (
	// kCFXMLParserAddImpliedAttributes - Where the DTD specifies implied attribute-value pairs for a particular element, add those pairs to any occurrences of the element in the element tree. Currently not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/addImpliedAttributes
	kCFXMLParserAddImpliedAttributes XMLParserOptions = 0
	// kCFXMLParserAllOptions - Makes the parser do the most work, returning only the pure elementtree.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/allOptions
	kCFXMLParserAllOptions XMLParserOptions = 0
	// kCFXMLParserNoOptions - Leaves the XML as “intact” as possible (reports all structures; performs no replacements).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/kCFXMLParserNoOptions
	kCFXMLParserNoOptions XMLParserOptions = 0
	// kCFXMLParserReplacePhysicalEntities - Replaces declared entities like  ;. Note that other than the 5 predefined entities ( ,  ,  ,  ,  ), these must be defined in the DTD. Currently not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/replacePhysicalEntities
	kCFXMLParserReplacePhysicalEntities XMLParserOptions = 0
	// kCFXMLParserResolveExternalEntities - Resolves all external entities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/resolveExternalEntities
	kCFXMLParserResolveExternalEntities XMLParserOptions = 0
	// kCFXMLParserSkipMetaData - Silently skip over metadata constructs (the DTD and comments).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/skipMetaData
	kCFXMLParserSkipMetaData XMLParserOptions = 0
	// kCFXMLParserSkipWhitespace - Skip over all whitespace that does not abut non-whitespace character data. In other words, given “ ,” the whitespace between foo’s open tag and bar’s open tag would be suppressed, but the whitespace around   would be preserved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/skipWhitespace
	kCFXMLParserSkipWhitespace XMLParserOptions = 0
	// kCFXMLParserValidateDocument - Validates the document against its grammar from the DTD, reporting any errors. Currently not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/validateDocument
	kCFXMLParserValidateDocument XMLParserOptions = 0
)

// XMLParserStatusCode - The various status and error flags that can be returned by the parser.
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
type CGRectEdge uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge/maxXEdge
	CGRectMaxXEdge CGRectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge/maxYEdge
	CGRectMaxYEdge CGRectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge/minXEdge
	CGRectMinXEdge CGRectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge/minYEdge
	CGRectMinYEdge CGRectEdge = 0
)


