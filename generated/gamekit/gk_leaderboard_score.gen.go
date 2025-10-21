// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [LeaderboardScore] class.
var (
	LeaderboardScoreClass     _LeaderboardScoreClass
	LeaderboardScoreClassOnce sync.Once
)

func getLeaderboardScoreClass() _LeaderboardScoreClass {
	LeaderboardScoreClassOnce.Do(func() {
		LeaderboardScoreClass = _LeaderboardScoreClass{objc.GetClass("GKLeaderboardScore")}
	})
	return LeaderboardScoreClass
}

type _LeaderboardScoreClass struct {
	class objc.Class
}

// An interface definition for the [LeaderboardScore] class.
type ILeaderboardScore interface {
	objectivec.IObject
}

// Information about a player’s score on a leaderboard.
//
// A object represents a score on a leaderboard for scores you report for challenges or turn-based games. When you create a object, set the property to the associated leaderboard, the property to the player who earns the score, and the property to the score. Make sure the score is compatible with the score format that you configure in App Store Connect. Then use either the or method to report one or more scores. For details about the score format, see in App Store Connect Help.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardScore
type LeaderboardScore struct {
	objectivec.Object
}

// LeaderboardScoreFrom constructs a [LeaderboardScore] from an unsafe.Pointer.
//
// Information about a player’s score on a leaderboard.
func LeaderboardScoreFrom(ptr unsafe.Pointer) LeaderboardScore {
	return LeaderboardScore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LeaderboardScoreClass) Alloc() LeaderboardScore {
	rv := objc.Send[LeaderboardScore](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LeaderboardScoreClass) New() LeaderboardScore {
	rv := objc.Send[LeaderboardScore](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LeaderboardScore) Init() LeaderboardScore {
	rv := objc.Send[LeaderboardScore](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LeaderboardScore) Autorelease() LeaderboardScore {
	rv := objc.Send[LeaderboardScore](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLeaderboardScore creates a new LeaderboardScore instance.
func NewLeaderboardScore() LeaderboardScore {
	return getLeaderboardScoreClass().New()
}




