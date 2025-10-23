// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [MatchmakerViewController] class.
type IMatchmakerViewController interface {
	appkit.IViewController
	DefaultInvitationMessage() string
	SetDefaultInvitationMessage(value string)
	Hosted() bool
	SetHosted(value bool)
	MatchmakerDelegate() objc.ID
	SetMatchmakerDelegate(value objc.ID)
	CanStartWithMinimumPlayers() bool
	SetCanStartWithMinimumPlayers(value bool)
	IsHosted() bool
	SetIsHosted(value bool)
	MatchRequest() GKMatchRequest
	SetMatchRequest(value IGKMatchRequest)
	MatchmakingMode() unsafe.Pointer
	SetMatchmakingMode(value unsafe.Pointer)
}

// An interface that allows a player to invite other players to a real-time game and automatch to fill any empty slots.
//
// Before you create a object, create a object and configure it according to the parameters of your game. Then pass the match request to the initializer to create the view controller. Configure the view controller and set its delegate before you present it to the local player. The view controller allows the local player to choose other players and, optionally, fill empty slots using automatch. If you add the Group Activities capability to your Xcode project, the player can invite others using SharePlay. See . Implement the and protocols to handle when players send and accept invitations. Implement the delegate method to present a object, which you create using the initializer, to the player who accepts an invitation. Then, implement the delegate method to dismiss the view controller and start the game when all players accept their invitations. In iOS, you present and dismiss the view controller from another view controller in your game, using the methods from the class. If you use SwiftUI, you can get the root view controller from the object. For visionOS games, the view controller appears anchored to the window, scene, or view relative to where you present the view controller. For immersive games, set the parent window to a separate window group than the immersive space window group. For macOS games, use the class to present and dismiss the view controller. For the complete matchmaking flow with code fragments, see .


// An interface that allows a player to invite other players to a real-time game and automatch to fill any empty slots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController
type MatchmakerViewController struct {
	appkit.ViewController
}

// MatchmakerViewControllerFrom constructs a [MatchmakerViewController] from an unsafe.Pointer.
//
// An interface that allows a player to invite other players to a real-time game and automatch to fill any empty slots.
func MatchmakerViewControllerFrom(ptr unsafe.Pointer) MatchmakerViewController {
	return MatchmakerViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MatchmakerViewControllerClass) Alloc() MatchmakerViewController {
	rv := objc.Send[MatchmakerViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a matchmaker view controller for the local player to start inviting other players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/init(matchRequest:)
func NewMatchmakerViewControllerWithMatchRequest(request IGKMatchRequest) MatchmakerViewController {
	instance := getMatchmakerViewControllerClass().Alloc()
	rv := objc.Send[MatchmakerViewController](instance.ID, objc.Sel("initWithMatchRequest:"), request)
	rv.Autorelease()
	return rv
}



// The default invitation message sent to a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/defaultInvitationMessage
func (m_ MatchmakerViewController) DefaultInvitationMessage() string {
	rv := objc.Send[string](m_.ID, objc.Sel("defaultInvitationMessage"))
	return rv
}


// The default invitation message sent to a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/defaultInvitationMessage
func (m_ MatchmakerViewController) SetDefaultInvitationMessage(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultInvitationMessage:"), objc.String(value))
}


// A Boolean value that indicates whether the match is hosted or peer-to-peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/isHosted
func (m_ MatchmakerViewController) Hosted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hosted"))
	return rv
}


// A Boolean value that indicates whether the match is hosted or peer-to-peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/isHosted
func (m_ MatchmakerViewController) SetHosted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHosted:"), value)
}


// The object that handles matchmaker view controller changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/matchmakerDelegate
func (m_ MatchmakerViewController) MatchmakerDelegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("matchmakerDelegate"))
	return rv
}


// The object that handles matchmaker view controller changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakerViewController/matchmakerDelegate
func (m_ MatchmakerViewController) SetMatchmakerDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatchmakerDelegate:"), value)
}


// A Boolean value that indicates whether your game can start after a minimum number of players join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/canstartwithminimumplayers
func (m_ MatchmakerViewController) CanStartWithMinimumPlayers() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canStartWithMinimumPlayers"))
	return rv
}


// A Boolean value that indicates whether your game can start after a minimum number of players join a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/canstartwithminimumplayers
func (m_ MatchmakerViewController) SetCanStartWithMinimumPlayers(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCanStartWithMinimumPlayers:"), value)
}


// A Boolean value that indicates whether the match is hosted or peer-to-peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/ishosted
func (m_ MatchmakerViewController) IsHosted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isHosted"))
	return rv
}


// A Boolean value that indicates whether the match is hosted or peer-to-peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/ishosted
func (m_ MatchmakerViewController) SetIsHosted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsHosted:"), value)
}


// The configuration for the desired match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/matchrequest
func (m_ MatchmakerViewController) MatchRequest() GKMatchRequest {
	rv := objc.Send[GKMatchRequest](m_.ID, objc.Sel("matchRequest"))
	return rv
}


// The configuration for the desired match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/matchrequest
func (m_ MatchmakerViewController) SetMatchRequest(value IGKMatchRequest) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatchRequest:"), value)
}


// The mode that a multiplayer game uses to find players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/matchmakingmode
func (m_ MatchmakerViewController) MatchmakingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("matchmakingMode"))
	return rv
}


// The mode that a multiplayer game uses to find players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatchmakerviewcontroller/matchmakingmode
func (m_ MatchmakerViewController) SetMatchmakingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatchmakingMode:"), value)
}


