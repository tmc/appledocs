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
	// properties:
	Attributes() unsafe.Pointer
	SetAttributes(value unsafe.Pointer)
	ConversationIdentifiers() string /* primitive/slice/pointer. */
	SetConversationIdentifiers(value string /* primitive/slice/pointer. */)
	ConversationIdentifiersOperator() unsafe.Pointer
	SetConversationIdentifiersOperator(value unsafe.Pointer)
	DateTimeRange() INDateComponentsRange /* already interface */
	SetDateTimeRange(value INDateComponentsRange /* already interface */)
	GroupNames() string /* primitive/slice/pointer. */
	SetGroupNames(value string /* primitive/slice/pointer. */)
	GroupNamesOperator() unsafe.Pointer
	SetGroupNamesOperator(value unsafe.Pointer)
	Identifiers() string /* primitive/slice/pointer. */
	SetIdentifiers(value string /* primitive/slice/pointer. */)
	IdentifiersOperator() unsafe.Pointer
	SetIdentifiersOperator(value unsafe.Pointer)
	NotificationIdentifiers() string /* primitive/slice/pointer. */
	SetNotificationIdentifiers(value string /* primitive/slice/pointer. */)
	NotificationIdentifiersOperator() unsafe.Pointer
	SetNotificationIdentifiersOperator(value unsafe.Pointer)
	Recipients() INPerson /* already interface */
	SetRecipients(value INPerson /* already interface */)
	RecipientsOperator() unsafe.Pointer
	SetRecipientsOperator(value unsafe.Pointer)
	SearchTerms() string /* primitive/slice/pointer. */
	SetSearchTerms(value string /* primitive/slice/pointer. */)
	SearchTermsOperator() unsafe.Pointer
	SetSearchTermsOperator(value unsafe.Pointer)
	Senders() INPerson /* already interface */
	SetSenders(value INPerson /* already interface */)
	SendersOperator() unsafe.Pointer
	SetSendersOperator(value unsafe.Pointer)
	SpeakableGroupNames() INSpeakableString /* already interface */
	SetSpeakableGroupNames(value INSpeakableString /* already interface */)
	SpeakableGroupNamesOperator() unsafe.Pointer
	SetSpeakableGroupNamesOperator(value unsafe.Pointer)
	// methods:
}

// A request to list the messages that match the specified criteria.
//
// Siri creates objects when the user asks to see sent or received messages. You must implement this intent to support the reading of messages by Siri. This intent object contains the values for you to match when searching the user’s messages. Users can search for messages involving a specific person, messages with specific sent or received dates, or messages containing specific terms. When performing the search, use only the provided parameters, and ignore any that have no values. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler confirms the request and creates an object with the results of the search. For successful searches, Siri offers the user a way to launch your app and see the results. To allow a user wearing AirPods to automatically hear messages, you must implement both and . Add to the options when calling . Finally, add to the category option and to the category intent identifier.


// A request to list the messages that match the specified criteria.
//
// [Full Topic]
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



// The attributes that must be present on a message to yield a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/attributes
func (i_ INSearchForMessagesIntent) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("attributes"))
	return rv
}


// The attributes that must be present on a message to yield a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/attributes
func (i_ INSearchForMessagesIntent) SetAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAttributes:"), value)
}


// The conversation identifiers to locate in your search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/conversationidentifiers
func (i_ INSearchForMessagesIntent) ConversationIdentifiers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("conversationIdentifiers"))
	return rv
}


// The conversation identifiers to locate in your search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/conversationidentifiers
func (i_ INSearchForMessagesIntent) SetConversationIdentifiers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setConversationIdentifiers:"), objc.String(value))
}


// The operator that defines how to use the conversation identifiers in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/conversationidentifiersoperator
func (i_ INSearchForMessagesIntent) ConversationIdentifiersOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("conversationIdentifiersOperator"))
	return rv
}


// The operator that defines how to use the conversation identifiers in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/conversationidentifiersoperator
func (i_ INSearchForMessagesIntent) SetConversationIdentifiersOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setConversationIdentifiersOperator:"), value)
}


// The range of dates in which to search for messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/datetimerange
func (i_ INSearchForMessagesIntent) DateTimeRange() INDateComponentsRange /* already interface */ {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("dateTimeRange"))
	return rv
}


// The range of dates in which to search for messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/datetimerange
func (i_ INSearchForMessagesIntent) SetDateTimeRange(value INDateComponentsRange /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDateTimeRange:"), value)
}


// The names of any groups associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/groupnames
func (i_ INSearchForMessagesIntent) GroupNames() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("groupNames"))
	return rv
}


// The names of any groups associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/groupnames
func (i_ INSearchForMessagesIntent) SetGroupNames(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupNames:"), objc.String(value))
}


// The operator that defines how to use the group names in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/groupnamesoperator
func (i_ INSearchForMessagesIntent) GroupNamesOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("groupNamesOperator"))
	return rv
}


// The operator that defines how to use the group names in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/groupnamesoperator
func (i_ INSearchForMessagesIntent) SetGroupNamesOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupNamesOperator:"), value)
}


// The message identifiers to locate in your search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/identifiers
func (i_ INSearchForMessagesIntent) Identifiers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("identifiers"))
	return rv
}


// The message identifiers to locate in your search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/identifiers
func (i_ INSearchForMessagesIntent) SetIdentifiers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIdentifiers:"), objc.String(value))
}


// The operator that defines how to use the identifiers in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/identifiersoperator
func (i_ INSearchForMessagesIntent) IdentifiersOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("identifiersOperator"))
	return rv
}


// The operator that defines how to use the identifiers in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/identifiersoperator
func (i_ INSearchForMessagesIntent) SetIdentifiersOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIdentifiersOperator:"), value)
}


// The notification identifiers to locate in your search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/notificationidentifiers
func (i_ INSearchForMessagesIntent) NotificationIdentifiers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("notificationIdentifiers"))
	return rv
}


// The notification identifiers to locate in your search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/notificationidentifiers
func (i_ INSearchForMessagesIntent) SetNotificationIdentifiers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNotificationIdentifiers:"), objc.String(value))
}


// The operator that defines how to use the notification identifiers in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/notificationidentifiersoperator
func (i_ INSearchForMessagesIntent) NotificationIdentifiersOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("notificationIdentifiersOperator"))
	return rv
}


// The operator that defines how to use the notification identifiers in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/notificationidentifiersoperator
func (i_ INSearchForMessagesIntent) SetNotificationIdentifiersOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNotificationIdentifiersOperator:"), value)
}


// The contacts who are the recipients of the messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/recipients
func (i_ INSearchForMessagesIntent) Recipients() INPerson /* already interface */ {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("recipients"))
	return rv
}


// The contacts who are the recipients of the messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/recipients
func (i_ INSearchForMessagesIntent) SetRecipients(value INPerson /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRecipients:"), value)
}


// The operator that defines how to use the recipients in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/recipientsoperator
func (i_ INSearchForMessagesIntent) RecipientsOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("recipientsOperator"))
	return rv
}


// The operator that defines how to use the recipients in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/recipientsoperator
func (i_ INSearchForMessagesIntent) SetRecipientsOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRecipientsOperator:"), value)
}


// The terms to look for in the messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/searchterms
func (i_ INSearchForMessagesIntent) SearchTerms() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("searchTerms"))
	return rv
}


// The terms to look for in the messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/searchterms
func (i_ INSearchForMessagesIntent) SetSearchTerms(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchTerms:"), objc.String(value))
}


// The operator that defines how to use the set of terms in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/searchtermsoperator
func (i_ INSearchForMessagesIntent) SearchTermsOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("searchTermsOperator"))
	return rv
}


// The operator that defines how to use the set of terms in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/searchtermsoperator
func (i_ INSearchForMessagesIntent) SetSearchTermsOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchTermsOperator:"), value)
}


// The senders to include or exclude when you search for messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/senders
func (i_ INSearchForMessagesIntent) Senders() INPerson /* already interface */ {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("senders"))
	return rv
}


// The senders to include or exclude when you search for messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/senders
func (i_ INSearchForMessagesIntent) SetSenders(value INPerson /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSenders:"), value)
}


// The operator that defines how to use the senders in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/sendersoperator
func (i_ INSearchForMessagesIntent) SendersOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("sendersOperator"))
	return rv
}


// The operator that defines how to use the senders in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/sendersoperator
func (i_ INSearchForMessagesIntent) SetSendersOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSendersOperator:"), value)
}


// The names of any groups associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/speakablegroupnames
func (i_ INSearchForMessagesIntent) SpeakableGroupNames() INSpeakableString /* already interface */ {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("speakableGroupNames"))
	return rv
}


// The names of any groups associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/speakablegroupnames
func (i_ INSearchForMessagesIntent) SetSpeakableGroupNames(value INSpeakableString /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSpeakableGroupNames:"), value)
}


// The operator that defines how to use the group names in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/speakablegroupnamesoperator
func (i_ INSearchForMessagesIntent) SpeakableGroupNamesOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("speakableGroupNamesOperator"))
	return rv
}


// The operator that defines how to use the group names in the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchformessagesintent/speakablegroupnamesoperator
func (i_ INSearchForMessagesIntent) SetSpeakableGroupNamesOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSpeakableGroupNamesOperator:"), value)
}



