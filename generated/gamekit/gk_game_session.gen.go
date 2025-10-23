// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GameSession] class.
var (
	GameSessionClass     _GameSessionClass
	GameSessionClassOnce sync.Once
)

func getGameSessionClass() _GameSessionClass {
	GameSessionClassOnce.Do(func() {
		GameSessionClass = _GameSessionClass{objc.GetClass("GKGameSession")}
	})
	return GameSessionClass
}

type _GameSessionClass struct {
	class objc.Class
}

// An interface definition for the [GameSession] class.
type IGameSession interface {
	objectivec.IObject
	// properties:
	BadgedPlayers() objc.IObject /* cross-framework: CloudPlayer */
	SetBadgedPlayers(value objc.IObject /* cross-framework: CloudPlayer */)
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	LastModifiedDate() foundation.objc.IObject /* cross-framework: Date */
	SetLastModifiedDate(value foundation.objc.IObject /* cross-framework: Date */)
	LastModifiedPlayer() objc.IObject /* cross-framework: CloudPlayer */
	SetLastModifiedPlayer(value objc.IObject /* cross-framework: CloudPlayer */)
	MaxNumberOfConnectedPlayers() int /* primitive/slice/pointer. */
	SetMaxNumberOfConnectedPlayers(value int /* primitive/slice/pointer. */)
	Owner() objc.IObject /* cross-framework: CloudPlayer */
	SetOwner(value objc.IObject /* cross-framework: CloudPlayer */)
	Players() objc.IObject /* cross-framework: CloudPlayer */
	SetPlayers(value objc.IObject /* cross-framework: CloudPlayer */)
	Title() string /* primitive/slice/pointer. */
	SetTitle(value string /* primitive/slice/pointer. */)
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
	// methods:
}

// A game session you can use to save game data, invite other players, and create turn-based and real-time game apps.
//
// Use a object to play turn-based and real-time games in iCloud. Every instance of a game session resides inside of an iCloud container. You can create multiple sessions for a single app, allowing players to play several games at once. All of the information for a game session is saved in the owner’s iCloud. Each session can contain a maximum of 100 players. Inside of a session, up to 16 of those players can be connected to each other in real-time. The 16 connected players can be selected from any of the 100 players in the session. You can change a connected player with another player in the session at any time. After a game session is created, you can save game data in iCloud. Each game session can save a maximum of 512KB data. This prevents games from using a large about of space in a user’s iCloud account. This data can be loaded, edited, and saved by anyone in the game session, providing your app provides this behavior. You must ensure that you delete a game session from a user’s iCloud after a game is over, otherwise the session will stay in the user’s iCloud forever. Game sessions are not automatically removed after a set amount of time. They can only be actively removed.


// A game session you can use to save game data, invite other players, and create turn-based and real-time game apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession
type GameSession struct {
	objectivec.Object
}

// GameSessionFrom constructs a [GameSession] from an unsafe.Pointer.
//
// A game session you can use to save game data, invite other players, and create turn-based and real-time game apps.
func GameSessionFrom(ptr unsafe.Pointer) GameSession {
	return GameSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GameSessionClass) Alloc() GameSession {
	rv := objc.Send[GameSession](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GameSessionClass) New() GameSession {
	rv := objc.Send[GameSession](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GameSession) Init() GameSession {
	rv := objc.Send[GameSession](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GameSession) Autorelease() GameSession {
	rv := objc.Send[GameSession](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGameSession creates a new GameSession instance.
func NewGameSession() GameSession {
	return getGameSessionClass().New()
}



// An array containing all of the currently badged players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/badgedplayers
func (g_ GameSession) BadgedPlayers() objc.IObject /* cross-framework: CloudPlayer */ {
	rv := objc.Send[CloudPlayer](g_.ID, objc.Sel("badgedPlayers"))
	return rv
}


// An array containing all of the currently badged players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/badgedplayers
func (g_ GameSession) SetBadgedPlayers(value objc.IObject /* cross-framework: CloudPlayer */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBadgedPlayers:"), value)
}


// A unique identifier for a game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/identifier
func (g_ GameSession) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](g_.ID, objc.Sel("identifier"))
	return rv
}


// A unique identifier for a game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/identifier
func (g_ GameSession) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The date that the game session was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/lastmodifieddate
func (g_ GameSession) LastModifiedDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](g_.ID, objc.Sel("lastModifiedDate"))
	return rv
}


// The date that the game session was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/lastmodifieddate
func (g_ GameSession) SetLastModifiedDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLastModifiedDate:"), value)
}


// The last player to modify the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/lastmodifiedplayer
func (g_ GameSession) LastModifiedPlayer() objc.IObject /* cross-framework: CloudPlayer */ {
	rv := objc.Send[CloudPlayer](g_.ID, objc.Sel("lastModifiedPlayer"))
	return rv
}


// The last player to modify the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/lastmodifiedplayer
func (g_ GameSession) SetLastModifiedPlayer(value objc.IObject /* cross-framework: CloudPlayer */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLastModifiedPlayer:"), value)
}


// The maximum number of players allowed to connect to the game session at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/maxnumberofconnectedplayers
func (g_ GameSession) MaxNumberOfConnectedPlayers() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](g_.ID, objc.Sel("maxNumberOfConnectedPlayers"))
	return rv
}


// The maximum number of players allowed to connect to the game session at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/maxnumberofconnectedplayers
func (g_ GameSession) SetMaxNumberOfConnectedPlayers(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMaxNumberOfConnectedPlayers:"), value)
}


// A player object that represents the owner of the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/owner
func (g_ GameSession) Owner() objc.IObject /* cross-framework: CloudPlayer */ {
	rv := objc.Send[CloudPlayer](g_.ID, objc.Sel("owner"))
	return rv
}


// A player object that represents the owner of the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/owner
func (g_ GameSession) SetOwner(value objc.IObject /* cross-framework: CloudPlayer */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOwner:"), value)
}


// An array of player objects associated with the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/players
func (g_ GameSession) Players() objc.IObject /* cross-framework: CloudPlayer */ {
	rv := objc.Send[CloudPlayer](g_.ID, objc.Sel("players"))
	return rv
}


// An array of player objects associated with the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/players
func (g_ GameSession) SetPlayers(value objc.IObject /* cross-framework: CloudPlayer */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlayers:"), value)
}


// The title of the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/title
func (g_ GameSession) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](g_.ID, objc.Sel("title"))
	return rv
}


// The title of the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamesession/title
func (g_ GameSession) SetTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameSession) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](g_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameSession) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}



