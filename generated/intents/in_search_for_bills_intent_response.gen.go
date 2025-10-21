// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchForBillsIntentResponse] class.
var (
	INSearchForBillsIntentResponseClass     _INSearchForBillsIntentResponseClass
	INSearchForBillsIntentResponseClassOnce sync.Once
)

func getINSearchForBillsIntentResponseClass() _INSearchForBillsIntentResponseClass {
	INSearchForBillsIntentResponseClassOnce.Do(func() {
		INSearchForBillsIntentResponseClass = _INSearchForBillsIntentResponseClass{objc.GetClass("INSearchForBillsIntentResponse")}
	})
	return INSearchForBillsIntentResponseClass
}

type _INSearchForBillsIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INSearchForBillsIntentResponse] class.
type IINSearchForBillsIntentResponse interface {
	IINIntentResponse
}

// Your app’s response to a request to a search for bills.
//
// Use an object to return the list of bills found during a search operation. After performing a search using the criteria specified in an object, create an instance of this class and fill it with the results of that search. Siri communicates the information from your response to the user at appropriate times. You create an object in the and methods of your search for bills handler object. For more information about implementing your handler object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForBillsIntentResponse
type INSearchForBillsIntentResponse struct {
	INIntentResponse
}

// INSearchForBillsIntentResponseFrom constructs a [INSearchForBillsIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a request to a search for bills.
func INSearchForBillsIntentResponseFrom(ptr unsafe.Pointer) INSearchForBillsIntentResponse {
	return INSearchForBillsIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchForBillsIntentResponseClass) Alloc() INSearchForBillsIntentResponse {
	rv := objc.Send[INSearchForBillsIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchForBillsIntentResponseClass) New() INSearchForBillsIntentResponse {
	rv := objc.Send[INSearchForBillsIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchForBillsIntentResponse) Init() INSearchForBillsIntentResponse {
	rv := objc.Send[INSearchForBillsIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchForBillsIntentResponse) Autorelease() INSearchForBillsIntentResponse {
	rv := objc.Send[INSearchForBillsIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchForBillsIntentResponse creates a new INSearchForBillsIntentResponse instance.
func NewINSearchForBillsIntentResponse() INSearchForBillsIntentResponse {
	return getINSearchForBillsIntentResponseClass().New()
}


// The bills found during the search.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintentresponse/bills
func (i_ INSearchForBillsIntentResponse) Bills() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("bills"))
	return rv
}


// SetBills sets the value of the bills property.
// The bills found during the search.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintentresponse/bills
func (i_ INSearchForBillsIntentResponse) SetBills(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBills:"), value)
}

// The code indicating whether you successfully handled the intent.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintentresponse/code
func (i_ INSearchForBillsIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}


// SetCode sets the value of the code property.
// The code indicating whether you successfully handled the intent.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintentresponse/code
func (i_ INSearchForBillsIntentResponse) SetCode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCode:"), value)
}



