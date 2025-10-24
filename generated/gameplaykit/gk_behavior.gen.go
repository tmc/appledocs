// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKBehavior */


/* debug [class_header]: Header for GKBehavior */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Behavior */
// An interface definition for the [Behavior] class.
type IBehavior interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Behavior */
	// properties:
	GoalCount() int
	Behavior() IGKBehavior
	SetBehavior(value IGKBehavior)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Behavior */
	// methods:
	RemoveGoal(goal IGKGoal)
	RemoveAllGoals()
	SetObjectForKeyedSubscript(weight objc.IObject /* cross-framework: NSNumber */, goal IGKGoal)
	SetWeightForGoal(weight float32, goal IGKGoal)
	ObjectForKeyedSubscript(goal IGKGoal) foundation.Number
	ObjectAtIndexedSubscript(idx uint) IGoal
	WeightForGoal(goal IGKGoal) float32
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Behavior */
// Alloc allocates a new instance without initialization.
func (bc _BehaviorClass) Alloc() Behavior {
	rv := objc.Send[Behavior](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Behavior */
// A set of goals that together influence the movement of an agent.
//
// By combining multiple goals ( objects) you can create complex behavior, such as groups of agents ( objects) that move together naturally. To assign a set of goals to an agent, use its property. To learn more about using goals and agents, see in .


// A set of goals that together influence the movement of an agent.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Behavior */

// Creates a behavior with a single goal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goal:weight:)
func NewBehaviorWithGoalWeight(goal IGKGoal, weight float32) Behavior {
	rv := objc.Send[Behavior](objc.ID(getBehaviorClass().class), objc.Sel("behaviorWithGoal:weight:"), goal, weight)
	return rv
}/* debug [class_init_methods/constructor]: NewBehaviorWithGoalWeight */


// Creates a behavior with the specified goals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goals:)
func NewBehaviorWithGoals(goals []Goal) Behavior {
	rv := objc.Send[Behavior](objc.ID(getBehaviorClass().class), objc.Sel("behaviorWithGoals:"), goals)
	return rv
}/* debug [class_init_methods/constructor]: NewBehaviorWithGoals */


// Creates a behavior with the specified goals and weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goals:andWeights:)
func NewBehaviorWithGoalsAndWeights(goals []Goal, weights []foundation.Number) Behavior {
	rv := objc.Send[Behavior](objc.ID(getBehaviorClass().class), objc.Sel("behaviorWithGoals:andWeights:"), goals, weights)
	return rv
}/* debug [class_init_methods/constructor]: NewBehaviorWithGoalsAndWeights */


// Creates a behavior with the specified mapping of goals to their weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(weightedGoals:)
func NewBehaviorWithWeightedGoals(weightedGoals foundation.IDictionary) Behavior {
	rv := objc.Send[Behavior](objc.ID(getBehaviorClass().class), objc.Sel("behaviorWithWeightedGoals:"), weightedGoals)
	return rv
}/* debug [class_init_methods/constructor]: NewBehaviorWithWeightedGoals */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Behavior */

// Creates a behavior with a single goal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goal:weight:)
func (bc _BehaviorClass) BehaviorWithGoalWeight(goal IGKGoal, weight float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("behaviorWithGoal:weight:"), goal, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BehaviorWithGoalWeight) */


// Creates a behavior with the specified goals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goals:)
func (bc _BehaviorClass) BehaviorWithGoals(goals []Goal) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("behaviorWithGoals:"), goals)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BehaviorWithGoals) */


// Creates a behavior with the specified goals and weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(goals:andWeights:)
func (bc _BehaviorClass) BehaviorWithGoalsAndWeights(goals []Goal, weights []foundation.Number) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("behaviorWithGoals:andWeights:"), goals, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BehaviorWithGoalsAndWeights) */


// Creates a behavior with the specified mapping of goals to their weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/init(weightedGoals:)
func (bc _BehaviorClass) BehaviorWithWeightedGoals(weightedGoals foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("behaviorWithWeightedGoals:"), weightedGoals)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BehaviorWithWeightedGoals) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Behavior */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Behavior */

// Removes the specified goal from the behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/remove(_:)
func (b_ Behavior) RemoveGoal(goal IGKGoal) {
	objc.Send[objc.ID](b_.ID, objc.Sel("removeGoal:"), goal)
}/* debug [instance_methods/method]: RemoveGoal */


// Removes all goals from the behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/removeAllGoals()
func (b_ Behavior) RemoveAllGoals() {
	objc.Send[objc.ID](b_.ID, objc.Sel("removeAllGoals"))
}/* debug [instance_methods/method]: RemoveAllGoals */


// Sets the weight for the goal specified by subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/setObject:forKeyedSubscript:
func (b_ Behavior) SetObjectForKeyedSubscript(weight objc.IObject /* cross-framework: NSNumber */, goal IGKGoal) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setObject:forKeyedSubscript:"), weight, goal)
}/* debug [instance_methods/method]: SetObjectForKeyedSubscript */


// Sets the weight for the specified goal’s influence on agents, adding that goal to the behavior if not already present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/setWeight(_:for:)
func (b_ Behavior) SetWeightForGoal(weight float32, goal IGKGoal) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setWeight:forGoal:"), weight, goal)
}/* debug [instance_methods/method]: SetWeightForGoal */


// Returns the weight associated with the goal specified by subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/subscript(_:)-2yvko
func (b_ Behavior) ObjectForKeyedSubscript(goal IGKGoal) foundation.Number {
	rv := objc.Send[foundation.Number](b_.ID, objc.Sel("objectForKeyedSubscript:"), goal)
	return rv
}/* debug [instance_methods/method]: ObjectForKeyedSubscript */


// Returns the goal at the specified index in the behavior’s list of goals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/subscript(_:)-997a9
func (b_ Behavior) ObjectAtIndexedSubscript(idx uint) IGoal {
	rv := objc.Send[Goal](b_.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */


// Returns the weight for the specified goal’s influence on agents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/weight(for:)
func (b_ Behavior) WeightForGoal(goal IGKGoal) float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("weightForGoal:"), goal)
	return rv
}/* debug [instance_methods/method]: WeightForGoal */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Behavior */

// The number of goals in the behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBehavior/goalCount
func (b_ Behavior) GoalCount() int {
	rv := objc.Send[int](b_.ID, objc.Sel("goalCount"))
	return rv
}/* debug [instance_properties/getter]: goalCount */


// A weighted collection of goals that influence the agent’s movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkagent/behavior
func (b_ Behavior) Behavior() IGKBehavior {
	rv := objc.Send[Behavior](b_.ID, objc.Sel("behavior"))
	return rv
}/* debug [instance_properties/getter]: behavior */


// A weighted collection of goals that influence the agent’s movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkagent/behavior
func (b_ Behavior) SetBehavior(value IGKBehavior) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBehavior:"), value)
}/* debug [instance_properties/setter]: behavior */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKBehavior */


