// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XPCListenerEndpoint] class.
var XPCListenerEndpointClass = _XPCListenerEndpointClass{objc.GetClass("NSXPCListenerEndpoint")}

type _XPCListenerEndpointClass struct {
	class objc.Class
}

type XPCListenerEndpoint struct {
	objc.ID
}

func XPCListenerEndpointFrom(ptr unsafe.Pointer) XPCListenerEndpoint {
	return XPCListenerEndpoint{
		ID: objc.ID(ptr),
	}
}




