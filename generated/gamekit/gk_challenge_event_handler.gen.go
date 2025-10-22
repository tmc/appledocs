// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ChallengeEventHandler] class.
var (
	ChallengeEventHandlerClass     _ChallengeEventHandlerClass
	ChallengeEventHandlerClassOnce sync.Once
)

func getChallengeEventHandlerClass() _ChallengeEventHandlerClass {
	ChallengeEventHandlerClassOnce.Do(func() {
		ChallengeEventHandlerClass = _ChallengeEventHandlerClass{objc.GetClass("GKChallengeEventHandler")}
	})
	return ChallengeEventHandlerClass
}

type _ChallengeEventHandlerClass struct {
	class objc.Class
}

// An interface definition for the [ChallengeEventHandler] class.
type IChallengeEventHandler interface {
	objectivec.IObject
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
}

// The class is used to respond to events related to challenges sent or received by the local player.
//
// To use it, call the class method to get the instance and assign an object that implements the protocol to its property. You should assign a challenge event handler immediately after initializing the local player, because your game may have launched in response to a challenge notification being received by the player.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeEventHandler
type ChallengeEventHandler struct {
	objectivec.Object
}

// ChallengeEventHandlerFrom constructs a [ChallengeEventHandler] from an unsafe.Pointer.
//
// The class is used to respond to events related to challenges sent or received by the local player.
func ChallengeEventHandlerFrom(ptr unsafe.Pointer) ChallengeEventHandler {
	return ChallengeEventHandler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ChallengeEventHandlerClass) Alloc() ChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChallengeEventHandlerClass) New() ChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChallengeEventHandler) Init() ChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChallengeEventHandler) Autorelease() ChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChallengeEventHandler creates a new ChallengeEventHandler instance.
func NewChallengeEventHandler() ChallengeEventHandler {
	return getChallengeEventHandlerClass().New()
}


// Returns the shared instance of the event handler
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeEventHandler/challengeEventHandler
func (cc _ChallengeEventHandlerClass) ChallengeEventHandler() ChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](objc.ID(cc.class), objc.Sel("challengeEventHandler"))
	return rv
}

// The delegate for the event handler.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengeeventhandler/delegate
func (c_ ChallengeEventHandler) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for the event handler.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengeeventhandler/delegate
func (c_ ChallengeEventHandler) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}



