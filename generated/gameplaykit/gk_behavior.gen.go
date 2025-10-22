// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Behavior] class.
var (
	BehaviorClass     _BehaviorClass
	BehaviorClassOnce sync.Once
)

func getBehaviorClass() _BehaviorClass {
	BehaviorClassOnce.Do(func() {
		BehaviorClass = _BehaviorClass{objc.GetClass("GKBehavior")}
	})
	return BehaviorClass
}

type _BehaviorClass struct {
	class objc.Class
}

// An interface definition for the [Behavior] class.
type IBehavior interface {
	objectivec.IObject
	RemoveGoal(goal IGKGoal)
	RemoveAllGoals()
	SetObjectForKeyedSubscript(weight foundation.INumber, goal IGKGoal)
	SetWeightForGoal(weight float32, goal IGKGoal)
	ObjectForKeyedSubscript(goal IGKGoal) foundation.Number
	ObjectAtIndexedSubscript(idx uint) Goal
	WeightForGoal(goal IGKGoal) float32
	GoalCount() int
	Behavior() GKBehavior
	SetBehavior(value GKBehavior)
}

// A set of goals that together influence the movement of an agent.
//
// By combining multiple goals ( objects) you can create complex behavior, such as groups of agents ( objects) that move together naturally. To assign a set of goals to an agent, use its property. To learn more about using goals and agents, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior
type Behavior struct {
	objectivec.Object
}

// BehaviorFrom constructs a [Behavior] from an unsafe.Pointer.
//
// A set of goals that together influence the movement of an agent.
func BehaviorFrom(ptr unsafe.Pointer) Behavior {
	return Behavior{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BehaviorClass) Alloc() Behavior {
	rv := objc.Send[Behavior](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BehaviorClass) New() Behavior {
	rv := objc.Send[Behavior](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ Behavior) Init() Behavior {
	rv := objc.Send[Behavior](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ Behavior) Autorelease() Behavior {
	rv := objc.Send[Behavior](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBehavior creates a new Behavior instance.
func NewBehavior() Behavior {
	return getBehaviorClass().New()
}




// Creates a behavior with a single goal.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goal:weight:)
func NewBehaviorWithGoalWeight(goal IGKGoal, weight float32) Behavior {
	rv := objc.Send[Behavior](objc.ID(getBehaviorClass().class), objc.Sel("behaviorWithGoal:weight:"), goal, weight)
	return rv
}



// Creates a behavior with the specified goals.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goals:)
func NewBehaviorWithGoals(goals []Goal) Behavior {
	rv := objc.Send[Behavior](objc.ID(getBehaviorClass().class), objc.Sel("behaviorWithGoals:"), goals)
	return rv
}



// Creates a behavior with the specified goals and weights.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goals:andWeights:)
func NewBehaviorWithGoalsAndWeights(goals []Goal, weights []foundation.INumber) Behavior {
	rv := objc.Send[Behavior](objc.ID(getBehaviorClass().class), objc.Sel("behaviorWithGoals:andWeights:"), goals, weights)
	return rv
}



// Creates a behavior with the specified mapping of goals to their weights.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(weightedGoals:)
func NewBehaviorWithWeightedGoals(weightedGoals unsafe.Pointer) Behavior {
	rv := objc.Send[Behavior](objc.ID(getBehaviorClass().class), objc.Sel("behaviorWithWeightedGoals:"), weightedGoals)
	return rv
}


// Creates a behavior with a single goal.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goal:weight:)
func (bc _BehaviorClass) BehaviorWithGoalWeight(goal IGKGoal, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("behaviorWithGoal:weight:"), goal, weight)
	return rv
}

// Creates a behavior with the specified goals.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goals:)
func (bc _BehaviorClass) BehaviorWithGoals(goals []Goal) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("behaviorWithGoals:"), goals)
	return rv
}

// Creates a behavior with the specified goals and weights.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goals:andWeights:)
func (bc _BehaviorClass) BehaviorWithGoalsAndWeights(goals []Goal, weights []foundation.INumber) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("behaviorWithGoals:andWeights:"), goals, weights)
	return rv
}

// Creates a behavior with the specified mapping of goals to their weights.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(weightedGoals:)
func (bc _BehaviorClass) BehaviorWithWeightedGoals(weightedGoals unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("behaviorWithWeightedGoals:"), weightedGoals)
	return rv
}

// Removes the specified goal from the behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/remove(_:)
func (b_ Behavior) RemoveGoal(goal IGKGoal) {
	objc.Send[objc.ID](b_.ID, objc.Sel("removeGoal:"), goal)
}

// Removes all goals from the behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/removeAllGoals()
func (b_ Behavior) RemoveAllGoals() {
	objc.Send[objc.ID](b_.ID, objc.Sel("removeAllGoals"))
}

// Sets the weight for the goal specified by subscript syntax.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/setObject:forKeyedSubscript:
func (b_ Behavior) SetObjectForKeyedSubscript(weight foundation.INumber, goal IGKGoal) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setObject:forKeyedSubscript:"), weight, goal)
}

// Sets the weight for the specified goal’s influence on agents, adding that goal to the behavior if not already present.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/setWeight(_:for:)
func (b_ Behavior) SetWeightForGoal(weight float32, goal IGKGoal) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setWeight:forGoal:"), weight, goal)
}

// Returns the weight associated with the goal specified by subscript syntax.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/subscript(_:)-2yvko
func (b_ Behavior) ObjectForKeyedSubscript(goal IGKGoal) foundation.Number {
	rv := objc.Send[foundation.Number](b_.ID, objc.Sel("objectForKeyedSubscript:"), goal)
	return rv
}

// Returns the goal at the specified index in the behavior’s list of goals.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/subscript(_:)-997a9
func (b_ Behavior) ObjectAtIndexedSubscript(idx uint) Goal {
	rv := objc.Send[Goal](b_.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}

// Returns the weight for the specified goal’s influence on agents.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/weight(for:)
func (b_ Behavior) WeightForGoal(goal IGKGoal) float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("weightForGoal:"), goal)
	return rv
}

// The number of goals in the behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/goalCount
func (b_ Behavior) GoalCount() int {
	rv := objc.Send[int](b_.ID, objc.Sel("goalCount"))
	return rv
}

// A weighted collection of goals that influence the agent’s movement.
//
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkagent/behavior
func (b_ Behavior) Behavior() GKBehavior {
	rv := objc.Send[GKBehavior](b_.ID, objc.Sel("behavior"))
	return rv
}


// SetBehavior sets the value of the behavior property.
// A weighted collection of goals that influence the agent’s movement.

//
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkagent/behavior
func (b_ Behavior) SetBehavior(value GKBehavior) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBehavior:"), value)
}


