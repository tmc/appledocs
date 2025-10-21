// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchForMessagesIntentResponse] class.
var (
	INSearchForMessagesIntentResponseClass     _INSearchForMessagesIntentResponseClass
	INSearchForMessagesIntentResponseClassOnce sync.Once
)

func getINSearchForMessagesIntentResponseClass() _INSearchForMessagesIntentResponseClass {
	INSearchForMessagesIntentResponseClassOnce.Do(func() {
		INSearchForMessagesIntentResponseClass = _INSearchForMessagesIntentResponseClass{objc.GetClass("INSearchForMessagesIntentResponse")}
	})
	return INSearchForMessagesIntentResponseClass
}

type _INSearchForMessagesIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INSearchForMessagesIntentResponse] class.
type IINSearchForMessagesIntentResponse interface {
	IINIntentResponse
}

// Your app’s response to a search for messages intent.
//
// Use an object to specify the results of searching the user’s messages. After performing a search using the criteria specified in the object, create an instance of this class with the results of that search. Siri communicates the status from your response to the user at appropriate times. You create an object in the and methods of your search for messages handler object. For more information about implementing your handler object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntentResponse
type INSearchForMessagesIntentResponse struct {
	INIntentResponse
}

// INSearchForMessagesIntentResponseFrom constructs a [INSearchForMessagesIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a search for messages intent.
func INSearchForMessagesIntentResponseFrom(ptr unsafe.Pointer) INSearchForMessagesIntentResponse {
	return INSearchForMessagesIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchForMessagesIntentResponseClass) Alloc() INSearchForMessagesIntentResponse {
	rv := objc.Send[INSearchForMessagesIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchForMessagesIntentResponseClass) New() INSearchForMessagesIntentResponse {
	rv := objc.Send[INSearchForMessagesIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchForMessagesIntentResponse) Init() INSearchForMessagesIntentResponse {
	rv := objc.Send[INSearchForMessagesIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchForMessagesIntentResponse) Autorelease() INSearchForMessagesIntentResponse {
	rv := objc.Send[INSearchForMessagesIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchForMessagesIntentResponse creates a new INSearchForMessagesIntentResponse instance.
func NewINSearchForMessagesIntentResponse() INSearchForMessagesIntentResponse {
	return getINSearchForMessagesIntentResponseClass().New()
}




// Initializes the response object with the specified code and user activity object.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntentResponse/init(code:userActivity:)
func NewINSearchForMessagesIntentResponseWithCodeUserActivity(code unsafe.Pointer, userActivity unsafe.Pointer) INSearchForMessagesIntentResponse {
	instance := getINSearchForMessagesIntentResponseClass().Alloc()
	rv := objc.Send[INSearchForMessagesIntentResponse](instance.ID, objc.Sel("initWithCode:userActivity:"), code, userActivity)
	rv.Autorelease()
	return rv
}


// The code indicating whether you successfully handled the intent.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntentResponse/code
func (i_ INSearchForMessagesIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}

// The array of messages matching the search parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntentResponse/messages
func (i_ INSearchForMessagesIntentResponse) Messages() []INMessage {
	rv := objc.Send[[]INMessage](i_.ID, objc.Sel("messages"))
	return rv
}


// SetMessages sets the value of the messages property.
// The array of messages matching the search parameters.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntentResponse/messages
func (i_ INSearchForMessagesIntentResponse) SetMessages(value []INMessage) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](i_.ID, objc.Sel("setMessages:"), nsArray)
}


