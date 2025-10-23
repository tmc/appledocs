// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEMessage] class.
var (
	MEMessageClass     _MEMessageClass
	MEMessageClassOnce sync.Once
)

func getMEMessageClass() _MEMessageClass {
	MEMessageClassOnce.Do(func() {
		MEMessageClass = _MEMessageClass{objc.GetClass("MEMessage")}
	})
	return MEMessageClass
}

type _MEMessageClass struct {
	class objc.Class
}

// An interface definition for the [MEMessage] class.
type IMEMessage interface {
	objectivec.IObject
	// properties:
	AllRecipientAddresses() []MEEmailAddress
	State() MEMessageState
	Subject() string
	BccAddresses() IMEEmailAddress
	SetBccAddresses(value IMEEmailAddress)
	CcAddresses() IMEEmailAddress
	SetCcAddresses(value IMEEmailAddress)
	EncryptionState() unsafe.Pointer
	SetEncryptionState(value unsafe.Pointer)
	FromAddress() IMEEmailAddress
	SetFromAddress(value IMEEmailAddress)
	Headers() string
	SetHeaders(value string)
	RawData() foundation.Data
	SetRawData(value foundation.Data)
	ReplyToAddresses() IMEEmailAddress
	SetReplyToAddresses(value IMEEmailAddress)
	ToAddresses() IMEEmailAddress
	SetToAddresses(value IMEEmailAddress)
	// methods:
}

// An object that contains information about a mail message, such as the subject, addressees, date sent, and the message contents.


// An object that contains information about a mail message, such as the subject, addressees, date sent, and the message contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage
type MEMessage struct {
	objectivec.Object
}

// MEMessageFrom constructs a [MEMessage] from an unsafe.Pointer.
//
// An object that contains information about a mail message, such as the subject, addressees, date sent, and the message contents.
func MEMessageFrom(ptr unsafe.Pointer) MEMessage {
	return MEMessage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEMessageClass) Alloc() MEMessage {
	rv := objc.Send[MEMessage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEMessageClass) New() MEMessage {
	rv := objc.Send[MEMessage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEMessage) Init() MEMessage {
	rv := objc.Send[MEMessage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEMessage) Autorelease() MEMessage {
	rv := objc.Send[MEMessage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEMessage creates a new MEMessage instance.
func NewMEMessage() MEMessage {
	return getMEMessageClass().New()
}



// An array of email addresses for all recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/allRecipientAddresses
func (m_ MEMessage) AllRecipientAddresses() []MEEmailAddress {
	rv := objc.Send[[]MEEmailAddress](m_.ID, objc.Sel("allRecipientAddresses"))
	return rv
}


// The state of the mail message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/state
func (m_ MEMessage) State() MEMessageState {
	rv := objc.Send[MEMessageState](m_.ID, objc.Sel("state"))
	return rv
}


// The subject of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/subject
func (m_ MEMessage) Subject() string {
	rv := objc.Send[string](m_.ID, objc.Sel("subject"))
	return rv
}


// An array of email addresses for the concealed tertiary recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/bccaddresses
func (m_ MEMessage) BccAddresses() IMEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("bccAddresses"))
	return rv
}


// An array of email addresses for the concealed tertiary recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/bccaddresses
func (m_ MEMessage) SetBccAddresses(value IMEEmailAddress) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBccAddresses:"), value)
}


// An array of email addresses for the secondary recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/ccaddresses
func (m_ MEMessage) CcAddresses() IMEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("ccAddresses"))
	return rv
}


// An array of email addresses for the secondary recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/ccaddresses
func (m_ MEMessage) SetCcAddresses(value IMEEmailAddress) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCcAddresses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/encryptionstate
func (m_ MEMessage) EncryptionState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("encryptionState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/encryptionstate
func (m_ MEMessage) SetEncryptionState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEncryptionState:"), value)
}


// The sender’s email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/fromaddress
func (m_ MEMessage) FromAddress() IMEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("fromAddress"))
	return rv
}


// The sender’s email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/fromaddress
func (m_ MEMessage) SetFromAddress(value IMEEmailAddress) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFromAddress:"), value)
}


// A dictionary that contains the message’s header values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/headers
func (m_ MEMessage) Headers() string {
	rv := objc.Send[string](m_.ID, objc.Sel("headers"))
	return rv
}


// A dictionary that contains the message’s header values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/headers
func (m_ MEMessage) SetHeaders(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeaders:"), objc.String(value))
}


// The raw RFC 2822 header and body content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/rawdata
func (m_ MEMessage) RawData() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rawData"))
	return rv
}


// The raw RFC 2822 header and body content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/rawdata
func (m_ MEMessage) SetRawData(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRawData:"), value)
}


// An array of email addresses to use when replying to the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/replytoaddresses
func (m_ MEMessage) ReplyToAddresses() IMEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("replyToAddresses"))
	return rv
}


// An array of email addresses to use when replying to the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/replytoaddresses
func (m_ MEMessage) SetReplyToAddresses(value IMEEmailAddress) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReplyToAddresses:"), value)
}


// An array of email addresses for the primary recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/toaddresses
func (m_ MEMessage) ToAddresses() IMEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("toAddresses"))
	return rv
}


// An array of email addresses for the primary recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/memessage/toaddresses
func (m_ MEMessage) SetToAddresses(value IMEEmailAddress) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setToAddresses:"), value)
}




