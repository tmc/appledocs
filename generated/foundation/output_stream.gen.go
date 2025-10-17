// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OutputStream] class.
var outputStreamClass = _OutputStreamClass{objc.GetClass("NSOutputStream")}

type _OutputStreamClass struct {
	class objc.Class
}

// A stream that provides write-only stream functionality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream

type OutputStream struct {
	Stream
}

// OutputStreamFrom constructs a [OutputStream] from an unsafe.Pointer.
//
// A stream that provides write-only stream functionality.
func OutputStreamFrom(ptr unsafe.Pointer) OutputStream {
	return OutputStream{
		Stream: StreamFrom(ptr),
	}
}



