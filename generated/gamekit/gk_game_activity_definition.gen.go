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
	DefaultProperties() unsafe.Pointer
	GroupIdentifier() string
	Details() string
	SetDetails(value string)
	FallbackURL() foundation.URL
	SetFallbackURL(value foundation.IURL)
	Identifier() string
	SetIdentifier(value string)
	PlayStyle() unsafe.Pointer
	SetPlayStyle(value unsafe.Pointer)
	PlayerRange() unsafe.Pointer
	SetPlayerRange(value unsafe.Pointer)
	ReleaseState() unsafe.Pointer
	SetReleaseState(value unsafe.Pointer)
	SupportsPartyCode() bool
	SetSupportsPartyCode(value bool)
	SupportsUnlimitedPlayers() bool
	SetSupportsUnlimitedPlayers(value bool)
	Title() string
	SetTitle(value string)
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



// Default properties defined by the developer for this type of game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/defaultProperties

func (g_ GameActivityDefinition) DefaultProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("defaultProperties"))
	return rv
}


// The group identifier for the activity, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityDefinition/groupIdentifier

func (g_ GameActivityDefinition) GroupIdentifier() string {
	rv := objc.Send[string](g_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// A more detailed description of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/details

func (g_ GameActivityDefinition) Details() string {
	rv := objc.Send[string](g_.ID, objc.Sel("details"))
	return rv
}


// A more detailed description of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/details

func (g_ GameActivityDefinition) SetDetails(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDetails:"), objc.String(value))
}


// A fallback URL that can be used to construct a game-specific URL for players to share or join, if the joining device does not support the default URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/fallbackurl

func (g_ GameActivityDefinition) FallbackURL() foundation.URL {
	rv := objc.Send[foundation.URL](g_.ID, objc.Sel("fallbackURL"))
	return rv
}


// A fallback URL that can be used to construct a game-specific URL for players to share or join, if the joining device does not support the default URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/fallbackurl

func (g_ GameActivityDefinition) SetFallbackURL(value foundation.IURL) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFallbackURL:"), value)
}


// The developer defined identifier for a given game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/identifier

func (g_ GameActivityDefinition) Identifier() string {
	rv := objc.Send[string](g_.ID, objc.Sel("identifier"))
	return rv
}


// The developer defined identifier for a given game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/identifier

func (g_ GameActivityDefinition) SetIdentifier(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The play style of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playstyle

func (g_ GameActivityDefinition) PlayStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("playStyle"))
	return rv
}


// The play style of the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/playstyle

func (g_ GameActivityDefinition) SetPlayStyle(value unsafe.Pointer) {
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

func (g_ GameActivityDefinition) ReleaseState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("releaseState"))
	return rv
}


// The release state of the game activity definition in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/releasestate

func (g_ GameActivityDefinition) SetReleaseState(value unsafe.Pointer) {
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

func (g_ GameActivityDefinition) Title() string {
	rv := objc.Send[string](g_.ID, objc.Sel("title"))
	return rv
}


// A short title for the game activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgameactivitydefinition/title

func (g_ GameActivityDefinition) SetTitle(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTitle:"), objc.String(value))
}



