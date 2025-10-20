// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DateComponents] class.
var (
	dateComponentsClass     _DateComponentsClass
	dateComponentsClassOnce sync.Once
)

func getDateComponentsClass() _DateComponentsClass {
	dateComponentsClassOnce.Do(func() {
		dateComponentsClass = _DateComponentsClass{objc.GetClass("NSDateComponents")}
	})
	return dateComponentsClass
}

type _DateComponentsClass struct {
	class objc.Class
}

// An interface definition for the [DateComponents] class.
type IDateComponents interface {
	objectivec.IObject
}

// An object that specifies a date or time in terms of units (such as year, month, day, hour, and minute) to be evaluated in a calendar system and time zone.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. encapsulates the components of a date in an extendable, object-oriented manner. It’s used to specify a date by providing the temporal components that make up a date and time: hour, minutes, seconds, day, month, year, and so on. You can also use it to specify a duration of time, for example, 5 hours and 16 minutes. An object is not required to define all the component fields. When a new instance of is created, the date components are set to . An instance of is not responsible for answering questions about a date beyond the information with which it was initialized. For example, if you initialize one with May 4, 2017, its weekday is , not Thursday. To get the correct day of the week, you must create a suitable instance of , create an object using and then use to retrieve the weekday—as illustrated in the following example. For more details, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents
type DateComponents struct {
	objectivec.Object
}

// DateComponentsFrom constructs a [DateComponents] from an unsafe.Pointer.
//
// An object that specifies a date or time in terms of units (such as year, month, day, hour, and minute) to be evaluated in a calendar system and time zone.
func DateComponentsFrom(ptr unsafe.Pointer) DateComponents {
	return DateComponents{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DateComponentsClass) Alloc() DateComponents {
	rv := objc.Send[DateComponents](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DateComponentsClass) New() DateComponents {
	rv := objc.Send[DateComponents](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DateComponents) Init() DateComponents {
	rv := objc.Send[DateComponents](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DateComponents) Autorelease() DateComponents {
	rv := objc.Send[DateComponents](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateComponents creates a new DateComponents instance.
func NewDateComponents() DateComponents {
	return getDateComponentsClass().New()
}




