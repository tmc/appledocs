// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [dataCodewordCount] class.
var (
	DataCodewordCountClass     _dataCodewordCountClass
	DataCodewordCountClassOnce sync.Once
)

func getdataCodewordCountClass() _dataCodewordCountClass {
	DataCodewordCountClassOnce.Do(func() {
		DataCodewordCountClass = _dataCodewordCountClass{objc.GetClass("dataCodewordCount")}
	})
	return DataCodewordCountClass
}

type _dataCodewordCountClass struct {
	class objc.Class
}

// An interface definition for the [dataCodewordCount] class.
type IdataCodewordCount interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/dataCodewordCount-c.ivar
type dataCodewordCount struct {
	objectivec.Object
}

// dataCodewordCountFrom constructs a [dataCodewordCount] from an unsafe.Pointer.
func dataCodewordCountFrom(ptr unsafe.Pointer) dataCodewordCount {
	return dataCodewordCount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _dataCodewordCountClass) Alloc() dataCodewordCount {
	rv := objc.Send[dataCodewordCount](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _dataCodewordCountClass) New() dataCodewordCount {
	rv := objc.Send[dataCodewordCount](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ dataCodewordCount) Init() dataCodewordCount {
	rv := objc.Send[dataCodewordCount](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ dataCodewordCount) Autorelease() dataCodewordCount {
	rv := objc.Send[dataCodewordCount](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdataCodewordCount creates a new dataCodewordCount instance.
func NewdataCodewordCount() dataCodewordCount {
	return getdataCodewordCountClass().New()
}




