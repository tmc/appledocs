// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NEDNSProxyProviderProtocol */


/* debug [class_header]: Header for NEDNSProxyProviderProtocol */
// The class instance for the [NEDNSProxyProviderProtocol] class.
var (
	NEDNSProxyProviderProtocolClass     _NEDNSProxyProviderProtocolClass
	NEDNSProxyProviderProtocolClassOnce sync.Once
)

func getNEDNSProxyProviderProtocolClass() _NEDNSProxyProviderProtocolClass {
	NEDNSProxyProviderProtocolClassOnce.Do(func() {
		NEDNSProxyProviderProtocolClass = _NEDNSProxyProviderProtocolClass{objc.GetClass("NEDNSProxyProviderProtocol")}
	})
	return NEDNSProxyProviderProtocolClass
}

type _NEDNSProxyProviderProtocolClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEDNSProxyProviderProtocol */
// An interface definition for the [NEDNSProxyProviderProtocol] class.
type INEDNSProxyProviderProtocol interface {
	INEVPNProtocol
	
/* debug [class_interface_properties]: Properties for NEDNSProxyProviderProtocol */
	// properties:
	ProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */
	SetProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */)
	ProviderConfiguration() foundation.IDictionary
	SetProviderConfiguration(value foundation.IDictionary)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEDNSProxyProviderProtocol */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEDNSProxyProviderProtocol */
// Alloc allocates a new instance without initialization.
func (nc _NEDNSProxyProviderProtocolClass) Alloc() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEDNSProxyProviderProtocolClass) New() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSProxyProviderProtocol) Init() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSProxyProviderProtocol) Autorelease() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSProxyProviderProtocol creates a new NEDNSProxyProviderProtocol instance.
func NewNEDNSProxyProviderProtocol() NEDNSProxyProviderProtocol {
	return getNEDNSProxyProviderProtocolClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEDNSProxyProviderProtocol */
// Configuration parameters for a DNS proxy.


// Configuration parameters for a DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol
type NEDNSProxyProviderProtocol struct {
	NEVPNProtocol
}

// NEDNSProxyProviderProtocolFrom constructs a [NEDNSProxyProviderProtocol] from an unsafe.Pointer.
//
// Configuration parameters for a DNS proxy.
func NEDNSProxyProviderProtocolFrom(ptr unsafe.Pointer) NEDNSProxyProviderProtocol {
	return NEDNSProxyProviderProtocol{
		NEVPNProtocol: NEVPNProtocolFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEDNSProxyProviderProtocol *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEDNSProxyProviderProtocol */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEDNSProxyProviderProtocol */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEDNSProxyProviderProtocol */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEDNSProxyProviderProtocol */

// A string containing the bundle identifier of the proxy provider to be used by this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerBundleIdentifier
func (n_ NEDNSProxyProviderProtocol) ProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: providerBundleIdentifier */


// A string containing the bundle identifier of the proxy provider to be used by this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerBundleIdentifier
func (n_ NEDNSProxyProviderProtocol) SetProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), value)
}/* debug [instance_properties/setter]: providerBundleIdentifier */


// A dictionary containing vendor-specific configuration parameters for a proxy provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerConfiguration
func (n_ NEDNSProxyProviderProtocol) ProviderConfiguration() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}/* debug [instance_properties/getter]: providerConfiguration */


// A dictionary containing vendor-specific configuration parameters for a proxy provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerConfiguration
func (n_ NEDNSProxyProviderProtocol) SetProviderConfiguration(value foundation.IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), value)
}/* debug [instance_properties/setter]: providerConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEDNSProxyProviderProtocol */



