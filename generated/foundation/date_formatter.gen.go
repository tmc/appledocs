// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateFormatter] class.
var (
	dateFormatterClass     _DateFormatterClass
	dateFormatterClassOnce sync.Once
)

func getDateFormatterClass() _DateFormatterClass {
	dateFormatterClassOnce.Do(func() {
		dateFormatterClass = _DateFormatterClass{objc.GetClass("NSDateFormatter")}
	})
	return dateFormatterClass
}

type _DateFormatterClass struct {
	class objc.Class
}

// An interface definition for the [DateFormatter] class.
type IDateFormatter interface {
	IFormatter
}

// A formatter that converts between dates and their textual representations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter
type DateFormatter struct {
	Formatter
}

// DateFormatterFrom constructs a [DateFormatter] from an unsafe.Pointer.
//
// A formatter that converts between dates and their textual representations.
func DateFormatterFrom(ptr unsafe.Pointer) DateFormatter {
	return DateFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DateFormatterClass) Alloc() DateFormatter {
	rv := objc.Send[DateFormatter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DateFormatterClass) New() DateFormatter {
	rv := objc.Send[DateFormatter](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DateFormatter) Init() DateFormatter {
	rv := objc.Send[DateFormatter](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DateFormatter) Autorelease() DateFormatter {
	rv := objc.Send[DateFormatter](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateFormatter creates a new DateFormatter instance.
func NewDateFormatter() DateFormatter {
	return getDateFormatterClass().New()
}




