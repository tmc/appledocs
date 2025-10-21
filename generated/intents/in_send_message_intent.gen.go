// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSendMessageIntent] class.
var (
	INSendMessageIntentClass     _INSendMessageIntentClass
	INSendMessageIntentClassOnce sync.Once
)

func getINSendMessageIntentClass() _INSendMessageIntentClass {
	INSendMessageIntentClassOnce.Do(func() {
		INSendMessageIntentClass = _INSendMessageIntentClass{objc.GetClass("INSendMessageIntent")}
	})
	return INSendMessageIntentClass
}

type _INSendMessageIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSendMessageIntent] class.
type IINSendMessageIntent interface {
	IINIntent
}

// A request to send a message to the designated recipients.
//
// Siri creates an object when the user asks to send a message to one or more users. This intent object contains the message to send and the recipients of the message, which can include groups of users. Use the information in this object to construct and send the message. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler confirms the request and creates an object with the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent
type INSendMessageIntent struct {
	INIntent
}

// INSendMessageIntentFrom constructs a [INSendMessageIntent] from an unsafe.Pointer.
//
// A request to send a message to the designated recipients.
func INSendMessageIntentFrom(ptr unsafe.Pointer) INSendMessageIntent {
	return INSendMessageIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSendMessageIntentClass) Alloc() INSendMessageIntent {
	rv := objc.Send[INSendMessageIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSendMessageIntentClass) New() INSendMessageIntent {
	rv := objc.Send[INSendMessageIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSendMessageIntent) Init() INSendMessageIntent {
	rv := objc.Send[INSendMessageIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSendMessageIntent) Autorelease() INSendMessageIntent {
	rv := objc.Send[INSendMessageIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSendMessageIntent creates a new INSendMessageIntent instance.
func NewINSendMessageIntent() INSendMessageIntent {
	return getINSendMessageIntentClass().New()
}




// Initializes a send message intent object with the specified content and recipients.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/init(recipients:content:groupName:serviceName:sender:)
func NewINSendMessageIntentWithRecipientsContentGroupNameServiceNameSender(recipients unsafe.Pointer, content string, groupName string, serviceName string, sender unsafe.Pointer) INSendMessageIntent {
	instance := getINSendMessageIntentClass().Alloc()
	rv := objc.Send[INSendMessageIntent](instance.ID, objc.Sel("initWithRecipients:content:groupName:serviceName:sender:"), recipients, objc.String(content), objc.String(groupName), objc.String(serviceName), sender)
	rv.Autorelease()
	return rv
}



// Initializes a send message intent object with the specified content and recipients.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/init(recipients:content:speakableGroupName:conversationIdentifier:serviceName:sender:)
func NewINSendMessageIntentWithRecipientsContentSpeakableGroupNameConversationIdentifierServiceNameSender(recipients unsafe.Pointer, content string, speakableGroupName unsafe.Pointer, conversationIdentifier string, serviceName string, sender unsafe.Pointer) INSendMessageIntent {
	instance := getINSendMessageIntentClass().Alloc()
	rv := objc.Send[INSendMessageIntent](instance.ID, objc.Sel("initWithRecipients:content:speakableGroupName:conversationIdentifier:serviceName:sender:"), recipients, objc.String(content), speakableGroupName, objc.String(conversationIdentifier), objc.String(serviceName), sender)
	rv.Autorelease()
	return rv
}



// Initializes a send message intent object with the specified content and recipients.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/init(recipients:outgoingMessageType:content:speakableGroupName:conversationIdentifier:serviceName:sender:)
func NewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSender(recipients unsafe.Pointer, outgoingMessageType unsafe.Pointer, content string, speakableGroupName unsafe.Pointer, conversationIdentifier string, serviceName string, sender unsafe.Pointer) INSendMessageIntent {
	instance := getINSendMessageIntentClass().Alloc()
	rv := objc.Send[INSendMessageIntent](instance.ID, objc.Sel("initWithRecipients:outgoingMessageType:content:speakableGroupName:conversationIdentifier:serviceName:sender:"), recipients, outgoingMessageType, objc.String(content), speakableGroupName, objc.String(conversationIdentifier), objc.String(serviceName), sender)
	rv.Autorelease()
	return rv
}



// Creates a send message intent object with the specified content and recipients.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/init(recipients:outgoingMessageType:content:speakableGroupName:conversationIdentifier:serviceName:sender:attachments:)
func NewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSenderAttachments(recipients unsafe.Pointer, outgoingMessageType unsafe.Pointer, content string, speakableGroupName unsafe.Pointer, conversationIdentifier string, serviceName string, sender unsafe.Pointer, attachments unsafe.Pointer) INSendMessageIntent {
	instance := getINSendMessageIntentClass().Alloc()
	rv := objc.Send[INSendMessageIntent](instance.ID, objc.Sel("initWithRecipients:outgoingMessageType:content:speakableGroupName:conversationIdentifier:serviceName:sender:attachments:"), recipients, outgoingMessageType, objc.String(content), speakableGroupName, objc.String(conversationIdentifier), objc.String(serviceName), sender, attachments)
	rv.Autorelease()
	return rv
}


// Audio files to include in the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/attachments
func (i_ INSendMessageIntent) Attachments() []INSendMessageAttachment {
	rv := objc.Send[[]INSendMessageAttachment](i_.ID, objc.Sel("attachments"))
	return rv
}

// The content of the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/content
func (i_ INSendMessageIntent) Content() string {
	rv := objc.Send[string](i_.ID, objc.Sel("content"))
	return rv
}

// The identifier for the conversation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/conversationIdentifier
func (i_ INSendMessageIntent) ConversationIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("conversationIdentifier"))
	return rv
}

// The name of the group to receive the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/groupName
func (i_ INSendMessageIntent) GroupName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("groupName"))
	return rv
}

// The format of the message contents.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/outgoingMessageType
func (i_ INSendMessageIntent) OutgoingMessageType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("outgoingMessageType"))
	return rv
}

// The array of users to receive the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/recipients
func (i_ INSendMessageIntent) Recipients() []INPerson {
	rv := objc.Send[[]INPerson](i_.ID, objc.Sel("recipients"))
	return rv
}

// The person or account that’s sending the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/sender
func (i_ INSendMessageIntent) Sender() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("sender"))
	return rv
}

// The service to use when sending the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/serviceName
func (i_ INSendMessageIntent) ServiceName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("serviceName"))
	return rv
}

// The name of the group to receive the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/speakableGroupName
func (i_ INSendMessageIntent) SpeakableGroupName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("speakableGroupName"))
	return rv
}


