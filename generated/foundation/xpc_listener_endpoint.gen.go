// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XPCListenerEndpoint] class.
var XPCListenerEndpointClass objc.Class

func init() {
	XPCListenerEndpointClass = objc.GetClass("NSXPCListenerEndpoint")
}

type XPCListenerEndpoint struct {
	objc.ID
}

func XPCListenerEndpointFrom(ptr unsafe.Pointer) XPCListenerEndpoint {
	return XPCListenerEndpoint{
		ID: objc.ID(ptr),
	}
}



