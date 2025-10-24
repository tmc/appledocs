// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKTurnBasedExchange */


/* debug [class_header]: Header for GKTurnBasedExchange */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TurnBasedExchange */
// An interface definition for the [TurnBasedExchange] class.
type ITurnBasedExchange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TurnBasedExchange */
	// properties:
	CompletionDate() objc.IObject /* cross-framework: NSDate */
	Data() objc.IObject /* cross-framework: NSData */
	ExchangeID() objc.IObject /* cross-framework: NSString */
	Message() objc.IObject /* cross-framework: NSString */
	Recipients() []TurnBasedParticipant
	Replies() []TurnBasedExchangeReply
	SendDate() objc.IObject /* cross-framework: NSDate */
	Sender() IGKTurnBasedParticipant
	Status() TurnBasedExchangeStatus
	TimeoutDate() objc.IObject /* cross-framework: NSDate */
	CompletedExchanges() IGKTurnBasedExchange
	SetCompletedExchanges(value IGKTurnBasedExchange)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TurnBasedExchange */
	// methods:
	CancelWithLocalizableMessageKeyArgumentsCompletionHandler(key objc.IObject /* cross-framework: NSString */, arguments []string, completionHandler unsafe.Pointer)
	ReplyWithLocalizableMessageKeyArgumentsDataCompletionHandler(key objc.IObject /* cross-framework: NSString */, arguments []string, data objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TurnBasedExchange */
// Alloc allocates a new instance without initialization.
func (tc _TurnBasedExchangeClass) Alloc() TurnBasedExchange {
	rv := objc.Send[TurnBasedExchange](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TurnBasedExchange */
// Exchange request information that participants send in a turn-based match.
//
// GameKit sends exchange objects to protocol methods when the local player receives an exchange request or recipients reply to an exchange request. The exchange object encapsulates your custom game data that you want to communicate to other players. You initiate an exchange request using the method. Then GameKit sends the request to the recipients passing the exchange object to the protocol method. GameKit sets the status of the exchange object to . After all recipients respond to the request, using the method, or exceed the time out specified in the request, GameKit sends the exchange to the sender and the current participant. GameKit sets the exchange status to and then passes it to the method. Before the current participant ends their turn, save the completed exchanges using the method. Get the exchanges from the match object using the property. Alternatively, save exchange data in the protocol method when all recipients reply to specific exchange requests. To cancel an active or complete exchange, use the method. GameKit notifies the recipients when the player cancels an exchange, using the protocol method.


// Exchange request information that participants send in a turn-based match.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TurnBasedExchange *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TurnBasedExchange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TurnBasedExchange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TurnBasedExchange */

// Cancels an exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/cancel(withLocalizableMessageKey:arguments:completionHandler:)
func (t_ TurnBasedExchange) CancelWithLocalizableMessageKeyArgumentsCompletionHandler(key objc.IObject /* cross-framework: NSString */, arguments []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("cancelWithLocalizableMessageKey:arguments:completionHandler:"), key, arguments, completionHandler)
}/* debug [instance_methods/method]: CancelWithLocalizableMessageKeyArgumentsCompletionHandler */


// Replies to an exchange request on behalf of a recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/reply(withLocalizableMessageKey:arguments:data:completionHandler:)
func (t_ TurnBasedExchange) ReplyWithLocalizableMessageKeyArgumentsDataCompletionHandler(key objc.IObject /* cross-framework: NSString */, arguments []string, data objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replyWithLocalizableMessageKey:arguments:data:completionHandler:"), key, arguments, data, completionHandler)
}/* debug [instance_methods/method]: ReplyWithLocalizableMessageKeyArgumentsDataCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TurnBasedExchange */

// The date when all recipients of the exchange request reply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/completionDate
func (t_ TurnBasedExchange) CompletionDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("completionDate"))
	return rv
}/* debug [instance_properties/getter]: completionDate */


// The game-specific exchange data that GameKit sends to participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/data
func (t_ TurnBasedExchange) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The identifier for the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/exchangeID
func (t_ TurnBasedExchange) ExchangeID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("exchangeID"))
	return rv
}/* debug [instance_properties/getter]: exchangeID */


// A localized message from the sender to the recipients of an exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/message
func (t_ TurnBasedExchange) Message() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("message"))
	return rv
}/* debug [instance_properties/getter]: message */


// The participants who receives the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/recipients
func (t_ TurnBasedExchange) Recipients() []TurnBasedParticipant {
	rv := objc.Send[[]TurnBasedParticipant](t_.ID, objc.Sel("recipients"))
	return rv
}/* debug [instance_properties/getter]: recipients */


// The replies from recipients of the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/replies
func (t_ TurnBasedExchange) Replies() []TurnBasedExchangeReply {
	rv := objc.Send[[]TurnBasedExchangeReply](t_.ID, objc.Sel("replies"))
	return rv
}/* debug [instance_properties/getter]: replies */


// The date that the sender initiates the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/sendDate
func (t_ TurnBasedExchange) SendDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("sendDate"))
	return rv
}/* debug [instance_properties/getter]: sendDate */


// The participant who sends the exchange request to recipients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/sender
func (t_ TurnBasedExchange) Sender() IGKTurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](t_.ID, objc.Sel("sender"))
	return rv
}/* debug [instance_properties/getter]: sender */


// The status of the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/status
func (t_ TurnBasedExchange) Status() TurnBasedExchangeStatus {
	rv := objc.Send[TurnBasedExchangeStatus](t_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// The date that the recipients must reply by before the exchange request times out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchange/timeoutDate
func (t_ TurnBasedExchange) TimeoutDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("timeoutDate"))
	return rv
}/* debug [instance_properties/getter]: timeoutDate */


// The exchange requests that all recipients replied to and the current participant needs to save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/completedexchanges
func (t_ TurnBasedExchange) CompletedExchanges() IGKTurnBasedExchange {
	rv := objc.Send[TurnBasedExchange](t_.ID, objc.Sel("completedExchanges"))
	return rv
}/* debug [instance_properties/getter]: completedExchanges */


// The exchange requests that all recipients replied to and the current participant needs to save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/completedexchanges
func (t_ TurnBasedExchange) SetCompletedExchanges(value IGKTurnBasedExchange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompletedExchanges:"), value)
}/* debug [instance_properties/setter]: completedExchanges */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKTurnBasedExchange */



