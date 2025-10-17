
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PredicateEditor] class.
var PredicateEditorClass _PredicateEditorClass

func init() {
	PredicateEditorClass = _PredicateEditorClass{objc.GetClass("NSPredicateEditor")}
}

type _PredicateEditorClass struct {
	objc.Class
}

// An interface definition for the [PredicateEditor] class.
type IPredicateEditor interface {
	ID() objc.ID
}

type PredicateEditor struct {
	id objc.ID
}

func PredicateEditorFrom(ptr unsafe.Pointer) PredicateEditor {
	return PredicateEditor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PredicateEditor) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PredicateEditorClass) Alloc() PredicateEditor {
	rv := objc.Send[PredicateEditor](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PredicateEditorClass) New() PredicateEditor {
	rv := objc.Send[PredicateEditor](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPredicateEditor creates and returns a new initialized instance.
func NewPredicateEditor() PredicateEditor {
	return PredicateEditorClass.New()
}

// Init initializes the instance.
func (p_ PredicateEditor) Init() PredicateEditor {
	rv := objc.Send[PredicateEditor](p_.ID(), selInit)
	return rv
}
