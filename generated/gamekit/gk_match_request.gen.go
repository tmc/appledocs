// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that encapsulates the parameters to create a real-time or turn-based match.
//
// To request a match, set the properties of the match request, such as the number of players, the invitation message, and whether to use automatch to fill the player slots. You’re required to set the minimum and maximum number of players allowed in the match. Then, pass the match request to the appropriate class, depending on the type of game and whether you implement your own user interface. To use the matchmaking user interface that GameKit provides, pass the match request to the class for real-time games, or the class for turn-based games. GameKit sends messages to the delegates of these classes when players receive and accept invitations to the match. If you implement your own interface for finding players, pass the match request to the class for real-time games, or the class for turn-based games. If the player selects the other players to invite in your interface, set the , the , and the properties before creating the match.
//
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
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/maxPlayersAllowedForMatch(of:)
func (mc _MatchRequestClass) MaxPlayersAllowedForMatchOfType(matchType unsafe.Pointer) uint {
	rv := objc.Send[uint](objc.ID(mc.class), objc.Sel("maxPlayersAllowedForMatchOfType:"), matchType)
	return rv
}

// A number identifying a subset of players invited to join a match.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playergroup
func (m_ MatchRequest) PlayerGroup() int {
	rv := objc.Send[int](m_.ID, objc.Sel("playerGroup"))
	return rv
}


// SetPlayerGroup sets the value of the playerGroup property.
// A number identifying a subset of players invited to join a match.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playergroup
func (m_ MatchRequest) SetPlayerGroup(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayerGroup:"), value)
}

// A list of player identifiers for players to invite to the match.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playerstoinvite
func (m_ MatchRequest) PlayersToInvite() string {
	rv := objc.Send[string](m_.ID, objc.Sel("playersToInvite"))
	return rv
}


// SetPlayersToInvite sets the value of the playersToInvite property.
// A list of player identifiers for players to invite to the match.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playerstoinvite
func (m_ MatchRequest) SetPlayersToInvite(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayersToInvite:"), objc.String(value))
}

// The name of the queue that Game Center places the match request in.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/queuename
func (m_ MatchRequest) QueueName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("queueName"))
	return rv
}


// SetQueueName sets the value of the queueName property.
// The name of the queue that Game Center places the match request in.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/queuename
func (m_ MatchRequest) SetQueueName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueName:"), objc.String(value))
}

// A mask that specifies the role that the local player would like to play in the game.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playerattributes
func (m_ MatchRequest) PlayerAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("playerAttributes"))
	return rv
}


// SetPlayerAttributes sets the value of the playerAttributes property.
// A mask that specifies the role that the local player would like to play in the game.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/playerattributes
func (m_ MatchRequest) SetPlayerAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayerAttributes:"), value)
}

// The players to invite to the match.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/recipients
func (m_ MatchRequest) Recipients() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("recipients"))
	return rv
}


// SetRecipients sets the value of the recipients property.
// The players to invite to the match.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/recipients
func (m_ MatchRequest) SetRecipients(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecipients:"), value)
}

// The criteria for the local player that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/properties
func (m_ MatchRequest) Properties() string {
	rv := objc.Send[string](m_.ID, objc.Sel("properties"))
	return rv
}


// SetProperties sets the value of the properties property.
// The criteria for the local player that Game Center uses to find other players when using matchmaking rules.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/properties
func (m_ MatchRequest) SetProperties(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProperties:"), objc.String(value))
}

// Handles when a player responds to an invitation.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/inviteeresponsehandler
func (m_ MatchRequest) InviteeResponseHandler() string {
	rv := objc.Send[string](m_.ID, objc.Sel("inviteeResponseHandler"))
	return rv
}


// SetInviteeResponseHandler sets the value of the inviteeResponseHandler property.
// Handles when a player responds to an invitation.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/inviteeresponsehandler
func (m_ MatchRequest) SetInviteeResponseHandler(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInviteeResponseHandler:"), objc.String(value))
}

// A Boolean value that determines whether a game uses automatch to find players or the local player invites players.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/restricttoautomatch
func (m_ MatchRequest) RestrictToAutomatch() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("restrictToAutomatch"))
	return rv
}


// SetRestrictToAutomatch sets the value of the restrictToAutomatch property.
// A Boolean value that determines whether a game uses automatch to find players or the local player invites players.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchrequest/restricttoautomatch
func (m_ MatchRequest) SetRestrictToAutomatch(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRestrictToAutomatch:"), value)
}

// The default number of players for the match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/defaultNumberOfPlayers
func (m_ MatchRequest) DefaultNumberOfPlayers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("defaultNumberOfPlayers"))
	return rv
}


// SetDefaultNumberOfPlayers sets the value of the defaultNumberOfPlayers property.
// The default number of players for the match.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/defaultNumberOfPlayers
func (m_ MatchRequest) SetDefaultNumberOfPlayers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultNumberOfPlayers:"), value)
}

// The message sent to other players when the local player invites them to join a match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/inviteMessage
func (m_ MatchRequest) InviteMessage() string {
	rv := objc.Send[string](m_.ID, objc.Sel("inviteMessage"))
	return rv
}


// SetInviteMessage sets the value of the inviteMessage property.
// The message sent to other players when the local player invites them to join a match.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/inviteMessage
func (m_ MatchRequest) SetInviteMessage(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInviteMessage:"), objc.String(value))
}

// The maximum number of players that can join the match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/maxPlayers
func (m_ MatchRequest) MaxPlayers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxPlayers"))
	return rv
}


// SetMaxPlayers sets the value of the maxPlayers property.
// The maximum number of players that can join the match.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/maxPlayers
func (m_ MatchRequest) SetMaxPlayers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPlayers:"), value)
}

// The minimum number of players that can join the match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/minPlayers
func (m_ MatchRequest) MinPlayers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("minPlayers"))
	return rv
}


// SetMinPlayers sets the value of the minPlayers property.
// The minimum number of players that can join the match.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/minPlayers
func (m_ MatchRequest) SetMinPlayers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPlayers:"), value)
}

// The criteria for recipients of the match request that Game Center uses to find other players when using matchmaking rules.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientProperties
func (m_ MatchRequest) RecipientProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("recipientProperties"))
	return rv
}


// SetRecipientProperties sets the value of the recipientProperties property.
// The criteria for recipients of the match request that Game Center uses to find other players when using matchmaking rules.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchRequest/recipientProperties
func (m_ MatchRequest) SetRecipientProperties(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecipientProperties:"), value)
}



