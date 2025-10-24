// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKGameSession */


/* debug [class_header]: Header for GKGameSession */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GameSession */
// An interface definition for the [GameSession] class.
type IGameSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GameSession */
	// properties:
	BadgedPlayers() []CloudPlayer
	Identifier() objc.IObject /* cross-framework: NSString */
	LastModifiedDate() objc.IObject /* cross-framework: NSDate */
	LastModifiedPlayer() IGKCloudPlayer
	MaxNumberOfConnectedPlayers() int
	Owner() IGKCloudPlayer
	Players() []CloudPlayer
	Title() objc.IObject /* cross-framework: NSString */
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GameSession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GameSession */
// Alloc allocates a new instance without initialization.
func (gc _GameSessionClass) Alloc() GameSession {
	rv := objc.Send[GameSession](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GameSession */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GameSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GameSession */

// Adds a new event listener object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/add(listener:)
func (gc _GameSessionClass) AddEventListener(listener unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("addEventListener:"), listener)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AddEventListener) */


// Creates a new game session inside of an iCloud container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/createSession(inContainer:withTitle:maxConnectedPlayers:completionHandler:)
func (gc _GameSessionClass) CreateSessionInContainerWithTitleMaxConnectedPlayersCompletionHandler(containerName objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, maxPlayers int, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("createSessionInContainer:withTitle:maxConnectedPlayers:completionHandler:"), containerName, title, maxPlayers, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateSessionInContainerWithTitleMaxConnectedPlayersCompletionHandler) */


// Loads a specific game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/load(withIdentifier:completionHandler:)
func (gc _GameSessionClass) LoadSessionWithIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("loadSessionWithIdentifier:completionHandler:"), identifier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadSessionWithIdentifierCompletionHandler) */


// Retrieves all of the game sessions associated with a container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/loadSessions(inContainer:completionHandler:)
func (gc _GameSessionClass) LoadSessionsInContainerCompletionHandler(containerName objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("loadSessionsInContainer:completionHandler:"), containerName, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadSessionsInContainerCompletionHandler) */


// Stops listening to the event listener object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/remove(listener:)
func (gc _GameSessionClass) RemoveEventListener(listener unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("removeEventListener:"), listener)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemoveEventListener) */


// Removes the specified game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/remove(withIdentifier:completionHandler:)
func (gc _GameSessionClass) RemoveSessionWithIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("removeSessionWithIdentifier:completionHandler:"), identifier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemoveSessionWithIdentifierCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GameSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GameSession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GameSession */

// An array containing all of the currently badged players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/badgedPlayers
func (g_ GameSession) BadgedPlayers() []CloudPlayer {
	rv := objc.Send[[]CloudPlayer](g_.ID, objc.Sel("badgedPlayers"))
	return rv
}/* debug [instance_properties/getter]: badgedPlayers */


// A unique identifier for a game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/identifier
func (g_ GameSession) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The date that the game session was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/lastModifiedDate
func (g_ GameSession) LastModifiedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](g_.ID, objc.Sel("lastModifiedDate"))
	return rv
}/* debug [instance_properties/getter]: lastModifiedDate */


// The last player to modify the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/lastModifiedPlayer
func (g_ GameSession) LastModifiedPlayer() IGKCloudPlayer {
	rv := objc.Send[CloudPlayer](g_.ID, objc.Sel("lastModifiedPlayer"))
	return rv
}/* debug [instance_properties/getter]: lastModifiedPlayer */


// The maximum number of players allowed to connect to the game session at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/maxNumberOfConnectedPlayers
func (g_ GameSession) MaxNumberOfConnectedPlayers() int {
	rv := objc.Send[int](g_.ID, objc.Sel("maxNumberOfConnectedPlayers"))
	return rv
}/* debug [instance_properties/getter]: maxNumberOfConnectedPlayers */


// A player object that represents the owner of the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/owner
func (g_ GameSession) Owner() IGKCloudPlayer {
	rv := objc.Send[CloudPlayer](g_.ID, objc.Sel("owner"))
	return rv
}/* debug [instance_properties/getter]: owner */


// An array of player objects associated with the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/players
func (g_ GameSession) Players() []CloudPlayer {
	rv := objc.Send[[]CloudPlayer](g_.ID, objc.Sel("players"))
	return rv
}/* debug [instance_properties/getter]: players */


// The title of the game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSession/title
func (g_ GameSession) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameSession) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](g_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameSession) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGameSession */



