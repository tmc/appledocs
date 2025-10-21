// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TurnBasedExchange] class.
var (
	TurnBasedExchangeClass     _TurnBasedExchangeClass
	TurnBasedExchangeClassOnce sync.Once
)

func getTurnBasedExchangeClass() _TurnBasedExchangeClass {
	TurnBasedExchangeClassOnce.Do(func() {
		TurnBasedExchangeClass = _TurnBasedExchangeClass{objc.GetClass("GKTurnBasedExchange")}
	})
	return TurnBasedExchangeClass
}

type _TurnBasedExchangeClass struct {
	class objc.Class
}

// An interface definition for the [TurnBasedExchange] class.
type ITurnBasedExchange interface {
	objectivec.IObject
}

// Exchange request information that participants send in a turn-based match.
//
// GameKit sends exchange objects to protocol methods when the local player receives an exchange request or recipients reply to an exchange request. The exchange object encapsulates your custom game data that you want to communicate to other players. You initiate an exchange request using the method. Then GameKit sends the request to the recipients passing the exchange object to the protocol method. GameKit sets the status of the exchange object to . After all recipients respond to the request, using the method, or exceed the time out specified in the request, GameKit sends the exchange to the sender and the current participant. GameKit sets the exchange status to and then passes it to the method. Before the current participant ends their turn, save the completed exchanges using the method. Get the exchanges from the match object using the property. Alternatively, save exchange data in the protocol method when all recipients reply to specific exchange requests. To cancel an active or complete exchange, use the method. GameKit notifies the recipients when the player cancels an exchange, using the protocol method.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange
type TurnBasedExchange struct {
	objectivec.Object
}

// TurnBasedExchangeFrom constructs a [TurnBasedExchange] from an unsafe.Pointer.
//
// Exchange request information that participants send in a turn-based match.
func TurnBasedExchangeFrom(ptr unsafe.Pointer) TurnBasedExchange {
	return TurnBasedExchange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TurnBasedExchangeClass) Alloc() TurnBasedExchange {
	rv := objc.Send[TurnBasedExchange](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TurnBasedExchangeClass) New() TurnBasedExchange {
	rv := objc.Send[TurnBasedExchange](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TurnBasedExchange) Init() TurnBasedExchange {
	rv := objc.Send[TurnBasedExchange](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TurnBasedExchange) Autorelease() TurnBasedExchange {
	rv := objc.Send[TurnBasedExchange](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTurnBasedExchange creates a new TurnBasedExchange instance.
func NewTurnBasedExchange() TurnBasedExchange {
	return getTurnBasedExchangeClass().New()
}


// The identifier for the exchange request.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/exchangeID
func (t_ TurnBasedExchange) ExchangeID() string {
	rv := objc.Send[string](t_.ID, objc.Sel("exchangeID"))
	return rv
}

// The date when all recipients of the exchange request reply.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/completiondate
func (t_ TurnBasedExchange) CompletionDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("completionDate"))
	return rv
}


// SetCompletionDate sets the value of the completionDate property.
// The date when all recipients of the exchange request reply.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/completiondate
func (t_ TurnBasedExchange) SetCompletionDate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompletionDate:"), value)
}

// The game-specific exchange data that GameKit sends to participants.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/data
func (t_ TurnBasedExchange) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
// The game-specific exchange data that GameKit sends to participants.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/data
func (t_ TurnBasedExchange) SetData(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setData:"), value)
}

// A localized message from the sender to the recipients of an exchange request.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/message
func (t_ TurnBasedExchange) Message() string {
	rv := objc.Send[string](t_.ID, objc.Sel("message"))
	return rv
}


// SetMessage sets the value of the message property.
// A localized message from the sender to the recipients of an exchange request.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/message
func (t_ TurnBasedExchange) SetMessage(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMessage:"), objc.String(value))
}

// The participants who receives the exchange request.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/recipients
func (t_ TurnBasedExchange) Recipients() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("recipients"))
	return rv
}


// SetRecipients sets the value of the recipients property.
// The participants who receives the exchange request.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/recipients
func (t_ TurnBasedExchange) SetRecipients(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRecipients:"), value)
}

// The replies from recipients of the exchange request.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/replies
func (t_ TurnBasedExchange) Replies() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("replies"))
	return rv
}


// SetReplies sets the value of the replies property.
// The replies from recipients of the exchange request.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/replies
func (t_ TurnBasedExchange) SetReplies(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReplies:"), value)
}

// The date that the sender initiates the exchange request.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/senddate
func (t_ TurnBasedExchange) SendDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("sendDate"))
	return rv
}


// SetSendDate sets the value of the sendDate property.
// The date that the sender initiates the exchange request.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/senddate
func (t_ TurnBasedExchange) SetSendDate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSendDate:"), value)
}

// The participant who sends the exchange request to recipients.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/sender
func (t_ TurnBasedExchange) Sender() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("sender"))
	return rv
}


// SetSender sets the value of the sender property.
// The participant who sends the exchange request to recipients.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/sender
func (t_ TurnBasedExchange) SetSender(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSender:"), value)
}

// The status of the exchange request.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/status
func (t_ TurnBasedExchange) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
// The status of the exchange request.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/status
func (t_ TurnBasedExchange) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStatus:"), value)
}

// The date that the recipients must reply by before the exchange request times out.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/timeoutdate
func (t_ TurnBasedExchange) TimeoutDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("timeoutDate"))
	return rv
}


// SetTimeoutDate sets the value of the timeoutDate property.
// The date that the recipients must reply by before the exchange request times out.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/timeoutdate
func (t_ TurnBasedExchange) SetTimeoutDate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTimeoutDate:"), value)
}

// The exchange requests that all recipients replied to and the current participant needs to save.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/completedexchanges
func (t_ TurnBasedExchange) CompletedExchanges() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("completedExchanges"))
	return rv
}


// SetCompletedExchanges sets the value of the completedExchanges property.
// The exchange requests that all recipients replied to and the current participant needs to save.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/completedexchanges
func (t_ TurnBasedExchange) SetCompletedExchanges(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompletedExchanges:"), value)
}



