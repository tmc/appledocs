// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKGoal */


/* debug [class_header]: Header for GKGoal */
// The class instance for the [Goal] class.
var (
	GoalClass     _GoalClass
	GoalClassOnce sync.Once
)

func getGoalClass() _GoalClass {
	GoalClassOnce.Do(func() {
		GoalClass = _GoalClass{objc.GetClass("GKGoal")}
	})
	return GoalClass
}

type _GoalClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Goal */
// An interface definition for the [Goal] class.
type IGoal interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Goal */
	// properties:
	Behavior() IGKBehavior
	SetBehavior(value IGKBehavior)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Goal */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Goal */
// Alloc allocates a new instance without initialization.
func (gc _GoalClass) Alloc() Goal {
	rv := objc.Send[Goal](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GoalClass) New() Goal {
	rv := objc.Send[Goal](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Goal) Init() Goal {
	rv := objc.Send[Goal](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Goal) Autorelease() Goal {
	rv := objc.Send[Goal](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGoal creates a new Goal instance.
func NewGoal() Goal {
	return getGoalClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Goal */
// An influence that motivates the movement of one or more agents.
//
// Goals can motivate agents ( objects) to actions such as moving toward a target, following a path, or staying aligned with a group of other agents. To give an agent one or more goals, combine those goals in a object (which includes weights for the relative influence of each goal) and assign that object to the agent’s property. Each time an agent’s method runs, the agent evaluates each goal in its behavior to find the change in direction and speed necessary to move toward fulfilling that goal (within the limits of the time delta and the agent’s maximum speed and turn rate). It then combines the effects from all the goals in its behavior, using the weights in the behavior to modulate the influence of each goal, to produce a total change in its direction and speed. To learn more about using goals and agents, see in .


// An influence that motivates the movement of one or more agents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal
type Goal struct {
	objectivec.Object
}

// GoalFrom constructs a [Goal] from an unsafe.Pointer.
//
// An influence that motivates the movement of one or more agents.
func GoalFrom(ptr unsafe.Pointer) Goal {
	return Goal{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Goal */

// Creates a goal whose effect is to make an agent align its orientation with that of other agents in a specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAlignWith:maxDistance:maxAngle:)
func NewGoalToAlignWithAgentsMaxDistanceMaxAngle(agents []Agent, maxDistance float32, maxAngle float32) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToAlignWithAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToAlignWithAgentsMaxDistanceMaxAngle */


// Creates a goal whose effect is to make an agent avoid colliding with the specified other agents, taking into account the other agents’ movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAvoid:maxPredictionTime:)-96a0i
func NewGoalToAvoidAgentsMaxPredictionTime(agents []Agent, maxPredictionTime float64) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToAvoidAgents:maxPredictionTime:"), agents, maxPredictionTime)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToAvoidAgentsMaxPredictionTime */


// Creates a goal whose effect is to make an agent avoid colliding with the specified static obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAvoid:maxPredictionTime:)-7oslq
func NewGoalToAvoidObstaclesMaxPredictionTime(obstacles []Obstacle, maxPredictionTime float64) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToAvoidObstacles:maxPredictionTime:"), obstacles, maxPredictionTime)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToAvoidObstaclesMaxPredictionTime */


// Creates a goal whose effect is to make an agent stay near the other agents in a specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toCohereWith:maxDistance:maxAngle:)
func NewGoalToCohereWithAgentsMaxDistanceMaxAngle(agents []Agent, maxDistance float32, maxAngle float32) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToCohereWithAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToCohereWithAgentsMaxDistanceMaxAngle */


// Creates a goal whose effect is to move an agent away from the current position of the specified other agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toFleeAgent:)
func NewGoalToFleeAgent(agent IGKAgent) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToFleeAgent:"), agent)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToFleeAgent */


// Creates a goal whose effect is to both maintain position on and traverse the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toFollow:maxPredictionTime:forward:)
func NewGoalToFollowPathMaxPredictionTimeForward(path IGKPath, maxPredictionTime float64, forward bool) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToFollowPath:maxPredictionTime:forward:"), path, maxPredictionTime, forward)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToFollowPathMaxPredictionTimeForward */


// Creates a goal whose effect is to make an agent pursue the specified other agent, taking into account the target’s movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toInterceptAgent:maxPredictionTime:)
func NewGoalToInterceptAgentMaxPredictionTime(target IGKAgent, maxPredictionTime float64) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToInterceptAgent:maxPredictionTime:"), target, maxPredictionTime)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToInterceptAgentMaxPredictionTime */


// Creates a goal whose effect is to accelerate or decelerate an agent until it reaches the specified speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toReachTargetSpeed:)
func NewGoalToReachTargetSpeed(targetSpeed float32) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToReachTargetSpeed:"), targetSpeed)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToReachTargetSpeed */


// Creates a goal whose effect is to move an agent toward the current position of the specified other agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toSeekAgent:)
func NewGoalToSeekAgent(agent IGKAgent) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToSeekAgent:"), agent)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToSeekAgent */


// Creates a goal whose effect is to make an agent maintain the specified distance from other agents in a specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toSeparateFrom:maxDistance:maxAngle:)
func NewGoalToSeparateFromAgentsMaxDistanceMaxAngle(agents []Agent, maxDistance float32, maxAngle float32) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToSeparateFromAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToSeparateFromAgentsMaxDistanceMaxAngle */


// Creates a goal whose effect is to maintain an agent’s position within the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toStayOn:maxPredictionTime:)
func NewGoalToStayOnPathMaxPredictionTime(path IGKPath, maxPredictionTime float64) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToStayOnPath:maxPredictionTime:"), path, maxPredictionTime)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToStayOnPathMaxPredictionTime */


// Creates a goal whose effect is to make an agent wander aimlessly, moving forward and turning at random.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toWander:)
func NewGoalToWander(speed float32) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToWander:"), speed)
	return rv
}/* debug [class_init_methods/constructor]: NewGoalToWander */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Goal */

// Creates a goal whose effect is to make an agent align its orientation with that of other agents in a specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAlignWith:maxDistance:maxAngle:)
func (gc _GoalClass) GoalToAlignWithAgentsMaxDistanceMaxAngle(agents []Agent, maxDistance float32, maxAngle float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToAlignWithAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToAlignWithAgentsMaxDistanceMaxAngle) */


// Creates a goal whose effect is to make an agent avoid colliding with the specified static obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAvoid:maxPredictionTime:)-7oslq
func (gc _GoalClass) GoalToAvoidObstaclesMaxPredictionTime(obstacles []Obstacle, maxPredictionTime float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToAvoidObstacles:maxPredictionTime:"), obstacles, maxPredictionTime)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToAvoidObstaclesMaxPredictionTime) */


// Creates a goal whose effect is to make an agent avoid colliding with the specified other agents, taking into account the other agents’ movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAvoid:maxPredictionTime:)-96a0i
func (gc _GoalClass) GoalToAvoidAgentsMaxPredictionTime(agents []Agent, maxPredictionTime float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToAvoidAgents:maxPredictionTime:"), agents, maxPredictionTime)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToAvoidAgentsMaxPredictionTime) */


// Creates a goal whose effect is to make an agent stay near the other agents in a specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toCohereWith:maxDistance:maxAngle:)
func (gc _GoalClass) GoalToCohereWithAgentsMaxDistanceMaxAngle(agents []Agent, maxDistance float32, maxAngle float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToCohereWithAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToCohereWithAgentsMaxDistanceMaxAngle) */


// Creates a goal whose effect is to move an agent away from the current position of the specified other agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toFleeAgent:)
func (gc _GoalClass) GoalToFleeAgent(agent IGKAgent) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToFleeAgent:"), agent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToFleeAgent) */


// Creates a goal whose effect is to both maintain position on and traverse the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toFollow:maxPredictionTime:forward:)
func (gc _GoalClass) GoalToFollowPathMaxPredictionTimeForward(path IGKPath, maxPredictionTime float64, forward bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToFollowPath:maxPredictionTime:forward:"), path, maxPredictionTime, forward)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToFollowPathMaxPredictionTimeForward) */


// Creates a goal whose effect is to make an agent pursue the specified other agent, taking into account the target’s movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toInterceptAgent:maxPredictionTime:)
func (gc _GoalClass) GoalToInterceptAgentMaxPredictionTime(target IGKAgent, maxPredictionTime float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToInterceptAgent:maxPredictionTime:"), target, maxPredictionTime)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToInterceptAgentMaxPredictionTime) */


// Creates a goal whose effect is to accelerate or decelerate an agent until it reaches the specified speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toReachTargetSpeed:)
func (gc _GoalClass) GoalToReachTargetSpeed(targetSpeed float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToReachTargetSpeed:"), targetSpeed)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToReachTargetSpeed) */


// Creates a goal whose effect is to move an agent toward the current position of the specified other agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toSeekAgent:)
func (gc _GoalClass) GoalToSeekAgent(agent IGKAgent) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToSeekAgent:"), agent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToSeekAgent) */


// Creates a goal whose effect is to make an agent maintain the specified distance from other agents in a specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toSeparateFrom:maxDistance:maxAngle:)
func (gc _GoalClass) GoalToSeparateFromAgentsMaxDistanceMaxAngle(agents []Agent, maxDistance float32, maxAngle float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToSeparateFromAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToSeparateFromAgentsMaxDistanceMaxAngle) */


// Creates a goal whose effect is to maintain an agent’s position within the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toStayOn:maxPredictionTime:)
func (gc _GoalClass) GoalToStayOnPathMaxPredictionTime(path IGKPath, maxPredictionTime float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToStayOnPath:maxPredictionTime:"), path, maxPredictionTime)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToStayOnPathMaxPredictionTime) */


// Creates a goal whose effect is to make an agent wander aimlessly, moving forward and turning at random.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toWander:)
func (gc _GoalClass) GoalToWander(speed float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("goalToWander:"), speed)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GoalToWander) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Goal */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Goal */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Goal */

// A weighted collection of goals that influence the agent’s movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkagent/behavior
func (g_ Goal) Behavior() IGKBehavior {
	rv := objc.Send[Behavior](g_.ID, objc.Sel("behavior"))
	return rv
}/* debug [instance_properties/getter]: behavior */


// A weighted collection of goals that influence the agent’s movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkagent/behavior
func (g_ Goal) SetBehavior(value IGKBehavior) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBehavior:"), value)
}/* debug [instance_properties/setter]: behavior */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGoal */


