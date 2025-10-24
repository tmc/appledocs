// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

// Enum types and constants
// NEDNSProxyManagerError - The possible DNS proxy manager errors.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError
type NEDNSProxyManagerError uint

// NEFilterPacketProviderVerdict - The verdict returned by a packet handler indicating what the framework should do with a packet.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/Verdict
type NEFilterPacketProviderVerdict uint

// NEHotspotConfigurationError - Error values returned by hotspot configuration manager methods.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError
type NEHotspotConfigurationError uint

// NEHotspotConfigurationEAPType - The EAP types that may be specified in 
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/EAPType
type NEHotspotConfigurationEAPType uint

// NEHotspotHelperCommandType - An enumeration of hotspot command types.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType
type NEHotspotHelperCommandType uint

// NERelayManagerClientError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError
type NERelayManagerClientError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/certificateExpired
	NERelayManagerClientErrorCertificateExpired NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/certificateMissing
	NERelayManagerClientErrorCertificateMissing NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/none
	NERelayManagerClientErrorNone NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/other
	NERelayManagerClientErrorOther NERelayManagerClientError = 0
)

// NETunnelProviderError - Error codes that the tunnel provider declares.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderError-swift.struct/Code
type NETunnelProviderError uint

const (
	// NETunnelProviderErrorNetworkSettingsCanceled - The request to set or clear the tunnel network settings was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderError-swift.struct/Code/networkSettingsCanceled
	NETunnelProviderErrorNetworkSettingsCanceled NETunnelProviderError = 0
)

// NEURLFilterVerdict - A verdict returned by a URL filter.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter/Verdict
type NEURLFilterVerdict uint

// NEVPNError - Codes that indicate the source of an error.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNError-swift.struct/Code
type NEVPNError uint

// NEVPNStatus - The possible states of a VPN connection.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus
type NEVPNStatus uint


