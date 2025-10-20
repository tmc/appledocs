// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var xPCListenerClass _XPCListenerClass

func init() {
	xPCListenerClass = _XPCListenerClass{objc.GetClass("NSXPCListener")}
}

type _XPCListenerClass struct {
	class objc.Class
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/anonymous()
func (xc _XPCListenerClass) AnonymousListener() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(xc.class), objc.Sel("anonymousListener"))
	return rv
}


