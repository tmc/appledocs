
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [safeAreaInsets] class.
var safeAreaInsetsClass _safeAreaInsetsClass

func init() {
	safeAreaInsetsClass = _safeAreaInsetsClass{objc.GetClass("safeAreaInsets")}
}

type _safeAreaInsetsClass struct {
	objc.Class
}

// An interface definition for the [safeAreaInsets] class.
type IsafeAreaInsets interface {
	ID() objc.ID
}

type safeAreaInsets struct {
	id objc.ID
}

func safeAreaInsetsFrom(ptr unsafe.Pointer) safeAreaInsets {
	return safeAreaInsets{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ safeAreaInsets) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _safeAreaInsetsClass) Alloc() safeAreaInsets {
	rv := objc.Send[safeAreaInsets](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _safeAreaInsetsClass) New() safeAreaInsets {
	rv := objc.Send[safeAreaInsets](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsafeAreaInsets creates and returns a new initialized instance.
func NewsafeAreaInsets() safeAreaInsets {
	return safeAreaInsetsClass.New()
}

// Init initializes the instance.
func (s_ safeAreaInsets) Init() safeAreaInsets {
	rv := objc.Send[safeAreaInsets](s_.ID(), selInit)
	return rv
}
