// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ISO8601DateFormatter] class.
var iSO8601DateFormatterClass = _ISO8601DateFormatterClass{objc.GetClass("NSISO8601DateFormatter")}

type _ISO8601DateFormatterClass struct {
	class objc.Class
}

// A formatter that converts between dates and their ISO 8601 string representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter

type ISO8601DateFormatter struct {
	Formatter
}

// ISO8601DateFormatterFrom constructs a [ISO8601DateFormatter] from an unsafe.Pointer.
//
// A formatter that converts between dates and their ISO 8601 string representations.
func ISO8601DateFormatterFrom(ptr unsafe.Pointer) ISO8601DateFormatter {
	return ISO8601DateFormatter{
		Formatter: FormatterFrom(ptr),
	}
}



