// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [DecisionNode] class.
type IDecisionNode interface {
	objectivec.IObject
	CreateBranchWithPredicateAttribute(predicate foundation.IPredicate, attribute objectivec.IObject) unsafe.Pointer
	CreateBranchWithValueAttribute(value foundation.INumber, attribute objectivec.IObject) unsafe.Pointer
	CreateBranchWithWeightAttribute(weight int, attribute objectivec.IObject) unsafe.Pointer
}

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

// Alloc allocates a new instance without initialization.
func (dc _DecisionNodeClass) Alloc() DecisionNode {
	rv := objc.Send[DecisionNode](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a child node that the decision tree should use when the current node’s attribute satisfies the specified predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionNode/createBranch(predicate:attribute:)
func (d_ DecisionNode) CreateBranchWithPredicateAttribute(predicate foundation.IPredicate, attribute objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("createBranchWithPredicate:attribute:"), predicate, attribute)
	return rv
}


// Creates a child node that the decision tree should use when the current node’s attribute has the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionNode/createBranch(value:attribute:)
func (d_ DecisionNode) CreateBranchWithValueAttribute(value foundation.INumber, attribute objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("createBranchWithValue:attribute:"), value, attribute)
	return rv
}


// Creates a child node that the decision tree should use as the result of a random choice, biased by the specified weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionNode/createBranch(weight:attribute:)
func (d_ DecisionNode) CreateBranchWithWeightAttribute(weight int, attribute objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("createBranchWithWeight:attribute:"), weight, attribute)
	return rv
}



