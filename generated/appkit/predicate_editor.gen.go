// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PredicateEditor] class.
var predicateEditorClass = _PredicateEditorClass{objc.GetClass("NSPredicateEditor")}

type _PredicateEditorClass struct {
	class objc.Class
}

// An interface definition for the [PredicateEditor] class.
type IPredicateEditor interface {
	IRuleEditor
}

// A defined set of rules that allows the editing of predicate objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditor

type PredicateEditor struct {
	RuleEditor
}

// PredicateEditorFrom constructs a [PredicateEditor] from an unsafe.Pointer.
//
// A defined set of rules that allows the editing of predicate objects.
func PredicateEditorFrom(ptr unsafe.Pointer) PredicateEditor {
	return PredicateEditor{
		RuleEditor: RuleEditorFrom(ptr),
	}
}



