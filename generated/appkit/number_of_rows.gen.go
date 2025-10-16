
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [numberOfRows] class.
var numberOfRowsClass _numberOfRowsClass

func init() {
	numberOfRowsClass = _numberOfRowsClass{objc.GetClass("numberOfRows")}
}

type _numberOfRowsClass struct {
	objc.Class
}

// An interface definition for the [numberOfRows] class.
type InumberOfRows interface {
	ID() objc.ID
}

type numberOfRows struct {
	id objc.ID
}

func numberOfRowsFrom(ptr unsafe.Pointer) numberOfRows {
	return numberOfRows{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ numberOfRows) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _numberOfRowsClass) Alloc() numberOfRows {
	rv := objc.Send[numberOfRows](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _numberOfRowsClass) New() numberOfRows {
	rv := objc.Send[numberOfRows](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnumberOfRows creates and returns a new initialized instance.
func NewnumberOfRows() numberOfRows {
	return numberOfRowsClass.New()
}

// Init initializes the instance.
func (n_ numberOfRows) Init() numberOfRows {
	rv := objc.Send[numberOfRows](n_.ID(), selInit)
	return rv
}
