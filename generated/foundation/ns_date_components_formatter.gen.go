// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateComponentsFormatter] class.
var (
	DateComponentsFormatterClass     _DateComponentsFormatterClass
	DateComponentsFormatterClassOnce sync.Once
)

func getDateComponentsFormatterClass() _DateComponentsFormatterClass {
	DateComponentsFormatterClassOnce.Do(func() {
		DateComponentsFormatterClass = _DateComponentsFormatterClass{objc.GetClass("NSDateComponentsFormatter")}
	})
	return DateComponentsFormatterClass
}

type _DateComponentsFormatterClass struct {
	class objc.Class
}

// An interface definition for the [DateComponentsFormatter] class.
type IDateComponentsFormatter interface {
	IFormatter
	StringForObjectValue(obj objc.ID) string
}

// A formatter that creates string representations of quantities of time.
//
// An object takes quantities of time and formats them as a user-readable string. Use a date components formatter to create strings for your app’s interface. The formatter object has many options for creating both abbreviated and expanded strings. The formatter takes the current user’s locale and language into account when generating strings. To use this class, create an instance, configure its properties, and call one of its methods to generate an appropriate string. The properties of this class let you configure the calendar and specify the date and time units you want displayed in the resulting string. The listing below shows how to configure a formatter to create the string “About 5 minutes remaining”. The methods of this class may be called safely from any thread of your app. It is also safe to share a single instance of this class from multiple threads, with the caveat that you should not change the configuration of the object while another thread is using it to generate a string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter
type DateComponentsFormatter struct {
	Formatter
}

// DateComponentsFormatterFrom constructs a [DateComponentsFormatter] from an unsafe.Pointer.
//
// A formatter that creates string representations of quantities of time.
func DateComponentsFormatterFrom(ptr unsafe.Pointer) DateComponentsFormatter {
	return DateComponentsFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DateComponentsFormatterClass) Alloc() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DateComponentsFormatterClass) New() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DateComponentsFormatter) Init() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DateComponentsFormatter) Autorelease() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateComponentsFormatter creates a new DateComponentsFormatter instance.
func NewDateComponentsFormatter() DateComponentsFormatter {
	return getDateComponentsFormatterClass().New()
}

// Returns a formatted string based on the date information in the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/string(for:)
func (d_ DateComponentsFormatter) StringForObjectValue(obj objc.ID) string {
	rv := objc.Send[string](d_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}

// The formatting style for units whose value is 0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/zeroFormattingBehavior-swift.property
func (d_ DateComponentsFormatter) ZeroFormattingBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("zeroFormattingBehavior"))
	return rv
}

// SetZeroFormattingBehavior sets the value of the zeroFormattingBehavior property.
// The formatting style for units whose value is 0.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/zeroFormattingBehavior-swift.property
func (d_ DateComponentsFormatter) SetZeroFormattingBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setZeroFormattingBehavior:"), value)
}
