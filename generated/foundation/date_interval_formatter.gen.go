// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateIntervalFormatter] class.
var (
	dateIntervalFormatterClass     _DateIntervalFormatterClass
	dateIntervalFormatterClassOnce sync.Once
)

func getDateIntervalFormatterClass() _DateIntervalFormatterClass {
	dateIntervalFormatterClassOnce.Do(func() {
		dateIntervalFormatterClass = _DateIntervalFormatterClass{objc.GetClass("NSDateIntervalFormatter")}
	})
	return dateIntervalFormatterClass
}

type _DateIntervalFormatterClass struct {
	class objc.Class
}

// An interface definition for the [DateIntervalFormatter] class.
type IDateIntervalFormatter interface {
	IFormatter
}

// A formatter that creates string representations of time intervals.
//
// A object creates user-readable strings from pairs of dates. Use a date interval formatter to create user-readable strings of the form for your app’s interface, where and are date values that you supply. The formatter uses locale and language information, along with custom formatting options, to define the content of the resulting string. You can specify different styles for the date and time information in each date value. To use this class, create an instance, configure its properties, and call the method to generate a string. The properties of this class let you configure the calendar and specify the style to apply to date and time values. Given a current date of January 16, 2015, Configuring the Formatter Options shows how to configure a formatter object and generate the string “1/16/15 - 1/17/15”. Configuring a formatter object The method may be called safely from any thread of your app. It is also safe to share a single instance of this class from multiple threads, with the caveat that you should not change the configuration of the object while another thread is using it to generate a string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter
type DateIntervalFormatter struct {
	Formatter
}

// DateIntervalFormatterFrom constructs a [DateIntervalFormatter] from an unsafe.Pointer.
//
// A formatter that creates string representations of time intervals.
func DateIntervalFormatterFrom(ptr unsafe.Pointer) DateIntervalFormatter {
	return DateIntervalFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DateIntervalFormatterClass) Alloc() DateIntervalFormatter {
	rv := objc.Send[DateIntervalFormatter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DateIntervalFormatterClass) New() DateIntervalFormatter {
	rv := objc.Send[DateIntervalFormatter](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DateIntervalFormatter) Init() DateIntervalFormatter {
	rv := objc.Send[DateIntervalFormatter](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DateIntervalFormatter) Autorelease() DateIntervalFormatter {
	rv := objc.Send[DateIntervalFormatter](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateIntervalFormatter creates a new DateIntervalFormatter instance.
func NewDateIntervalFormatter() DateIntervalFormatter {
	return getDateIntervalFormatterClass().New()
}




