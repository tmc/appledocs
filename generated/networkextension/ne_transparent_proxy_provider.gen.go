// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NETransparentProxyProvider */


/* debug [class_header]: Header for NETransparentProxyProvider */
// The class instance for the [NETransparentProxyProvider] class.
var (
	NETransparentProxyProviderClass     _NETransparentProxyProviderClass
	NETransparentProxyProviderClassOnce sync.Once
)

func getNETransparentProxyProviderClass() _NETransparentProxyProviderClass {
	NETransparentProxyProviderClassOnce.Do(func() {
		NETransparentProxyProviderClass = _NETransparentProxyProviderClass{objc.GetClass("NETransparentProxyProvider")}
	})
	return NETransparentProxyProviderClass
}

type _NETransparentProxyProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NETransparentProxyProvider */
// An interface definition for the [NETransparentProxyProvider] class.
type INETransparentProxyProvider interface {
	INEAppProxyProvider
	
/* debug [class_interface_properties]: Properties for NETransparentProxyProvider */
	// properties:
	IncludedNetworkRules() INENetworkRule
	SetIncludedNetworkRules(value INENetworkRule)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NETransparentProxyProvider */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NETransparentProxyProvider */
// Alloc allocates a new instance without initialization.
func (nc _NETransparentProxyProviderClass) Alloc() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NETransparentProxyProviderClass) New() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETransparentProxyProvider) Init() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETransparentProxyProvider) Autorelease() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETransparentProxyProvider creates a new NETransparentProxyProvider instance.
func NewNETransparentProxyProvider() NETransparentProxyProvider {
	return getNETransparentProxyProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NETransparentProxyProvider */
// An object that implements the client side of a custom transparent network proxy solution.
//
// The class has the following behavior differences from its superclass : Returning from and causes the flow to proceed to communicate directly with the flow’s ultimate destination, instead of closing the flow with a “Connection Refused” error. This provider ignores and specified within . Flows that match the within use the same DNS and proxy settings that other flows on the system currently use. Flows that are created using a “connect by name” API (such as framework or ) that match the don’t bypass DNS resolution.


// An object that implements the client side of a custom transparent network proxy solution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyProvider
type NETransparentProxyProvider struct {
	NEAppProxyProvider
}

// NETransparentProxyProviderFrom constructs a [NETransparentProxyProvider] from an unsafe.Pointer.
//
// An object that implements the client side of a custom transparent network proxy solution.
func NETransparentProxyProviderFrom(ptr unsafe.Pointer) NETransparentProxyProvider {
	return NETransparentProxyProvider{
		NEAppProxyProvider: NEAppProxyProviderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NETransparentProxyProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NETransparentProxyProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NETransparentProxyProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NETransparentProxyProvider */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NETransparentProxyProvider */

// An array of rules that collectively specify what traffic to route through the transparent proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netransparentproxynetworksettings/includednetworkrules
func (n_ NETransparentProxyProvider) IncludedNetworkRules() INENetworkRule {
	rv := objc.Send[NENetworkRule](n_.ID, objc.Sel("includedNetworkRules"))
	return rv
}/* debug [instance_properties/getter]: includedNetworkRules */


// An array of rules that collectively specify what traffic to route through the transparent proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netransparentproxynetworksettings/includednetworkrules
func (n_ NETransparentProxyProvider) SetIncludedNetworkRules(value INENetworkRule) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedNetworkRules:"), value)
}/* debug [instance_properties/setter]: includedNetworkRules */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NETransparentProxyProvider */



