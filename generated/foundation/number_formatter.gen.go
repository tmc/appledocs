// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NumberFormatter] class.
var numberFormatterClass = _NumberFormatterClass{objc.GetClass("NSNumberFormatter")}

type _NumberFormatterClass struct {
	class objc.Class
}

// An interface definition for the [NumberFormatter] class.
type INumberFormatter interface {
	IFormatter
}

// A formatter that converts between numeric values and their textual representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter

type NumberFormatter struct {
	Formatter
}

// NumberFormatterFrom constructs a [NumberFormatter] from an unsafe.Pointer.
//
// A formatter that converts between numeric values and their textual representations.
func NumberFormatterFrom(ptr unsafe.Pointer) NumberFormatter {
	return NumberFormatter{
		Formatter: FormatterFrom(ptr),
	}
}



