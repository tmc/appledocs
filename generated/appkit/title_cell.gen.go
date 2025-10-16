
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [titleCell] class.
var titleCellClass _titleCellClass

func init() {
	titleCellClass = _titleCellClass{objc.GetClass("titleCell")}
}

type _titleCellClass struct {
	objc.Class
}

// An interface definition for the [titleCell] class.
type ItitleCell interface {
	ID() objc.ID
}

type titleCell struct {
	id objc.ID
}

func titleCellFrom(ptr unsafe.Pointer) titleCell {
	return titleCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ titleCell) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _titleCellClass) Alloc() titleCell {
	rv := objc.Send[titleCell](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _titleCellClass) New() titleCell {
	rv := objc.Send[titleCell](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtitleCell creates and returns a new initialized instance.
func NewtitleCell() titleCell {
	return titleCellClass.New()
}

// Init initializes the instance.
func (t_ titleCell) Init() titleCell {
	rv := objc.Send[titleCell](t_.ID(), selInit)
	return rv
}
