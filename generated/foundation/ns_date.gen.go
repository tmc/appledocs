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
	CustomPlaygroundQuickLook() unsafe.Pointer
	SetCustomPlaygroundQuickLook(value unsafe.Pointer)
	Description() string
	SetDescription(value string)
	SrAbsoluteTime() unsafe.Pointer
	SetSrAbsoluteTime(value unsafe.Pointer)
	TimeIntervalSince1970() TimeInterval
	SetTimeIntervalSince1970(value ITimeInterval)
	TimeIntervalSinceNow() TimeInterval
	SetTimeIntervalSinceNow(value ITimeInterval)
	NSTimeIntervalSince1970() float64
	SetNSTimeIntervalSince1970(value float64)
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



// A date object representing a date in the distant future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/distantFuture

func (dc _DateClass) DistantFuture() Date {
	rv := objc.Send[NSDate](objc.ID(dc.class), objc.Sel("distantFuture"))
	return rv
}

// A date object representing a date in the distant past.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/distantPast

func (dc _DateClass) DistantPast() Date {
	rv := objc.Send[NSDate](objc.ID(dc.class), objc.Sel("distantPast"))
	return rv
}

// The interval between 00:00:00 UTC on 1 January 2001 and the current date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSinceReferenceDate-swift.type.property

func (dc _DateClass) TimeIntervalSinceReferenceDate() TimeInterval {
	rv := objc.Send[TimeInterval](objc.ID(dc.class), objc.Sel("timeIntervalSinceReferenceDate"))
	return rv
}

// A date object representing a date in the distant future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/distantFuture

func (d_ Date) DistantFuture() NSDate {
	rv := objc.Send[NSDate](d_.ID, objc.Sel("distantFuture"))
	return rv
}


// A date object representing a date in the distant past.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/distantPast

func (d_ Date) DistantPast() NSDate {
	rv := objc.Send[NSDate](d_.ID, objc.Sel("distantPast"))
	return rv
}


// The interval between 00:00:00 UTC on 1 January 2001 and the current date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSinceReferenceDate-swift.type.property

func (d_ Date) TimeIntervalSinceReferenceDate() TimeInterval {
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


// A string representation of the date object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/description

func (d_ Date) Description() string {
	rv := objc.Send[string](d_.ID, objc.Sel("description"))
	return rv
}


// A string representation of the date object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/description

func (d_ Date) SetDescription(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDescription:"), objc.String(value))
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/srabsolutetime

func (d_ Date) SrAbsoluteTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("srAbsoluteTime"))
	return rv
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/srabsolutetime

func (d_ Date) SetSrAbsoluteTime(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSrAbsoluteTime:"), value)
}


// The interval between the date object and 00:00:00 UTC on 1 January 1970.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/timeintervalsince1970

func (d_ Date) TimeIntervalSince1970() TimeInterval {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("timeIntervalSince1970"))
	return rv
}


// The interval between the date object and 00:00:00 UTC on 1 January 1970.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/timeintervalsince1970

func (d_ Date) SetTimeIntervalSince1970(value ITimeInterval) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeIntervalSince1970:"), value)
}


// The interval between the date object and the current date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/timeintervalsincenow

func (d_ Date) TimeIntervalSinceNow() TimeInterval {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("timeIntervalSinceNow"))
	return rv
}


// The interval between the date object and the current date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/timeintervalsincenow

func (d_ Date) SetTimeIntervalSinceNow(value ITimeInterval) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeIntervalSinceNow:"), value)
}


// The number of seconds from 1 January 1970 to the reference date, 1 January 2001.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimeintervalsince1970

func (d_ Date) NSTimeIntervalSince1970() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("NSTimeIntervalSince1970"))
	return rv
}


// The number of seconds from 1 January 1970 to the reference date, 1 January 2001.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimeintervalsince1970

func (d_ Date) SetNSTimeIntervalSince1970(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSTimeIntervalSince1970:"), value)
}



