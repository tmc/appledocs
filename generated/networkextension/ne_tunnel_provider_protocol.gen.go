// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NETunnelProviderProtocol */


/* debug [class_header]: Header for NETunnelProviderProtocol */
// The class instance for the [NETunnelProviderProtocol] class.
var (
	NETunnelProviderProtocolClass     _NETunnelProviderProtocolClass
	NETunnelProviderProtocolClassOnce sync.Once
)

func getNETunnelProviderProtocolClass() _NETunnelProviderProtocolClass {
	NETunnelProviderProtocolClassOnce.Do(func() {
		NETunnelProviderProtocolClass = _NETunnelProviderProtocolClass{objc.GetClass("NETunnelProviderProtocol")}
	})
	return NETunnelProviderProtocolClass
}

type _NETunnelProviderProtocolClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NETunnelProviderProtocol */
// An interface definition for the [NETunnelProviderProtocol] class.
type INETunnelProviderProtocol interface {
	INEVPNProtocol
	
/* debug [class_interface_properties]: Properties for NETunnelProviderProtocol */
	// properties:
	ProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */
	SetProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */)
	ProviderConfiguration() foundation.IDictionary
	SetProviderConfiguration(value foundation.IDictionary)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NETunnelProviderProtocol */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NETunnelProviderProtocol */
// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderProtocolClass) Alloc() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NETunnelProviderProtocolClass) New() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETunnelProviderProtocol) Init() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETunnelProviderProtocol) Autorelease() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProviderProtocol creates a new NETunnelProviderProtocol instance.
func NewNETunnelProviderProtocol() NETunnelProviderProtocol {
	return getNETunnelProviderProtocolClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NETunnelProviderProtocol */
// Configuration parameters for a VPN tunnel.
//
// objects are used to specify configuration parameters for Tunnel Provider extensions.


// Configuration parameters for a VPN tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol
type NETunnelProviderProtocol struct {
	NEVPNProtocol
}

// NETunnelProviderProtocolFrom constructs a [NETunnelProviderProtocol] from an unsafe.Pointer.
//
// Configuration parameters for a VPN tunnel.
func NETunnelProviderProtocolFrom(ptr unsafe.Pointer) NETunnelProviderProtocol {
	return NETunnelProviderProtocol{
		NEVPNProtocol: NEVPNProtocolFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NETunnelProviderProtocol *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NETunnelProviderProtocol */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NETunnelProviderProtocol */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NETunnelProviderProtocol */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NETunnelProviderProtocol */

// A string identifying the specific Tunnel Provider extension that should be used with this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerBundleIdentifier
func (n_ NETunnelProviderProtocol) ProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: providerBundleIdentifier */


// A string identifying the specific Tunnel Provider extension that should be used with this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerBundleIdentifier
func (n_ NETunnelProviderProtocol) SetProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), value)
}/* debug [instance_properties/setter]: providerBundleIdentifier */


// A dictionary containing keys and values defined by the Tunnel Provider developer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerConfiguration
func (n_ NETunnelProviderProtocol) ProviderConfiguration() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}/* debug [instance_properties/getter]: providerConfiguration */


// A dictionary containing keys and values defined by the Tunnel Provider developer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerConfiguration
func (n_ NETunnelProviderProtocol) SetProviderConfiguration(value foundation.IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), value)
}/* debug [instance_properties/setter]: providerConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NETunnelProviderProtocol */



