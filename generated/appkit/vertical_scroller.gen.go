
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [verticalScroller] class.
var verticalScrollerClass _verticalScrollerClass

func init() {
	verticalScrollerClass = _verticalScrollerClass{objc.GetClass("verticalScroller")}
}

type _verticalScrollerClass struct {
	objc.Class
}

// An interface definition for the [verticalScroller] class.
type IverticalScroller interface {
	ID() objc.ID
}

type verticalScroller struct {
	id objc.ID
}

func verticalScrollerFrom(ptr unsafe.Pointer) verticalScroller {
	return verticalScroller{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ verticalScroller) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _verticalScrollerClass) Alloc() verticalScroller {
	rv := objc.Send[verticalScroller](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _verticalScrollerClass) New() verticalScroller {
	rv := objc.Send[verticalScroller](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewverticalScroller creates and returns a new initialized instance.
func NewverticalScroller() verticalScroller {
	return verticalScrollerClass.New()
}

// Init initializes the instance.
func (v_ verticalScroller) Init() verticalScroller {
	rv := objc.Send[verticalScroller](v_.ID(), selInit)
	return rv
}
