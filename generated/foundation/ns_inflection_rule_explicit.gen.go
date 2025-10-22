// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InflectionRuleExplicit] class.
var (
	InflectionRuleExplicitClass     _InflectionRuleExplicitClass
	InflectionRuleExplicitClassOnce sync.Once
)

func getInflectionRuleExplicitClass() _InflectionRuleExplicitClass {
	InflectionRuleExplicitClassOnce.Do(func() {
		InflectionRuleExplicitClass = _InflectionRuleExplicitClass{objc.GetClass("NSInflectionRuleExplicit")}
	})
	return InflectionRuleExplicitClass
}

type _InflectionRuleExplicitClass struct {
	class objc.Class
}

// An interface definition for the [InflectionRuleExplicit] class.
type IInflectionRuleExplicit interface {
	IInflectionRule
}

// An inflection rule that uses a morphology instance to determine how to inflect attribued strings.


// An inflection rule that uses a morphology instance to determine how to inflect attribued strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit
type InflectionRuleExplicit struct {
	InflectionRule
}

// InflectionRuleExplicitFrom constructs a [InflectionRuleExplicit] from an unsafe.Pointer.
//
// An inflection rule that uses a morphology instance to determine how to inflect attribued strings.
func InflectionRuleExplicitFrom(ptr unsafe.Pointer) InflectionRuleExplicit {
	return InflectionRuleExplicit{
		InflectionRule: InflectionRuleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InflectionRuleExplicitClass) Alloc() InflectionRuleExplicit {
	rv := objc.Send[InflectionRuleExplicit](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InflectionRuleExplicitClass) New() InflectionRuleExplicit {
	rv := objc.Send[InflectionRuleExplicit](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InflectionRuleExplicit) Init() InflectionRuleExplicit {
	rv := objc.Send[InflectionRuleExplicit](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InflectionRuleExplicit) Autorelease() InflectionRuleExplicit {
	rv := objc.Send[InflectionRuleExplicit](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInflectionRuleExplicit creates a new InflectionRuleExplicit instance.
func NewInflectionRuleExplicit() InflectionRuleExplicit {
	return getInflectionRuleExplicitClass().New()
}




