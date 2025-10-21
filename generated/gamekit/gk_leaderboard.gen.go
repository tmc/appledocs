// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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


// The identifier of the game activity associated with this leaderboard, as configured by the developer in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/activityidentifier
func (l_ Leaderboard) ActivityIdentifier() string {
	rv := objc.Send[string](l_.ID, objc.Sel("activityIdentifier"))
	return rv
}


// SetActivityIdentifier sets the value of the activityIdentifier property.
// The identifier of the game activity associated with this leaderboard, as configured by the developer in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/activityidentifier
func (l_ Leaderboard) SetActivityIdentifier(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setActivityIdentifier:"), objc.String(value))
}

// The description of this Leaderboard as configured by the developer in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/leaderboarddescription
func (l_ Leaderboard) LeaderboardDescription() string {
	rv := objc.Send[string](l_.ID, objc.Sel("leaderboardDescription"))
	return rv
}


// SetLeaderboardDescription sets the value of the leaderboardDescription property.
// The description of this Leaderboard as configured by the developer in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/leaderboarddescription
func (l_ Leaderboard) SetLeaderboardDescription(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLeaderboardDescription:"), objc.String(value))
}

// The type of leaderboard, classic or recurring.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/type
func (l_ Leaderboard) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of leaderboard, classic or recurring.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/type
func (l_ Leaderboard) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setType:"), value)
}

// The release state of the leaderboard in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/releasestate
func (l_ Leaderboard) ReleaseState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("releaseState"))
	return rv
}


// SetReleaseState sets the value of the releaseState property.
// The release state of the leaderboard in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/releasestate
func (l_ Leaderboard) SetReleaseState(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReleaseState:"), value)
}

// The properties when associating this leaderboard with a game activity, as configured by the developer in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/activityproperties
func (l_ Leaderboard) ActivityProperties() string {
	rv := objc.Send[string](l_.ID, objc.Sel("activityProperties"))
	return rv
}


// SetActivityProperties sets the value of the activityProperties property.
// The properties when associating this leaderboard with a game activity, as configured by the developer in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/activityproperties
func (l_ Leaderboard) SetActivityProperties(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setActivityProperties:"), objc.String(value))
}

// The identifier for the group the leaderboard belongs to.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/groupidentifier
func (l_ Leaderboard) GroupIdentifier() string {
	rv := objc.Send[string](l_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// SetGroupIdentifier sets the value of the groupIdentifier property.
// The identifier for the group the leaderboard belongs to.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/groupidentifier
func (l_ Leaderboard) SetGroupIdentifier(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGroupIdentifier:"), objc.String(value))
}

// The date and time the next recurring leaderboard occurrence starts accepting scores.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/nextstartdate
func (l_ Leaderboard) NextStartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("nextStartDate"))
	return rv
}


// SetNextStartDate sets the value of the nextStartDate property.
// The date and time the next recurring leaderboard occurrence starts accepting scores.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/nextstartdate
func (l_ Leaderboard) SetNextStartDate(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNextStartDate:"), value)
}

// A Boolean value that indicates whether the current leaderboard isn’t visible in Game Center views.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/ishidden
func (l_ Leaderboard) IsHidden() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
// A Boolean value that indicates whether the current leaderboard isn’t visible in Game Center views.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/ishidden
func (l_ Leaderboard) SetIsHidden(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsHidden:"), value)
}

// The duration from the start date that a recurring leaderboard occurrence accepts scores.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/duration
func (l_ Leaderboard) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// The duration from the start date that a recurring leaderboard occurrence accepts scores.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/duration
func (l_ Leaderboard) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDuration:"), value)
}

// The localized title for the leaderboard.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/title
func (l_ Leaderboard) Title() string {
	rv := objc.Send[string](l_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The localized title for the leaderboard.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/title
func (l_ Leaderboard) SetTitle(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The ID that Game Center uses to identify this leaderboard.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/baseleaderboardid
func (l_ Leaderboard) BaseLeaderboardID() string {
	rv := objc.Send[string](l_.ID, objc.Sel("baseLeaderboardID"))
	return rv
}


// SetBaseLeaderboardID sets the value of the baseLeaderboardID property.
// The ID that Game Center uses to identify this leaderboard.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/baseleaderboardid
func (l_ Leaderboard) SetBaseLeaderboardID(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBaseLeaderboardID:"), objc.String(value))
}

// The date and time a recurring leaderboard occurrence starts accepting scores.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/startDate
func (l_ Leaderboard) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("startDate"))
	return rv
}



