// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LocalizedNumberFormatRule] class.
var localizedNumberFormatRuleClass = _LocalizedNumberFormatRuleClass{objc.GetClass("NSLocalizedNumberFormatRule")}

type _LocalizedNumberFormatRuleClass struct {
	class objc.Class
}

// An interface definition for the [LocalizedNumberFormatRule] class.
type ILocalizedNumberFormatRule interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocalizedNumberFormatRule

type LocalizedNumberFormatRule struct {
	objectivec.Object
}

// LocalizedNumberFormatRuleFrom constructs a [LocalizedNumberFormatRule] from an unsafe.Pointer.
func LocalizedNumberFormatRuleFrom(ptr unsafe.Pointer) LocalizedNumberFormatRule {
	return LocalizedNumberFormatRule{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (lc _LocalizedNumberFormatRuleClass) Alloc() LocalizedNumberFormatRule {
	rv := objc.Send[LocalizedNumberFormatRule](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (lc _LocalizedNumberFormatRuleClass) New() LocalizedNumberFormatRule {
	rv := objc.Send[LocalizedNumberFormatRule](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LocalizedNumberFormatRule) Init() LocalizedNumberFormatRule {
	rv := objc.Send[LocalizedNumberFormatRule](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LocalizedNumberFormatRule) Autorelease() LocalizedNumberFormatRule {
	rv := objc.Send[LocalizedNumberFormatRule](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocalizedNumberFormatRule creates a new LocalizedNumberFormatRule instance.
func NewLocalizedNumberFormatRule() LocalizedNumberFormatRule {
	return localizedNumberFormatRuleClass.New()
}




