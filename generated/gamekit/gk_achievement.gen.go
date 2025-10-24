// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKAchievement */


/* debug [class_header]: Header for GKAchievement */
// The class instance for the [Achievement] class.
var (
	AchievementClass     _AchievementClass
	AchievementClassOnce sync.Once
)

func getAchievementClass() _AchievementClass {
	AchievementClassOnce.Do(func() {
		AchievementClass = _AchievementClass{objc.GetClass("GKAchievement")}
	})
	return AchievementClass
}

type _AchievementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Achievement */
// An interface definition for the [Achievement] class.
type IAchievement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Achievement */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	Completed() bool
	Hidden() bool
	LastReportedDate() objc.IObject /* cross-framework: NSDate */
	PercentComplete() float64
	SetPercentComplete(value float64)
	Player() IGKPlayer
	PlayerID() objc.IObject /* cross-framework: NSString */
	ShowsCompletionBanner() bool
	SetShowsCompletionBanner(value bool)
	IsCompleted() bool
	SetIsCompleted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Achievement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Achievement */
// Alloc allocates a new instance without initialization.
func (ac _AchievementClass) Alloc() Achievement {
	rv := objc.Send[Achievement](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AchievementClass) New() Achievement {
	rv := objc.Send[Achievement](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Achievement) Init() Achievement {
	rv := objc.Send[Achievement](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Achievement) Autorelease() Achievement {
	rv := objc.Send[Achievement](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAchievement creates a new Achievement instance.
func NewAchievement() Achievement {
	return getAchievementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Achievement */
// An achievement you can award a player as they make progress toward and reach a goal in your game.
//
// Before using this class, configure your game achievements in App Store Connect. Then the dashboard shows the achievements initially locked and you can access them in your code. Use the method to load all the achievements that the local player is progressing toward. If an achievement doesn’t load, then it’s the first time you’re reporting the player’s progress toward it, and you must create a object to represent it. Next, set the percentage complete of the achievement using the property. You can report the player’s progress for one or more achievements to Game Center using the method. The dashboard changes the appearance of the achievements to show the current percentages. If you set the percentage of an achievement to 100, the dashboard shows it as completed. To reset the player’s progress on all achievements, use the class method.


// An achievement you can award a player as they make progress toward and reach a goal in your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement
type Achievement struct {
	objectivec.Object
}

// AchievementFrom constructs a [Achievement] from an unsafe.Pointer.
//
// An achievement you can award a player as they make progress toward and reach a goal in your game.
func AchievementFrom(ptr unsafe.Pointer) Achievement {
	return Achievement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Achievement */

// Initializes an achievement for the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/init(identifier:)
func NewAchievementWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) Achievement {
	instance := getAchievementClass().Alloc()
	rv := objc.Send[Achievement](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAchievementWithIdentifier */


// Initializes an achievement for a specific player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/init(identifier:forPlayer:)
func NewAchievementWithIdentifierForPlayer(identifier objc.IObject /* cross-framework: NSString */, playerID objc.IObject /* cross-framework: NSString */) Achievement {
	instance := getAchievementClass().Alloc()
	rv := objc.Send[Achievement](instance.ID, objc.Sel("initWithIdentifier:forPlayer:"), identifier, playerID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAchievementWithIdentifierForPlayer */


// Initializes an achievement for a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/init(identifier:player:)
func NewAchievementWithIdentifierPlayer(identifier objc.IObject /* cross-framework: NSString */, player IGKPlayer) Achievement {
	instance := getAchievementClass().Alloc()
	rv := objc.Send[Achievement](instance.ID, objc.Sel("initWithIdentifier:player:"), identifier, player)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAchievementWithIdentifierPlayer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Achievement */

// Loads the achievements that you previously reported the player making progress toward.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/loadAchievements(completionHandler:)
func (ac _AchievementClass) LoadAchievementsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("loadAchievementsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadAchievementsWithCompletionHandler) */


// Reports the player’s progress of players toward one or more achievements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/report(_:withCompletionHandler:)
func (ac _AchievementClass) ReportAchievementsWithCompletionHandler(achievements []Achievement, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("reportAchievements:withCompletionHandler:"), achievements, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReportAchievementsWithCompletionHandler) */


// Reports the player’s progress on achievements and limits the challenges, associated with those achievements, that the player may complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/report(_:withEligibleChallenges:withCompletionHandler:)
func (ac _AchievementClass) ReportAchievementsWithEligibleChallengesWithCompletionHandler(achievements []Achievement, challenges []Challenge, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("reportAchievements:withEligibleChallenges:withCompletionHandler:"), achievements, challenges, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReportAchievementsWithEligibleChallengesWithCompletionHandler) */


// Resets the percentage completed for all of the player’s achievements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/resetAchievements(completionHandler:)
func (ac _AchievementClass) ResetAchievementsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("resetAchievementsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ResetAchievementsWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Achievement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Achievement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Achievement */

// The identifier for the achievement that you enter in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/identifier
func (a_ Achievement) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The identifier for the achievement that you enter in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/identifier
func (a_ Achievement) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// A Boolean value that states whether the player has completed the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/isCompleted
func (a_ Achievement) Completed() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("completed"))
	return rv
}/* debug [instance_properties/getter]: completed */


// A Boolean value that indicates whether the system hides this achievement from the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/isHidden
func (a_ Achievement) Hidden() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// The last time your game reported progress on the achievement for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/lastReportedDate
func (a_ Achievement) LastReportedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](a_.ID, objc.Sel("lastReportedDate"))
	return rv
}/* debug [instance_properties/getter]: lastReportedDate */


// A percentage value that states how far the player has progressed on the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/percentComplete
func (a_ Achievement) PercentComplete() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("percentComplete"))
	return rv
}/* debug [instance_properties/getter]: percentComplete */


// A percentage value that states how far the player has progressed on the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/percentComplete
func (a_ Achievement) SetPercentComplete(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPercentComplete:"), value)
}/* debug [instance_properties/setter]: percentComplete */


// The player who earned the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/player
func (a_ Achievement) Player() IGKPlayer {
	rv := objc.Send[Player](a_.ID, objc.Sel("player"))
	return rv
}/* debug [instance_properties/getter]: player */


// A string that identifies the player who earned the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/playerID
func (a_ Achievement) PlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("playerID"))
	return rv
}/* debug [instance_properties/getter]: playerID */


// A Boolean value that indicates whether GameKit displays a banner when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/showsCompletionBanner
func (a_ Achievement) ShowsCompletionBanner() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("showsCompletionBanner"))
	return rv
}/* debug [instance_properties/getter]: showsCompletionBanner */


// A Boolean value that indicates whether GameKit displays a banner when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/showsCompletionBanner
func (a_ Achievement) SetShowsCompletionBanner(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShowsCompletionBanner:"), value)
}/* debug [instance_properties/setter]: showsCompletionBanner */


// A Boolean value that states whether the player has completed the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievement/iscompleted
func (a_ Achievement) IsCompleted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompleted"))
	return rv
}/* debug [instance_properties/getter]: isCompleted */


// A Boolean value that states whether the player has completed the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievement/iscompleted
func (a_ Achievement) SetIsCompleted(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCompleted:"), value)
}/* debug [instance_properties/setter]: isCompleted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKAchievement */


