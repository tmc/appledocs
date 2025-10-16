
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canDraw] class.
var canDrawClass _canDrawClass

func init() {
	canDrawClass = _canDrawClass{objc.GetClass("canDraw")}
}

type _canDrawClass struct {
	objc.Class
}

// An interface definition for the [canDraw] class.
type IcanDraw interface {
	ID() objc.ID
}

type canDraw struct {
	id objc.ID
}

func canDrawFrom(ptr unsafe.Pointer) canDraw {
	return canDraw{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canDraw) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canDrawClass) Alloc() canDraw {
	rv := objc.Send[canDraw](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canDrawClass) New() canDraw {
	rv := objc.Send[canDraw](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanDraw creates and returns a new initialized instance.
func NewcanDraw() canDraw {
	return canDrawClass.New()
}

// Init initializes the instance.
func (c_ canDraw) Init() canDraw {
	rv := objc.Send[canDraw](c_.ID(), selInit)
	return rv
}
