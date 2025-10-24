// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

/* debug [enums.gen.go]: Generating 7 enums for MailKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MEMessageSecurityErrorCode (2 cases) */
// MEMessageSecurityErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityError/Code
type MEMessageSecurityErrorCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityError/Code/decodingError
	MEMessageSecurityDecodingError MEMessageSecurityErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityError/Code/encodingError
	MEMessageSecurityEncodingError MEMessageSecurityErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum MEMessageState (3 cases) */
// MEMessageState - The state of a message: sent, unsent, or received.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageState
type MEMessageState uint

const (
	// MEMessageStateDraft - A state that indicates the user is composing the message, and hasn’t sent it yet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageState/draft
	MEMessageStateDraft MEMessageState = 0
	// MEMessageStateReceived - A state that indicates the system has received and stored the message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageState/received
	MEMessageStateReceived MEMessageState = 0
	// MEMessageStateSending - A state that indicates the system is in the process of sending the message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageState/sending
	MEMessageStateSending MEMessageState = 0
)

/* debug [enums.gen.go]: Processing enum MEComposeSessionErrorCode (3 cases) */
// MEComposeSessionErrorCode - Errors that indicate invalid compose session states.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSessionError/Code
type MEComposeSessionErrorCode uint

const (
	// MEComposeSessionErrorCodeInvalidBody - An error code that indicates the message’s body is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSessionError/Code/invalidBody
	MEComposeSessionErrorCodeInvalidBody MEComposeSessionErrorCode = 0
	// MEComposeSessionErrorCodeInvalidHeaders - An error code that indicates one or more of the message’s headers are invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSessionError/Code/invalidHeaders
	MEComposeSessionErrorCodeInvalidHeaders MEComposeSessionErrorCode = 0
	// MEComposeSessionErrorCodeInvalidRecipients - An error code that indicates one or more of the message’s recipients are invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSessionError/Code/invalidRecipients
	MEComposeSessionErrorCodeInvalidRecipients MEComposeSessionErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum MEComposeUserAction (4 cases) */
// MEComposeUserAction enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeUserAction
type MEComposeUserAction uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeUserAction/forward
	MEComposeUserActionForward MEComposeUserAction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeUserAction/newMessage
	MEComposeUserActionNewMessage MEComposeUserAction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeUserAction/reply
	MEComposeUserActionReply MEComposeUserAction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeUserAction/replyAll
	MEComposeUserActionReplyAll MEComposeUserAction = 0
)

/* debug [enums.gen.go]: Processing enum MEMessageActionFlag (9 cases) */
// MEMessageActionFlag enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag
type MEMessageActionFlag uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/blue
	MEMessageActionFlagBlue MEMessageActionFlag = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/defaultColor
	MEMessageActionFlagDefaultColor MEMessageActionFlag = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/gray
	MEMessageActionFlagGray MEMessageActionFlag = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/green
	MEMessageActionFlagGreen MEMessageActionFlag = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/none
	MEMessageActionFlagNone MEMessageActionFlag = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/orange
	MEMessageActionFlagOrange MEMessageActionFlag = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/purple
	MEMessageActionFlagPurple MEMessageActionFlag = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/red
	MEMessageActionFlagRed MEMessageActionFlag = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/yellow
	MEMessageActionFlagYellow MEMessageActionFlag = 0
)

/* debug [enums.gen.go]: Processing enum MEMessageActionMessageColor (8 cases) */
// MEMessageActionMessageColor - A color that the system uses to display a message in the message list.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/MessageColor
type MEMessageActionMessageColor uint

const (
	// MEMessageActionMessageColorBlue - Sets the color of the message to blue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/MessageColor/blue
	MEMessageActionMessageColorBlue MEMessageActionMessageColor = 0
	// MEMessageActionMessageColorGray - Sets the color of the message to gray.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/MessageColor/gray
	MEMessageActionMessageColorGray MEMessageActionMessageColor = 0
	// MEMessageActionMessageColorGreen - Sets the color of the message to green.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/MessageColor/green
	MEMessageActionMessageColorGreen MEMessageActionMessageColor = 0
	// MEMessageActionMessageColorNone - Clears the color of the message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/MessageColor/none
	MEMessageActionMessageColorNone MEMessageActionMessageColor = 0
	// MEMessageActionMessageColorOrange - Sets the color of the message to orange.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/MessageColor/orange
	MEMessageActionMessageColorOrange MEMessageActionMessageColor = 0
	// MEMessageActionMessageColorPurple - Sets the color of the message to purple.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/MessageColor/purple
	MEMessageActionMessageColorPurple MEMessageActionMessageColor = 0
	// MEMessageActionMessageColorRed - Sets the color of the message to red.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/MessageColor/red
	MEMessageActionMessageColorRed MEMessageActionMessageColor = 0
	// MEMessageActionMessageColorYellow - Sets the color of the message to yellow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/MessageColor/yellow
	MEMessageActionMessageColorYellow MEMessageActionMessageColor = 0
)

/* debug [enums.gen.go]: Processing enum MEMessageEncryptionState (3 cases) */
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


