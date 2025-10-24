// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKLeaderboardSet */


/* debug [class_header]: Header for GKLeaderboardSet */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LeaderboardSet */
// An interface definition for the [LeaderboardSet] class.
type ILeaderboardSet interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LeaderboardSet */
	// properties:
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LeaderboardSet */
	// methods:
	LoadImageWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadLeaderboardsWithHandler(handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LeaderboardSet */
// Alloc allocates a new instance without initialization.
func (lc _LeaderboardSetClass) Alloc() LeaderboardSet {
	rv := objc.Send[LeaderboardSet](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LeaderboardSet */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LeaderboardSet *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LeaderboardSet */

// Loads all of the leaderboard sets you configure for your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardSet/loadLeaderboardSets(completionHandler:)
func (lc _LeaderboardSetClass) LoadLeaderboardSetsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("loadLeaderboardSetsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadLeaderboardSetsWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LeaderboardSet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LeaderboardSet */

// Loads the localized image that you associate with the leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardSet/loadImage(completionHandler:)
func (l_ LeaderboardSet) LoadImageWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadImageWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadImageWithCompletionHandler */


// Loads the leaderboards in the leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardSet/loadLeaderboards(handler:)
func (l_ LeaderboardSet) LoadLeaderboardsWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadLeaderboardsWithHandler:"), handler)
}/* debug [instance_methods/method]: LoadLeaderboardsWithHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LeaderboardSet */

// The identifier for the group that the leaderboard set belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardSet/groupIdentifier
func (l_ LeaderboardSet) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("groupIdentifier"))
	return rv
}/* debug [instance_properties/getter]: groupIdentifier */


// The identifier for the leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardSet/identifier
func (l_ LeaderboardSet) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The identifier for the leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardSet/identifier
func (l_ LeaderboardSet) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// The localized title for the leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardSet/title
func (l_ LeaderboardSet) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKLeaderboardSet */



