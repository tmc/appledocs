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
	IsDaylightSavingTimeForDate(aDate unsafe.Pointer) bool
	IsEqualToTimeZone(aTimeZone unsafe.Pointer) bool
	NextDaylightSavingTimeTransitionAfterDate(aDate unsafe.Pointer) unsafe.Pointer
}

// Information about standard time conventions associated with a specific geopolitical region.
//
// In Swift, this type bridges to ; use when you need reference semantics or other Foundation-specific behavior. Time zones represent the standard time policies for a geopolitical region. Time zones have identifiers like “America/Los_Angeles” and can also be identified by abbreviations, such as PST for Pacific Standard Time. You can create time zone objects by ID with and by abbreviation with . Time zones can also represent a temporal offset—either plus or minus—from Greenwich Mean Time (GMT). For example, the temporal offset of Pacific Standard Time is 8 hours behind Greenwich Mean Time (GMT-8). You can create time zone objects with a temporal offset by using . You typically work with system time zones rather than creating time zones by identifier or by offset. The class property returns the time zone currently used by the system, if known. This value is cached once the property is accessed and doesn’t reflect any system time zone changes until you call the method. The class property returns an autoupdating proxy object that always returns the current time zone used by the system. You can also set the class property to make your app run as if it were in a different time zone than the system. is with its Core Foundation counterpart, . See for more information on toll-free bridging.
//
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




// Returns a time zone object offset from Greenwich Mean Time by a given number of seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(forSecondsFromGMT:)
func NewTimeZoneForSecondsFromGMT(seconds int) TimeZone {
	rv := objc.Send[TimeZone](objc.ID(getTimeZoneClass().class), objc.Sel("timeZoneForSecondsFromGMT:"), seconds)
	return rv
}



// Returns the time zone object identified by a given abbreviation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(abbreviation:)
func NewTimeZoneWithAbbreviation(abbreviation string) TimeZone {
	rv := objc.Send[TimeZone](objc.ID(getTimeZoneClass().class), objc.Sel("timeZoneWithAbbreviation:"), objc.String(abbreviation))
	return rv
}



// Returns a time zone initialized with a given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(name:)
func NewTimeZoneWithName(tzName string) TimeZone {
	instance := getTimeZoneClass().Alloc()
	rv := objc.Send[TimeZone](instance.ID, objc.Sel("initWithName:"), objc.String(tzName))
	rv.Autorelease()
	return rv
}


// Returns the time zone object identified by a given abbreviation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(abbreviation:)
func (tc _TimeZoneClass) TimeZoneWithAbbreviation(abbreviation string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("timeZoneWithAbbreviation:"), objc.String(abbreviation))
	return rv
}

// Returns a time zone object offset from Greenwich Mean Time by a given number of seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(forSecondsFromGMT:)
func (tc _TimeZoneClass) TimeZoneForSecondsFromGMT(seconds int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("timeZoneForSecondsFromGMT:"), seconds)
	return rv
}

// Clears any time zone value cached for the property.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/resetSystemTimeZone()
func (tc _TimeZoneClass) ResetSystemTimeZone() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("resetSystemTimeZone"))
}

// Returns the time zone object identified by a given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneWithName:
func (tc _TimeZoneClass) TimeZoneWithName(tzName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("timeZoneWithName:"), objc.String(tzName))
	return rv
}

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviationDictionary
func (tc _TimeZoneClass) AbbreviationDictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("abbreviationDictionary"))
	return rv
}
// The default time zone for the current app.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/default
func (tc _TimeZoneClass) DefaultTimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("defaultTimeZone"))
	return rv
}
// An object that tracks the current system time zone.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/local
func (tc _TimeZoneClass) LocalTimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("localTimeZone"))
	return rv
}
// The time zone currently used by the system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/system
func (tc _TimeZoneClass) SystemTimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("systemTimeZone"))
	return rv
}
// Indicates whether the receiver uses daylight saving time on a given date.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/isDaylightSavingTime(for:)
func (t_ TimeZone) IsDaylightSavingTimeForDate(aDate unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isDaylightSavingTimeForDate:"), aDate)
	return rv
}

// Indicates whether the receiver has the same name and data as the specified time zone.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/isEqual(to:)
func (t_ TimeZone) IsEqualToTimeZone(aTimeZone unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEqualToTimeZone:"), aTimeZone)
	return rv
}

// Returns the next daylight saving time transition after a given date.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/nextDaylightSavingTimeTransition(after:)
func (t_ TimeZone) NextDaylightSavingTimeTransitionAfterDate(aDate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("nextDaylightSavingTimeTransitionAfterDate:"), aDate)
	return rv
}

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviationDictionary
func (t_ TimeZone) AbbreviationDictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("abbreviationDictionary"))
	return rv
}


// SetAbbreviationDictionary sets the value of the abbreviationDictionary property.
// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviationDictionary
func (t_ TimeZone) SetAbbreviationDictionary(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAbbreviationDictionary:"), value)
}

// The data that stores the information used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/data
func (t_ TimeZone) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("data"))
	return rv
}

// The default time zone for the current app.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/default
func (t_ TimeZone) DefaultTimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("defaultTimeZone"))
	return rv
}


// SetDefaultTimeZone sets the value of the defaultTimeZone property.
// The default time zone for the current app.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/default
func (t_ TimeZone) SetDefaultTimeZone(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultTimeZone:"), value)
}

// A textual description of the time zone including the name, abbreviation, offset from GMT, and whether or not daylight saving time is currently in effect.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/description
func (t_ TimeZone) Description() string {
	rv := objc.Send[string](t_.ID, objc.Sel("description"))
	return rv
}

// An object that tracks the current system time zone.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/local
func (t_ TimeZone) LocalTimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("localTimeZone"))
	return rv
}

// The geopolitical region ID that identifies the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/name
func (t_ TimeZone) Name() string {
	rv := objc.Send[string](t_.ID, objc.Sel("name"))
	return rv
}

// The time zone currently used by the system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/system
func (t_ TimeZone) SystemTimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("systemTimeZone"))
	return rv
}

// The abbreviation for the receiver, such as “EDT” (Eastern Daylight Time).
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/abbreviation
func (t_ TimeZone) Abbreviation() string {
	rv := objc.Send[string](t_.ID, objc.Sel("abbreviation"))
	return rv
}


// SetAbbreviation sets the value of the abbreviation property.
// The abbreviation for the receiver, such as “EDT” (Eastern Daylight Time).

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/abbreviation
func (t_ TimeZone) SetAbbreviation(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAbbreviation:"), objc.String(value))
}

// The current daylight saving time offset of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/daylightsavingtimeoffset
func (t_ TimeZone) DaylightSavingTimeOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("daylightSavingTimeOffset"))
	return rv
}


// SetDaylightSavingTimeOffset sets the value of the daylightSavingTimeOffset property.
// The current daylight saving time offset of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/daylightsavingtimeoffset
func (t_ TimeZone) SetDaylightSavingTimeOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDaylightSavingTimeOffset:"), value)
}

// A Boolean value that indicates whether the receiver is currently using daylight saving time.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/isdaylightsavingtime
func (t_ TimeZone) IsDaylightSavingTime() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isDaylightSavingTime"))
	return rv
}


// SetIsDaylightSavingTime sets the value of the isDaylightSavingTime property.
// A Boolean value that indicates whether the receiver is currently using daylight saving time.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/isdaylightsavingtime
func (t_ TimeZone) SetIsDaylightSavingTime(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsDaylightSavingTime:"), value)
}

// The date of the next daylight saving time transition for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/nextdaylightsavingtimetransition
func (t_ TimeZone) NextDaylightSavingTimeTransition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("nextDaylightSavingTimeTransition"))
	return rv
}


// SetNextDaylightSavingTimeTransition sets the value of the nextDaylightSavingTimeTransition property.
// The date of the next daylight saving time transition for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/nextdaylightsavingtimetransition
func (t_ TimeZone) SetNextDaylightSavingTimeTransition(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNextDaylightSavingTimeTransition:"), value)
}

// The current difference in seconds between the receiver and Greenwich Mean Time.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/secondsfromgmt
func (t_ TimeZone) SecondsFromGMT() int {
	rv := objc.Send[int](t_.ID, objc.Sel("secondsFromGMT"))
	return rv
}


// SetSecondsFromGMT sets the value of the secondsFromGMT property.
// The current difference in seconds between the receiver and Greenwich Mean Time.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/secondsfromgmt
func (t_ TimeZone) SetSecondsFromGMT(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSecondsFromGMT:"), value)
}


