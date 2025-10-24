// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDistantObject */


/* debug [class_header]: Header for NSDistantObject */
// The class instance for the [DistantObject] class.
var (
	DistantObjectClass     _DistantObjectClass
	DistantObjectClassOnce sync.Once
)

func getDistantObjectClass() _DistantObjectClass {
	DistantObjectClassOnce.Do(func() {
		DistantObjectClass = _DistantObjectClass{objc.GetClass("NSDistantObject")}
	})
	return DistantObjectClass
}

type _DistantObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DistantObject */
// An interface definition for the [DistantObject] class.
type IDistantObject interface {
	IProxy
	
/* debug [class_interface_properties]: Properties for DistantObject */
	// properties:
	ConnectionForProxy() IConnection
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DistantObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DistantObject */
// Alloc allocates a new instance without initialization.
func (dc _DistantObjectClass) Alloc() DistantObject {
	rv := objc.Send[DistantObject](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DistantObjectClass) New() DistantObject {
	rv := objc.Send[DistantObject](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DistantObject) Init() DistantObject {
	rv := objc.Send[DistantObject](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DistantObject) Autorelease() DistantObject {
	rv := objc.Send[DistantObject](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDistantObject creates a new DistantObject instance.
func NewDistantObject() DistantObject {
	return getDistantObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DistantObject */
// A proxy for objects in other applications or threads.
//
// When a distant object receives a message, in most cases it forwards the message through its object to the real object in another application, supplying the return value to the sender of the message if one is received, and propagating any exception back to the invoker of the method that raised it. is a concrete subclass of , adding two useful instance methods of its own: returns the object that handles the receiver; establishes the set of methods the real object is known to respond to, saving the network traffic required to determine the argument and return types the first time a particular selector is forwarded to the remote proxy. There are two kinds of distant object: local proxies and remote proxies. A local proxy is created by an object the first time an object is sent to another application. It is used by the connection for bookkeeping purposes and should be considered private. The local proxy is transmitted over the network using the protocol to create the remote proxy, which is the object that the other application uses. defines methods for an object to create instances, but they’re intended only for subclasses to override—you should never invoke them directly. Use the method of , which sets up all the required state for an object-proxy pair.


// A proxy for objects in other applications or threads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObject
type DistantObject struct {
	Proxy
}

// DistantObjectFrom constructs a [DistantObject] from an unsafe.Pointer.
//
// A proxy for objects in other applications or threads.
func DistantObjectFrom(ptr unsafe.Pointer) DistantObject {
	return DistantObject{
		Proxy: ProxyFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DistantObject */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObject/initWithCoder:
func NewDistantObjectWithCoder(inCoder ICoder) DistantObject {
	instance := getDistantObjectClass().Alloc()
	rv := objc.Send[DistantObject](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDistantObjectWithCoder */


// Initializes an object as a local proxy for a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObject/initWithLocal:connection:
func NewDistantObjectWithLocalConnection(target objc.IObject, connection IConnection) DistantObject {
	instance := getDistantObjectClass().Alloc()
	rv := objc.Send[DistantObject](instance.ID, objc.Sel("initWithLocal:connection:"), target, connection)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDistantObjectWithLocalConnection */


// Initializes a newly allocated NSDistantObject as a remote proxy for , which is an id in another thread or another application’s address space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObject/initWithTarget:connection:
func NewDistantObjectWithTargetConnection(target objc.IObject, connection IConnection) DistantObject {
	instance := getDistantObjectClass().Alloc()
	rv := objc.Send[DistantObject](instance.ID, objc.Sel("initWithTarget:connection:"), target, connection)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDistantObjectWithTargetConnection */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DistantObject */

// Returns a local proxy for a given object and connection, creating the proxy if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObject/proxyWithLocal:connection:
func (dc _DistantObjectClass) ProxyWithLocalConnection(target objc.IObject, connection IConnection) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("proxyWithLocal:connection:"), target, connection)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ProxyWithLocalConnection) */


// Returns a remote proxy for a given object and connection, creating the proxy if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObject/proxyWithTarget:connection:
func (dc _DistantObjectClass) ProxyWithTargetConnection(target objc.IObject, connection IConnection) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("proxyWithTarget:connection:"), target, connection)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ProxyWithTargetConnection) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DistantObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DistantObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DistantObject */

// Returns the connection used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObject/connectionForProxy
func (d_ DistantObject) ConnectionForProxy() IConnection {
	rv := objc.Send[Connection](d_.ID, objc.Sel("connectionForProxy"))
	return rv
}/* debug [instance_properties/getter]: connectionForProxy */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDistantObject */


