// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/securityfoundation"
)

/* debug [class.gen.go]: Generating class CWInterface */


/* debug [class_header]: Header for CWInterface */
// The class instance for the [CWInterface] class.
var (
	CWInterfaceClass     _CWInterfaceClass
	CWInterfaceClassOnce sync.Once
)

func getCWInterfaceClass() _CWInterfaceClass {
	CWInterfaceClassOnce.Do(func() {
		CWInterfaceClass = _CWInterfaceClass{objc.GetClass("CWInterface")}
	})
	return CWInterfaceClass
}

type _CWInterfaceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CWInterface */
// An interface definition for the [CWInterface] class.
type ICWInterface interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CWInterface */
	// properties:
	InterfaceName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CWInterface */
	// methods:
	ActivePHYMode() CWPHYMode
	AssociateToNetworkPasswordError(network ICWNetwork, password objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool
	AssociateToEnterpriseNetworkIdentityUsernamePasswordError(network ICWNetwork, identity unsafe.Pointer, username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool
	Bssid() foundation.String
	CachedScanResults() unsafe.Pointer
	CommitConfigurationAuthorizationError(configuration ICWConfiguration, authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool
	Configuration() ICWConfiguration
	CountryCode() foundation.String
	Disassociate()
	HardwareAddress() foundation.String
	InterfaceMode() CWInterfaceMode
	NoiseMeasurement() int
	PowerOn() bool
	RssiValue() int
	ScanForNetworksWithNameError(networkName objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) unsafe.Pointer
	ScanForNetworksWithNameIncludeHiddenError(networkName objc.IObject /* cross-framework: NSString */, includeHidden bool, error_ unsafe.Pointer) unsafe.Pointer
	ScanForNetworksWithSSIDError(ssid objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) unsafe.Pointer
	ScanForNetworksWithSSIDIncludeHiddenError(ssid objc.IObject /* cross-framework: NSData */, includeHidden bool, error_ unsafe.Pointer) unsafe.Pointer
	Security() CWSecurity
	ServiceActive() bool
	SetPairwiseMasterKeyError(key objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) bool
	SetPowerError(power bool, error_ unsafe.Pointer) bool
	SetWEPKeyFlagsIndexError(key objc.IObject /* cross-framework: NSData */, flags CWCipherKeyFlags, index int, error_ unsafe.Pointer) bool
	SetWLANChannelError(channel ICWChannel, error_ unsafe.Pointer) bool
	Ssid() foundation.String
	SsidData() foundation.Data
	SupportedWLANChannels() unsafe.Pointer
	TransmitPower() int
	TransmitRate() float64
	WlanChannel() ICWChannel
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CWInterface */
// Alloc allocates a new instance without initialization.
func (cc _CWInterfaceClass) Alloc() CWInterface {
	rv := objc.Send[CWInterface](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CWInterfaceClass) New() CWInterface {
	rv := objc.Send[CWInterface](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWInterface) Init() CWInterface {
	rv := objc.Send[CWInterface](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWInterface) Autorelease() CWInterface {
	rv := objc.Send[CWInterface](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWInterface creates a new CWInterface instance.
func NewCWInterface() CWInterface {
	return getCWInterfaceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CWInterface */
// Encapsulates an IEEE 802.11 interface.
//
// Provides access to various WLAN interface parameters, and operations such as scanning for networks, association, and creating computer-to-computer (ad-hoc) networks.


// Encapsulates an IEEE 802.11 interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface
type CWInterface struct {
	objectivec.Object
}

// CWInterfaceFrom constructs a [CWInterface] from an unsafe.Pointer.
//
// Encapsulates an IEEE 802.11 interface.
func CWInterfaceFrom(ptr unsafe.Pointer) CWInterface {
	return CWInterface{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CWInterface */

// Convenience method for getting an CWInterface object with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/init(interfaceName:)
func NewCWInterfaceWithInterfaceName(name objc.IObject /* cross-framework: NSString */) CWInterface {
	instance := getCWInterfaceClass().Alloc()
	rv := objc.Send[CWInterface](instance.ID, objc.Sel("initWithInterfaceName:"), name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCWInterfaceWithInterfaceName */


// An instance method for obtaining an CWInterface object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/init(name:)
func NewCWInterfaceWithName(name objc.IObject /* cross-framework: NSString */) CWInterface {
	rv := objc.Send[CWInterface](objc.ID(getCWInterfaceClass().class), objc.Sel("interfaceWithName:"), name)
	return rv
}/* debug [class_init_methods/constructor]: NewCWInterfaceWithName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CWInterface */

// An instance method for obtaining an CWInterface object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/init(name:)
func (cc _CWInterfaceClass) InterfaceWithName(name objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("interfaceWithName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterfaceWithName) */


// Convenience method for getting an CWInterface object for the default WLAN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interface
func (cc _CWInterfaceClass) Interface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("interface"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Interface) */


// Returns the list of BSD names for WLAN interfaces available on the current system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interfaceNames()
func (cc _CWInterfaceClass) InterfaceNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("interfaceNames"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterfaceNames) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CWInterface */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CWInterface */

// The current active PHY modes for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/activePHYMode()
func (c_ CWInterface) ActivePHYMode() CWPHYMode {
	rv := objc.Send[CWPHYMode](c_.ID, objc.Sel("activePHYMode"))
	return rv
}/* debug [instance_methods/method]: ActivePHYMode */


// Associates to a given network using the given network passphrase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/associate(to:password:)
func (c_ CWInterface) AssociateToNetworkPasswordError(network ICWNetwork, password objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("associateToNetwork:password:error:"), network, password, error_)
	return rv
}/* debug [instance_methods/method]: AssociateToNetworkPasswordError */


// Connects to the given enterprise network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/associate(toEnterpriseNetwork:identity:username:password:)
func (c_ CWInterface) AssociateToEnterpriseNetworkIdentityUsernamePasswordError(network ICWNetwork, identity unsafe.Pointer, username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("associateToEnterpriseNetwork:identity:username:password:error:"), network, identity, username, password, error_)
	return rv
}/* debug [instance_methods/method]: AssociateToEnterpriseNetworkIdentityUsernamePasswordError */


// The current basic service set identifier (BSSID) for the interface, returned as a UTF-8 string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/bssid()
func (c_ CWInterface) Bssid() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("bssid"))
	return rv
}/* debug [instance_methods/method]: Bssid */


// The networks currently in the scan cache for the WLAN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/cachedScanResults()
func (c_ CWInterface) CachedScanResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cachedScanResults"))
	return rv
}/* debug [instance_methods/method]: CachedScanResults */


// Commit a configuration for the given WLAN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/commitConfiguration(_:authorization:)
func (c_ CWInterface) CommitConfigurationAuthorizationError(configuration ICWConfiguration, authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("commitConfiguration:authorization:error:"), configuration, authorization, error_)
	return rv
}/* debug [instance_methods/method]: CommitConfigurationAuthorizationError */


// The current configuration for the given WLAN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/configuration()
func (c_ CWInterface) Configuration() ICWConfiguration {
	rv := objc.Send[CWConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_methods/method]: Configuration */


// The current country code (ISO/IEC 3166-1:1997) for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/countryCode()
func (c_ CWInterface) CountryCode() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("countryCode"))
	return rv
}/* debug [instance_methods/method]: CountryCode */


// Disassociates from the current network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/disassociate()
func (c_ CWInterface) Disassociate() {
	objc.Send[objc.ID](c_.ID, objc.Sel("disassociate"))
}/* debug [instance_methods/method]: Disassociate */


// The hardware media access control (MAC) address for the interface, returned as a UTF-8 string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/hardwareAddress()
func (c_ CWInterface) HardwareAddress() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("hardwareAddress"))
	return rv
}/* debug [instance_methods/method]: HardwareAddress */


// The current mode for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interfaceMode()
func (c_ CWInterface) InterfaceMode() CWInterfaceMode {
	rv := objc.Send[CWInterfaceMode](c_.ID, objc.Sel("interfaceMode"))
	return rv
}/* debug [instance_methods/method]: InterfaceMode */


// The current aggregate noise measurement (dBm) for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/noiseMeasurement()
func (c_ CWInterface) NoiseMeasurement() int {
	rv := objc.Send[int](c_.ID, objc.Sel("noiseMeasurement"))
	return rv
}/* debug [instance_methods/method]: NoiseMeasurement */


// The interface power state is set to “ON”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/powerOn()
func (c_ CWInterface) PowerOn() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("powerOn"))
	return rv
}/* debug [instance_methods/method]: PowerOn */


// The current aggregate received signal strength indication (RSSI) measurement (dBm) for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/rssiValue()
func (c_ CWInterface) RssiValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("rssiValue"))
	return rv
}/* debug [instance_methods/method]: RssiValue */


// Scans for networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withName:)
func (c_ CWInterface) ScanForNetworksWithNameError(networkName objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scanForNetworksWithName:error:"), networkName, error_)
	return rv
}/* debug [instance_methods/method]: ScanForNetworksWithNameError */


// Scans for networks with the name you specify, optionally including hidden networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withName:includeHidden:)
func (c_ CWInterface) ScanForNetworksWithNameIncludeHiddenError(networkName objc.IObject /* cross-framework: NSString */, includeHidden bool, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scanForNetworksWithName:includeHidden:error:"), networkName, includeHidden, error_)
	return rv
}/* debug [instance_methods/method]: ScanForNetworksWithNameIncludeHiddenError */


// Scans for networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withSSID:)
func (c_ CWInterface) ScanForNetworksWithSSIDError(ssid objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scanForNetworksWithSSID:error:"), ssid, error_)
	return rv
}/* debug [instance_methods/method]: ScanForNetworksWithSSIDError */


// Scans for networks with the SSID you specify, optionally including hidden networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withSSID:includeHidden:)
func (c_ CWInterface) ScanForNetworksWithSSIDIncludeHiddenError(ssid objc.IObject /* cross-framework: NSData */, includeHidden bool, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scanForNetworksWithSSID:includeHidden:error:"), ssid, includeHidden, error_)
	return rv
}/* debug [instance_methods/method]: ScanForNetworksWithSSIDIncludeHiddenError */


// The current security mode for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/security()
func (c_ CWInterface) Security() CWSecurity {
	rv := objc.Send[CWSecurity](c_.ID, objc.Sel("security"))
	return rv
}/* debug [instance_methods/method]: Security */


// The interface has its corresponding network service enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/serviceActive()
func (c_ CWInterface) ServiceActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("serviceActive"))
	return rv
}/* debug [instance_methods/method]: ServiceActive */


// Sets the interface pairwise primary key (PMK).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setPairwiseMasterKey(_:)
func (c_ CWInterface) SetPairwiseMasterKeyError(key objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setPairwiseMasterKey:error:"), key, error_)
	return rv
}/* debug [instance_methods/method]: SetPairwiseMasterKeyError */


// Sets the interface power state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setPower(_:)
func (c_ CWInterface) SetPowerError(power bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setPower:error:"), power, error_)
	return rv
}/* debug [instance_methods/method]: SetPowerError */


// Sets the interface WEP key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setWEPKey(_:flags:index:)
func (c_ CWInterface) SetWEPKeyFlagsIndexError(key objc.IObject /* cross-framework: NSData */, flags CWCipherKeyFlags, index int, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setWEPKey:flags:index:error:"), key, flags, index, error_)
	return rv
}/* debug [instance_methods/method]: SetWEPKeyFlagsIndexError */


// Sets the interface channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setWLANChannel(_:)
func (c_ CWInterface) SetWLANChannelError(channel ICWChannel, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setWLANChannel:error:"), channel, error_)
	return rv
}/* debug [instance_methods/method]: SetWLANChannelError */


// The current service set identifier (SSID) for the interface, encoded as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/ssid()
func (c_ CWInterface) Ssid() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("ssid"))
	return rv
}/* debug [instance_methods/method]: Ssid */


// The current service set identifier (SSID) for the interface, returned as data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/ssidData()
func (c_ CWInterface) SsidData() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("ssidData"))
	return rv
}/* debug [instance_methods/method]: SsidData */


// An array of channels supported by the interface for the active country code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/supportedWLANChannels()
func (c_ CWInterface) SupportedWLANChannels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supportedWLANChannels"))
	return rv
}/* debug [instance_methods/method]: SupportedWLANChannels */


// The current transmit power (mW) for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/transmitPower()
func (c_ CWInterface) TransmitPower() int {
	rv := objc.Send[int](c_.ID, objc.Sel("transmitPower"))
	return rv
}/* debug [instance_methods/method]: TransmitPower */


// The current transmit rate (Mbps) for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/transmitRate()
func (c_ CWInterface) TransmitRate() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("transmitRate"))
	return rv
}/* debug [instance_methods/method]: TransmitRate */


// The current channel for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/wlanChannel()
func (c_ CWInterface) WlanChannel() ICWChannel {
	rv := objc.Send[CWChannel](c_.ID, objc.Sel("wlanChannel"))
	return rv
}/* debug [instance_methods/method]: WlanChannel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CWInterface */

// The BSD name of the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interfaceName
func (c_ CWInterface) InterfaceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("interfaceName"))
	return rv
}/* debug [instance_properties/getter]: interfaceName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CWInterface */


