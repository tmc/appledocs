// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKMatchmaker */


/* debug [class_header]: Header for GKMatchmaker */
// The class instance for the [Matchmaker] class.
var (
	MatchmakerClass     _MatchmakerClass
	MatchmakerClassOnce sync.Once
)

func getMatchmakerClass() _MatchmakerClass {
	MatchmakerClassOnce.Do(func() {
		MatchmakerClass = _MatchmakerClass{objc.GetClass("GKMatchmaker")}
	})
	return MatchmakerClass
}

type _MatchmakerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Matchmaker */
// An interface definition for the [Matchmaker] class.
type IMatchmaker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Matchmaker */
	// properties:
	InviteHandler() unsafe.Pointer
	SetInviteHandler(value unsafe.Pointer)
	ExpectedPlayerCount() int
	SetExpectedPlayerCount(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Matchmaker */
	// methods:
	AddPlayersToMatchMatchRequestCompletionHandler(match IGKMatch, matchRequest IGKMatchRequest, completionHandler unsafe.Pointer)
	Cancel()
	CancelPendingInviteToPlayer(player IGKPlayer)
	FindMatchForRequestWithCompletionHandler(request IGKMatchRequest, completionHandler unsafe.Pointer)
	FindMatchedPlayersWithCompletionHandler(request IGKMatchRequest, completionHandler unsafe.Pointer)
	FindPlayersForHostedRequestWithCompletionHandler(request IGKMatchRequest, completionHandler unsafe.Pointer)
	FinishMatchmakingForMatch(match IGKMatch)
	MatchForInviteCompletionHandler(invite IGKInvite, completionHandler unsafe.Pointer)
	QueryActivityWithCompletionHandler(completionHandler unsafe.Pointer)
	QueryPlayerGroupActivityWithCompletionHandler(playerGroup uint, completionHandler unsafe.Pointer)
	QueryQueueActivityWithCompletionHandler(queueName objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	StartBrowsingForNearbyPlayersWithHandler(reachableHandler unsafe.Pointer)
	StartGroupActivityWithPlayerHandler(handler unsafe.Pointer)
	StopBrowsingForNearbyPlayers()
	StopGroupActivity()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Matchmaker */
// Alloc allocates a new instance without initialization.
func (mc _MatchmakerClass) Alloc() Matchmaker {
	rv := objc.Send[Matchmaker](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatchmakerClass) New() Matchmaker {
	rv := objc.Send[Matchmaker](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Matchmaker) Init() Matchmaker {
	rv := objc.Send[Matchmaker](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Matchmaker) Autorelease() Matchmaker {
	rv := objc.Send[Matchmaker](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatchmaker creates a new Matchmaker instance.
func NewMatchmaker() Matchmaker {
	return getMatchmakerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Matchmaker */
// An object that creates matches with other players without presenting an interface to the players.
//
// Use the class to auto-match players for a quicker game start, programmatically invite specific players, or implement your own interface for players to invite other players. If you want to present a familiar matchmaking GameKit interface to players, instead use either the or class. If you host a game on your own server, you can also use this class to find Game Center players. That is, you implement the networking and communication between the players through your own servers not Game Center. To find players using this class, create a object and configure it according to the parameters of your game. Then, pass the match request and a handler using the method, or the method for hosted games, to the shared object. GameKit calls the handler when players accept their invitations. Implement the handler to set the delegate of the object that GameKit sends and start the game when there are enough players. If the match doesn’t have enough players (for example, some players decline their invitations), you can create another match request and call the method repeatedly until the match’s property is zero. When you have enough players to start the match, call the method to end the matchmaking process. If you provide a SharePlay interface for inviting players, use the and methods to create a group activity on behalf of the player.


// An object that creates matches with other players without presenting an interface to the players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker
type Matchmaker struct {
	objectivec.Object
}

// MatchmakerFrom constructs a [Matchmaker] from an unsafe.Pointer.
//
// An object that creates matches with other players without presenting an interface to the players.
func MatchmakerFrom(ptr unsafe.Pointer) Matchmaker {
	return Matchmaker{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Matchmaker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Matchmaker */

// Returns the singleton matchmaker instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/shared()
func (mc _MatchmakerClass) SharedMatchmaker() IMatchmaker {
	rv := objc.Send[Matchmaker](objc.ID(mc.class), objc.Sel("sharedMatchmaker"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedMatchmaker) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Matchmaker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Matchmaker */

// Invites additional players to an existing match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/addPlayers(to:matchRequest:completionHandler:)
func (m_ Matchmaker) AddPlayersToMatchMatchRequestCompletionHandler(match IGKMatch, matchRequest IGKMatchRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addPlayersToMatch:matchRequest:completionHandler:"), match, matchRequest, completionHandler)
}/* debug [instance_methods/method]: AddPlayersToMatchMatchRequestCompletionHandler */


// Cancels a matchmaking request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/cancel()
func (m_ Matchmaker) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Cancels a pending invitation to another player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/cancelPendingInvite(to:)
func (m_ Matchmaker) CancelPendingInviteToPlayer(player IGKPlayer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelPendingInviteToPlayer:"), player)
}/* debug [instance_methods/method]: CancelPendingInviteToPlayer */


// Initiates a request to find players for a peer-to-peer match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/findMatch(for:withCompletionHandler:)
func (m_ Matchmaker) FindMatchForRequestWithCompletionHandler(request IGKMatchRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("findMatchForRequest:withCompletionHandler:"), request, completionHandler)
}/* debug [instance_methods/method]: FindMatchForRequestWithCompletionHandler */


// Initiates a request to find players for a hosted match that uses matchmaking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/findMatchedPlayers(_:withCompletionHandler:)
func (m_ Matchmaker) FindMatchedPlayersWithCompletionHandler(request IGKMatchRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("findMatchedPlayers:withCompletionHandler:"), request, completionHandler)
}/* debug [instance_methods/method]: FindMatchedPlayersWithCompletionHandler */


// Initiates a request to find players for a hosted match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/findPlayers(forHostedRequest:withCompletionHandler:)
func (m_ Matchmaker) FindPlayersForHostedRequestWithCompletionHandler(request IGKMatchRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("findPlayersForHostedRequest:withCompletionHandler:"), request, completionHandler)
}/* debug [instance_methods/method]: FindPlayersForHostedRequestWithCompletionHandler */


// Informs the server when programmatic matchmaking finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/finishMatchmaking(for:)
func (m_ Matchmaker) FinishMatchmakingForMatch(match IGKMatch) {
	objc.Send[objc.ID](m_.ID, objc.Sel("finishMatchmakingForMatch:"), match)
}/* debug [instance_methods/method]: FinishMatchmakingForMatch */


// Creates a match from an invitation that the local player accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/match(for:completionHandler:)
func (m_ Matchmaker) MatchForInviteCompletionHandler(invite IGKInvite, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("matchForInvite:completionHandler:"), invite, completionHandler)
}/* debug [instance_methods/method]: MatchForInviteCompletionHandler */


// Finds the number of players, across player groups, who recently requested a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/queryActivity(completionHandler:)
func (m_ Matchmaker) QueryActivityWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryActivityWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: QueryActivityWithCompletionHandler */


// Finds the number of players in a player group who recently requested a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/queryPlayerGroupActivity(_:withCompletionHandler:)
func (m_ Matchmaker) QueryPlayerGroupActivityWithCompletionHandler(playerGroup uint, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryPlayerGroupActivity:withCompletionHandler:"), playerGroup, completionHandler)
}/* debug [instance_methods/method]: QueryPlayerGroupActivityWithCompletionHandler */


// Finds the number of players in a specific queue who recently requested a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/queryQueueActivity(_:withCompletionHandler:)
func (m_ Matchmaker) QueryQueueActivityWithCompletionHandler(queueName objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryQueueActivity:withCompletionHandler:"), queueName, completionHandler)
}/* debug [instance_methods/method]: QueryQueueActivityWithCompletionHandler */


// Finds nearby players through Bluetooth or WiFi on the same subnet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/startBrowsingForNearbyPlayers(handler:)
func (m_ Matchmaker) StartBrowsingForNearbyPlayersWithHandler(reachableHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startBrowsingForNearbyPlayersWithHandler:"), reachableHandler)
}/* debug [instance_methods/method]: StartBrowsingForNearbyPlayersWithHandler */


// Begins a SharePlay activity for your game when a FaceTime call is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/startGroupActivity(playerHandler:)
func (m_ Matchmaker) StartGroupActivityWithPlayerHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startGroupActivityWithPlayerHandler:"), handler)
}/* debug [instance_methods/method]: StartGroupActivityWithPlayerHandler */


// Stops finding nearby players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/stopBrowsingForNearbyPlayers()
func (m_ Matchmaker) StopBrowsingForNearbyPlayers() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopBrowsingForNearbyPlayers"))
}/* debug [instance_methods/method]: StopBrowsingForNearbyPlayers */


// Ends a SharePlay activity for the entire group, which the local player activates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/stopGroupActivity()
func (m_ Matchmaker) StopGroupActivity() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopGroupActivity"))
}/* debug [instance_methods/method]: StopGroupActivity */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Matchmaker */

// A block that GameKit calls when an invitation to join a match is accepted by the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/inviteHandler
func (m_ Matchmaker) InviteHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("inviteHandler"))
	return rv
}/* debug [instance_properties/getter]: inviteHandler */


// A block that GameKit calls when an invitation to join a match is accepted by the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/inviteHandler
func (m_ Matchmaker) SetInviteHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInviteHandler:"), value)
}/* debug [instance_properties/setter]: inviteHandler */


// The remaining number of players invited but not yet connected to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/expectedplayercount
func (m_ Matchmaker) ExpectedPlayerCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("expectedPlayerCount"))
	return rv
}/* debug [instance_properties/getter]: expectedPlayerCount */


// The remaining number of players invited but not yet connected to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/expectedplayercount
func (m_ Matchmaker) SetExpectedPlayerCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExpectedPlayerCount:"), value)
}/* debug [instance_properties/setter]: expectedPlayerCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKMatchmaker */



