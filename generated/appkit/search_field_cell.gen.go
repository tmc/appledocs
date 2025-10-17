
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SearchFieldCell] class.
var SearchFieldCellClass _SearchFieldCellClass

func init() {
	SearchFieldCellClass = _SearchFieldCellClass{objc.GetClass("NSSearchFieldCell")}
}

type _SearchFieldCellClass struct {
	objc.Class
}

// An interface definition for the [SearchFieldCell] class.
type ISearchFieldCell interface {
	ID() objc.ID
}

type SearchFieldCell struct {
	id objc.ID
}

func SearchFieldCellFrom(ptr unsafe.Pointer) SearchFieldCell {
	return SearchFieldCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SearchFieldCell) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SearchFieldCellClass) Alloc() SearchFieldCell {
	rv := objc.Send[SearchFieldCell](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SearchFieldCellClass) New() SearchFieldCell {
	rv := objc.Send[SearchFieldCell](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSearchFieldCell creates and returns a new initialized instance.
func NewSearchFieldCell() SearchFieldCell {
	return SearchFieldCellClass.New()
}

// Init initializes the instance.
func (s_ SearchFieldCell) Init() SearchFieldCell {
	rv := objc.Send[SearchFieldCell](s_.ID(), selInit)
	return rv
}
