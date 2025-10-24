// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKMatchRequest */


/* debug [class_header]: Header for GKMatchRequest */
// The class instance for the [MatchRequest] class.
var (
	MatchRequestClass     _MatchRequestClass
	MatchRequestClassOnce sync.Once
)

func getMatchRequestClass() _MatchRequestClass {
	MatchRequestClassOnce.Do(func() {
		MatchRequestClass = _MatchRequestClass{objc.GetClass("GKMatchRequest")}
	})
	return MatchRequestClass
}

type _MatchRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatchRequest */
// An interface definition for the [MatchRequest] class.
type IMatchRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MatchRequest */
	// properties:
	DefaultNumberOfPlayers() uint
	SetDefaultNumberOfPlayers(value uint)
	InviteeResponseHandler() func(unsafe.Pointer, unsafe.Pointer)
	SetInviteeResponseHandler(value func(unsafe.Pointer, unsafe.Pointer))
	InviteMessage() objc.IObject /* cross-framework: NSString */
	SetInviteMessage(value objc.IObject /* cross-framework: NSString */)
	MaxPlayers() uint
	SetMaxPlayers(value uint)
	MinPlayers() uint
	SetMinPlayers(value uint)
	PlayerAttributes() uint32 /* not a class type */
	SetPlayerAttributes(value uint32 /* not a class type */)
	PlayerGroup() uint
	SetPlayerGroup(value uint)
	PlayersToInvite() []string
	SetPlayersToInvite(value []string)
	Properties() MatchProperties /* not a class type */
	SetProperties(value MatchProperties /* not a class type */)
	QueueName() objc.IObject /* cross-framework: NSString */
	SetQueueName(value objc.IObject /* cross-framework: NSString */)
	RecipientProperties() foundation.IDictionary
	SetRecipientProperties(value foundation.IDictionary)
	RecipientResponseHandler() unsafe.Pointer
	SetRecipientResponseHandler(value unsafe.Pointer)
	Recipients() []Player
	SetRecipients(value []Player)
	RestrictToAutomatch() bool
	SetRestrictToAutomatch(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatchRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatchRequest */
// Alloc allocates a new instance without initialization.
func (mc _MatchRequestClass) Alloc() MatchRequest {
	rv := objc.Send[MatchRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatchRequestClass) New() MatchRequest {
	rv := objc.Send[MatchRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatchRequest) Init() MatchRequest {
	rv := objc.Send[MatchRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatchRequest) Autorelease() MatchRequest {
	rv := objc.Send[MatchRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatchRequest creates a new MatchRequest instance.
func NewMatchRequest() MatchRequest {
	return getMatchRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatchRequest */
// An object that encapsulates the parameters to create a real-time or turn-based match.
//
// To request a match, set the properties of the match request, such as the number of players, the invitation message, and whether to use automatch to fill the player slots. You’re required to set the minimum and maximum number of players allowed in the match. Then, pass the match request to the appropriate class, depending on the type of game and whether you implement your own user interface. To use the matchmaking user interface that GameKit provides, pass the match request to the class for real-time games, or the class for turn-based games. GameKit sends messages to the delegates of these classes when players receive and accept invitations to the match. If you implement your own interface for finding players, pass the match request to the class for real-time games, or the class for turn-based games. If the player selects the other players to invite in your interface, set the , the , and the properties before creating the match.


// An object that encapsulates the parameters to create a real-time or turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest
type MatchRequest struct {
	objectivec.Object
}

// MatchRequestFrom constructs a [MatchRequest] from an unsafe.Pointer.
//
// An object that encapsulates the parameters to create a real-time or turn-based match.
func MatchRequestFrom(ptr unsafe.Pointer) MatchRequest {
	return MatchRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatchRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatchRequest */

// Returns the maximum number of players allowed in the match request for a given match type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/maxPlayersAllowedForMatch(of:)
func (mc _MatchRequestClass) MaxPlayersAllowedForMatchOfType(matchType MatchType) uint {
	rv := objc.Send[uint](objc.ID(mc.class), objc.Sel("maxPlayersAllowedForMatchOfType:"), matchType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MaxPlayersAllowedForMatchOfType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatchRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatchRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatchRequest */

// The default number of players for the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/defaultNumberOfPlayers
func (m_ MatchRequest) DefaultNumberOfPlayers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("defaultNumberOfPlayers"))
	return rv
}/* debug [instance_properties/getter]: defaultNumberOfPlayers */


// The default number of players for the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/defaultNumberOfPlayers
func (m_ MatchRequest) SetDefaultNumberOfPlayers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultNumberOfPlayers:"), value)
}/* debug [instance_properties/setter]: defaultNumberOfPlayers */


// Handles when a player responds to an invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/inviteeResponseHandler
func (m_ MatchRequest) InviteeResponseHandler() func(unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer)](m_.ID, objc.Sel("inviteeResponseHandler"))
	return rv
}/* debug [instance_properties/getter]: inviteeResponseHandler */


// Handles when a player responds to an invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/inviteeResponseHandler
func (m_ MatchRequest) SetInviteeResponseHandler(value func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInviteeResponseHandler:"), value)
}/* debug [instance_properties/setter]: inviteeResponseHandler */


// The message sent to other players when the local player invites them to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/inviteMessage
func (m_ MatchRequest) InviteMessage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("inviteMessage"))
	return rv
}/* debug [instance_properties/getter]: inviteMessage */


// The message sent to other players when the local player invites them to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/inviteMessage
func (m_ MatchRequest) SetInviteMessage(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInviteMessage:"), value)
}/* debug [instance_properties/setter]: inviteMessage */


// The maximum number of players that can join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/maxPlayers
func (m_ MatchRequest) MaxPlayers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxPlayers"))
	return rv
}/* debug [instance_properties/getter]: maxPlayers */


// The maximum number of players that can join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/maxPlayers
func (m_ MatchRequest) SetMaxPlayers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPlayers:"), value)
}/* debug [instance_properties/setter]: maxPlayers */


// The minimum number of players that can join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/minPlayers
func (m_ MatchRequest) MinPlayers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("minPlayers"))
	return rv
}/* debug [instance_properties/getter]: minPlayers */


// The minimum number of players that can join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/minPlayers
func (m_ MatchRequest) SetMinPlayers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPlayers:"), value)
}/* debug [instance_properties/setter]: minPlayers */


// A mask that specifies the role that the local player would like to play in the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/playerAttributes
func (m_ MatchRequest) PlayerAttributes() uint32 /* not a class type */ {
	rv := objc.Send[uint32](m_.ID, objc.Sel("playerAttributes"))
	return rv
}/* debug [instance_properties/getter]: playerAttributes */


// A mask that specifies the role that the local player would like to play in the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/playerAttributes
func (m_ MatchRequest) SetPlayerAttributes(value uint32 /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayerAttributes:"), value)
}/* debug [instance_properties/setter]: playerAttributes */


// A number identifying a subset of players invited to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/playerGroup
func (m_ MatchRequest) PlayerGroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("playerGroup"))
	return rv
}/* debug [instance_properties/getter]: playerGroup */


// A number identifying a subset of players invited to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/playerGroup
func (m_ MatchRequest) SetPlayerGroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayerGroup:"), value)
}/* debug [instance_properties/setter]: playerGroup */


// A list of player identifiers for players to invite to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/playersToInvite
func (m_ MatchRequest) PlayersToInvite() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("playersToInvite"))
	return rv
}/* debug [instance_properties/getter]: playersToInvite */


// A list of player identifiers for players to invite to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/playersToInvite
func (m_ MatchRequest) SetPlayersToInvite(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayersToInvite:"), nsArray)
}/* debug [instance_properties/setter]: playersToInvite */


// The criteria for the local player that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/properties
func (m_ MatchRequest) Properties() MatchProperties /* not a class type */ {
	rv := objc.Send[MatchProperties](m_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */


// The criteria for the local player that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/properties
func (m_ MatchRequest) SetProperties(value MatchProperties /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProperties:"), value)
}/* debug [instance_properties/setter]: properties */


// The name of the queue that Game Center places the match request in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/queueName
func (m_ MatchRequest) QueueName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("queueName"))
	return rv
}/* debug [instance_properties/getter]: queueName */


// The name of the queue that Game Center places the match request in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/queueName
func (m_ MatchRequest) SetQueueName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueName:"), value)
}/* debug [instance_properties/setter]: queueName */


// The criteria for recipients of the match request that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientProperties
func (m_ MatchRequest) RecipientProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("recipientProperties"))
	return rv
}/* debug [instance_properties/getter]: recipientProperties */


// The criteria for recipients of the match request that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientProperties
func (m_ MatchRequest) SetRecipientProperties(value foundation.IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecipientProperties:"), value)
}/* debug [instance_properties/setter]: recipientProperties */


// A method that handles when a player responds to an invitation to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientResponseHandler
func (m_ MatchRequest) RecipientResponseHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("recipientResponseHandler"))
	return rv
}/* debug [instance_properties/getter]: recipientResponseHandler */


// A method that handles when a player responds to an invitation to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientResponseHandler
func (m_ MatchRequest) SetRecipientResponseHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecipientResponseHandler:"), value)
}/* debug [instance_properties/setter]: recipientResponseHandler */


// The players to invite to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipients
func (m_ MatchRequest) Recipients() []Player {
	rv := objc.Send[[]Player](m_.ID, objc.Sel("recipients"))
	return rv
}/* debug [instance_properties/getter]: recipients */


// The players to invite to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipients
func (m_ MatchRequest) SetRecipients(value []Player) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecipients:"), nsArray)
}/* debug [instance_properties/setter]: recipients */


// A Boolean value that determines whether a game uses automatch to find players or the local player invites players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/restrictToAutomatch
func (m_ MatchRequest) RestrictToAutomatch() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("restrictToAutomatch"))
	return rv
}/* debug [instance_properties/getter]: restrictToAutomatch */


// A Boolean value that determines whether a game uses automatch to find players or the local player invites players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/restrictToAutomatch
func (m_ MatchRequest) SetRestrictToAutomatch(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRestrictToAutomatch:"), value)
}/* debug [instance_properties/setter]: restrictToAutomatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKMatchRequest */



