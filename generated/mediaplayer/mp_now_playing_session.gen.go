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
	// properties:
	IsActive() bool
	SetIsActive(value bool)
	// methods:
}

// An object that manages Now Playing information and remote commands for multiple players.
//
// An object can have only one Now Playing session. An manages its own player and Now Playing session, so you can’t add your own Now Playing session.


// An object that manages Now Playing information and remote commands for multiple players.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingSession/init(players:)
func NewNowPlayingSessionWithPlayers(players []objc.IObject /* cross-framework: Player */) NowPlayingSession {
	instance := getNowPlayingSessionClass().Alloc()
	rv := objc.Send[NowPlayingSession](instance.ID, objc.Sel("initWithPlayers:"), players)
	rv.Autorelease()
	return rv
}



// A Boolean value that indicates whether the session is the app’s active Now Playing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayingsession/isactive
func (n_ NowPlayingSession) IsActive() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean value that indicates whether the session is the app’s active Now Playing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayingsession/isactive
func (n_ NowPlayingSession) SetIsActive(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsActive:"), value)
}


