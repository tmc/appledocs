// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents_test

import (
	"github.com/tmc/appledocs/generated/intents"
)

// Suppress unused import errors
var _ = intents.NewINSearchForMessagesIntent

// ExampleNewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersGroupNames demonstrates how to create a INSearchForMessagesIntent instance using NewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersGroupNames.
// Creates a   intent object with the specified search criteria.
func ExampleNewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersGroupNames() {
	_ = intents.NewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersGroupNames(
		[]intents.INPerson{}, // recipients []INPerson
		[]intents.INPerson{}, // senders []INPerson
		[]intents.string{}, // searchTerms []string
		intents.INMessageAttributeOptions{}, // attributes INMessageAttributeOptions
		intents.INDateComponentsRange{}, // dateTimeRange INDateComponentsRange
		[]intents.string{}, // identifiers []string
		[]intents.string{}, // notificationIdentifiers []string
		[]intents.string{}, // groupNames []string
	)
	// Output:
}
// ExampleNewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNames demonstrates how to create a INSearchForMessagesIntent instance using NewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNames.
// Creates a   intent object with the specified search criteria.
func ExampleNewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNames() {
	_ = intents.NewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNames(
		[]intents.INPerson{}, // recipients []INPerson
		[]intents.INPerson{}, // senders []INPerson
		[]intents.string{}, // searchTerms []string
		intents.INMessageAttributeOptions{}, // attributes INMessageAttributeOptions
		intents.INDateComponentsRange{}, // dateTimeRange INDateComponentsRange
		[]intents.string{}, // identifiers []string
		[]intents.string{}, // notificationIdentifiers []string
		[]intents.INSpeakableString{}, // speakableGroupNames []INSpeakableString
	)
	// Output:
}
// ExampleNewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNamesConversationIdentifiers demonstrates how to create a INSearchForMessagesIntent instance using NewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNamesConversationIdentifiers.
// Creates a search messages intent object with the specified search criteria.
func ExampleNewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNamesConversationIdentifiers() {
	_ = intents.NewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNamesConversationIdentifiers(
		[]intents.INPerson{}, // recipients []INPerson
		[]intents.INPerson{}, // senders []INPerson
		[]intents.string{}, // searchTerms []string
		intents.INMessageAttributeOptions{}, // attributes INMessageAttributeOptions
		intents.INDateComponentsRange{}, // dateTimeRange INDateComponentsRange
		[]intents.string{}, // identifiers []string
		[]intents.string{}, // notificationIdentifiers []string
		[]intents.INSpeakableString{}, // speakableGroupNames []INSpeakableString
		[]intents.string{}, // conversationIdentifiers []string
	)
	// Output:
}
