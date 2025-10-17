// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XPCCoder] class.
var xPCCoderClass = _XPCCoderClass{objc.GetClass("NSXPCCoder")}

type _XPCCoderClass struct {
	class objc.Class
}

// A coder that encodes and decodes objects that your app sends over an XPC connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder

type XPCCoder struct {
	Coder
}

// XPCCoderFrom constructs a [XPCCoder] from an unsafe.Pointer.
//
// A coder that encodes and decodes objects that your app sends over an XPC connection.
func XPCCoderFrom(ptr unsafe.Pointer) XPCCoder {
	return XPCCoder{
		Coder: CoderFrom(ptr),
	}
}



