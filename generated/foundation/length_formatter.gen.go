// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LengthFormatter] class.
var lengthFormatterClass = _LengthFormatterClass{objc.GetClass("NSLengthFormatter")}

type _LengthFormatterClass struct {
	class objc.Class
}

// An interface definition for the [LengthFormatter] class.
type ILengthFormatter interface {
	IFormatter
}

// A formatter that provides localized descriptions of linear distances, such as length and height measurements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter

type LengthFormatter struct {
	Formatter
}

// LengthFormatterFrom constructs a [LengthFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized descriptions of linear distances, such as length and height measurements.
func LengthFormatterFrom(ptr unsafe.Pointer) LengthFormatter {
	return LengthFormatter{
		Formatter: FormatterFrom(ptr),
	}
}



