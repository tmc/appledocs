// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ProtocolChecker] class.
var (
	ProtocolCheckerClass     _ProtocolCheckerClass
	ProtocolCheckerClassOnce sync.Once
)

func getProtocolCheckerClass() _ProtocolCheckerClass {
	ProtocolCheckerClassOnce.Do(func() {
		ProtocolCheckerClass = _ProtocolCheckerClass{objc.GetClass("NSProtocolChecker")}
	})
	return ProtocolCheckerClass
}

type _ProtocolCheckerClass struct {
	class objc.Class
}

// An interface definition for the [ProtocolChecker] class.
type IProtocolChecker interface {
	IProxy
}

// An object that restricts the messages that can be sent to another object (referred to as the checker’s delegate).
//
// A object can be particularly useful when an object with many methods, only a few of which ought to be remotely accessible, is made available using the distributed objects system. A protocol checker acts as a kind of proxy; when it receives a message that is in its designated protocol, it forwards the message to its target and consequently appears to be the target object itself. However, when it receives a message not in its protocol, it raises an to indicate that the message isn’t allowed, whether or not the target object implements the method. Typically, an object that is to be distributed (yet must restrict messages) creates an for itself and returns the checker rather than returning itself in response to any messages. The object might also register the checker as the root object of an NSConnection. The object should be careful about vending references to —the protocol checker will convert a return value of to indicate the checker rather than the object for any messages forwarded by the checker, but direct references to the object (bypassing the checker) could be passed around by other objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProtocolChecker
type ProtocolChecker struct {
	Proxy
}

// ProtocolCheckerFrom constructs a [ProtocolChecker] from an unsafe.Pointer.
//
// An object that restricts the messages that can be sent to another object (referred to as the checker’s delegate).
func ProtocolCheckerFrom(ptr unsafe.Pointer) ProtocolChecker {
	return ProtocolChecker{
		Proxy: ProxyFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _ProtocolCheckerClass) Alloc() ProtocolChecker {
	rv := objc.Send[ProtocolChecker](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ProtocolCheckerClass) New() ProtocolChecker {
	rv := objc.Send[ProtocolChecker](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProtocolChecker) Init() ProtocolChecker {
	rv := objc.Send[ProtocolChecker](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProtocolChecker) Autorelease() ProtocolChecker {
	rv := objc.Send[ProtocolChecker](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProtocolChecker creates a new ProtocolChecker instance.
func NewProtocolChecker() ProtocolChecker {
	return getProtocolCheckerClass().New()
}


// Returns the protocol object the receiver uses.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprotocolchecker/protocol
func (p_ ProtocolChecker) Protocol() objectivec.Protocol {
	rv := objc.Send[objectivec.Protocol](p_.ID, objc.Sel("protocol"))
	return rv
}


// SetProtocol sets the value of the protocol property.
// Returns the protocol object the receiver uses.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprotocolchecker/protocol
func (p_ ProtocolChecker) SetProtocol(value objectivec.Protocol) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProtocol:"), value)
}

// Returns the target of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprotocolchecker/target
func (p_ ProtocolChecker) Target() NSObject {
	rv := objc.Send[NSObject](p_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// Returns the target of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprotocolchecker/target
func (p_ ProtocolChecker) SetTarget(value IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTarget:"), value)
}



