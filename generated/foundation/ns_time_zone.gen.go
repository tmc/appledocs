// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TimeZone] class.
var (
	TimeZoneClass     _TimeZoneClass
	TimeZoneClassOnce sync.Once
)

func getTimeZoneClass() _TimeZoneClass {
	TimeZoneClassOnce.Do(func() {
		TimeZoneClass = _TimeZoneClass{objc.GetClass("NSTimeZone")}
	})
	return TimeZoneClass
}

type _TimeZoneClass struct {
	class objc.Class
}

// An interface definition for the [TimeZone] class.
type ITimeZone interface {
	objectivec.IObject
	Abbreviation() string
	Data() IData
	Name() string
	SecondsFromGMT() int
	DaylightSavingTimeOffset() TimeInterval
	SetDaylightSavingTimeOffset(value TimeInterval)
	Description() string
	SetDescription(value string)
	IsDaylightSavingTime() bool
	SetIsDaylightSavingTime(value bool)
	NextDaylightSavingTimeTransition() IDate
	SetNextDaylightSavingTimeTransition(value IDate)
	AbbreviationForDate(aDate IDate) IString
	SecondsFromGMTForDate(aDate IDate) int
}

// Information about standard time conventions associated with a specific geopolitical region.
//
// In Swift, this type bridges to ; use when you need reference semantics or other Foundation-specific behavior. Time zones represent the standard time policies for a geopolitical region. Time zones have identifiers like “America/Los_Angeles” and can also be identified by abbreviations, such as PST for Pacific Standard Time. You can create time zone objects by ID with and by abbreviation with . Time zones can also represent a temporal offset—either plus or minus—from Greenwich Mean Time (GMT). For example, the temporal offset of Pacific Standard Time is 8 hours behind Greenwich Mean Time (GMT-8). You can create time zone objects with a temporal offset by using . You typically work with system time zones rather than creating time zones by identifier or by offset. The class property returns the time zone currently used by the system, if known. This value is cached once the property is accessed and doesn’t reflect any system time zone changes until you call the method. The class property returns an autoupdating proxy object that always returns the current time zone used by the system. You can also set the class property to make your app run as if it were in a different time zone than the system. is with its Core Foundation counterpart, . See for more information on toll-free bridging.


// Information about standard time conventions associated with a specific geopolitical region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone
type TimeZone struct {
	objectivec.Object
}

// TimeZoneFrom constructs a [TimeZone] from an unsafe.Pointer.
//
// Information about standard time conventions associated with a specific geopolitical region.
func TimeZoneFrom(ptr unsafe.Pointer) TimeZone {
	return TimeZone{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TimeZoneClass) Alloc() TimeZone {
	rv := objc.Send[TimeZone](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TimeZoneClass) New() TimeZone {
	rv := objc.Send[TimeZone](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TimeZone) Init() TimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TimeZone) Autorelease() TimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTimeZone creates a new TimeZone instance.
func NewTimeZone() TimeZone {
	return getTimeZoneClass().New()
}



// Returns the time zone data version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneDataVersion
func (tc _TimeZoneClass) TimeZoneDataVersion() string {
	rv := objc.Send[string](objc.ID(tc.class), objc.Sel("timeZoneDataVersion"))
	return rv
}

// Returns the abbreviation for the receiver at a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviation(for:)
func (t_ TimeZone) AbbreviationForDate(aDate IDate) IString {
	rv := objc.Send[String](t_.ID, objc.Sel("abbreviationForDate:"), aDate)
	return rv
}


// Returns the difference in seconds between the receiver and Greenwich Mean Time at a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/secondsFromGMT(for:)
func (t_ TimeZone) SecondsFromGMTForDate(aDate IDate) int {
	rv := objc.Send[int](t_.ID, objc.Sel("secondsFromGMTForDate:"), aDate)
	return rv
}


// The abbreviation for the receiver, such as “EDT” (Eastern Daylight Time).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviation
func (t_ TimeZone) Abbreviation() string {
	rv := objc.Send[string](t_.ID, objc.Sel("abbreviation"))
	return rv
}


// The data that stores the information used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/data
func (t_ TimeZone) Data() IData {
	rv := objc.Send[Data](t_.ID, objc.Sel("data"))
	return rv
}


// The geopolitical region ID that identifies the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/name
func (t_ TimeZone) Name() string {
	rv := objc.Send[string](t_.ID, objc.Sel("name"))
	return rv
}


// The current difference in seconds between the receiver and Greenwich Mean Time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/secondsFromGMT
func (t_ TimeZone) SecondsFromGMT() int {
	rv := objc.Send[int](t_.ID, objc.Sel("secondsFromGMT"))
	return rv
}


// Returns the time zone data version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneDataVersion
func (t_ TimeZone) TimeZoneDataVersion() string {
	rv := objc.Send[string](t_.ID, objc.Sel("timeZoneDataVersion"))
	return rv
}


// The current daylight saving time offset of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/daylightsavingtimeoffset
func (t_ TimeZone) DaylightSavingTimeOffset() TimeInterval {
	rv := objc.Send[TimeInterval](t_.ID, objc.Sel("daylightSavingTimeOffset"))
	return rv
}


// The current daylight saving time offset of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/daylightsavingtimeoffset
func (t_ TimeZone) SetDaylightSavingTimeOffset(value TimeInterval) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDaylightSavingTimeOffset:"), value)
}


// A textual description of the time zone including the name, abbreviation, offset from GMT, and whether or not daylight saving time is currently in effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/description
func (t_ TimeZone) Description() string {
	rv := objc.Send[string](t_.ID, objc.Sel("description"))
	return rv
}


// A textual description of the time zone including the name, abbreviation, offset from GMT, and whether or not daylight saving time is currently in effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/description
func (t_ TimeZone) SetDescription(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDescription:"), objc.String(value))
}


// A Boolean value that indicates whether the receiver is currently using daylight saving time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/isdaylightsavingtime
func (t_ TimeZone) IsDaylightSavingTime() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isDaylightSavingTime"))
	return rv
}


// A Boolean value that indicates whether the receiver is currently using daylight saving time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/isdaylightsavingtime
func (t_ TimeZone) SetIsDaylightSavingTime(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsDaylightSavingTime:"), value)
}


// The date of the next daylight saving time transition for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/nextdaylightsavingtimetransition
func (t_ TimeZone) NextDaylightSavingTimeTransition() IDate {
	rv := objc.Send[Date](t_.ID, objc.Sel("nextDaylightSavingTimeTransition"))
	return rv
}


// The date of the next daylight saving time transition for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/nextdaylightsavingtimetransition
func (t_ TimeZone) SetNextDaylightSavingTimeTransition(value IDate) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNextDaylightSavingTimeTransition:"), value)
}



