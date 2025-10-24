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
	ColumnCountClass     _columnCountClass
	ColumnCountClassOnce sync.Once
)

func getcolumnCountClass() _columnCountClass {
	ColumnCountClassOnce.Do(func() {
		ColumnCountClass = _columnCountClass{objc.GetClass("columnCount")}
	})
	return ColumnCountClass
}

type _columnCountClass struct {
	class objc.Class
}





// An interface definition for the [columnCount] class.
type IcolumnCount interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _columnCountClass) Alloc() columnCount {
	rv := objc.Send[columnCount](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/columnCount-c.ivar
type columnCount struct {
	objectivec.Object
}

// columnCountFrom constructs a [columnCount] from an unsafe.Pointer.
func columnCountFrom(ptr unsafe.Pointer) columnCount {
	return columnCount{objectivec.Object{objc.ID(ptr)}}
}































