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
	SetProgressOnAchievementToPercentComplete(achievement IGKAchievement, percentComplete unsafe.Pointer)
	SetScoreOnLeaderboardToScore(leaderboard IGKLeaderboard, score int)
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
func (g_ GameActivity) SetProgressOnAchievementToPercentComplete(achievement IGKAchievement, percentComplete unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProgressOnAchievement:toPercentComplete:"), achievement, percentComplete)
}

// Set a score of a leaderboard for a player.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/setScore(on:to:)
func (g_ GameActivity) SetScoreOnLeaderboardToScore(leaderboard IGKLeaderboard, score int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScoreOnLeaderboard:toScore:"), leaderboard, score)
}

// Total time elapsed while in active state.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/duration
func (g_ GameActivity) Duration() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](g_.ID, objc.Sel("duration"))
	return rv
}

// All achievements that have been associated with this activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/achievements
func (g_ GameActivity) Achievements() GKAchievement {
	rv := objc.Send[GKAchievement](g_.ID, objc.Sel("achievements"))
	return rv
}


// SetAchievements sets the value of the achievements property.
// All achievements that have been associated with this activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/achievements
func (g_ GameActivity) SetAchievements(value IGKAchievement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAchievements:"), value)
}

// The activity definition that this activity instance is based on.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/activitydefinition
func (g_ GameActivity) ActivityDefinition() GKGameActivityDefinition {
	rv := objc.Send[GKGameActivityDefinition](g_.ID, objc.Sel("activityDefinition"))
	return rv
}


// SetActivityDefinition sets the value of the activityDefinition property.
// The activity definition that this activity instance is based on.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/activitydefinition
func (g_ GameActivity) SetActivityDefinition(value IGKGameActivityDefinition) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setActivityDefinition:"), value)
}

// The date when the activity was created.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/creationdate
func (g_ GameActivity) CreationDate() foundation.Date {
	rv := objc.Send[foundation.Date](g_.ID, objc.Sel("creationDate"))
	return rv
}


// SetCreationDate sets the value of the creationDate property.
// The date when the activity was created.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/creationdate
func (g_ GameActivity) SetCreationDate(value foundation.IDate) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCreationDate:"), value)
}

// The date when the activity was officially ended.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/enddate
func (g_ GameActivity) EndDate() foundation.Date {
	rv := objc.Send[foundation.Date](g_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The date when the activity was officially ended.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/enddate
func (g_ GameActivity) SetEndDate(value foundation.IDate) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEndDate:"), value)
}

// The identifier of this activity instance.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/identifier
func (g_ GameActivity) Identifier() appkit.string {
	rv := objc.Send[appkit.string](g_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The identifier of this activity instance.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/identifier
func (g_ GameActivity) SetIdentifier(value appkit.string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIdentifier:"), value)
}

// The date when the activity was last resumed.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/lastresumedate
func (g_ GameActivity) LastResumeDate() foundation.Date {
	rv := objc.Send[foundation.Date](g_.ID, objc.Sel("lastResumeDate"))
	return rv
}


// SetLastResumeDate sets the value of the lastResumeDate property.
// The date when the activity was last resumed.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/lastresumedate
func (g_ GameActivity) SetLastResumeDate(value foundation.IDate) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLastResumeDate:"), value)
}

// All leaderboard scores that have been associated with this activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/leaderboardscores
func (g_ GameActivity) LeaderboardScores() GKLeaderboardScore {
	rv := objc.Send[GKLeaderboardScore](g_.ID, objc.Sel("leaderboardScores"))
	return rv
}


// SetLeaderboardScores sets the value of the leaderboardScores property.
// All leaderboard scores that have been associated with this activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/leaderboardscores
func (g_ GameActivity) SetLeaderboardScores(value IGKLeaderboardScore) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeaderboardScores:"), value)
}

// If the game supports party code, this is the party code that can be shared among players to join the party.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partycode
func (g_ GameActivity) PartyCode() appkit.string {
	rv := objc.Send[appkit.string](g_.ID, objc.Sel("partyCode"))
	return rv
}


// SetPartyCode sets the value of the partyCode property.
// If the game supports party code, this is the party code that can be shared among players to join the party.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partycode
func (g_ GameActivity) SetPartyCode(value appkit.string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPartyCode:"), value)
}

// If the game supports party code, this is the URL that can be shared among players to join the party.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partyurl
func (g_ GameActivity) PartyURL() foundation.URL {
	rv := objc.Send[foundation.URL](g_.ID, objc.Sel("partyURL"))
	return rv
}


// SetPartyURL sets the value of the partyURL property.
// If the game supports party code, this is the URL that can be shared among players to join the party.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partyurl
func (g_ GameActivity) SetPartyURL(value foundation.IURL) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPartyURL:"), value)
}

// Properties that contain additional information about the activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/properties
func (g_ GameActivity) Properties() appkit.string {
	rv := objc.Send[appkit.string](g_.ID, objc.Sel("properties"))
	return rv
}


// SetProperties sets the value of the properties property.
// Properties that contain additional information about the activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/properties
func (g_ GameActivity) SetProperties(value appkit.string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProperties:"), value)
}

// The date when the activity was initially started.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/startdate
func (g_ GameActivity) StartDate() foundation.Date {
	rv := objc.Send[foundation.Date](g_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The date when the activity was initially started.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/startdate
func (g_ GameActivity) SetStartDate(value foundation.IDate) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStartDate:"), value)
}

// The state of the game activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/state-swift.property
func (g_ GameActivity) State() coreml.State {
	rv := objc.Send[coreml.State](g_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The state of the game activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/state-swift.property
func (g_ GameActivity) SetState(value coreml.State) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setState:"), value)
}



