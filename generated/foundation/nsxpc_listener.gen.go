// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSXPCListener */


/* debug [class_header]: Header for NSXPCListener */
// The class instance for the [XPCListener] class.
var (
	XPCListenerClass     _XPCListenerClass
	XPCListenerClassOnce sync.Once
)

func getXPCListenerClass() _XPCListenerClass {
	XPCListenerClassOnce.Do(func() {
		XPCListenerClass = _XPCListenerClass{objc.GetClass("NSXPCListener")}
	})
	return XPCListenerClass
}

type _XPCListenerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for XPCListener */
// An interface definition for the [XPCListener] class.
type IXPCListener interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for XPCListener */
	// properties:
	Endpoint() IXPCListenerEndpoint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for XPCListener */
	// methods:
	Activate()
	Invalidate()
	Resume()
	SetConnectionCodeSigningRequirement(requirement IString)
	Suspend()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for XPCListener */
// Alloc allocates a new instance without initialization.
func (xc _XPCListenerClass) Alloc() XPCListener {
	rv := objc.Send[XPCListener](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (xc _XPCListenerClass) New() XPCListener {
	rv := objc.Send[XPCListener](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XPCListener) Init() XPCListener {
	rv := objc.Send[XPCListener](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XPCListener) Autorelease() XPCListener {
	rv := objc.Send[XPCListener](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXPCListener creates a new XPCListener instance.
func NewXPCListener() XPCListener {
	return getXPCListenerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for XPCListener */
// A listener that waits for new incoming connections, configures them, and accepts or rejects them.
//
// Each XPC service, launchd agent, or launchd daemon typically has at least one object that listens for connections to a specified service name. Each listener must have a delegate that conforms to the protocol. When the listener receives a new connection request, it creates a new object, then asks the delegate to inspect, configure, and resume the connection object by calling the delegate’s method.


// A listener that waits for new incoming connections, configures them, and accepts or rejects them.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for XPCListener */

// Initializes a listener in a LaunchAgent or LaunchDaemon which has a name advertised in a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/init(machServiceName:)
func NewXPCListenerWithMachServiceName(name IString) XPCListener {
	instance := getXPCListenerClass().Alloc()
	rv := objc.Send[XPCListener](instance.ID, objc.Sel("initWithMachServiceName:"), name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewXPCListenerWithMachServiceName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for XPCListener */

// Returns a new anonymous listener connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/anonymous()
func (xc _XPCListenerClass) AnonymousListener() IXPCListener {
	rv := objc.Send[XPCListener](objc.ID(xc.class), objc.Sel("anonymousListener"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AnonymousListener) */


// Returns the singleton listener used to listen for incoming connections in an XPC service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/service()
func (xc _XPCListenerClass) ServiceListener() IXPCListener {
	rv := objc.Send[XPCListener](objc.ID(xc.class), objc.Sel("serviceListener"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ServiceListener) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for XPCListener */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for XPCListener */

// Activates the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/activate()
func (x_ XPCListener) Activate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("activate"))
}/* debug [instance_methods/method]: Activate */


// Invalidates the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/invalidate()
func (x_ XPCListener) Invalidate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("invalidate"))
}/* debug [instance_methods/method]: Invalidate */


// Starts processing of incoming requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/resume()
func (x_ XPCListener) Resume() {
	objc.Send[objc.ID](x_.ID, objc.Sel("resume"))
}/* debug [instance_methods/method]: Resume */


// Sets the code signing requirement for connections to this listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/setConnectionCodeSigningRequirement(_:)
func (x_ XPCListener) SetConnectionCodeSigningRequirement(requirement IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setConnectionCodeSigningRequirement:"), requirement)
}/* debug [instance_methods/method]: SetConnectionCodeSigningRequirement */


// Suspends the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/suspend()
func (x_ XPCListener) Suspend() {
	objc.Send[objc.ID](x_.ID, objc.Sel("suspend"))
}/* debug [instance_methods/method]: Suspend */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for XPCListener */

// Returns an endpoint object that may be sent over an existing connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/endpoint
func (x_ XPCListener) Endpoint() IXPCListenerEndpoint {
	rv := objc.Send[XPCListenerEndpoint](x_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSXPCListener */


