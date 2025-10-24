// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEMessage */


/* debug [class_header]: Header for MEMessage */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEMessage */
// An interface definition for the [MEMessage] class.
type IMEMessage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEMessage */
	// properties:
	AllRecipientAddresses() []MEEmailAddress
	BccAddresses() []MEEmailAddress
	CcAddresses() []MEEmailAddress
	DateReceived() objc.IObject /* cross-framework: NSDate */
	DateSent() objc.IObject /* cross-framework: NSDate */
	EncryptionState() MEMessageEncryptionState
	FromAddress() IMEEmailAddress
	Headers() foundation.IDictionary
	RawData() objc.IObject /* cross-framework: NSData */
	ReplyToAddresses() []MEEmailAddress
	State() MEMessageState
	Subject() objc.IObject /* cross-framework: NSString */
	ToAddresses() []MEEmailAddress
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEMessage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEMessage */
// Alloc allocates a new instance without initialization.
func (mc _MEMessageClass) Alloc() MEMessage {
	rv := objc.Send[MEMessage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEMessage */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEMessage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEMessage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEMessage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEMessage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEMessage */

// An array of email addresses for all recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/allRecipientAddresses
func (m_ MEMessage) AllRecipientAddresses() []MEEmailAddress {
	rv := objc.Send[[]MEEmailAddress](m_.ID, objc.Sel("allRecipientAddresses"))
	return rv
}/* debug [instance_properties/getter]: allRecipientAddresses */


// An array of email addresses for the concealed tertiary recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/bccAddresses
func (m_ MEMessage) BccAddresses() []MEEmailAddress {
	rv := objc.Send[[]MEEmailAddress](m_.ID, objc.Sel("bccAddresses"))
	return rv
}/* debug [instance_properties/getter]: bccAddresses */


// An array of email addresses for the secondary recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/ccAddresses
func (m_ MEMessage) CcAddresses() []MEEmailAddress {
	rv := objc.Send[[]MEEmailAddress](m_.ID, objc.Sel("ccAddresses"))
	return rv
}/* debug [instance_properties/getter]: ccAddresses */


// The date that the recipient received the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/dateReceived
func (m_ MEMessage) DateReceived() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("dateReceived"))
	return rv
}/* debug [instance_properties/getter]: dateReceived */


// The date the sender sent the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/dateSent
func (m_ MEMessage) DateSent() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("dateSent"))
	return rv
}/* debug [instance_properties/getter]: dateSent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/encryptionState
func (m_ MEMessage) EncryptionState() MEMessageEncryptionState {
	rv := objc.Send[MEMessageEncryptionState](m_.ID, objc.Sel("encryptionState"))
	return rv
}/* debug [instance_properties/getter]: encryptionState */


// The sender’s email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/fromAddress
func (m_ MEMessage) FromAddress() IMEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("fromAddress"))
	return rv
}/* debug [instance_properties/getter]: fromAddress */


// A dictionary that contains the message’s header values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/headers
func (m_ MEMessage) Headers() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("headers"))
	return rv
}/* debug [instance_properties/getter]: headers */


// The raw RFC 2822 header and body content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/rawData
func (m_ MEMessage) RawData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("rawData"))
	return rv
}/* debug [instance_properties/getter]: rawData */


// An array of email addresses to use when replying to the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/replyToAddresses
func (m_ MEMessage) ReplyToAddresses() []MEEmailAddress {
	rv := objc.Send[[]MEEmailAddress](m_.ID, objc.Sel("replyToAddresses"))
	return rv
}/* debug [instance_properties/getter]: replyToAddresses */


// The state of the mail message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/state
func (m_ MEMessage) State() MEMessageState {
	rv := objc.Send[MEMessageState](m_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The subject of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/subject
func (m_ MEMessage) Subject() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subject"))
	return rv
}/* debug [instance_properties/getter]: subject */


// An array of email addresses for the primary recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessage/toAddresses
func (m_ MEMessage) ToAddresses() []MEEmailAddress {
	rv := objc.Send[[]MEEmailAddress](m_.ID, objc.Sel("toAddresses"))
	return rv
}/* debug [instance_properties/getter]: toAddresses */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEMessage */



