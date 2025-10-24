// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEEncodedOutgoingMessage */


/* debug [class_header]: Header for MEEncodedOutgoingMessage */
// The class instance for the [MEEncodedOutgoingMessage] class.
var (
	MEEncodedOutgoingMessageClass     _MEEncodedOutgoingMessageClass
	MEEncodedOutgoingMessageClassOnce sync.Once
)

func getMEEncodedOutgoingMessageClass() _MEEncodedOutgoingMessageClass {
	MEEncodedOutgoingMessageClassOnce.Do(func() {
		MEEncodedOutgoingMessageClass = _MEEncodedOutgoingMessageClass{objc.GetClass("MEEncodedOutgoingMessage")}
	})
	return MEEncodedOutgoingMessageClass
}

type _MEEncodedOutgoingMessageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEEncodedOutgoingMessage */
// An interface definition for the [MEEncodedOutgoingMessage] class.
type IMEEncodedOutgoingMessage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEEncodedOutgoingMessage */
	// properties:
	IsEncrypted() bool
	IsSigned() bool
	RawData() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEEncodedOutgoingMessage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEEncodedOutgoingMessage */
// Alloc allocates a new instance without initialization.
func (mc _MEEncodedOutgoingMessageClass) Alloc() MEEncodedOutgoingMessage {
	rv := objc.Send[MEEncodedOutgoingMessage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEEncodedOutgoingMessageClass) New() MEEncodedOutgoingMessage {
	rv := objc.Send[MEEncodedOutgoingMessage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEEncodedOutgoingMessage) Init() MEEncodedOutgoingMessage {
	rv := objc.Send[MEEncodedOutgoingMessage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEEncodedOutgoingMessage) Autorelease() MEEncodedOutgoingMessage {
	rv := objc.Send[MEEncodedOutgoingMessage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEEncodedOutgoingMessage creates a new MEEncodedOutgoingMessage instance.
func NewMEEncodedOutgoingMessage() MEEncodedOutgoingMessage {
	return getMEEncodedOutgoingMessageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEEncodedOutgoingMessage */
// An object that contains the signed or encrypted representation of a message’s RFC 2822 data.
//
// When MailKit invokes your message security handler’s method, it digitally signs and encrypts the message. After encoding the message data, create an instance of to pass back to MailKit. Set the and values to indicate how you encoded the message.


// An object that contains the signed or encrypted representation of a message’s RFC 2822 data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEncodedOutgoingMessage
type MEEncodedOutgoingMessage struct {
	objectivec.Object
}

// MEEncodedOutgoingMessageFrom constructs a [MEEncodedOutgoingMessage] from an unsafe.Pointer.
//
// An object that contains the signed or encrypted representation of a message’s RFC 2822 data.
func MEEncodedOutgoingMessageFrom(ptr unsafe.Pointer) MEEncodedOutgoingMessage {
	return MEEncodedOutgoingMessage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEEncodedOutgoingMessage */

// Creates an object that contains the outgoing message’s encoded data, and indicates if the encoder encrypted or signed the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEncodedOutgoingMessage/init(rawData:isSigned:isEncrypted:)
func NewMEEncodedOutgoingMessageWithRawDataIsSignedIsEncrypted(rawData objc.IObject /* cross-framework: NSData */, isSigned bool, isEncrypted bool) MEEncodedOutgoingMessage {
	instance := getMEEncodedOutgoingMessageClass().Alloc()
	rv := objc.Send[MEEncodedOutgoingMessage](instance.ID, objc.Sel("initWithRawData:isSigned:isEncrypted:"), rawData, isSigned, isEncrypted)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEEncodedOutgoingMessageWithRawDataIsSignedIsEncrypted */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEEncodedOutgoingMessage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEEncodedOutgoingMessage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEEncodedOutgoingMessage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEEncodedOutgoingMessage */

// A Boolean value that indicates if the message encoder encrypted the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEncodedOutgoingMessage/isEncrypted
func (m_ MEEncodedOutgoingMessage) IsEncrypted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEncrypted"))
	return rv
}/* debug [instance_properties/getter]: isEncrypted */


// A Boolean value that indicates if the message encoder signed the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEncodedOutgoingMessage/isSigned
func (m_ MEEncodedOutgoingMessage) IsSigned() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSigned"))
	return rv
}/* debug [instance_properties/getter]: isSigned */


// The encrypted, signed, or both encrypted and signed data for the outgoing message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEncodedOutgoingMessage/rawData
func (m_ MEEncodedOutgoingMessage) RawData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("rawData"))
	return rv
}/* debug [instance_properties/getter]: rawData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEEncodedOutgoingMessage */


