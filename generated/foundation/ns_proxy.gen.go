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
	// properties:
	DebugDescription() IString
	Description() IString
	// methods:
	AllowsWeakReference() bool
	Dealloc()
	Finalize()
	ForwardInvocation(invocation IInvocation)
	MethodSignatureForSelector(sel objc.SEL) MethodSignature /* not a class type */
	RetainWeakReference() bool
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



// Returns (the class object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/class()
func (pc _ProxyClass) Class() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(pc.class), objc.Sel("class"))
	return rv
}


// Returns a Boolean value that indicates whether the receiving class responds to a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/responds(to:)
func (pc _ProxyClass) RespondsToSelector(aSelector objc.SEL) bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("respondsToSelector:"), aSelector)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/allowsWeakReference
func (p_ Proxy) AllowsWeakReference() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsWeakReference"))
	return rv
}


// Deallocates the memory occupied by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/dealloc()
func (p_ Proxy) Dealloc() {
	objc.Send[objc.ID](p_.ID, objc.Sel("dealloc"))
}


// The garbage collector invokes this method on the receiver before disposing of the memory it uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/finalize()
func (p_ Proxy) Finalize() {
	objc.Send[objc.ID](p_.ID, objc.Sel("finalize"))
}


// Passes a given invocation to the real object the proxy represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/forwardInvocation(_:)
func (p_ Proxy) ForwardInvocation(invocation IInvocation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("forwardInvocation:"), invocation)
}


// Raises . Override this method in your concrete subclass to return a proper object for the given selector and the class your proxy objects stand in for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/methodSignatureForSelector:
func (p_ Proxy) MethodSignatureForSelector(sel objc.SEL) MethodSignature /* not a class type */ {
	rv := objc.Send[MethodSignature](p_.ID, objc.Sel("methodSignatureForSelector:"), sel)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/retainWeakReference
func (p_ Proxy) RetainWeakReference() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("retainWeakReference"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/debugDescription
func (p_ Proxy) DebugDescription() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("debugDescription"))
	return rv
}


// A string containing the real class name and the id of the receiver as a hexadecimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/description
func (p_ Proxy) Description() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("description"))
	return rv
}



