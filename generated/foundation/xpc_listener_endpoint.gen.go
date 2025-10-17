// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XPCListenerEndpoint] class.
var xPCListenerEndpointClass = _XPCListenerEndpointClass{objc.GetClass("NSXPCListenerEndpoint")}

type _XPCListenerEndpointClass struct {
	class objc.Class
}

// An interface definition for the [XPCListenerEndpoint] class.
type IXPCListenerEndpoint interface {
	objectivec.IObject
}

// An object that names a specific XPC listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListenerEndpoint

type XPCListenerEndpoint struct {
	objectivec.Object
}

// XPCListenerEndpointFrom constructs a [XPCListenerEndpoint] from an unsafe.Pointer.
//
// An object that names a specific XPC listener.
func XPCListenerEndpointFrom(ptr unsafe.Pointer) XPCListenerEndpoint {
	return XPCListenerEndpoint{objectivec.Object{objc.ID(ptr)}}
}



