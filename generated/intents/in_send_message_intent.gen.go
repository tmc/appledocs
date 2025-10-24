// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	SpeakableGroupName() INSpeakableString
	Attachments() INSendMessageAttachment
	SetAttachments(value INSendMessageAttachment)
	Content() objc.IObject /* cross-framework: NSString */
	SetContent(value objc.IObject /* cross-framework: NSString */)
	ConversationIdentifier() objc.IObject /* cross-framework: NSString */
	SetConversationIdentifier(value objc.IObject /* cross-framework: NSString */)
	GroupName() objc.IObject /* cross-framework: NSString */
	SetGroupName(value objc.IObject /* cross-framework: NSString */)
	OutgoingMessageType() unsafe.Pointer
	SetOutgoingMessageType(value unsafe.Pointer)
	Recipients() INPerson
	SetRecipients(value INPerson)
	Sender() INPerson
	SetSender(value INPerson)
	ServiceName() objc.IObject /* cross-framework: NSString */
	SetServiceName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A request to send a message to the designated recipients.
//
// Siri creates an object when the user asks to send a message to one or more users. This intent object contains the message to send and the recipients of the message, which can include groups of users. Use the information in this object to construct and send the message. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler confirms the request and creates an object with the results.

// A request to send a message to the designated recipients.
//
// [Full Topic]
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

// The name of the group to receive the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntent/speakableGroupName
func (i_ INSendMessageIntent) SpeakableGroupName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("speakableGroupName"))
	return rv
}

// Audio files to include in the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/attachments
func (i_ INSendMessageIntent) Attachments() INSendMessageAttachment {
	rv := objc.Send[INSendMessageAttachment](i_.ID, objc.Sel("attachments"))
	return rv
}

// Audio files to include in the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/attachments
func (i_ INSendMessageIntent) SetAttachments(value INSendMessageAttachment) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAttachments:"), value)
}

// The content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/content
func (i_ INSendMessageIntent) Content() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("content"))
	return rv
}

// The content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/content
func (i_ INSendMessageIntent) SetContent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContent:"), value)
}

// The identifier for the conversation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/conversationidentifier
func (i_ INSendMessageIntent) ConversationIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("conversationIdentifier"))
	return rv
}

// The identifier for the conversation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/conversationidentifier
func (i_ INSendMessageIntent) SetConversationIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setConversationIdentifier:"), value)
}

// The name of the group to receive the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/groupname
func (i_ INSendMessageIntent) GroupName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("groupName"))
	return rv
}

// The name of the group to receive the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/groupname
func (i_ INSendMessageIntent) SetGroupName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupName:"), value)
}

// The format of the message contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/outgoingmessagetype
func (i_ INSendMessageIntent) OutgoingMessageType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("outgoingMessageType"))
	return rv
}

// The format of the message contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/outgoingmessagetype
func (i_ INSendMessageIntent) SetOutgoingMessageType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOutgoingMessageType:"), value)
}

// The array of users to receive the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/recipients
func (i_ INSendMessageIntent) Recipients() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("recipients"))
	return rv
}

// The array of users to receive the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/recipients
func (i_ INSendMessageIntent) SetRecipients(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRecipients:"), value)
}

// The person or account that’s sending the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/sender
func (i_ INSendMessageIntent) Sender() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("sender"))
	return rv
}

// The person or account that’s sending the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/sender
func (i_ INSendMessageIntent) SetSender(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSender:"), value)
}

// The service to use when sending the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/servicename
func (i_ INSendMessageIntent) ServiceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("serviceName"))
	return rv
}

// The service to use when sending the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintent/servicename
func (i_ INSendMessageIntent) SetServiceName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setServiceName:"), value)
}
