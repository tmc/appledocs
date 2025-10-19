// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RuleEditor] class.
var (
	ruleEditorClass     _RuleEditorClass
	ruleEditorClassOnce sync.Once
)

func getRuleEditorClass() _RuleEditorClass {
	ruleEditorClassOnce.Do(func() {
		ruleEditorClass = _RuleEditorClass{objc.GetClass("NSRuleEditor")}
	})
	return ruleEditorClass
}

type _RuleEditorClass struct {
	class objc.Class
}

// An interface definition for the [RuleEditor] class.
type IRuleEditor interface {
	IControl
}

// An interface for configuring a rule-based list of options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor

type RuleEditor struct {
	Control
}

// RuleEditorFrom constructs a [RuleEditor] from an unsafe.Pointer.
//
// An interface for configuring a rule-based list of options.
func RuleEditorFrom(ptr unsafe.Pointer) RuleEditor {
	return RuleEditor{
		Control: ControlFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (rc _RuleEditorClass) Alloc() RuleEditor {
	rv := objc.Send[RuleEditor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RuleEditorClass) New() RuleEditor {
	rv := objc.Send[RuleEditor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RuleEditor) Init() RuleEditor {
	rv := objc.Send[RuleEditor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RuleEditor) Autorelease() RuleEditor {
	rv := objc.Send[RuleEditor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRuleEditor creates a new RuleEditor instance.
func NewRuleEditor() RuleEditor {
	return getRuleEditorClass().New()
}




