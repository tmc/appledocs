// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ListFormatter] class.
var listFormatterClass = _ListFormatterClass{objc.GetClass("NSListFormatter")}

type _ListFormatterClass struct {
	class objc.Class
}

// An object that provides locale-correct formatting of a list of items using the appropriate separator and conjunction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ListFormatter

type ListFormatter struct {
	Formatter
}

// ListFormatterFrom constructs a [ListFormatter] from an unsafe.Pointer.
//
// An object that provides locale-correct formatting of a list of items using the appropriate separator and conjunction.
func ListFormatterFrom(ptr unsafe.Pointer) ListFormatter {
	return ListFormatter{
		Formatter: FormatterFrom(ptr),
	}
}



