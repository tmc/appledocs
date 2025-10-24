// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKDecisionNode */


/* debug [class_header]: Header for GKDecisionNode */
// The class instance for the [DecisionNode] class.
var (
	DecisionNodeClass     _DecisionNodeClass
	DecisionNodeClassOnce sync.Once
)

func getDecisionNodeClass() _DecisionNodeClass {
	DecisionNodeClassOnce.Do(func() {
		DecisionNodeClass = _DecisionNodeClass{objc.GetClass("GKDecisionNode")}
	})
	return DecisionNodeClass
}

type _DecisionNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DecisionNode */
// An interface definition for the [DecisionNode] class.
type IDecisionNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DecisionNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DecisionNode */
	// methods:
	CreateBranchWithPredicateAttribute(predicate foundation.Predicate, attribute unsafe.Pointer) objectivec.IObject
	CreateBranchWithValueAttribute(value objc.IObject /* cross-framework: NSNumber */, attribute unsafe.Pointer) objectivec.IObject
	CreateBranchWithWeightAttribute(weight int, attribute unsafe.Pointer) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DecisionNode */
// Alloc allocates a new instance without initialization.
func (dc _DecisionNodeClass) Alloc() DecisionNode {
	rv := objc.Send[DecisionNode](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DecisionNodeClass) New() DecisionNode {
	rv := objc.Send[DecisionNode](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DecisionNode) Init() DecisionNode {
	rv := objc.Send[DecisionNode](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DecisionNode) Autorelease() DecisionNode {
	rv := objc.Send[DecisionNode](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDecisionNode creates a new DecisionNode instance.
func NewDecisionNode() DecisionNode {
	return getDecisionNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DecisionNode */
// A node for use in manually creating decision trees, representing a specific question and possible answers, or an action that follows from answering other questions.
//
// A instance represents an element in a decision tree (a object). Decision trees contain two kinds of nodes. Some nodes, including the tree’s root node, represent individual decisions to be made (also called a question or ) and reference child nodes for each possible outcome of (or from) that decision. Each branch can lead to another question node, or to a leaf node—nodes that have no branches represent a final outcome (or ) to result from the tree’s decision-making process. After creating a decision tree from a set of nodes, you can present the tree with a set of inputs (values for attributes, or answers to questions) and the tree provides a final action that follows from the branches corresponding to each attribute. There are two ways to create a decision tree. You use the class directly only when you want to define an entire decision tree manually—that is, to specify each question, the possible branches from each question, and the possible final actions. To create such a decision tree, start with the initializer, then use the methods listed in Creating Child Nodes for Decision Branches to add branches to the tree. To instead automatically learn a decision tree given a set of questions and example answers, use the method.


// A node for use in manually creating decision trees, representing a specific question and possible answers, or an action that follows from answering other questions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionNode
type DecisionNode struct {
	objectivec.Object
}

// DecisionNodeFrom constructs a [DecisionNode] from an unsafe.Pointer.
//
// A node for use in manually creating decision trees, representing a specific question and possible answers, or an action that follows from answering other questions.
func DecisionNodeFrom(ptr unsafe.Pointer) DecisionNode {
	return DecisionNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DecisionNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DecisionNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DecisionNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DecisionNode */

// Creates a child node that the decision tree should use when the current node’s attribute satisfies the specified predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionNode/createBranch(predicate:attribute:)
func (d_ DecisionNode) CreateBranchWithPredicateAttribute(predicate foundation.Predicate, attribute unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("createBranchWithPredicate:attribute:"), predicate, attribute)
	return rv
}/* debug [instance_methods/method]: CreateBranchWithPredicateAttribute */


// Creates a child node that the decision tree should use when the current node’s attribute has the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionNode/createBranch(value:attribute:)
func (d_ DecisionNode) CreateBranchWithValueAttribute(value objc.IObject /* cross-framework: NSNumber */, attribute unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("createBranchWithValue:attribute:"), value, attribute)
	return rv
}/* debug [instance_methods/method]: CreateBranchWithValueAttribute */


// Creates a child node that the decision tree should use as the result of a random choice, biased by the specified weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionNode/createBranch(weight:attribute:)
func (d_ DecisionNode) CreateBranchWithWeightAttribute(weight int, attribute unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("createBranchWithWeight:attribute:"), weight, attribute)
	return rv
}/* debug [instance_methods/method]: CreateBranchWithWeightAttribute */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DecisionNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKDecisionNode */



