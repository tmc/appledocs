// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	StartDate() objc.IObject /* cross-framework: NSDate */
	ActivityIdentifier() objc.IObject /* cross-framework: NSString */
	SetActivityIdentifier(value objc.IObject /* cross-framework: NSString */)
	ActivityProperties() objc.IObject /* cross-framework: NSString */
	SetActivityProperties(value objc.IObject /* cross-framework: NSString */)
	BaseLeaderboardID() objc.IObject /* cross-framework: NSString */
	SetBaseLeaderboardID(value objc.IObject /* cross-framework: NSString */)
	Duration() float64
	SetDuration(value float64)
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
	SetGroupIdentifier(value objc.IObject /* cross-framework: NSString */)
	IsHidden() bool
	SetIsHidden(value bool)
	LeaderboardDescription() objc.IObject /* cross-framework: NSString */
	SetLeaderboardDescription(value objc.IObject /* cross-framework: NSString */)
	NextStartDate() objc.IObject /* cross-framework: Date */
	SetNextStartDate(value objc.IObject /* cross-framework: Date */)
	ReleaseState() ReleaseState /* not a class type */
	SetReleaseState(value ReleaseState /* not a class type */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
	// methods:
	SubmitScoreContextPlayerCompletionHandler(score int, context uint, player IGKPlayer, completionHandler unsafe.Pointer)
}

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



// Submits a score to the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/submitScore(_:context:player:completionHandler:)
func (l_ Leaderboard) SubmitScoreContextPlayerCompletionHandler(score int, context uint, player IGKPlayer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("submitScore:context:player:completionHandler:"), score, context, player, completionHandler)
}


// The date and time a recurring leaderboard occurrence starts accepting scores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/startDate
func (l_ Leaderboard) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](l_.ID, objc.Sel("startDate"))
	return rv
}


// The identifier of the game activity associated with this leaderboard, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/activityidentifier
func (l_ Leaderboard) ActivityIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("activityIdentifier"))
	return rv
}


// The identifier of the game activity associated with this leaderboard, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/activityidentifier
func (l_ Leaderboard) SetActivityIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setActivityIdentifier:"), value)
}


// The properties when associating this leaderboard with a game activity, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/activityproperties
func (l_ Leaderboard) ActivityProperties() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("activityProperties"))
	return rv
}


// The properties when associating this leaderboard with a game activity, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/activityproperties
func (l_ Leaderboard) SetActivityProperties(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setActivityProperties:"), value)
}


// The ID that Game Center uses to identify this leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/baseleaderboardid
func (l_ Leaderboard) BaseLeaderboardID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("baseLeaderboardID"))
	return rv
}


// The ID that Game Center uses to identify this leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/baseleaderboardid
func (l_ Leaderboard) SetBaseLeaderboardID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBaseLeaderboardID:"), value)
}


// The duration from the start date that a recurring leaderboard occurrence accepts scores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/duration
func (l_ Leaderboard) Duration() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("duration"))
	return rv
}


// The duration from the start date that a recurring leaderboard occurrence accepts scores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/duration
func (l_ Leaderboard) SetDuration(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDuration:"), value)
}


// The identifier for the group the leaderboard belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/groupidentifier
func (l_ Leaderboard) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// The identifier for the group the leaderboard belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/groupidentifier
func (l_ Leaderboard) SetGroupIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGroupIdentifier:"), value)
}


// A Boolean value that indicates whether the current leaderboard isn’t visible in Game Center views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/ishidden
func (l_ Leaderboard) IsHidden() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isHidden"))
	return rv
}


// A Boolean value that indicates whether the current leaderboard isn’t visible in Game Center views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/ishidden
func (l_ Leaderboard) SetIsHidden(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsHidden:"), value)
}


// The description of this Leaderboard as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/leaderboarddescription
func (l_ Leaderboard) LeaderboardDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("leaderboardDescription"))
	return rv
}


// The description of this Leaderboard as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/leaderboarddescription
func (l_ Leaderboard) SetLeaderboardDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLeaderboardDescription:"), value)
}


// The date and time the next recurring leaderboard occurrence starts accepting scores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/nextstartdate
func (l_ Leaderboard) NextStartDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](l_.ID, objc.Sel("nextStartDate"))
	return rv
}


// The date and time the next recurring leaderboard occurrence starts accepting scores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/nextstartdate
func (l_ Leaderboard) SetNextStartDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNextStartDate:"), value)
}


// The release state of the leaderboard in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/releasestate
func (l_ Leaderboard) ReleaseState() ReleaseState /* not a class type */ {
	rv := objc.Send[ReleaseState](l_.ID, objc.Sel("releaseState"))
	return rv
}


// The release state of the leaderboard in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/releasestate
func (l_ Leaderboard) SetReleaseState(value ReleaseState /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReleaseState:"), value)
}


// The localized title for the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/title
func (l_ Leaderboard) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("title"))
	return rv
}


// The localized title for the leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/title
func (l_ Leaderboard) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTitle:"), value)
}


// The type of leaderboard, classic or recurring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/type
func (l_ Leaderboard) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("type"))
	return rv
}


// The type of leaderboard, classic or recurring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboard/type
func (l_ Leaderboard) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setType:"), value)
}



