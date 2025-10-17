// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Protocol] class.
var protocolClass = _ProtocolClass{objc.GetClass("Protocol")}

type _ProtocolClass struct {
	class objc.Class
}

// An interface definition for the [Protocol] class.
type IProtocol interface {
	IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/Protocol

type Protocol struct {
	Object
}

// ProtocolFrom constructs a [Protocol] from an unsafe.Pointer.
func ProtocolFrom(ptr unsafe.Pointer) Protocol {
	return Protocol{Object{objc.ID(ptr)}}
}



