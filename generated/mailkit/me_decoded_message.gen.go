// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEDecodedMessage */


/* debug [class_header]: Header for MEDecodedMessage */
// The class instance for the [MEDecodedMessage] class.
var (
	MEDecodedMessageClass     _MEDecodedMessageClass
	MEDecodedMessageClassOnce sync.Once
)

func getMEDecodedMessageClass() _MEDecodedMessageClass {
	MEDecodedMessageClassOnce.Do(func() {
		MEDecodedMessageClass = _MEDecodedMessageClass{objc.GetClass("MEDecodedMessage")}
	})
	return MEDecodedMessageClass
}

type _MEDecodedMessageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEDecodedMessage */
// An interface definition for the [MEDecodedMessage] class.
type IMEDecodedMessage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEDecodedMessage */
	// properties:
	Banner() IMEDecodedMessageBanner
	Context() objc.IObject /* cross-framework: NSData */
	RawData() objc.IObject /* cross-framework: NSData */
	SecurityInformation() IMEMessageSecurityInformation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEDecodedMessage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEDecodedMessage */
// Alloc allocates a new instance without initialization.
func (mc _MEDecodedMessageClass) Alloc() MEDecodedMessage {
	rv := objc.Send[MEDecodedMessage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEDecodedMessageClass) New() MEDecodedMessage {
	rv := objc.Send[MEDecodedMessage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEDecodedMessage) Init() MEDecodedMessage {
	rv := objc.Send[MEDecodedMessage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEDecodedMessage) Autorelease() MEDecodedMessage {
	rv := objc.Send[MEDecodedMessage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEDecodedMessage creates a new MEDecodedMessage instance.
func NewMEDecodedMessage() MEDecodedMessage {
	return getMEDecodedMessageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEDecodedMessage */
// An object that contains the RFC 2822 data for a message, without encryption or digital signatures.
//
// When MailKit invokes your message security handler’s method, you decode the message data and return an instance of that contains unencrypted MIME data.


// An object that contains the RFC 2822 data for a message, without encryption or digital signatures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessage
type MEDecodedMessage struct {
	objectivec.Object
}

// MEDecodedMessageFrom constructs a [MEDecodedMessage] from an unsafe.Pointer.
//
// An object that contains the RFC 2822 data for a message, without encryption or digital signatures.
func MEDecodedMessageFrom(ptr unsafe.Pointer) MEDecodedMessage {
	return MEDecodedMessage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEDecodedMessage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessage/init(data:securityInformation:context:)
func NewMEDecodedMessageWithDataSecurityInformationContext(rawData objc.IObject /* cross-framework: NSData */, securityInformation IMEMessageSecurityInformation, context objc.IObject /* cross-framework: NSData */) MEDecodedMessage {
	instance := getMEDecodedMessageClass().Alloc()
	rv := objc.Send[MEDecodedMessage](instance.ID, objc.Sel("initWithData:securityInformation:context:"), rawData, securityInformation, context)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEDecodedMessageWithDataSecurityInformationContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessage/init(data:securityInformation:context:banner:)
func NewMEDecodedMessageWithDataSecurityInformationContextBanner(rawData objc.IObject /* cross-framework: NSData */, securityInformation IMEMessageSecurityInformation, context objc.IObject /* cross-framework: NSData */, banner IMEDecodedMessageBanner) MEDecodedMessage {
	instance := getMEDecodedMessageClass().Alloc()
	rv := objc.Send[MEDecodedMessage](instance.ID, objc.Sel("initWithData:securityInformation:context:banner:"), rawData, securityInformation, context, banner)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEDecodedMessageWithDataSecurityInformationContextBanner */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEDecodedMessage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEDecodedMessage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEDecodedMessage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEDecodedMessage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessage/banner
func (m_ MEDecodedMessage) Banner() IMEDecodedMessageBanner {
	rv := objc.Send[MEDecodedMessageBanner](m_.ID, objc.Sel("banner"))
	return rv
}/* debug [instance_properties/getter]: banner */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessage/context
func (m_ MEDecodedMessage) Context() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */


// The decoded MIME data for a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessage/rawData
func (m_ MEDecodedMessage) RawData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("rawData"))
	return rv
}/* debug [instance_properties/getter]: rawData */


// An object that contains encryption and digital signature information about the message content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessage/securityInformation
func (m_ MEDecodedMessage) SecurityInformation() IMEMessageSecurityInformation {
	rv := objc.Send[MEMessageSecurityInformation](m_.ID, objc.Sel("securityInformation"))
	return rv
}/* debug [instance_properties/getter]: securityInformation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEDecodedMessage */


