
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RuleEditor] class.
var RuleEditorClass _RuleEditorClass

func init() {
	RuleEditorClass = _RuleEditorClass{objc.GetClass("NSRuleEditor")}
}

type _RuleEditorClass struct {
	objc.Class
}

// An interface definition for the [RuleEditor] class.
type IRuleEditor interface {
	ID() objc.ID
}

type RuleEditor struct {
	id objc.ID
}

func RuleEditorFrom(ptr unsafe.Pointer) RuleEditor {
	return RuleEditor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ RuleEditor) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _RuleEditorClass) Alloc() RuleEditor {
	rv := objc.Send[RuleEditor](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _RuleEditorClass) New() RuleEditor {
	rv := objc.Send[RuleEditor](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewRuleEditor creates and returns a new initialized instance.
func NewRuleEditor() RuleEditor {
	return RuleEditorClass.New()
}

// Init initializes the instance.
func (r_ RuleEditor) Init() RuleEditor {
	rv := objc.Send[RuleEditor](r_.ID(), selInit)
	return rv
}
