// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EKRecurrenceEnd */


/* debug [class_header]: Header for EKRecurrenceEnd */
// The class instance for the [EKRecurrenceEnd] class.
var (
	EKRecurrenceEndClass     _EKRecurrenceEndClass
	EKRecurrenceEndClassOnce sync.Once
)

func getEKRecurrenceEndClass() _EKRecurrenceEndClass {
	EKRecurrenceEndClassOnce.Do(func() {
		EKRecurrenceEndClass = _EKRecurrenceEndClass{objc.GetClass("EKRecurrenceEnd")}
	})
	return EKRecurrenceEndClass
}

type _EKRecurrenceEndClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKRecurrenceEnd */
// An interface definition for the [EKRecurrenceEnd] class.
type IEKRecurrenceEnd interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EKRecurrenceEnd */
	// properties:
	EndDate() objc.IObject /* cross-framework: NSDate */
	OccurrenceCount() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKRecurrenceEnd */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKRecurrenceEnd */
// Alloc allocates a new instance without initialization.
func (ec _EKRecurrenceEndClass) Alloc() EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EKRecurrenceEndClass) New() EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKRecurrenceEnd) Init() EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKRecurrenceEnd) Autorelease() EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKRecurrenceEnd creates a new EKRecurrenceEnd instance.
func NewEKRecurrenceEnd() EKRecurrenceEnd {
	return getEKRecurrenceEndClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKRecurrenceEnd */
// A class that defines the end of a recurrence rule.
//
// The class defines the end of a recurrence rule defined by an object. The recurrence end can be specified by a date (date-based) or by a maximum count of occurrences (count-based). An event that is intended to continue indefinitely should have its set to .


// A class that defines the end of a recurrence rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd
type EKRecurrenceEnd struct {
	objectivec.Object
}

// EKRecurrenceEndFrom constructs a [EKRecurrenceEnd] from an unsafe.Pointer.
//
// A class that defines the end of a recurrence rule.
func EKRecurrenceEndFrom(ptr unsafe.Pointer) EKRecurrenceEnd {
	return EKRecurrenceEnd{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKRecurrenceEnd */

// Initializes and returns a date-based recurrence end with a given end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/init(end:)
func NewEKRecurrenceEndWithEndDate(endDate objc.IObject /* cross-framework: NSDate */) EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](objc.ID(getEKRecurrenceEndClass().class), objc.Sel("recurrenceEndWithEndDate:"), endDate)
	return rv
}/* debug [class_init_methods/constructor]: NewEKRecurrenceEndWithEndDate */


// Initializes and returns a count-based recurrence end with a given maximum occurrence count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/init(occurrenceCount:)
func NewEKRecurrenceEndWithOccurrenceCount(occurrenceCount uint) EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](objc.ID(getEKRecurrenceEndClass().class), objc.Sel("recurrenceEndWithOccurrenceCount:"), occurrenceCount)
	return rv
}/* debug [class_init_methods/constructor]: NewEKRecurrenceEndWithOccurrenceCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKRecurrenceEnd */

// Initializes and returns a date-based recurrence end with a given end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/init(end:)
func (ec _EKRecurrenceEndClass) RecurrenceEndWithEndDate(endDate objc.IObject /* cross-framework: NSDate */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("recurrenceEndWithEndDate:"), endDate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RecurrenceEndWithEndDate) */


// Initializes and returns a count-based recurrence end with a given maximum occurrence count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/init(occurrenceCount:)
func (ec _EKRecurrenceEndClass) RecurrenceEndWithOccurrenceCount(occurrenceCount uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("recurrenceEndWithOccurrenceCount:"), occurrenceCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RecurrenceEndWithOccurrenceCount) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKRecurrenceEnd */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKRecurrenceEnd */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKRecurrenceEnd */

// The end date of the recurrence end, or if the recurrence end is count-based.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/endDate
func (e_ EKRecurrenceEnd) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The occurrence count of the recurrence end, or if the recurrence end is date-based.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/occurrenceCount
func (e_ EKRecurrenceEnd) OccurrenceCount() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("occurrenceCount"))
	return rv
}/* debug [instance_properties/getter]: occurrenceCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKRecurrenceEnd */


