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
}

// An object that represents the static metadata you define for the activity.
//
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


// The range of players supported by this type of game activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playerrange
func (g_ GameActivityDefinition) PlayerRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("playerRange"))
	return rv
}


// SetPlayerRange sets the value of the playerRange property.
// The range of players supported by this type of game activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playerrange
func (g_ GameActivityDefinition) SetPlayerRange(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlayerRange:"), value)
}

// The play style of the game activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playstyle
func (g_ GameActivityDefinition) PlayStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("playStyle"))
	return rv
}


// SetPlayStyle sets the value of the playStyle property.
// The play style of the game activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playstyle
func (g_ GameActivityDefinition) SetPlayStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlayStyle:"), value)
}

// The release state of the game activity definition in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/releasestate
func (g_ GameActivityDefinition) ReleaseState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("releaseState"))
	return rv
}


// SetReleaseState sets the value of the releaseState property.
// The release state of the game activity definition in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/releasestate
func (g_ GameActivityDefinition) SetReleaseState(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReleaseState:"), value)
}

// A fallback URL that can be used to construct a game-specific URL for players to share or join, if the joining device does not support the default URL.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/fallbackurl
func (g_ GameActivityDefinition) FallbackURL() foundation.URL {
	rv := objc.Send[foundation.URL](g_.ID, objc.Sel("fallbackURL"))
	return rv
}


// SetFallbackURL sets the value of the fallbackURL property.
// A fallback URL that can be used to construct a game-specific URL for players to share or join, if the joining device does not support the default URL.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/fallbackurl
func (g_ GameActivityDefinition) SetFallbackURL(value foundation.URL) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFallbackURL:"), value)
}

// A more detailed description of the game activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/details
func (g_ GameActivityDefinition) Details() string {
	rv := objc.Send[string](g_.ID, objc.Sel("details"))
	return rv
}


// SetDetails sets the value of the details property.
// A more detailed description of the game activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/details
func (g_ GameActivityDefinition) SetDetails(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDetails:"), objc.String(value))
}

// A short title for the game activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/title
func (g_ GameActivityDefinition) Title() string {
	rv := objc.Send[string](g_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// A short title for the game activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/title
func (g_ GameActivityDefinition) SetTitle(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// True if the activity supports an unlimited number of players. False if maxPlayers is set to a defined limit or if no player range is provided.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/supportsunlimitedplayers
func (g_ GameActivityDefinition) SupportsUnlimitedPlayers() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("supportsUnlimitedPlayers"))
	return rv
}


// SetSupportsUnlimitedPlayers sets the value of the supportsUnlimitedPlayers property.
// True if the activity supports an unlimited number of players. False if maxPlayers is set to a defined limit or if no player range is provided.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/supportsunlimitedplayers
func (g_ GameActivityDefinition) SetSupportsUnlimitedPlayers(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSupportsUnlimitedPlayers:"), value)
}

// Whether the activity can be joined by others via a party code.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/supportspartycode
func (g_ GameActivityDefinition) SupportsPartyCode() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("supportsPartyCode"))
	return rv
}


// SetSupportsPartyCode sets the value of the supportsPartyCode property.
// Whether the activity can be joined by others via a party code.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/supportspartycode
func (g_ GameActivityDefinition) SetSupportsPartyCode(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSupportsPartyCode:"), value)
}

// The developer defined identifier for a given game activity.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/identifier
func (g_ GameActivityDefinition) Identifier() string {
	rv := objc.Send[string](g_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The developer defined identifier for a given game activity.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/identifier
func (g_ GameActivityDefinition) SetIdentifier(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// Default properties defined by the developer for this type of game activity.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/defaultProperties
func (g_ GameActivityDefinition) DefaultProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("defaultProperties"))
	return rv
}

// The group identifier for the activity, if one exists.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/groupIdentifier
func (g_ GameActivityDefinition) GroupIdentifier() string {
	rv := objc.Send[string](g_.ID, objc.Sel("groupIdentifier"))
	return rv
}



