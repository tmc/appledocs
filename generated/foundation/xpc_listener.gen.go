// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XPCListener] class.
var xPCListenerClass = _XPCListenerClass{objc.GetClass("NSXPCListener")}

type _XPCListenerClass struct {
	class objc.Class
}

// A listener that waits for new incoming connections, configures them, and accepts or rejects them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener

type XPCListener struct {
	objectivec.Object
}

// XPCListenerFrom constructs a [XPCListener] from an unsafe.Pointer.
//
// A listener that waits for new incoming connections, configures them, and accepts or rejects them.
func XPCListenerFrom(ptr unsafe.Pointer) XPCListener {
	return XPCListener{objectivec.Object{objc.ID(ptr)}}
}

// Returns a new anonymous listener connection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/anonymous()
func (xc _XPCListenerClass) AnonymousListener() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(xc.class), objc.Sel("anonymousListener"))
	return rv
}


