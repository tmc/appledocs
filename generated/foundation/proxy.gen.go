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

// An interface definition for the [Proxy] class.
type IProxy interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (pc _ProxyClass) Alloc() Proxy {
	rv := objc.Send[Proxy](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _ProxyClass) New() Proxy {
	rv := objc.Send[Proxy](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Proxy) Init() Proxy {
	rv := objc.Send[Proxy](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Proxy) Autorelease() Proxy {
	rv := objc.Send[Proxy](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProxy creates a new Proxy instance.
func NewProxy() Proxy {
	return proxyClass.New()
}




