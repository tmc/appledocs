// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

/* debug [enums.gen.go]: Generating 9 enums for CryptoTokenKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum TKErrorCode (12 cases) */
// TKErrorCode - Error codes from CryptoTokenKit.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code
type TKErrorCode uint

const (
	// TKErrorCodeAuthenticationFailed - Authentication failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/authenticationFailed
	TKErrorCodeAuthenticationFailed TKErrorCode = 0
	// TKErrorCodeAuthenticationNeeded - Authentication is needed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/authenticationNeeded
	TKErrorCodeAuthenticationNeeded TKErrorCode = 0
	// TKErrorCodeBadParameter - An invalid parameter was provided.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/badParameter
	TKErrorCodeBadParameter TKErrorCode = 0
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
	// TKErrorCodeNotImplemented - The functionality is not implemented.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/notImplemented
	TKErrorCodeNotImplemented TKErrorCode = 0
	// TKErrorCodeObjectNotFound - The object was not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/objectNotFound
	TKErrorCodeObjectNotFound TKErrorCode = 0
	// TKErrorCodeTokenNotFound - The token was not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/tokenNotFound
	TKErrorCodeTokenNotFound TKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKErrorCode/TKErrorAuthenticationFailed
	TKErrorAuthenticationFailed TKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKErrorCode/TKErrorObjectNotFound
	TKErrorObjectNotFound TKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKErrorCode/TKErrorTokenNotFound
	TKErrorTokenNotFound TKErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum TKSmartCardSlotState (5 cases) */
// TKSmartCardSlotState - All smart card slot states.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum
type TKSmartCardSlotState uint

const (
	// TKSmartCardSlotStateEmpty - The Smart Card reader slot is empty; no card is inserted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/empty
	TKSmartCardSlotStateEmpty TKSmartCardSlotState = 0
	// TKSmartCardSlotStateMissing - The Smart Card reader slot is no longer known to the system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/missing
	TKSmartCardSlotStateMissing TKSmartCardSlotState = 0
	// TKSmartCardSlotStateMuteCard - A Smart Card is inserted, but is mute, or does not provide responses to commands.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/muteCard
	TKSmartCardSlotStateMuteCard TKSmartCardSlotState = 0
	// TKSmartCardSlotStateProbing - A Smart Card was inserted into the slot and an initial probe is in underway.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/probing
	TKSmartCardSlotStateProbing TKSmartCardSlotState = 0
	// TKSmartCardSlotStateValidCard - A Smart Card is inserted and properly answered to a reset command.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum/validCard
	TKSmartCardSlotStateValidCard TKSmartCardSlotState = 0
)

/* debug [enums.gen.go]: Processing enum TKSmartCardPINCharset (3 cases) */
// TKSmartCardPINCharset - Possible PIN character sets.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Charset-swift.enum
type TKSmartCardPINCharset uint

const (
	// TKSmartCardPINCharsetAlphanumeric - PIN can be composed of digits and letters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Charset-swift.enum/alphanumeric
	TKSmartCardPINCharsetAlphanumeric TKSmartCardPINCharset = 0
	// TKSmartCardPINCharsetNumeric - PIN is only composed of digits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Charset-swift.enum/numeric
	TKSmartCardPINCharsetNumeric TKSmartCardPINCharset = 0
	// TKSmartCardPINCharsetUpperAlphanumeric - PIN can be composed of digits and uppercase letters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Charset-swift.enum/upperAlphanumeric
	TKSmartCardPINCharsetUpperAlphanumeric TKSmartCardPINCharset = 0
)

/* debug [enums.gen.go]: Processing enum TKSmartCardPINEncoding (3 cases) */
// TKSmartCardPINEncoding - Possible PIN encoding types.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Encoding-swift.enum
type TKSmartCardPINEncoding uint

const (
	// TKSmartCardPINEncodingASCII - Characters are encoded in ASCII format (for example,   is encoded as  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Encoding-swift.enum/ascii
	TKSmartCardPINEncodingASCII TKSmartCardPINEncoding = 0
	// TKSmartCardPINEncodingBCD - Characters (only digits) are encoded in BCD format (for example,   is encoded as  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Encoding-swift.enum/bcd
	TKSmartCardPINEncodingBCD TKSmartCardPINEncoding = 0
	// TKSmartCardPINEncodingBinary - Characters are encoded in Binary format (for example,   is encoded as  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Encoding-swift.enum/binary
	TKSmartCardPINEncodingBinary TKSmartCardPINEncoding = 0
)

/* debug [enums.gen.go]: Processing enum TKSmartCardPINJustification (2 cases) */
// TKSmartCardPINJustification - Possible PIN justification types
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Justification
type TKSmartCardPINJustification uint

const (
	// TKSmartCardPINJustificationLeft - Justify to the left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Justification/left
	TKSmartCardPINJustificationLeft TKSmartCardPINJustification = 0
	// TKSmartCardPINJustificationRight - Justify to the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Justification/right
	TKSmartCardPINJustificationRight TKSmartCardPINJustification = 0
)

/* debug [enums.gen.go]: Processing enum TKSmartCardProtocol (5 cases) */
// TKSmartCardProtocol - Smart Card transmission protocols.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol
type TKSmartCardProtocol uint

const (
	// TKSmartCardProtocolAny - Any available transmission protocols.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/any
	TKSmartCardProtocolAny TKSmartCardProtocol = 0
	// TKSmartCardProtocolT0 - T=0 transmission protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/t0
	TKSmartCardProtocolT0 TKSmartCardProtocol = 0
	// TKSmartCardProtocolT1 - T=1 transmission protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/t1
	TKSmartCardProtocolT1 TKSmartCardProtocol = 0
	// TKSmartCardProtocolT15 - T=15 transmission protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/t15
	TKSmartCardProtocolT15 TKSmartCardProtocol = 0
	// TKSmartCardProtocolNone - No transmission protocols.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/TKSmartCardProtocolNone
	TKSmartCardProtocolNone TKSmartCardProtocol = 0
)

/* debug [enums.gen.go]: Processing enum TKSmartCardPINCompletion (3 cases) */
// TKSmartCardPINCompletion enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/Completion
type TKSmartCardPINCompletion uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/Completion/key
	TKSmartCardPINCompletionKey TKSmartCardPINCompletion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/Completion/maxLength
	TKSmartCardPINCompletionMaxLength TKSmartCardPINCompletion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/Completion/timeout
	TKSmartCardPINCompletionTimeout TKSmartCardPINCompletion = 0
)

/* debug [enums.gen.go]: Processing enum TKSmartCardPINConfirmation (3 cases) */
// TKSmartCardPINConfirmation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange/Confirmation
type TKSmartCardPINConfirmation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINConfirmation/TKSmartCardPINConfirmationNone
	TKSmartCardPINConfirmationNone TKSmartCardPINConfirmation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange/Confirmation/current
	TKSmartCardPINConfirmationCurrent TKSmartCardPINConfirmation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange/Confirmation/new
	TKSmartCardPINConfirmationNew TKSmartCardPINConfirmation = 0
)

/* debug [enums.gen.go]: Processing enum TKTokenOperation (5 cases) */
// TKTokenOperation - Operations that can be performed with a token’s keys and certificates.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation
type TKTokenOperation uint

const (
	// TKTokenOperationDecryptData - Decrypt data using a private key
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation/decryptData
	TKTokenOperationDecryptData TKTokenOperation = 0
	// TKTokenOperationNone - No operation
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation/none
	TKTokenOperationNone TKTokenOperation = 0
	// TKTokenOperationPerformKeyExchange - Perform a Diffie-Hellman style cryptographic key exchange using a private key
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation/performKeyExchange
	TKTokenOperationPerformKeyExchange TKTokenOperation = 0
	// TKTokenOperationReadData - Read raw data of a certificate
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation/readData
	TKTokenOperationReadData TKTokenOperation = 0
	// TKTokenOperationSignData - Create a cryptographic signature using a private key
	//
	// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation/signData
	TKTokenOperationSignData TKTokenOperation = 0
)


