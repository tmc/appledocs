// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKSavedGame */


/* debug [class_header]: Header for GKSavedGame */
// The class instance for the [SavedGame] class.
var (
	SavedGameClass     _SavedGameClass
	SavedGameClassOnce sync.Once
)

func getSavedGameClass() _SavedGameClass {
	SavedGameClassOnce.Do(func() {
		SavedGameClass = _SavedGameClass{objc.GetClass("GKSavedGame")}
	})
	return SavedGameClass
}

type _SavedGameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SavedGame */
// An interface definition for the [SavedGame] class.
type ISavedGame interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SavedGame */
	// properties:
	DeviceName() objc.IObject /* cross-framework: NSString */
	ModificationDate() objc.IObject /* cross-framework: NSDate */
	Name() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SavedGame */
	// methods:
	LoadDataWithCompletionHandler(handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SavedGame */
// Alloc allocates a new instance without initialization.
func (sc _SavedGameClass) Alloc() SavedGame {
	rv := objc.Send[SavedGame](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SavedGameClass) New() SavedGame {
	rv := objc.Send[SavedGame](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SavedGame) Init() SavedGame {
	rv := objc.Send[SavedGame](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SavedGame) Autorelease() SavedGame {
	rv := objc.Send[SavedGame](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSavedGame creates a new SavedGame instance.
func NewSavedGame() SavedGame {
	return getSavedGameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SavedGame */
// An object that represents a file containing saved game data.
//
// A object represents the file that contains game data you saved using the method. You don’t create objects directly. Instead use the method to get the games you saved. Then get the filename, its modification date, and the name of the device the player used to save the game from the returned objects. Use the method to get the actual game data you saved in the file. To delete saved games, use the method.


// An object that represents a file containing saved game data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSavedGame
type SavedGame struct {
	objectivec.Object
}

// SavedGameFrom constructs a [SavedGame] from an unsafe.Pointer.
//
// An object that represents a file containing saved game data.
func SavedGameFrom(ptr unsafe.Pointer) SavedGame {
	return SavedGame{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SavedGame *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SavedGame */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SavedGame */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SavedGame */

// Loads the game data from the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSavedGame/loadData(completionHandler:)
func (s_ SavedGame) LoadDataWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("loadDataWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: LoadDataWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SavedGame */

// The name of the device that the player uses to save the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSavedGame/deviceName
func (s_ SavedGame) DeviceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("deviceName"))
	return rv
}/* debug [instance_properties/getter]: deviceName */


// The date when you saved the game data or modified it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSavedGame/modificationDate
func (s_ SavedGame) ModificationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](s_.ID, objc.Sel("modificationDate"))
	return rv
}/* debug [instance_properties/getter]: modificationDate */


// The name of the saved game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSavedGame/name
func (s_ SavedGame) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKSavedGame */



