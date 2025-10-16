
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [cellSize] class.
var cellSizeClass _cellSizeClass

func init() {
	cellSizeClass = _cellSizeClass{objc.GetClass("cellSize")}
}

type _cellSizeClass struct {
	objc.Class
}

// An interface definition for the [cellSize] class.
type IcellSize interface {
	ID() objc.ID
}

type cellSize struct {
	id objc.ID
}

func cellSizeFrom(ptr unsafe.Pointer) cellSize {
	return cellSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ cellSize) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _cellSizeClass) Alloc() cellSize {
	rv := objc.Send[cellSize](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _cellSizeClass) New() cellSize {
	rv := objc.Send[cellSize](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcellSize creates and returns a new initialized instance.
func NewcellSize() cellSize {
	return cellSizeClass.New()
}

// Init initializes the instance.
func (c_ cellSize) Init() cellSize {
	rv := objc.Send[cellSize](c_.ID(), selInit)
	return rv
}
