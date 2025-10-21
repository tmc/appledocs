// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSendMessageIntentResponse] class.
var (
	INSendMessageIntentResponseClass     _INSendMessageIntentResponseClass
	INSendMessageIntentResponseClassOnce sync.Once
)

func getINSendMessageIntentResponseClass() _INSendMessageIntentResponseClass {
	INSendMessageIntentResponseClassOnce.Do(func() {
		INSendMessageIntentResponseClass = _INSendMessageIntentResponseClass{objc.GetClass("INSendMessageIntentResponse")}
	})
	return INSendMessageIntentResponseClass
}

type _INSendMessageIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INSendMessageIntentResponse] class.
type IINSendMessageIntentResponse interface {
	IINIntentResponse
}

// Your app’s response to a send message intent.
//
// Use an object to specify the results of sending a message to another user. You create instances of this class when confirming or handling a send message intent. Use this object to communicate whether the message was successfully sent or whether an error occurred. You create an object in the and methods of your send message handler object. For more information about implementing your handler object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentResponse
type INSendMessageIntentResponse struct {
	INIntentResponse
}

// INSendMessageIntentResponseFrom constructs a [INSendMessageIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a send message intent.
func INSendMessageIntentResponseFrom(ptr unsafe.Pointer) INSendMessageIntentResponse {
	return INSendMessageIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSendMessageIntentResponseClass) Alloc() INSendMessageIntentResponse {
	rv := objc.Send[INSendMessageIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSendMessageIntentResponseClass) New() INSendMessageIntentResponse {
	rv := objc.Send[INSendMessageIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSendMessageIntentResponse) Init() INSendMessageIntentResponse {
	rv := objc.Send[INSendMessageIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSendMessageIntentResponse) Autorelease() INSendMessageIntentResponse {
	rv := objc.Send[INSendMessageIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSendMessageIntentResponse creates a new INSendMessageIntentResponse instance.
func NewINSendMessageIntentResponse() INSendMessageIntentResponse {
	return getINSendMessageIntentResponseClass().New()
}


// The code indicating whether you successfully handled the intent.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintentresponse/code
func (i_ INSendMessageIntentResponse) Code() INSendMessageIntentResponseCode {
	rv := objc.Send[INSendMessageIntentResponseCode](i_.ID, objc.Sel("code"))
	return rv
}


// SetCode sets the value of the code property.
// The code indicating whether you successfully handled the intent.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintentresponse/code
func (i_ INSendMessageIntentResponse) SetCode(value INSendMessageIntentResponseCode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCode:"), value)
}

// The message sent by the intent.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintentresponse/sentmessage
func (i_ INSendMessageIntentResponse) SentMessage() INMessage {
	rv := objc.Send[INMessage](i_.ID, objc.Sel("sentMessage"))
	return rv
}


// SetSentMessage sets the value of the sentMessage property.
// The message sent by the intent.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintentresponse/sentmessage
func (i_ INSendMessageIntentResponse) SetSentMessage(value INMessage) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSentMessage:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintentresponse/sentmessages
func (i_ INSendMessageIntentResponse) SentMessages() INMessage {
	rv := objc.Send[INMessage](i_.ID, objc.Sel("sentMessages"))
	return rv
}


// SetSentMessages sets the value of the sentMessages property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendmessageintentresponse/sentmessages
func (i_ INSendMessageIntentResponse) SetSentMessages(value INMessage) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSentMessages:"), value)
}



