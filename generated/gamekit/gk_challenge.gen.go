// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKChallenge */


/* debug [class_header]: Header for GKChallenge */
// The class instance for the [Challenge] class.
var (
	ChallengeClass     _ChallengeClass
	ChallengeClassOnce sync.Once
)

func getChallengeClass() _ChallengeClass {
	ChallengeClassOnce.Do(func() {
		ChallengeClass = _ChallengeClass{objc.GetClass("GKChallenge")}
	})
	return ChallengeClass
}

type _ChallengeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Challenge */
// An interface definition for the [Challenge] class.
type IChallenge interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Challenge */
	// properties:
	CompletionDate() objc.IObject /* cross-framework: NSDate */
	IssueDate() objc.IObject /* cross-framework: NSDate */
	IssuingPlayer() IGKPlayer
	IssuingPlayerID() objc.IObject /* cross-framework: NSString */
	Message() objc.IObject /* cross-framework: NSString */
	ReceivingPlayer() IGKPlayer
	ReceivingPlayerID() objc.IObject /* cross-framework: NSString */
	State() ChallengeState
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Challenge */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Challenge */
// Alloc allocates a new instance without initialization.
func (cc _ChallengeClass) Alloc() Challenge {
	rv := objc.Send[Challenge](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChallengeClass) New() Challenge {
	rv := objc.Send[Challenge](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Challenge) Init() Challenge {
	rv := objc.Send[Challenge](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Challenge) Autorelease() Challenge {
	rv := objc.Send[Challenge](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChallenge creates a new Challenge instance.
func NewChallenge() Challenge {
	return getChallengeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Challenge */
// A challenge issued by the local player to another player.
//
// Players can use Game Center to challenge other players to beat their scores and achievements that they earn in your game. When a player issues a challenge to another player, Game Center sends a push notification to the other player. That player can then accept or refuse the challenge. If the player accepts the challenge, Game Center adds the challenge to the player’s list of challenges. Later, if the player beats the challenge, Game Center notifies both players. Game Center supports two kinds of challenges: You never subclass the class directly. However, you can subclass or to create specific kinds of challenges.


// A challenge issued by the local player to another player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge
type Challenge struct {
	objectivec.Object
}

// ChallengeFrom constructs a [Challenge] from an unsafe.Pointer.
//
// A challenge issued by the local player to another player.
func ChallengeFrom(ptr unsafe.Pointer) Challenge {
	return Challenge{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Challenge *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Challenge */

// Loads the list of outstanding challenges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge/loadReceivedChallenges(completionHandler:)
func (cc _ChallengeClass) LoadReceivedChallengesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadReceivedChallengesWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadReceivedChallengesWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Challenge */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Challenge */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Challenge */

// The date the challenged player completed the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge/completionDate
func (c_ Challenge) CompletionDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("completionDate"))
	return rv
}/* debug [instance_properties/getter]: completionDate */


// The date the player issued the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge/issueDate
func (c_ Challenge) IssueDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("issueDate"))
	return rv
}/* debug [instance_properties/getter]: issueDate */


// The player who issues the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge/issuingPlayer
func (c_ Challenge) IssuingPlayer() IGKPlayer {
	rv := objc.Send[Player](c_.ID, objc.Sel("issuingPlayer"))
	return rv
}/* debug [instance_properties/getter]: issuingPlayer */


// The player who issues the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge/issuingPlayerID
func (c_ Challenge) IssuingPlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("issuingPlayerID"))
	return rv
}/* debug [instance_properties/getter]: issuingPlayerID */


// A text message that describes the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge/message
func (c_ Challenge) Message() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("message"))
	return rv
}/* debug [instance_properties/getter]: message */


// The player who receives the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge/receivingPlayer
func (c_ Challenge) ReceivingPlayer() IGKPlayer {
	rv := objc.Send[Player](c_.ID, objc.Sel("receivingPlayer"))
	return rv
}/* debug [instance_properties/getter]: receivingPlayer */


// The player who receives the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge/receivingPlayerID
func (c_ Challenge) ReceivingPlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("receivingPlayerID"))
	return rv
}/* debug [instance_properties/getter]: receivingPlayerID */


// The current state of the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallenge/state
func (c_ Challenge) State() ChallengeState {
	rv := objc.Send[ChallengeState](c_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (c_ Challenge) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (c_ Challenge) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKChallenge */



