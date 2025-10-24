// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MatchRequest] class.
type IMatchRequest interface {
	objectivec.IObject
	// properties:
	InviteMessage() objc.IObject /* cross-framework: NSString */
	SetInviteMessage(value objc.IObject /* cross-framework: NSString */)
	RecipientProperties() foundation.IDictionary
	SetRecipientProperties(value foundation.IDictionary)
	RecipientResponseHandler() unsafe.Pointer
	SetRecipientResponseHandler(value unsafe.Pointer)
	DefaultNumberOfPlayers() int
	SetDefaultNumberOfPlayers(value int)
	InviteeResponseHandler() unsafe.Pointer
	SetInviteeResponseHandler(value unsafe.Pointer)
	MaxPlayers() int
	SetMaxPlayers(value int)
	MinPlayers() int
	SetMinPlayers(value int)
	PlayerAttributes() unsafe.Pointer
	SetPlayerAttributes(value unsafe.Pointer)
	PlayerGroup() int
	SetPlayerGroup(value int)
	PlayersToInvite() objc.IObject /* cross-framework: NSString */
	SetPlayersToInvite(value objc.IObject /* cross-framework: NSString */)
	Properties() objc.IObject /* cross-framework: NSString */
	SetProperties(value objc.IObject /* cross-framework: NSString */)
	QueueName() objc.IObject /* cross-framework: NSString */
	SetQueueName(value objc.IObject /* cross-framework: NSString */)
	Recipients() IGKPlayer
	SetRecipients(value IGKPlayer)
	RestrictToAutomatch() bool
	SetRestrictToAutomatch(value bool)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MatchRequestClass) Alloc() MatchRequest {
	rv := objc.Send[MatchRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns the maximum number of players allowed in the match request for a given match type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/maxPlayersAllowedForMatch(of:)
func (mc _MatchRequestClass) MaxPlayersAllowedForMatchOfType(matchType MatchType /* not a class type */) uint {
	rv := objc.Send[uint](objc.ID(mc.class), objc.Sel("maxPlayersAllowedForMatchOfType:"), matchType)
	return rv
}


// The message sent to other players when the local player invites them to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/inviteMessage
func (m_ MatchRequest) InviteMessage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("inviteMessage"))
	return rv
}


// The message sent to other players when the local player invites them to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/inviteMessage
func (m_ MatchRequest) SetInviteMessage(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInviteMessage:"), value)
}


// The criteria for recipients of the match request that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientProperties
func (m_ MatchRequest) RecipientProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("recipientProperties"))
	return rv
}


// The criteria for recipients of the match request that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientProperties
func (m_ MatchRequest) SetRecipientProperties(value foundation.IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecipientProperties:"), value)
}


// A method that handles when a player responds to an invitation to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientResponseHandler
func (m_ MatchRequest) RecipientResponseHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("recipientResponseHandler"))
	return rv
}


// A method that handles when a player responds to an invitation to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientResponseHandler
func (m_ MatchRequest) SetRecipientResponseHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecipientResponseHandler:"), value)
}


// The default number of players for the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/defaultnumberofplayers
func (m_ MatchRequest) DefaultNumberOfPlayers() int {
	rv := objc.Send[int](m_.ID, objc.Sel("defaultNumberOfPlayers"))
	return rv
}


// The default number of players for the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/defaultnumberofplayers
func (m_ MatchRequest) SetDefaultNumberOfPlayers(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultNumberOfPlayers:"), value)
}


// Handles when a player responds to an invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/inviteeresponsehandler
func (m_ MatchRequest) InviteeResponseHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("inviteeResponseHandler"))
	return rv
}


// Handles when a player responds to an invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/inviteeresponsehandler
func (m_ MatchRequest) SetInviteeResponseHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInviteeResponseHandler:"), value)
}


// The maximum number of players that can join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/maxplayers
func (m_ MatchRequest) MaxPlayers() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maxPlayers"))
	return rv
}


// The maximum number of players that can join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/maxplayers
func (m_ MatchRequest) SetMaxPlayers(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPlayers:"), value)
}


// The minimum number of players that can join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/minplayers
func (m_ MatchRequest) MinPlayers() int {
	rv := objc.Send[int](m_.ID, objc.Sel("minPlayers"))
	return rv
}


// The minimum number of players that can join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/minplayers
func (m_ MatchRequest) SetMinPlayers(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPlayers:"), value)
}


// A mask that specifies the role that the local player would like to play in the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playerattributes
func (m_ MatchRequest) PlayerAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("playerAttributes"))
	return rv
}


// A mask that specifies the role that the local player would like to play in the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playerattributes
func (m_ MatchRequest) SetPlayerAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayerAttributes:"), value)
}


// A number identifying a subset of players invited to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playergroup
func (m_ MatchRequest) PlayerGroup() int {
	rv := objc.Send[int](m_.ID, objc.Sel("playerGroup"))
	return rv
}


// A number identifying a subset of players invited to join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playergroup
func (m_ MatchRequest) SetPlayerGroup(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayerGroup:"), value)
}


// A list of player identifiers for players to invite to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playerstoinvite
func (m_ MatchRequest) PlayersToInvite() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("playersToInvite"))
	return rv
}


// A list of player identifiers for players to invite to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playerstoinvite
func (m_ MatchRequest) SetPlayersToInvite(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayersToInvite:"), value)
}


// The criteria for the local player that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/properties
func (m_ MatchRequest) Properties() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("properties"))
	return rv
}


// The criteria for the local player that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/properties
func (m_ MatchRequest) SetProperties(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProperties:"), value)
}


// The name of the queue that Game Center places the match request in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/queuename
func (m_ MatchRequest) QueueName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("queueName"))
	return rv
}


// The name of the queue that Game Center places the match request in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/queuename
func (m_ MatchRequest) SetQueueName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueName:"), value)
}


// The players to invite to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/recipients
func (m_ MatchRequest) Recipients() IGKPlayer {
	rv := objc.Send[Player](m_.ID, objc.Sel("recipients"))
	return rv
}


// The players to invite to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/recipients
func (m_ MatchRequest) SetRecipients(value IGKPlayer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecipients:"), value)
}


// A Boolean value that determines whether a game uses automatch to find players or the local player invites players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/restricttoautomatch
func (m_ MatchRequest) RestrictToAutomatch() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("restrictToAutomatch"))
	return rv
}


// A Boolean value that determines whether a game uses automatch to find players or the local player invites players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/restricttoautomatch
func (m_ MatchRequest) SetRestrictToAutomatch(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRestrictToAutomatch:"), value)
}



