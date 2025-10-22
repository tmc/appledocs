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

// An interface definition for the [CWInterface] class.
type ICWInterface interface {
	objectivec.IObject
	ActivePHYMode() CWPHYMode
	AssociateToNetworkPasswordError(network ICWNetwork, password string, error_ unsafe.Pointer) bool
	AssociateToEnterpriseNetworkIdentityUsernamePasswordError(network ICWNetwork, identity unsafe.Pointer, username string, password string, error_ unsafe.Pointer) bool
	Bssid() foundation.String
	CachedScanResults() unsafe.Pointer
	CommitConfigurationAuthorizationError(configuration ICWConfiguration, authorization securityfoundation.ISFAuthorization, error_ unsafe.Pointer) bool
	Configuration() CWConfiguration
	CountryCode() foundation.String
	Disassociate()
	HardwareAddress() foundation.String
	InterfaceMode() CWInterfaceMode
	NoiseMeasurement() int
	PowerOn() bool
	RssiValue() int
	ScanForNetworksWithNameError(networkName string, error_ unsafe.Pointer) unsafe.Pointer
	ScanForNetworksWithNameIncludeHiddenError(networkName string, includeHidden bool, error_ unsafe.Pointer) unsafe.Pointer
	ScanForNetworksWithSSIDError(ssid foundation.IData, error_ unsafe.Pointer) unsafe.Pointer
	ScanForNetworksWithSSIDIncludeHiddenError(ssid foundation.IData, includeHidden bool, error_ unsafe.Pointer) unsafe.Pointer
	Security() CWSecurity
	ServiceActive() bool
	SetPairwiseMasterKeyError(key foundation.IData, error_ unsafe.Pointer) bool
	SetPowerError(power bool, error_ unsafe.Pointer) bool
	SetWEPKeyFlagsIndexError(key foundation.IData, flags CWCipherKeyFlags, index int, error_ unsafe.Pointer) bool
	SetWLANChannelError(channel ICWChannel, error_ unsafe.Pointer) bool
	Ssid() foundation.String
	SsidData() foundation.Data
	StartIBSSModeWithSSIDSecurityChannelPasswordError(ssidData foundation.IData, security ICWIBSSModeSecurity, channel uint, password string, error_ unsafe.Pointer) bool
	SupportedWLANChannels() unsafe.Pointer
	TransmitPower() int
	TransmitRate() float64
	WlanChannel() CWChannel
	InterfaceName() string
}

// Encapsulates an IEEE 802.11 interface.
//
// Provides access to various WLAN interface parameters, and operations such as scanning for networks, association, and creating computer-to-computer (ad-hoc) networks.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CWInterfaceClass) Alloc() CWInterface {
	rv := objc.Send[CWInterface](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Convenience method for getting an CWInterface object with the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/init(interfaceName:)
func NewCWInterfaceWithInterfaceName(name string) CWInterface {
	instance := getCWInterfaceClass().Alloc()
	rv := objc.Send[CWInterface](instance.ID, objc.Sel("initWithInterfaceName:"), objc.String(name))
	rv.Autorelease()
	return rv
}



// An instance method for obtaining an CWInterface object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/init(name:)
func NewCWInterfaceWithName(name string) CWInterface {
	rv := objc.Send[CWInterface](objc.ID(getCWInterfaceClass().class), objc.Sel("interfaceWithName:"), objc.String(name))
	return rv
}


// An instance method for obtaining an CWInterface object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/init(name:)
func (cc _CWInterfaceClass) InterfaceWithName(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("interfaceWithName:"), objc.String(name))
	return rv
}

// Convenience method for getting an CWInterface object for the default WLAN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interface
func (cc _CWInterfaceClass) Interface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("interface"))
	return rv
}

// Returns the list of BSD names for WLAN interfaces available on the current system.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interfaceNames()
func (cc _CWInterfaceClass) InterfaceNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("interfaceNames"))
	return rv
}

// The current active PHY modes for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/activePHYMode()
func (c_ CWInterface) ActivePHYMode() CWPHYMode {
	rv := objc.Send[CWPHYMode](c_.ID, objc.Sel("activePHYMode"))
	return rv
}

// Associates to a given network using the given network passphrase.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/associate(to:password:)
func (c_ CWInterface) AssociateToNetworkPasswordError(network ICWNetwork, password string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("associateToNetwork:password:error:"), network, objc.String(password), error_)
	return rv
}

// Connects to the given enterprise network.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/associate(toEnterpriseNetwork:identity:username:password:)
func (c_ CWInterface) AssociateToEnterpriseNetworkIdentityUsernamePasswordError(network ICWNetwork, identity unsafe.Pointer, username string, password string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("associateToEnterpriseNetwork:identity:username:password:error:"), network, identity, objc.String(username), objc.String(password), error_)
	return rv
}

// The current basic service set identifier (BSSID) for the interface, returned as a UTF-8 string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/bssid()
func (c_ CWInterface) Bssid() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("bssid"))
	return rv
}

// The networks currently in the scan cache for the WLAN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/cachedScanResults()
func (c_ CWInterface) CachedScanResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cachedScanResults"))
	return rv
}

// Commit a configuration for the given WLAN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/commitConfiguration(_:authorization:)
func (c_ CWInterface) CommitConfigurationAuthorizationError(configuration ICWConfiguration, authorization securityfoundation.ISFAuthorization, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("commitConfiguration:authorization:error:"), configuration, authorization, error_)
	return rv
}

// The current configuration for the given WLAN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/configuration()
func (c_ CWInterface) Configuration() CWConfiguration {
	rv := objc.Send[CWConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}

// The current country code (ISO/IEC 3166-1:1997) for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/countryCode()
func (c_ CWInterface) CountryCode() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("countryCode"))
	return rv
}

// Disassociates from the current network.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/disassociate()
func (c_ CWInterface) Disassociate() {
	objc.Send[objc.ID](c_.ID, objc.Sel("disassociate"))
}

// The hardware media access control (MAC) address for the interface, returned as a UTF-8 string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/hardwareAddress()
func (c_ CWInterface) HardwareAddress() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("hardwareAddress"))
	return rv
}

// The current mode for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interfaceMode()
func (c_ CWInterface) InterfaceMode() CWInterfaceMode {
	rv := objc.Send[CWInterfaceMode](c_.ID, objc.Sel("interfaceMode"))
	return rv
}

// The current aggregate noise measurement (dBm) for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/noiseMeasurement()
func (c_ CWInterface) NoiseMeasurement() int {
	rv := objc.Send[int](c_.ID, objc.Sel("noiseMeasurement"))
	return rv
}

// The interface power state is set to “ON”.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/powerOn()
func (c_ CWInterface) PowerOn() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("powerOn"))
	return rv
}

// The current aggregate received signal strength indication (RSSI) measurement (dBm) for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/rssiValue()
func (c_ CWInterface) RssiValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("rssiValue"))
	return rv
}

// Scans for networks.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withName:)
func (c_ CWInterface) ScanForNetworksWithNameError(networkName string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scanForNetworksWithName:error:"), objc.String(networkName), error_)
	return rv
}

// Scans for networks with the name you specify, optionally including hidden networks.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withName:includeHidden:)
func (c_ CWInterface) ScanForNetworksWithNameIncludeHiddenError(networkName string, includeHidden bool, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scanForNetworksWithName:includeHidden:error:"), objc.String(networkName), includeHidden, error_)
	return rv
}

// Scans for networks.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withSSID:)
func (c_ CWInterface) ScanForNetworksWithSSIDError(ssid foundation.IData, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scanForNetworksWithSSID:error:"), ssid, error_)
	return rv
}

// Scans for networks with the SSID you specify, optionally including hidden networks.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withSSID:includeHidden:)
func (c_ CWInterface) ScanForNetworksWithSSIDIncludeHiddenError(ssid foundation.IData, includeHidden bool, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scanForNetworksWithSSID:includeHidden:error:"), ssid, includeHidden, error_)
	return rv
}

// The current security mode for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/security()
func (c_ CWInterface) Security() CWSecurity {
	rv := objc.Send[CWSecurity](c_.ID, objc.Sel("security"))
	return rv
}

// The interface has its corresponding network service enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/serviceActive()
func (c_ CWInterface) ServiceActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("serviceActive"))
	return rv
}

// Sets the interface pairwise primary key (PMK).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setPairwiseMasterKey(_:)
func (c_ CWInterface) SetPairwiseMasterKeyError(key foundation.IData, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setPairwiseMasterKey:error:"), key, error_)
	return rv
}

// Sets the interface power state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setPower(_:)
func (c_ CWInterface) SetPowerError(power bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setPower:error:"), power, error_)
	return rv
}

// Sets the interface WEP key.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setWEPKey(_:flags:index:)
func (c_ CWInterface) SetWEPKeyFlagsIndexError(key foundation.IData, flags CWCipherKeyFlags, index int, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setWEPKey:flags:index:error:"), key, flags, index, error_)
	return rv
}

// Sets the interface channel.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setWLANChannel(_:)
func (c_ CWInterface) SetWLANChannelError(channel ICWChannel, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setWLANChannel:error:"), channel, error_)
	return rv
}

// The current service set identifier (SSID) for the interface, encoded as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/ssid()
func (c_ CWInterface) Ssid() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("ssid"))
	return rv
}

// The current service set identifier (SSID) for the interface, returned as data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/ssidData()
func (c_ CWInterface) SsidData() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("ssidData"))
	return rv
}

// Creates a computer-to-computer (ad-hoc) network with the given network name, security type, and password on the specified channel.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/startIBSSMode(withSSID:security:channel:password:)
func (c_ CWInterface) StartIBSSModeWithSSIDSecurityChannelPasswordError(ssidData foundation.IData, security ICWIBSSModeSecurity, channel uint, password string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("startIBSSModeWithSSID:security:channel:password:error:"), ssidData, security, channel, objc.String(password), error_)
	return rv
}

// An array of channels supported by the interface for the active country code.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/supportedWLANChannels()
func (c_ CWInterface) SupportedWLANChannels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supportedWLANChannels"))
	return rv
}

// The current transmit power (mW) for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/transmitPower()
func (c_ CWInterface) TransmitPower() int {
	rv := objc.Send[int](c_.ID, objc.Sel("transmitPower"))
	return rv
}

// The current transmit rate (Mbps) for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/transmitRate()
func (c_ CWInterface) TransmitRate() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("transmitRate"))
	return rv
}

// The current channel for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/wlanChannel()
func (c_ CWInterface) WlanChannel() CWChannel {
	rv := objc.Send[CWChannel](c_.ID, objc.Sel("wlanChannel"))
	return rv
}

// The BSD name of the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interfaceName
func (c_ CWInterface) InterfaceName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("interfaceName"))
	return rv
}


