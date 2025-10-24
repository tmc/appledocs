// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InflectionRule] class.
var (
	InflectionRuleClass     _InflectionRuleClass
	InflectionRuleClassOnce sync.Once
)

func getInflectionRuleClass() _InflectionRuleClass {
	InflectionRuleClassOnce.Do(func() {
		InflectionRuleClass = _InflectionRuleClass{objc.GetClass("NSInflectionRule")}
	})
	return InflectionRuleClass
}

type _InflectionRuleClass struct {
	class objc.Class
}

// An interface definition for the [InflectionRule] class.
type IInflectionRule interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A rule that affects how an attributed string performs automatic grammatical agreement.
//
// Most apps can rely on loading localized strings to perform automatic grammar agreement. Typically, your app’s strings files use the Markdown extension syntax to indicate portions of the string that may require inflection to agree grammatically. This transformation occurs when you load the attributed string with methods like . However, if the system lacks information about the words in the string, you may need to apply an inflection rule programmatically. For example, a social networking app may have gender information about other users that you want to apply at runtime. When performing manual inflection at runtime, you use an inflection rule to indicate to the system what portions of a string should be automatically edited, and what to match. Add the attribute with an on an , then call to perform the grammar agreement and produce an edited string.


// A rule that affects how an attributed string performs automatic grammatical agreement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRule
type InflectionRule struct {
	objectivec.Object
}

// InflectionRuleFrom constructs a [InflectionRule] from an unsafe.Pointer.
//
// A rule that affects how an attributed string performs automatic grammatical agreement.
func InflectionRuleFrom(ptr unsafe.Pointer) InflectionRule {
	return InflectionRule{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _InflectionRuleClass) Alloc() InflectionRule {
	rv := objc.Send[InflectionRule](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InflectionRuleClass) New() InflectionRule {
	rv := objc.Send[InflectionRule](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InflectionRule) Init() InflectionRule {
	rv := objc.Send[InflectionRule](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InflectionRule) Autorelease() InflectionRule {
	rv := objc.Send[InflectionRule](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInflectionRule creates a new InflectionRule instance.
func NewInflectionRule() InflectionRule {
	return getInflectionRuleClass().New()
}



// Returns a Boolean value that indicates whether the rule can inflect a given language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRule/canInflectLanguage:
func (ic _InflectionRuleClass) CanInflectLanguage(language IString) bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("canInflectLanguage:"), language)
	return rv
}


// An inflection rule that performs automatic grammar agreement with default transformations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRule/automaticRule
func (ic _InflectionRuleClass) AutomaticRule() InflectionRule {
	rv := objc.Send[InflectionRule](objc.ID(ic.class), objc.Sel("automaticRule"))
	return rv
}

// A Boolean value that indicates whether the rule can inflect the user’s current preferred localization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRule/canInflectPreferredLocalization
func (ic _InflectionRuleClass) CanInflectPreferredLocalization() bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("canInflectPreferredLocalization"))
	return rv
}

// An inflection rule that performs automatic grammar agreement with default transformations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRule/automaticRule
func (i_ InflectionRule) AutomaticRule() IInflectionRule {
	rv := objc.Send[InflectionRule](i_.ID, objc.Sel("automaticRule"))
	return rv
}


// A Boolean value that indicates whether the rule can inflect the user’s current preferred localization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRule/canInflectPreferredLocalization
func (i_ InflectionRule) CanInflectPreferredLocalization() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canInflectPreferredLocalization"))
	return rv
}



