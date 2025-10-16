
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StringDrawingContext] class.
var StringDrawingContextClass _StringDrawingContextClass

func init() {
	StringDrawingContextClass = _StringDrawingContextClass{objc.GetClass("NSStringDrawingContext")}
}

type _StringDrawingContextClass struct {
	objc.Class
}

// An interface definition for the [StringDrawingContext] class.
type IStringDrawingContext interface {
	ID() objc.ID
}

type StringDrawingContext struct {
	id objc.ID
}

func StringDrawingContextFrom(ptr unsafe.Pointer) StringDrawingContext {
	return StringDrawingContext{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ StringDrawingContext) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StringDrawingContextClass) Alloc() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StringDrawingContextClass) New() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStringDrawingContext creates and returns a new initialized instance.
func NewStringDrawingContext() StringDrawingContext {
	return StringDrawingContextClass.New()
}

// Init initializes the instance.
func (s_ StringDrawingContext) Init() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](s_.ID(), selInit)
	return rv
}
