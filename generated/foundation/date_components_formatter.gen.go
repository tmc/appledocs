// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateComponentsFormatter] class.
var dateComponentsFormatterClass = _DateComponentsFormatterClass{objc.GetClass("NSDateComponentsFormatter")}

type _DateComponentsFormatterClass struct {
	class objc.Class
}

// An interface definition for the [DateComponentsFormatter] class.
type IDateComponentsFormatter interface {
	IFormatter
}

// A formatter that creates string representations of quantities of time. [Full Topic]
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



