// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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

// The identifier of this activity instance.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/identifier
func (g_ GameActivity) Identifier() string {
	rv := objc.Send[string](g_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The identifier of this activity instance.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/identifier
func (g_ GameActivity) SetIdentifier(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// The date when the activity was initially started.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/startdate
func (g_ GameActivity) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The date when the activity was initially started.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/startdate
func (g_ GameActivity) SetStartDate(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStartDate:"), value)
}

// If the game supports party code, this is the party code that can be shared among players to join the party.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partycode
func (g_ GameActivity) PartyCode() string {
	rv := objc.Send[string](g_.ID, objc.Sel("partyCode"))
	return rv
}


// SetPartyCode sets the value of the partyCode property.
// If the game supports party code, this is the party code that can be shared among players to join the party.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/partycode
func (g_ GameActivity) SetPartyCode(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPartyCode:"), objc.String(value))
}

// The date when the activity was created.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/creationdate
func (g_ GameActivity) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("creationDate"))
	return rv
}


// SetCreationDate sets the value of the creationDate property.
// The date when the activity was created.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/creationdate
func (g_ GameActivity) SetCreationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCreationDate:"), value)
}

// Properties that contain additional information about the activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/properties
func (g_ GameActivity) Properties() string {
	rv := objc.Send[string](g_.ID, objc.Sel("properties"))
	return rv
}


// SetProperties sets the value of the properties property.
// Properties that contain additional information about the activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/properties
func (g_ GameActivity) SetProperties(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProperties:"), objc.String(value))
}

// All leaderboard scores that have been associated with this activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/leaderboardscores
func (g_ GameActivity) LeaderboardScores() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("leaderboardScores"))
	return rv
}


// SetLeaderboardScores sets the value of the leaderboardScores property.
// All leaderboard scores that have been associated with this activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/leaderboardscores
func (g_ GameActivity) SetLeaderboardScores(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeaderboardScores:"), value)
}

// All achievements that have been associated with this activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/achievements
func (g_ GameActivity) Achievements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("achievements"))
	return rv
}


// SetAchievements sets the value of the achievements property.
// All achievements that have been associated with this activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/achievements
func (g_ GameActivity) SetAchievements(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAchievements:"), value)
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
func (g_ GameActivity) SetPartyURL(value foundation.URL) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPartyURL:"), value)
}

// The date when the activity was last resumed.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/lastresumedate
func (g_ GameActivity) LastResumeDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("lastResumeDate"))
	return rv
}


// SetLastResumeDate sets the value of the lastResumeDate property.
// The date when the activity was last resumed.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/lastresumedate
func (g_ GameActivity) SetLastResumeDate(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLastResumeDate:"), value)
}

// The date when the activity was officially ended.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/enddate
func (g_ GameActivity) EndDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The date when the activity was officially ended.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/enddate
func (g_ GameActivity) SetEndDate(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEndDate:"), value)
}

// The activity definition that this activity instance is based on.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/activitydefinition
func (g_ GameActivity) ActivityDefinition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("activityDefinition"))
	return rv
}


// SetActivityDefinition sets the value of the activityDefinition property.
// The activity definition that this activity instance is based on.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/activitydefinition
func (g_ GameActivity) SetActivityDefinition(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setActivityDefinition:"), value)
}

// The state of the game activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/state-swift.property
func (g_ GameActivity) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The state of the game activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivity/state-swift.property
func (g_ GameActivity) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setState:"), value)
}

// Total time elapsed while in active state.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/duration
func (g_ GameActivity) Duration() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](g_.ID, objc.Sel("duration"))
	return rv
}



