// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSXPCListenerEndpoint */


/* debug [class_header]: Header for NSXPCListenerEndpoint */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for XPCListenerEndpoint */
// An interface definition for the [XPCListenerEndpoint] class.
type IXPCListenerEndpoint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for XPCListenerEndpoint */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for XPCListenerEndpoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for XPCListenerEndpoint */
// Alloc allocates a new instance without initialization.
func (xc _XPCListenerEndpointClass) Alloc() XPCListenerEndpoint {
	rv := objc.Send[XPCListenerEndpoint](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for XPCListenerEndpoint */
// An object that names a specific XPC listener.
//
// An instance of may be retrieved from an instance and sent over existing s. A process may then use the endpoint to create a new to the original . This pattern is useful if you have a service which multiplexes work to other services. The service can act as an intermediate helper. The requesting application does not need to know specifically which service it is connecting to, just that it implements a known .


// An object that names a specific XPC listener.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for XPCListenerEndpoint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for XPCListenerEndpoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for XPCListenerEndpoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for XPCListenerEndpoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for XPCListenerEndpoint */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSXPCListenerEndpoint */



