
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [toolTip] class.
var toolTipClass _toolTipClass

func init() {
	toolTipClass = _toolTipClass{objc.GetClass("toolTip")}
}

type _toolTipClass struct {
	objc.Class
}

// An interface definition for the [toolTip] class.
type ItoolTip interface {
	ID() objc.ID
}

type toolTip struct {
	id objc.ID
}

func toolTipFrom(ptr unsafe.Pointer) toolTip {
	return toolTip{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ toolTip) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _toolTipClass) Alloc() toolTip {
	rv := objc.Send[toolTip](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _toolTipClass) New() toolTip {
	rv := objc.Send[toolTip](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtoolTip creates and returns a new initialized instance.
func NewtoolTip() toolTip {
	return toolTipClass.New()
}

// Init initializes the instance.
func (t_ toolTip) Init() toolTip {
	rv := objc.Send[toolTip](t_.ID(), selInit)
	return rv
}
