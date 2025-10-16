
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [frame] class.
var frameClass _frameClass

func init() {
	frameClass = _frameClass{objc.GetClass("frame")}
}

type _frameClass struct {
	objc.Class
}

// An interface definition for the [frame] class.
type Iframe interface {
	ID() objc.ID
}

type frame struct {
	id objc.ID
}

func frameFrom(ptr unsafe.Pointer) frame {
	return frame{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ frame) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _frameClass) Alloc() frame {
	rv := objc.Send[frame](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _frameClass) New() frame {
	rv := objc.Send[frame](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newframe creates and returns a new initialized instance.
func Newframe() frame {
	return frameClass.New()
}

// Init initializes the instance.
func (f_ frame) Init() frame {
	rv := objc.Send[frame](f_.ID(), selInit)
	return rv
}
