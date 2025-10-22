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
	DebugDescription() string
	SetDebugDescription(value string)
	Description() string
	SetDescription(value string)
}

// An abstract superclass defining an API for objects that act as stand-ins for other objects or for objects that don’t exist yet.
//
// Typically, a message to a proxy is forwarded to the real object or causes the proxy to load (or transform itself into) the real object. Subclasses of can be used to implement transparent distributed messaging (for example, ) or for lazy instantiation of objects that are expensive to create. implements the basic methods required of a root class, including those defined in the protocol. However, as an abstract class it doesn’t provide an initialization method, and it raises an exception upon receiving any message it doesn’t respond to. A concrete subclass must therefore provide an initialization or creation method and override the and methods to handle messages that it doesn’t implement itself. A subclass’s implementation of should do whatever is needed to process the invocation, such as forwarding the invocation over the network or loading the real object and passing it the invocation. is required to provide argument type information for a given message; a subclass’s implementation should be able to determine the argument types for the messages it needs to forward and should construct an object accordingly. See the , , and class specifications for more information.


// An abstract superclass defining an API for objects that act as stand-ins for other objects or for objects that don’t exist yet.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsproxy/debugdescription

func (p_ Proxy) DebugDescription() string {
	rv := objc.Send[string](p_.ID, objc.Sel("debugDescription"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsproxy/debugdescription

func (p_ Proxy) SetDebugDescription(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDebugDescription:"), objc.String(value))
}


// A string containing the real class name and the id of the receiver as a hexadecimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsproxy/description

func (p_ Proxy) Description() string {
	rv := objc.Send[string](p_.ID, objc.Sel("description"))
	return rv
}


// A string containing the real class name and the id of the receiver as a hexadecimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsproxy/description

func (p_ Proxy) SetDescription(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDescription:"), objc.String(value))
}



