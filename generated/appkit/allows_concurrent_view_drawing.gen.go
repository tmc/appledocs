
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsConcurrentViewDrawing] class.
var allowsConcurrentViewDrawingClass _allowsConcurrentViewDrawingClass

func init() {
	allowsConcurrentViewDrawingClass = _allowsConcurrentViewDrawingClass{objc.GetClass("allowsConcurrentViewDrawing")}
}

type _allowsConcurrentViewDrawingClass struct {
	objc.Class
}

// An interface definition for the [allowsConcurrentViewDrawing] class.
type IallowsConcurrentViewDrawing interface {
	ID() objc.ID
}

type allowsConcurrentViewDrawing struct {
	id objc.ID
}

func allowsConcurrentViewDrawingFrom(ptr unsafe.Pointer) allowsConcurrentViewDrawing {
	return allowsConcurrentViewDrawing{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsConcurrentViewDrawing) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsConcurrentViewDrawingClass) Alloc() allowsConcurrentViewDrawing {
	rv := objc.Send[allowsConcurrentViewDrawing](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsConcurrentViewDrawingClass) New() allowsConcurrentViewDrawing {
	rv := objc.Send[allowsConcurrentViewDrawing](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsConcurrentViewDrawing creates and returns a new initialized instance.
func NewallowsConcurrentViewDrawing() allowsConcurrentViewDrawing {
	return allowsConcurrentViewDrawingClass.New()
}

// Init initializes the instance.
func (a_ allowsConcurrentViewDrawing) Init() allowsConcurrentViewDrawing {
	rv := objc.Send[allowsConcurrentViewDrawing](a_.ID(), selInit)
	return rv
}
