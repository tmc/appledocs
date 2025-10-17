// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XPCListener] class.
var XPCListenerClass objc.Class

func init() {
	XPCListenerClass = objc.GetClass("NSXPCListener")
}

type XPCListener struct {
	objc.ID
}

func XPCListenerFrom(ptr unsafe.Pointer) XPCListener {
	return XPCListener{
		ID: objc.ID(ptr),
	}
}


// Returns a new anonymous listener connection. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCListener/anonymous()
func (xc XPCListener) AnonymousListener() unsafe.Pointer {
	sel := objc.RegisterName("anonymousListener")
	ret := objc.ID(XPCListenerClass).Send(sel)
	return unsafe.Pointer(ret)
}


