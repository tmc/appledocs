
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backingScaleFactor] class.
var backingScaleFactorClass _backingScaleFactorClass

func init() {
	backingScaleFactorClass = _backingScaleFactorClass{objc.GetClass("backingScaleFactor")}
}

type _backingScaleFactorClass struct {
	objc.Class
}

// An interface definition for the [backingScaleFactor] class.
type IbackingScaleFactor interface {
	ID() objc.ID
}

type backingScaleFactor struct {
	id objc.ID
}

func backingScaleFactorFrom(ptr unsafe.Pointer) backingScaleFactor {
	return backingScaleFactor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backingScaleFactor) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backingScaleFactorClass) Alloc() backingScaleFactor {
	rv := objc.Send[backingScaleFactor](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backingScaleFactorClass) New() backingScaleFactor {
	rv := objc.Send[backingScaleFactor](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackingScaleFactor creates and returns a new initialized instance.
func NewbackingScaleFactor() backingScaleFactor {
	return backingScaleFactorClass.New()
}

// Init initializes the instance.
func (b_ backingScaleFactor) Init() backingScaleFactor {
	rv := objc.Send[backingScaleFactor](b_.ID(), selInit)
	return rv
}
