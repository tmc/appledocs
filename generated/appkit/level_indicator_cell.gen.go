
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LevelIndicatorCell] class.
var LevelIndicatorCellClass _LevelIndicatorCellClass

func init() {
	LevelIndicatorCellClass = _LevelIndicatorCellClass{objc.GetClass("NSLevelIndicatorCell")}
}

type _LevelIndicatorCellClass struct {
	objc.Class
}

// An interface definition for the [LevelIndicatorCell] class.
type ILevelIndicatorCell interface {
	ID() objc.ID
}

type LevelIndicatorCell struct {
	id objc.ID
}

func LevelIndicatorCellFrom(ptr unsafe.Pointer) LevelIndicatorCell {
	return LevelIndicatorCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ LevelIndicatorCell) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _LevelIndicatorCellClass) Alloc() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _LevelIndicatorCellClass) New() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewLevelIndicatorCell creates and returns a new initialized instance.
func NewLevelIndicatorCell() LevelIndicatorCell {
	return LevelIndicatorCellClass.New()
}

// Init initializes the instance.
func (l_ LevelIndicatorCell) Init() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](l_.ID(), selInit)
	return rv
}
