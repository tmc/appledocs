// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTimeZone */


/* debug [class_header]: Header for NSTimeZone */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TimeZone */
// An interface definition for the [TimeZone] class.
type ITimeZone interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TimeZone */
	// properties:
	Abbreviation() IString
	Data() IData
	DaylightSavingTimeOffset() float64
	Description() IString
	DaylightSavingTime() bool
	Name() IString
	NextDaylightSavingTimeTransition() IDate
	SecondsFromGMT() int
	IsDaylightSavingTime() bool
	SetIsDaylightSavingTime(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TimeZone */
	// methods:
	AbbreviationForDate(aDate IDate) IString
	DaylightSavingTimeOffsetForDate(aDate IDate) float64
	IsDaylightSavingTimeForDate(aDate IDate) bool
	IsEqualToTimeZone(aTimeZone ITimeZone) bool
	LocalizedNameLocale(style TimeZoneNameStyle, locale ILocale) IString
	NextDaylightSavingTimeTransitionAfterDate(aDate IDate) IDate
	SecondsFromGMTForDate(aDate IDate) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TimeZone */
// Alloc allocates a new instance without initialization.
func (tc _TimeZoneClass) Alloc() TimeZone {
	rv := objc.Send[TimeZone](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TimeZone */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TimeZone */

// Returns a time zone object offset from Greenwich Mean Time by a given number of seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(forSecondsFromGMT:)
func NewTimeZoneForSecondsFromGMT(seconds int) TimeZone {
	rv := objc.Send[TimeZone](objc.ID(getTimeZoneClass().class), objc.Sel("timeZoneForSecondsFromGMT:"), seconds)
	return rv
}/* debug [class_init_methods/constructor]: NewTimeZoneForSecondsFromGMT */


// Returns the time zone object identified by a given abbreviation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(abbreviation:)
func NewTimeZoneWithAbbreviation(abbreviation IString) TimeZone {
	rv := objc.Send[TimeZone](objc.ID(getTimeZoneClass().class), objc.Sel("timeZoneWithAbbreviation:"), abbreviation)
	return rv
}/* debug [class_init_methods/constructor]: NewTimeZoneWithAbbreviation */


// Returns a time zone initialized with a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(name:)
func NewTimeZoneWithName(tzName IString) TimeZone {
	instance := getTimeZoneClass().Alloc()
	rv := objc.Send[TimeZone](instance.ID, objc.Sel("initWithName:"), tzName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTimeZoneWithName */


// Initializes a time zone with a given identifier and time zone data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(name:data:)
func NewTimeZoneWithNameData(tzName IString, aData IData) TimeZone {
	instance := getTimeZoneClass().Alloc()
	rv := objc.Send[TimeZone](instance.ID, objc.Sel("initWithName:data:"), tzName, aData)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTimeZoneWithNameData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TimeZone */

// Returns the time zone object identified by a given abbreviation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(abbreviation:)
func (tc _TimeZoneClass) TimeZoneWithAbbreviation(abbreviation IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("timeZoneWithAbbreviation:"), abbreviation)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimeZoneWithAbbreviation) */


// Returns a time zone object offset from Greenwich Mean Time by a given number of seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(forSecondsFromGMT:)
func (tc _TimeZoneClass) TimeZoneForSecondsFromGMT(seconds int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("timeZoneForSecondsFromGMT:"), seconds)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimeZoneForSecondsFromGMT) */


// Clears any time zone value cached for the property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/resetSystemTimeZone()
func (tc _TimeZoneClass) ResetSystemTimeZone() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("resetSystemTimeZone"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ResetSystemTimeZone) */


// Returns the time zone object identified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneWithName:
func (tc _TimeZoneClass) TimeZoneWithName(tzName IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("timeZoneWithName:"), tzName)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimeZoneWithName) */


// Returns the time zone with a given identifier whose data has been initialized using given data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneWithName:data:
func (tc _TimeZoneClass) TimeZoneWithNameData(tzName IString, aData IData) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("timeZoneWithName:data:"), tzName, aData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimeZoneWithNameData) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TimeZone */

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviationDictionary
func (tc _TimeZoneClass) AbbreviationDictionary() IDictionary {
	rv := objc.Send[Dictionary](objc.ID(tc.class), objc.Sel("abbreviationDictionary"))
	return rv
}/* debug [class_properties_class/property]: abbreviationDictionary */

// The default time zone for the current app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/default
func (tc _TimeZoneClass) DefaultTimeZone() TimeZone {
	rv := objc.Send[TimeZone](objc.ID(tc.class), objc.Sel("defaultTimeZone"))
	return rv
}/* debug [class_properties_class/property]: defaultTimeZone */

// Returns an array of strings listing the IDs of all the time zones known to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/knownTimeZoneNames
func (tc _TimeZoneClass) KnownTimeZoneNames() []string {
	rv := objc.Send[[]string](objc.ID(tc.class), objc.Sel("knownTimeZoneNames"))
	return rv
}/* debug [class_properties_class/property]: knownTimeZoneNames */

// An object that tracks the current system time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/local
func (tc _TimeZoneClass) LocalTimeZone() TimeZone {
	rv := objc.Send[TimeZone](objc.ID(tc.class), objc.Sel("localTimeZone"))
	return rv
}/* debug [class_properties_class/property]: localTimeZone */

// The time zone currently used by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/system
func (tc _TimeZoneClass) SystemTimeZone() TimeZone {
	rv := objc.Send[TimeZone](objc.ID(tc.class), objc.Sel("systemTimeZone"))
	return rv
}/* debug [class_properties_class/property]: systemTimeZone */

// Returns the time zone data version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneDataVersion
func (tc _TimeZoneClass) TimeZoneDataVersion() IString {
	rv := objc.Send[String](objc.ID(tc.class), objc.Sel("timeZoneDataVersion"))
	return rv
}/* debug [class_properties_class/property]: timeZoneDataVersion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TimeZone */

// Returns the abbreviation for the receiver at a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviation(for:)
func (t_ TimeZone) AbbreviationForDate(aDate IDate) IString {
	rv := objc.Send[String](t_.ID, objc.Sel("abbreviationForDate:"), aDate)
	return rv
}/* debug [instance_methods/method]: AbbreviationForDate */


// Returns the daylight saving time offset for a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/daylightSavingTimeOffset(for:)
func (t_ TimeZone) DaylightSavingTimeOffsetForDate(aDate IDate) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("daylightSavingTimeOffsetForDate:"), aDate)
	return rv
}/* debug [instance_methods/method]: DaylightSavingTimeOffsetForDate */


// Indicates whether the receiver uses daylight saving time on a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/isDaylightSavingTime(for:)
func (t_ TimeZone) IsDaylightSavingTimeForDate(aDate IDate) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isDaylightSavingTimeForDate:"), aDate)
	return rv
}/* debug [instance_methods/method]: IsDaylightSavingTimeForDate */


// Indicates whether the receiver has the same name and data as the specified time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/isEqual(to:)
func (t_ TimeZone) IsEqualToTimeZone(aTimeZone ITimeZone) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEqualToTimeZone:"), aTimeZone)
	return rv
}/* debug [instance_methods/method]: IsEqualToTimeZone */


// Returns the localized name of the time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/localizedName(_:locale:)
func (t_ TimeZone) LocalizedNameLocale(style TimeZoneNameStyle, locale ILocale) IString {
	rv := objc.Send[String](t_.ID, objc.Sel("localizedName:locale:"), style, locale)
	return rv
}/* debug [instance_methods/method]: LocalizedNameLocale */


// Returns the next daylight saving time transition after a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/nextDaylightSavingTimeTransition(after:)
func (t_ TimeZone) NextDaylightSavingTimeTransitionAfterDate(aDate IDate) IDate {
	rv := objc.Send[Date](t_.ID, objc.Sel("nextDaylightSavingTimeTransitionAfterDate:"), aDate)
	return rv
}/* debug [instance_methods/method]: NextDaylightSavingTimeTransitionAfterDate */


// Returns the difference in seconds between the receiver and Greenwich Mean Time at a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/secondsFromGMT(for:)
func (t_ TimeZone) SecondsFromGMTForDate(aDate IDate) int {
	rv := objc.Send[int](t_.ID, objc.Sel("secondsFromGMTForDate:"), aDate)
	return rv
}/* debug [instance_methods/method]: SecondsFromGMTForDate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TimeZone */

// The abbreviation for the receiver, such as “EDT” (Eastern Daylight Time).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviation
func (t_ TimeZone) Abbreviation() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("abbreviation"))
	return rv
}/* debug [instance_properties/getter]: abbreviation */


// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviationDictionary
func (t_ TimeZone) AbbreviationDictionary() IDictionary {
	rv := objc.Send[Dictionary](t_.ID, objc.Sel("abbreviationDictionary"))
	return rv
}/* debug [instance_properties/getter]: abbreviationDictionary */


// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviationDictionary
func (t_ TimeZone) SetAbbreviationDictionary(value IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAbbreviationDictionary:"), value)
}/* debug [instance_properties/setter]: abbreviationDictionary */


// The data that stores the information used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/data
func (t_ TimeZone) Data() IData {
	rv := objc.Send[Data](t_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The current daylight saving time offset of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/daylightSavingTimeOffset
func (t_ TimeZone) DaylightSavingTimeOffset() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("daylightSavingTimeOffset"))
	return rv
}/* debug [instance_properties/getter]: daylightSavingTimeOffset */


// The default time zone for the current app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/default
func (t_ TimeZone) DefaultTimeZone() ITimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("defaultTimeZone"))
	return rv
}/* debug [instance_properties/getter]: defaultTimeZone */


// The default time zone for the current app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/default
func (t_ TimeZone) SetDefaultTimeZone(value ITimeZone) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultTimeZone:"), value)
}/* debug [instance_properties/setter]: defaultTimeZone */


// A textual description of the time zone including the name, abbreviation, offset from GMT, and whether or not daylight saving time is currently in effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/description
func (t_ TimeZone) Description() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("description"))
	return rv
}/* debug [instance_properties/getter]: description */


// A Boolean value that indicates whether the receiver is currently using daylight saving time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/isDaylightSavingTime
func (t_ TimeZone) DaylightSavingTime() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("daylightSavingTime"))
	return rv
}/* debug [instance_properties/getter]: daylightSavingTime */


// Returns an array of strings listing the IDs of all the time zones known to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/knownTimeZoneNames
func (t_ TimeZone) KnownTimeZoneNames() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("knownTimeZoneNames"))
	return rv
}/* debug [instance_properties/getter]: knownTimeZoneNames */


// An object that tracks the current system time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/local
func (t_ TimeZone) LocalTimeZone() ITimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("localTimeZone"))
	return rv
}/* debug [instance_properties/getter]: localTimeZone */


// The geopolitical region ID that identifies the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/name
func (t_ TimeZone) Name() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The date of the next daylight saving time transition for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/nextDaylightSavingTimeTransition
func (t_ TimeZone) NextDaylightSavingTimeTransition() IDate {
	rv := objc.Send[Date](t_.ID, objc.Sel("nextDaylightSavingTimeTransition"))
	return rv
}/* debug [instance_properties/getter]: nextDaylightSavingTimeTransition */


// The current difference in seconds between the receiver and Greenwich Mean Time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/secondsFromGMT
func (t_ TimeZone) SecondsFromGMT() int {
	rv := objc.Send[int](t_.ID, objc.Sel("secondsFromGMT"))
	return rv
}/* debug [instance_properties/getter]: secondsFromGMT */


// The time zone currently used by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/system
func (t_ TimeZone) SystemTimeZone() ITimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("systemTimeZone"))
	return rv
}/* debug [instance_properties/getter]: systemTimeZone */


// Returns the time zone data version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneDataVersion
func (t_ TimeZone) TimeZoneDataVersion() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("timeZoneDataVersion"))
	return rv
}/* debug [instance_properties/getter]: timeZoneDataVersion */


// A Boolean value that indicates whether the receiver is currently using daylight saving time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/isdaylightsavingtime
func (t_ TimeZone) IsDaylightSavingTime() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isDaylightSavingTime"))
	return rv
}/* debug [instance_properties/getter]: isDaylightSavingTime */


// A Boolean value that indicates whether the receiver is currently using daylight saving time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/isdaylightsavingtime
func (t_ TimeZone) SetIsDaylightSavingTime(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsDaylightSavingTime:"), value)
}/* debug [instance_properties/setter]: isDaylightSavingTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTimeZone */


