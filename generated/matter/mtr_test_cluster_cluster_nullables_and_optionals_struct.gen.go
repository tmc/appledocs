// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterNullablesAndOptionalsStruct] class.
var (
	MTRTestClusterClusterNullablesAndOptionalsStructClass     _MTRTestClusterClusterNullablesAndOptionalsStructClass
	MTRTestClusterClusterNullablesAndOptionalsStructClassOnce sync.Once
)

func getMTRTestClusterClusterNullablesAndOptionalsStructClass() _MTRTestClusterClusterNullablesAndOptionalsStructClass {
	MTRTestClusterClusterNullablesAndOptionalsStructClassOnce.Do(func() {
		MTRTestClusterClusterNullablesAndOptionalsStructClass = _MTRTestClusterClusterNullablesAndOptionalsStructClass{objc.GetClass("MTRTestClusterClusterNullablesAndOptionalsStruct")}
	})
	return MTRTestClusterClusterNullablesAndOptionalsStructClass
}

type _MTRTestClusterClusterNullablesAndOptionalsStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterNullablesAndOptionalsStruct] class.
type IMTRTestClusterClusterNullablesAndOptionalsStruct interface {
	IMTRUnitTestingClusterNullablesAndOptionalsStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterNullablesAndOptionalsStruct
type MTRTestClusterClusterNullablesAndOptionalsStruct struct {
	MTRUnitTestingClusterNullablesAndOptionalsStruct
}

// MTRTestClusterClusterNullablesAndOptionalsStructFrom constructs a [MTRTestClusterClusterNullablesAndOptionalsStruct] from an unsafe.Pointer.
func MTRTestClusterClusterNullablesAndOptionalsStructFrom(ptr unsafe.Pointer) MTRTestClusterClusterNullablesAndOptionalsStruct {
	return MTRTestClusterClusterNullablesAndOptionalsStruct{
		MTRUnitTestingClusterNullablesAndOptionalsStruct: MTRUnitTestingClusterNullablesAndOptionalsStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterNullablesAndOptionalsStructClass) Alloc() MTRTestClusterClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRTestClusterClusterNullablesAndOptionalsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterNullablesAndOptionalsStructClass) New() MTRTestClusterClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRTestClusterClusterNullablesAndOptionalsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) Init() MTRTestClusterClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRTestClusterClusterNullablesAndOptionalsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) Autorelease() MTRTestClusterClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRTestClusterClusterNullablesAndOptionalsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterNullablesAndOptionalsStruct creates a new MTRTestClusterClusterNullablesAndOptionalsStruct instance.
func NewMTRTestClusterClusterNullablesAndOptionalsStruct() MTRTestClusterClusterNullablesAndOptionalsStruct {
	return getMTRTestClusterClusterNullablesAndOptionalsStructClass().New()
}




