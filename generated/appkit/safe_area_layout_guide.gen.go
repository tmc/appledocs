
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [safeAreaLayoutGuide] class.
var safeAreaLayoutGuideClass _safeAreaLayoutGuideClass

func init() {
	safeAreaLayoutGuideClass = _safeAreaLayoutGuideClass{objc.GetClass("safeAreaLayoutGuide")}
}

type _safeAreaLayoutGuideClass struct {
	objc.Class
}

// An interface definition for the [safeAreaLayoutGuide] class.
type IsafeAreaLayoutGuide interface {
	ID() objc.ID
}

type safeAreaLayoutGuide struct {
	id objc.ID
}

func safeAreaLayoutGuideFrom(ptr unsafe.Pointer) safeAreaLayoutGuide {
	return safeAreaLayoutGuide{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ safeAreaLayoutGuide) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _safeAreaLayoutGuideClass) Alloc() safeAreaLayoutGuide {
	rv := objc.Send[safeAreaLayoutGuide](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _safeAreaLayoutGuideClass) New() safeAreaLayoutGuide {
	rv := objc.Send[safeAreaLayoutGuide](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsafeAreaLayoutGuide creates and returns a new initialized instance.
func NewsafeAreaLayoutGuide() safeAreaLayoutGuide {
	return safeAreaLayoutGuideClass.New()
}

// Init initializes the instance.
func (s_ safeAreaLayoutGuide) Init() safeAreaLayoutGuide {
	rv := objc.Send[safeAreaLayoutGuide](s_.ID(), selInit)
	return rv
}
