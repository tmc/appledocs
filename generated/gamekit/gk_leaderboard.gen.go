// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Leaderboard] class.
var (
	LeaderboardClass     _LeaderboardClass
	LeaderboardClassOnce sync.Once
)

func getLeaderboardClass() _LeaderboardClass {
	LeaderboardClassOnce.Do(func() {
		LeaderboardClass = _LeaderboardClass{objc.GetClass("GKLeaderboard")}
	})
	return LeaderboardClass
}

type _LeaderboardClass struct {
	class objc.Class
}

// An interface definition for the [Leaderboard] class.
type ILeaderboard interface {
	objectivec.IObject
}

// A leaderboard for a game that Game Center stores.
//
// Leaderboards allow players to compare their scores against other players in your game. You configure a classic or recurring leaderboard in App Store Connect and then access the localized information for a leaderboard in your code using objects. A is persistent, that is, the scores never reset unless you delete the leaderboard. A contains scores for a period of time useful for competitions and encouraging players to try for higher scores. You configure the duration, frequency, and delay between occurrences that Game Center uses to automatically restart the leaderboard in App Store Connect. In your code, you use the identifier you set for the leaderboard in App Store Connect to submit scores or load leaderboards. Use the class method to submit a score to one or more leaderboards. Alternatively, load a recurring leaderboard using the class method and then submit a score using the method. To learn more about recurring leaderboards, see . To retrieve information about all leaderboards in your game, use the class method. To fetch the scores for a leaderboard, use the or method. Use the parameters of these methods to filter the scores to the player’s friends, a rank, and time period when the score occurs. You must create leaderboard objects using one of the load methods above. If the request is successful, GameKit passes corresponding objects to the handler. GameKit doesn’t load the images you add to App Store Connect when it loads the leaderboards. Use the method to get the image for a leaderboard.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard
type Leaderboard struct {
	objectivec.Object
}

// LeaderboardFrom constructs a [Leaderboard] from an unsafe.Pointer.
//
// A leaderboard for a game that Game Center stores.
func LeaderboardFrom(ptr unsafe.Pointer) Leaderboard {
	return Leaderboard{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LeaderboardClass) Alloc() Leaderboard {
	rv := objc.Send[Leaderboard](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LeaderboardClass) New() Leaderboard {
	rv := objc.Send[Leaderboard](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Leaderboard) Init() Leaderboard {
	rv := objc.Send[Leaderboard](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Leaderboard) Autorelease() Leaderboard {
	rv := objc.Send[Leaderboard](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLeaderboard creates a new Leaderboard instance.
func NewLeaderboard() Leaderboard {
	return getLeaderboardClass().New()
}


// The date and time a recurring leaderboard occurrence starts accepting scores.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/startDate
func (l_ Leaderboard) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("startDate"))
	return rv
}



