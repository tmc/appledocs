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


