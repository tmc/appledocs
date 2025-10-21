// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Match] class.
var (
	MatchClass     _MatchClass
	MatchClassOnce sync.Once
)

func getMatchClass() _MatchClass {
	MatchClassOnce.Do(func() {
		MatchClass = _MatchClass{objc.GetClass("GKMatch")}
	})
	return MatchClass
}

type _MatchClass struct {
	class objc.Class
}

// An interface definition for the [Match] class.
type IMatch interface {
	objectivec.IObject
	VoiceChatWithName(name string) unsafe.Pointer
}

// A peer-to-peer network between a group of players that sign into Game Center.
//
// Matches provide a mechanism for a player to send both game and voice data to other players. You never create a object directly. Instead, GameKit passes a match object to a method or a handler when you set up a multiplayer game. For details, see . If you use the class to find players, implement the delegate method to set the match delegate. If you use the class, set the match delegate in the handler you pass to the method. You can begin exchanging data when two or more players join the match. Implement the delegate method to track when players connect or disconnect from the match. Then use either the or the method to send data. To process the data on the recipient side, implement the delegate method. To implement voice chat, use the method to create one or more voice channels represented by the returned object. When you’re finished with a match, call the method and set the match’s delegate to . Otherwise, GameKit may send to the delegate until all players disconnect from the match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch
type Match struct {
	objectivec.Object
}

// MatchFrom constructs a [Match] from an unsafe.Pointer.
//
// A peer-to-peer network between a group of players that sign into Game Center.
func MatchFrom(ptr unsafe.Pointer) Match {
	return Match{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MatchClass) Alloc() Match {
	rv := objc.Send[Match](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MatchClass) New() Match {
	rv := objc.Send[Match](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Match) Init() Match {
	rv := objc.Send[Match](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Match) Autorelease() Match {
	rv := objc.Send[Match](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatch creates a new Match instance.
func NewMatch() Match {
	return getMatchClass().New()
}


// Joins the local player to a voice channel.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/voiceChat(withName:)
func (m_ Match) VoiceChatWithName(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("voiceChatWithName:"), objc.String(name))
	return rv
}



