// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GameActivityDefinition] class.
type IGameActivityDefinition interface {
	objectivec.IObject
	// properties:
	DefaultProperties() foundation.IDictionary
	MaxPlayers() objc.IObject /* cross-framework: NSNumber */
	Details() objc.IObject /* cross-framework: NSString */
	SetDetails(value objc.IObject /* cross-framework: NSString */)
	FallbackURL() objc.IObject /* cross-framework: URL */
	SetFallbackURL(value objc.IObject /* cross-framework: URL */)
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
	SetGroupIdentifier(value objc.IObject /* cross-framework: NSString */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	PlayStyle() GameActivityPlayStyle /* not a class type */
	SetPlayStyle(value GameActivityPlayStyle /* not a class type */)
	PlayerRange() unsafe.Pointer
	SetPlayerRange(value unsafe.Pointer)
	ReleaseState() ReleaseState /* not a class type */
	SetReleaseState(value ReleaseState /* not a class type */)
	SupportsPartyCode() bool
	SetSupportsPartyCode(value bool)
	SupportsUnlimitedPlayers() bool
	SetSupportsUnlimitedPlayers(value bool)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	// methods:
	LoadAchievementDescriptionsWithCompletionHandler(completionHandler unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (gc _GameActivityDefinitionClass) Alloc() GameActivityDefinition {
	rv := objc.Send[GameActivityDefinition](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Loads all associated achievements that have defined deep links to this game activity definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/loadAchievementDescriptions(completionHandler:)
func (g_ GameActivityDefinition) LoadAchievementDescriptionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("loadAchievementDescriptionsWithCompletionHandler:"), completionHandler)
}


// Default properties defined by the developer for this type of game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/defaultProperties
func (g_ GameActivityDefinition) DefaultProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("defaultProperties"))
	return rv
}


// The maximum number of participants that can join the activity. Returns nil when no maximum is set (unlimited players) or when player range is undefined. When not nil, the value is always greater than or equal to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/maxPlayers
func (g_ GameActivityDefinition) MaxPlayers() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("maxPlayers"))
	return rv
}


// A more detailed description of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/details
func (g_ GameActivityDefinition) Details() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("details"))
	return rv
}


// A more detailed description of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/details
func (g_ GameActivityDefinition) SetDetails(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDetails:"), value)
}


// A fallback URL that can be used to construct a game-specific URL for players to share or join, if the joining device does not support the default URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/fallbackurl
func (g_ GameActivityDefinition) FallbackURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](g_.ID, objc.Sel("fallbackURL"))
	return rv
}


// A fallback URL that can be used to construct a game-specific URL for players to share or join, if the joining device does not support the default URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/fallbackurl
func (g_ GameActivityDefinition) SetFallbackURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFallbackURL:"), value)
}


// The group identifier for the activity, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/groupidentifier
func (g_ GameActivityDefinition) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// The group identifier for the activity, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/groupidentifier
func (g_ GameActivityDefinition) SetGroupIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupIdentifier:"), value)
}


// The developer defined identifier for a given game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/identifier
func (g_ GameActivityDefinition) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("identifier"))
	return rv
}


// The developer defined identifier for a given game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/identifier
func (g_ GameActivityDefinition) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIdentifier:"), value)
}


// The play style of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playstyle
func (g_ GameActivityDefinition) PlayStyle() GameActivityPlayStyle /* not a class type */ {
	rv := objc.Send[GameActivityPlayStyle](g_.ID, objc.Sel("playStyle"))
	return rv
}


// The play style of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playstyle
func (g_ GameActivityDefinition) SetPlayStyle(value GameActivityPlayStyle /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlayStyle:"), value)
}


// The range of players supported by this type of game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playerrange
func (g_ GameActivityDefinition) PlayerRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("playerRange"))
	return rv
}


// The range of players supported by this type of game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playerrange
func (g_ GameActivityDefinition) SetPlayerRange(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlayerRange:"), value)
}


// The release state of the game activity definition in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/releasestate
func (g_ GameActivityDefinition) ReleaseState() ReleaseState /* not a class type */ {
	rv := objc.Send[ReleaseState](g_.ID, objc.Sel("releaseState"))
	return rv
}


// The release state of the game activity definition in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/releasestate
func (g_ GameActivityDefinition) SetReleaseState(value ReleaseState /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReleaseState:"), value)
}


// Whether the activity can be joined by others via a party code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/supportspartycode
func (g_ GameActivityDefinition) SupportsPartyCode() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("supportsPartyCode"))
	return rv
}


// Whether the activity can be joined by others via a party code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/supportspartycode
func (g_ GameActivityDefinition) SetSupportsPartyCode(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSupportsPartyCode:"), value)
}


// True if the activity supports an unlimited number of players. False if maxPlayers is set to a defined limit or if no player range is provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/supportsunlimitedplayers
func (g_ GameActivityDefinition) SupportsUnlimitedPlayers() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("supportsUnlimitedPlayers"))
	return rv
}


// True if the activity supports an unlimited number of players. False if maxPlayers is set to a defined limit or if no player range is provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/supportsunlimitedplayers
func (g_ GameActivityDefinition) SetSupportsUnlimitedPlayers(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSupportsUnlimitedPlayers:"), value)
}


// A short title for the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/title
func (g_ GameActivityDefinition) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("title"))
	return rv
}


// A short title for the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/title
func (g_ GameActivityDefinition) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTitle:"), value)
}



