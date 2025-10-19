// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Date] class.
var dateClass = _DateClass{objc.GetClass("NSDate")}

type _DateClass struct {
	class objc.Class
}

// An interface definition for the [Date] class.
type IDate interface {
	objectivec.IObject
}

// A representation of a specific point in time, independent of any calendar or time zone. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return dateClass.New()
}




