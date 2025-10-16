
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [numberOfColumns] class.
var numberOfColumnsClass _numberOfColumnsClass

func init() {
	numberOfColumnsClass = _numberOfColumnsClass{objc.GetClass("numberOfColumns")}
}

type _numberOfColumnsClass struct {
	objc.Class
}

// An interface definition for the [numberOfColumns] class.
type InumberOfColumns interface {
	ID() objc.ID
}

type numberOfColumns struct {
	id objc.ID
}

func numberOfColumnsFrom(ptr unsafe.Pointer) numberOfColumns {
	return numberOfColumns{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ numberOfColumns) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _numberOfColumnsClass) Alloc() numberOfColumns {
	rv := objc.Send[numberOfColumns](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _numberOfColumnsClass) New() numberOfColumns {
	rv := objc.Send[numberOfColumns](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnumberOfColumns creates and returns a new initialized instance.
func NewnumberOfColumns() numberOfColumns {
	return numberOfColumnsClass.New()
}

// Init initializes the instance.
func (n_ numberOfColumns) Init() numberOfColumns {
	rv := objc.Send[numberOfColumns](n_.ID(), selInit)
	return rv
}
