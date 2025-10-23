// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [rowCount] class.
var (
	RowCountClass     _rowCountClass
	RowCountClassOnce sync.Once
)

func getrowCountClass() _rowCountClass {
	RowCountClassOnce.Do(func() {
		RowCountClass = _rowCountClass{objc.GetClass("rowCount")}
	})
	return RowCountClass
}

type _rowCountClass struct {
	class objc.Class
}

// An interface definition for the [rowCount] class.
type IrowCount interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/rowCount-c.ivar
type rowCount struct {
	objectivec.Object
}

// rowCountFrom constructs a [rowCount] from an unsafe.Pointer.
func rowCountFrom(ptr unsafe.Pointer) rowCount {
	return rowCount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _rowCountClass) Alloc() rowCount {
	rv := objc.Send[rowCount](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _rowCountClass) New() rowCount {
	rv := objc.Send[rowCount](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ rowCount) Init() rowCount {
	rv := objc.Send[rowCount](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ rowCount) Autorelease() rowCount {
	rv := objc.Send[rowCount](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrowCount creates a new rowCount instance.
func NewrowCount() rowCount {
	return getrowCountClass().New()
}




