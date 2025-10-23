// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CompositeBehavior] class.
var (
	CompositeBehaviorClass     _CompositeBehaviorClass
	CompositeBehaviorClassOnce sync.Once
)

func getCompositeBehaviorClass() _CompositeBehaviorClass {
	CompositeBehaviorClassOnce.Do(func() {
		CompositeBehaviorClass = _CompositeBehaviorClass{objc.GetClass("GKCompositeBehavior")}
	})
	return CompositeBehaviorClass
}

type _CompositeBehaviorClass struct {
	class objc.Class
}

// An interface definition for the [CompositeBehavior] class.
type ICompositeBehavior interface {
	IBehavior
	RemoveBehavior(behavior GKBehavior)
	RemoveAllBehaviors()
	SetObjectForKeyedSubscript(weight foundation.INumber, behavior GKBehavior)
	SetWeightForBehavior(weight float32, behavior GKBehavior)
	ObjectForKeyedSubscript(behavior GKBehavior) foundation.Number
	ObjectAtIndexedSubscript(idx uint) Behavior
	WeightForBehavior(behavior GKBehavior) float32
	BehaviorCount() int
}

// A set of behaviors, each of which is a set of goals, that together influence the movement of an agent.
//
// By composing objects into subgroups ( objects) and composing those behaviors into composite behaviors, you can control certain aspects of a object’s movement in concert. To assign a behavior to an agent, use its property. For example, you might create a behavior for a set of agents to stay together as a flock (with cohesion, alignment, and separation goals) while loosely following a path. With a single object, whenever you want to change the importance of the flocking goals relative to the path-following goals, you’d need to individually change the weight of each goal. With a composite behavior, you can adjust the relative influence of a group of goals together, as in the following code. After constructing this behavior, you can use the method to increase or decrease the influence of the and behaviors relative to one another. To learn more about using goals and agents, see in .


// A set of behaviors, each of which is a set of goals, that together influence the movement of an agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior
type CompositeBehavior struct {
	Behavior
}

// CompositeBehaviorFrom constructs a [CompositeBehavior] from an unsafe.Pointer.
//
// A set of behaviors, each of which is a set of goals, that together influence the movement of an agent.
func CompositeBehaviorFrom(ptr unsafe.Pointer) CompositeBehavior {
	return CompositeBehavior{
		Behavior: BehaviorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CompositeBehaviorClass) Alloc() CompositeBehavior {
	rv := objc.Send[CompositeBehavior](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CompositeBehaviorClass) New() CompositeBehavior {
	rv := objc.Send[CompositeBehavior](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CompositeBehavior) Init() CompositeBehavior {
	rv := objc.Send[CompositeBehavior](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CompositeBehavior) Autorelease() CompositeBehavior {
	rv := objc.Send[CompositeBehavior](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompositeBehavior creates a new CompositeBehavior instance.
func NewCompositeBehavior() CompositeBehavior {
	return getCompositeBehaviorClass().New()
}



// Creates a composite behavior from the specified individual behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/init(behaviors:)
func NewCompositeBehaviorWithBehaviors(behaviors []Behavior) CompositeBehavior {
	rv := objc.Send[CompositeBehavior](objc.ID(getCompositeBehaviorClass().class), objc.Sel("behaviorWithBehaviors:"), behaviors)
	return rv
}


// Creates a behavior with the specified behaviors and weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/init(behaviors:andWeights:)
func NewCompositeBehaviorWithBehaviorsAndWeights(behaviors []Behavior, weights []foundation.INumber) CompositeBehavior {
	rv := objc.Send[CompositeBehavior](objc.ID(getCompositeBehaviorClass().class), objc.Sel("behaviorWithBehaviors:andWeights:"), behaviors, weights)
	return rv
}



// Creates a composite behavior from the specified individual behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/init(behaviors:)
func (cc _CompositeBehaviorClass) BehaviorWithBehaviors(behaviors []Behavior) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("behaviorWithBehaviors:"), behaviors)
	return rv
}


// Creates a behavior with the specified behaviors and weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/init(behaviors:andWeights:)
func (cc _CompositeBehaviorClass) BehaviorWithBehaviorsAndWeights(behaviors []Behavior, weights []foundation.INumber) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("behaviorWithBehaviors:andWeights:"), behaviors, weights)
	return rv
}


// Removes the specified individual behavior from the composite behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/remove(_:)
func (c_ CompositeBehavior) RemoveBehavior(behavior GKBehavior) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeBehavior:"), behavior)
}


// Removes all individual behaviors from the composite behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/removeAllBehaviors()
func (c_ CompositeBehavior) RemoveAllBehaviors() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllBehaviors"))
}


// Sets the weight for the behavior specified by subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/setObject:forKeyedSubscript:
func (c_ CompositeBehavior) SetObjectForKeyedSubscript(weight foundation.INumber, behavior GKBehavior) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObject:forKeyedSubscript:"), weight, behavior)
}


// Sets the weight for the specified individual behavior’s influence on agents, adding that behavior to the composite behavior if it is not already present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/setWeight(_:for:)
func (c_ CompositeBehavior) SetWeightForBehavior(weight float32, behavior GKBehavior) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeight:forBehavior:"), weight, behavior)
}


// Returns the weight associated with the behavior specified by subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/subscript(_:)-6jng9
func (c_ CompositeBehavior) ObjectForKeyedSubscript(behavior GKBehavior) foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("objectForKeyedSubscript:"), behavior)
	return rv
}


// Returns the individual behavior at the specified index in the composite behavior’s list of behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/subscript(_:)-6krdg
func (c_ CompositeBehavior) ObjectAtIndexedSubscript(idx uint) Behavior {
	rv := objc.Send[Behavior](c_.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}


// Returns the weight for the specified individual behavior’s influence on agents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/weight(for:)
func (c_ CompositeBehavior) WeightForBehavior(behavior GKBehavior) float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("weightForBehavior:"), behavior)
	return rv
}


// The number of individual behaviors in the composite behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/behaviorCount
func (c_ CompositeBehavior) BehaviorCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("behaviorCount"))
	return rv
}


