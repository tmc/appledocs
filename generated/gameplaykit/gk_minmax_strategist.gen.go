// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MinmaxStrategist] class.
type IMinmaxStrategist interface {
	objectivec.IObject
	// properties:
	MaxLookAheadDepth() int
	SetMaxLookAheadDepth(value int)
	GameModel() GameModel /* not a class type */
	SetGameModel(value GameModel /* not a class type */)
	RandomSource() Random /* not a class type */
	SetRandomSource(value Random /* not a class type */)
	// methods:
	BestMoveForPlayer(player objectivec.IObject) objc.ID
	RandomMoveForPlayerFromNumberOfBestMoves(player objectivec.IObject, numMovesToConsider int) objc.ID
}

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

// Alloc allocates a new instance without initialization.
func (mc _MinmaxStrategistClass) Alloc() MinmaxStrategist {
	rv := objc.Send[MinmaxStrategist](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Computes and returns the best possible move for the specified player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMinmaxStrategist/bestMove(for:)
func (m_ MinmaxStrategist) BestMoveForPlayer(player objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("bestMoveForPlayer:"), player)
	return rv
}


// Computes several of the best possible moves for the specified player, and returns a move randomly selected from among them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMinmaxStrategist/randomMove(for:fromNumberOfBestMoves:)
func (m_ MinmaxStrategist) RandomMoveForPlayerFromNumberOfBestMoves(player objectivec.IObject, numMovesToConsider int) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("randomMoveForPlayer:fromNumberOfBestMoves:"), player, numMovesToConsider)
	return rv
}


// The number of future turns for the strategist to consider when planning moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMinmaxStrategist/maxLookAheadDepth
func (m_ MinmaxStrategist) MaxLookAheadDepth() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maxLookAheadDepth"))
	return rv
}


// The number of future turns for the strategist to consider when planning moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMinmaxStrategist/maxLookAheadDepth
func (m_ MinmaxStrategist) SetMaxLookAheadDepth(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxLookAheadDepth:"), value)
}


// The model representing the current state of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/gamemodel
func (m_ MinmaxStrategist) GameModel() GameModel /* not a class type */ {
	rv := objc.Send[GameModel](m_.ID, objc.Sel("gameModel"))
	return rv
}


// The model representing the current state of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/gamemodel
func (m_ MinmaxStrategist) SetGameModel(value GameModel /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGameModel:"), value)
}


// A randomizer object to be used when the strategist randomly selects a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/randomsource
func (m_ MinmaxStrategist) RandomSource() Random /* not a class type */ {
	rv := objc.Send[Random](m_.ID, objc.Sel("randomSource"))
	return rv
}


// A randomizer object to be used when the strategist randomly selects a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/randomsource
func (m_ MinmaxStrategist) SetRandomSource(value Random /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRandomSource:"), value)
}



