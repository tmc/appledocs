// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InputStream] class.
var inputStreamClass = _InputStreamClass{objc.GetClass("NSInputStream")}

type _InputStreamClass struct {
	class objc.Class
}

// A stream that provides read-only stream functionality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream

type InputStream struct {
	Stream
}

// InputStreamFrom constructs a [InputStream] from an unsafe.Pointer.
//
// A stream that provides read-only stream functionality.
func InputStreamFrom(ptr unsafe.Pointer) InputStream {
	return InputStream{
		Stream: StreamFrom(ptr),
	}
}



