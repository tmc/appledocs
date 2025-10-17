// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateIntervalFormatter] class.
var dateIntervalFormatterClass = _DateIntervalFormatterClass{objc.GetClass("NSDateIntervalFormatter")}

type _DateIntervalFormatterClass struct {
	class objc.Class
}

// An interface definition for the [DateIntervalFormatter] class.
type IDateIntervalFormatter interface {
	IFormatter
}

// A formatter that creates string representations of time intervals. [Full Topic]
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



