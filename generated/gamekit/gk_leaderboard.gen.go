// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKLeaderboard */


/* debug [class_header]: Header for GKLeaderboard */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Leaderboard */
// An interface definition for the [Leaderboard] class.
type ILeaderboard interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Leaderboard */
	// properties:
	ActivityIdentifier() objc.IObject /* cross-framework: NSString */
	ActivityProperties() foundation.IDictionary
	BaseLeaderboardID() objc.IObject /* cross-framework: NSString */
	Category() objc.IObject /* cross-framework: NSString */
	SetCategory(value objc.IObject /* cross-framework: NSString */)
	Duration() float64
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	IsHidden() bool
	Loading() bool
	LeaderboardDescription() objc.IObject /* cross-framework: NSString */
	LocalPlayerScore() IGKScore
	MaxRange() uint
	NextStartDate() objc.IObject /* cross-framework: NSDate */
	PlayerScope() LeaderboardPlayerScope
	SetPlayerScope(value LeaderboardPlayerScope)
	Range() corefoundation.Range
	SetRange(value corefoundation.Range)
	ReleaseState() ReleaseState
	Scores() []Score
	StartDate() objc.IObject /* cross-framework: NSDate */
	TimeScope() LeaderboardTimeScope
	SetTimeScope(value LeaderboardTimeScope)
	Title() objc.IObject /* cross-framework: NSString */
	Type() LeaderboardType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Leaderboard */
	// methods:
	LoadEntriesForPlayersTimeScopeCompletionHandler(players []Player, timeScope LeaderboardTimeScope, completionHandler unsafe.Pointer)
	LoadEntriesForPlayerScopeTimeScopeRangeCompletionHandler(playerScope LeaderboardPlayerScope, timeScope LeaderboardTimeScope, range_ corefoundation.Range, completionHandler unsafe.Pointer)
	LoadImageWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadPreviousOccurrenceWithCompletionHandler(completionHandler unsafe.Pointer)
	SubmitScoreContextPlayerCompletionHandler(score int, context uint, player IGKPlayer, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Leaderboard */
// Alloc allocates a new instance without initialization.
func (lc _LeaderboardClass) Alloc() Leaderboard {
	rv := objc.Send[Leaderboard](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Leaderboard */
// A leaderboard for a game that Game Center stores.
//
// Leaderboards allow players to compare their scores against other players in your game. You configure a classic or recurring leaderboard in App Store Connect and then access the localized information for a leaderboard in your code using objects. A is persistent, that is, the scores never reset unless you delete the leaderboard. A contains scores for a period of time useful for competitions and encouraging players to try for higher scores. You configure the duration, frequency, and delay between occurrences that Game Center uses to automatically restart the leaderboard in App Store Connect. In your code, you use the identifier you set for the leaderboard in App Store Connect to submit scores or load leaderboards. Use the class method to submit a score to one or more leaderboards. Alternatively, load a recurring leaderboard using the class method and then submit a score using the method. To learn more about recurring leaderboards, see . To retrieve information about all leaderboards in your game, use the class method. To fetch the scores for a leaderboard, use the or method. Use the parameters of these methods to filter the scores to the player’s friends, a rank, and time period when the score occurs. You must create leaderboard objects using one of the load methods above. If the request is successful, GameKit passes corresponding objects to the handler. GameKit doesn’t load the images you add to App Store Connect when it loads the leaderboards. Use the method to get the image for a leaderboard.


// A leaderboard for a game that Game Center stores.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Leaderboard */

// Initializes a leaderboard request to retrieve the scores of a specific group of players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/init(playerIDs:)
func NewLeaderboardWithPlayerIDs(playerIDs []string) Leaderboard {
	instance := getLeaderboardClass().Alloc()
	rv := objc.Send[Leaderboard](instance.ID, objc.Sel("initWithPlayerIDs:"), playerIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLeaderboardWithPlayerIDs */


// Initializes a leaderboard request to retrieve the scores of a specific group of players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/init(players:)
func NewLeaderboardWithPlayers(players []Player) Leaderboard {
	instance := getLeaderboardClass().Alloc()
	rv := objc.Send[Leaderboard](instance.ID, objc.Sel("initWithPlayers:"), players)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLeaderboardWithPlayers */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Leaderboard */

// Loads the list of leaderboard categories along with their corresponding localized titles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/loadCategories(completionHandler:)
func (lc _LeaderboardClass) LoadCategoriesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("loadCategoriesWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadCategoriesWithCompletionHandler) */


// Loads a list of leaderboards from Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/loadLeaderboards(completionHandler:)
func (lc _LeaderboardClass) LoadLeaderboardsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("loadLeaderboardsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadLeaderboardsWithCompletionHandler) */


// Loads leaderboards for the specified leaderboard IDs that Game Center uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/loadLeaderboards(IDs:completionHandler:)
func (lc _LeaderboardClass) LoadLeaderboardsWithIDsCompletionHandler(leaderboardIDs []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("loadLeaderboardsWithIDs:completionHandler:"), leaderboardIDs, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadLeaderboardsWithIDsCompletionHandler) */


// Sets the default leaderboard for the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/setDefault(_:withCompletionHandler:)
func (lc _LeaderboardClass) SetDefaultLeaderboardWithCompletionHandler(leaderboardIdentifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("setDefaultLeaderboard:withCompletionHandler:"), leaderboardIdentifier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetDefaultLeaderboardWithCompletionHandler) */


// Submits a score to multiple leaderboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/submitScore(_:context:player:leaderboardIDs:completionHandler:)
func (lc _LeaderboardClass) SubmitScoreContextPlayerLeaderboardIDsCompletionHandler(score int, context uint, player IGKPlayer, leaderboardIDs []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("submitScore:context:player:leaderboardIDs:completionHandler:"), score, context, player, leaderboardIDs, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SubmitScoreContextPlayerLeaderboardIDsCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Leaderboard */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Leaderboard */

// Returns the scores for the local player and other players for the specified time period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/loadEntries(for:timeScope:completionHandler:)
func (l_ Leaderboard) LoadEntriesForPlayersTimeScopeCompletionHandler(players []Player, timeScope LeaderboardTimeScope, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadEntriesForPlayers:timeScope:completionHandler:"), players, timeScope, completionHandler)
}/* debug [instance_methods/method]: LoadEntriesForPlayersTimeScopeCompletionHandler */


// Returns the scores for the local player and other players for the specified type of player, time period, and ranks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/loadEntries(for:timeScope:range:completionHandler:)
func (l_ Leaderboard) LoadEntriesForPlayerScopeTimeScopeRangeCompletionHandler(playerScope LeaderboardPlayerScope, timeScope LeaderboardTimeScope, range_ corefoundation.Range, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadEntriesForPlayerScope:timeScope:range:completionHandler:"), playerScope, timeScope, range_, completionHandler)
}/* debug [instance_methods/method]: LoadEntriesForPlayerScopeTimeScopeRangeCompletionHandler */


// Loads the image for the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/loadImage(completionHandler:)
func (l_ Leaderboard) LoadImageWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadImageWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadImageWithCompletionHandler */


// Loads the previous recurring leaderboard occurrence that the player submits a score to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/loadPreviousOccurrence(completionHandler:)
func (l_ Leaderboard) LoadPreviousOccurrenceWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadPreviousOccurrenceWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadPreviousOccurrenceWithCompletionHandler */


// Submits a score to the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/submitScore(_:context:player:completionHandler:)
func (l_ Leaderboard) SubmitScoreContextPlayerCompletionHandler(score int, context uint, player IGKPlayer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("submitScore:context:player:completionHandler:"), score, context, player, completionHandler)
}/* debug [instance_methods/method]: SubmitScoreContextPlayerCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Leaderboard */

// The identifier of the game activity associated with this leaderboard, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/activityIdentifier
func (l_ Leaderboard) ActivityIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("activityIdentifier"))
	return rv
}/* debug [instance_properties/getter]: activityIdentifier */


// The properties when associating this leaderboard with a game activity, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/activityProperties
func (l_ Leaderboard) ActivityProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](l_.ID, objc.Sel("activityProperties"))
	return rv
}/* debug [instance_properties/getter]: activityProperties */


// The ID that Game Center uses to identify this leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/baseLeaderboardID
func (l_ Leaderboard) BaseLeaderboardID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("baseLeaderboardID"))
	return rv
}/* debug [instance_properties/getter]: baseLeaderboardID */


// The named leaderboard to retrieve information from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/category
func (l_ Leaderboard) Category() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("category"))
	return rv
}/* debug [instance_properties/getter]: category */


// The named leaderboard to retrieve information from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/category
func (l_ Leaderboard) SetCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCategory:"), value)
}/* debug [instance_properties/setter]: category */


// The duration from the start date that a recurring leaderboard occurrence accepts scores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/duration
func (l_ Leaderboard) Duration() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The identifier for the group the leaderboard belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/groupIdentifier
func (l_ Leaderboard) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("groupIdentifier"))
	return rv
}/* debug [instance_properties/getter]: groupIdentifier */


// The named leaderboard to retrieve information from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/identifier
func (l_ Leaderboard) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The named leaderboard to retrieve information from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/identifier
func (l_ Leaderboard) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// A Boolean value that indicates whether the current leaderboard isn’t visible in Game Center views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/isHidden
func (l_ Leaderboard) IsHidden() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// A Boolean value that indicates whether the leaderboard object is retrieving scores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/isLoading
func (l_ Leaderboard) Loading() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("loading"))
	return rv
}/* debug [instance_properties/getter]: loading */


// The description of this Leaderboard as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/leaderboardDescription
func (l_ Leaderboard) LeaderboardDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("leaderboardDescription"))
	return rv
}/* debug [instance_properties/getter]: leaderboardDescription */


// The score that the local player earns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/localPlayerScore
func (l_ Leaderboard) LocalPlayerScore() IGKScore {
	rv := objc.Send[Score](l_.ID, objc.Sel("localPlayerScore"))
	return rv
}/* debug [instance_properties/getter]: localPlayerScore */


// The size of the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/maxRange
func (l_ Leaderboard) MaxRange() uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("maxRange"))
	return rv
}/* debug [instance_properties/getter]: maxRange */


// The date and time the next recurring leaderboard occurrence starts accepting scores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/nextStartDate
func (l_ Leaderboard) NextStartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](l_.ID, objc.Sel("nextStartDate"))
	return rv
}/* debug [instance_properties/getter]: nextStartDate */


// A filter that restricts the search to a subset of the players in Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/playerScope-swift.property
func (l_ Leaderboard) PlayerScope() LeaderboardPlayerScope {
	rv := objc.Send[LeaderboardPlayerScope](l_.ID, objc.Sel("playerScope"))
	return rv
}/* debug [instance_properties/getter]: playerScope */


// A filter that restricts the search to a subset of the players in Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/playerScope-swift.property
func (l_ Leaderboard) SetPlayerScope(value LeaderboardPlayerScope) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPlayerScope:"), value)
}/* debug [instance_properties/setter]: playerScope */


// The numerical score rankings to return from the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/range
func (l_ Leaderboard) Range() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](l_.ID, objc.Sel("range"))
	return rv
}/* debug [instance_properties/getter]: range */


// The numerical score rankings to return from the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/range
func (l_ Leaderboard) SetRange(value corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRange:"), value)
}/* debug [instance_properties/setter]: range */


// The release state of the leaderboard in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/releaseState
func (l_ Leaderboard) ReleaseState() ReleaseState {
	rv := objc.Send[ReleaseState](l_.ID, objc.Sel("releaseState"))
	return rv
}/* debug [instance_properties/getter]: releaseState */


// An array of scores that contains the scores that the search returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/scores
func (l_ Leaderboard) Scores() []Score {
	rv := objc.Send[[]Score](l_.ID, objc.Sel("scores"))
	return rv
}/* debug [instance_properties/getter]: scores */


// The date and time a recurring leaderboard occurrence starts accepting scores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/startDate
func (l_ Leaderboard) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](l_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// A filter that restricts the search to scores within a specific period of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/timeScope-swift.property
func (l_ Leaderboard) TimeScope() LeaderboardTimeScope {
	rv := objc.Send[LeaderboardTimeScope](l_.ID, objc.Sel("timeScope"))
	return rv
}/* debug [instance_properties/getter]: timeScope */


// A filter that restricts the search to scores within a specific period of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/timeScope-swift.property
func (l_ Leaderboard) SetTimeScope(value LeaderboardTimeScope) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTimeScope:"), value)
}/* debug [instance_properties/setter]: timeScope */


// The localized title for the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/title
func (l_ Leaderboard) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The type of leaderboard, classic or recurring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/type
func (l_ Leaderboard) Type() LeaderboardType {
	rv := objc.Send[LeaderboardType](l_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKLeaderboard */


