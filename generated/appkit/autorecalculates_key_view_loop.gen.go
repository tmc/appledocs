
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [autorecalculatesKeyViewLoop] class.
var autorecalculatesKeyViewLoopClass _autorecalculatesKeyViewLoopClass

func init() {
	autorecalculatesKeyViewLoopClass = _autorecalculatesKeyViewLoopClass{objc.GetClass("autorecalculatesKeyViewLoop")}
}

type _autorecalculatesKeyViewLoopClass struct {
	objc.Class
}

// An interface definition for the [autorecalculatesKeyViewLoop] class.
type IautorecalculatesKeyViewLoop interface {
	ID() objc.ID
}

type autorecalculatesKeyViewLoop struct {
	id objc.ID
}

func autorecalculatesKeyViewLoopFrom(ptr unsafe.Pointer) autorecalculatesKeyViewLoop {
	return autorecalculatesKeyViewLoop{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ autorecalculatesKeyViewLoop) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _autorecalculatesKeyViewLoopClass) Alloc() autorecalculatesKeyViewLoop {
	rv := objc.Send[autorecalculatesKeyViewLoop](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _autorecalculatesKeyViewLoopClass) New() autorecalculatesKeyViewLoop {
	rv := objc.Send[autorecalculatesKeyViewLoop](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewautorecalculatesKeyViewLoop creates and returns a new initialized instance.
func NewautorecalculatesKeyViewLoop() autorecalculatesKeyViewLoop {
	return autorecalculatesKeyViewLoopClass.New()
}

// Init initializes the instance.
func (a_ autorecalculatesKeyViewLoop) Init() autorecalculatesKeyViewLoop {
	rv := objc.Send[autorecalculatesKeyViewLoop](a_.ID(), selInit)
	return rv
}
