
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [depth] class.
var depthClass _depthClass

func init() {
	depthClass = _depthClass{objc.GetClass("depth")}
}

type _depthClass struct {
	objc.Class
}

// An interface definition for the [depth] class.
type Idepth interface {
	ID() objc.ID
}

type depth struct {
	id objc.ID
}

func depthFrom(ptr unsafe.Pointer) depth {
	return depth{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ depth) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _depthClass) Alloc() depth {
	rv := objc.Send[depth](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _depthClass) New() depth {
	rv := objc.Send[depth](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newdepth creates and returns a new initialized instance.
func Newdepth() depth {
	return depthClass.New()
}

// Init initializes the instance.
func (d_ depth) Init() depth {
	rv := objc.Send[depth](d_.ID(), selInit)
	return rv
}
