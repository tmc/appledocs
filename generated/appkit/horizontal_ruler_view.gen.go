
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [horizontalRulerView] class.
var horizontalRulerViewClass _horizontalRulerViewClass

func init() {
	horizontalRulerViewClass = _horizontalRulerViewClass{objc.GetClass("horizontalRulerView")}
}

type _horizontalRulerViewClass struct {
	objc.Class
}

// An interface definition for the [horizontalRulerView] class.
type IhorizontalRulerView interface {
	ID() objc.ID
}

type horizontalRulerView struct {
	id objc.ID
}

func horizontalRulerViewFrom(ptr unsafe.Pointer) horizontalRulerView {
	return horizontalRulerView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ horizontalRulerView) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _horizontalRulerViewClass) Alloc() horizontalRulerView {
	rv := objc.Send[horizontalRulerView](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _horizontalRulerViewClass) New() horizontalRulerView {
	rv := objc.Send[horizontalRulerView](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhorizontalRulerView creates and returns a new initialized instance.
func NewhorizontalRulerView() horizontalRulerView {
	return horizontalRulerViewClass.New()
}

// Init initializes the instance.
func (h_ horizontalRulerView) Init() horizontalRulerView {
	rv := objc.Send[horizontalRulerView](h_.ID(), selInit)
	return rv
}
