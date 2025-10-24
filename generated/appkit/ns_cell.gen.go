// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Cell] class.
var (
	CellClass     _CellClass
	CellClassOnce sync.Once
)

func getCellClass() _CellClass {
	CellClassOnce.Do(func() {
		CellClass = _CellClass{objc.GetClass("NSCell")}
	})
	return CellClass
}

type _CellClass struct {
	class objc.Class
}

// An interface definition for the [Cell] class.
type ICell interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type Cell struct {
	objectivec.Object
}

// CellFrom constructs a [Cell] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func CellFrom(ptr unsafe.Pointer) Cell {
	return Cell{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CellClass) Alloc() Cell {
	rv := objc.Send[Cell](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CellClass) New() Cell {
	rv := objc.Send[Cell](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Cell) Init() Cell {
	rv := objc.Send[Cell](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Cell) Autorelease() Cell {
	rv := objc.Send[Cell](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCell creates a new Cell instance.
func NewCell() Cell {
	return getCellClass().New()
}




