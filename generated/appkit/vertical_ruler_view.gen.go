
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [verticalRulerView] class.
var verticalRulerViewClass _verticalRulerViewClass

func init() {
	verticalRulerViewClass = _verticalRulerViewClass{objc.GetClass("verticalRulerView")}
}

type _verticalRulerViewClass struct {
	objc.Class
}

// An interface definition for the [verticalRulerView] class.
type IverticalRulerView interface {
	ID() objc.ID
}

type verticalRulerView struct {
	id objc.ID
}

func verticalRulerViewFrom(ptr unsafe.Pointer) verticalRulerView {
	return verticalRulerView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ verticalRulerView) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _verticalRulerViewClass) Alloc() verticalRulerView {
	rv := objc.Send[verticalRulerView](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _verticalRulerViewClass) New() verticalRulerView {
	rv := objc.Send[verticalRulerView](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewverticalRulerView creates and returns a new initialized instance.
func NewverticalRulerView() verticalRulerView {
	return verticalRulerViewClass.New()
}

// Init initializes the instance.
func (v_ verticalRulerView) Init() verticalRulerView {
	rv := objc.Send[verticalRulerView](v_.ID(), selInit)
	return rv
}
