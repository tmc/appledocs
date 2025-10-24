// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKAccessPoint */


/* debug [class_header]: Header for GKAccessPoint */
// The class instance for the [AccessPoint] class.
var (
	AccessPointClass     _AccessPointClass
	AccessPointClassOnce sync.Once
)

func getAccessPointClass() _AccessPointClass {
	AccessPointClassOnce.Do(func() {
		AccessPointClass = _AccessPointClass{objc.GetClass("GKAccessPoint")}
	})
	return AccessPointClass
}

type _AccessPointClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccessPoint */
// An interface definition for the [AccessPoint] class.
type IAccessPoint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccessPoint */
	// properties:
	FrameInScreenCoordinates() Rect /* not a class type */
	Active() bool
	SetActive(value bool)
	IsPresentingGameCenter() bool
	Visible() bool
	Location() AccessPointLocation
	SetLocation(value AccessPointLocation)
	ParentWindow() appkit.Window
	SetParentWindow(value appkit.Window)
	ShowHighlights() bool
	SetShowHighlights(value bool)
	IsActive() bool
	SetIsActive(value bool)
	IsFocused() bool
	SetIsFocused(value bool)
	IsVisible() bool
	SetIsVisible(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccessPoint */
	// methods:
	TriggerAccessPointWithAchievementIDHandler(achievementID objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	TriggerAccessPointWithChallengeDefinitionIDHandler(challengeDefinitionID objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	Trigger()
	TriggerAccessPointWithGameActivityHandler(gameActivity IGKGameActivity, handler unsafe.Pointer)
	TriggerAccessPointWithGameActivityDefinitionIDHandler(gameActivityDefinitionID objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	TriggerAccessPointWithHandler(handler unsafe.Pointer)
	TriggerAccessPointWithLeaderboardIDPlayerScopeTimeScopeHandler(leaderboardID objc.IObject /* cross-framework: NSString */, playerScope LeaderboardPlayerScope, timeScope LeaderboardTimeScope, handler unsafe.Pointer)
	TriggerAccessPointWithLeaderboardSetIDHandler(leaderboardSetID objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	TriggerAccessPointWithPlayerHandler(player IGKPlayer, handler unsafe.Pointer)
	TriggerAccessPointWithStateHandler(state GameCenterViewControllerState, handler unsafe.Pointer)
	TriggerAccessPointForChallengesWithHandler(handler unsafe.Pointer)
	TriggerAccessPointForFriendingWithHandler(handler unsafe.Pointer)
	TriggerAccessPointForPlayTogetherWithHandler(handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccessPoint */
// Alloc allocates a new instance without initialization.
func (ac _AccessPointClass) Alloc() AccessPoint {
	rv := objc.Send[AccessPoint](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccessPointClass) New() AccessPoint {
	rv := objc.Send[AccessPoint](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccessPoint) Init() AccessPoint {
	rv := objc.Send[AccessPoint](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccessPoint) Autorelease() AccessPoint {
	rv := objc.Send[AccessPoint](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccessPoint creates a new AccessPoint instance.
func NewAccessPoint() AccessPoint {
	return getAccessPointClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccessPoint */
// An object that allows players to view and manage their Game Center information from within your game.
//
// The access point displays a control in a corner of your game that opens a Game Center dashboard when the player taps or clicks it. Use the property to get the shared access point object. GameKit attaches the access point to the window you specify in the property, in the corner you specify using the property. If you don’t specify a parent window, GameKit infers an appropriate location. For the location of the access point on visionOS, see . To display highlights, set the property to . Then set to to display the access point control.


// An object that allows players to view and manage their Game Center information from within your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint
type AccessPoint struct {
	objectivec.Object
}

// AccessPointFrom constructs a [AccessPoint] from an unsafe.Pointer.
//
// An object that allows players to view and manage their Game Center information from within your game.
func AccessPointFrom(ptr unsafe.Pointer) AccessPoint {
	return AccessPoint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccessPoint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccessPoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccessPoint */

// The shared access point object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/shared
func (ac _AccessPointClass) Shared() AccessPoint {
	rv := objc.Send[AccessPoint](objc.ID(ac.class), objc.Sel("shared"))
	return rv
}/* debug [class_properties_class/property]: shared */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccessPoint */

// Displays the Game Center dashboard in a state that shows a specific achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(achievementID:handler:)
func (a_ AccessPoint) TriggerAccessPointWithAchievementIDHandler(achievementID objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointWithAchievementID:handler:"), achievementID, handler)
}/* debug [instance_methods/method]: TriggerAccessPointWithAchievementIDHandler */


// Displays the challenge creation view for the provided challenge definition ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(challengeDefinitionID:handler:)
func (a_ AccessPoint) TriggerAccessPointWithChallengeDefinitionIDHandler(challengeDefinitionID objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointWithChallengeDefinitionID:handler:"), challengeDefinitionID, handler)
}/* debug [instance_methods/method]: TriggerAccessPointWithChallengeDefinitionIDHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(gameActivity:handler:)-6lnz8
func (a_ AccessPoint) Trigger() {
	objc.Send[objc.ID](a_.ID, objc.Sel("trigger"))
}/* debug [instance_methods/method]: Trigger */


// Displays the game activity view for the provided activity instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(gameActivity:handler:)-8i6w7
func (a_ AccessPoint) TriggerAccessPointWithGameActivityHandler(gameActivity IGKGameActivity, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointWithGameActivity:handler:"), gameActivity, handler)
}/* debug [instance_methods/method]: TriggerAccessPointWithGameActivityHandler */


// Displays the game activity creation view for the provided activity definition ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(gameActivityDefinitionID:handler:)-9hemd
func (a_ AccessPoint) TriggerAccessPointWithGameActivityDefinitionIDHandler(gameActivityDefinitionID objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointWithGameActivityDefinitionID:handler:"), gameActivityDefinitionID, handler)
}/* debug [instance_methods/method]: TriggerAccessPointWithGameActivityDefinitionIDHandler */


// Displays the Game Center dashboard as if the player taps or presses the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(handler:)
func (a_ AccessPoint) TriggerAccessPointWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointWithHandler:"), handler)
}/* debug [instance_methods/method]: TriggerAccessPointWithHandler */


// Displays the Game Center dashboard in a state that shows a specific leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(leaderboardID:playerScope:timeScope:handler:)
func (a_ AccessPoint) TriggerAccessPointWithLeaderboardIDPlayerScopeTimeScopeHandler(leaderboardID objc.IObject /* cross-framework: NSString */, playerScope LeaderboardPlayerScope, timeScope LeaderboardTimeScope, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointWithLeaderboardID:playerScope:timeScope:handler:"), leaderboardID, playerScope, timeScope, handler)
}/* debug [instance_methods/method]: TriggerAccessPointWithLeaderboardIDPlayerScopeTimeScopeHandler */


// Displays the Game Center dashboard in a state that shows a specific leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(leaderboardSetID:handler:)
func (a_ AccessPoint) TriggerAccessPointWithLeaderboardSetIDHandler(leaderboardSetID objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointWithLeaderboardSetID:handler:"), leaderboardSetID, handler)
}/* debug [instance_methods/method]: TriggerAccessPointWithLeaderboardSetIDHandler */


// Displays the Game Center dashboard in a state that shows a player profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(player:handler:)
func (a_ AccessPoint) TriggerAccessPointWithPlayerHandler(player IGKPlayer, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointWithPlayer:handler:"), player, handler)
}/* debug [instance_methods/method]: TriggerAccessPointWithPlayerHandler */


// Displays the Game Center dashboard in the specified state as if the player taps or presses the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/trigger(state:handler:)
func (a_ AccessPoint) TriggerAccessPointWithStateHandler(state GameCenterViewControllerState, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointWithState:handler:"), state, handler)
}/* debug [instance_methods/method]: TriggerAccessPointWithStateHandler */


// Displays the view that allows players to engage each other with challenges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/triggerForChallenges(handler:)
func (a_ AccessPoint) TriggerAccessPointForChallengesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointForChallengesWithHandler:"), handler)
}/* debug [instance_methods/method]: TriggerAccessPointForChallengesWithHandler */


// Brings up the invite friends view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/triggerForFriending(handler:)
func (a_ AccessPoint) TriggerAccessPointForFriendingWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointForFriendingWithHandler:"), handler)
}/* debug [instance_methods/method]: TriggerAccessPointForFriendingWithHandler */


// Displays the view that allows players to engage each other with activities and challenges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/triggerForPlayTogether(handler:)
func (a_ AccessPoint) TriggerAccessPointForPlayTogetherWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointForPlayTogetherWithHandler:"), handler)
}/* debug [instance_methods/method]: TriggerAccessPointForPlayTogetherWithHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccessPoint */

// The frame of the access point in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/frameInScreenCoordinates
func (a_ AccessPoint) FrameInScreenCoordinates() Rect /* not a class type */ {
	rv := objc.Send[Rect](a_.ID, objc.Sel("frameInScreenCoordinates"))
	return rv
}/* debug [instance_properties/getter]: frameInScreenCoordinates */


// A Boolean value that determines whether to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/isActive
func (a_ AccessPoint) Active() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A Boolean value that determines whether to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/isActive
func (a_ AccessPoint) SetActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setActive:"), value)
}/* debug [instance_properties/setter]: active */


// A Boolean value that indicates whether the game is presenting the Game Center dashboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/isPresentingGameCenter
func (a_ AccessPoint) IsPresentingGameCenter() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPresentingGameCenter"))
	return rv
}/* debug [instance_properties/getter]: isPresentingGameCenter */


// A Boolean value that indicates whether the access point is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/isVisible
func (a_ AccessPoint) Visible() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("visible"))
	return rv
}/* debug [instance_properties/getter]: visible */


// The corner of the screen to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/location-swift.property
func (a_ AccessPoint) Location() AccessPointLocation {
	rv := objc.Send[AccessPointLocation](a_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// The corner of the screen to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/location-swift.property
func (a_ AccessPoint) SetLocation(value AccessPointLocation) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocation:"), value)
}/* debug [instance_properties/setter]: location */


// The window that contains the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/parentWindow
func (a_ AccessPoint) ParentWindow() appkit.Window {
	rv := objc.Send[appkit.Window](a_.ID, objc.Sel("parentWindow"))
	return rv
}/* debug [instance_properties/getter]: parentWindow */


// The window that contains the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/parentWindow
func (a_ AccessPoint) SetParentWindow(value appkit.Window) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParentWindow:"), value)
}/* debug [instance_properties/setter]: parentWindow */


// The shared access point object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/shared
func (a_ AccessPoint) Shared() IGKAccessPoint {
	rv := objc.Send[AccessPoint](a_.ID, objc.Sel("shared"))
	return rv
}/* debug [instance_properties/getter]: shared */


// A Boolean value that indicates whether to display highlights for achievements and current ranks for leaderboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/showHighlights
func (a_ AccessPoint) ShowHighlights() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("showHighlights"))
	return rv
}/* debug [instance_properties/getter]: showHighlights */


// A Boolean value that indicates whether to display highlights for achievements and current ranks for leaderboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/showHighlights
func (a_ AccessPoint) SetShowHighlights(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShowHighlights:"), value)
}/* debug [instance_properties/setter]: showHighlights */


// A Boolean value that determines whether to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isactive
func (a_ AccessPoint) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean value that determines whether to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isactive
func (a_ AccessPoint) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */


// A Boolean value that indicates whether the access point is in focus on tvOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isfocused
func (a_ AccessPoint) IsFocused() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isFocused"))
	return rv
}/* debug [instance_properties/getter]: isFocused */


// A Boolean value that indicates whether the access point is in focus on tvOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isfocused
func (a_ AccessPoint) SetIsFocused(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsFocused:"), value)
}/* debug [instance_properties/setter]: isFocused */


// A Boolean value that indicates whether the access point is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isvisible
func (a_ AccessPoint) IsVisible() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isVisible"))
	return rv
}/* debug [instance_properties/getter]: isVisible */


// A Boolean value that indicates whether the access point is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isvisible
func (a_ AccessPoint) SetIsVisible(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsVisible:"), value)
}/* debug [instance_properties/setter]: isVisible */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKAccessPoint */


