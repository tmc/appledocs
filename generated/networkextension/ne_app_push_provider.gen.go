// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NEAppPushProvider */


/* debug [class_header]: Header for NEAppPushProvider */
// The class instance for the [NEAppPushProvider] class.
var (
	NEAppPushProviderClass     _NEAppPushProviderClass
	NEAppPushProviderClassOnce sync.Once
)

func getNEAppPushProviderClass() _NEAppPushProviderClass {
	NEAppPushProviderClassOnce.Do(func() {
		NEAppPushProviderClass = _NEAppPushProviderClass{objc.GetClass("NEAppPushProvider")}
	})
	return NEAppPushProviderClass
}

type _NEAppPushProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEAppPushProvider */
// An interface definition for the [NEAppPushProvider] class.
type INEAppPushProvider interface {
	INEProvider
	
/* debug [class_interface_properties]: Properties for NEAppPushProvider */
	// properties:
	Delegate() objc.IObject /* cross-framework: NEAppPushDelegate */
	SetDelegate(value objc.IObject /* cross-framework: NEAppPushDelegate */)
	ProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */
	SetProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEAppPushProvider */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEAppPushProvider */
// Alloc allocates a new instance without initialization.
func (nc _NEAppPushProviderClass) Alloc() NEAppPushProvider {
	rv := objc.Send[NEAppPushProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEAppPushProviderClass) New() NEAppPushProvider {
	rv := objc.Send[NEAppPushProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppPushProvider) Init() NEAppPushProvider {
	rv := objc.Send[NEAppPushProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppPushProvider) Autorelease() NEAppPushProvider {
	rv := objc.Send[NEAppPushProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppPushProvider creates a new NEAppPushProvider instance.
func NewNEAppPushProvider() NEAppPushProvider {
	return getNEAppPushProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEAppPushProvider */
// An object that creates and maintains a persistent network connection to a local push server.
//
// Subclass to provide the connection to your local push server. A creates instances of your provider class based on the in the manager’s configuration. The manager then calls methods on your provider to start and stop communication with the server, and periodically check the provider’s status. When your provider receives an incoming call from your server, call the provider’s method to alert the manager’s .


// An object that creates and maintains a persistent network connection to a local push server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider
type NEAppPushProvider struct {
	NEProvider
}

// NEAppPushProviderFrom constructs a [NEAppPushProvider] from an unsafe.Pointer.
//
// An object that creates and maintains a persistent network connection to a local push server.
func NEAppPushProviderFrom(ptr unsafe.Pointer) NEAppPushProvider {
	return NEAppPushProvider{
		NEProvider: NEProviderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEAppPushProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEAppPushProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEAppPushProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEAppPushProvider */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEAppPushProvider */

// A delegate that receives incoming call information from the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/delegate
func (n_ NEAppPushProvider) Delegate() objc.IObject /* cross-framework: NEAppPushDelegate */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A delegate that receives incoming call information from the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/delegate
func (n_ NEAppPushProvider) SetDelegate(value objc.IObject /* cross-framework: NEAppPushDelegate */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A string that contains the bundle identifier of the push provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/providerbundleidentifier
func (n_ NEAppPushProvider) ProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: providerBundleIdentifier */


// A string that contains the bundle identifier of the push provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/providerbundleidentifier
func (n_ NEAppPushProvider) SetProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), value)
}/* debug [instance_properties/setter]: providerBundleIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEAppPushProvider */


