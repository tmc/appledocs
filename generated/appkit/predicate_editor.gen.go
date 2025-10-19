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
// Alloc allocates a new instance without initialization.
func (pc _PredicateEditorClass) Alloc() PredicateEditor {
	rv := objc.Send[PredicateEditor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PredicateEditorClass) New() PredicateEditor {
	rv := objc.Send[PredicateEditor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PredicateEditor) Init() PredicateEditor {
	rv := objc.Send[PredicateEditor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PredicateEditor) Autorelease() PredicateEditor {
	rv := objc.Send[PredicateEditor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPredicateEditor creates a new PredicateEditor instance.
func NewPredicateEditor() PredicateEditor {
	return predicateEditorClass.New()
}




