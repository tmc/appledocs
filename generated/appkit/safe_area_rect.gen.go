
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [safeAreaRect] class.
var safeAreaRectClass _safeAreaRectClass

func init() {
	safeAreaRectClass = _safeAreaRectClass{objc.GetClass("safeAreaRect")}
}

type _safeAreaRectClass struct {
	objc.Class
}

// An interface definition for the [safeAreaRect] class.
type IsafeAreaRect interface {
	ID() objc.ID
}

type safeAreaRect struct {
	id objc.ID
}

func safeAreaRectFrom(ptr unsafe.Pointer) safeAreaRect {
	return safeAreaRect{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ safeAreaRect) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _safeAreaRectClass) Alloc() safeAreaRect {
	rv := objc.Send[safeAreaRect](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _safeAreaRectClass) New() safeAreaRect {
	rv := objc.Send[safeAreaRect](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsafeAreaRect creates and returns a new initialized instance.
func NewsafeAreaRect() safeAreaRect {
	return safeAreaRectClass.New()
}

// Init initializes the instance.
func (s_ safeAreaRect) Init() safeAreaRect {
	rv := objc.Send[safeAreaRect](s_.ID(), selInit)
	return rv
}
