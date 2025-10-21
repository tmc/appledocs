// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INMessage] class.
var (
	INMessageClass     _INMessageClass
	INMessageClassOnce sync.Once
)

func getINMessageClass() _INMessageClass {
	INMessageClassOnce.Do(func() {
		INMessageClass = _INMessageClass{objc.GetClass("INMessage")}
	})
	return INMessageClass
}

type _INMessageClass struct {
	class objc.Class
}

// An interface definition for the [INMessage] class.
type IINMessage interface {
	objectivec.IObject
}

// An object that describes a sent or received message.
//
// When your app responds to an , you create instances of this class to provide Siri with information about the messages in your app. Each message contains a unique identifier, the participants’ details, and the content. You can group messages into larger discussions and, if you app supports multiple services, identify the service that delivers each message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMessage
type INMessage struct {
	objectivec.Object
}

// INMessageFrom constructs a [INMessage] from an unsafe.Pointer.
//
// An object that describes a sent or received message.
func INMessageFrom(ptr unsafe.Pointer) INMessage {
	return INMessage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INMessageClass) Alloc() INMessage {
	rv := objc.Send[INMessage](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INMessageClass) New() INMessage {
	rv := objc.Send[INMessage](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INMessage) Init() INMessage {
	rv := objc.Send[INMessage](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INMessage) Autorelease() INMessage {
	rv := objc.Send[INMessage](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINMessage creates a new INMessage instance.
func NewINMessage() INMessage {
	return getINMessageClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMessage/init(identifier:conversationIdentifier:content:dateSent:sender:recipients:groupName:serviceName:messageType:referencedMessage:reaction:)
func NewINMessageWithIdentifierConversationIdentifierContentDateSentSenderRecipientsGroupNameServiceNameMessageTypeReferencedMessageReaction(identifier string, conversationIdentifier string, content string, dateSent unsafe.Pointer, sender unsafe.Pointer, recipients unsafe.Pointer, groupName unsafe.Pointer, serviceName string, messageType unsafe.Pointer, referencedMessage unsafe.Pointer, reaction unsafe.Pointer) INMessage {
	instance := getINMessageClass().Alloc()
	rv := objc.Send[INMessage](instance.ID, objc.Sel("initWithIdentifier:conversationIdentifier:content:dateSent:sender:recipients:groupName:serviceName:messageType:referencedMessage:reaction:"), objc.String(identifier), objc.String(conversationIdentifier), objc.String(content), dateSent, sender, recipients, groupName, objc.String(serviceName), messageType, referencedMessage, reaction)
	rv.Autorelease()
	return rv
}


// The date and time the app sent the message to each recipient.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMessage/dateSent
func (i_ INMessage) DateSent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dateSent"))
	return rv
}

// The type of content the message contains.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMessage/messageType
func (i_ INMessage) MessageType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("messageType"))
	return rv
}

// The people who received the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMessage/recipients
func (i_ INMessage) Recipients() []INPerson {
	rv := objc.Send[[]INPerson](i_.ID, objc.Sel("recipients"))
	return rv
}


