// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

// Enum types and constants
// MEMessageActionFlag enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag
type MEMessageActionFlag uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/none
	MEMessageActionFlagNone MEMessageActionFlag = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/Flag/yellow
	MEMessageActionFlagYellow MEMessageActionFlag = 0
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
	// MEMessageStateReceived - A state that indicates the system has received and stored the message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageState/received
	MEMessageStateReceived MEMessageState = 0
)


