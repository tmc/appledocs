// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKCompositeBehavior */


/* debug [class_header]: Header for GKCompositeBehavior */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CompositeBehavior */
// An interface definition for the [CompositeBehavior] class.
type ICompositeBehavior interface {
	IBehavior
	
/* debug [class_interface_properties]: Properties for CompositeBehavior */
	// properties:
	BehaviorCount() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CompositeBehavior */
	// methods:
	RemoveBehavior(behavior IGKBehavior)
	RemoveAllBehaviors()
	SetObjectForKeyedSubscript(weight objc.IObject /* cross-framework: NSNumber */, behavior IGKBehavior)
	SetWeightForBehavior(weight float32, behavior IGKBehavior)
	ObjectForKeyedSubscript(behavior IGKBehavior) foundation.Number
	ObjectAtIndexedSubscript(idx uint) IBehavior
	WeightForBehavior(behavior IGKBehavior) float32
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CompositeBehavior */
// Alloc allocates a new instance without initialization.
func (cc _CompositeBehaviorClass) Alloc() CompositeBehavior {
	rv := objc.Send[CompositeBehavior](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CompositeBehavior */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CompositeBehavior */

// Creates a composite behavior from the specified individual behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/init(behaviors:)
func NewCompositeBehaviorWithBehaviors(behaviors []Behavior) CompositeBehavior {
	rv := objc.Send[CompositeBehavior](objc.ID(getCompositeBehaviorClass().class), objc.Sel("behaviorWithBehaviors:"), behaviors)
	return rv
}/* debug [class_init_methods/constructor]: NewCompositeBehaviorWithBehaviors */


// Creates a behavior with the specified behaviors and weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/init(behaviors:andWeights:)
func NewCompositeBehaviorWithBehaviorsAndWeights(behaviors []Behavior, weights []foundation.Number) CompositeBehavior {
	rv := objc.Send[CompositeBehavior](objc.ID(getCompositeBehaviorClass().class), objc.Sel("behaviorWithBehaviors:andWeights:"), behaviors, weights)
	return rv
}/* debug [class_init_methods/constructor]: NewCompositeBehaviorWithBehaviorsAndWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CompositeBehavior */

// Creates a composite behavior from the specified individual behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/init(behaviors:)
func (cc _CompositeBehaviorClass) BehaviorWithBehaviors(behaviors []Behavior) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("behaviorWithBehaviors:"), behaviors)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BehaviorWithBehaviors) */


// Creates a behavior with the specified behaviors and weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/init(behaviors:andWeights:)
func (cc _CompositeBehaviorClass) BehaviorWithBehaviorsAndWeights(behaviors []Behavior, weights []foundation.Number) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("behaviorWithBehaviors:andWeights:"), behaviors, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BehaviorWithBehaviorsAndWeights) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CompositeBehavior */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CompositeBehavior */

// Removes the specified individual behavior from the composite behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/remove(_:)
func (c_ CompositeBehavior) RemoveBehavior(behavior IGKBehavior) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeBehavior:"), behavior)
}/* debug [instance_methods/method]: RemoveBehavior */


// Removes all individual behaviors from the composite behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/removeAllBehaviors()
func (c_ CompositeBehavior) RemoveAllBehaviors() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllBehaviors"))
}/* debug [instance_methods/method]: RemoveAllBehaviors */


// Sets the weight for the behavior specified by subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/setObject:forKeyedSubscript:
func (c_ CompositeBehavior) SetObjectForKeyedSubscript(weight objc.IObject /* cross-framework: NSNumber */, behavior IGKBehavior) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObject:forKeyedSubscript:"), weight, behavior)
}/* debug [instance_methods/method]: SetObjectForKeyedSubscript */


// Sets the weight for the specified individual behavior’s influence on agents, adding that behavior to the composite behavior if it is not already present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/setWeight(_:for:)
func (c_ CompositeBehavior) SetWeightForBehavior(weight float32, behavior IGKBehavior) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeight:forBehavior:"), weight, behavior)
}/* debug [instance_methods/method]: SetWeightForBehavior */


// Returns the weight associated with the behavior specified by subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/subscript(_:)-6jng9
func (c_ CompositeBehavior) ObjectForKeyedSubscript(behavior IGKBehavior) foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("objectForKeyedSubscript:"), behavior)
	return rv
}/* debug [instance_methods/method]: ObjectForKeyedSubscript */


// Returns the individual behavior at the specified index in the composite behavior’s list of behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/subscript(_:)-6krdg
func (c_ CompositeBehavior) ObjectAtIndexedSubscript(idx uint) IBehavior {
	rv := objc.Send[Behavior](c_.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */


// Returns the weight for the specified individual behavior’s influence on agents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/weight(for:)
func (c_ CompositeBehavior) WeightForBehavior(behavior IGKBehavior) float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("weightForBehavior:"), behavior)
	return rv
}/* debug [instance_methods/method]: WeightForBehavior */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CompositeBehavior */

// The number of individual behaviors in the composite behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCompositeBehavior/behaviorCount
func (c_ CompositeBehavior) BehaviorCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("behaviorCount"))
	return rv
}/* debug [instance_properties/getter]: behaviorCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKCompositeBehavior */


