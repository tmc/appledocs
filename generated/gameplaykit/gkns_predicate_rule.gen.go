// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NSPredicateRule] class.
var (
	NSPredicateRuleClass     _NSPredicateRuleClass
	NSPredicateRuleClassOnce sync.Once
)

func getNSPredicateRuleClass() _NSPredicateRuleClass {
	NSPredicateRuleClassOnce.Do(func() {
		NSPredicateRuleClass = _NSPredicateRuleClass{objc.GetClass("GKNSPredicateRule")}
	})
	return NSPredicateRuleClass
}

type _NSPredicateRuleClass struct {
	class objc.Class
}

// An interface definition for the [NSPredicateRule] class.
type INSPredicateRule interface {
	IRule
	EvaluatePredicateWithSystem(system unsafe.Pointer) bool
}

// A rule for use in a rule system that uses a Foundation object to evaluate itself.
//
// The class is a specialized subclass of the class (which represents rules to be used by objects). Custom subclasses of use an object to evaluate a rule, rather than requiring custom logic for evaluation as is the case with custom subclasses. For more information about rules and rule systems, read in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKNSPredicateRule
type NSPredicateRule struct {
	Rule
}

// NSPredicateRuleFrom constructs a [NSPredicateRule] from an unsafe.Pointer.
//
// A rule for use in a rule system that uses a Foundation object to evaluate itself.
func NSPredicateRuleFrom(ptr unsafe.Pointer) NSPredicateRule {
	return NSPredicateRule{
		Rule: RuleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _NSPredicateRuleClass) Alloc() NSPredicateRule {
	rv := objc.Send[NSPredicateRule](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _NSPredicateRuleClass) New() NSPredicateRule {
	rv := objc.Send[NSPredicateRule](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ NSPredicateRule) Init() NSPredicateRule {
	rv := objc.Send[NSPredicateRule](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ NSPredicateRule) Autorelease() NSPredicateRule {
	rv := objc.Send[NSPredicateRule](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPredicateRule creates a new NSPredicateRule instance.
func NewNSPredicateRule() NSPredicateRule {
	return getNSPredicateRuleClass().New()
}




// Initializes a rule with the specified predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKNSPredicateRule/init(predicate:)
func NewNSPredicateRuleWithPredicate(predicate unsafe.Pointer) NSPredicateRule {
	instance := getNSPredicateRuleClass().Alloc()
	rv := objc.Send[NSPredicateRule](instance.ID, objc.Sel("initWithPredicate:"), predicate)
	rv.Autorelease()
	return rv
}


// Returns a Boolean value indicating whether the rule’s predicate has been satisfied in the context of the specified rule system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKNSPredicateRule/evaluatePredicate(in:)
func (p_ NSPredicateRule) EvaluatePredicateWithSystem(system unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("evaluatePredicateWithSystem:"), system)
	return rv
}

// A predicate to be tested when evaluating the rule.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKNSPredicateRule/predicate
func (p_ NSPredicateRule) Predicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("predicate"))
	return rv
}


