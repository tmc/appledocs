// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GKScoreChallenge */


/* debug [class_header]: Header for GKScoreChallenge */
// The class instance for the [ScoreChallenge] class.
var (
	ScoreChallengeClass     _ScoreChallengeClass
	ScoreChallengeClassOnce sync.Once
)

func getScoreChallengeClass() _ScoreChallengeClass {
	ScoreChallengeClassOnce.Do(func() {
		ScoreChallengeClass = _ScoreChallengeClass{objc.GetClass("GKScoreChallenge")}
	})
	return ScoreChallengeClass
}

type _ScoreChallengeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScoreChallenge */
// An interface definition for the [ScoreChallenge] class.
type IScoreChallenge interface {
	IChallenge
	
/* debug [class_interface_properties]: Properties for ScoreChallenge */
	// properties:
	LeaderboardEntry() IGKLeaderboardEntry
	Score() IGKScore
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScoreChallenge */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScoreChallenge */
// Alloc allocates a new instance without initialization.
func (sc _ScoreChallengeClass) Alloc() ScoreChallenge {
	rv := objc.Send[ScoreChallenge](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScoreChallengeClass) New() ScoreChallenge {
	rv := objc.Send[ScoreChallenge](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScoreChallenge) Init() ScoreChallenge {
	rv := objc.Send[ScoreChallenge](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScoreChallenge) Autorelease() ScoreChallenge {
	rv := objc.Send[ScoreChallenge](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScoreChallenge creates a new ScoreChallenge instance.
func NewScoreChallenge() ScoreChallenge {
	return getScoreChallengeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScoreChallenge */
// A type of challenge where a player must beat the leaderboard score of another player.
//
// To complete the challenge, the player must score an equal or better score than the other player. When the player completes the challenge, Game Center issues a new score challenge to the player who initiated the challenge and continues issuing challenges between the players until a player refuses the challenge.


// A type of challenge where a player must beat the leaderboard score of another player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScoreChallenge
type ScoreChallenge struct {
	Challenge
}

// ScoreChallengeFrom constructs a [ScoreChallenge] from an unsafe.Pointer.
//
// A type of challenge where a player must beat the leaderboard score of another player.
func ScoreChallengeFrom(ptr unsafe.Pointer) ScoreChallenge {
	return ScoreChallenge{
		Challenge: ChallengeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScoreChallenge *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScoreChallenge */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScoreChallenge */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScoreChallenge */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScoreChallenge */

// The challenger’s leaderboard score that the player must beat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScoreChallenge/leaderboardEntry
func (s_ ScoreChallenge) LeaderboardEntry() IGKLeaderboardEntry {
	rv := objc.Send[LeaderboardEntry](s_.ID, objc.Sel("leaderboardEntry"))
	return rv
}/* debug [instance_properties/getter]: leaderboardEntry */


// The challenger’s leaderboard score that the player must beat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScoreChallenge/score
func (s_ ScoreChallenge) Score() IGKScore {
	rv := objc.Send[Score](s_.ID, objc.Sel("score"))
	return rv
}/* debug [instance_properties/getter]: score */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (s_ ScoreChallenge) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (s_ ScoreChallenge) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKScoreChallenge */



