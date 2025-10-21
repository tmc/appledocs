// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

// Enum types and constants
// NEAppProxyFlowError - Error codes that the app proxy flow API declares.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code
type NEAppProxyFlowError uint

// NEDNSSettingsManagerError - Error codes specific to DNS managers.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManagerError
type NEDNSSettingsManagerError uint

// NEHotspotConfigurationError - Error values returned by hotspot configuration manager methods.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError
type NEHotspotConfigurationError uint

const (
	// NEHotspotConfigurationErrorAlreadyAssociated - The configuration is already associated with the hotspot.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/alreadyAssociated
	NEHotspotConfigurationErrorAlreadyAssociated NEHotspotConfigurationError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/systemDenied
	NEHotspotConfigurationErrorSystemDenied NEHotspotConfigurationError = 0
)

// NENetworkRuleProtocol - A type to represent network protocols used by routing rules.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/Protocol
type NENetworkRuleProtocol uint

// NEProviderStopReason - Reasons why the provider extension was stopped.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason
type NEProviderStopReason uint

// NERelayManagerClientError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError
type NERelayManagerClientError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/certificateExpired
	NERelayManagerClientErrorCertificateExpired NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/certificateInvalid
	NERelayManagerClientErrorCertificateInvalid NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/certificateMissing
	NERelayManagerClientErrorCertificateMissing NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/dnsFailed
	NERelayManagerClientErrorDNSFailed NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/none
	NERelayManagerClientErrorNone NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/other
	NERelayManagerClientErrorOther NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/serverCertificateExpired
	NERelayManagerClientErrorServerCertificateExpired NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/serverCertificateInvalid
	NERelayManagerClientErrorServerCertificateInvalid NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/serverDisconnected
	NERelayManagerClientErrorServerDisconnected NERelayManagerClientError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError/serverUnreachable
	NERelayManagerClientErrorServerUnreachable NERelayManagerClientError = 0
)

// NERelayManagerError - Error codes specific to relay managers.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerError
type NERelayManagerError uint

// NETrafficDirection - A type to represent the direction of network traffic.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETrafficDirection
type NETrafficDirection uint

// NETunnelProviderError - Error codes that the tunnel provider declares.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderError-swift.struct/Code
type NETunnelProviderError uint

const (
	// NETunnelProviderErrorNetworkSettingsCanceled - The request to set or clear the tunnel network settings was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderError-swift.struct/Code/networkSettingsCanceled
	NETunnelProviderErrorNetworkSettingsCanceled NETunnelProviderError = 0
	// NETunnelProviderErrorNetworkSettingsFailed - The request to set or clear the tunnel network settings failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderError-swift.struct/Code/networkSettingsFailed
	NETunnelProviderErrorNetworkSettingsFailed NETunnelProviderError = 0
	// NETunnelProviderErrorNetworkSettingsInvalid - The provided tunnel network settings are invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderError-swift.struct/Code/networkSettingsInvalid
	NETunnelProviderErrorNetworkSettingsInvalid NETunnelProviderError = 0
)

// NETunnelProviderRoutingMethod enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderRoutingMethod
type NETunnelProviderRoutingMethod uint

// NEVPNConnectionError - Error codes specific to VPN connections.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError
type NEVPNConnectionError uint

// NEVPNError - Codes that indicate the source of an error.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNError-swift.struct/Code
type NEVPNError uint

const (
	// NEVPNErrorConfigurationDisabled - An error code indicating the VPN configuration associated with the VPN manager isn’t enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNError-swift.struct/Code/configurationDisabled
	NEVPNErrorConfigurationDisabled NEVPNError = 0
	// NEVPNErrorConfigurationInvalid - An error code indicating the VPN configuration associated with the VPN manager object is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNError-swift.struct/Code/configurationInvalid
	NEVPNErrorConfigurationInvalid NEVPNError = 0
	// NEVPNErrorConfigurationReadWriteFailed - An error code that indicates an error occurred while reading or writing the Network Extension preferences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNError-swift.struct/Code/configurationReadWriteFailed
	NEVPNErrorConfigurationReadWriteFailed NEVPNError = 0
	// NEVPNErrorConfigurationStale - An error code that indicates another process modfied the VPN configuration since the last time the app loaded the configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNError-swift.struct/Code/configurationStale
	NEVPNErrorConfigurationStale NEVPNError = 0
	// NEVPNErrorConfigurationUnknown - An error code that indicates that unspecified error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNError-swift.struct/Code/configurationUnknown
	NEVPNErrorConfigurationUnknown NEVPNError = 0
	// NEVPNErrorConnectionFailed - The connection to the VPN server failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNError-swift.struct/Code/connectionFailed
	NEVPNErrorConnectionFailed NEVPNError = 0
)

// NEVPNIKEAuthenticationMethod - Internet Key Exchange (IKE) authentication methods used to authenticate with the IPSec server.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEAuthenticationMethod
type NEVPNIKEAuthenticationMethod uint

// NEVPNIKEv2CertificateType - An enumeration of certificate type values.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType
type NEVPNIKEv2CertificateType uint

// NEVPNIKEv2DeadPeerDetectionRate - An enumeration of values for the frequency at which the IKEv2 client runs the dead peer detection algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DeadPeerDetectionRate
type NEVPNIKEv2DeadPeerDetectionRate uint

// NEVPNIKEv2TLSVersion - An enumeration of TLS Versions for use in EAP-TLS.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2TLSVersion
type NEVPNIKEv2TLSVersion uint

// NEVPNStatus - The possible states of a VPN connection.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus
type NEVPNStatus uint

const (
	// NEVPNStatusConnecting - The VPN is in the process of connecting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus/connecting
	NEVPNStatusConnecting NEVPNStatus = 0
	// NEVPNStatusDisconnected - The VPN is disconnected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus/disconnected
	NEVPNStatusDisconnected NEVPNStatus = 0
)


