
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasHorizontalRuler] class.
var hasHorizontalRulerClass _hasHorizontalRulerClass

func init() {
	hasHorizontalRulerClass = _hasHorizontalRulerClass{objc.GetClass("hasHorizontalRuler")}
}

type _hasHorizontalRulerClass struct {
	objc.Class
}

// An interface definition for the [hasHorizontalRuler] class.
type IhasHorizontalRuler interface {
	ID() objc.ID
}

type hasHorizontalRuler struct {
	id objc.ID
}

func hasHorizontalRulerFrom(ptr unsafe.Pointer) hasHorizontalRuler {
	return hasHorizontalRuler{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasHorizontalRuler) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasHorizontalRulerClass) Alloc() hasHorizontalRuler {
	rv := objc.Send[hasHorizontalRuler](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasHorizontalRulerClass) New() hasHorizontalRuler {
	rv := objc.Send[hasHorizontalRuler](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasHorizontalRuler creates and returns a new initialized instance.
func NewhasHorizontalRuler() hasHorizontalRuler {
	return hasHorizontalRulerClass.New()
}

// Init initializes the instance.
func (h_ hasHorizontalRuler) Init() hasHorizontalRuler {
	rv := objc.Send[hasHorizontalRuler](h_.ID(), selInit)
	return rv
}
