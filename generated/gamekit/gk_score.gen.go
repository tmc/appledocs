// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKScore */


/* debug [class_header]: Header for GKScore */
// The class instance for the [Score] class.
var (
	ScoreClass     _ScoreClass
	ScoreClassOnce sync.Once
)

func getScoreClass() _ScoreClass {
	ScoreClassOnce.Do(func() {
		ScoreClass = _ScoreClass{objc.GetClass("GKScore")}
	})
	return ScoreClass
}

type _ScoreClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Score */
// An interface definition for the [Score] class.
type IScore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Score */
	// properties:
	Category() objc.IObject /* cross-framework: NSString */
	SetCategory(value objc.IObject /* cross-framework: NSString */)
	Context() uint64
	SetContext(value uint64)
	Date() objc.IObject /* cross-framework: NSDate */
	FormattedValue() objc.IObject /* cross-framework: NSString */
	LeaderboardIdentifier() objc.IObject /* cross-framework: NSString */
	SetLeaderboardIdentifier(value objc.IObject /* cross-framework: NSString */)
	Player() IGKPlayer
	PlayerID() objc.IObject /* cross-framework: NSString */
	Rank() int
	ShouldSetDefaultLeaderboard() bool
	SetShouldSetDefaultLeaderboard(value bool)
	Value() int64
	SetValue(value int64)
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Score */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Score */
// Alloc allocates a new instance without initialization.
func (sc _ScoreClass) Alloc() Score {
	rv := objc.Send[Score](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScoreClass) New() Score {
	rv := objc.Send[Score](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Score) Init() Score {
	rv := objc.Send[Score](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Score) Autorelease() Score {
	rv := objc.Send[Score](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScore creates a new Score instance.
func NewScore() Score {
	return getScoreClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Score */
// An object containing information for a score that was earned by the player.
//
// Your game creates objects to post scores to a leaderboard on Game Center. When your game retrieves score information from a leaderboard, those scores are returned as objects. Scores and leaderboards work together to help you create a better game. Whenever a new object is created, it is associated with a leaderboard. You must ensure that the score being sent to a leaderboard is compatible with the leaderboard scoring format set in App Store Connect. See for information on how to create a leaderboard in App Store Connect. To report a score to Game Center, your game allocates and initializes a new object, sets the property to the score the player earned, and then calls the method. The mechanism your game uses to calculate scores is up to you to design; scores are only compared within your game.


// An object containing information for a score that was earned by the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore
type Score struct {
	objectivec.Object
}

// ScoreFrom constructs a [Score] from an unsafe.Pointer.
//
// An object containing information for a score that was earned by the player.
func ScoreFrom(ptr unsafe.Pointer) Score {
	return Score{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Score */

// Returns an initialized score object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/init(category:)
func NewScoreWithCategory(category objc.IObject /* cross-framework: NSString */) Score {
	instance := getScoreClass().Alloc()
	rv := objc.Send[Score](instance.ID, objc.Sel("initWithCategory:"), category)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScoreWithCategory */


// Returns an initialized score object using the local player and the current date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/init(leaderboardIdentifier:)
func NewScoreWithLeaderboardIdentifier(identifier objc.IObject /* cross-framework: NSString */) Score {
	instance := getScoreClass().Alloc()
	rv := objc.Send[Score](instance.ID, objc.Sel("initWithLeaderboardIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScoreWithLeaderboardIdentifier */


// Returns an initialized score object for the specified leaderboard and player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/init(leaderboardIdentifier:forPlayer:)
func NewScoreWithLeaderboardIdentifierForPlayer(identifier objc.IObject /* cross-framework: NSString */, playerID objc.IObject /* cross-framework: NSString */) Score {
	instance := getScoreClass().Alloc()
	rv := objc.Send[Score](instance.ID, objc.Sel("initWithLeaderboardIdentifier:forPlayer:"), identifier, playerID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScoreWithLeaderboardIdentifierForPlayer */


// Returns an initialized score object for the specified leaderboard and player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/init(leaderboardIdentifier:player:)
func NewScoreWithLeaderboardIdentifierPlayer(identifier objc.IObject /* cross-framework: NSString */, player IGKPlayer) Score {
	instance := getScoreClass().Alloc()
	rv := objc.Send[Score](instance.ID, objc.Sel("initWithLeaderboardIdentifier:player:"), identifier, player)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScoreWithLeaderboardIdentifierPlayer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Score */

// Reports a list of scores to Game Center
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/report(_:withCompletionHandler:)
func (sc _ScoreClass) ReportScoresWithCompletionHandler(scores []Score, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("reportScores:withCompletionHandler:"), scores, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReportScoresWithCompletionHandler) */


// Submits a list of scores and all eligible challenges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/report(_:withEligibleChallenges:withCompletionHandler:)-2tycl
func (sc _ScoreClass) ReportLeaderboardScoresWithEligibleChallengesWithCompletionHandler(scores []LeaderboardScore, challenges []Challenge, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("reportLeaderboardScores:withEligibleChallenges:withCompletionHandler:"), scores, challenges, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReportLeaderboardScoresWithEligibleChallengesWithCompletionHandler) */


// Submits a list of scores and all eligible challenges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/report(_:withEligibleChallenges:withCompletionHandler:)-3c5lh
func (sc _ScoreClass) ReportScoresWithEligibleChallengesWithCompletionHandler(scores []Score, challenges []Challenge, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("reportScores:withEligibleChallenges:withCompletionHandler:"), scores, challenges, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReportScoresWithEligibleChallengesWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Score */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Score */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Score */

// The leaderboard that this score belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/category
func (s_ Score) Category() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("category"))
	return rv
}/* debug [instance_properties/getter]: category */


// The leaderboard that this score belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/category
func (s_ Score) SetCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCategory:"), value)
}/* debug [instance_properties/setter]: category */


// An integer value used by your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/context
func (s_ Score) Context() uint64 {
	rv := objc.Send[uint64](s_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */


// An integer value used by your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/context
func (s_ Score) SetContext(value uint64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContext:"), value)
}/* debug [instance_properties/setter]: context */


// The date and time when the score was earned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/date
func (s_ Score) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](s_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// Returns the player’s score as a localized string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/formattedValue
func (s_ Score) FormattedValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("formattedValue"))
	return rv
}/* debug [instance_properties/getter]: formattedValue */


// The identifier for the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/leaderboardIdentifier
func (s_ Score) LeaderboardIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("leaderboardIdentifier"))
	return rv
}/* debug [instance_properties/getter]: leaderboardIdentifier */


// The identifier for the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/leaderboardIdentifier
func (s_ Score) SetLeaderboardIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLeaderboardIdentifier:"), value)
}/* debug [instance_properties/setter]: leaderboardIdentifier */


// The player who earned the score.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/player
func (s_ Score) Player() IGKPlayer {
	rv := objc.Send[Player](s_.ID, objc.Sel("player"))
	return rv
}/* debug [instance_properties/getter]: player */


// The player identifier for the player that earned the score.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/playerID
func (s_ Score) PlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("playerID"))
	return rv
}/* debug [instance_properties/getter]: playerID */


// The position of the score in the results of a leaderboard search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/rank
func (s_ Score) Rank() int {
	rv := objc.Send[int](s_.ID, objc.Sel("rank"))
	return rv
}/* debug [instance_properties/getter]: rank */


// A Boolean value that indicates whether this score should also update the default leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/shouldSetDefaultLeaderboard
func (s_ Score) ShouldSetDefaultLeaderboard() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldSetDefaultLeaderboard"))
	return rv
}/* debug [instance_properties/getter]: shouldSetDefaultLeaderboard */


// A Boolean value that indicates whether this score should also update the default leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/shouldSetDefaultLeaderboard
func (s_ Score) SetShouldSetDefaultLeaderboard(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldSetDefaultLeaderboard:"), value)
}/* debug [instance_properties/setter]: shouldSetDefaultLeaderboard */


// The score earned by the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/value
func (s_ Score) Value() int64 {
	rv := objc.Send[int64](s_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The score earned by the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKScore/value
func (s_ Score) SetValue(value int64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (s_ Score) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (s_ Score) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKScore */


