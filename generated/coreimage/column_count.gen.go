// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [columnCount] class.
var (
	columnCountClass     _columnCountClass
	columnCountClassOnce sync.Once
)

func getcolumnCountClass() _columnCountClass {
	columnCountClassOnce.Do(func() {
		columnCountClass = _columnCountClass{objc.GetClass("columnCount")}
	})
	return columnCountClass
}

type _columnCountClass struct {
	class objc.Class
}

// An interface definition for the [columnCount] class.
type IcolumnCount interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/columnCount-c.ivar
type columnCount struct {
	objectivec.Object
}

// columnCountFrom constructs a [columnCount] from an unsafe.Pointer.
func columnCountFrom(ptr unsafe.Pointer) columnCount {
	return columnCount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _columnCountClass) Alloc() columnCount {
	rv := objc.Send[columnCount](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _columnCountClass) New() columnCount {
	rv := objc.Send[columnCount](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ columnCount) Init() columnCount {
	rv := objc.Send[columnCount](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ columnCount) Autorelease() columnCount {
	rv := objc.Send[columnCount](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcolumnCount creates a new columnCount instance.
func NewcolumnCount() columnCount {
	return getcolumnCountClass().New()
}




