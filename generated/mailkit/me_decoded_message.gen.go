// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MEDecodedMessage] class.
type IMEDecodedMessage interface {
	objectivec.IObject
	// properties:
	Banner() IMEDecodedMessageBanner
	SetBanner(value IMEDecodedMessageBanner)
	Context() objc.IObject /* cross-framework: Data */
	SetContext(value objc.IObject /* cross-framework: Data */)
	RawData() objc.IObject /* cross-framework: Data */
	SetRawData(value objc.IObject /* cross-framework: Data */)
	SecurityInformation() unsafe.Pointer
	SetSecurityInformation(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MEDecodedMessageClass) Alloc() MEDecodedMessage {
	rv := objc.Send[MEDecodedMessage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessage/banner
func (m_ MEDecodedMessage) Banner() IMEDecodedMessageBanner {
	rv := objc.Send[MEDecodedMessageBanner](m_.ID, objc.Sel("banner"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessage/banner
func (m_ MEDecodedMessage) SetBanner(value IMEDecodedMessageBanner) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBanner:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessage/context
func (m_ MEDecodedMessage) Context() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("context"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessage/context
func (m_ MEDecodedMessage) SetContext(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContext:"), value)
}


// The decoded MIME data for a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessage/rawdata
func (m_ MEDecodedMessage) RawData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rawData"))
	return rv
}


// The decoded MIME data for a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessage/rawdata
func (m_ MEDecodedMessage) SetRawData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRawData:"), value)
}


// An object that contains encryption and digital signature information about the message content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessage/securityinformation
func (m_ MEDecodedMessage) SecurityInformation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("securityInformation"))
	return rv
}


// An object that contains encryption and digital signature information about the message content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessage/securityinformation
func (m_ MEDecodedMessage) SetSecurityInformation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecurityInformation:"), value)
}



