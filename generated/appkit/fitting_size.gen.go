
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [fittingSize] class.
var fittingSizeClass _fittingSizeClass

func init() {
	fittingSizeClass = _fittingSizeClass{objc.GetClass("fittingSize")}
}

type _fittingSizeClass struct {
	objc.Class
}

// An interface definition for the [fittingSize] class.
type IfittingSize interface {
	ID() objc.ID
}

type fittingSize struct {
	id objc.ID
}

func fittingSizeFrom(ptr unsafe.Pointer) fittingSize {
	return fittingSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ fittingSize) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _fittingSizeClass) Alloc() fittingSize {
	rv := objc.Send[fittingSize](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _fittingSizeClass) New() fittingSize {
	rv := objc.Send[fittingSize](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfittingSize creates and returns a new initialized instance.
func NewfittingSize() fittingSize {
	return fittingSizeClass.New()
}

// Init initializes the instance.
func (f_ fittingSize) Init() fittingSize {
	rv := objc.Send[fittingSize](f_.ID(), selInit)
	return rv
}
