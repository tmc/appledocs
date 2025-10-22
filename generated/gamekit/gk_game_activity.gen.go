// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
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
	SetProgressOnAchievementToPercentComplete(achievement IGKAchievement, percentComplete float64)
	SetScoreOnLeaderboardToScore(leaderboard IGKLeaderboard, score int)
	Duration() foundation.TimeInterval
	Achievements() GKAchievement
	SetAchievements(value IGKAchievement)
	ActivityDefinition() GKGameActivityDefinition
	SetActivityDefinition(value IGKGameActivityDefinition)
	CreationDate() foundation.Date
	SetCreationDate(value foundation.IDate)
	EndDate() foundation.Date
	SetEndDate(value foundation.IDate)
	Identifier() string
	SetIdentifier(value string)
	LastResumeDate() foundation.Date
	SetLastResumeDate(value foundation.IDate)
	LeaderboardScores() GKLeaderboardScore
	SetLeaderboardScores(value IGKLeaderboardScore)
	PartyCode() string
	SetPartyCode(value string)
	PartyURL() foundation.URL
	SetPartyURL(value foundation.IURL)
	Properties() string
	SetProperties(value string)
	StartDate() foundation.Date
	SetStartDate(value foundation.IDate)
	State() coreml.State
	SetState(value coreml.State)
}

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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/setProgress(on:to:)

func (g_ GameActivity) SetProgressOnAchievementToPercentComplete(achievement IGKAchievement, percentComplete float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProgressOnAchievement:toPercentComplete:"), achievement, percentComplete)
}



// Set a score of a leaderboard for a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/setScore(on:to:)

func (g_ GameActivity) SetScoreOnLeaderboardToScore(leaderboard IGKLeaderboard, score int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScoreOnLeaderboard:toScore:"), leaderboard, score)
}


// Total time elapsed while in active state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/duration

func (g_ GameActivity) Duration() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](g_.ID, objc.Sel("duration"))
	return rv
}


// All achievements that have been associated with this activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/achievements

func (g_ GameActivity) Achievements() GKAchievement {
	rv := objc.Send[GKAchievement](g_.ID, objc.Sel("achievements"))
	return rv
}


// All achievements that have been associated with this activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/achievements

func (g_ GameActivity) SetAchievements(value IGKAchievement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAchievements:"), value)
}


// The activity definition that this activity instance is based on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/activitydefinition

func (g_ GameActivity) ActivityDefinition() GKGameActivityDefinition {
	rv := objc.Send[GKGameActivityDefinition](g_.ID, objc.Sel("activityDefinition"))
	return rv
}


// The activity definition that this activity instance is based on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/activitydefinition

func (g_ GameActivity) SetActivityDefinition(value IGKGameActivityDefinition) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setActivityDefinition:"), value)
}


// The date when the activity was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/creationdate

func (g_ GameActivity) CreationDate() foundation.Date {
	rv := objc.Send[foundation.Date](g_.ID, objc.Sel("creationDate"))
	return rv
}


// The date when the activity was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/creationdate

func (g_ GameActivity) SetCreationDate(value foundation.IDate) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCreationDate:"), value)
}


// The date when the activity was officially ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/enddate

func (g_ GameActivity) EndDate() foundation.Date {
	rv := objc.Send[foundation.Date](g_.ID, objc.Sel("endDate"))
	return rv
}


// The date when the activity was officially ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/enddate

func (g_ GameActivity) SetEndDate(value foundation.IDate) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEndDate:"), value)
}


// The identifier of this activity instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/identifier

func (g_ GameActivity) Identifier() string {
	rv := objc.Send[string](g_.ID, objc.Sel("identifier"))
	return rv
}


// The identifier of this activity instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/identifier

func (g_ GameActivity) SetIdentifier(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The date when the activity was last resumed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/lastresumedate

func (g_ GameActivity) LastResumeDate() foundation.Date {
	rv := objc.Send[foundation.Date](g_.ID, objc.Sel("lastResumeDate"))
	return rv
}


// The date when the activity was last resumed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/lastresumedate

func (g_ GameActivity) SetLastResumeDate(value foundation.IDate) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLastResumeDate:"), value)
}


// All leaderboard scores that have been associated with this activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/leaderboardscores

func (g_ GameActivity) LeaderboardScores() GKLeaderboardScore {
	rv := objc.Send[GKLeaderboardScore](g_.ID, objc.Sel("leaderboardScores"))
	return rv
}


// All leaderboard scores that have been associated with this activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/leaderboardscores

func (g_ GameActivity) SetLeaderboardScores(value IGKLeaderboardScore) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeaderboardScores:"), value)
}


// If the game supports party code, this is the party code that can be shared among players to join the party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partycode

func (g_ GameActivity) PartyCode() string {
	rv := objc.Send[string](g_.ID, objc.Sel("partyCode"))
	return rv
}


// If the game supports party code, this is the party code that can be shared among players to join the party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partycode

func (g_ GameActivity) SetPartyCode(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPartyCode:"), objc.String(value))
}


// If the game supports party code, this is the URL that can be shared among players to join the party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partyurl

func (g_ GameActivity) PartyURL() foundation.URL {
	rv := objc.Send[foundation.URL](g_.ID, objc.Sel("partyURL"))
	return rv
}


// If the game supports party code, this is the URL that can be shared among players to join the party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partyurl

func (g_ GameActivity) SetPartyURL(value foundation.IURL) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPartyURL:"), value)
}


// Properties that contain additional information about the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/properties

func (g_ GameActivity) Properties() string {
	rv := objc.Send[string](g_.ID, objc.Sel("properties"))
	return rv
}


// Properties that contain additional information about the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/properties

func (g_ GameActivity) SetProperties(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProperties:"), objc.String(value))
}


// The date when the activity was initially started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/startdate

func (g_ GameActivity) StartDate() foundation.Date {
	rv := objc.Send[foundation.Date](g_.ID, objc.Sel("startDate"))
	return rv
}


// The date when the activity was initially started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/startdate

func (g_ GameActivity) SetStartDate(value foundation.IDate) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStartDate:"), value)
}


// The state of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/state-swift.property

func (g_ GameActivity) State() coreml.State {
	rv := objc.Send[coreml.State](g_.ID, objc.Sel("state"))
	return rv
}


// The state of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/state-swift.property

func (g_ GameActivity) SetState(value coreml.State) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setState:"), value)
}



