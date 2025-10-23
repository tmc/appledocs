// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

// Enum types and constants
// CWChannelBand - CoreWLAN channel bands.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelBand
type CWChannelBand uint

// CWChannelWidth - CoreWLAN channel widths.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannelWidth
type CWChannelWidth uint

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

// CWErr enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr
type CWErr uint

const (
	// kCWChallengeFailureErr - Authentication was rejected because of a challenge failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwChallengeFailureErr
	kCWChallengeFailureErr CWErr = 0
	// kCWInvalidFormatErr - Invalid protocol element field detected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidFormatErr
	kCWInvalidFormatErr CWErr = 0
	// kCWInvalidInformationElementErr - Invalid information element included in association request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidInformationElementErr
	kCWInvalidInformationElementErr CWErr = 0
	// kCWInvalidPMKErr - PMK rejected by the access point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidPMKErr
	kCWInvalidPMKErr CWErr = 0
	// kCWInvalidParameterErr - Parameter error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwInvalidParameterErr
	kCWInvalidParameterErr CWErr = 0
	// kCWNotSupportedErr - Operation not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwNotSupportedErr
	kCWNotSupportedErr CWErr = 0
	// kCWOperationNotPermittedErr - Calling process does not have permission to perform this operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwOperationNotPermittedErr
	kCWOperationNotPermittedErr CWErr = 0
	// kCWTimeoutErr - Authentication/Association timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwTimeoutErr
	kCWTimeoutErr CWErr = 0
	// kCWEAPOLErr - EAPOL-related error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cweapolErr
	kCWEAPOLErr CWErr = 0
	// kCWPCOTransitionTimeNotSupportedErr - Association was denied because the requesting station does not support the PCO transition time required by the AP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWErr/cwpcoTransitionTimeNotSupportedErr
	kCWPCOTransitionTimeNotSupportedErr CWErr = 0
)

// CWEventType - Wi-Fi event types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWEventType
type CWEventType uint

// CWIBSSModeSecurity - IBSS mode security types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWIBSSModeSecurity
type CWIBSSModeSecurity uint

// CWInterfaceMode - Wi-Fi interface operating modes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode
type CWInterfaceMode uint

const (
	// kCWInterfaceModeIBSS - Interface is participating in an IBSS network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode/IBSS
	kCWInterfaceModeIBSS CWInterfaceMode = 0
	// kCWInterfaceModeHostAP - Interface is participating in an infrastructure network as an access point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode/hostAP
	kCWInterfaceModeHostAP CWInterfaceMode = 0
	// kCWInterfaceModeNone - Interface is not in any mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode/none
	kCWInterfaceModeNone CWInterfaceMode = 0
	// kCWInterfaceModeStation - Interface is participating in an infrastructure network as a non-AP station.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode/station
	kCWInterfaceModeStation CWInterfaceMode = 0
)

// CWKeychainDomain - Keychain domain types that CoreWLAN keychain methods use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDomain
type CWKeychainDomain uint

// CWPHYMode - CoreWLAN physical layer modes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode
type CWPHYMode uint

// CWSecurity - CoreWLAN security types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWSecurity
type CWSecurity uint


