// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKLeaderboardEntry */


/* debug [class_header]: Header for GKLeaderboardEntry */
// The class instance for the [LeaderboardEntry] class.
var (
	LeaderboardEntryClass     _LeaderboardEntryClass
	LeaderboardEntryClassOnce sync.Once
)

func getLeaderboardEntryClass() _LeaderboardEntryClass {
	LeaderboardEntryClassOnce.Do(func() {
		LeaderboardEntryClass = _LeaderboardEntryClass{objc.GetClass("GKLeaderboardEntry")}
	})
	return LeaderboardEntryClass
}

type _LeaderboardEntryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LeaderboardEntry */
// An interface definition for the [LeaderboardEntry] class.
type ILeaderboardEntry interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LeaderboardEntry */
	// properties:
	Context() uint
	Date() objc.IObject /* cross-framework: NSDate */
	FormattedScore() objc.IObject /* cross-framework: NSString */
	Player() IGKPlayer
	Rank() int
	Score() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LeaderboardEntry */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LeaderboardEntry */
// Alloc allocates a new instance without initialization.
func (lc _LeaderboardEntryClass) Alloc() LeaderboardEntry {
	rv := objc.Send[LeaderboardEntry](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LeaderboardEntryClass) New() LeaderboardEntry {
	rv := objc.Send[LeaderboardEntry](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LeaderboardEntry) Init() LeaderboardEntry {
	rv := objc.Send[LeaderboardEntry](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LeaderboardEntry) Autorelease() LeaderboardEntry {
	rv := objc.Send[LeaderboardEntry](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLeaderboardEntry creates a new LeaderboardEntry instance.
func NewLeaderboardEntry() LeaderboardEntry {
	return getLeaderboardEntryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LeaderboardEntry */
// Information about a single score by a player on a leaderboard.


// Information about a single score by a player on a leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/Entry
type LeaderboardEntry struct {
	objectivec.Object
}

// LeaderboardEntryFrom constructs a [LeaderboardEntry] from an unsafe.Pointer.
//
// Information about a single score by a player on a leaderboard.
func LeaderboardEntryFrom(ptr unsafe.Pointer) LeaderboardEntry {
	return LeaderboardEntry{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LeaderboardEntry *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LeaderboardEntry */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LeaderboardEntry */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LeaderboardEntry */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LeaderboardEntry */

// An integer value that your game uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/Entry/context
func (l_ LeaderboardEntry) Context() uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */


// The date and time when the player earns the score.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/Entry/date
func (l_ LeaderboardEntry) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](l_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// The player’s score as a localized string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/Entry/formattedScore
func (l_ LeaderboardEntry) FormattedScore() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("formattedScore"))
	return rv
}/* debug [instance_properties/getter]: formattedScore */


// The player who earns the score.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/Entry/player
func (l_ LeaderboardEntry) Player() IGKPlayer {
	rv := objc.Send[Player](l_.ID, objc.Sel("player"))
	return rv
}/* debug [instance_properties/getter]: player */


// The position of the score in the results of a leaderboard search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/Entry/rank
func (l_ LeaderboardEntry) Rank() int {
	rv := objc.Send[int](l_.ID, objc.Sel("rank"))
	return rv
}/* debug [instance_properties/getter]: rank */


// The score that the player earns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/Entry/score
func (l_ LeaderboardEntry) Score() int {
	rv := objc.Send[int](l_.ID, objc.Sel("score"))
	return rv
}/* debug [instance_properties/getter]: score */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKLeaderboardEntry */



