// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XPCInterface] class.
var xPCInterfaceClass = _XPCInterfaceClass{objc.GetClass("NSXPCInterface")}

type _XPCInterfaceClass struct {
	class objc.Class
}

// An interface definition for the [XPCInterface] class.
type IXPCInterface interface {
	objectivec.IObject
}

// An interface that may be sent to an exported object or remote object proxy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface

type XPCInterface struct {
	objectivec.Object
}

// XPCInterfaceFrom constructs a [XPCInterface] from an unsafe.Pointer.
//
// An interface that may be sent to an exported object or remote object proxy.
func XPCInterfaceFrom(ptr unsafe.Pointer) XPCInterface {
	return XPCInterface{objectivec.Object{objc.ID(ptr)}}
}



