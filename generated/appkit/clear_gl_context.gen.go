
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [clearGLContext] class.
var clearGLContextClass _clearGLContextClass

func init() {
	clearGLContextClass = _clearGLContextClass{objc.GetClass("clearGLContext")}
}

type _clearGLContextClass struct {
	objc.Class
}

// An interface definition for the [clearGLContext] class.
type IclearGLContext interface {
	ID() objc.ID
}

type clearGLContext struct {
	id objc.ID
}

func clearGLContextFrom(ptr unsafe.Pointer) clearGLContext {
	return clearGLContext{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ clearGLContext) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _clearGLContextClass) Alloc() clearGLContext {
	rv := objc.Send[clearGLContext](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _clearGLContextClass) New() clearGLContext {
	rv := objc.Send[clearGLContext](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewclearGLContext creates and returns a new initialized instance.
func NewclearGLContext() clearGLContext {
	return clearGLContextClass.New()
}

// Init initializes the instance.
func (c_ clearGLContext) Init() clearGLContext {
	rv := objc.Send[clearGLContext](c_.ID(), selInit)
	return rv
}
