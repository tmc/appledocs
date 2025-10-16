
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [recalculateKeyViewLoop] class.
var recalculateKeyViewLoopClass _recalculateKeyViewLoopClass

func init() {
	recalculateKeyViewLoopClass = _recalculateKeyViewLoopClass{objc.GetClass("recalculateKeyViewLoop")}
}

type _recalculateKeyViewLoopClass struct {
	objc.Class
}

// An interface definition for the [recalculateKeyViewLoop] class.
type IrecalculateKeyViewLoop interface {
	ID() objc.ID
}

type recalculateKeyViewLoop struct {
	id objc.ID
}

func recalculateKeyViewLoopFrom(ptr unsafe.Pointer) recalculateKeyViewLoop {
	return recalculateKeyViewLoop{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ recalculateKeyViewLoop) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _recalculateKeyViewLoopClass) Alloc() recalculateKeyViewLoop {
	rv := objc.Send[recalculateKeyViewLoop](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _recalculateKeyViewLoopClass) New() recalculateKeyViewLoop {
	rv := objc.Send[recalculateKeyViewLoop](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrecalculateKeyViewLoop creates and returns a new initialized instance.
func NewrecalculateKeyViewLoop() recalculateKeyViewLoop {
	return recalculateKeyViewLoopClass.New()
}

// Init initializes the instance.
func (r_ recalculateKeyViewLoop) Init() recalculateKeyViewLoop {
	rv := objc.Send[recalculateKeyViewLoop](r_.ID(), selInit)
	return rv
}
