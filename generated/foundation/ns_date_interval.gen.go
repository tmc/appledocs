// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDateInterval */


/* debug [class_header]: Header for NSDateInterval */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DateInterval */
// An interface definition for the [DateInterval] class.
type IDateInterval interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DateInterval */
	// properties:
	Duration() float64
	EndDate() IDate
	StartDate() IDate
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DateInterval */
	// methods:
	Compare(dateInterval IDateInterval) ComparisonResult
	ContainsDate(date IDate) bool
	IntersectionWithDateInterval(dateInterval IDateInterval) IDateInterval
	IntersectsDateInterval(dateInterval IDateInterval) bool
	IsEqualToDateInterval(dateInterval IDateInterval) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DateInterval */
// Alloc allocates a new instance without initialization.
func (dc _DateIntervalClass) Alloc() DateInterval {
	rv := objc.Send[DateInterval](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DateInterval */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DateInterval */

// Returns a date interval initialized from data in the given unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(coder:)
func NewDateIntervalWithCoder(coder ICoder) DateInterval {
	instance := getDateIntervalClass().Alloc()
	rv := objc.Send[DateInterval](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDateIntervalWithCoder */


// Initializes a date interval with a given start date and duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(start:duration:)
func NewDateIntervalWithStartDateDuration(startDate IDate, duration float64) DateInterval {
	instance := getDateIntervalClass().Alloc()
	rv := objc.Send[DateInterval](instance.ID, objc.Sel("initWithStartDate:duration:"), startDate, duration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDateIntervalWithStartDateDuration */


// Initializes a date interval from a given start date and end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(start:end:)
func NewDateIntervalWithStartDateEndDate(startDate IDate, endDate IDate) DateInterval {
	instance := getDateIntervalClass().Alloc()
	rv := objc.Send[DateInterval](instance.ID, objc.Sel("initWithStartDate:endDate:"), startDate, endDate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDateIntervalWithStartDateEndDate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DateInterval */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DateInterval */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DateInterval */

// Compares the receiver with the specified date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/compare(_:)
func (d_ DateInterval) Compare(dateInterval IDateInterval) ComparisonResult {
	rv := objc.Send[ComparisonResult](d_.ID, objc.Sel("compare:"), dateInterval)
	return rv
}/* debug [instance_methods/method]: Compare */


// Indicates whether the receiver contains the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/contains(_:)
func (d_ DateInterval) ContainsDate(date IDate) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("containsDate:"), date)
	return rv
}/* debug [instance_methods/method]: ContainsDate */


// Returns the intersection between the receiver and the specified date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/intersection(with:)
func (d_ DateInterval) IntersectionWithDateInterval(dateInterval IDateInterval) IDateInterval {
	rv := objc.Send[DateInterval](d_.ID, objc.Sel("intersectionWithDateInterval:"), dateInterval)
	return rv
}/* debug [instance_methods/method]: IntersectionWithDateInterval */


// Indicates whether the receiver intersects with the specified date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/intersects(_:)
func (d_ DateInterval) IntersectsDateInterval(dateInterval IDateInterval) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("intersectsDateInterval:"), dateInterval)
	return rv
}/* debug [instance_methods/method]: IntersectsDateInterval */


// Indicates whether the receiver is equal to the specified date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/isEqual(to:)
func (d_ DateInterval) IsEqualToDateInterval(dateInterval IDateInterval) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEqualToDateInterval:"), dateInterval)
	return rv
}/* debug [instance_methods/method]: IsEqualToDateInterval */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DateInterval */

// The duration of the date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/duration
func (d_ DateInterval) Duration() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The end date of the date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/endDate
func (d_ DateInterval) EndDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The start date of the date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval/startDate
func (d_ DateInterval) StartDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDateInterval */


