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
	ClearBadgeForPlayersCompletionHandler(players []CloudPlayer, completionHandler unsafe.Pointer)
	GetShareURLWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadDataWithCompletionHandler(completionHandler unsafe.Pointer)
	PlayersWithConnectionState(state unsafe.Pointer) []CloudPlayer
	SaveDataCompletionHandler(data foundation.IData, completionHandler unsafe.Pointer)
	SendDataWithTransportTypeCompletionHandler(data foundation.IData, transport TransportType, completionHandler unsafe.Pointer)
	SendMessageWithLocalizedFormatKeyArgumentsDataToPlayersBadgePlayersCompletionHandler(key string, arguments []string, data foundation.IData, players []CloudPlayer, badgePlayers bool, completionHandler unsafe.Pointer)
	SetConnectionStateCompletionHandler(state unsafe.Pointer, completionHandler unsafe.Pointer)
	BadgedPlayers() []CloudPlayer
	Identifier() string
	LastModifiedDate() foundation.NSDate
	LastModifiedPlayer() GKCloudPlayer
	MaxNumberOfConnectedPlayers() int
	Owner() GKCloudPlayer
	Players() []CloudPlayer
	Title() string
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
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



// Adds a new event listener object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/add(listener:)
func (gc _GameSessionClass) AddEventListener(listener unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("addEventListener:"), listener)
}


// Creates a new game session inside of an iCloud container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/createSession(inContainer:withTitle:maxConnectedPlayers:completionHandler:)
func (gc _GameSessionClass) CreateSessionInContainerWithTitleMaxConnectedPlayersCompletionHandler(containerName string, title string, maxPlayers int, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("createSessionInContainer:withTitle:maxConnectedPlayers:completionHandler:"), objc.String(containerName), objc.String(title), maxPlayers, completionHandler)
}


// Loads a specific game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/load(withIdentifier:completionHandler:)
func (gc _GameSessionClass) LoadSessionWithIdentifierCompletionHandler(identifier string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("loadSessionWithIdentifier:completionHandler:"), objc.String(identifier), completionHandler)
}


// Retrieves all of the game sessions associated with a container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/loadSessions(inContainer:completionHandler:)
func (gc _GameSessionClass) LoadSessionsInContainerCompletionHandler(containerName string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("loadSessionsInContainer:completionHandler:"), objc.String(containerName), completionHandler)
}


// Stops listening to the event listener object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/remove(listener:)
func (gc _GameSessionClass) RemoveEventListener(listener unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("removeEventListener:"), listener)
}


// Removes the specified game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/remove(withIdentifier:completionHandler:)
func (gc _GameSessionClass) RemoveSessionWithIdentifierCompletionHandler(identifier string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("removeSessionWithIdentifier:completionHandler:"), objc.String(identifier), completionHandler)
}


// Clears the badge from the designated players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/clearBadge(for:completionHandler:)
func (g_ GameSession) ClearBadgeForPlayersCompletionHandler(players []CloudPlayer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("clearBadgeForPlayers:completionHandler:"), players, completionHandler)
}


// Retrieves the URL used to share a game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/getShareURL(completionHandler:)
func (g_ GameSession) GetShareURLWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("getShareURLWithCompletionHandler:"), completionHandler)
}


// Retrieves the game data from the current game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/loadData(completionHandler:)
func (g_ GameSession) LoadDataWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("loadDataWithCompletionHandler:"), completionHandler)
}


// Retrieves a list of players with the specified connection state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/players(with:)
func (g_ GameSession) PlayersWithConnectionState(state unsafe.Pointer) []CloudPlayer {
	rv := objc.Send[[]CloudPlayer](g_.ID, objc.Sel("playersWithConnectionState:"), state)
	return rv
}


// Saves the current game session data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/save(_:completionHandler:)
func (g_ GameSession) SaveDataCompletionHandler(data foundation.IData, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("saveData:completionHandler:"), data, completionHandler)
}


// Sends the indicated data to all connected players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/send(_:with:completionHandler:)
func (g_ GameSession) SendDataWithTransportTypeCompletionHandler(data foundation.IData, transport TransportType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("sendData:withTransportType:completionHandler:"), data, transport, completionHandler)
}


// Sends a message to players in a game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/sendMessage(withLocalizedFormatKey:arguments:data:to:badgePlayers:completionHandler:)
func (g_ GameSession) SendMessageWithLocalizedFormatKeyArgumentsDataToPlayersBadgePlayersCompletionHandler(key string, arguments []string, data foundation.IData, players []CloudPlayer, badgePlayers bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("sendMessageWithLocalizedFormatKey:arguments:data:toPlayers:badgePlayers:completionHandler:"), objc.String(key), arguments, data, players, badgePlayers, completionHandler)
}


// Sets the connection state for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/setConnectionState(_:completionHandler:)
func (g_ GameSession) SetConnectionStateCompletionHandler(state unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setConnectionState:completionHandler:"), state, completionHandler)
}


// An array containing all of the currently badged players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/badgedPlayers
func (g_ GameSession) BadgedPlayers() []CloudPlayer {
	rv := objc.Send[[]CloudPlayer](g_.ID, objc.Sel("badgedPlayers"))
	return rv
}


// A unique identifier for a game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/identifier
func (g_ GameSession) Identifier() string {
	rv := objc.Send[string](g_.ID, objc.Sel("identifier"))
	return rv
}


// The date that the game session was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/lastModifiedDate
func (g_ GameSession) LastModifiedDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](g_.ID, objc.Sel("lastModifiedDate"))
	return rv
}


// The last player to modify the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/lastModifiedPlayer
func (g_ GameSession) LastModifiedPlayer() GKCloudPlayer {
	rv := objc.Send[GKCloudPlayer](g_.ID, objc.Sel("lastModifiedPlayer"))
	return rv
}


// The maximum number of players allowed to connect to the game session at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/maxNumberOfConnectedPlayers
func (g_ GameSession) MaxNumberOfConnectedPlayers() int {
	rv := objc.Send[int](g_.ID, objc.Sel("maxNumberOfConnectedPlayers"))
	return rv
}


// A player object that represents the owner of the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/owner
func (g_ GameSession) Owner() GKCloudPlayer {
	rv := objc.Send[GKCloudPlayer](g_.ID, objc.Sel("owner"))
	return rv
}


// An array of player objects associated with the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/players
func (g_ GameSession) Players() []CloudPlayer {
	rv := objc.Send[[]CloudPlayer](g_.ID, objc.Sel("players"))
	return rv
}


// The title of the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/title
func (g_ GameSession) Title() string {
	rv := objc.Send[string](g_.ID, objc.Sel("title"))
	return rv
}


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameSession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameSession) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}



