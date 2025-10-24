// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKMonteCarloStrategist */


/* debug [class_header]: Header for GKMonteCarloStrategist */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MonteCarloStrategist */
// An interface definition for the [MonteCarloStrategist] class.
type IMonteCarloStrategist interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MonteCarloStrategist */
	// properties:
	Budget() uint
	SetBudget(value uint)
	ExplorationParameter() uint
	SetExplorationParameter(value uint)
	GameModel() GameModel /* not a class type */
	SetGameModel(value GameModel /* not a class type */)
	RandomSource() Random /* not a class type */
	SetRandomSource(value Random /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MonteCarloStrategist */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MonteCarloStrategist */
// Alloc allocates a new instance without initialization.
func (mc _MonteCarloStrategistClass) Alloc() MonteCarloStrategist {
	rv := objc.Send[MonteCarloStrategist](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MonteCarloStrategist */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MonteCarloStrategist *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MonteCarloStrategist */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MonteCarloStrategist */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MonteCarloStrategist */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MonteCarloStrategist */

// The maximum number of game model states the strategist will examine when searching for a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMonteCarloStrategist/budget
func (m_ MonteCarloStrategist) Budget() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("budget"))
	return rv
}/* debug [instance_properties/getter]: budget */


// The maximum number of game model states the strategist will examine when searching for a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMonteCarloStrategist/budget
func (m_ MonteCarloStrategist) SetBudget(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBudget:"), value)
}/* debug [instance_properties/setter]: budget */


// A value that influences whether the strategist searches more broadly or more deeply for winning game model states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMonteCarloStrategist/explorationParameter
func (m_ MonteCarloStrategist) ExplorationParameter() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("explorationParameter"))
	return rv
}/* debug [instance_properties/getter]: explorationParameter */


// A value that influences whether the strategist searches more broadly or more deeply for winning game model states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMonteCarloStrategist/explorationParameter
func (m_ MonteCarloStrategist) SetExplorationParameter(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExplorationParameter:"), value)
}/* debug [instance_properties/setter]: explorationParameter */


// The model representing the current state of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/gamemodel
func (m_ MonteCarloStrategist) GameModel() GameModel /* not a class type */ {
	rv := objc.Send[GameModel](m_.ID, objc.Sel("gameModel"))
	return rv
}/* debug [instance_properties/getter]: gameModel */


// The model representing the current state of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/gamemodel
func (m_ MonteCarloStrategist) SetGameModel(value GameModel /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGameModel:"), value)
}/* debug [instance_properties/setter]: gameModel */


// A randomizer object to be used when the strategist randomly selects a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/randomsource
func (m_ MonteCarloStrategist) RandomSource() Random /* not a class type */ {
	rv := objc.Send[Random](m_.ID, objc.Sel("randomSource"))
	return rv
}/* debug [instance_properties/getter]: randomSource */


// A randomizer object to be used when the strategist randomly selects a move.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkstrategist/randomsource
func (m_ MonteCarloStrategist) SetRandomSource(value Random /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRandomSource:"), value)
}/* debug [instance_properties/setter]: randomSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKMonteCarloStrategist */



