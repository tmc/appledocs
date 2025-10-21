// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
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
func NewINMessageWithIdentifierConversationIdentifierContentDateSentSenderRecipientsGroupNameServiceNameMessageTypeReferencedMessageReaction(identifier appkit.string, conversationIdentifier appkit.string, content appkit.string, dateSent foundation.IDate, sender INPerson, recipients []INPerson, groupName INSpeakableString, serviceName appkit.string, messageType unsafe.Pointer, referencedMessage INMessage, reaction unsafe.Pointer) INMessage {
	instance := getINMessageClass().Alloc()
	rv := objc.Send[INMessage](instance.ID, objc.Sel("initWithIdentifier:conversationIdentifier:content:dateSent:sender:recipients:groupName:serviceName:messageType:referencedMessage:reaction:"), identifier, conversationIdentifier, content, dateSent, sender, recipients, groupName, serviceName, messageType, referencedMessage, reaction)
	rv.Autorelease()
	return rv
}


// The date and time the app sent the message to each recipient.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMessage/dateSent
func (i_ INMessage) DateSent() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](i_.ID, objc.Sel("dateSent"))
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

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/attachmentfiles
func (i_ INMessage) AttachmentFiles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("attachmentFiles"))
	return rv
}


// SetAttachmentFiles sets the value of the attachmentFiles property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/attachmentfiles
func (i_ INMessage) SetAttachmentFiles(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAttachmentFiles:"), value)
}

// An audio recording that Siri plays to the message recipient.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/audiomessagefile
func (i_ INMessage) AudioMessageFile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("audioMessageFile"))
	return rv
}


// SetAudioMessageFile sets the value of the audioMessageFile property.
// An audio recording that Siri plays to the message recipient.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/audiomessagefile
func (i_ INMessage) SetAudioMessageFile(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAudioMessageFile:"), value)
}

// The text that Siri recites to the message recipient.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/content
func (i_ INMessage) Content() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("content"))
	return rv
}


// SetContent sets the value of the content property.
// The text that Siri recites to the message recipient.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/content
func (i_ INMessage) SetContent(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContent:"), value)
}

// The identifier of the conversation that contains this message.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/conversationidentifier
func (i_ INMessage) ConversationIdentifier() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("conversationIdentifier"))
	return rv
}


// SetConversationIdentifier sets the value of the conversationIdentifier property.
// The identifier of the conversation that contains this message.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/conversationidentifier
func (i_ INMessage) SetConversationIdentifier(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setConversationIdentifier:"), value)
}

// The name of the grouped conversation.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/groupname
func (i_ INMessage) GroupName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("groupName"))
	return rv
}


// SetGroupName sets the value of the groupName property.
// The name of the grouped conversation.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/groupname
func (i_ INMessage) SetGroupName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupName:"), value)
}

// The message’s unique identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/identifier
func (i_ INMessage) Identifier() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The message’s unique identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/identifier
func (i_ INMessage) SetIdentifier(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIdentifier:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/linkmetadata
func (i_ INMessage) LinkMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("linkMetadata"))
	return rv
}


// SetLinkMetadata sets the value of the linkMetadata property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/linkmetadata
func (i_ INMessage) SetLinkMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLinkMetadata:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/numberofattachments
func (i_ INMessage) NumberOfAttachments() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("numberOfAttachments"))
	return rv
}


// SetNumberOfAttachments sets the value of the numberOfAttachments property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/numberofattachments
func (i_ INMessage) SetNumberOfAttachments(value foundation.INumber) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNumberOfAttachments:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/reaction
func (i_ INMessage) Reaction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("reaction"))
	return rv
}


// SetReaction sets the value of the reaction property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/reaction
func (i_ INMessage) SetReaction(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReaction:"), value)
}

// The person who sent the message.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/sender
func (i_ INMessage) Sender() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("sender"))
	return rv
}


// SetSender sets the value of the sender property.
// The person who sent the message.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/sender
func (i_ INMessage) SetSender(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSender:"), value)
}

// The name of the service that delivers the message.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/servicename
func (i_ INMessage) ServiceName() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("serviceName"))
	return rv
}


// SetServiceName sets the value of the serviceName property.
// The name of the service that delivers the message.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/servicename
func (i_ INMessage) SetServiceName(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setServiceName:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/sticker
func (i_ INMessage) Sticker() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("sticker"))
	return rv
}


// SetSticker sets the value of the sticker property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmessage/sticker
func (i_ INMessage) SetSticker(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSticker:"), value)
}


