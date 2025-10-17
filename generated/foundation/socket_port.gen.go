// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SocketPort] class.
var socketPortClass = _SocketPortClass{objc.GetClass("NSSocketPort")}

type _SocketPortClass struct {
	class objc.Class
}

// An interface definition for the [SocketPort] class.
type ISocketPort interface {
	IPort
}

// A port that represents a BSD socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort

type SocketPort struct {
	Port
}

// SocketPortFrom constructs a [SocketPort] from an unsafe.Pointer.
//
// A port that represents a BSD socket.
func SocketPortFrom(ptr unsafe.Pointer) SocketPort {
	return SocketPort{
		Port: PortFrom(ptr),
	}
}



