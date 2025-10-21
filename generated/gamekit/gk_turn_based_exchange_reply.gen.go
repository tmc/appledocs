// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [TurnBasedExchangeReply] class.
type ITurnBasedExchangeReply interface {
	objectivec.IObject
}

// Details about a recipient’s response to an exchange request.
//
// When you accept an exchange request using the method, GameKit sends a object to participants using the protocol method. You can also get responses to exchange requests from the object using the parameter.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TurnBasedExchangeReplyClass) Alloc() TurnBasedExchangeReply {
	rv := objc.Send[TurnBasedExchangeReply](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The participant who replies to the exchange request.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeReply/recipient
func (t_ TurnBasedExchangeReply) Recipient() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("recipient"))
	return rv
}



