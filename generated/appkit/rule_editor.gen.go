// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RuleEditor] class.
var ruleEditorClass = _RuleEditorClass{objc.GetClass("NSRuleEditor")}

type _RuleEditorClass struct {
	class objc.Class
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



