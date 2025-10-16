
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [stringWithSavedFrame] class.
var stringWithSavedFrameClass _stringWithSavedFrameClass

func init() {
	stringWithSavedFrameClass = _stringWithSavedFrameClass{objc.GetClass("stringWithSavedFrame")}
}

type _stringWithSavedFrameClass struct {
	objc.Class
}

// An interface definition for the [stringWithSavedFrame] class.
type IstringWithSavedFrame interface {
	ID() objc.ID
}

type stringWithSavedFrame struct {
	id objc.ID
}

func stringWithSavedFrameFrom(ptr unsafe.Pointer) stringWithSavedFrame {
	return stringWithSavedFrame{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ stringWithSavedFrame) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _stringWithSavedFrameClass) Alloc() stringWithSavedFrame {
	rv := objc.Send[stringWithSavedFrame](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _stringWithSavedFrameClass) New() stringWithSavedFrame {
	rv := objc.Send[stringWithSavedFrame](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewstringWithSavedFrame creates and returns a new initialized instance.
func NewstringWithSavedFrame() stringWithSavedFrame {
	return stringWithSavedFrameClass.New()
}

// Init initializes the instance.
func (s_ stringWithSavedFrame) Init() stringWithSavedFrame {
	rv := objc.Send[stringWithSavedFrame](s_.ID(), selInit)
	return rv
}
