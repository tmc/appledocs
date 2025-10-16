
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [prepareOpenGL] class.
var prepareOpenGLClass _prepareOpenGLClass

func init() {
	prepareOpenGLClass = _prepareOpenGLClass{objc.GetClass("prepareOpenGL")}
}

type _prepareOpenGLClass struct {
	objc.Class
}

// An interface definition for the [prepareOpenGL] class.
type IprepareOpenGL interface {
	ID() objc.ID
}

type prepareOpenGL struct {
	id objc.ID
}

func prepareOpenGLFrom(ptr unsafe.Pointer) prepareOpenGL {
	return prepareOpenGL{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ prepareOpenGL) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _prepareOpenGLClass) Alloc() prepareOpenGL {
	rv := objc.Send[prepareOpenGL](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _prepareOpenGLClass) New() prepareOpenGL {
	rv := objc.Send[prepareOpenGL](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewprepareOpenGL creates and returns a new initialized instance.
func NewprepareOpenGL() prepareOpenGL {
	return prepareOpenGLClass.New()
}

// Init initializes the instance.
func (p_ prepareOpenGL) Init() prepareOpenGL {
	rv := objc.Send[prepareOpenGL](p_.ID(), selInit)
	return rv
}
