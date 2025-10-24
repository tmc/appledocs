// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NEAppProxyProviderManager */


/* debug [class_header]: Header for NEAppProxyProviderManager */
// The class instance for the [NEAppProxyProviderManager] class.
var (
	NEAppProxyProviderManagerClass     _NEAppProxyProviderManagerClass
	NEAppProxyProviderManagerClassOnce sync.Once
)

func getNEAppProxyProviderManagerClass() _NEAppProxyProviderManagerClass {
	NEAppProxyProviderManagerClassOnce.Do(func() {
		NEAppProxyProviderManagerClass = _NEAppProxyProviderManagerClass{objc.GetClass("NEAppProxyProviderManager")}
	})
	return NEAppProxyProviderManagerClass
}

type _NEAppProxyProviderManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEAppProxyProviderManager */
// An interface definition for the [NEAppProxyProviderManager] class.
type INEAppProxyProviderManager interface {
	INETunnelProviderManager
	
/* debug [class_interface_properties]: Properties for NEAppProxyProviderManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEAppProxyProviderManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEAppProxyProviderManager */
// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyProviderManagerClass) Alloc() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEAppProxyProviderManagerClass) New() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyProviderManager) Init() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyProviderManager) Autorelease() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyProviderManager creates a new NEAppProxyProviderManager instance.
func NewNEAppProxyProviderManager() NEAppProxyProviderManager {
	return getNEAppProxyProviderManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEAppProxyProviderManager */
// An object to create and manage the app proxy provider’s VPN configuration.
//
// Objects cannot be directly instantiated. Instead, App Proxy configurations are created exclusively from payloads in configuration profiles. App Proxy configurations can only be used with Per-App VPN routing rules. For more details about how to create App Proxy configurations and configure Per-App VPN, see .


// An object to create and manage the app proxy provider’s VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProviderManager
type NEAppProxyProviderManager struct {
	NETunnelProviderManager
}

// NEAppProxyProviderManagerFrom constructs a [NEAppProxyProviderManager] from an unsafe.Pointer.
//
// An object to create and manage the app proxy provider’s VPN configuration.
func NEAppProxyProviderManagerFrom(ptr unsafe.Pointer) NEAppProxyProviderManager {
	return NEAppProxyProviderManager{
		NETunnelProviderManager: NETunnelProviderManagerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEAppProxyProviderManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEAppProxyProviderManager */

// Load all of the App Proxy configurations associated with the calling app that have previously been saved to the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProviderManager/loadAllFromPreferences(completionHandler:)
func (nc _NEAppProxyProviderManagerClass) LoadAllFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("loadAllFromPreferencesWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadAllFromPreferencesWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEAppProxyProviderManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEAppProxyProviderManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEAppProxyProviderManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEAppProxyProviderManager */



