// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MessagePort] class.
var messagePortClass = _MessagePortClass{objc.GetClass("NSMessagePort")}

type _MessagePortClass struct {
	class objc.Class
}

// An interface definition for the [MessagePort] class.
type IMessagePort interface {
	IPort
}

// A port that can be used as an endpoint for distributed object connections (or raw messaging). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MessagePort

type MessagePort struct {
	Port
}

// MessagePortFrom constructs a [MessagePort] from an unsafe.Pointer.
//
// A port that can be used as an endpoint for distributed object connections (or raw messaging).
func MessagePortFrom(ptr unsafe.Pointer) MessagePort {
	return MessagePort{
		Port: PortFrom(ptr),
	}
}



