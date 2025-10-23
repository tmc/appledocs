// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DateInterval] class.
var (
	DateIntervalClass     _DateIntervalClass
	DateIntervalClassOnce sync.Once
)

func getDateIntervalClass() _DateIntervalClass {
	DateIntervalClassOnce.Do(func() {
		DateIntervalClass = _DateIntervalClass{objc.GetClass("NSDateInterval")}
	})
	return DateIntervalClass
}

type _DateIntervalClass struct {
	class objc.Class
}

// An interface definition for the [DateInterval] class.
type IDateInterval interface {
	objectivec.IObject
	// properties:
	Duration() objc.IObject /* cross-framework: TimeInterval */
	EndDate() IDate
	StartDate() IDate
	// methods:
	Compare(dateInterval IDateInterval) ComparisonResult /* not a class type */
	ContainsDate(date IDate) bool /* primitive/slice/pointer. */
	IntersectionWithDateInterval(dateInterval IDateInterval) IDateInterval
	IntersectsDateInterval(dateInterval IDateInterval) bool /* primitive/slice/pointer. */
	IsEqualToDateInterval(dateInterval IDateInterval) bool /* primitive/slice/pointer. */
}

// An object representing the span of time between a specific start date and end date.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. An object represents a closed interval between two dates. The class provides a programmatic interface for calculating the duration of a time interval and determining whether a date falls within it, as well as comparing date intervals and checking to see whether they intersect. An object consists of a and an . The and of a date interval can be equal, in which case its is . However, cannot occur earlier than . You can use the class to create string representations of objects that are suitable for display in the current locale.


// An object representing the span of time between a specific start date and end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval
type DateInterval struct {
	objectivec.Object
}

// DateIntervalFrom constructs a [DateInterval] from an unsafe.Pointer.
//
// An object representing the span of time between a specific start date and end date.
func DateIntervalFrom(ptr unsafe.Pointer) DateInterval {
	return DateInterval{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DateIntervalClass) Alloc() DateInterval {
	rv := objc.Send[DateInterval](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DateIntervalClass) New() DateInterval {
	rv := objc.Send[DateInterval](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DateInterval) Init() DateInterval {
	rv := objc.Send[DateInterval](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DateInterval) Autorelease() DateInterval {
	rv := objc.Send[DateInterval](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateInterval creates a new DateInterval instance.
func NewDateInterval() DateInterval {
	return getDateIntervalClass().New()
}



// Returns a date interval initialized from data in the given unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(coder:)
func NewDateIntervalWithCoder(coder ICoder) DateInterval {
	instance := getDateIntervalClass().Alloc()
	rv := objc.Send[DateInterval](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes a date interval with a given start date and duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(start:duration:)
func NewDateIntervalWithStartDateDuration(startDate IDate, duration objc.IObject /* cross-framework TimeInterval */) DateInterval {
	instance := getDateIntervalClass().Alloc()
	rv := objc.Send[DateInterval](instance.ID, objc.Sel("initWithStartDate:duration:"), startDate, duration)
	rv.Autorelease()
	return rv
}


// Initializes a date interval from a given start date and end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(start:end:)
func NewDateIntervalWithStartDateEndDate(startDate IDate, endDate IDate) DateInterval {
	instance := getDateIntervalClass().Alloc()
	rv := objc.Send[DateInterval](instance.ID, objc.Sel("initWithStartDate:endDate:"), startDate, endDate)
	rv.Autorelease()
	return rv
}



// Compares the receiver with the specified date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/compare(_:)
func (d_ DateInterval) Compare(dateInterval IDateInterval) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](d_.ID, objc.Sel("compare:"), dateInterval)
	return rv
}


// Indicates whether the receiver contains the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/contains(_:)
func (d_ DateInterval) ContainsDate(date IDate) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("containsDate:"), date)
	return rv
}


// Returns the intersection between the receiver and the specified date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/intersection(with:)
func (d_ DateInterval) IntersectionWithDateInterval(dateInterval IDateInterval) IDateInterval {
	rv := objc.Send[DateInterval](d_.ID, objc.Sel("intersectionWithDateInterval:"), dateInterval)
	return rv
}


// Indicates whether the receiver intersects with the specified date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/intersects(_:)
func (d_ DateInterval) IntersectsDateInterval(dateInterval IDateInterval) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("intersectsDateInterval:"), dateInterval)
	return rv
}


// Indicates whether the receiver is equal to the specified date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/isEqual(to:)
func (d_ DateInterval) IsEqualToDateInterval(dateInterval IDateInterval) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEqualToDateInterval:"), dateInterval)
	return rv
}


// The duration of the date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/duration
func (d_ DateInterval) Duration() objc.IObject /* cross-framework: TimeInterval */ {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("duration"))
	return rv
}


// The end date of the date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/endDate
func (d_ DateInterval) EndDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("endDate"))
	return rv
}


// The start date of the date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/startDate
func (d_ DateInterval) StartDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("startDate"))
	return rv
}


