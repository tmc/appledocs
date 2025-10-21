// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RuleSystem] class.
var (
	RuleSystemClass     _RuleSystemClass
	RuleSystemClassOnce sync.Once
)

func getRuleSystemClass() _RuleSystemClass {
	RuleSystemClassOnce.Do(func() {
		RuleSystemClass = _RuleSystemClass{objc.GetClass("GKRuleSystem")}
	})
	return RuleSystemClass
}

type _RuleSystemClass struct {
	class objc.Class
}

// An interface definition for the [RuleSystem] class.
type IRuleSystem interface {
	objectivec.IObject
	AddRule(rule unsafe.Pointer)
	AddRulesFromArray(rules unsafe.Pointer)
	AssertFact(fact objc.ID)
	AssertFactGrade(fact objc.ID, grade unsafe.Pointer)
	Evaluate()
	GradeForFact(fact objc.ID) unsafe.Pointer
	MaximumGradeForFacts(facts objc.ID) unsafe.Pointer
	MinimumGradeForFacts(facts objc.ID) unsafe.Pointer
	RemoveAllRules()
	Reset()
	RetractFact(fact objc.ID)
	RetractFactGrade(fact objc.ID, grade unsafe.Pointer)
}

// A list of rules, together with a context for evaluating them and interpreting results, for use in constructing data-driven logic or fuzzy logic systems.
//
// A object manages a list of rules ( objects). A rule system also offers methods for evaluating its list of rules in a context defined by two features: a dictionary containing information to be tested by rules, and a set of representing the conclusions drawn as a result of rule evaluation. You can evaluate facts based on a binary truth state—that is, a fact either is or is not in the set—or on a continuously variable membership grade, representing different levels of veracity, confidence, or strength for use in fuzzy logic. You construct a rule system by creating objects and adding them to the system’s list of rules. There are multiple ways to construct rules: for greater reusability, use the methods listed in Creating Data-Driven Rules; or for greater flexibility, use the method or create a custom subclass of or . Then, add rules to the system with the methods listed in Managing a System’s List of Rules below. To evaluate a system, call the method. This method processes each rule in the system in the order it appears in the system’s list. You set this order with the property of each rule, or with the order in which you add rules to the system. As the system processes each rule, it tests the rule’s method to determine whether the rule is satisfied in the context of the system. If the rule’s predicate is satisfied, the system executes the rule’s method and moves the rule to the list (so the further evaluation of the agenda doesn’t repeatedly trigger the rule’s action). Rules typically use the system’s dictionary as input and its set of as output. (However, more complex systems can include sets of rules whose predicates test facts or whose actions mutate the system’s state.) After evaluating a rule system, you can examine the set of facts it has produced using the methods listed in Drawing Conclusions from Facts below. You can then use the presence of a fact in the set, the value of its membership grade, or the combined membership grades of a group of facts to influence the behaviors in your game. For more information about rules and rule systems, read in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem
type RuleSystem struct {
	objectivec.Object
}

// RuleSystemFrom constructs a [RuleSystem] from an unsafe.Pointer.
//
// A list of rules, together with a context for evaluating them and interpreting results, for use in constructing data-driven logic or fuzzy logic systems.
func RuleSystemFrom(ptr unsafe.Pointer) RuleSystem {
	return RuleSystem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RuleSystemClass) Alloc() RuleSystem {
	rv := objc.Send[RuleSystem](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RuleSystemClass) New() RuleSystem {
	rv := objc.Send[RuleSystem](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RuleSystem) Init() RuleSystem {
	rv := objc.Send[RuleSystem](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RuleSystem) Autorelease() RuleSystem {
	rv := objc.Send[RuleSystem](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRuleSystem creates a new RuleSystem instance.
func NewRuleSystem() RuleSystem {
	return getRuleSystemClass().New()
}



// Adds the specified rule to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/add(_:)-76jb5
func (r_ RuleSystem) AddRule(rule unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addRule:"), rule)
}

// Adds the specified list of rules to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/add(_:)-7u5zw
func (r_ RuleSystem) AddRulesFromArray(rules unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addRulesFromArray:"), rules)
}

// Adds the specified fact to the fact set with a membership grade of 1.0, and reevaluates the rules in the system’s agenda.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/assertFact(_:)
func (r_ RuleSystem) AssertFact(fact objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("assertFact:"), fact)
}

// Increases the membership grade of the specified fact by the specified amount, adding it to the fact set if necessary, and reevaluates the rules in the system’s agenda.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/assertFact(_:grade:)
func (r_ RuleSystem) AssertFactGrade(fact objc.ID, grade unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("assertFact:grade:"), fact, grade)
}

// Evaluates the rule system, executing the list of rules in its agenda.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/evaluate()
func (r_ RuleSystem) Evaluate() {
	objc.Send[objc.ID](r_.ID, objc.Sel("evaluate"))
}

// Returns the membership grade of the specified fact.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/grade(forFact:)
func (r_ RuleSystem) GradeForFact(fact objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("gradeForFact:"), fact)
	return rv
}

// Returns the highest membership grade among the specified facts.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/maximumGrade(forFacts:)
func (r_ RuleSystem) MaximumGradeForFacts(facts objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("maximumGradeForFacts:"), facts)
	return rv
}

// Returns the lowest membership grade among the specified facts.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/minimumGrade(forFacts:)
func (r_ RuleSystem) MinimumGradeForFacts(facts objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("minimumGradeForFacts:"), facts)
	return rv
}

// Removes all rules from the system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/removeAllRules()
func (r_ RuleSystem) RemoveAllRules() {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeAllRules"))
}

// Returns the rule system to its original agenda and clears all facts.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/reset()
func (r_ RuleSystem) Reset() {
	objc.Send[objc.ID](r_.ID, objc.Sel("reset"))
}

// Removes the specified fact from the fact set, and reevaluates the rules in the system’s agenda.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/retractFact(_:)
func (r_ RuleSystem) RetractFact(fact objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("retractFact:"), fact)
}

// Reduces the membership grade of the specified fact by the specified amount, removing it from the fact set if necessary, and reevaluates the rules in the system’s agenda.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/retractFact(_:grade:)
func (r_ RuleSystem) RetractFactGrade(fact objc.ID, grade unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("retractFact:grade:"), fact, grade)
}

// The list of rules to be considered when evaluating the system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/agenda
func (r_ RuleSystem) Agenda() []Rule {
	rv := objc.Send[[]Rule](r_.ID, objc.Sel("agenda"))
	return rv
}

// The list of rules whose actions have been performed during evaluation of the system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/executed
func (r_ RuleSystem) Executed() []Rule {
	rv := objc.Send[[]Rule](r_.ID, objc.Sel("executed"))
	return rv
}

// The list of facts claimed by the rule system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/facts
func (r_ RuleSystem) Facts() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("facts"))
	return rv
}

// The list of rules to be executed when evaluating the system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/rules
func (r_ RuleSystem) Rules() []Rule {
	rv := objc.Send[[]Rule](r_.ID, objc.Sel("rules"))
	return rv
}

// A dictionary of state information to be evaluated by the system’s rules.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRuleSystem/state
func (r_ RuleSystem) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("state"))
	return rv
}


