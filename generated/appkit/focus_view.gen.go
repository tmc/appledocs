
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [focusView] class.
var focusViewClass _focusViewClass

func init() {
	focusViewClass = _focusViewClass{objc.GetClass("focusView")}
}

type _focusViewClass struct {
	objc.Class
}

// An interface definition for the [focusView] class.
type IfocusView interface {
	ID() objc.ID
}

type focusView struct {
	id objc.ID
}

func focusViewFrom(ptr unsafe.Pointer) focusView {
	return focusView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ focusView) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _focusViewClass) Alloc() focusView {
	rv := objc.Send[focusView](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _focusViewClass) New() focusView {
	rv := objc.Send[focusView](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfocusView creates and returns a new initialized instance.
func NewfocusView() focusView {
	return focusViewClass.New()
}

// Init initializes the instance.
func (f_ focusView) Init() focusView {
	rv := objc.Send[focusView](f_.ID(), selInit)
	return rv
}
