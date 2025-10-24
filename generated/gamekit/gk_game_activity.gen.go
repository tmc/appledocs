// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKGameActivity */


/* debug [class_header]: Header for GKGameActivity */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GameActivity */
// An interface definition for the [GameActivity] class.
type IGameActivity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GameActivity */
	// properties:
	Achievements() unsafe.Pointer
	ActivityDefinition() IGKGameActivityDefinition
	CreationDate() objc.IObject /* cross-framework: NSDate */
	Duration() float64
	EndDate() objc.IObject /* cross-framework: NSDate */
	Identifier() objc.IObject /* cross-framework: NSString */
	LastResumeDate() objc.IObject /* cross-framework: NSDate */
	LeaderboardScores() unsafe.Pointer
	PartyCode() objc.IObject /* cross-framework: NSString */
	PartyURL() objc.IObject /* cross-framework: NSURL */
	Properties() foundation.IDictionary
	SetProperties(value foundation.IDictionary)
	StartDate() objc.IObject /* cross-framework: NSDate */
	State() GameActivityState
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GameActivity */
	// methods:
	End()
	FindMatchWithCompletionHandler(completionHandler unsafe.Pointer)
	FindPlayersForHostedMatchWithCompletionHandler(completionHandler unsafe.Pointer)
	MakeMatchRequest() IMatchRequest
	Pause()
	GetProgressOnAchievement(achievement IGKAchievement) float64
	RemoveAchievements(achievements []Achievement)
	RemoveScoresFromLeaderboards(leaderboards []Leaderboard)
	Resume()
	GetScoreOnLeaderboard(leaderboard IGKLeaderboard) ILeaderboardScore
	SetAchievementCompleted(achievement IGKAchievement)
	SetProgressOnAchievementToPercentComplete(achievement IGKAchievement, percentComplete float64)
	SetScoreOnLeaderboardToScore(leaderboard IGKLeaderboard, score int)
	SetScoreOnLeaderboardToScoreContext(leaderboard IGKLeaderboard, score int, context uint)
	Start()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GameActivity */
// Alloc allocates a new instance without initialization.
func (gc _GameActivityClass) Alloc() GameActivity {
	rv := objc.Send[GameActivity](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GameActivity */
// An object that represents a single instance of a game activity for the current game.


// An object that represents a single instance of a game activity for the current game.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GameActivity */

// Creates a game activity with definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/init(definition:)
func NewGameActivityWithDefinition(activityDefinition IGKGameActivityDefinition) GameActivity {
	instance := getGameActivityClass().Alloc()
	rv := objc.Send[GameActivity](instance.ID, objc.Sel("initWithDefinition:"), activityDefinition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGameActivityWithDefinition */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GameActivity */

// Checks whether there is a pending activity to handle for the current game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/checkPendingGameActivityExistence(completionHandler:)
func (gc _GameActivityClass) CheckPendingGameActivityExistenceWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("checkPendingGameActivityExistenceWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CheckPendingGameActivityExistenceWithCompletionHandler) */


// Checks whether a party code is in valid format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/isValidPartyCode(_:)
func (gc _GameActivityClass) IsValidPartyCode(partyCode objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(gc.class), objc.Sel("isValidPartyCode:"), partyCode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsValidPartyCode) */


// Creates and starts a game activity with a definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/start(definition:)
func (gc _GameActivityClass) StartWithDefinitionError(activityDefinition IGKGameActivityDefinition, error_ unsafe.Pointer) IGameActivity {
	rv := objc.Send[GameActivity](objc.ID(gc.class), objc.Sel("startWithDefinition:error:"), activityDefinition, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StartWithDefinitionError) */


// Creates and starts a new game activity with a custom party code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/start(definition:partyCode:)
func (gc _GameActivityClass) StartWithDefinitionPartyCodeError(activityDefinition IGKGameActivityDefinition, partyCode objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) IGameActivity {
	rv := objc.Send[GameActivity](objc.ID(gc.class), objc.Sel("startWithDefinition:partyCode:error:"), activityDefinition, partyCode, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StartWithDefinitionPartyCodeError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GameActivity */

// Allowed characters for the party code to be used to share this activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/validPartyCodeAlphabet
func (gc _GameActivityClass) ValidPartyCodeAlphabet() []string {
	rv := objc.Send[[]string](objc.ID(gc.class), objc.Sel("validPartyCodeAlphabet"))
	return rv
}/* debug [class_properties_class/property]: validPartyCodeAlphabet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GameActivity */

// Ends the game activity if it’s not already ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/end()
func (g_ GameActivity) End() {
	objc.Send[objc.ID](g_.ID, objc.Sel("end"))
}/* debug [instance_methods/method]: End */


// Use information from the activity to find matches for the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/findMatch(completionHandler:)
func (g_ GameActivity) FindMatchWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("findMatchWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: FindMatchWithCompletionHandler */


// Use information from the activity to find server hosted players for the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/findPlayersForHostedMatch(completionHandler:)
func (g_ GameActivity) FindPlayersForHostedMatchWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("findPlayersForHostedMatchWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: FindPlayersForHostedMatchWithCompletionHandler */


// Makes a match request object with information from the activity, which you can use to find matches for the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/makeMatchRequest()
func (g_ GameActivity) MakeMatchRequest() IMatchRequest {
	rv := objc.Send[MatchRequest](g_.ID, objc.Sel("makeMatchRequest"))
	return rv
}/* debug [instance_methods/method]: MakeMatchRequest */


// Pauses the game activity if it’s not already paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/pause()
func (g_ GameActivity) Pause() {
	objc.Send[objc.ID](g_.ID, objc.Sel("pause"))
}/* debug [instance_methods/method]: Pause */


// Get the achievement progress from a specific achievement of the local player if previously set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/progress(on:)
func (g_ GameActivity) GetProgressOnAchievement(achievement IGKAchievement) float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("getProgressOnAchievement:"), achievement)
	return rv
}/* debug [instance_methods/method]: GetProgressOnAchievement */


// Removes all achievements if they exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/removeAchievements(_:)
func (g_ GameActivity) RemoveAchievements(achievements []Achievement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("removeAchievements:"), achievements)
}/* debug [instance_methods/method]: RemoveAchievements */


// Removes all scores from leaderboards for a player if exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/removeScores(from:)
func (g_ GameActivity) RemoveScoresFromLeaderboards(leaderboards []Leaderboard) {
	objc.Send[objc.ID](g_.ID, objc.Sel("removeScoresFromLeaderboards:"), leaderboards)
}/* debug [instance_methods/method]: RemoveScoresFromLeaderboards */


// Resumes the game activity if it was paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/resume()
func (g_ GameActivity) Resume() {
	objc.Send[objc.ID](g_.ID, objc.Sel("resume"))
}/* debug [instance_methods/method]: Resume */


// Get the leaderboard score from a specific leaderboard of the local player if previously set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/score(on:)
func (g_ GameActivity) GetScoreOnLeaderboard(leaderboard IGKLeaderboard) ILeaderboardScore {
	rv := objc.Send[LeaderboardScore](g_.ID, objc.Sel("getScoreOnLeaderboard:"), leaderboard)
	return rv
}/* debug [instance_methods/method]: GetScoreOnLeaderboard */


// Set progress to 100% for an achievement for a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/setAchievementCompleted(_:)
func (g_ GameActivity) SetAchievementCompleted(achievement IGKAchievement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAchievementCompleted:"), achievement)
}/* debug [instance_methods/method]: SetAchievementCompleted */


// Set a progress for an achievement for a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/setProgress(on:to:)
func (g_ GameActivity) SetProgressOnAchievementToPercentComplete(achievement IGKAchievement, percentComplete float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProgressOnAchievement:toPercentComplete:"), achievement, percentComplete)
}/* debug [instance_methods/method]: SetProgressOnAchievementToPercentComplete */


// Set a score of a leaderboard for a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/setScore(on:to:)
func (g_ GameActivity) SetScoreOnLeaderboardToScore(leaderboard IGKLeaderboard, score int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScoreOnLeaderboard:toScore:"), leaderboard, score)
}/* debug [instance_methods/method]: SetScoreOnLeaderboardToScore */


// Set a score of a leaderboard with a context for a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/setScore(on:to:context:)
func (g_ GameActivity) SetScoreOnLeaderboardToScoreContext(leaderboard IGKLeaderboard, score int, context uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScoreOnLeaderboard:toScore:context:"), leaderboard, score, context)
}/* debug [instance_methods/method]: SetScoreOnLeaderboardToScoreContext */


// Starts the game activity if it’s not already started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/start()
func (g_ GameActivity) Start() {
	objc.Send[objc.ID](g_.ID, objc.Sel("start"))
}/* debug [instance_methods/method]: Start */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GameActivity */

// All achievements that have been associated with this activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/achievements
func (g_ GameActivity) Achievements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("achievements"))
	return rv
}/* debug [instance_properties/getter]: achievements */


// The activity definition that this activity instance is based on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/activityDefinition
func (g_ GameActivity) ActivityDefinition() IGKGameActivityDefinition {
	rv := objc.Send[GameActivityDefinition](g_.ID, objc.Sel("activityDefinition"))
	return rv
}/* debug [instance_properties/getter]: activityDefinition */


// The date when the activity was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/creationDate
func (g_ GameActivity) CreationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](g_.ID, objc.Sel("creationDate"))
	return rv
}/* debug [instance_properties/getter]: creationDate */


// The total time elapsed while in active state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/duration
func (g_ GameActivity) Duration() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The date when the activity was officially ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/endDate
func (g_ GameActivity) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](g_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The identifier of this activity instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/identifier
func (g_ GameActivity) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The date when the activity was last resumed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/lastResumeDate
func (g_ GameActivity) LastResumeDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](g_.ID, objc.Sel("lastResumeDate"))
	return rv
}/* debug [instance_properties/getter]: lastResumeDate */


// All leaderboard scores that have been associated with this activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/leaderboardScores
func (g_ GameActivity) LeaderboardScores() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("leaderboardScores"))
	return rv
}/* debug [instance_properties/getter]: leaderboardScores */


// If the game supports party code, this is the party code that can be shared among players to join the party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/partyCode
func (g_ GameActivity) PartyCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("partyCode"))
	return rv
}/* debug [instance_properties/getter]: partyCode */


// If the game supports party code, this is the URL that can be shared among players to join the party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/partyURL
func (g_ GameActivity) PartyURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](g_.ID, objc.Sel("partyURL"))
	return rv
}/* debug [instance_properties/getter]: partyURL */


// Properties that contain additional information about the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/properties
func (g_ GameActivity) Properties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */


// Properties that contain additional information about the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/properties
func (g_ GameActivity) SetProperties(value foundation.IDictionary) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProperties:"), value)
}/* debug [instance_properties/setter]: properties */


// The date when the activity was initially started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/startDate
func (g_ GameActivity) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](g_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The state of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/state-swift.property
func (g_ GameActivity) State() GameActivityState {
	rv := objc.Send[GameActivityState](g_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// Allowed characters for the party code to be used to share this activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/validPartyCodeAlphabet
func (g_ GameActivity) ValidPartyCodeAlphabet() []string {
	rv := objc.Send[[]string](g_.ID, objc.Sel("validPartyCodeAlphabet"))
	return rv
}/* debug [instance_properties/getter]: validPartyCodeAlphabet */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGameActivity */


