// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Goal] class.
type IGoal interface {
	objectivec.IObject
}

// An influence that motivates the movement of one or more agents.
//
// Goals can motivate agents ( objects) to actions such as moving toward a target, following a path, or staying aligned with a group of other agents. To give an agent one or more goals, combine those goals in a object (which includes weights for the relative influence of each goal) and assign that object to the agent’s property. Each time an agent’s method runs, the agent evaluates each goal in its behavior to find the change in direction and speed necessary to move toward fulfilling that goal (within the limits of the time delta and the agent’s maximum speed and turn rate). It then combines the effects from all the goals in its behavior, using the weights in the behavior to modulate the influence of each goal, to produce a total change in its direction and speed. To learn more about using goals and agents, see in .
//
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

// Alloc allocates a new instance without initialization.
func (gc _GoalClass) Alloc() Goal {
	rv := objc.Send[Goal](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a goal whose effect is to make an agent align its orientation with that of other agents in a specified group.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAlignWith:maxDistance:maxAngle:)
func NewGoalToAlignWithAgentsMaxDistanceMaxAngle(agents unsafe.Pointer, maxDistance unsafe.Pointer, maxAngle unsafe.Pointer) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToAlignWithAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}



// Creates a goal whose effect is to make an agent avoid colliding with the specified other agents, taking into account the other agents’ movement.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAvoid:maxPredictionTime:)-96a0i
func NewGoalToAvoidAgentsMaxPredictionTime(agents unsafe.Pointer, maxPredictionTime foundation.TimeInterval) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToAvoidAgents:maxPredictionTime:"), agents, maxPredictionTime)
	return rv
}



// Creates a goal whose effect is to make an agent avoid colliding with the specified static obstacles.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAvoid:maxPredictionTime:)-7oslq
func NewGoalToAvoidObstaclesMaxPredictionTime(obstacles unsafe.Pointer, maxPredictionTime foundation.TimeInterval) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToAvoidObstacles:maxPredictionTime:"), obstacles, maxPredictionTime)
	return rv
}



// Creates a goal whose effect is to make an agent stay near the other agents in a specified group.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toCohereWith:maxDistance:maxAngle:)
func NewGoalToCohereWithAgentsMaxDistanceMaxAngle(agents unsafe.Pointer, maxDistance unsafe.Pointer, maxAngle unsafe.Pointer) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToCohereWithAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}



// Creates a goal whose effect is to move an agent away from the current position of the specified other agent.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toFleeAgent:)
func NewGoalToFleeAgent(agent unsafe.Pointer) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToFleeAgent:"), agent)
	return rv
}



// Creates a goal whose effect is to both maintain position on and traverse the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toFollow:maxPredictionTime:forward:)
func NewGoalToFollowPathMaxPredictionTimeForward(path unsafe.Pointer, maxPredictionTime foundation.TimeInterval, forward bool) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToFollowPath:maxPredictionTime:forward:"), path, maxPredictionTime, forward)
	return rv
}



// Creates a goal whose effect is to make an agent pursue the specified other agent, taking into account the target’s movement.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toInterceptAgent:maxPredictionTime:)
func NewGoalToInterceptAgentMaxPredictionTime(target unsafe.Pointer, maxPredictionTime foundation.TimeInterval) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToInterceptAgent:maxPredictionTime:"), target, maxPredictionTime)
	return rv
}



// Creates a goal whose effect is to accelerate or decelerate an agent until it reaches the specified speed.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toReachTargetSpeed:)
func NewGoalToReachTargetSpeed(targetSpeed unsafe.Pointer) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToReachTargetSpeed:"), targetSpeed)
	return rv
}



// Creates a goal whose effect is to move an agent toward the current position of the specified other agent.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toSeekAgent:)
func NewGoalToSeekAgent(agent unsafe.Pointer) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToSeekAgent:"), agent)
	return rv
}



// Creates a goal whose effect is to make an agent maintain the specified distance from other agents in a specified group.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toSeparateFrom:maxDistance:maxAngle:)
func NewGoalToSeparateFromAgentsMaxDistanceMaxAngle(agents unsafe.Pointer, maxDistance unsafe.Pointer, maxAngle unsafe.Pointer) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToSeparateFromAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}



// Creates a goal whose effect is to maintain an agent’s position within the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toStayOn:maxPredictionTime:)
func NewGoalToStayOnPathMaxPredictionTime(path unsafe.Pointer, maxPredictionTime foundation.TimeInterval) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToStayOnPath:maxPredictionTime:"), path, maxPredictionTime)
	return rv
}



// Creates a goal whose effect is to make an agent wander aimlessly, moving forward and turning at random.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toWander:)
func NewGoalToWander(speed unsafe.Pointer) Goal {
	rv := objc.Send[Goal](objc.ID(getGoalClass().class), objc.Sel("goalToWander:"), speed)
	return rv
}


// Creates a goal whose effect is to make an agent align its orientation with that of other agents in a specified group.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAlignWith:maxDistance:maxAngle:)
func (gc _GoalClass) GoalToAlignWithAgentsMaxDistanceMaxAngle(agents unsafe.Pointer, maxDistance unsafe.Pointer, maxAngle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToAlignWithAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}

// Creates a goal whose effect is to make an agent avoid colliding with the specified static obstacles.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAvoid:maxPredictionTime:)-7oslq
func (gc _GoalClass) GoalToAvoidObstaclesMaxPredictionTime(obstacles unsafe.Pointer, maxPredictionTime foundation.TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToAvoidObstacles:maxPredictionTime:"), obstacles, maxPredictionTime)
	return rv
}

// Creates a goal whose effect is to make an agent avoid colliding with the specified other agents, taking into account the other agents’ movement.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toAvoid:maxPredictionTime:)-96a0i
func (gc _GoalClass) GoalToAvoidAgentsMaxPredictionTime(agents unsafe.Pointer, maxPredictionTime foundation.TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToAvoidAgents:maxPredictionTime:"), agents, maxPredictionTime)
	return rv
}

// Creates a goal whose effect is to make an agent stay near the other agents in a specified group.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toCohereWith:maxDistance:maxAngle:)
func (gc _GoalClass) GoalToCohereWithAgentsMaxDistanceMaxAngle(agents unsafe.Pointer, maxDistance unsafe.Pointer, maxAngle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToCohereWithAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}

// Creates a goal whose effect is to move an agent away from the current position of the specified other agent.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toFleeAgent:)
func (gc _GoalClass) GoalToFleeAgent(agent unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToFleeAgent:"), agent)
	return rv
}

// Creates a goal whose effect is to both maintain position on and traverse the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toFollow:maxPredictionTime:forward:)
func (gc _GoalClass) GoalToFollowPathMaxPredictionTimeForward(path unsafe.Pointer, maxPredictionTime foundation.TimeInterval, forward bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToFollowPath:maxPredictionTime:forward:"), path, maxPredictionTime, forward)
	return rv
}

// Creates a goal whose effect is to make an agent pursue the specified other agent, taking into account the target’s movement.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toInterceptAgent:maxPredictionTime:)
func (gc _GoalClass) GoalToInterceptAgentMaxPredictionTime(target unsafe.Pointer, maxPredictionTime foundation.TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToInterceptAgent:maxPredictionTime:"), target, maxPredictionTime)
	return rv
}

// Creates a goal whose effect is to accelerate or decelerate an agent until it reaches the specified speed.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toReachTargetSpeed:)
func (gc _GoalClass) GoalToReachTargetSpeed(targetSpeed unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToReachTargetSpeed:"), targetSpeed)
	return rv
}

// Creates a goal whose effect is to move an agent toward the current position of the specified other agent.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toSeekAgent:)
func (gc _GoalClass) GoalToSeekAgent(agent unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToSeekAgent:"), agent)
	return rv
}

// Creates a goal whose effect is to make an agent maintain the specified distance from other agents in a specified group.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toSeparateFrom:maxDistance:maxAngle:)
func (gc _GoalClass) GoalToSeparateFromAgentsMaxDistanceMaxAngle(agents unsafe.Pointer, maxDistance unsafe.Pointer, maxAngle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToSeparateFromAgents:maxDistance:maxAngle:"), agents, maxDistance, maxAngle)
	return rv
}

// Creates a goal whose effect is to maintain an agent’s position within the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toStayOn:maxPredictionTime:)
func (gc _GoalClass) GoalToStayOnPathMaxPredictionTime(path unsafe.Pointer, maxPredictionTime foundation.TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToStayOnPath:maxPredictionTime:"), path, maxPredictionTime)
	return rv
}

// Creates a goal whose effect is to make an agent wander aimlessly, moving forward and turning at random.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGoal/init(toWander:)
func (gc _GoalClass) GoalToWander(speed unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("goalToWander:"), speed)
	return rv
}

// A weighted collection of goals that influence the agent’s movement.
//
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkagent/behavior
func (g_ Goal) Behavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("behavior"))
	return rv
}


// SetBehavior sets the value of the behavior property.
// A weighted collection of goals that influence the agent’s movement.

//
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkagent/behavior
func (g_ Goal) SetBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBehavior:"), value)
}


