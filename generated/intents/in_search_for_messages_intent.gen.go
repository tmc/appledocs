// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchForMessagesIntent] class.
var (
	INSearchForMessagesIntentClass     _INSearchForMessagesIntentClass
	INSearchForMessagesIntentClassOnce sync.Once
)

func getINSearchForMessagesIntentClass() _INSearchForMessagesIntentClass {
	INSearchForMessagesIntentClassOnce.Do(func() {
		INSearchForMessagesIntentClass = _INSearchForMessagesIntentClass{objc.GetClass("INSearchForMessagesIntent")}
	})
	return INSearchForMessagesIntentClass
}

type _INSearchForMessagesIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSearchForMessagesIntent] class.
type IINSearchForMessagesIntent interface {
	IINIntent
	Attributes() INMessageAttributeOptions
	ConversationIdentifiers() []string
	ConversationIdentifiersOperator() INConditionalOperator
	DateTimeRange() INDateComponentsRange
	GroupNames() []string
	GroupNamesOperator() INConditionalOperator
	Identifiers() []string
	IdentifiersOperator() INConditionalOperator
	NotificationIdentifiers() []string
	NotificationIdentifiersOperator() INConditionalOperator
	Recipients() []INPerson
	RecipientsOperator() INConditionalOperator
	SearchTerms() []string
	SearchTermsOperator() INConditionalOperator
	Senders() []INPerson
	SendersOperator() INConditionalOperator
	SpeakableGroupNames() []INSpeakableString
	SpeakableGroupNamesOperator() INConditionalOperator
}

// A request to list the messages that match the specified criteria.
//
// Siri creates objects when the user asks to see sent or received messages. You must implement this intent to support the reading of messages by Siri. This intent object contains the values for you to match when searching the user’s messages. Users can search for messages involving a specific person, messages with specific sent or received dates, or messages containing specific terms. When performing the search, use only the provided parameters, and ignore any that have no values. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler confirms the request and creates an object with the results of the search. For successful searches, Siri offers the user a way to launch your app and see the results. To allow a user wearing AirPods to automatically hear messages, you must implement both and . Add to the options when calling . Finally, add to the category option and to the category intent identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent
type INSearchForMessagesIntent struct {
	INIntent
}

// INSearchForMessagesIntentFrom constructs a [INSearchForMessagesIntent] from an unsafe.Pointer.
//
// A request to list the messages that match the specified criteria.
func INSearchForMessagesIntentFrom(ptr unsafe.Pointer) INSearchForMessagesIntent {
	return INSearchForMessagesIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchForMessagesIntentClass) Alloc() INSearchForMessagesIntent {
	rv := objc.Send[INSearchForMessagesIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchForMessagesIntentClass) New() INSearchForMessagesIntent {
	rv := objc.Send[INSearchForMessagesIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchForMessagesIntent) Init() INSearchForMessagesIntent {
	rv := objc.Send[INSearchForMessagesIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchForMessagesIntent) Autorelease() INSearchForMessagesIntent {
	rv := objc.Send[INSearchForMessagesIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchForMessagesIntent creates a new INSearchForMessagesIntent instance.
func NewINSearchForMessagesIntent() INSearchForMessagesIntent {
	return getINSearchForMessagesIntentClass().New()
}




// Creates a intent object with the specified search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/init(recipients:senders:searchTerms:attributes:dateTimeRange:identifiers:notificationIdentifiers:groupNames:)
func NewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersGroupNames(recipients []INPerson, senders []INPerson, searchTerms []string, attributes INMessageAttributeOptions, dateTimeRange INDateComponentsRange, identifiers []string, notificationIdentifiers []string, groupNames []string) INSearchForMessagesIntent {
	instance := getINSearchForMessagesIntentClass().Alloc()
	rv := objc.Send[INSearchForMessagesIntent](instance.ID, objc.Sel("initWithRecipients:senders:searchTerms:attributes:dateTimeRange:identifiers:notificationIdentifiers:groupNames:"), recipients, senders, searchTerms, attributes, dateTimeRange, identifiers, notificationIdentifiers, groupNames)
	rv.Autorelease()
	return rv
}



// Creates a intent object with the specified search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/init(recipients:senders:searchTerms:attributes:dateTimeRange:identifiers:notificationIdentifiers:speakableGroupNames:)
func NewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNames(recipients []INPerson, senders []INPerson, searchTerms []string, attributes INMessageAttributeOptions, dateTimeRange INDateComponentsRange, identifiers []string, notificationIdentifiers []string, speakableGroupNames []INSpeakableString) INSearchForMessagesIntent {
	instance := getINSearchForMessagesIntentClass().Alloc()
	rv := objc.Send[INSearchForMessagesIntent](instance.ID, objc.Sel("initWithRecipients:senders:searchTerms:attributes:dateTimeRange:identifiers:notificationIdentifiers:speakableGroupNames:"), recipients, senders, searchTerms, attributes, dateTimeRange, identifiers, notificationIdentifiers, speakableGroupNames)
	rv.Autorelease()
	return rv
}



// Creates a search messages intent object with the specified search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/init(recipients:senders:searchTerms:attributes:dateTime:identifiers:notificationIdentifiers:speakableGroupNames:conversationIdentifiers:)
func NewINSearchForMessagesIntentWithRecipientsSendersSearchTermsAttributesDateTimeRangeIdentifiersNotificationIdentifiersSpeakableGroupNamesConversationIdentifiers(recipients []INPerson, senders []INPerson, searchTerms []string, attributes INMessageAttributeOptions, dateTimeRange INDateComponentsRange, identifiers []string, notificationIdentifiers []string, speakableGroupNames []INSpeakableString, conversationIdentifiers []string) INSearchForMessagesIntent {
	instance := getINSearchForMessagesIntentClass().Alloc()
	rv := objc.Send[INSearchForMessagesIntent](instance.ID, objc.Sel("initWithRecipients:senders:searchTerms:attributes:dateTimeRange:identifiers:notificationIdentifiers:speakableGroupNames:conversationIdentifiers:"), recipients, senders, searchTerms, attributes, dateTimeRange, identifiers, notificationIdentifiers, speakableGroupNames, conversationIdentifiers)
	rv.Autorelease()
	return rv
}


// The attributes that must be present on a message to yield a match.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/attributes
func (i_ INSearchForMessagesIntent) Attributes() INMessageAttributeOptions {
	rv := objc.Send[INMessageAttributeOptions](i_.ID, objc.Sel("attributes"))
	return rv
}

// The conversation identifiers to locate in your search.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/conversationIdentifiers
func (i_ INSearchForMessagesIntent) ConversationIdentifiers() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("conversationIdentifiers"))
	return rv
}

// The operator that defines how to use the conversation identifiers in the search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/conversationIdentifiersOperator
func (i_ INSearchForMessagesIntent) ConversationIdentifiersOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("conversationIdentifiersOperator"))
	return rv
}

// The range of dates in which to search for messages.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/dateTimeRange
func (i_ INSearchForMessagesIntent) DateTimeRange() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("dateTimeRange"))
	return rv
}

// The names of any groups associated with the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/groupNames
func (i_ INSearchForMessagesIntent) GroupNames() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("groupNames"))
	return rv
}

// The operator that defines how to use the group names in the search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/groupNamesOperator
func (i_ INSearchForMessagesIntent) GroupNamesOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("groupNamesOperator"))
	return rv
}

// The message identifiers to locate in your search.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/identifiers
func (i_ INSearchForMessagesIntent) Identifiers() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("identifiers"))
	return rv
}

// The operator that defines how to use the identifiers in the search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/identifiersOperator
func (i_ INSearchForMessagesIntent) IdentifiersOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("identifiersOperator"))
	return rv
}

// The notification identifiers to locate in your search.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/notificationIdentifiers
func (i_ INSearchForMessagesIntent) NotificationIdentifiers() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("notificationIdentifiers"))
	return rv
}

// The operator that defines how to use the notification identifiers in the search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/notificationIdentifiersOperator
func (i_ INSearchForMessagesIntent) NotificationIdentifiersOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("notificationIdentifiersOperator"))
	return rv
}

// The contacts who are the recipients of the messages.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/recipients
func (i_ INSearchForMessagesIntent) Recipients() []INPerson {
	rv := objc.Send[[]INPerson](i_.ID, objc.Sel("recipients"))
	return rv
}

// The operator that defines how to use the recipients in the search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/recipientsOperator
func (i_ INSearchForMessagesIntent) RecipientsOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("recipientsOperator"))
	return rv
}

// The terms to look for in the messages.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/searchTerms
func (i_ INSearchForMessagesIntent) SearchTerms() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("searchTerms"))
	return rv
}

// The operator that defines how to use the set of terms in the search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/searchTermsOperator
func (i_ INSearchForMessagesIntent) SearchTermsOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("searchTermsOperator"))
	return rv
}

// The senders to include or exclude when you search for messages.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/senders
func (i_ INSearchForMessagesIntent) Senders() []INPerson {
	rv := objc.Send[[]INPerson](i_.ID, objc.Sel("senders"))
	return rv
}

// The operator that defines how to use the senders in the search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/sendersOperator
func (i_ INSearchForMessagesIntent) SendersOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("sendersOperator"))
	return rv
}

// The names of any groups associated with the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/speakableGroupNames
func (i_ INSearchForMessagesIntent) SpeakableGroupNames() []INSpeakableString {
	rv := objc.Send[[]INSpeakableString](i_.ID, objc.Sel("speakableGroupNames"))
	return rv
}

// The operator that defines how to use the group names in the search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntent/speakableGroupNamesOperator
func (i_ INSearchForMessagesIntent) SpeakableGroupNamesOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("speakableGroupNamesOperator"))
	return rv
}


