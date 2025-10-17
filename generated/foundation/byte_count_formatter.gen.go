// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ByteCountFormatter] class.
var byteCountFormatterClass = _ByteCountFormatterClass{objc.GetClass("NSByteCountFormatter")}

type _ByteCountFormatterClass struct {
	class objc.Class
}

// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter

type ByteCountFormatter struct {
	Formatter
}

// ByteCountFormatterFrom constructs a [ByteCountFormatter] from an unsafe.Pointer.
//
// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on).
func ByteCountFormatterFrom(ptr unsafe.Pointer) ByteCountFormatter {
	return ByteCountFormatter{
		Formatter: FormatterFrom(ptr),
	}
}



