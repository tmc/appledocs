// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NWHostEndpoint */


/* debug [class_header]: Header for NWHostEndpoint */
// The class instance for the [NWHostEndpoint] class.
var (
	NWHostEndpointClass     _NWHostEndpointClass
	NWHostEndpointClassOnce sync.Once
)

func getNWHostEndpointClass() _NWHostEndpointClass {
	NWHostEndpointClassOnce.Do(func() {
		NWHostEndpointClass = _NWHostEndpointClass{objc.GetClass("NWHostEndpoint")}
	})
	return NWHostEndpointClass
}

type _NWHostEndpointClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NWHostEndpoint */
// An interface definition for the [NWHostEndpoint] class.
type INWHostEndpoint interface {
	INWEndpoint
	
/* debug [class_interface_properties]: Properties for NWHostEndpoint */
	// properties:
	Hostname() objc.IObject /* cross-framework: NSString */
	Port() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NWHostEndpoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NWHostEndpoint */
// Alloc allocates a new instance without initialization.
func (nc _NWHostEndpointClass) Alloc() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NWHostEndpointClass) New() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWHostEndpoint) Init() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWHostEndpoint) Autorelease() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWHostEndpoint creates a new NWHostEndpoint instance.
func NewNWHostEndpoint() NWHostEndpoint {
	return getNWHostEndpointClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NWHostEndpoint */
// A network endpoint specified by DNS name (or IP address) and port.


// A network endpoint specified by DNS name (or IP address) and port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint
type NWHostEndpoint struct {
	NWEndpoint
}

// NWHostEndpointFrom constructs a [NWHostEndpoint] from an unsafe.Pointer.
//
// A network endpoint specified by DNS name (or IP address) and port.
func NWHostEndpointFrom(ptr unsafe.Pointer) NWHostEndpoint {
	return NWHostEndpoint{
		NWEndpoint: NWEndpointFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NWHostEndpoint */

// Create a host endpoint with a hostname and port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/init(hostname:port:)
func NewNWHostEndpointWithHostnamePort(hostname objc.IObject /* cross-framework: NSString */, port objc.IObject /* cross-framework: NSString */) NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](objc.ID(getNWHostEndpointClass().class), objc.Sel("endpointWithHostname:port:"), hostname, port)
	return rv
}/* debug [class_init_methods/constructor]: NewNWHostEndpointWithHostnamePort */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NWHostEndpoint */

// Create a host endpoint with a hostname and port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/init(hostname:port:)
func (nc _NWHostEndpointClass) EndpointWithHostnamePort(hostname objc.IObject /* cross-framework: NSString */, port objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("endpointWithHostname:port:"), hostname, port)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EndpointWithHostnamePort) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NWHostEndpoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NWHostEndpoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NWHostEndpoint */

// The endpoint’s hostname.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/hostname
func (n_ NWHostEndpoint) Hostname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("hostname"))
	return rv
}/* debug [instance_properties/getter]: hostname */


// The endpoint’s port, represented as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/port
func (n_ NWHostEndpoint) Port() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("port"))
	return rv
}/* debug [instance_properties/getter]: port */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NWHostEndpoint */


