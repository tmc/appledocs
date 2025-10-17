// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Proxy] class.
var proxyClass = _ProxyClass{objc.GetClass("NSProxy")}

type _ProxyClass struct {
	class objc.Class
}

// An abstract superclass defining an API for objects that act as stand-ins for other objects or for objects that don’t exist yet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy

type Proxy struct {
	objectivec.Object
}

// ProxyFrom constructs a [Proxy] from an unsafe.Pointer.
//
// An abstract superclass defining an API for objects that act as stand-ins for other objects or for objects that don’t exist yet.
func ProxyFrom(ptr unsafe.Pointer) Proxy {
	return Proxy{objectivec.Object{objc.ID(ptr)}}
}



