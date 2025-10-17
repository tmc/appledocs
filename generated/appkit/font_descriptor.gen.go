
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontDescriptor] class.
var FontDescriptorClass _FontDescriptorClass

func init() {
	FontDescriptorClass = _FontDescriptorClass{objc.GetClass("NSFontDescriptor")}
}

type _FontDescriptorClass struct {
	objc.Class
}

// An interface definition for the [FontDescriptor] class.
type IFontDescriptor interface {
	ID() objc.ID
}

type FontDescriptor struct {
	id objc.ID
}

func FontDescriptorFrom(ptr unsafe.Pointer) FontDescriptor {
	return FontDescriptor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ FontDescriptor) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FontDescriptorClass) Alloc() FontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FontDescriptorClass) New() FontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewFontDescriptor creates and returns a new initialized instance.
func NewFontDescriptor() FontDescriptor {
	return FontDescriptorClass.New()
}

// Init initializes the instance.
func (f_ FontDescriptor) Init() FontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID(), selInit)
	return rv
}
