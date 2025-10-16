
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GraphicsContext] class.
var GraphicsContextClass _GraphicsContextClass

func init() {
	GraphicsContextClass = _GraphicsContextClass{objc.GetClass("NSGraphicsContext")}
}

type _GraphicsContextClass struct {
	objc.Class
}

// An interface definition for the [GraphicsContext] class.
type IGraphicsContext interface {
	ID() objc.ID
}

type GraphicsContext struct {
	id objc.ID
}

func GraphicsContextFrom(ptr unsafe.Pointer) GraphicsContext {
	return GraphicsContext{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ GraphicsContext) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GraphicsContextClass) Alloc() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GraphicsContextClass) New() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGraphicsContext creates and returns a new initialized instance.
func NewGraphicsContext() GraphicsContext {
	return GraphicsContextClass.New()
}

// Init initializes the instance.
func (g_ GraphicsContext) Init() GraphicsContext {
	rv := objc.Send[GraphicsContext](g_.ID(), selInit)
	return rv
}
