
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [customizationPaletteIsRunning] class.
var customizationPaletteIsRunningClass _customizationPaletteIsRunningClass

func init() {
	customizationPaletteIsRunningClass = _customizationPaletteIsRunningClass{objc.GetClass("customizationPaletteIsRunning")}
}

type _customizationPaletteIsRunningClass struct {
	objc.Class
}

// An interface definition for the [customizationPaletteIsRunning] class.
type IcustomizationPaletteIsRunning interface {
	ID() objc.ID
}

type customizationPaletteIsRunning struct {
	id objc.ID
}

func customizationPaletteIsRunningFrom(ptr unsafe.Pointer) customizationPaletteIsRunning {
	return customizationPaletteIsRunning{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ customizationPaletteIsRunning) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _customizationPaletteIsRunningClass) Alloc() customizationPaletteIsRunning {
	rv := objc.Send[customizationPaletteIsRunning](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _customizationPaletteIsRunningClass) New() customizationPaletteIsRunning {
	rv := objc.Send[customizationPaletteIsRunning](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcustomizationPaletteIsRunning creates and returns a new initialized instance.
func NewcustomizationPaletteIsRunning() customizationPaletteIsRunning {
	return customizationPaletteIsRunningClass.New()
}

// Init initializes the instance.
func (c_ customizationPaletteIsRunning) Init() customizationPaletteIsRunning {
	rv := objc.Send[customizationPaletteIsRunning](c_.ID(), selInit)
	return rv
}
