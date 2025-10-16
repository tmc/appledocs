
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ignoresMultiClick] class.
var ignoresMultiClickClass _ignoresMultiClickClass

func init() {
	ignoresMultiClickClass = _ignoresMultiClickClass{objc.GetClass("ignoresMultiClick")}
}

type _ignoresMultiClickClass struct {
	objc.Class
}

// An interface definition for the [ignoresMultiClick] class.
type IignoresMultiClick interface {
	ID() objc.ID
}

type ignoresMultiClick struct {
	id objc.ID
}

func ignoresMultiClickFrom(ptr unsafe.Pointer) ignoresMultiClick {
	return ignoresMultiClick{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ ignoresMultiClick) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _ignoresMultiClickClass) Alloc() ignoresMultiClick {
	rv := objc.Send[ignoresMultiClick](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _ignoresMultiClickClass) New() ignoresMultiClick {
	rv := objc.Send[ignoresMultiClick](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewignoresMultiClick creates and returns a new initialized instance.
func NewignoresMultiClick() ignoresMultiClick {
	return ignoresMultiClickClass.New()
}

// Init initializes the instance.
func (i_ ignoresMultiClick) Init() ignoresMultiClick {
	rv := objc.Send[ignoresMultiClick](i_.ID(), selInit)
	return rv
}
