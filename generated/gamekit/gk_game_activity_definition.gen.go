// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKGameActivityDefinition */


/* debug [class_header]: Header for GKGameActivityDefinition */
// The class instance for the [GameActivityDefinition] class.
var (
	GameActivityDefinitionClass     _GameActivityDefinitionClass
	GameActivityDefinitionClassOnce sync.Once
)

func getGameActivityDefinitionClass() _GameActivityDefinitionClass {
	GameActivityDefinitionClassOnce.Do(func() {
		GameActivityDefinitionClass = _GameActivityDefinitionClass{objc.GetClass("GKGameActivityDefinition")}
	})
	return GameActivityDefinitionClass
}

type _GameActivityDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GameActivityDefinition */
// An interface definition for the [GameActivityDefinition] class.
type IGameActivityDefinition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GameActivityDefinition */
	// properties:
	DefaultProperties() foundation.IDictionary
	Details() objc.IObject /* cross-framework: NSString */
	FallbackURL() objc.IObject /* cross-framework: NSURL */
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	MaxPlayers() objc.IObject /* cross-framework: NSNumber */
	MinPlayers() objc.IObject /* cross-framework: NSNumber */
	PlayStyle() GameActivityPlayStyle
	ReleaseState() ReleaseState
	SupportsPartyCode() bool
	SupportsUnlimitedPlayers() bool
	Title() objc.IObject /* cross-framework: NSString */
	PlayerRange() unsafe.Pointer
	SetPlayerRange(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GameActivityDefinition */
	// methods:
	LoadAchievementDescriptionsWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadImageWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadLeaderboardsWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GameActivityDefinition */
// Alloc allocates a new instance without initialization.
func (gc _GameActivityDefinitionClass) Alloc() GameActivityDefinition {
	rv := objc.Send[GameActivityDefinition](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GameActivityDefinitionClass) New() GameActivityDefinition {
	rv := objc.Send[GameActivityDefinition](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GameActivityDefinition) Init() GameActivityDefinition {
	rv := objc.Send[GameActivityDefinition](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GameActivityDefinition) Autorelease() GameActivityDefinition {
	rv := objc.Send[GameActivityDefinition](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGameActivityDefinition creates a new GameActivityDefinition instance.
func NewGameActivityDefinition() GameActivityDefinition {
	return getGameActivityDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GameActivityDefinition */
// An object that represents the static metadata you define for the activity.


// An object that represents the static metadata you define for the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition
type GameActivityDefinition struct {
	objectivec.Object
}

// GameActivityDefinitionFrom constructs a [GameActivityDefinition] from an unsafe.Pointer.
//
// An object that represents the static metadata you define for the activity.
func GameActivityDefinitionFrom(ptr unsafe.Pointer) GameActivityDefinition {
	return GameActivityDefinition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GameActivityDefinition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GameActivityDefinition */

// Loads all the game activity definitions for the current game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/loadGameActivityDefinitions(completionHandler:)
func (gc _GameActivityDefinitionClass) LoadGameActivityDefinitionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("loadGameActivityDefinitionsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadGameActivityDefinitionsWithCompletionHandler) */


// Loads game activity definitions with the supplied App Store Connect identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/loadGameActivityDefinitions(IDs:completionHandler:)
func (gc _GameActivityDefinitionClass) LoadGameActivityDefinitionsWithIDsCompletionHandler(activityDefinitionIDs []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("loadGameActivityDefinitionsWithIDs:completionHandler:"), activityDefinitionIDs, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadGameActivityDefinitionsWithIDsCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GameActivityDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GameActivityDefinition */

// Loads all associated achievements that have defined deep links to this game activity definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/loadAchievementDescriptions(completionHandler:)
func (g_ GameActivityDefinition) LoadAchievementDescriptionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("loadAchievementDescriptionsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadAchievementDescriptionsWithCompletionHandler */


// Asynchronously load the image. Error will be nil on success.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/loadImage(completionHandler:)
func (g_ GameActivityDefinition) LoadImageWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("loadImageWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadImageWithCompletionHandler */


// Loads all associated leaderboards that have defined deep links to this game activity definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/loadLeaderboards(completionHandler:)
func (g_ GameActivityDefinition) LoadLeaderboardsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("loadLeaderboardsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadLeaderboardsWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GameActivityDefinition */

// Default properties defined by the developer for this type of game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/defaultProperties
func (g_ GameActivityDefinition) DefaultProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("defaultProperties"))
	return rv
}/* debug [instance_properties/getter]: defaultProperties */


// A more detailed description of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/details
func (g_ GameActivityDefinition) Details() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("details"))
	return rv
}/* debug [instance_properties/getter]: details */


// A fallback URL that can be used to construct a game-specific URL for players to share or join, if the joining device does not support the default URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/fallbackURL
func (g_ GameActivityDefinition) FallbackURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](g_.ID, objc.Sel("fallbackURL"))
	return rv
}/* debug [instance_properties/getter]: fallbackURL */


// The group identifier for the activity, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/groupIdentifier
func (g_ GameActivityDefinition) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("groupIdentifier"))
	return rv
}/* debug [instance_properties/getter]: groupIdentifier */


// The developer defined identifier for a given game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/identifier
func (g_ GameActivityDefinition) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The maximum number of participants that can join the activity. Returns nil when no maximum is set (unlimited players) or when player range is undefined. When not nil, the value is always greater than or equal to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/maxPlayers
func (g_ GameActivityDefinition) MaxPlayers() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("maxPlayers"))
	return rv
}/* debug [instance_properties/getter]: maxPlayers */


// The minimum number of participants that can join the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/minPlayers
func (g_ GameActivityDefinition) MinPlayers() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("minPlayers"))
	return rv
}/* debug [instance_properties/getter]: minPlayers */


// The play style of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/playStyle
func (g_ GameActivityDefinition) PlayStyle() GameActivityPlayStyle {
	rv := objc.Send[GameActivityPlayStyle](g_.ID, objc.Sel("playStyle"))
	return rv
}/* debug [instance_properties/getter]: playStyle */


// The release state of the game activity definition in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/releaseState
func (g_ GameActivityDefinition) ReleaseState() ReleaseState {
	rv := objc.Send[ReleaseState](g_.ID, objc.Sel("releaseState"))
	return rv
}/* debug [instance_properties/getter]: releaseState */


// Whether the activity can be joined by others via a party code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/supportsPartyCode
func (g_ GameActivityDefinition) SupportsPartyCode() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("supportsPartyCode"))
	return rv
}/* debug [instance_properties/getter]: supportsPartyCode */


// True if the activity supports an unlimited number of players. False if maxPlayers is set to a defined limit or if no player range is provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/supportsUnlimitedPlayers
func (g_ GameActivityDefinition) SupportsUnlimitedPlayers() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("supportsUnlimitedPlayers"))
	return rv
}/* debug [instance_properties/getter]: supportsUnlimitedPlayers */


// A short title for the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/title
func (g_ GameActivityDefinition) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The range of players supported by this type of game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playerrange
func (g_ GameActivityDefinition) PlayerRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("playerRange"))
	return rv
}/* debug [instance_properties/getter]: playerRange */


// The range of players supported by this type of game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playerrange
func (g_ GameActivityDefinition) SetPlayerRange(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlayerRange:"), value)
}/* debug [instance_properties/setter]: playerRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGameActivityDefinition */



