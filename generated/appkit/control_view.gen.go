
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [controlView] class.
var controlViewClass _controlViewClass

func init() {
	controlViewClass = _controlViewClass{objc.GetClass("controlView")}
}

type _controlViewClass struct {
	objc.Class
}

// An interface definition for the [controlView] class.
type IcontrolView interface {
	ID() objc.ID
}

type controlView struct {
	id objc.ID
}

func controlViewFrom(ptr unsafe.Pointer) controlView {
	return controlView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ controlView) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _controlViewClass) Alloc() controlView {
	rv := objc.Send[controlView](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _controlViewClass) New() controlView {
	rv := objc.Send[controlView](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontrolView creates and returns a new initialized instance.
func NewcontrolView() controlView {
	return controlViewClass.New()
}

// Init initializes the instance.
func (c_ controlView) Init() controlView {
	rv := objc.Send[controlView](c_.ID(), selInit)
	return rv
}
