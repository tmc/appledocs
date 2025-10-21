// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [NEAppPushProvider] class.
type INEAppPushProvider interface {
	INEProvider
	ReportIncomingCallWithUserInfo(userInfo objc.ID)
	StartWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An object that creates and maintains a persistent network connection to a local push server.
//
// Subclass to provide the connection to your local push server. A creates instances of your provider class based on the in the manager’s configuration. The manager then calls methods on your provider to start and stop communication with the server, and periodically check the provider’s status. When your provider receives an incoming call from your server, call the provider’s method to alert the manager’s .
//
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

// Alloc allocates a new instance without initialization.
func (nc _NEAppPushProviderClass) Alloc() NEAppPushProvider {
	rv := objc.Send[NEAppPushProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Informs the manager about an incoming call.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/reportIncomingCall(userInfo:)
func (n_ NEAppPushProvider) ReportIncomingCallWithUserInfo(userInfo objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("reportIncomingCallWithUserInfo:"), userInfo)
}

// Indicates that the framework has started the provider, and provides a completion handler for subclasses to signal their readiness.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/start(completionHandler:)
func (n_ NEAppPushProvider) StartWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("startWithCompletionHandler:"), completionHandler)
}

// A dictionary that contains current vendor-specific configuration parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushprovider/providerconfiguration
func (n_ NEAppPushProvider) ProviderConfiguration() string {
	rv := objc.Send[string](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}


// SetProviderConfiguration sets the value of the providerConfiguration property.
// A dictionary that contains current vendor-specific configuration parameters.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushprovider/providerconfiguration
func (n_ NEAppPushProvider) SetProviderConfiguration(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), objc.String(value))
}

// A delegate that receives incoming call information from the provider.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/delegate
func (n_ NEAppPushProvider) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate that receives incoming call information from the provider.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/delegate
func (n_ NEAppPushProvider) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}

// A string that contains the bundle identifier of the push provider.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/providerbundleidentifier
func (n_ NEAppPushProvider) ProviderBundleIdentifier() string {
	rv := objc.Send[string](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}


// SetProviderBundleIdentifier sets the value of the providerBundleIdentifier property.
// A string that contains the bundle identifier of the push provider.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/providerbundleidentifier
func (n_ NEAppPushProvider) SetProviderBundleIdentifier(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), objc.String(value))
}



