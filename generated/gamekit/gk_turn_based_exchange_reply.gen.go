// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKTurnBasedExchangeReply */


/* debug [class_header]: Header for GKTurnBasedExchangeReply */
// The class instance for the [TurnBasedExchangeReply] class.
var (
	TurnBasedExchangeReplyClass     _TurnBasedExchangeReplyClass
	TurnBasedExchangeReplyClassOnce sync.Once
)

func getTurnBasedExchangeReplyClass() _TurnBasedExchangeReplyClass {
	TurnBasedExchangeReplyClassOnce.Do(func() {
		TurnBasedExchangeReplyClass = _TurnBasedExchangeReplyClass{objc.GetClass("GKTurnBasedExchangeReply")}
	})
	return TurnBasedExchangeReplyClass
}

type _TurnBasedExchangeReplyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TurnBasedExchangeReply */
// An interface definition for the [TurnBasedExchangeReply] class.
type ITurnBasedExchangeReply interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TurnBasedExchangeReply */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	Message() objc.IObject /* cross-framework: NSString */
	Recipient() IGKTurnBasedParticipant
	ReplyDate() objc.IObject /* cross-framework: NSDate */
	Replies() IGKTurnBasedExchangeReply
	SetReplies(value IGKTurnBasedExchangeReply)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TurnBasedExchangeReply */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TurnBasedExchangeReply */
// Alloc allocates a new instance without initialization.
func (tc _TurnBasedExchangeReplyClass) Alloc() TurnBasedExchangeReply {
	rv := objc.Send[TurnBasedExchangeReply](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TurnBasedExchangeReplyClass) New() TurnBasedExchangeReply {
	rv := objc.Send[TurnBasedExchangeReply](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TurnBasedExchangeReply) Init() TurnBasedExchangeReply {
	rv := objc.Send[TurnBasedExchangeReply](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TurnBasedExchangeReply) Autorelease() TurnBasedExchangeReply {
	rv := objc.Send[TurnBasedExchangeReply](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTurnBasedExchangeReply creates a new TurnBasedExchangeReply instance.
func NewTurnBasedExchangeReply() TurnBasedExchangeReply {
	return getTurnBasedExchangeReplyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TurnBasedExchangeReply */
// Details about a recipient’s response to an exchange request.
//
// When you accept an exchange request using the method, GameKit sends a object to participants using the protocol method. You can also get responses to exchange requests from the object using the parameter.


// Details about a recipient’s response to an exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeReply
type TurnBasedExchangeReply struct {
	objectivec.Object
}

// TurnBasedExchangeReplyFrom constructs a [TurnBasedExchangeReply] from an unsafe.Pointer.
//
// Details about a recipient’s response to an exchange request.
func TurnBasedExchangeReplyFrom(ptr unsafe.Pointer) TurnBasedExchangeReply {
	return TurnBasedExchangeReply{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TurnBasedExchangeReply *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TurnBasedExchangeReply */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TurnBasedExchangeReply */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TurnBasedExchangeReply */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TurnBasedExchangeReply */

// The game-specific data that the recipent provides in the exchange request reply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeReply/data
func (t_ TurnBasedExchangeReply) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// A message from the recipient to the sender of the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeReply/message
func (t_ TurnBasedExchangeReply) Message() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("message"))
	return rv
}/* debug [instance_properties/getter]: message */


// The participant who replies to the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeReply/recipient
func (t_ TurnBasedExchangeReply) Recipient() IGKTurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](t_.ID, objc.Sel("recipient"))
	return rv
}/* debug [instance_properties/getter]: recipient */


// The date the recipient replies to the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeReply/replyDate
func (t_ TurnBasedExchangeReply) ReplyDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("replyDate"))
	return rv
}/* debug [instance_properties/getter]: replyDate */


// The replies from recipients of the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/replies
func (t_ TurnBasedExchangeReply) Replies() IGKTurnBasedExchangeReply {
	rv := objc.Send[TurnBasedExchangeReply](t_.ID, objc.Sel("replies"))
	return rv
}/* debug [instance_properties/getter]: replies */


// The replies from recipients of the exchange request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedexchange/replies
func (t_ TurnBasedExchangeReply) SetReplies(value IGKTurnBasedExchangeReply) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReplies:"), value)
}/* debug [instance_properties/setter]: replies */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKTurnBasedExchangeReply */



