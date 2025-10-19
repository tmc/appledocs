// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InflectionRule] class.
var inflectionRuleClass = _InflectionRuleClass{objc.GetClass("NSInflectionRule")}

type _InflectionRuleClass struct {
	class objc.Class
}

// An interface definition for the [InflectionRule] class.
type IInflectionRule interface {
	objectivec.IObject
}

// A rule that affects how an attributed string performs automatic grammatical agreement. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return inflectionRuleClass.New()
}




