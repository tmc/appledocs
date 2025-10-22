// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MonteCarloStrategist] class.
var (
	MonteCarloStrategistClass     _MonteCarloStrategistClass
	MonteCarloStrategistClassOnce sync.Once
)

func getMonteCarloStrategistClass() _MonteCarloStrategistClass {
	MonteCarloStrategistClassOnce.Do(func() {
		MonteCarloStrategistClass = _MonteCarloStrategistClass{objc.GetClass("GKMonteCarloStrategist")}
	})
	return MonteCarloStrategistClass
}

type _MonteCarloStrategistClass struct {
	class objc.Class
}

// An interface definition for the [MonteCarloStrategist] class.
type IMonteCarloStrategist interface {
	objectivec.IObject
	Budget() uint
	SetBudget(value uint)
	ExplorationParameter() uint
	SetExplorationParameter(value uint)
	GameModel() unsafe.Pointer
	SetGameModel(value unsafe.Pointer)
	RandomSource() unsafe.Pointer
	SetRandomSource(value unsafe.Pointer)
}

// An AI that chooses moves in turn-based games using a strategy.
//
// To use this strategy, you indicate whether a possible states of your game model represents a win, and the strategist randomly searches possible game model states in order to find moves that will likely result in winning the game. You provide information about your game model to the strategist by implementing the , , and protocols in your custom classes, then use the strategist’s methods to find optimal moves.


// An AI that chooses moves in turn-based games using a strategy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMonteCarloStrategist

type MonteCarloStrategist struct {
	objectivec.Object
}

// MonteCarloStrategistFrom constructs a [MonteCarloStrategist] from an unsafe.Pointer.
//
// An AI that chooses moves in turn-based games using a strategy.
func MonteCarloStrategistFrom(ptr unsafe.Pointer) MonteCarloStrategist {
	return MonteCarloStrategist{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MonteCarloStrategistClass) Alloc() MonteCarloStrategist {
	rv := objc.Send[MonteCarloStrategist](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MonteCarloStrategistClass) New() MonteCarloStrategist {
	rv := objc.Send[MonteCarloStrategist](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MonteCarloStrategist) Init() MonteCarloStrategist {
	rv := objc.Send[MonteCarloStrategist](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MonteCarloStrategist) Autorelease() MonteCarloStrategist {
	rv := objc.Send[MonteCarloStrategist](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMonteCarloStrategist creates a new MonteCarloStrategist instance.
func NewMonteCarloStrategist() MonteCarloStrategist {
	return getMonteCarloStrategistClass().New()
}



// The maximum number of game model states the strategist will examine when searching for a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMonteCarloStrategist/budget

func (m_ MonteCarloStrategist) Budget() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("budget"))
	return rv
}


// The maximum number of game model states the strategist will examine when searching for a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMonteCarloStrategist/budget

func (m_ MonteCarloStrategist) SetBudget(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBudget:"), value)
}


// A value that influences whether the strategist searches more broadly or more deeply for winning game model states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMonteCarloStrategist/explorationParameter

func (m_ MonteCarloStrategist) ExplorationParameter() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("explorationParameter"))
	return rv
}


// A value that influences whether the strategist searches more broadly or more deeply for winning game model states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMonteCarloStrategist/explorationParameter

func (m_ MonteCarloStrategist) SetExplorationParameter(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExplorationParameter:"), value)
}


// The model representing the current state of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/gamemodel

func (m_ MonteCarloStrategist) GameModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("gameModel"))
	return rv
}


// The model representing the current state of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/gamemodel

func (m_ MonteCarloStrategist) SetGameModel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGameModel:"), value)
}


// A randomizer object to be used when the strategist randomly selects a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/randomsource

func (m_ MonteCarloStrategist) RandomSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("randomSource"))
	return rv
}


// A randomizer object to be used when the strategist randomly selects a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/randomsource

func (m_ MonteCarloStrategist) SetRandomSource(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRandomSource:"), value)
}



