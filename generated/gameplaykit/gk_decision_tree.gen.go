// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DecisionTree] class.
var (
	DecisionTreeClass     _DecisionTreeClass
	DecisionTreeClassOnce sync.Once
)

func getDecisionTreeClass() _DecisionTreeClass {
	DecisionTreeClassOnce.Do(func() {
		DecisionTreeClass = _DecisionTreeClass{objc.GetClass("GKDecisionTree")}
	})
	return DecisionTreeClass
}

type _DecisionTreeClass struct {
	class objc.Class
}

// An interface definition for the [DecisionTree] class.
type IDecisionTree interface {
	objectivec.IObject
	// properties:
	RandomSource() IGKRandomSource
	SetRandomSource(value IGKRandomSource)
	RootNode() IGKDecisionNode
	Description() string /* primitive/slice/pointer. */
	SetDescription(value string /* primitive/slice/pointer. */)
	// methods:
	ExportToURLError(url foundation.objc.IObject /* cross-framework URL */, error_ Error /* not a class type */) bool /* primitive/slice/pointer. */
	FindActionForAnswers(answers foundation.IDictionary /* already interface */) objc.ID
}

// A data structure that models a set of specific questions, their possible answers, and the actions that follow from a series of answers.
//
// You can define a decision tree manually, by specifying questions, answers, and actions, or you can allow the class to automatically learn a predictive model based on example data. A decision tree has several elements: represent individual questions to be answered or choices to be made. are the possible answers to the questions or choices posed by each attribute. are the final outcomes of the tree’s decision-making process. Each branch from an attribute leads either to another attribute or to an action. When you use the class, attributes and actions can be any object type relevant to your app or game. You can define branches for specific answer values, using predicates, or with weights that influence a random decision. For example, a strategy combat game might use a decision tree to choose what a character should do on its turn, based on several criteria about the match in progress. In this case: For attributes, you might use (non-user-visible) strings that represent those criteria, such as (what type of enemy is the character’s opponent?), (how much health does the opponent have remaining?), and (is the character’s special move available for use?). For branches, you’d use an appropriate style for each attribute. The attribute might have a branch for each possible enemy type, but the attribute could use predicates to determine whether the enemy’s health is above or below a certain threshold value. For actions, you might define your own enumerated type representing the kinds of attacks the character can choose (such as , , and ). Alternately, you might use instances of your own custom classes representing items or spells available to the character. illustrates a possible tree structure based on the above example attributes, branches, and actions.


// A data structure that models a set of specific questions, their possible answers, and the actions that follow from a series of answers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionTree
type DecisionTree struct {
	objectivec.Object
}

// DecisionTreeFrom constructs a [DecisionTree] from an unsafe.Pointer.
//
// A data structure that models a set of specific questions, their possible answers, and the actions that follow from a series of answers.
func DecisionTreeFrom(ptr unsafe.Pointer) DecisionTree {
	return DecisionTree{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DecisionTreeClass) Alloc() DecisionTree {
	rv := objc.Send[DecisionTree](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DecisionTreeClass) New() DecisionTree {
	rv := objc.Send[DecisionTree](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DecisionTree) Init() DecisionTree {
	rv := objc.Send[DecisionTree](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DecisionTree) Autorelease() DecisionTree {
	rv := objc.Send[DecisionTree](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDecisionTree creates a new DecisionTree instance.
func NewDecisionTree() DecisionTree {
	return getDecisionTreeClass().New()
}



// Creates a decision tree starting with the specified initial attribute to test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionTree/init(attribute:)
func NewDecisionTreeWithAttribute(attribute objectivec.IObject) DecisionTree {
	instance := getDecisionTreeClass().Alloc()
	rv := objc.Send[DecisionTree](instance.ID, objc.Sel("initWithAttribute:"), attribute)
	rv.Autorelease()
	return rv
}


// Creates an automatically learned decision tree using the specified attributes, example items, and actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionTree/init(examples:actions:attributes:)
func NewDecisionTreeWithExamplesActionsAttributes(examples []foundation.objc.IObject /* cross-framework Array */, actions []objc.ID /* already interface */, attributes []objc.ID /* already interface */) DecisionTree {
	instance := getDecisionTreeClass().Alloc()
	rv := objc.Send[DecisionTree](instance.ID, objc.Sel("initWithExamples:actions:attributes:"), examples, actions, attributes)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionTree/init(url:error:)
func NewDecisionTreeWithURLError(url foundation.objc.IObject /* cross-framework URL */, error_ Error /* not a class type */) DecisionTree {
	instance := getDecisionTreeClass().Alloc()
	rv := objc.Send[DecisionTree](instance.ID, objc.Sel("initWithURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionTree/export(to:error:)
func (d_ DecisionTree) ExportToURLError(url foundation.objc.IObject /* cross-framework URL */, error_ Error /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("exportToURL:error:"), url, error_)
	return rv
}


// Searches the decision tree, following the branches corresponding to each of the specified answers, and returns the resulting action object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionTree/findAction(forAnswers:)
func (d_ DecisionTree) FindActionForAnswers(answers foundation.IDictionary /* already interface */) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("findActionForAnswers:"), answers)
	return rv
}


// The randomizer to be used when evaluating parts of the tree that branch randomly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionTree/randomSource
func (d_ DecisionTree) RandomSource() IGKRandomSource {
	rv := objc.Send[RandomSource](d_.ID, objc.Sel("randomSource"))
	return rv
}


// The randomizer to be used when evaluating parts of the tree that branch randomly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionTree/randomSource
func (d_ DecisionTree) SetRandomSource(value IGKRandomSource) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRandomSource:"), value)
}


// The decision node at the root of the decision tree, representing the first attribute to test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKDecisionTree/rootNode
func (d_ DecisionTree) RootNode() IGKDecisionNode {
	rv := objc.Send[DecisionNode](d_.ID, objc.Sel("rootNode"))
	return rv
}


// A textual representation of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description
func (d_ DecisionTree) Description() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("description"))
	return rv
}


// A textual representation of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description
func (d_ DecisionTree) SetDescription(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDescription:"), objc.String(value))
}


