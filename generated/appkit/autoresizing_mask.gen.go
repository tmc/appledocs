
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [autoresizingMask] class.
var autoresizingMaskClass _autoresizingMaskClass

func init() {
	autoresizingMaskClass = _autoresizingMaskClass{objc.GetClass("autoresizingMask")}
}

type _autoresizingMaskClass struct {
	objc.Class
}

// An interface definition for the [autoresizingMask] class.
type IautoresizingMask interface {
	ID() objc.ID
}

type autoresizingMask struct {
	id objc.ID
}

func autoresizingMaskFrom(ptr unsafe.Pointer) autoresizingMask {
	return autoresizingMask{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ autoresizingMask) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _autoresizingMaskClass) Alloc() autoresizingMask {
	rv := objc.Send[autoresizingMask](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _autoresizingMaskClass) New() autoresizingMask {
	rv := objc.Send[autoresizingMask](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewautoresizingMask creates and returns a new initialized instance.
func NewautoresizingMask() autoresizingMask {
	return autoresizingMaskClass.New()
}

// Init initializes the instance.
func (a_ autoresizingMask) Init() autoresizingMask {
	rv := objc.Send[autoresizingMask](a_.ID(), selInit)
	return rv
}
