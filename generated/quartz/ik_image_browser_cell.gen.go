// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IKImageBrowserCell] class.
var (
	IKImageBrowserCellClass     _IKImageBrowserCellClass
	IKImageBrowserCellClassOnce sync.Once
)

func getIKImageBrowserCellClass() _IKImageBrowserCellClass {
	IKImageBrowserCellClassOnce.Do(func() {
		IKImageBrowserCellClass = _IKImageBrowserCellClass{objc.GetClass("IKImageBrowserCell")}
	})
	return IKImageBrowserCellClass
}

type _IKImageBrowserCellClass struct {
	class objc.Class
}

// An interface definition for the [IKImageBrowserCell] class.
type IIKImageBrowserCell interface {
	objectivec.IObject
	// properties:
	// methods:
	CellState() unsafe.Pointer
}

// A class used to display a cell.
//
// class that is used to display a cell conforming to the in an .


// A class used to display a cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell
type IKImageBrowserCell struct {
	objectivec.Object
}

// IKImageBrowserCellFrom constructs a [IKImageBrowserCell] from an unsafe.Pointer.
//
// A class used to display a cell.
func IKImageBrowserCellFrom(ptr unsafe.Pointer) IKImageBrowserCell {
	return IKImageBrowserCell{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IKImageBrowserCellClass) Alloc() IKImageBrowserCell {
	rv := objc.Send[IKImageBrowserCell](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKImageBrowserCellClass) New() IKImageBrowserCell {
	rv := objc.Send[IKImageBrowserCell](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKImageBrowserCell) Init() IKImageBrowserCell {
	rv := objc.Send[IKImageBrowserCell](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKImageBrowserCell) Autorelease() IKImageBrowserCell {
	rv := objc.Send[IKImageBrowserCell](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKImageBrowserCell creates a new IKImageBrowserCell instance.
func NewIKImageBrowserCell() IKImageBrowserCell {
	return getIKImageBrowserCellClass().New()
}



// Returns the current cell state of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/cellState()
func (i_ IKImageBrowserCell) CellState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("cellState"))
	return rv
}



