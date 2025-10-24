// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Date] class.
var (
	DateClass     _DateClass
	DateClassOnce sync.Once
)

func getDateClass() _DateClass {
	DateClassOnce.Do(func() {
		DateClass = _DateClass{objc.GetClass("NSDate")}
	})
	return DateClass
}

type _DateClass struct {
	class objc.Class
}

// An interface definition for the [Date] class.
type IDate interface {
	objectivec.IObject
	// properties:
	Description() IString
	SrAbsoluteTime() unsafe.Pointer
	TimeIntervalSince1970() objc.IObject /* cross-framework: TimeInterval */
	TimeIntervalSinceNow() objc.IObject /* cross-framework: TimeInterval */
	TimeIntervalSinceReferenceDate() objc.IObject /* cross-framework: TimeInterval */
	CustomPlaygroundQuickLook() unsafe.Pointer
	SetCustomPlaygroundQuickLook(value unsafe.Pointer)
	NSTimeIntervalSince1970() float64 /* primitive/slice/pointer. */
	SetNSTimeIntervalSince1970(value float64 /* primitive/slice/pointer. */)
	// methods:
	DateByAddingTimeInterval(ti objc.IObject /* cross-framework: TimeInterval */) unsafe.Pointer
	Compare(other IDate) ComparisonResult
	DescriptionWithLocale(locale objectivec.IObject) IString
	EarlierDate(anotherDate IDate) IDate
	IsEqualToDate(otherDate IDate) bool /* primitive/slice/pointer. */
	LaterDate(anotherDate IDate) IDate
	TimeIntervalSinceDate(anotherDate IDate) objc.IObject /* cross-framework: TimeInterval */
}

// A representation of a specific point in time, independent of any calendar or time zone.
//
// In Swift, use this type when you need reference semantics or other Foundation-specific behavior. objects encapsulate a single point in time, independent of any particular calendrical system or time zone. Date objects are immutable, representing an invariant time interval relative to an absolute reference date (00:00:00 UTC on 1 January 2001). The class provides methods for comparing dates, calculating the time interval between two dates, and creating a new date from a time interval relative to another date. objects can be used in conjunction with objects to create localized representations of dates and times, as well as with objects to perform calendar arithmetic. is with its Core Foundation counterpart, . See for more information on toll-free bridging.


// A representation of a specific point in time, independent of any calendar or time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate
type Date struct {
	objectivec.Object
}

// DateFrom constructs a [Date] from an unsafe.Pointer.
//
// A representation of a specific point in time, independent of any calendar or time zone.
func DateFrom(ptr unsafe.Pointer) Date {
	return Date{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DateClass) Alloc() Date {
	rv := objc.Send[Date](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DateClass) New() Date {
	rv := objc.Send[Date](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Date) Init() Date {
	rv := objc.Send[Date](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Date) Autorelease() Date {
	rv := objc.Send[Date](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDate creates a new Date instance.
func NewDate() Date {
	return getDateClass().New()
}



// Returns a date object initialized from data in the given unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/init(coder:)
func NewDateWithCoder(coder ICoder) Date {
	instance := getDateClass().Alloc()
	rv := objc.Send[Date](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/init(SRAbsoluteTime:)-886t8
func NewDateWithSRAbsoluteTime(time unsafe.Pointer) Date {
	instance := getDateClass().Alloc()
	rv := objc.Send[Date](instance.ID, objc.Sel("initWithSRAbsoluteTime:"), time)
	rv.Autorelease()
	return rv
}


// Returns a date object initialized with a date and time value specified by a given string in the international string representation format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/init(string:)
func NewDateWithString(description IString) Date {
	instance := getDateClass().Alloc()
	rv := objc.Send[Date](instance.ID, objc.Sel("initWithString:"), description)
	rv.Autorelease()
	return rv
}


// Returns a date object initialized relative to 00:00:00 UTC on 1 January 1970 by a given number of seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/init(timeIntervalSince1970:)
func NewDateWithTimeIntervalSince1970(secs objc.IObject /* cross-framework: TimeInterval */) Date {
	instance := getDateClass().Alloc()
	rv := objc.Send[Date](instance.ID, objc.Sel("initWithTimeIntervalSince1970:"), secs)
	rv.Autorelease()
	return rv
}


// Returns a date object initialized relative to another given date by a given number of seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/init(timeInterval:sinceDate:)-71m1f
func NewDateWithTimeIntervalSinceDate(secsToBeAdded objc.IObject /* cross-framework: TimeInterval */, date IDate) Date {
	instance := getDateClass().Alloc()
	rv := objc.Send[Date](instance.ID, objc.Sel("initWithTimeInterval:sinceDate:"), secsToBeAdded, date)
	rv.Autorelease()
	return rv
}


// Returns a date object initialized relative to the current date and time by a given number of seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/init(timeIntervalSinceNow:)
func NewDateWithTimeIntervalSinceNow(secs objc.IObject /* cross-framework: TimeInterval */) Date {
	instance := getDateClass().Alloc()
	rv := objc.Send[Date](instance.ID, objc.Sel("initWithTimeIntervalSinceNow:"), secs)
	rv.Autorelease()
	return rv
}


// Returns a date object initialized relative to 00:00:00 UTC on 1 January 2001 by a given number of seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/init(timeIntervalSinceReferenceDate:)
func NewDateWithTimeIntervalSinceReferenceDate(ti objc.IObject /* cross-framework: TimeInterval */) Date {
	instance := getDateClass().Alloc()
	rv := objc.Send[Date](instance.ID, objc.Sel("initWithTimeIntervalSinceReferenceDate:"), ti)
	rv.Autorelease()
	return rv
}



// Creates and returns a date object with a date and time value specified by a given string in the international string representation format ( ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/date(with:)
func (dc _DateClass) DateWithString(aString IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("dateWithString:"), aString)
	return rv
}


// Creates and returns a date object set to the date and time specified by a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/date(withNaturalLanguageString:)
func (dc _DateClass) DateWithNaturalLanguageString(string_ IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("dateWithNaturalLanguageString:"), string_)
	return rv
}


// Creates and returns a date object set to the date and time specified by a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/date(withNaturalLanguageString:locale:)
func (dc _DateClass) DateWithNaturalLanguageStringLocale(string_ IString, locale objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("dateWithNaturalLanguageString:locale:"), string_, locale)
	return rv
}


// Creates and returns a new date object set to the current date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/date
func (dc _DateClass) Date() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("date"))
	return rv
}


// Creates and returns a date object set to the given number of seconds from 00:00:00 UTC on 1 January 1970.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/dateWithTimeIntervalSince1970:
func (dc _DateClass) DateWithTimeIntervalSince1970(secs objc.IObject /* cross-framework: TimeInterval */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dateWithTimeIntervalSince1970:"), secs)
	return rv
}


// Creates and returns a date object set to a given number of seconds from the current date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/dateWithTimeIntervalSinceNow:
func (dc _DateClass) DateWithTimeIntervalSinceNow(secs objc.IObject /* cross-framework: TimeInterval */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dateWithTimeIntervalSinceNow:"), secs)
	return rv
}


// Creates and returns a date object set to a given number of seconds from 00:00:00 UTC on 1 January 2001.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/dateWithTimeIntervalSinceReferenceDate:
func (dc _DateClass) DateWithTimeIntervalSinceReferenceDate(ti objc.IObject /* cross-framework: TimeInterval */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dateWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/init(SRAbsoluteTime:)-9wpl1
func (dc _DateClass) DateWithSRAbsoluteTime(time unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dateWithSRAbsoluteTime:"), time)
	return rv
}


// Creates and returns a date object set to a given number of seconds from the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/init(timeInterval:sinceDate:)-49cea
func (dc _DateClass) DateWithTimeIntervalSinceDate(secsToBeAdded objc.IObject /* cross-framework: TimeInterval */, date IDate) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dateWithTimeInterval:sinceDate:"), secsToBeAdded, date)
	return rv
}


// A date object representing a date in the distant future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/distantFuture
func (dc _DateClass) DistantFuture() Date {
	rv := objc.Send[Date](objc.ID(dc.class), objc.Sel("distantFuture"))
	return rv
}

// A date object representing a date in the distant past.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/distantPast
func (dc _DateClass) DistantPast() Date {
	rv := objc.Send[Date](objc.ID(dc.class), objc.Sel("distantPast"))
	return rv
}

// The current date and time, as of the time of access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/now
func (dc _DateClass) Now() Date {
	rv := objc.Send[Date](objc.ID(dc.class), objc.Sel("now"))
	return rv
}

// Returns a new date object that is set to a given number of seconds relative to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/addingTimeInterval(_:)
func (d_ Date) DateByAddingTimeInterval(ti objc.IObject /* cross-framework: TimeInterval */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("dateByAddingTimeInterval:"), ti)
	return rv
}


// Indicates the temporal ordering of the receiver and another given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/compare(_:)
func (d_ Date) Compare(other IDate) ComparisonResult {
	rv := objc.Send[ComparisonResult](d_.ID, objc.Sel("compare:"), other)
	return rv
}


// Returns a string representation of the date using the given locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/description(with:)
func (d_ Date) DescriptionWithLocale(locale objectivec.IObject) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}


// Returns the earlier of the receiver and another given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/earlierDate(_:)
func (d_ Date) EarlierDate(anotherDate IDate) IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("earlierDate:"), anotherDate)
	return rv
}


// Returns a Boolean value that indicates whether a given object is a date that is exactly equal the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/isEqual(to:)
func (d_ Date) IsEqualToDate(otherDate IDate) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEqualToDate:"), otherDate)
	return rv
}


// Returns the later of the receiver and another given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/laterDate(_:)
func (d_ Date) LaterDate(anotherDate IDate) IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("laterDate:"), anotherDate)
	return rv
}


// Returns the interval between the receiver and another given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSince(_:)
func (d_ Date) TimeIntervalSinceDate(anotherDate IDate) objc.IObject /* cross-framework: TimeInterval */ {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("timeIntervalSinceDate:"), anotherDate)
	return rv
}


// A string representation of the date object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/description
func (d_ Date) Description() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("description"))
	return rv
}


// A date object representing a date in the distant future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/distantFuture
func (d_ Date) DistantFuture() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("distantFuture"))
	return rv
}


// A date object representing a date in the distant past.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/distantPast
func (d_ Date) DistantPast() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("distantPast"))
	return rv
}


// The current date and time, as of the time of access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/now
func (d_ Date) Now() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("now"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/srAbsoluteTime
func (d_ Date) SrAbsoluteTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("srAbsoluteTime"))
	return rv
}


// The interval between the date object and 00:00:00 UTC on 1 January 1970.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSince1970
func (d_ Date) TimeIntervalSince1970() objc.IObject /* cross-framework: TimeInterval */ {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("timeIntervalSince1970"))
	return rv
}


// The interval between the date object and the current date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSinceNow
func (d_ Date) TimeIntervalSinceNow() objc.IObject /* cross-framework: TimeInterval */ {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("timeIntervalSinceNow"))
	return rv
}


// The interval between the date object and 00:00:00 UTC on 1 January 2001.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSinceReferenceDate-swift.property
func (d_ Date) TimeIntervalSinceReferenceDate() objc.IObject /* cross-framework: TimeInterval */ {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("timeIntervalSinceReferenceDate"))
	return rv
}


// A custom playground Quick Look for this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/customplaygroundquicklook
func (d_ Date) CustomPlaygroundQuickLook() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("customPlaygroundQuickLook"))
	return rv
}


// A custom playground Quick Look for this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/customplaygroundquicklook
func (d_ Date) SetCustomPlaygroundQuickLook(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCustomPlaygroundQuickLook:"), value)
}


// The number of seconds from 1 January 1970 to the reference date, 1 January 2001.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimeintervalsince1970
func (d_ Date) NSTimeIntervalSince1970() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](d_.ID, objc.Sel("NSTimeIntervalSince1970"))
	return rv
}


// The number of seconds from 1 January 1970 to the reference date, 1 January 2001.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimeintervalsince1970
func (d_ Date) SetNSTimeIntervalSince1970(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSTimeIntervalSince1970:"), value)
}


