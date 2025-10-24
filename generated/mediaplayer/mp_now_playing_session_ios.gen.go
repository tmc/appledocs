//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NowPlayingSession


// Adds a player to the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/addPlayer(_:)
func (n_ NowPlayingSession) AddPlayer(player objc.IObject /* cross-framework: Player */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("addPlayer:"), player)
}

// Tells the system to make the session the active Now Playing session if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/becomeActiveIfPossible(completion:)
func (n_ NowPlayingSession) BecomeActiveIfPossibleWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("becomeActiveIfPossibleWithCompletion:"), completion)
}

// Removes a player from the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/removePlayer(_:)
func (n_ NowPlayingSession) RemovePlayer(player objc.IObject /* cross-framework: Player */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removePlayer:"), player)
}

// iOS-only properties

// A Boolean that indicates whether Now Playing info automatically publishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/automaticallyPublishesNowPlayingInfo
func (n_ NowPlayingSession) AutomaticallyPublishesNowPlayingInfo() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("automaticallyPublishesNowPlayingInfo"))
	return rv
}
func (n_ NowPlayingSession) SetAutomaticallyPublishesNowPlayingInfo(value bool) {
	n_.ID.Send(objc.RegisterName("setAutomaticallyPublishesNowPlayingInfo:"), value)
}

// A Boolean value that indicates whether the session can become the app’s active Now Playing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/canBecomeActive
func (n_ NowPlayingSession) CanBecomeActive() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("canBecomeActive"))
	return rv
}

// The Now Playing session’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/delegate
func (n_ NowPlayingSession) Delegate() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("delegate"))
	return rv
}
func (n_ NowPlayingSession) SetDelegate(value objc.ID) {
	n_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// A Boolean value that indicates whether the session is the app’s active Now Playing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/isActive
func (n_ NowPlayingSession) Active() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("active"))
	return rv
}

// The Now Playing information center associated with the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/nowPlayingInfoCenter
func (n_ NowPlayingSession) NowPlayingInfoCenter() IMPNowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](n_.ID, objc.Sel("nowPlayingInfoCenter"))
	return rv
}

// The array of players associated with the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/players
func (n_ NowPlayingSession) Players() []objc.IObject /* cross-framework: Player */ {
	rv := objc.Send[[]avfoundation.Player](n_.ID, objc.Sel("players"))
	return rv
}

// The remote command center associated with the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/remoteCommandCenter
func (n_ NowPlayingSession) RemoteCommandCenter() IMPRemoteCommandCenter {
	rv := objc.Send[RemoteCommandCenter](n_.ID, objc.Sel("remoteCommandCenter"))
	return rv
}




