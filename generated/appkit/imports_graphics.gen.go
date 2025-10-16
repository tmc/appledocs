
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [importsGraphics] class.
var importsGraphicsClass _importsGraphicsClass

func init() {
	importsGraphicsClass = _importsGraphicsClass{objc.GetClass("importsGraphics")}
}

type _importsGraphicsClass struct {
	objc.Class
}

// An interface definition for the [importsGraphics] class.
type IimportsGraphics interface {
	ID() objc.ID
}

type importsGraphics struct {
	id objc.ID
}

func importsGraphicsFrom(ptr unsafe.Pointer) importsGraphics {
	return importsGraphics{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ importsGraphics) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _importsGraphicsClass) Alloc() importsGraphics {
	rv := objc.Send[importsGraphics](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _importsGraphicsClass) New() importsGraphics {
	rv := objc.Send[importsGraphics](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewimportsGraphics creates and returns a new initialized instance.
func NewimportsGraphics() importsGraphics {
	return importsGraphicsClass.New()
}

// Init initializes the instance.
func (i_ importsGraphics) Init() importsGraphics {
	rv := objc.Send[importsGraphics](i_.ID(), selInit)
	return rv
}
