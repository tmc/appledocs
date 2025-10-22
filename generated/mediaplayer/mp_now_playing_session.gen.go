// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NowPlayingSession] class.
var (
	NowPlayingSessionClass     _NowPlayingSessionClass
	NowPlayingSessionClassOnce sync.Once
)

func getNowPlayingSessionClass() _NowPlayingSessionClass {
	NowPlayingSessionClassOnce.Do(func() {
		NowPlayingSessionClass = _NowPlayingSessionClass{objc.GetClass("MPNowPlayingSession")}
	})
	return NowPlayingSessionClass
}

type _NowPlayingSessionClass struct {
	class objc.Class
}

// An interface definition for the [NowPlayingSession] class.
type INowPlayingSession interface {
	objectivec.IObject
	AddPlayer(player avfoundation.IPlayer)
	BecomeActiveIfPossibleWithCompletion(completion unsafe.Pointer)
	RemovePlayer(player avfoundation.IPlayer)
	AutomaticallyPublishesNowPlayingInfo() bool
	SetAutomaticallyPublishesNowPlayingInfo(value bool)
	CanBecomeActive() bool
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Active() bool
	NowPlayingInfoCenter() MPNowPlayingInfoCenter
	Players() []avfoundation.Player
	RemoteCommandCenter() MPRemoteCommandCenter
	IsActive() bool
	SetIsActive(value bool)
}

// An object that manages Now Playing information and remote commands for multiple players.
//
// An object can have only one Now Playing session. An manages its own player and Now Playing session, so you can’t add your own Now Playing session.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession
type NowPlayingSession struct {
	objectivec.Object
}

// NowPlayingSessionFrom constructs a [NowPlayingSession] from an unsafe.Pointer.
//
// An object that manages Now Playing information and remote commands for multiple players.
func NowPlayingSessionFrom(ptr unsafe.Pointer) NowPlayingSession {
	return NowPlayingSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NowPlayingSessionClass) Alloc() NowPlayingSession {
	rv := objc.Send[NowPlayingSession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NowPlayingSessionClass) New() NowPlayingSession {
	rv := objc.Send[NowPlayingSession](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NowPlayingSession) Init() NowPlayingSession {
	rv := objc.Send[NowPlayingSession](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NowPlayingSession) Autorelease() NowPlayingSession {
	rv := objc.Send[NowPlayingSession](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNowPlayingSession creates a new NowPlayingSession instance.
func NewNowPlayingSession() NowPlayingSession {
	return getNowPlayingSessionClass().New()
}




// Creates a Now Playing session object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/init(players:)
func NewNowPlayingSessionWithPlayers(players []avfoundation.IPlayer) NowPlayingSession {
	instance := getNowPlayingSessionClass().Alloc()
	rv := objc.Send[NowPlayingSession](instance.ID, objc.Sel("initWithPlayers:"), players)
	rv.Autorelease()
	return rv
}


// Adds a player to the session.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/addPlayer(_:)
func (n_ NowPlayingSession) AddPlayer(player avfoundation.IPlayer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("addPlayer:"), player)
}

// Tells the system to make the session the active Now Playing session if possible.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/becomeActiveIfPossible(completion:)
func (n_ NowPlayingSession) BecomeActiveIfPossibleWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("becomeActiveIfPossibleWithCompletion:"), completion)
}

// Removes a player from the session.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/removePlayer(_:)
func (n_ NowPlayingSession) RemovePlayer(player avfoundation.IPlayer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removePlayer:"), player)
}

// A Boolean that indicates whether Now Playing info automatically publishes.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/automaticallyPublishesNowPlayingInfo
func (n_ NowPlayingSession) AutomaticallyPublishesNowPlayingInfo() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("automaticallyPublishesNowPlayingInfo"))
	return rv
}


// SetAutomaticallyPublishesNowPlayingInfo sets the value of the automaticallyPublishesNowPlayingInfo property.
// A Boolean that indicates whether Now Playing info automatically publishes.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/automaticallyPublishesNowPlayingInfo
func (n_ NowPlayingSession) SetAutomaticallyPublishesNowPlayingInfo(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAutomaticallyPublishesNowPlayingInfo:"), value)
}

// A Boolean value that indicates whether the session can become the app’s active Now Playing session.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/canBecomeActive
func (n_ NowPlayingSession) CanBecomeActive() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("canBecomeActive"))
	return rv
}

// The Now Playing session’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/delegate
func (n_ NowPlayingSession) Delegate() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The Now Playing session’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/delegate
func (n_ NowPlayingSession) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the session is the app’s active Now Playing session.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/isActive
func (n_ NowPlayingSession) Active() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("active"))
	return rv
}

// The Now Playing information center associated with the session.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/nowPlayingInfoCenter
func (n_ NowPlayingSession) NowPlayingInfoCenter() MPNowPlayingInfoCenter {
	rv := objc.Send[MPNowPlayingInfoCenter](n_.ID, objc.Sel("nowPlayingInfoCenter"))
	return rv
}

// The array of players associated with the session.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/players
func (n_ NowPlayingSession) Players() []avfoundation.Player {
	rv := objc.Send[[]avfoundation.Player](n_.ID, objc.Sel("players"))
	return rv
}

// The remote command center associated with the session.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/remoteCommandCenter
func (n_ NowPlayingSession) RemoteCommandCenter() MPRemoteCommandCenter {
	rv := objc.Send[MPRemoteCommandCenter](n_.ID, objc.Sel("remoteCommandCenter"))
	return rv
}

// A Boolean value that indicates whether the session is the app’s active Now Playing session.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayingsession/isactive
func (n_ NowPlayingSession) IsActive() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isActive"))
	return rv
}


// SetIsActive sets the value of the isActive property.
// A Boolean value that indicates whether the session is the app’s active Now Playing session.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayingsession/isactive
func (n_ NowPlayingSession) SetIsActive(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsActive:"), value)
}


