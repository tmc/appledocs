// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateFormatter] class.
var dateFormatterClass = _DateFormatterClass{objc.GetClass("NSDateFormatter")}

type _DateFormatterClass struct {
	class objc.Class
}

// An interface definition for the [DateFormatter] class.
type IDateFormatter interface {
	IFormatter
}

// A formatter that converts between dates and their textual representations. [Full Topic]
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



