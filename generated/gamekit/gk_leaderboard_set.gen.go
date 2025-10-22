// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LeaderboardSet] class.
var (
	LeaderboardSetClass     _LeaderboardSetClass
	LeaderboardSetClassOnce sync.Once
)

func getLeaderboardSetClass() _LeaderboardSetClass {
	LeaderboardSetClassOnce.Do(func() {
		LeaderboardSetClass = _LeaderboardSetClass{objc.GetClass("GKLeaderboardSet")}
	})
	return LeaderboardSetClass
}

type _LeaderboardSetClass struct {
	class objc.Class
}

// An interface definition for the [LeaderboardSet] class.
type ILeaderboardSet interface {
	objectivec.IObject
	GroupIdentifier() string
	SetGroupIdentifier(value string)
	Identifier() string
	SetIdentifier(value string)
	Title() string
	SetTitle(value string)
}

// Organizes leaderboards into logical and coherent groups.
//
// A object represents a group of leaderboards that you configure in App Store Connect. For example, if your game has different worlds or levels, you can organize the leaderboards into sets for each world or level. In the Game Center dashboard, players navigate from the leaderboard sets to the individual leaderboards. If you use leaderboard sets, you must have one or more leaderboards and then place each leaderboard in a set, which can be a mix of classic and recurring leaderboards. To load all the leaderboard sets for your game, use the class method. Then use the , , and properties to access the data for each leaderboard set. If you localize the leaderboard set in App Store Connect, the property localizes. GameKit only sets the property when your game is in a game group. To load the images you add to App Store Connect for each set, use the method. Then use the method to get the leaderboards in each set. To organize leaderboards into sets, see in App Store Connect Help.


// Organizes leaderboards into logical and coherent groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardSet

type LeaderboardSet struct {
	objectivec.Object
}

// LeaderboardSetFrom constructs a [LeaderboardSet] from an unsafe.Pointer.
//
// Organizes leaderboards into logical and coherent groups.
func LeaderboardSetFrom(ptr unsafe.Pointer) LeaderboardSet {
	return LeaderboardSet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LeaderboardSetClass) Alloc() LeaderboardSet {
	rv := objc.Send[LeaderboardSet](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LeaderboardSetClass) New() LeaderboardSet {
	rv := objc.Send[LeaderboardSet](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LeaderboardSet) Init() LeaderboardSet {
	rv := objc.Send[LeaderboardSet](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LeaderboardSet) Autorelease() LeaderboardSet {
	rv := objc.Send[LeaderboardSet](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLeaderboardSet creates a new LeaderboardSet instance.
func NewLeaderboardSet() LeaderboardSet {
	return getLeaderboardSetClass().New()
}



// Loads all of the leaderboard sets you configure for your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardSet/loadLeaderboardSets(completionHandler:)

func (lc _LeaderboardSetClass) LoadLeaderboardSetsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("loadLeaderboardSetsWithCompletionHandler:"), completionHandler)
}


// The identifier for the group that the leaderboard set belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboardset/groupidentifier

func (l_ LeaderboardSet) GroupIdentifier() string {
	rv := objc.Send[string](l_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// The identifier for the group that the leaderboard set belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboardset/groupidentifier

func (l_ LeaderboardSet) SetGroupIdentifier(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGroupIdentifier:"), objc.String(value))
}


// The identifier for the leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboardset/identifier

func (l_ LeaderboardSet) Identifier() string {
	rv := objc.Send[string](l_.ID, objc.Sel("identifier"))
	return rv
}


// The identifier for the leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboardset/identifier

func (l_ LeaderboardSet) SetIdentifier(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The localized title for the leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboardset/title

func (l_ LeaderboardSet) Title() string {
	rv := objc.Send[string](l_.ID, objc.Sel("title"))
	return rv
}


// The localized title for the leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkleaderboardset/title

func (l_ LeaderboardSet) SetTitle(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTitle:"), objc.String(value))
}



