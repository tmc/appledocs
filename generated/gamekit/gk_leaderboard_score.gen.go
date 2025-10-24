// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKLeaderboardScore */


/* debug [class_header]: Header for GKLeaderboardScore */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LeaderboardScore */
// An interface definition for the [LeaderboardScore] class.
type ILeaderboardScore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LeaderboardScore */
	// properties:
	Context() uint
	SetContext(value uint)
	LeaderboardID() objc.IObject /* cross-framework: NSString */
	SetLeaderboardID(value objc.IObject /* cross-framework: NSString */)
	Player() IGKPlayer
	SetPlayer(value IGKPlayer)
	Value() int
	SetValue(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LeaderboardScore */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LeaderboardScore */
// Alloc allocates a new instance without initialization.
func (lc _LeaderboardScoreClass) Alloc() LeaderboardScore {
	rv := objc.Send[LeaderboardScore](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LeaderboardScore */
// Information about a player’s score on a leaderboard.
//
// A object represents a score on a leaderboard for scores you report for challenges or turn-based games. When you create a object, set the property to the associated leaderboard, the property to the player who earns the score, and the property to the score. Make sure the score is compatible with the score format that you configure in App Store Connect. Then use either the or method to report one or more scores. For details about the score format, see in App Store Connect Help.


// Information about a player’s score on a leaderboard.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LeaderboardScore *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LeaderboardScore */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LeaderboardScore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LeaderboardScore */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LeaderboardScore */

// An integer value that your game uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardScore/context
func (l_ LeaderboardScore) Context() uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */


// An integer value that your game uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardScore/context
func (l_ LeaderboardScore) SetContext(value uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContext:"), value)
}/* debug [instance_properties/setter]: context */


// The ID that Game Center uses for the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardScore/leaderboardID
func (l_ LeaderboardScore) LeaderboardID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("leaderboardID"))
	return rv
}/* debug [instance_properties/getter]: leaderboardID */


// The ID that Game Center uses for the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardScore/leaderboardID
func (l_ LeaderboardScore) SetLeaderboardID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLeaderboardID:"), value)
}/* debug [instance_properties/setter]: leaderboardID */


// The player who earns the score.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardScore/player
func (l_ LeaderboardScore) Player() IGKPlayer {
	rv := objc.Send[Player](l_.ID, objc.Sel("player"))
	return rv
}/* debug [instance_properties/getter]: player */


// The player who earns the score.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardScore/player
func (l_ LeaderboardScore) SetPlayer(value IGKPlayer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPlayer:"), value)
}/* debug [instance_properties/setter]: player */


// The score that the player earns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardScore/value
func (l_ LeaderboardScore) Value() int {
	rv := objc.Send[int](l_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The score that the player earns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardScore/value
func (l_ LeaderboardScore) SetValue(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKLeaderboardScore */



