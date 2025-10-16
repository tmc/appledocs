
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tile] class.
var tileClass _tileClass

func init() {
	tileClass = _tileClass{objc.GetClass("tile")}
}

type _tileClass struct {
	objc.Class
}

// An interface definition for the [tile] class.
type Itile interface {
	ID() objc.ID
}

type tile struct {
	id objc.ID
}

func tileFrom(ptr unsafe.Pointer) tile {
	return tile{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tile) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tileClass) Alloc() tile {
	rv := objc.Send[tile](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tileClass) New() tile {
	rv := objc.Send[tile](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newtile creates and returns a new initialized instance.
func Newtile() tile {
	return tileClass.New()
}

// Init initializes the instance.
func (t_ tile) Init() tile {
	rv := objc.Send[tile](t_.ID(), selInit)
	return rv
}
