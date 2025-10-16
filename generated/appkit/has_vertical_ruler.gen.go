
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasVerticalRuler] class.
var hasVerticalRulerClass _hasVerticalRulerClass

func init() {
	hasVerticalRulerClass = _hasVerticalRulerClass{objc.GetClass("hasVerticalRuler")}
}

type _hasVerticalRulerClass struct {
	objc.Class
}

// An interface definition for the [hasVerticalRuler] class.
type IhasVerticalRuler interface {
	ID() objc.ID
}

type hasVerticalRuler struct {
	id objc.ID
}

func hasVerticalRulerFrom(ptr unsafe.Pointer) hasVerticalRuler {
	return hasVerticalRuler{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasVerticalRuler) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasVerticalRulerClass) Alloc() hasVerticalRuler {
	rv := objc.Send[hasVerticalRuler](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasVerticalRulerClass) New() hasVerticalRuler {
	rv := objc.Send[hasVerticalRuler](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasVerticalRuler creates and returns a new initialized instance.
func NewhasVerticalRuler() hasVerticalRuler {
	return hasVerticalRulerClass.New()
}

// Init initializes the instance.
func (h_ hasVerticalRuler) Init() hasVerticalRuler {
	rv := objc.Send[hasVerticalRuler](h_.ID(), selInit)
	return rv
}
