// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

// Enum types and constants
// MEComposeSessionErrorCode - Errors that indicate invalid compose session states.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSessionError/Code
type MEComposeSessionErrorCode uint

// MEComposeUserAction enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeUserAction
type MEComposeUserAction uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeUserAction/replyAll
MEComposeUserActionReplyAll MEComposeUserAction = 0
)

// MEMessageActionFlag enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag
type MEMessageActionFlag uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/orange
MEMessageActionFlagOrange MEMessageActionFlag = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/yellow
MEMessageActionFlagYellow MEMessageActionFlag = 0
)

// MEMessageEncryptionState enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageEncryptionState
type MEMessageEncryptionState uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageEncryptionState/encrypted
MEMessageEncryptionStateEncrypted MEMessageEncryptionState = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageEncryptionState/notEncrypted
MEMessageEncryptionStateNotEncrypted MEMessageEncryptionState = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageEncryptionState/unknown
MEMessageEncryptionStateUnknown MEMessageEncryptionState = 0
)

// MEMessageSecurityErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityError/Code
type MEMessageSecurityErrorCode uint

// MEMessageState - The state of a message: sent, unsent, or received.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageState
type MEMessageState uint

const (
// MEMessageStateDraft - A state that indicates the user is composing the message, and hasn’t sent it yet.
//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageState/draft
MEMessageStateDraft MEMessageState = 0
// MEMessageStateSending - A state that indicates the system is in the process of sending the message.
//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageState/sending
MEMessageStateSending MEMessageState = 0
)


