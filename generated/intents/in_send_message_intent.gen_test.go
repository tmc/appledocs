// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents_test

import (
	"github.com/tmc/appledocs/generated/intents"
)

// Suppress unused import errors
var _ = intents.NewINSendMessageIntent

// ExampleNewINSendMessageIntentWithRecipientsContentGroupNameServiceNameSender demonstrates how to create a INSendMessageIntent instance using NewINSendMessageIntentWithRecipientsContentGroupNameServiceNameSender.
// Initializes a send message intent object with the specified content and recipients.
func ExampleNewINSendMessageIntentWithRecipientsContentGroupNameServiceNameSender() {
	_ = intents.NewINSendMessageIntentWithRecipientsContentGroupNameServiceNameSender(
		[]intents.INPerson{}, // recipients []INPerson
		"content",            // content string
		"groupName",          // groupName string
		"serviceName",        // serviceName string
		intents.INPerson{},   // sender INPerson
	)
	// Output:
}

// ExampleNewINSendMessageIntentWithRecipientsContentSpeakableGroupNameConversationIdentifierServiceNameSender demonstrates how to create a INSendMessageIntent instance using NewINSendMessageIntentWithRecipientsContentSpeakableGroupNameConversationIdentifierServiceNameSender.
// Initializes a send message intent object with the specified content and recipients.
func ExampleNewINSendMessageIntentWithRecipientsContentSpeakableGroupNameConversationIdentifierServiceNameSender() {
	_ = intents.NewINSendMessageIntentWithRecipientsContentSpeakableGroupNameConversationIdentifierServiceNameSender(
		[]intents.INPerson{},        // recipients []INPerson
		"content",                   // content string
		intents.INSpeakableString{}, // speakableGroupName INSpeakableString
		"conversationIdentifier",    // conversationIdentifier string
		"serviceName",               // serviceName string
		intents.INPerson{},          // sender INPerson
	)
	// Output:
}

// ExampleNewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSender demonstrates how to create a INSendMessageIntent instance using NewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSender.
// Initializes a send message intent object with the specified content and recipients.
func ExampleNewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSender() {
	_ = intents.NewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSender(
		[]intents.INPerson{},            // recipients []INPerson
		intents.INOutgoingMessageType{}, // outgoingMessageType INOutgoingMessageType
		"content",                       // content string
		intents.INSpeakableString{},     // speakableGroupName INSpeakableString
		"conversationIdentifier",        // conversationIdentifier string
		"serviceName",                   // serviceName string
		intents.INPerson{},              // sender INPerson
	)
	// Output:
}

// ExampleNewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSenderAttachments demonstrates how to create a INSendMessageIntent instance using NewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSenderAttachments.
// Creates a send message intent object with the specified content and recipients.
func ExampleNewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSenderAttachments() {
	_ = intents.NewINSendMessageIntentWithRecipientsOutgoingMessageTypeContentSpeakableGroupNameConversationIdentifierServiceNameSenderAttachments(
		[]intents.INPerson{},                // recipients []INPerson
		intents.INOutgoingMessageType{},     // outgoingMessageType INOutgoingMessageType
		"content",                           // content string
		intents.INSpeakableString{},         // speakableGroupName INSpeakableString
		"conversationIdentifier",            // conversationIdentifier string
		"serviceName",                       // serviceName string
		intents.INPerson{},                  // sender INPerson
		[]intents.INSendMessageAttachment{}, // attachments []INSendMessageAttachment
	)
	// Output:
}
