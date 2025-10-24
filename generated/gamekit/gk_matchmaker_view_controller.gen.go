// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GKMatchmakerViewController */


/* debug [class_header]: Header for GKMatchmakerViewController */
// The class instance for the [MatchmakerViewController] class.
var (
	MatchmakerViewControllerClass     _MatchmakerViewControllerClass
	MatchmakerViewControllerClassOnce sync.Once
)

func getMatchmakerViewControllerClass() _MatchmakerViewControllerClass {
	MatchmakerViewControllerClassOnce.Do(func() {
		MatchmakerViewControllerClass = _MatchmakerViewControllerClass{objc.GetClass("GKMatchmakerViewController")}
	})
	return MatchmakerViewControllerClass
}

type _MatchmakerViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatchmakerViewController */
// An interface definition for the [MatchmakerViewController] class.
type IMatchmakerViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for MatchmakerViewController */
	// properties:
	CanStartWithMinimumPlayers() bool
	SetCanStartWithMinimumPlayers(value bool)
	DefaultInvitationMessage() objc.IObject /* cross-framework: NSString */
	SetDefaultInvitationMessage(value objc.IObject /* cross-framework: NSString */)
	Hosted() bool
	SetHosted(value bool)
	MatchmakerDelegate() unsafe.Pointer
	SetMatchmakerDelegate(value unsafe.Pointer)
	MatchmakingMode() MatchmakingMode
	SetMatchmakingMode(value MatchmakingMode)
	MatchRequest() IGKMatchRequest
	IsHosted() bool
	SetIsHosted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatchmakerViewController */
	// methods:
	AddPlayersToMatch(match IGKMatch)
	SetHostedPlayerDidConnect(player IGKPlayer, connected bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatchmakerViewController */
// Alloc allocates a new instance without initialization.
func (mc _MatchmakerViewControllerClass) Alloc() MatchmakerViewController {
	rv := objc.Send[MatchmakerViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatchmakerViewControllerClass) New() MatchmakerViewController {
	rv := objc.Send[MatchmakerViewController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatchmakerViewController) Init() MatchmakerViewController {
	rv := objc.Send[MatchmakerViewController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatchmakerViewController) Autorelease() MatchmakerViewController {
	rv := objc.Send[MatchmakerViewController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatchmakerViewController creates a new MatchmakerViewController instance.
func NewMatchmakerViewController() MatchmakerViewController {
	return getMatchmakerViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatchmakerViewController */
// An interface that allows a player to invite other players to a real-time game and automatch to fill any empty slots.
//
// Before you create a object, create a object and configure it according to the parameters of your game. Then pass the match request to the initializer to create the view controller. Configure the view controller and set its delegate before you present it to the local player. The view controller allows the local player to choose other players and, optionally, fill empty slots using automatch. If you add the Group Activities capability to your Xcode project, the player can invite others using SharePlay. See . Implement the and protocols to handle when players send and accept invitations. Implement the delegate method to present a object, which you create using the initializer, to the player who accepts an invitation. Then, implement the delegate method to dismiss the view controller and start the game when all players accept their invitations. In iOS, you present and dismiss the view controller from another view controller in your game, using the methods from the class. If you use SwiftUI, you can get the root view controller from the object. For visionOS games, the view controller appears anchored to the window, scene, or view relative to where you present the view controller. For immersive games, set the parent window to a separate window group than the immersive space window group. For macOS games, use the class to present and dismiss the view controller. For the complete matchmaking flow with code fragments, see .


// An interface that allows a player to invite other players to a real-time game and automatch to fill any empty slots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController
type MatchmakerViewController struct {
	ViewController
}

// MatchmakerViewControllerFrom constructs a [MatchmakerViewController] from an unsafe.Pointer.
//
// An interface that allows a player to invite other players to a real-time game and automatch to fill any empty slots.
func MatchmakerViewControllerFrom(ptr unsafe.Pointer) MatchmakerViewController {
	return MatchmakerViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatchmakerViewController */

// Creates a matchmaker view controller to present to a player who accepts an invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/init(invite:)
func NewMatchmakerViewControllerWithInvite(invite IGKInvite) MatchmakerViewController {
	instance := getMatchmakerViewControllerClass().Alloc()
	rv := objc.Send[MatchmakerViewController](instance.ID, objc.Sel("initWithInvite:"), invite)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatchmakerViewControllerWithInvite */


// Creates a matchmaker view controller for the local player to start inviting other players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/init(matchRequest:)
func NewMatchmakerViewControllerWithMatchRequest(request IGKMatchRequest) MatchmakerViewController {
	instance := getMatchmakerViewControllerClass().Alloc()
	rv := objc.Send[MatchmakerViewController](instance.ID, objc.Sel("initWithMatchRequest:"), request)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatchmakerViewControllerWithMatchRequest */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatchmakerViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatchmakerViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatchmakerViewController */

// Invites additional players to join an existing match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/addPlayers(to:)
func (m_ MatchmakerViewController) AddPlayersToMatch(match IGKMatch) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addPlayersToMatch:"), match)
}/* debug [instance_methods/method]: AddPlayersToMatch */


// Updates the connection status of a player in a hosted game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/setHostedPlayer(_:didConnect:)
func (m_ MatchmakerViewController) SetHostedPlayerDidConnect(player IGKPlayer, connected bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHostedPlayer:didConnect:"), player, connected)
}/* debug [instance_methods/method]: SetHostedPlayerDidConnect */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatchmakerViewController */

// A Boolean value that indicates whether your game can start after a minimum number of players join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/canStartWithMinimumPlayers
func (m_ MatchmakerViewController) CanStartWithMinimumPlayers() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canStartWithMinimumPlayers"))
	return rv
}/* debug [instance_properties/getter]: canStartWithMinimumPlayers */


// A Boolean value that indicates whether your game can start after a minimum number of players join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/canStartWithMinimumPlayers
func (m_ MatchmakerViewController) SetCanStartWithMinimumPlayers(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCanStartWithMinimumPlayers:"), value)
}/* debug [instance_properties/setter]: canStartWithMinimumPlayers */


// The default invitation message sent to a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/defaultInvitationMessage
func (m_ MatchmakerViewController) DefaultInvitationMessage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("defaultInvitationMessage"))
	return rv
}/* debug [instance_properties/getter]: defaultInvitationMessage */


// The default invitation message sent to a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/defaultInvitationMessage
func (m_ MatchmakerViewController) SetDefaultInvitationMessage(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultInvitationMessage:"), value)
}/* debug [instance_properties/setter]: defaultInvitationMessage */


// A Boolean value that indicates whether the match is hosted or peer-to-peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/isHosted
func (m_ MatchmakerViewController) Hosted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hosted"))
	return rv
}/* debug [instance_properties/getter]: hosted */


// A Boolean value that indicates whether the match is hosted or peer-to-peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/isHosted
func (m_ MatchmakerViewController) SetHosted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHosted:"), value)
}/* debug [instance_properties/setter]: hosted */


// The object that handles matchmaker view controller changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/matchmakerDelegate
func (m_ MatchmakerViewController) MatchmakerDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("matchmakerDelegate"))
	return rv
}/* debug [instance_properties/getter]: matchmakerDelegate */


// The object that handles matchmaker view controller changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/matchmakerDelegate
func (m_ MatchmakerViewController) SetMatchmakerDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatchmakerDelegate:"), value)
}/* debug [instance_properties/setter]: matchmakerDelegate */


// The mode that a multiplayer game uses to find players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/matchmakingMode
func (m_ MatchmakerViewController) MatchmakingMode() MatchmakingMode {
	rv := objc.Send[MatchmakingMode](m_.ID, objc.Sel("matchmakingMode"))
	return rv
}/* debug [instance_properties/getter]: matchmakingMode */


// The mode that a multiplayer game uses to find players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/matchmakingMode
func (m_ MatchmakerViewController) SetMatchmakingMode(value MatchmakingMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatchmakingMode:"), value)
}/* debug [instance_properties/setter]: matchmakingMode */


// The configuration for the desired match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/matchRequest
func (m_ MatchmakerViewController) MatchRequest() IGKMatchRequest {
	rv := objc.Send[MatchRequest](m_.ID, objc.Sel("matchRequest"))
	return rv
}/* debug [instance_properties/getter]: matchRequest */


// A Boolean value that indicates whether the match is hosted or peer-to-peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/ishosted
func (m_ MatchmakerViewController) IsHosted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isHosted"))
	return rv
}/* debug [instance_properties/getter]: isHosted */


// A Boolean value that indicates whether the match is hosted or peer-to-peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/ishosted
func (m_ MatchmakerViewController) SetIsHosted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsHosted:"), value)
}/* debug [instance_properties/setter]: isHosted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKMatchmakerViewController */


