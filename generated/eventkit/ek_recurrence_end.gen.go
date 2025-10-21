// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [EKRecurrenceEnd] class.
type IEKRecurrenceEnd interface {
	objectivec.IObject
}

// A class that defines the end of a recurrence rule.
//
// The class defines the end of a recurrence rule defined by an object. The recurrence end can be specified by a date (date-based) or by a maximum count of occurrences (count-based). An event that is intended to continue indefinitely should have its set to .
//
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

// Alloc allocates a new instance without initialization.
func (ec _EKRecurrenceEndClass) Alloc() EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes and returns a date-based recurrence end with a given end date.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/init(end:)
func NewEKRecurrenceEndWithEndDate(endDate unsafe.Pointer) EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](objc.ID(getEKRecurrenceEndClass().class), objc.Sel("recurrenceEndWithEndDate:"), endDate)
	return rv
}



// Initializes and returns a count-based recurrence end with a given maximum occurrence count.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/init(occurrenceCount:)
func NewEKRecurrenceEndWithOccurrenceCount(occurrenceCount uint) EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](objc.ID(getEKRecurrenceEndClass().class), objc.Sel("recurrenceEndWithOccurrenceCount:"), occurrenceCount)
	return rv
}


// Initializes and returns a date-based recurrence end with a given end date.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/init(end:)
func (ec _EKRecurrenceEndClass) RecurrenceEndWithEndDate(endDate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("recurrenceEndWithEndDate:"), endDate)
	return rv
}

// Initializes and returns a count-based recurrence end with a given maximum occurrence count.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/init(occurrenceCount:)
func (ec _EKRecurrenceEndClass) RecurrenceEndWithOccurrenceCount(occurrenceCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("recurrenceEndWithOccurrenceCount:"), occurrenceCount)
	return rv
}

// The end date of the recurrence end, or if the recurrence end is count-based.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/endDate
func (e_ EKRecurrenceEnd) EndDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("endDate"))
	return rv
}

// The occurrence count of the recurrence end, or if the recurrence end is date-based.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceEnd/occurrenceCount
func (e_ EKRecurrenceEnd) OccurrenceCount() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("occurrenceCount"))
	return rv
}


