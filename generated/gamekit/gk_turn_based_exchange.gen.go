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



