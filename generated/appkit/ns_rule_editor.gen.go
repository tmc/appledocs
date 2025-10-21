// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RuleEditor] class.
var (
	RuleEditorClass     _RuleEditorClass
	RuleEditorClassOnce sync.Once
)

func getRuleEditorClass() _RuleEditorClass {
	RuleEditorClassOnce.Do(func() {
		RuleEditorClass = _RuleEditorClass{objc.GetClass("NSRuleEditor")}
	})
	return RuleEditorClass
}

type _RuleEditorClass struct {
	class objc.Class
}

// An interface definition for the [RuleEditor] class.
type IRuleEditor interface {
	IControl
}

// An interface for configuring a rule-based list of options.
//
// A rule editor lets the user visually create and configure a list of options that are expressed as a predicate (as described in ). Each row displayed by the rule editor represents a particular path down a tree of choices. The rule editor’s delegate provides the tree of choices to be displayed. The rule editor presents those choices to the user as a row of popup buttons, static text fields, and custom views. exposes one binding, . You can bind to an ordered collection (such as an instance of ). Each object in the collection should have the following properties:
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


// The key path for the subrows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/subrowsKeyPath
func (r_ RuleEditor) SubrowsKeyPath() string {
	rv := objc.Send[string](r_.ID, objc.Sel("subrowsKeyPath"))
	return rv
}


// SetSubrowsKeyPath sets the value of the subrowsKeyPath property.
// The key path for the subrows.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/subrowsKeyPath
func (r_ RuleEditor) SetSubrowsKeyPath(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSubrowsKeyPath:"), objc.String(value))
}


