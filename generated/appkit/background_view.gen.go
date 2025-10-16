
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backgroundView] class.
var backgroundViewClass _backgroundViewClass

func init() {
	backgroundViewClass = _backgroundViewClass{objc.GetClass("backgroundView")}
}

type _backgroundViewClass struct {
	objc.Class
}

// An interface definition for the [backgroundView] class.
type IbackgroundView interface {
	ID() objc.ID
}

type backgroundView struct {
	id objc.ID
}

func backgroundViewFrom(ptr unsafe.Pointer) backgroundView {
	return backgroundView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backgroundView) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backgroundViewClass) Alloc() backgroundView {
	rv := objc.Send[backgroundView](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backgroundViewClass) New() backgroundView {
	rv := objc.Send[backgroundView](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackgroundView creates and returns a new initialized instance.
func NewbackgroundView() backgroundView {
	return backgroundViewClass.New()
}

// Init initializes the instance.
func (b_ backgroundView) Init() backgroundView {
	rv := objc.Send[backgroundView](b_.ID(), selInit)
	return rv
}
