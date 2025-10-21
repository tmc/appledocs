// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterSimpleStruct] class.
var (
	MTRTestClusterClusterSimpleStructClass     _MTRTestClusterClusterSimpleStructClass
	MTRTestClusterClusterSimpleStructClassOnce sync.Once
)

func getMTRTestClusterClusterSimpleStructClass() _MTRTestClusterClusterSimpleStructClass {
	MTRTestClusterClusterSimpleStructClassOnce.Do(func() {
		MTRTestClusterClusterSimpleStructClass = _MTRTestClusterClusterSimpleStructClass{objc.GetClass("MTRTestClusterClusterSimpleStruct")}
	})
	return MTRTestClusterClusterSimpleStructClass
}

type _MTRTestClusterClusterSimpleStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterSimpleStruct] class.
type IMTRTestClusterClusterSimpleStruct interface {
	IMTRUnitTestingClusterSimpleStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterSimpleStruct
type MTRTestClusterClusterSimpleStruct struct {
	MTRUnitTestingClusterSimpleStruct
}

// MTRTestClusterClusterSimpleStructFrom constructs a [MTRTestClusterClusterSimpleStruct] from an unsafe.Pointer.
func MTRTestClusterClusterSimpleStructFrom(ptr unsafe.Pointer) MTRTestClusterClusterSimpleStruct {
	return MTRTestClusterClusterSimpleStruct{
		MTRUnitTestingClusterSimpleStruct: MTRUnitTestingClusterSimpleStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterSimpleStructClass) Alloc() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterSimpleStructClass) New() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterSimpleStruct) Init() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterSimpleStruct) Autorelease() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterSimpleStruct creates a new MTRTestClusterClusterSimpleStruct instance.
func NewMTRTestClusterClusterSimpleStruct() MTRTestClusterClusterSimpleStruct {
	return getMTRTestClusterClusterSimpleStructClass().New()
}




