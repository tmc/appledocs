// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
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

// ByteCountFormatterCountStyle - Specifies display of file or storage byte counts. The display style is platform specific.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/CountStyle-swift.enum
type ByteCountFormatterCountStyle uint

const (
	// ByteCountFormatterCountStyleBinary - Causes 1024 bytes to be shown as 1 KB. It is better to use   or   in most cases.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/CountStyle-swift.enum/binary
	ByteCountFormatterCountStyleBinary ByteCountFormatterCountStyle = 3
	// ByteCountFormatterCountStyleDecimal - Causes 1000 bytes to be shown as 1 KB. It is better to use   or   in most cases.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/CountStyle-swift.enum/decimal
	ByteCountFormatterCountStyleDecimal ByteCountFormatterCountStyle = 2
	// ByteCountFormatterCountStyleFile - Specifies display of file byte counts. The actual behavior for this is platform-specific; in macOS 10.8, this uses the decimal style, but that may change over time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/CountStyle-swift.enum/file
	ByteCountFormatterCountStyleFile ByteCountFormatterCountStyle = 0
	// ByteCountFormatterCountStyleMemory - Specifies display of memory byte counts. The actual behavior for this is platform-specific; in macOS 10.8, this uses the binary style, but that may change over time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/CountStyle-swift.enum/memory
	ByteCountFormatterCountStyleMemory ByteCountFormatterCountStyle = 1
)

// ByteCountFormatterUnits - Specifies the units appropriate for the formatter to display. Specifying any units explicitly causes just those units to be used in showing the number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units
type ByteCountFormatterUnits uint

const (
	// ByteCountFormatterUseAll - Can use any unit in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/useAll
	ByteCountFormatterUseAll ByteCountFormatterUnits = 65535
	// ByteCountFormatterUseBytes - Displays bytes in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/useBytes
	ByteCountFormatterUseBytes ByteCountFormatterUnits = 1
	// ByteCountFormatterUseEB - Displays exabytes in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/useEB
	ByteCountFormatterUseEB ByteCountFormatterUnits = 64
	// ByteCountFormatterUseGB - Displays gigabytes in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/useGB
	ByteCountFormatterUseGB ByteCountFormatterUnits = 8
	// ByteCountFormatterUseKB - Displays kilobytes in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/useKB
	ByteCountFormatterUseKB ByteCountFormatterUnits = 2
	// ByteCountFormatterUseMB - Displays megabytes in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/useMB
	ByteCountFormatterUseMB ByteCountFormatterUnits = 4
	// ByteCountFormatterUsePB - Displays petabyte in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/usePB
	ByteCountFormatterUsePB ByteCountFormatterUnits = 32
	// ByteCountFormatterUseTB - Displays terabytes in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/useTB
	ByteCountFormatterUseTB ByteCountFormatterUnits = 16
	// ByteCountFormatterUseYBOrHigher - Displays yottabytes in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/useYBOrHigher
	ByteCountFormatterUseYBOrHigher ByteCountFormatterUnits = 255
	// ByteCountFormatterUseZB - Displays zettabytes in the formatter content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units/useZB
	ByteCountFormatterUseZB ByteCountFormatterUnits = 128
	// ByteCountFormatterUseDefault - This causes default units appropriate for the platform to be used. This is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSByteCountFormatterUnits/NSByteCountFormatterUseDefault
	ByteCountFormatterUseDefault ByteCountFormatterUnits = 0
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
	// DateComponentsFormatterZeroFormattingBehaviorNone - No formatting behavior. This behavior prevents the dropping of zero values or adding of zeroes for padding. For example, with hours, minutes, and seconds displayed, the abbreviated value for one hour and 10 seconds is “1h 0m 10s”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponentsFormatterZeroFormattingBehavior/NSDateComponentsFormatterZeroFormattingBehaviorNone
	DateComponentsFormatterZeroFormattingBehaviorNone DateComponentsFormatterZeroFormattingBehavior = 0
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

// DistributedNotificationOptions - These constants specify the behavior of notifications posted using the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/Options
type DistributedNotificationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/Options/deliverImmediately
	DistributedNotificationDeliverImmediately DistributedNotificationOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/Options/postToAllSessions
	DistributedNotificationPostToAllSessions DistributedNotificationOptions = 2
)

// NotificationSuspensionBehavior - These constants specify the types of notification delivery suspension behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior
type NotificationSuspensionBehavior uint

const (
	// NotificationSuspensionBehaviorCoalesce - The server only queues the last notification of the specified name and object; earlier notifications are dropped. In cover methods for which suspension behavior is not an explicit argument,   is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior/coalesce
	NotificationSuspensionBehaviorCoalesce NotificationSuspensionBehavior = 2
	// NotificationSuspensionBehaviorDeliverImmediately - The server delivers notifications matching this registration irrespective of whether   with an argument of   has been called. When a notification with this suspension behavior is matched, it has the effect of first flushing any queued notifications. The effect is as if   with an argument of   were first called if the application is suspended, followed by the notification in question being delivered, followed by a transition back to the previous suspended or unsuspended state.
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

// EnergyFormatterUnit - The units supported by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit
type EnergyFormatterUnit uint

const (
	// EnergyFormatterUnitCalorie - The calorie unit. This unit is often used in chemistry. It is abbreviated as “cal.”
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit/calorie
	EnergyFormatterUnitCalorie EnergyFormatterUnit = 1792
	// EnergyFormatterUnitJoule - The joule unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit/joule
	EnergyFormatterUnitJoule EnergyFormatterUnit = 11
	// EnergyFormatterUnitKilocalorie - The kilocalorie unit. This unit is used for food calories in some locales. In general, it is abbreviated as “kcal.” However, it may be abbreviated as “C” when used to represent food calories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit/kilocalorie
	EnergyFormatterUnitKilocalorie EnergyFormatterUnit = 1792
	// EnergyFormatterUnitKilojoule - The kilojoule unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit/kilojoule
	EnergyFormatterUnitKilojoule EnergyFormatterUnit = 14
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

// FormattingUnitStyle - Specifies the width of the unit, determining the textual representation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/UnitStyle
type FormattingUnitStyle uint

const (
	// FormattingUnitStyleLong - Specifies a long unit style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/UnitStyle/long
	FormattingUnitStyleLong FormattingUnitStyle = 3
	// FormattingUnitStyleMedium - Specifies a medium unit style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/UnitStyle/medium
	FormattingUnitStyleMedium FormattingUnitStyle = 2
	// FormattingUnitStyleShort - Specifies a short unit style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/UnitStyle/short
	FormattingUnitStyleShort FormattingUnitStyle = 1
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

// JSONReadingOptions - Options used when creating Foundation objects from JSON data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions
type JSONReadingOptions uint

const (
	// JSONReadingAllowFragments - A deprecated option that specifies that the parser should allow top-level objects that aren’t arrays or dictionaries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions/allowFragments
	JSONReadingAllowFragments JSONReadingOptions = 7
	// JSONReadingFragmentsAllowed - Specifies that the parser allows top-level objects that aren’t arrays or dictionaries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions/fragmentsAllowed
	JSONReadingFragmentsAllowed JSONReadingOptions = 4
	// JSONReadingJSON5Allowed - Specifies that reading serialized JSON data supports the JSON5 syntax.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions/json5Allowed
	JSONReadingJSON5Allowed JSONReadingOptions = 5
	// JSONReadingMutableContainers - Specifies that arrays and dictionaries in the returned object are mutable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions/mutableContainers
	JSONReadingMutableContainers JSONReadingOptions = 1
	// JSONReadingMutableLeaves - Specifies that leaf strings in the JSON object graph are mutable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions/mutableLeaves
	JSONReadingMutableLeaves JSONReadingOptions = 2
	// JSONReadingTopLevelDictionaryAssumed - Specifies that the parser assumes the top level of the JSON data is a dictionary, even if it doesn’t begin and end with curly braces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions/topLevelDictionaryAssumed
	JSONReadingTopLevelDictionaryAssumed JSONReadingOptions = 6
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
	// JSONWritingPrettyPrinted - Specifies that the output uses white space and indentation to make the resulting data more readable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/WritingOptions/prettyPrinted
	JSONWritingPrettyPrinted JSONWritingOptions = 1
	// JSONWritingSortedKeys - Specifies that the output sorts keys in lexicographic order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/WritingOptions/sortedKeys
	JSONWritingSortedKeys JSONWritingOptions = 2
	// JSONWritingWithoutEscapingSlashes - Specifies that the output doesn’t prefix slash characters with escape characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/WritingOptions/withoutEscapingSlashes
	JSONWritingWithoutEscapingSlashes JSONWritingOptions = 5
)

// LengthFormatterUnit - The units supported by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit
type LengthFormatterUnit uint

const (
	// LengthFormatterUnitCentimeter - The centimeter unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit/centimeter
	LengthFormatterUnitCentimeter LengthFormatterUnit = 9
	// LengthFormatterUnitFoot - The foot unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit/foot
	LengthFormatterUnitFoot LengthFormatterUnit = 1280
	// LengthFormatterUnitInch - The inch unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit/inch
	LengthFormatterUnitInch LengthFormatterUnit = 1280
	// LengthFormatterUnitKilometer - The kilometer unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit/kilometer
	LengthFormatterUnitKilometer LengthFormatterUnit = 14
	// LengthFormatterUnitMeter - The meter unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit/meter
	LengthFormatterUnitMeter LengthFormatterUnit = 11
	// LengthFormatterUnitMile - The mile unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit/mile
	LengthFormatterUnitMile LengthFormatterUnit = 1280
	// LengthFormatterUnitMillimeter - The millimeter unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit/millimeter
	LengthFormatterUnitMillimeter LengthFormatterUnit = 8
	// LengthFormatterUnitYard - The yard unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit/yard
	LengthFormatterUnitYard LengthFormatterUnit = 1280
)

// MassFormatterUnit - The units supported by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/Unit
type MassFormatterUnit uint

const (
	// MassFormatterUnitGram - The gram unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/Unit/gram
	MassFormatterUnitGram MassFormatterUnit = 11
	// MassFormatterUnitKilogram - The kilogram unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/Unit/kilogram
	MassFormatterUnitKilogram MassFormatterUnit = 14
	// MassFormatterUnitOunce - The ounce unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/Unit/ounce
	MassFormatterUnitOunce MassFormatterUnit = 1536
	// MassFormatterUnitPound - The pound unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/Unit/pound
	MassFormatterUnitPound MassFormatterUnit = 1536
	// MassFormatterUnitStone - The stone unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/Unit/stone
	MassFormatterUnitStone MassFormatterUnit = 1536
)

// MeasurementFormatterUnitOptions - Measurement formatter options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/UnitOptions-swift.struct
type MeasurementFormatterUnitOptions uint

const (
	// MeasurementFormatterUnitOptionsNaturalScale - Specifies that representations of measurements are reduced and converted into a more convenient unit, when possible. For example, a quantity of 12000 meters would be represented as  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/UnitOptions-swift.struct/naturalScale
	MeasurementFormatterUnitOptionsNaturalScale MeasurementFormatterUnitOptions = 2
	// MeasurementFormatterUnitOptionsProvidedUnit - Specifies that the provided unit for the measurement should be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/UnitOptions-swift.struct/providedUnit
	MeasurementFormatterUnitOptionsProvidedUnit MeasurementFormatterUnitOptions = 1
	// MeasurementFormatterUnitOptionsTemperatureWithoutUnit - Specifies that representations of a measurement with the   unit omit the letter denoting the temperature scale. For example, a temperature measurement with value equal to 72 using the   would be represented as   rather than  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/UnitOptions-swift.struct/temperatureWithoutUnit
	MeasurementFormatterUnitOptionsTemperatureWithoutUnit MeasurementFormatterUnitOptions = 4
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
	AttributedStringFormattingInsertArgumentAttributesWithoutMerging AttributedStringFormattingOptions = 0
	AttributedStringFormattingApplyReplacementIndexAttribute AttributedStringFormattingOptions = 1
)

// AttributedStringMarkdownInterpretedSyntax - A type that represents the syntax for intepreting a Markdown string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownInterpretedSyntax
type AttributedStringMarkdownInterpretedSyntax int

const (
	AttributedStringMarkdownInterpretedSyntaxFull AttributedStringMarkdownInterpretedSyntax = 0
	AttributedStringMarkdownInterpretedSyntaxInlineOnly AttributedStringMarkdownInterpretedSyntax = 1
	AttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace AttributedStringMarkdownInterpretedSyntax = 2
)

// AttributedStringMarkdownParsingFailurePolicy - A type that represents policies for handling parsing failures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingFailurePolicy
type AttributedStringMarkdownParsingFailurePolicy int

const (
	AttributedStringMarkdownParsingFailureReturnError AttributedStringMarkdownParsingFailurePolicy = 0
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

// CollectionChangeType - The type of change represented in computing the difference of an ordered collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCollectionChangeType
type CollectionChangeType uint

const (
	CollectionChangeInsert CollectionChangeType = 0
	CollectionChangeRemove CollectionChangeType = 1
)

// ComparisonPredicateModifier - Constants that describe the possible types of modifier for a comparison predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier
type ComparisonPredicateModifier uint

const (
	// AllPredicateModifier - A predicate to compare all entries in the destination of a to-many relationship.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier/all
	AllPredicateModifier ComparisonPredicateModifier = 1
	// AnyPredicateModifier - A predicate to match with any entry in the destination of a to-many relationship.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier/any
	AnyPredicateModifier ComparisonPredicateModifier = 2
	// DirectPredicateModifier - A predicate to compare directly the left and right hand sides.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier/direct
	DirectPredicateModifier ComparisonPredicateModifier = 0
)

// PredicateOperatorType - Defines the type of comparison for a comparison predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator
type PredicateOperatorType uint

const (
	// BeginsWithPredicateOperatorType - A begins-with predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/beginsWith
	BeginsWithPredicateOperatorType PredicateOperatorType = 8
	// BetweenPredicateOperatorType - A predicate to determine if the left hand side lies at or between bounds specified by the right hand side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/between
	BetweenPredicateOperatorType PredicateOperatorType = 13
	// ContainsPredicateOperatorType - A predicate to determine if the left hand side contains the right hand side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/contains
	ContainsPredicateOperatorType PredicateOperatorType = 12
	// CustomSelectorPredicateOperatorType - A predicate that uses a custom selector that takes a single argument and returns a   value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/customSelector
	CustomSelectorPredicateOperatorType PredicateOperatorType = 11
	// EndsWithPredicateOperatorType - An ends-with predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/endsWith
	EndsWithPredicateOperatorType PredicateOperatorType = 9
	// EqualToPredicateOperatorType - An equal-to predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/equalTo
	EqualToPredicateOperatorType PredicateOperatorType = 4
	// GreaterThanPredicateOperatorType - A greater-than predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/greaterThan
	GreaterThanPredicateOperatorType PredicateOperatorType = 2
	// GreaterThanOrEqualToPredicateOperatorType - A greater-than-or-equal-to predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/greaterThanOrEqualTo
	GreaterThanOrEqualToPredicateOperatorType PredicateOperatorType = 3
	// InPredicateOperatorType - A predicate to determine if the left hand side is in the right hand side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/in
	InPredicateOperatorType PredicateOperatorType = 10
	// LessThanPredicateOperatorType - A less-than predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/lessThan
	LessThanPredicateOperatorType PredicateOperatorType = 0
	// LessThanOrEqualToPredicateOperatorType - A less-than-or-equal-to predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/lessThanOrEqualTo
	LessThanOrEqualToPredicateOperatorType PredicateOperatorType = 1
	// LikePredicateOperatorType - A simple subset of the MATCHES predicate, similar in behavior to SQL  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/like
	LikePredicateOperatorType PredicateOperatorType = 7
	// MatchesPredicateOperatorType - A full regular expression matching predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/matches
	MatchesPredicateOperatorType PredicateOperatorType = 6
	// NotEqualToPredicateOperatorType - A not-equal-to predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator/notEqualTo
	NotEqualToPredicateOperatorType PredicateOperatorType = 5
)

// ComparisonPredicateOptions - Constants that describe the possible types of string comparison for comparison predicates.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
type ComparisonPredicateOptions uint

const (
	// CaseInsensitivePredicateOption - A case-insensitive predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct/caseInsensitive
	CaseInsensitivePredicateOption ComparisonPredicateOptions = 1
	// DiacriticInsensitivePredicateOption - A diacritic-insensitive predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct/diacriticInsensitive
	DiacriticInsensitivePredicateOption ComparisonPredicateOptions = 2
	// NormalizedPredicateOption - A predicate that indicates you’ve preprocessed the strings to compare.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct/normalized
	NormalizedPredicateOption ComparisonPredicateOptions = 3
)

// CompoundPredicateType - Constants that describe the possible types of a compound predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType
type CompoundPredicateType uint

const (
	// AndPredicateType - A logical AND predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType/and
	AndPredicateType CompoundPredicateType = 1
	// NotPredicateType - A logical NOT predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType/not
	NotPredicateType CompoundPredicateType = 0
	// OrPredicateType - A logical OR predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType/or
	OrPredicateType CompoundPredicateType = 2
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

// ExpressionType - Defines the possible types of an expression.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum
type ExpressionType uint

const (
	// AggregateExpressionType - An expression that defines an aggregate of   objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/aggregate
	AggregateExpressionType ExpressionType = 9
	// AnyKeyExpressionType - An expression that represents any key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/anyKey
	AnyKeyExpressionType ExpressionType = 10
	// BlockExpressionType - An expression that uses a Block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/block
	BlockExpressionType ExpressionType = 19
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/conditional
	ConditionalExpressionType ExpressionType = 20
	// ConstantValueExpressionType - An expression that always returns the same value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/constantValue
	ConstantValueExpressionType ExpressionType = 0
	// EvaluatedObjectExpressionType - An expression that always returns the parameter object itself.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/evaluatedObject
	EvaluatedObjectExpressionType ExpressionType = 1
	// FunctionExpressionType - An expression that returns the result of evaluating a function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/function
	FunctionExpressionType ExpressionType = 4
	// IntersectSetExpressionType - An expression that creates an intersection of the results of two nested expressions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/intersectSet
	IntersectSetExpressionType ExpressionType = 6
	// KeyPathExpressionType - An expression that returns something that can be used as a key path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/keyPath
	KeyPathExpressionType ExpressionType = 3
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
	// VariableExpressionType - An expression that always returns whatever value is associated with the key specified by ‘variable’ in the bindings dictionary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum/variable
	VariableExpressionType ExpressionType = 2
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
	// GrammaticalGenderFeminine - The feminine grammatical gender.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalGender/feminine
	GrammaticalGenderFeminine GrammaticalGender = 1
	// GrammaticalGenderMasculine - The masculine grammatical gender.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalGender/masculine
	GrammaticalGenderMasculine GrammaticalGender = 2
	// GrammaticalGenderNeuter - A value to not specify gender when inflecting a string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalGender/neuter
	GrammaticalGenderNeuter GrammaticalGender = 3
	// GrammaticalGenderNotSet - A value that indicates the gender is unset.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalGender/notSet
	GrammaticalGenderNotSet GrammaticalGender = 0
)

// GrammaticalNumber - A representation of grammatical number, used for inflecting strings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber
type GrammaticalNumber uint

const (
	// GrammaticalNumberNotSet - A value that indicates the number is unset.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber/notSet
	GrammaticalNumberNotSet GrammaticalNumber = 0
	// GrammaticalNumberPlural - Multiple persons or things, as used for a grammatical number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber/plural
	GrammaticalNumberPlural GrammaticalNumber = 3
	// GrammaticalNumberPluralFew - A small number of persons or things, as used for a grammatical number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber/pluralFew
	GrammaticalNumberPluralFew GrammaticalNumber = 5
	// GrammaticalNumberPluralMany - A large number of persons or things, as used for a grammatical number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber/pluralMany
	GrammaticalNumberPluralMany GrammaticalNumber = 6
	// GrammaticalNumberPluralTwo - Two persons or things, as used for a grammatical number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber/pluralTwo
	GrammaticalNumberPluralTwo GrammaticalNumber = 4
	// GrammaticalNumberSingular - A single person or thing, as used for a grammatical number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber/singular
	GrammaticalNumberSingular GrammaticalNumber = 1
	// GrammaticalNumberZero - Zero persons or things, as used for a grammatical number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber/zero
	GrammaticalNumberZero GrammaticalNumber = 2
)

// GrammaticalPartOfSpeech - A representation of grammatical parts of speech, used for inflecting strings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech
type GrammaticalPartOfSpeech uint

const (
	// GrammaticalPartOfSpeechAbbreviation - An abbreviation, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/abbreviation
	GrammaticalPartOfSpeechAbbreviation GrammaticalPartOfSpeech = 14
	// GrammaticalPartOfSpeechAdjective - An adjective, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/adjective
	GrammaticalPartOfSpeechAdjective GrammaticalPartOfSpeech = 6
	// GrammaticalPartOfSpeechAdposition - An adposition, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/adposition
	GrammaticalPartOfSpeechAdposition GrammaticalPartOfSpeech = 7
	// GrammaticalPartOfSpeechAdverb - An adverb, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/adverb
	GrammaticalPartOfSpeechAdverb GrammaticalPartOfSpeech = 4
	// GrammaticalPartOfSpeechConjunction - A conjunction, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/conjunction
	GrammaticalPartOfSpeechConjunction GrammaticalPartOfSpeech = 10
	// GrammaticalPartOfSpeechDeterminer - A determiner, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/determiner
	GrammaticalPartOfSpeechDeterminer GrammaticalPartOfSpeech = 1
	// GrammaticalPartOfSpeechInterjection - An interjection, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/interjection
	GrammaticalPartOfSpeechInterjection GrammaticalPartOfSpeech = 12
	// GrammaticalPartOfSpeechLetter - A letter, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/letter
	GrammaticalPartOfSpeechLetter GrammaticalPartOfSpeech = 3
	// GrammaticalPartOfSpeechNotSet - A value that indicates the part of speech is unset.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/notSet
	GrammaticalPartOfSpeechNotSet GrammaticalPartOfSpeech = 0
	// GrammaticalPartOfSpeechNoun - A noun, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/noun
	GrammaticalPartOfSpeechNoun GrammaticalPartOfSpeech = 9
	// GrammaticalPartOfSpeechNumeral - A numeral, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/numeral
	GrammaticalPartOfSpeechNumeral GrammaticalPartOfSpeech = 11
	// GrammaticalPartOfSpeechParticle - A particle, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/particle
	GrammaticalPartOfSpeechParticle GrammaticalPartOfSpeech = 5
	// GrammaticalPartOfSpeechPreposition - A preposition, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/preposition
	GrammaticalPartOfSpeechPreposition GrammaticalPartOfSpeech = 13
	// GrammaticalPartOfSpeechPronoun - A pronoun, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/pronoun
	GrammaticalPartOfSpeechPronoun GrammaticalPartOfSpeech = 2
	// GrammaticalPartOfSpeechVerb - A verb, as used as a part of speech.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech/verb
	GrammaticalPartOfSpeechVerb GrammaticalPartOfSpeech = 8
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
	// LocaleLanguageDirectionUnknown - The direction of the language is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection/unknown
	LocaleLanguageDirectionUnknown LocaleLanguageDirection = 0
)

// MachPortOptions - Used to remove access rights to a mach port when the 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/Options
type MachPortOptions uint

const (
	MachPortDeallocateNone MachPortOptions = 0
	MachPortDeallocateSendRight MachPortOptions = 1
	MachPortDeallocateReceiveRight MachPortOptions = 2
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
	// MatchingWithTransparentBounds - Specifies that matching may examine parts of the string beyond the bounds of the search range, for purposes such as word boundary detection, lookahead, etc. This constant has no effect if the search range contains the entire string. See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/withTransparentBounds
	MatchingWithTransparentBounds MatchingOptions = 8
	// MatchingWithoutAnchoringBounds - Specifies that   and   will not automatically match the beginning and end of the search range, but will still match the beginning and end of the entire string. This constant has no effect if the search range contains the entire string. See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/withoutAnchoringBounds
	MatchingWithoutAnchoringBounds MatchingOptions = 16
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

// TestComparisonOperation - These are passed to  
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSpecifierTest/TestComparisonOperation
type TestComparisonOperation uint

const (
	// BeginsWithComparison - Binary containment operator that results in true if the test object is a list or string that matches the beginning of the other object (which is also a list or string).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSpecifierTest/TestComparisonOperation/beginsWith
	BeginsWithComparison TestComparisonOperation = 5
	// GreaterThanComparison - Binary comparison operator that results in true if the value of the test object is greater than the value of the other object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSpecifierTest/TestComparisonOperation/greaterThan
	GreaterThanComparison TestComparisonOperation = 4
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

// URLSessionWebSocketMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType
type URLSessionWebSocketMessageType int

const (
	URLSessionWebSocketMessageTypeData URLSessionWebSocketMessageType = 0
	URLSessionWebSocketMessageTypeString URLSessionWebSocketMessageType = 1
)

// XPCConnectionOptions - Options that you can pass to a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/Options
type XPCConnectionOptions uint

const (
	// XPCConnectionPrivileged - Use this option if connecting to a service in the privileged Mach bootstrap (for example, a daemon with a   in  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/Options/privileged
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
	// NotificationCoalescingOnName - Coalesce notifications with the same name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing/onName
	NotificationCoalescingOnName NotificationCoalescing = 1
	// NotificationCoalescingOnSender - Coalesce notifications with the same object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing/onSender
	NotificationCoalescingOnSender NotificationCoalescing = 2
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
	// PersonNameComponentsFormatterPhonetic - The formatter should format the component object’s   components instead of its own components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Options/phonetic
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
	// TaskTerminationReasonExit - The task exited normally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/TerminationReason-swift.enum/exit
	TaskTerminationReasonExit TaskTerminationReason = 1
	// TaskTerminationReasonUncaughtSignal - The task exited due to an uncaught signal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/TerminationReason-swift.enum/uncaughtSignal
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
	// ActivityAutomaticTerminationDisabled - A flag to prevent automatic termination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/automaticTerminationDisabled
	ActivityAutomaticTerminationDisabled ActivityOptions = 32768
	// ActivityBackground - A flag to indicate the app has initiated some kind of work, but not as the direct result of user request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/background
	ActivityBackground ActivityOptions = 0
	// ActivityIdleDisplaySleepDisabled - A flag to require the screen to stay powered on.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/idleDisplaySleepDisabled
	ActivityIdleDisplaySleepDisabled ActivityOptions = 1099511627776
	// ActivityIdleSystemSleepDisabled - A flag to prevent idle sleep.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/idleSystemSleepDisabled
	ActivityIdleSystemSleepDisabled ActivityOptions = 1048576
	// ActivityLatencyCritical - A flag to indicate the activity requires the highest amount of timer and I/O precision available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/latencyCritical
	ActivityLatencyCritical ActivityOptions = 0
	// ActivitySuddenTerminationDisabled - A flag to prevent sudden termination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/suddenTerminationDisabled
	ActivitySuddenTerminationDisabled ActivityOptions = 16384
	// ActivityTrackingEnabled - A flag to track the activity with a signpost interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions/trackingEnabled
	ActivityTrackingEnabled ActivityOptions = 32770
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

// PropertyListMutabilityOptions - These constants specify mutability options in property lists.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/MutabilityOptions
type PropertyListMutabilityOptions uint

const (
	// PropertyListImmutable - Causes the returned property list to contain immutable objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPropertyListMutabilityOptions/NSPropertyListImmutable
	PropertyListImmutable PropertyListMutabilityOptions = 0
	// PropertyListMutableContainers - Causes the returned property list to have mutable containers but immutable leaves.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/MutabilityOptions/mutableContainers
	PropertyListMutableContainers PropertyListMutabilityOptions = 0
	// PropertyListMutableContainersAndLeaves - Causes the returned property list to have mutable containers and leaves.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/MutabilityOptions/mutableContainersAndLeaves
	PropertyListMutableContainersAndLeaves PropertyListMutabilityOptions = 0
)

// PropertyListFormat - These constants are used to specify a property list serialization format.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat
type PropertyListFormat uint

const (
	// PropertyListBinaryFormat_v1_0 - Specifies the binary property list format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat/binary
	PropertyListBinaryFormat_v1_0 PropertyListFormat = 0
	// PropertyListOpenStepFormat - Specifies the ASCII property list format inherited from the OpenStep APIs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat/openStep
	PropertyListOpenStepFormat PropertyListFormat = 0
	// PropertyListXMLFormat_v1_0 - Specifies the XML property list format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat/xml
	PropertyListXMLFormat_v1_0 PropertyListFormat = 0
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
	// RelativeDateTimeFormatterStyleNamed - A style that uses named styles to describe relative dates, such as “yesterday”, “last week”, or “next week”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/DateTimeStyle-swift.enum/named
	RelativeDateTimeFormatterStyleNamed RelativeDateTimeFormatterStyle = 1
	// RelativeDateTimeFormatterStyleNumeric - A style that uses a numeric style to describe relative dates, such as “1 day ago” or “in 3 weeks”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/DateTimeStyle-swift.enum/numeric
	RelativeDateTimeFormatterStyleNumeric RelativeDateTimeFormatterStyle = 0
)

// RelativeDateTimeFormatterUnitsStyle - A type that represents the style to use when formatting the units of relative dates.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/UnitsStyle-swift.enum
type RelativeDateTimeFormatterUnitsStyle uint

const (
	// RelativeDateTimeFormatterUnitsStyleAbbreviated - A style that uses abbreviated units, such as “2 mo. ago”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/UnitsStyle-swift.enum/abbreviated
	RelativeDateTimeFormatterUnitsStyleAbbreviated RelativeDateTimeFormatterUnitsStyle = 3
	// RelativeDateTimeFormatterUnitsStyleFull - A style that uses full units, such as “2 months ago”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/UnitsStyle-swift.enum/full
	RelativeDateTimeFormatterUnitsStyleFull RelativeDateTimeFormatterUnitsStyle = 0
	// RelativeDateTimeFormatterUnitsStyleShort - A style that uses shortened units, such as “2 mo. ago”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/UnitsStyle-swift.enum/short
	RelativeDateTimeFormatterUnitsStyleShort RelativeDateTimeFormatterUnitsStyle = 2
	// RelativeDateTimeFormatterUnitsStyleSpellOut - A style that spells out units such as “two months ago”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/UnitsStyle-swift.enum/spellOut
	RelativeDateTimeFormatterUnitsStyleSpellOut RelativeDateTimeFormatterUnitsStyle = 1
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

// URLSessionWebSocketCloseCode - A code that indicates why a WebSocket connection closed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/CloseCode-swift.enum
type URLSessionWebSocketCloseCode uint

const (
	URLSessionWebSocketCloseCodeInvalid URLSessionWebSocketCloseCode = 0
	URLSessionWebSocketCloseCodeNormalClosure URLSessionWebSocketCloseCode = 1000
	URLSessionWebSocketCloseCodeGoingAway URLSessionWebSocketCloseCode = 1001
	URLSessionWebSocketCloseCodeProtocolError URLSessionWebSocketCloseCode = 1002
	URLSessionWebSocketCloseCodeUnsupportedData URLSessionWebSocketCloseCode = 1003
	URLSessionWebSocketCloseCodeNoStatusReceived URLSessionWebSocketCloseCode = 1005
	URLSessionWebSocketCloseCodeAbnormalClosure URLSessionWebSocketCloseCode = 1006
	URLSessionWebSocketCloseCodeInvalidFramePayloadData URLSessionWebSocketCloseCode = 1007
	URLSessionWebSocketCloseCodePolicyViolation URLSessionWebSocketCloseCode = 1008
	URLSessionWebSocketCloseCodeMessageTooBig URLSessionWebSocketCloseCode = 1009
	URLSessionWebSocketCloseCodeMandatoryExtensionMissing URLSessionWebSocketCloseCode = 1010
	URLSessionWebSocketCloseCodeInternalServerError URLSessionWebSocketCloseCode = 1011
	URLSessionWebSocketCloseCodeTLSHandshakeFailure URLSessionWebSocketCloseCode = 1015
)

// XMLDTDNodeKind - The type defined for the constants that specify the kind and subkind of DTD declaration represented by an 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum
type XMLDTDNodeKind uint

const (
	// XMLEntityParameterKind - Identifies a parameter entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/parameter
	XMLEntityParameterKind XMLDTDNodeKind = 4
	// XMLEntityUnparsedKind - Identifies an unparsed entity declaration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum/unparsed
	XMLEntityUnparsedKind XMLDTDNodeKind = 3
)

// XMLNodeOptions - These constants are input and output options for all 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/Options
type XMLNodeOptions uint

const (
	// XMLNodeOptionsNone - No options are requested for this input or output action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXMLNodeOptions/NSXMLNodeOptionsNone
	XMLNodeOptionsNone XMLNodeOptions = 0
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


