
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PathComponentCell] class.
var PathComponentCellClass _PathComponentCellClass

func init() {
	PathComponentCellClass = _PathComponentCellClass{objc.GetClass("NSPathComponentCell")}
}

type _PathComponentCellClass struct {
	objc.Class
}

// An interface definition for the [PathComponentCell] class.
type IPathComponentCell interface {
	ID() objc.ID
}

type PathComponentCell struct {
	id objc.ID
}

func PathComponentCellFrom(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PathComponentCell) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PathComponentCellClass) Alloc() PathComponentCell {
	rv := objc.Send[PathComponentCell](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PathComponentCellClass) New() PathComponentCell {
	rv := objc.Send[PathComponentCell](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPathComponentCell creates and returns a new initialized instance.
func NewPathComponentCell() PathComponentCell {
	return PathComponentCellClass.New()
}

// Init initializes the instance.
func (p_ PathComponentCell) Init() PathComponentCell {
	rv := objc.Send[PathComponentCell](p_.ID(), selInit)
	return rv
}
