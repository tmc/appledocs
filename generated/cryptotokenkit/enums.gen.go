// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

// Enum types and constants
// TKErrorCode - Error codes from CryptoTokenKit.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code
type TKErrorCode uint

const (
// TKErrorCodeCanceledByUser - The operation was canceled by the user.
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/canceledByUser
TKErrorCodeCanceledByUser TKErrorCode = 0
// TKErrorCodeCommunicationError - A communication error occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/communicationError
TKErrorCodeCommunicationError TKErrorCode = 0
// TKErrorCodeCorruptedData - The data was corrupted.
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/corruptedData
TKErrorCodeCorruptedData TKErrorCode = 0
// TKErrorCodeObjectNotFound - The object was not found.
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/objectNotFound
TKErrorCodeObjectNotFound TKErrorCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKErrorCode/TKErrorObjectNotFound
TKErrorObjectNotFound TKErrorCode = 0
)

// TKSmartCardSlotState - All smart card slot states.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum
type TKSmartCardSlotState uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/empty
TKSmartCardSlotStateEmpty TKSmartCardSlotState = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/missing
TKSmartCardSlotStateMissing TKSmartCardSlotState = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/muteCard
TKSmartCardSlotStateMuteCard TKSmartCardSlotState = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/probing
TKSmartCardSlotStateProbing TKSmartCardSlotState = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/validCard
TKSmartCardSlotStateValidCard TKSmartCardSlotState = 0
)


