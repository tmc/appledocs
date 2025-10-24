// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

/* debug [enums.gen.go]: Generating 45 enums for NetworkExtension */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum NERelayManagerClientError (10 cases) */
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

/* debug [enums.gen.go]: Processing enum NETunnelProviderError (3 cases) */
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

/* debug [enums.gen.go]: Processing enum NEVPNError (6 cases) */
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

/* debug [enums.gen.go]: Processing enum NEAppProxyFlowError (10 cases) */
// NEAppProxyFlowError - Error codes that the app proxy flow API declares.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code
type NEAppProxyFlowError uint

const (
	// NEAppProxyFlowErrorAborted - The flow was aborted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/aborted
	NEAppProxyFlowErrorAborted NEAppProxyFlowError = 0
	// NEAppProxyFlowErrorDatagramTooLarge - A caller attempted to write a datagram that was larger than the socket’s receive window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/datagramTooLarge
	NEAppProxyFlowErrorDatagramTooLarge NEAppProxyFlowError = 0
	// NEAppProxyFlowErrorHostUnreachable - An attempt to reach the remote endpoint of the flow failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/hostUnreachable
	NEAppProxyFlowErrorHostUnreachable NEAppProxyFlowError = 0
	// NEAppProxyFlowErrorInternal - An internal error occurred while handling the flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/internal
	NEAppProxyFlowErrorInternal NEAppProxyFlowError = 0
	// NEAppProxyFlowErrorInvalidArgument - A proxy flow method received an invalid argument.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/invalidArgument
	NEAppProxyFlowErrorInvalidArgument NEAppProxyFlowError = 0
	// NEAppProxyFlowErrorNotConnected - The flow is not fully opened.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/notConnected
	NEAppProxyFlowErrorNotConnected NEAppProxyFlowError = 0
	// NEAppProxyFlowErrorPeerReset - The remote peer closed the flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/peerReset
	NEAppProxyFlowErrorPeerReset NEAppProxyFlowError = 0
	// NEAppProxyFlowErrorReadAlreadyPending - A read operation on the flow is already pending.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/readAlreadyPending
	NEAppProxyFlowErrorReadAlreadyPending NEAppProxyFlowError = 0
	// NEAppProxyFlowErrorRefused - Connecting the flow to its remote endpoint failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/refused
	NEAppProxyFlowErrorRefused NEAppProxyFlowError = 0
	// NEAppProxyFlowErrorTimedOut - The flow timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code/timedOut
	NEAppProxyFlowErrorTimedOut NEAppProxyFlowError = 0
)

/* debug [enums.gen.go]: Processing enum NEAppPushManagerError (4 cases) */
// NEAppPushManagerError - Error codes that the local push API declares.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManagerError-swift.struct/Code
type NEAppPushManagerError uint

const (
	// NEAppPushManagerErrorConfigurationInvalid - An error code that indicates the app push configuration is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManagerError-swift.struct/Code/configurationInvalid
	NEAppPushManagerErrorConfigurationInvalid NEAppPushManagerError = 0
	// NEAppPushManagerErrorConfigurationNotLoaded - An error code that indicates the manager hasn’t loaded the app push configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManagerError-swift.struct/Code/configurationNotLoaded
	NEAppPushManagerErrorConfigurationNotLoaded NEAppPushManagerError = 0
	// NEAppPushManagerErrorInactiveSession - An error code that indicates an invalid attempt to perform an operation on an inactive session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManagerError-swift.struct/Code/inactiveSession
	NEAppPushManagerErrorInactiveSession NEAppPushManagerError = 0
	// NEAppPushManagerErrorInternalError - An error code that indicates an internal error in the local push connectivity framework.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManagerError-swift.struct/Code/internalError
	NEAppPushManagerErrorInternalError NEAppPushManagerError = 0
)

/* debug [enums.gen.go]: Processing enum NEDNSProtocol (3 cases) */
// NEDNSProtocol enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProtocol
type NEDNSProtocol uint

const (
	// NEDNSProtocolCleartext - The DNS server uses cleartext UDP or TCP over port 53.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProtocol/cleartext
	NEDNSProtocolCleartext NEDNSProtocol = 0
	// NEDNSProtocolHTTPS - The DNS server uses DNS-over-HTTPS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProtocol/HTTPS
	NEDNSProtocolHTTPS NEDNSProtocol = 0
	// NEDNSProtocolTLS - The DNS server uses DNS-over-TLS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProtocol/TLS
	NEDNSProtocolTLS NEDNSProtocol = 0
)

/* debug [enums.gen.go]: Processing enum NEDNSProxyManagerError (4 cases) */
// NEDNSProxyManagerError - The possible DNS proxy manager errors.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError
type NEDNSProxyManagerError uint

const (
	// NEDNSProxyManagerErrorConfigurationCannotBeRemoved - Unremovable DNS proxy configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError/configurationCannotBeRemoved
	NEDNSProxyManagerErrorConfigurationCannotBeRemoved NEDNSProxyManagerError = 0
	// NEDNSProxyManagerErrorConfigurationDisabled - Disabled DNS proxy configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError/configurationDisabled
	NEDNSProxyManagerErrorConfigurationDisabled NEDNSProxyManagerError = 0
	// NEDNSProxyManagerErrorConfigurationInvalid - Invalid DNS proxy configuration that cannot be stored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError/configurationInvalid
	NEDNSProxyManagerErrorConfigurationInvalid NEDNSProxyManagerError = 0
	// NEDNSProxyManagerErrorConfigurationStale - Outdated DNS proxy configuration that needs to be loaded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError/configurationStale
	NEDNSProxyManagerErrorConfigurationStale NEDNSProxyManagerError = 0
)

/* debug [enums.gen.go]: Processing enum NEDNSSettingsManagerError (4 cases) */
// NEDNSSettingsManagerError - Error codes specific to DNS managers.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManagerError
type NEDNSSettingsManagerError uint

const (
	// NEDNSSettingsManagerErrorConfigurationCannotBeRemoved - An error code that indicates removing the DNS settings manager failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManagerError/configurationCannotBeRemoved
	NEDNSSettingsManagerErrorConfigurationCannotBeRemoved NEDNSSettingsManagerError = 0
	// NEDNSSettingsManagerErrorConfigurationDisabled - An error code that indicates the DNS settings manager isn’t enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManagerError/configurationDisabled
	NEDNSSettingsManagerErrorConfigurationDisabled NEDNSSettingsManagerError = 0
	// NEDNSSettingsManagerErrorConfigurationInvalid - An error code that indicates the DNS settings manager is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManagerError/configurationInvalid
	NEDNSSettingsManagerErrorConfigurationInvalid NEDNSSettingsManagerError = 0
	// NEDNSSettingsManagerErrorConfigurationStale - An error code that indicates the DNS settings manager isn’t loaded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManagerError/configurationStale
	NEDNSSettingsManagerErrorConfigurationStale NEDNSSettingsManagerError = 0
)

/* debug [enums.gen.go]: Processing enum NEEvaluateConnectionRuleAction (2 cases) */
// NEEvaluateConnectionRuleAction enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRuleAction
type NEEvaluateConnectionRuleAction uint

const (
	// NEEvaluateConnectionRuleActionConnectIfNeeded - Start the VPN if connections to the matching hostname cannot be resolved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRuleAction/connectIfNeeded
	NEEvaluateConnectionRuleActionConnectIfNeeded NEEvaluateConnectionRuleAction = 0
	// NEEvaluateConnectionRuleActionNeverConnect - Do not start the VPN.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRuleAction/neverConnect
	NEEvaluateConnectionRuleActionNeverConnect NEEvaluateConnectionRuleAction = 0
)

/* debug [enums.gen.go]: Processing enum NEFilterAction (5 cases) */
// NEFilterAction - The actions a data provider can take on a filter flow.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction
type NEFilterAction uint

const (
	// NEFilterActionAllow - Allow the flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction/allow
	NEFilterActionAllow NEFilterAction = 0
	// NEFilterActionDrop - Drop the flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction/drop
	NEFilterActionDrop NEFilterAction = 0
	// NEFilterActionFilterData - Filter data on the flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction/filterData
	NEFilterActionFilterData NEFilterAction = 0
	// NEFilterActionInvalid - Invalid action used to represent an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction/invalid
	NEFilterActionInvalid NEFilterAction = 0
	// NEFilterActionRemediate - Remediate the flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction/remediate
	NEFilterActionRemediate NEFilterAction = 0
)

/* debug [enums.gen.go]: Processing enum NEFilterDataAttribute (1 cases) */
// NEFilterDataAttribute - Attribute flags that describe the data handled by a filter.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataAttribute
type NEFilterDataAttribute uint

const (
	// NEFilterDataAttributeHasIPHeader - An attribute that indicates the data includes an IP header.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataAttribute/hasIPHeader
	NEFilterDataAttributeHasIPHeader NEFilterDataAttribute = 0
)

/* debug [enums.gen.go]: Processing enum NEFilterManagerGrade (2 cases) */
// NEFilterManagerGrade - A type for the grade or priority of the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/Grade-swift.enum
type NEFilterManagerGrade uint

const (
	// NEFilterManagerGradeFirewall - A grade for filters that act as firewalls, blocking some network traffic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/Grade-swift.enum/firewall
	NEFilterManagerGradeFirewall NEFilterManagerGrade = 0
	// NEFilterManagerGradeInspector - A grade for filters that act as inspectors of network traffic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/Grade-swift.enum/inspector
	NEFilterManagerGradeInspector NEFilterManagerGrade = 0
)

/* debug [enums.gen.go]: Processing enum NEFilterManagerError (6 cases) */
// NEFilterManagerError - Error codes specific to filter managers.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManagerError
type NEFilterManagerError uint

const (
	// NEFilterManagerErrorConfigurationCannotBeRemoved - An error code that indicates removing the configuration isn’t allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManagerError/configurationCannotBeRemoved
	NEFilterManagerErrorConfigurationCannotBeRemoved NEFilterManagerError = 0
	// NEFilterManagerErrorConfigurationDisabled - An error code that indicates the filter configuration isn’t enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManagerError/configurationDisabled
	NEFilterManagerErrorConfigurationDisabled NEFilterManagerError = 0
	// NEFilterManagerErrorConfigurationInternalError - An error code that indicates an internal configuration error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManagerError/configurationInternalError
	NEFilterManagerErrorConfigurationInternalError NEFilterManagerError = 0
	// NEFilterManagerErrorConfigurationInvalid - An error code that indicates the filter configuration is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManagerError/configurationInvalid
	NEFilterManagerErrorConfigurationInvalid NEFilterManagerError = 0
	// NEFilterManagerErrorConfigurationPermissionDenied - An error code that indicates the configuration lacks permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManagerError/configurationPermissionDenied
	NEFilterManagerErrorConfigurationPermissionDenied NEFilterManagerError = 0
	// NEFilterManagerErrorConfigurationStale - An error code that indicates another process modfied the filter configuration since the last time the app loaded the configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManagerError/configurationStale
	NEFilterManagerErrorConfigurationStale NEFilterManagerError = 0
)

/* debug [enums.gen.go]: Processing enum NEFilterPacketProviderVerdict (3 cases) */
// NEFilterPacketProviderVerdict - The verdict returned by a packet handler indicating what the framework should do with a packet.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/Verdict
type NEFilterPacketProviderVerdict uint

const (
	// NEFilterPacketProviderVerdictAllow - A verdict to allow a packet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/Verdict/allow
	NEFilterPacketProviderVerdictAllow NEFilterPacketProviderVerdict = 0
	// NEFilterPacketProviderVerdictDelay - A verdict to delay a packet until a future verdict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/Verdict/delay
	NEFilterPacketProviderVerdictDelay NEFilterPacketProviderVerdict = 0
	// NEFilterPacketProviderVerdictDrop - A verdict to drop a packet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/Verdict/drop
	NEFilterPacketProviderVerdictDrop NEFilterPacketProviderVerdict = 0
)

/* debug [enums.gen.go]: Processing enum NEFilterReportEvent (4 cases) */
// NEFilterReportEvent - A type that represents the kind of event indicated by a report.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Event-swift.enum
type NEFilterReportEvent uint

const (
	// NEFilterReportEventDataDecision - A type of event indicating the report is about a pass/block decision made after analyzing some amount of a flow’s data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Event-swift.enum/dataDecision
	NEFilterReportEventDataDecision NEFilterReportEvent = 0
	// NEFilterReportEventFlowClosed - A type of event indicating the report is for a flow’s closing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Event-swift.enum/flowClosed
	NEFilterReportEventFlowClosed NEFilterReportEvent = 0
	// NEFilterReportEventNewFlow - A type of event indicating the report is for a new flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Event-swift.enum/newFlow
	NEFilterReportEventNewFlow NEFilterReportEvent = 0
	// NEFilterReportEventStatistics - A type of event indicating the report is for the latest statistics of the flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Event-swift.enum/statistics
	NEFilterReportEventStatistics NEFilterReportEvent = 0
)

/* debug [enums.gen.go]: Processing enum NEFilterReportFrequency (4 cases) */
// NEFilterReportFrequency - An enumeration that represents the frequency of filter report delivery.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Frequency
type NEFilterReportFrequency uint

const (
	// NEFilterReportFrequencyHigh - A low frequency of reports, about once every half-second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Frequency/high
	NEFilterReportFrequencyHigh NEFilterReportFrequency = 0
	// NEFilterReportFrequencyLow - A low frequency of reports, about once every five seconds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Frequency/low
	NEFilterReportFrequencyLow NEFilterReportFrequency = 0
	// NEFilterReportFrequencyMedium - A low frequency of reports, about once every second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Frequency/medium
	NEFilterReportFrequencyMedium NEFilterReportFrequency = 0
	// NEFilterReportFrequencyNone - A frequency value that indicates no report delivery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Frequency/none
	NEFilterReportFrequencyNone NEFilterReportFrequency = 0
)

/* debug [enums.gen.go]: Processing enum NEHotspotConfigurationError (18 cases) */
// NEHotspotConfigurationError - Error values returned by hotspot configuration manager methods.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError
type NEHotspotConfigurationError uint

const (
	// NEHotspotConfigurationErrorAlreadyAssociated - The configuration is already associated with the hotspot.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/alreadyAssociated
	NEHotspotConfigurationErrorAlreadyAssociated NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorApplicationIsNotInForeground - The application is not running in the foreground.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/applicationIsNotInForeground
	NEHotspotConfigurationErrorApplicationIsNotInForeground NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInternal - Internal error, otherwise undefined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/internal
	NEHotspotConfigurationErrorInternal NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInvalid - The configuration is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/invalid
	NEHotspotConfigurationErrorInvalid NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInvalidEAPSettings - EAP settings are not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/invalidEAPSettings
	NEHotspotConfigurationErrorInvalidEAPSettings NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInvalidHS20DomainName - The HS 2.0 domain name is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/invalidHS20DomainName
	NEHotspotConfigurationErrorInvalidHS20DomainName NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInvalidHS20Settings - The HS 2.0 settings are not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/invalidHS20Settings
	NEHotspotConfigurationErrorInvalidHS20Settings NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInvalidSSID - The SSID value is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/invalidSSID
	NEHotspotConfigurationErrorInvalidSSID NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInvalidSSIDPrefix - The SSID prefix used to create the hotspot configuration is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/invalidSSIDPrefix
	NEHotspotConfigurationErrorInvalidSSIDPrefix NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInvalidWEPPassphrase - The WEP passphrase is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/invalidWEPPassphrase
	NEHotspotConfigurationErrorInvalidWEPPassphrase NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInvalidWPAPassphrase - The WPA passphrase is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/invalidWPAPassphrase
	NEHotspotConfigurationErrorInvalidWPAPassphrase NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorJoinOnceNotSupported - The join-once option isn’t support for EAP configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/joinOnceNotSupported
	NEHotspotConfigurationErrorJoinOnceNotSupported NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorPending - The network configuration action has not completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/pending
	NEHotspotConfigurationErrorPending NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorSystemConfiguration - The system configuration is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/systemConfiguration
	NEHotspotConfigurationErrorSystemConfiguration NEHotspotConfigurationError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/systemDenied
	NEHotspotConfigurationErrorSystemDenied NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorUnknown - An unknown error has occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/unknown
	NEHotspotConfigurationErrorUnknown NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorUserDenied - The user has refused the network configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/userDenied
	NEHotspotConfigurationErrorUserDenied NEHotspotConfigurationError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError/userUnauthorized
	NEHotspotConfigurationErrorUserUnauthorized NEHotspotConfigurationError = 0
)

/* debug [enums.gen.go]: Processing enum NEHotspotConfigurationEAPType (4 cases) */
// NEHotspotConfigurationEAPType - The EAP types that may be specified in 
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/EAPType
type NEHotspotConfigurationEAPType uint

const (
	// NEHotspotConfigurationEAPTypeEAPFAST - Network EAP type is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/EAPType/EAPFAST
	NEHotspotConfigurationEAPTypeEAPFAST NEHotspotConfigurationEAPType = 0
	// NEHotspotConfigurationEAPTypeEAPPEAP - Network EAP type is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/EAPType/EAPPEAP
	NEHotspotConfigurationEAPTypeEAPPEAP NEHotspotConfigurationEAPType = 0
	// NEHotspotConfigurationEAPTypeEAPTLS - Network EAP type is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/EAPType/EAPTLS
	NEHotspotConfigurationEAPTypeEAPTLS NEHotspotConfigurationEAPType = 0
	// NEHotspotConfigurationEAPTypeEAPTTLS - Network EAP type is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/EAPType/EAPTTLS
	NEHotspotConfigurationEAPTypeEAPTTLS NEHotspotConfigurationEAPType = 0
)

/* debug [enums.gen.go]: Processing enum NEHotspotConfigurationEAPTLSVersion (3 cases) */
// NEHotspotConfigurationEAPTLSVersion - The EAPTLS Version identifiers that may be specified by 
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/TLSVersion
type NEHotspotConfigurationEAPTLSVersion uint

const (
	// NEHotspotConfigurationEAPTLSVersion_1_0 - Network EAPTLS version 1.0.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationEAPTLSVersion/NEHotspotConfigurationEAPTLSVersion_1_0
	NEHotspotConfigurationEAPTLSVersion_1_0 NEHotspotConfigurationEAPTLSVersion = 0
	// NEHotspotConfigurationEAPTLSVersion_1_1 - Network EAPTLS version 1.1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationEAPTLSVersion/NEHotspotConfigurationEAPTLSVersion_1_1
	NEHotspotConfigurationEAPTLSVersion_1_1 NEHotspotConfigurationEAPTLSVersion = 0
	// NEHotspotConfigurationEAPTLSVersion_1_2 - Network EAPTLS version 1.2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationEAPTLSVersion/NEHotspotConfigurationEAPTLSVersion_1_2
	NEHotspotConfigurationEAPTLSVersion_1_2 NEHotspotConfigurationEAPTLSVersion = 0
)

/* debug [enums.gen.go]: Processing enum NEHotspotConfigurationTTLSInnerAuthenticationType (5 cases) */
// NEHotspotConfigurationTTLSInnerAuthenticationType - The TTLS Inner Authentication Types that may be specified by 
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/TTLSInnerAuthenticationType-swift.enum
type NEHotspotConfigurationTTLSInnerAuthenticationType uint

const (
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationCHAP - Network EAPTTLS inner authentication type is CHAP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/TTLSInnerAuthenticationType-swift.enum/eapttlsInnerAuthenticationCHAP
	NEHotspotConfigurationEAPTTLSInnerAuthenticationCHAP NEHotspotConfigurationTTLSInnerAuthenticationType = 0
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationEAP - Network EAPTTLS inner authentication type is EAP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/TTLSInnerAuthenticationType-swift.enum/eapttlsInnerAuthenticationEAP
	NEHotspotConfigurationEAPTTLSInnerAuthenticationEAP NEHotspotConfigurationTTLSInnerAuthenticationType = 0
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAP - Network EAPTTLS inner authentication type is MSCHAP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/TTLSInnerAuthenticationType-swift.enum/eapttlsInnerAuthenticationMSCHAP
	NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAP NEHotspotConfigurationTTLSInnerAuthenticationType = 0
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAPv2 - Network EAPTTLS inner authentication type is MSCHAP, version 2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/TTLSInnerAuthenticationType-swift.enum/eapttlsInnerAuthenticationMSCHAPv2
	NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAPv2 NEHotspotConfigurationTTLSInnerAuthenticationType = 0
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationPAP - Network EAPTTLS inner authentication type is PAP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/TTLSInnerAuthenticationType-swift.enum/eapttlsInnerAuthenticationPAP
	NEHotspotConfigurationEAPTTLSInnerAuthenticationPAP NEHotspotConfigurationTTLSInnerAuthenticationType = 0
)

/* debug [enums.gen.go]: Processing enum NEHotspotHelperCommandType (7 cases) */
// NEHotspotHelperCommandType - An enumeration of hotspot command types.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType
type NEHotspotHelperCommandType uint

const (
	// kNEHotspotHelperCommandTypeAuthenticate - Authenticate to the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType/authenticate
	kNEHotspotHelperCommandTypeAuthenticate NEHotspotHelperCommandType = 0
	// kNEHotspotHelperCommandTypeEvaluate - Evaluate the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType/evaluate
	kNEHotspotHelperCommandTypeEvaluate NEHotspotHelperCommandType = 0
	// kNEHotspotHelperCommandTypeFilterScanList - Filter the Wi-Fi scan list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType/filterScanList
	kNEHotspotHelperCommandTypeFilterScanList NEHotspotHelperCommandType = 0
	// kNEHotspotHelperCommandTypeLogoff - Logoff the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType/logoff
	kNEHotspotHelperCommandTypeLogoff NEHotspotHelperCommandType = 0
	// kNEHotspotHelperCommandTypeMaintain - Maintain the connection to the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType/maintain
	kNEHotspotHelperCommandTypeMaintain NEHotspotHelperCommandType = 0
	// kNEHotspotHelperCommandTypeNone - Placeholder for the null command.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType/none
	kNEHotspotHelperCommandTypeNone NEHotspotHelperCommandType = 0
	// kNEHotspotHelperCommandTypePresentUI - Present user interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType/presentUI
	kNEHotspotHelperCommandTypePresentUI NEHotspotHelperCommandType = 0
)

/* debug [enums.gen.go]: Processing enum NEHotspotHelperConfidence (3 cases) */
// NEHotspotHelperConfidence - A type that indicates the hotspot helper’s confidence in its ability to handle the network.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperConfidence
type NEHotspotHelperConfidence uint

const (
	// kNEHotspotHelperConfidenceHigh - The helper has high confidence in being able to handle the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperConfidence/high
	kNEHotspotHelperConfidenceHigh NEHotspotHelperConfidence = 0
	// kNEHotspotHelperConfidenceLow - The helper has some confidence in being able to handle the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperConfidence/low
	kNEHotspotHelperConfidenceLow NEHotspotHelperConfidence = 0
	// kNEHotspotHelperConfidenceNone - The helper is unable to handle the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperConfidence/none
	kNEHotspotHelperConfidenceNone NEHotspotHelperConfidence = 0
)

/* debug [enums.gen.go]: Processing enum NEHotspotHelperResult (7 cases) */
// NEHotspotHelperResult - The result of handling a hotspot command.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResult
type NEHotspotHelperResult uint

const (
	// kNEHotspotHelperResultAuthenticationRequired - The network requires authentication again. This result is only valid in response to a command with type  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResult/authenticationRequired
	kNEHotspotHelperResultAuthenticationRequired NEHotspotHelperResult = 0
	// kNEHotspotHelperResultCommandNotRecognized - The helper did not recognize the command type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResult/commandNotRecognized
	kNEHotspotHelperResultCommandNotRecognized NEHotspotHelperResult = 0
	// kNEHotspotHelperResultFailure - The command failed to be handled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResult/failure
	kNEHotspotHelperResultFailure NEHotspotHelperResult = 0
	// kNEHotspotHelperResultSuccess - The command was handled successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResult/success
	kNEHotspotHelperResultSuccess NEHotspotHelperResult = 0
	// kNEHotspotHelperResultTemporaryFailure - The Hotspot Helper app determined that it is temporarily unable to perform the authentication. This result is only valid in response to commands of type   and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResult/temporaryFailure
	kNEHotspotHelperResultTemporaryFailure NEHotspotHelperResult = 0
	// kNEHotspotHelperResultUIRequired - The operation requires user interaction. This result is only valid in response to a command with type  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResult/uiRequired
	kNEHotspotHelperResultUIRequired NEHotspotHelperResult = 0
	// kNEHotspotHelperResultUnsupportedNetwork - After attempting to authenticate, the Hotspot Helper app determined that it can’t perform the authentication. This result is only valid in response to commands of type   and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResult/unsupportedNetwork
	kNEHotspotHelperResultUnsupportedNetwork NEHotspotHelperResult = 0
)

/* debug [enums.gen.go]: Processing enum NEHotspotNetworkSecurityType (5 cases) */
// NEHotspotNetworkSecurityType - An enumeration of constants that define Wi-Fi hotspot network security types.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetworkSecurityType
type NEHotspotNetworkSecurityType uint

const (
	// NEHotspotNetworkSecurityTypeEnterprise - A security type to represent use of Wi-Fi protected access (WPA), WPA2, and WPA3 standards using enterprise-level seciurity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetworkSecurityType/enterprise
	NEHotspotNetworkSecurityTypeEnterprise NEHotspotNetworkSecurityType = 0
	// NEHotspotNetworkSecurityTypeOpen - A security type to represent an open network with no security protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetworkSecurityType/open
	NEHotspotNetworkSecurityTypeOpen NEHotspotNetworkSecurityType = 0
	// NEHotspotNetworkSecurityTypePersonal - A security type to represent use of Wi-Fi protected access (WPA), WPA2, and WPA3 standards using a pre-shared secret.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetworkSecurityType/personal
	NEHotspotNetworkSecurityTypePersonal NEHotspotNetworkSecurityType = 0
	// NEHotspotNetworkSecurityTypeUnknown - A value that represents an unknown security type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetworkSecurityType/unknown
	NEHotspotNetworkSecurityTypeUnknown NEHotspotNetworkSecurityType = 0
	// NEHotspotNetworkSecurityTypeWEP - A security type to represent use of Wired Equivalent Privacy (WEP).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetworkSecurityType/WEP
	NEHotspotNetworkSecurityTypeWEP NEHotspotNetworkSecurityType = 0
)

/* debug [enums.gen.go]: Processing enum NENetworkRuleProtocol (3 cases) */
// NENetworkRuleProtocol - A type to represent network protocols used by routing rules.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/Protocol
type NENetworkRuleProtocol uint

const (
	// NENetworkRuleProtocolAny - A rule protocol to match TCP and UDP traffic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/Protocol/any
	NENetworkRuleProtocolAny NENetworkRuleProtocol = 0
	// NENetworkRuleProtocolTCP - A rule protocol to match TCP traffic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/Protocol/TCP
	NENetworkRuleProtocolTCP NENetworkRuleProtocol = 0
	// NENetworkRuleProtocolUDP - A rule protocol to match UDP traffic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/Protocol/UDP
	NENetworkRuleProtocolUDP NENetworkRuleProtocol = 0
)

/* debug [enums.gen.go]: Processing enum NEOnDemandRuleAction (4 cases) */
// NEOnDemandRuleAction enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleAction
type NEOnDemandRuleAction uint

const (
	// NEOnDemandRuleActionConnect - Start the VPN connection for every connection attempt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleAction/connect
	NEOnDemandRuleActionConnect NEOnDemandRuleAction = 0
	// NEOnDemandRuleActionDisconnect - Do not start the VPN connection, and disconnect the VPN connection if it is not currently disconnected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleAction/disconnect
	NEOnDemandRuleActionDisconnect NEOnDemandRuleAction = 0
	// NEOnDemandRuleActionEvaluateConnection - Start the VPN after evaluating the destination host being accessed against the rule’s parameters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleAction/evaluateConnection
	NEOnDemandRuleActionEvaluateConnection NEOnDemandRuleAction = 0
	// NEOnDemandRuleActionIgnore - Do not start the VPN connection, but do not disconnect it if it is currently connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleAction/ignore
	NEOnDemandRuleActionIgnore NEOnDemandRuleAction = 0
)

/* debug [enums.gen.go]: Processing enum NEOnDemandRuleInterfaceType (4 cases) */
// NEOnDemandRuleInterfaceType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleInterfaceType
type NEOnDemandRuleInterfaceType uint

const (
	// NEOnDemandRuleInterfaceTypeAny - Match any interface type
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleInterfaceType/any
	NEOnDemandRuleInterfaceTypeAny NEOnDemandRuleInterfaceType = 0
	// NEOnDemandRuleInterfaceTypeCellular - Match cellular data interfaces
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleInterfaceType/cellular
	NEOnDemandRuleInterfaceTypeCellular NEOnDemandRuleInterfaceType = 0
	// NEOnDemandRuleInterfaceTypeEthernet - Match wired ethernet interfaces
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleInterfaceType/ethernet
	NEOnDemandRuleInterfaceTypeEthernet NEOnDemandRuleInterfaceType = 0
	// NEOnDemandRuleInterfaceTypeWiFi - Match Wi-Fi interfaces
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleInterfaceType/wiFi
	NEOnDemandRuleInterfaceTypeWiFi NEOnDemandRuleInterfaceType = 0
)

/* debug [enums.gen.go]: Processing enum NEProviderStopReason (18 cases) */
// NEProviderStopReason - Reasons why the provider extension was stopped.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason
type NEProviderStopReason uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/appUpdate
	NEProviderStopReasonAppUpdate NEProviderStopReason = 0
	// NEProviderStopReasonAuthenticationCanceled - The authentication process was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/authenticationCanceled
	NEProviderStopReasonAuthenticationCanceled NEProviderStopReason = 0
	// NEProviderStopReasonConfigurationDisabled - The configuration was disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/configurationDisabled
	NEProviderStopReasonConfigurationDisabled NEProviderStopReason = 0
	// NEProviderStopReasonConfigurationFailed - The configuration is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/configurationFailed
	NEProviderStopReasonConfigurationFailed NEProviderStopReason = 0
	// NEProviderStopReasonConfigurationRemoved - The configuration was removed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/configurationRemoved
	NEProviderStopReasonConfigurationRemoved NEProviderStopReason = 0
	// NEProviderStopReasonConnectionFailed - The connection failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/connectionFailed
	NEProviderStopReasonConnectionFailed NEProviderStopReason = 0
	// NEProviderStopReasonIdleTimeout - The session timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/idleTimeout
	NEProviderStopReasonIdleTimeout NEProviderStopReason = 0
	// NEProviderStopReasonInternalError - The provider encountered an internal error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/internalError
	NEProviderStopReasonInternalError NEProviderStopReason = 0
	// NEProviderStopReasonNone - No specific reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/none
	NEProviderStopReasonNone NEProviderStopReason = 0
	// NEProviderStopReasonNoNetworkAvailable - No network connectivity is currently available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/noNetworkAvailable
	NEProviderStopReasonNoNetworkAvailable NEProviderStopReason = 0
	// NEProviderStopReasonProviderDisabled - The provider was disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/providerDisabled
	NEProviderStopReasonProviderDisabled NEProviderStopReason = 0
	// NEProviderStopReasonProviderFailed - The provider failed to function correctly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/providerFailed
	NEProviderStopReasonProviderFailed NEProviderStopReason = 0
	// NEProviderStopReasonSleep - A stop reason indicating the configuration enabled disconnect on sleep and the device went to sleep.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/sleep
	NEProviderStopReasonSleep NEProviderStopReason = 0
	// NEProviderStopReasonSuperceded - The configuration was superceded by a higher-priority configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/superceded
	NEProviderStopReasonSuperceded NEProviderStopReason = 0
	// NEProviderStopReasonUnrecoverableNetworkChange - The device’s network connectivity changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/unrecoverableNetworkChange
	NEProviderStopReasonUnrecoverableNetworkChange NEProviderStopReason = 0
	// NEProviderStopReasonUserInitiated - The user stopped the provider extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/userInitiated
	NEProviderStopReasonUserInitiated NEProviderStopReason = 0
	// NEProviderStopReasonUserLogout - The user logged out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/userLogout
	NEProviderStopReasonUserLogout NEProviderStopReason = 0
	// NEProviderStopReasonUserSwitch - The current console user changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason/userSwitch
	NEProviderStopReasonUserSwitch NEProviderStopReason = 0
)

/* debug [enums.gen.go]: Processing enum NERelayManagerError (4 cases) */
// NERelayManagerError - Error codes specific to relay managers.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerError
type NERelayManagerError uint

const (
	// NERelayManagerErrorConfigurationCannotBeRemoved - An error code that indicates removing the relay manager failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerError/configurationCannotBeRemoved
	NERelayManagerErrorConfigurationCannotBeRemoved NERelayManagerError = 0
	// NERelayManagerErrorConfigurationDisabled - An error code that indicates the relay manager isn’t enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerError/configurationDisabled
	NERelayManagerErrorConfigurationDisabled NERelayManagerError = 0
	// NERelayManagerErrorConfigurationInvalid - An error code that indicates the relay manager is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerError/configurationInvalid
	NERelayManagerErrorConfigurationInvalid NERelayManagerError = 0
	// NERelayManagerErrorConfigurationStale - An error code that indicates the relay manager isn’t loaded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerError/configurationStale
	NERelayManagerErrorConfigurationStale NERelayManagerError = 0
)

/* debug [enums.gen.go]: Processing enum NETrafficDirection (3 cases) */
// NETrafficDirection - A type to represent the direction of network traffic.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETrafficDirection
type NETrafficDirection uint

const (
	// NETrafficDirectionAny - A direction that matches either inbound or outbound traffic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETrafficDirection/any
	NETrafficDirectionAny NETrafficDirection = 0
	// NETrafficDirectionInbound - The inbound traffic direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETrafficDirection/inbound
	NETrafficDirectionInbound NETrafficDirection = 0
	// NETrafficDirectionOutbound - The outbound traffic direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETrafficDirection/outbound
	NETrafficDirectionOutbound NETrafficDirection = 0
)

/* debug [enums.gen.go]: Processing enum NETunnelProviderRoutingMethod (3 cases) */
// NETunnelProviderRoutingMethod enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderRoutingMethod
type NETunnelProviderRoutingMethod uint

const (
	// NETunnelProviderRoutingMethodDestinationIP - Route network traffic to the tunnel based on destination IP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderRoutingMethod/destinationIP
	NETunnelProviderRoutingMethodDestinationIP NETunnelProviderRoutingMethod = 0
	// NETunnelProviderRoutingMethodNetworkRule - A routing method that routes traffic based on network rule objects specified by the provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderRoutingMethod/networkRule
	NETunnelProviderRoutingMethodNetworkRule NETunnelProviderRoutingMethod = 0
	// NETunnelProviderRoutingMethodSourceApplication - Route network traffic to the tunnel based on source application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderRoutingMethod/sourceApplication
	NETunnelProviderRoutingMethodSourceApplication NETunnelProviderRoutingMethod = 0
)

/* debug [enums.gen.go]: Processing enum NEURLFilterVerdict (3 cases) */
// NEURLFilterVerdict - A verdict returned by a URL filter.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter/Verdict
type NEURLFilterVerdict uint

const (
	// NEURLFilterVerdictAllow - A verdict that indicates that accessing the URL is allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter/Verdict/allow
	NEURLFilterVerdictAllow NEURLFilterVerdict = 0
	// NEURLFilterVerdictDeny - A verdict that indicates that accessing the URL is denied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter/Verdict/deny
	NEURLFilterVerdictDeny NEURLFilterVerdict = 0
	// NEURLFilterVerdictUnknown - A verdict that indicates URL validation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter/Verdict/unknown
	NEURLFilterVerdictUnknown NEURLFilterVerdict = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNConnectionError (19 cases) */
// NEVPNConnectionError - Error codes specific to VPN connections.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError
type NEVPNConnectionError uint

const (
	// NEVPNConnectionErrorAuthenticationFailed - An error code that indicates the VPN connection failed because the VPN server rejected the user credentials.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/authenticationFailed
	NEVPNConnectionErrorAuthenticationFailed NEVPNConnectionError = 0
	// NEVPNConnectionErrorClientCertificateExpired - An error code that indicates the client certfiicate’s validity period has passed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/clientCertificateExpired
	NEVPNConnectionErrorClientCertificateExpired NEVPNConnectionError = 0
	// NEVPNConnectionErrorClientCertificateInvalid - An error code that indicates the client certfiicate is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/clientCertificateInvalid
	NEVPNConnectionErrorClientCertificateInvalid NEVPNConnectionError = 0
	// NEVPNConnectionErrorClientCertificateNotYetValid - An error code that indicates the client certfiicate won’t be valid until some time in the future.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/clientCertificateNotYetValid
	NEVPNConnectionErrorClientCertificateNotYetValid NEVPNConnectionError = 0
	// NEVPNConnectionErrorConfigurationFailed - An error code that indicates the VPN connection failed because the configuration is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/configurationFailed
	NEVPNConnectionErrorConfigurationFailed NEVPNConnectionError = 0
	// NEVPNConnectionErrorConfigurationNotFound - An error code that indicates the VPN connection failed because the system couldn’t find a configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/configurationNotFound
	NEVPNConnectionErrorConfigurationNotFound NEVPNConnectionError = 0
	// NEVPNConnectionErrorNegotiationFailed - An error code that indicates the VPN connection failed because the negotiation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/negotiationFailed
	NEVPNConnectionErrorNegotiationFailed NEVPNConnectionError = 0
	// NEVPNConnectionErrorNoNetworkAvailable - An error code that indicates the VPN connection failed because the system isn’t connected to a network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/noNetworkAvailable
	NEVPNConnectionErrorNoNetworkAvailable NEVPNConnectionError = 0
	// NEVPNConnectionErrorOverslept - An error code that indicates the system slept for an extended period of time, causing the VPN connection to terminate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/overslept
	NEVPNConnectionErrorOverslept NEVPNConnectionError = 0
	// NEVPNConnectionErrorPluginDisabled - An error code that indicates the VPN plugin isn’t available or needs an update.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/pluginDisabled
	NEVPNConnectionErrorPluginDisabled NEVPNConnectionError = 0
	// NEVPNConnectionErrorPluginFailed - An error code that indicates the VPN plugin failed unexpectedly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/pluginFailed
	NEVPNConnectionErrorPluginFailed NEVPNConnectionError = 0
	// NEVPNConnectionErrorServerAddressResolutionFailed - An error code that indicates the VPN connection failed because the system couldn’t determine the VPN server address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/serverAddressResolutionFailed
	NEVPNConnectionErrorServerAddressResolutionFailed NEVPNConnectionError = 0
	// NEVPNConnectionErrorServerCertificateExpired - An error code that indicates the server certfiicate’s validity period has passed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/serverCertificateExpired
	NEVPNConnectionErrorServerCertificateExpired NEVPNConnectionError = 0
	// NEVPNConnectionErrorServerCertificateInvalid - An error code that indicates the server certfiicate is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/serverCertificateInvalid
	NEVPNConnectionErrorServerCertificateInvalid NEVPNConnectionError = 0
	// NEVPNConnectionErrorServerCertificateNotYetValid - An error code that indicates the server certfiicate won’t be valid until some time in the future.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/serverCertificateNotYetValid
	NEVPNConnectionErrorServerCertificateNotYetValid NEVPNConnectionError = 0
	// NEVPNConnectionErrorServerDead - An error code that indicates the VPN connection failed because the VPN server has stopped responding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/serverDead
	NEVPNConnectionErrorServerDead NEVPNConnectionError = 0
	// NEVPNConnectionErrorServerDisconnected - An error code that indicates the VPN connection failed because the VPN server terminated the connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/serverDisconnected
	NEVPNConnectionErrorServerDisconnected NEVPNConnectionError = 0
	// NEVPNConnectionErrorServerNotResponding - An error code that indicates the VPN connection failed because the VPN server isn’t responding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/serverNotResponding
	NEVPNConnectionErrorServerNotResponding NEVPNConnectionError = 0
	// NEVPNConnectionErrorUnrecoverableNetworkChange - An error code that indicates network conditions changed such that the VPN connection needed to terminate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError/unrecoverableNetworkChange
	NEVPNConnectionErrorUnrecoverableNetworkChange NEVPNConnectionError = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNIKEAuthenticationMethod (3 cases) */
// NEVPNIKEAuthenticationMethod - Internet Key Exchange (IKE) authentication methods used to authenticate with the IPSec server.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEAuthenticationMethod
type NEVPNIKEAuthenticationMethod uint

const (
	// NEVPNIKEAuthenticationMethodCertificate - Use a certificate and private key as the authentication credential. The certificate and private key set in the   or   property will be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEAuthenticationMethod/certificate
	NEVPNIKEAuthenticationMethodCertificate NEVPNIKEAuthenticationMethod = 0
	// NEVPNIKEAuthenticationMethodNone - Do not authenticate with the IPSec server. Note that extended authentication may still be performed if the   property is set. This value is only valid for IKE version 2 (IKEv2)
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEAuthenticationMethod/none
	NEVPNIKEAuthenticationMethodNone NEVPNIKEAuthenticationMethod = 0
	// NEVPNIKEAuthenticationMethodSharedSecret - Use a shared secret as the authentication credential. The shared secret set in the   property will be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEAuthenticationMethod/sharedSecret
	NEVPNIKEAuthenticationMethodSharedSecret NEVPNIKEAuthenticationMethod = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNIKEv2CertificateType (6 cases) */
// NEVPNIKEv2CertificateType - An enumeration of certificate type values.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType
type NEVPNIKEv2CertificateType uint

const (
	// NEVPNIKEv2CertificateTypeECDSA256 - The ECDSA with p-256 curve certificate type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType/ECDSA256
	NEVPNIKEv2CertificateTypeECDSA256 NEVPNIKEv2CertificateType = 0
	// NEVPNIKEv2CertificateTypeECDSA384 - The ECDSA with p-384 curve certificate type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType/ECDSA384
	NEVPNIKEv2CertificateTypeECDSA384 NEVPNIKEv2CertificateType = 0
	// NEVPNIKEv2CertificateTypeECDSA521 - The ECDSA with p-521 curve certificate type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType/ECDSA521
	NEVPNIKEv2CertificateTypeECDSA521 NEVPNIKEv2CertificateType = 0
	// NEVPNIKEv2CertificateTypeEd25519 - The Edwards 25519 curve certificate type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType/ed25519
	NEVPNIKEv2CertificateTypeEd25519 NEVPNIKEv2CertificateType = 0
	// NEVPNIKEv2CertificateTypeRSA - The RSA certificate type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType/RSA
	NEVPNIKEv2CertificateTypeRSA NEVPNIKEv2CertificateType = 0
	// NEVPNIKEv2CertificateTypeRSAPSS - The RSA-PSS certificate type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType/RSAPSS
	NEVPNIKEv2CertificateTypeRSAPSS NEVPNIKEv2CertificateType = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNIKEv2DeadPeerDetectionRate (4 cases) */
// NEVPNIKEv2DeadPeerDetectionRate - An enumeration of values for the frequency at which the IKEv2 client runs the dead peer detection algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DeadPeerDetectionRate
type NEVPNIKEv2DeadPeerDetectionRate uint

const (
	// NEVPNIKEv2DeadPeerDetectionRateHigh - Run dead peer detection once every 1 minute. If the peer does not respond, retry 5 times at 1 second intervals before declaring the peer dead and terminating the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DeadPeerDetectionRate/high
	NEVPNIKEv2DeadPeerDetectionRateHigh NEVPNIKEv2DeadPeerDetectionRate = 0
	// NEVPNIKEv2DeadPeerDetectionRateLow - Run dead peer detection once every 30 minutes. If the peer does not respond, retry 5 times at 1 second intervals before declaring the peer dead and terminating the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DeadPeerDetectionRate/low
	NEVPNIKEv2DeadPeerDetectionRateLow NEVPNIKEv2DeadPeerDetectionRate = 0
	// NEVPNIKEv2DeadPeerDetectionRateMedium - Run dead peer detection once every 10 minutes. If the peer does not respond, retry 5 times at 1 second intervals before declaring the peer dead and terminating the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DeadPeerDetectionRate/medium
	NEVPNIKEv2DeadPeerDetectionRateMedium NEVPNIKEv2DeadPeerDetectionRate = 0
	// NEVPNIKEv2DeadPeerDetectionRateNone - Do not perform dead peer detection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DeadPeerDetectionRate/none
	NEVPNIKEv2DeadPeerDetectionRateNone NEVPNIKEv2DeadPeerDetectionRate = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNIKEv2DiffieHellmanGroup (14 cases) */
// NEVPNIKEv2DiffieHellmanGroup - An enumeration of Diffie-Hellman group values.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup
type NEVPNIKEv2DiffieHellmanGroup uint

const (
	// NEVPNIKEv2DiffieHellmanGroup1 - Diffie Hellman group 1 (768-bit modular exponential [MODP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group1
	NEVPNIKEv2DiffieHellmanGroup1 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup14 - Diffie Hellman group 14 (2048-bit modular exponential [MODP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group14
	NEVPNIKEv2DiffieHellmanGroup14 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup15 - Diffie Hellman group 15 (3072-bit modular exponential [MODP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group15
	NEVPNIKEv2DiffieHellmanGroup15 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup16 - Diffie Hellman group 16 (4096-bit modular exponential [MODP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group16
	NEVPNIKEv2DiffieHellmanGroup16 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup17 - Diffie Hellman group 17 (6144-bit modular exponential [MODP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group17
	NEVPNIKEv2DiffieHellmanGroup17 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup18 - Diffie Hellman group 18 (8192-bit modular exponential [MODP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group18
	NEVPNIKEv2DiffieHellmanGroup18 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup19 - Diffie Hellman group 19 (256-bit random elliptic curve group over GF[P] [ECP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group19
	NEVPNIKEv2DiffieHellmanGroup19 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup2 - Diffie Hellman group 2 (1024-bit modular exponential [MODP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group2
	NEVPNIKEv2DiffieHellmanGroup2 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup20 - Diffie Hellman group 20 (384-bit random elliptic curve group over GF[P] [ECP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group20
	NEVPNIKEv2DiffieHellmanGroup20 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup21 - Diffie Hellman group 21 (521-bit random elliptic curve group over GF[P] [ECP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group21
	NEVPNIKEv2DiffieHellmanGroup21 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup31 - Diffie Hellman group 31 (Curve 25519).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group31
	NEVPNIKEv2DiffieHellmanGroup31 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup32 - Diffie Hellman group 32 (Curve 448).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group32
	NEVPNIKEv2DiffieHellmanGroup32 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroup5 - Diffie Hellman group 5 (1536-bit modular exponential [MODP]).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group5
	NEVPNIKEv2DiffieHellmanGroup5 NEVPNIKEv2DiffieHellmanGroup = 0
	// NEVPNIKEv2DiffieHellmanGroupInvalid - A value indicating the group is not a valid Diffie-Hellman group.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/groupInvalid
	NEVPNIKEv2DiffieHellmanGroupInvalid NEVPNIKEv2DiffieHellmanGroup = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNIKEv2EncryptionAlgorithm (7 cases) */
// NEVPNIKEv2EncryptionAlgorithm - An enumeration of encryption algorithm values.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm
type NEVPNIKEv2EncryptionAlgorithm uint

const (
	// NEVPNIKEv2EncryptionAlgorithm3DES - Triple Data Encryption Algorithm (aka 3DES)
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm/algorithm3DES
	NEVPNIKEv2EncryptionAlgorithm3DES NEVPNIKEv2EncryptionAlgorithm = 0
	// NEVPNIKEv2EncryptionAlgorithmAES128 - Advanced Encryption Standard 256-bit (AES256).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm/algorithmAES128
	NEVPNIKEv2EncryptionAlgorithmAES128 NEVPNIKEv2EncryptionAlgorithm = 0
	// NEVPNIKEv2EncryptionAlgorithmAES128GCM - Advanced Encryption Standard 128-bit Galois/Counter Mode (AES128GCM).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm/algorithmAES128GCM
	NEVPNIKEv2EncryptionAlgorithmAES128GCM NEVPNIKEv2EncryptionAlgorithm = 0
	// NEVPNIKEv2EncryptionAlgorithmAES256 - Advanced Encryption Standard 256 bit (AES256).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm/algorithmAES256
	NEVPNIKEv2EncryptionAlgorithmAES256 NEVPNIKEv2EncryptionAlgorithm = 0
	// NEVPNIKEv2EncryptionAlgorithmAES256GCM - Advanced Encryption Standard 256-bit Galois/Counter Mode (AES256GCM).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm/algorithmAES256GCM
	NEVPNIKEv2EncryptionAlgorithmAES256GCM NEVPNIKEv2EncryptionAlgorithm = 0
	// NEVPNIKEv2EncryptionAlgorithmChaCha20Poly1305 - ChaCha20 and Poly1305 (ChaCha20Poly1305).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm/algorithmChaCha20Poly1305
	NEVPNIKEv2EncryptionAlgorithmChaCha20Poly1305 NEVPNIKEv2EncryptionAlgorithm = 0
	// NEVPNIKEv2EncryptionAlgorithmDES - Data Encryption Standard (DES)
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm/algorithmDES
	NEVPNIKEv2EncryptionAlgorithmDES NEVPNIKEv2EncryptionAlgorithm = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNIKEv2IntegrityAlgorithm (5 cases) */
// NEVPNIKEv2IntegrityAlgorithm enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2IntegrityAlgorithm
type NEVPNIKEv2IntegrityAlgorithm uint

const (
	// NEVPNIKEv2IntegrityAlgorithmSHA160 - SHA-1 160-bit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2IntegrityAlgorithm/SHA160
	NEVPNIKEv2IntegrityAlgorithmSHA160 NEVPNIKEv2IntegrityAlgorithm = 0
	// NEVPNIKEv2IntegrityAlgorithmSHA256 - SHA-2 256-bit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2IntegrityAlgorithm/SHA256
	NEVPNIKEv2IntegrityAlgorithmSHA256 NEVPNIKEv2IntegrityAlgorithm = 0
	// NEVPNIKEv2IntegrityAlgorithmSHA384 - SHA-2 384-bit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2IntegrityAlgorithm/SHA384
	NEVPNIKEv2IntegrityAlgorithmSHA384 NEVPNIKEv2IntegrityAlgorithm = 0
	// NEVPNIKEv2IntegrityAlgorithmSHA512 - SHA-2 512-bit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2IntegrityAlgorithm/SHA512
	NEVPNIKEv2IntegrityAlgorithmSHA512 NEVPNIKEv2IntegrityAlgorithm = 0
	// NEVPNIKEv2IntegrityAlgorithmSHA96 - SHA-1 96-bit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2IntegrityAlgorithm/SHA96
	NEVPNIKEv2IntegrityAlgorithmSHA96 NEVPNIKEv2IntegrityAlgorithm = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNIKEv2PostQuantumKeyExchangeMethod (3 cases) */
// NEVPNIKEv2PostQuantumKeyExchangeMethod - Quantum-secure key exchange methods you use with IKEv2 servers.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PostQuantumKeyExchangeMethod
type NEVPNIKEv2PostQuantumKeyExchangeMethod uint

const (
	// NEVPNIKEv2PostQuantumKeyExchangeMethod36 - Instructs the server to use the ML-KEM-768 key exchange method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PostQuantumKeyExchangeMethod/method36
	NEVPNIKEv2PostQuantumKeyExchangeMethod36 NEVPNIKEv2PostQuantumKeyExchangeMethod = 0
	// NEVPNIKEv2PostQuantumKeyExchangeMethod37 - Instructs the server to use the ML-KEM-1024 key exchange method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PostQuantumKeyExchangeMethod/method37
	NEVPNIKEv2PostQuantumKeyExchangeMethod37 NEVPNIKEv2PostQuantumKeyExchangeMethod = 0
	// NEVPNIKEv2PostQuantumKeyExchangeMethodNone - Instructs the server not to use a quantum-secure key exchange method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PostQuantumKeyExchangeMethod/methodNone
	NEVPNIKEv2PostQuantumKeyExchangeMethodNone NEVPNIKEv2PostQuantumKeyExchangeMethod = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNIKEv2TLSVersion (4 cases) */
// NEVPNIKEv2TLSVersion - An enumeration of TLS Versions for use in EAP-TLS.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2TLSVersion
type NEVPNIKEv2TLSVersion uint

const (
	// NEVPNIKEv2TLSVersion1_0 - A value to use TLS version 1.0.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2TLSVersion/version1_0
	NEVPNIKEv2TLSVersion1_0 NEVPNIKEv2TLSVersion = 0
	// NEVPNIKEv2TLSVersion1_1 - A value to use TLS version 1.1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2TLSVersion/version1_1
	NEVPNIKEv2TLSVersion1_1 NEVPNIKEv2TLSVersion = 0
	// NEVPNIKEv2TLSVersion1_2 - A value to use TLS version 1.2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2TLSVersion/version1_2
	NEVPNIKEv2TLSVersion1_2 NEVPNIKEv2TLSVersion = 0
	// NEVPNIKEv2TLSVersionDefault - A value to use the default TLS configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2TLSVersion/versionDefault
	NEVPNIKEv2TLSVersionDefault NEVPNIKEv2TLSVersion = 0
)

/* debug [enums.gen.go]: Processing enum NEVPNStatus (6 cases) */
// NEVPNStatus - The possible states of a VPN connection.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus
type NEVPNStatus uint

const (
	// NEVPNStatusConnected - The VPN is connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus/connected
	NEVPNStatusConnected NEVPNStatus = 0
	// NEVPNStatusConnecting - The VPN is in the process of connecting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus/connecting
	NEVPNStatusConnecting NEVPNStatus = 0
	// NEVPNStatusDisconnected - The VPN is disconnected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus/disconnected
	NEVPNStatusDisconnected NEVPNStatus = 0
	// NEVPNStatusDisconnecting - The VPN is in the process of disconnecting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus/disconnecting
	NEVPNStatusDisconnecting NEVPNStatus = 0
	// NEVPNStatusInvalid - The associated VPN configuration doesn’t exist in the Network Extension preferences or isn’t enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus/invalid
	NEVPNStatusInvalid NEVPNStatus = 0
	// NEVPNStatusReasserting - The VPN is in the process of reconnecting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus/reasserting
	NEVPNStatusReasserting NEVPNStatus = 0
)

/* debug [enums.gen.go]: Processing enum NWPathStatus (4 cases) */
// NWPathStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPathStatus
type NWPathStatus uint

const (
	// NWPathStatusInvalid - The path cannot be evaluated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPathStatus/invalid
	NWPathStatusInvalid NWPathStatus = 0
	// NWPathStatusSatisfiable - The path is not currently satisfied, but may become satisfied upon a connection attempt. This can be due to a service, such as a VPN or a cellular data connection not being activated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPathStatus/satisfiable
	NWPathStatusSatisfiable NWPathStatus = 0
	// NWPathStatusSatisfied - The path is ready to be used for network connections.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPathStatus/satisfied
	NWPathStatusSatisfied NWPathStatus = 0
	// NWPathStatusUnsatisfied - The path for network connections is not available, either due to lack of network connectivity or being prohibited by system policy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPathStatus/unsatisfied
	NWPathStatusUnsatisfied NWPathStatus = 0
)

/* debug [enums.gen.go]: Processing enum NWTCPConnectionState (6 cases) */
// NWTCPConnectionState - Defined connection states. New types may be defined in the future.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionState
type NWTCPConnectionState uint

const (
	// NWTCPConnectionStateCancelled - The connection has been cancelled by the client calling  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionState/cancelled
	NWTCPConnectionStateCancelled NWTCPConnectionState = 0
	// NWTCPConnectionStateConnected - The connection is established. It is now possible to transfer data. If TLS is in use, the TLS handshake has finished.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionState/connected
	NWTCPConnectionStateConnected NWTCPConnectionState = 0
	// NWTCPConnectionStateConnecting - The connection is attempting to connect. This includes endpoint resolution when applicable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionState/connecting
	NWTCPConnectionStateConnecting NWTCPConnectionState = 0
	// NWTCPConnectionStateDisconnected - The connection is disconnected. It is no longer possible to transfer data. The application should call   to clean up resources.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionState/disconnected
	NWTCPConnectionStateDisconnected NWTCPConnectionState = 0
	// NWTCPConnectionStateInvalid - The connection is in an invalid or uninitialized state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionState/invalid
	NWTCPConnectionStateInvalid NWTCPConnectionState = 0
	// NWTCPConnectionStateWaiting - The connection has attempted to connect but failed. It is now waiting for better conditions before trying again.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionState/waiting
	NWTCPConnectionStateWaiting NWTCPConnectionState = 0
)

/* debug [enums.gen.go]: Processing enum NWUDPSessionState (6 cases) */
// NWUDPSessionState enum type
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSessionState
type NWUDPSessionState uint

const (
	// NWUDPSessionStateCancelled - The session has been cancelled by the client calling  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSessionState/cancelled
	NWUDPSessionStateCancelled NWUDPSessionState = 0
	// NWUDPSessionStateFailed - None of the currently resolved endpoints can be used at this time, either due to problems with the path or the client rejecting the endpoints.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSessionState/failed
	NWUDPSessionStateFailed NWUDPSessionState = 0
	// NWUDPSessionStateInvalid - The session is in an invalid or uninitialized state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSessionState/invalid
	NWUDPSessionStateInvalid NWUDPSessionState = 0
	// NWUDPSessionStatePreparing - The remote endpoint is being resolved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSessionState/preparing
	NWUDPSessionStatePreparing NWUDPSessionState = 0
	// NWUDPSessionStateReady - The session is ready for reading and writing data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSessionState/ready
	NWUDPSessionStateReady NWUDPSessionState = 0
	// NWUDPSessionStateWaiting - The session is waiting for better conditions before attempting to make the session ready.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSessionState/waiting
	NWUDPSessionStateWaiting NWUDPSessionState = 0
)


