// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that contains information about a mail message, such as the subject, addressees, date sent, and the message contents.
//
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
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/allRecipientAddresses
func (m_ MEMessage) AllRecipientAddresses() []MEEmailAddress {
	rv := objc.Send[[]MEEmailAddress](m_.ID, objc.Sel("allRecipientAddresses"))
	return rv
}

// The date that the recipient received the message.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/dateReceived
func (m_ MEMessage) DateReceived() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dateReceived"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/encryptionState
func (m_ MEMessage) EncryptionState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("encryptionState"))
	return rv
}

// A dictionary that contains the message’s header values.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/headers
func (m_ MEMessage) Headers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("headers"))
	return rv
}

// The subject of the message.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/subject
func (m_ MEMessage) Subject() string {
	rv := objc.Send[string](m_.ID, objc.Sel("subject"))
	return rv
}




