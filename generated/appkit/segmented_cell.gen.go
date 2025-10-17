
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SegmentedCell] class.
var SegmentedCellClass _SegmentedCellClass

func init() {
	SegmentedCellClass = _SegmentedCellClass{objc.GetClass("NSSegmentedCell")}
}

type _SegmentedCellClass struct {
	objc.Class
}

// An interface definition for the [SegmentedCell] class.
type ISegmentedCell interface {
	ID() objc.ID
}

type SegmentedCell struct {
	id objc.ID
}

func SegmentedCellFrom(ptr unsafe.Pointer) SegmentedCell {
	return SegmentedCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SegmentedCell) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SegmentedCellClass) Alloc() SegmentedCell {
	rv := objc.Send[SegmentedCell](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SegmentedCellClass) New() SegmentedCell {
	rv := objc.Send[SegmentedCell](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSegmentedCell creates and returns a new initialized instance.
func NewSegmentedCell() SegmentedCell {
	return SegmentedCellClass.New()
}

// Init initializes the instance.
func (s_ SegmentedCell) Init() SegmentedCell {
	rv := objc.Send[SegmentedCell](s_.ID(), selInit)
	return rv
}
