
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PathCell] class.
var PathCellClass _PathCellClass

func init() {
	PathCellClass = _PathCellClass{objc.GetClass("NSPathCell")}
}

type _PathCellClass struct {
	objc.Class
}

// An interface definition for the [PathCell] class.
type IPathCell interface {
	ID() objc.ID
}

type PathCell struct {
	id objc.ID
}

func PathCellFrom(ptr unsafe.Pointer) PathCell {
	return PathCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PathCell) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PathCellClass) Alloc() PathCell {
	rv := objc.Send[PathCell](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PathCellClass) New() PathCell {
	rv := objc.Send[PathCell](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPathCell creates and returns a new initialized instance.
func NewPathCell() PathCell {
	return PathCellClass.New()
}

// Init initializes the instance.
func (p_ PathCell) Init() PathCell {
	rv := objc.Send[PathCell](p_.ID(), selInit)
	return rv
}
