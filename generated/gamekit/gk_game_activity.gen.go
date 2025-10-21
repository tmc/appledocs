// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GameActivity] class.
var (
	GameActivityClass     _GameActivityClass
	GameActivityClassOnce sync.Once
)

func getGameActivityClass() _GameActivityClass {
	GameActivityClassOnce.Do(func() {
		GameActivityClass = _GameActivityClass{objc.GetClass("GKGameActivity")}
	})
	return GameActivityClass
}

type _GameActivityClass struct {
	class objc.Class
}

// An interface definition for the [GameActivity] class.
type IGameActivity interface {
	objectivec.IObject
	SetProgressOnAchievementToPercentComplete(achievement unsafe.Pointer, percentComplete unsafe.Pointer)
	SetScoreOnLeaderboardToScore(leaderboard unsafe.Pointer, score int)
}

// An object that represents a single instance of a game activity for the current game.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity
type GameActivity struct {
	objectivec.Object
}

// GameActivityFrom constructs a [GameActivity] from an unsafe.Pointer.
//
// An object that represents a single instance of a game activity for the current game.
func GameActivityFrom(ptr unsafe.Pointer) GameActivity {
	return GameActivity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GameActivityClass) Alloc() GameActivity {
	rv := objc.Send[GameActivity](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GameActivityClass) New() GameActivity {
	rv := objc.Send[GameActivity](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GameActivity) Init() GameActivity {
	rv := objc.Send[GameActivity](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GameActivity) Autorelease() GameActivity {
	rv := objc.Send[GameActivity](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGameActivity creates a new GameActivity instance.
func NewGameActivity() GameActivity {
	return getGameActivityClass().New()
}


// Set a progress for an achievement for a player.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/setProgress(on:to:)
func (g_ GameActivity) SetProgressOnAchievementToPercentComplete(achievement unsafe.Pointer, percentComplete unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProgressOnAchievement:toPercentComplete:"), achievement, percentComplete)
}

// Set a score of a leaderboard for a player.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/setScore(on:to:)
func (g_ GameActivity) SetScoreOnLeaderboardToScore(leaderboard unsafe.Pointer, score int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScoreOnLeaderboard:toScore:"), leaderboard, score)
}

// Total time elapsed while in active state.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/duration
func (g_ GameActivity) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](g_.ID, objc.Sel("duration"))
	return rv
}



