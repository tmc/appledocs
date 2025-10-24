// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKRule */


/* debug [class_header]: Header for GKRule */
// The class instance for the [Rule] class.
var (
	RuleClass     _RuleClass
	RuleClassOnce sync.Once
)

func getRuleClass() _RuleClass {
	RuleClassOnce.Do(func() {
		RuleClass = _RuleClass{objc.GetClass("GKRule")}
	})
	return RuleClass
}

type _RuleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Rule */
// An interface definition for the [Rule] class.
type IRule interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Rule */
	// properties:
	Salience() int
	SetSalience(value int)
	Facts() objectivec.IObject
	SetFacts(value objectivec.IObject)
	State() foundation.MutableDictionary
	SetState(value foundation.MutableDictionary)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Rule */
	// methods:
	EvaluatePredicateWithSystem(system IGKRuleSystem) bool
	PerformActionWithSystem(system IGKRuleSystem)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Rule */
// Alloc allocates a new instance without initialization.
func (rc _RuleClass) Alloc() Rule {
	rv := objc.Send[Rule](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RuleClass) New() Rule {
	rv := objc.Send[Rule](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Rule) Init() Rule {
	rv := objc.Send[Rule](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Rule) Autorelease() Rule {
	rv := objc.Send[Rule](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRule creates a new Rule instance.
func NewRule() Rule {
	return getRuleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Rule */
// A rule to be used in the context of a rule system, with a predicate to be tested and an action to be executed when the test succeeds.
//
// Evaluating a object tests each of its rules, which typically examine the state or facts associated with the rule system, and executes the actions specified by each rule whose test passes, such as asserting or retracting facts in the rule system or modifying its state. A rule has two parts: a predicate and an action. The rule’s determines whether the rule has been satisfied, within the context of a given rule system. Evaluating a rule’s predicate typically involves examining information in the rule sytem’s dictionary or testing the membership grade of facts claimed by the system (see the property in for details). The rule’s is executed if and only if the rule’s predicate is satisfied. Rule actions typically involve asserting or retracting facts in the system (see the methods listed in Asserting and Retracting Facts) or modifying information in the system’s dictionary. There are multiple ways to create rules for use in a rule system, each with its own advantages. Typical rule predicates involve conditional logic tests on the properties of the containing rule system, and typical rule actions assert or retract facts. If your rules fit this pattern, you can use the and methods to create rules that are entirely data-driven—that is, they can be easily archived for later reuse, edited without compiling source code, and created at runtime. To create rules with entirely custom logic for both predicate and action, use the method. This method creates rules that are very flexible, but that cannot be archived for reuse. To create rules with more complex custom logic, implement your own rule classes: subclass to build custom logic for both the rule’s predicate and its action, or subclass to use an object for the rule’s predicate and build custom logic only for the rule’s action. The reusability of custom rule classes depends on your implementation of such classes. For more information about rules and rule systems, read in .


// A rule to be used in the context of a rule system, with a predicate to be tested and an action to be executed when the test succeeds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule
type Rule struct {
	objectivec.Object
}

// RuleFrom constructs a [Rule] from an unsafe.Pointer.
//
// A rule to be used in the context of a rule system, with a predicate to be tested and an action to be executed when the test succeeds.
func RuleFrom(ptr unsafe.Pointer) Rule {
	return Rule{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Rule */

// Creates a rule whose predicate is evaluated and action is executed through the specified blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/init(blockPredicate:action:)
func NewRuleWithBlockPredicateAction(predicate unsafe.Pointer, action unsafe.Pointer) Rule {
	rv := objc.Send[Rule](objc.ID(getRuleClass().class), objc.Sel("ruleWithBlockPredicate:action:"), predicate, action)
	return rv
}/* debug [class_init_methods/constructor]: NewRuleWithBlockPredicateAction */


// Creates a data-driven rule with the specified predicate, whose action asserts a fact in the rule system evaluating the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/init(predicate:assertingFact:grade:)
func NewRuleWithPredicateAssertingFactGrade(predicate foundation.Predicate, fact unsafe.Pointer, grade float32) Rule {
	rv := objc.Send[Rule](objc.ID(getRuleClass().class), objc.Sel("ruleWithPredicate:assertingFact:grade:"), predicate, fact, grade)
	return rv
}/* debug [class_init_methods/constructor]: NewRuleWithPredicateAssertingFactGrade */


// Creates a data-driven rule with the specified predicate, whose action retracts a fact in the rule system evaluating the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/init(predicate:retractingFact:grade:)
func NewRuleWithPredicateRetractingFactGrade(predicate foundation.Predicate, fact unsafe.Pointer, grade float32) Rule {
	rv := objc.Send[Rule](objc.ID(getRuleClass().class), objc.Sel("ruleWithPredicate:retractingFact:grade:"), predicate, fact, grade)
	return rv
}/* debug [class_init_methods/constructor]: NewRuleWithPredicateRetractingFactGrade */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Rule */

// Creates a rule whose predicate is evaluated and action is executed through the specified blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/init(blockPredicate:action:)
func (rc _RuleClass) RuleWithBlockPredicateAction(predicate unsafe.Pointer, action unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("ruleWithBlockPredicate:action:"), predicate, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RuleWithBlockPredicateAction) */


// Creates a data-driven rule with the specified predicate, whose action asserts a fact in the rule system evaluating the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/init(predicate:assertingFact:grade:)
func (rc _RuleClass) RuleWithPredicateAssertingFactGrade(predicate foundation.Predicate, fact unsafe.Pointer, grade float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("ruleWithPredicate:assertingFact:grade:"), predicate, fact, grade)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RuleWithPredicateAssertingFactGrade) */


// Creates a data-driven rule with the specified predicate, whose action retracts a fact in the rule system evaluating the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/init(predicate:retractingFact:grade:)
func (rc _RuleClass) RuleWithPredicateRetractingFactGrade(predicate foundation.Predicate, fact unsafe.Pointer, grade float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("ruleWithPredicate:retractingFact:grade:"), predicate, fact, grade)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RuleWithPredicateRetractingFactGrade) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Rule */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Rule */

// Returns a Boolean value indicating whether the rule has been satisfied in the context of the specified rule system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/evaluatePredicate(in:)
func (r_ Rule) EvaluatePredicateWithSystem(system IGKRuleSystem) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("evaluatePredicateWithSystem:"), system)
	return rv
}/* debug [instance_methods/method]: EvaluatePredicateWithSystem */


// Performs actions that should result when the rule is satisfied in the context of the specified rule system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/performAction(in:)
func (r_ Rule) PerformActionWithSystem(system IGKRuleSystem) {
	objc.Send[objc.ID](r_.ID, objc.Sel("performActionWithSystem:"), system)
}/* debug [instance_methods/method]: PerformActionWithSystem */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Rule */

// The importance of the rule relative to others in a rule system’s agenda.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/salience
func (r_ Rule) Salience() int {
	rv := objc.Send[int](r_.ID, objc.Sel("salience"))
	return rv
}/* debug [instance_properties/getter]: salience */


// The importance of the rule relative to others in a rule system’s agenda.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRule/salience
func (r_ Rule) SetSalience(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSalience:"), value)
}/* debug [instance_properties/setter]: salience */


// The list of facts claimed by the rule system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkrulesystem/facts
func (r_ Rule) Facts() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("facts"))
	return rv
}/* debug [instance_properties/getter]: facts */


// The list of facts claimed by the rule system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkrulesystem/facts
func (r_ Rule) SetFacts(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFacts:"), value)
}/* debug [instance_properties/setter]: facts */


// A dictionary of state information to be evaluated by the system’s rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkrulesystem/state
func (r_ Rule) State() foundation.MutableDictionary {
	rv := objc.Send[foundation.MutableDictionary](r_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// A dictionary of state information to be evaluated by the system’s rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkrulesystem/state
func (r_ Rule) SetState(value foundation.MutableDictionary) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKRule */


