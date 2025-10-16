
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontManager] class.
var FontManagerClass _FontManagerClass

func init() {
	FontManagerClass = _FontManagerClass{objc.GetClass("NSFontManager")}
}

type _FontManagerClass struct {
	objc.Class
}

// An interface definition for the [FontManager] class.
type IFontManager interface {
	ID() objc.ID
}

type FontManager struct {
	id objc.ID
}

func FontManagerFrom(ptr unsafe.Pointer) FontManager {
	return FontManager{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ FontManager) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FontManagerClass) Alloc() FontManager {
	rv := objc.Send[FontManager](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FontManagerClass) New() FontManager {
	rv := objc.Send[FontManager](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewFontManager creates and returns a new initialized instance.
func NewFontManager() FontManager {
	return FontManagerClass.New()
}

// Init initializes the instance.
func (f_ FontManager) Init() FontManager {
	rv := objc.Send[FontManager](f_.ID(), selInit)
	return rv
}
