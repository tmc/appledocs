// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKMinmaxStrategist */


/* debug [class_header]: Header for GKMinmaxStrategist */
// The class instance for the [MinmaxStrategist] class.
var (
	MinmaxStrategistClass     _MinmaxStrategistClass
	MinmaxStrategistClassOnce sync.Once
)

func getMinmaxStrategistClass() _MinmaxStrategistClass {
	MinmaxStrategistClassOnce.Do(func() {
		MinmaxStrategistClass = _MinmaxStrategistClass{objc.GetClass("GKMinmaxStrategist")}
	})
	return MinmaxStrategistClass
}

type _MinmaxStrategistClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MinmaxStrategist */
// An interface definition for the [MinmaxStrategist] class.
type IMinmaxStrategist interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MinmaxStrategist */
	// properties:
	MaxLookAheadDepth() int
	SetMaxLookAheadDepth(value int)
	GameModel() GameModel /* not a class type */
	SetGameModel(value GameModel /* not a class type */)
	RandomSource() Random /* not a class type */
	SetRandomSource(value Random /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MinmaxStrategist */
	// methods:
	BestMoveForPlayer(player unsafe.Pointer) unsafe.Pointer
	RandomMoveForPlayerFromNumberOfBestMoves(player unsafe.Pointer, numMovesToConsider int) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MinmaxStrategist */
// Alloc allocates a new instance without initialization.
func (mc _MinmaxStrategistClass) Alloc() MinmaxStrategist {
	rv := objc.Send[MinmaxStrategist](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MinmaxStrategistClass) New() MinmaxStrategist {
	rv := objc.Send[MinmaxStrategist](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MinmaxStrategist) Init() MinmaxStrategist {
	rv := objc.Send[MinmaxStrategist](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MinmaxStrategist) Autorelease() MinmaxStrategist {
	rv := objc.Send[MinmaxStrategist](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMinmaxStrategist creates a new MinmaxStrategist instance.
func NewMinmaxStrategist() MinmaxStrategist {
	return getMinmaxStrategistClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MinmaxStrategist */
// An AI that chooses moves in turn-based games using a strategy.
//
// To use this strategy, you provide scores that rate possible states of your game model for their desirability to a player, and the strategist exhaustively searches all possible game model states in order to make choices that maximize the rating for its own moves and minimize the rating for an opponent’s moves. You provide information about your game model to the strategist by implementing the , , and protocols in your custom classes, and then use the strategist’s methods to find optimal moves.


// An AI that chooses moves in turn-based games using a strategy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMinmaxStrategist
type MinmaxStrategist struct {
	objectivec.Object
}

// MinmaxStrategistFrom constructs a [MinmaxStrategist] from an unsafe.Pointer.
//
// An AI that chooses moves in turn-based games using a strategy.
func MinmaxStrategistFrom(ptr unsafe.Pointer) MinmaxStrategist {
	return MinmaxStrategist{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MinmaxStrategist *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MinmaxStrategist */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MinmaxStrategist */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MinmaxStrategist */

// Computes and returns the best possible move for the specified player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMinmaxStrategist/bestMove(for:)
func (m_ MinmaxStrategist) BestMoveForPlayer(player unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bestMoveForPlayer:"), player)
	return rv
}/* debug [instance_methods/method]: BestMoveForPlayer */


// Computes several of the best possible moves for the specified player, and returns a move randomly selected from among them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMinmaxStrategist/randomMove(for:fromNumberOfBestMoves:)
func (m_ MinmaxStrategist) RandomMoveForPlayerFromNumberOfBestMoves(player unsafe.Pointer, numMovesToConsider int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("randomMoveForPlayer:fromNumberOfBestMoves:"), player, numMovesToConsider)
	return rv
}/* debug [instance_methods/method]: RandomMoveForPlayerFromNumberOfBestMoves */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MinmaxStrategist */

// The number of future turns for the strategist to consider when planning moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMinmaxStrategist/maxLookAheadDepth
func (m_ MinmaxStrategist) MaxLookAheadDepth() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maxLookAheadDepth"))
	return rv
}/* debug [instance_properties/getter]: maxLookAheadDepth */


// The number of future turns for the strategist to consider when planning moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMinmaxStrategist/maxLookAheadDepth
func (m_ MinmaxStrategist) SetMaxLookAheadDepth(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxLookAheadDepth:"), value)
}/* debug [instance_properties/setter]: maxLookAheadDepth */


// The model representing the current state of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/gamemodel
func (m_ MinmaxStrategist) GameModel() GameModel /* not a class type */ {
	rv := objc.Send[GameModel](m_.ID, objc.Sel("gameModel"))
	return rv
}/* debug [instance_properties/getter]: gameModel */


// The model representing the current state of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/gamemodel
func (m_ MinmaxStrategist) SetGameModel(value GameModel /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGameModel:"), value)
}/* debug [instance_properties/setter]: gameModel */


// A randomizer object to be used when the strategist randomly selects a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/randomsource
func (m_ MinmaxStrategist) RandomSource() Random /* not a class type */ {
	rv := objc.Send[Random](m_.ID, objc.Sel("randomSource"))
	return rv
}/* debug [instance_properties/getter]: randomSource */


// A randomizer object to be used when the strategist randomly selects a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/randomsource
func (m_ MinmaxStrategist) SetRandomSource(value Random /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRandomSource:"), value)
}/* debug [instance_properties/setter]: randomSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKMinmaxStrategist */



