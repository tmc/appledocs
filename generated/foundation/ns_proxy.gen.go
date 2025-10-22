// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Proxy] class.
var (
	ProxyClass     _ProxyClass
	ProxyClassOnce sync.Once
)

func getProxyClass() _ProxyClass {
	ProxyClassOnce.Do(func() {
		ProxyClass = _ProxyClass{objc.GetClass("NSProxy")}
	})
	return ProxyClass
}

type _ProxyClass struct {
	class objc.Class
}

// An interface definition for the [Proxy] class.
type IProxy interface {
	objectivec.IObject
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type Proxy struct {
	objectivec.Object
}

// ProxyFrom constructs a [Proxy] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func ProxyFrom(ptr unsafe.Pointer) Proxy {
	return Proxy{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ProxyClass) Alloc() Proxy {
	rv := objc.Send[Proxy](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getProxyClass().New()
}




