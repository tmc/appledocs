// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XPCListenerEndpoint] class.
var (
	XPCListenerEndpointClass     _XPCListenerEndpointClass
	XPCListenerEndpointClassOnce sync.Once
)

func getXPCListenerEndpointClass() _XPCListenerEndpointClass {
	XPCListenerEndpointClassOnce.Do(func() {
		XPCListenerEndpointClass = _XPCListenerEndpointClass{objc.GetClass("NSXPCListenerEndpoint")}
	})
	return XPCListenerEndpointClass
}

type _XPCListenerEndpointClass struct {
	class objc.Class
}

// An interface definition for the [XPCListenerEndpoint] class.
type IXPCListenerEndpoint interface {
	objectivec.IObject
}

// An object that names a specific XPC listener.
//
// An instance of may be retrieved from an instance and sent over existing s. A process may then use the endpoint to create a new to the original . This pattern is useful if you have a service which multiplexes work to other services. The service can act as an intermediate helper. The requesting application does not need to know specifically which service it is connecting to, just that it implements a known .
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

// Alloc allocates a new instance without initialization.
func (xc _XPCListenerEndpointClass) Alloc() XPCListenerEndpoint {
	rv := objc.Send[XPCListenerEndpoint](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XPCListenerEndpointClass) New() XPCListenerEndpoint {
	rv := objc.Send[XPCListenerEndpoint](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XPCListenerEndpoint) Init() XPCListenerEndpoint {
	rv := objc.Send[XPCListenerEndpoint](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XPCListenerEndpoint) Autorelease() XPCListenerEndpoint {
	rv := objc.Send[XPCListenerEndpoint](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXPCListenerEndpoint creates a new XPCListenerEndpoint instance.
func NewXPCListenerEndpoint() XPCListenerEndpoint {
	return getXPCListenerEndpointClass().New()
}




