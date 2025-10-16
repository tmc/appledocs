
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canBecomeKeyView] class.
var canBecomeKeyViewClass _canBecomeKeyViewClass

func init() {
	canBecomeKeyViewClass = _canBecomeKeyViewClass{objc.GetClass("canBecomeKeyView")}
}

type _canBecomeKeyViewClass struct {
	objc.Class
}

// An interface definition for the [canBecomeKeyView] class.
type IcanBecomeKeyView interface {
	ID() objc.ID
}

type canBecomeKeyView struct {
	id objc.ID
}

func canBecomeKeyViewFrom(ptr unsafe.Pointer) canBecomeKeyView {
	return canBecomeKeyView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canBecomeKeyView) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canBecomeKeyViewClass) Alloc() canBecomeKeyView {
	rv := objc.Send[canBecomeKeyView](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canBecomeKeyViewClass) New() canBecomeKeyView {
	rv := objc.Send[canBecomeKeyView](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanBecomeKeyView creates and returns a new initialized instance.
func NewcanBecomeKeyView() canBecomeKeyView {
	return canBecomeKeyViewClass.New()
}

// Init initializes the instance.
func (c_ canBecomeKeyView) Init() canBecomeKeyView {
	rv := objc.Send[canBecomeKeyView](c_.ID(), selInit)
	return rv
}
