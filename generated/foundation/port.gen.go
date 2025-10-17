// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Port] class.
var portClass = _PortClass{objc.GetClass("NSPort")}

type _PortClass struct {
	class objc.Class
}

// An abstract class that represents a communication channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Port

type Port struct {
	objectivec.Object
}

// PortFrom constructs a [Port] from an unsafe.Pointer.
//
// An abstract class that represents a communication channel.
func PortFrom(ptr unsafe.Pointer) Port {
	return Port{objectivec.Object{objc.ID(ptr)}}
}



