// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

// Enum types and constants
// CFCalendarUnit - CFCalendarUnit constants are used to specify calendrical units, such as day or month, in various calendar calculations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit
type CFCalendarUnit uint

const (
	// kCFCalendarUnitDay - Specifies the day unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/day
	kCFCalendarUnitDay CFCalendarUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit/dayOfYear
	kCFCalendarUnitDayOfYear CFCalendarUnit = 0
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

const (
	// kCFCharacterSetAlphaNumeric - Alpha Numeric character set (Unicode General Category L*, M*, & N*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/alphaNumeric
	kCFCharacterSetAlphaNumeric CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetCapitalizedLetter - Titlecase character set (Unicode General Category Lt).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/capitalizedLetter
	kCFCharacterSetCapitalizedLetter CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetControl - Control character set (Unicode General Category Cc and Cf).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/control
	kCFCharacterSetControl CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetDecimalDigit - Decimal digit character set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/decimalDigit
	kCFCharacterSetDecimalDigit CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetDecomposable - Canonically decomposable character set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/decomposable
	kCFCharacterSetDecomposable CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetIllegal - Illegal character set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/illegal
	kCFCharacterSetIllegal CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetLetter - Letter character set (Unicode General Category L* & M*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/letter
	kCFCharacterSetLetter CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetLowercaseLetter - Lowercase character set (Unicode General Category Ll).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/lowercaseLetter
	kCFCharacterSetLowercaseLetter CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetNewline - Newline character set ( ,  ,  , and  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/newline
	kCFCharacterSetNewline CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetNonBase - Non-base character set (Unicode General Category M*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/nonBase
	kCFCharacterSetNonBase CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetPunctuation - Punctuation character set (Unicode General Category P*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/punctuation
	kCFCharacterSetPunctuation CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetSymbol - Symbol character set (Unicode General Category S*).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/symbol
	kCFCharacterSetSymbol CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetUppercaseLetter - Uppercase character set (Unicode General Category Lu and Lt).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/uppercaseLetter
	kCFCharacterSetUppercaseLetter CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetWhitespace - Whitespace character set (Unicode General Category Zs and U0009 CHARACTER TABULATION).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/whitespace
	kCFCharacterSetWhitespace CFCharacterSetPredefinedSet = 0
	// kCFCharacterSetWhitespaceAndNewline - Whitespace and Newline character set (Unicode General Category Z*,  , and  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet/whitespaceAndNewline
	kCFCharacterSetWhitespaceAndNewline CFCharacterSetPredefinedSet = 0
)

// CFComparisonResult - Constants returned by comparison functions, indicating whether a value is equal to, less than, or greater than another value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult
type CFComparisonResult uint

const (
	// kCFCompareEqualTo - Returned by a comparison function if the first value is equal to the second value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult/compareEqualTo
	kCFCompareEqualTo CFComparisonResult = 0
	// kCFCompareGreaterThan - Returned by a comparison function if the first value is greater than the second value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult/compareGreaterThan
	kCFCompareGreaterThan CFComparisonResult = 0
	// kCFCompareLessThan - Returned by a comparison function if the first value is less than the second value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult/compareLessThan
	kCFCompareLessThan CFComparisonResult = 0
)

// CFDataSearchFlags - A 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags
type CFDataSearchFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags/anchored
	kCFDataSearchAnchored CFDataSearchFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags/backwards
	kCFDataSearchBackwards CFDataSearchFlags = 0
)

// CFDateFormatterStyle - Data type for predefined date and time format styles.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle
type CFDateFormatterStyle uint

const (
	// kCFDateFormatterFullStyle - Specifies a full style with complete details, such as “Tuesday, April 12, 1952 AD” or “3:30:42pm PST”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/fullStyle
	kCFDateFormatterFullStyle CFDateFormatterStyle = 0
	// kCFDateFormatterLongStyle - Specifies a long style, typically with full text, such as “November 23, 1937” or “3:30:32pm”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/longStyle
	kCFDateFormatterLongStyle CFDateFormatterStyle = 0
	// kCFDateFormatterMediumStyle - Specifies a medium style, typically with abbreviated text, such as “Nov 23, 1937”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/mediumStyle
	kCFDateFormatterMediumStyle CFDateFormatterStyle = 0
	// kCFDateFormatterNoStyle - Specifies no output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/noStyle
	kCFDateFormatterNoStyle CFDateFormatterStyle = 0
	// kCFDateFormatterShortStyle - Specifies a short style, typically numeric only, such as “11/23/37” or “3:30pm”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle/shortStyle
	kCFDateFormatterShortStyle CFDateFormatterStyle = 0
)

// CFFileSecurityClearOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions
type CFFileSecurityClearOptions uint

const (
	// kCFFileSecurityClearAccessControlList - Clear the access control list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/accessControlList
	kCFFileSecurityClearAccessControlList CFFileSecurityClearOptions = 0
	// kCFFileSecurityClearGroup - Clear the (POSIX) group ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/group
	kCFFileSecurityClearGroup CFFileSecurityClearOptions = 0
	// kCFFileSecurityClearGroupUUID - Clear the group UUID (for the access control list).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/groupUUID
	kCFFileSecurityClearGroupUUID CFFileSecurityClearOptions = 0
	// kCFFileSecurityClearMode - Clear the file’s mode (POSIX permissions).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/mode
	kCFFileSecurityClearMode CFFileSecurityClearOptions = 0
	// kCFFileSecurityClearOwner - Clear the (POSIX) owner ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/owner
	kCFFileSecurityClearOwner CFFileSecurityClearOptions = 0
	// kCFFileSecurityClearOwnerUUID - Clear the owner UUID (for the access control list).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions/ownerUUID
	kCFFileSecurityClearOwnerUUID CFFileSecurityClearOptions = 0
)

// CFGregorianUnitFlags - These option flags are used as a mask to indicate a specific set of fields in the CFGregorianDate or CFGregorianUnits structures.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags
type CFGregorianUnitFlags uint

const (
	// kCFGregorianAllUnits - Specifies all fields.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/allUnits
	kCFGregorianAllUnits CFGregorianUnitFlags = 0
	// kCFGregorianUnitsDays - Specifies the day field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsDays
	kCFGregorianUnitsDays CFGregorianUnitFlags = 0
	// kCFGregorianUnitsHours - Specifies the hours field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsHours
	kCFGregorianUnitsHours CFGregorianUnitFlags = 0
	// kCFGregorianUnitsMinutes - Specifies the minutes field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsMinutes
	kCFGregorianUnitsMinutes CFGregorianUnitFlags = 0
	// kCFGregorianUnitsMonths - Specifies the month field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsMonths
	kCFGregorianUnitsMonths CFGregorianUnitFlags = 0
	// kCFGregorianUnitsSeconds - Specifies the seconds field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsSeconds
	kCFGregorianUnitsSeconds CFGregorianUnitFlags = 0
	// kCFGregorianUnitsYears - Specifies the year field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags/unitsYears
	kCFGregorianUnitsYears CFGregorianUnitFlags = 0
)

// CFISO8601DateFormatOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions
type CFISO8601DateFormatOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withColonSeparatorInTime
	kCFISO8601DateFormatWithColonSeparatorInTime CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withColonSeparatorInTimeZone
	kCFISO8601DateFormatWithColonSeparatorInTimeZone CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withDashSeparatorInDate
	kCFISO8601DateFormatWithDashSeparatorInDate CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withDay
	kCFISO8601DateFormatWithDay CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withFractionalSeconds
	kCFISO8601DateFormatWithFractionalSeconds CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withFullDate
	kCFISO8601DateFormatWithFullDate CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withFullTime
	kCFISO8601DateFormatWithFullTime CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withInternetDateTime
	kCFISO8601DateFormatWithInternetDateTime CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withMonth
	kCFISO8601DateFormatWithMonth CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withSpaceBetweenDateAndTime
	kCFISO8601DateFormatWithSpaceBetweenDateAndTime CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withTime
	kCFISO8601DateFormatWithTime CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withTimeZone
	kCFISO8601DateFormatWithTimeZone CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withWeekOfYear
	kCFISO8601DateFormatWithWeekOfYear CFISO8601DateFormatOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions/withYear
	kCFISO8601DateFormatWithYear CFISO8601DateFormatOptions = 0
)

// CFLocaleLanguageDirection - These constants describe the text direction for a language. They are returned by the functions 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection
type CFLocaleLanguageDirection uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/bottomToTop
	kCFLocaleLanguageDirectionBottomToTop CFLocaleLanguageDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/leftToRight
	kCFLocaleLanguageDirectionLeftToRight CFLocaleLanguageDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/rightToLeft
	kCFLocaleLanguageDirectionRightToLeft CFLocaleLanguageDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/topToBottom
	kCFLocaleLanguageDirectionTopToBottom CFLocaleLanguageDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection/unknown
	kCFLocaleLanguageDirectionUnknown CFLocaleLanguageDirection = 0
)

// CFNotificationSuspensionBehavior - Suspension flags that indicate how distributed notifications should be handled when the receiving application is in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior
type CFNotificationSuspensionBehavior uint

const (
	// CFNotificationSuspensionBehaviorCoalesce - The server will only queue the last notification of the specified name and object; earlier notifications are dropped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior/coalesce
	CFNotificationSuspensionBehaviorCoalesce CFNotificationSuspensionBehavior = 0
	// CFNotificationSuspensionBehaviorDeliverImmediately - The server will deliver notifications of the specified name and object whether or not the application is in the background. When a notification with this suspension behavior is matched, it has the effect of first flushing any queued notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior/deliverImmediately
	CFNotificationSuspensionBehaviorDeliverImmediately CFNotificationSuspensionBehavior = 0
	// CFNotificationSuspensionBehaviorDrop - The server will not queue any notifications of the specified name and object while the receiving application is in the background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior/drop
	CFNotificationSuspensionBehaviorDrop CFNotificationSuspensionBehavior = 0
	// CFNotificationSuspensionBehaviorHold - The server will hold all matching notifications until the queue has been filled (queue size determined by the server) at which point the server may flush queued notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior/hold
	CFNotificationSuspensionBehaviorHold CFNotificationSuspensionBehavior = 0
)

// CFNumberFormatterOptionFlags - Type for constants specifying how numbers should be parsed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterOptionFlags
type CFNumberFormatterOptionFlags uint

const (
	// kCFNumberFormatterParseIntegersOnly - Specifies that only integers should be parsed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterOptionFlags/parseIntegersOnly
	kCFNumberFormatterParseIntegersOnly CFNumberFormatterOptionFlags = 0
)

// CFNumberFormatterPadPosition - Type for constants specifying how numbers should be padded.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition
type CFNumberFormatterPadPosition uint

const (
	// kCFNumberFormatterPadAfterPrefix - Specifies the number of padding characters after the prefix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition/afterPrefix
	kCFNumberFormatterPadAfterPrefix CFNumberFormatterPadPosition = 0
	// kCFNumberFormatterPadAfterSuffix - Specifies the number of padding characters after the suffix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition/afterSuffix
	kCFNumberFormatterPadAfterSuffix CFNumberFormatterPadPosition = 0
	// kCFNumberFormatterPadBeforePrefix - Specifies the number of padding characters before the prefix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition/beforePrefix
	kCFNumberFormatterPadBeforePrefix CFNumberFormatterPadPosition = 0
	// kCFNumberFormatterPadBeforeSuffix - Specifies the number of padding characters before the suffix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition/beforeSuffix
	kCFNumberFormatterPadBeforeSuffix CFNumberFormatterPadPosition = 0
)

// CFNumberFormatterRoundingMode - These constants are used to specify how numbers should be rounded.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode
type CFNumberFormatterRoundingMode uint

const (
	// kCFNumberFormatterRoundCeiling - Round towards positive infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundCeiling
	kCFNumberFormatterRoundCeiling CFNumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundDown - Round towards zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundDown
	kCFNumberFormatterRoundDown CFNumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundFloor - Round towards negative infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundFloor
	kCFNumberFormatterRoundFloor CFNumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundHalfDown - Round towards the nearest integer, or towards zero if equidistant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundHalfDown
	kCFNumberFormatterRoundHalfDown CFNumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundHalfEven - Round towards the nearest integer, or towards an even number if equidistant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundHalfEven
	kCFNumberFormatterRoundHalfEven CFNumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundHalfUp - Round towards the nearest integer, or away from zero if equidistant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundHalfUp
	kCFNumberFormatterRoundHalfUp CFNumberFormatterRoundingMode = 0
	// kCFNumberFormatterRoundUp - Round away from zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode/roundUp
	kCFNumberFormatterRoundUp CFNumberFormatterRoundingMode = 0
)

// CFNumberFormatterStyle - Type for constants specifying a formatter style.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle
type CFNumberFormatterStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/currencyAccountingStyle
	kCFNumberFormatterCurrencyAccountingStyle CFNumberFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/currencyISOCodeStyle
	kCFNumberFormatterCurrencyISOCodeStyle CFNumberFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/currencyPluralStyle
	kCFNumberFormatterCurrencyPluralStyle CFNumberFormatterStyle = 0
	// kCFNumberFormatterCurrencyStyle - Specifies a currency style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/currencyStyle
	kCFNumberFormatterCurrencyStyle CFNumberFormatterStyle = 0
	// kCFNumberFormatterDecimalStyle - Specifies a decimal style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/decimalStyle
	kCFNumberFormatterDecimalStyle CFNumberFormatterStyle = 0
	// kCFNumberFormatterNoStyle - Specifies no style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/noStyle
	kCFNumberFormatterNoStyle CFNumberFormatterStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/ordinalStyle
	kCFNumberFormatterOrdinalStyle CFNumberFormatterStyle = 0
	// kCFNumberFormatterPercentStyle - Specifies a percent style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/percentStyle
	kCFNumberFormatterPercentStyle CFNumberFormatterStyle = 0
	// kCFNumberFormatterScientificStyle - Specifies a scientific style format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/scientificStyle
	kCFNumberFormatterScientificStyle CFNumberFormatterStyle = 0
	// kCFNumberFormatterSpellOutStyle - Specifies a spelled out format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle/spellOutStyle
	kCFNumberFormatterSpellOutStyle CFNumberFormatterStyle = 0
)

// CFNumberType - Flags used by CFNumber to indicate the data type of a value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType
type CFNumberType uint

const (
	// kCFNumberCFIndexType - CFIndex value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/cfIndexType
	kCFNumberCFIndexType CFNumberType = 0
	// kCFNumberCGFloatType -  value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/cgFloatType
	kCFNumberCGFloatType CFNumberType = 0
	// kCFNumberCharType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/charType
	kCFNumberCharType CFNumberType = 0
	// kCFNumberDoubleType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/doubleType
	kCFNumberDoubleType CFNumberType = 0
	// kCFNumberFloat32Type - Thirty-two-bit real. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/float32Type
	kCFNumberFloat32Type CFNumberType = 0
	// kCFNumberFloat64Type - Sixty-four-bit real. The   data type is defined in   and conforms to the 64-bit IEEE 754 standard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/float64Type
	kCFNumberFloat64Type CFNumberType = 0
	// kCFNumberFloatType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/floatType
	kCFNumberFloatType CFNumberType = 0
	// kCFNumberIntType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/intType
	kCFNumberIntType CFNumberType = 0
	// kCFNumberLongLongType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/longLongType
	kCFNumberLongLongType CFNumberType = 0
	// kCFNumberLongType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/longType
	kCFNumberLongType CFNumberType = 0
	// kCFNumberMaxType - Same as  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/maxType
	kCFNumberMaxType CFNumberType = 0
	// kCFNumberNSIntegerType -  value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/nsIntegerType
	kCFNumberNSIntegerType CFNumberType = 0
	// kCFNumberSInt16Type - Sixteen-bit, signed integer. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/sInt16Type
	kCFNumberSInt16Type CFNumberType = 0
	// kCFNumberSInt32Type - Thirty-two-bit, signed integer. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/sInt32Type
	kCFNumberSInt32Type CFNumberType = 0
	// kCFNumberSInt64Type - Sixty-four-bit, signed integer. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/sInt64Type
	kCFNumberSInt64Type CFNumberType = 0
	// kCFNumberSInt8Type - Eight-bit, signed integer. The   data type is defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/sInt8Type
	kCFNumberSInt8Type CFNumberType = 0
	// kCFNumberShortType - Basic C   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberType/shortType
	kCFNumberShortType CFNumberType = 0
)

// CFPropertyListFormat - Specifies the format of a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListFormat
type CFPropertyListFormat uint

const (
	// kCFPropertyListBinaryFormat_v1_0 - Binary format version 1.0.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListFormat/binaryFormat_v1_0
	kCFPropertyListBinaryFormat_v1_0 CFPropertyListFormat = 0
	// kCFPropertyListOpenStepFormat - OpenStep format (use of this format is discouraged).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListFormat/openStepFormat
	kCFPropertyListOpenStepFormat CFPropertyListFormat = 0
	// kCFPropertyListXMLFormat_v1_0 - XML format version 1.0.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListFormat/xmlFormat_v1_0
	kCFPropertyListXMLFormat_v1_0 CFPropertyListFormat = 0
)

// CFPropertyListMutabilityOptions - Type for flags that determine the degree of mutability of newly created property lists.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListMutabilityOptions
type CFPropertyListMutabilityOptions uint

const (
	// kCFPropertyListImmutable - Specifies that the property list should be immutable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListMutabilityOptions/kCFPropertyListImmutable
	kCFPropertyListImmutable CFPropertyListMutabilityOptions = 0
	// kCFPropertyListMutableContainers - Specifies that the property list should have mutable containers but immutable leaves.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListMutabilityOptions/mutableContainers
	kCFPropertyListMutableContainers CFPropertyListMutabilityOptions = 0
	// kCFPropertyListMutableContainersAndLeaves - Specifies that the property list should have mutable containers and mutable leaves.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListMutabilityOptions/mutableContainersAndLeaves
	kCFPropertyListMutableContainersAndLeaves CFPropertyListMutabilityOptions = 0
)

// CFRunLoopActivity - Run loop activity stages in which run loop observers can be scheduled.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity
type CFRunLoopActivity uint

const (
	// kCFRunLoopAfterWaiting - Inside the event processing loop after the run loop wakes up, but before processing the event that woke it up. This activity occurs only if the run loop did in fact go to sleep during the current loop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/afterWaiting
	kCFRunLoopAfterWaiting CFRunLoopActivity = 0
	// kCFRunLoopAllActivities - A combination of all the preceding stages.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/allActivities
	kCFRunLoopAllActivities CFRunLoopActivity = 0
	// kCFRunLoopBeforeSources - Inside the event processing loop before any sources are processed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/beforeSources
	kCFRunLoopBeforeSources CFRunLoopActivity = 0
	// kCFRunLoopBeforeTimers - Inside the event processing loop before any timers are processed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/beforeTimers
	kCFRunLoopBeforeTimers CFRunLoopActivity = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/beforeWaiting
	kCFRunLoopBeforeWaiting CFRunLoopActivity = 0
	// kCFRunLoopEntry - The entrance of the run loop, before entering the event processing loop. This activity occurs once for each call to   and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/entry
	kCFRunLoopEntry CFRunLoopActivity = 0
	// kCFRunLoopExit - The exit of the run loop, after exiting the event processing loop. This activity occurs once for each call to   and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity/exit
	kCFRunLoopExit CFRunLoopActivity = 0
)

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

// CFSocketCallBackType - Types of socket activity that can cause the callback function of a CFSocket object to be called.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType
type CFSocketCallBackType uint

const (
	// kCFSocketAcceptCallBack - New connections will be automatically accepted and the callback is called with the data argument being a pointer to a   of the child socket. This callback is usable only with listening sockets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/acceptCallBack
	kCFSocketAcceptCallBack CFSocketCallBackType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/connectCallBack
	kCFSocketConnectCallBack CFSocketCallBackType = 0
	// kCFSocketDataCallBack - Incoming data will be read in chunks in the background and the callback is called with the data argument being a CFData object containing the read data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/dataCallBack
	kCFSocketDataCallBack CFSocketCallBackType = 0
	// kCFSocketNoCallBack - No callback should be made for any activity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/kCFSocketNoCallBack
	kCFSocketNoCallBack CFSocketCallBackType = 0
	// kCFSocketReadCallBack - The callback is called when data is available to be read or a new connection is waiting to be accepted. The data is not automatically read; the callback must read the data itself.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/readCallBack
	kCFSocketReadCallBack CFSocketCallBackType = 0
	// kCFSocketWriteCallBack - The callback is called when the socket is writable. This callback type may be useful when large amounts of data are being sent rapidly over the socket and you want a notification when there is space in the kernel buffers for more data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType/writeCallBack
	kCFSocketWriteCallBack CFSocketCallBackType = 0
)

// CFSocketError - Error codes for many CFSocket functions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketError
type CFSocketError uint

const (
	// kCFSocketError - The socket operation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketError/error
	kCFSocketError CFSocketError = 0
	// kCFSocketSuccess - The socket operation succeeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketError/success
	kCFSocketSuccess CFSocketError = 0
	// kCFSocketTimeout - The socket operation timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketError/timeout
	kCFSocketTimeout CFSocketError = 0
)

// CFStreamErrorDomain - Defines constants for values returned in the domain field of the 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamErrorDomain
type CFStreamErrorDomain uint

const (
	// kCFStreamErrorDomainPOSIX - The error code is an error code defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamErrorDomain/POSIX
	kCFStreamErrorDomainPOSIX CFStreamErrorDomain = 0
	// kCFStreamErrorDomainCustom - The error code is a custom error code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamErrorDomain/custom
	kCFStreamErrorDomainCustom CFStreamErrorDomain = 0
	// kCFStreamErrorDomainMacOSStatus - The error is an OSStatus value defined in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamErrorDomain/macOSStatus
	kCFStreamErrorDomainMacOSStatus CFStreamErrorDomain = 0
)

// CFStreamEventType - Defines constants for stream-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType
type CFStreamEventType uint

const (
	// kCFStreamEventCanAcceptBytes - The stream can accept bytes for writing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/canAcceptBytes
	kCFStreamEventCanAcceptBytes CFStreamEventType = 0
	// kCFStreamEventEndEncountered - The end of the stream has been reached.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/endEncountered
	kCFStreamEventEndEncountered CFStreamEventType = 0
	// kCFStreamEventErrorOccurred - An error has occurred on the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/errorOccurred
	kCFStreamEventErrorOccurred CFStreamEventType = 0
	// kCFStreamEventHasBytesAvailable - The stream has bytes to be read.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/hasBytesAvailable
	kCFStreamEventHasBytesAvailable CFStreamEventType = 0
	// kCFStreamEventNone - No event has occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/kCFStreamEventNone
	kCFStreamEventNone CFStreamEventType = 0
	// kCFStreamEventOpenCompleted - The open has completed successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType/openCompleted
	kCFStreamEventOpenCompleted CFStreamEventType = 0
)

// CFStreamStatus - Constants that describe the status of a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus
type CFStreamStatus uint

const (
	// kCFStreamStatusAtEnd - There is no more data to read, or no more data can be written.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/atEnd
	kCFStreamStatusAtEnd CFStreamStatus = 0
	// kCFStreamStatusClosed - The stream is closed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/closed
	kCFStreamStatusClosed CFStreamStatus = 0
	// kCFStreamStatusError - An error occurred on the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/error
	kCFStreamStatusError CFStreamStatus = 0
	// kCFStreamStatusNotOpen - The stream is not open for reading or writing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/notOpen
	kCFStreamStatusNotOpen CFStreamStatus = 0
	// kCFStreamStatusOpen - The stream is open.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/open
	kCFStreamStatusOpen CFStreamStatus = 0
	// kCFStreamStatusOpening - The stream is being opened for reading or for writing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/opening
	kCFStreamStatusOpening CFStreamStatus = 0
	// kCFStreamStatusReading - The stream is being read from.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/reading
	kCFStreamStatusReading CFStreamStatus = 0
	// kCFStreamStatusWriting - The stream is being written to.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus/writing
	kCFStreamStatusWriting CFStreamStatus = 0
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
	// kCFStringEncodingUTF16 - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF16Format encoding (alias of kCFStringEncodingUnicode).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF16
	kCFStringEncodingUTF16 CFStringBuiltInEncodings = 0
	// kCFStringEncodingUTF16BE - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF16BEFormat encoding. This constant specifies big-endian byte order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF16BE
	kCFStringEncodingUTF16BE CFStringBuiltInEncodings = 0
	// kCFStringEncodingUTF16LE - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF16LEFormat encoding. This constant specifies little-endian byte order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF16LE
	kCFStringEncodingUTF16LE CFStringBuiltInEncodings = 0
	// kCFStringEncodingUTF32 - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF32Format encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF32
	kCFStringEncodingUTF32 CFStringBuiltInEncodings = 0
	// kCFStringEncodingUTF32BE - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF32BEFormat encoding. This constant specifies big-endian byte order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF32BE
	kCFStringEncodingUTF32BE CFStringBuiltInEncodings = 0
	// kCFStringEncodingUTF32LE - An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF32LEFormat encoding. This constant specifies little-endian byte order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/UTF32LE
	kCFStringEncodingUTF32LE CFStringBuiltInEncodings = 0
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
	// kCFStringEncodingNonLossyASCII - An encoding constant that identifies non-lossy ASCII encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/nonLossyASCII
	kCFStringEncodingNonLossyASCII CFStringBuiltInEncodings = 0
	// kCFStringEncodingUnicode - An encoding constant that identifies the Unicode encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/unicode
	kCFStringEncodingUnicode CFStringBuiltInEncodings = 0
	// kCFStringEncodingWindowsLatin1 - An encoding constant that identifies the Windows Latin 1 encoding (ANSI codepage 1252).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings/windowsLatin1
	kCFStringEncodingWindowsLatin1 CFStringBuiltInEncodings = 0
)

// CFStringCompareFlags - A 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags
type CFStringCompareFlags uint

const (
	// kCFCompareAnchored - Performs searching only on characters at the beginning or end of the range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareAnchored
	kCFCompareAnchored CFStringCompareFlags = 0
	// kCFCompareBackwards - Specifies that the comparison should start at the last elements of the entities being compared (for example, strings or arrays).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareBackwards
	kCFCompareBackwards CFStringCompareFlags = 0
	// kCFCompareCaseInsensitive - Specifies that the comparison should ignore differences in case between alphabetical characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareCaseInsensitive
	kCFCompareCaseInsensitive CFStringCompareFlags = 0
	// kCFCompareDiacriticInsensitive - Specifies that the comparison should ignore diacritic markers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareDiacriticInsensitive
	kCFCompareDiacriticInsensitive CFStringCompareFlags = 0
	// kCFCompareForcedOrdering - Specifies that the comparison is forced to return either   or   if the strings are equivalent but not strictly equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareForcedOrdering
	kCFCompareForcedOrdering CFStringCompareFlags = 0
	// kCFCompareLocalized - Specifies that the comparison should take into account differences related to locale, such as the thousands separator character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareLocalized
	kCFCompareLocalized CFStringCompareFlags = 0
	// kCFCompareNonliteral - Specifies that loose equivalence is acceptable, especially as pertains to diacritical marks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareNonliteral
	kCFCompareNonliteral CFStringCompareFlags = 0
	// kCFCompareNumerically - Specifies that represented numeric values should be used as the basis for comparison and not the actual character values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareNumerically
	kCFCompareNumerically CFStringCompareFlags = 0
	// kCFCompareWidthInsensitive - Specifies that the comparison should ignore width differences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags/compareWidthInsensitive
	kCFCompareWidthInsensitive CFStringCompareFlags = 0
)

// CFStringEncodings - Index type for constants used to specify external string encodings.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings
type CFStringEncodings uint

const (
	// kCFStringEncodingANSEL - ANSEL (ANSI Z39.47).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ANSEL
	kCFStringEncodingANSEL CFStringEncodings = 0
	// kCFStringEncodingCNS_11643_92_P1 - CNS 11643-1992 plane 1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/CNS_11643_92_P1
	kCFStringEncodingCNS_11643_92_P1 CFStringEncodings = 0
	// kCFStringEncodingCNS_11643_92_P2 - CNS 11643-1992 plane 2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/CNS_11643_92_P2
	kCFStringEncodingCNS_11643_92_P2 CFStringEncodings = 0
	// kCFStringEncodingCNS_11643_92_P3 - CNS 11643-1992 plane 3 (was plane 14 in 1986 version).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/CNS_11643_92_P3
	kCFStringEncodingCNS_11643_92_P3 CFStringEncodings = 0
	// kCFStringEncodingEBCDIC_CP037 - code page 037, extended EBCDIC (Latin-1 set) for US, Canada.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EBCDIC_CP037
	kCFStringEncodingEBCDIC_CP037 CFStringEncodings = 0
	// kCFStringEncodingEBCDIC_US - basic EBCDIC-US
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EBCDIC_US
	kCFStringEncodingEBCDIC_US CFStringEncodings = 0
	// kCFStringEncodingEUC_CN - ISO 646, GB 2312-80.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EUC_CN
	kCFStringEncodingEUC_CN CFStringEncodings = 0
	// kCFStringEncodingEUC_JP - ISO 646, 1-byte katakana, JIS 208, JIS 212.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EUC_JP
	kCFStringEncodingEUC_JP CFStringEncodings = 0
	// kCFStringEncodingEUC_KR - ISO 646, KS C 5601-1987.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EUC_KR
	kCFStringEncodingEUC_KR CFStringEncodings = 0
	// kCFStringEncodingEUC_TW - ISO 646, CNS 11643-1992 Planes 1-16.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/EUC_TW
	kCFStringEncodingEUC_TW CFStringEncodings = 0
	// kCFStringEncodingGBK_95 - Annex to GB 13000-93; for Windows 95.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/GBK_95
	kCFStringEncodingGBK_95 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/GB_18030_2000
	kCFStringEncodingGB_18030_2000 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/GB_2312_80
	kCFStringEncodingGB_2312_80 CFStringEncodings = 0
	// kCFStringEncodingHZ_GB_2312 - HZ (RFC 1842, for Chinese mail & news).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/HZ_GB_2312
	kCFStringEncodingHZ_GB_2312 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_CN
	kCFStringEncodingISO_2022_CN CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_CN_EXT
	kCFStringEncodingISO_2022_CN_EXT CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_JP
	kCFStringEncodingISO_2022_JP CFStringEncodings = 0
	// kCFStringEncodingISO_2022_JP_1 - RFC 2237.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_JP_1
	kCFStringEncodingISO_2022_JP_1 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_JP_2
	kCFStringEncodingISO_2022_JP_2 CFStringEncodings = 0
	// kCFStringEncodingISO_2022_JP_3 - JIS X0213.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_JP_3
	kCFStringEncodingISO_2022_JP_3 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ISO_2022_KR
	kCFStringEncodingISO_2022_KR CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_C6226_78
	kCFStringEncodingJIS_C6226_78 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_X0201_76
	kCFStringEncodingJIS_X0201_76 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_X0208_83
	kCFStringEncodingJIS_X0208_83 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_X0208_90
	kCFStringEncodingJIS_X0208_90 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/JIS_X0212_90
	kCFStringEncodingJIS_X0212_90 CFStringEncodings = 0
	// kCFStringEncodingKOI8_R - Russian internet standard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/KOI8_R
	kCFStringEncodingKOI8_R CFStringEncodings = 0
	// kCFStringEncodingKOI8_U - RFC 2319, Ukrainian.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/KOI8_U
	kCFStringEncodingKOI8_U CFStringEncodings = 0
	// kCFStringEncodingKSC_5601_87 - Same as KSC 5601-92 without Johab annex.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/KSC_5601_87
	kCFStringEncodingKSC_5601_87 CFStringEncodings = 0
	// kCFStringEncodingUTF7 - kTextEncodingUnicodeDefault + kUnicodeUTF7Format RFC2152.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/UTF7
	kCFStringEncodingUTF7 CFStringEncodings = 0
	// kCFStringEncodingUTF7_IMAP - UTF-7 (IMAP folder variant) RFC3501.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/UTF7_IMAP
	kCFStringEncodingUTF7_IMAP CFStringEncodings = 0
	// kCFStringEncodingVISCII - RFC 1456, Vietnamese.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/VISCII
	kCFStringEncodingVISCII CFStringEncodings = 0
	// kCFStringEncodingBig5 - Big-5 (has variants)
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/big5
	kCFStringEncodingBig5 CFStringEncodings = 0
	// kCFStringEncodingBig5_E - Taiwan Big-5E standard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/big5_E
	kCFStringEncodingBig5_E CFStringEncodings = 0
	// kCFStringEncodingBig5_HKSCS_1999 - Big-5 with Hong Kong special char set supplement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/big5_HKSCS_1999
	kCFStringEncodingBig5_HKSCS_1999 CFStringEncodings = 0
	// kCFStringEncodingDOSArabic - Code page 864.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosArabic
	kCFStringEncodingDOSArabic CFStringEncodings = 0
	// kCFStringEncodingDOSBalticRim - Code page 775.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosBalticRim
	kCFStringEncodingDOSBalticRim CFStringEncodings = 0
	// kCFStringEncodingDOSCanadianFrench - Code page 863.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosCanadianFrench
	kCFStringEncodingDOSCanadianFrench CFStringEncodings = 0
	// kCFStringEncodingDOSChineseSimplif - Code page 936, also for Windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosChineseSimplif
	kCFStringEncodingDOSChineseSimplif CFStringEncodings = 0
	// kCFStringEncodingDOSChineseTrad - Code page 950, also for Windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosChineseTrad
	kCFStringEncodingDOSChineseTrad CFStringEncodings = 0
	// kCFStringEncodingDOSCyrillic - Code page 855, IBM Cyrillic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosCyrillic
	kCFStringEncodingDOSCyrillic CFStringEncodings = 0
	// kCFStringEncodingDOSGreek - Code page 737 (formerly code page 437G).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosGreek
	kCFStringEncodingDOSGreek CFStringEncodings = 0
	// kCFStringEncodingDOSGreek1 - Code page 851.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosGreek1
	kCFStringEncodingDOSGreek1 CFStringEncodings = 0
	// kCFStringEncodingDOSGreek2 - Code page 869, IBM Modern Greek.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosGreek2
	kCFStringEncodingDOSGreek2 CFStringEncodings = 0
	// kCFStringEncodingDOSHebrew - Code page 862.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosHebrew
	kCFStringEncodingDOSHebrew CFStringEncodings = 0
	// kCFStringEncodingDOSIcelandic - Code page 861.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosIcelandic
	kCFStringEncodingDOSIcelandic CFStringEncodings = 0
	// kCFStringEncodingDOSJapanese - Code page 932, also for Windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosJapanese
	kCFStringEncodingDOSJapanese CFStringEncodings = 0
	// kCFStringEncodingDOSKorean - Code page 949, also for Windows; Unified Hangul Code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosKorean
	kCFStringEncodingDOSKorean CFStringEncodings = 0
	// kCFStringEncodingDOSLatin1 - Code page 850, “Multilingual”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosLatin1
	kCFStringEncodingDOSLatin1 CFStringEncodings = 0
	// kCFStringEncodingDOSLatin2 - Code page 852, Slavic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosLatin2
	kCFStringEncodingDOSLatin2 CFStringEncodings = 0
	// kCFStringEncodingDOSLatinUS - Code page 437.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosLatinUS
	kCFStringEncodingDOSLatinUS CFStringEncodings = 0
	// kCFStringEncodingDOSNordic - Code page 865.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosNordic
	kCFStringEncodingDOSNordic CFStringEncodings = 0
	// kCFStringEncodingDOSPortuguese - Code page 860.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosPortuguese
	kCFStringEncodingDOSPortuguese CFStringEncodings = 0
	// kCFStringEncodingDOSRussian - Code page 866.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosRussian
	kCFStringEncodingDOSRussian CFStringEncodings = 0
	// kCFStringEncodingDOSThai - Code page 874, also for Windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosThai
	kCFStringEncodingDOSThai CFStringEncodings = 0
	// kCFStringEncodingDOSTurkish - Code page 857, IBM Turkish.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/dosTurkish
	kCFStringEncodingDOSTurkish CFStringEncodings = 0
	// kCFStringEncodingISOLatin10 - ISO 8859-16.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin10
	kCFStringEncodingISOLatin10 CFStringEncodings = 0
	// kCFStringEncodingISOLatin2 - ISO 8859-2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin2
	kCFStringEncodingISOLatin2 CFStringEncodings = 0
	// kCFStringEncodingISOLatin3 - ISO 8859-3.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin3
	kCFStringEncodingISOLatin3 CFStringEncodings = 0
	// kCFStringEncodingISOLatin4 - ISO 8859-4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin4
	kCFStringEncodingISOLatin4 CFStringEncodings = 0
	// kCFStringEncodingISOLatin5 - ISO 8859-9.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin5
	kCFStringEncodingISOLatin5 CFStringEncodings = 0
	// kCFStringEncodingISOLatin6 - ISO 8859-10.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin6
	kCFStringEncodingISOLatin6 CFStringEncodings = 0
	// kCFStringEncodingISOLatin7 - ISO 8859-13.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin7
	kCFStringEncodingISOLatin7 CFStringEncodings = 0
	// kCFStringEncodingISOLatin8 - ISO 8859-14.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin8
	kCFStringEncodingISOLatin8 CFStringEncodings = 0
	// kCFStringEncodingISOLatin9 - ISO 8859-15.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatin9
	kCFStringEncodingISOLatin9 CFStringEncodings = 0
	// kCFStringEncodingISOLatinArabic - ISO 8859-6, =ASMO 708, =DOS CP 708.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinArabic
	kCFStringEncodingISOLatinArabic CFStringEncodings = 0
	// kCFStringEncodingISOLatinCyrillic - ISO 8859-5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinCyrillic
	kCFStringEncodingISOLatinCyrillic CFStringEncodings = 0
	// kCFStringEncodingISOLatinGreek - ISO 8859-7.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinGreek
	kCFStringEncodingISOLatinGreek CFStringEncodings = 0
	// kCFStringEncodingISOLatinHebrew - ISO 8859-8.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinHebrew
	kCFStringEncodingISOLatinHebrew CFStringEncodings = 0
	// kCFStringEncodingISOLatinThai - ISO 8859-11.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/isoLatinThai
	kCFStringEncodingISOLatinThai CFStringEncodings = 0
	// kCFStringEncodingKSC_5601_92_Johab - KSC 5601-92 Johab annex.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/ksc_5601_92_Johab
	kCFStringEncodingKSC_5601_92_Johab CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macArabic
	kCFStringEncodingMacArabic CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macArmenian
	kCFStringEncodingMacArmenian CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macBengali
	kCFStringEncodingMacBengali CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macBurmese
	kCFStringEncodingMacBurmese CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macCeltic
	kCFStringEncodingMacCeltic CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macCentralEurRoman
	kCFStringEncodingMacCentralEurRoman CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macChineseSimp
	kCFStringEncodingMacChineseSimp CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macChineseTrad
	kCFStringEncodingMacChineseTrad CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macCroatian
	kCFStringEncodingMacCroatian CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macCyrillic
	kCFStringEncodingMacCyrillic CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macDevanagari
	kCFStringEncodingMacDevanagari CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macDingbats
	kCFStringEncodingMacDingbats CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macEthiopic
	kCFStringEncodingMacEthiopic CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macExtArabic
	kCFStringEncodingMacExtArabic CFStringEncodings = 0
	// kCFStringEncodingMacFarsi - Like MacArabic but uses Farsi digits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macFarsi
	kCFStringEncodingMacFarsi CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGaelic
	kCFStringEncodingMacGaelic CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGeorgian
	kCFStringEncodingMacGeorgian CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGreek
	kCFStringEncodingMacGreek CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGujarati
	kCFStringEncodingMacGujarati CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macGurmukhi
	kCFStringEncodingMacGurmukhi CFStringEncodings = 0
	// kCFStringEncodingMacHFS - Meta-value, should never appear in a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macHFS
	kCFStringEncodingMacHFS CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macHebrew
	kCFStringEncodingMacHebrew CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macIcelandic
	kCFStringEncodingMacIcelandic CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macInuit
	kCFStringEncodingMacInuit CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macJapanese
	kCFStringEncodingMacJapanese CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macKannada
	kCFStringEncodingMacKannada CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macKhmer
	kCFStringEncodingMacKhmer CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macKorean
	kCFStringEncodingMacKorean CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macLaotian
	kCFStringEncodingMacLaotian CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macMalayalam
	kCFStringEncodingMacMalayalam CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macMongolian
	kCFStringEncodingMacMongolian CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macOriya
	kCFStringEncodingMacOriya CFStringEncodings = 0
	// kCFStringEncodingMacRomanLatin1 - Mac OS Roman permuted to align with ISO Latin-1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macRomanLatin1
	kCFStringEncodingMacRomanLatin1 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macRomanian
	kCFStringEncodingMacRomanian CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macSinhalese
	kCFStringEncodingMacSinhalese CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macSymbol
	kCFStringEncodingMacSymbol CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macTamil
	kCFStringEncodingMacTamil CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macTelugu
	kCFStringEncodingMacTelugu CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macThai
	kCFStringEncodingMacThai CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macTibetan
	kCFStringEncodingMacTibetan CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macTurkish
	kCFStringEncodingMacTurkish CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macUkrainian
	kCFStringEncodingMacUkrainian CFStringEncodings = 0
	// kCFStringEncodingMacVT100 - VT100102 font from Comm Toolbox: Latin-1 repertoire + box drawing etc.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macVT100
	kCFStringEncodingMacVT100 CFStringEncodings = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/macVietnamese
	kCFStringEncodingMacVietnamese CFStringEncodings = 0
	// kCFStringEncodingNextStepJapanese - NextStep Japanese encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/nextStepJapanese
	kCFStringEncodingNextStepJapanese CFStringEncodings = 0
	// kCFStringEncodingShiftJIS - Plain Shift-JIS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/shiftJIS
	kCFStringEncodingShiftJIS CFStringEncodings = 0
	// kCFStringEncodingShiftJIS_X0213 - Shift-JIS format encoding of JIS X0213 planes 1 and 2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/shiftJIS_X0213
	kCFStringEncodingShiftJIS_X0213 CFStringEncodings = 0
	// kCFStringEncodingShiftJIS_X0213_00 - Shift-JIS format encoding of JIS X0213 planes 1 and 2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/shiftJIS_X0213_00
	kCFStringEncodingShiftJIS_X0213_00 CFStringEncodings = 0
	// kCFStringEncodingShiftJIS_X0213_MenKuTen - JIS X0213 in plane-row-column notation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/shiftJIS_X0213_MenKuTen
	kCFStringEncodingShiftJIS_X0213_MenKuTen CFStringEncodings = 0
	// kCFStringEncodingWindowsArabic - Code page 1256.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsArabic
	kCFStringEncodingWindowsArabic CFStringEncodings = 0
	// kCFStringEncodingWindowsBalticRim - Code page 1257.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsBalticRim
	kCFStringEncodingWindowsBalticRim CFStringEncodings = 0
	// kCFStringEncodingWindowsCyrillic - Code page 1251, Slavic Cyrillic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsCyrillic
	kCFStringEncodingWindowsCyrillic CFStringEncodings = 0
	// kCFStringEncodingWindowsGreek - Code page 1253.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsGreek
	kCFStringEncodingWindowsGreek CFStringEncodings = 0
	// kCFStringEncodingWindowsHebrew - Code page 1255.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsHebrew
	kCFStringEncodingWindowsHebrew CFStringEncodings = 0
	// kCFStringEncodingWindowsKoreanJohab - Code page 1361, for Windows NT.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsKoreanJohab
	kCFStringEncodingWindowsKoreanJohab CFStringEncodings = 0
	// kCFStringEncodingWindowsLatin2 - Code page 1250, Central Europe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsLatin2
	kCFStringEncodingWindowsLatin2 CFStringEncodings = 0
	// kCFStringEncodingWindowsLatin5 - Code page 1254, Turkish.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsLatin5
	kCFStringEncodingWindowsLatin5 CFStringEncodings = 0
	// kCFStringEncodingWindowsVietnamese - Code page 1258.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings/windowsVietnamese
	kCFStringEncodingWindowsVietnamese CFStringEncodings = 0
)

// CFStringNormalizationForm - Unicode normalization forms as described in Unicode Technical Report #15.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm
type CFStringNormalizationForm uint

const (
	// kCFStringNormalizationFormC - Canonical decomposition followed by canonical composition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm/C
	kCFStringNormalizationFormC CFStringNormalizationForm = 0
	// kCFStringNormalizationFormD - Canonical decomposition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm/D
	kCFStringNormalizationFormD CFStringNormalizationForm = 0
	// kCFStringNormalizationFormKC - Compatibility decomposition followed by canonical composition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm/KC
	kCFStringNormalizationFormKC CFStringNormalizationForm = 0
	// kCFStringNormalizationFormKD - Compatibility decomposition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm/KD
	kCFStringNormalizationFormKD CFStringNormalizationForm = 0
)

// CFStringTokenizerTokenType - Token types returned by 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType
type CFStringTokenizerTokenType uint

const (
	// kCFStringTokenizerTokenHasDerivedSubTokensMask - Compound token which may contain derived subtokens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/hasDerivedSubTokensMask
	kCFStringTokenizerTokenHasDerivedSubTokensMask CFStringTokenizerTokenType = 0
	// kCFStringTokenizerTokenHasHasNumbersMask - Appears to contain a number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/hasHasNumbersMask
	kCFStringTokenizerTokenHasHasNumbersMask CFStringTokenizerTokenType = 0
	// kCFStringTokenizerTokenHasNonLettersMask - Contains punctuation, symbols, and so on.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/hasNonLettersMask
	kCFStringTokenizerTokenHasNonLettersMask CFStringTokenizerTokenType = 0
	// kCFStringTokenizerTokenHasSubTokensMask - Compound token which may contain subtokens but with no derived subtokens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/hasSubTokensMask
	kCFStringTokenizerTokenHasSubTokensMask CFStringTokenizerTokenType = 0
	// kCFStringTokenizerTokenIsCJWordMask - Contains kana and/or ideographs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/isCJWordMask
	kCFStringTokenizerTokenIsCJWordMask CFStringTokenizerTokenType = 0
	// kCFStringTokenizerTokenNone - Has no token.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/kCFStringTokenizerTokenNone
	kCFStringTokenizerTokenNone CFStringTokenizerTokenType = 0
	// kCFStringTokenizerTokenNormal - Has a normal token.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType/normal
	kCFStringTokenizerTokenNormal CFStringTokenizerTokenType = 0
)

// CFTimeZoneNameStyle - Index type for constants used to specify styles of time zone names.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle
type CFTimeZoneNameStyle uint

const (
	// kCFTimeZoneNameStyleDaylightSaving - Specifies the daylight saving name style; for example, “Central Daylight Time” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/daylightSaving
	kCFTimeZoneNameStyleDaylightSaving CFTimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleGeneric - Specifies the generic name style, which does not distinguish between daylight saving and standard time; for example, “Central Time” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/generic
	kCFTimeZoneNameStyleGeneric CFTimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleShortDaylightSaving - Specifies the short daylight saving name style; for example, “CDT” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/shortDaylightSaving
	kCFTimeZoneNameStyleShortDaylightSaving CFTimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleShortGeneric - Specifies the short generic name style, which does not distinguish between daylight saving and standard time; for example, “CT” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/shortGeneric
	kCFTimeZoneNameStyleShortGeneric CFTimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleShortStandard - Specifies the short standard name style; for example, “CST” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/shortStandard
	kCFTimeZoneNameStyleShortStandard CFTimeZoneNameStyle = 0
	// kCFTimeZoneNameStyleStandard - Specifies the standard name style; for example, “Central Standard Time” for the Central time zone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle/standard
	kCFTimeZoneNameStyleStandard CFTimeZoneNameStyle = 0
)

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
	// kCFBookmarkResolutionWithoutMountingMask - Specifies that no volume should be mounted during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfBookmarkResolutionWithoutMountingMask
	kCFBookmarkResolutionWithoutMountingMask CFURLBookmarkResolutionOptions = 0
	// kCFBookmarkResolutionWithoutUIMask - Specifies that no UI feedback accompany resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfBookmarkResolutionWithoutUIMask
	kCFBookmarkResolutionWithoutUIMask CFURLBookmarkResolutionOptions = 0
	// kCFURLBookmarkResolutionWithSecurityScope - Specifies that the security scope, applied to the bookmark when it was created, should be used during resolution of the bookmark data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfurlBookmarkResolutionWithSecurityScope
	kCFURLBookmarkResolutionWithSecurityScope CFURLBookmarkResolutionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfurlBookmarkResolutionWithoutImplicitStartAccessing
	kCFURLBookmarkResolutionWithoutImplicitStartAccessing CFURLBookmarkResolutionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions/cfurlBookmarkResolutionWithoutMountingMask
	kCFURLBookmarkResolutionWithoutMountingMask CFURLBookmarkResolutionOptions = 0
)

// CFURLComponentType - The types of components in a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType
type CFURLComponentType uint

const (
	// kCFURLComponentFragment - The URL’s fragment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/fragment
	kCFURLComponentFragment CFURLComponentType = 0
	// kCFURLComponentHost - The URL’s host.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/host
	kCFURLComponentHost CFURLComponentType = 0
	// kCFURLComponentNetLocation - The URL’s network location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/netLocation
	kCFURLComponentNetLocation CFURLComponentType = 0
	// kCFURLComponentParameterString - The URL’s parameter string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/parameterString
	kCFURLComponentParameterString CFURLComponentType = 0
	// kCFURLComponentPassword - The user’s password.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/password
	kCFURLComponentPassword CFURLComponentType = 0
	// kCFURLComponentPath - The URL’s path component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/path
	kCFURLComponentPath CFURLComponentType = 0
	// kCFURLComponentPort - The URL’s port.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/port
	kCFURLComponentPort CFURLComponentType = 0
	// kCFURLComponentQuery - The URL’s query.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/query
	kCFURLComponentQuery CFURLComponentType = 0
	// kCFURLComponentResourceSpecifier - The URL’s resource specifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/resourceSpecifier
	kCFURLComponentResourceSpecifier CFURLComponentType = 0
	// kCFURLComponentScheme - The URL’s scheme.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/scheme
	kCFURLComponentScheme CFURLComponentType = 0
	// kCFURLComponentUser - The URL’s user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/user
	kCFURLComponentUser CFURLComponentType = 0
	// kCFURLComponentUserInfo - The user’s information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType/userInfo
	kCFURLComponentUserInfo CFURLComponentType = 0
)

// CFURLEnumeratorOptions - Options for controlling enumerator behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions
type CFURLEnumeratorOptions uint

const (
	// kCFURLEnumeratorDescendRecursively - The enumerator recurses into each subdirectory enumerated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/descendRecursively
	kCFURLEnumeratorDescendRecursively CFURLEnumeratorOptions = 0
	// kCFURLEnumeratorGenerateFileReferenceURLs - The enumerator generates file reference URLs instead of file path URLs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/generateFileReferenceURLs
	kCFURLEnumeratorGenerateFileReferenceURLs CFURLEnumeratorOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/generateRelativePathURLs
	kCFURLEnumeratorGenerateRelativePathURLs CFURLEnumeratorOptions = 0
	// kCFURLEnumeratorIncludeDirectoriesPostOrder - If provided along with the   option, the recursive enumerator returns a directory’s URL after returning the URLs of the directory’s descendents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/includeDirectoriesPostOrder
	kCFURLEnumeratorIncludeDirectoriesPostOrder CFURLEnumeratorOptions = 0
	// kCFURLEnumeratorIncludeDirectoriesPreOrder - If provided along with the   option, the recursive enumerator returns a directory’s URL before returning the URLs of the directory’s descendents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/includeDirectoriesPreOrder
	kCFURLEnumeratorIncludeDirectoriesPreOrder CFURLEnumeratorOptions = 0
	// kCFURLEnumeratorDefaultBehavior - The enumerator performs its default behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/kCFURLEnumeratorDefaultBehavior
	kCFURLEnumeratorDefaultBehavior CFURLEnumeratorOptions = 0
	// kCFURLEnumeratorSkipInvisibles - The enumerator skips “hidden” or “invisible” objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/skipInvisibles
	kCFURLEnumeratorSkipInvisibles CFURLEnumeratorOptions = 0
	// kCFURLEnumeratorSkipPackageContents - The enumerator skips package directory contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions/skipPackageContents
	kCFURLEnumeratorSkipPackageContents CFURLEnumeratorOptions = 0
)

// CFURLEnumeratorResult - Result codes from the 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult
type CFURLEnumeratorResult uint

const (
	// kCFURLEnumeratorDirectoryPostOrderSuccess - The recursive post-order enumerator returned the URL for a directory after having returned the URLs for all of the directory’s descendents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult/directoryPostOrderSuccess
	kCFURLEnumeratorDirectoryPostOrderSuccess CFURLEnumeratorResult = 0
	// kCFURLEnumeratorEnd - The enumeration is complete.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult/end
	kCFURLEnumeratorEnd CFURLEnumeratorResult = 0
	// kCFURLEnumeratorError - An error occurred during enumeration. The   parameter of the function is populated with error information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult/error
	kCFURLEnumeratorError CFURLEnumeratorResult = 0
	// kCFURLEnumeratorSuccess - The enumerator was advanced successfully and returned a valid URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult/success
	kCFURLEnumeratorSuccess CFURLEnumeratorResult = 0
)

// CFURLError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError
type CFURLError uint

const (
	// kCFURLImproperArgumentsError - Indicates one or more arguments are improper.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/improperArgumentsError
	kCFURLImproperArgumentsError CFURLError = 0
	// kCFURLPropertyKeyUnavailableError - Indicates a property key was unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/propertyKeyUnavailableError
	kCFURLPropertyKeyUnavailableError CFURLError = 0
	// kCFURLRemoteHostUnavailableError - Indicates a remote host is unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/remoteHostUnavailableError
	kCFURLRemoteHostUnavailableError CFURLError = 0
	// kCFURLResourceAccessViolationError - Indicates an error in accessing a resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/resourceAccessViolationError
	kCFURLResourceAccessViolationError CFURLError = 0
	// kCFURLResourceNotFoundError - Indicates a resource was not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/resourceNotFoundError
	kCFURLResourceNotFoundError CFURLError = 0
	// kCFURLTimeoutError - Indicates a timeout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/timeoutError
	kCFURLTimeoutError CFURLError = 0
	// kCFURLUnknownError - Indicates an unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/unknownError
	kCFURLUnknownError CFURLError = 0
	// kCFURLUnknownPropertyKeyError - Indicates a property key is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/unknownPropertyKeyError
	kCFURLUnknownPropertyKeyError CFURLError = 0
	// kCFURLUnknownSchemeError - Indicates that the scheme is not recognized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLError/unknownSchemeError
	kCFURLUnknownSchemeError CFURLError = 0
)

// CFURLPathStyle - Options you can use to determine how CFURL functions parse a file system path name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle
type CFURLPathStyle uint

const (
	// kCFURLWindowsPathStyle - Indicates a Windows style path name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle/cfurlWindowsPathStyle
	kCFURLWindowsPathStyle CFURLPathStyle = 0
	// kCFURLHFSPathStyle - Indicates a HFS style path name. Components are colon delimited. A leading colon indicates a relative path, otherwise the first path component denotes the volume.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle/cfurlhfsPathStyle
	kCFURLHFSPathStyle CFURLPathStyle = 0
	// kCFURLPOSIXPathStyle - Indicates a POSIX style path name. Components are slash delimited. A leading slash indicates an absolute path; a trailing slash is not significant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle/cfurlposixPathStyle
	kCFURLPOSIXPathStyle CFURLPathStyle = 0
)

// CFXMLEntityTypeCode - The entity type identification codes that the parser uses to describe XML entities.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode
type CFXMLEntityTypeCode uint

const (
	// kCFXMLEntityTypeCharacter - Indicates a character entity type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/character
	kCFXMLEntityTypeCharacter CFXMLEntityTypeCode = 0
	// kCFXMLEntityTypeParameter - Implies a parsed, internal entity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/parameter
	kCFXMLEntityTypeParameter CFXMLEntityTypeCode = 0
	// kCFXMLEntityTypeParsedExternal - Indicates a parsed, external entity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/parsedExternal
	kCFXMLEntityTypeParsedExternal CFXMLEntityTypeCode = 0
	// kCFXMLEntityTypeParsedInternal - Indicates a parsed, internal entity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/parsedInternal
	kCFXMLEntityTypeParsedInternal CFXMLEntityTypeCode = 0
	// kCFXMLEntityTypeUnparsed - Indicates an unparsed entity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode/unparsed
	kCFXMLEntityTypeUnparsed CFXMLEntityTypeCode = 0
)

// CFXMLNodeTypeCode - The various XML data type identification codes that the parser uses to describe XML structures.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode
type CFXMLNodeTypeCode uint

const (
	// kCFXMLNodeTypeAttribute - Currently not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/attribute
	kCFXMLNodeTypeAttribute CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeAttributeListDeclaration - Indicates an attribute list declaration where the data string is the tag name and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/attributeListDeclaration
	kCFXMLNodeTypeAttributeListDeclaration CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeCDATASection - Indicates a CDATA section where the data string is the text of the CDATA and the additional information is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/cdataSection
	kCFXMLNodeTypeCDATASection CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeComment - Indicates a comment section where the data string is the text of the comment and the additional information is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/comment
	kCFXMLNodeTypeComment CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeDocument - Indicates a document where the data string is   and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/document
	kCFXMLNodeTypeDocument CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeDocumentFragment - Currently not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/documentFragment
	kCFXMLNodeTypeDocumentFragment CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeDocumentType - Indicates a document type where the data string is the name given to the top-level element and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/documentType
	kCFXMLNodeTypeDocumentType CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeElement - Indicates an element where the data string is the name of the tag and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/element
	kCFXMLNodeTypeElement CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeElementTypeDeclaration - Indicates an element type declaration where the data string is the tag name and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/elementTypeDeclaration
	kCFXMLNodeTypeElementTypeDeclaration CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeEntity - Indicates an entity where the data string is the name of the entity and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/entity
	kCFXMLNodeTypeEntity CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeEntityReference - Indicates an entity reference where the data string is the name of the referenced entity and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/entityReference
	kCFXMLNodeTypeEntityReference CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeNotation - Indicates a notation where the data string is the notation name and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/notation
	kCFXMLNodeTypeNotation CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeProcessingInstruction - Indicates a processing instruction where the data string is the name of the target and the additional information is a pointer to a   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/processingInstruction
	kCFXMLNodeTypeProcessingInstruction CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeText - Indicates a text section where the data string is the text’s contents and the additional information is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/text
	kCFXMLNodeTypeText CFXMLNodeTypeCode = 0
	// kCFXMLNodeTypeWhitespace - Indicates white space where the data string is the text of the white space and the additional information is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeTypeCode/whitespace
	kCFXMLNodeTypeWhitespace CFXMLNodeTypeCode = 0
)

// CFXMLParserOptions - Options you can use to control the parser’s treatment of an XML document.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions
type CFXMLParserOptions uint

const (
	// kCFXMLParserAddImpliedAttributes - Where the DTD specifies implied attribute-value pairs for a particular element, add those pairs to any occurrences of the element in the element tree. Currently not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/addImpliedAttributes
	kCFXMLParserAddImpliedAttributes CFXMLParserOptions = 0
	// kCFXMLParserAllOptions - Makes the parser do the most work, returning only the pure elementtree.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/allOptions
	kCFXMLParserAllOptions CFXMLParserOptions = 0
	// kCFXMLParserNoOptions - Leaves the XML as “intact” as possible (reports all structures; performs no replacements).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/kCFXMLParserNoOptions
	kCFXMLParserNoOptions CFXMLParserOptions = 0
	// kCFXMLParserReplacePhysicalEntities - Replaces declared entities like  ;. Note that other than the 5 predefined entities ( ,  ,  ,  ,  ), these must be defined in the DTD. Currently not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/replacePhysicalEntities
	kCFXMLParserReplacePhysicalEntities CFXMLParserOptions = 0
	// kCFXMLParserResolveExternalEntities - Resolves all external entities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/resolveExternalEntities
	kCFXMLParserResolveExternalEntities CFXMLParserOptions = 0
	// kCFXMLParserSkipMetaData - Silently skip over metadata constructs (the DTD and comments).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/skipMetaData
	kCFXMLParserSkipMetaData CFXMLParserOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/skipWhitespace
	kCFXMLParserSkipWhitespace CFXMLParserOptions = 0
	// kCFXMLParserValidateDocument - Validates the document against its grammar from the DTD, reporting any errors. Currently not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions/validateDocument
	kCFXMLParserValidateDocument CFXMLParserOptions = 0
)

// CFXMLParserStatusCode - The various status and error flags that can be returned by the parser.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode
type CFXMLParserStatusCode uint

const (
	// kCFXMLErrorElementlessDocument - Indicates a document containing no elements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorElementlessDocument
	kCFXMLErrorElementlessDocument CFXMLParserStatusCode = 0
	// kCFXMLErrorEncodingConversionFailure - Indicates an encoding conversion error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorEncodingConversionFailure
	kCFXMLErrorEncodingConversionFailure CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedCDSect - Indicates a malformed CDATA section.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedCDSect
	kCFXMLErrorMalformedCDSect CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedCharacterReference - Indicates a malformed character reference.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedCharacterReference
	kCFXMLErrorMalformedCharacterReference CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedCloseTag - Indicates a malformed close tag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedCloseTag
	kCFXMLErrorMalformedCloseTag CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedComment - Indicates a malformed comment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedComment
	kCFXMLErrorMalformedComment CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedDTD - Indicates a malformed DTD.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedDTD
	kCFXMLErrorMalformedDTD CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedDocument - Indicates a malformed document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedDocument
	kCFXMLErrorMalformedDocument CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedName - Indicates a malformed name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedName
	kCFXMLErrorMalformedName CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedParsedCharacterData - Indicates malformed character data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedParsedCharacterData
	kCFXMLErrorMalformedParsedCharacterData CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedProcessingInstruction - Indicates a malformed processing instruction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedProcessingInstruction
	kCFXMLErrorMalformedProcessingInstruction CFXMLParserStatusCode = 0
	// kCFXMLErrorMalformedStartTag - Indicates a malformed start tag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorMalformedStartTag
	kCFXMLErrorMalformedStartTag CFXMLParserStatusCode = 0
	// kCFXMLErrorNoData - Indicates a no data error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorNoData
	kCFXMLErrorNoData CFXMLParserStatusCode = 0
	// kCFXMLErrorUnexpectedEOF - Indicates an unexpected EOF occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorUnexpectedEOF
	kCFXMLErrorUnexpectedEOF CFXMLParserStatusCode = 0
	// kCFXMLErrorUnknownEncoding - Indicates an unknown encoding error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/errorUnknownEncoding
	kCFXMLErrorUnknownEncoding CFXMLParserStatusCode = 0
	// kCFXMLStatusParseSuccessful - Indicates the parser was successful.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/kCFXMLStatusParseSuccessful
	kCFXMLStatusParseSuccessful CFXMLParserStatusCode = 0
	// kCFXMLStatusParseInProgress - Indicates the parser is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/statusParseInProgress
	kCFXMLStatusParseInProgress CFXMLParserStatusCode = 0
	// kCFXMLStatusParseNotBegun - Indicates the parser has not begun.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode/statusParseNotBegun
	kCFXMLStatusParseNotBegun CFXMLParserStatusCode = 0
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


