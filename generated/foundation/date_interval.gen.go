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
	dateIntervalClass     _DateIntervalClass
	dateIntervalClassOnce sync.Once
)

func getDateIntervalClass() _DateIntervalClass {
	dateIntervalClassOnce.Do(func() {
		dateIntervalClass = _DateIntervalClass{objc.GetClass("NSDateInterval")}
	})
	return dateIntervalClass
}

type _DateIntervalClass struct {
	class objc.Class
}

// An interface definition for the [DateInterval] class.
type IDateInterval interface {
	objectivec.IObject
}

// An object representing the span of time between a specific start date and end date.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. An object represents a closed interval between two dates. The class provides a programmatic interface for calculating the duration of a time interval and determining whether a date falls within it, as well as comparing date intervals and checking to see whether they intersect. An object consists of a and an . The and of a date interval can be equal, in which case its is . However, cannot occur earlier than . You can use the class to create string representations of objects that are suitable for display in the current locale.
//
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




