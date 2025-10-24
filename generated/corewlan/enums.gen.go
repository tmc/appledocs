// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

/* debug [enums.gen.go]: Generating 10 enums for CoreWLAN */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum CWCipherKeyFlags (5 cases) */
// CWCipherKeyFlags - Cipher key flags.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWCipherKeyFlags
type CWCipherKeyFlags uint

const (
	// kCWCipherKeyFlagsNone - Open System authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWCipherKeyFlags/kCWCipherKeyFlagsNone
	kCWCipherKeyFlagsNone CWCipherKeyFlags = 0
	// kCWCipherKeyFlagsMulticast - A flag that indicates to use the cipher key for multicast packets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWCipherKeyFlags/multicast
	kCWCipherKeyFlagsMulticast CWCipherKeyFlags = 0
	// kCWCipherKeyFlagsRx - A flag that indicates to use the cipher key for packets received by the interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWCipherKeyFlags/rx
	kCWCipherKeyFlagsRx CWCipherKeyFlags = 0
	// kCWCipherKeyFlagsTx - A flag that indicates to use the cipher key for packets sent from the interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWCipherKeyFlags/tx
	kCWCipherKeyFlagsTx CWCipherKeyFlags = 0
	// kCWCipherKeyFlagsUnicast - A flag that indicates to use the cipher key for unicast packets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWCipherKeyFlags/unicast
	kCWCipherKeyFlagsUnicast CWCipherKeyFlags = 0
)

/* debug [enums.gen.go]: Processing enum CWChannelBand (4 cases) */
// CWChannelBand - CoreWLAN channel bands.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelBand
type CWChannelBand uint

const (
	// kCWChannelBand2GHz - 2.4GHz channel band.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelBand/band2GHz
	kCWChannelBand2GHz CWChannelBand = 0
	// kCWChannelBand5GHz - 5GHz channel band.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelBand/band5GHz
	kCWChannelBand5GHz CWChannelBand = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelBand/band6GHz
	kCWChannelBand6GHz CWChannelBand = 0
	// kCWChannelBandUnknown - Unknown channel band.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelBand/bandUnknown
	kCWChannelBandUnknown CWChannelBand = 0
)

/* debug [enums.gen.go]: Processing enum CWChannelWidth (5 cases) */
// CWChannelWidth - CoreWLAN channel widths.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelWidth
type CWChannelWidth uint

const (
	// kCWChannelWidth160MHz - 160MHz channel width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelWidth/width160MHz
	kCWChannelWidth160MHz CWChannelWidth = 0
	// kCWChannelWidth20MHz - 20MHz channel width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelWidth/width20MHz
	kCWChannelWidth20MHz CWChannelWidth = 0
	// kCWChannelWidth40MHz - 40MHz channel width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelWidth/width40MHz
	kCWChannelWidth40MHz CWChannelWidth = 0
	// kCWChannelWidth80MHz - 80MHz channel width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelWidth/width80MHz
	kCWChannelWidth80MHz CWChannelWidth = 0
	// kCWChannelWidthUnknown - Unknown channel width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelWidth/widthUnknown
	kCWChannelWidthUnknown CWChannelWidth = 0
)

/* debug [enums.gen.go]: Processing enum CWErr (34 cases) */
// CWErr enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr
type CWErr uint

const (
	// kCWAPFullErr - Access point is unable to handle another associated station.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwapFullErr
	kCWAPFullErr CWErr = 0
	// kCWAssociationDeniedErr - Association was denied for an unspecified reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwAssociationDeniedErr
	kCWAssociationDeniedErr CWErr = 0
	// kCWAuthenticationAlgorithmUnsupportedErr - Specified authentication algorithm is not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwAuthenticationAlgorithmUnsupportedErr
	kCWAuthenticationAlgorithmUnsupportedErr CWErr = 0
	// kCWChallengeFailureErr - Authentication was rejected because of a challenge failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwChallengeFailureErr
	kCWChallengeFailureErr CWErr = 0
	// kCWCipherSuiteRejectedErr - Cipher suite rejected due to network security policy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwCipherSuiteRejectedErr
	kCWCipherSuiteRejectedErr CWErr = 0
	// kCWDSSSOFDMUnsupportedErr - Association denied because DSSS-OFDM is not supported by requesting station.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwdsssofdmUnsupportedErr
	kCWDSSSOFDMUnsupportedErr CWErr = 0
	// kCWEAPOLErr - EAPOL-related error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cweapolErr
	kCWEAPOLErr CWErr = 0
	// kCWErr - Generic error, no specific error code exists to describe the error condition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwErr
	kCWErr CWErr = 0
	// kCWHTFeaturesNotSupportedErr - Association was denied because the requesting station does not support HT features.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwhtFeaturesNotSupportedErr
	kCWHTFeaturesNotSupportedErr CWErr = 0
	// kCWInvalidAKMPErr - Invalid authentication selector requested.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidAKMPErr
	kCWInvalidAKMPErr CWErr = 0
	// kCWInvalidAuthenticationSequenceNumberErr - Authentication frame received with an authentication sequence number out of expected sequence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidAuthenticationSequenceNumberErr
	kCWInvalidAuthenticationSequenceNumberErr CWErr = 0
	// kCWInvalidFormatErr - Invalid protocol element field detected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidFormatErr
	kCWInvalidFormatErr CWErr = 0
	// kCWInvalidGroupCipherErr - Invalid group cipher requested.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidGroupCipherErr
	kCWInvalidGroupCipherErr CWErr = 0
	// kCWInvalidInformationElementErr - Invalid information element included in association request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidInformationElementErr
	kCWInvalidInformationElementErr CWErr = 0
	// kCWInvalidPairwiseCipherErr - Invalid pairwise cipher requested.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidPairwiseCipherErr
	kCWInvalidPairwiseCipherErr CWErr = 0
	// kCWInvalidParameterErr - Parameter error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidParameterErr
	kCWInvalidParameterErr CWErr = 0
	// kCWInvalidPMKErr - PMK rejected by the access point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidPMKErr
	kCWInvalidPMKErr CWErr = 0
	// kCWInvalidRSNCapabilitiesErr - Invalid RSN capabilities specified in association request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidRSNCapabilitiesErr
	kCWInvalidRSNCapabilitiesErr CWErr = 0
	// kCWIPCFailureErr - Error communicating with a separate process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwipcFailureErr
	kCWIPCFailureErr CWErr = 0
	// kCWNoErr - Success.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwNoErr
	kCWNoErr CWErr = 0
	// kCWNoMemoryErr - Memory allocation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwNoMemoryErr
	kCWNoMemoryErr CWErr = 0
	// kCWNotSupportedErr - Operation not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwNotSupportedErr
	kCWNotSupportedErr CWErr = 0
	// kCWOperationNotPermittedErr - Calling process does not have permission to perform this operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwOperationNotPermittedErr
	kCWOperationNotPermittedErr CWErr = 0
	// kCWPCOTransitionTimeNotSupportedErr - Association was denied because the requesting station does not support the PCO transition time required by the AP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwpcoTransitionTimeNotSupportedErr
	kCWPCOTransitionTimeNotSupportedErr CWErr = 0
	// kCWReassociationDeniedErr - Reassociation was denied because the access point was unable to determine that an association exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwReassociationDeniedErr
	kCWReassociationDeniedErr CWErr = 0
	// kCWReferenceNotBoundErr - No interface is bound to the CWInterface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwReferenceNotBoundErr
	kCWReferenceNotBoundErr CWErr = 0
	// kCWShortSlotUnsupportedErr - Association denied because short slot time option is not supported by requesting station.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwShortSlotUnsupportedErr
	kCWShortSlotUnsupportedErr CWErr = 0
	// kCWSupplicantTimeoutErr - WPA/WPA2 handshake timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwSupplicantTimeoutErr
	kCWSupplicantTimeoutErr CWErr = 0
	// kCWTimeoutErr - Authentication/Association timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwTimeoutErr
	kCWTimeoutErr CWErr = 0
	// kCWUnknownErr - Unexpected error condition encountered for which no error code exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwUnknownErr
	kCWUnknownErr CWErr = 0
	// kCWUnspecifiedFailureErr - Access point did not specify a reason for authentication/association failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwUnspecifiedFailureErr
	kCWUnspecifiedFailureErr CWErr = 0
	// kCWUnsupportedCapabilitiesErr - Access point cannot support all requested capabilities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwUnsupportedCapabilitiesErr
	kCWUnsupportedCapabilitiesErr CWErr = 0
	// kCWUnsupportedRateSetErr - Interface does not support all of the rates in the access point’s basic rate set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwUnsupportedRateSetErr
	kCWUnsupportedRateSetErr CWErr = 0
	// kCWUnsupportedRSNVersionErr - Invalid WPA/WPA2 version specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwUnsupportedRSNVersionErr
	kCWUnsupportedRSNVersionErr CWErr = 0
)

/* debug [enums.gen.go]: Processing enum CWEventType (11 cases) */
// CWEventType - Wi-Fi event types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType
type CWEventType uint

const (
	// CWEventTypeBSSIDDidChange - Posts when the current BSSID of any Wi-Fi interface changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/bssidDidChange
	CWEventTypeBSSIDDidChange CWEventType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/btCoexStats
	CWEventTypeBtCoexStats CWEventType = 0
	// CWEventTypeCountryCodeDidChange - Posts when the adopted country code of any Wi-Fi interface changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/countryCodeDidChange
	CWEventTypeCountryCodeDidChange CWEventType = 0
	// CWEventTypeLinkDidChange - Posts when the link state for any Wi-Fi interface changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/linkDidChange
	CWEventTypeLinkDidChange CWEventType = 0
	// CWEventTypeLinkQualityDidChange - Posts when the RSSI or transmit rate for any Wi-Fi interface changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/linkQualityDidChange
	CWEventTypeLinkQualityDidChange CWEventType = 0
	// CWEventTypeModeDidChange - Posts when the operating mode of any Wi-Fi interface changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/modeDidChange
	CWEventTypeModeDidChange CWEventType = 0
	// CWEventTypeNone - No specified event type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/none
	CWEventTypeNone CWEventType = 0
	// CWEventTypePowerDidChange - Posts when the power state of any Wi-Fi interface changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/powerDidChange
	CWEventTypePowerDidChange CWEventType = 0
	// CWEventTypeScanCacheUpdated - Posts when the scan cache of any Wi-Fi interface is updated with new scan results.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/scanCacheUpdated
	CWEventTypeScanCacheUpdated CWEventType = 0
	// CWEventTypeSSIDDidChange - Posts when the current SSID of any Wi-Fi interface changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/ssidDidChange
	CWEventTypeSSIDDidChange CWEventType = 0
	// CWEventTypeUnknown - Unknown event type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/unknown
	CWEventTypeUnknown CWEventType = 0
)

/* debug [enums.gen.go]: Processing enum CWIBSSModeSecurity (3 cases) */
// CWIBSSModeSecurity - IBSS mode security types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWIBSSModeSecurity
type CWIBSSModeSecurity uint

const (
	// kCWIBSSModeSecurityNone - Open System authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWIBSSModeSecurity/none
	kCWIBSSModeSecurityNone CWIBSSModeSecurity = 0
	// kCWIBSSModeSecurityWEP104 - WPA Personal authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWIBSSModeSecurity/WEP104
	kCWIBSSModeSecurityWEP104 CWIBSSModeSecurity = 0
	// kCWIBSSModeSecurityWEP40 - WEP security.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWIBSSModeSecurity/WEP40
	kCWIBSSModeSecurityWEP40 CWIBSSModeSecurity = 0
)

/* debug [enums.gen.go]: Processing enum CWInterfaceMode (4 cases) */
// CWInterfaceMode - Wi-Fi interface operating modes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode
type CWInterfaceMode uint

const (
	// kCWInterfaceModeHostAP - Interface is participating in an infrastructure network as an access point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode/hostAP
	kCWInterfaceModeHostAP CWInterfaceMode = 0
	// kCWInterfaceModeIBSS - Interface is participating in an IBSS network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode/IBSS
	kCWInterfaceModeIBSS CWInterfaceMode = 0
	// kCWInterfaceModeNone - Interface is not in any mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode/none
	kCWInterfaceModeNone CWInterfaceMode = 0
	// kCWInterfaceModeStation - Interface is participating in an infrastructure network as a non-AP station.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode/station
	kCWInterfaceModeStation CWInterfaceMode = 0
)

/* debug [enums.gen.go]: Processing enum CWKeychainDomain (3 cases) */
// CWKeychainDomain - Keychain domain types that CoreWLAN keychain methods use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDomain
type CWKeychainDomain uint

const (
	// kCWKeychainDomainNone - No keychain domain specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDomain/none
	kCWKeychainDomainNone CWKeychainDomain = 0
	// kCWKeychainDomainSystem - The system keychain domain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDomain/system
	kCWKeychainDomainSystem CWKeychainDomain = 0
	// kCWKeychainDomainUser - The user keychain domain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDomain/user
	kCWKeychainDomainUser CWKeychainDomain = 0
)

/* debug [enums.gen.go]: Processing enum CWPHYMode (7 cases) */
// CWPHYMode - CoreWLAN physical layer modes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode
type CWPHYMode uint

const (
	// kCWPHYMode11a - IEEE 802.11a PHY.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode/mode11a
	kCWPHYMode11a CWPHYMode = 0
	// kCWPHYMode11ac - IEEE 802.11ac PHY.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode/mode11ac
	kCWPHYMode11ac CWPHYMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode/mode11ax
	kCWPHYMode11ax CWPHYMode = 0
	// kCWPHYMode11b - IEEE 802.11b PHY.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode/mode11b
	kCWPHYMode11b CWPHYMode = 0
	// kCWPHYMode11g - IEEE 802.11g PHY.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode/mode11g
	kCWPHYMode11g CWPHYMode = 0
	// kCWPHYMode11n - IEEE 802.11n PHY.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode/mode11n
	kCWPHYMode11n CWPHYMode = 0
	// kCWPHYModeNone - No specified mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode/modeNone
	kCWPHYModeNone CWPHYMode = 0
)

/* debug [enums.gen.go]: Processing enum CWSecurity (17 cases) */
// CWSecurity - CoreWLAN security types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity
type CWSecurity uint

const (
	// kCWSecurityDynamicWEP - Dynamic WEP security.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/dynamicWEP
	kCWSecurityDynamicWEP CWSecurity = 0
	// kCWSecurityEnterprise - Enterprise authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/enterprise
	kCWSecurityEnterprise CWSecurity = 0
	// kCWSecurityNone - Open System authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/none
	kCWSecurityNone CWSecurity = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/OWE
	kCWSecurityOWE CWSecurity = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/oweTransition
	kCWSecurityOWETransition CWSecurity = 0
	// kCWSecurityPersonal - Personal authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/personal
	kCWSecurityPersonal CWSecurity = 0
	// kCWSecurityUnknown - Unknown security type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/unknown
	kCWSecurityUnknown CWSecurity = 0
	// kCWSecurityWEP - WEP security.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/WEP
	kCWSecurityWEP CWSecurity = 0
	// kCWSecurityWPA2Enterprise - WPA2 Enterprise authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/wpa2Enterprise
	kCWSecurityWPA2Enterprise CWSecurity = 0
	// kCWSecurityWPA2Personal - WPA2 Personal authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/wpa2Personal
	kCWSecurityWPA2Personal CWSecurity = 0
	// kCWSecurityWPA3Enterprise - WPA3 Enterprise authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/wpa3Enterprise
	kCWSecurityWPA3Enterprise CWSecurity = 0
	// kCWSecurityWPA3Personal - WPA3 Personal authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/wpa3Personal
	kCWSecurityWPA3Personal CWSecurity = 0
	// kCWSecurityWPA3Transition - WPA3 Transition (WPA3/WPA2 Personal) authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/wpa3Transition
	kCWSecurityWPA3Transition CWSecurity = 0
	// kCWSecurityWPAEnterprise - WPA Enterprise authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/wpaEnterprise
	kCWSecurityWPAEnterprise CWSecurity = 0
	// kCWSecurityWPAEnterpriseMixed - WPA/WPA2 Enterprise authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/wpaEnterpriseMixed
	kCWSecurityWPAEnterpriseMixed CWSecurity = 0
	// kCWSecurityWPAPersonal - WPA Personal authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/wpaPersonal
	kCWSecurityWPAPersonal CWSecurity = 0
	// kCWSecurityWPAPersonalMixed - WPA/WPA2 Personal authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity/wpaPersonalMixed
	kCWSecurityWPAPersonalMixed CWSecurity = 0
)


