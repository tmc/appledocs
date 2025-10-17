// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MassFormatter] class.
var massFormatterClass = _MassFormatterClass{objc.GetClass("NSMassFormatter")}

type _MassFormatterClass struct {
	class objc.Class
}

// A formatter that provides localized descriptions of mass and weight values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter

type MassFormatter struct {
	Formatter
}

// MassFormatterFrom constructs a [MassFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized descriptions of mass and weight values.
func MassFormatterFrom(ptr unsafe.Pointer) MassFormatter {
	return MassFormatter{
		Formatter: FormatterFrom(ptr),
	}
}



