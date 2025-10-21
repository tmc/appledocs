// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TKSmartCard] class.
var (
	TKSmartCardClass     _TKSmartCardClass
	TKSmartCardClassOnce sync.Once
)

func getTKSmartCardClass() _TKSmartCardClass {
	TKSmartCardClassOnce.Do(func() {
		TKSmartCardClass = _TKSmartCardClass{objc.GetClass("TKSmartCard")}
	})
	return TKSmartCardClass
}

type _TKSmartCardClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCard] class.
type ITKSmartCard interface {
	objectivec.IObject
	BeginSessionWithReply(reply unsafe.Pointer)
	EndSession()
}

// A representation of a smart card.
//
// This class provides an interface for managing sessions with a smart card, transmitting requests, and facilitating user interaction. You can create a object when a smart card is inserted into a slot, by calling the method on the corresponding object. To start communicating with the smart card, call the method on the object. Once an exclusive session has been established, you transmit data using the method. After you’ve finished communicating with a smart card, you call the method. If the smart card is physically removed from its slot, the session object becomes invalid, and any further calls to will return an error. You can use Key-Value Observing on the property to be notified when a smart card is invalidated, due to being removed from the slot or another reason.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard
type TKSmartCard struct {
	objectivec.Object
}

// TKSmartCardFrom constructs a [TKSmartCard] from an unsafe.Pointer.
//
// A representation of a smart card.
func TKSmartCardFrom(ptr unsafe.Pointer) TKSmartCard {
	return TKSmartCard{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardClass) Alloc() TKSmartCard {
	rv := objc.Send[TKSmartCard](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardClass) New() TKSmartCard {
	rv := objc.Send[TKSmartCard](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCard) Init() TKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCard) Autorelease() TKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCard creates a new TKSmartCard instance.
func NewTKSmartCard() TKSmartCard {
	return getTKSmartCardClass().New()
}


// Begins a session with the Smart Card.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/beginSession(reply:)
func (t_ TKSmartCard) BeginSessionWithReply(reply unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("beginSessionWithReply:"), reply)
}

// Completes any pending transmissions and ends the session to the Smart Card.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/endSession()
func (t_ TKSmartCard) EndSession() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endSession"))
}

// User-specified information. This property is automatically set to if the Smart Card is removed or another object begins a session.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/context
func (t_ TKSmartCard) Context() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("context"))
	return rv
}


// SetContext sets the value of the context property.
// User-specified information. This property is automatically set to if the Smart Card is removed or another object begins a session.

//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/context
func (t_ TKSmartCard) SetContext(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContext:"), value)
}


