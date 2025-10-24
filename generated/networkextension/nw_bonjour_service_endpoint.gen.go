// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NWBonjourServiceEndpoint */


/* debug [class_header]: Header for NWBonjourServiceEndpoint */
// The class instance for the [NWBonjourServiceEndpoint] class.
var (
	NWBonjourServiceEndpointClass     _NWBonjourServiceEndpointClass
	NWBonjourServiceEndpointClassOnce sync.Once
)

func getNWBonjourServiceEndpointClass() _NWBonjourServiceEndpointClass {
	NWBonjourServiceEndpointClassOnce.Do(func() {
		NWBonjourServiceEndpointClass = _NWBonjourServiceEndpointClass{objc.GetClass("NWBonjourServiceEndpoint")}
	})
	return NWBonjourServiceEndpointClass
}

type _NWBonjourServiceEndpointClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NWBonjourServiceEndpoint */
// An interface definition for the [NWBonjourServiceEndpoint] class.
type INWBonjourServiceEndpoint interface {
	INWEndpoint
	
/* debug [class_interface_properties]: Properties for NWBonjourServiceEndpoint */
	// properties:
	Domain() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	Type() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NWBonjourServiceEndpoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NWBonjourServiceEndpoint */
// Alloc allocates a new instance without initialization.
func (nc _NWBonjourServiceEndpointClass) Alloc() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NWBonjourServiceEndpointClass) New() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWBonjourServiceEndpoint) Init() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWBonjourServiceEndpoint) Autorelease() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWBonjourServiceEndpoint creates a new NWBonjourServiceEndpoint instance.
func NewNWBonjourServiceEndpoint() NWBonjourServiceEndpoint {
	return getNWBonjourServiceEndpointClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NWBonjourServiceEndpoint */
// A network endpoint specified as a Bonjour service name, type, and domain.
//
// For example, the Bonjour service has the name , the type , and the domain .


// A network endpoint specified as a Bonjour service name, type, and domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint
type NWBonjourServiceEndpoint struct {
	NWEndpoint
}

// NWBonjourServiceEndpointFrom constructs a [NWBonjourServiceEndpoint] from an unsafe.Pointer.
//
// A network endpoint specified as a Bonjour service name, type, and domain.
func NWBonjourServiceEndpointFrom(ptr unsafe.Pointer) NWBonjourServiceEndpoint {
	return NWBonjourServiceEndpoint{
		NWEndpoint: NWEndpointFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NWBonjourServiceEndpoint */

// Create an endpoint with a Bonjour service name, type, and domain. All fields must be specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/init(name:type:domain:)
func NewNWBonjourServiceEndpointWithNameTypeDomain(name objc.IObject /* cross-framework: NSString */, type_ objc.IObject /* cross-framework: NSString */, domain objc.IObject /* cross-framework: NSString */) NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](objc.ID(getNWBonjourServiceEndpointClass().class), objc.Sel("endpointWithName:type:domain:"), name, type_, domain)
	return rv
}/* debug [class_init_methods/constructor]: NewNWBonjourServiceEndpointWithNameTypeDomain */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NWBonjourServiceEndpoint */

// Create an endpoint with a Bonjour service name, type, and domain. All fields must be specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/init(name:type:domain:)
func (nc _NWBonjourServiceEndpointClass) EndpointWithNameTypeDomain(name objc.IObject /* cross-framework: NSString */, type_ objc.IObject /* cross-framework: NSString */, domain objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("endpointWithName:type:domain:"), name, type_, domain)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EndpointWithNameTypeDomain) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NWBonjourServiceEndpoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NWBonjourServiceEndpoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NWBonjourServiceEndpoint */

// The endpoint’s Bonjour service domain, such as .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/domain
func (n_ NWBonjourServiceEndpoint) Domain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("domain"))
	return rv
}/* debug [instance_properties/getter]: domain */


// The endpoint’s Bonjour service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/name
func (n_ NWBonjourServiceEndpoint) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The endpoint’s Bonjour service type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/type
func (n_ NWBonjourServiceEndpoint) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NWBonjourServiceEndpoint */


