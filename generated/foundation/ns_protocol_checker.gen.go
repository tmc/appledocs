// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSProtocolChecker */


/* debug [class_header]: Header for NSProtocolChecker */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ProtocolChecker */
// An interface definition for the [ProtocolChecker] class.
type IProtocolChecker interface {
	IProxy
	
/* debug [class_interface_properties]: Properties for ProtocolChecker */
	// properties:
	Protocol() objectivec.Protocol
	Target() objc.IObject /* cross-framework: NSObject */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ProtocolChecker */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ProtocolChecker */
// Alloc allocates a new instance without initialization.
func (pc _ProtocolCheckerClass) Alloc() ProtocolChecker {
	rv := objc.Send[ProtocolChecker](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ProtocolChecker */
// An object that restricts the messages that can be sent to another object (referred to as the checker’s delegate).
//
// A object can be particularly useful when an object with many methods, only a few of which ought to be remotely accessible, is made available using the distributed objects system. A protocol checker acts as a kind of proxy; when it receives a message that is in its designated protocol, it forwards the message to its target and consequently appears to be the target object itself. However, when it receives a message not in its protocol, it raises an to indicate that the message isn’t allowed, whether or not the target object implements the method. Typically, an object that is to be distributed (yet must restrict messages) creates an for itself and returns the checker rather than returning itself in response to any messages. The object might also register the checker as the root object of an NSConnection. The object should be careful about vending references to —the protocol checker will convert a return value of to indicate the checker rather than the object for any messages forwarded by the checker, but direct references to the object (bypassing the checker) could be passed around by other objects.


// An object that restricts the messages that can be sent to another object (referred to as the checker’s delegate).
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ProtocolChecker */

// Initializes a newly allocated instance that will forward any messages in to , the protocol checker’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProtocolChecker/init(target:protocol:)
func NewProtocolCheckerWithTargetProtocol(anObject objc.IObject /* cross-framework: NSObject */, aProtocol objectivec.Protocol) ProtocolChecker {
	instance := getProtocolCheckerClass().Alloc()
	rv := objc.Send[ProtocolChecker](instance.ID, objc.Sel("initWithTarget:protocol:"), anObject, aProtocol)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewProtocolCheckerWithTargetProtocol */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ProtocolChecker */

// Allocates and initializes an instance that will forward any messages in to , the protocol checker’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProtocolChecker/protocolCheckerWithTarget:protocol:
func (pc _ProtocolCheckerClass) ProtocolCheckerWithTargetProtocol(anObject objc.IObject /* cross-framework: NSObject */, aProtocol objectivec.Protocol) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("protocolCheckerWithTarget:protocol:"), anObject, aProtocol)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ProtocolCheckerWithTargetProtocol) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ProtocolChecker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ProtocolChecker */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ProtocolChecker */

// Returns the protocol object the receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProtocolChecker/protocol
func (p_ ProtocolChecker) Protocol() objectivec.Protocol {
	rv := objc.Send[objectivec.Protocol](p_.ID, objc.Sel("protocol"))
	return rv
}/* debug [instance_properties/getter]: protocol */


// Returns the target of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProtocolChecker/target
func (p_ ProtocolChecker) Target() objc.IObject /* cross-framework: NSObject */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSProtocolChecker */


