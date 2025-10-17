
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextInputContext] class.
var TextInputContextClass _TextInputContextClass

func init() {
	TextInputContextClass = _TextInputContextClass{objc.GetClass("NSTextInputContext")}
}

type _TextInputContextClass struct {
	objc.Class
}

// An interface definition for the [TextInputContext] class.
type ITextInputContext interface {
	ID() objc.ID
}

type TextInputContext struct {
	id objc.ID
}

func TextInputContextFrom(ptr unsafe.Pointer) TextInputContext {
	return TextInputContext{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextInputContext) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextInputContextClass) Alloc() TextInputContext {
	rv := objc.Send[TextInputContext](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextInputContextClass) New() TextInputContext {
	rv := objc.Send[TextInputContext](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextInputContext creates and returns a new initialized instance.
func NewTextInputContext() TextInputContext {
	return TextInputContextClass.New()
}

// Init initializes the instance.
func (t_ TextInputContext) Init() TextInputContext {
	rv := objc.Send[TextInputContext](t_.ID(), selInit)
	return rv
}
