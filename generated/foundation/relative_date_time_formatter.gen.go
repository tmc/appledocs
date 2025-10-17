// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RelativeDateTimeFormatter] class.
var relativeDateTimeFormatterClass = _RelativeDateTimeFormatterClass{objc.GetClass("NSRelativeDateTimeFormatter")}

type _RelativeDateTimeFormatterClass struct {
	class objc.Class
}

// A formatter that creates locale-aware string representations of a relative date or time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter

type RelativeDateTimeFormatter struct {
	Formatter
}

// RelativeDateTimeFormatterFrom constructs a [RelativeDateTimeFormatter] from an unsafe.Pointer.
//
// A formatter that creates locale-aware string representations of a relative date or time.
func RelativeDateTimeFormatterFrom(ptr unsafe.Pointer) RelativeDateTimeFormatter {
	return RelativeDateTimeFormatter{
		Formatter: FormatterFrom(ptr),
	}
}



