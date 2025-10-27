// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NEVPNManager] class.
var (
	NEVPNManagerClass     _NEVPNManagerClass
	NEVPNManagerClassOnce sync.Once
)

func getNEVPNManagerClass() _NEVPNManagerClass {
	NEVPNManagerClassOnce.Do(func() {
		NEVPNManagerClass = _NEVPNManagerClass{objc.GetClass("NEVPNManager")}
	})
	return NEVPNManagerClass
}

type _NEVPNManagerClass struct {
	class objc.Class
}





// An interface definition for the [NEVPNManager] class.
type INEVPNManager interface {
	objectivec.IObject
	

	// properties:
	Connection() INEVPNConnection
	Enabled() bool
	SetEnabled(value bool)
	OnDemandEnabled() bool
	SetOnDemandEnabled(value bool)
	LocalizedDescription() foundation.foundation.INSString
	SetLocalizedDescription(value foundation.foundation.INSString)
	OnDemandRules() []NEOnDemandRule
	SetOnDemandRules(value []NEOnDemandRule)
	Protocol() INEVPNProtocol
	SetProtocol(value INEVPNProtocol)
	ProtocolConfiguration() INEVPNProtocol
	SetProtocolConfiguration(value INEVPNProtocol)
	NEVPNErrorDomain() foundation.foundation.INSString
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsOnDemandEnabled() bool
	SetIsOnDemandEnabled(value bool)


	

	// methods:
	LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	SetAuthorization(authorization objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (nc _NEVPNManagerClass) Alloc() NEVPNManager {
	rv := objc.Send[NEVPNManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEVPNManagerClass) New() NEVPNManager {
	rv := objc.Send[NEVPNManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNManager) Init() NEVPNManager {
	rv := objc.Send[NEVPNManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNManager) Autorelease() NEVPNManager {
	rv := objc.Send[NEVPNManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNManager creates a new NEVPNManager instance.
func NewNEVPNManager() NEVPNManager {
	return getNEVPNManagerClass().New()
}





// An object to create and manage a Personal VPN configuration.
//
// The API gives apps the ability to create and manage a Personal VPN configuration on iOS and macOS. Personal VPN configurations are typically used to provide a service to users that protects their Internet browsing activity on insecure networks such as public Wi-Fi networks.


// An object to create and manage a Personal VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager
type NEVPNManager struct {
	objectivec.Object
}

// NEVPNManagerFrom constructs a [NEVPNManager] from an unsafe.Pointer.
//
// An object to create and manage a Personal VPN configuration.
func NEVPNManagerFrom(ptr unsafe.Pointer) NEVPNManager {
	return NEVPNManager{objectivec.Object{objc.ID(ptr)}}
}










// Access the single instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/shared()
func (nc _NEVPNManagerClass) SharedManager() NEVPNManager {
	rv := objc.Send[NEVPNManager](objc.ID(nc.class), objc.Sel("sharedManager"))
	return rv
}












// Load the VPN configuration from the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/loadFromPreferences(completionHandler:)
func (n_ NEVPNManager) LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}


// Remove the VPN configuration from the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/removeFromPreferences(completionHandler:)
func (n_ NEVPNManager) RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), completionHandler)
}


// Save the VPN configuration in the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/saveToPreferences(completionHandler:)
func (n_ NEVPNManager) SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/setAuthorization(_:)
func (n_ NEVPNManager) SetAuthorization(authorization objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAuthorization:"), authorization)
}







// An object that is used to control the VPN tunnel specified by the VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/connection
func (n_ NEVPNManager) Connection() INEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("connection"))
	return rv
}


// A Boolean used to toggle the enabled state of the VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/isEnabled
func (n_ NEVPNManager) Enabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean used to toggle the enabled state of the VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/isEnabled
func (n_ NEVPNManager) SetEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnabled:"), value)
}


// A Boolean used to toggle the Connect On Demand capability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/isOnDemandEnabled
func (n_ NEVPNManager) OnDemandEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("onDemandEnabled"))
	return rv
}


// A Boolean used to toggle the Connect On Demand capability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/isOnDemandEnabled
func (n_ NEVPNManager) SetOnDemandEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOnDemandEnabled:"), value)
}


// A string containing the display name of the VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/localizedDescription
func (n_ NEVPNManager) LocalizedDescription() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A string containing the display name of the VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/localizedDescription
func (n_ NEVPNManager) SetLocalizedDescription(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}


// An ordered list of Connect On Demand rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/onDemandRules
func (n_ NEVPNManager) OnDemandRules() []NEOnDemandRule {
	rv := objc.Send[[]NEOnDemandRule](n_.ID, objc.Sel("onDemandRules"))
	return rv
}


// An ordered list of Connect On Demand rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/onDemandRules
func (n_ NEVPNManager) SetOnDemandRules(value []NEOnDemandRule) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setOnDemandRules:"), nsArray)
}


// An object containing the configuration settings of the VPN tunneling protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/protocol
func (n_ NEVPNManager) Protocol() INEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("protocol"))
	return rv
}


// An object containing the configuration settings of the VPN tunneling protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/protocol
func (n_ NEVPNManager) SetProtocol(value INEVPNProtocol) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProtocol:"), value)
}


// An object containing the configuration settings of the VPN tunneling protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/protocolConfiguration
func (n_ NEVPNManager) ProtocolConfiguration() INEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("protocolConfiguration"))
	return rv
}


// An object containing the configuration settings of the VPN tunneling protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/protocolConfiguration
func (n_ NEVPNManager) SetProtocolConfiguration(value INEVPNProtocol) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProtocolConfiguration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnerrordomain
func (n_ NEVPNManager) NEVPNErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEVPNErrorDomain"))
	return rv
}


// A Boolean used to toggle the enabled state of the VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/isenabled
func (n_ NEVPNManager) IsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean used to toggle the enabled state of the VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/isenabled
func (n_ NEVPNManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean used to toggle the Connect On Demand capability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/isondemandenabled
func (n_ NEVPNManager) IsOnDemandEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isOnDemandEnabled"))
	return rv
}


// A Boolean used to toggle the Connect On Demand capability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/isondemandenabled
func (n_ NEVPNManager) SetIsOnDemandEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsOnDemandEnabled:"), value)
}








